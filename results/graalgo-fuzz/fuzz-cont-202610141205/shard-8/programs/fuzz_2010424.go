package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint8 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint16) uint64 {
    var v5 int32 = 2147483647
    p0 += p0
    switch 42 {
    case -153:
        fmt.Println(v5)
    case 2147483647:
        if true {
            var v6 int = 9223372036854775807
            _ = v6
        } else {
            v7 := GenS1{F0: -9223372036854775808, F1: -785, F2: 127}
            fmt.Println(p0)
            var v8 uint8 = 64
            _ = v7
            _ = v8
        }
    default:
        switch 32767 {
        case 1:
            switch 307 {
            case -9223372036854775808:
                var v9 uint16 = 0
                _ = v9
            case 32767:
                var v10 bool = true
                _ = v10
            default:
                v11 := GenS1{F0: 417, F1: 1, F2: 129}
                _ = v11
            }
        }
    }
    v5 = (v5 + 790)
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 int16, p1 int32, p2 uint16) uint64 {
    p1 = p1
    switch 9223372036854775807 {
    case -386:
        v12 := GenS1{F0: 9223372036854775807, F1: (p0 | (p0 ^ p0)), F2: 255}
        _ = v12
    case 9223372036854775807:
        v13 := []int{12, -2147483648}
        _ = v13
    case 719:
        var v14 uint16 = p2
        _ = v14
    default:
        fmt.Println(p2, p2, p1)
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 float64 = 0.1
        fmt.Println(p1, p1, v16)
        var v17 bool = false
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS2{F0: 935}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return 42
}

func main() {
    var v1 uint = 2
    var v2 int8 = -105
    var v3 int16 = -1
    var v4 bool = (v2 < v2)
    v19 := GenS2{F0: 0}
    switch -2147483648 {
    case 2147483647:
        if v4 {
            v20 := GenS2{F0: -9223372036854775808}
            _ = v20
        } else {
            fmt.Println(v4, v2)
            var v21 int8 = (v2 * (-(v2)))
            switch "foo bar" {
            case "hello":
                v21 = (v2 << ((v1) & 63))
            default:
                fmt.Println(v2, v4)
            }
            _ = v21
        }
    }
    v22 := []int{}
    v23 := GenS2{F0: -9223372036854775808}
    if v4 {
        for fi24 := 0; fi24 < 7; fi24++ {
            fmt.Println(v1)
            for fi25 := 0; fi25 < 15; fi25++ {
                v26 := GenS2{F0: 591}
                fmt.Println(v4, fi25)
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
        fmt.Println(v3, v4, v1)
        fmt.Println(v4, v1)
    }
    switch -2147483648 {
    case 32768:
        v27 := GenS2{F0: 9223372036854775807}
        _ = v27
    case 414:
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 float64 = 1.5
            _ = fi28
            _ = v29
        }
    default:
        _ = len(v22)
    }
    var v30 float32 = 0.0
    if true {
        switch 42 {
        case 42:
            v30 += 0.0
        }
    } else {
        v31 := GenS2{F0: 42}
        _ = v31
    }
    v1 = (v1 << ((v1) & 63))
    for fi32 := 0; fi32 < 17; fi32++ {
        v3 += 792
        var v33 int = fi32
        v3 += (^(v3))
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v4
}
