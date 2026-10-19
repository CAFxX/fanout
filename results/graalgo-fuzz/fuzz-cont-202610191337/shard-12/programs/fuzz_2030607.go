package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]bool)
    var v3 float32 = 1.5
    var v4 int8 = 1
    var v5 float32 = (0.0 - 1e10)
    v4 = (v4 << ((4294967295) & 63))
    v3 += v5
    if true {
        var v6 float32 = v5
        var v7 int16 = 1
        _ = v6
        _ = v7
    }
    var v8 string = "foo bar"
    v9 := GenS1{F0: 4294967295}
    if (1 <= 0) {
        v10 := []int{128, 1, 0}
        _ = v10
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := []int{(fi11 >> ((65535) & 63)), fi11, fi11}
        var v13 float32 = (v5 + v3)
        _ = fi11
        _ = v12
        _ = v13
    }
    switch 128 {
    case 2147483647:
        var v14 int32 = 0
        _ = v14
    case -9223372036854775808:
        if (v1 > v1) {
            v5 -= (((((v5 + (0.1 - v3)) - -2.25) + v5) + v5) * 0.0)
            v8 = v8
        } else {
            var v15 string = ((v8 + ("hello" + v8)) + (v8 + (("foo bar" + v8) + v8)))
            var v16 bool = false
            v1 = (v1 / ((v1) | 1))
            _ = v15
            _ = v16
        }
    case -2147483648:
        v17 := GenS2{F0: (v8 + ("" + ((v8 + (v8 + v8)) + v8)))}
        _ = v17
    }
    delete(v2, 738)
    var v18 uint16 = 484
    v18 -= (^(v18))
    var v19 int = -124
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := []int{fi20}
        v4 = v4
        if true {
            var v22 int64 = -9223372036854775808
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    if (691 < -325) {
        for fi23 := 0; fi23 < 18; fi23++ {
            v24 := []int{int(v4)}
            _ = fi23
            _ = v24
        }
        switch (v19 & v19) {
        case -1:
            var v25 bool = false
            _ = v25
        default:
            if ((168 >= 127) || ((true && false) && false)) {
                fmt.Println(v1)
                v4 += -4
                var v26 float64 = float64(v5)
                _ = v26
            }
        }
    } else {
        if (!false) {
            v27 := GenS1{F0: 65535}
            fmt.Println(v4)
            for fi28 := 0; fi28 < 3; fi28++ {
                var v29 float64 = (-2.25 / (float64(v3) + 3.14159))
                if true {
                    if true {
                        switch "hello" {
                        case "\n\t":
                            v30 := v2[2147483648]
                            _ = v30
                        default:
                            v1 = (v1 / ((2) | 1))
                        }
                    }
                } else {
                    for fi31 := 0; fi31 < 2; fi31++ {
                        v4 = (v4 << ((255) & 63))
                        _ = fi31
                    }
                    v1 -= (v1 >> ((65535) & 63))
                }
                v29 = v29
                _ = fi28
                _ = v29
            }
            _ = v27
        } else {
            v32 := v2[2147483647]
            _ = v32
        }
    }
    var v33 string = v8
    var v34 int64 = -1
    v33 = "\n\t"
    var v35 int8 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
