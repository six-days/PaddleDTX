#!/bin/bash

# Copyright (c) 2021 PaddlePaddle Authors. All Rights Reserved.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Script to test PaddleDTX service, start train or predict task with docker-compose.
# Usage: ./paddledtx_test.sh {upload_sample_files | start_vl_linear_train | start_vl_linear_predict | start_vl_logistic_train | start_vl_logistic_predict | tasklist | gettaskbyid}

# Requester private key 
# 计算需求方私钥
REQUESTERKEY="40816c779f624a8fbc4e37be1ef8bbddc6c07b5f91e704953a599f9080458f60"
# Requester public key 
# 计算需求方公钥
REQUESTER_PUBLICKEY="6cb69efc0439032b0d0f52bae1c9aada3f8fb46a5f24fa99065910055b77a1174d4afbac3c0529c8927587bb0e2ad90a85eaa600cfddd6b99f1212112135ef2b"
# Executor1 private key, same as network_up.sh's EXECUTOR1_PRIVATEKEY
# 任务执行节点1私钥，同network_up.sh的EXECUTOR1_PRIVATEKEY
EXECUTOR1_PRIVATEKEY="14a54c188d0071bc1b161a50fe7eacb74dcd016993bb7ad0d5449f72a8780e21"
EXECUTOR1_PUBLICKEY="4637ef79f14b036ced59b76408b0d88453ac9e5baa523a86890aa547eac3e3a0f4a3c005178f021c1b060d916f42082c18e1d57505cdaaeef106729e6442f4e5"
# Executor2 private key, same as network_up.sh's EXECUTOR2_PRIVATEKEY
# 任务执行节点2私钥，同network_up.sh的EXECUTOR2_PRIVATEKEY
EXECUTOR2_PRIVATEKEY="858843291fe4ed4bd2afc1120efd7315f3cae2d3f79e582f7df843ac6eb0543b"
EXECUTOR2_PUBLICKEY="e4530d81ccddc478978070e8f9fcc9f101dfc3b5c3ca1519c522c5e9698f394a35aab9145f242765185689a64b7338e9929c6a32e09050ff15645bb121ce1754"
CONFIG="./conf/config.toml"
# The namespace of the sample file store, same as network_up.sh's NAMESPACE
# 样本文件存储的命名空间，同network_up.sh的NAMESPACE
NAMESPACE=paddlempc
# Sample file description, used for training or prediction task
# 样本文件描述说明，用于任务训练或预测
LINEAR_TRAIN_SAMPLE_FILE_DES="vertical linear regression training sample file of Boston house price"
LINEAR_PREDICT_SAMPLE_FILE_DES="vertical linear regression prediction sample file of Boston house price"
LOGISTIC_TRAIN_SAMPLE_FILE_DES="vertical logistic regression training sample file of Iris plants"
LOGISTIC_PREDICT_SAMPLE_FILE_DES="vertical logistic regression prediction sample file of Iris plants"
# Expiration time of file storage
# 文件存储的到期时间
ARCH=$(uname -s | grep Darwin)
if [ "$ARCH" == "Darwin" ]; then
  FILE_EXPIRE_TIME=`date -v +6m +"%Y-%m-%d %H:%M:%S"`
else
  FILE_EXPIRE_TIME=`date -d "+6 month" +"%Y-%m-%d %H:%M:%S"`
fi

# Parameters required for training or prediction task
# 训练及预测任务所需的参数
PSI="id,id"
VLLINALGO="linear-vl"
VLLOGALGO="logistic-vl"
VLLINLABEL="MEDV"
VLLOGLABEL="Label"
VLLOGLABELName="Iris-setosa"
VLLINTASKTRAINNAME="boston_housing_train"
VLLINTASKPREDICTNAME="boston_housing_predict"
VLLOGTASKTRAINNAME="iris_plants_train"
VLLOGTASKPREDICTNAME="iris_plants_predict"
TASKNUM=1
#alpha=0.1
AMPLITUDE=0.0001
#batch=4

function uploadSampleFiles() {
  # 1. Create a namespace for the sample file store
  # 1. 创建文件存储的命名空间
  createNamespace
  
  # 2. Upload linear training file
  # 2. 上传线性训练文件
  sleep 1
  uploadLinearTrainSampleFile

  # 3. Upload linear prediction file
  # 3. 上传线性预测文件
  sleep 1
  uploadLinearPredictSampleFile

  # 4. Upload logic training files
  # 4. 上传逻辑训练文件
  sleep 1
  uploadLogisticTrainSampleFile

  # 5. Upload logic prediction file
  # 5. 上传逻辑预测文件
  sleep 1
  uploadLogisticPredictSampleFile
}

function createNamespace() {
  # 1. Create namespace for dataOwner1
  # 1. 数据持有节点1创建命名空间
  data1AddNsResult=`docker exec -it dataowner1.node.com sh -c "
    ./xdb-cli files addns  --host http://dataowner1.node.com:80 -k $EXECUTOR1_PRIVATEKEY -n $NAMESPACE -r 2"`
  echo "======> DataOwner1 create files storage namespace result: $data1AddNsResult"
  isData1AddNsOk=$(echo $data1AddNsResult | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1')
  if [ "$isData1AddNsOk" != "OK" ]; then
    exit 1
  fi
  # 2. Create namespace for dataOwner2
  # 2. 数据持有节点2创建命名空间
  data2AddNsResult=`docker exec -it dataowner2.node.com sh -c "
    ./xdb-cli files addns --host http://dataowner2.node.com:80 -k $EXECUTOR2_PRIVATEKEY  -n $NAMESPACE  -r 2 
  "`
  echo "======> DataOwner2 create files storage namespace result: $data2AddNsResult"
  isData2AddNsOk=$(echo $data2AddNsResult | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1')
  if [ "$isData2AddNsOk" != "OK" ]; then
    exit 1
  fi
}

# uploadLinearTrainSampleFile dataOwner1 and dataOwner2 upload vertical linear train sample files
# 数据持有节点1和数据持有节点2上传纵向线性训练样本文件
function uploadLinearTrainSampleFile() {
  sampleFileAName=train_dataA.csv
  sampleFileBName=train_dataB.csv
  fileAName="linear_"$sampleFileAName
  fileBName="linear_"$sampleFileBName

  # DataOwner1 uploads linear train sample files
  # 数据持有节点1上传纵向线性训练样本文件
  data1Samples=`docker exec -it dataowner1.node.com sh -c "
    ./xdb-cli files upload --host http://dataowner1.node.com:80  -e '$FILE_EXPIRE_TIME' -n $NAMESPACE -m $fileAName -k $EXECUTOR1_PRIVATEKEY \
    --ext '{\"FileType\":\"csv\",\"Features\":\"id,CRIM,ZN,INDUS,CHAS,NOX,RM\", \"TotalRows\":456}' -i /home/mpc-data/linear_boston_housing/$sampleFileAName \
    -d '$LINEAR_TRAIN_SAMPLE_FILE_DES'
  "`
  echo "======> DataOwner1 upload vertical_linear_train sample file: $data1Samples"

  data1FileUploadRes=$(echo $data1Samples | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1')
  data1FileId=${data1FileUploadRes##*: }

  sleep 3
  # DataOwner2 uploads linear train sample files
  # 数据持有节点2上传纵向线性训练样本文件
  data2Samples=`docker exec -it dataowner2.node.com sh -c "
    ./xdb-cli files upload --host http://dataowner2.node.com:80  -e '$FILE_EXPIRE_TIME' -n $NAMESPACE -m $fileBName -k $EXECUTOR2_PRIVATEKEY \
    --ext '{\"FileType\":\"csv\",\"Features\":\"id,AGE,DIS,RAD,TAX,PTRATIO,B,LSTAT,MEDV\",\"TotalRows\":456}' -i /home/mpc-data/linear_boston_housing/$sampleFileBName \
    -d '$LINEAR_TRAIN_SAMPLE_FILE_DES'
  "`
  echo "======> DataOwner2 upload vertical_linear_train sample file: $data2Samples"

  data2FileUploadRes=$(echo $data2Samples | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1')
  data2FileId=${data2FileUploadRes##*: }

  files="$data1FileId,$data2FileId"

  printf "\033[0;32m%s\033[0m\n" "======> Vertical linear train sample files: $files"
}

# uploadLinearPredictSampleFile dataOwner1 and dataOwner2 upload vertical linear prediction sample files
# 数据持有节点1和数据持有节点2上传纵向线性预测样本文件
function uploadLinearPredictSampleFile() {
  sampleFileAName=predict_dataA.csv
  sampleFileBName=predict_dataB.csv
  fileAName="linear_"$sampleFileAName
  fileBName="linear_"$sampleFileBName

  # DataOwner1 uploads linear prediction sample files
  # 数据持有节点1上传纵向线性预测样本文件
  data1Samples=`docker exec -it dataowner1.node.com sh -c "
    ./xdb-cli files upload --host http://dataowner1.node.com:80  -e '$FILE_EXPIRE_TIME' -n $NAMESPACE -m $fileAName -k $EXECUTOR1_PRIVATEKEY \
    --ext '{\"FileType\":\"csv\",\"Features\":\"id,CRIM,ZN,INDUS,CHAS,NOX,RM\", \"TotalRows\":50}' -i /home/mpc-data/linear_boston_housing/$sampleFileAName \
    -d '$LINEAR_PREDICT_SAMPLE_FILE_DES'
  "`
  echo "======> DataOwner1 upload vertical_linear_prediction sample file: $data1Samples"

  data1FileUploadRes=$(echo $data1Samples | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1')
  data1FileId=${data1FileUploadRes##*: }

  # DataOwner2 uploads linear prediction sample files
  # 数据持有节点2上传纵向线性预测样本文件
  data2Samples=`docker exec -it dataowner2.node.com sh -c "
    ./xdb-cli files upload --host http://dataowner2.node.com:80  -e '$FILE_EXPIRE_TIME' -n $NAMESPACE -m $fileBName -k $EXECUTOR2_PRIVATEKEY \
    --ext '{\"FileType\":\"csv\",\"Features\":\"id,AGE,DIS,RAD,TAX,PTRATIO,B,LSTAT\",\"TotalRows\":50}' -i /home/mpc-data/linear_boston_housing/$sampleFileBName \
    -d '$LINEAR_PREDICT_SAMPLE_FILE_DES'
  "`
  echo "======> DataOwner2 upload vertical_linear_prediction sample file: $data2Samples"

  data2FileUploadRes=$(echo $data2Samples | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1')
  data2FileId=${data2FileUploadRes##*: }

  files="$data1FileId,$data2FileId"
  printf "\033[0;32m%s\033[0m\n" "======> Vertical linear prediction sample files: $files"
}

# uploadLogisticTrainSampleFile dataOwner1 and dataOwner2 upload vertical logistic train sample files
# 数据持有节点1和数据持有节点2上传纵向逻辑训练样本文件
function uploadLogisticTrainSampleFile() {
  sampleFileAName=train_dataA.csv
  sampleFileBName=train_dataB.csv
  fileAName="logistic_"$sampleFileAName
  fileBName="logistic_"$sampleFileBName

  # DataOwner1 uploads logistic train sample files
  # 数据持有节点1上传纵向逻辑训练样本文件
  data1Samples=`docker exec -it dataowner1.node.com sh -c "
    ./xdb-cli files upload --host http://dataowner1.node.com:80  -e '$FILE_EXPIRE_TIME' -n $NAMESPACE -m $fileAName -k $EXECUTOR1_PRIVATEKEY \
    --ext '{\"FileType\":\"csv\",\"Features\":\"id,Sepal Length,Sepal Width\", \"TotalRows\":135}' -i /home/mpc-data/logic_iris_plants/$sampleFileAName \
    -d $LOGISTIC_TRAIN_SAMPLE_FILE_DES
  "`
  echo "======> DataOwner1 upload vertical_logistic_train sample file: $data1Samples"

  data1FileUploadRes=$(echo $data1Samples | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1')
  data1FileId=${data1FileUploadRes##*: }

  # DataOwner2 uploads logistic train sample files
  # 数据持有节点2上传纵向逻辑训练样本文件
  data2Samples=`docker exec -it dataowner2.node.com sh -c "
    ./xdb-cli files upload --host http://dataowner2.node.com:80  -e '$FILE_EXPIRE_TIME' -n $NAMESPACE -m $fileBName -k $EXECUTOR2_PRIVATEKEY \
    --ext '{\"FileType\":\"csv\",\"Features\":\"id,Petal Length,Petal Width,Label\", \"TotalRows\":135}' -i /home/mpc-data/logic_iris_plants/$sampleFileBName \
    -d $LOGISTIC_TRAIN_SAMPLE_FILE_DES
  "`
  echo "======> DataOwner2 upload vertical_logistic_train sample file: $data2Samples"

  data2FileUploadRes=$(echo $data2Samples | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1')
  data2FileId=${data2FileUploadRes##*: }

  files="$data1FileId,$data2FileId"

  printf "\033[0;32m%s\033[0m\n" "======> Vertical logistic train sample files: $files"
}

# uploadLogisticPredictSampleFile dataOwner1 and dataOwner2 upload vertical logistic prediction sample files
# 数据持有节点1和数据持有节点2上传纵向逻辑预测样本文件
function uploadLogisticPredictSampleFile() {
  sampleFileAName=predict_dataA.csv
  sampleFileBName=predict_dataB.csv
  fileAName="logistic_"$sampleFileAName
  fileBName="logistic_"$sampleFileBName

  # DataOwner1 uploads logistic prediction sample files
  # 数据持有节点1上传纵向逻辑预测样本文件
  data1Samples=`docker exec -it dataowner1.node.com sh -c "
    ./xdb-cli files upload --host http://dataowner1.node.com:80  -e '$FILE_EXPIRE_TIME' -n $NAMESPACE -m $fileAName -k $EXECUTOR1_PRIVATEKEY \
    --ext '{\"FileType\":\"csv\",\"Features\":\"id,Petal Length,Petal Width\", \"TotalRows\":15}' -i /home/mpc-data/logic_iris_plants/$sampleFileAName \
    -d $LOGISTIC_PREDICT_SAMPLE_FILE_DES
  "`
  echo "======> DataOwner1 upload vertical_logistic_prediction sample file: $data1Samples"

  data1FileUploadRes=$(echo $data1Samples | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1')
  data1FileId=${data1FileUploadRes##*: }

  # DataOwner2 uploads logistic prediction sample files
  # 数据持有节点2上传纵向逻辑预测样本文件
  data2Samples=`docker exec -it dataowner2.node.com sh -c "
    ./xdb-cli files upload --host http://dataowner2.node.com:80  -e '$FILE_EXPIRE_TIME' -n $NAMESPACE -m $fileBName -k $EXECUTOR2_PRIVATEKEY \
    --ext '{\"FileType\":\"csv\",\"Features\":\"id,Petal Length,Petal Width\", \"TotalRows\":15}' -i /home/mpc-data/logic_iris_plants/$sampleFileBName \
    -d $LOGISTIC_PREDICT_SAMPLE_FILE_DES
  "`
  echo "======> DataOwner2 upload vertical_logistic_prediction sample file: $data2Samples"

  data2FileUploadRes=$(echo $data2Samples | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1')
  data2FileId=${data2FileUploadRes##*: }

  files="$data1FileId,$data2FileId"

  printf "\033[0;32m%s\033[0m\n" "======> Vertical logistic prediction sample files: $files"
}

# taskConfirmAndStart used Executor1 and Executor2 confirm task, then requester start ready task
# 任务执行节点分别确认任务后，计算需求方启动任务
function taskConfirmAndStart() {
  sleep 4
  # Executor1 confirms the task published by the requester
  # 任务执行节点1确认任务
  executor1ConfirmResult=`docker exec -it executor1.node.com sh -c "
    ./executor-cli task --host executor1.node.com:80 confirm -k $EXECUTOR1_PRIVATEKEY -i $1"`
  echo "======> DataOwner1 authorizes Requester to use its data to train or predict and Executor1 confirms the task: $executor1ConfirmResult"
  sleep 4

  # Executor2 confirms the task published by the requester
  # 任务执行节点2确认任务
  executor2ConfirmResult=`docker exec -it executor2.node.com sh -c "
    ./executor-cli task --host executor2.node.com:80 confirm  -k $EXECUTOR2_PRIVATEKEY -i $1
    "`
  echo "======> DataOwner2 authorizes Requester to use its data to train or predict and Executor2 confirms the task: $executor1ConfirmResult"
  sleep 4

  # Requester starts the task when train or prediction task is confirmed
  # 计算方需求方启动任务
  requesterStartResult=`docker exec -it executor1.node.com sh -c "
  ./requester-cli task start -k $REQUESTERKEY -c ./conf/config.toml -i $1
  "`
  echo "======> Requester started task: $executor1ConfirmResult"
}

function linearVlTrain() {
  # List of sample files involved in linear train
  # 纵向线性训练任务所需的样本文件
  paramCheck "$VLLIN_TRAIN_FILES" "Sample files of linear train cannot be empty"
  for ((i=1; i<=$TASKNUM; i++))
  do
  # Requester published linear training task
  # 计算需求方发布纵向线性训练任务
  result=`docker exec -it executor1.node.com sh -c "./requester-cli task publish -p $PSI -a $VLLINALGO -f $VLLIN_TRAIN_FILES \
  -l $VLLINLABEL -k $REQUESTERKEY -t train -n $VLLINTASKTRAINNAME -c $CONFIG --amplitude $AMPLITUDE" | awk 'BEGIN{RS="\r";ORS="";}{print $0}'| awk '$1=$1'`
  checkOperateResult "$result"
  echo "======> Requester published linear train task: $result "
  taskid=${result##*: }

  taskConfirmAndStart $taskid
  done
}

function linearVlPredict() {
  # List of sample files involved in linear prediction
  # 纵向线性预测所需的预测样本文件
  paramCheck "$VLLIN_PREDICT_FILES" "Sample files of linear prediction cannot be empty"
  # Training task model ID required for linear prediction
  # 纵向线性预测所需的模型ID
  paramCheck "$LINEAR_MODEL_TASKID" "Training task model ID cannot be empty"
  # Requester published linear prediction task
  # 计算需求方发布纵向线性预测任务
  result=`docker exec -it executor1.node.com sh -c " 
    ./requester-cli task publish -p $PSI -a $VLLINALGO -f $VLLIN_PREDICT_FILES -k $REQUESTERKEY -t predict -n $VLLINTASKPREDICTNAME -c $CONFIG -i $LINEAR_MODEL_TASKID
  " | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1'`
  checkOperateResult "$result"
  echo "======> Requester published linear prediction task: $result "
  taskid=${result##*: }

  taskConfirmAndStart $taskid

  sleep 30
  # Get linear prediction results
  # 获取线性预测任务的预测结果
  linearVlPredictRes=`docker exec -it executor1.node.com sh -c "
  ./requester-cli task result -k $REQUESTERKEY -o ./linear_output.csv \
  --conf ./conf/config.toml  -i $taskid" | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1'`
  echo "======> Requester get linear prediction task result: $linearVlPredictRes "
  # Copy linear prediction results to the current directory
  # 将线性预测结果拷贝到当前目录
  docker cp executor1.node.com:/home/linear_output.csv ./
  echo "======> LinearVlPrediction file path: ./linear_output.csv"

  # Calculate root mean square error
  # 计算均方根误差
  housePricePredictFile=`cat linear_output.csv | awk 'NR>1'`
  housePriceTargetFile=`docker exec -it dataowner1.node.com sh -c "
  cat /home/mpc-data/linear_boston_housing/targetValues.csv | awk 'NR>1'
  " | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1'`
  calculateRMSE "$housePricePredictFile" "$housePriceTargetFile" "HousePrice" "Root mean square error of Boston house price prediction is"
}

function logisticVlTrain() {
  # List of sample files involved in logistic train
  # 纵向逻辑训练任务所需的样本文件
  paramCheck "$VLLOG_TRAIN_FILES" "Sample files of logistic train cannot be empty"
  for ((i=1; i<=$TASKNUM; i++))
  do
  # Requester published logistic training task
  # 计算需求方发布纵向逻辑训练任务
  result=`docker exec -it executor1.node.com sh -c "./requester-cli task publish -p $PSI -a $VLLOGALGO -f $VLLOG_TRAIN_FILES \
  -l $VLLOGLABEL -k $REQUESTERKEY -t train -n $VLLOGTASKTRAINNAME -c $CONFIG --labelName $VLLOGLABELName
  " | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1'`

  checkOperateResult "$result"
  echo "======> Requester published logistic train task: $result "
  taskid=${result##*: }

  taskConfirmAndStart $taskid
  done
}

function logisticVlPredict() {
  # List of sample files involved in logistic prediction
  # 纵向逻辑预测所需的预测样本文件
  paramCheck "$VLLOG_PREDICT_FILES" "Sample files of logistic prediction cannot be empty"
  # Training task model ID required for logistic prediction
  # 纵向逻辑预测所需的模型ID
  paramCheck "$LOGISTIC_MODEL_TASKID" "Training task model ID cannot be empty"
  # Requester published logistic training task
  # 计算需求方发布纵向逻辑预测任务
  result=`docker exec -it executor1.node.com sh -c " 
    ./requester-cli task publish -p $PSI -a $VLLOGALGO -f $VLLOG_PREDICT_FILES -k $REQUESTERKEY -t predict -n $VLLOGTASKPREDICTNAME -c $CONFIG -i $LOGISTIC_MODEL_TASKID
    " | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1'`

  checkOperateResult "$result"
  echo "======> Requester published logistic prediction task: $result "
  taskid=${result##*: }

  taskConfirmAndStart $taskid

  sleep 10
  # Get logistic prediction results
  # 获取逻辑预测任务的预测结果
  logisticVlPredictRes=`docker exec -it executor1.node.com sh -c "
  ./requester-cli task result -k $REQUESTERKEY -o ./logistic_output.csv \
  --conf ./conf/config.toml  -i $taskid
  " | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1'`
  echo "======> Requester get logistic prediction task result: $logisticVlPredictRes "
  # Copy logistic prediction results to the current directory
  # 将逻辑预测结果拷贝到当前目录
  docker cp executor1.node.com:/home/logistic_output.csv ./
  echo "======> LogisticVlPrediction file path: ./logistic_output.csv"

  # Calculate root mean square error
  # 计算均方根误差
  irisPredictFile=`cat logistic_output.csv | awk 'NR>1'`
  irisTargetFile=`docker exec -it dataowner1.node.com sh -c "
  cat /home/mpc-data/logic_iris_plants/targetValues.csv | awk 'NR>1'
  " | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1'`
  calculateLogisticPredictAccuracy "$irisPredictFile" "$irisTargetFile" "Accuracy of Iris plants prediction is"
}

function taskList() {
  docker exec -it executor1.node.com sh -c "
  ./requester-cli task list --conf ./conf/config.toml  -p $REQUESTER_PUBLICKEY"
}

function getTaskById() {
  paramCheck "$TASKID" "TaskId cannot be empty"
  docker exec -it executor1.node.com sh -c " 
  ./executor-cli task getbyid --host 127.0.0.1:80 -i $TASKID"
}

function checkOperateResult() {
  errMessage=`echo "$1" | grep -i "Error\|Fail\|Failed\|err"`
  if [ "$errMessage" ]; then
    printf "\033[0;31m%s\033[0m\n" "$1"
    exit 1
  fi
}

function paramCheck() {
  if [ "$1" = "" ];then
    printf "\033[0;31m======> ERROR !!!! %s\033[0m\n" "$2"
    exit 1
  fi
}

# Calculate Root Mean Squared Error
# 均方根误差, 用于衡量模型的误差, 真实值-预测值，然后平方之后求和，再计算平均值, 最后开平方
function calculateRMSE() {
  sumOfSquares=0
  logisticPredictAccuracy=0
  for line in $2
  do
    IFS=',' read -r -a targetArray <<<"$line"
    targetKey=${targetArray[1]}
    targetValue=${targetArray[2]}
    if [ "$3" = "IrisPlants" ];then
      if [ "$targetValue" = "Iris-setosa" ];then
        targetValue=1
      else
        targetValue=0
      fi
    fi
    # Read the predicted value of the same sample ID
    # 读取相同样本ID的预测值
    predictValueLine=`echo "$1" | grep "^$targetKey,"`
    IFS=',' read -r -a predictArray <<<"$predictValueLine"
    predictValue=${predictArray[1]}
   
    predictError=`echo "scale=15;$targetValue-$predictValue"|bc|awk '{printf "%.15f", $0}'`
    sumOfSquare=`echo "scale=15;$predictError*$predictError"|bc|awk '{printf "%.15f", $0}'`
    sumOfSquares=`echo "scale=15;$sumOfSquares+$sumOfSquare"|bc|awk '{printf "%.15f", $0}'`
  done
  num=`echo "$2"| wc -l | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1'`
  avg=`echo "scale=15;$sumOfSquares/$num"|bc|awk '{printf "%.15f", $0}'`
  rmse=`echo "scale=15;sqrt($avg)"|bc|awk '{printf "%.15f", $0}'`
  echo "======> $4: $rmse"
}

# Calculate the accuracy of iris prediction
# 计算鸢尾花预测的准确率
function calculateLogisticPredictAccuracy() {
  predictCorrectNum=0
  for line in $2
  do
    IFS=',' read -r -a targetArray <<<"$line"
    targetKey=${targetArray[1]}
    targetValue=${targetArray[2]}

    # Read the predicted value of the same sample ID
    # 读取相同样本ID的预测值
    predictValueLine=`echo "$1" | grep "^$targetKey,"`
    IFS=',' read -r -a predictArray <<<"$predictValueLine"
    predictValue=${predictArray[1]}
    if [ `echo "$predictValue >= 0.5" | bc` -eq 1 ]&&[ "$targetValue" = "Iris-setosa" ];then
      predictCorrectNum=$((predictCorrectNum+1))
    elif [ `echo "$predictValue < 0.5" | bc` -eq 1 ]&&[ "$targetValue" != "Iris-setosa" ];then
      predictCorrectNum=$((predictCorrectNum+1))
    else
      predictCorrectNum=$((predictCorrectNum-1))
    fi
  done
  num=`echo "$2"| wc -l | awk 'BEGIN{RS="\r";ORS="";}{print $0}' | awk '$1=$1'`
  accuracy=`echo "scale=2;$predictCorrectNum/$num"|bc|awk '{printf "%.2f", $0}'`
  echo "======> $3: $accuracy"
}

# Print the usage message
function printHelp() {
  echo "Usage: "
  echo "  ./paddledtx_test.sh <mode> [-f <sample files>] [-m <model task id>] [-i <task id>]"
  echo "    <mode> - one of 'upload_sample_files', 'start_vl_linear_train', 'start_vl_linear_predict', 'start_vl_logistic_train'"
  echo "         'start_vl_logistic_predict', 'tasklist', 'gettaskbyid'"
  echo "      - 'upload_sample_files' - save linear and logistic sample files into XuperDB"
  echo "      - 'start_vl_linear_train' - start vertical linear training task"
  echo "      - 'start_vl_linear_predict' - start vertical linear prediction task"
  echo "      - 'start_vl_logistic_train' - start vertical logistic training task"
  echo "      - 'start_vl_logistic_predict' - start vertical logistic prediction task"
  echo "      - 'tasklist' - list task in PaddleDTX"
  echo "      - 'gettaskbyid' - get task by id from PaddleDTX"
  echo "    -f <sample files> - linear or logistic sample files"
  echo "    -m <model task id> - finished train task ID from which obtain the model, required for predict task"
  echo "    -i <task id> - training or prediction task id"
  echo
  echo "  ./paddledtx_test.sh -h (print this message), e.g.:"
  echo
  echo "  ./paddledtx_test.sh upload_sample_files"
  echo "  ./paddledtx_test.sh start_vl_linear_train -f 1ffc4504-6a62-45be-a7e3-191c708b901f,f8439128-bebb-47c2-a04d-1121dbc087a4"
  echo "  ./paddledtx_test.sh start_vl_linear_predict -f cb40b8ad-db08-447f-a9d9-628b69d01660,2a8a45ab-3c5d-482e-b945-bc45b7e28bf9 -m 9b3ff4be-bfcd-4520-a23b-4aa6ea4d59f1"
  echo "  ./paddledtx_test.sh start_vl_logistic_train -f b31f53a5-0f8b-4f57-a7ea-956f1c7f7991,f3dddade-1f52-4b9e-9253-835e9fc81901"
  echo "  ./paddledtx_test.sh start_vl_logistic_predict -f 1e97d684-722f-4798-aaf0-dffe955a94ba,b51a927c-f73e-4b8f-a81c-491b9e938b4d -m d8c8865c-a837-41fd-802b-8bd754b648eb"
  echo "  ./paddledtx_test.sh gettaskbyid -i 9b3ff4be-bfcd-4520-a23b-4aa6ea4d59f1"
  echo "  ./paddledtx_test.sh tasklist"
  echo
}

VLLIN_TRAIN_FILES=""
VLLIN_PREDICT_FILES=""
LINEAR_MODEL_TASKID=""
VLLOG_TRAIN_FILES=""
VLLOG_PREDICT_FILES=""
LOGISTIC_MODEL_TASKID=""
TASKID=""

action=$1
shift
while getopts "h?f:m:i:" opt; do
  case "$opt" in
  h | \?)
    printHelp
    exit 0
    ;;
  f)
    if [ "$action" == "start_vl_linear_train" ]; then
      VLLIN_TRAIN_FILES=$OPTARG
    elif [ "$action" == "start_vl_linear_predict" ]; then
      VLLIN_PREDICT_FILES=$OPTARG
    elif [ "$action" == "start_vl_logistic_train" ]; then
      VLLOG_TRAIN_FILES=$OPTARG
    elif [ "$action" == "start_vl_logistic_predict" ]; then
      VLLOG_PREDICT_FILES=$OPTARG
    else
      printHelp
      exit 0
    fi
    ;;
  m)
    if [ "$action" == "start_vl_linear_predict" ]; then
      LINEAR_MODEL_TASKID=$OPTARG
    elif [ "$action" == "start_vl_logistic_predict" ]; then
      LOGISTIC_MODEL_TASKID=$OPTARG
    else
      printHelp
      exit 0
    fi
    ;;
  i)
    TASKID=$OPTARG
    ;;
  esac
done

case $action in
upload_sample_files)
  uploadSampleFiles $@
  ;;
start_vl_linear_train)
  linearVlTrain $@
  ;;
start_vl_linear_predict)
  linearVlPredict $@
  ;;
start_vl_logistic_train)
  logisticVlTrain $@
  ;;
start_vl_logistic_predict)
  logisticVlPredict $@
  ;;
tasklist)
  taskList $@
  ;;
gettaskbyid)
  getTaskById $@
  ;;
*)
  printHelp
  exit 1
  ;;
esac
