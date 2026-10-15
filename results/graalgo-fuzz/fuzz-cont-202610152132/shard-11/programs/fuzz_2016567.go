package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 string; F2 bool }
type GenS4 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint64, p2 int8) int32 {
    if (566 < 4294967295) {
        for fi4 := 0; fi4 < 17; fi4++ {
            v5 := []int{}
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 uint = 65535
                var v8 uint64 = (^(p1))
                p1 = (^(v8))
                _ = fi6
                _ = v7
                _ = v8
            }
            p1 = (p1 >> ((662) & 63))
            _ = fi4
            _ = v5
        }
        p1 += 42
    } else {
        for fi9 := 0; fi9 < 2; fi9++ {
            p1 += p1
            _ = fi9
        }
    }
    p2 -= 0
    v10 := GenS1{F0: 32767, F1: 32767}
    v11 := []int{-2147483648, 0}
    var v12 uint = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 0
}
func genFn1(p0 uint) int8 {
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 bool = true
        v15 := GenS2{F0: (p0 << ((255) & 63))}
        for fi16 := 0; fi16 < 13; fi16++ {
            v17 := GenS2{F0: (p0 - (p0 & p0))}
            v18 := GenS2{F0: (p0 - (p0 | 272))}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    var v19 string = ("a" + "a")
    fmt.Println(p0, p0)
    p0 += p0
    var v20 int = -9223372036854775808
    _ = p0
    _ = v19
    _ = v20
    return -1
}

func main() {
    var v1 string = ("hello" + "foo bar")
    var v2 uint8 = 128
    var v3 float64 = 1e10
    v21 := []int{}
    v21 = append(v21, 1)
    v22 := GenS3{F0: 9223372036854775807, F1: "hello", F2: (495 < 732)}
    v21 = append(v21, 128)
    var v23 int8 = 42
    var v24 uint32 = 545
    var v25 int = -562
    v26 := GenS3{F0: 1, F1: "", F2: true}
    var v27 uint16 = 2
    v27 += (v27 | 32767)
    v28 := GenS4{F0: 255, F1: (^(v2)), F2: 0}
    v29 := GenS5{F0: (v25 / (((v25 & v25)) | 1)), F1: (v24 & v24)}
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    v30 := GenS6{F0: v24}
    switch (v1 + ((v1 + v1) + v1)) {
    case "foo bar":
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
    case "a":
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
    }
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
