package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint8 }
func genFn0(p0 int64, p1 uint64, p2 uint8) uint8 {
    var v7 uint8 = p2
    var v8 uint64 = (p1 / ((18446744073709551615) | 1))
    switch ("\n\t" + "hello") {
    case "":
        p1 = (v8 & v8)
    case "hello":
        switch -32 {
        case -1:
            v9 := []int{754, -564, 0}
            _ = v9
        case 9223372036854775807:
            p1 -= p1
        case -378:
            if true {
                v8 += uint64(p2)
                var v10 uint = 1
                v11 := []int{9223372036854775807}
                _ = v10
                _ = v11
            }
        }
    case "a":
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int16, p1 int8) uint {
    switch ("hello" + "hello") {
    case "hello":
        v12 := []int{-9223372036854775808, 12, -1}
        _ = v12
    case "foo bar":
        if false {
            v13 := []int{1, 539}
            _ = v13
        } else {
            v14 := GenS1{F0: 0}
            var v15 uint16 = 968
            fmt.Println(p0, v15, p0)
            _ = v14
            _ = v15
        }
    }
    p1 -= p1
    var v16 int16 = (p0 | (p0 | (p0 ^ (p0 * -1))))
    var v17 uint64 = 1
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 2
}

func main() {
    var v1 uint8 = 127
    v2 := []int{412, -855, 32768, -745}
    var v3 uint64 = 0
    var v4 int = -583
    v5 := make(map[uint32]uint32)
    v6 := []int{v4, 127}
    v18 := GenS1{F0: 0}
    v19 := v5[42]
    if len(v2) > 0 {
        v2[0] = (v4 >> ((18446744073709551615) & 63))
    }
    var v20 int32 = -445
    for fi21 := 0; fi21 < 10; fi21++ {
        switch "a" {
        case "hello":
            var v22 uint = 923
            _ = v22
        }
        _ = fi21
    }
    var v23 uint = 861
    for fi24 := 0; fi24 < 11; fi24++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v25 float64 = 0.1
        if (1000 == 42) {
            var v26 int = (v4 * fi24)
            v3 = (v3 | uint64(v19))
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v20)
    v19 += v19
    v3 += 18446744073709551615
    var v27 uint8 = 1
    delete(v5, v19)
    var v28 int64 = -646
    var v29 uint = 126
    for fi30 := 0; fi30 < 18; fi30++ {
        v4 = (fi30 << ((uint(v1)) & 63))
        v31 := GenS2{F0: int32(v3), F1: 2, F2: (v1 | v27)}
        v20 = v20
        _ = fi30
        _ = v31
    }
    _ = len(v2)
    if true {
        switch "a" {
        case "\n\t":
            var v32 uint64 = (^(v3))
            _ = v32
        case "a":
            v33 := GenS1{F0: v28}
            _ = v33
        default:
            v34 := GenS2{F0: 948, F1: 2, F2: (v1 ^ v27)}
            _ = v34
        }
        v35 := v5[uint32(v3)]
        for fi36 := 0; fi36 < 19; fi36++ {
            for fi37 := 0; fi37 < 17; fi37++ {
                if (640 != 42) {
                    var v38 uint8 = 154
                    v39 := GenS1{F0: int64(v23)}
                    var v40 float32 = (3.14159 + -2.25)
                    v1 = (v38 - v1)
                    var v41 int8 = 24
                    _ = v38
                    _ = v39
                    _ = v40
                    _ = v41
                } else {
                    var v42 int64 = 1
                    v35 += v35
                    _ = v42
                }
                switch "" {
                case "a":
                    v20 -= (^(v20))
                case "\n\t":
                    var v43 string = (("hello" + ((("foo bar" + "foo bar") + ((("foo bar" + ((("\n\t" + "foo bar") + "a") + ("a" + "hello"))) + ("hello" + "")) + "\n\t")) + "a")) + "")
                    _ = v43
                default:
                    var v44 uint16 = uint16(v1)
                    _ = v44
                }
                v23 -= (v23 + 959)
                _ = fi37
            }
            v6 = append(v6, v4)
            _ = fi36
        }
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
