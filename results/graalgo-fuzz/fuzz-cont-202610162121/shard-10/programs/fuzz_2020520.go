package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int; F1 int8; F2 uint32 }
type GenS3 struct { F0 int32; F1 string }
func genFn0() int {
    switch 9223372036854775807 {
    case 32768:
        var v7 uint = 18446744073709551615
        _ = v7
    default:
        var v8 int64 = 42
        _ = v8
    }
    switch -1 {
    case 401:
        v9 := []int{42, 383}
        _ = v9
    case 32767:
        v10 := []int{32768, 0, 42, -2147483648}
        _ = v10
    case -210:
        v11 := GenS1{F0: (("a" + "") + "foo bar"), F1: false}
        _ = v11
    default:
        if (!true) {
            if (("foo bar" == ("hello" + "hello")) && false) {
                switch -1 {
                case 9223372036854775807:
                    var v12 float64 = 1e10
                    _ = v12
                case 0:
                    v13 := []int{2147483647}
                    _ = v13
                case 32767:
                    var v14 string = "\n\t"
                    _ = v14
                }
                var v15 int64 = 9223372036854775807
                v15 = (^(v15))
                _ = v15
            }
        } else {
            v16 := GenS1{F0: (("\n\t" + ("foo bar" + "foo bar")) + "\n\t"), F1: (18446744073709551615 != 2)}
            _ = v16
        }
    }
    var v17 uint = 0
    if (-128 >= 42) {
        v18 := []int{-9223372036854775808, -8, 1}
        var v19 uint = 4294967295
        v17 = (v17 * (v17 ^ v19))
        _ = v18
        _ = v19
    } else {
        if false {
            switch -1 {
            case 0:
                if false {
                    v17 -= (v17 * v17)
                } else {
                    var v20 float64 = -2.25
                    v20 -= 1e10
                    if true {
                        v21 := GenS1{F0: ("hello" + "hello"), F1: true}
                        var v22 int8 = -78
                        _ = v21
                        _ = v22
                    } else {
                        fmt.Println(v17, v17)
                    }
                    _ = v20
                }
            default:
                var v23 uint8 = 160
                _ = v23
            }
            if true {
                v17 += (v17 << ((v17) & 63))
                var v24 uint = (^(v17))
                _ = v24
            } else {
                v17 -= (v17 >> ((v17) & 63))
            }
        }
        if true {
            fmt.Println(v17)
            v25 := GenS2{F0: -83, F1: -1, F2: 2147483647}
            _ = v25
        } else {
            v17 = v17
            v17 = (v17 % ((v17) | 1))
            var v26 uint32 = 96
            _ = v26
        }
        for fi27 := 0; fi27 < 2; fi27++ {
            v28 := GenS2{F0: (fi27 % ((42) | 1)), F1: 127, F2: 2147483647}
            _ = fi27
            _ = v28
        }
    }
    _ = v17
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    v3 := []uint32{185}
    v4 := make(map[uint32]uint32)
    var v5 int8 = 42
    v6 := []string{(("hello" + "") + (("" + "hello") + "")), "", "\n\t"}
    var v29 int16 = 0
    var v30 float32 = -2.25
    fmt.Println(v29, v5, v5)
    v30 += v30
    v31 := GenS3{F0: -448, F1: ("hello" + (("" + "foo bar") + "foo bar"))}
    switch 901 {
    case 554:
        v5 = v5
    default:
        for fi32 := 0; fi32 < 19; fi32++ {
            if true {
                fmt.Println(fi32)
            }
            _ = fi32
        }
    }
    v33 := v1[2147483647]
    switch v33 {
    case "foo bar":
        var v34 float64 = 1e10
        _ = v34
    case "\n\t":
        var v35 uint16 = 32767
        _ = v35
    case "a":
        fmt.Println(v5, v30)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
