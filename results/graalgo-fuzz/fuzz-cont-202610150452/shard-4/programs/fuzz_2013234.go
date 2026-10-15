package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 int8 }
func genFn0() int {
    var v4 int32 = -1
    v5 := GenS1{F0: 42, F1: 2}
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 uint16, p1 int, p2 uint8) uint16 {
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint64 = 331
        _ = fi6
        _ = v7
    }
    fmt.Println(p1, p0)
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p0
}

func main() {
    var v1 uint8 = 207
    var v2 int8 = 127
    v3 := []string{(((("" + ("hello" + (((("\n\t" + ((("hello" + ("foo bar" + "hello")) + "") + ("" + ("" + "\n\t")))) + "\n\t") + "foo bar") + "hello"))) + "a") + "\n\t") + ""), "", ((("hello" + ("" + "a")) + "") + "hello"), "hello"}
    v2 += 127
    v9 := GenS1{F0: 9223372036854775807, F1: 243}
    fmt.Println(v1, v2)
    if ((((1.5 - 1.5) - float32(0.1)) < -2.25) || (true || (float32((0.1 * 3.14159)) < 1e10))) {
        if (42 >= 42) {
            var v10 float64 = (((1e10 / ((-2.25 - float64((-2.25 - 1e10))) - 0.1)) - ((-2.25 + float64(1.5)) + 1.5)) + float64((3.14159 + 1.5)))
            _ = v10
        } else {
            for fi11 := 0; fi11 < 4; fi11++ {
                if true {
                    for fi12 := 0; fi12 < 6; fi12++ {
                        switch (fi12 - fi12) {
                        case -419:
                            var v13 int8 = v2
                            _ = v13
                        case 1:
                            if false {
                                v14 := GenS2{F0: 255, F1: 1, F2: 2147483647}
                                v1 = (v1 & (v1 / (((v1 % (((v1 / (((v1 ^ (v1 + v1))) | 1))) | 1))) | 1)))
                                fmt.Println(fi12, v2)
                                _ = v14
                            }
                        case 32767:
                            v15 := GenS2{F0: 876, F1: 255, F2: 85}
                            _ = v15
                        }
                        var v16 uint = 18446744073709551615
                        switch "hello" {
                        case "a":
                            var v17 uint16 = 0
                            _ = v17
                        case "hello":
                            v2 += 42
                        case "foo bar":
                            v1 = 127
                        }
                        v3 = append(v3, "foo bar")
                        _ = fi12
                        _ = v16
                    }
                    for fi18 := 0; fi18 < 11; fi18++ {
                        var v19 int64 = 42
                        v3 = append(v3, "a")
                        v19 += (-(v19))
                        _ = fi18
                        _ = v19
                    }
                } else {
                    v2 = v2
                    v2 -= 102
                    v1 -= v1
                }
                var v20 int16 = -418
                _ = fi11
                _ = v20
            }
            v1 -= v1
            v1 = v1
        }
        if (true || (42 >= genFn1(0, -29, (^(v1))))) {
            for fi21 := 0; fi21 < 4; fi21++ {
                fmt.Println(v1, v2, v1)
                var v22 int = (fi21 >> ((4294967295) & 63))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi21
                _ = v22
            }
        }
        v23 := GenS1{F0: 9223372036854775807, F1: 1}
        _ = v23
    } else {
        for fi24 := 0; fi24 < 4; fi24++ {
            if ((9223372036854775807 != -1) || (572 >= 919)) {
                v1 = (v1 * v1)
                v2 -= (v2 & (v2 ^ v2))
                v1 -= 146
            } else {
                var v25 int64 = 9223372036854775807
                _ = v25
            }
            _ = fi24
        }
        for fi26 := 0; fi26 < 3; fi26++ {
            v2 += 1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi26
        }
    }
    if (!false) {
        v2 = v2
        var v27 int8 = v2
        _ = v27
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        v29 := GenS1{F0: -964, F1: 1}
        _ = fi28
        _ = v29
    }
    if false {
        v2 = (^(v2))
    }
    if false {
        var v30 int16 = 1
        var v31 int16 = 547
        v32 := GenS3{F0: 589}
        _ = v30
        _ = v31
        _ = v32
    }
    var v33 int8 = -128
    var v34 uint = 42
    switch "foo bar" {
    case "a":
        for fi35 := 0; fi35 < 4; fi35++ {
            v33 = -128
            _ = fi35
        }
    default:
        v36 := GenS2{F0: genFn1(2, 9223372036854775807, v1), F1: 65535, F2: 427}
        _ = v36
    }
    var v37 int = -97
    v34 = (v34 / (((v34 % ((0) | 1))) | 1))
    var v38 int64 = 630
    v38 = v38
    var v39 uint64 = 311
    switch "\n\t" {
    case "":
        v38 += v38
    case "\n\t":
        v39 += (v39 - (v39 | 2))
    case "a":
        v40 := GenS4{F0: -695, F1: int8(v34)}
        _ = v40
    }
    v41 := GenS3{F0: genFn1(255, (v37 << (((v34 - 4294967295)) & 63)), v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v41
    _ = v9
}
