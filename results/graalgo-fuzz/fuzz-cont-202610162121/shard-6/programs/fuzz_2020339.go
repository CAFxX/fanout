package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int32 }
func genFn0() uint32 {
    var v6 int = 32767
    var v7 int16 = 152
    fmt.Println(v7, v7)
    var v8 uint = 0
    var v9 float64 = 3.14159
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint64, p1 bool) uint32 {
    var v10 string = "hello"
    for fi11 := 0; fi11 < 15; fi11++ {
        v12 := []int{fi11, (fi11 - (fi11 % ((-2147483648) | 1)))}
        var v13 bool = (!p1)
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int32 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v14
    return 2147483647
}

func main() {
    var v1 uint32 = 2
    var v2 int = 32767
    v3 := []int{}
    v4 := make(map[string]uint32)
    v5 := []int{(-(v2)), (^(v2))}
    switch (v2 << ((42) & 63)) {
    case 127:
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS1{F0: 4294967295, F1: true}
            _ = fi15
            _ = v16
        }
    case -2147483648:
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 int8 = 100
            if true {
                v18 -= v18
            } else {
                fmt.Println(v18)
            }
            v19 := GenS1{F0: (^(v1)), F1: false}
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    v20 := GenS1{F0: (v1 - (v1 * (v1 * genFn0()))), F1: true}
    v3 = append(v3, v2)
    var v21 float64 = ((0.0 - float64(1.5)) * 0.0)
    fmt.Println(v21)
    if true {
        switch ("" + "foo bar") {
        case "\n\t":
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 float32 = -2.25
                var v24 int16 = 32767
                if true {
                    v2 = (v2 + (^(fi22)))
                    v25 := GenS1{F0: v1, F1: false}
                    _ = v25
                }
                _ = fi22
                _ = v23
                _ = v24
            }
        }
    } else {
        v1 = (^(v1))
    }
    v26 := GenS2{F0: 0}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v27 uint16 = 137
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
