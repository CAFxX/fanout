package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
func genFn0() int {
    switch -790 {
    case -9223372036854775808:
        for fi5 := 0; fi5 < 20; fi5++ {
            if ("\n\t" > "") {
                if false {
                    var v6 uint = 301
                    _ = v6
                } else {
                    var v7 uint8 = 128
                    _ = v7
                }
                v8 := []int{fi5}
                var v9 uint16 = 967
                _ = v8
                _ = v9
            }
            _ = fi5
        }
    }
    var v10 int = 0
    v10 += 128
    var v11 int32 = -982
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 int32 = -2147483648
    var v2 int32 = 42
    var v3 int16 = -32768
    var v4 uint16 = 255
    switch 32768 {
    case -9223372036854775808:
        v2 -= (v2 ^ v1)
    case -1:
        for fi12 := 0; fi12 < 15; fi12++ {
            for fi13 := 0; fi13 < 16; fi13++ {
                v14 := GenS1{F0: v2}
                v4 -= 42
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
    default:
        v3 = -32768
    }
    v1 -= v1
    for fi15 := 0; fi15 < 12; fi15++ {
        v4 = v4
        if ("a" > "hello") {
            var v16 float64 = (0.0 + 1.5)
            for fi17 := 0; fi17 < 5; fi17++ {
                if true {
                    v18 := []int{(fi17 << ((255) & 63))}
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    var v19 float64 = (v16 + v16)
                    _ = v18
                    _ = v19
                }
                var v20 int32 = v2
                if false {
                    v2 -= -743
                } else {
                    v4 += v4
                }
                _ = fi17
                _ = v20
            }
            _ = v16
        } else {
            v21 := GenS2{F0: 1}
            _ = v21
        }
        _ = fi15
    }
    v3 -= v3
    var v22 int8 = -1
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 int16 = v3
        _ = fi23
        _ = v24
    }
    v1 = v1
    v25 := []int{9223372036854775807}
    v26 := GenS1{F0: v2}
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    switch 132 {
    case 2147483647:
        v3 = v3
    }
    if ((^(v22)) < v22) {
        v27 := GenS2{F0: 2}
        var v28 int16 = 1000
        _ = v27
        _ = v28
    } else {
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 int16 = (v3 * v3)
            var v31 int = (fi29 & (fi29 << ((4294967295) & 63)))
            v22 -= 73
            _ = fi29
            _ = v30
            _ = v31
        }
        v3 = v3
    }
    var v32 uint64 = 301
    var v33 int32 = 42
    v34 := GenS1{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
