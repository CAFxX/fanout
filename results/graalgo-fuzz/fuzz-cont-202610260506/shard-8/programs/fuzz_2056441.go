package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint8; F1 uint8; F2 int8 }
func genFn0() int64 {
    if true {
        v4 := GenS1{F0: "", F1: 21}
        var v5 int = 32768
        fmt.Println(v5)
        _ = v4
        _ = v5
    } else {
        v6 := GenS2{F0: ""}
        var v7 string = ""
        if true {
            v8 := []int{783, 2147483647, -176, 4294967295}
            v9 := GenS3{F0: 1, F1: 0}
            var v10 int32 = 2147483647
            _ = v10
            _ = v8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    var v11 uint8 = 42
    v11 = (v11 - v11)
    var v12 uint16 = 339
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 uint8 = 127
    v2 := []uint32{2, 2147483647, 998, 2147483647}
    var v3 int32 = -1
    var v13 float64 = 1.5
    v14 := GenS1{F0: ("\n\t" + "hello"), F1: v1}
    fmt.Println(v13, v13, v13)
    v15 := GenS4{F0: 128, F1: (v1 / (((^(v1))) | 1)), F2: -1}
    v16 := GenS2{F0: "\n\t"}
    for fi17 := 0; fi17 < 5; fi17++ {
        if (!(false || true)) {
            var v18 int8 = 0
            v2 = append(v2, 65535)
            _ = v18
        }
        _ = fi17
    }
    fmt.Println(v1, v1)
    switch "\n\t" {
    case "":
        if (v1 < 255) {
            var v19 int8 = 127
            v19 = v19
            _ = v19
        } else {
            var v20 uint8 = 127
            for fi21 := 0; fi21 < 20; fi21++ {
                fmt.Println(v13, fi21)
                _ = fi21
            }
            v20 += (^(v20))
            _ = v20
        }
    default:
        v22 := GenS2{F0: "foo bar"}
        _ = v22
    }
    _ = len(v2)
    v23 := GenS3{F0: 560, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
}
