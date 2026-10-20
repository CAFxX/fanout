package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 int16; F2 uint }
func genFn0(p0 bool, p1 int64, p2 uint16) int8 {
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := []int{}
            if p0 {
                v9 := GenS1{F0: -2147483648, F1: 0}
                _ = v9
            } else {
                var v10 int16 = 32767
                switch "hello" {
                case "":
                    p1 = 1
                case "a":
                    if (p1 >= p1) {
                        v11 := GenS2{F0: 771, F1: (p1 >> ((65535) & 63))}
                        v8 = append(v8, fi6)
                        _ = v11
                    }
                }
                switch fi6 {
                case -1:
                    if false {
                        _ = len(v8)
                        v12 := GenS1{F0: -2147483648, F1: 21}
                        _ = v12
                    }
                case 982:
                    var v13 int8 = -1
                    _ = v13
                default:
                    p0 = (-1 <= 42)
                }
                _ = v10
            }
            var v14 uint8 = 1
            _ = fi7
            _ = v14
            _ = v8
        }
        _ = fi6
    }
    switch -2147483648 {
    case 32767:
        if p0 {
            p1 += 1
        } else {
            p2 = (p2 + (p2 & (p2 << ((255) & 63))))
        }
    case 2147483647:
        for fi15 := 0; fi15 < 17; fi15++ {
            if (p0 && (p0 || (!(-1 == 1000)))) {
                p1 -= p1
                v16 := []int{fi15, fi15}
                var v17 bool = p0
                _ = v16
                _ = v17
            }
            _ = fi15
        }
    case 4294967295:
        var v18 uint32 = 0
        _ = v18
    }
    p2 += (p2 % (((p2 + p2)) | 1))
    for fi19 := 0; fi19 < 6; fi19++ {
        fmt.Println(fi19, p2, p0)
        _ = fi19
    }
    fmt.Println(p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return -79
}
func genFn1(p0 int32, p1 int, p2 int8) uint32 {
    switch "foo bar" {
    case "":
        for fi20 := 0; fi20 < 7; fi20++ {
            if (false && (-647 != 42)) {
                var v21 uint16 = 812
                _ = v21
            } else {
                if false {
                    if (false && (3.14159 < 1e10)) {
                        if true {
                            fmt.Println(fi20, p1)
                            v22 := []int{}
                            switch fi20 {
                            case 175:
                                for fi23 := 0; fi23 < 1; fi23++ {
                                    fmt.Println(fi23, fi23)
                                    _ = fi23
                                }
                            default:
                                v22 = append(v22, (fi20 % ((-275) | 1)))
                            }
                            _ = v22
                        }
                    } else {
                        var v24 int32 = p0
                        v25 := []int{32768, -626, (fi20 & int(v24)), fi20}
                        var v26 int = (p1 << ((1) & 63))
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    fmt.Println(p0)
                } else {
                    fmt.Println(fi20, fi20)
                    switch "\n\t" {
                    case "foo bar":
                        v27 := []int{-2147483648}
                        _ = v27
                    case "a":
                        var v28 int = 32768
                        _ = v28
                    case "\n\t":
                        var v29 uint16 = 434
                        _ = v29
                    default:
                        v30 := []int{fi20, fi20, 42}
                        _ = v30
                    }
                }
                var v31 uint32 = 42
                _ = v31
            }
            _ = fi20
        }
    case "foo bar":
        switch p1 {
        case 488:
            p0 += 0
        default:
            for fi32 := 0; fi32 < 1; fi32++ {
                p2 = (p2 >> ((1) & 63))
                _ = fi32
            }
        }
    case "hello":
        v33 := GenS3{F0: 726}
        _ = v33
    default:
        p0 += p0
    }
    var v34 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    return 60
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 537
    v3 := make(map[int]int)
    var v4 int16 = 1
    var v5 float32 = (1e10 - 1.5)
    v35 := GenS4{F0: "foo bar", F1: 130}
    v5 += v5
    if (!(2 > 44)) {
        v36 := GenS3{F0: 1}
        switch 32767 {
        case 0:
            v37 := GenS1{F0: 816, F1: 18446744073709551615}
            _ = v37
        case 586:
            v2 += genFn1(42, 32768, 127)
        }
        v38 := GenS5{F0: 65535}
        _ = v36
        _ = v38
    } else {
        v5 = (v5 / (((v5 + v5) * float32(0.0)) * 1.5))
    }
    v39 := GenS6{F0: 2, F1: (-(v4)), F2: 65535}
    var v40 uint = 2
    if false {
        var v41 bool = true
        _ = v41
    }
    var v42 float32 = (float32((0.0 - 1.5)) * v5)
    fmt.Println(v42, v42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v42
    _ = v5
}
