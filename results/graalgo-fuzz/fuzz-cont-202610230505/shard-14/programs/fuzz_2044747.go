package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int64 }
func genFn0(p0 uint64) uint64 {
    switch -963 {
    case -2147483648:
        if true {
            v6 := []int{0}
            var v7 uint = 955
            _ = v6
            _ = v7
        }
    case 4294967295:
        p0 -= p0
    case -295:
        switch "a" {
        case "hello":
            p0 = p0
        case "a":
            switch -172 {
            case 127:
                if (false && (("a" + "\n\t") > "")) {
                    v8 := []int{990}
                    _ = v8
                } else {
                    v9 := GenS1{F0: 4294967295, F1: 1, F2: -9223372036854775808}
                    fmt.Println(p0)
                    _ = v9
                }
            case 731:
                v10 := []int{-1, 2147483647, -2147483648}
                _ = v10
            }
        case "foo bar":
            var v11 uint64 = 18446744073709551615
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 4; fi12++ {
            if false {
                v13 := GenS1{F0: fi12, F1: 0, F2: 42}
                v14 := []int{(fi12 << ((0) & 63)), (fi12 << ((0) & 63)), fi12, fi12}
                var v15 uint = 42
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
    }
    fmt.Println(p0)
    var v16 int32 = 1
    _ = p0
    _ = v16
    return (p0 << ((2) & 63))
}

func main() {
    var v1 int16 = -702
    var v2 int32 = 42
    v3 := []string{"foo bar", ""}
    var v4 uint64 = 420
    v5 := make(map[uint32]bool)
    var v17 uint32 = 0
    if (genFn0(v4) == genFn0(v4)) {
        v18 := GenS1{F0: 32767, F1: 72, F2: 732}
        if (!true) {
            _ = len(v3)
        } else {
            v4 -= v4
        }
        _ = v18
    } else {
        for fi19 := 0; fi19 < 1; fi19++ {
            switch fi19 {
            case 127:
                var v20 uint32 = v17
                _ = v20
            }
            v21 := GenS1{F0: 128, F1: 121, F2: 0}
            var v22 uint64 = v4
            _ = fi19
            _ = v21
            _ = v22
        }
        var v23 uint16 = 440
        _ = v23
    }
    v24 := v5[v17]
    v17 = 633
    v4 = v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v25 int8 = -128
    v1 -= (-(v1))
    var v26 float32 = float32(((0.1 - 1e10) / (1.5 + float64((0.0 * 1.5)))))
    v26 = (((v26 * v26) + v26) * 3.14159)
    if ((((v2 % ((-2147483648) | 1)) != -2147483648) && true) && ((!(42 == 219)) && v24)) {
        var v27 int32 = (v2 + v2)
        _ = len(v3)
        _ = v27
    }
    var v28 bool = v24
    v24 = v28
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
