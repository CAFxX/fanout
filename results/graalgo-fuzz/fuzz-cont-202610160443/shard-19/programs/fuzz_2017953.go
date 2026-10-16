package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint16; F2 int64 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0(p0 int32, p1 uint16) string {
    v4 := GenS1{F0: 1}
    var v5 uint = 1
    v5 = (^(v5))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return "\n\t"
}

func main() {
    var v1 int8 = -128
    v2 := []string{}
    var v3 float64 = (1e10 - 0.1)
    v2 = append(v2, "\n\t")
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        switch "hello" {
        case "hello":
            var v7 int16 = 42
            _ = v7
        default:
            if len(v2) > 0 {
                v2[0] = ((("a" + ("\n\t" + ("foo bar" + ""))) + "hello") + ("hello" + ("a" + "")))
            }
        }
        if true {
            var v8 float64 = v3
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v8
        } else {
            v2 = append(v2, "a")
            v3 -= v3
        }
        if true {
            v3 -= v3
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 int8 = (v1 - (v1 + (v1 << ((255) & 63))))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v11 := GenS1{F0: 2}
                _ = fi9
                _ = v10
                _ = v11
            }
        } else {
            fmt.Println(v3, fi6, v1)
            var v12 int32 = -1
            for fi13 := 0; fi13 < 8; fi13++ {
                v14 := GenS2{F0: v1, F1: 1, F2: -1}
                fmt.Println(v3)
                var v15 uint64 = 0
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        }
        _ = fi6
    }
    if ("hello" <= "a") {
        if (18446744073709551615 < 4294967295) {
            var v16 uint32 = 4294967295
            var v17 int64 = 41
            _ = v16
            _ = v17
        } else {
            v3 = v3
        }
    }
    v2 = append(v2, "")
    v1 += 0
    var v18 uint8 = 0
    fmt.Println(v18, v18)
    if true {
        var v19 int8 = 0
        v19 -= v1
        _ = v19
    } else {
        if (4294967295 == 42) {
            if len(v2) > 0 {
                v2[0] = ("\n\t" + ("foo bar" + ((("hello" + "") + "hello") + ((("a" + ("hello" + ("foo bar" + ""))) + "foo bar") + ((("" + "") + "") + "foo bar")))))
            }
            var v20 uint32 = 2147483648
            for fi21 := 0; fi21 < 4; fi21++ {
                if false {
                    for fi22 := 0; fi22 < 19; fi22++ {
                        var v23 string = "a"
                        var v24 int16 = -126
                        switch v23 {
                        case "foo bar":
                            fmt.Println(v20, v18, v24)
                        case "a":
                            var v25 uint32 = v20
                            _ = v25
                        case "":
                            v3 = 3.14159
                        }
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    v1 -= (v1 + (v1 & (v1 << ((42) & 63))))
                    v20 -= (^(v20))
                }
                v26 := GenS2{F0: v1, F1: 255, F2: -1}
                var v27 float64 = (v3 - v3)
                _ = fi21
                _ = v26
                _ = v27
            }
            _ = v20
        }
    }
    fmt.Println(v1, v18)
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := GenS3{F0: 1000, F1: -1}
        v1 += v1
        _ = fi28
        _ = v29
    }
    v1 -= 119
    v30 := GenS4{F0: int8(v18), F1: 373}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
}
