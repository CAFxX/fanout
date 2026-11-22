package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int64; F2 uint8 }
func genFn0(p0 uint, p1 uint64, p2 int) uint64 {
    var v7 int32 = 0
    if false {
        if true {
            v8 := GenS1{F0: 159}
            switch p2 {
            case 9223372036854775807:
                p0 -= p0
            case 32768:
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 uint = (p0 - 636)
                    fmt.Println(fi9, p1, fi9)
                    for fi11 := 0; fi11 < 12; fi11++ {
                        p0 += v10
                        _ = fi11
                    }
                    _ = fi9
                    _ = v10
                }
            case 340:
                if (1 <= 42) {
                    p0 = p0
                } else {
                    v12 := []int{(-(p2)), p2, p2, -9223372036854775808}
                    _ = v12
                }
            }
            _ = v8
        } else {
            var v13 int8 = int8(p1)
            var v14 bool = false
            v15 := GenS1{F0: -127}
            _ = v13
            _ = v14
            _ = v15
        }
        var v16 uint = p0
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := []int{p2, p2, (fi17 ^ p2)}
            var v19 float64 = (1e10 - 1e10)
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v16
    } else {
        var v20 string = "foo bar"
        _ = v20
    }
    var v21 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v7
    return (^(p1))
}
func genFn1(p0 bool, p1 int) uint16 {
    var v22 uint64 = genFn0(1, 18446744073709551615, p1)
    v23 := GenS1{F0: 1}
    p1 -= 2147483647
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return 309
}

func main() {
    v1 := []int{4294967295}
    var v2 bool = true
    v3 := []uint32{214, 48, 2, 625}
    var v4 uint32 = 0
    var v5 bool = v2
    v6 := make(map[int]int)
    v4 += v4
    if v5 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v5 = true
    var v24 int16 = -135
    var v25 int32 = 287
    fmt.Println(v2)
    if v5 {
        v2 = (1 == 70)
    } else {
        v26 := GenS1{F0: v24}
        v27 := GenS2{F0: (!v2), F1: 42, F2: 210}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v26
        _ = v27
    }
    delete(v6, 1)
    v25 = (v25 | v25)
    if (0 >= (v25 | int32(v4))) {
        var v28 int16 = (v24 | -32768)
        _ = v28
    }
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
