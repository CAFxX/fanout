package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 string }
type GenS2 struct { F0 int64; F1 bool; F2 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 int32 }
func genFn0(p0 int32, p1 string) int16 {
    for fi6 := 0; fi6 < 10; fi6++ {
        if ((2 > 158) && ((42 > 41) || false)) {
            p1 = "a"
            for fi7 := 0; fi7 < 12; fi7++ {
                v8 := []int{-1, (fi7 * (fi6 * fi7)), -171}
                _ = len(v8)
                _ = fi7
                _ = v8
            }
            if ((((float32(((-2.25 / 1e10) * float64(1e10))) * float32(0.0)) - ((float32(float64(1.5)) / 1.0) * 0.0)) <= (0.1 + 1e10)) || false) {
                fmt.Println(p1, p1)
                for fi9 := 0; fi9 < 20; fi9++ {
                    switch fi9 {
                    case 2147483647:
                        var v10 float32 = (1e10 + (1e10 * 0.1))
                        _ = v10
                    case 1000:
                        p1 = p1
                    case 827:
                        switch 237 {
                        case -2147483648:
                            v11 := GenS1{F0: 42, F1: 0, F2: "a"}
                            _ = v11
                        case 128:
                            v12 := []int{(-(fi9)), (fi9 | fi6), (fi9 + fi9)}
                            _ = v12
                        }
                    default:
                        fmt.Println(p1, fi9)
                    }
                    _ = fi9
                }
            }
        }
        switch fi6 {
        case 812:
            var v13 int8 = -1
            _ = v13
        case -498:
            switch p1 {
            case "foo bar":
                v14 := []int{fi6, (fi6 * (fi6 + fi6)), fi6}
                _ = v14
            case "hello":
                for fi15 := 0; fi15 < 4; fi15++ {
                    v16 := GenS1{F0: 726, F1: 491, F2: p1}
                    _ = fi15
                    _ = v16
                }
            case "":
                switch "foo bar" {
                case "a":
                    if true {
                        for fi17 := 0; fi17 < 15; fi17++ {
                            p1 = "\n\t"
                            var v18 uint32 = 4294967295
                            p0 += (^(p0))
                            p1 = ((p1 + "hello") + p1)
                            v19 := GenS2{F0: 9223372036854775807, F1: (!(!false)), F2: 2147483647}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    }
                case "\n\t":
                    var v20 bool = (-2.25 >= -2.25)
                    _ = v20
                case "hello":
                    fmt.Println(p1, fi6)
                }
            }
        case 2147483647:
            if false {
                var v21 uint64 = 495
                var v22 uint64 = v21
                _ = v21
                _ = v22
            }
        }
        if ((float32(float64(0.0)) > 0.1) || false) {
            v23 := GenS3{F0: fi6}
            v24 := GenS3{F0: fi6}
            _ = v23
            _ = v24
        }
        _ = fi6
    }
    if false {
        var v25 uint8 = 128
        var v26 int32 = (p0 << ((65535) & 63))
        v27 := GenS3{F0: -1}
        _ = v25
        _ = v26
        _ = v27
    }
    v28 := GenS3{F0: -1}
    _ = p0
    _ = p1
    _ = v28
    return 461
}

func main() {
    var v1 string = "hello"
    var v2 int64 = -9223372036854775808
    v3 := make(map[uint32]int)
    v4 := []uint32{2147483647, 639}
    v5 := make(map[int]bool)
    v2 = -1
    var v29 uint32 = 447
    for fi30 := 0; fi30 < 2; fi30++ {
        if false {
            fmt.Println(v1)
            var v31 bool = (220 > 174)
            switch (fi30 % ((fi30) | 1)) {
            case 32767:
                _ = len(v5)
            case 0:
                v2 += (v2 % (((v2 / ((v2) | 1))) | 1))
            }
            _ = v31
        } else {
            v29 = (v29 + (v29 - (v29 >> ((65535) & 63))))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v4) > 0 {
                v4[0] = 960
            }
        }
        delete(v5, fi30)
        _ = fi30
    }
    v2 = (v2 % ((-1) | 1))
    v29 += v29
    v32 := GenS3{F0: 32767}
    for fi33 := 0; fi33 < 6; fi33++ {
        if (522 < 1) {
            v34 := GenS4{F0: 2147483647, F1: 127}
            _ = v34
        }
        _ = fi33
    }
    v29 = v29
    switch "foo bar" {
    case "foo bar":
        v35 := GenS2{F0: v2, F1: false, F2: 936}
        _ = v35
    }
    v36 := GenS5{F0: -1, F1: 251}
    _ = len(v4)
    v3[(v29 & (v29 | (^(v29))))] = 867
    for fi37 := 0; fi37 < 17; fi37++ {
        v1 = "\n\t"
        _ = len(v3)
        switch (fi37 << ((247) & 63)) {
        case 127:
            v29 -= (v29 ^ 2)
        }
        _ = fi37
    }
    v38 := GenS6{F0: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v36
    _ = v38
    _ = v4
    _ = v5
}
