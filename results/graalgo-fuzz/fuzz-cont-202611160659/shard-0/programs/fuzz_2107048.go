package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0(p0 string, p1 uint8) uint8 {
    if (1 == -32768) {
        if ((0.0 + 1e10) > (0.1 - 1.5)) {
            v7 := []int{-650, 276, 42}
            _ = v7
        } else {
            v8 := GenS1{F0: p1}
            switch 4294967295 {
            case 1:
                p0 = p0
            case 42:
                for fi9 := 0; fi9 < 10; fi9++ {
                    p1 = (^(p1))
                    _ = fi9
                }
            case 594:
                v10 := []int{32767}
                _ = v10
            }
            v11 := GenS1{F0: p1}
            _ = v11
            _ = v8
        }
        if false {
            if (877 <= 412) {
                var v12 uint32 = 1
                var v13 uint64 = 18446744073709551615
                v14 := GenS1{F0: p1}
                _ = v12
                _ = v13
                _ = v14
            }
            if false {
                var v15 float64 = (1.5 - 1e10)
                _ = v15
            }
        }
        fmt.Println(p0, p1)
    }
    v16 := GenS1{F0: (p1 << ((1) & 63))}
    for fi17 := 0; fi17 < 13; fi17++ {
        switch "a" {
        case "foo bar":
            v18 := GenS1{F0: p1}
            _ = v18
        case "":
            for fi19 := 0; fi19 < 7; fi19++ {
                p1 += 42
                var v20 uint64 = 1
                var v21 int16 = 42
                if ((true || false) || true) {
                    v20 = 18446744073709551615
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        case "\n\t":
            p0 = ("\n\t" + (("" + p0) + ("hello" + (p0 + p0))))
        }
        v22 := []int{fi17, 721, fi17}
        p1 = 1
        _ = fi17
        _ = v22
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v16
    return p1
}

func main() {
    var v1 int16 = 1000
    var v2 int8 = -94
    v3 := make(map[uint32]bool)
    v4 := []string{("\n\t" + ""), (("" + "") + ""), ((("\n\t" + ("" + "foo bar")) + "foo bar") + "foo bar")}
    v5 := []string{"\n\t"}
    var v6 uint16 = 950
    fmt.Println(v6, v6, v6)
    v1 = (v1 | 0)
    for fi23 := 0; fi23 < 20; fi23++ {
        fmt.Println(v1)
        _ = fi23
    }
    v24 := GenS2{F0: 2, F1: -1}
    v2 -= v2
    switch "foo bar" {
    case "foo bar":
        var v25 uint8 = 2
        _ = v25
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        v5 = append(v5, ((("\n\t" + (("a" + "hello") + "foo bar")) + "\n\t") + "a"))
        _ = fi26
    }
    switch -136 {
    case 32767:
        var v27 int8 = 100
        _ = v27
    default:
        v1 += v1
    }
    _ = len(v4)
    fmt.Println(v6, v2, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
