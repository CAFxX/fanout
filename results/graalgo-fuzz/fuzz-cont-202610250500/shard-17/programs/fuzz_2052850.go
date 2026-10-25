package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int32 {
    if (980 == -731) {
        v7 := GenS1{F0: 794}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 3; fi8++ {
            if ("" <= "\n\t") {
                var v9 uint8 = 42
                var v10 uint64 = 2
                _ = v10
                _ = v9
            }
            _ = fi8
        }
    }
    var v11 int32 = 504
    _ = v11
    return v11
}
func genFn1() int {
    if ((-9223372036854775808 > 32767) && (!(-9223372036854775808 != -1))) {
        var v12 uint8 = 127
        _ = v12
    }
    var v13 float64 = float64(3.14159)
    v14 := []int{-9223372036854775808, 127}
    _ = v13
    _ = v14
    return 127
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (((0.1 + 3.14159) + 0.1) * 1.5)
    var v3 uint32 = 2
    var v4 uint32 = (v3 ^ v3)
    var v5 int16 = -895
    v6 := []uint32{}
    delete(v1, "")
    switch 737 {
    case 42:
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 uint = 65535
            switch "" {
            case "\n\t":
                var v17 uint32 = (^(v3))
                _ = v17
            case "foo bar":
                for fi18 := 0; fi18 < 2; fi18++ {
                    var v19 float32 = 1e10
                    _ = fi18
                    _ = v19
                }
            default:
                for fi20 := 0; fi20 < 20; fi20++ {
                    var v21 float64 = -2.25
                    switch (fi15 ^ genFn1()) {
                    case 2147483647:
                        if len(v6) > 0 {
                            v6[0] = v4
                        }
                    case -1:
                        fmt.Println(v4, v5)
                    default:
                        var v22 uint = v16
                        _ = v22
                    }
                    if (2 >= 1) {
                        v23 := GenS1{F0: (^(v16))}
                        _ = v23
                    } else {
                        var v24 int32 = 42
                        fmt.Println(v4, fi20, v16)
                        v25 := GenS1{F0: 255}
                        _ = v24
                        _ = v25
                    }
                    fmt.Println(v2)
                    var v26 int = genFn1()
                    _ = fi20
                    _ = v21
                    _ = v26
                }
            }
            _ = fi15
            _ = v16
        }
    case 32767:
        if (false && true) {
            fmt.Println(v4)
        }
    case 127:
        var v27 string = "foo bar"
        _ = v27
    default:
        fmt.Println(v4, v5)
    }
    v28 := GenS1{F0: 901}
    var v29 float64 = 3.14159
    if false {
        var v30 float64 = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 uint64 = 4294967295
        _ = fi31
        _ = v32
    }
    switch ("hello" + "hello") {
    case "\n\t":
        for fi33 := 0; fi33 < 12; fi33++ {
            fmt.Println(v3, v4, fi33)
            v29 -= v29
            var v34 string = "a"
            _ = fi33
            _ = v34
        }
    case "":
        var v35 bool = ((1e10 + 1e10) < 1e10)
        _ = v35
    default:
        for fi36 := 0; fi36 < 16; fi36++ {
            var v37 string = "foo bar"
            v38 := GenS1{F0: 4294967295}
            v4 -= 65535
            _ = fi36
            _ = v37
            _ = v38
        }
    }
    var v39 int8 = 42
    var v40 int = 127
    var v41 uint = 0
    for fi42 := 0; fi42 < 20; fi42++ {
        v43 := GenS1{F0: 2}
        _ = fi42
        _ = v43
    }
    for fi44 := 0; fi44 < 5; fi44++ {
        var v45 int32 = -1
        _ = len(v6)
        _ = fi44
        _ = v45
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v2 = (0.0 + ((3.14159 * 1e10) / v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
