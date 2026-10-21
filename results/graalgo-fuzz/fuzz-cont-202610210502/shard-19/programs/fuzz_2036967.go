package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint8; F2 int8 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 string; F1 uint64 }
func genFn0(p0 uint, p1 uint32) int8 {
    fmt.Println(p1, p0)
    if ((uint(p1) < (p0 % (((^(p0))) | 1))) && false) {
        v5 := []int{}
        p0 += 4294967295
        fmt.Println(p1, p0)
        _ = v5
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 127
}

func main() {
    var v1 float64 = 1e10
    var v2 int8 = -128
    var v3 int16 = -308
    v4 := make(map[string]uint32)
    var v6 uint = 4
    if true {
        var v7 string = ("a" + "foo bar")
        _ = v7
    } else {
        v8 := GenS1{F0: 239, F1: (v2 & 42), F2: (v3 ^ v3)}
        var v9 int16 = v3
        _ = v8
        _ = v9
    }
    if (!((("hello" + "a") > "") && (v3 > v3))) {
        v3 -= (v3 << ((958) & 63))
        fmt.Println(v2, v2, v2)
    } else {
        var v10 int32 = 1
        v1 += (v1 * v1)
        if true {
            v11 := GenS1{F0: 144, F1: (v2 / (((v2 + v2)) | 1)), F2: v3}
            _ = v11
        } else {
            v12 := GenS2{F0: int64(v3)}
            v13 := []int{}
            for fi14 := 0; fi14 < 4; fi14++ {
                v15 := GenS3{F0: -32768, F1: 42, F2: int8(fi14)}
                var v16 uint8 = 2
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v12
            _ = v13
        }
        _ = v10
    }
    fmt.Println(v2, v3, v3)
    fmt.Println(v6, v2, v1)
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 int = 9223372036854775807
        v19 := []int{}
        _ = fi17
        _ = v18
        _ = v19
    }
    switch "hello" {
    case "foo bar":
        v20 := []int{128, -157, -680}
        _ = v20
    case "hello":
        var v21 string = "hello"
        _ = v21
    case "":
        v22 := GenS4{F0: 0, F1: -2147483648}
        _ = v22
    }
    if (v6 > (v6 >> (((v6 | (^(v6)))) & 63))) {
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := GenS3{F0: v3, F1: 0, F2: v2}
            var v25 uint32 = 2147483647
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 int8 = genFn0(977, 65535)
                _ = fi26
                _ = v27
            }
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    if (65535 >= 592) {
        v6 = v6
        for fi28 := 0; fi28 < 12; fi28++ {
            v6 -= 42
            _ = fi28
        }
        var v29 string = "\n\t"
        _ = v29
    } else {
        var v30 int64 = -9223372036854775808
        _ = v30
    }
    fmt.Println(v6, v3)
    if (("a" + ("" + ((("hello" + "foo bar") + "hello") + "a"))) != "hello") {
        var v31 uint16 = 65535
        if true {
            if true {
                for fi32 := 0; fi32 < 2; fi32++ {
                    switch 9223372036854775807 {
                    case 439:
                        fmt.Println(v1)
                    case -256:
                        fmt.Println(v2, v2, v2)
                    case 0:
                        for fi33 := 0; fi33 < 10; fi33++ {
                            fmt.Println(v3)
                            _ = fi33
                        }
                    }
                    _ = fi32
                }
            } else {
                var v34 string = "\n\t"
                for fi35 := 0; fi35 < 14; fi35++ {
                    switch fi35 {
                    case 2147483647:
                        v36 := GenS5{F0: (((v34 + v34) + v34) + ""), F1: 824}
                        _ = v36
                    }
                    v31 = 380
                    _ = fi35
                }
                _ = v34
            }
        } else {
            v6 -= (v6 / ((18446744073709551615) | 1))
        }
        _ = len(v4)
        _ = v31
    } else {
        fmt.Println(v3, v2, v2)
    }
    var v37 uint16 = 926
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v6
}
