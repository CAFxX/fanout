package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 bool; F1 int16; F2 int }
type GenS4 struct { F0 int64; F1 int64; F2 uint32 }

func main() {
    var v1 bool = (-9223372036854775808 > 9223372036854775807)
    var v2 int64 = -699
    var v3 int = 32768
    v1 = false
    v4 := []int{1, v3, (-(v3))}
    v5 := GenS1{F0: 127}
    v6 := GenS1{F0: 0}
    for fi7 := 0; fi7 < 16; fi7++ {
        switch 2147483647 {
        case 917:
            if len(v4) > 0 {
                v4[0] = (fi7 | (fi7 >> ((18446744073709551615) & 63)))
            }
        case -775:
            switch (-(v3)) {
            case 42:
                if v1 {
                    v1 = true
                } else {
                    if ((^(v2)) <= -9223372036854775808) {
                        v2 = 9223372036854775807
                        fmt.Println(v2, fi7, v1)
                    } else {
                        switch (-(fi7)) {
                        case 9223372036854775807:
                            var v8 int = v3
                            _ = v8
                        case 0:
                            var v9 int64 = 42
                            _ = v9
                        case 1:
                            var v10 float64 = 1.5
                            _ = v10
                        }
                        v11 := GenS2{F0: 424, F1: 798, F2: 18446744073709551615}
                        _ = v11
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v12 := GenS1{F0: -1}
                    _ = v12
                }
            case -2147483648:
                if (389 > 1) {
                    v13 := GenS2{F0: 0, F1: 402, F2: 18446744073709551615}
                    v14 := GenS3{F0: (v1 && (138 >= 1)), F1: 97, F2: v3}
                    fmt.Println(fi7, v3)
                    _ = v13
                    _ = v14
                }
            }
        case 9223372036854775807:
            if (((float64(3.14159) * 0.1) / 3.14159) != ((3.14159 + 0.0) + 3.14159)) {
                var v15 uint64 = 42
                _ = v15
            } else {
                var v16 uint16 = 0
                var v17 int = v3
                _ = v16
                _ = v17
            }
        default:
            fmt.Println(fi7, fi7, v1)
        }
        v3 = (fi7 * (fi7 | fi7))
        var v18 int16 = 32767
        _ = fi7
        _ = v18
    }
    var v19 int = v3
    v20 := GenS4{F0: v2, F1: v2, F2: 2147483647}
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 string = ((("a" + "\n\t") + "hello") + ((("" + ("" + ("\n\t" + "foo bar"))) + ("foo bar" + "")) + ("" + ("" + ("\n\t" + (("" + ("a" + (("" + ("hello" + "hello")) + "a"))) + "hello"))))))
        _ = fi21
        _ = v22
    }
    fmt.Println(v19, v3)
    v23 := GenS3{F0: (1 > -2147483648), F1: 1000, F2: (v3 >> ((331) & 63))}
    v1 = (!(!(3.14159 != (0.1 + ((3.14159 - ((0.1 * -2.25) * 1e10)) - 1e10)))))
    var v24 int = (v19 / ((v3) | 1))
    var v25 int8 = 58
    if (1.5 > (-2.25 / (1e10 - ((0.1 - (3.14159 / 1.0)) / (3.14159 - 1.5))))) {
        for fi26 := 0; fi26 < 19; fi26++ {
            v19 = -1
            var v27 int32 = -1
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            var v29 int32 = 2147483647
            _ = fi28
            _ = v29
        }
    }
    var v30 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
