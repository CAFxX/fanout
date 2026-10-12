package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 uint; F1 int }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 49
    var v3 int8 = 34
    var v4 float32 = 1.5
    var v5 int = -491
    var v6 int16 = -922
    switch "" {
    case "\n\t":
        v3 -= v3
    default:
        var v7 bool = (false && (false || true))
        _ = v7
    }
    var v8 int64 = 0
    fmt.Println(v3, v6, v6)
    v6 += 1
    v4 = v4
    for fi9 := 0; fi9 < 13; fi9++ {
        if false {
            if (!true) {
                v8 += v8
            }
            var v10 float32 = v4
            for fi11 := 0; fi11 < 20; fi11++ {
                switch v5 {
                case -177:
                    var v12 uint32 = 0
                    _ = v12
                case -2147483648:
                    fmt.Println(v8, v8, fi9)
                case 32767:
                    v13 := GenS1{F0: 1, F1: (0 <= 4294967295), F2: 2}
                    _ = v13
                default:
                    v14 := GenS1{F0: -2147483648, F1: false, F2: 42}
                    _ = v14
                }
                _ = fi11
            }
            _ = v10
        } else {
            fmt.Println(v4)
            var v15 int64 = -9223372036854775808
            _ = v15
        }
        _ = fi9
    }
    if false {
        switch "hello" {
        case "":
            delete(v1, "\n\t")
        case "\n\t":
            if (false || true) {
                var v16 bool = (true || (!true))
                _ = v16
            } else {
                if (-2147483648 <= 1) {
                    var v17 float32 = v4
                    _ = v17
                }
            }
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS1{F0: 257, F1: ((v8 + (v8 - (-(v8)))) > (v8 << ((18446744073709551615) & 63))), F2: 450}
            v20 := GenS1{F0: 236, F1: ((fi18 * (^(fi18))) != 0), F2: 2}
            _ = fi18
            _ = v19
            _ = v20
        }
    } else {
        v21 := v1["a"]
        if false {
            v22 := []int{}
            _ = v22
        } else {
            var v23 int16 = 467
            var v24 int8 = v3
            v25 := GenS2{F0: 387, F1: ("\n\t" + "\n\t")}
            _ = v23
            _ = v24
            _ = v25
        }
        v26 := GenS3{F0: 0, F1: -9223372036854775808}
        _ = v21
        _ = v26
    }
    v3 = (-(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
