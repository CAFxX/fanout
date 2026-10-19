package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }

func main() {
    var v1 int64 = 42
    var v2 uint32 = 1
    var v3 int64 = -1
    v4 := []int{32768}
    var v5 uint8 = 109
    v6 := []string{("" + "foo bar"), ("\n\t" + (("foo bar" + "a") + "a")), ((("a" + "foo bar") + "\n\t") + "")}
    if false {
        var v7 float64 = (-2.25 + 1e10)
        _ = v7
    }
    switch ((("foo bar" + "foo bar") + ((("hello" + (("hello" + "") + "\n\t")) + "foo bar") + "foo bar")) + "a") {
    case "hello":
        var v8 string = "a"
        _ = v8
    }
    var v9 float64 = 0.1
    if true {
        switch 128 {
        case 1:
            fmt.Println(v5, v9, v1)
        }
        var v10 bool = false
        v9 = v9
        _ = v10
    } else {
        var v11 int32 = -884
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v11
    }
    if true {
        var v12 string = ""
        v13 := GenS1{F0: 880, F1: 127}
        _ = v12
        _ = v13
    }
    v6 = append(v6, (("\n\t" + "\n\t") + "hello"))
    fmt.Println(v1, v1)
    v14 := GenS1{F0: 37, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
