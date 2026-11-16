package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int64; F1 uint32; F2 string }
type GenS4 struct { F0 int32; F1 int32; F2 int8 }
type GenS5 struct { F0 int64 }
func genFn0(p0 string) uint32 {
    v6 := GenS1{F0: 223, F1: (749 > 2)}
    var v7 uint16 = 32767
    _ = p0
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 int, p1 uint) uint32 {
    p1 -= 2
    var v8 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v8
    return 319
}

func main() {
    v1 := []uint32{2}
    v2 := make(map[int]bool)
    var v3 uint64 = 42
    v4 := []uint32{588, 250, 409, 4294967295}
    var v5 uint8 = 127
    _ = len(v2)
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS2{F0: (v3 ^ v3), F1: -12}
        switch "foo bar" {
        case "hello":
            v11 := GenS1{F0: 2147483647, F1: false}
            _ = v11
        case "":
            fmt.Println(v3)
        case "a":
            v5 = 42
        default:
            fmt.Println(v3, v3, v5)
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v3, v3)
    if (42 <= -244) {
        if true {
            for fi12 := 0; fi12 < 19; fi12++ {
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := v2[32767]
                    v15 := GenS2{F0: 418, F1: 100}
                    fmt.Println(v14, v14, v14)
                    if ((true || (v14 || ((v3 > v3) && true))) || (true && false)) {
                        var v16 int32 = -2147483648
                        v17 := GenS1{F0: v16, F1: (false && (!false))}
                        _ = v16
                        _ = v17
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                switch (fi12 / ((fi12) | 1)) {
                case 42:
                    delete(v2, (fi12 >> ((975) & 63)))
                }
                _ = fi12
            }
        } else {
            switch (("foo bar" + "foo bar") + "foo bar") {
            case "":
                switch "foo bar" {
                case "\n\t":
                    v18 := GenS3{F0: 577, F1: 65535, F2: "a"}
                    _ = v18
                case "":
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v5 -= v5
                        _ = fi19
                    }
                case "a":
                    v5 = (^(v5))
                default:
                    var v20 uint64 = (v3 / ((v3) | 1))
                    _ = v20
                }
            case "a":
                for fi21 := 0; fi21 < 8; fi21++ {
                    _ = len(v2)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if false {
                        var v22 float32 = 1.5
                        var v23 string = "\n\t"
                        v3 += uint64(v5)
                        _ = v22
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 6; fi24++ {
                            var v25 float64 = 0.1
                            v3 = v3
                            _ = fi24
                            _ = v25
                        }
                    }
                    _ = fi21
                }
            case "\n\t":
                v26 := GenS4{F0: 884, F1: -2147483648, F2: -37}
                _ = v26
            default:
                v27 := v2[0]
                _ = v27
            }
            switch 37 {
            case 138:
                v28 := GenS5{F0: -9223372036854775808}
                _ = v28
            }
            v5 -= (^(v5))
        }
    } else {
        var v29 float32 = float32(float64(0.0))
        _ = v29
    }
    var v30 uint64 = v3
    fmt.Println(v30)
    v5 -= v5
    var v31 uint = 99
    var v32 int32 = -1
    v31 += (^(v31))
    var v33 uint8 = (v5 % (((^(v5))) | 1))
    v30 -= (v3 * uint64(v31))
    v32 -= v32
    var v34 uint = v31
    v30 = (^(v3))
    if false {
        if false {
            if true {
                for fi35 := 0; fi35 < 2; fi35++ {
                    for fi36 := 0; fi36 < 5; fi36++ {
                        var v37 int32 = v32
                        _ = fi36
                        _ = v37
                    }
                    _ = fi35
                }
            } else {
                if (false && true) {
                    if true {
                        v5 += (v5 * 8)
                    } else {
                        var v38 uint64 = (v30 | 42)
                        _ = v38
                    }
                    v3 += v30
                } else {
                    delete(v2, 664)
                    fmt.Println(v32, v30)
                    if (false || false) {
                        if false {
                            var v39 int = 32767
                            _ = v39
                        }
                        var v40 int8 = 42
                        v3 += 4294967295
                        _ = v40
                    } else {
                        fmt.Println(v3, v31, v3)
                        for fi41 := 0; fi41 < 13; fi41++ {
                            var v42 uint64 = v30
                            _ = fi41
                            _ = v42
                        }
                    }
                }
                var v43 uint64 = (v30 % ((v3) | 1))
                _ = v43
            }
        }
        v44 := GenS2{F0: v3, F1: 1}
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
