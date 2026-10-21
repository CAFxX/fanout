package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int8 }
type GenS2 struct { F0 string; F1 uint16; F2 int8 }
type GenS3 struct { F0 int8; F1 uint32; F2 string }
func genFn0(p0 string, p1 string, p2 uint16) int32 {
    var v6 uint16 = 415
    if (0 != 1) {
        var v7 int64 = -1
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int16 = -316
            v10 := GenS1{F0: 393, F1: (-(v9)), F2: 0}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    } else {
        v11 := []int{32768, -570}
        _ = v11
    }
    v12 := []int{}
    var v13 bool = ((!false) && true)
    var v14 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    _ = v6
    return -200
}
func genFn1(p0 uint8, p1 int64) int16 {
    p0 = p0
    v15 := []int{559, 32768}
    p0 += (p0 / (((p0 << ((4294967295) & 63))) | 1))
    _ = p0
    _ = p1
    _ = v15
    return -831
}

func main() {
    var v1 uint8 = 0
    v2 := []string{("\n\t" + "\n\t"), "", "foo bar"}
    v3 := make(map[uint32]bool)
    var v4 int64 = 916
    v5 := make(map[int]bool)
    fmt.Println(v4, v4)
    v5[9223372036854775807] = false
    var v16 uint = 4
    var v17 bool = (!((true && (2 >= 155)) && (0.0 != 1e10)))
    switch 1 {
    case -2147483648:
        v2 = append(v2, ("foo bar" + (((("\n\t" + (("hello" + "\n\t") + "a")) + ("\n\t" + "")) + (("\n\t" + "hello") + "foo bar")) + "foo bar")))
    case 128:
        switch 1 {
        case -9223372036854775808:
            fmt.Println(v16)
        default:
            if (0 >= 65535) {
                switch "hello" {
                case "hello":
                    if v17 {
                        v4 = (v4 ^ 478)
                        if ("" > "foo bar") {
                            if (310 > -1) {
                                switch "" {
                                case "a":
                                    v16 = (^(v16))
                                case "hello":
                                    var v18 string = ("" + "a")
                                    _ = v18
                                default:
                                    var v19 uint16 = 222
                                    _ = v19
                                }
                                delete(v3, 448)
                            } else {
                                if v17 {
                                    for fi20 := 0; fi20 < 20; fi20++ {
                                        v21 := GenS1{F0: 2147483647, F1: -1, F2: -17}
                                        if (((false && (1000 > 1)) && true) || (((v16 / ((v16) | 1)) <= 65535) || (!v17))) {
                                            fmt.Println(v16, fi20, v4)
                                        }
                                        _ = fi20
                                        _ = v21
                                    }
                                } else {
                                    for fi22 := 0; fi22 < 14; fi22++ {
                                        fmt.Println(v1, v16, fi22)
                                        fmt.Println(v16)
                                        _ = fi22
                                    }
                                    for fi23 := 0; fi23 < 13; fi23++ {
                                        if v17 {
                                            v24 := GenS1{F0: 65535, F1: 1000, F2: -46}
                                            v3[403] = true
                                            _ = v24
                                        }
                                        _ = fi23
                                    }
                                }
                                v1 -= v1
                            }
                        } else {
                            v25 := GenS1{F0: 1, F1: 1, F2: -1}
                            _ = v25
                        }
                    } else {
                        if false {
                            v17 = false
                        }
                        v26 := GenS2{F0: ("a" + ("foo bar" + "a")), F1: 42, F2: 127}
                        _ = v26
                    }
                case "a":
                    var v27 float32 = (1.5 - (float32(1e10) * 1.5))
                    _ = v27
                }
            }
        }
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        switch ("foo bar" + ("hello" + "a")) {
        case "hello":
            var v29 float32 = 0.0
            _ = v29
        case "":
            var v30 int = (^(fi28))
            _ = v30
        case "a":
            v31 := GenS2{F0: "\n\t", F1: 32767, F2: 100}
            _ = v31
        }
        v32 := GenS1{F0: 876, F1: -811, F2: 0}
        var v33 int8 = -12
        _ = fi28
        _ = v32
        _ = v33
    }
    v34 := v5[32768]
    var v35 uint = 623
    fmt.Println(v16)
    if v34 {
        fmt.Println(v4)
    } else {
        v4 = -576
        v36 := GenS3{F0: -1, F1: 2147483648, F2: ""}
        _ = v36
    }
    v17 = true
    v4 = v4
    var v37 int8 = 127
    v35 -= v35
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
