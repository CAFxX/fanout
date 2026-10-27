package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint64; F2 string }
func genFn0() uint8 {
    var v4 int = -9223372036854775808
    for fi5 := 0; fi5 < 20; fi5++ {
        if false {
            var v6 float64 = 3.14159
            _ = v6
        }
        switch (fi5 % ((v4) | 1)) {
        case 9223372036854775807:
            var v7 bool = (!true)
            _ = v7
        case 0:
            var v8 uint32 = 2147483647
            _ = v8
        case -951:
            var v9 uint64 = 4294967295
            _ = v9
        }
        v10 := []int{fi5, fi5, v4}
        _ = fi5
        _ = v10
    }
    v4 += v4
    _ = v4
    return 237
}
func genFn1(p0 uint64) int {
    fmt.Println(p0)
    p0 -= 446
    _ = p0
    return 0
}

func main() {
    var v1 uint32 = 557
    var v2 uint64 = 725
    var v3 int = 9223372036854775807
    if (((float64(float32((3.14159 * 0.0))) - 1e10) / 0.1) < (3.14159 / 1.0)) {
        v11 := GenS1{F0: 9223372036854775807, F1: (v3 * (v3 / ((-2147483648) | 1))), F2: genFn0()}
        if true {
            var v12 uint = uint(v2)
            var v13 string = "hello"
            _ = v12
            _ = v13
        } else {
            v2 = 75
            var v14 int = (v3 | 2147483647)
            fmt.Println(v3)
            _ = v14
        }
        _ = v11
    } else {
        if true {
            var v15 int8 = -1
            if (v1 != v1) {
                if (0 <= 42) {
                    switch "foo bar" {
                    case "foo bar":
                        v3 = 128
                    }
                    v15 += 100
                    v16 := []int{genFn1((^(v2)))}
                    _ = v16
                }
                v17 := GenS2{F0: genFn0()}
                if true {
                    var v18 string = "hello"
                    var v19 uint16 = 698
                    _ = v18
                    _ = v19
                }
                _ = v17
            }
            _ = v15
        } else {
            v20 := []int{32767, v3}
            var v21 float64 = 0.0
            _ = v20
            _ = v21
        }
    }
    if false {
        if ("hello" == "") {
            v22 := []int{v3, genFn1(v2), 971}
            _ = v22
        } else {
            v3 += v3
            v23 := GenS3{F0: 32767, F1: 1, F2: v2}
            _ = v23
        }
        v24 := []int{}
        var v25 int8 = -29
        _ = v24
        _ = v25
    } else {
        v26 := GenS2{F0: genFn0()}
        _ = v26
    }
    v27 := []int{v3, (^(v3)), v3, (v3 & 1)}
    var v28 uint64 = (v2 / ((v2) | 1))
    if (0 > 705) {
        v1 = 2147483647
    } else {
        var v29 uint64 = 2
        var v30 uint = 255
        _ = v29
        _ = v30
    }
    fmt.Println(v1)
    switch -40 {
    case -990:
        if false {
            v31 := GenS1{F0: 1, F1: 2147483647, F2: 127}
            _ = v31
        }
    case 42:
        v27 = append(v27, v3)
    case 127:
        v2 -= 686
    default:
        for fi32 := 0; fi32 < 11; fi32++ {
            v2 -= v2
            var v33 uint = 42
            if (((42 == (fi32 << (((^(v33))) & 63))) || ((649 >= 4) && (0.1 != 1.5))) || (v1 >= v1)) {
                v34 := GenS4{F0: uint8(v1), F1: uint64(v1), F2: ("foo bar" + "a")}
                v27 = append(v27, (^(v3)))
                v2 -= uint64(v1)
                _ = v34
            }
            _ = fi32
            _ = v33
        }
    }
    fmt.Println(v2, v2)
    v1 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
