package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint }
func genFn0(p0 int16) uint64 {
    v7 := []int{}
    if (true || false) {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 bool = false
            switch "a" {
            case "\n\t":
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case "a":
                v9 = v9
            case "hello":
                var v10 int64 = -834
                _ = v10
            }
            _ = fi8
            _ = v9
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            fmt.Println(fi11, p0)
            p0 -= (p0 - 0)
            _ = fi11
        }
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 int16) uint64 {
    p0 -= (p0 ^ p0)
    var v12 uint16 = 65535
    _ = p0
    _ = v12
    return genFn0(-1)
}

func main() {
    var v1 int64 = -673
    var v2 int = 128
    var v3 int64 = v1
    var v4 uint64 = 475
    var v5 uint32 = 4294967295
    var v6 int = (v2 ^ v2)
    v13 := GenS1{F0: 706, F1: 42, F2: 255}
    v4 -= 42
    var v14 float32 = -2.25
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 float64 = 1.5
        fmt.Println(v5, v2)
        switch "" {
        case "foo bar":
            switch "hello" {
            case "\n\t":
                switch v6 {
                case -636:
                    var v17 int16 = -370
                    _ = v17
                case 300:
                    for fi18 := 0; fi18 < 20; fi18++ {
                        fmt.Println(v4, v6, v14)
                        v19 := GenS1{F0: 255, F1: uint16(v4), F2: 1}
                        v14 = (float32(3.14159) / 3.14159)
                        _ = fi18
                        _ = v19
                    }
                case -961:
                    var v20 int32 = -1
                    _ = v20
                default:
                    var v21 uint8 = 255
                    _ = v21
                }
            case "":
                v22 := GenS1{F0: 0, F1: 32767, F2: 42}
                _ = v22
            case "a":
                var v23 uint = 42
                _ = v23
            }
        case "hello":
            for fi24 := 0; fi24 < 4; fi24++ {
                v5 = 254
                _ = fi24
            }
        case "\n\t":
            v25 := GenS1{F0: 2, F1: uint16(v5), F2: 2}
            _ = v25
        default:
            var v26 float32 = (v14 - float32((1e10 - ((0.0 * -2.25) / v16))))
            _ = v26
        }
        _ = fi15
        _ = v16
    }
    v27 := GenS1{F0: 255, F1: 647, F2: 65535}
    var v28 float32 = v14
    if (237 != 1) {
        if true {
            v29 := GenS1{F0: 598, F1: uint16(v4), F2: uint16(v4)}
            var v30 int16 = 42
            v6 = (-(v2))
            _ = v29
            _ = v30
        }
    } else {
        var v31 uint16 = 1
        v32 := GenS1{F0: 4294967295, F1: (v31 % ((uint16(v4)) | 1)), F2: v31}
        if ((!true) || ((v2 % ((-1) | 1)) < -145)) {
            v33 := GenS2{F0: uint8(v4), F1: uint(v5)}
            _ = v33
        }
        _ = v31
        _ = v32
    }
    fmt.Println(v5, v5, v5)
    var v34 int = (v6 - int(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
