package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int32; F1 int16; F2 int64 }
type GenS3 struct { F0 int; F1 uint32; F2 int16 }
func genFn0(p0 int8) int8 {
    switch 9223372036854775807 {
    case 42:
        v7 := GenS1{F0: "", F1: -1}
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        p0 -= (p0 << ((1) & 63))
        if (!(true && (0 < -843))) {
            p0 = p0
            var v9 bool = true
            _ = v9
        } else {
            fmt.Println(p0)
            v10 := GenS2{F0: 2147483647, F1: -775, F2: -9223372036854775808}
            _ = v10
        }
        for fi11 := 0; fi11 < 12; fi11++ {
            p0 -= (-(p0))
            var v12 bool = ((255 <= 65535) || (!("foo bar" > "\n\t")))
            _ = fi11
            _ = v12
        }
        _ = fi8
    }
    p0 -= (p0 + 100)
    _ = p0
    return -128
}

func main() {
    var v1 uint8 = 128
    var v2 int64 = -1
    v3 := make(map[uint32]bool)
    var v4 uint16 = 2
    var v5 uint64 = 42
    var v6 int64 = -1
    if (!(true || true)) {
        for fi13 := 0; fi13 < 18; fi13++ {
            switch 32767 {
            case -212:
                for fi14 := 0; fi14 < 16; fi14++ {
                    var v15 uint = 2
                    v2 += (v6 + (^(v2)))
                    _ = fi14
                    _ = v15
                }
            default:
                v16 := GenS3{F0: (fi13 / ((502) | 1)), F1: 2147483648, F2: 171}
                _ = v16
            }
            switch fi13 {
            case 2147483647:
                var v17 uint64 = v5
                _ = v17
            }
            var v18 int16 = -45
            _ = fi13
            _ = v18
        }
        v19 := GenS2{F0: 436, F1: 32767, F2: (v2 - -241)}
        _ = v19
    } else {
        v3[4294967295] = false
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := GenS3{F0: int(v5), F1: 2147483648, F2: 32767}
            v3[uint32(v4)] = true
            v6 = int64(fi20)
            _ = fi20
            _ = v21
        }
        v22 := []int{499, 347, 4294967295, 420}
        _ = v22
    }
    v23 := []int{}
    v1 -= v1
    if false {
        v6 -= (^(v6))
        if ((v4 > v4) || (!(-1 >= -76))) {
            v6 += -9223372036854775808
            var v24 uint32 = 989
            _ = v24
        }
    } else {
        _ = len(v3)
        v3[65535] = (!((100 == 0) || false))
        v2 += (v6 / (((^(v2))) | 1))
    }
    fmt.Println(v2, v5)
    var v25 int64 = (-(v2))
    if false {
        var v26 int64 = v25
        var v27 uint64 = (v5 + (v5 * uint64(v4)))
        _ = v26
        _ = v27
    }
    v28 := GenS2{F0: 0, F1: -968, F2: v6}
    if true {
        v6 += -9223372036854775808
        fmt.Println(v2)
        var v29 string = ("\n\t" + "foo bar")
        _ = v29
    } else {
        if (true && (-69 < 0)) {
            var v30 uint8 = v1
            _ = v30
        }
        fmt.Println(v5)
    }
    var v31 uint64 = 0
    v1 -= 3
    var v32 int64 = v6
    var v33 int32 = -179
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
