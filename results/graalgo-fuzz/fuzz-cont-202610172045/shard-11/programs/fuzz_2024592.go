package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int32; F2 int32 }

func main() {
    var v1 int32 = 124
    var v2 int32 = (v1 - (v1 % ((v1) | 1)))
    var v3 uint16 = 2
    var v4 uint = 0
    v5 := []string{""}
    v6 := make(map[string]int)
    fmt.Println(v3, v1)
    v6["foo bar"] = 1
    switch 32767 {
    case 830:
        for fi7 := 0; fi7 < 18; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 int16 = -983
                var v10 uint16 = (v3 + 1)
                v11 := GenS1{F0: 42}
                v10 = v3
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 int64 = 0
            _ = fi7
            _ = v12
        }
    case 42:
        v13 := GenS2{F0: 0, F1: (v1 * (-(v2))), F2: v2}
        _ = v13
    case 9223372036854775807:
        delete(v6, "foo bar")
    }
    v3 += v3
    if true {
        v1 += 42
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS2{F0: 47, F1: (-(v1)), F2: (v2 * (v2 | v2))}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        delete(v6, "hello")
        _ = len(v6)
        v17 := GenS1{F0: 9223372036854775807}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        for fi19 := 0; fi19 < 20; fi19++ {
            v1 = -406
            _ = fi19
        }
        _ = len(v6)
        for fi20 := 0; fi20 < 14; fi20++ {
            if (100 < 127) {
                fmt.Println(v1, v2)
                var v21 uint8 = 42
                var v22 float32 = (-2.25 * (3.14159 / ((1e10 / 0.1) * 1.5)))
                _ = v21
                _ = v22
            }
            delete(v6, ("foo bar" + "foo bar"))
            _ = fi20
        }
        _ = fi18
    }
    v23 := GenS1{F0: 42}
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            switch 42 {
            case -9223372036854775808:
                if len(v5) > 0 {
                    v5[0] = "foo bar"
                }
            default:
                var v26 uint8 = 4
                _ = v26
            }
            v1 = -787
            _ = fi25
        }
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
