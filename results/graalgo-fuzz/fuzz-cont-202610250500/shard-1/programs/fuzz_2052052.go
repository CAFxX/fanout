package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int16; F1 int64; F2 int64 }
func genFn0() int {
    var v4 string = "foo bar"
    v5 := []int{-9223372036854775808}
    _ = v4
    _ = v5
    return 2147483647
}
func genFn1(p0 int16) int64 {
    v6 := GenS1{F0: 1000}
    p0 = (^(p0))
    v7 := GenS2{F0: -128, F1: -13}
    _ = p0
    _ = v6
    _ = v7
    return 10
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[int]uint32)
    v3 := []string{("\n\t" + ""), "hello"}
    if (4294967295 > 65535) {
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(fi8, v1)
            _ = fi8
        }
        v1 = 908
        if true {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 int = 32768
                v1 += 4294967295
                _ = fi9
                _ = v10
            }
        } else {
            v1 += 42
            if (2 < 888) {
                var v11 int8 = 100
                switch "" {
                case "a":
                    if len(v3) > 0 {
                        v3[0] = (("a" + "foo bar") + "a")
                    }
                default:
                    v12 := GenS1{F0: 1000}
                    _ = v12
                }
                _ = v11
            } else {
                v2[32767] = 42
            }
            if ((!(!(285 != 4294967295))) && false) {
                for fi13 := 0; fi13 < 14; fi13++ {
                    if (v1 == (v1 % (((v1 - (v1 ^ (v1 / ((v1) | 1))))) | 1))) {
                        v1 = (v1 + 0)
                        for fi14 := 0; fi14 < 7; fi14++ {
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                            v15 := GenS2{F0: -128, F1: 0}
                            _ = fi14
                            _ = v15
                        }
                    } else {
                        if false {
                            fmt.Println(v1)
                            v1 -= 730
                        } else {
                            v16 := GenS3{F0: (^(fi13))}
                            _ = v16
                        }
                        fmt.Println(fi13, fi13)
                        var v17 int8 = 127
                        _ = v17
                    }
                    v18 := GenS4{F0: 541, F1: 2, F2: 42}
                    v1 += v1
                    _ = fi13
                    _ = v18
                }
                var v19 float32 = (0.1 * 0.0)
                _ = v19
            } else {
                switch 42 {
                case 9223372036854775807:
                    fmt.Println(v1)
                case -269:
                    v1 -= v1
                }
                switch 32768 {
                case 32768:
                    v20 := GenS5{F0: 32767, F1: -9223372036854775808, F2: -5}
                    _ = v20
                case -895:
                    var v21 int8 = -128
                    _ = v21
                case 42:
                    var v22 bool = ((0.1 / 3.14159) <= 1e10)
                    _ = v22
                }
                var v23 int32 = 1
                _ = v23
            }
        }
    }
    fmt.Println(v1)
    if false {
        delete(v2, -1)
    } else {
        v1 += 1
        v1 -= v1
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 += (^(v1))
    if true {
        v24 := GenS4{F0: 2147483648, F1: 2147483647, F2: 179}
        _ = v24
    } else {
        var v25 int16 = 0
        v1 -= v1
        _ = v25
    }
    if true {
        if false {
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 uint16 = 539
                v27 += v27
                fmt.Println(v27, v27)
                _ = fi26
                _ = v27
            }
            if true {
                var v28 int8 = -36
                _ = v28
            } else {
                if false {
                    if false {
                        v29 := GenS1{F0: 32767}
                        var v30 int8 = -1
                        v30 = (v30 - v30)
                        _ = v29
                        _ = v30
                    }
                    fmt.Println(v1)
                } else {
                    var v31 string = ("a" + ("\n\t" + "hello"))
                    _ = v31
                }
            }
        } else {
            if len(v3) > 0 {
                v3[0] = ((("" + "foo bar") + "hello") + "a")
            }
            switch genFn0() {
            case -165:
                for fi32 := 0; fi32 < 3; fi32++ {
                    switch ("" + "a") {
                    case "":
                        _ = len(v3)
                    }
                    v33 := GenS1{F0: -32768}
                    _ = fi32
                    _ = v33
                }
            case 1:
                var v34 uint = 573
                _ = v34
            case -766:
                var v35 uint8 = 1
                _ = v35
            default:
                var v36 float32 = 0.1
                _ = v36
            }
            if false {
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    v37 := GenS4{F0: 2147483648, F1: 42, F2: 1}
                    _ = v37
                case "":
                    v1 -= 18446744073709551615
                }
                v38 := v2[127]
                _ = v38
            } else {
                fmt.Println(v1)
            }
        }
    } else {
        fmt.Println(v1)
        v1 = v1
    }
    switch (("" + "") + "hello") {
    case "\n\t":
        _ = len(v3)
    case "foo bar":
        var v39 uint64 = (v1 & v1)
        _ = v39
    case "hello":
        var v40 int64 = 0
        _ = v40
    }
    for fi41 := 0; fi41 < 7; fi41++ {
        v3 = append(v3, "foo bar")
        fmt.Println(v1)
        _ = fi41
    }
    v42 := GenS3{F0: 128}
    v3 = append(v3, (("a" + "\n\t") + ("hello" + "foo bar")))
    var v43 bool = (-2147483648 < 1)
    for fi44 := 0; fi44 < 2; fi44++ {
        v45 := GenS3{F0: genFn0()}
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 2; fi46++ {
        v43 = (9223372036854775807 != 9223372036854775807)
        v47 := GenS5{F0: 32767, F1: -1, F2: -263}
        v1 -= (v1 - (v1 << ((14) & 63)))
        _ = fi46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
}
