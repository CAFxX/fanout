package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS4 struct { F0 string; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint16, p1 uint16) bool {
    var v4 int64 = -1
    var v5 string = ("hello" + "hello")
    p0 += (^(p1))
    var v6 uint64 = 4294967295
    var v7 int64 = 42
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return false
}
func genFn1(p0 uint8, p1 uint64) int32 {
    if (p0 >= (p0 >> ((4294967295) & 63))) {
        if (42 != 65535) {
            var v8 int16 = -1
            var v9 bool = (false || false)
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 2, F1: 4294967295}
        _ = v10
    } else {
        var v11 float64 = 1e10
        _ = v11
    }
    if true {
        switch "\n\t" {
        case "hello":
            var v12 int32 = -2147483648
            _ = v12
        case "a":
            var v13 uint32 = 0
            _ = v13
        case "foo bar":
            fmt.Println(p1)
        default:
            switch 127 {
            case 32768:
                var v14 bool = (!true)
                _ = v14
            case -556:
                v15 := GenS1{F0: 2, F1: 793}
                _ = v15
            case -741:
                switch "" {
                case "foo bar":
                    fmt.Println(p0)
                default:
                    var v16 int = -9223372036854775808
                    _ = v16
                }
            default:
                v17 := GenS1{F0: 895, F1: 65535}
                _ = v17
            }
        }
        var v18 string = (("\n\t" + ("" + "foo bar")) + "a")
        var v19 int32 = 0
        _ = v18
        _ = v19
    }
    var v20 int8 = 100
    v21 := []int{}
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return -668
}

func main() {
    var v1 float64 = ((1.5 - 1.5) + ((-2.25 - 3.14159) / ((3.14159 - 1e10) - float64(0.0))))
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    v1 -= 1.5
    switch "hello" {
    case "foo bar":
        v1 -= (v1 - (v1 - v1))
    }
    var v22 float32 = 1e10
    v23 := GenS2{F0: 42, F1: 686, F2: 1}
    v22 = 0.0
    if (!((1 == 1) && true)) {
        if true {
            for fi24 := 0; fi24 < 20; fi24++ {
                v1 = v1
                v22 -= v22
                v22 -= 0.0
                _ = fi24
            }
            v1 -= 3.14159
        }
    }
    if false {
        var v25 int64 = 1
        if false {
            for fi26 := 0; fi26 < 18; fi26++ {
                v27 := GenS3{F0: 2, F1: -32768, F2: 830}
                var v28 int32 = 42
                v29 := GenS4{F0: "hello", F1: 25, F2: 106}
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            var v30 int16 = 0
            v31 := GenS1{F0: 854, F1: 2}
            _ = v30
            _ = v31
        } else {
            switch -9223372036854775808 {
            case 834:
                for fi32 := 0; fi32 < 5; fi32++ {
                    v1 -= float64(v22)
                    for fi33 := 0; fi33 < 5; fi33++ {
                        v34 := []int{(^(fi33)), 523}
                        if (v25 == (-(v25))) {
                            v35 := GenS1{F0: 42, F1: 0}
                            v36 := GenS5{F0: 171, F1: 107}
                            _ = v35
                            _ = v36
                        } else {
                            v25 += (v25 * (v25 | 42))
                            var v37 int16 = -776
                            _ = v37
                        }
                        v22 -= (v22 + ((v22 - (v22 / float32((float64((-2.25 - 0.0)) + 0.0)))) * 0.1))
                        _ = fi33
                        _ = v34
                    }
                    _ = fi32
                }
            case 4294967295:
                var v38 string = "hello"
                _ = v38
            }
        }
        var v39 float32 = (float32((float64(0.1) / float64(3.14159))) + (v22 - 1e10))
        _ = v25
        _ = v39
    } else {
        var v40 int8 = 127
        var v41 int32 = -495
        v42 := []int{32768}
        _ = v40
        _ = v41
        _ = v42
    }
    fmt.Println(v1, v22)
    v22 -= v22
    v1 = (1e10 * (v1 - float64(v22)))
    var v43 bool = false
    var v44 int64 = 0
    var v45 int = -310
    v46 := GenS4{F0: "foo bar", F1: 120, F2: 127}
    fmt.Println(v44, v1)
    fmt.Println(v22, v45)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
