package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint16; F2 int64 }
func genFn0(p0 string, p1 uint64, p2 bool) uint64 {
    var v7 int32 = 1
    p1 += p1
    p2 = (float64(0.0) < (1e10 - 1e10))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 805
}
func genFn1(p0 int, p1 uint, p2 string) bool {
    if true {
        v8 := GenS1{F0: 366, F1: 42, F2: 4294967295}
        v9 := []int{p0}
        _ = v8
        _ = v9
    } else {
        v10 := []int{973, p0, (p0 << ((922) & 63))}
        _ = v10
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return (-1 <= 42)
}

func main() {
    var v1 int = 32768
    var v2 string = "a"
    var v3 bool = ((true && ((0.0 - -2.25) > ((((0.1 / 1.0) - 0.0) / 3.14159) - 3.14159))) && false)
    var v4 uint64 = 1
    v5 := make(map[int]string)
    var v6 int = 32768
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: 42, F1: 144, F2: (v4 ^ 304)}
        _ = fi11
        _ = v12
    }
    if v3 {
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := v5[-9223372036854775808]
            _ = fi13
            _ = v14
        }
        v6 += v6
    } else {
        v15 := GenS2{F0: -870, F1: 1, F2: 42}
        fmt.Println(v6)
        _ = v15
    }
    v4 -= v4
    if true {
        if v3 {
            var v16 int32 = -721
            _ = v16
        } else {
            var v17 uint64 = 42
            fmt.Println(v6, v17, v4)
            if false {
                v18 := GenS2{F0: 42, F1: 255, F2: 1}
                _ = v18
            } else {
                v19 := GenS1{F0: 9223372036854775807, F1: 2147483648, F2: v4}
                var v20 uint64 = v17
                fmt.Println(v6)
                _ = v19
                _ = v20
            }
            _ = v17
        }
    }
    var v21 uint32 = 2
    v22 := []int{v6}
    if (!(-2.25 >= 1e10)) {
        for fi23 := 0; fi23 < 1; fi23++ {
            fmt.Println(v1)
            _ = fi23
        }
        switch (v2 + ("foo bar" + "a")) {
        case "a":
            var v24 uint32 = v21
            _ = v24
        default:
            for fi25 := 0; fi25 < 9; fi25++ {
                for fi26 := 0; fi26 < 12; fi26++ {
                    var v27 int16 = 483
                    var v28 bool = true
                    v29 := v5[fi26]
                    v27 += (v27 % ((v27) | 1))
                    v28 = v28
                    v27 -= -32768
                    v6 = fi25
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                _ = fi25
            }
        }
    } else {
        if (v3 || v3) {
            var v30 int64 = 962
            var v31 uint32 = v21
            _ = v30
            _ = v31
        } else {
            v4 += v4
        }
        if (v2 <= ("\n\t" + v2)) {
            v4 -= uint64(v21)
            for fi32 := 0; fi32 < 16; fi32++ {
                v4 = (v4 ^ genFn0(v2, (^(v4)), v3))
                _ = fi32
            }
            delete(v5, v1)
        }
        var v33 int32 = -1
        _ = v33
    }
    var v34 string = ("\n\t" + "a")
    v3 = ((!true) && true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
