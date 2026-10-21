package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int64 }
type GenS2 struct { F0 int64; F1 bool; F2 int16 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint64 = 1
    var v2 float64 = 0.1
    var v3 float64 = (v2 / v2)
    v4 := GenS1{F0: "a", F1: 2, F2: -9223372036854775808}
    if true {
        var v5 uint64 = v1
        var v6 int16 = -32768
        var v7 bool = (42 > 1)
        _ = v5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := []int{(fi8 & fi8), (fi8 & fi8), fi8, (fi8 & (fi8 & fi8))}
        if (int(v1) < fi8) {
            v10 := GenS2{F0: 1, F1: (v1 > (v1 / (((^(v1))) | 1))), F2: 1}
            v1 = v1
            if len(v9) > 0 {
                v9[0] = fi8
            }
            _ = v10
        } else {
            v11 := GenS3{F0: 94, F1: 846}
            v12 := GenS1{F0: ("\n\t" + "hello"), F1: 0, F2: -375}
            v13 := GenS4{F0: -9223372036854775808}
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi8
        _ = v9
    }
    var v14 int = -9223372036854775808
    v15 := GenS5{F0: 0}
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 uint16 = 708
        _ = fi16
        _ = v17
    }
    v1 = v1
    for fi18 := 0; fi18 < 9; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := []int{-1}
            var v21 float32 = 1.5
            if true {
                if len(v20) > 0 {
                    v20[0] = (fi19 * 927)
                }
                switch ((("\n\t" + "\n\t") + "") + "foo bar") {
                case "\n\t":
                    fmt.Println(fi18, v21, fi18)
                default:
                    fmt.Println(v21, v14, v21)
                }
            }
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            v3 = 1e10
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    var v22 string = "\n\t"
    var v23 float64 = (v3 / -2.25)
    v2 = v23
    v24 := []int{v14, (v14 << ((4294967295) & 63)), (v14 / ((v14) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
