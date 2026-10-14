package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint; F2 int }
type GenS5 struct { F0 int32; F1 uint32 }
func genFn0(p0 string, p1 uint16, p2 uint8) uint64 {
    var v7 uint16 = p1
    var v8 int8 = -57
    v9 := GenS1{F0: 2147483647, F1: (p2 / ((p2) | 1)), F2: p2}
    if (!false) {
        var v10 uint8 = (p2 << ((18446744073709551615) & 63))
        for fi11 := 0; fi11 < 12; fi11++ {
            p0 = "hello"
            if false {
                for fi12 := 0; fi12 < 7; fi12++ {
                    v13 := []int{(fi11 ^ fi11), (fi11 * (fi11 ^ fi11)), fi11, int(v8)}
                    var v14 int32 = 2147483647
                    v15 := GenS1{F0: 2147483648, F1: 101, F2: v10}
                    v13 = append(v13, fi12)
                    fmt.Println(p2, p1, fi11)
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
                v16 := []int{(fi11 + 127), -9223372036854775808}
                _ = v16
            }
            if true {
                if (false && true) {
                    v17 := []int{127}
                    var v18 int = (fi11 >> ((2) & 63))
                    _ = v17
                    _ = v18
                }
                for fi19 := 0; fi19 < 11; fi19++ {
                    v20 := []int{}
                    p2 = v10
                    _ = fi19
                    _ = v20
                }
            } else {
                if (0.0 < (0.0 + ((1.5 / 1.5) / -2.25))) {
                    v21 := GenS2{F0: 178}
                    var v22 int8 = v8
                    v23 := []int{}
                    _ = v21
                    _ = v22
                    _ = v23
                }
                for fi24 := 0; fi24 < 12; fi24++ {
                    v25 := []int{fi11}
                    v25 = append(v25, 508)
                    _ = fi24
                    _ = v25
                }
            }
            _ = fi11
        }
        fmt.Println(v10, p2, v10)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 float32 = (3.14159 / float32(float64(1e10)))
    var v2 uint64 = 148
    v3 := make(map[int]uint32)
    var v4 bool = (true || (!false))
    v5 := make(map[uint32]uint32)
    v6 := []uint32{481}
    if (0 <= -902) {
        if v4 {
            if true {
                v26 := v5[65535]
                fmt.Println(v2, v1)
                _ = v26
            } else {
                for fi27 := 0; fi27 < 17; fi27++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v1 = v1
                    _ = fi27
                }
                v28 := GenS3{F0: 2, F1: 646}
                var v29 uint32 = 42
                _ = v28
                _ = v29
            }
            v30 := v5[2147483647]
            _ = v30
        } else {
            if len(v6) > 0 {
                v6[0] = 0
            }
        }
        var v31 uint = 377
        _ = v31
    } else {
        v2 += v2
        v32 := GenS4{F0: 306, F1: 42, F2: 9223372036854775807}
        _ = v32
    }
    var v33 int64 = -1
    _ = len(v5)
    for fi34 := 0; fi34 < 16; fi34++ {
        for fi35 := 0; fi35 < 9; fi35++ {
            fmt.Println(v4, v2)
            switch 42 {
            case 780:
                var v36 float64 = (3.14159 + -2.25)
                fmt.Println(v33, v36, fi34)
                _ = v36
            case -1:
                switch int(v2) {
                case 2147483647:
                    if (89 > 255) {
                        var v37 int8 = 127
                        _ = v37
                    }
                default:
                    var v38 string = (("hello" + ("foo bar" + ("\n\t" + "foo bar"))) + "")
                    _ = v38
                }
            case -457:
                var v39 int32 = -1
                _ = v39
            }
            v33 -= -601
            _ = fi35
        }
        _ = fi34
    }
    if false {
        var v40 uint8 = 255
        if true {
            fmt.Println(v1, v2)
            switch 553 {
            case 32768:
                if true {
                    v33 += v33
                    for fi41 := 0; fi41 < 20; fi41++ {
                        v42 := v5[2]
                        var v43 uint8 = 255
                        _ = fi41
                        _ = v42
                        _ = v43
                    }
                    v44 := GenS2{F0: (v2 | (^(v2)))}
                    _ = v44
                } else {
                    for fi45 := 0; fi45 < 18; fi45++ {
                        var v46 uint = 699
                        fmt.Println(v1)
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi45
                        _ = v46
                    }
                }
            case 294:
                v47 := GenS1{F0: 65535, F1: v40, F2: (^(v40))}
                _ = v47
            case 911:
                v4 = (v1 <= float32(-2.25))
            default:
                for fi48 := 0; fi48 < 2; fi48++ {
                    v33 = (v33 << ((uint(v2)) & 63))
                    fmt.Println(v1)
                    _ = fi48
                }
            }
        }
        _ = v40
    } else {
        fmt.Println(v4)
    }
    var v49 string = ("" + ("hello" + "foo bar"))
    v1 = float32((((float64(1.5) / (3.14159 - float64(v1))) / ((-2.25 / ((1.5 * 3.14159) / 1e10)) + 1.5)) / ((((((3.14159 * (0.0 + ((0.1 / (-2.25 / (0.1 + 1e10))) / float64(v1)))) - 1e10) + ((0.1 * 1.5) - 3.14159)) + float64(3.14159)) * ((((3.14159 * (float64(1.5) + (0.0 * 1e10))) * (0.1 + 1e10)) / 0.1) / 1.0)) / float64((1e10 + (float32(0.1) - (float32((((0.0 - 0.1) - 3.14159) * -2.25)) / -2.25)))))))
    v50 := GenS5{F0: 0, F1: 333}
    var v51 int64 = (v33 & (v33 << ((7) & 63)))
    var v52 uint = 255
    delete(v3, 4294967295)
    var v53 int = 2147483647
    var v54 uint = 888
    v55 := GenS2{F0: 2}
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
