package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint; F1 int; F2 string }
type GenS6 struct { F0 uint16; F1 uint; F2 int8 }
type GenS7 struct { F0 uint; F1 bool; F2 int8 }
func genFn0(p0 int64) uint64 {
    v7 := GenS1{F0: 42}
    if true {
        switch "" {
        case "hello":
            switch 76 {
            case 128:
                v8 := []int{-1, -1, -812, 9223372036854775807}
                _ = v8
            }
        case "\n\t":
            if true {
                fmt.Println(p0)
                for fi9 := 0; fi9 < 3; fi9++ {
                    for fi10 := 0; fi10 < 12; fi10++ {
                        v11 := GenS1{F0: 32767}
                        _ = fi10
                        _ = v11
                    }
                    var v12 float64 = 3.14159
                    for fi13 := 0; fi13 < 7; fi13++ {
                        var v14 int32 = -879
                        v15 := GenS1{F0: 32767}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    _ = fi9
                    _ = v12
                }
            } else {
                var v16 float32 = ((0.0 * (1.5 + 1e10)) - 1e10)
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 uint8 = 127
            v19 := []int{fi17}
            switch ("\n\t" + "") {
            case "":
                v20 := GenS1{F0: 72}
                _ = v20
            case "foo bar":
                switch "" {
                case "a":
                    v21 := GenS2{F0: 32767}
                    _ = v21
                case "foo bar":
                    switch "a" {
                    case "a":
                        v22 := GenS3{F0: 42, F1: (((("hello" + ("foo bar" + (("foo bar" + ("foo bar" + "a")) + "foo bar"))) + "foo bar") + "\n\t") + ("hello" + "\n\t"))}
                        _ = v22
                    case "":
                        var v23 int32 = -2147483648
                        _ = v23
                    case "hello":
                        v24 := GenS4{F0: 42, F1: 42}
                        _ = v24
                    default:
                        var v25 bool = (p0 != 9223372036854775807)
                        _ = v25
                    }
                default:
                    fmt.Println(p0, p0)
                }
            case "a":
                v19 = append(v19, fi17)
            default:
                var v26 uint32 = 2147483648
                _ = v26
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        var v27 uint64 = 4294967295
        _ = v27
    } else {
        v28 := []int{-1, 503, -820}
        _ = v28
    }
    v29 := []int{-9223372036854775808, 421, -390}
    _ = p0
    _ = v29
    _ = v7
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -904
    var v3 uint32 = 0
    v4 := make(map[int]string)
    var v5 int64 = 42
    v6 := make(map[int]uint32)
    v3 = v3
    var v30 uint8 = 127
    for fi31 := 0; fi31 < 4; fi31++ {
        if (3.14159 != 1e10) {
            var v32 float64 = (3.14159 * float64(1e10))
            _ = v32
        } else {
            v30 = (v30 % ((v30) | 1))
        }
        v33 := []int{32768}
        var v34 uint16 = 1
        _ = fi31
        _ = v33
        _ = v34
    }
    var v35 float64 = (0.0 + (float64(1.5) + (((1.5 * ((0.0 - 0.1) * 3.14159)) - 0.0) - -2.25)))
    v36 := GenS3{F0: v2, F1: (("hello" + "a") + "a")}
    v37 := GenS4{F0: genFn0(v5), F1: 2}
    v38 := GenS5{F0: 1, F1: 1, F2: "\n\t"}
    for fi39 := 0; fi39 < 8; fi39++ {
        fmt.Println(v5)
        v35 += (v35 * (v35 * v35))
        for fi40 := 0; fi40 < 10; fi40++ {
            if false {
                v2 = v2
                if true {
                    var v41 string = "a"
                    v42 := []int{fi39, (fi40 ^ fi39), (fi39 / (((fi40 & (^(fi39)))) | 1)), (fi39 + -695)}
                    v3 += v3
                    _ = v41
                    _ = v42
                }
            } else {
                v43 := GenS1{F0: 65535}
                if true {
                    var v44 int64 = v5
                    var v45 bool = (9223372036854775807 <= (fi39 / ((fi40) | 1)))
                    fmt.Println(v44, v2)
                    _ = v44
                    _ = v45
                }
                _ = v43
            }
            v5 = v5
            _ = fi40
        }
        _ = fi39
    }
    if (false || true) {
        var v46 uint32 = 0
        var v47 int8 = 100
        if (((v2 < v2) || true) && (false && ((1e10 * float32((1e10 - v35))) == -2.25))) {
            v47 += 127
        } else {
            var v48 uint8 = (^(v30))
            v4[42] = "a"
            if true {
                v35 -= float64(float32(v35))
            }
            _ = v48
        }
        _ = v46
        _ = v47
    }
    switch "foo bar" {
    case "hello":
        fmt.Println(v5, v35, v35)
    case "foo bar":
        v4[int(v2)] = (("a" + "foo bar") + ("" + (("" + "hello") + "a")))
    case "":
        if true {
            var v49 bool = (88 < 895)
            _ = v49
        } else {
            v50 := []int{127, 4294967295, 1, -1}
            for fi51 := 0; fi51 < 13; fi51++ {
                v5 = (v5 & v5)
                _ = fi51
            }
            _ = v50
        }
    }
    v30 -= (^(v30))
    switch 846 {
    case -1:
        var v52 int16 = v2
        _ = v52
    case 2147483647:
        _ = len(v4)
    default:
        v53 := []int{32767, -1}
        _ = v53
    }
    for fi54 := 0; fi54 < 12; fi54++ {
        for fi55 := 0; fi55 < 11; fi55++ {
            var v56 int = fi54
            if false {
                fmt.Println(v30, v3)
                v57 := GenS5{F0: 42, F1: 32768, F2: ("hello" + "hello")}
                _ = v57
            }
            v5 -= -1
            v58 := GenS6{F0: 32767, F1: 973, F2: 1}
            fmt.Println(v35, v30, v56)
            _ = fi55
            _ = v56
            _ = v58
        }
        v59 := GenS7{F0: 65535, F1: true, F2: 0}
        _ = fi54
        _ = v59
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
