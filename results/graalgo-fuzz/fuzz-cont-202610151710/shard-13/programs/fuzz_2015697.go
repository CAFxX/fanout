package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int8; F1 uint64; F2 uint32 }
func genFn0(p0 string) uint8 {
    v4 := GenS1{F0: 65535, F1: -1}
    p0 = p0
    _ = p0
    _ = v4
    return 2
}

func main() {
    var v1 uint32 = 42
    var v2 float64 = (1.5 / 1e10)
    var v3 int16 = -1
    v5 := GenS1{F0: 255, F1: 1}
    v6 := []int{4294967295}
    var v7 int32 = -240
    var v8 bool = false
    if (v8 || false) {
        var v9 uint = 42
        _ = len(v6)
        var v10 int64 = -9223372036854775808
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: 255, F1: 42}
        var v12 uint64 = 42
        var v13 int = 1
        _ = v11
        _ = v12
        _ = v13
    }
    fmt.Println(v2, v3, v7)
    var v14 int16 = (v3 % ((32767) | 1))
    if v8 {
        if v8 {
            var v15 int8 = 42
            fmt.Println(v15, v14, v8)
            var v16 string = "hello"
            _ = v15
            _ = v16
        } else {
            var v17 uint = 28
            _ = v17
        }
    } else {
        for fi18 := 0; fi18 < 11; fi18++ {
            if (!(v8 && v8)) {
                for fi19 := 0; fi19 < 2; fi19++ {
                    if len(v6) > 0 {
                        v6[0] = (fi18 + (fi19 * fi18))
                    }
                    _ = fi19
                }
                if (!v8) {
                    var v20 int64 = 423
                    if (v20 == v20) {
                        for fi21 := 0; fi21 < 18; fi21++ {
                            switch ("a" + "\n\t") {
                            case "":
                                v22 := GenS1{F0: 2, F1: v20}
                                _ = v22
                            case "hello":
                                v7 = 1
                            case "a":
                                v23 := GenS1{F0: 65535, F1: v20}
                                _ = v23
                            default:
                                v24 := GenS2{F0: 0, F1: 1}
                                _ = v24
                            }
                            _ = fi21
                        }
                    } else {
                        fmt.Println(v2)
                        var v25 uint8 = 121
                        _ = v25
                    }
                    _ = v20
                }
                v7 -= v7
            }
            var v26 int32 = (v7 * (v7 / (((v7 & v7)) | 1)))
            _ = fi18
            _ = v26
        }
        var v27 uint32 = v1
        if v8 {
            v7 = (v7 ^ v7)
            v27 -= 4294967295
            var v28 string = ("\n\t" + "hello")
            _ = v28
        } else {
            var v29 int16 = v14
            for fi30 := 0; fi30 < 14; fi30++ {
                v2 += ((1.5 - v2) - v2)
                var v31 float32 = 1e10
                _ = fi30
                _ = v31
            }
            _ = v29
        }
        _ = v27
    }
    if v8 {
        v32 := GenS3{F0: 4, F1: 42, F2: v1}
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
