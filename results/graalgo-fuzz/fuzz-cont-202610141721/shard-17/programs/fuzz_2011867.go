package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
func genFn0(p0 int64, p1 bool, p2 int64) int8 {
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 float32 = 3.14159
        _ = fi4
        _ = v5
    }
    var v6 int = -330
    v7 := []int{(^(v6)), v6, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 100
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int16 = 32767
    v3 := []int{-159, v1, v1}
    var v8 int64 = 1
    var v9 uint8 = 127
    switch ("a" + "foo bar") {
    case "hello":
        var v10 uint16 = 0
        _ = v10
    case "a":
        v2 = (v2 * (v2 % (((v2 % (((v2 >> ((255) & 63))) | 1))) | 1)))
    default:
        if (4294967295 >= 364) {
            v11 := GenS1{F0: v8, F1: false}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 5; fi12++ {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                if (v2 != (v2 & v2)) {
                    v9 = (v9 ^ 255)
                    v8 += (v8 & (v8 << ((255) & 63)))
                    var v13 float64 = -2.25
                    _ = v13
                }
                v14 := GenS1{F0: 752, F1: true}
                _ = fi12
                _ = v14
            }
            v15 := GenS1{F0: (v8 >> ((65535) & 63)), F1: false}
            fmt.Println(v8)
            _ = v15
        }
    }
    if true {
        v2 = 32767
        var v16 uint32 = 1
        fmt.Println(v16)
        _ = v16
    }
    fmt.Println(v9, v9, v8)
    v17 := GenS1{F0: (v8 & v8), F1: true}
    v8 = v8
    v8 = v8
    var v18 uint8 = 163
    if (true || false) {
        if true {
            if false {
                v8 -= (^(v8))
                v19 := GenS1{F0: 503, F1: (!(((false && false) || (-2147483648 != 471)) || false))}
                var v20 float32 = 3.14159
                _ = v19
                _ = v20
            }
        }
        if (!true) {
            switch "hello" {
            case "foo bar":
                if ((612 != 447) || ((v1 * 913) != (^(v1)))) {
                    var v21 uint8 = v18
                    v18 = (v21 + 0)
                    _ = v21
                }
            case "":
                var v22 string = "hello"
                _ = v22
            case "a":
                for fi23 := 0; fi23 < 2; fi23++ {
                    var v24 int32 = 0
                    v24 -= 1
                    var v25 uint64 = 0
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
        }
        for fi26 := 0; fi26 < 19; fi26++ {
            for fi27 := 0; fi27 < 8; fi27++ {
                v18 -= v9
                var v28 int16 = -994
                v28 = v28
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        v3 = append(v3, 127)
        if false {
            v18 = (v18 - (v9 * (v18 / ((v18) | 1))))
        } else {
            v9 -= 1
            var v30 float64 = 0.1
            _ = v30
        }
        _ = len(v3)
        _ = fi29
    }
    fmt.Println(v8)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
