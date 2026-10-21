package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 uint32; F2 bool }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint16 }
func genFn0(p0 uint64, p1 uint8) uint {
    var v5 uint = 0
    var v6 int32 = 625
    switch "\n\t" {
    case "hello":
        v7 := GenS1{F0: (p1 > (p1 >> ((65535) & 63)))}
        _ = v7
    case "\n\t":
        var v8 float32 = 0.0
        _ = v8
    case "":
        v9 := GenS1{F0: false}
        _ = v9
    default:
        v5 = 255
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return v5
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v4 float64 = 0.1
    fmt.Println(v4)
    v4 -= (v4 + float64(float32(((float64(0.0) * -2.25) * v4))))
    fmt.Println(v4)
    switch "\n\t" {
    case "a":
        var v10 int8 = 127
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        v4 -= float64(float32(v4))
        if ((42 < 2) && (false && (1 == 0))) {
            if true {
                for fi12 := 0; fi12 < 18; fi12++ {
                    switch fi11 {
                    case 767:
                        fmt.Println(fi12)
                    case -2147483648:
                        var v13 bool = ((((0 <= 120) || ((1 >= 67) && true)) || true) && ((!(("hello" + (("\n\t" + ("hello" + ("a" + "a"))) + "")) >= "hello")) || false))
                        _ = v13
                    case 32767:
                        v14 := v2[127]
                        _ = v14
                    default:
                        var v15 float32 = (1.5 / 1.5)
                        _ = v15
                    }
                    v4 = (v4 * v4)
                    v4 += (v4 / 1.0)
                    _ = fi12
                }
                for fi16 := 0; fi16 < 12; fi16++ {
                    var v17 uint64 = 18446744073709551615
                    v4 = v4
                    _ = fi16
                    _ = v17
                }
                var v18 int64 = -227
                _ = v18
            }
            v19 := GenS2{F0: -9223372036854775808, F1: 918, F2: false}
            _ = v19
        } else {
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 string = ("foo bar" + ("\n\t" + "a"))
                _ = fi20
                _ = v21
            }
            if true {
                v22 := v3[""]
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 bool = (1 < 255)
                    _ = fi23
                    _ = v24
                }
                v4 = 3.14159
                _ = v22
            } else {
                for fi25 := 0; fi25 < 1; fi25++ {
                    v26 := []int{fi25}
                    if (1 <= 127) {
                        v4 = v4
                        var v27 uint = 65535
                        delete(v3, "foo bar")
                        _ = v27
                    }
                    v4 = float64((3.14159 * 3.14159))
                    v3["\n\t"] = fi11
                    v4 = 1.5
                    _ = fi25
                    _ = v26
                }
            }
        }
        _ = fi11
    }
    v4 = 3.14159
    fmt.Println(v4)
    v4 = v4
    if (-108 <= 127) {
        if true {
            var v28 float32 = 0.1
            var v29 uint8 = 127
            _ = v28
            _ = v29
        } else {
            v4 += v4
            var v30 int = -626
            var v31 bool = true
            _ = v30
            _ = v31
        }
    } else {
        v4 = 1.5
        var v32 int8 = 48
        for fi33 := 0; fi33 < 7; fi33++ {
            v34 := []int{(fi33 * (fi33 | (^(fi33)))), (fi33 << ((genFn0(42, 127)) & 63))}
            _ = fi33
            _ = v34
        }
        _ = v32
    }
    _ = len(v2)
    v35 := GenS3{F0: 138, F1: 42}
    var v36 int16 = 101
    if (v36 >= 42) {
        v37 := GenS4{F0: ((42 > 42) && true)}
        v38 := []int{-2147483648, -1, -9223372036854775808}
        _ = v37
        _ = v38
    } else {
        v39 := GenS4{F0: false}
        _ = v39
    }
    var v40 int16 = v36
    fmt.Println(v36, v40, v40)
    if false {
        v41 := GenS4{F0: true}
        _ = v41
    } else {
        var v42 float32 = (((-2.25 - 1e10) / 1.0) / 0.1)
        var v43 uint8 = 42
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 4; fi44++ {
        _ = len(v2)
        if ((1e10 > v4) && (0 != 412)) {
            v45 := GenS5{F0: fi44, F1: 0}
            _ = v45
        }
        _ = fi44
    }
    var v46 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v46
}
