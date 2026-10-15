package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int; F2 string }
func genFn0(p0 int16, p1 uint, p2 uint16) uint16 {
    for fi6 := 0; fi6 < 14; fi6++ {
        if false {
            p0 -= p0
            p1 += (p1 >> (((^(p1))) & 63))
        }
        _ = fi6
    }
    if false {
        if ((32768 == int(p2)) && (3.14159 > 3.14159)) {
            v7 := []int{-496}
            _ = v7
        }
        var v8 float32 = 1.5
        p2 += p2
        _ = v8
    }
    p0 = p0
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return p2
}
func genFn1(p0 uint8, p1 int32) uint32 {
    var v9 int64 = 9223372036854775807
    v10 := GenS1{F0: "\n\t", F1: 2, F2: 319}
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS1{F0: "\n\t", F1: 588, F2: 42}
        if (32767 == 2) {
            var v13 uint = 830
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    if true {
        for fi14 := 0; fi14 < 15; fi14++ {
            if (("a" + "") == "\n\t") {
                var v15 int = fi14
                _ = v15
            }
            _ = fi14
        }
        for fi16 := 0; fi16 < 2; fi16++ {
            p1 += (p1 >> ((255) & 63))
            _ = fi16
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 775
}

func main() {
    var v1 int8 = 0
    var v2 bool = true
    var v3 string = "\n\t"
    var v4 uint32 = 42
    var v5 int64 = -9223372036854775808
    v1 = 127
    var v17 uint = 366
    var v18 uint32 = v4
    v1 -= v1
    if v2 {
        fmt.Println(v17, v4)
        var v19 int32 = -1
        _ = v19
    } else {
        if true {
            var v20 string = ""
            v5 = v5
            v21 := GenS2{F0: 2147483647}
            _ = v20
            _ = v21
        } else {
            var v22 uint64 = 2
            v23 := []int{128, 127}
            _ = v22
            _ = v23
        }
        var v24 uint32 = (^(v18))
        v25 := []int{-1}
        _ = v24
        _ = v25
    }
    v2 = (199 != 0)
    if v2 {
        var v26 uint32 = genFn1(127, 2147483647)
        v27 := []int{}
        for fi28 := 0; fi28 < 6; fi28++ {
            if v2 {
                _ = len(v27)
                v29 := GenS1{F0: v3, F1: 65535, F2: 553}
                _ = v29
            }
            fmt.Println(v17)
            _ = fi28
        }
        _ = v26
        _ = v27
    }
    switch 0 {
    case 293:
        var v30 uint16 = 0
        _ = v30
    case 595:
        if (v3 > v3) {
            for fi31 := 0; fi31 < 18; fi31++ {
                switch "hello" {
                case "hello":
                    switch "foo bar" {
                    case "hello":
                        v32 := GenS3{F0: ((v18 >> ((v17) & 63)) > (v4 | (v4 | (^(v4)))))}
                        _ = v32
                    default:
                        for fi33 := 0; fi33 < 15; fi33++ {
                            v34 := GenS2{F0: (fi31 - (fi33 >> (((v17 / ((v17) | 1))) & 63)))}
                            var v35 string = "hello"
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                    }
                case "a":
                    v5 = (v5 / (((v5 & v5)) | 1))
                case "\n\t":
                    v36 := GenS4{F0: (^(fi31)), F1: (fi31 & (fi31 & (-(fi31)))), F2: v3}
                    _ = v36
                }
                _ = fi31
            }
            v18 = 198
            v37 := []int{-135, int(v17), -1}
            _ = v37
        } else {
            v17 = v17
            switch 128 {
            case 0:
                v38 := []int{2147483647, 42, 908}
                _ = v38
            case 1:
                var v39 string = v3
                _ = v39
            default:
                for fi40 := 0; fi40 < 9; fi40++ {
                    v41 := GenS1{F0: v3, F1: 937, F2: 0}
                    _ = fi40
                    _ = v41
                }
            }
        }
    case 32767:
        v42 := []int{}
        _ = v42
    default:
        var v43 float32 = ((-2.25 - -2.25) / (((float32(0.1) - 3.14159) - 3.14159) + 1e10))
        _ = v43
    }
    var v44 int8 = -128
    for fi45 := 0; fi45 < 17; fi45++ {
        v2 = true
        switch fi45 {
        case 127:
            if false {
                switch fi45 {
                case 0:
                    if (v2 && v2) {
                        fmt.Println(v44, v17)
                    }
                }
                if v2 {
                    if v2 {
                        var v46 int32 = -1
                        _ = v46
                    }
                    v47 := GenS1{F0: "\n\t", F1: 42, F2: 1}
                    if (v4 > genFn1(128, -922)) {
                        var v48 int = 128
                        _ = v48
                    } else {
                        v17 = v17
                    }
                    _ = v47
                }
                v49 := GenS2{F0: fi45}
                _ = v49
            }
        case 32767:
            v2 = false
        }
        fmt.Println(v17, v5)
        _ = fi45
    }
    v50 := []int{-9223372036854775808, 32768, -991, -9223372036854775808}
    if len(v50) > 0 {
        v50[0] = 9223372036854775807
    }
    var v51 uint16 = 820
    v52 := GenS3{F0: (((1 >= 1) && true) || (float32(float64(3.14159)) == (float32(1.5) + 1e10)))}
    var v53 uint64 = 0
    for fi54 := 0; fi54 < 10; fi54++ {
        v51 += uint16(v18)
        _ = len(v50)
        _ = fi54
    }
    v55 := GenS3{F0: (v4 < genFn1(0, 1))}
    for fi56 := 0; fi56 < 13; fi56++ {
        v18 = 2147483647
        for fi57 := 0; fi57 < 11; fi57++ {
            v4 = (v4 % ((2147483647) | 1))
            _ = fi57
        }
        v58 := GenS2{F0: (^(fi56))}
        _ = fi56
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v55
}
