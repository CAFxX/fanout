package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 int32 }

func main() {
    var v1 uint32 = 1
    v2 := []int{-996, -2147483648, 507, 679}
    v3 := []int{298, -1}
    fmt.Println(v1)
    var v4 int16 = 1
    if len(v2) > 0 {
        v2[0] = 42
    }
    v1 += v1
    fmt.Println(v1, v1)
    if true {
        _ = len(v2)
        var v5 uint32 = (^(v1))
        _ = v5
    } else {
        v6 := GenS1{F0: (v4 % ((v4) | 1))}
        _ = v6
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int = fi7
        v1 -= v1
        v1 = (v1 | (v1 | v1))
        _ = fi7
        _ = v8
    }
    if len(v3) > 0 {
        v3[0] = -757
    }
    v1 = (v1 * v1)
    v2 = append(v2, -770)
    v9 := GenS1{F0: 160}
    switch ("a" + "") {
    case "":
        var v10 int8 = -1
        _ = v10
    case "foo bar":
        v1 = (v1 ^ 274)
    case "\n\t":
        v1 = (v1 << ((4294967295) & 63))
    default:
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(v4, fi11, v4)
            _ = fi11
        }
    }
    if true {
        var v12 string = ""
        fmt.Println(v1)
        if true {
            var v13 string = v12
            _ = v13
        } else {
            var v14 uint16 = 1
            _ = v14
        }
        _ = v12
    } else {
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 int64 = 638
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: 2147483647, F1: (v1 << ((182) & 63)), F2: -2147483648}
        fmt.Println(v4, v4)
        _ = v17
    }
    var v18 int8 = 114
    if ((v1 % (((v1 * (v1 << ((255) & 63)))) | 1)) <= (v1 * (v1 & 97))) {
        if ((v18 % ((-128) | 1)) < v18) {
            var v19 string = "foo bar"
            var v20 float32 = (1.5 / float32(((3.14159 - (1e10 - (0.0 - 1e10))) / 1.5)))
            v21 := GenS1{F0: (v4 * 32767)}
            _ = v19
            _ = v20
            _ = v21
        } else {
            switch "foo bar" {
            case "a":
                if (255 != 1) {
                    var v22 uint = 42
                    var v23 float64 = ((3.14159 / (1.5 * float64(-2.25))) - 0.1)
                    _ = v22
                    _ = v23
                }
            case "":
                if len(v3) > 0 {
                    v3[0] = 485
                }
            }
            switch -193 {
            case 901:
                for fi24 := 0; fi24 < 7; fi24++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v25 int8 = (v18 / (((v18 % ((v18) | 1))) | 1))
                    _ = fi24
                    _ = v25
                }
            }
        }
        if false {
            v1 = 341
        } else {
            v4 = -32768
        }
    } else {
        var v26 float32 = -2.25
        v27 := GenS2{F0: (v1 | 873), F1: 940, F2: 526}
        _ = v26
        _ = v27
    }
    v18 = v18
    for fi28 := 0; fi28 < 7; fi28++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v29 bool = false
        v29 = (1 >= 42)
        _ = fi28
        _ = v29
    }
    v3 = append(v3, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
