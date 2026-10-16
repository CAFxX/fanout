package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint32; F1 int8; F2 int64 }
func genFn0(p0 uint32, p1 string, p2 uint64) int64 {
    p2 -= p2
    switch 562 {
    case 42:
        if false {
            if true {
                var v5 int64 = 9223372036854775807
                v5 = (v5 ^ (-(v5)))
                _ = v5
            }
            var v6 int8 = 127
            p1 = p1
            _ = v6
        } else {
            v7 := GenS1{F0: 1, F1: -1, F2: p0}
            if false {
                var v8 string = p1
                switch "\n\t" {
                case "hello":
                    fmt.Println(p0, v8)
                }
                if false {
                    p1 = (v8 + "a")
                    if (true && false) {
                        p1 = "foo bar"
                        fmt.Println(p1)
                        p0 -= uint32(p2)
                    }
                } else {
                    fmt.Println(p0, p2, p1)
                }
                _ = v8
            }
            switch 4294967295 {
            case 32768:
                var v9 string = p1
                _ = v9
            case 32767:
                p2 = 0
            case 128:
                if false {
                    p0 = p0
                    fmt.Println(p1, p1)
                } else {
                    var v10 uint32 = (p0 - p0)
                    _ = v10
                }
            default:
                fmt.Println(p0)
            }
            _ = v7
        }
    default:
        if ((127 >= 42) || false) {
            v11 := []int{}
            _ = v11
        }
    }
    p1 = "hello"
    var v12 uint8 = 57
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -9223372036854775808
}
func genFn1(p0 uint16) uint16 {
    v13 := []int{0, -2147483648}
    var v14 int32 = -478
    var v15 uint = 42
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    return 894
}

func main() {
    var v1 bool = ((-9223372036854775808 != -9223372036854775808) || (-818 == 0))
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    v4 := make(map[int]uint32)
    if (1 > 9223372036854775807) {
        v1 = (v1 || v1)
        var v16 int16 = 0
        _ = v16
    }
    var v17 uint16 = 376
    if false {
        var v18 uint32 = 204
        switch 1 {
        case 9223372036854775807:
            v19 := GenS2{F0: 460, F1: v18}
            _ = v19
        default:
            v20 := GenS2{F0: -1, F1: v18}
            _ = v20
        }
        _ = v18
    } else {
        v21 := GenS3{F0: 127, F1: 4294967295}
        v1 = v1
        _ = v21
    }
    _ = len(v4)
    switch ("a" + "foo bar") {
    case "a":
        var v22 uint64 = 1
        _ = v22
    case "hello":
        v1 = (127 < 100)
    case "foo bar":
        switch 9223372036854775807 {
        case -688:
            for fi23 := 0; fi23 < 9; fi23++ {
                v2 = append(v2, 1)
                var v24 float32 = (0.1 / 1.0)
                v24 = (1e10 / (v24 * v24))
                _ = fi23
                _ = v24
            }
        case 32768:
            for fi25 := 0; fi25 < 7; fi25++ {
                switch fi25 {
                case 32767:
                    v17 = (^(v17))
                case 2147483647:
                    if len(v2) > 0 {
                        v2[0] = 74
                    }
                default:
                    v26 := GenS3{F0: -90, F1: 927}
                    _ = v26
                }
                _ = fi25
            }
        default:
            v17 += genFn1((v17 + (v17 | 384)))
        }
    default:
        v27 := GenS4{F0: 2147483648, F1: 0, F2: genFn0(4294967295, "foo bar", 204)}
        _ = v27
    }
    switch ("foo bar" + "hello") {
    case "":
        v1 = (42 != 127)
    default:
        if v1 {
            var v28 uint = 65535
            var v29 uint64 = 744
            _ = v28
            _ = v29
        } else {
            v30 := GenS2{F0: 323, F1: 2}
            if v1 {
                var v31 float32 = 3.14159
                v32 := GenS2{F0: 218, F1: 42}
                if (4294967295 == 65535) {
                    v1 = v1
                } else {
                    v17 = 65535
                    v1 = false
                    var v33 int32 = -2147483648
                    _ = v33
                }
                _ = v31
                _ = v32
            } else {
                if ((((((((0.1 / (float64(((float32((1.5 + (0.1 + -2.25))) - 1.5) / 3.14159)) - 0.0)) * ((1.5 - 0.0) - (1.5 - 1.5))) * 0.0) + 0.0) / -2.25) - (0.1 - float64((0.1 / (1e10 + 3.14159))))) / 3.14159) == (float64(float32(0.1)) + 1e10)) {
                    _ = len(v2)
                    fmt.Println(v1, v1)
                    v17 = v17
                }
                for fi34 := 0; fi34 < 13; fi34++ {
                    var v35 uint = 18446744073709551615
                    _ = fi34
                    _ = v35
                }
            }
            if v1 {
                v17 += genFn1((^(v17)))
                if len(v2) > 0 {
                    v2[0] = uint32(v17)
                }
                fmt.Println(v1, v1)
            } else {
                var v36 int = 4294967295
                _ = v36
            }
            _ = v30
        }
    }
    var v37 int64 = -9223372036854775808
    v38 := GenS1{F0: -368, F1: 931, F2: 510}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
