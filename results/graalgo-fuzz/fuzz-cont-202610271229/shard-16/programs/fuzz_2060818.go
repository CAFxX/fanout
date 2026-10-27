package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint64; F1 string }
func genFn0(p0 string, p1 uint16) int32 {
    var v7 uint64 = 18446744073709551615
    var v8 int16 = 32767
    for fi9 := 0; fi9 < 14; fi9++ {
        v8 -= (^(v8))
        v10 := []int{fi9, fi9, fi9}
        _ = fi9
        _ = v10
    }
    var v11 float32 = 0.0
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1(p0 uint8, p1 int64, p2 int8) int64 {
    fmt.Println(p1)
    var v12 uint32 = 587
    var v13 int16 = 42
    v14 := GenS1{F0: (p0 << ((1) & 63)), F1: -181}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return (p1 + (p1 - (p1 % ((int64(v12)) | 1))))
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]uint32)
    var v3 int32 = 42
    v4 := make(map[string]uint32)
    var v5 uint8 = 35
    var v6 uint = 18446744073709551615
    if (42 <= 32768) {
        if (((v5 | (v5 >> ((v6) & 63))) < 128) && (!((!true) && false))) {
            v15 := GenS2{F0: (^(v5)), F1: 402}
            if true {
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 bool = false
                    _ = fi16
                    _ = v17
                }
                var v18 int32 = v3
                _ = v18
            }
            v5 += 42
            _ = v15
        } else {
            for fi19 := 0; fi19 < 17; fi19++ {
                v3 = (v3 | (v3 / ((genFn0(("a" + "foo bar"), 551)) | 1)))
                _ = fi19
            }
            v1 = 846
        }
    } else {
        var v20 uint64 = 963
        var v21 int32 = (v3 / ((v3) | 1))
        _ = v20
        _ = v21
    }
    var v22 float32 = float32(float64((0.1 - (0.1 + float32(3.14159)))))
    for fi23 := 0; fi23 < 2; fi23++ {
        delete(v2, "")
        v6 += v6
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := GenS3{F0: (!false), F1: 0}
            if false {
                var v26 uint16 = (v1 << ((v6) & 63))
                fmt.Println(v22, v5)
                var v27 bool = (v22 <= v22)
                _ = v26
                _ = v27
            }
            v22 -= v22
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    v28 := []int{-2147483648, 508, 1, 1}
    fmt.Println(v1, v5)
    var v29 uint = 946
    var v30 int16 = 183
    v1 = (v1 % (((v1 - uint16(v5))) | 1))
    v5 = v5
    v31 := GenS4{F0: 332, F1: "foo bar"}
    v29 -= uint(v5)
    for fi32 := 0; fi32 < 1; fi32++ {
        var v33 bool = ((!((v1 << (((v6 % ((42) | 1))) & 63)) == (v1 + v1))) || (v5 >= (v5 >> ((v6) & 63))))
        _ = fi32
        _ = v33
    }
    v34 := GenS1{F0: (v5 >> (((^(v6))) & 63)), F1: 42}
    var v35 uint16 = v1
    v1 = v35
    v1 -= v35
    v4[""] = 0
    v29 -= (^(v29))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
