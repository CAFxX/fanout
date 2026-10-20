package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 int32; F2 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 uint8; F2 uint }
func genFn0(p0 uint8, p1 int8) int64 {
    v4 := GenS1{F0: 1, F1: "foo bar", F2: 255}
    p1 += (p1 & p1)
    if true {
        p0 += (p0 * (p0 / ((127) | 1)))
    }
    var v5 uint32 = 0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{127, -9223372036854775808, -9223372036854775808}
    v3 := make(map[uint32]string)
    var v6 uint32 = 2147483648
    var v7 int8 = -1
    delete(v1, ("\n\t" + ""))
    var v8 uint = 865
    var v9 int32 = -1
    var v10 bool = (!false)
    v11 := GenS2{F0: 127}
    v10 = (true && v10)
    for fi12 := 0; fi12 < 4; fi12++ {
        if false {
            v13 := GenS3{F0: 111, F1: v9, F2: v7}
            for fi14 := 0; fi14 < 4; fi14++ {
                if v10 {
                    var v15 uint8 = 186
                    _ = v15
                }
                _ = fi14
            }
            v7 = (v7 + (v7 << ((v8) & 63)))
            _ = v13
        } else {
            v16 := GenS4{F0: 2}
            if len(v2) > 0 {
                v2[0] = fi12
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS3{F0: 0, F1: -936, F2: (v7 % (((v7 & (v7 * (v7 ^ v7)))) | 1))}
                v19 := GenS5{F0: 9223372036854775807}
                fmt.Println(fi12)
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = v16
        }
        var v20 int8 = v7
        fmt.Println(v10, v8, v6)
        _ = fi12
        _ = v20
    }
    var v21 int16 = 0
    if v10 {
        switch 32767 {
        case 0:
            v7 = v7
        }
        switch "hello" {
        case "a":
            v1["foo bar"] = (v6 + (^(v6)))
        case "foo bar":
            v3[(^(v6))] = "hello"
        }
    } else {
        v22 := GenS6{F0: 42, F1: 40, F2: (v8 | 1)}
        if true {
            v8 = v8
        }
        _ = v22
    }
    if (!(2 != 0)) {
        v8 -= v8
        v3[v6] = "a"
    } else {
        switch "" {
        case "foo bar":
            if false {
                v23 := GenS2{F0: -1}
                var v24 int = -9223372036854775808
                _ = v23
                _ = v24
            }
        case "":
            var v25 bool = v10
            _ = v25
        }
        var v26 float64 = 1e10
        fmt.Println(v26, v8)
        _ = v26
    }
    v6 += 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
