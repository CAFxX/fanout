package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 string; F2 uint8 }
type GenS4 struct { F0 string }
func genFn0(p0 int64, p1 int) int64 {
    v7 := []int{}
    v8 := GenS1{F0: 18446744073709551615, F1: "foo bar"}
    switch (("\n\t" + "hello") + "") {
    case "\n\t":
        fmt.Println(p1, p0)
    default:
        p1 -= (p1 | 1)
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]int)
    v4 := []uint32{42}
    var v5 float32 = 0.1
    var v6 uint = 200
    var v9 float64 = (1.5 / 0.1)
    if (((v6 ^ v6) != (v6 / ((v6) | 1))) && ((-83 > 42) || true)) {
        for fi10 := 0; fi10 < 20; fi10++ {
            v4 = append(v4, 0)
            _ = fi10
        }
        fmt.Println(v6, v5, v2)
    } else {
        if false {
            if (!false) {
                v3[1] = 972
                var v11 string = ("foo bar" + "a")
                var v12 float64 = v9
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 17; fi13++ {
                    v1 = (-(v1))
                    var v14 int = -35
                    _ = fi13
                    _ = v14
                }
                if (!(false && (!(!false)))) {
                    var v15 string = "a"
                    v9 += v9
                    _ = v15
                } else {
                    switch -452 {
                    case 9223372036854775807:
                        switch "\n\t" {
                        case "hello":
                            if false {
                                _ = len(v3)
                            }
                        case "foo bar":
                            fmt.Println(v1, v6, v6)
                        case "":
                            var v16 string = "hello"
                            _ = v16
                        default:
                            var v17 uint16 = 374
                            _ = v17
                        }
                    case 42:
                        v18 := v3[32768]
                        _ = v18
                    }
                }
            }
            for fi19 := 0; fi19 < 1; fi19++ {
                v20 := GenS2{F0: 2147483647}
                v1 = v1
                _ = fi19
                _ = v20
            }
        }
        v21 := GenS2{F0: 42}
        _ = v21
    }
    v22 := GenS3{F0: -1, F1: ((("" + "\n\t") + "foo bar") + "hello"), F2: 2}
    var v23 float64 = v9
    var v24 float64 = ((v9 / v9) - (v23 + v23))
    var v25 uint16 = 65535
    v24 = 3.14159
    switch 32767 {
    case 128:
        v26 := GenS4{F0: "\n\t"}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
