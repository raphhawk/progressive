### Progressive.go

A minimal concurrent progress bar to get feedback on your running go processes.

#### Examples

##### Single Process

```go
    // initializing channel
    ch := make(chan int)

    // creating a valid progress
    var p ValidProgress = &ProgressChannel{ch, 0, "", ""}

    // creating a progress bar on different go routine
    go progressBar(25, p, colorMap["Green"], colorMap["Gray"])

    // mocking multiple steps and progress in a process
    p.passProgress(20, "Initializing Containers")
    p.passProgress(50, "Creating Kube Clusters")
    p.passProgress(75, "Deploying Application")
    p.passProgress(100, "Operation Complete")

    //closing channel
    p.closeProgress()
```

##### Multiple Processes

```go
    // initializing channel
    ch := make(chan int)

    // creating multiple valid progress using same channel
    var p ValidProgress = &ProgressChannel{ch, 0, "", "Process 1"}
    var pt ValidProgress = &ProgressChannel{ch, 0, "", "Process 2"}

    // creating multiple progress bars on different go routines
    go progressBar(25, p, colorMap["Green"], colorMap["Gray"])
    go progressBar(25, pt, colorMap["Purple"], colorMap["Gray"])

    // mocking multiple steps and progress in multiple processes
    p.passProgress(20, "Initializing Containers")
    p.passProgress(50, "Creating Kube Clusters")

    pt.passProgress(28, "Initializing Services")

    p.passProgress(75, "Deploying Application")
    pt.passProgress(34, "Monitorings Pod Health")
    p.passProgress(100, "Operation Complete")

    pt.passProgress(56, "Service #1 Deployed")
    pt.passProgress(82, "Service #2 Deployed")
    pt.passProgress(100, "Operation Complete")

    //closing channel
    pt.closeProgress()
```

#### Goal

Intented for any external application tasks whose progress have to be pre-calculated and needs to be tracked visually and concurrently.
