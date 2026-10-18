package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint32; F1 int; F2 string }
func genFn0(p0 bool) int64 {
    v5 := []int{-452, -356}
    p0 = p0
    fmt.Println(p0)
    v5 = append(v5, 128)
    _ = p0
    _ = v5
    return -449
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "\n\t"
    var v3 uint32 = 2147483647
    var v4 uint32 = v3
    v4 += v3
    var v6 uint = 0
    var v7 uint64 = 271
    if (!(true && (false && (1 <= 1)))) {
        var v8 uint16 = 690
        _ = v8
    }
    v1[(v4 - v4)] = (v2 + v2)
    delete(v1, v4)
    v9 := GenS1{F0: 9223372036854775807}
    if (false && false) {
        var v10 uint16 = uint16(v6)
        v2 = "\n\t"
        switch "hello" {
        case "foo bar":
            fmt.Println(v6, v3)
        case "a":
            if false {
                for fi11 := 0; fi11 < 20; fi11++ {
                    v12 := GenS1{F0: genFn0(true)}
                    fmt.Println(v3)
                    if (-762 < 0) {
                        v1[65535] = v2
                        var v13 uint = 42
                        _ = v13
                    } else {
                        v14 := GenS2{F0: -80, F1: 4294967295, F2: int8(fi11)}
                        v15 := []int{(fi11 << (((v6 + 1)) & 63)), fi11}
                        delete(v1, uint32(v10))
                        _ = v14
                        _ = v15
                    }
                    _ = fi11
                    _ = v12
                }
                _ = len(v1)
            } else {
                var v16 int8 = 1
                for fi17 := 0; fi17 < 2; fi17++ {
                    var v18 int64 = 9223372036854775807
                    var v19 uint64 = (v7 / (((v7 - v7)) | 1))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = v16
            }
        case "hello":
            fmt.Println(v7)
        default:
            v20 := GenS3{F0: 142, F1: (v6 ^ v6)}
            _ = v20
        }
        _ = v10
    } else {
        fmt.Println(v7)
        v21 := GenS4{F0: 2, F1: -961, F2: "hello"}
        _ = v21
    }
    var v22 string = ("hello" + v2)
    if true {
        v4 += (^(v4))
        _ = len(v1)
    }
    fmt.Println(v22)
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := GenS4{F0: 4294967295, F1: fi23, F2: v22}
        _ = fi23
        _ = v24
    }
    var v25 float64 = float64((((1.5 / -2.25) + float32(1.5)) / 0.1))
    var v26 bool = false
    var v27 uint = 350
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
