package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 string; F1 int32 }

func main() {
    var v1 int8 = -1
    v2 := make(map[int]string)
    v3 := []string{"a", "\n\t", ("hello" + ""), "foo bar"}
    var v4 uint32 = 65535
    v4 += 548
    v5 := GenS1{F0: 42, F1: true, F2: 1}
    v3 = append(v3, "hello")
    fmt.Println(v1, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v6 := GenS1{F0: 42, F1: false, F2: 255}
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint8 = 1
        var v9 uint32 = (^(v4))
        v4 += 2147483648
        _ = fi7
        _ = v8
        _ = v9
    }
    v4 = v4
    if (true && (true || false)) {
        v10 := GenS2{F0: -1, F1: 277}
        v4 = v4
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := GenS1{F0: 2, F1: (("" + "foo bar") < "foo bar"), F2: 235}
            v4 = 0
            _ = fi11
            _ = v12
        }
        _ = v10
    } else {
        v4 = (v4 + v4)
    }
    if false {
        v13 := GenS2{F0: 42, F1: 2147483647}
        fmt.Println(v4, v1)
        v14 := GenS3{F0: "hello", F1: 755}
        _ = v13
        _ = v14
    }
    var v15 int32 = -19
    switch 127 {
    case 464:
        v1 += (^(v1))
    case 777:
        var v16 uint8 = 226
        _ = v16
    }
    var v17 float32 = 1.5
    if ((^(v1)) < (v1 + v1)) {
        fmt.Println(v15)
    }
    v17 = v17
    switch ("a" + "") {
    case "foo bar":
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
