package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int16; F1 int32; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 string; F2 bool }
type GenS5 struct { F0 uint; F1 string }
func genFn0(p0 int32) bool {
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 18446744073709551615, F1: 531}
        v9 := GenS2{F0: 1, F1: (p0 ^ (p0 << ((516) & 63))), F2: 100}
        _ = fi7
        _ = v8
        _ = v9
    }
    p0 += p0
    _ = p0
    return ("" >= "")
}
func genFn1(p0 uint64, p1 uint8) int32 {
    var v10 string = "foo bar"
    v11 := []int{-9223372036854775808, 342, -2147483648, -2147483648}
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 4294967295
    v3 := []int{42, 42}
    v4 := []string{"foo bar"}
    var v5 uint8 = 42
    var v6 bool = (65535 == 1)
    fmt.Println(v5, v6)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        var v12 int32 = 2147483647
        var v13 int64 = 1
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        fmt.Println(v2, fi14, fi14)
        fmt.Println(v5)
        _ = fi14
    }
    fmt.Println(v6, v6, v6)
    var v15 uint64 = 366
    v5 += (v5 | 128)
    switch "a" {
    case "":
        var v16 int16 = -1
        _ = v16
    default:
        delete(v1, 0)
    }
    var v17 int64 = 42
    var v18 bool = v6
    if (2 == v15) {
        v5 = v5
        var v19 uint64 = v15
        if v18 {
            for fi20 := 0; fi20 < 2; fi20++ {
                switch ("a" + "\n\t") {
                case "hello":
                    if (true && (!(!(!(!(-9223372036854775808 >= (v17 % ((v17) | 1)))))))) {
                        if v6 {
                            if len(v3) > 0 {
                                v3[0] = (fi20 >> ((42) & 63))
                            }
                            var v21 int = (fi20 * (fi20 % ((fi20) | 1)))
                            _ = v21
                        }
                        fmt.Println(v18)
                    } else {
                        v22 := GenS3{F0: -2147483648}
                        _ = v22
                    }
                }
                _ = fi20
            }
            var v23 float32 = (((1.5 + (0.0 / 1.0)) + -2.25) - -2.25)
            _ = v23
        } else {
            v24 := GenS4{F0: 357, F1: "a", F2: (345 != 1)}
            for fi25 := 0; fi25 < 8; fi25++ {
                fmt.Println(v19, v19, fi25)
                for fi26 := 0; fi26 < 9; fi26++ {
                    v27 := GenS4{F0: 18446744073709551615, F1: ((((("foo bar" + "\n\t") + ("foo bar" + "a")) + "hello") + ("" + "a")) + (("hello" + "\n\t") + "hello")), F2: (-2147483648 != genFn1(v15, v5))}
                    switch 2147483647 {
                    case 32768:
                        v28 := GenS3{F0: -2147483648}
                        _ = v28
                    default:
                        var v29 uint32 = 1
                        _ = v29
                    }
                    switch ("" + ("" + "\n\t")) {
                    case "foo bar":
                        v15 = (v19 - 249)
                    }
                    _ = fi26
                    _ = v27
                }
                _ = fi25
            }
            if (!false) {
                var v30 string = (("a" + "a") + "a")
                _ = v30
            } else {
                fmt.Println(v5, v19, v6)
                var v31 uint8 = (v5 - (v5 / ((235) | 1)))
                for fi32 := 0; fi32 < 17; fi32++ {
                    v6 = (!(!(!v6)))
                    switch -2147483648 {
                    case -9223372036854775808:
                        v33 := GenS5{F0: 255, F1: ("" + ("foo bar" + (("a" + "") + ("\n\t" + ((("hello" + "hello") + "\n\t") + "a")))))}
                        _ = v33
                    }
                    _ = fi32
                }
                _ = v31
            }
            _ = v24
        }
        _ = v19
    } else {
        v17 = v17
    }
    var v34 uint16 = 42
    fmt.Println(v17, v34)
    var v35 uint = 184
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
