package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int16; F1 string }
func genFn0() int {
    var v7 float64 = 3.14159
    v7 = 0.0
    var v8 float64 = (-2.25 + 0.0)
    fmt.Println(v8)
    var v9 uint8 = 1
    _ = v7
    _ = v8
    _ = v9
    return 32768
}

func main() {
    var v1 uint32 = 0
    var v2 uint16 = 1
    v3 := make(map[int]int)
    var v4 uint16 = 42
    var v5 int64 = -438
    v6 := make(map[uint32]string)
    var v10 float64 = 0.0
    fmt.Println(v10)
    v5 = v5
    fmt.Println(v5, v1)
    v11 := []int{}
    v12 := v6[v1]
    v13 := GenS1{F0: 42, F1: 0}
    switch "" {
    case "\n\t":
        if false {
            var v14 uint = 0
            _ = v14
        } else {
            v12 = "a"
        }
    case "a":
        v4 = 42
    }
    v12 = "foo bar"
    for fi15 := 0; fi15 < 15; fi15++ {
        if true {
            if len(v11) > 0 {
                v11[0] = genFn0()
            }
            v16 := v3[(fi15 % ((fi15) | 1))]
            fmt.Println(fi15, v5, fi15)
            _ = v16
        } else {
            var v17 int32 = 860
            _ = v17
        }
        _ = fi15
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 7; fi19++ {
            v4 = 32767
            _ = fi19
        }
        fmt.Println(v1, v2, v1)
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 uint32 = v1
            switch v12 {
            case "a":
                if (false && (!true)) {
                    switch "a" {
                    case "a":
                        fmt.Println(v1)
                    case "":
                        v22 := GenS2{F0: 1000, F1: (v12 + (v12 + ((v12 + (v12 + (v12 + "hello"))) + ("" + (("hello" + "\n\t") + (v12 + ""))))))}
                        _ = v22
                    case "\n\t":
                        var v23 uint16 = 0
                        v5 = (v5 >> ((42) & 63))
                        v24 := v3[2147483647]
                        _ = v23
                        _ = v24
                    default:
                        v25 := GenS2{F0: -1, F1: v12}
                        _ = v25
                    }
                    v5 = int64(v2)
                    v2 -= v2
                } else {
                    var v26 uint32 = 711
                    v2 = 1
                    v26 -= 42
                    _ = v26
                }
            }
            fmt.Println(fi18, fi20, v2)
            _ = fi20
            _ = v21
        }
        _ = fi18
    }
    v10 = (3.14159 + v10)
    v5 = (v5 & 0)
    var v27 int64 = 1
    switch -1 {
    case 381:
        if (!true) {
            for fi28 := 0; fi28 < 19; fi28++ {
                v1 += uint32(v4)
                v12 = v12
                if (v10 == (v10 + -2.25)) {
                    v1 = v1
                    for fi29 := 0; fi29 < 20; fi29++ {
                        fmt.Println(v12)
                        _ = fi29
                    }
                }
                _ = fi28
            }
            v11 = append(v11, 386)
        }
    case 9223372036854775807:
        v11 = append(v11, 4294967295)
    default:
        v30 := GenS1{F0: 4294967295, F1: 0}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
