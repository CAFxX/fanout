package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int64 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 string }
func genFn0(p0 uint8) uint8 {
    v5 := GenS1{F0: 32767, F1: -544, F2: -1}
    for fi6 := 0; fi6 < 19; fi6++ {
        if (true && (-193 >= 42)) {
            v7 := GenS1{F0: 42, F1: 1, F2: 42}
            var v8 int8 = -1
            switch (fi6 >> ((2) & 63)) {
            case 1:
                switch "foo bar" {
                case "hello":
                    p0 = 1
                case "a":
                    v9 := []int{(fi6 % ((591) | 1)), fi6, 32768, fi6}
                    _ = v9
                }
            case -885:
                var v10 uint8 = p0
                _ = v10
            default:
                if true {
                    for fi11 := 0; fi11 < 9; fi11++ {
                        v12 := GenS2{F0: 722, F1: 1}
                        v13 := GenS3{F0: 252}
                        switch fi11 {
                        case 1:
                            p0 = (p0 + (p0 % ((1) | 1)))
                        default:
                            var v14 int = (fi11 + (fi6 << ((666) & 63)))
                            _ = v14
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    for fi15 := 0; fi15 < 19; fi15++ {
                        var v16 uint = 0
                        v16 = (v16 + (v16 >> (((v16 & 514)) & 63)))
                        _ = fi15
                        _ = v16
                    }
                    if true {
                        for fi17 := 0; fi17 < 2; fi17++ {
                            if true {
                                v18 := GenS4{F0: 448, F1: "\n\t"}
                                _ = v18
                            } else {
                                fmt.Println(fi6, fi6, fi6)
                                var v19 float64 = ((1.5 - 0.0) + -2.25)
                                _ = v19
                            }
                            _ = fi17
                        }
                    }
                }
            }
            _ = v7
            _ = v8
        }
        _ = fi6
    }
    _ = p0
    _ = v5
    return p0
}

func main() {
    var v1 int32 = -1
    v2 := []string{}
    v3 := make(map[uint32]string)
    v4 := []uint32{65535}
    v1 -= v1
    for fi20 := 0; fi20 < 5; fi20++ {
        fmt.Println(v1)
        _ = fi20
    }
    var v21 float32 = -2.25
    switch "\n\t" {
    case "":
        var v22 uint32 = 2147483648
        _ = v22
    case "\n\t":
        v4 = append(v4, 1)
    default:
        var v23 float64 = ((-2.25 - 1.5) * 0.1)
        _ = v23
    }
    var v24 uint = 255
    var v25 int = 0
    if true {
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 string = "\n\t"
            _ = fi26
            _ = v27
        }
        var v28 uint = (v24 / (((v24 ^ 255)) | 1))
        _ = v28
    }
    if false {
        _ = len(v2)
    }
    v24 = (v24 % (((v24 & 18446744073709551615)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
