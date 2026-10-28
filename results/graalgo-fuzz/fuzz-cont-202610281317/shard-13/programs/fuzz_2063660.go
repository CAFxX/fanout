package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
func genFn0() uint8 {
    var v7 uint32 = 4294967295
    var v8 int32 = 1
    v8 -= v8
    _ = v7
    _ = v8
    return 102
}

func main() {
    var v1 int8 = 127
    v2 := []string{("hello" + ("" + ("\n\t" + "foo bar")))}
    var v3 uint64 = 1
    v4 := []int{32768, 32768}
    var v5 uint16 = 255
    v6 := make(map[string]bool)
    _ = len(v6)
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 bool = (14 <= v1)
        v3 = (v3 - 4294967295)
        if (!(0.0 < 0.0)) {
            if true {
                v11 := GenS1{F0: genFn0(), F1: 128}
                _ = v11
            } else {
                var v12 int8 = (^(v1))
                _ = v12
            }
        }
        _ = fi9
        _ = v10
    }
    v5 = 65535
    var v13 int64 = -1
    switch ("hello" + "a") {
    case "":
        var v14 int64 = 42
        _ = v14
    }
    var v15 float32 = 0.1
    if (!(!true)) {
        var v16 int32 = -2147483648
        v15 -= v15
        var v17 uint64 = (v3 % ((735) | 1))
        _ = v16
        _ = v17
    } else {
        v18 := GenS2{F0: 33}
        _ = v18
    }
    switch 128 {
    case 519:
        v19 := GenS3{F0: int64(v5)}
        _ = v19
    case 32768:
        if len(v2) > 0 {
            v2[0] = "hello"
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
