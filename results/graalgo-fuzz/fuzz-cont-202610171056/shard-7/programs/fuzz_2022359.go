package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    v3 := []string{("a" + "\n\t"), ("" + "a"), (("hello" + "hello") + "hello"), ("hello" + ("foo bar" + ("\n\t" + (("foo bar" + "") + ("hello" + "a")))))}
    var v4 uint32 = 403
    if false {
        fmt.Println(v4)
        if true {
            v4 = (v4 * (v4 | (^(v4))))
        }
        fmt.Println(v4)
    } else {
        var v5 uint64 = 970
        var v6 float64 = ((float64(float32(((1e10 / (1e10 + (0.1 / 1e10))) + 1.5))) + -2.25) + 3.14159)
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 1, F1: 42}
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 int16 = -32768
        _ = fi8
        _ = v9
    }
    fmt.Println(v4)
    var v10 int16 = 1000
    if (-1 > 1) {
        v4 = v4
        v11 := v2["foo bar"]
        _ = v11
    } else {
        v4 = v4
        fmt.Println(v4)
        for fi12 := 0; fi12 < 2; fi12++ {
            v4 = 27
            var v13 int64 = 9223372036854775807
            _ = fi12
            _ = v13
        }
    }
    v14 := GenS1{F0: 42, F1: 0}
    if true {
        for fi15 := 0; fi15 < 16; fi15++ {
            v10 = v10
            _ = fi15
        }
        v10 = (v10 & v10)
        var v16 uint32 = v4
        _ = v16
    } else {
        for fi17 := 0; fi17 < 3; fi17++ {
            v10 += -1
            _ = fi17
        }
    }
    if false {
        delete(v2, "a")
    } else {
        v10 += 275
        switch "hello" {
        case "hello":
            switch 877 {
            case 32768:
                fmt.Println(v10, v10)
            case -9223372036854775808:
                v18 := GenS1{F0: 255, F1: (v10 / (((v10 >> ((2) & 63))) | 1))}
                _ = v18
            case 127:
                fmt.Println(v10, v10)
            }
        }
        fmt.Println(v10, v10)
    }
    if ((true && true) || true) {
        v19 := GenS2{F0: 0}
        _ = v19
    } else {
        switch -495 {
        case 144:
            var v20 float32 = 3.14159
            _ = v20
        case -665:
            v4 -= 2147483647
        case 2147483647:
            for fi21 := 0; fi21 < 18; fi21++ {
                var v22 uint16 = uint16(v4)
                _ = fi21
                _ = v22
            }
        default:
            var v23 uint32 = 434
            _ = v23
        }
    }
    if (927 >= 9223372036854775807) {
        v4 += v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
