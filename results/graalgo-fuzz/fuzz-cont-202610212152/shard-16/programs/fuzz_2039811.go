package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint32; F2 int16 }
func genFn0(p0 int32, p1 int16) string {
    var v5 float64 = 1e10
    v6 := []int{4294967295, 1}
    if false {
        var v7 uint16 = 65535
        _ = v7
    } else {
        fmt.Println(p1, p0, p0)
    }
    if false {
        if (!false) {
            for fi8 := 0; fi8 < 8; fi8++ {
                v5 += v5
                _ = fi8
            }
            if len(v6) > 0 {
                v6[0] = -9223372036854775808
            }
            var v9 int16 = p1
            _ = v9
        }
        fmt.Println(p1, v5)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "a"
}
func genFn1(p0 bool, p1 int) int32 {
    v10 := GenS1{F0: 815, F1: 128}
    switch -1 {
    case 837:
        var v11 int16 = -32768
        _ = v11
    case 9223372036854775807:
        var v12 uint64 = 0
        _ = v12
    case 32768:
        v13 := GenS1{F0: 0, F1: 42}
        _ = v13
    }
    var v14 string = "hello"
    p1 = p1
    switch p1 {
    case 884:
        switch v14 {
        case "hello":
            if (1 >= 1000) {
                p0 = p0
                switch p1 {
                case -2147483648:
                    v15 := []int{-757, p1, (p1 | 1)}
                    _ = v15
                default:
                    for fi16 := 0; fi16 < 12; fi16++ {
                        if p0 {
                            p1 = p1
                            for fi17 := 0; fi17 < 8; fi17++ {
                                if p0 {
                                    if (p0 && p0) {
                                        v18 := []int{p1, 2147483647, (^(fi17))}
                                        _ = v18
                                    }
                                    var v19 int = 9223372036854775807
                                    _ = v19
                                } else {
                                    var v20 float64 = (float64(0.0) + ((1e10 * (((-2.25 * -2.25) + 3.14159) * ((float64(3.14159) / 0.1) - 1.5))) / 3.14159))
                                    v21 := []int{}
                                    v21 = append(v21, 2147483647)
                                    _ = v20
                                    _ = v21
                                }
                                var v22 float64 = 0.0
                                v23 := []int{(^(fi16))}
                                _ = fi17
                                _ = v22
                                _ = v23
                            }
                            var v24 string = v14
                            _ = v24
                        } else {
                            var v25 int8 = 127
                            switch (v14 + "a") {
                            case "a":
                                v26 := GenS1{F0: 9223372036854775807, F1: 255}
                                _ = v26
                            default:
                                v27 := []int{9223372036854775807, 128, (p1 * 923)}
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 14; fi28++ {
                                v29 := GenS1{F0: -909, F1: 2}
                                fmt.Println(v25)
                                var v30 int16 = -32768
                                _ = fi28
                                _ = v29
                                _ = v30
                            }
                            _ = v25
                        }
                        var v31 float64 = 1e10
                        var v32 uint64 = 732
                        _ = fi16
                        _ = v31
                        _ = v32
                    }
                }
                if false {
                    v14 = "\n\t"
                } else {
                    var v33 int16 = 42
                    v34 := []int{(p1 - (p1 % ((p1) | 1))), (p1 | -238), p1}
                    _ = v33
                    _ = v34
                }
            }
        case "\n\t":
            var v35 uint16 = 1
            _ = v35
        }
    case -797:
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v14
    return -1
}

func main() {
    v1 := []string{("a" + "\n\t"), "\n\t", ""}
    v2 := make(map[int]bool)
    var v3 int32 = 0
    v4 := []string{"foo bar"}
    fmt.Println(v3)
    v3 -= (v3 - (v3 >> ((1) & 63)))
    v3 = v3
    if (0 > 255) {
        var v36 float64 = float64(((0.1 + (-2.25 + float32(1e10))) * ((1e10 + 0.0) / 1.5)))
        if len(v4) > 0 {
            v4[0] = (("" + "") + "foo bar")
        }
        _ = v36
    } else {
        var v37 bool = (42 < 255)
        var v38 string = ("a" + "hello")
        _ = v37
        _ = v38
    }
    var v39 uint = 812
    var v40 uint = v39
    v41 := GenS2{F0: 1}
    switch "hello" {
    case "\n\t":
        if true {
            fmt.Println(v39, v39, v40)
            v42 := GenS2{F0: 0}
            v3 = (v3 << ((v40) & 63))
            _ = v42
        } else {
            fmt.Println(v39, v3, v39)
            v3 = int32(v40)
        }
    case "hello":
        var v43 bool = true
        _ = v43
    case "":
        v44 := GenS3{F0: (v3 & (^(v3))), F1: 2147483647, F2: 32767}
        _ = v44
    default:
        v45 := v2[670]
        _ = v45
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
