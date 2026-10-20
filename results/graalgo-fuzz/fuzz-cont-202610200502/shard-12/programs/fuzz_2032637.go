package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int64 }
func genFn0(p0 bool) int16 {
    var v7 float64 = (3.14159 / -2.25)
    v8 := []int{9223372036854775807}
    switch -101 {
    case -409:
        if true {
            var v9 int8 = 1
            var v10 float32 = (((((1e10 * 1e10) / 1.5) + 1.5) + 3.14159) / 1.0)
            _ = len(v8)
            _ = v10
            _ = v9
        } else {
            if p0 {
                p0 = p0
                switch ((("hello" + "") + ("foo bar" + ("hello" + "foo bar"))) + "a") {
                case "a":
                    var v11 uint64 = 18446744073709551615
                    _ = v11
                }
                var v12 int32 = 0
                _ = v12
            } else {
                switch 1 {
                case 32768:
                    var v13 uint32 = 2
                    _ = v13
                case 660:
                    switch 32768 {
                    case 127:
                        var v14 string = "hello"
                        _ = v14
                    case 128:
                        v8 = append(v8, 0)
                    }
                default:
                    for fi15 := 0; fi15 < 19; fi15++ {
                        var v16 int32 = 1
                        if (581 >= 1) {
                            switch fi15 {
                            case -2147483648:
                                var v17 uint64 = 18446744073709551615
                                _ = v17
                            case -9223372036854775808:
                                var v18 bool = false
                                _ = v18
                            default:
                                p0 = false
                            }
                            var v19 int32 = v16
                            var v20 int = fi15
                            _ = v19
                            _ = v20
                        }
                        _ = fi15
                        _ = v16
                    }
                }
            }
            v7 -= v7
        }
    }
    if true {
        _ = len(v8)
    }
    _ = p0
    _ = v7
    _ = v8
    return 1000
}
func genFn1(p0 string, p1 uint16, p2 uint16) int {
    v21 := GenS1{F0: 35}
    for fi22 := 0; fi22 < 17; fi22++ {
        for fi23 := 0; fi23 < 2; fi23++ {
            fmt.Println(fi22, p0)
            _ = fi23
        }
        _ = fi22
    }
    switch 9223372036854775807 {
    case 32768:
        fmt.Println(p2)
    case -665:
        var v24 uint = 920
        _ = v24
    case 1:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 4294967295
}

func main() {
    var v1 uint16 = 1
    var v2 float64 = ((0.0 + (-2.25 / -2.25)) + 0.1)
    var v3 uint16 = (^(v1))
    v4 := make(map[int]uint32)
    var v5 int16 = 986
    var v6 uint8 = 74
    var v25 int32 = -1
    switch (("" + ("" + "hello")) + ("" + ("hello" + "\n\t"))) {
    case "foo bar":
        v1 += v3
    }
    if (true && ((127 <= 82) || (false || false))) {
        fmt.Println(v1, v5, v3)
    } else {
        v1 -= (^(v3))
        var v26 uint = 42
        _ = v26
    }
    v3 += (^(v1))
    var v27 float64 = (v2 * 0.0)
    v3 -= v3
    switch 1 {
    case 9223372036854775807:
        v4[2147483647] = 4294967295
    case 186:
        v2 = v2
    case -137:
        switch -961 {
        case 790:
            v28 := []int{}
            _ = v28
        case -230:
            v6 -= (^(v6))
        default:
            var v29 int8 = -30
            _ = v29
        }
    }
    v30 := GenS1{F0: (v6 >> ((769) & 63))}
    if true {
        v3 = (v3 & (v3 | 32767))
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        v32 := []int{fi31, (fi31 >> ((65535) & 63)), (^(fi31)), fi31}
        if len(v32) >= 2 {
            v32 = v32[1:len(v32)]
        }
        _ = fi31
        _ = v32
    }
    if (false || false) {
        fmt.Println(v5, v3, v25)
        v6 = v6
        v33 := GenS2{F0: 4294967295, F1: 636}
        _ = v33
    } else {
        for fi34 := 0; fi34 < 12; fi34++ {
            v35 := []int{}
            var v36 int8 = -1
            switch (fi34 << ((462) & 63)) {
            case 455:
                switch (fi34 + (fi34 * fi34)) {
                case 42:
                    v37 := v4[fi34]
                    _ = v37
                case 32767:
                    var v38 uint16 = (v1 << ((90) & 63))
                    _ = v38
                default:
                    if false {
                        switch 32767 {
                        case 32767:
                            v2 = v27
                        case 0:
                            if (18446744073709551615 == 695) {
                                v39 := GenS3{F0: 2, F1: (v36 - (v36 & (^(v36)))), F2: 236}
                                _ = len(v4)
                                _ = v39
                            }
                        }
                    }
                }
            case 2147483647:
                var v40 bool = (18446744073709551615 != 42)
                _ = v40
            case 32767:
                if false {
                    if ((v25 * v25) >= v25) {
                        var v41 uint16 = 32767
                        _ = v41
                    }
                    if false {
                        var v42 uint16 = v1
                        _ = v42
                    }
                }
            default:
                v43 := GenS4{F0: (v6 | 128), F1: 0}
                _ = v43
            }
            _ = fi34
            _ = v35
            _ = v36
        }
        for fi44 := 0; fi44 < 5; fi44++ {
            if (4294967295 > 358) {
                v6 = v6
                v45 := GenS2{F0: 961, F1: 39}
                _ = v45
            }
            v46 := []int{fi44}
            v47 := GenS1{F0: v6}
            _ = fi44
            _ = v46
            _ = v47
        }
    }
    fmt.Println(v6)
    switch "a" {
    case "foo bar":
        var v48 bool = (!(uint64(v6) == 547))
        _ = v48
    case "hello":
        if true {
            var v49 uint16 = 255
            var v50 uint = 255
            v50 += (v50 | (^(v50)))
            _ = v49
            _ = v50
        } else {
            if false {
                v51 := v4[genFn1("hello", 255, 0)]
                if ("\n\t" >= "a") {
                    v52 := []int{2147483647, 0}
                    fmt.Println(v6)
                    _ = v52
                }
                _ = v51
            }
            if true {
                switch ("foo bar" + "\n\t") {
                case "\n\t":
                    v53 := GenS2{F0: 4294967295, F1: 1}
                    _ = v53
                default:
                    if false {
                        fmt.Println(v25)
                    }
                }
                _ = len(v4)
            }
            v27 = (((3.14159 / (float64(float32((float64(0.1) + ((1e10 * 1e10) - float64(1e10))))) / (((v2 * v2) * (3.14159 - 3.14159)) * ((((v2 - 1e10) - v2) - v2) / ((0.0 * float64(1.5)) - (1e10 - v2)))))) + ((((v2 - v27) - v2) + float64(0.1)) + (float64((3.14159 - float32((float64(0.0) - ((0.1 / 1.5) + v27))))) * float64((-2.25 * (0.0 / 0.1)))))) / ((((v2 - float64((3.14159 * ((((((1e10 - -2.25) + ((((3.14159 + 3.14159) - float32(v27)) - 1e10) + 1e10)) * (-2.25 / float32(3.14159))) + -2.25) - 3.14159) + (3.14159 + (3.14159 + 1.5)))))) + v27) * v2) - v27))
        }
    default:
        v2 -= v27
    }
    v27 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
