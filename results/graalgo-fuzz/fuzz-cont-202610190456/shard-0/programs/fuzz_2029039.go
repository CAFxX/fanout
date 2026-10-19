package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 128
    var v3 float32 = 1.5
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]uint32)
    var v6 int = -789
    v3 -= (v3 / 1e10)
    var v7 uint64 = 0
    if (545 < 743) {
        v3 -= -2.25
        if false {
            v8 := GenS1{F0: (v2 + v2), F1: v6}
            _ = v8
        }
    }
    switch (v6 ^ v6) {
    case 32768:
        fmt.Println(v2)
    case 1:
        v7 = 866
    }
    var v9 int16 = -768
    var v10 int16 = int16(v2)
    if (-9223372036854775808 > 914) {
        v3 = 1.5
    } else {
        switch (v6 - 127) {
        case -1:
            v3 = v3
        case -176:
            for fi11 := 0; fi11 < 17; fi11++ {
                v6 += v6
                v2 = v2
                var v12 int16 = v9
                _ = fi11
                _ = v12
            }
        case 42:
            for fi13 := 0; fi13 < 19; fi13++ {
                if true {
                    v3 -= v3
                    if true {
                        var v14 int8 = 112
                        v15 := GenS1{F0: v2, F1: (v6 | (fi13 % (((v6 << ((2) & 63))) | 1)))}
                        var v16 int64 = -803
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        var v17 uint16 = 216
                        for fi18 := 0; fi18 < 11; fi18++ {
                            var v19 uint64 = (v7 / (((v7 / (((^(v7))) | 1))) | 1))
                            _ = fi18
                            _ = v19
                        }
                        _ = v17
                    }
                    switch "hello" {
                    case "":
                        v3 = (((((3.14159 - 0.1) - (((v3 - (v3 / v3)) / v3) - 3.14159)) + (1.5 + v3)) * v3) + (v3 * 0.0))
                    }
                }
                _ = fi13
            }
        default:
            switch "hello" {
            case "":
                var v20 uint64 = v7
                _ = v20
            case "\n\t":
                v7 = (v7 ^ (v7 % ((1) | 1)))
            }
        }
        fmt.Println(v2, v2)
    }
    var v21 string = ""
    v21 = v21
    if (v3 > (v3 / v3)) {
        if false {
            v2 -= v2
            switch v6 {
            case -161:
                v22 := GenS1{F0: v2, F1: v6}
                _ = v22
            default:
                v7 = v7
            }
            fmt.Println(v21)
        } else {
            v23 := []int{v6, (-(v6)), (^(v6))}
            _ = v23
        }
    } else {
        fmt.Println(v7)
    }
    delete(v4, (v6 << ((65535) & 63)))
    fmt.Println(v6)
    for fi24 := 0; fi24 < 11; fi24++ {
        _ = len(v4)
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 uint8 = (v2 + v2)
            switch (v21 + "\n\t") {
            case "a":
                v3 += v3
            case "foo bar":
                v27 := []int{fi24, (^(fi24)), (v6 | 0)}
                _ = v27
            case "":
                v26 += 58
            default:
                switch (v6 % ((128) | 1)) {
                case 42:
                    var v28 uint32 = 1
                    _ = v28
                case 0:
                    _ = len(v1)
                case 623:
                    v21 = "hello"
                }
            }
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    if true {
        var v29 int32 = -2147483648
        v30 := GenS2{F0: v7}
        _ = v29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
