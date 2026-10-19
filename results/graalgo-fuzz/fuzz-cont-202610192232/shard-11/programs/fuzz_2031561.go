package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
func genFn0() uint32 {
    v6 := GenS1{F0: 1}
    v7 := []int{}
    _ = v6
    _ = v7
    return 162
}
func genFn1(p0 bool, p1 int32) string {
    v8 := []int{2147483647, 15, 128, 9223372036854775807}
    for fi9 := 0; fi9 < 4; fi9++ {
        if len(v8) > 0 {
            v8[0] = (fi9 % ((fi9) | 1))
        }
        if p0 {
            p0 = (!(765 > 955))
            if (1.5 != 1e10) {
                for fi10 := 0; fi10 < 13; fi10++ {
                    v11 := GenS1{F0: 2}
                    _ = fi10
                    _ = v11
                }
                switch fi9 {
                case -9223372036854775808:
                    var v12 float32 = 1e10
                    _ = v12
                case 42:
                    v13 := GenS1{F0: 0}
                    _ = v13
                case -466:
                    p0 = ((125 == 279) && true)
                }
            } else {
                if true {
                    v14 := GenS1{F0: 18446744073709551615}
                    var v15 uint = 65535
                    v8 = append(v8, (fi9 % (((fi9 ^ (fi9 >> (((v15 % (((^(v15))) | 1))) & 63)))) | 1)))
                    _ = v14
                    _ = v15
                }
                var v16 uint64 = 1
                _ = v16
            }
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v8
    return "\n\t"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 bool = (-253 != -646)
    var v5 int8 = -3
    switch 709 {
    case 2147483647:
        v5 -= (-(v5))
    }
    var v17 uint32 = 42
    if (229 >= 1000) {
        _ = len(v3)
        if (978 == 0) {
            v17 += v17
            var v18 uint32 = (^(v17))
            _ = v18
        }
        var v19 int8 = (v5 - v5)
        _ = v19
    } else {
        fmt.Println(v17, v17, v5)
        delete(v3, ("" + ("foo bar" + (("\n\t" + "hello") + "hello"))))
        if true {
            v20 := GenS2{F0: "foo bar"}
            _ = v20
        }
    }
    if (false || (v4 || (((v17 << ((18446744073709551615) & 63)) <= v17) || false))) {
        v4 = (v4 || (((v17 << ((255) & 63)) < v17) || true))
        if (((v4 && v4) || true) || false) {
            v21 := GenS1{F0: 1}
            _ = v21
        } else {
            v4 = true
            v5 -= v5
            v22 := []int{-503, -556, 32767}
            _ = v22
        }
    } else {
        var v23 bool = false
        v23 = ((genFn0() != v17) && v4)
        v4 = ((((3.14159 != (float32(0.0) / 0.1)) || (true && false)) || (-2147483648 == 2147483647)) || (-1 > (v5 ^ v5)))
        _ = v23
    }
    v24 := GenS3{F0: -67}
    v25 := []int{}
    var v26 int = 9223372036854775807
    v26 += (v26 << ((48) & 63))
    v17 -= genFn0()
    for fi27 := 0; fi27 < 1; fi27++ {
        delete(v1, "foo bar")
        switch "foo bar" {
        case "hello":
            v26 = 1
        case "a":
            if v4 {
                switch v26 {
                case 786:
                    var v28 uint32 = (v17 / ((v17) | 1))
                    _ = v28
                case -721:
                    v29 := GenS3{F0: (v26 * v26)}
                    _ = v29
                case 2147483647:
                    var v30 bool = true
                    _ = v30
                }
            }
        }
        _ = fi27
    }
    fmt.Println(v4)
    v5 += v5
    v5 = (v5 % (((v5 >> ((1) & 63))) | 1))
    v26 += (v26 << ((4294967295) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
