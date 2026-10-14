package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 string }
type GenS2 struct { F0 int32; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint16; F2 int }
func genFn0() string {
    var v4 uint16 = 84
    var v5 string = "hello"
    if (!true) {
        switch -268 {
        case 2147483647:
            var v6 uint16 = (^(v4))
            _ = v6
        default:
            var v7 float32 = ((0.0 - (0.1 / 1e10)) + 3.14159)
            _ = v7
        }
        v4 += (v4 % (((v4 ^ 2)) | 1))
    } else {
        v8 := []int{-659}
        if (97 <= 86) {
            var v9 bool = ((-403 > 42) || false)
            _ = v9
        } else {
            v8 = append(v8, 127)
            var v10 uint = 42
            _ = v10
        }
        _ = v8
    }
    var v11 bool = (32767 > -32768)
    v12 := GenS1{F0: v4, F1: 610, F2: "hello"}
    _ = v11
    _ = v12
    _ = v4
    _ = v5
    return v5
}
func genFn1() string {
    v13 := []int{}
    v13 = append(v13, 128)
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 uint16 = 0
        v15 += 672
        v15 -= v15
        _ = fi14
        _ = v15
    }
    _ = v13
    return ""
}

func main() {
    var v1 int8 = -55
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v16 uint8 = 1
    fmt.Println(v1, v1)
    v17 := GenS2{F0: -896, F1: 255, F2: 32767}
    if (!(-32768 < 1)) {
        for fi18 := 0; fi18 < 9; fi18++ {
            v19 := []int{(fi18 >> ((1) & 63)), fi18, (-(fi18))}
            v16 = v16
            _ = fi18
            _ = v19
        }
        var v20 uint8 = v16
        var v21 uint64 = 272
        _ = v20
        _ = v21
    }
    v2[2147483648] = -1
    switch -2147483648 {
    case -2147483648:
        v22 := []int{359, 715, 65}
        _ = v22
    case 32767:
        if false {
            for fi23 := 0; fi23 < 20; fi23++ {
                if (false && false) {
                    var v24 uint16 = 42
                    if (!(0 != 255)) {
                        v25 := []int{(fi23 % ((42) | 1)), fi23}
                        v16 += v16
                        v26 := GenS2{F0: 2147483647, F1: 18446744073709551615, F2: 255}
                        _ = v25
                        _ = v26
                    } else {
                        v27 := []int{fi23}
                        _ = v27
                    }
                    _ = v24
                } else {
                    for fi28 := 0; fi28 < 14; fi28++ {
                        switch fi23 {
                        case -1:
                            switch ("foo bar" + ("foo bar" + (("\n\t" + "a") + ("" + ("foo bar" + "\n\t"))))) {
                            case "foo bar":
                                if (!false) {
                                    v29 := []int{}
                                    _ = len(v3)
                                    v1 -= v1
                                    var v30 int = fi28
                                    _ = v29
                                    _ = v30
                                } else {
                                    v31 := GenS3{F0: 42}
                                    _ = v31
                                }
                            case "hello":
                                fmt.Println(fi23, fi28)
                            }
                        default:
                            v32 := GenS1{F0: 518, F1: uint32(v16), F2: (("" + "\n\t") + "foo bar")}
                            _ = v32
                        }
                        fmt.Println(fi28, v16, v1)
                        _ = fi28
                    }
                    v33 := GenS1{F0: 42, F1: 2, F2: "hello"}
                    switch 32768 {
                    case 4294967295:
                        var v34 int64 = 9223372036854775807
                        _ = v34
                    default:
                        var v35 string = "hello"
                        _ = v35
                    }
                    _ = v33
                }
                v36 := []int{(fi23 | -9223372036854775808), fi23, fi23}
                _ = fi23
                _ = v36
            }
            var v37 string = ("" + ((("\n\t" + "\n\t") + "hello") + "hello"))
            v38 := GenS4{F0: false, F1: 206, F2: 128}
            _ = v37
            _ = v38
        } else {
            fmt.Println(v16)
        }
    default:
        v1 = (^(v1))
    }
    fmt.Println(v16, v1)
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
