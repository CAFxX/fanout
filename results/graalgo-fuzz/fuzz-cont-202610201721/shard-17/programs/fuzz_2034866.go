package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 string }
type GenS2 struct { F0 uint16; F1 int8; F2 bool }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 int32; F1 bool }
func genFn0(p0 int8, p1 int, p2 int) uint16 {
    fmt.Println(p2, p1, p2)
    if (65535 <= 431) {
        v7 := GenS1{F0: 0, F1: -583, F2: "hello"}
        _ = v7
    }
    p0 = (p0 % ((127) | 1))
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int8) int {
    v8 := GenS1{F0: -566, F1: 0, F2: "a"}
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 65535
    v3 := make(map[string]uint32)
    var v4 int64 = 9223372036854775807
    var v5 bool = (0.0 == 3.14159)
    var v6 string = "foo bar"
    v6 = "hello"
    var v9 int8 = int8(v2)
    _ = len(v1)
    if ((v5 && (v6 < (v6 + v6))) && false) {
        fmt.Println(v9, v2, v9)
        if (v5 && (!true)) {
            v3[v6] = 2147483648
            v10 := []int{1, 127}
            _ = v10
        } else {
            v11 := GenS2{F0: v2, F1: int8(v4), F2: v5}
            switch (v6 + ((v6 + v6) + ((((v6 + "\n\t") + (v6 + v6)) + "") + "hello"))) {
            case "foo bar":
                v12 := GenS2{F0: 255, F1: v9, F2: v5}
                _ = v12
            case "a":
                if v5 {
                    v4 = v4
                } else {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        for fi14 := 0; fi14 < 13; fi14++ {
                            v15 := GenS3{F0: 56, F1: genFn0(123, fi14, fi13)}
                            v3[(v6 + ((v6 + "hello") + "\n\t"))] = 2147483647
                            _ = fi14
                            _ = v15
                        }
                        _ = fi13
                    }
                    v4 -= v4
                }
            default:
                if v5 {
                    switch v6 {
                    case "hello":
                        fmt.Println(v6)
                    case "foo bar":
                        for fi16 := 0; fi16 < 13; fi16++ {
                            if v5 {
                                v6 = (v6 + v6)
                            }
                            v4 = v4
                            var v17 float32 = 1.5
                            _ = fi16
                            _ = v17
                        }
                    }
                    if false {
                        switch v6 {
                        case "\n\t":
                            v4 -= (v4 | (v4 | (v4 - (-(v4)))))
                        case "":
                            if false {
                                if (!v5) {
                                    fmt.Println(v6, v5)
                                    fmt.Println(v2)
                                }
                                v5 = v5
                                v18 := GenS4{F0: 1}
                                _ = v18
                            }
                        case "hello":
                            fmt.Println(v9, v2, v2)
                        }
                        fmt.Println(v9)
                        var v19 uint32 = 1
                        _ = v19
                    } else {
                        v20 := []int{-728}
                        _ = v20
                    }
                    var v21 float64 = -2.25
                    _ = v21
                }
            }
            _ = v11
        }
    }
    fmt.Println(v6)
    if v5 {
        if (!(-403 != 32767)) {
            switch 42 {
            case 32768:
                v5 = (v4 <= (v4 ^ v4))
            case 2147483647:
                v2 -= v2
            case 4294967295:
                var v22 int = 4294967295
                _ = v22
            default:
                v23 := []int{2147483647, int(v2), 646, 2147483647}
                _ = v23
            }
        } else {
            v9 = (v9 - (v9 >> ((1) & 63)))
        }
        v9 += -1
        if v5 {
            for fi24 := 0; fi24 < 2; fi24++ {
                v6 = "hello"
                v6 = v6
                _ = fi24
            }
        } else {
            if v5 {
                v9 -= v9
                if v5 {
                    var v25 int8 = v9
                    v4 += (v4 / ((-457) | 1))
                    delete(v3, v6)
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 13; fi26++ {
                        if (0.0 < 0.0) {
                            var v27 float64 = (float64(1e10) + (3.14159 - 0.0))
                            _ = v27
                        } else {
                            fmt.Println(v6, fi26)
                        }
                        fmt.Println(v4)
                        _ = fi26
                    }
                }
                v2 -= (v2 % (((v2 | (v2 * 0))) | 1))
            }
            fmt.Println(v6, v4, v4)
            var v28 bool = (42 == -751)
            _ = v28
        }
    }
    var v29 uint64 = 42
    v30 := GenS2{F0: 255, F1: 1, F2: true}
    var v31 int = 49
    switch ("a" + ((v6 + ((("hello" + "\n\t") + "a") + v6)) + "")) {
    case "":
        var v32 int8 = v9
        _ = v32
    case "foo bar":
        for fi33 := 0; fi33 < 8; fi33++ {
            v6 = "hello"
            _ = fi33
        }
    default:
        v2 += genFn0((-(v9)), (v31 * v31), (v31 - (v31 - (v31 - v31))))
    }
    var v34 int64 = (v4 & -59)
    for fi35 := 0; fi35 < 12; fi35++ {
        for fi36 := 0; fi36 < 1; fi36++ {
            if (v5 || v5) {
                v37 := GenS1{F0: -1, F1: v34, F2: v6}
                v38 := GenS5{F0: 18446744073709551615, F1: int32(v29)}
                _ = v37
                _ = v38
            }
            if ("hello" != v6) {
                if v5 {
                    if (42 != 0) {
                        v39 := []int{32768, fi35, (-(fi36)), genFn1(v9)}
                        var v40 float32 = 1e10
                        switch (v31 | fi35) {
                        case -311:
                            v2 = 255
                        case 4294967295:
                            v41 := GenS6{F0: -1, F1: (!false)}
                            _ = v41
                        default:
                            if len(v39) > 0 {
                                v39[0] = (fi36 - (fi36 - -1))
                            }
                        }
                        _ = v39
                        _ = v40
                    }
                    v42 := GenS4{F0: 1}
                    v43 := []int{genFn1(-116), (fi35 % ((fi36) | 1)), int(v2), (fi36 >> ((42) & 63))}
                    _ = v42
                    _ = v43
                }
            }
            _ = fi36
        }
        var v44 bool = false
        var v45 string = v6
        _ = fi35
        _ = v44
        _ = v45
    }
    v5 = ((v2 / ((v2) | 1)) < (^(v2)))
    var v46 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v9
}
