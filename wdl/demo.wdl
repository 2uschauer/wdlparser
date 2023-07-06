version 1.0

task run {
  input {
    String task_input
  }
  command {
    java -jar task.jar -id ${task_input} -out ${task_input}.out
  }
  output {
    File task_output = "${task_input}.out"
  }
}
workflow test {
  input {
    String workflow_input
  }

  call run {
    input: task_input=workflow_input
  }

  output {
    File workflow_output = run.task_output
  }
}