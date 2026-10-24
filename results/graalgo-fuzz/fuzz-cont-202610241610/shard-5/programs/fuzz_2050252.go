package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint32 }
type GenS2 struct { F0 int; F1 uint64; F2 uint32 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int8; F1 int8 }
func genFn0(p0 int8) int8 {
    v6 := GenS1{F0: 2147483647, F1: 42, F2: 4294967295}
    p0 += (p0 >> ((42) & 63))
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int32 = -1
        v9 := []int{(fi7 & -3)}
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{2147483647, -9223372036854775808, 32768}
    _ = p0
    _ = v10
    _ = v6
    return (p0 * (p0 & p0))
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[uint32]bool)
    v3 := []uint32{}
    var v4 float32 = (0.0 + (0.1 - 0.1))
    var v5 string = "hello"
    if true {
        v1 += (v1 | v1)
        v5 = ((v5 + "a") + (v5 + v5))
        v1 = 2147483648
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        switch (("\n\t" + v5) + ("" + "a")) {
        case "":
            fmt.Println(v5)
        }
        switch ("a" + v5) {
        case "foo bar":
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 int8 = 0
                _ = fi12
                _ = v13
            }
        default:
            switch ((v5 + v5) + (v5 + v5)) {
            case "foo bar":
                fmt.Println(v5, v1)
            case "\n\t":
                if (!(false && true)) {
                    for fi14 := 0; fi14 < 2; fi14++ {
                        v4 -= ((((((v4 + (v4 / (v4 / (v4 / v4)))) * v4) + float32(float64(v4))) + ((v4 - (-2.25 * (v4 * v4))) + v4)) + 3.14159) + v4)
                        switch v5 {
                        case "a":
                            v1 += (v1 << ((65535) & 63))
                        case "foo bar":
                            switch fi11 {
                            case -796:
                                if (!true) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    fmt.Println(fi11)
                                }
                            case -1:
                                if true {
                                    fmt.Println(v4)
                                }
                            case 472:
                                var v15 float64 = float64((v4 - (float32(1e10) * (v4 - v4))))
                                fmt.Println(v4)
                                v2[(v1 | v1)] = true
                                _ = v15
                            }
                        case "\n\t":
                            var v16 bool = (!false)
                            _ = v16
                        }
                        _ = fi14
                    }
                    if (!(true && (-32768 != -1))) {
                        v17 := GenS2{F0: (fi11 ^ (fi11 - -437)), F1: 42, F2: (v1 * (v1 - (v1 | 615)))}
                        _ = v17
                    }
                } else {
                    switch fi11 {
                    case 514:
                        var v18 uint64 = 18446744073709551615
                        _ = v18
                    case 481:
                        v1 = v1
                    case -196:
                        var v19 uint = 65535
                        _ = v19
                    }
                    fmt.Println(fi11, v4)
                }
            }
        }
        var v20 uint8 = 130
        _ = fi11
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        if true {
            v4 = 0.0
            for fi22 := 0; fi22 < 5; fi22++ {
                if true {
                    switch (fi22 + (fi22 >> ((42) & 63))) {
                    case 42:
                        fmt.Println(fi21)
                    case 230:
                        v23 := GenS3{F0: "foo bar", F1: v1}
                        _ = v23
                    default:
                        v4 = (v4 - v4)
                        fmt.Println(fi22, v1, v1)
                        v4 = -2.25
                        v1 = v1
                        switch (fi22 | fi21) {
                        case 2147483647:
                            if false {
                                switch fi21 {
                                case 42:
                                    v1 = (^(v1))
                                case 127:
                                    var v24 string = (v5 + v5)
                                    _ = v24
                                default:
                                    _ = len(v3)
                                }
                                var v25 int8 = -13
                                _ = v25
                            } else {
                                v26 := GenS4{F0: 2, F1: 1}
                                var v27 uint16 = 2
                                _ = v26
                                _ = v27
                            }
                        case -9223372036854775808:
                            _ = len(v3)
                        }
                    }
                    fmt.Println(fi21, fi21, fi22)
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v1 -= 154
                _ = fi22
            }
            v4 = v4
        }
        _ = fi21
    }
    var v28 uint8 = 156
    var v29 int = 2147483647
    for fi30 := 0; fi30 < 6; fi30++ {
        if ((-1 != -1) || (true && (255 == 960))) {
            if ((genFn0(127) <= 1) && true) {
                var v31 int16 = -1
                var v32 int = fi30
                if true {
                    fmt.Println(v32, v4)
                }
                _ = v31
                _ = v32
            } else {
                v33 := GenS1{F0: -1, F1: -1, F2: v1}
                _ = v33
            }
        }
        v29 += fi30
        _ = fi30
    }
    fmt.Println(v4, v29)
    delete(v2, (v1 >> ((uint(v28)) & 63)))
    switch v29 {
    case 2147483647:
        if false {
            v4 += v4
            for fi34 := 0; fi34 < 19; fi34++ {
                v1 = (v1 & v1)
                var v35 int = fi34
                _ = fi34
                _ = v35
            }
            v36 := v2[v1]
            _ = v36
        }
    case 0:
        fmt.Println(v5, v5)
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        if false {
            var v37 float32 = v4
            _ = v37
        } else {
            for fi38 := 0; fi38 < 17; fi38++ {
                switch v5 {
                case "hello":
                    var v39 uint32 = v1
                    _ = v39
                default:
                    v40 := GenS2{F0: fi38, F1: 18446744073709551615, F2: v1}
                    _ = v40
                }
                v29 -= fi38
                _ = fi38
            }
            switch v29 {
            case 733:
                var v41 int32 = -1
                _ = v41
            case 0:
                var v42 int32 = -2147483648
                _ = v42
            default:
                if len(v3) > 0 {
                    v3[0] = v1
                }
            }
        }
        if (false || (v1 <= (v1 ^ (v1 / ((34) | 1))))) {
            v43 := GenS1{F0: 0, F1: 42, F2: (v1 << ((380) & 63))}
            v44 := GenS5{F0: -1, F1: 1}
            switch "hello" {
            case "hello":
                var v45 string = v5
                _ = v45
            }
            _ = v43
            _ = v44
        }
        var v46 bool = (uint8(v1) == (v28 << ((0) & 63)))
        _ = v46
    }
    v2[v1] = (true || (((^(v1)) == (v1 << ((976) & 63))) || true))
    _ = len(v2)
    v4 = 3.14159
    var v47 uint8 = v28
    for fi48 := 0; fi48 < 5; fi48++ {
        for fi49 := 0; fi49 < 11; fi49++ {
            var v50 int64 = -50
            v50 += 9223372036854775807
            _ = fi49
            _ = v50
        }
        v1 -= (v1 >> ((2) & 63))
        var v51 int32 = -63
        _ = fi48
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
