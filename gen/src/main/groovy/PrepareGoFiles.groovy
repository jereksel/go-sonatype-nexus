class PrepareGoFiles {

    public static void main(args) {

        def scripts = ["GetAll", "Remove", "CreateHostedMaven", "CreateProxyMaven"]

        def dest = new File("../scripts")

        dest.listFiles().toList().forEach { File file -> file.delete() }

        scripts.forEach { String scriptName ->

            def text = new File("src/main/groovy/${scriptName}.groovy").text

            text = text.split("// ###")[0]

            text = text.replace("// #", "")

            def file = """package scripts

var ${scriptName}Script = `${text}`


"""

            new File(dest, "${scriptName}.go").setText(file)

        }

//        println(dest)

    }

}
