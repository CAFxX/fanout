package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
func genFn0() uint8 {
    var v7 int8 = 1
    v7 = (v7 << ((18446744073709551615) & 63))
    _ = v7
    return 45
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint16 = 197
    var v3 uint = 1
    v4 := []string{"foo bar", "\n\t"}
    var v5 int64 = -1
    var v6 int = 271
    for fi8 := 0; fi8 < 7; fi8++ {
        v6 += v6
        for fi9 := 0; fi9 < 4; fi9++ {
            if false {
                if len(v4) > 0 {
                    v4[0] = ((((((((("foo bar" + "hello") + "") + "foo bar") + "foo bar") + "foo bar") + "a") + "\n\t") + "") + "")
                }
                var v10 uint8 = genFn0()
                v4 = append(v4, "hello")
                _ = v10
            }
            v6 = fi9
            fmt.Println(v3, v5)
            _ = fi9
        }
        var v11 int16 = -32768
        _ = fi8
        _ = v11
    }
    v2 = (v2 ^ v2)
    if true {
        fmt.Println(v1, v6, v2)
        v12 := GenS1{F0: int64(v6), F1: 42}
        _ = v12
    } else {
        var v13 uint32 = v1
        _ = v13
    }
    var v14 uint64 = 18446744073709551615
    v15 := GenS1{F0: v5, F1: -1}
    for fi16 := 0; fi16 < 2; fi16++ {
        v1 += v1
        _ = fi16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        if (0 > 155) {
            var v18 int = fi17
            switch 32767 {
            case 2147483647:
                switch fi17 {
                case 702:
                    for fi19 := 0; fi19 < 18; fi19++ {
                        var v20 int8 = -1
                        _ = fi19
                        _ = v20
                    }
                }
            default:
                v5 = v5
            }
            _ = len(v4)
            _ = v18
        } else {
            var v21 uint32 = v1
            fmt.Println(v2)
            _ = v21
        }
        v2 -= v2
        _ = fi17
    }
    if ((v14 / (((v14 & 4294967295)) | 1)) > (v14 + (v14 * (^(v14))))) {
        v3 = uint(v1)
    } else {
        v6 += (v6 / ((v6) | 1))
    }
    fmt.Println(v14, v2, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
