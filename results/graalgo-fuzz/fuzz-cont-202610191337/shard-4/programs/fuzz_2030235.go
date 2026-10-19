package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int; F1 int32 }
func genFn0(p0 uint8, p1 int64, p2 int) uint16 {
    for fi6 := 0; fi6 < 3; fi6++ {
        if true {
            p0 += 177
            if true {
                p0 += 1
            } else {
                switch p2 {
                case 32768:
                    v7 := []int{(fi6 ^ (^(p2))), (p2 >> ((0) & 63))}
                    _ = v7
                }
                p0 -= 42
                switch "a" {
                case "hello":
                    switch (^(fi6)) {
                    case 910:
                        p0 -= p0
                    }
                case "foo bar":
                    var v8 int16 = 10
                    _ = v8
                default:
                    var v9 int = -1
                    _ = v9
                }
            }
        } else {
            fmt.Println(p1, p2)
            for fi10 := 0; fi10 < 20; fi10++ {
                if (!(350 >= 65535)) {
                    switch (fi6 ^ p2) {
                    case 32768:
                        p1 = (p1 % ((42) | 1))
                    case -2147483648:
                        var v11 int = 42
                        _ = v11
                    default:
                        if false {
                            var v12 uint16 = 32767
                            _ = v12
                        }
                    }
                } else {
                    var v13 int64 = (p1 - p1)
                    v14 := GenS1{F0: 0, F1: "a"}
                    var v15 int16 = 0
                    var v16 uint = 65535
                    v17 := GenS2{F0: fi10, F1: int32(v15)}
                    _ = v13
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                }
                _ = fi10
            }
            p0 = 42
        }
        p0 = 0
        _ = fi6
    }
    p2 -= (p2 + (p2 - p2))
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 uint = 4294967295
        _ = fi18
        _ = v19
    }
    if ((p0 + p0) == p0) {
        var v20 uint8 = p0
        var v21 string = "hello"
        for fi22 := 0; fi22 < 19; fi22++ {
            var v23 uint64 = 0
            fmt.Println(p2, v23)
            _ = fi22
            _ = v23
        }
        _ = v20
        _ = v21
    } else {
        p2 += p2
    }
    var v24 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 2
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]uint32)
    var v3 int = 127
    v4 := make(map[uint32]int)
    var v5 uint8 = 128
    v1 -= (v1 & (v1 / (((^(v1))) | 1)))
    v5 = v5
    if true {
        var v25 int32 = -1
        if true {
            var v26 int64 = 0
            _ = v26
        }
        v1 = -1
        _ = v25
    }
    var v27 int16 = 42
    fmt.Println(v5)
    if (42 > 0) {
        if false {
            for fi28 := 0; fi28 < 4; fi28++ {
                fmt.Println(v1, v5, v3)
                _ = fi28
            }
        } else {
            v27 = v27
            _ = len(v2)
            var v29 int64 = 1
            _ = v29
        }
        var v30 int16 = (v27 & v27)
        _ = v30
    } else {
        var v31 int16 = (v27 ^ v27)
        if true {
            fmt.Println(v31)
        } else {
            var v32 float64 = (3.14159 * float64(3.14159))
            var v33 int16 = 0
            fmt.Println(v3)
            _ = v32
            _ = v33
        }
        _ = v31
    }
    fmt.Println(v1, v1, v1)
    v34 := GenS1{F0: 18446744073709551615, F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
