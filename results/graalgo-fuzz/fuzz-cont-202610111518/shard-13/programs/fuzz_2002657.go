package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint) uint8 {
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    fmt.Println(p1, p1)
    p1 += p1
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 255
}

func main() {
    v1 := []string{"", "foo bar", "foo bar"}
    var v2 int64 = -566
    v3 := make(map[string]uint32)
    var v4 int = 9223372036854775807
    var v5 uint32 = 1
    var v7 uint8 = 91
    v4 = (v4 + 2147483647)
    switch ("foo bar" + (("\n\t" + "hello") + "\n\t")) {
    case "a":
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 int = 32767
            switch v9 {
            case 9223372036854775807:
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 int16 = -32768
                    var v12 int8 = 0
                    var v13 uint16 = 551
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            case -270:
                v9 = v4
            default:
                if true {
                    for fi14 := 0; fi14 < 6; fi14++ {
                        v5 = 2147483647
                        _ = fi14
                    }
                    fmt.Println(v9, v2)
                } else {
                    var v15 bool = false
                    _ = v15
                }
            }
            _ = fi8
            _ = v9
        }
    case "foo bar":
        delete(v3, "")
    }
    var v16 uint32 = (v5 / ((v5) | 1))
    switch "foo bar" {
    case "":
        var v17 uint64 = 1
        _ = v17
    case "\n\t":
        for fi18 := 0; fi18 < 13; fi18++ {
            v5 = v5
            var v19 uint32 = (v5 ^ (^(v5)))
            v5 = (v5 + v19)
            _ = fi18
            _ = v19
        }
    case "hello":
        if (!(2147483647 < -1)) {
            v20 := v3["\n\t"]
            fmt.Println(v4, v20, v16)
            _ = v20
        }
    default:
        v2 -= int64(v16)
    }
    v7 = (^(v7))
    v4 += v4
    v16 -= v5
    fmt.Println(v5, v2, v2)
    var v21 float64 = 3.14159
    switch v4 {
    case 128:
        v4 -= 128
    case 42:
        switch v4 {
        case 4294967295:
            for fi22 := 0; fi22 < 6; fi22++ {
                v23 := GenS1{F0: true, F1: 127, F2: (^(v16))}
                if true {
                    v24 := GenS1{F0: (!(4294967295 > 18446744073709551615)), F1: 127, F2: 42}
                    _ = v24
                }
                var v25 int32 = -306
                _ = fi22
                _ = v23
                _ = v25
            }
        case 42:
            if true {
                v3[("" + "hello")] = 538
            }
        case 680:
            var v26 uint64 = 42
            _ = v26
        default:
            v5 += v16
        }
    case -2147483648:
        v27 := GenS2{F0: 42}
        _ = v27
    default:
        var v28 int = (v4 ^ 2147483647)
        _ = v28
    }
    if (42 >= v16) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v4 -= v4
    delete(v3, "a")
    if true {
        var v29 int32 = -31
        _ = v29
    } else {
        v7 = (v7 - v7)
        var v30 int16 = 746
        _ = v30
    }
    if (360 >= 0) {
        var v31 bool = (true || true)
        v32 := GenS1{F0: true, F1: 0, F2: (v16 + (^(v5)))}
        v33 := GenS3{F0: 746}
        _ = v31
        _ = v32
        _ = v33
    }
    v34 := GenS2{F0: 92}
    fmt.Println(v7, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v7
}
