package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int64 }
func genFn0(p0 int16) int8 {
    switch "a" {
    case "hello":
        p0 = -971
    case "\n\t":
        v6 := GenS1{F0: "a", F1: -1}
        _ = v6
    default:
        p0 -= -52
    }
    p0 = p0
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint16 = 65
        v8 += v8
        for fi9 := 0; fi9 < 20; fi9++ {
            if false {
                v8 -= 65535
            } else {
                if (42 > 0) {
                    p0 -= 1
                    if (false || (true || false)) {
                        p0 -= (p0 * 1)
                    }
                    v10 := []int{}
                    _ = v10
                }
            }
            v11 := GenS1{F0: "a", F1: -1}
            v8 += v8
            v12 := GenS2{F0: 42}
            _ = fi9
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        if true {
            p0 += p0
        }
        v14 := []int{fi13, fi13, 182, (fi13 * fi13)}
        _ = fi13
        _ = v14
    }
    _ = p0
    return int8(p0)
}

func main() {
    var v1 float64 = -2.25
    var v2 uint = 18446744073709551615
    var v3 uint64 = 0
    var v4 int = 128
    var v5 uint = (v2 / ((75) | 1))
    var v15 int16 = -32768
    v4 = v4
    var v16 float64 = (v1 * v1)
    switch "foo bar" {
    case "a":
        var v17 bool = (((v15 >> (((v5 & v5)) & 63)) == v15) && true)
        _ = v17
    case "":
        var v18 int8 = -1
        _ = v18
    default:
        v19 := GenS3{F0: v15}
        _ = v19
    }
    v20 := []int{(v4 | (v4 & (v4 | 2147483647))), v4, (^(v4))}
    v21 := GenS1{F0: ("hello" + ""), F1: 871}
    var v22 uint64 = 42
    if false {
        _ = len(v20)
        v23 := GenS4{F0: 65535, F1: 196}
        _ = v23
    } else {
        var v24 uint32 = 2
        v25 := GenS5{F0: (^(v3)), F1: -1}
        if (65535 > 0) {
            if (true && false) {
                var v26 uint16 = uint16(v24)
                _ = v26
            }
            fmt.Println(v4)
            var v27 int = (v4 % ((v4) | 1))
            _ = v27
        } else {
            fmt.Println(v24, v15)
            var v28 string = ("a" + ("\n\t" + "a"))
            _ = v28
        }
        _ = v24
        _ = v25
    }
    v29 := GenS5{F0: 4294967295, F1: 9223372036854775807}
    v15 = v15
    v22 = (^(v3))
    var v30 uint = (v5 ^ (^(v5)))
    if len(v20) > 0 {
        v20[0] = (v4 << ((398) & 63))
    }
    if len(v20) > 0 {
        v20[0] = (^(v4))
    }
    switch -471 {
    case 321:
        for fi31 := 0; fi31 < 16; fi31++ {
            var v32 float32 = 1e10
            _ = fi31
            _ = v32
        }
    case 360:
        var v33 int64 = 1
        _ = v33
    default:
        var v34 bool = false
        _ = v34
    }
    v30 += (v30 - v2)
    if true {
        switch 0 {
        case -692:
            _ = len(v20)
        case 127:
            var v35 uint = v30
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
