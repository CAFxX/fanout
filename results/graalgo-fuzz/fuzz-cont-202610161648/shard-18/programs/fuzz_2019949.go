package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0() int {
    v6 := GenS1{F0: 175}
    var v7 int = 0
    v8 := GenS1{F0: 1}
    v7 = 9223372036854775807
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 uint, p1 uint32, p2 uint64) int16 {
    p0 -= 4294967295
    v9 := GenS2{F0: -9223372036854775808, F1: p0}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1000
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]string)
    var v3 int64 = -437
    var v4 uint = 255
    var v5 uint8 = 0
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        var v10 uint32 = 4294967295
        _ = v10
    case "foo bar":
        var v11 int = -1
        _ = v11
    case "hello":
        v4 += 739
    default:
        switch 80 {
        case -193:
            var v12 float64 = ((((float64(0.0) * float64(1e10)) * v1) - 3.14159) * 0.0)
            _ = v12
        }
    }
    v13 := v2["foo bar"]
    v14 := GenS1{F0: 1}
    fmt.Println(v5, v3, v3)
    v15 := v2[v13]
    switch "" {
    case "":
        v5 -= (v5 ^ 174)
    case "foo bar":
        switch genFn0() {
        case -539:
            var v16 int32 = -1
            _ = v16
        case 0:
            if false {
                fmt.Println(v3)
                for fi17 := 0; fi17 < 11; fi17++ {
                    var v18 float64 = 1.5
                    _ = fi17
                    _ = v18
                }
                v3 = v3
            }
        }
    }
    switch (("hello" + "hello") + "foo bar") {
    case "hello":
        _ = len(v2)
    case "\n\t":
        v15 = "a"
    }
    if (true || (false && true)) {
        var v19 uint32 = 1
        _ = v19
    } else {
        var v20 bool = false
        var v21 uint8 = (v5 ^ v5)
        v2[v13] = "foo bar"
        _ = v20
        _ = v21
    }
    fmt.Println(v13)
    if false {
        if (int8(v3) == 127) {
            v22 := []int{}
            var v23 int16 = 196
            var v24 string = (v13 + (v13 + v13))
            _ = v22
            _ = v23
            _ = v24
        } else {
            v25 := GenS1{F0: 0}
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 bool = (!(!((!(196 > (v5 & (v5 | (v5 / ((255) | 1)))))) || false)))
                _ = fi26
                _ = v27
            }
            v28 := GenS3{F0: 1, F1: -76}
            _ = v25
            _ = v28
        }
    } else {
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 uint = (v4 / ((v4) | 1))
            delete(v2, v13)
            switch (^(fi29)) {
            case 872:
                var v31 uint64 = 957
                _ = v31
            case -828:
                for fi32 := 0; fi32 < 9; fi32++ {
                    fmt.Println(v3, v30)
                    _ = fi32
                }
            default:
                var v33 int64 = v3
                _ = v33
            }
            _ = fi29
            _ = v30
        }
    }
    _ = len(v2)
    var v34 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
