
pipeline {
  agent {
      kubernetes {
        label 'hello'
        yaml '''
apiVersion: v1
kind: Pod
metadata:
   name: clean-ci
spec:
   containers:
   - name: docker
     image: 'docker:stable-dind'
     command:
     - dockerd
     - --host=unix:///var/run/docker.sock
     - --host=tcp://0.0.0.0:8000
     - --insecure-registry=157.230.248.65:30002
     securityContext:
       privileged: true
     volumeMounts:
     - mountPath: /var/run
       name: cache-dir
   - name: clean-ci
     image: 'docker:stable'
     command: ["/bin/sh"]
     args: ["-c","while true; do sleep 86400; done"]
     volumeMounts:
     - mountPath: /var/run
       name: cache-dir

   - name: go-lint
     image: 'golangci/golangci-lint'
     command: ["/bin/sh"]
     args: ["-c","while true; do sleep 86400; done"]
     volumeMounts:
     - mountPath: /var/run
       name: cache-dir


   - name: trivy
     image: 'aquasec/trivy:0.21.1'
     command: ["/bin/sh"]
     args: ["-c","while true; do sleep 86400; done"]
     volumeMounts:
     - mountPath: /var/run
       name: cache-dir


   volumes:
   - name: cache-dir
     emptyDir: {}
        '''.stripIndent()
          }
          }

    stages {
        stage ('checkout scm') {
            steps {
                checkout(scm)
            }
       }


		stage("build & SonarQube analysis") {
            steps {
                script {
                scannerHome = tool 'sonarsc'
                }
                withSonarQubeEnv('SonarQube') {
                sh "${scannerHome}/bin/sonar-scanner"
                }
            }
        }

        stage("Quality Gate"){
			steps{
				timeout(time: 15, unit: 'MINUTES') {
					waitForQualityGate abortPipeline: true
				}
			}
		}


        stage('scan with trivy') {
            steps {
                container ('trivy')
                sh "trivy image -f json -o results.json nginx:1.18"
                recordIssues(tools: [trivy(pattern: 'results.json')])
            }
        }




 /*       stage('deploy to dev'){
 *         steps {
 *            container ('kubectl'){
 *                   sh '  kubectl apply -f ./yml/my.yaml'
 *            }
 *         }
 */      }


        //直接通过定义两个image scanner，自动修改指定镜像，完成测试环境与生产环境的操作。
        //git commit测试环境代码库，(考虑自动commit或者手动commit)，人工测试，通过就commit生产代码库。

/*           stage('push with tag'){
*             steps {
*                container ('docker'){
*                    sh '''
*                    echo "git push"
*                    '''
*                }
*             }
*           }

*           stage('deploy to prod'){
*             steps {
*                //是否部署到生产环境。
*                input(id: 'deploy-to-dev', message: 'deploy to dev?')
*                container ('docker'){
*

        sh '''
*                    echo "deploy"
*                    '''
*                }
*           }
*            //添加回滚功能
*/     }

