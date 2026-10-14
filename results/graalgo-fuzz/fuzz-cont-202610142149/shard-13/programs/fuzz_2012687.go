package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int64 }
type GenS2 struct { F0 bool; F1 uint8; F2 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 0
    v3 := make(map[int]int)
    var v4 uint32 = 357
    for fi5 := 0; fi5 < 8; fi5++ {
        v2 += v2
        _ = fi5
    }
    v4 = (v4 | v4)
    v6 := []int{1, -9223372036854775808, 139}
    for fi7 := 0; fi7 < 20; fi7++ {
        v4 = v4
        for fi8 := 0; fi8 < 20; fi8++ {
            if (!true) {
                var v9 int32 = 0
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                switch fi8 {
                case 128:
                    var v10 float64 = 0.0
                    _ = v10
                }
                v4 -= v4
                _ = v9
            }
            v4 = v4
            switch 696 {
            case 0:
                v3[(fi8 + int(v4))] = 127
            case 320:
                v11 := GenS1{F0: "a", F1: 229, F2: 9223372036854775807}
                _ = v11
            default:
                fmt.Println(v4, fi8)
            }
            _ = fi8
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            v13 := GenS2{F0: (!(-5 > 100)), F1: 2, F2: -9223372036854775808}
            v14 := GenS2{F0: true, F1: 130, F2: -641}
            delete(v1, 32768)
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi7
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int64 = 1
            v18 := GenS2{F0: (true && true), F1: 71, F2: v17}
            v2 = v2
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    var v19 bool = false
    switch "\n\t" {
    case "":
        if true {
            _ = len(v6)
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 uint64 = 6
                if true {
                    var v22 int = -1
                    fmt.Println(v22, v19)
                    _ = v22
                } else {
                    v2 = v2
                    v23 := GenS1{F0: ("" + "\n\t"), F1: 255, F2: -1}
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
        }
    }
    if false {
        fmt.Println(v19)
        var v24 float32 = (-2.25 - 1e10)
        v4 -= (^(v4))
        _ = v24
    } else {
        fmt.Println(v2)
        v1[32767] = v19
    }
    v19 = (1 <= int16(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
