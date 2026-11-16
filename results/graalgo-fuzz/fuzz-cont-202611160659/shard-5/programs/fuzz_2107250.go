package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int64; F1 bool }
func genFn0(p0 int8) int8 {
    p0 -= p0
    for fi5 := 0; fi5 < 2; fi5++ {
        p0 += (p0 / ((123) | 1))
        _ = fi5
    }
    v6 := []int{42, 953, 128, 888}
    fmt.Println(p0)
    _ = p0
    _ = v6
    return -66
}
func genFn1() int32 {
    var v7 bool = false
    var v8 uint = 1
    var v9 int8 = 1
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int8 = 42
    var v2 int8 = (v1 >> ((42) & 63))
    var v3 uint = 0
    var v4 int8 = (v2 | (v1 + (v1 / ((v1) | 1))))
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 int64 = 1
        fmt.Println(v3, v11)
        _ = fi10
        _ = v11
    }
    if (true || false) {
        v12 := GenS1{F0: false, F1: (v1 | v4), F2: 6}
        var v13 int16 = int16(v4)
        _ = v12
        _ = v13
    }
    var v14 string = ""
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 float64 = 0.1
        _ = fi15
        _ = v16
    }
    if true {
        if (122 < 128) {
            v3 = (v3 << ((v3) & 63))
            v4 -= 109
        } else {
            v4 = 69
            var v17 float64 = float64(0.0)
            v18 := GenS2{F0: 479, F1: (v3 << (((v3 - (v3 >> ((v3) & 63)))) & 63))}
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 int32 = 1
        _ = fi19
        _ = v20
    }
    var v21 string = "\n\t"
    var v22 bool = (815 <= 291)
    for fi23 := 0; fi23 < 8; fi23++ {
        for fi24 := 0; fi24 < 14; fi24++ {
            fmt.Println(v2, v22)
            var v25 int = fi23
            _ = fi24
            _ = v25
        }
        v26 := []int{128, (fi23 | fi23), 457, (^(fi23))}
        _ = fi23
        _ = v26
    }
    var v27 uint16 = 2
    fmt.Println(v22)
    fmt.Println(v22, v22)
    v28 := []int{127, -1}
    if true {
        fmt.Println(v22)
        for fi29 := 0; fi29 < 7; fi29++ {
            if len(v28) >= 2 {
                v28 = v28[1:len(v28)]
            }
            for fi30 := 0; fi30 < 8; fi30++ {
                if false {
                    var v31 int8 = (v4 >> ((2) & 63))
                    v4 = -1
                    var v32 int32 = -1
                    v33 := GenS3{F0: 249, F1: (243 <= 1)}
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = fi30
            }
            v14 = ("\n\t" + ("" + "\n\t"))
            _ = fi29
        }
    }
    if len(v28) > 0 {
        v28[0] = 136
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
