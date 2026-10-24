package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 bool; F1 string }
func genFn0(p0 bool, p1 uint8, p2 uint64) bool {
    v6 := []int{1, 0, -575, -2147483648}
    switch 2147483647 {
    case 1:
        v6 = append(v6, 32768)
    case 128:
        var v7 bool = (!false)
        _ = v7
    case 0:
        fmt.Println(p1, p1)
    default:
        p1 = (p1 << ((255) & 63))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (true && (!p0))
}

func main() {
    var v1 int8 = -128
    v2 := make(map[int]bool)
    var v3 uint = 18446744073709551615
    v4 := []uint32{2}
    v5 := []string{(("" + "a") + ("foo bar" + ((((("a" + "hello") + "") + ("foo bar" + ("\n\t" + "\n\t"))) + "hello") + (("" + "") + "\n\t")))), ("hello" + ("foo bar" + (("\n\t" + "foo bar") + ("a" + ("" + "\n\t"))))), "foo bar", "\n\t"}
    fmt.Println(v1, v3)
    v3 = 18446744073709551615
    var v8 int32 = 838
    var v9 int32 = (v8 >> ((v3) & 63))
    v10 := GenS1{F0: -422, F1: uint64(v3)}
    v3 = v3
    v8 -= (v8 & v9)
    for fi11 := 0; fi11 < 14; fi11++ {
        v2[(^(fi11))] = false
        _ = fi11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(v3, fi12)
        _ = len(v5)
        if (495 <= fi12) {
            for fi13 := 0; fi13 < 10; fi13++ {
                v1 = (v1 ^ 1)
                _ = fi13
            }
            switch "" {
            case "foo bar":
                var v14 uint8 = 11
                _ = v14
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            switch "hello" {
            case "":
                v15 := GenS1{F0: 891, F1: 1}
                _ = v15
            case "a":
                var v16 float64 = float64(1.5)
                _ = v16
            }
        }
        _ = fi12
    }
    v17 := GenS2{F0: false, F1: "hello"}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
