package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 string; F2 uint64 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 uint; F1 int16; F2 uint32 }
type GenS7 struct { F0 int32; F1 bool; F2 int32 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (!true)
    var v3 uint16 = 2
    var v4 int = 42
    var v5 int64 = 1
    v3 = 909
    delete(v1, 32768)
    if false {
        v6 := GenS1{F0: 121}
        if (2147483647 >= -2147483648) {
            fmt.Println(v3, v3, v2)
        }
        fmt.Println(v2, v3, v3)
        _ = v6
    } else {
        v3 = 2
        delete(v1, 407)
    }
    var v7 uint = 42
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(v3)
        _ = fi8
    }
    var v9 uint8 = 26
    for fi10 := 0; fi10 < 3; fi10++ {
        fmt.Println(v4)
        if true {
            v4 -= 32767
            v3 += v3
            if v2 {
                v11 := GenS1{F0: 46}
                v12 := GenS1{F0: -53}
                _ = v11
                _ = v12
            }
        }
        var v13 int32 = -2147483648
        _ = fi10
        _ = v13
    }
    var v14 uint = v7
    v15 := GenS2{F0: int(v9), F1: ("a" + "")}
    var v16 float64 = 0.0
    fmt.Println(v4)
    if v2 {
        fmt.Println(v4, v16, v2)
        v17 := []int{}
        v2 = (!(!(!false)))
        _ = v17
    }
    if false {
        v7 = (v7 - v14)
        if v2 {
            var v18 uint64 = 0
            _ = v18
        } else {
            v19 := GenS2{F0: -423, F1: ("" + "\n\t")}
            _ = v19
        }
        v5 = (v5 << (((^(v14))) & 63))
    }
    var v20 bool = false
    v7 -= 255
    if ((0 < -100) || v20) {
        v3 = 21
        fmt.Println(v14)
        switch ("foo bar" + "foo bar") {
        case "\n\t":
            switch "" {
            case "a":
                for fi21 := 0; fi21 < 2; fi21++ {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        var v23 float32 = (3.14159 / float32(((v16 / v16) - 1e10)))
                        var v24 int = (fi22 ^ (fi22 << (((^(v7))) & 63)))
                        v14 = (v7 << ((v14) & 63))
                        var v25 int = fi22
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    v1[(fi21 - (v4 % ((int(v9)) | 1)))] = v2
                    for fi26 := 0; fi26 < 19; fi26++ {
                        v7 += (v14 - (^(v14)))
                        _ = fi26
                    }
                    _ = fi21
                }
            case "foo bar":
                fmt.Println(v14, v7)
            case "":
                v14 = (v14 << ((v7) & 63))
            }
        case "foo bar":
            v2 = ((!((v20 && v2) || v20)) || true)
        case "a":
            v27 := GenS3{F0: 1}
            _ = v27
        default:
            v28 := []int{v4}
            _ = v28
        }
    } else {
        if (0 > 487) {
            var v29 uint8 = (v9 | v9)
            v30 := []int{}
            _ = v29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        switch (v4 & fi31) {
        case 721:
            for fi32 := 0; fi32 < 1; fi32++ {
                if true {
                    v4 += fi32
                }
                _ = fi32
            }
        case 9223372036854775807:
            for fi33 := 0; fi33 < 16; fi33++ {
                v5 = v5
                if ((v7 - v7) > v7) {
                    _ = len(v1)
                    switch fi31 {
                    case -303:
                        v14 = (v7 * 1)
                        v34 := GenS4{F0: 0, F1: "a", F2: 1}
                        v35 := v1[(v4 >> ((4294967295) & 63))]
                        _ = v34
                        _ = v35
                    case 32768:
                        v36 := []int{fi31, fi33, (fi31 + 9223372036854775807)}
                        _ = v36
                    default:
                        switch "" {
                        case "":
                            if (!false) {
                                v37 := []int{465, v4, fi33}
                                v7 = v7
                                _ = v37
                            } else {
                                switch "" {
                                case "foo bar":
                                    v2 = ("\n\t" <= "foo bar")
                                default:
                                    v4 = (v4 + fi33)
                                }
                                v38 := GenS5{F0: (false && v2), F1: 471}
                                v39 := GenS5{F0: v2, F1: -55}
                                _ = v38
                                _ = v39
                            }
                        }
                    }
                } else {
                    v9 = 136
                    fmt.Println(v16)
                    var v40 int32 = -703
                    v41 := GenS6{F0: v7, F1: int16(v3), F2: 42}
                    v42 := GenS7{F0: v40, F1: true, F2: v40}
                    v3 += v3
                    _ = v40
                    _ = v41
                    _ = v42
                }
                v16 += v16
                _ = fi33
            }
        }
        fmt.Println(v7, v3)
        v5 -= 9223372036854775807
        _ = fi31
    }
    v2 = (v16 > (v16 * (-2.25 * v16)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
