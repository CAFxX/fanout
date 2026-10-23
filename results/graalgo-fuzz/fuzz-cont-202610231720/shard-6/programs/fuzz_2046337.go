package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
func genFn0(p0 bool, p1 string) int16 {
    v4 := []int{-648, 127, -350, -1}
    p1 = "\n\t"
    p1 = (("hello" + p1) + p1)
    _ = p0
    _ = p1
    _ = v4
    return 42
}
func genFn1() int {
    if (false && (1 <= 65535)) {
        var v5 bool = (1 < 82)
        if (!(!false)) {
            for fi6 := 0; fi6 < 14; fi6++ {
                switch fi6 {
                case -840:
                    var v7 uint16 = 718
                    _ = v7
                case 9223372036854775807:
                    v8 := []int{(fi6 - 42), 699, (^(fi6))}
                    _ = v8
                case 42:
                    v5 = false
                default:
                    v9 := []int{(^(fi6)), fi6, (fi6 * 965), (fi6 + fi6)}
                    _ = v9
                }
                v10 := []int{(fi6 << ((1) & 63)), 0, (fi6 * -2147483648)}
                _ = fi6
                _ = v10
            }
        } else {
            v11 := GenS1{F0: (9223372036854775807 <= 1), F1: 32767}
            switch ("hello" + ("" + ("hello" + ("foo bar" + ("\n\t" + (("a" + "foo bar") + "")))))) {
            case "a":
                var v12 int64 = 0
                _ = v12
            default:
                var v13 uint32 = 0
                _ = v13
            }
            _ = v11
        }
        _ = v5
    }
    v14 := GenS1{F0: false, F1: 1}
    for fi15 := 0; fi15 < 11; fi15++ {
        if (128 != (fi15 * fi15)) {
            var v16 int8 = 0
            v16 -= v16
            fmt.Println(v16)
            _ = v16
        }
        _ = fi15
    }
    _ = v14
    return 42
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    var v17 uint64 = 42
    var v18 bool = (("" + "foo bar") <= (("hello" + ("foo bar" + "foo bar")) + ("a" + "")))
    var v19 float64 = 3.14159
    for fi20 := 0; fi20 < 16; fi20++ {
        if true {
            v17 -= 18446744073709551615
            v21 := GenS1{F0: (uint16(v17) >= 745), F1: (fi20 * (-(fi20)))}
            fmt.Println(v19, v1)
            _ = v21
        } else {
            if (1.5 != 0.1) {
                var v22 uint8 = 0
                _ = v22
            }
            v18 = (!("\n\t" >= "a"))
        }
        _ = fi20
    }
    if (!(0 >= 690)) {
        for fi23 := 0; fi23 < 15; fi23++ {
            for fi24 := 0; fi24 < 13; fi24++ {
                fmt.Println(v17, v1, fi24)
                _ = fi24
            }
            v25 := GenS1{F0: false, F1: 2147483647}
            _ = fi23
            _ = v25
        }
    } else {
        v26 := GenS2{F0: (v17 << ((255) & 63)), F1: genFn1()}
        _ = v26
    }
    if (!(!true)) {
        v18 = v18
        for fi27 := 0; fi27 < 3; fi27++ {
            switch "foo bar" {
            case "\n\t":
                _ = len(v2)
            default:
                if (!v18) {
                    switch "hello" {
                    case "hello":
                        v28 := GenS1{F0: true, F1: 2147483647}
                        _ = v28
                    default:
                        var v29 int32 = 2147483647
                        _ = v29
                    }
                    var v30 string = ("foo bar" + ("hello" + ("foo bar" + ("foo bar" + "\n\t"))))
                    v31 := v3[v30]
                    _ = v30
                    _ = v31
                } else {
                    v17 = 1
                }
            }
            for fi32 := 0; fi32 < 2; fi32++ {
                switch (-(fi32)) {
                case 32768:
                    v1 -= -275
                }
                _ = fi32
            }
            _ = fi27
        }
    } else {
        delete(v2, -305)
        v33 := []int{genFn1(), genFn1(), 2147483647, -427}
        switch 128 {
        case 9223372036854775807:
            v34 := GenS1{F0: v18, F1: -2147483648}
            _ = v34
        case 127:
            v17 -= (v17 ^ v17)
        }
        _ = v33
    }
    fmt.Println(v1, v19, v19)
    v35 := GenS3{F0: 333}
    if v18 {
        var v36 float32 = 0.0
        _ = v36
    }
    if (0 > 2) {
        var v37 float64 = (v19 * (((float64(3.14159) - v19) - 0.1) - v19))
        v37 -= float64(1e10)
        _ = v37
    }
    if (!(!v18)) {
        var v38 int = genFn1()
        var v39 int64 = -1
        _ = v38
        _ = v39
    }
    v19 -= v19
    var v40 float32 = (float32(v19) + 3.14159)
    fmt.Println(v40, v18, v19)
    v41 := GenS3{F0: genFn1()}
    if ("a" >= "\n\t") {
        v42 := GenS2{F0: v17, F1: genFn1()}
        switch "a" {
        case "":
            if v18 {
                v43 := GenS3{F0: -815}
                _ = len(v3)
                v17 += 42
                _ = v43
            }
        }
        _ = v42
    }
    v44 := GenS4{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v40
    _ = v41
    _ = v44
}
