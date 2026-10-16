package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint32; F2 bool }
type GenS3 struct { F0 int8; F1 string; F2 string }
type GenS4 struct { F0 int8; F1 int16 }
func genFn0(p0 uint64, p1 uint64, p2 int16) uint {
    p1 = p1
    p2 = (p2 - 0)
    for fi4 := 0; fi4 < 9; fi4++ {
        if false {
            v5 := []int{fi4}
            var v6 uint = 909
            _ = v5
            _ = v6
        }
        var v7 uint32 = 2147483647
        _ = fi4
        _ = v7
    }
    v8 := []int{-426, 352}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}
func genFn1(p0 uint8) uint8 {
    p0 = p0
    v9 := GenS1{F0: 223}
    p0 -= p0
    switch "\n\t" {
    case "hello":
        v10 := []int{32768}
        _ = v10
    case "\n\t":
        var v11 float32 = ((1.5 - 1e10) + float32(3.14159))
        _ = v11
    default:
        v12 := []int{-2147483648, 128, 9223372036854775807, 4294967295}
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        p0 = 52
        _ = fi13
    }
    _ = p0
    _ = v9
    return (^(p0))
}

func main() {
    v1 := []int{-812, 4294967295, 9223372036854775807, 184}
    v2 := []string{}
    v3 := []uint32{65535}
    if false {
        switch "foo bar" {
        case "foo bar":
            if (2 <= 1) {
                switch -345 {
                case 444:
                    for fi14 := 0; fi14 < 9; fi14++ {
                        for fi15 := 0; fi15 < 1; fi15++ {
                            var v16 int64 = 145
                            var v17 int16 = -32768
                            if true {
                                if (!true) {
                                    v17 = v17
                                    v17 = v17
                                    v17 = v17
                                }
                                v18 := GenS2{F0: (^(v17)), F1: 1, F2: (!false)}
                                _ = v18
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(fi14, v17)
                            }
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        v19 := GenS3{F0: 8, F1: ("a" + "\n\t"), F2: ""}
                        _ = fi14
                        _ = v19
                    }
                case -2147483648:
                    if false {
                        if (false || false) {
                            v3 = append(v3, 2147483648)
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                        }
                        var v20 bool = (!(!false))
                        fmt.Println(v20)
                        _ = v20
                    } else {
                        var v21 uint = 18446744073709551615
                        var v22 int8 = 127
                        v23 := GenS3{F0: v22, F1: ("\n\t" + ""), F2: "a"}
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                case 0:
                    var v24 string = "foo bar"
                    _ = v24
                default:
                    v25 := GenS3{F0: 127, F1: "a", F2: ""}
                    _ = v25
                }
                var v26 int32 = -224
                _ = v26
            }
        }
        v27 := GenS4{F0: 127, F1: 1000}
        _ = v27
    }
    var v28 bool = (!false)
    v29 := GenS1{F0: 2}
    v28 = v28
    for fi30 := 0; fi30 < 8; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            fmt.Println(v28, fi31, v28)
            v32 := GenS3{F0: 122, F1: "a", F2: "\n\t"}
            var v33 bool = true
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    var v34 uint32 = 1
    var v35 float32 = 1e10
    v34 = (v34 / ((v34) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v35
}
