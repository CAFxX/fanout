package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint8 }
func genFn0() uint32 {
    v7 := GenS1{F0: 49}
    var v8 float64 = -2.25
    v8 += float64((1.5 * ((float32(((v8 + v8) * 3.14159)) / float32(0.1)) - float32(float64((3.14159 / (1.5 + 0.1)))))))
    fmt.Println(v8)
    if (!true) {
        if (0 >= -32768) {
            var v9 int32 = 403
            v9 += v9
            _ = v9
        }
        var v10 uint8 = 255
        _ = v10
    }
    _ = v7
    _ = v8
    return 700
}
func genFn1() int16 {
    var v11 uint32 = genFn0()
    v11 = v11
    v12 := GenS1{F0: -9223372036854775808}
    switch "" {
    case "a":
        var v13 bool = (32767 <= -1)
        _ = v13
    case "\n\t":
        var v14 bool = true
        _ = v14
    }
    switch "" {
    case "":
        v11 = (v11 ^ v11)
    case "\n\t":
        var v15 uint64 = 0
        _ = v15
    }
    _ = v11
    _ = v12
    return 32767
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]bool)
    var v3 int8 = 0
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    v6 := make(map[int]string)
    if (!false) {
        v1 -= (v1 % ((v1) | 1))
        if false {
            v3 = v3
            v5[2] = (!true)
        }
        v3 += -63
    } else {
        v1 = (v1 % ((2) | 1))
        var v16 uint64 = 4294967295
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := []int{0, fi17}
            fmt.Println(v3)
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    _ = len(v5)
    var v19 bool = false
    if v19 {
        v3 += 0
    }
    if v19 {
        v1 = v1
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := []int{(fi20 & (fi20 - fi20)), fi20, -413}
            _ = fi20
            _ = v21
        }
    }
    fmt.Println(v1, v19)
    if (((false && (2 >= 42)) && true) && true) {
        v2["\n\t"] = false
        v19 = false
        v19 = (v19 && v19)
    }
    switch 32767 {
    case 1:
        if false {
            var v22 int64 = 1
            _ = v22
        } else {
            var v23 uint8 = 127
            var v24 float32 = 3.14159
            _ = v23
            _ = v24
        }
    case 2147483647:
        delete(v2, "\n\t")
    case -692:
        var v25 bool = false
        _ = v25
    default:
        v3 = v3
    }
    if v19 {
        if v19 {
            var v26 int8 = 1
            var v27 float64 = (0.0 - (float64(1.5) - 0.1))
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 18; fi28++ {
                v19 = false
                var v29 int = 682
                var v30 int8 = (v3 & v3)
                _ = fi28
                _ = v29
                _ = v30
            }
            v31 := GenS2{F0: -624, F1: 42}
            _ = v31
        }
        for fi32 := 0; fi32 < 18; fi32++ {
            if (3.14159 > 3.14159) {
                var v33 uint8 = (^(v1))
                v3 = (v3 + (v3 + v3))
                _ = v33
            } else {
                v3 -= 1
                var v34 uint = 65535
                for fi35 := 0; fi35 < 12; fi35++ {
                    v1 += v1
                    v1 += (v1 >> (((v34 << (((^(v34))) & 63))) & 63))
                    _ = fi35
                }
                _ = v34
            }
            _ = fi32
        }
    }
    var v36 uint = 255
    v37 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
