package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint32; F2 int16 }
type GenS4 struct { F0 int64; F1 int64; F2 uint }
type GenS5 struct { F0 uint8; F1 int32 }

func main() {
    v1 := []uint32{42, 4294967295}
    var v2 int = 2147483647
    var v3 uint = 18446744073709551615
    v4 := []int{(v2 / ((-9223372036854775808) | 1)), 127, v2}
    v5 := GenS1{F0: 759}
    switch "hello" {
    case "\n\t":
        if (v3 > v3) {
            var v6 bool = ((1e10 - float64(0.0)) == -2.25)
            _ = v6
        }
    default:
        fmt.Println(v2, v2)
    }
    v2 = v2
    switch ("hello" + "") {
    case "foo bar":
        fmt.Println(v2, v2)
    case "a":
        for fi7 := 0; fi7 < 5; fi7++ {
            for fi8 := 0; fi8 < 10; fi8++ {
                var v9 uint8 = 255
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    case "":
        _ = len(v4)
    }
    v10 := GenS2{F0: ("foo bar" == ("" + "hello")), F1: 139}
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS2{F0: ("hello" != "a"), F1: 303}
        var v13 int32 = int32(v3)
        var v14 int32 = -2147483648
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v2 += (-(v2))
    var v15 bool = false
    if (-2.25 > 1e10) {
        for fi16 := 0; fi16 < 18; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 int64 = 1
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    } else {
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS1{F0: 0}
            var v21 uint8 = 238
            v3 -= 667
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v22 := GenS3{F0: v3, F1: 0, F2: -339}
    v23 := GenS3{F0: v3, F1: 0, F2: -32768}
    if v15 {
        switch -1 {
        case -1:
            switch "foo bar" {
            case "foo bar":
                var v24 int = (v2 + v2)
                _ = v24
            case "a":
                v3 = (v3 - v3)
            case "":
                var v25 int = v2
                _ = v25
            }
        default:
            _ = len(v1)
        }
    } else {
        var v26 float64 = (0.1 * 0.0)
        v26 += float64(((1e10 * 0.0) - -2.25))
        if ((42 >= 42) || (v2 != v2)) {
            for fi27 := 0; fi27 < 4; fi27++ {
                var v28 uint8 = 128
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi27
                _ = v28
            }
            fmt.Println(v15)
            v29 := GenS4{F0: 9223372036854775807, F1: -822, F2: v3}
            _ = v29
        } else {
            v3 += v3
            for fi30 := 0; fi30 < 16; fi30++ {
                var v31 int64 = 0
                v32 := GenS5{F0: 1, F1: -2147483648}
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        _ = v26
    }
    switch v2 {
    case 325:
        switch (v2 << ((v3) & 63)) {
        case 128:
            fmt.Println(v2, v2, v3)
        }
    default:
        v3 = (v3 << ((v3) & 63))
    }
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
