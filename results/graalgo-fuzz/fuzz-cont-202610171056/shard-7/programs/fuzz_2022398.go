package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }

func main() {
    var v1 uint64 = 672
    var v2 uint64 = 478
    var v3 bool = (481 < -454)
    if true {
        if ((v3 && (!(v3 || (false && (!(true || true)))))) && (v3 || (1 > 42))) {
            v4 := GenS1{F0: (!false)}
            _ = v4
        }
        var v5 int = 4294967295
        if (42 <= 0) {
            v2 = (^(v2))
            v6 := GenS2{F0: 2}
            v3 = (!(!(-2.25 >= (0.1 - float32(1e10)))))
            _ = v6
        }
        _ = v5
    }
    fmt.Println(v3, v1)
    v2 = v2
    if false {
        var v7 int8 = 0
        _ = v7
    }
    if (-2147483648 <= 798) {
        fmt.Println(v1)
        v3 = (!(65535 <= 589))
        if true {
            var v8 int = int(v1)
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        if (24 >= 32767) {
            var v10 int32 = 1
            var v11 uint16 = 65535
            var v12 int16 = 0
            _ = v10
            _ = v11
            _ = v12
        } else {
            v13 := []int{(fi9 << ((4294967295) & 63)), (fi9 - (fi9 * fi9)), fi9}
            var v14 int64 = -1
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 int = (fi9 | (fi15 & (fi9 << ((255) & 63))))
                v17 := GenS2{F0: 2}
                v1 -= (^(v1))
                var v18 int16 = 1
                var v19 uint = 18446744073709551615
                v20 := GenS3{F0: 427}
                var v21 float32 = ((0.0 + (0.1 - float32(1e10))) * ((3.14159 - (1.5 * (0.0 * (float32(1e10) * (float32(float64(-2.25)) + 1.5))))) / 1.5))
                var v22 int = int(v14)
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
                _ = v19
                _ = v20
                _ = v21
                _ = v22
            }
            _ = v13
            _ = v14
        }
        if v3 {
            v23 := GenS1{F0: (!v3)}
            _ = v23
        }
        v24 := []int{9223372036854775807, (fi9 - (-(fi9)))}
        _ = fi9
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        switch (fi25 | fi25) {
        case 0:
            var v26 int16 = 1
            _ = v26
        default:
            v27 := []int{(fi25 >> ((65535) & 63))}
            _ = v27
        }
        v28 := GenS1{F0: v3}
        _ = fi25
        _ = v28
    }
    v3 = v3
    var v29 int32 = 580
    var v30 int32 = v29
    v31 := []int{127}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
