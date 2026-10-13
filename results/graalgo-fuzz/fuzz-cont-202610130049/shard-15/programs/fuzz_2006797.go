package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 uint8; F2 int }
func genFn0(p0 int16, p1 int64) uint8 {
    p0 += (p0 + 756)
    v6 := []int{127, -444}
    p1 = -9223372036854775808
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v6
    return 127
}
func genFn1(p0 int32, p1 uint8) int32 {
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 255, F1: 2}
        var v9 int16 = 42
        v9 = (v9 * 32767)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: 2}
    p0 = (^(p0))
    switch "\n\t" {
    case "":
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS2{F0: 67}
            if false {
                v13 := GenS1{F0: 42, F1: 2147483647}
                if false {
                    switch "foo bar" {
                    case "\n\t":
                        v14 := []int{fi11, (fi11 & -9223372036854775808), 127, fi11}
                        _ = v14
                    default:
                        p1 = (p1 * (p1 * (p1 / ((genFn0(1000, 42)) | 1))))
                    }
                    if (((1.5 / -2.25) - 1.5) == (float32(1e10) + 3.14159)) {
                        v15 := GenS2{F0: (^(p1))}
                        var v16 uint64 = 1
                        _ = v15
                        _ = v16
                    } else {
                        switch 2147483647 {
                        case -2147483648:
                            v17 := GenS2{F0: 127}
                            _ = v17
                        case 127:
                            var v18 int16 = -32768
                            _ = v18
                        }
                        var v19 float32 = -2.25
                        _ = v19
                    }
                    var v20 string = "hello"
                    _ = v20
                }
                fmt.Println(p1, p0, p0)
                _ = v13
            } else {
                var v21 uint16 = 65535
                _ = v21
            }
            _ = fi11
            _ = v12
        }
    }
    var v22 uint = 314
    _ = p0
    _ = p1
    _ = v10
    _ = v22
    return p0
}

func main() {
    var v1 uint64 = 729
    var v2 uint32 = 0
    var v3 uint64 = (v1 * 18446744073709551615)
    v4 := []uint32{}
    var v5 uint = uint(v2)
    var v23 int = 9223372036854775807
    for fi24 := 0; fi24 < 7; fi24++ {
        if true {
            if true {
                switch (^(v23)) {
                case 9223372036854775807:
                    for fi25 := 0; fi25 < 12; fi25++ {
                        fmt.Println(v3, fi25)
                        _ = fi25
                    }
                case 153:
                    v26 := GenS3{F0: (v2 ^ v2)}
                    _ = v26
                case 32768:
                    if true {
                        v27 := GenS2{F0: 241}
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                        var v28 uint32 = (v2 & v2)
                        _ = v27
                        _ = v28
                    } else {
                        v29 := GenS1{F0: 322, F1: (v2 << ((v5) & 63))}
                        _ = v29
                    }
                }
            }
            v30 := GenS1{F0: 552, F1: 2}
            _ = v30
        }
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        _ = fi24
    }
    fmt.Println(v5, v3)
    v2 = uint32(v3)
    var v31 uint32 = v2
    v31 -= (v31 << ((v5) & 63))
    v32 := GenS4{F0: 1000, F1: genFn0(1000, 404), F2: 9223372036854775807}
    var v33 int64 = 9223372036854775807
    var v34 uint32 = (v2 / ((v31) | 1))
    if (((("" + "a") + "") + "foo bar") != ("" + "a")) {
        for fi35 := 0; fi35 < 18; fi35++ {
            v5 -= (v5 ^ (v5 << ((864) & 63)))
            _ = fi35
        }
    } else {
        v3 = v1
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if len(v4) > 0 {
        v4[0] = v34
    }
    if ("" == ("" + "a")) {
        _ = len(v4)
    }
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
