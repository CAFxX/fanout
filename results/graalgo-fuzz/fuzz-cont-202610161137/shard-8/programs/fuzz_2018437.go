package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 bool; F1 int }
func genFn0() int64 {
    var v7 uint64 = 389
    for fi8 := 0; fi8 < 10; fi8++ {
        v7 = v7
        _ = fi8
    }
    if (-9223372036854775808 == -1) {
        for fi9 := 0; fi9 < 1; fi9++ {
            v7 = (v7 & (v7 % ((170) | 1)))
            var v10 int32 = 2147483647
            _ = fi9
            _ = v10
        }
    } else {
        var v11 uint16 = 1
        switch "" {
        case "":
            for fi12 := 0; fi12 < 19; fi12++ {
                for fi13 := 0; fi13 < 14; fi13++ {
                    v14 := []int{fi12, fi13, 42, fi12}
                    v15 := GenS1{F0: (v7 << ((953) & 63)), F1: 1, F2: ""}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v16 float64 = 1.5
                _ = fi12
                _ = v16
            }
        case "foo bar":
            v11 = (^(v11))
        case "hello":
            v11 -= 209
        }
        _ = v11
    }
    _ = v7
    return -1
}
func genFn1(p0 string, p1 int64, p2 int8) int32 {
    var v17 float32 = 0.1
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return -423
}

func main() {
    var v1 float32 = (1.5 - 1e10)
    var v2 uint16 = 32767
    v3 := make(map[uint32]bool)
    var v4 bool = true
    var v5 uint16 = v2
    var v6 int8 = 0
    var v18 int = 32767
    fmt.Println(v5)
    v6 += (v6 - (v6 << ((4294967295) & 63)))
    var v19 uint16 = (v2 / ((v2) | 1))
    fmt.Println(v4)
    if (654 <= 135) {
        v18 = v18
        switch ((("" + "foo bar") + "\n\t") + ("\n\t" + "\n\t")) {
        case "":
            v20 := GenS2{F0: 255}
            _ = v20
        case "a":
            v19 = (v2 % ((v5) | 1))
        default:
            v21 := []int{v18}
            _ = v21
        }
    }
    if true {
        v22 := GenS3{F0: v5, F1: 4294967295}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 5; fi23++ {
            fmt.Println(v2, v1)
            for fi24 := 0; fi24 < 2; fi24++ {
                v25 := GenS3{F0: (v5 ^ (v5 + (v2 & (^(v2))))), F1: 942}
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS3{F0: v19, F1: 2}
            for fi28 := 0; fi28 < 9; fi28++ {
                v4 = v4
                _ = fi28
            }
            var v29 int8 = 1
            _ = fi26
            _ = v27
            _ = v29
        }
    }
    var v30 bool = (4294967295 <= 0)
    v31 := GenS4{F0: true, F1: int(v2)}
    v6 = 60
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 int64 = genFn0()
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        fmt.Println(fi34, v30, v1)
        delete(v3, 2)
        _ = fi34
    }
    v35 := GenS1{F0: 18446744073709551615, F1: 0, F2: ""}
    v2 -= v19
    fmt.Println(v2, v18, v1)
    var v36 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
