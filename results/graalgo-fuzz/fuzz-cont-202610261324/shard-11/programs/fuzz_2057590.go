package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 2147483647
    var v3 uint32 = 2147483648
    var v4 float64 = (((1e10 / float64((1.5 * -2.25))) / (1e10 - 0.0)) / 0.1)
    v5 := make(map[string]bool)
    v6 := []uint32{v3, (v3 - 42), (v2 % ((v3) | 1)), v3}
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 int64 = 42
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v3 += v2
    if ((100 <= -7) || (!(42 != 18446744073709551615))) {
        var v10 int8 = -1
        if false {
            v6 = append(v6, (^(v3)))
        } else {
            for fi11 := 0; fi11 < 5; fi11++ {
                switch "hello" {
                case "hello":
                    if false {
                        var v12 string = ""
                        _ = v12
                    }
                case "foo bar":
                    v10 -= (v10 & v10)
                case "a":
                    var v13 uint32 = v3
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 12; fi14++ {
                        v3 += (^(v3))
                        if true {
                            delete(v1, fi11)
                            v15 := GenS1{F0: "hello", F1: 1}
                            _ = len(v6)
                            fmt.Println(v3, fi11)
                            _ = v15
                        } else {
                            v10 += (v10 << ((415) & 63))
                            delete(v1, -386)
                            v16 := GenS1{F0: "\n\t", F1: -32768}
                            _ = v16
                        }
                        var v17 bool = ((4294967295 >= 42) || false)
                        _ = fi14
                        _ = v17
                    }
                }
                fmt.Println(v4)
                _ = fi11
            }
            var v18 uint32 = 385
            v10 = (v10 / ((v10) | 1))
            _ = v18
        }
        _ = v10
    } else {
        for fi19 := 0; fi19 < 17; fi19++ {
            v2 -= 2147483648
            _ = fi19
        }
        var v20 float64 = 3.14159
        _ = v20
    }
    var v21 float64 = 3.14159
    if false {
        var v22 bool = ("a" < "\n\t")
        _ = v22
    } else {
        fmt.Println(v4, v4, v4)
        v3 = v3
        _ = len(v6)
    }
    var v23 uint16 = 42
    v6 = append(v6, (v3 << ((42) & 63)))
    v24 := GenS1{F0: (("" + (("\n\t" + "foo bar") + "")) + (("hello" + ((("a" + "a") + "a") + (("a" + "foo bar") + "hello"))) + "hello")), F1: -1}
    v25 := GenS1{F0: (("foo bar" + "foo bar") + ("a" + "")), F1: -1}
    var v26 int64 = 1
    var v27 int = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
