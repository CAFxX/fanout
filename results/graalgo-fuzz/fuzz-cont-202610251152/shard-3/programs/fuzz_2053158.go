package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 0.0
    var v3 int = 1
    var v4 uint8 = 127
    var v5 uint64 = 42
    v6 := GenS1{F0: (^(v4)), F1: (!(42 >= 2147483647)), F2: 841}
    v1[("a" + ("foo bar" + "foo bar"))] = true
    fmt.Println(v3)
    if (1 == 429) {
        switch ((("a" + "a") + "a") + "foo bar") {
        case "hello":
            v4 -= 0
        case "a":
            if (!((65535 > 42) && false)) {
                v7 := GenS1{F0: v4, F1: false, F2: 275}
                var v8 string = "a"
                _ = v7
                _ = v8
            }
        default:
            switch v3 {
            case 1:
                v9 := []int{}
                _ = v9
            default:
                var v10 string = "\n\t"
                _ = v10
            }
        }
        v11 := []int{}
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 uint = 4294967295
            if (1 < 607) {
                v14 := GenS1{F0: v4, F1: ((false || (v13 > v13)) && true), F2: v13}
                var v15 bool = (-146 <= 1000)
                _ = v14
                _ = v15
            } else {
                switch fi12 {
                case 470:
                    var v16 float32 = float32(1e10)
                    _ = v16
                }
                var v17 int64 = 0
                fmt.Println(v17)
                _ = v17
            }
            _ = fi12
            _ = v13
        }
        _ = v11
    }
    v4 = v4
    _ = len(v1)
    if false {
        switch ("\n\t" + "\n\t") {
        case "hello":
            for fi18 := 0; fi18 < 7; fi18++ {
                fmt.Println(v3, fi18, v2)
                _ = fi18
            }
        case "a":
            fmt.Println(v4)
        case "foo bar":
            fmt.Println(v4)
        default:
            v19 := GenS1{F0: (v4 - (^(v4))), F1: false, F2: 0}
            _ = v19
        }
        v20 := GenS1{F0: v4, F1: ("foo bar" != "hello"), F2: 624}
        var v21 uint32 = 65535
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 13; fi22++ {
            v4 += v4
            _ = fi22
        }
        v4 = (v4 | 255)
        var v23 float64 = v2
        _ = v23
    }
    v5 += (^(v5))
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 uint8 = (v4 << ((499) & 63))
        _ = fi24
        _ = v25
    }
    var v26 uint16 = 65535
    var v27 uint64 = 4294967295
    v2 = 0.0
    v28 := GenS1{F0: v4, F1: (32 != 128), F2: 4294967295}
    var v29 uint32 = 394
    var v30 float64 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
