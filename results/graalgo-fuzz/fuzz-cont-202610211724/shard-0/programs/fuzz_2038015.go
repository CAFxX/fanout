package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int }
func genFn0(p0 bool) int8 {
    var v5 uint16 = 98
    v5 -= 967
    fmt.Println(v5)
    var v6 uint64 = 1
    p0 = p0
    _ = p0
    _ = v5
    _ = v6
    return -128
}

func main() {
    var v1 float32 = (float32(0.1) / float32((float64(0.1) - (1e10 / 3.14159))))
    v2 := []string{("a" + "a"), "hello", "hello"}
    var v3 uint64 = 0
    v4 := []string{(((("foo bar" + "") + "\n\t") + "hello") + "hello"), "a", (("a" + "a") + (("foo bar" + ((("foo bar" + "hello") + ("a" + "hello")) + ("\n\t" + (("" + "") + "foo bar")))) + (("foo bar" + "hello") + "\n\t")))}
    v1 = v1
    if len(v4) > 0 {
        v4[0] = ("hello" + "a")
    }
    if (false || true) {
        var v7 float64 = 0.0
        switch 42 {
        case 128:
            var v8 float32 = 0.0
            _ = v8
        }
        _ = v7
    }
    if len(v4) > 0 {
        v4[0] = ((("\n\t" + "\n\t") + ("a" + "")) + "")
    }
    if true {
        switch ("\n\t" + ("\n\t" + "\n\t")) {
        case "":
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint16 = 205
                _ = fi9
                _ = v10
            }
        default:
            v3 = (v3 & 2)
        }
        if (255 < 255) {
            switch -13 {
            case 127:
                var v11 int16 = -755
                _ = v11
            case 9223372036854775807:
                switch -964 {
                case 865:
                    switch 32767 {
                    case 1:
                        v1 += (v1 + (1e10 / v1))
                    case 128:
                        switch ("a" + "") {
                        case "hello":
                            v2 = append(v2, "\n\t")
                        default:
                            var v12 int32 = -624
                            _ = v12
                        }
                    case 144:
                        fmt.Println(v3, v1)
                    default:
                        v1 = ((v1 * v1) - 3.14159)
                    }
                case 128:
                    v3 = v3
                default:
                    if false {
                        var v13 uint16 = 132
                        var v14 int16 = 624
                        v2 = append(v2, "hello")
                        _ = v13
                        _ = v14
                    }
                }
            }
            fmt.Println(v3, v3)
            var v15 string = "\n\t"
            _ = v15
        }
        if false {
            var v16 uint32 = 2147483648
            v17 := GenS1{F0: 0, F1: 32767, F2: 2147483647}
            _ = v16
            _ = v17
        }
    }
    var v18 int16 = -146
    if true {
        var v19 uint = 750
        switch -131 {
        case -499:
            v18 = (^(v18))
        }
        _ = v19
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 uint32 = 4294967295
        for fi22 := 0; fi22 < 8; fi22++ {
            if ((!false) || (false && (1e10 == 1.5))) {
                v21 = 0
                fmt.Println(fi20, fi20)
                v1 += float32(0.1)
            }
            _ = fi22
        }
        v3 += (v3 / ((v3) | 1))
        _ = fi20
        _ = v21
    }
    if (0 > 0) {
        if true {
            var v23 uint8 = 187
            v18 += v18
            _ = v23
        }
        switch "" {
        case "hello":
            for fi24 := 0; fi24 < 15; fi24++ {
                v2 = append(v2, "hello")
                var v25 int32 = 331
                _ = fi24
                _ = v25
            }
        }
    }
    var v26 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
