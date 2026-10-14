package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS5 struct { F0 int8; F1 int64; F2 bool }
func genFn0(p0 string) int {
    v6 := GenS1{F0: 805, F1: 0}
    p0 = (p0 + p0)
    p0 = p0
    var v7 int = -9223372036854775808
    _ = p0
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 int32, p1 uint8, p2 uint) int16 {
    var v8 uint8 = p1
    var v9 uint8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 2147483648
    var v3 bool = (0 >= 42)
    var v4 int = 165
    v5 := make(map[string]string)
    var v10 uint16 = 130
    v4 += genFn0("foo bar")
    switch "hello" {
    case "\n\t":
        switch ("hello" + "") {
        case "hello":
            for fi11 := 0; fi11 < 5; fi11++ {
                switch v4 {
                case 0:
                    if (!v3) {
                        for fi12 := 0; fi12 < 16; fi12++ {
                            if (((float32((0.0 / 1.0)) - 3.14159) / 0.1) < 0.0) {
                                v3 = v3
                                fmt.Println(v2)
                                var v13 uint64 = 42
                                _ = v13
                            } else {
                                v4 = genFn0(("\n\t" + "a"))
                                if (1 == 0) {
                                    if ((("\n\t" + "a") + (("" + ("a" + (("\n\t" + ("\n\t" + ("hello" + ""))) + "foo bar"))) + ("hello" + "\n\t"))) > "foo bar") {
                                        fmt.Println(v4)
                                    }
                                } else {
                                    v14 := GenS2{F0: -1, F1: ((v4 >> ((1) & 63)) < fi11)}
                                    v4 = genFn0("hello")
                                    _ = v14
                                }
                                v15 := GenS3{F0: 127}
                                var v16 int = genFn0((("a" + ("\n\t" + ("a" + ((((("hello" + "a") + "\n\t") + "foo bar") + ("hello" + "a")) + "hello")))) + "a"))
                                fmt.Println(v16, v2)
                                _ = v15
                                _ = v16
                            }
                            delete(v1, genFn0(""))
                            _ = fi12
                        }
                    } else {
                        fmt.Println(v10, v10)
                        v17 := GenS3{F0: 1}
                        _ = v17
                    }
                case 470:
                    v18 := GenS4{F0: false, F1: (v10 * 32767), F2: 0}
                    _ = v18
                }
                _ = fi11
            }
        case "":
            if v3 {
                v10 = (^(v10))
                v19 := GenS5{F0: 89, F1: 409, F2: (!(0.0 == 0.0))}
                _ = v19
            }
        }
    default:
        fmt.Println(v4, v4, v3)
    }
    var v20 uint16 = (v10 - 2)
    v3 = (!v3)
    switch v4 {
    case 32768:
        fmt.Println(v3, v2)
    case -516:
        if v3 {
            var v21 int = -400
            if v3 {
                var v22 uint64 = 790
                _ = v22
            } else {
                v23 := []int{(v4 ^ 9223372036854775807)}
                _ = v23
            }
            _ = v21
        } else {
            v2 += 4294967295
            v24 := GenS3{F0: 171}
            _ = v24
        }
    }
    if (32767 == -32768) {
        for fi25 := 0; fi25 < 4; fi25++ {
            fmt.Println(v2)
            v26 := GenS5{F0: 83, F1: 9223372036854775807, F2: (!(2147483648 != v2))}
            var v27 bool = (v3 && ((2 < 0) || ((v2 / ((v2) | 1)) != v2)))
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    v28 := v5[("" + "a")]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
