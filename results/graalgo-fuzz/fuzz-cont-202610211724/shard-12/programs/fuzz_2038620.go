package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 bool; F1 uint; F2 int }
func genFn0(p0 int) uint {
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            p0 += p0
            var v9 float32 = -2.25
            var v10 float32 = ((v9 * float32(0.1)) * (-2.25 / v9))
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v11 := GenS1{F0: 974, F1: 42}
    v12 := GenS1{F0: 572, F1: 9223372036854775807}
    _ = p0
    _ = v11
    _ = v12
    return 2
}

func main() {
    v1 := []uint32{0}
    var v2 int8 = -1
    var v3 uint64 = 4294967295
    var v4 string = ""
    var v5 uint64 = v3
    v6 := []int{0, 45, 9223372036854775807, 42}
    v4 = v4
    var v13 bool = false
    v14 := GenS1{F0: 42, F1: 0}
    _ = len(v1)
    fmt.Println(v5, v13)
    fmt.Println(v2)
    v15 := GenS2{F0: 1, F1: 0}
    if ((117 < 1) || ((!(false && (69 > 1))) && (42 < 42))) {
        if (v13 && (v2 <= (v2 << ((4294967295) & 63)))) {
            v6 = append(v6, 4294967295)
            v13 = ((v3 << ((2) & 63)) == 0)
            _ = len(v6)
        }
    } else {
        if (v4 > (v4 + v4)) {
            v4 = v4
        } else {
            var v16 float32 = (1e10 + 0.1)
            if v13 {
                for fi17 := 0; fi17 < 17; fi17++ {
                    v18 := GenS3{F0: v13, F1: 1, F2: (fi17 - fi17)}
                    v13 = v13
                    v19 := GenS2{F0: 1, F1: -196}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v20 uint16 = 32767
                _ = v20
            }
            for fi21 := 0; fi21 < 1; fi21++ {
                var v22 bool = false
                var v23 uint16 = 65535
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v16
        }
        v5 = (^(v5))
        var v24 bool = v13
        _ = v24
    }
    if (v13 && (v13 || (!(4294967295 >= 9223372036854775807)))) {
        v5 = v5
        var v25 float32 = 0.0
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v25
    }
    v5 += 18446744073709551615
    v13 = ((1 <= 65535) || false)
    v4 = v4
    var v26 int16 = 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
