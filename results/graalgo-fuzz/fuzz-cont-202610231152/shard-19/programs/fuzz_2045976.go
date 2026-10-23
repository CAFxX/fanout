package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64 }

func main() {
    var v1 float32 = 1e10
    var v2 bool = false
    var v3 uint16 = 255
    v4 := []string{"a", ""}
    var v5 int8 = 0
    var v6 int = 133
    v7 := GenS1{F0: 127, F1: ("\n\t" + (("foo bar" + ("foo bar" + "\n\t")) + "hello")), F2: ""}
    if (0 <= 2) {
        v4 = append(v4, "")
        var v8 uint = 28
        _ = v8
    }
    var v9 uint16 = (v3 - (v3 | 42))
    var v10 int64 = 42
    v4 = append(v4, ("\n\t" + "hello"))
    v11 := GenS2{F0: 961}
    fmt.Println(v9, v3, v3)
    fmt.Println(v2, v5, v2)
    v5 = 58
    for fi12 := 0; fi12 < 7; fi12++ {
        if (125 >= 1) {
            var v13 uint16 = v9
            var v14 uint16 = 65535
            for fi15 := 0; fi15 < 5; fi15++ {
                if v2 {
                    v16 := GenS3{F0: 4294967295, F1: 83}
                    v3 = v9
                    v13 = 32767
                    _ = v16
                }
                var v17 int16 = 0
                if true {
                    switch ("" + ("\n\t" + "a")) {
                    case "a":
                        _ = len(v4)
                    case "\n\t":
                        v13 = (^(v13))
                    case "hello":
                        v14 -= (v14 * v13)
                    }
                } else {
                    var v18 uint32 = 65535
                    var v19 int64 = 848
                    _ = len(v4)
                    _ = v18
                    _ = v19
                }
                var v20 uint = 42
                _ = fi15
                _ = v17
                _ = v20
            }
            _ = v13
            _ = v14
        } else {
            for fi21 := 0; fi21 < 12; fi21++ {
                switch "hello" {
                case "foo bar":
                    v6 -= fi12
                case "":
                    v3 = v9
                case "hello":
                    switch fi12 {
                    case 42:
                        var v22 float32 = (1e10 / (float32(0.1) * float32(((float64((1e10 - (float32((3.14159 / ((0.0 / 1e10) + 1.5))) / v1))) - float64((float32(-2.25) * 1.5))) / 1e10))))
                        _ = len(v4)
                        _ = len(v4)
                        _ = v22
                    case 4294967295:
                        v10 += 42
                    case 262:
                        if true {
                            var v23 uint32 = 4294967295
                            fmt.Println(v5)
                            _ = v23
                        }
                    default:
                        if false {
                            if (!true) {
                                var v24 uint64 = 4294967295
                                var v25 int64 = v10
                                _ = v24
                                _ = v25
                            }
                            fmt.Println(v1)
                            v26 := GenS4{F0: v10}
                            _ = v26
                        } else {
                            v5 -= -1
                            if v2 {
                                v27 := GenS5{F0: v5, F1: "hello"}
                                if false {
                                    if len(v4) > 0 {
                                        v4[0] = ((("a" + "") + "\n\t") + "\n\t")
                                    }
                                }
                                _ = len(v4)
                                _ = v27
                            }
                            v4 = append(v4, ("foo bar" + ("foo bar" + "hello")))
                        }
                    }
                default:
                    if (!v2) {
                        v2 = (v2 || v2)
                        _ = len(v4)
                        if (!(-1 > (fi12 | v6))) {
                            v5 = v5
                            fmt.Println(v2, v10)
                            if false {
                                v4 = append(v4, (("foo bar" + "") + "\n\t"))
                            }
                        }
                    } else {
                        v28 := GenS6{F0: v5}
                        _ = v28
                    }
                }
                _ = fi21
            }
        }
        _ = fi12
    }
    var v29 float64 = -2.25
    v2 = v2
    switch v6 {
    case 42:
        v30 := GenS3{F0: 542, F1: 127}
        _ = v30
    default:
        var v31 int8 = 85
        _ = v31
    }
    fmt.Println(v5, v3, v2)
    var v32 string = (("" + (("foo bar" + "foo bar") + "hello")) + "")
    if v2 {
        for fi33 := 0; fi33 < 14; fi33++ {
            fmt.Println(v3, v9)
            _ = fi33
        }
        v34 := GenS6{F0: v5}
        v6 = (v6 - (v6 ^ v6))
        _ = v34
    } else {
        fmt.Println(v9)
        if (true && (v2 && (!false))) {
            if v2 {
                switch v6 {
                case 4294967295:
                    v35 := GenS7{F0: (v10 - (-(v10))), F1: 4294967295}
                    _ = v35
                case 128:
                    var v36 int = (v6 / (((v6 | (v6 >> ((80) & 63)))) | 1))
                    _ = v36
                case 0:
                    var v37 int = (v6 + 2147483647)
                    _ = v37
                default:
                    v38 := GenS8{F0: 184}
                    _ = v38
                }
            } else {
                if v2 {
                    fmt.Println(v5)
                    fmt.Println(v32, v1)
                    var v39 int8 = v5
                    _ = v39
                }
                v40 := GenS9{F0: (v10 & v10)}
                _ = v40
            }
        }
        var v41 float64 = float64((0.1 - v1))
        _ = v41
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = (v2 || (((1.5 > (v29 * v29)) || true) && v2))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
