package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 bool }
type GenS2 struct { F0 int8; F1 string; F2 uint64 }
type GenS3 struct { F0 int32; F1 int64; F2 int64 }
type GenS4 struct { F0 uint8; F1 int8; F2 string }
func genFn0(p0 int8) uint {
    p0 += (p0 - 127)
    fmt.Println(p0)
    switch 32768 {
    case 9223372036854775807:
        fmt.Println(p0)
    default:
        p0 = 0
    }
    var v6 int32 = -973
    v7 := GenS1{F0: 1, F1: 32767, F2: ((65535 != 1) || false)}
    _ = p0
    _ = v6
    _ = v7
    return 197
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 32767
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    var v5 uint16 = 65535
    for fi8 := 0; fi8 < 9; fi8++ {
        v1 = (^(v1))
        _ = fi8
    }
    var v9 bool = true
    fmt.Println(v9)
    var v10 uint64 = 0
    for fi11 := 0; fi11 < 19; fi11++ {
        v5 = 255
        _ = fi11
    }
    var v12 int64 = 9223372036854775807
    v10 = v10
    if (!(32767 < (v5 ^ v5))) {
        for fi13 := 0; fi13 < 4; fi13++ {
            v14 := []int{(fi13 + 1), fi13, (fi13 << ((uint(v5)) & 63)), fi13}
            _ = fi13
            _ = v14
        }
        v15 := GenS1{F0: (v2 + -32768), F1: 1, F2: (v9 || (v5 > v5))}
        var v16 uint16 = (v5 + v5)
        _ = v15
        _ = v16
    }
    if ((1.5 + 1.5) < 1.5) {
        if false {
            var v17 uint8 = 2
            v18 := GenS1{F0: v2, F1: 326, F2: false}
            _ = v17
            _ = v18
        }
    } else {
        var v19 uint64 = 506
        if v9 {
            v20 := GenS2{F0: 42, F1: "foo bar", F2: v10}
            v5 = 1
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 uint64 = v10
                if v9 {
                    var v23 int32 = int32(v5)
                    if (42 < 1) {
                        v19 -= v10
                        v2 -= int16(v23)
                    }
                    _ = v23
                } else {
                    v24 := []int{32768, (-(fi21)), fi21}
                    v25 := GenS1{F0: v2, F1: v2, F2: v9}
                    _ = len(v3)
                    _ = v24
                    _ = v25
                }
                _ = fi21
                _ = v22
            }
            _ = v20
        }
        _ = v19
    }
    var v26 uint16 = (v5 << ((0) & 63))
    if true {
        var v27 uint16 = (v26 / ((788) | 1))
        switch 32767 {
        case 127:
            for fi28 := 0; fi28 < 8; fi28++ {
                for fi29 := 0; fi29 < 20; fi29++ {
                    v30 := GenS1{F0: (v2 + 962), F1: v2, F2: v9}
                    _ = fi29
                    _ = v30
                }
                var v31 float32 = (((-2.25 - (1.5 - float32(float64(1.5)))) + float32(((1.5 - 1.5) * (3.14159 - 1e10)))) - 1e10)
                _ = fi28
                _ = v31
            }
        }
        _ = v27
    } else {
        for fi32 := 0; fi32 < 15; fi32++ {
            fmt.Println(v5)
            var v33 uint64 = v10
            _ = fi32
            _ = v33
        }
        v34 := v4[""]
        _ = v34
    }
    v2 += v2
    for fi35 := 0; fi35 < 16; fi35++ {
        v36 := []int{}
        _ = fi35
        _ = v36
    }
    var v37 int = -2147483648
    var v38 uint16 = (v26 * (v5 | (v5 & v26)))
    var v39 int32 = 1
    v40 := GenS3{F0: (^(v39)), F1: (v12 / ((v1) | 1)), F2: (^(v1))}
    if v9 {
        v41 := GenS4{F0: 255, F1: -127, F2: "\n\t"}
        _ = v41
    } else {
        v39 -= 1
        var v42 uint16 = v5
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v9
}
