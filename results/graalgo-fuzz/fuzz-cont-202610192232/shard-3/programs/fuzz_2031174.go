package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
func genFn0(p0 int16, p1 uint32, p2 bool) string {
    var v7 uint8 = 241
    var v8 int64 = 806
    if p2 {
        fmt.Println(v7, p0, p2)
        var v9 int8 = -24
        v10 := []int{0}
        _ = v10
        _ = v9
    }
    var v11 uint = 729
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    var v1 uint64 = 209
    var v2 uint16 = 232
    v3 := make(map[int]int)
    v4 := make(map[uint32]uint32)
    var v5 int64 = 0
    var v6 int = -467
    var v12 float64 = 0.0
    _ = len(v4)
    v13 := []int{}
    if true {
        _ = len(v13)
        var v14 int16 = 1000
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 15; fi16++ {
            switch (fi15 << ((470) & 63)) {
            case 514:
                var v17 string = "foo bar"
                _ = v17
            case -1:
                if false {
                    var v18 uint8 = 255
                    if false {
                        v1 -= v1
                        v19 := v4[885]
                        v5 = v5
                        v20 := GenS1{F0: "\n\t", F1: v5}
                        _ = v19
                        _ = v20
                    }
                    _ = v18
                } else {
                    v6 = (fi15 & 2147483647)
                    var v21 float64 = (v12 / 1.0)
                    var v22 uint = uint(v1)
                    _ = v21
                    _ = v22
                }
            case 1:
                v12 = v12
            }
            v5 += v5
            v12 = float64((0.0 + float32(1.5)))
            _ = fi16
        }
        _ = len(v3)
        v23 := GenS1{F0: ("" + "a"), F1: v5}
        _ = fi15
        _ = v23
    }
    var v24 float64 = 1.5
    v25 := v4[65535]
    if ("a" < "hello") {
        var v26 int16 = 1000
        _ = v26
    }
    var v27 float32 = 3.14159
    var v28 uint16 = (^(v2))
    for fi29 := 0; fi29 < 5; fi29++ {
        fmt.Println(v27, v25)
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = fi29
    }
    fmt.Println(v25, v25, v27)
    v5 = (v5 - -135)
    v5 -= (-(v5))
    if ((42 < 921) || (true || (v25 != 2))) {
        switch ("\n\t" + "foo bar") {
        case "":
            var v30 uint64 = v1
            _ = v30
        case "foo bar":
            switch ("\n\t" + "a") {
            case "":
                v6 -= v6
            case "\n\t":
                v31 := GenS1{F0: "a", F1: int64(v25)}
                _ = v31
            }
        case "a":
            if false {
                _ = len(v13)
                v32 := GenS1{F0: ("" + ""), F1: (v5 >> ((18446744073709551615) & 63))}
                _ = v32
            }
        default:
            var v33 int8 = int8(v6)
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
