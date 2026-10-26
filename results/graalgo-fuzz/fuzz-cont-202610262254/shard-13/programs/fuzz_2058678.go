package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int32 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint32 = 42
    var v2 float32 = 1.5
    var v3 int64 = 0
    var v4 float64 = 0.0
    for fi5 := 0; fi5 < 7; fi5++ {
        if (0.1 > 0.1) {
            v1 = (v1 * v1)
            v6 := []int{fi5, fi5, fi5, 9223372036854775807}
            fmt.Println(v3)
            _ = v6
        }
        v7 := GenS1{F0: -1, F1: 905, F2: 2147483647}
        _ = fi5
        _ = v7
    }
    if (!false) {
        v3 -= (v3 / (((-(v3))) | 1))
        for fi8 := 0; fi8 < 16; fi8++ {
            v4 -= ((1.5 + 3.14159) / 1.5)
            var v9 uint = 326
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: -512, F1: 716, F2: 673}
        _ = v10
    } else {
        v11 := []int{int(v3)}
        var v12 int32 = -585
        var v13 bool = false
        _ = v11
        _ = v12
        _ = v13
    }
    switch (((("" + "\n\t") + "\n\t") + "hello") + "") {
    case "hello":
        var v14 string = ""
        _ = v14
    case "a":
        v15 := GenS1{F0: 2147483647, F1: 638, F2: 0}
        _ = v15
    case "\n\t":
        v1 = (v1 & (v1 ^ (^(v1))))
    default:
        fmt.Println(v1)
    }
    v16 := []int{}
    if false {
        fmt.Println(v1, v4)
    }
    v4 -= ((v4 * v4) * (v4 + (v4 / v4)))
    var v17 float64 = v4
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 uint16 = 716
        fmt.Println(v19, fi18)
        if len(v16) > 0 {
            v16[0] = (fi18 | (fi18 | fi18))
        }
        _ = fi18
        _ = v19
    }
    var v20 int16 = -842
    var v21 float64 = v4
    _ = len(v16)
    fmt.Println(v2, v1, v1)
    var v22 int = 2147483647
    if false {
        fmt.Println(v4)
    }
    if false {
        if (163 != 0) {
            for fi23 := 0; fi23 < 20; fi23++ {
                v22 -= v22
                v21 = ((3.14159 * ((v17 + v21) - v4)) * v4)
                _ = fi23
            }
            var v24 uint32 = 2
            v25 := GenS1{F0: 203, F1: 163, F2: 214}
            _ = v24
            _ = v25
        } else {
            var v26 string = ("\n\t" + ("foo bar" + "foo bar"))
            if true {
                v20 = v20
                switch (v22 << ((4294967295) & 63)) {
                case 4294967295:
                    var v27 float64 = 3.14159
                    _ = v27
                default:
                    var v28 uint8 = 42
                    _ = v28
                }
            }
            _ = v26
        }
        if (v22 == -9223372036854775808) {
            var v29 int = 0
            _ = v29
        } else {
            switch "\n\t" {
            case "":
                v20 = (^(v20))
            default:
                v30 := GenS2{F0: 1}
                _ = v30
            }
            v31 := GenS2{F0: 605}
            _ = v31
        }
    }
    switch v22 {
    case 4294967295:
        for fi32 := 0; fi32 < 5; fi32++ {
            for fi33 := 0; fi33 < 4; fi33++ {
                fmt.Println(v17)
                v1 = v1
                _ = fi33
            }
            switch "\n\t" {
            case "\n\t":
                var v34 uint32 = v1
                _ = v34
            default:
                v3 += (v3 + v3)
            }
            _ = fi32
        }
    case -2147483648:
        for fi35 := 0; fi35 < 6; fi35++ {
            v17 = 0.0
            _ = fi35
        }
    }
    v36 := GenS1{F0: 1, F1: 18446744073709551615, F2: -2147483648}
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v36
    _ = v4
}
