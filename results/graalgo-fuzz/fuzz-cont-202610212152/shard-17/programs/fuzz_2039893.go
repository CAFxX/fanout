package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int }
func genFn0() uint8 {
    v7 := GenS1{F0: 52, F1: 820, F2: -161}
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int16 = 454
        if (false && (true && false)) {
            v10 := []int{}
            _ = v10
        } else {
            var v11 bool = (true || false)
            v11 = (!(v11 && ((((("foo bar" + ("hello" + "foo bar")) + ("a" + (("" + ("" + ("a" + ((("\n\t" + "\n\t") + ("hello" + "hello")) + (("hello" + ("\n\t" + (("a" + ("foo bar" + "a")) + "foo bar"))) + "\n\t"))))) + "a"))) + "a") + "foo bar") <= "foo bar")))
            v11 = (127 <= 239)
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    var v12 uint64 = 42
    var v13 uint = 490
    _ = v12
    _ = v13
    _ = v7
    return 255
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 1
    v3 := make(map[uint32]int)
    v4 := make(map[string]int)
    v5 := []uint32{}
    v6 := make(map[uint32]int)
    var v14 bool = (!false)
    if len(v5) > 0 {
        v5[0] = 56
    }
    v14 = false
    var v15 int64 = 42
    v14 = true
    fmt.Println(v14)
    v15 = v15
    switch -2147483648 {
    case -173:
        if (1 >= 1) {
            var v16 uint8 = 42
            v17 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: 32768}
            _ = v16
            _ = v17
        } else {
            fmt.Println(v14, v15, v2)
            var v18 uint64 = 1
            _ = v18
        }
    case -9223372036854775808:
        if true {
            v2 -= (v2 % (((^(v2))) | 1))
            v2 = (v2 << ((2) & 63))
            var v19 uint = 4294967295
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 uint8 = 255
        v22 := GenS1{F0: 42, F1: 0, F2: fi20}
        var v23 int = (fi20 ^ 879)
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    if (127 <= 127) {
        var v24 int32 = 1
        v14 = (1.5 < (float64(-2.25) * 1e10))
        _ = v24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        v26 := GenS1{F0: 1, F1: 2, F2: 42}
        switch ("" + "\n\t") {
        case "foo bar":
            var v27 uint64 = 2
            _ = v27
        case "":
            v2 = v2
        case "a":
            for fi28 := 0; fi28 < 16; fi28++ {
                switch (fi25 << ((2) & 63)) {
                case 9223372036854775807:
                    v29 := GenS1{F0: 582, F1: 72, F2: (fi25 * -164)}
                    v15 = (-(v15))
                    _ = v29
                case 127:
                    var v30 float32 = (0.0 + 1.5)
                    _ = v30
                default:
                    v2 -= (v2 % ((v2) | 1))
                }
                v15 = 1
                v14 = v14
                _ = fi28
            }
        }
        _ = fi25
        _ = v26
    }
    v14 = v14
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
