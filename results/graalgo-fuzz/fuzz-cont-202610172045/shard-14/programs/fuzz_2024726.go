package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint8; F1 int }

func main() {
    var v1 string = "a"
    v2 := []uint32{42, 2147483648, 42, 0}
    var v3 uint8 = 143
    v4 := make(map[uint32]bool)
    v5 := make(map[int]int)
    if true {
        if ((1 > 611) && (42 > 255)) {
            v6 := GenS1{F0: (19 <= 244), F1: 744, F2: 1}
            _ = v6
        }
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(v3, fi7, v3)
            if (-60 == 127) {
                for fi8 := 0; fi8 < 8; fi8++ {
                    _ = len(v2)
                    _ = fi8
                }
            } else {
                switch 1 {
                case 32767:
                    for fi9 := 0; fi9 < 19; fi9++ {
                        var v10 bool = (65535 >= 120)
                        var v11 string = ((v1 + "\n\t") + v1)
                        v3 += v3
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                case 2:
                    delete(v5, fi7)
                }
            }
            for fi12 := 0; fi12 < 6; fi12++ {
                v3 = (v3 / ((127) | 1))
                var v13 uint64 = 500
                _ = fi12
                _ = v13
            }
            _ = fi7
        }
        v14 := GenS1{F0: true, F1: 2, F2: 0}
        _ = v14
    }
    var v15 uint64 = 2
    v3 += v3
    var v16 uint = 787
    v17 := GenS2{F0: v3, F1: 144}
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := v5[int(v3)]
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    v21 := GenS1{F0: false, F1: 42, F2: -32768}
    switch 128 {
    case 203:
        var v22 uint16 = 0
        _ = v22
    default:
        for fi23 := 0; fi23 < 1; fi23++ {
            v1 = "foo bar"
            var v24 bool = true
            fmt.Println(fi23, v15)
            _ = fi23
            _ = v24
        }
    }
    v15 = (v15 / ((337) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
