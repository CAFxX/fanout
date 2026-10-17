package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 int32; F2 uint8 }
func genFn0() int8 {
    var v6 uint32 = 685
    var v7 int32 = -1
    for fi8 := 0; fi8 < 12; fi8++ {
        v6 = 681
        var v9 int8 = 100
        _ = fi8
        _ = v9
    }
    v6 = (v6 & (^(v6)))
    _ = v6
    _ = v7
    return 100
}
func genFn1(p0 uint32) int64 {
    p0 += (^(p0))
    var v10 bool = (!true)
    var v11 uint16 = 517
    _ = p0
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 string = "a"
    var v2 int16 = 77
    var v3 int64 = 0
    var v4 uint64 = 0
    var v5 uint64 = v4
    if ((false || false) && true) {
        v12 := GenS1{F0: v4, F1: 272, F2: 4294967295}
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int64 = v3
            v16 := GenS1{F0: (v4 % ((1) | 1)), F1: 42, F2: 0}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    fmt.Println(v5, v1, v1)
    for fi17 := 0; fi17 < 12; fi17++ {
        if (!false) {
            if false {
                var v18 int64 = (v3 + genFn1(0))
                switch 4294967295 {
                case 9223372036854775807:
                    if true {
                        for fi19 := 0; fi19 < 1; fi19++ {
                            switch ((v1 + (((v1 + v1) + (v1 + (v1 + (v1 + (v1 + (v1 + "hello")))))) + ("a" + v1))) + ((v1 + v1) + v1)) {
                            case "":
                                v2 += 0
                            case "a":
                                v2 = (v2 * (v2 << ((844) & 63)))
                            case "hello":
                                fmt.Println(v18, fi19)
                            default:
                                if ((false || (0 <= -1)) && false) {
                                    v20 := []int{}
                                    var v21 uint8 = 42
                                    v2 = (v2 | (v2 - (v2 & (v2 + int16(v18)))))
                                    _ = v20
                                    _ = v21
                                } else {
                                    v2 = (^(v2))
                                    fmt.Println(v3, v5)
                                    var v22 float32 = float32(0.0)
                                    _ = v22
                                }
                            }
                            _ = fi19
                        }
                    }
                case 1:
                    if true {
                        v18 = (v3 + genFn1(42))
                    }
                case 2147483647:
                    v23 := []int{fi17, fi17}
                    _ = v23
                }
                fmt.Println(v4, v4, v1)
                _ = v18
            }
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 float32 = 1e10
                _ = fi24
                _ = v25
            }
        } else {
            v3 += (v3 + (v3 & (v3 % ((v3) | 1))))
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 float32 = (0.1 / (1.5 + 3.14159))
                var v28 uint64 = (^(v4))
                _ = fi26
                _ = v27
                _ = v28
            }
            v5 += (v4 & (v5 << ((42) & 63)))
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            v30 := []int{fi17, fi29, (fi29 | (^(fi17)))}
            var v31 int8 = genFn0()
            fmt.Println(v3, v4)
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi17
    }
    if (32767 == v2) {
        v32 := []int{42, 127, 1, -1}
        v33 := GenS1{F0: 42, F1: 1, F2: 832}
        v3 += 1
        _ = v32
        _ = v33
    } else {
        switch ("\n\t" + (v1 + ((((("\n\t" + v1) + v1) + (v1 + ("\n\t" + (v1 + v1)))) + v1) + v1))) {
        case "\n\t":
            v1 = "a"
        case "":
            v34 := []int{4294967295}
            _ = v34
        default:
            var v35 bool = (0 >= 597)
            _ = v35
        }
        fmt.Println(v1)
        if true {
            var v36 uint8 = 144
            for fi37 := 0; fi37 < 8; fi37++ {
                v38 := GenS1{F0: (v4 & (v5 - (v4 | (v4 + (^(v4)))))), F1: 2147483648, F2: 1}
                v4 = (v5 / ((0) | 1))
                if false {
                    v1 = "foo bar"
                    v39 := GenS1{F0: v5, F1: 849, F2: uint(v36)}
                    _ = v39
                }
                _ = fi37
                _ = v38
            }
            _ = v36
        }
    }
    v40 := []int{32767}
    var v41 uint64 = (^(v4))
    v42 := GenS2{F0: genFn1(2147483647)}
    for fi43 := 0; fi43 < 13; fi43++ {
        _ = len(v40)
        v44 := GenS3{F0: 816, F1: 10, F2: 255}
        fmt.Println(v5)
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
