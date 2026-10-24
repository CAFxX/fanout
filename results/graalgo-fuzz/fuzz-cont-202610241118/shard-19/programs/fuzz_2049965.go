package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0(p0 uint8, p1 uint16, p2 uint) uint {
    var v4 uint64 = 4294967295
    if true {
        var v5 uint64 = (v4 - v4)
        _ = v5
    }
    v6 := GenS1{F0: -955, F1: 2147483647}
    if true {
        p0 += (p0 * (^(p0)))
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return p2
}

func main() {
    var v1 string = ""
    v2 := []int{-9223372036854775808}
    var v3 int16 = -32768
    for fi7 := 0; fi7 < 14; fi7++ {
        v3 = (v3 + (v3 & v3))
        if (2147483647 >= 0) {
            var v8 int32 = 2147483647
            _ = v8
        } else {
            switch (fi7 ^ 4294967295) {
            case -9223372036854775808:
                if (42 < 2147483648) {
                    v1 = ""
                    var v9 int = fi7
                    v1 = v1
                    _ = v9
                }
            default:
                v3 = (v3 << ((0) & 63))
            }
            _ = len(v2)
        }
        switch v1 {
        case "hello":
            if len(v2) > 0 {
                v2[0] = fi7
            }
        default:
            v3 -= v3
        }
        _ = fi7
    }
    var v10 int64 = 1
    if ((v3 / ((v3) | 1)) < (-(v3))) {
        fmt.Println(v1)
    } else {
        if len(v2) > 0 {
            v2[0] = -1
        }
        var v11 uint32 = 611
        _ = v11
    }
    fmt.Println(v3, v1)
    switch ("a" + v1) {
    case "a":
        for fi12 := 0; fi12 < 16; fi12++ {
            v13 := GenS2{F0: v1, F1: 65535}
            _ = fi12
            _ = v13
        }
    }
    if true {
        for fi14 := 0; fi14 < 6; fi14++ {
            if false {
                var v15 uint16 = 32767
                _ = v15
            }
            _ = fi14
        }
    } else {
        v3 += -32768
    }
    v1 = v1
    if true {
        v1 = ("\n\t" + "\n\t")
        var v16 int16 = 1000
        if false {
            var v17 int = 127
            var v18 uint64 = 42
            _ = v17
            _ = v18
        }
        _ = v16
    } else {
        var v19 uint32 = 2
        switch -90 {
        case 976:
            var v20 uint8 = 2
            _ = v20
        case 127:
            v21 := GenS1{F0: -139, F1: 32767}
            _ = v21
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            for fi23 := 0; fi23 < 12; fi23++ {
                v2 = append(v2, (fi22 % (((-(fi22))) | 1)))
                _ = fi23
            }
            _ = fi22
        }
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
