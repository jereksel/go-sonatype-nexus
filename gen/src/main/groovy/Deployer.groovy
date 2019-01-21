import static groovyx.net.http.HttpBuilder.configure
import static groovy.json.JsonOutput.toJson

class Deployer {

    static public void deploy(String name) {

//        println(new File(".").listFiles())

        def text = new File("src/main/groovy/${name}.groovy").text

        text = text.split("// ###")[0]

        text = text.replace("// #", "")

        println(text)

        def id = UUID.randomUUID().toString()

        def http = configure {
            request.uri = "http://localhost:8081/service/rest/v1/script"
            request.headers = [Authorization: "Basic YWRtaW46YWRtaW4xMjM="]
        }

        http.post {
            request.contentType = "application/json"
            request.body = toJson(name: id, content: text, type: "groovy")
        }

        def result = http.post {
            request.contentType = "text/plain"
            request.uri = "http://localhost:8081/service/rest/v1/script/${id}/run"
            request.body = "{}"
        }

        

    }

}
