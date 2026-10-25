package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int }
func genFn0() int32 {
    var v7 uint64 = 42
    var v8 uint16 = 688
    var v9 bool = false
    var v10 uint16 = (v8 - v8)
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -293
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 42
    v3 := make(map[uint32]uint32)
    var v4 uint16 = 255
    v5 := []uint32{}
    var v6 int16 = v2
    var v11 int8 = 100
    v12 := GenS1{F0: -196, F1: genFn0()}
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS1{F0: fi13, F1: -1}
        if true {
            v15 := GenS1{F0: fi13, F1: 694}
            _ = v15
        } else {
            v6 = v2
            for fi16 := 0; fi16 < 14; fi16++ {
                v4 += (v4 >> ((18446744073709551615) & 63))
                var v17 float32 = 0.0
                v18 := GenS1{F0: (-(fi13)), F1: 0}
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        _ = fi13
        _ = v14
    }
    var v19 uint8 = 159
    fmt.Println(v2, v19, v11)
    v4 = (v4 * v4)
    switch 4294967295 {
    case 32767:
        switch ("\n\t" + "a") {
        case "a":
            var v20 uint8 = (^(v19))
            _ = v20
        case "foo bar":
            v4 += (v4 ^ v4)
        default:
            var v21 uint64 = 0
            _ = v21
        }
    case -1:
        v22 := GenS1{F0: 464, F1: genFn0()}
        _ = v22
    default:
        var v23 uint64 = 4294967295
        _ = v23
    }
    var v24 int32 = -545
    v25 := GenS2{F0: 4294967295}
    var v26 uint64 = 42
    _ = len(v5)
    var v27 string = ("foo bar" + ("a" + "foo bar"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
