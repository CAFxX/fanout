package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 bool = (!true)
    var v2 int16 = 1
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    v5 := []int{}
    v6 := make(map[int]string)
    if v1 {
        v1 = (42 < 2147483647)
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int8 = 47
            v9 := GenS1{F0: 911, F1: 0, F2: 1}
            _ = fi7
            _ = v8
            _ = v9
        }
        if (!(!(!v1))) {
            var v10 int16 = (^(v2))
            _ = v10
        }
    }
    var v11 int64 = 42
    v2 -= (v2 / ((v2) | 1))
    var v12 uint = 42
    if v1 {
        v2 -= (v2 * (^(v2)))
        v11 += -9223372036854775808
        var v13 int16 = 1000
        _ = v13
    } else {
        v1 = v1
        var v14 int16 = v2
        switch ("foo bar" + "foo bar") {
        case "\n\t":
            switch 32768 {
            case 128:
                v11 -= v11
            }
        default:
            v12 -= 255
        }
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        if (1e10 >= 3.14159) {
            for fi16 := 0; fi16 < 1; fi16++ {
                switch (((((("a" + "") + "\n\t") + "") + "hello") + "\n\t") + "") {
                case "a":
                    v17 := v6[(fi16 & fi16)]
                    _ = v17
                case "foo bar":
                    v18 := GenS1{F0: 874, F1: -305, F2: v11}
                    _ = v18
                case "hello":
                    var v19 uint64 = 1
                    _ = v19
                }
                _ = fi16
            }
            for fi20 := 0; fi20 < 3; fi20++ {
                delete(v6, fi15)
                _ = fi20
            }
            var v21 int16 = -32768
            _ = v21
        }
        var v22 uint32 = 4294967295
        for fi23 := 0; fi23 < 17; fi23++ {
            if (false || v1) {
                var v24 bool = (!false)
                v11 = v11
                var v25 uint64 = 4294967295
                _ = v24
                _ = v25
            }
            v2 -= int16(v22)
            v26 := GenS1{F0: 4294967295, F1: fi23, F2: (v11 | 160)}
            _ = fi23
            _ = v26
        }
        _ = fi15
        _ = v22
    }
    switch "a" {
    case "":
        v11 = v11
    case "hello":
        if v1 {
            v27 := GenS2{F0: 488}
            if v1 {
                v28 := GenS2{F0: 4294967295}
                _ = v28
            } else {
                switch "\n\t" {
                case "foo bar":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
            }
            if len(v5) > 0 {
                v5[0] = 32767
            }
            _ = v27
        }
    case "foo bar":
        var v29 string = "\n\t"
        _ = v29
    default:
        var v30 bool = v1
        _ = v30
    }
    v1 = v1
    v2 -= (v2 + (v2 ^ 673))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
