package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int8 }
func genFn0(p0 int64) uint32 {
    for fi7 := 0; fi7 < 19; fi7++ {
        fmt.Println(p0, fi7)
        p0 = (-(p0))
        _ = fi7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := []int{fi8, (fi8 | (^(fi8))), (fi8 | (fi8 * 42)), fi8}
        _ = fi8
        _ = v9
    }
    _ = p0
    return 2147483648
}
func genFn1(p0 int64, p1 string) uint {
    var v10 uint64 = 449
    v11 := []int{42, 793}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 18446744073709551615
}

func main() {
    var v1 string = (("a" + ((("foo bar" + ("a" + "\n\t")) + "hello") + "hello")) + "")
    var v2 uint32 = 42
    v3 := make(map[int]int)
    var v4 uint32 = 2147483648
    var v5 uint = 69
    var v6 int64 = 1
    switch "foo bar" {
    case "hello":
        var v12 int32 = -357
        _ = v12
    case "a":
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 bool = ((!(v5 <= (v5 % ((v5) | 1)))) || (!true))
            var v15 uint32 = 65535
            _ = fi13
            _ = v14
            _ = v15
        }
    case "":
        if (4294967295 >= 663) {
            var v16 uint = 51
            v6 = 42
            _ = v16
        } else {
            for fi17 := 0; fi17 < 8; fi17++ {
                v18 := GenS1{F0: ("a" + v1)}
                v19 := GenS1{F0: "\n\t"}
                var v20 float64 = 3.14159
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            var v21 int8 = 121
            _ = v21
        }
    default:
        if true {
            v22 := GenS2{F0: 128, F1: false}
            v4 -= (^(v2))
            _ = v22
        }
    }
    if true {
        v2 -= v4
        var v23 int32 = 0
        delete(v3, 9223372036854775807)
        _ = v23
    }
    v2 = v4
    v24 := []int{-765, 42, 987}
    fmt.Println(v2, v2)
    v25 := GenS1{F0: v1}
    if false {
        v4 = v2
        var v26 uint8 = 42
        if ("a" == "\n\t") {
            var v27 float32 = 0.0
            if ((!(!false)) && (v1 >= v1)) {
                switch ("foo bar" + v1) {
                case "\n\t":
                    v5 += (v5 << (((^(v5))) & 63))
                }
                for fi28 := 0; fi28 < 11; fi28++ {
                    switch -56 {
                    case -1:
                        if ((^(v5)) > (^(v5))) {
                            v29 := v3[fi28]
                            _ = v29
                        }
                    default:
                        var v30 int64 = (v6 & v6)
                        _ = v30
                    }
                    var v31 int = -2147483648
                    switch "a" {
                    case "":
                        fmt.Println(v26)
                    case "hello":
                        var v32 float32 = v27
                        _ = v32
                    case "a":
                        for fi33 := 0; fi33 < 16; fi33++ {
                            var v34 float32 = 1e10
                            if true {
                                switch v1 {
                                case "a":
                                    var v35 int64 = v6
                                    _ = v35
                                case "":
                                    if false {
                                        fmt.Println(v34, v2, v5)
                                        fmt.Println(v6)
                                        fmt.Println(v26, v4)
                                    }
                                case "hello":
                                    v27 = 3.14159
                                default:
                                    v2 += genFn0((-(v6)))
                                }
                            } else {
                                fmt.Println(v6)
                                v6 = -393
                            }
                            var v36 int32 = -993
                            _ = fi33
                            _ = v34
                            _ = v36
                        }
                    default:
                        if false {
                            for fi37 := 0; fi37 < 17; fi37++ {
                                v27 = ((v27 + (1e10 / v27)) * float32(0.1))
                                v2 = 42
                                _ = fi37
                            }
                            v38 := GenS1{F0: "foo bar"}
                            _ = v38
                        }
                    }
                    _ = fi28
                    _ = v31
                }
                var v39 int16 = 1000
                _ = v39
            }
            _ = v27
        }
        _ = v26
    }
    if len(v24) > 0 {
        v24[0] = 823
    }
    v6 -= v6
    switch 895 {
    case -1:
        switch 42 {
        case 192:
            v40 := GenS2{F0: 42, F1: false}
            _ = v40
        case 127:
            v41 := GenS3{F0: 1}
            _ = v41
        default:
            v42 := v3[128]
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 3; fi43++ {
        for fi44 := 0; fi44 < 20; fi44++ {
            v45 := GenS3{F0: 3}
            if (86 != -128) {
                var v46 int = fi44
                _ = v46
            }
            var v47 uint8 = 193
            _ = fi44
            _ = v45
            _ = v47
        }
        v24 = append(v24, int(v5))
        _ = fi43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
