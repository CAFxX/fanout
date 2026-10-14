package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int16 }
func genFn0() uint16 {
    if (2147483648 < 90) {
        var v5 uint = 4294967295
        fmt.Println(v5)
        if (!((!(42 <= 127)) || false)) {
            v5 = (v5 * (v5 >> (((v5 % (((^(v5))) | 1))) & 63)))
        }
        _ = v5
    } else {
        v6 := []int{1}
        v7 := GenS1{F0: (42 <= 52)}
        v8 := GenS1{F0: (!false)}
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint32 = 0
    switch "\n\t" {
    case "foo bar":
        var v10 float64 = (-2.25 / 3.14159)
        _ = v10
    }
    v9 += 45
    _ = v9
    return 111
}
func genFn1(p0 uint64, p1 string) string {
    p1 = p1
    p1 = ("hello" + p1)
    p0 = p0
    if false {
        p1 = (p1 + p1)
        var v11 int32 = 285
        _ = v11
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        p1 = p1
        _ = fi12
    }
    _ = p0
    _ = p1
    return "a"
}

func main() {
    var v1 int32 = -1
    var v2 uint8 = 150
    var v3 int32 = v1
    var v4 uint8 = 127
    var v13 float32 = 1e10
    v14 := GenS1{F0: (false && (-123 > 1))}
    v15 := []int{}
    fmt.Println(v13)
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    v15 = append(v15, 32767)
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS2{F0: 42, F1: 42}
        if (180 == (v2 & 0)) {
            var v18 int32 = 42
            v19 := GenS2{F0: 832, F1: -1}
            _ = v18
            _ = v19
        } else {
            var v20 int32 = v3
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    v21 := GenS1{F0: (!(1.5 <= 0.0))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
