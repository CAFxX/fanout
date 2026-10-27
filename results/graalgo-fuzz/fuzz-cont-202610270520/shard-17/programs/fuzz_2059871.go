package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 bool }
func genFn0(p0 int32, p1 int64) uint64 {
    v4 := []int{}
    if false {
        fmt.Println(p0, p1)
    }
    if (!(p1 == p1)) {
        if true {
            fmt.Println(p1, p0)
        }
        p1 = (p1 >> ((255) & 63))
    }
    if true {
        p0 = -617
        p1 = 9223372036854775807
        var v5 int8 = 0
        _ = v5
    }
    p0 = (p0 - (p0 ^ p0))
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    var v1 int32 = 1
    v2 := []uint32{1, 332, 366, 2147483648}
    var v3 uint32 = 0
    v1 = (-(v1))
    switch 42 {
    case 32768:
        var v6 float64 = (1e10 / 1.0)
        _ = v6
    case 127:
        switch "hello" {
        case "foo bar":
            for fi7 := 0; fi7 < 2; fi7++ {
                var v8 int64 = 1
                v3 += 4294967295
                _ = fi7
                _ = v8
            }
        default:
            switch ("hello" + "\n\t") {
            case "a":
                v9 := GenS1{F0: (^(v3))}
                _ = v9
            default:
                v10 := GenS1{F0: (v3 << ((255) & 63))}
                _ = v10
            }
        }
    case 42:
        switch -403 {
        case -9223372036854775808:
            v1 = v1
        default:
            v1 += v1
        }
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS2{F0: fi11, F1: true}
        for fi13 := 0; fi13 < 3; fi13++ {
            fmt.Println(fi13)
            var v14 int64 = -9223372036854775808
            _ = fi13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    fmt.Println(v1, v3)
    var v15 bool = false
    var v16 uint32 = (v3 * (v3 & v3))
    var v17 uint16 = 0
    var v18 int64 = 1
    fmt.Println(v1)
    v18 += (v18 | (v18 + -9223372036854775808))
    for fi19 := 0; fi19 < 3; fi19++ {
        v2 = append(v2, (v3 & v3))
        _ = fi19
    }
    if false {
        if len(v2) > 0 {
            v2[0] = 2
        }
        var v20 uint64 = 2
        _ = v20
    } else {
        fmt.Println(v15)
        for fi21 := 0; fi21 < 8; fi21++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v22 uint64 = 0
            if false {
                switch "foo bar" {
                case "hello":
                    v3 -= v16
                case "a":
                    v1 += v1
                case "foo bar":
                    var v23 int16 = 430
                    _ = v23
                }
                for fi24 := 0; fi24 < 1; fi24++ {
                    v25 := GenS2{F0: (-(fi24)), F1: (888 < 0)}
                    v15 = v15
                    _ = fi24
                    _ = v25
                }
                switch (fi21 * fi21) {
                case -771:
                    v17 = (^(v17))
                case 0:
                    var v26 uint32 = (v3 / ((v3) | 1))
                    _ = v26
                }
            } else {
                var v27 int32 = (v1 % ((207) | 1))
                _ = v27
            }
            _ = fi21
            _ = v22
        }
        v3 = v3
    }
    v28 := GenS2{F0: 475, F1: (2 < 128)}
    v3 = v3
    for fi29 := 0; fi29 < 9; fi29++ {
        v30 := GenS1{F0: v16}
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
}
