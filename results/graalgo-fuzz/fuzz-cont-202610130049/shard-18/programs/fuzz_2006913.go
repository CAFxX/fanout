package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint64; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 int }
func genFn0(p0 bool, p1 int8) uint8 {
    var v4 int32 = 42
    v5 := []int{258, 596, -1, 32768}
    v6 := GenS1{F0: "foo bar"}
    for fi7 := 0; fi7 < 10; fi7++ {
        switch "a" {
        case "foo bar":
            for fi8 := 0; fi8 < 8; fi8++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if true {
                    fmt.Println(p0)
                    var v9 uint = 65535
                    var v10 uint32 = 0
                    _ = v10
                    _ = v9
                } else {
                    v4 += v4
                    p1 += 1
                }
                var v11 float64 = (3.14159 - (0.0 + 3.14159))
                fmt.Println(v4, p1, fi7)
                v12 := GenS1{F0: ("" + ("a" + "a"))}
                _ = fi8
                _ = v11
                _ = v12
            }
        }
        switch fi7 {
        case 9223372036854775807:
            p0 = p0
        case 1:
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 float64 = 1e10
                var v15 uint = 18446744073709551615
                switch (-(fi7)) {
                case 4294967295:
                    p0 = (int32(fi13) > 42)
                case 0:
                    if (!true) {
                        var v16 uint8 = 42
                        var v17 float32 = (1.5 * float32((3.14159 + (-2.25 - (v14 * (v14 + v14))))))
                        _ = v16
                        _ = v17
                    } else {
                        var v18 int8 = 100
                        p1 += (-(p1))
                        v19 := GenS1{F0: ("foo bar" + "")}
                        _ = v18
                        _ = v19
                    }
                case 42:
                    v20 := GenS1{F0: "hello"}
                    _ = v20
                default:
                    var v21 int = 0
                    _ = v21
                }
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        v5 = append(v5, -2147483648)
        _ = fi7
    }
    if (!true) {
        if p0 {
            var v22 uint64 = 324
            if p0 {
                v4 += -674
                switch 42 {
                case -9223372036854775808:
                    var v23 uint8 = 27
                    _ = v23
                case 32767:
                    p0 = (2 <= 2147483647)
                case -898:
                    v24 := GenS2{F0: 32767}
                    _ = v24
                }
                switch "hello" {
                case "":
                    p0 = false
                case "hello":
                    var v25 uint32 = 0
                    _ = v25
                }
            }
            _ = v22
        } else {
            p1 = (-(p1))
            for fi26 := 0; fi26 < 11; fi26++ {
                if true {
                    for fi27 := 0; fi27 < 11; fi27++ {
                        var v28 float64 = ((1.5 - 3.14159) / (1e10 * 0.1))
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(p1, fi26)
                        default:
                            v29 := GenS1{F0: "a"}
                            _ = v29
                        }
                        _ = fi27
                        _ = v28
                    }
                    p0 = true
                } else {
                    v5 = append(v5, (fi26 - 1))
                }
                if len(v5) > 0 {
                    v5[0] = fi26
                }
                _ = fi26
            }
            var v30 float64 = (1.5 - float64((float32((((0.0 - ((3.14159 + float64(0.0)) + 3.14159)) / 1e10) + 0.0)) / -2.25)))
            _ = v30
        }
        var v31 uint32 = 2147483648
        v32 := GenS3{F0: 255, F1: 213, F2: -110}
        _ = v31
        _ = v32
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 141
}
func genFn1(p0 uint16) uint32 {
    v33 := GenS4{F0: 18446744073709551615}
    v34 := GenS4{F0: 1}
    _ = p0
    _ = v33
    _ = v34
    return 2147483647
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 1
    v3 := make(map[string]int)
    v2 += v2
    v35 := GenS4{F0: 42}
    v2 += (v2 / ((4294967295) | 1))
    v36 := GenS5{F0: 42, F1: -1}
    var v37 bool = (198 == 323)
    if (!(!false)) {
        v37 = false
    }
    fmt.Println(v37)
    v38 := []int{42}
    fmt.Println(v2, v37)
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
}
