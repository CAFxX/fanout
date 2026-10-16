package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
func genFn0() bool {
    if (-9223372036854775808 == 9223372036854775807) {
        var v5 int32 = -1
        for fi6 := 0; fi6 < 3; fi6++ {
            fmt.Println(fi6)
            var v7 int16 = 1
            fmt.Println(v5)
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    var v8 uint = 65535
    switch -59 {
    case 128:
        var v9 uint16 = 1
        _ = v9
    case 1:
        var v10 string = ""
        _ = v10
    }
    var v11 float32 = 0.0
    var v12 uint32 = 2147483647
    _ = v11
    _ = v12
    _ = v8
    return (-95 < 42)
}

func main() {
    var v1 int64 = -426
    v2 := make(map[uint32]string)
    var v3 uint = 0
    v4 := make(map[uint32]int)
    switch "hello" {
    case "hello":
        v3 -= (^(v3))
    case "":
        v4[0] = 127
    case "\n\t":
        v13 := GenS1{F0: -1, F1: -773, F2: ""}
        _ = v13
    default:
        var v14 int8 = 0
        _ = v14
    }
    v3 -= 42
    for fi15 := 0; fi15 < 5; fi15++ {
        if false {
            _ = len(v4)
            v3 += v3
        } else {
            fmt.Println(fi15)
            var v16 int = 683
            fmt.Println(v3)
            _ = v16
        }
        fmt.Println(fi15, v3, v3)
        _ = fi15
    }
    var v17 uint64 = 744
    if (true || true) {
        fmt.Println(v3, v3, v1)
        v3 -= 2
        for fi18 := 0; fi18 < 7; fi18++ {
            switch fi18 {
            case 127:
                fmt.Println(v17)
            }
            var v19 uint64 = v17
            _ = fi18
            _ = v19
        }
    } else {
        var v20 string = "\n\t"
        _ = v20
    }
    if true {
        v21 := []int{}
        var v22 uint8 = 255
        _ = v21
        _ = v22
    } else {
        if false {
            v23 := GenS2{F0: v3}
            v17 = (v17 << (((v3 ^ (v3 ^ 0))) & 63))
            var v24 int16 = 42
            _ = v23
            _ = v24
        } else {
            if true {
                fmt.Println(v17, v3, v17)
            } else {
                v25 := v4[42]
                _ = v25
            }
            var v26 int32 = -2147483648
            _ = v26
        }
    }
    v1 += v1
    v27 := GenS3{F0: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
