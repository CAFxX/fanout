package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
func genFn0() uint32 {
    var v4 int8 = 1
    v5 := GenS1{F0: ("\n\t" + ((("a" + "") + "hello") + "foo bar"))}
    v4 -= -60
    v4 = v4
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[uint32]string)
    var v3 int16 = -1
    var v6 uint8 = 1
    v7 := GenS2{F0: 65535, F1: 1, F2: 42}
    v1 += 0.0
    v8 := GenS3{F0: 65535}
    switch "a" {
    case "foo bar":
        v9 := []int{0, 42, 32768}
        _ = v9
    case "":
        fmt.Println(v6, v1)
    case "hello":
        if false {
            var v10 uint32 = 588
            _ = v10
        } else {
            if (false || (true && (v6 != (^(v6))))) {
                if (255 == 587) {
                    var v11 string = ("hello" + "foo bar")
                    var v12 int32 = 42
                    for fi13 := 0; fi13 < 16; fi13++ {
                        v14 := GenS2{F0: 475, F1: 458, F2: 2}
                        var v15 uint32 = 2147483648
                        var v16 float32 = ((0.1 + v1) + v1)
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    _ = v11
                    _ = v12
                } else {
                    v3 = v3
                    if true {
                        fmt.Println(v6)
                        v3 -= v3
                        v6 -= v6
                    }
                }
            } else {
                v17 := v2[0]
                switch "foo bar" {
                case "a":
                    if false {
                        if true {
                            v17 = v17
                        }
                        v18 := []int{-1, 805, 549, 42}
                        _ = len(v18)
                        _ = v18
                    } else {
                        for fi19 := 0; fi19 < 7; fi19++ {
                            fmt.Println(v3, v6, v17)
                            for fi20 := 0; fi20 < 12; fi20++ {
                                v21 := GenS4{F0: ""}
                                fmt.Println(v17, fi19)
                                fmt.Println(v17, v17)
                                _ = fi20
                                _ = v21
                            }
                            _ = fi19
                        }
                    }
                case "hello":
                    v3 = (v3 + -32768)
                case "":
                    v6 -= (v6 + (^(v6)))
                }
                _ = v17
            }
        }
    }
    v3 += 1000
    for fi22 := 0; fi22 < 15; fi22++ {
        v1 -= (v1 - ((0.1 * (v1 / 1e10)) / 3.14159))
        _ = fi22
    }
    v6 += 212
    var v23 uint8 = (v6 | (^(v6)))
    var v24 float64 = (0.1 + 0.1)
    fmt.Println(v23, v1)
    var v25 int64 = 9223372036854775807
    v2[2147483648] = "a"
    var v26 uint64 = 918
    v25 += (-(v25))
    v27 := []int{1, 127, -9223372036854775808, 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
