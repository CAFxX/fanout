package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int8; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int64 }
func genFn0() string {
    v4 := GenS1{F0: 1, F1: -9223372036854775808}
    switch "hello" {
    case "hello":
        var v5 int32 = 1
        _ = v5
    case "":
        v6 := GenS1{F0: 65535, F1: 9223372036854775807}
        _ = v6
    case "a":
        var v7 uint = 18446744073709551615
        _ = v7
    default:
        var v8 uint64 = 2
        _ = v8
    }
    switch "a" {
    case "foo bar":
        switch 128 {
        case -189:
            v9 := []int{}
            _ = v9
        default:
            var v10 int64 = -570
            _ = v10
        }
    }
    _ = v4
    return "hello"
}
func genFn1(p0 int8, p1 uint32) uint64 {
    p0 = 85
    var v11 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = v11
    return v11
}

func main() {
    var v1 uint8 = 178
    var v2 int8 = -72
    var v3 int = 128
    v12 := []int{(v3 / ((91) | 1)), (v3 << ((2) & 63)), v3, (v3 << ((503) & 63))}
    if ((!true) && (!(false || true))) {
        v12 = append(v12, -185)
    } else {
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := GenS1{F0: 825, F1: -1}
            v15 := GenS2{F0: (^(v2)), F1: (v1 / ((v1) | 1)), F2: 32767}
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(v1, v3)
        if false {
            for fi16 := 0; fi16 < 19; fi16++ {
                _ = len(v12)
                var v17 int8 = (v2 * (v2 * 42))
                _ = fi16
                _ = v17
            }
        } else {
            v3 -= (v3 << ((42) & 63))
        }
    }
    switch "a" {
    case "":
        var v18 bool = (false || (-2.25 < 0.0))
        _ = v18
    case "\n\t":
        for fi19 := 0; fi19 < 4; fi19++ {
            fmt.Println(v2)
            var v20 float64 = 0.1
            _ = fi19
            _ = v20
        }
    default:
        v1 += (v1 << ((131) & 63))
    }
    var v21 float64 = 3.14159
    fmt.Println(v1, v3, v1)
    if false {
        v22 := GenS3{F0: 572}
        v21 = float64(3.14159)
        _ = v22
    } else {
        fmt.Println(v1, v21)
    }
    switch v3 {
    case 129:
        v23 := GenS1{F0: 2147483647, F1: 852}
        _ = v23
    }
    if true {
        var v24 int8 = 41
        for fi25 := 0; fi25 < 1; fi25++ {
            switch 127 {
            case -2147483648:
                if false {
                    if len(v12) > 0 {
                        v12[0] = (fi25 * v3)
                    }
                    if ((v21 * v21) >= v21) {
                        if (2147483648 == 42) {
                            for fi26 := 0; fi26 < 19; fi26++ {
                                _ = len(v12)
                                var v27 int = 4294967295
                                _ = fi26
                                _ = v27
                            }
                        }
                        var v28 int64 = -9223372036854775808
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        _ = v28
                    }
                }
            case -99:
                if len(v12) > 0 {
                    v12[0] = (fi25 - (-(v3)))
                }
            }
            v2 = v2
            if len(v12) > 0 {
                v12[0] = (fi25 & (fi25 + v3))
            }
            _ = fi25
        }
        if true {
            fmt.Println(v1, v21)
            if len(v12) > 0 {
                v12[0] = (v3 - -1)
            }
            v29 := GenS4{F0: 156, F1: -9223372036854775808}
            _ = v29
        }
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
}
