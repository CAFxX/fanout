package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint32, p1 uint8, p2 uint16) int16 {
    v4 := GenS1{F0: 9223372036854775807, F1: 150, F2: -128}
    for fi5 := 0; fi5 < 5; fi5++ {
        if (false || true) {
            for fi6 := 0; fi6 < 2; fi6++ {
                fmt.Println(fi6, fi5, fi6)
                v7 := GenS1{F0: -9223372036854775808, F1: 9223372036854775807, F2: 100}
                _ = fi6
                _ = v7
            }
            fmt.Println(fi5, fi5)
            if (42 <= 361) {
                if true {
                    p0 = (p0 >> ((0) & 63))
                    for fi8 := 0; fi8 < 10; fi8++ {
                        fmt.Println(fi5, p1, fi8)
                        switch 1 {
                        case 851:
                            var v9 int64 = -97
                            _ = v9
                        case 1:
                            var v10 uint32 = 171
                            _ = v10
                        default:
                            p0 -= 2
                        }
                        p1 += (^(p1))
                        _ = fi8
                    }
                    fmt.Println(p2, p1)
                }
            }
        }
        p2 = (p2 * (p2 + (p2 << ((255) & 63))))
        v11 := []int{fi5, (fi5 >> ((46) & 63))}
        _ = fi5
        _ = v11
    }
    var v12 uint32 = 354
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v4
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 bool = false
    v13 := []int{}
    v14 := v2[2147483647]
    switch "foo bar" {
    case "foo bar":
        v14 = false
    case "a":
        if (1 == -1) {
            fmt.Println(v14, v3)
            v15 := GenS2{F0: ("foo bar" + "foo bar")}
            fmt.Println(v14, v14)
            _ = v15
        }
    case "\n\t":
        v3 = false
    default:
        var v16 int64 = 42
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        fmt.Println(v14, v3)
        var v18 uint = 65535
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: "a"}
    v20 := GenS1{F0: 0, F1: -9223372036854775808, F2: 100}
    for fi21 := 0; fi21 < 17; fi21++ {
        v14 = v14
        v22 := GenS3{F0: 1, F1: 146}
        _ = fi21
        _ = v22
    }
    fmt.Println(v3, v3)
    v23 := GenS1{F0: -1, F1: 836, F2: 100}
    v3 = ((0.0 / -2.25) < (3.14159 + 1.5))
    var v24 uint = 63
    var v25 uint16 = 32767
    switch (("hello" + "a") + "\n\t") {
    case "a":
        if len(v13) > 0 {
            v13[0] = 1
        }
    case "foo bar":
        v13 = append(v13, -9223372036854775808)
    default:
        for fi26 := 0; fi26 < 1; fi26++ {
            fmt.Println(fi26, v14)
            if v14 {
                if (false && false) {
                    v27 := v1[("" + (((("a" + ("" + "foo bar")) + ("a" + (("\n\t" + (((("foo bar" + "foo bar") + "a") + ("\n\t" + "a")) + "hello")) + "a"))) + ("a" + "a")) + ("foo bar" + ("\n\t" + "a"))))]
                    switch "\n\t" {
                    case "\n\t":
                        for fi28 := 0; fi28 < 7; fi28++ {
                            if ((42 > 42) || ((v27 >= 32768) || false)) {
                                v29 := GenS1{F0: -1, F1: 9223372036854775807, F2: 127}
                                v14 = (v24 >= (v24 >> (((v24 & v24)) & 63)))
                                _ = v29
                            }
                            v1["hello"] = (v27 ^ fi28)
                            _ = fi28
                        }
                    }
                    v3 = ((9223372036854775807 < -40) || ((0.1 + (0.0 - 0.0)) <= -2.25))
                    _ = v27
                }
                v25 += v25
                for fi30 := 0; fi30 < 8; fi30++ {
                    v31 := GenS2{F0: "hello"}
                    var v32 int8 = 127
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            } else {
                if v14 {
                    for fi33 := 0; fi33 < 19; fi33++ {
                        v34 := GenS3{F0: -1, F1: (v25 & (v25 / ((v25) | 1)))}
                        v24 = (^(v24))
                        v35 := GenS3{F0: -32768, F1: (v25 & (v25 % ((v25) | 1)))}
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        if v3 {
                            delete(v2, 2)
                            v3 = (0 < 1)
                        }
                    }
                }
            }
            _ = fi26
        }
    }
    v36 := GenS2{F0: (("a" + "hello") + (("" + "hello") + (("\n\t" + ("" + ("" + "\n\t"))) + "a")))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v36
}
