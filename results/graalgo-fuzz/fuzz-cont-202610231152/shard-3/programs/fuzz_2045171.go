package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0(p0 int64, p1 int, p2 uint32) uint {
    var v7 string = ""
    v8 := []int{(p1 ^ (p1 | p1)), (p1 & p1), 1, (p1 / (((p1 << ((2) & 63))) | 1))}
    switch 128 {
    case 0:
        p2 += (p2 & (^(p2)))
    default:
        switch (v7 + v7) {
        case "\n\t":
            v7 = (v7 + v7)
        }
    }
    if false {
        fmt.Println(p0, p0, p1)
        var v9 float64 = 0.0
        _ = v9
    }
    var v10 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    return v10
}

func main() {
    v1 := []int{591, 127}
    var v2 float64 = ((1.5 - 0.1) * -2.25)
    v3 := []string{}
    v4 := make(map[string]string)
    var v5 uint = 523
    v6 := []string{}
    var v11 bool = true
    fmt.Println(v2, v5, v11)
    if len(v3) > 0 {
        v3[0] = ("foo bar" + (("hello" + ("foo bar" + ("a" + ("" + "foo bar")))) + (("foo bar" + "hello") + (("" + ("\n\t" + "\n\t")) + "foo bar"))))
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 int16 = 0
        _ = fi12
        _ = v13
    }
    v11 = false
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS1{F0: 1, F1: 255, F2: 1}
        _ = fi14
        _ = v15
    }
    var v16 uint8 = 53
    v5 += (^(v5))
    switch "\n\t" {
    case "foo bar":
        if true {
            fmt.Println(v5)
            v11 = true
        } else {
            if (42 == 1) {
                v17 := GenS1{F0: v16, F1: uint16(v16), F2: 901}
                if v11 {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                switch ((((("hello" + "") + "hello") + ("" + "\n\t")) + "a") + "hello") {
                case "hello":
                    fmt.Println(v11, v11)
                case "a":
                    switch "" {
                    case "hello":
                        v16 -= 127
                    case "\n\t":
                        _ = len(v1)
                    case "a":
                        if (-686 != -9223372036854775808) {
                            for fi18 := 0; fi18 < 5; fi18++ {
                                var v19 uint = 42
                                _ = fi18
                                _ = v19
                            }
                            var v20 uint = v5
                            _ = v20
                        }
                    default:
                        v16 -= (^(v16))
                    }
                case "foo bar":
                    var v21 float64 = v2
                    _ = v21
                default:
                    var v22 string = (("hello" + ((("\n\t" + ((("hello" + "foo bar") + "a") + "")) + "\n\t") + "\n\t")) + "hello")
                    _ = v22
                }
                _ = v17
            } else {
                var v23 uint16 = 32767
                v24 := GenS1{F0: v16, F1: 2, F2: 4294967295}
                _ = v23
                _ = v24
            }
            for fi25 := 0; fi25 < 2; fi25++ {
                switch fi25 {
                case 9223372036854775807:
                    if (fi25 != (fi25 - (^(fi25)))) {
                        var v26 int16 = 0
                        var v27 uint = genFn0(991, fi25, 439)
                        _ = v26
                        _ = v27
                    }
                default:
                    v28 := GenS2{F0: 32767, F1: 2147483647}
                    _ = v28
                }
                _ = fi25
            }
        }
    }
    v5 = v5
    delete(v4, "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
