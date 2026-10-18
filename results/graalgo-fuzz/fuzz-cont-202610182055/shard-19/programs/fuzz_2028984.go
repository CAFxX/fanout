package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 string; F1 int8; F2 string }
func genFn0() uint8 {
    var v4 uint8 = 42
    var v5 float64 = 1.5
    v5 -= 0.0
    v6 := GenS1{F0: 1, F1: 268}
    _ = v4
    _ = v5
    _ = v6
    return v4
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 4294967295
    var v3 int8 = 115
    var v7 uint32 = 0
    for fi8 := 0; fi8 < 18; fi8++ {
        if false {
            delete(v1, "a")
            switch fi8 {
            case 127:
                var v9 uint32 = 2
                _ = v9
            case -807:
                if (1.5 <= 3.14159) {
                    var v10 int16 = int16(v2)
                    v2 = (v2 % ((v2) | 1))
                    _ = v10
                } else {
                    var v11 uint16 = 2
                    _ = v11
                }
            case 578:
                v12 := GenS1{F0: 157, F1: -32768}
                _ = v12
            }
        } else {
            v3 -= (v3 << ((421) & 63))
            var v13 uint = 255
            if (!(1e10 > ((1e10 - 1.5) - float32(0.0)))) {
                v14 := GenS1{F0: v13, F1: -32768}
                switch "" {
                case "hello":
                    var v15 float64 = 1.5
                    _ = v15
                case "foo bar":
                    var v16 uint32 = (v7 * 836)
                    _ = v16
                case "":
                    v17 := GenS1{F0: 308, F1: -1}
                    _ = v17
                }
                _ = v14
            }
            _ = v13
        }
        switch (fi8 % ((fi8) | 1)) {
        case 127:
            var v18 int8 = (v3 ^ (-(v3)))
            _ = v18
        case 1:
            var v19 uint8 = 2
            _ = v19
        case -897:
            v3 = (v3 | int8(v7))
        }
        if false {
            var v20 string = (("\n\t" + "hello") + "a")
            _ = v20
        }
        _ = fi8
    }
    fmt.Println(v3)
    v21 := GenS2{F0: "", F1: -47, F2: "a"}
    fmt.Println(v2, v2)
    v3 += v3
    if false {
        fmt.Println(v3, v2, v3)
        for fi22 := 0; fi22 < 12; fi22++ {
            if false {
                fmt.Println(v7, fi22)
            } else {
                fmt.Println(v7)
            }
            v7 = (v7 % (((v7 % ((uint32(v2)) | 1))) | 1))
            var v23 uint = 18446744073709551615
            _ = fi22
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        v25 := GenS1{F0: 42, F1: -827}
        v26 := GenS2{F0: "\n\t", F1: 42, F2: "hello"}
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 19; fi27++ {
        var v28 int16 = 400
        _ = fi27
        _ = v28
    }
    v2 -= v2
    v3 = (v3 - v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v7
}
