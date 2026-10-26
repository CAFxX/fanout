package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32, p1 int8) int32 {
    p1 = p1
    var v4 float32 = ((0.0 / -2.25) + 1e10)
    var v5 float32 = ((0.1 * 0.0) + 1e10)
    v6 := GenS1{F0: 32767, F1: 42, F2: (p0 << ((135) & 63))}
    for fi7 := 0; fi7 < 18; fi7++ {
        p1 = (p1 & (^(p1)))
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    var v1 uint = 4294967295
    var v2 int64 = 9223372036854775807
    v3 := []string{}
    v8 := GenS2{F0: 2}
    var v9 bool = false
    v9 = (0 < 127)
    var v10 int8 = 100
    var v11 int32 = genFn0(2147483647, v10)
    fmt.Println(v1, v2)
    if (v1 < v1) {
        _ = len(v3)
        v12 := GenS1{F0: 1, F1: 42, F2: 4294967295}
        _ = v12
    }
    v11 = genFn0(2, 42)
    for fi13 := 0; fi13 < 1; fi13++ {
        v2 -= 958
        switch "\n\t" {
        case "\n\t":
            if true {
                switch "hello" {
                case "hello":
                    for fi14 := 0; fi14 < 12; fi14++ {
                        var v15 bool = true
                        switch "a" {
                        case "":
                            var v16 int8 = v10
                            _ = v16
                        case "\n\t":
                            fmt.Println(v15)
                        }
                        v11 = v11
                        _ = fi14
                        _ = v15
                    }
                case "a":
                    var v17 uint8 = 1
                    _ = v17
                default:
                    _ = len(v3)
                }
                if (fi13 <= (fi13 * (^(fi13)))) {
                    var v18 string = ""
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if len(v3) > 0 {
                        v3[0] = (v18 + v18)
                    }
                    _ = v18
                }
            }
        case "hello":
            var v19 int16 = 1000
            _ = v19
        case "a":
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 float32 = 3.14159
                _ = fi20
                _ = v21
            }
        }
        v22 := GenS1{F0: 32767, F1: uint16(v1), F2: 0}
        _ = fi13
        _ = v22
    }
    fmt.Println(v1, v2)
    var v23 int32 = (-(v11))
    v10 -= v10
    fmt.Println(v23)
    var v24 uint32 = 1
    for fi25 := 0; fi25 < 17; fi25++ {
        switch fi25 {
        case 127:
            if ((-(v2)) <= (^(v2))) {
                v26 := GenS1{F0: 65535, F1: 2, F2: v24}
                _ = v26
            } else {
                var v27 uint = v1
                if (!v9) {
                    v28 := GenS2{F0: 255}
                    switch (((("foo bar" + "\n\t") + "a") + "foo bar") + ("" + "")) {
                    case "foo bar":
                        if (0.1 == 0.1) {
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                        } else {
                            fmt.Println(v27, v2, fi25)
                            var v29 bool = ((v2 * (v2 & (-(v2)))) > v2)
                            _ = v29
                        }
                    default:
                        v30 := GenS1{F0: 65535, F1: 0, F2: (v24 & 2147483648)}
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 20; fi31++ {
                        var v32 uint8 = 42
                        _ = fi31
                        _ = v32
                    }
                    _ = v28
                } else {
                    for fi33 := 0; fi33 < 15; fi33++ {
                        v10 -= v10
                        _ = fi33
                    }
                    v9 = v9
                    fmt.Println(v11)
                }
                v23 += int32(v10)
                _ = v27
            }
        case 42:
            if (!((1 >= 1) && v9)) {
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                switch "foo bar" {
                case "hello":
                    var v34 string = ("foo bar" + ("hello" + (("a" + "") + "")))
                    _ = v34
                default:
                    v2 -= -514
                }
                fmt.Println(v9, v9, v11)
            } else {
                var v35 uint16 = 249
                for fi36 := 0; fi36 < 19; fi36++ {
                    v37 := GenS1{F0: 42, F1: v35, F2: 1}
                    _ = fi36
                    _ = v37
                }
                for fi38 := 0; fi38 < 6; fi38++ {
                    var v39 int16 = 42
                    if ((((v24 - v24) != v24) && (!(v9 || (!(false || true))))) && v9) {
                        v11 = v23
                        v2 -= (v2 | -172)
                    } else {
                        var v40 float64 = (0.0 / (-2.25 / float64((float32(-2.25) * 1e10))))
                        var v41 uint16 = 255
                        _ = v40
                        _ = v41
                    }
                    _ = fi38
                    _ = v39
                }
                _ = v35
            }
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v8
    _ = v9
}
