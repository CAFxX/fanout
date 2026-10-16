package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 bool; F2 int }
func genFn0(p0 int32) int8 {
    var v7 int8 = 42
    p0 -= p0
    v7 += v7
    for fi8 := 0; fi8 < 5; fi8++ {
        p0 = p0
        var v9 uint16 = 354
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v7
    return v7
}

func main() {
    var v1 uint16 = 32767
    var v2 uint16 = v1
    var v3 int64 = -9223372036854775808
    var v4 string = "foo bar"
    v5 := []string{(((v4 + (v4 + ("\n\t" + (v4 + ((v4 + v4) + v4))))) + "foo bar") + v4), v4}
    var v6 int16 = 1000
    if (!(375 >= 42)) {
        v10 := GenS1{F0: 1}
        _ = len(v5)
        var v11 float64 = -2.25
        _ = v10
        _ = v11
    }
    switch ("foo bar" + v4) {
    case "hello":
        v12 := GenS1{F0: 146}
        _ = v12
    case "":
        switch "foo bar" {
        case "":
            var v13 int64 = v3
            _ = v13
        }
    case "\n\t":
        v2 = (v1 % ((42) | 1))
    }
    switch 4294967295 {
    case -343:
        for fi14 := 0; fi14 < 11; fi14++ {
            for fi15 := 0; fi15 < 8; fi15++ {
                var v16 float64 = 1e10
                var v17 float64 = v16
                if true {
                    if true {
                        v18 := GenS1{F0: 1}
                        _ = v18
                    }
                } else {
                    var v19 float64 = (v16 + (1e10 - v17))
                    fmt.Println(v3, fi15, fi14)
                    fmt.Println(v2, v19)
                    _ = v19
                }
                if true {
                    v20 := GenS1{F0: -1}
                    v1 = v1
                    _ = v20
                }
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi14
        }
    }
    v4 = ("hello" + "a")
    v3 = 893
    for fi21 := 0; fi21 < 15; fi21++ {
        switch "a" {
        case "":
            fmt.Println(v6, v4, v4)
        case "foo bar":
            var v22 uint16 = v1
            _ = v22
        }
        v23 := GenS2{F0: 1, F1: (805 >= 669), F2: fi21}
        v24 := GenS1{F0: 932}
        _ = fi21
        _ = v23
        _ = v24
    }
    if ((0.0 / 0.1) > -2.25) {
        v25 := GenS2{F0: -2147483648, F1: false, F2: -483}
        _ = v25
    } else {
        var v26 string = v4
        _ = v26
    }
    v5 = append(v5, v4)
    for fi27 := 0; fi27 < 8; fi27++ {
        v28 := GenS2{F0: 0, F1: (-2.25 > float32((1.5 - -2.25))), F2: 1}
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
