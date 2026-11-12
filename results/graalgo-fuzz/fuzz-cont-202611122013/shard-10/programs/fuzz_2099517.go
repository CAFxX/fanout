package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int16; F1 int; F2 uint8 }
type GenS3 struct { F0 uint16; F1 bool; F2 int8 }
type GenS4 struct { F0 uint16; F1 int64; F2 int64 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 int32; F1 int32; F2 int16 }
type GenS8 struct { F0 int16; F1 int32 }
type GenS9 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint64, p1 int64) uint16 {
    v6 := []int{32768, 635, 816}
    var v7 int32 = int32(p1)
    var v8 uint8 = 255
    if true {
        for fi9 := 0; fi9 < 4; fi9++ {
            if (!(0 <= 42)) {
                var v10 float32 = 3.14159
                fmt.Println(v10, v8, v7)
                p1 = (p1 >> ((uint(p0)) & 63))
                _ = v10
            } else {
                v11 := GenS1{F0: v8, F1: 32767}
                _ = v11
            }
            var v12 float64 = 0.1
            v13 := GenS2{F0: 42, F1: (-(fi9)), F2: (^(v8))}
            _ = fi9
            _ = v12
            _ = v13
        }
    } else {
        if ((-9223372036854775808 >= p1) || true) {
            switch 127 {
            case -9223372036854775808:
                fmt.Println(v7, v8)
            case 9223372036854775807:
                v14 := GenS3{F0: 481, F1: (-888 <= 1000), F2: 76}
                _ = v14
            case 42:
                switch int(p0) {
                case 897:
                    var v15 uint16 = 32767
                    _ = v15
                case 682:
                    v16 := GenS4{F0: 1, F1: (p1 | p1), F2: (p1 ^ 42)}
                    _ = v16
                }
            }
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            fmt.Println(fi17)
            _ = fi17
        }
    }
    if (false && (true && true)) {
        for fi18 := 0; fi18 < 8; fi18++ {
            if (v7 != v7) {
                var v19 uint = 1
                _ = v19
            }
            for fi20 := 0; fi20 < 18; fi20++ {
                switch ("foo bar" + "hello") {
                case "a":
                    if false {
                        if false {
                            p0 += (p0 | p0)
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            p0 += (p0 % ((42) | 1))
                        }
                    } else {
                        var v21 int32 = -304
                        _ = v21
                    }
                case "\n\t":
                    v8 += v8
                case "":
                    v22 := GenS5{F0: ((3.14159 * (((((-2.25 / 1.0) / (float64(0.0) - (0.1 - -2.25))) + -2.25) - 1e10) * 3.14159)) < 3.14159), F1: "hello"}
                    _ = v22
                }
                fmt.Println(fi20, v8)
                switch fi20 {
                case 128:
                    v23 := GenS6{F0: fi18, F1: 4294967295}
                    _ = v23
                case -84:
                    v7 -= v7
                case 9223372036854775807:
                    if true {
                        if true {
                            var v24 uint32 = 65535
                            p1 += 42
                            v25 := GenS3{F0: 2, F1: true, F2: -1}
                            _ = v24
                            _ = v25
                        }
                        var v26 bool = true
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = v26
                    }
                }
                _ = fi20
            }
            var v27 string = "hello"
            _ = fi18
            _ = v27
        }
        fmt.Println(v8)
        v8 -= 42
    } else {
        fmt.Println(v8, p1)
        p1 = (p1 | p1)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint64, p1 int16) uint16 {
    var v28 int32 = -2147483648
    var v29 float32 = 0.1
    v30 := []int{127, 32768, 0, 2147483647}
    fmt.Println(p0, v29, v28)
    _ = p0
    _ = p1
    _ = v28
    _ = v29
    _ = v30
    return 10
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    var v4 uint16 = 939
    var v5 int = 1
    v4 = (^(v4))
    for fi31 := 0; fi31 < 3; fi31++ {
        fmt.Println(v4)
        for fi32 := 0; fi32 < 4; fi32++ {
            var v33 int64 = 0
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    var v34 float64 = 3.14159
    switch "" {
    case "a":
        v34 = 3.14159
    }
    var v35 float32 = (float32((float64((-2.25 - ((0.1 + 1.5) + float32(v34)))) * v34)) / 1.5)
    var v36 int32 = 353
    for fi37 := 0; fi37 < 4; fi37++ {
        v38 := []int{}
        if true {
            v39 := GenS4{F0: 42, F1: 42, F2: 9223372036854775807}
            _ = v39
        } else {
            var v40 int32 = v36
            fmt.Println(v40, v40, v4)
            _ = v40
        }
        _ = fi37
        _ = v38
    }
    v41 := GenS7{F0: 0, F1: v36, F2: 0}
    var v42 uint64 = 192
    v36 += v36
    for fi43 := 0; fi43 < 11; fi43++ {
        v44 := GenS6{F0: -9223372036854775808, F1: 255}
        _ = fi43
        _ = v44
    }
    v45 := GenS8{F0: 546, F1: (v36 / (((v36 / ((v36) | 1))) | 1))}
    switch v5 {
    case 42:
        for fi46 := 0; fi46 < 9; fi46++ {
            v47 := GenS9{F0: 272, F1: 417}
            var v48 int16 = 42
            v48 = (v48 & v48)
            _ = fi46
            _ = v47
            _ = v48
        }
    case 271:
        switch ("\n\t" + "\n\t") {
        case "hello":
            var v49 int8 = 127
            _ = v49
        case "":
            v50 := []int{-975, 9223372036854775807, 465}
            _ = v50
        default:
            if true {
                switch 127 {
                case -558:
                    v51 := []int{(-(v5)), 128, 32768}
                    _ = v51
                case -769:
                    v52 := []int{}
                    _ = v52
                }
                var v53 uint32 = 2147483648
                _ = v53
            }
        }
    case 824:
        for fi54 := 0; fi54 < 3; fi54++ {
            switch (^(fi54)) {
            case 42:
                v34 = ((v34 * (((float64((float32(v34) / 3.14159)) * (0.0 / v34)) + 1e10) / v34)) / v34)
            }
            _ = fi54
        }
    default:
        v55 := GenS7{F0: 711, F1: v36, F2: 0}
        _ = v55
    }
    for fi56 := 0; fi56 < 11; fi56++ {
        for fi57 := 0; fi57 < 6; fi57++ {
            var v58 string = "hello"
            _ = fi57
            _ = v58
        }
        _ = fi56
    }
    v4 = (v4 % (((v4 | (v4 / (((^(v4))) | 1)))) | 1))
    fmt.Println(v5, v4)
    v34 += 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v5
}
