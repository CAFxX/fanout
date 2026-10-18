package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int16; F1 int64; F2 bool }
type GenS6 struct { F0 int16; F1 string }
func genFn0(p0 int, p1 string) int32 {
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: 42, F1: false, F2: 42}
        v8 := GenS1{F0: 68, F1: ((652 != 42) && (2147483647 < 378)), F2: 255}
        _ = fi6
        _ = v7
        _ = v8
    }
    if (((1 == 511) || false) && false) {
        var v9 uint16 = 381
        _ = v9
    } else {
        var v10 int32 = 142
        switch (p0 % (((-(p0))) | 1)) {
        case 445:
            for fi11 := 0; fi11 < 14; fi11++ {
                switch (p0 ^ 9223372036854775807) {
                case 32768:
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v13 := []int{fi12, fi12, fi12}
                        _ = fi12
                        _ = v13
                    }
                case -2147483648:
                    v14 := GenS2{F0: 470, F1: 0}
                    _ = v14
                case 14:
                    switch "hello" {
                    case "\n\t":
                        switch p1 {
                        case "":
                            v15 := GenS3{F0: 9223372036854775807, F1: v10}
                            _ = v15
                        default:
                            if false {
                                p1 = p1
                                if true {
                                    if (328 <= 362) {
                                        fmt.Println(p1, v10)
                                    }
                                }
                            }
                        }
                    default:
                        v16 := []int{(-(p0))}
                        _ = v16
                    }
                }
                p0 = fi11
                _ = fi11
            }
        case 128:
            var v17 float64 = 1e10
            _ = v17
        }
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 int = fi18
            _ = fi18
            _ = v19
        }
        _ = v10
    }
    if (642 <= 4294967295) {
        var v20 uint16 = 42
        var v21 uint8 = 42
        _ = v20
        _ = v21
    }
    v22 := []int{p0, -1, 32768, (p0 ^ p0)}
    _ = p0
    _ = p1
    _ = v22
    return -2147483648
}
func genFn1(p0 int) int8 {
    var v23 int16 = 42
    v24 := GenS1{F0: 128, F1: false, F2: 42}
    if true {
        var v25 float64 = 0.0
        _ = v25
    } else {
        var v26 bool = false
        _ = v26
    }
    v27 := GenS3{F0: -9223372036854775808, F1: -45}
    _ = p0
    _ = v23
    _ = v24
    _ = v27
    return 115
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := []string{}
    v4 := make(map[string]int)
    var v5 uint8 = 255
    var v28 int8 = -128
    if (v28 >= (v28 ^ v28)) {
        v5 -= v5
    } else {
        switch 32768 {
        case 0:
            if false {
                v29 := GenS4{F0: "\n\t", F1: 2}
                _ = v29
            }
        }
        var v30 float32 = 0.1
        _ = v30
    }
    v28 = (v28 + genFn1(int(v28)))
    v5 = (v5 - 252)
    if true {
        if (!true) {
            fmt.Println(v28, v28)
        }
    }
    switch "\n\t" {
    case "":
        var v31 int16 = 1
        _ = v31
    default:
        v5 = v5
    }
    switch 32767 {
    case 127:
        fmt.Println(v5, v5)
    case 53:
        delete(v4, ("" + "foo bar"))
    case 845:
        v5 += (v5 % (((v5 ^ 0)) | 1))
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        v5 = 1
        v33 := GenS2{F0: 0, F1: 0}
        v34 := GenS3{F0: 0, F1: 1}
        _ = fi32
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        fmt.Println(v5, v28, v5)
        _ = fi35
    }
    if false {
        v5 = (v5 << ((535) & 63))
        v36 := GenS5{F0: 32767, F1: 9223372036854775807, F2: false}
        _ = v36
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        v28 = (-(v28))
        var v38 uint32 = 0
        v39 := GenS6{F0: 42, F1: "a"}
        _ = fi37
        _ = v38
        _ = v39
    }
    v5 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
