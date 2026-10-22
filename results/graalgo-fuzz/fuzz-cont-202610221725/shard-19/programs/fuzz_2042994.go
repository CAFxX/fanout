package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int16, p1 uint16, p2 string) int {
    fmt.Println(p2, p2, p2)
    v7 := GenS1{F0: 0, F1: 1000}
    fmt.Println(p1, p1)
    p2 = "foo bar"
    var v8 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]int)
    var v3 int = 32768
    var v4 int64 = 1
    var v5 uint = 4294967295
    v6 := []int{}
    var v9 int = (v3 + genFn0(-962, 285, ("\n\t" + "a")))
    switch (v9 ^ (v3 * 975)) {
    case -881:
        fmt.Println(v5, v5, v9)
    default:
        if true {
            if (true || (!true)) {
                switch (v9 + (-(v9))) {
                case 4294967295:
                    v3 = 1
                }
                v5 -= v5
            }
            if true {
                fmt.Println(v1)
                fmt.Println(v5, v4, v9)
                if (!(true || (42 == 229))) {
                    v10 := GenS1{F0: v9, F1: -659}
                    v3 += (-(v3))
                    _ = v10
                }
            } else {
                switch "a" {
                case "\n\t":
                    _ = len(v6)
                }
                var v11 int32 = 42
                for fi12 := 0; fi12 < 17; fi12++ {
                    switch v3 {
                    case 295:
                        if (!(!(401 >= (^(v4))))) {
                            if true {
                                var v13 string = "a"
                                v3 = v9
                                v13 = v13
                                _ = v13
                            } else {
                                for fi14 := 0; fi14 < 11; fi14++ {
                                    var v15 int32 = v11
                                    v16 := GenS1{F0: int(v15), F1: (v1 / (((v1 / ((0) | 1))) | 1))}
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                            }
                            var v17 string = "a"
                            _ = v17
                        } else {
                            v18 := GenS2{F0: 128}
                            _ = v18
                        }
                    default:
                        _ = len(v6)
                    }
                    _ = fi12
                }
                _ = v11
            }
        } else {
            for fi19 := 0; fi19 < 11; fi19++ {
                for fi20 := 0; fi20 < 14; fi20++ {
                    v4 = (v4 << ((933) & 63))
                    _ = fi20
                }
                _ = fi19
            }
            var v21 string = ((("" + "a") + "hello") + (("foo bar" + "foo bar") + "foo bar"))
            v5 = 42
            _ = v21
        }
    }
    v9 = (v3 & v3)
    switch (v9 % ((v9) | 1)) {
    case -467:
        v5 -= 0
    default:
        var v22 uint = (v5 | 553)
        _ = v22
    }
    v23 := v2["foo bar"]
    v1 = v1
    var v24 bool = ((-(v3)) != genFn0(v1, 2, ("\n\t" + (("\n\t" + "") + "hello"))))
    v24 = (42 > 844)
    switch "a" {
    case "":
        var v25 int16 = (v1 << (((v5 | 18446744073709551615)) & 63))
        _ = v25
    case "hello":
        v4 = v4
    default:
        v26 := GenS2{F0: 127}
        _ = v26
    }
    var v27 float32 = (1e10 / 0.1)
    var v28 int32 = 42
    v1 += (v1 * (v1 - (v1 * (v1 >> ((v5) & 63)))))
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 int = -9223372036854775808
        _ = fi29
        _ = v30
    }
    var v31 uint16 = 2
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
