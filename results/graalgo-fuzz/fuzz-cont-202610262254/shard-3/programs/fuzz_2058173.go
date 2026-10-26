package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint16 }

func main() {
    v1 := []string{((("hello" + "") + ("a" + "foo bar")) + ("\n\t" + "a")), ""}
    v2 := []string{}
    var v3 string = "foo bar"
    v4 := []string{}
    if true {
        var v5 float64 = float64((1e10 / 1.5))
        v3 = (("hello" + ((v3 + (v3 + "\n\t")) + v3)) + "foo bar")
        var v6 int = 2147483647
        _ = v5
        _ = v6
    }
    switch ("hello" + ((v3 + ((v3 + "\n\t") + v3)) + (v3 + v3))) {
    case "foo bar":
        var v7 string = v3
        _ = v7
    case "hello":
        var v8 uint64 = 4294967295
        _ = v8
    }
    if len(v4) > 0 {
        v4[0] = ((v3 + ((((v3 + "a") + "") + (("" + v3) + (v3 + v3))) + (v3 + v3))) + v3)
    }
    v9 := GenS1{F0: false, F1: -579, F2: 207}
    if true {
        var v10 float32 = (((3.14159 + 1e10) / (1e10 / (3.14159 + 0.1))) / 1e10)
        v10 -= ((v10 / ((1e10 / (((0.1 + (1e10 + 1e10)) - ((v10 * float32((float64(v10) / 1e10))) - ((v10 + -2.25) - (0.0 - 1.5)))) + 1.5)) - ((v10 + v10) * v10))) - (v10 / (-2.25 + float32((1.5 * ((float64(v10) + (3.14159 + 0.1)) / 3.14159))))))
        v3 = v3
        _ = v10
    } else {
        fmt.Println(v3)
    }
    v11 := GenS1{F0: (-513 < 128), F1: -67, F2: 65535}
    for fi12 := 0; fi12 < 13; fi12++ {
        v3 = (v3 + v3)
        if true {
            var v13 uint = 2
            _ = v13
        }
        v2 = append(v2, v3)
        _ = fi12
    }
    var v14 float64 = 0.1
    fmt.Println(v3, v3)
    for fi15 := 0; fi15 < 5; fi15++ {
        if false {
            v3 = (((v3 + (v3 + v3)) + "foo bar") + v3)
        } else {
            _ = len(v4)
            fmt.Println(fi15, v14)
            v1 = append(v1, v3)
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi15
    }
    var v16 uint32 = 2147483647
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
