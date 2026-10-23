package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint }
type GenS3 struct { F0 string; F1 string; F2 uint32 }
type GenS4 struct { F0 int8; F1 uint32; F2 int }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int32) uint64 {
    v6 := GenS1{F0: 212, F1: 1, F2: 646}
    fmt.Println(p0)
    var v7 float32 = float32(-2.25)
    if true {
        switch ("a" + "a") {
        case "hello":
            fmt.Println(p0)
        }
    } else {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 0, F1: 439, F2: 42}
            if true {
                p0 = p0
            } else {
                fmt.Println(p0)
                var v10 uint16 = 159
                if (!((p0 >= p0) || (!(((1e10 * (1e10 / float64(0.1))) * float64(v7)) >= 3.14159)))) {
                    p0 += p0
                } else {
                    v11 := GenS2{F0: 248, F1: v10, F2: 44}
                    _ = v11
                }
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        if (!(42 != 2147483647)) {
            if (!false) {
                switch "a" {
                case "":
                    var v12 int64 = 42
                    _ = v12
                case "a":
                    var v13 int = 4294967295
                    _ = v13
                case "\n\t":
                    var v14 uint16 = 0
                    _ = v14
                default:
                    var v15 float32 = v7
                    _ = v15
                }
            }
            var v16 uint64 = 767
            var v17 int8 = 1
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            fmt.Println(fi18, fi18)
            var v19 int8 = 32
            v20 := GenS2{F0: 128, F1: 268, F2: 2}
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int, p1 string) int64 {
    var v21 int64 = 9223372036854775807
    v22 := GenS1{F0: 525, F1: 4294967295, F2: 0}
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    return (-(v21))
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    var v3 int32 = 0
    v4 := []string{""}
    var v5 int64 = 0
    switch "" {
    case "a":
        if (true || (!(false && false))) {
            var v23 bool = true
            var v24 uint64 = 42
            if (1e10 >= ((0.1 - -2.25) + 1e10)) {
                for fi25 := 0; fi25 < 11; fi25++ {
                    switch fi25 {
                    case 128:
                        var v26 int8 = 42
                        _ = v26
                    default:
                        v23 = false
                    }
                    fmt.Println(v24)
                    v5 = 9223372036854775807
                    _ = fi25
                }
            } else {
                var v27 string = "a"
                for fi28 := 0; fi28 < 9; fi28++ {
                    var v29 int16 = 32767
                    _ = fi28
                    _ = v29
                }
                var v30 uint = 808
                _ = v27
                _ = v30
            }
            _ = v23
            _ = v24
        }
    case "hello":
        if false {
            fmt.Println(v5, v5)
            if (!false) {
                v31 := GenS3{F0: "\n\t", F1: "a", F2: 705}
                for fi32 := 0; fi32 < 18; fi32++ {
                    v3 -= 1
                    _ = fi32
                }
                _ = v31
            } else {
                v33 := GenS4{F0: 100, F1: 850, F2: 32767}
                var v34 float64 = 1.5
                _ = v33
                _ = v34
            }
        }
    }
    v5 -= v5
    v3 = v3
    fmt.Println(v5)
    v35 := GenS1{F0: 65535, F1: 4294967295, F2: v5}
    for fi36 := 0; fi36 < 1; fi36++ {
        v37 := v2["hello"]
        var v38 string = v37
        v39 := GenS5{F0: 65535}
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    v3 = (v3 ^ (v3 - v3))
    _ = len(v4)
    if true {
        var v40 uint16 = 693
        if (!(0 > uint32(v40))) {
            if (((false && (489 < (-(v5)))) && false) || (!true)) {
                v5 = v5
                switch -2147483648 {
                case 1:
                    switch "" {
                    case "a":
                        v41 := GenS5{F0: 0}
                        _ = v41
                    }
                case 128:
                    fmt.Println(v3, v3, v5)
                case 9223372036854775807:
                    fmt.Println(v3)
                default:
                    fmt.Println(v40, v40, v5)
                }
            }
            v42 := GenS6{F0: 1}
            v40 -= 0
            _ = v42
        } else {
            fmt.Println(v5, v40)
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        }
        v40 -= (v40 ^ (v40 >> ((255) & 63)))
        _ = v40
    }
    v43 := GenS7{F0: 366}
    var v44 int8 = -1
    var v45 int16 = -78
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
