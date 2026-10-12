package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 bool) uint64 {
    var v4 uint8 = 124
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 bool = (127 <= (v4 >> ((166) & 63)))
        v7 := []int{-9223372036854775808, (fi5 | fi5), (fi5 << ((42) & 63))}
        _ = fi5
        _ = v6
        _ = v7
    }
    v4 = v4
    if (164 == 128) {
        var v8 uint64 = uint64(v4)
        v9 := []int{28}
        switch "" {
        case "\n\t":
            switch (("\n\t" + ("foo bar" + ("a" + "\n\t"))) + "\n\t") {
            case "a":
                v10 := GenS1{F0: 1, F1: -771, F2: 42}
                _ = v10
            case "":
                if false {
                    v4 += (v4 / ((v4) | 1))
                }
            case "hello":
                v11 := GenS1{F0: 1, F1: 1, F2: -32768}
                _ = v11
            }
        default:
            var v12 float32 = float32(float64(3.14159))
            _ = v12
        }
        _ = v8
        _ = v9
    } else {
        if (!(!p0)) {
            for fi13 := 0; fi13 < 13; fi13++ {
                switch "" {
                case "a":
                    if p0 {
                        var v14 int = (fi13 | (-(fi13)))
                        var v15 int = -600
                        _ = v14
                        _ = v15
                    }
                case "\n\t":
                    v16 := GenS1{F0: 317, F1: -1, F2: 1}
                    _ = v16
                }
                var v17 int = (^(fi13))
                if p0 {
                    v4 -= 0
                }
                _ = fi13
                _ = v17
            }
            p0 = p0
        }
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 int16 = 32767
            for fi20 := 0; fi20 < 6; fi20++ {
                p0 = p0
                v4 = (v4 << ((0) & 63))
                _ = fi20
            }
            var v21 uint8 = v4
            _ = fi18
            _ = v19
            _ = v21
        }
    }
    _ = p0
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 float32 = (1.5 * (0.0 * 1.5))
    var v2 bool = true
    var v3 int16 = 1000
    v22 := []int{}
    if len(v22) > 0 {
        v22[0] = -739
    }
    fmt.Println(v2, v2, v1)
    if ((((float64(v1) >= (0.1 * (1.5 * -2.25))) && (!(42 > -1))) && true) || (!v2)) {
        fmt.Println(v1, v2)
    } else {
        fmt.Println(v1, v1, v2)
    }
    if ((1e10 / -2.25) >= 0.1) {
        var v23 float32 = v1
        var v24 int64 = 0
        v25 := GenS2{F0: v24, F1: -2147483648}
        _ = v23
        _ = v24
        _ = v25
    }
    v26 := GenS3{F0: 2}
    v1 = v1
    if v2 {
        v3 -= -1
    }
    v3 += (^(v3))
    v27 := GenS3{F0: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v27
    _ = v3
}
