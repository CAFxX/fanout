package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0() int64 {
    var v5 uint16 = 65535
    v6 := []int{-550}
    _ = v5
    _ = v6
    return 42
}
func genFn1() int {
    var v7 int = 2147483647
    if (42 == 88) {
        var v8 uint8 = 128
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        v7 += fi9
        v10 := []int{-325}
        v11 := GenS1{F0: true}
        _ = fi9
        _ = v10
        _ = v11
    }
    v7 -= v7
    _ = v7
    return v7
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    var v3 uint8 = 255
    v4 := make(map[int]bool)
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := GenS1{F0: ((!true) || (349 < 32767))}
        fmt.Println(fi12, v3)
        for fi14 := 0; fi14 < 10; fi14++ {
            switch (fi12 / (((fi14 | genFn1())) | 1)) {
            case 42:
                var v15 bool = true
                _ = v15
            default:
                v3 -= v3
            }
            v16 := GenS1{F0: false}
            _ = fi14
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    var v17 int64 = 1
    v17 = (v17 - (v17 + genFn0()))
    v17 = v17
    if (!true) {
        var v18 float32 = 1.5
        _ = v18
    }
    v3 = (v3 / ((v3) | 1))
    switch genFn1() {
    case 4294967295:
        v17 = (v17 << ((0) & 63))
    case 127:
        fmt.Println(v17)
    }
    fmt.Println(v3, v3)
    if true {
        switch -1 {
        case 42:
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 uint16 = uint16(v3)
                switch int(v20) {
                case -227:
                    v1["a"] = false
                default:
                    fmt.Println(v3)
                }
                _ = fi19
                _ = v20
            }
        case 328:
            switch 2147483647 {
            case 32767:
                if true {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        v17 = v17
                        v22 := GenS2{F0: 285, F1: -1, F2: 883}
                        _ = fi21
                        _ = v22
                    }
                }
            case -9223372036854775808:
                fmt.Println(v17, v17)
            default:
                fmt.Println(v3, v17)
            }
        default:
            switch "\n\t" {
            case "a":
                fmt.Println(v17, v17)
            case "hello":
                var v23 bool = (!false)
                _ = v23
            case "":
                fmt.Println(v3)
            default:
                for fi24 := 0; fi24 < 15; fi24++ {
                    var v25 uint16 = 439
                    v26 := GenS3{F0: -1, F1: 246}
                    v27 := GenS3{F0: 537, F1: -489}
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
            }
        }
    } else {
        v3 = v3
        var v28 uint32 = 42
        _ = v28
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        v3 = v3
        _ = fi29
    }
    v30 := GenS2{F0: 2147483648, F1: 127, F2: 65535}
    var v31 uint16 = 65535
    v17 = -202
    var v32 int = 2147483647
    var v33 uint8 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
