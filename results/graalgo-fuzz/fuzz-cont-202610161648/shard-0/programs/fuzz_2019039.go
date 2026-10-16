package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int16, p1 uint, p2 int64) int16 {
    var v4 uint64 = 42
    switch 32768 {
    case -625:
        if true {
            p0 = (^(p0))
        }
    }
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := []int{(fi5 & (fi5 % ((128) | 1))), fi5}
        var v7 float64 = 0.0
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p0
}
func genFn1(p0 uint, p1 int) uint16 {
    var v8 int64 = -9223372036854775808
    for fi9 := 0; fi9 < 6; fi9++ {
        v8 += (^(v8))
        v10 := []int{}
        _ = fi9
        _ = v10
    }
    v11 := []int{p1, (p1 - -2147483648), p1}
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return 1
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []string{}
    v3 := []int{4294967295, -114, 127, -264}
    for fi12 := 0; fi12 < 14; fi12++ {
        if (!false) {
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := GenS1{F0: 127}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi13
                _ = v14
            }
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        if (false && (v1 >= 931)) {
            v1 -= v1
            var v16 uint16 = 65535
            _ = v16
        } else {
            var v17 float64 = 3.14159
            fmt.Println(v17, v17)
            _ = v17
        }
        v1 -= v1
        _ = fi15
    }
    v1 = v1
    if false {
        fmt.Println(v1)
        var v18 int64 = -217
        _ = v18
    }
    v1 += v1
    var v19 int64 = 223
    for fi20 := 0; fi20 < 10; fi20++ {
        v1 = (v1 >> ((65535) & 63))
        _ = len(v2)
        if true {
            v1 = 65535
        } else {
            var v21 uint = 141
            _ = len(v2)
            _ = v21
        }
        _ = fi20
    }
    v19 = (v19 + (v19 >> ((255) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
