package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 bool; F2 uint64 }
func genFn0(p0 string) int16 {
    var v5 float32 = ((1.5 + ((1e10 + float32(((((3.14159 + (0.1 * 0.1)) - (1.5 / 3.14159)) / 1.5) / ((1.5 / 3.14159) * 0.1)))) + 0.1)) - 1.5)
    v5 = v5
    var v6 bool = true
    if (true && (!v6)) {
        fmt.Println(v6)
        if false {
            v7 := []int{}
            _ = v7
        }
    } else {
        var v8 int16 = 0
        v9 := []int{-9223372036854775808, 2147483647, -139}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    return -32768
}
func genFn1(p0 int, p1 string) string {
    p0 = (p0 | p0)
    fmt.Println(p1, p1)
    var v10 bool = (-937 > genFn0("\n\t"))
    _ = p0
    _ = p1
    _ = v10
    return p1
}

func main() {
    var v1 int = -669
    var v2 uint16 = 2
    var v3 int64 = -540
    var v4 string = "\n\t"
    if true {
        switch v4 {
        case "a":
            v11 := GenS1{F0: "hello"}
            _ = v11
        case "":
            v12 := []int{}
            _ = v12
        }
        v13 := []int{v1, (v1 >> ((776) & 63)), -9223372036854775808}
        _ = v13
    }
    if true {
        if false {
            var v14 uint = 52
            v2 -= (v2 + 273)
            _ = v14
        } else {
            v4 = ("\n\t" + (v4 + ((("a" + ("\n\t" + v4)) + v4) + v4)))
            for fi15 := 0; fi15 < 9; fi15++ {
                for fi16 := 0; fi16 < 18; fi16++ {
                    if (783 >= 0) {
                        fmt.Println(fi16, v3)
                        fmt.Println(v2, fi15, fi16)
                        v17 := GenS1{F0: ""}
                        _ = v17
                    } else {
                        var v18 int = -623
                        var v19 int16 = 42
                        var v20 int16 = v19
                        fmt.Println(fi15, v18, v2)
                        v4 = ((v4 + v4) + (v4 + "foo bar"))
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    v21 := []int{(fi16 << ((18446744073709551615) & 63)), 790, fi15}
                    _ = fi16
                    _ = v21
                }
                var v22 float32 = 1.5
                var v23 int8 = -1
                _ = fi15
                _ = v22
                _ = v23
            }
        }
    } else {
        var v24 uint16 = v2
        v25 := []int{v1, 0}
        if len(v25) > 0 {
            v25[0] = (v1 % ((v1) | 1))
        }
        _ = v24
        _ = v25
    }
    fmt.Println(v4, v3, v1)
    v26 := GenS1{F0: ""}
    v27 := GenS2{F0: 9223372036854775807, F1: false, F2: 667}
    var v28 uint32 = 337
    fmt.Println(v4, v3, v4)
    v3 -= -68
    switch v4 {
    case "hello":
        v29 := GenS2{F0: (v3 * (v3 - (^(v3)))), F1: true, F2: 0}
        _ = v29
    case "foo bar":
        var v30 uint8 = 1
        _ = v30
    case "\n\t":
        v1 = int(v28)
    default:
        fmt.Println(v28, v28)
    }
    fmt.Println(v3, v4)
    switch v1 {
    case 0:
        fmt.Println(v28, v4)
    }
    var v31 int32 = -1
    var v32 int8 = -28
    fmt.Println(v32)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
