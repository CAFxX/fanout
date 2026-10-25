package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int16; F2 int8 }
type GenS4 struct { F0 bool; F1 bool }
func genFn0(p0 uint32, p1 string, p2 int64) uint64 {
    if true {
        p0 = 4294967295
    }
    var v5 uint8 = 30
    if (1 == 825) {
        var v6 int16 = 547
        var v7 float32 = float32(1.5)
        var v8 float32 = v7
        _ = v6
        _ = v7
        _ = v8
    } else {
        if true {
            for fi9 := 0; fi9 < 11; fi9++ {
                fmt.Println(fi9)
                _ = fi9
            }
        } else {
            for fi10 := 0; fi10 < 1; fi10++ {
                for fi11 := 0; fi11 < 14; fi11++ {
                    v12 := GenS1{F0: 856, F1: 6, F2: (^(p0))}
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
        }
        if (42 > p2) {
            if ((-(p2)) < (p2 * 9223372036854775807)) {
                switch (p1 + ((p1 + p1) + (("foo bar" + "foo bar") + ((p1 + ("\n\t" + p1)) + ((p1 + p1) + "foo bar"))))) {
                case "hello":
                    p2 = p2
                case "a":
                    p0 -= p0
                }
            }
        }
        var v13 float64 = (float64(1.5) - 3.14159)
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]int)
    var v3 int8 = 100
    var v4 uint64 = 113
    fmt.Println(v4, v1)
    v14 := GenS1{F0: 0, F1: v4, F2: 298}
    switch 2147483647 {
    case -9223372036854775808:
        var v15 int32 = -2147483648
        _ = v15
    case 32768:
        v16 := GenS1{F0: 2147483647, F1: (v4 ^ genFn0(42, ("" + "hello"), -282)), F2: 932}
        _ = v16
    }
    v17 := GenS2{F0: 0}
    switch "" {
    case "foo bar":
        fmt.Println(v4, v1)
    case "a":
        switch 310 {
        case 32768:
            for fi18 := 0; fi18 < 20; fi18++ {
                v4 -= 2
                _ = fi18
            }
        default:
            v1 -= (^(v1))
        }
    }
    var v19 string = "hello"
    var v20 uint64 = v4
    for fi21 := 0; fi21 < 17; fi21++ {
        v20 += (v4 << ((986) & 63))
        v22 := GenS1{F0: 2147483647, F1: v4, F2: 2147483647}
        if (3.14159 > (1.5 + 1.5)) {
            delete(v2, "\n\t")
            switch fi21 {
            case 127:
                for fi23 := 0; fi23 < 18; fi23++ {
                    v24 := GenS3{F0: (fi21 >> ((1) & 63)), F1: 42, F2: (v3 | (v3 | 1))}
                    _ = fi23
                    _ = v24
                }
            case 9223372036854775807:
                v25 := GenS4{F0: ((2147483648 != 1) && (!(1 >= 127))), F1: ("" >= ((("a" + "") + "") + ("hello" + ((v19 + v19) + v19))))}
                _ = v25
            default:
                var v26 float64 = 3.14159
                _ = v26
            }
        } else {
            v20 = genFn0(uint32(v20), "hello", -755)
        }
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
