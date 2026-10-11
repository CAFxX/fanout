package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 uint32 }
func genFn0() uint64 {
    v7 := []int{32767, 987, 9223372036854775807, 127}
    var v8 uint8 = 51
    if false {
        var v9 uint8 = v8
        if (false && (true || true)) {
            v10 := GenS1{F0: 42, F1: 18446744073709551615, F2: ("\n\t" + "hello")}
            switch 1 {
            case -9223372036854775808:
                var v11 int64 = -709
                _ = v11
            case 32768:
                for fi12 := 0; fi12 < 12; fi12++ {
                    v13 := GenS1{F0: 302, F1: 265, F2: ""}
                    _ = fi12
                    _ = v13
                }
            }
            _ = v10
        } else {
            if len(v7) > 0 {
                v7[0] = -344
            }
        }
        if false {
            var v14 uint16 = 2
            fmt.Println(v14)
            _ = v14
        }
        _ = v9
    }
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 int8 = -1
    var v2 int32 = -2147483648
    var v3 int = -836
    v4 := []int{}
    v5 := make(map[string]uint32)
    v6 := make(map[string]string)
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 int8 = 100
        _ = fi15
        _ = v16
    }
    var v17 int64 = 572
    _ = len(v6)
    var v18 int32 = (-(v2))
    v18 += v2
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 int64 = 0
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        switch "hello" {
        case "a":
            v4 = append(v4, fi21)
        default:
            var v22 float32 = (float32(1e10) / 1.0)
            _ = v22
        }
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 float64 = (1e10 + ((1.5 + 0.1) * (-2.25 + 1e10)))
            if ((!false) && true) {
                switch "\n\t" {
                case "foo bar":
                    if false {
                        if len(v4) > 0 {
                            v4[0] = (v3 << ((834) & 63))
                        }
                    } else {
                        v18 += (^(v2))
                        v25 := GenS2{F0: 42}
                        _ = v25
                    }
                }
            }
            _ = fi23
            _ = v24
        }
        _ = fi21
    }
    fmt.Println(v1)
    v18 += (v2 + -734)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
