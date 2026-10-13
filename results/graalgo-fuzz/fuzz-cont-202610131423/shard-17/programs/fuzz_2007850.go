package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int8; F1 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 284
    var v3 string = (("" + "hello") + ("" + ("foo bar" + "\n\t")))
    var v4 int16 = -65
    var v5 float32 = (0.0 + 1.5)
    var v6 uint32 = (^(v2))
    var v7 int8 = 27
    v8 := GenS1{F0: (v7 - (v7 & v7)), F1: 2147483647}
    if false {
        v9 := []int{32768, -1, -9223372036854775808}
        var v10 int = 1
        if (((v7 * int8(v10)) < 0) && (((v5 - (float32(float64(float32(1e10))) - ((0.1 - 1.5) + -2.25))) - v5) > 1e10)) {
            if true {
                switch "hello" {
                case "foo bar":
                    for fi11 := 0; fi11 < 19; fi11++ {
                        v12 := GenS2{F0: 42, F1: 280}
                        fmt.Println(v4)
                        _ = fi11
                        _ = v12
                    }
                }
            } else {
                switch "" {
                case "hello":
                    v7 = v7
                case "":
                    var v13 float32 = ((v5 + v5) * 3.14159)
                    _ = v13
                }
            }
            v14 := GenS1{F0: 23, F1: 2147483647}
            _ = v14
        } else {
            var v15 uint16 = 1
            _ = v15
        }
        _ = v10
        _ = v9
    }
    var v16 int32 = -2147483648
    switch (v3 + (((((v3 + (v3 + v3)) + v3) + ("\n\t" + (((v3 + v3) + "") + v3))) + ("\n\t" + v3)) + ((v3 + (v3 + "")) + "foo bar"))) {
    case "a":
        v17 := []int{-8, -2147483648, 4294967295, 32767}
        _ = v17
    case "foo bar":
        if (false && (!false)) {
            v1[0] = true
        } else {
            v18 := GenS1{F0: (v7 + (v7 * v7)), F1: -549}
            v1[32768] = false
            _ = v18
        }
    case "":
        for fi19 := 0; fi19 < 14; fi19++ {
            for fi20 := 0; fi20 < 10; fi20++ {
                var v21 string = v3
                _ = fi20
                _ = v21
            }
            v1[fi19] = false
            _ = fi19
        }
    default:
        v6 -= v2
    }
    v22 := GenS2{F0: v7, F1: 0}
    v2 = 4294967295
    for fi23 := 0; fi23 < 19; fi23++ {
        for fi24 := 0; fi24 < 19; fi24++ {
            switch (fi24 | int(v4)) {
            case 1:
                var v25 uint16 = 42
                _ = v25
            }
            var v26 int64 = -376
            var v27 uint32 = 42
            _ = fi24
            _ = v26
            _ = v27
        }
        if ((v4 / (((^(v4))) | 1)) > -501) {
            v3 = v3
            fmt.Println(v6, v5, v6)
            v28 := []int{(-(fi23))}
            _ = v28
        }
        _ = fi23
    }
    var v29 string = v3
    v2 += (^(v2))
    v6 -= (v2 & v6)
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 uint32 = (v6 - (v6 + 4294967295))
        _ = fi30
        _ = v31
    }
    v16 += v16
    for fi32 := 0; fi32 < 13; fi32++ {
        if (818 <= -1) {
            var v33 int = fi32
            _ = v33
        }
        v4 = 1
        for fi34 := 0; fi34 < 20; fi34++ {
            var v35 float32 = v5
            fmt.Println(v2, v2)
            var v36 uint16 = 0
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi32
    }
    fmt.Println(v16, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
