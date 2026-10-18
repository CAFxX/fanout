package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0() uint {
    var v7 uint8 = 36
    var v8 int = 778
    var v9 int8 = -128
    v10 := []int{(v8 % (((^(v8))) | 1)), -1}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 204
}
func genFn1(p0 uint32) uint8 {
    switch "" {
    case "":
        if false {
            v11 := []int{-2147483648}
            _ = v11
        }
    case "foo bar":
        p0 = (p0 - (p0 >> ((42) & 63)))
    case "\n\t":
        switch 32767 {
        case 42:
            if false {
                v12 := []int{1, 2147483647, 966, 1}
                if len(v12) > 0 {
                    v12[0] = -2147483648
                }
                _ = v12
            } else {
                if (true || true) {
                    p0 = p0
                }
            }
        case 0:
            fmt.Println(p0)
        }
    default:
        for fi13 := 0; fi13 < 4; fi13++ {
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 uint32 = (p0 / ((2) | 1))
                _ = fi14
                _ = v15
            }
            var v16 string = "a"
            for fi17 := 0; fi17 < 19; fi17++ {
                v16 = ""
                v18 := GenS1{F0: -9223372036854775808, F1: "hello"}
                var v19 float32 = -2.25
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi13
            _ = v16
        }
    }
    switch "hello" {
    case "hello":
        fmt.Println(p0)
    case "foo bar":
        v20 := GenS1{F0: -1, F1: ("a" + "")}
        _ = v20
    case "":
        fmt.Println(p0)
    default:
        if (42 == 32767) {
            p0 = p0
        }
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 string = "\n\t"
    v3 := make(map[uint32]int)
    var v4 int8 = -5
    v5 := make(map[int]bool)
    var v6 uint = 255
    if (true && true) {
        delete(v5, 772)
        switch -2147483648 {
        case -1:
            var v21 uint = v6
            _ = v21
        default:
            v22 := v3[42]
            _ = v22
        }
    }
    if ((-686 > 0) || (0 >= 4294967295)) {
        switch v2 {
        case "":
            v4 -= v4
        }
        var v23 bool = (!true)
        for fi24 := 0; fi24 < 4; fi24++ {
            v25 := []int{int(v6), (fi24 + 32767), fi24, fi24}
            _ = len(v25)
            v2 = v2
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    switch v2 {
    case "a":
        if false {
            var v26 string = v2
            var v27 uint8 = genFn1(83)
            var v28 uint64 = 1
            _ = v26
            _ = v27
            _ = v28
        }
    case "\n\t":
        v29 := []int{1, 9223372036854775807}
        _ = v29
    case "":
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 uint64 = 878
            _ = fi30
            _ = v31
        }
    }
    var v32 int32 = 1
    v2 = v2
    var v33 int32 = (v32 + -383)
    v34 := GenS2{F0: 1}
    v35 := GenS1{F0: 977, F1: "foo bar"}
    fmt.Println(v2)
    _ = len(v3)
    switch 0 {
    case -2147483648:
        var v36 int = -2147483648
        _ = v36
    case 4294967295:
        if ((v4 - -1) >= v4) {
            for fi37 := 0; fi37 < 1; fi37++ {
                v38 := []int{(fi37 | 4294967295), (fi37 * (fi37 - (-(fi37))))}
                _ = fi37
                _ = v38
            }
        } else {
            for fi39 := 0; fi39 < 13; fi39++ {
                var v40 uint8 = 42
                if true {
                    v4 = (v4 & (-(v4)))
                } else {
                    if ((v2 + v2) >= "a") {
                        for fi41 := 0; fi41 < 15; fi41++ {
                            v32 = v33
                            v1[v2] = (^(fi41))
                            _ = fi41
                        }
                        if (!false) {
                            v42 := GenS2{F0: 290}
                            var v43 bool = true
                            var v44 float32 = 0.0
                            _ = v42
                            _ = v43
                            _ = v44
                        } else {
                            fmt.Println(v40)
                            v45 := []int{(fi39 + fi39)}
                            v46 := GenS3{F0: (v6 - (^(v6))), F1: 0}
                            _ = v45
                            _ = v46
                        }
                        fmt.Println(v2, v32)
                    }
                }
                v6 = v6
                _ = fi39
                _ = v40
            }
        }
    }
    if false {
        v4 += -19
        var v47 float32 = (0.0 * ((3.14159 / ((float32(0.0) / (0.0 + (3.14159 * float32(((float64(3.14159) / (float64(3.14159) / 1.0)) + (float64(float32((-2.25 / 3.14159))) + -2.25)))))) + (1e10 * 3.14159))) / 1e10))
        fmt.Println(v6)
        _ = v47
    } else {
        for fi48 := 0; fi48 < 1; fi48++ {
            fmt.Println(v4, v6, v32)
            _ = fi48
        }
        v49 := GenS1{F0: 727, F1: v2}
        _ = v49
    }
    v2 = "hello"
    v2 = "hello"
    v32 += int32(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
