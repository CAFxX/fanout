package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int; F1 bool; F2 string }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 int8; F1 int32; F2 uint }
func genFn0() bool {
    var v4 bool = (!true)
    v4 = (!(1.5 >= (3.14159 * 1e10)))
    _ = v4
    return v4
}

func main() {
    var v1 uint32 = 1
    v2 := []int{32768, 450, 32767}
    var v3 uint64 = 18446744073709551615
    var v5 int8 = 1
    var v6 float32 = (1e10 * (3.14159 * 0.1))
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v1 = v1
    default:
        _ = len(v2)
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 154, F1: -1}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 float32 = v6
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v5 += (v5 << ((65535) & 63))
            var v12 float64 = (1.5 + 1e10)
            _ = fi10
            _ = v11
            _ = v12
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi7
    }
    v13 := GenS1{F0: 60, F1: 9223372036854775807}
    var v14 uint = 662
    var v15 bool = (v14 == v14)
    switch 128 {
    case 2147483647:
        var v16 uint64 = v3
        _ = v16
    case 1:
        v1 += 2147483648
    default:
        if (v14 == v14) {
            v6 += v6
        } else {
            var v17 int16 = -665
            switch -988 {
            case 0:
                if (!false) {
                    switch "" {
                    case "a":
                        fmt.Println(v14, v6, v3)
                    case "hello":
                        for fi18 := 0; fi18 < 16; fi18++ {
                            if (true && (int16(v14) != v17)) {
                                if ((!v15) || v15) {
                                    var v19 float64 = (((0.0 + 1.5) - ((-2.25 / 1.0) + 1.5)) + 0.0)
                                    var v20 int16 = v17
                                    v14 = v14
                                    _ = v19
                                    _ = v20
                                } else {
                                    _ = len(v2)
                                    for fi21 := 0; fi21 < 14; fi21++ {
                                        v15 = (!(("a" <= "") && false))
                                        _ = fi21
                                    }
                                }
                                v5 += v5
                                if (!((3.14159 + (1e10 + float64((v6 * v6)))) == 1.5)) {
                                    var v22 uint32 = 42
                                    var v23 int = (fi18 | -744)
                                    if (false && (int16(fi18) > (v17 & v17))) {
                                        v17 += 32767
                                        for fi24 := 0; fi24 < 13; fi24++ {
                                            v14 += v14
                                            v2 = append(v2, int(v5))
                                            _ = fi24
                                        }
                                        v25 := GenS1{F0: 1, F1: 9223372036854775807}
                                        _ = v25
                                    }
                                    _ = v22
                                    _ = v23
                                } else {
                                    v2 = append(v2, 607)
                                    v26 := GenS2{F0: fi18, F1: true, F2: (("\n\t" + ("foo bar" + "hello")) + "")}
                                    _ = v26
                                }
                            } else {
                                v27 := GenS1{F0: 42, F1: -1}
                                if len(v2) > 0 {
                                    v2[0] = -538
                                }
                                if (v6 > v6) {
                                    var v28 float32 = (v6 * ((0.0 * (((3.14159 * 1e10) + 1.5) / 0.1)) * v6))
                                    var v29 uint8 = 23
                                    v3 = (v3 >> ((500) & 63))
                                    _ = v28
                                    _ = v29
                                }
                                _ = v27
                            }
                            _ = fi18
                        }
                    default:
                        var v30 uint8 = uint8(v1)
                        _ = v30
                    }
                    v31 := GenS3{F0: v5, F1: -197}
                    _ = v31
                }
            case 4294967295:
                if (0.0 < -2.25) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v15 = (v15 || false)
                }
            }
            if (v15 || (767 >= 65535)) {
                var v32 float32 = 3.14159
                for fi33 := 0; fi33 < 11; fi33++ {
                    v6 += v6
                    v5 += (-(v5))
                    _ = fi33
                }
                _ = v32
            }
            _ = v17
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v34 float64 = ((-2.25 * 1e10) + 1e10)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!(126 > 128)) {
        _ = len(v2)
        fmt.Println(v5, v34)
    } else {
        v35 := GenS1{F0: 1, F1: 42}
        _ = v35
    }
    if len(v2) > 0 {
        v2[0] = 32767
    }
    v36 := GenS4{F0: (-(v5)), F1: 2147483647, F2: 18446744073709551615}
    if v15 {
        v34 += ((3.14159 * v34) * v34)
    } else {
        v14 -= (v14 | 1)
    }
    var v37 int32 = 173
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v37
    _ = v5
    _ = v6
}
