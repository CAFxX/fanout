package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint64; F2 int32 }
type GenS4 struct { F0 int8; F1 int }
func genFn0(p0 int) uint64 {
    v6 := []int{p0, p0, 32768}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := []string{"hello", (("foo bar" + "hello") + ("foo bar" + "foo bar"))}
    var v2 float64 = 3.14159
    var v3 uint32 = 442
    var v4 uint32 = (^(v3))
    var v5 string = "\n\t"
    for fi7 := 0; fi7 < 13; fi7++ {
        switch fi7 {
        case 4294967295:
            v3 = (v4 & (v3 >> ((979) & 63)))
        case 42:
            var v8 float64 = v2
            _ = v8
        default:
            v9 := GenS1{F0: 2147483647}
            _ = v9
        }
        if ((false || (true && (127 <= 71))) && (!true)) {
            if (-1 != -1) {
                if true {
                    v10 := GenS1{F0: 1}
                    switch "foo bar" {
                    case "hello":
                        switch v5 {
                        case "a":
                            var v11 uint32 = (^(v4))
                            _ = v11
                        case "hello":
                            fmt.Println(v3, v2)
                        }
                    case "\n\t":
                        switch "\n\t" {
                        case "":
                            switch fi7 {
                            case 1:
                                var v12 float32 = 1.5
                                _ = v12
                            }
                        default:
                            v3 -= 304
                        }
                    case "foo bar":
                        v13 := GenS2{F0: 42, F1: (v3 | v3)}
                        _ = v13
                    default:
                        v3 -= 1
                    }
                    var v14 int = (fi7 + fi7)
                    _ = v10
                    _ = v14
                }
            }
        }
        _ = fi7
    }
    v1 = append(v1, "hello")
    var v15 int = 774
    switch v15 {
    case 9223372036854775807:
        if false {
            var v16 float64 = 0.1
            v1 = append(v1, (v5 + v5))
            for fi17 := 0; fi17 < 11; fi17++ {
                v18 := GenS3{F0: 0, F1: 649, F2: -763}
                for fi19 := 0; fi19 < 18; fi19++ {
                    var v20 uint16 = 255
                    var v21 uint64 = 4294967295
                    var v22 int8 = -1
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                _ = fi17
                _ = v18
            }
            _ = v16
        } else {
            v1 = append(v1, v5)
        }
    default:
        v2 = v2
    }
    v23 := GenS4{F0: 100, F1: 2147483647}
    v3 -= v4
    v15 -= (v15 - (v15 | (v15 + -987)))
    switch v5 {
    case "":
        if len(v1) > 0 {
            v1[0] = v5
        }
    case "hello":
        _ = len(v1)
    case "\n\t":
        v24 := GenS2{F0: v3, F1: (v3 * v4)}
        _ = v24
    default:
        v5 = v5
    }
    var v25 int32 = 42
    switch (v15 >> ((65535) & 63)) {
    case 1:
        if true {
            v15 -= v15
        }
    case 128:
        v1 = append(v1, (v5 + v5))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
