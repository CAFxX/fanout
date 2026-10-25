package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int32; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 int16, p1 int16) int8 {
    fmt.Println(p0, p0)
    if (42 <= 550) {
        var v6 uint = 0
        switch 0 {
        case -2147483648:
            v7 := []int{-9223372036854775808, 32768, 9223372036854775807, 42}
            _ = v7
        case 4294967295:
            v8 := GenS1{F0: 713, F1: -1}
            _ = v8
        default:
            switch 557 {
            case 9223372036854775807:
                for fi9 := 0; fi9 < 3; fi9++ {
                    for fi10 := 0; fi10 < 19; fi10++ {
                        v11 := []int{-9223372036854775808, 128, fi10, int(p1)}
                        p1 = p1
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(p0, fi10, fi10)
                        case "hello":
                            v6 -= v6
                        }
                        _ = fi10
                        _ = v11
                    }
                    v12 := []int{fi9, int(v6), 935, (fi9 >> ((v6) & 63))}
                    _ = fi9
                    _ = v12
                }
            }
        }
        _ = v6
    }
    _ = p0
    _ = p1
    return 100
}
func genFn1() int64 {
    var v13 string = ""
    if (!((!(float64(1e10) > 0.0)) || true)) {
        var v14 int32 = 596
        v14 = (^(v14))
        v13 = ((v13 + v13) + (("foo bar" + "\n\t") + "\n\t"))
        _ = v14
    } else {
        if (!(127 <= 481)) {
            v15 := GenS1{F0: 128, F1: 15}
            _ = v15
        } else {
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := GenS1{F0: (fi16 * fi16), F1: -128}
                v18 := GenS1{F0: 231, F1: 127}
                v13 = (v13 + "hello")
                _ = fi16
                _ = v17
                _ = v18
            }
        }
    }
    v13 = v13
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 1; fi20++ {
            v13 = (("a" + v13) + (v13 + (v13 + v13)))
            v13 = v13
            var v21 int = 32768
            _ = fi20
            _ = v21
        }
        v22 := GenS1{F0: -1, F1: genFn0(0, 42)}
        v23 := GenS2{F0: 2147483647, F1: -1, F2: 574}
        _ = fi19
        _ = v22
        _ = v23
    }
    _ = v13
    return 42
}

func main() {
    var v1 int64 = 42
    var v2 float64 = float64((-2.25 * 0.0))
    var v3 int16 = -32768
    v4 := make(map[string]uint32)
    v5 := []uint32{65535, 1, 1}
    if false {
        fmt.Println(v2)
        var v24 uint = 18446744073709551615
        _ = v24
    }
    v1 = genFn1()
    if false {
        v1 -= (v1 * v1)
        for fi25 := 0; fi25 < 16; fi25++ {
            var v26 int = (^(fi25))
            v26 += 42
            v27 := GenS1{F0: (v26 + fi25), F1: 1}
            _ = fi25
            _ = v26
            _ = v27
        }
        v3 -= (v3 << ((472) & 63))
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        for fi29 := 0; fi29 < 2; fi29++ {
            var v30 uint16 = 2
            v30 = v30
            fmt.Println(v1, fi28)
            _ = fi29
            _ = v30
        }
        v31 := GenS3{F0: 65535, F1: (65535 >= 2)}
        _ = fi28
        _ = v31
    }
    fmt.Println(v3)
    v32 := GenS2{F0: 1, F1: -128, F2: 2}
    v1 -= v1
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
