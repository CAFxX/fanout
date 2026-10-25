package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 int = 4294967295
    var v2 uint8 = 2
    var v3 string = "a"
    v4 := []uint32{503, 65535}
    var v5 string = v3
    v6 := GenS1{F0: (v1 | (v1 | v1)), F1: v5, F2: -32768}
    _ = len(v4)
    _ = len(v4)
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            fmt.Println(v1)
            v7 := GenS1{F0: (-(v1)), F1: (v5 + v3), F2: 42}
            v8 := GenS2{F0: 252}
            _ = v7
            _ = v8
        } else {
            var v9 uint = 42
            var v10 float64 = 0.0
            fmt.Println(v5, v1, v2)
            _ = v10
            _ = v9
        }
    }
    if (-13 == 1) {
        v4 = append(v4, 4294967295)
        var v11 uint32 = 689
        _ = v11
    } else {
        v4 = append(v4, 1)
        v1 = v1
        var v12 int = 32767
        _ = v12
    }
    switch "" {
    case "":
        fmt.Println(v2, v2, v2)
    default:
        switch "\n\t" {
        case "foo bar":
            v13 := GenS3{F0: 50}
            _ = v13
        }
    }
    switch (((v5 + (("" + ("a" + ((v5 + (("hello" + v3) + "foo bar")) + ("a" + (("foo bar" + "\n\t") + (v5 + "foo bar")))))) + v3)) + ("foo bar" + v5)) + v5) {
    case "\n\t":
        v14 := GenS1{F0: -25, F1: "a", F2: 32767}
        _ = v14
    case "a":
        v2 = v2
    default:
        var v15 int8 = 1
        _ = v15
    }
    v2 = 127
    var v16 uint = 0
    v2 = v2
    var v17 uint8 = 82
    if (v16 < v16) {
        v2 = (v17 & (v17 + (v17 & v17)))
    } else {
        v18 := GenS1{F0: 0, F1: "", F2: 0}
        v17 = v17
        _ = v18
    }
    v3 = v5
    var v19 uint16 = 1
    v4 = append(v4, uint32(v17))
    v20 := GenS3{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
