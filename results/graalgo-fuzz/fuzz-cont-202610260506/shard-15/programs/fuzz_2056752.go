package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 bool; F2 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int16, p1 int8, p2 int16) string {
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 42, F1: 4294967295}
        v9 := []int{(fi7 >> ((94) & 63)), (fi7 + (fi7 ^ fi7)), -493}
        _ = fi7
        _ = v8
        _ = v9
    }
    switch "" {
    case "hello":
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := []int{(fi10 & fi10), (^(fi10)), (^(fi10)), 9223372036854775807}
            _ = fi10
            _ = v11
        }
    case "":
        v12 := GenS1{F0: 255, F1: 18446744073709551615}
        _ = v12
    case "a":
        v13 := GenS1{F0: 32767, F1: 2}
        _ = v13
    default:
        fmt.Println(p1)
    }
    if true {
        p0 += 601
        p1 -= p1
    } else {
        fmt.Println(p2)
    }
    var v14 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (((((("a" + "foo bar") + ("hello" + "\n\t")) + ("a" + "\n\t")) + "a") + "a") + "hello")
}
func genFn1(p0 uint, p1 int8) uint64 {
    if false {
        var v15 uint16 = 42
        _ = v15
    }
    fmt.Println(p0)
    var v16 int = 32767
    var v17 uint16 = 0
    v16 += v16
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 359
}

func main() {
    var v1 int64 = -37
    var v2 int32 = 1
    var v3 int64 = 42
    var v4 uint16 = 98
    var v5 float64 = float64((0.0 - (0.1 - ((3.14159 * (0.1 * (((((float32((0.0 / 3.14159)) - -2.25) / -2.25) / 1.0) - (-2.25 / 0.1)) + (-2.25 * 1.5)))) * 0.0))))
    v6 := []uint32{42, 723, 2147483647, 2147483648}
    var v18 string = ("foo bar" + "foo bar")
    if len(v6) > 0 {
        v6[0] = 2147483648
    }
    fmt.Println(v2)
    v19 := GenS2{F0: 712}
    if ("foo bar" != v18) {
        v1 = 9223372036854775807
        v20 := GenS2{F0: 32767}
        _ = v20
    } else {
        if false {
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 int = (fi21 ^ (fi21 ^ fi21))
                var v23 uint = 18446744073709551615
                v3 -= (v1 / (((^(v3))) | 1))
                _ = fi21
                _ = v22
                _ = v23
            }
            v24 := GenS2{F0: -684}
            for fi25 := 0; fi25 < 15; fi25++ {
                switch v18 {
                case "\n\t":
                    for fi26 := 0; fi26 < 1; fi26++ {
                        v2 += (v2 - v2)
                        _ = fi26
                    }
                case "":
                    v27 := GenS3{F0: v18, F1: false, F2: fi25}
                    _ = v27
                case "foo bar":
                    for fi28 := 0; fi28 < 19; fi28++ {
                        v29 := GenS4{F0: 255}
                        _ = fi28
                        _ = v29
                    }
                }
                _ = fi25
            }
            _ = v24
        }
    }
    v30 := GenS1{F0: 255, F1: 584}
    v31 := GenS5{F0: 2}
    if true {
        if (!(float64(float32((-2.25 * 1.5))) > ((0.1 - float64(float32(((3.14159 + v5) - ((v5 - (v5 + v5)) / (((float64(0.1) - (float64(1e10) - (v5 / 3.14159))) + 1.5) * (3.14159 * 1e10))))))) / 1e10))) {
            for fi32 := 0; fi32 < 19; fi32++ {
                switch (v18 + v18) {
                case "":
                    fmt.Println(v18, v4, v1)
                case "\n\t":
                    if false {
                        var v33 uint32 = 0
                        _ = v33
                    } else {
                        var v34 uint16 = v4
                        v35 := GenS2{F0: -1}
                        _ = v34
                        _ = v35
                    }
                case "hello":
                    for fi36 := 0; fi36 < 1; fi36++ {
                        v5 -= v5
                        v37 := GenS5{F0: 0}
                        _ = fi36
                        _ = v37
                    }
                }
                for fi38 := 0; fi38 < 1; fi38++ {
                    var v39 float32 = (0.1 * ((3.14159 - 0.0) / float32((((float64((0.1 - 3.14159)) * 0.1) * (float64(3.14159) - (v5 * -2.25))) * (v5 + (1e10 - v5))))))
                    _ = fi38
                    _ = v39
                }
                _ = fi32
            }
        } else {
            var v40 int64 = v3
            for fi41 := 0; fi41 < 17; fi41++ {
                fmt.Println(v2, v5, v2)
                _ = fi41
            }
            _ = v40
        }
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
