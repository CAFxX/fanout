package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint32 }
func genFn0(p0 int64, p1 uint8, p2 bool) uint64 {
    p2 = (p0 != p0)
    var v4 uint64 = 2
    v5 := []int{128, 1, 2147483647, 128}
    p1 += (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 1
    var v2 uint64 = 1
    v3 := make(map[string]int)
    fmt.Println(v2, v1)
    v6 := GenS1{F0: 217}
    switch ("foo bar" + (("" + ("foo bar" + ("\n\t" + "a"))) + "\n\t")) {
    case "hello":
        v7 := []int{127}
        _ = v7
    case "":
        v8 := []int{}
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        if (!false) {
            v2 -= v2
        } else {
            fmt.Println(fi9)
            v10 := GenS1{F0: 0}
            _ = v10
        }
        switch fi9 {
        case 861:
            v11 := []int{fi9}
            _ = v11
        case 127:
            if false {
                var v12 float64 = 0.0
                var v13 uint8 = uint8(v1)
                _ = v12
                _ = v13
            }
        case 887:
            v14 := []int{(fi9 - (fi9 / ((fi9) | 1)))}
            _ = v14
        default:
            switch "a" {
            case "a":
                var v15 int = (fi9 * fi9)
                _ = v15
            default:
                if true {
                    v16 := GenS1{F0: 42}
                    v2 = uint64(v1)
                    _ = v16
                }
            }
        }
        fmt.Println(v2)
        _ = fi9
    }
    var v17 uint32 = 1
    if false {
        v18 := GenS1{F0: uint8(v17)}
        var v19 int = 32767
        _ = v18
        _ = v19
    }
    v2 += (v2 ^ (v2 % (((v2 << ((0) & 63))) | 1)))
    switch (("hello" + ("\n\t" + "\n\t")) + "\n\t") {
    case "foo bar":
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := []int{-9223372036854775808, (fi20 ^ fi20)}
            var v22 uint16 = (v1 << ((0) & 63))
            _ = fi20
            _ = v21
            _ = v22
        }
    case "\n\t":
        v23 := GenS2{F0: true, F1: 255, F2: (v17 / ((v17) | 1))}
        _ = v23
    default:
        var v24 int16 = 0
        _ = v24
    }
    var v25 float64 = (1e10 + (-2.25 + (0.0 / 0.1)))
    if (true || ((1 > 2147483647) && (127 != 128))) {
        v17 = v17
        v1 += 0
        fmt.Println(v25)
    }
    if false {
        v25 -= (v25 + (((v25 / 1.5) - -2.25) + float64(1.5)))
    }
    if (42 < -80) {
        v25 += v25
        delete(v3, "foo bar")
        v26 := GenS2{F0: (!((!(147 == 255)) || true)), F1: 128, F2: 2}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v6
}
