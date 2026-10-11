package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0(p0 uint64, p1 uint16) uint64 {
    for fi7 := 0; fi7 < 14; fi7++ {
        if (2147483648 >= 42) {
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := GenS1{F0: -673, F1: 926}
                var v10 int32 = 1
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 bool = true
            fmt.Println(v11, p1, fi7)
            _ = v11
        }
        p0 = p0
        _ = fi7
    }
    v12 := GenS2{F0: p0}
    _ = p0
    _ = p1
    _ = v12
    return p0
}
func genFn1(p0 int) int64 {
    p0 = p0
    var v13 int8 = 127
    _ = p0
    _ = v13
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 131
    v2 := []string{"\n\t", "foo bar", "\n\t", ""}
    var v3 float32 = 1.5
    v4 := make(map[uint32]int)
    var v5 int32 = 361
    var v6 uint16 = (v1 % (((v1 << ((42) & 63))) | 1))
    v14 := GenS3{F0: 32, F1: 32767}
    var v15 uint64 = 940
    if (42 <= (-(v5))) {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 uint8 = uint8(v15)
            var v18 uint32 = uint32(v6)
            v6 = (v1 | v6)
            _ = fi16
            _ = v17
            _ = v18
        }
        fmt.Println(v5, v1)
    }
    if true {
        var v19 uint16 = (^(v1))
        _ = v19
    }
    v15 = v15
    v4[uint32(v15)] = 611
    for fi20 := 0; fi20 < 2; fi20++ {
        v6 = v6
        v3 -= (v3 + v3)
        _ = fi20
    }
    var v21 int8 = 1
    var v22 uint64 = v15
    for fi23 := 0; fi23 < 16; fi23++ {
        if true {
            fmt.Println(v5, v6)
        } else {
            if len(v2) > 0 {
                v2[0] = "a"
            }
            v22 += 401
        }
        var v24 uint = 2
        switch fi23 {
        case 2147483647:
            v24 += v24
        case 32768:
            if len(v2) > 0 {
                v2[0] = (("foo bar" + "hello") + "foo bar")
            }
        }
        _ = fi23
        _ = v24
    }
    v21 += v21
    var v25 float32 = float32(0.1)
    var v26 int = -164
    v1 += 65535
    fmt.Println(v22, v26)
    v27 := GenS2{F0: (v15 & (v22 >> ((471) & 63)))}
    v25 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
