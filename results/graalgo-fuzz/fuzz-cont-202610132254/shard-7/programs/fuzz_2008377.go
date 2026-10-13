package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int8 }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0(p0 int32, p1 int, p2 int64) int8 {
    v6 := []int{}
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = len(v6)
    }
    if true {
        var v7 uint16 = 65535
        var v8 int8 = 127
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}
func genFn1() int16 {
    var v9 bool = false
    v9 = false
    v9 = v9
    for fi10 := 0; fi10 < 16; fi10++ {
        switch (fi10 % ((fi10) | 1)) {
        case 4294967295:
            var v11 int = (fi10 + 32767)
            _ = v11
        default:
            var v12 int = fi10
            _ = v12
        }
        v13 := []int{(fi10 * fi10), fi10, fi10, fi10}
        v14 := GenS1{F0: 591, F1: "", F2: 42}
        _ = fi10
        _ = v13
        _ = v14
    }
    _ = v9
    return 42
}

func main() {
    var v1 uint64 = 927
    var v2 int32 = 1
    v3 := []int{-9223372036854775808, 2147483647}
    var v4 uint16 = 42
    var v5 uint = 1
    for fi15 := 0; fi15 < 15; fi15++ {
        v4 = (v4 + (v4 >> ((v5) & 63)))
        v16 := GenS1{F0: v1, F1: ("foo bar" + "\n\t"), F2: genFn0(v2, fi15, -423)}
        _ = fi15
        _ = v16
    }
    var v17 uint = 4294967295
    if (32767 <= genFn1()) {
        var v18 int64 = 150
        v19 := GenS1{F0: (^(v1)), F1: "foo bar", F2: 116}
        _ = v18
        _ = v19
    } else {
        var v20 uint32 = 0
        _ = v20
    }
    if len(v3) > 0 {
        v3[0] = 32767
    }
    if true {
        var v21 uint = v17
        _ = v21
    } else {
        v22 := GenS2{F0: 0, F1: (v2 % ((v2) | 1))}
        _ = v22
    }
    v2 = v2
    v5 = (v5 >> ((v5) & 63))
    if false {
        if len(v3) > 0 {
            v3[0] = int(v4)
        }
        var v23 int = 4294967295
        if len(v3) > 0 {
            v3[0] = (-(v23))
        }
        _ = v23
    }
    v24 := GenS2{F0: 2147483647, F1: (v2 / ((-330) | 1))}
    v4 += 2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
