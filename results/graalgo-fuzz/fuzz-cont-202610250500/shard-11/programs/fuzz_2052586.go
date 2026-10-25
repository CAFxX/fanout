package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint8, p1 uint32) int16 {
    p1 = (p1 << ((4294967295) & 63))
    var v5 uint32 = p1
    v6 := []int{4294967295, 2147483647}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -32768
}
func genFn1(p0 int16, p1 uint8, p2 uint8) int {
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := []int{9223372036854775807, (fi8 / ((fi7) | 1)), fi8, (fi8 % ((fi8) | 1))}
            var v10 int = (fi7 + 128)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        switch (^(fi11)) {
        case -473:
            v12 := []int{}
            _ = v12
        case 32767:
            fmt.Println(p0)
        case 127:
            for fi13 := 0; fi13 < 4; fi13++ {
                v14 := []int{730, (fi11 ^ (^(fi13)))}
                if true {
                    var v15 int8 = 127
                    var v16 float32 = 0.0
                    _ = v15
                    _ = v16
                }
                fmt.Println(p1, fi13)
                _ = fi13
                _ = v14
            }
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    return -458
}

func main() {
    var v1 int64 = -198
    v2 := make(map[uint32]string)
    var v3 uint8 = 127
    var v4 uint32 = 406
    for fi17 := 0; fi17 < 8; fi17++ {
        v2[v4] = "\n\t"
        var v18 float32 = (1.5 * 0.1)
        v19 := GenS1{F0: (v3 * uint8(v4))}
        _ = fi17
        _ = v18
        _ = v19
    }
    v3 = v3
    v20 := GenS2{F0: 0}
    if true {
        if true {
            v1 -= v1
            var v21 int16 = -28
            v1 -= -714
            _ = v21
        }
        v22 := GenS1{F0: (v3 >> ((18446744073709551615) & 63))}
        var v23 int64 = v1
        _ = v22
        _ = v23
    }
    v4 = v4
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := GenS2{F0: 42}
        var v26 int64 = v1
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int8 = -128
    v28 := []int{-338, 127, 793}
    var v29 int16 = 967
    v30 := v2[(v4 | (v4 % (((v4 & 4294967295)) | 1)))]
    _ = len(v28)
    v27 += int8(v4)
    if true {
        fmt.Println(v29)
        var v31 int64 = (v1 - 0)
        fmt.Println(v30, v1, v27)
        _ = v31
    } else {
        var v32 uint8 = v3
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
