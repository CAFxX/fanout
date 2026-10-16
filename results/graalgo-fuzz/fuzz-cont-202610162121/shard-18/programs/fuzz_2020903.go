package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
func genFn0(p0 uint16, p1 uint, p2 int) int {
    p1 = (p1 ^ p1)
    var v6 int16 = 0
    if true {
        v7 := []int{p2, -238, -9223372036854775808, 9223372036854775807}
        p0 = 1
        _ = v7
    }
    fmt.Println(p1)
    v8 := GenS1{F0: (^(p2)), F1: -564}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return (^(p2))
}

func main() {
    var v1 uint8 = 255
    var v2 uint8 = (v1 - 255)
    v3 := make(map[int]int)
    v4 := []uint32{750, 0, 4294967295, 42}
    var v5 uint64 = 18446744073709551615
    var v9 string = "hello"
    var v10 bool = ((0 >= -32768) && ((0 != -567) || (-349 <= 1)))
    for fi11 := 0; fi11 < 18; fi11++ {
        if v10 {
            _ = len(v3)
            for fi12 := 0; fi12 < 17; fi12++ {
                fmt.Println(v5, v2, v10)
                v13 := v3[(fi11 << ((882) & 63))]
                v14 := GenS1{F0: genFn0(236, 127, v13), F1: 1}
                _ = fi12
                _ = v13
                _ = v14
            }
        }
        _ = fi11
    }
    if v10 {
        _ = len(v4)
    }
    var v15 uint = 4294967295
    delete(v3, 140)
    for fi16 := 0; fi16 < 2; fi16++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v17 uint = v15
        _ = fi16
        _ = v17
    }
    var v18 uint64 = v5
    var v19 int8 = 92
    for fi20 := 0; fi20 < 9; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := GenS1{F0: fi20, F1: -1}
            switch (fi21 & fi21) {
            case 32768:
                var v23 uint16 = 255
                _ = v23
            case 4294967295:
                if v10 {
                    var v24 int64 = 1
                    _ = v24
                }
            case -747:
                if (v10 && ((0 <= 1) || v10)) {
                    v1 = (^(v1))
                }
            }
            _ = fi21
            _ = v22
        }
        var v25 string = ((v9 + "") + (v9 + "\n\t"))
        _ = fi20
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
