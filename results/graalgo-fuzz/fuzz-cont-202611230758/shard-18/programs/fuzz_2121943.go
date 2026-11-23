package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 bool; F1 int16; F2 int8 }
func genFn0(p0 int8) string {
    v6 := GenS1{F0: 0, F1: (true || (128 <= 0)), F2: 32767}
    p0 = 0
    switch (("\n\t" + "\n\t") + "foo bar") {
    case "\n\t":
        p0 -= (p0 << ((1) & 63))
    }
    v7 := []int{-2147483648, -916}
    _ = p0
    _ = v6
    _ = v7
    return "foo bar"
}
func genFn1(p0 uint64, p1 uint, p2 int) uint32 {
    switch "foo bar" {
    case "foo bar":
        v8 := []int{int(p1)}
        _ = v8
    case "a":
        p1 += p1
    default:
        v9 := GenS1{F0: 42, F1: true, F2: p2}
        _ = v9
    }
    switch -668 {
    case 0:
        v10 := []int{}
        _ = v10
    case -265:
        for fi11 := 0; fi11 < 8; fi11++ {
            switch "a" {
            case "a":
                p1 = 65535
            case "":
                switch int(p0) {
                case 32767:
                    var v12 int64 = 799
                    _ = v12
                case 2147483647:
                    v13 := GenS2{F0: 0}
                    _ = v13
                case 32768:
                    v14 := GenS1{F0: 1, F1: false, F2: 4294967295}
                    _ = v14
                default:
                    v15 := []int{fi11}
                    _ = v15
                }
            default:
                switch p2 {
                case 1:
                    v16 := []int{(fi11 ^ (p2 | fi11)), -310, fi11}
                    _ = v16
                case 923:
                    fmt.Println(fi11)
                case 4294967295:
                    if true {
                        fmt.Println(p2)
                        var v17 int8 = -128
                        _ = v17
                    } else {
                        for fi18 := 0; fi18 < 5; fi18++ {
                            v19 := GenS1{F0: 42, F1: false, F2: p2}
                            if (1 != int16(p1)) {
                                v20 := GenS1{F0: 42, F1: true, F2: fi11}
                                p0 = 4294967295
                                _ = v20
                            } else {
                                var v21 float32 = float32((1.5 * 3.14159))
                                _ = v21
                            }
                            _ = fi18
                            _ = v19
                        }
                        var v22 int16 = -32768
                        p1 = p1
                        _ = v22
                    }
                }
            }
            var v23 int32 = 811
            _ = fi11
            _ = v23
        }
    case 4294967295:
        p2 += (p2 * p2)
    default:
        v24 := GenS3{F0: -1, F1: 384}
        _ = v24
    }
    var v25 string = ""
    for fi26 := 0; fi26 < 6; fi26++ {
        v27 := []int{}
        fmt.Println(v25)
        _ = fi26
        _ = v27
    }
    var v28 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v28
    return 2147483647
}

func main() {
    var v1 float64 = -2.25
    v2 := []int{2147483647, 858, -2147483648, 2147483647}
    var v3 int = 32767
    v4 := make(map[int]uint32)
    var v5 uint8 = 2
    v3 -= 32768
    if (true && (v1 != 3.14159)) {
        v29 := GenS1{F0: 255, F1: false, F2: -211}
        if false {
            switch (v3 | 4294967295) {
            case -9223372036854775808:
                switch -9223372036854775808 {
                case -2147483648:
                    fmt.Println(v3, v5)
                }
            case -107:
                if (!(2147483647 > 2147483647)) {
                    v3 += 128
                } else {
                    if (!true) {
                        var v30 float32 = 1.5
                        _ = v30
                    } else {
                        v31 := GenS1{F0: v5, F1: (("" + "hello") == "a"), F2: v3}
                        _ = v31
                    }
                    v3 += 0
                }
            }
        }
        fmt.Println(v3, v5, v3)
        _ = v29
    }
    fmt.Println(v5)
    for fi32 := 0; fi32 < 9; fi32++ {
        v5 += (v5 ^ v5)
        _ = fi32
    }
    v33 := GenS2{F0: 1}
    v2 = append(v2, (v3 / (((v3 & -967)) | 1)))
    var v34 int = (v3 ^ v3)
    v35 := GenS4{F0: true, F1: -135, F2: 1}
    v36 := GenS3{F0: -2147483648, F1: uint16(v5)}
    switch (v3 | 1) {
    case 32768:
        v1 = 0.0
    default:
        v4[v34] = genFn1(509, 803, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
