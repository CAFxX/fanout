package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }
type GenS2 struct { F0 uint; F1 uint8; F2 uint16 }

func main() {
    v1 := []uint32{1, 274}
    v2 := []uint32{42, 2147483647, 2}
    var v3 int16 = -1
    var v4 uint16 = 42
    v5 := make(map[string]uint32)
    var v6 uint64 = 42
    switch "a" {
    case "a":
        switch "" {
        case "\n\t":
            switch 0 {
            case 0:
                switch 634 {
                case 32767:
                    var v7 int16 = v3
                    _ = v7
                }
            case 32768:
                if false {
                    _ = len(v5)
                    v6 = v6
                    var v8 uint32 = 313
                    _ = v8
                } else {
                    var v9 float32 = 1e10
                    _ = v9
                }
            case -2147483648:
                v1 = append(v1, 65535)
            default:
                if len(v1) > 0 {
                    v1[0] = 278
                }
            }
        case "hello":
            v5[""] = 0
        default:
            v10 := GenS1{F0: 18446744073709551615, F1: 0, F2: 1}
            _ = v10
        }
    case "foo bar":
        for fi11 := 0; fi11 < 10; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 bool = ("a" < "foo bar")
                fmt.Println(fi11, v4, v3)
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    default:
        switch ("\n\t" + "a") {
        case "hello":
            v3 -= -32768
        default:
            v14 := v5[("hello" + "")]
            _ = v14
        }
    }
    v6 -= 222
    v15 := GenS2{F0: 65535, F1: uint8(v4), F2: 1}
    var v16 int8 = 127
    v17 := GenS2{F0: 42, F1: 182, F2: (v4 + 289)}
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 int32 = -1
        var v20 uint32 = 2147483648
        _ = fi18
        _ = v19
        _ = v20
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v21 int64 = -1
    v3 = (v3 % ((v3) | 1))
    fmt.Println(v16)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
