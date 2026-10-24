package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint64, p1 uint8, p2 uint8) uint8 {
    p1 = (p1 * 0)
    if (1.5 < 1e10) {
        v4 := []int{-1}
        fmt.Println(p0)
        _ = v4
    } else {
        p2 += p2
        v5 := []int{9223372036854775807}
        var v6 uint8 = 1
        _ = v5
        _ = v6
    }
    v7 := []int{2147483647, 2147483647, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}
func genFn1() uint32 {
    var v8 int8 = -32
    v9 := []int{-9223372036854775808, 1, -2147483648}
    v10 := GenS1{F0: 412, F1: 1}
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS2{F0: 128, F1: (!(!true))}
        v13 := GenS3{F0: ("foo bar" + "\n\t"), F1: 0}
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = v10
    _ = v8
    _ = v9
    return 381
}

func main() {
    var v1 uint64 = 376
    v2 := make(map[string]int)
    var v3 int64 = -1
    switch "" {
    case "":
        v2[(("" + "a") + ("" + (("" + ("foo bar" + "")) + "hello")))] = 32767
    default:
        var v14 int8 = -78
        _ = v14
    }
    switch (("a" + "hello") + "a") {
    case "hello":
        switch 1 {
        case -296:
            fmt.Println(v3, v3)
        case -266:
            var v15 float32 = (-2.25 * -2.25)
            _ = v15
        case 2147483647:
            v1 = v1
        default:
            _ = len(v2)
        }
    case "foo bar":
        var v16 bool = true
        _ = v16
    case "\n\t":
        if (1000 <= int16(v1)) {
            var v17 uint32 = 42
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 string = ("" + ("foo bar" + "a"))
                var v20 int16 = -32768
                v21 := GenS2{F0: 42, F1: false}
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            var v22 float64 = ((3.14159 / 1.0) - 0.0)
            _ = v17
            _ = v22
        } else {
            v23 := GenS4{F0: (-(v3)), F1: 42, F2: v1}
            if true {
                v3 -= v3
                switch 281 {
                case 32768:
                    v24 := []int{9223372036854775807, int(v1), 32767}
                    _ = v24
                }
            }
            _ = v23
        }
    }
    switch 32767 {
    case 127:
        v25 := []int{0, 128, 0}
        _ = v25
    case 920:
        var v26 bool = true
        _ = v26
    default:
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 int64 = (-(v3))
            var v29 int64 = (v28 * v28)
            v30 := GenS5{F0: 0, F1: 42}
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
    }
    var v31 int32 = 42
    v1 = 2
    switch -86 {
    case 128:
        for fi32 := 0; fi32 < 4; fi32++ {
            var v33 string = ("a" + ("a" + "foo bar"))
            _ = fi32
            _ = v33
        }
    case -2147483648:
        v34 := GenS4{F0: 9223372036854775807, F1: 916, F2: (v1 % ((2) | 1))}
        _ = v34
    default:
        v3 += (v3 << ((0) & 63))
    }
    v1 += (^(v1))
    switch -605 {
    case -2147483648:
        _ = len(v2)
    case 0:
        var v35 uint = 65535
        _ = v35
    case 627:
        var v36 int = -9223372036854775808
        _ = v36
    }
    switch "\n\t" {
    case "a":
        if (!true) {
            v31 = (v31 << ((42) & 63))
            var v37 uint = 18446744073709551615
            _ = v37
        } else {
            if false {
                var v38 int32 = v31
                _ = v38
            }
            var v39 uint16 = 545
            _ = v39
        }
    case "":
        var v40 uint16 = 2
        _ = v40
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        v42 := v2[("" + "a")]
        _ = fi41
        _ = v42
    }
    v43 := []int{4294967295, -365, 1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v43
}
