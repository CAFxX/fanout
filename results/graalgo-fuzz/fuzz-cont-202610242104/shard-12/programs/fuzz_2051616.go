package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint64 }
type GenS2 struct { F0 int32; F1 int32 }
func genFn0(p0 int16, p1 uint16, p2 int) uint8 {
    switch "a" {
    case "a":
        p0 = (p0 + 1000)
    case "foo bar":
        var v6 uint = 18446744073709551615
        _ = v6
    case "hello":
        v7 := []int{(p2 << ((0) & 63)), (p2 ^ int(p0))}
        _ = v7
    default:
        switch "\n\t" {
        case "\n\t":
            var v8 float64 = 0.1
            _ = v8
        case "":
            v9 := []int{(-(p2)), p2, (^(p2)), (p2 | 496)}
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        if (1 <= -1) {
            fmt.Println(p1, p1, p0)
        }
        p0 -= 1000
        _ = fi10
    }
    if false {
        var v11 uint32 = 1
        v12 := GenS1{F0: uint(p1), F1: 2, F2: 730}
        v13 := []int{}
        _ = v11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return 157
}

func main() {
    v1 := []int{128, -9223372036854775808}
    v2 := []int{42, 128, 999, 81}
    v3 := []uint32{332, 409, 1}
    var v4 int64 = 0
    var v5 float32 = 1.5
    v5 = -2.25
    switch 1 {
    case -132:
        var v14 int64 = 9223372036854775807
        _ = v14
    case -79:
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 string = ("hello" + "\n\t")
            v4 -= v4
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 int8 = -1
                _ = fi17
                _ = v18
            }
            _ = fi15
            _ = v16
        }
    case 32767:
        v4 += (v4 | (-(v4)))
    default:
        var v19 int64 = -9223372036854775808
        _ = v19
    }
    if true {
        var v20 int = -1
        if len(v2) > 0 {
            v2[0] = v20
        }
        _ = v20
    }
    var v21 bool = false
    var v22 uint32 = 65535
    if v21 {
        fmt.Println(v5)
        fmt.Println(v4, v4)
    }
    if (-473 >= -2147483648) {
        switch 2147483647 {
        case -1:
            if v21 {
                v22 -= v22
                var v23 int64 = -173
                for fi24 := 0; fi24 < 10; fi24++ {
                    fmt.Println(v4)
                    _ = fi24
                }
                _ = v23
            } else {
                for fi25 := 0; fi25 < 8; fi25++ {
                    v21 = (2 < 1)
                    var v26 float64 = 1e10
                    _ = len(v1)
                    _ = fi25
                    _ = v26
                }
                v22 = (v22 >> ((0) & 63))
                fmt.Println(v5)
            }
        case -2147483648:
            for fi27 := 0; fi27 < 8; fi27++ {
                v28 := GenS2{F0: int32(fi27), F1: -1}
                _ = fi27
                _ = v28
            }
        case 9223372036854775807:
            if true {
                fmt.Println(v4, v4, v4)
                if len(v2) > 0 {
                    v2[0] = 9223372036854775807
                }
            }
        }
        fmt.Println(v21)
        if v21 {
            if (!(!true)) {
                var v29 float64 = 1.5
                if len(v3) > 0 {
                    v3[0] = v22
                }
                for fi30 := 0; fi30 < 20; fi30++ {
                    if false {
                        v31 := GenS2{F0: 0, F1: 1}
                        v32 := GenS1{F0: 255, F1: 18446744073709551615, F2: 2}
                        _ = v31
                        _ = v32
                    }
                    v21 = v21
                    var v33 int = (fi30 * (fi30 / (((-(fi30))) | 1)))
                    _ = fi30
                    _ = v33
                }
                _ = v29
            } else {
                v4 += 178
                var v34 int64 = v4
                _ = v34
            }
            var v35 bool = (!(42 == 128))
            v4 += -1
            _ = v35
        }
    }
    switch 4294967295 {
    case -2147483648:
        var v36 uint = 18446744073709551615
        _ = v36
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v22 -= (v22 - (^(v22)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
