package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int64; F1 int; F2 uint8 }
func genFn0(p0 int16, p1 uint8, p2 uint16) int16 {
    v7 := GenS1{F0: 723, F1: (!false)}
    fmt.Println(p2, p1)
    v8 := GenS1{F0: 2147483647, F1: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ("hello" + ("\n\t" + "\n\t"))
    var v3 string = v2
    v4 := []int{}
    v5 := []uint32{4294967295, 42, 0, 2147483648}
    var v6 uint16 = 1
    var v9 float32 = float32(3.14159)
    _ = len(v1)
    v10 := GenS1{F0: 65535, F1: true}
    v9 = v9
    var v11 int8 = 100
    if true {
        if (!(100 != (^(v11)))) {
            if false {
                if (0 < -1) {
                    var v12 int32 = -386
                    fmt.Println(v9, v2)
                    _ = v12
                }
            }
        }
    }
    var v13 int64 = -1
    switch v2 {
    case "hello":
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 int64 = (v13 - (v13 & v13))
            _ = fi14
            _ = v15
        }
    case "":
        var v16 int8 = 52
        _ = v16
    case "\n\t":
        if false {
            var v17 float64 = (1.5 * 0.1)
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi18 := 0; fi18 < 7; fi18++ {
                _ = len(v1)
                v19 := GenS1{F0: 0, F1: false}
                _ = fi18
                _ = v19
            }
            _ = v17
        }
    default:
        switch 9223372036854775807 {
        case 947:
            v11 = (v11 ^ 42)
        case -307:
            v20 := GenS1{F0: 181, F1: true}
            _ = v20
        case -240:
            for fi21 := 0; fi21 < 11; fi21++ {
                if (float64(v9) == (((((1e10 - float64((v9 + -2.25))) / 0.1) - -2.25) - (3.14159 * (0.0 / 3.14159))) - 1e10)) {
                    switch ((("\n\t" + v3) + "") + ((("" + "foo bar") + "hello") + "a")) {
                    case "foo bar":
                        v13 = (v13 >> ((279) & 63))
                    default:
                        v1[(-(fi21))] = 630
                    }
                }
                fmt.Println(fi21, v9, v3)
                _ = fi21
            }
        }
    }
    v22 := GenS1{F0: 2147483647, F1: ((("hello" < (v2 + v2)) && true) && (!false))}
    if (2147483648 >= 2) {
        if true {
            for fi23 := 0; fi23 < 17; fi23++ {
                for fi24 := 0; fi24 < 13; fi24++ {
                    var v25 uint8 = 42
                    if (true && false) {
                        var v26 uint8 = v25
                        v27 := GenS2{F0: (-(v13)), F1: fi23, F2: v25}
                        _ = v26
                        _ = v27
                    }
                    _ = fi24
                    _ = v25
                }
                v11 += (v11 - v11)
                _ = fi23
            }
        } else {
            v3 = v3
            for fi28 := 0; fi28 < 5; fi28++ {
                v29 := GenS2{F0: v13, F1: (fi28 * (fi28 / ((fi28) | 1))), F2: 159}
                var v30 int64 = v13
                _ = fi28
                _ = v29
                _ = v30
            }
        }
        v11 = (-(v11))
        fmt.Println(v11, v6, v9)
    }
    fmt.Println(v6, v11, v11)
    v31 := GenS2{F0: (^(v13)), F1: 273, F2: 128}
    for fi32 := 0; fi32 < 19; fi32++ {
        var v33 int16 = -76
        fmt.Println(v3)
        _ = fi32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
