package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int64; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int8 }
func genFn0(p0 string, p1 uint8, p2 int64) int64 {
    fmt.Println(p1)
    p1 = 149
    if (32767 < 65535) {
        p1 = (p1 ^ (p1 ^ 2))
        v5 := []int{0}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return (^(p2))
}

func main() {
    v1 := []uint32{2147483648}
    var v2 bool = (!(4294967295 == 4294967295))
    var v3 uint = 1
    v4 := make(map[string]bool)
    var v6 int64 = -262
    v7 := GenS1{F0: 69}
    _ = len(v4)
    switch "\n\t" {
    case "foo bar":
        switch (("" + "\n\t") + "hello") {
        case "a":
            if v2 {
                v6 += -9
            }
        case "foo bar":
            if ((v6 - (v6 / ((-1) | 1))) <= v6) {
                if true {
                    v8 := GenS2{F0: v6, F1: 573, F2: 4294967295}
                    _ = v8
                } else {
                    v9 := GenS2{F0: v6, F1: 70, F2: 0}
                    v1 = append(v1, 1)
                    v10 := GenS1{F0: 93}
                    _ = v10
                    _ = v9
                }
            }
        default:
            if false {
                if true {
                    fmt.Println(v3, v6)
                    switch "" {
                    case "a":
                        switch 9223372036854775807 {
                        case 4294967295:
                            if (32767 != 697) {
                                v11 := GenS3{F0: v6, F1: -26, F2: 2}
                                var v12 float64 = float64(((-2.25 / 0.1) / 3.14159))
                                _ = v11
                                _ = v12
                            } else {
                                v1 = append(v1, 4294967295)
                                var v13 float64 = ((1.5 + 1.5) / 1.5)
                                _ = v13
                            }
                        }
                    case "\n\t":
                        for fi14 := 0; fi14 < 5; fi14++ {
                            v2 = v2
                            if len(v1) > 0 {
                                v1[0] = uint32(v3)
                            }
                            _ = fi14
                        }
                    }
                }
                if (230 >= 32767) {
                    v3 = (v3 % ((42) | 1))
                    if (42 > 222) {
                        for fi15 := 0; fi15 < 8; fi15++ {
                            var v16 uint = v3
                            _ = fi15
                            _ = v16
                        }
                        var v17 uint16 = 395
                        var v18 float32 = float32(-2.25)
                        _ = v17
                        _ = v18
                    }
                    if (v2 || v2) {
                        var v19 uint32 = 751
                        var v20 uint64 = 2
                        v2 = v2
                        _ = v19
                        _ = v20
                    } else {
                        var v21 int8 = -76
                        switch 4294967295 {
                        case 0:
                            for fi22 := 0; fi22 < 5; fi22++ {
                                if ("foo bar" != "\n\t") {
                                    v2 = true
                                    var v23 bool = false
                                    _ = v23
                                }
                                _ = fi22
                            }
                        case 4294967295:
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(v2, v3, v3)
                            case "foo bar":
                                fmt.Println(v6)
                            default:
                                v2 = (128 != 0)
                            }
                        default:
                            v24 := GenS4{F0: 65535, F1: v21}
                            _ = v24
                        }
                        v3 = 4294967295
                        _ = v21
                    }
                } else {
                    if (!((1e10 / 1e10) <= 0.1)) {
                        if (!(!(-390 < -1))) {
                            var v25 uint16 = 32767
                            for fi26 := 0; fi26 < 16; fi26++ {
                                if len(v1) > 0 {
                                    v1[0] = 2147483647
                                }
                                if (-1 == 42) {
                                    v3 = 255
                                    var v27 uint32 = 2147483648
                                    var v28 bool = false
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 uint32 = 874
                                    v29 = v29
                                    fmt.Println(fi26, v2, v29)
                                    _ = v29
                                }
                                v4[("foo bar" + "")] = true
                                _ = fi26
                            }
                            _ = v25
                        } else {
                            _ = len(v4)
                            _ = len(v1)
                        }
                    }
                    var v30 int16 = -32768
                    _ = v30
                }
                for fi31 := 0; fi31 < 17; fi31++ {
                    fmt.Println(v6, v6)
                    fmt.Println(v6, v6, fi31)
                    _ = fi31
                }
            } else {
                if false {
                    _ = len(v1)
                    v6 += (v6 % ((42) | 1))
                }
                var v32 uint64 = 42
                for fi33 := 0; fi33 < 12; fi33++ {
                    if v2 {
                        fmt.Println(v6)
                    }
                    var v34 uint64 = (v32 ^ 0)
                    _ = fi33
                    _ = v34
                }
                _ = v32
            }
        }
    }
    v35 := GenS2{F0: genFn0(((("\n\t" + "a") + "hello") + "foo bar"), 234, genFn0("\n\t", 127, (v6 | -1))), F1: 42, F2: 18446744073709551615}
    var v36 bool = (109 != 1)
    v6 = 1
    if ((("hello" + ("hello" + ("\n\t" + ""))) != "a") || true) {
        switch (("\n\t" + ("hello" + "a")) + "") {
        case "a":
            fmt.Println(v3)
        case "hello":
            for fi37 := 0; fi37 < 16; fi37++ {
                v36 = (!v2)
                switch "\n\t" {
                case "foo bar":
                    switch ("foo bar" + "") {
                    case "\n\t":
                        var v38 int32 = -470
                        _ = v38
                    case "a":
                        v36 = false
                    default:
                        fmt.Println(v2, v6, v6)
                    }
                case "\n\t":
                    var v39 string = "a"
                    _ = v39
                case "a":
                    v1 = append(v1, 4294967295)
                }
                delete(v4, "a")
                _ = fi37
            }
        case "":
            v40 := v4["a"]
            _ = v40
        }
    } else {
        v4[("a" + "")] = (false && (4294967295 <= 4294967295))
        v41 := GenS4{F0: 65535, F1: 0}
        _ = v41
    }
    v2 = (-123 < 100)
    var v42 int16 = int16(v3)
    v43 := GenS4{F0: 307, F1: 0}
    v6 -= v6
    var v44 uint = (v3 & (v3 | v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v6
    _ = v7
}
