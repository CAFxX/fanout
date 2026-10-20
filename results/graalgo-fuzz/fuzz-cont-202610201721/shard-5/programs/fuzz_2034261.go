package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 string }
func genFn0() int32 {
    var v7 uint64 = 2
    var v8 uint8 = 42
    fmt.Println(v8, v7)
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int16, p1 uint16) uint {
    v9 := GenS1{F0: 818, F1: 4294967295}
    p1 += 32767
    p1 += (p1 + (p1 & p1))
    v10 := GenS2{F0: false, F1: -2147483648}
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS3{F0: "\n\t"}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 814
}

func main() {
    var v1 int16 = -740
    var v2 int16 = -272
    var v3 float64 = (1e10 / 0.1)
    var v4 int = -2147483648
    v5 := []int{v4, v4, (v4 ^ (v4 / ((v4) | 1)))}
    v6 := []int{v4, v4, v4}
    if (v4 == (v4 | 1)) {
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 float64 = (0.1 + -2.25)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi13
            _ = v14
        }
        var v15 string = ("hello" + ("a" + "a"))
        _ = v15
    } else {
        v1 = v1
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v16 float64 = v3
    if false {
        var v17 int16 = v1
        v3 = (1.5 + 3.14159)
        fmt.Println(v4, v4, v3)
        _ = v17
    } else {
        v18 := GenS2{F0: (!false), F1: -2147483648}
        _ = v18
    }
    v3 -= v3
    var v19 uint = 18446744073709551615
    if (true || true) {
        var v20 int8 = 1
        v21 := GenS3{F0: (("" + "a") + "\n\t")}
        _ = v20
        _ = v21
    } else {
        v16 = v16
        var v22 int32 = 1
        for fi23 := 0; fi23 < 2; fi23++ {
            for fi24 := 0; fi24 < 8; fi24++ {
                var v25 bool = false
                v5 = append(v5, fi24)
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
        _ = v22
    }
    fmt.Println(v3, v3)
    fmt.Println(v16, v1)
    switch (v4 + v4) {
    case 0:
        v5 = append(v5, (v4 / ((v4) | 1)))
    case 742:
        if len(v6) > 0 {
            v6[0] = 127
        }
    }
    v16 += v3
    fmt.Println(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
