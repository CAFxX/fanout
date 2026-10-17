package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 bool }
func genFn0() string {
    var v7 int64 = -9223372036854775808
    fmt.Println(v7)
    var v8 string = (("foo bar" + "foo bar") + (("\n\t" + "foo bar") + "a"))
    v9 := []int{}
    _ = v7
    _ = v8
    _ = v9
    return v8
}
func genFn1(p0 uint64) int16 {
    var v10 int64 = 1
    p0 -= 1
    p0 -= 18446744073709551615
    _ = p0
    _ = v10
    return 1
}

func main() {
    var v1 uint = 65535
    v2 := []int{-1, 32768, 914}
    v3 := []int{}
    var v4 float64 = float64(0.0)
    v5 := []int{}
    var v6 uint64 = 4294967295
    for fi11 := 0; fi11 < 20; fi11++ {
        if true {
            v4 -= -2.25
            v4 = 3.14159
            var v12 int64 = 1
            _ = v12
        }
        _ = fi11
    }
    fmt.Println(v1, v6, v1)
    v6 -= (v6 / ((v6) | 1))
    _ = len(v5)
    v13 := GenS1{F0: -128}
    var v14 int64 = 9223372036854775807
    v1 += 303
    v14 = 9223372036854775807
    fmt.Println(v1, v1)
    v14 -= (^(v14))
    var v15 uint16 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v16 int16 = 1000
    v1 -= 255
    if true {
        if false {
            v4 = v4
        } else {
            var v17 int64 = v14
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 uint8 = 2
                v20 := GenS1{F0: 100}
                var v21 int8 = -128
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            _ = v17
        }
        if true {
            v6 -= v6
            for fi22 := 0; fi22 < 7; fi22++ {
                v23 := GenS1{F0: 91}
                _ = fi22
                _ = v23
            }
        }
        var v24 uint8 = 42
        _ = v24
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 uint16 = 0
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 bool = (((v15 % ((v26) | 1)) <= v26) && true)
            switch -437 {
            case 4294967295:
                v29 := GenS2{F0: 4294967295, F1: (v28 || false)}
                _ = v29
            case 32767:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case -9223372036854775808:
                if (!v28) {
                    v4 -= (v4 + float64(0.1))
                    v6 += (v6 / (((v6 % ((uint64(v1)) | 1))) | 1))
                    var v30 int16 = 32767
                    _ = v30
                } else {
                    v28 = (v28 || v28)
                }
            default:
                v31 := GenS1{F0: 1}
                _ = v31
            }
            if len(v2) > 0 {
                v2[0] = (fi27 & 284)
            }
            _ = fi27
            _ = v28
        }
        if true {
            v26 = (v15 & 255)
        } else {
            var v32 int = 4294967295
            _ = v32
        }
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
