package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 string }
type GenS2 struct { F0 uint; F1 string; F2 uint64 }

func main() {
    var v1 string = ("\n\t" + ("\n\t" + "\n\t"))
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    var v4 uint16 = 397
    var v5 bool = false
    var v6 int32 = 42
    var v7 int64 = 0
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{}
        v10 := GenS1{F0: 65535, F1: (-(v7)), F2: ("a" + "\n\t")}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: (v4 ^ 255), F1: v7, F2: v1}
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS1{F0: v4, F1: 0, F2: ("" + v1)}
        v14 := GenS1{F0: v4, F1: -1, F2: v1}
        v15 := GenS2{F0: 4294967295, F1: v1, F2: 42}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    fmt.Println(v5, v6)
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 float64 = 0.0
        var v18 uint = 65535
        v18 -= v18
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        switch (fi19 / (((fi19 & fi19)) | 1)) {
        case -865:
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 int = (^(fi20))
                fmt.Println(v4, v21)
                fmt.Println(v7, v6, v5)
                _ = fi20
                _ = v21
            }
        case 551:
            v7 = (v7 / ((-1) | 1))
        case -656:
            var v22 bool = (v5 && v5)
            _ = v22
        }
        v5 = false
        v7 -= (v7 << ((900) & 63))
        _ = fi19
    }
    if v5 {
        v6 = v6
        if v5 {
            switch v1 {
            case "a":
                v23 := GenS2{F0: uint(v4), F1: "a", F2: 18446744073709551615}
                _ = v23
            case "foo bar":
                var v24 float64 = (1e10 * (((-2.25 - 1.5) - (-2.25 / -2.25)) + ((3.14159 * 1.5) / 3.14159)))
                _ = v24
            }
        }
    }
    v5 = ((v6 << ((210) & 63)) <= (^(v6)))
    for fi25 := 0; fi25 < 10; fi25++ {
        v4 = (v4 ^ (^(v4)))
        fmt.Println(v7, v1, v7)
        _ = fi25
    }
    v1 = ("hello" + "hello")
    switch 128 {
    case -2147483648:
        v6 = v6
    case 900:
        v26 := GenS2{F0: 0, F1: v1, F2: 2}
        _ = v26
    case -678:
        var v27 float64 = (float64(0.0) / -2.25)
        _ = v27
    default:
        for fi28 := 0; fi28 < 4; fi28++ {
            v4 = (v4 + v4)
            var v29 uint8 = 85
            _ = fi28
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        if v5 {
            if v5 {
                v31 := []int{2147483647}
                _ = v31
            }
        }
        var v32 uint16 = v4
        _ = fi30
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
