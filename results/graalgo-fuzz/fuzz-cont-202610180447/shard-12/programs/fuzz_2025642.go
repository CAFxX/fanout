package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int; F2 uint }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint8; F1 uint; F2 int64 }
type GenS5 struct { F0 bool; F1 int16; F2 int32 }
func genFn0(p0 uint, p1 int64, p2 int64) int32 {
    if false {
        var v6 string = ("foo bar" + "a")
        for fi7 := 0; fi7 < 11; fi7++ {
            fmt.Println(v6, fi7)
            var v8 int64 = p1
            switch -1 {
            case -9223372036854775808:
                var v9 int32 = 227
                _ = v9
            case 127:
                var v10 float32 = (0.0 - (0.0 - 1.5))
                _ = v10
            default:
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 uint8 = 42
                    v6 = v6
                    v13 := []int{(fi11 << (((^(p0))) & 63)), (fi7 - fi7), 2147483647}
                    var v14 int = fi7
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
            }
            _ = fi7
            _ = v8
        }
        v15 := GenS1{F0: 0}
        _ = v15
        _ = v6
    }
    if (p0 > p0) {
        p2 -= int64(p0)
    } else {
        switch (("foo bar" + ("hello" + ("\n\t" + "\n\t"))) + "") {
        case "a":
            p0 = (p0 >> (((p0 % ((0) | 1))) & 63))
        }
        fmt.Println(p1)
    }
    if (9223372036854775807 < 128) {
        if true {
            v16 := GenS1{F0: 1}
            var v17 int8 = -1
            _ = v16
            _ = v17
        } else {
            var v18 uint64 = 0
            switch "" {
            case "":
                v18 -= 4294967295
            case "hello":
                var v19 uint32 = 2147483647
                _ = v19
            default:
                p0 += (p0 ^ 565)
            }
            _ = v18
        }
        fmt.Println(p2, p1, p1)
        p2 = -1
    } else {
        if false {
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 float32 = 0.1
                v22 := GenS2{F0: -32768, F1: (fi20 / (((fi20 * fi20)) | 1)), F2: 0}
                v21 = 1.5
                _ = fi20
                _ = v21
                _ = v22
            }
            v23 := []int{127, -2147483648, 32768}
            _ = v23
        }
    }
    v24 := GenS3{F0: 127, F1: -2147483648}
    v25 := GenS2{F0: -32768, F1: 32768, F2: (p0 >> ((p0) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    return 42
}

func main() {
    v1 := []string{(("hello" + "hello") + "foo bar"), ("a" + ""), "foo bar", ("foo bar" + "\n\t")}
    var v2 uint8 = 128
    v3 := []uint32{2, 2147483647, 0}
    var v4 float32 = -2.25
    var v5 int16 = -586
    v26 := GenS2{F0: v5, F1: 2147483647, F2: 890}
    var v27 uint = 18446744073709551615
    var v28 int32 = 0
    var v29 int8 = -128
    var v30 uint64 = 4294967295
    v31 := GenS4{F0: v2, F1: v27, F2: 694}
    var v32 bool = (116 < (v29 | (v29 * -128)))
    v30 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v33 bool = (v32 && false)
    var v34 uint16 = 0
    v1 = append(v1, ((("" + (((("a" + ("foo bar" + "\n\t")) + "hello") + "a") + "")) + "\n\t") + "a"))
    v35 := GenS5{F0: true, F1: (v5 ^ v5), F2: v28}
    var v36 uint32 = 2147483648
    var v37 bool = v33
    v38 := GenS3{F0: (^(v2)), F1: genFn0(223, 44, int64(v27))}
    for fi39 := 0; fi39 < 1; fi39++ {
        if false {
            if (v33 || v33) {
                for fi40 := 0; fi40 < 13; fi40++ {
                    var v41 int16 = v5
                    v4 -= -2.25
                    _ = fi40
                    _ = v41
                }
                v27 -= v27
            }
            if len(v3) > 0 {
                v3[0] = 966
            }
        }
        _ = fi39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
