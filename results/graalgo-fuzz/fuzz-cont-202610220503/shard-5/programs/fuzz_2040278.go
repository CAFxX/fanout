package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int }

func main() {
    var v1 string = ("hello" + ("hello" + "hello"))
    var v2 int64 = -1
    var v3 int32 = 42
    switch 32767 {
    case -306:
        var v4 float32 = 0.0
        _ = v4
    }
    v5 := []int{32767, 9223372036854775807}
    if (!((v3 - v3) != -2147483648)) {
        v1 = ("a" + "hello")
        if false {
            var v6 uint8 = 127
            v7 := GenS1{F0: 42, F1: 1000, F2: 471}
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 19; fi8++ {
                if true {
                    var v9 uint64 = 42
                    var v10 float64 = (float64(1e10) / ((-2.25 / 1.0) - 1e10))
                    v3 += v3
                    _ = v10
                    _ = v9
                }
                v5 = append(v5, (fi8 | (fi8 >> ((255) & 63))))
                _ = fi8
            }
            if false {
                if true {
                    if (!((v2 % (((^(v2))) | 1)) != v2)) {
                        v1 = "hello"
                        v3 = (v3 >> ((1) & 63))
                        _ = len(v5)
                    }
                    v2 = (v2 << ((755) & 63))
                }
            }
            v3 += 0
        }
    } else {
        v11 := GenS2{F0: 491, F1: 1, F2: 1}
        v5 = append(v5, 128)
        _ = len(v5)
        _ = v11
    }
    v12 := GenS1{F0: (^(v2)), F1: 0, F2: 42}
    if false {
        if (100 != -2) {
            fmt.Println(v3)
        }
        v2 = v2
    } else {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 uint = 0
            var v15 uint32 = 1
            var v16 uint = 1
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        v3 = (v3 % ((v3) | 1))
        var v17 bool = true
        _ = v17
    }
    if false {
        if (!false) {
            var v18 int8 = 1
            if len(v5) > 0 {
                v5[0] = int(v3)
            }
            var v19 float32 = ((-2.25 + 0.1) + float32((1e10 / (float64(0.0) + (1.5 / -2.25)))))
            _ = v18
            _ = v19
        } else {
            v20 := GenS2{F0: 42, F1: 829, F2: 42}
            for fi21 := 0; fi21 < 19; fi21++ {
                if len(v5) > 0 {
                    v5[0] = -1
                }
                if (2 == 64) {
                    v2 -= -1
                    v22 := GenS3{F0: 1, F1: fi21}
                    _ = v22
                }
                _ = fi21
            }
            v3 -= (v3 & (^(v3)))
            _ = v20
        }
        if (!false) {
            var v23 uint16 = 255
            _ = v23
        }
    } else {
        var v24 int64 = v2
        v25 := GenS3{F0: 253, F1: -57}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v24
        _ = v25
    }
    switch 0 {
    case -199:
        var v26 int32 = 42
        _ = v26
    default:
        if (!(2147483647 != 2147483647)) {
            var v27 int64 = v2
            _ = v27
        } else {
            fmt.Println(v2, v3, v3)
        }
    }
    v28 := GenS2{F0: 242, F1: 379, F2: 331}
    var v29 int32 = v3
    switch v1 {
    case "hello":
        var v30 float32 = 0.0
        _ = v30
    default:
        v5 = append(v5, 1)
    }
    if (true && (1 > 100)) {
        if (!(false && (255 >= 72))) {
            v3 = v29
        }
        v3 += v3
    } else {
        fmt.Println(v3, v2, v2)
        fmt.Println(v1, v1, v29)
        var v31 float32 = 1e10
        _ = v31
    }
    fmt.Println(v29, v29)
    if false {
        v32 := GenS3{F0: 2147483647, F1: 2147483647}
        _ = v32
    } else {
        v2 = v2
        var v33 uint32 = 65535
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v5
}
