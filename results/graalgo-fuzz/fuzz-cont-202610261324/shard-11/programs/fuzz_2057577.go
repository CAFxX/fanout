package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint8, p1 uint64, p2 int32) int32 {
    switch ("hello" + "") {
    case "\n\t":
        v6 := []int{int(p1), 127, -1}
        _ = v6
    case "":
        if false {
            v7 := GenS1{F0: 417}
            var v8 int8 = -1
            for fi9 := 0; fi9 < 1; fi9++ {
                switch (fi9 ^ (fi9 | fi9)) {
                case 403:
                    var v10 bool = (!false)
                    _ = v10
                case 4294967295:
                    v11 := GenS2{F0: (p1 < 18446744073709551615), F1: 32767, F2: 156}
                    _ = v11
                case -851:
                    var v12 bool = false
                    _ = v12
                }
                _ = fi9
            }
            _ = v7
            _ = v8
        }
    case "a":
        if false {
            v13 := GenS3{F0: -1}
            switch 42 {
            case -1:
                var v14 float64 = 0.0
                _ = v14
            default:
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 uint = 0
                    for fi17 := 0; fi17 < 11; fi17++ {
                        p2 -= -2147483648
                        _ = fi17
                    }
                    _ = fi15
                    _ = v16
                }
            }
            _ = v13
        } else {
            var v18 uint64 = p1
            for fi19 := 0; fi19 < 6; fi19++ {
                for fi20 := 0; fi20 < 14; fi20++ {
                    p1 -= v18
                    var v21 uint8 = (p0 / (((p0 - 255)) | 1))
                    v22 := []int{(fi19 + (fi19 * (fi20 % ((fi19) | 1))))}
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    var v23 int32 = p2
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = fi19
            }
            v24 := []int{500}
            _ = v18
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        if false {
            v26 := []int{fi25, fi25}
            _ = v26
        } else {
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 int32 = (p2 % (((p2 * p2)) | 1))
                _ = fi27
                _ = v28
            }
            p1 += (^(p1))
        }
        _ = fi25
    }
    var v29 uint = 65535
    v29 -= 0
    p0 = 101
    _ = p0
    _ = p1
    _ = p2
    _ = v29
    return p2
}

func main() {
    v1 := []int{-171}
    var v2 int8 = -1
    v3 := []int{4294967295}
    v4 := make(map[int]bool)
    var v5 uint64 = 0
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    v30 := GenS4{F0: 255, F1: 190}
    for fi31 := 0; fi31 < 18; fi31++ {
        v4[fi31] = false
        v2 += 0
        _ = fi31
    }
    var v32 string = ("a" + "hello")
    var v33 bool = (170 < 2)
    v32 = ("\n\t" + "a")
    if v33 {
        v5 -= v5
        v34 := GenS4{F0: 2, F1: 42}
        _ = v34
    } else {
        v2 -= (v2 | (v2 & v2))
        v35 := GenS5{F0: 1}
        var v36 int32 = -973
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        v38 := GenS1{F0: 1}
        _ = fi37
        _ = v38
    }
    v33 = (367 == genFn0(245, 42, -2147483648))
    for fi39 := 0; fi39 < 13; fi39++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
