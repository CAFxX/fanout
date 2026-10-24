package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 int8; F1 string; F2 uint64 }
func genFn0(p0 int64, p1 string) int32 {
    v5 := GenS1{F0: 706, F1: 4294967295, F2: 382}
    p0 -= (p0 & 424)
    v6 := []int{32767, -905, 0}
    switch 2147483647 {
    case 0:
        fmt.Println(p1, p0)
    case 32767:
        var v7 int = 32768
        _ = v7
    case 9223372036854775807:
        if len(v6) > 0 {
            v6[0] = 1
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 float32 = 3.14159
    var v2 int64 = -215
    var v3 uint8 = 1
    v4 := []uint32{4294967295, 195, 874, 65535}
    var v8 uint64 = 42
    if false {
        v9 := GenS1{F0: (v2 ^ v2), F1: 129, F2: 147}
        _ = v9
    } else {
        if (!(65535 <= 79)) {
            var v10 uint8 = 127
            v1 += 3.14159
            _ = v10
        } else {
            var v11 float64 = (-2.25 - 0.1)
            v8 = 18446744073709551615
            v3 -= v3
            _ = v11
        }
        v12 := GenS1{F0: (v2 & v2), F1: 616, F2: 1}
        for fi13 := 0; fi13 < 15; fi13++ {
            _ = len(v4)
            v8 = v8
            fmt.Println(v3, fi13, fi13)
            _ = fi13
        }
        _ = v12
    }
    v1 = -2.25
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v14 := GenS1{F0: (v2 % ((v2) | 1)), F1: 2147483648, F2: 439}
    v2 += (v2 & (v2 % (((v2 >> ((4294967295) & 63))) | 1)))
    v15 := GenS2{F0: 999}
    var v16 uint32 = 340
    for fi17 := 0; fi17 < 16; fi17++ {
        v3 = 1
        _ = fi17
    }
    var v18 bool = true
    v16 -= (v16 * 4294967295)
    _ = len(v4)
    if (65535 < 2) {
        switch "foo bar" {
        case "a":
            v16 = (v16 / ((1) | 1))
        }
        fmt.Println(v1, v8)
        var v19 uint8 = v3
        _ = v19
    } else {
        if v18 {
            v3 = (^(v3))
        } else {
            var v20 int32 = 1
            if (36 <= 32767) {
                for fi21 := 0; fi21 < 8; fi21++ {
                    v8 -= 1
                    _ = fi21
                }
                for fi22 := 0; fi22 < 19; fi22++ {
                    if (!v18) {
                        switch fi22 {
                        case 503:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case 0:
                            var v23 float64 = float64(v1)
                            _ = v23
                        case 128:
                            v1 = v1
                        }
                        if len(v4) > 0 {
                            v4[0] = (v16 ^ v16)
                        }
                    }
                    _ = fi22
                }
            } else {
                for fi24 := 0; fi24 < 1; fi24++ {
                    if len(v4) > 0 {
                        v4[0] = 919
                    }
                    var v25 float64 = (3.14159 * 1.5)
                    _ = fi24
                    _ = v25
                }
                v26 := GenS1{F0: v2, F1: v16, F2: 42}
                fmt.Println(v8)
                _ = v26
            }
            _ = v20
        }
        if (v1 > v1) {
            fmt.Println(v3)
            fmt.Println(v8)
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 uint = 42
                if true {
                    if true {
                        var v29 uint16 = 2
                        v1 = (float32(0.1) + v1)
                        _ = v29
                    } else {
                        v4 = append(v4, (v16 * v16))
                        fmt.Println(v1)
                        var v30 string = ""
                        _ = v30
                    }
                    var v31 float32 = v1
                    v32 := GenS3{F0: uint32(v8)}
                    _ = v31
                    _ = v32
                } else {
                    switch fi27 {
                    case 32768:
                        var v33 uint32 = (^(v16))
                        _ = v33
                    case 923:
                        var v34 int16 = 1000
                        _ = v34
                    case -1:
                        var v35 float64 = 3.14159
                        _ = v35
                    default:
                        v2 -= v2
                    }
                    v36 := GenS4{F0: 42, F1: (v3 << ((v28) & 63))}
                    fmt.Println(v2, v28)
                    _ = v36
                }
                for fi37 := 0; fi37 < 3; fi37++ {
                    var v38 int = int(v16)
                    var v39 int64 = 0
                    v8 = (v8 >> ((4294967295) & 63))
                    v40 := GenS5{F0: -128, F1: ("hello" + ("" + ("foo bar" + "hello"))), F2: 380}
                    switch "foo bar" {
                    case "hello":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case "\n\t":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    default:
                        if ((((0.0 <= 3.14159) || (!(-32768 == -32768))) && v18) && ((v28 != (v28 | (v28 & 1))) || (2147483648 == v16))) {
                            var v41 int8 = 127
                            switch (fi27 | v38) {
                            case 249:
                                v28 -= (^(v28))
                            }
                            _ = v41
                        } else {
                            fmt.Println(v3)
                            if ((!((float64((v1 + (1.5 * 0.0))) - (3.14159 * 1.5)) != 3.14159)) || ((v1 < v1) && v18)) {
                                v4 = append(v4, v16)
                                fmt.Println(fi27)
                                v38 = -9223372036854775808
                            } else {
                                fmt.Println(v8, v2, fi27)
                                var v42 uint8 = (^(v3))
                                _ = v42
                            }
                            if (1e10 != 1.5) {
                                v39 = v2
                                if len(v4) > 0 {
                                    v4[0] = v16
                                }
                            } else {
                                if (!(!(!v18))) {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    fmt.Println(v28, v8, v28)
                                    var v43 uint64 = uint64(v16)
                                    fmt.Println(fi27)
                                    _ = v43
                                }
                                var v44 int8 = -1
                                _ = v44
                            }
                            switch "hello" {
                            case "\n\t":
                                var v45 bool = true
                                _ = v45
                            }
                            var v46 int32 = genFn0(-9223372036854775808, "hello")
                            _ = v46
                        }
                    }
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = fi27
                _ = v28
            }
        } else {
            switch 128 {
            case 0:
                var v47 int32 = 958
                _ = v47
            case 32767:
                switch -12 {
                case -1:
                    v1 -= (v1 + v1)
                }
            case 42:
                v3 = v3
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
