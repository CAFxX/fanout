package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 bool; F1 bool }
func genFn0() string {
    for fi4 := 0; fi4 < 2; fi4++ {
        for fi5 := 0; fi5 < 6; fi5++ {
            v6 := GenS1{F0: 1, F1: 1}
            if false {
                v7 := []int{}
                _ = v7
            }
            _ = fi5
            _ = v6
        }
        switch (fi4 * (^(fi4))) {
        case -329:
            v8 := []int{(fi4 << ((255) & 63)), fi4, (fi4 & -195), 9223372036854775807}
            _ = v8
        }
        _ = fi4
    }
    var v9 int16 = -547
    _ = v9
    return ("hello" + "foo bar")
}
func genFn1(p0 uint8) int {
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := []int{42, (fi10 << ((18446744073709551615) & 63)), fi10}
        var v12 float64 = float64(float32(0.0))
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(p0)
    var v13 float32 = 1e10
    if true {
        v14 := GenS1{F0: 1, F1: -1}
        _ = v14
    }
    _ = p0
    _ = v13
    return -2147483648
}

func main() {
    v1 := []int{32767, 128, 0, 478}
    var v2 uint = 0
    var v3 int = 32767
    v3 += (-(v3))
    var v15 float64 = (0.0 * 1e10)
    v16 := GenS1{F0: 0, F1: -85}
    switch v3 {
    case -290:
        var v17 int64 = -1
        _ = v17
    case 127:
        switch ("foo bar" + "hello") {
        case "foo bar":
            if len(v1) > 0 {
                v1[0] = (v3 / ((v3) | 1))
            }
        case "a":
            var v18 uint8 = 255
            _ = v18
        }
    case 128:
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 uint = v2
            v21 := GenS2{F0: (true && (!((1000 >= 0) && true))), F1: (2 > 1)}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 = (^(fi22))
        v1 = append(v1, 9223372036854775807)
        _ = fi22
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        v2 -= 549
        _ = fi23
    }
    v3 -= (v3 - (v3 % ((genFn1(42)) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
