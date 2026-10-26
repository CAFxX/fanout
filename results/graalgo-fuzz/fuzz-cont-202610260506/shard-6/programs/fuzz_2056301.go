package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int; F1 uint16; F2 int }
type GenS3 struct { F0 string; F1 int8; F2 uint }
func genFn0() uint64 {
    var v4 int = 127
    v4 = (^(v4))
    v5 := []int{v4, (v4 & v4)}
    v4 += 222
    fmt.Println(v4)
    _ = v4
    _ = v5
    return 706
}

func main() {
    var v1 uint16 = 0
    var v2 bool = false
    var v3 int8 = 0
    v6 := []int{}
    v2 = v2
    var v7 uint8 = uint8(v1)
    for fi8 := 0; fi8 < 9; fi8++ {
        if (true && true) {
            var v9 bool = true
            _ = v9
        }
        var v10 uint16 = (v1 * 289)
        var v11 uint64 = 0
        _ = fi8
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        v2 = (!(((fi12 | -166) == 32768) && true))
        _ = fi12
    }
    var v13 bool = (18446744073709551615 == 4294967295)
    v7 += 2
    v6 = append(v6, 644)
    switch int(v7) {
    case 271:
        if false {
            if v2 {
                v13 = v2
                var v14 uint = 65535
                fmt.Println(v1, v1)
                _ = v14
            }
        } else {
            if (2 >= 782) {
                if true {
                    if v13 {
                        var v15 uint16 = v1
                        fmt.Println(v15, v7)
                        _ = v15
                    }
                    switch 2147483647 {
                    case -2147483648:
                        v3 -= 52
                    case 1:
                        if true {
                            v16 := GenS1{F0: 0, F1: 0}
                            _ = len(v6)
                            _ = v16
                        } else {
                            var v17 uint8 = v7
                            _ = v17
                        }
                    case 32768:
                        var v18 int64 = -1
                        _ = v18
                    default:
                        var v19 int = 127
                        _ = v19
                    }
                } else {
                    _ = len(v6)
                    if false {
                        var v20 float64 = (0.1 / -2.25)
                        var v21 uint8 = v7
                        _ = v20
                        _ = v21
                    }
                }
            }
            var v22 int = 32767
            v6 = append(v6, (v22 * (-(v22))))
            _ = v22
        }
    }
    switch 1 {
    case -2147483648:
        for fi23 := 0; fi23 < 4; fi23++ {
            v2 = (42 <= uint64(v1))
            var v24 int16 = 1000
            for fi25 := 0; fi25 < 9; fi25++ {
                v3 = int8(v7)
                if (v13 || (264 > uint(v1))) {
                    fmt.Println(v7, v7)
                    v26 := GenS2{F0: -449, F1: (v1 | v1), F2: fi25}
                    switch "hello" {
                    case "hello":
                        v27 := GenS3{F0: "", F1: 100, F2: 2}
                        _ = v27
                    }
                    _ = v26
                }
                switch "" {
                case "foo bar":
                    var v28 int64 = 9223372036854775807
                    _ = v28
                }
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
    case -1:
        for fi29 := 0; fi29 < 7; fi29++ {
            _ = len(v6)
            var v30 int16 = -1
            _ = fi29
            _ = v30
        }
    case 127:
        var v31 uint8 = v7
        _ = v31
    }
    v32 := GenS3{F0: "hello", F1: (v3 | v3), F2: 18446744073709551615}
    var v33 int16 = 299
    var v34 int32 = 1
    var v35 string = "hello"
    v36 := GenS2{F0: -9223372036854775808, F1: v1, F2: 703}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v6
    _ = v7
}
