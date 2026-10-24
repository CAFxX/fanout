package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint16 = 311
    v3 := make(map[string]int)
    var v4 uint64 = 1
    var v5 int32 = -2147483648
    v6 := make(map[int]string)
    if (1.5 == 1e10) {
        for fi7 := 0; fi7 < 1; fi7++ {
            if false {
                var v8 string = ("foo bar" + "a")
                _ = v8
            }
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint32 = 52
                v10 = 4294967295
                v11 := GenS1{F0: 117}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi7
        }
        switch 128 {
        case -857:
            v4 += v4
        case -333:
            for fi12 := 0; fi12 < 15; fi12++ {
                fmt.Println(v2, v1, v5)
                if (false && ((v1 + (v1 / ((v1) | 1))) == 42)) {
                    v4 = 4294967295
                    v13 := GenS1{F0: 354}
                    _ = v13
                }
                _ = fi12
            }
        case 2147483647:
            v2 += v2
        default:
            v14 := v6[-1]
            _ = v14
        }
    }
    var v15 float64 = float64(float32(0.1))
    v16 := GenS1{F0: 927}
    v17 := GenS1{F0: -748}
    if (-32768 == 42) {
        v1 = 65535
    }
    var v18 uint8 = 42
    v19 := GenS1{F0: -9223372036854775808}
    for fi20 := 0; fi20 < 17; fi20++ {
        for fi21 := 0; fi21 < 10; fi21++ {
            var v22 uint = (v1 - v1)
            fmt.Println(v4, v2)
            _ = fi21
            _ = v22
        }
        fmt.Println(v4, v1)
        _ = fi20
    }
    switch 127 {
    case -9223372036854775808:
        for fi23 := 0; fi23 < 4; fi23++ {
            for fi24 := 0; fi24 < 12; fi24++ {
                fmt.Println(fi23)
                var v25 bool = ((v15 / (v15 - -2.25)) <= v15)
                fmt.Println(v5, v25, v25)
                v2 -= uint16(v1)
                if ((!v25) || false) {
                    v26 := []int{(fi23 + 0), (fi24 % (((fi23 ^ int(v18))) | 1))}
                    v27 := v3["foo bar"]
                    v28 := GenS1{F0: 1}
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
    case 2147483647:
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 uint8 = 0
            _ = fi29
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
