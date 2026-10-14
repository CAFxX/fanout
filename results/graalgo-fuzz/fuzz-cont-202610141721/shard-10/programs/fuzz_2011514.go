package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int; F1 bool }
func genFn0() bool {
    var v5 uint32 = 569
    v5 = v5
    v6 := GenS1{F0: -1, F1: 1}
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 bool, p1 int, p2 uint32) uint32 {
    v7 := GenS1{F0: 99, F1: 1}
    p1 = (p1 + (-(p1)))
    if (p1 != 153) {
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS1{F0: 30, F1: 1}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 108, F1: 889}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{430, 42, 885, 1}
    v3 := make(map[string]int)
    var v4 uint = 820
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch "hello" {
    case "":
        if false {
            for fi11 := 0; fi11 < 16; fi11++ {
                if len(v2) > 0 {
                    v2[0] = genFn1(false, -672, 481)
                }
                _ = fi11
            }
            v1 -= (^(v1))
            for fi12 := 0; fi12 < 1; fi12++ {
                v1 -= v1
                var v13 string = "foo bar"
                var v14 uint8 = 2
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    case "hello":
        v1 += (^(v1))
    }
    if ((v1 ^ 131) >= v1) {
        if true {
            var v15 uint16 = 1
            switch ("" + "\n\t") {
            case "":
                if (553 != 300) {
                    v4 = v4
                    v3["a"] = -875
                } else {
                    var v16 int8 = 0
                    var v17 uint64 = v1
                    _ = v16
                    _ = v17
                }
            case "foo bar":
                v18 := GenS1{F0: -32, F1: 1000}
                _ = v18
            case "hello":
                v19 := GenS1{F0: 71, F1: -942}
                _ = v19
            }
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 uint8 = 127
                _ = fi20
                _ = v21
            }
            _ = v15
        }
    } else {
        v22 := GenS1{F0: -1, F1: 42}
        switch -918 {
        case 32768:
            for fi23 := 0; fi23 < 20; fi23++ {
                v24 := GenS1{F0: 127, F1: 120}
                _ = fi23
                _ = v24
            }
        case -1:
            fmt.Println(v4)
        default:
            v2 = append(v2, 42)
        }
        _ = v22
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        v26 := GenS1{F0: 42, F1: 32767}
        v27 := GenS2{F0: fi25, F1: true}
        _ = fi25
        _ = v26
        _ = v27
    }
    v4 -= v4
    if len(v2) > 0 {
        v2[0] = 31
    }
    fmt.Println(v4, v1)
    var v28 int = -267
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
