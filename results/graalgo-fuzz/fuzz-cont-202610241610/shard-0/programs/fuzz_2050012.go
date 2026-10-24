package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint, p1 int8, p2 string) uint {
    p0 = (p0 / ((p0) | 1))
    v4 := []int{-1, 4294967295}
    p2 = (("a" + (p2 + p2)) + "\n\t")
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 778
}

func main() {
    var v1 int = 32767
    v2 := make(map[uint32]bool)
    var v3 string = ("" + ((("foo bar" + "a") + "foo bar") + "\n\t"))
    var v5 bool = false
    var v6 int64 = -51
    v7 := GenS1{F0: 0, F1: 32767}
    v3 = v3
    v8 := GenS1{F0: 586, F1: 65535}
    if v5 {
        var v9 float32 = (((1e10 / 1.0) - 0.0) - 0.0)
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(v1, v3, v6)
            fmt.Println(v5, v1)
            _ = fi10
        }
        _ = v9
    } else {
        for fi11 := 0; fi11 < 8; fi11++ {
            if (42 <= 42) {
                v5 = false
            } else {
                var v12 int32 = -1
                _ = v12
            }
            if (v5 && v5) {
                for fi13 := 0; fi13 < 16; fi13++ {
                    var v14 uint16 = 1
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        var v15 uint32 = 2147483648
                        _ = v15
                    case "a":
                        var v16 int8 = 100
                        _ = v16
                    default:
                        var v17 uint8 = 0
                        _ = v17
                    }
                    var v18 int = (v1 % ((v1) | 1))
                    _ = fi13
                    _ = v14
                    _ = v18
                }
                v19 := []int{}
                fmt.Println(fi11, v5, fi11)
                _ = v19
            }
            v6 = (v6 | v6)
            _ = fi11
        }
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        v5 = false
        if ((-1 < 127) && (171 > -378)) {
            v21 := []int{(^(v1)), (fi20 % (((v1 ^ v1)) | 1)), v1}
            _ = v21
        }
        _ = fi20
    }
    v6 -= (^(v6))
    var v22 int16 = 685
    fmt.Println(v5, v5)
    var v23 float32 = 0.1
    var v24 uint8 = 1
    if ((v6 + v6) <= v6) {
        v25 := []int{v1, int(v6), v1}
        _ = v25
    }
    var v26 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
