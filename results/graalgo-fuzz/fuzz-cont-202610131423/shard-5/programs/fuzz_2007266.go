package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int32, p1 bool) uint32 {
    p1 = p1
    for fi7 := 0; fi7 < 6; fi7++ {
        p0 += p0
        _ = fi7
    }
    switch "\n\t" {
    case "hello":
        var v8 int16 = 587
        _ = v8
    case "\n\t":
        if (3.14159 > (0.1 * 0.0)) {
            p0 -= -997
        } else {
            var v9 int16 = 220
            v10 := GenS1{F0: (p0 * p0)}
            _ = v10
            _ = v9
        }
    }
    v11 := GenS1{F0: (p0 << ((251) & 63))}
    _ = p0
    _ = p1
    _ = v11
    return 2147483648
}

func main() {
    v1 := []string{""}
    v2 := []int{}
    v3 := make(map[string]uint32)
    var v4 int = 0
    v5 := []int{}
    v6 := make(map[int]uint32)
    v12 := v6[v4]
    for fi13 := 0; fi13 < 13; fi13++ {
        switch -9223372036854775808 {
        case 127:
            _ = len(v2)
        case -1:
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 float32 = ((float32((-2.25 + 3.14159)) * 0.0) + 0.0)
                _ = fi14
                _ = v15
            }
        default:
            v12 -= (v12 % ((v12) | 1))
        }
        v16 := v3[(("foo bar" + "foo bar") + "")]
        _ = fi13
        _ = v16
    }
    v12 = (v12 << ((18446744073709551615) & 63))
    var v17 int16 = 32767
    var v18 string = ("hello" + "\n\t")
    var v19 uint64 = 4294967295
    fmt.Println(v17)
    v12 = (v12 & 42)
    var v20 int8 = 100
    v4 -= int(v19)
    switch v18 {
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = (v4 * int(v20))
        }
    default:
        var v21 float32 = 1e10
        _ = v21
    }
    switch v4 {
    case 128:
        v17 = (^(v17))
    case -149:
        for fi22 := 0; fi22 < 10; fi22++ {
            v4 -= fi22
            for fi23 := 0; fi23 < 10; fi23++ {
                v17 = (v17 | v17)
                _ = fi23
            }
            var v24 bool = (-9223372036854775808 != 42)
            _ = fi22
            _ = v24
        }
    }
    v19 = 1
    for fi25 := 0; fi25 < 17; fi25++ {
        fmt.Println(v18, v4)
        fmt.Println(fi25)
        v12 += (v12 >> ((18446744073709551615) & 63))
        _ = fi25
    }
    v26 := GenS1{F0: -61}
    for fi27 := 0; fi27 < 2; fi27++ {
        if (fi27 <= -2147483648) {
            if len(v5) > 0 {
                v5[0] = fi27
            }
            if (708 <= -9223372036854775808) {
                v18 = v18
                var v28 uint16 = 1
                _ = v28
            }
        }
        if true {
            for fi29 := 0; fi29 < 11; fi29++ {
                fmt.Println(v19, v20)
                v20 = (v20 & (v20 + 0))
                _ = fi29
            }
            v30 := GenS1{F0: 0}
            _ = v30
        } else {
            _ = len(v6)
            var v31 int8 = (v20 & v20)
            _ = v31
        }
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
