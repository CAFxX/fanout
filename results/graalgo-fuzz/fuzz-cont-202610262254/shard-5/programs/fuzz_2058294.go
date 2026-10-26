package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS2 struct { F0 string; F1 int8 }
func genFn0(p0 string) int8 {
    var v4 uint16 = 1
    var v5 int = 461
    _ = p0
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 bool) int8 {
    p0 = (0.1 < -2.25)
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            fmt.Println(p0)
            _ = fi7
        }
        _ = fi6
    }
    v8 := GenS1{F0: 133, F1: 0, F2: 0}
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 uint64 = 4294967295
        var v11 int8 = -45
        v12 := []int{fi9}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v8
    return 124
}

func main() {
    var v1 float32 = ((1.5 / (0.0 + 1.5)) - 0.0)
    var v2 int64 = -981
    var v3 int8 = 127
    var v13 int64 = (v2 & v2)
    var v14 bool = (true || true)
    var v15 string = "foo bar"
    v16 := GenS1{F0: 2, F1: 497, F2: 1}
    var v17 uint = 2
    if v14 {
        if (!(v14 && (138 >= 855))) {
            var v18 float64 = 0.1
            var v19 uint16 = 1
            for fi20 := 0; fi20 < 20; fi20++ {
                v21 := []int{fi20, (-(fi20))}
                var v22 uint8 = 42
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v18
            _ = v19
        }
        v2 = (v13 % ((-394) | 1))
        var v23 int64 = v13
        _ = v23
    } else {
        for fi24 := 0; fi24 < 17; fi24++ {
            v25 := GenS1{F0: 0, F1: 0, F2: -307}
            _ = fi24
            _ = v25
        }
    }
    var v26 string = v15
    for fi27 := 0; fi27 < 2; fi27++ {
        for fi28 := 0; fi28 < 11; fi28++ {
            switch fi28 {
            case 9223372036854775807:
                v29 := GenS2{F0: v15, F1: int8(fi27)}
                _ = v29
            case 32767:
                fmt.Println(v26, v26, v17)
            }
            _ = fi28
        }
        _ = fi27
    }
    switch (v26 + v15) {
    case "\n\t":
        v26 = v15
    }
    if v14 {
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 int16 = -32768
            _ = fi30
            _ = v31
        }
    }
    v15 = v15
    v14 = v14
    switch v26 {
    case "":
        v32 := []int{42}
        _ = v32
    case "\n\t":
        var v33 uint16 = 0
        _ = v33
    default:
        v15 = v15
    }
    var v34 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v34
}
