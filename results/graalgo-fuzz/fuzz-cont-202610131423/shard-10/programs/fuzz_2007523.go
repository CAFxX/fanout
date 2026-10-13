package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int16; F1 string }
func genFn0() int8 {
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 string = ("hello" + "a")
            v7 := []int{(fi5 * fi4), (fi5 << ((0) & 63)), fi4, fi5}
            _ = fi5
            _ = v6
            _ = v7
        }
        switch ("hello" + "hello") {
        case "foo bar":
            var v8 int = fi4
            _ = v8
        case "\n\t":
            v9 := []int{(fi4 & (-(fi4))), fi4}
            _ = v9
        default:
            var v10 uint8 = 255
            _ = v10
        }
        _ = fi4
    }
    if true {
        var v11 uint64 = 42
        fmt.Println(v11)
        _ = v11
    } else {
        var v12 float32 = 1e10
        var v13 int8 = 100
        _ = v12
        _ = v13
    }
    var v14 bool = true
    if v14 {
        v14 = v14
        var v15 float64 = 3.14159
        for fi16 := 0; fi16 < 3; fi16++ {
            if v14 {
                var v17 int32 = -2147483648
                var v18 float32 = float32(v15)
                v19 := GenS1{F0: 4294967295, F1: 1}
                _ = v17
                _ = v18
                _ = v19
            } else {
                v20 := GenS2{F0: 42, F1: (("hello" + (("a" + "hello") + ("hello" + "foo bar"))) + "")}
                _ = v20
            }
            _ = fi16
        }
        _ = v15
    } else {
        v21 := []int{}
        for fi22 := 0; fi22 < 10; fi22++ {
            v14 = v14
            _ = fi22
        }
        _ = v21
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        var v24 uint16 = 709
        var v25 bool = true
        v26 := GenS2{F0: -569, F1: "\n\t"}
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    _ = v14
    return 100
}

func main() {
    var v1 uint8 = 0
    v2 := []uint32{4294967295, 2, 294}
    v3 := make(map[int]uint32)
    switch "" {
    case "a":
        for fi27 := 0; fi27 < 4; fi27++ {
            if ((!((-2.25 * -2.25) <= (0.0 - 1.5))) && (0.0 != 1.5)) {
                fmt.Println(fi27, v1)
            }
            v28 := GenS1{F0: uint(v1), F1: -861}
            v29 := GenS2{F0: 0, F1: "\n\t"}
            _ = fi27
            _ = v28
            _ = v29
        }
    default:
        var v30 int = 1
        _ = v30
    }
    var v31 int8 = -35
    v1 += v1
    v1 -= (^(v1))
    if true {
        _ = len(v3)
        var v32 int = -496
        _ = v32
    }
    v3[-975] = 0
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 int32 = 42
        _ = fi33
        _ = v34
    }
    fmt.Println(v31)
    var v35 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
}
