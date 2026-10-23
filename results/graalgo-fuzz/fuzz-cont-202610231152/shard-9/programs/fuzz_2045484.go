package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int16; F1 string }
func genFn0(p0 uint16) int8 {
    p0 = (p0 + (p0 / ((p0) | 1)))
    v7 := []int{}
    p0 -= p0
    _ = p0
    _ = v7
    return -6
}
func genFn1(p0 int64) int8 {
    p0 -= (p0 & p0)
    switch -860 {
    case 4294967295:
        v8 := []int{1, 128, 4294967295, 32768}
        _ = v8
    case 9223372036854775807:
        for fi9 := 0; fi9 < 1; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                p0 = (p0 << ((42) & 63))
                _ = fi10
            }
            var v11 int32 = -793
            var v12 int64 = p0
            _ = fi9
            _ = v11
            _ = v12
        }
    default:
        v13 := []int{}
        _ = v13
    }
    _ = p0
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 uint64 = 18446744073709551615
    var v4 uint32 = 65535
    var v5 uint64 = v3
    v6 := []uint32{v4}
    fmt.Println(v5, v5, v5)
    v14 := GenS1{F0: (!(80 != 42)), F1: 4294967295}
    v5 += v3
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            if (genFn0(255) > 40) {
                fmt.Println(v4, v3, fi15)
            } else {
                var v17 bool = true
                v18 := GenS1{F0: (42 >= 1), F1: 4294967295}
                _ = v17
                _ = v18
            }
            if false {
                delete(v2, (v4 * 0))
                v19 := GenS1{F0: true, F1: 0}
                _ = len(v6)
                v3 += (^(v5))
                v4 -= (^(v4))
                _ = v19
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = len(v6)
                _ = len(v6)
                v20 := v2[(v4 | v4)]
                if (42 != -40) {
                    v3 = v5
                }
                _ = v20
            }
            _ = fi16
        }
        v4 = 42
        _ = fi15
    }
    v3 -= v3
    v4 -= v4
    v3 -= (^(v3))
    switch -2147483648 {
    case -3:
        _ = len(v1)
    case 4294967295:
        for fi21 := 0; fi21 < 16; fi21++ {
            switch (("\n\t" + "\n\t") + ("" + "\n\t")) {
            case "\n\t":
                v4 -= v4
            case "hello":
                fmt.Println(v4)
            case "foo bar":
                v5 += v3
            }
            v5 = (v3 * v3)
            v3 += 4294967295
            _ = fi21
        }
    }
    var v22 uint = 451
    if true {
        var v23 uint = v22
        _ = v23
    } else {
        var v24 uint32 = 713
        _ = len(v6)
        _ = v24
    }
    v4 -= v4
    v25 := GenS2{F0: -1, F1: "a"}
    v26 := GenS2{F0: -32768, F1: ""}
    var v27 int16 = -17
    switch 0 {
    case 128:
        v27 = (v27 % ((42) | 1))
    }
    var v28 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
