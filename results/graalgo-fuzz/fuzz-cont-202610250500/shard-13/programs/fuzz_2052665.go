package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int16; F1 int; F2 uint64 }

func main() {
    var v1 bool = (false && (!false))
    var v2 bool = v1
    var v3 uint16 = 65535
    var v4 int64 = 9223372036854775807
    v5 := []uint32{345, 0, 83}
    v6 := make(map[string]uint32)
    v1 = (82 <= 127)
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    var v7 uint8 = 184
    v2 = (!true)
    if (!true) {
        switch 584 {
        case 128:
            v8 := v6["hello"]
            _ = v8
        case 1:
            if ((v1 && v1) || v2) {
                v3 = (v3 << ((42) & 63))
                for fi9 := 0; fi9 < 5; fi9++ {
                    v1 = ("foo bar" > "")
                    _ = fi9
                }
                if (!((v3 * 0) < (v3 | v3))) {
                    v2 = true
                }
            }
        case -9223372036854775808:
            var v10 float64 = ((1e10 * 1e10) * 3.14159)
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            fmt.Println(v3, v4, v1)
            var v13 int64 = v4
            if v2 {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v4 = -1
                switch fi11 {
                case 0:
                    var v14 bool = v1
                    _ = v14
                case 128:
                    var v15 uint32 = 65535
                    _ = v15
                default:
                    var v16 int16 = -1
                    _ = v16
                }
            }
            _ = fi12
            _ = v13
        }
        switch fi11 {
        case 32768:
            fmt.Println(v4)
        case 128:
            var v17 int8 = 127
            _ = v17
        default:
            v7 -= (^(v7))
        }
        if (v1 || (0 <= -2147483648)) {
            v4 += -1
            v2 = (v2 && v1)
        }
        _ = fi11
    }
    v5 = append(v5, 217)
    v5 = append(v5, 758)
    var v18 int8 = 42
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 uint64 = 42
        fmt.Println(v20)
        v7 = 128
        _ = fi19
        _ = v20
    }
    fmt.Println(v3, v3)
    for fi21 := 0; fi21 < 3; fi21++ {
        v7 = (^(v7))
        v22 := v6["hello"]
        _ = fi21
        _ = v22
    }
    delete(v6, "a")
    var v23 string = "a"
    v24 := GenS1{F0: 2147483647, F1: (2 <= 42)}
    var v25 int32 = 899
    if v2 {
        if true {
            v4 = v4
            switch 32767 {
            case 2147483647:
                fmt.Println(v3)
            case 820:
                if len(v5) > 0 {
                    v5[0] = 42
                }
            default:
                v1 = ("\n\t" <= "a")
            }
        }
        var v26 int8 = (v18 & v18)
        _ = v26
    }
    v27 := GenS2{F0: 32767, F1: -600, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
