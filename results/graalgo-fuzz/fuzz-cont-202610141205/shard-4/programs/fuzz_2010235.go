package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int }
func genFn0() int64 {
    var v6 int8 = 100
    for fi7 := 0; fi7 < 8; fi7++ {
        switch ("\n\t" + ("" + "foo bar")) {
        case "\n\t":
            var v8 uint64 = 0
            _ = v8
        }
        v9 := GenS1{F0: v6}
        fmt.Println(fi7, fi7)
        _ = fi7
        _ = v9
    }
    _ = v6
    return -1
}

func main() {
    var v1 uint = 1
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    v4 := make(map[uint32]int)
    v5 := []uint32{0}
    v1 = (v1 ^ (^(v1)))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v10 int16 = 42
    v10 = (v10 << ((v1) & 63))
    switch "hello" {
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "foo bar":
        v11 := GenS1{F0: -128}
        _ = v11
    case "\n\t":
        var v12 int16 = -752
        _ = v12
    }
    v1 = (v1 - 1)
    if len(v5) > 0 {
        v5[0] = 2
    }
    v10 = v10
    v13 := GenS1{F0: 25}
    if ("\n\t" > ("foo bar" + "a")) {
        v1 = v1
        for fi14 := 0; fi14 < 11; fi14++ {
            fmt.Println(v1)
            v1 = (v1 * 255)
            v1 = (v1 * 537)
            _ = fi14
        }
    } else {
        switch 32768 {
        case -439:
            for fi15 := 0; fi15 < 12; fi15++ {
                v1 -= 800
                _ = fi15
            }
        default:
            var v16 uint32 = 4294967295
            _ = v16
        }
        v10 = v10
        if len(v5) > 0 {
            v5[0] = 695
        }
    }
    if (((!(2147483647 == -289)) && true) || false) {
        v17 := GenS2{F0: 473, F1: 47, F2: 128}
        if false {
            if (true || ((!(0 < 42)) && (1 != 1))) {
                fmt.Println(v10, v10)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            switch 32767 {
            case 32768:
                _ = len(v4)
            case 32767:
                if (int8(v1) != 100) {
                    v10 = (^(v10))
                    var v18 float32 = -2.25
                    v19 := v4[880]
                    _ = v18
                    _ = v19
                } else {
                    fmt.Println(v10, v10)
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v10)
                    }
                    v20 := GenS1{F0: 21}
                    _ = v20
                }
            case 0:
                fmt.Println(v1, v1)
            }
            var v21 uint = v1
            _ = v21
        }
        v22 := GenS1{F0: -1}
        _ = v17
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
