package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint8) int32 {
    if true {
        fmt.Println(p0)
    }
    var v4 uint32 = 42
    v5 := []int{42, 9223372036854775807}
    _ = p0
    _ = v4
    _ = v5
    return 499
}
func genFn1() string {
    var v6 int = 128
    if ((!(false || false)) && (("a" + "hello") > "")) {
        if false {
            var v7 int32 = -815
            if false {
                v7 += 2147483647
                for fi8 := 0; fi8 < 7; fi8++ {
                    switch "hello" {
                    case "\n\t":
                        switch "a" {
                        case "a":
                            v9 := []int{}
                            _ = v9
                        case "\n\t":
                            if (!(!(true || (-811 != 365)))) {
                                v7 -= v7
                            }
                        case "foo bar":
                            v6 -= (fi8 / (((fi8 / ((-2147483648) | 1))) | 1))
                        default:
                            v6 = fi8
                        }
                    }
                    var v10 bool = (0 == 0)
                    _ = fi8
                    _ = v10
                }
            } else {
                var v11 uint32 = 2
                var v12 uint32 = (^(v11))
                _ = v11
                _ = v12
            }
            _ = v7
        } else {
            v13 := GenS1{F0: false}
            fmt.Println(v6)
            _ = v13
        }
        switch 288 {
        case 17:
            v14 := []int{v6}
            _ = v14
        case 9223372036854775807:
            var v15 uint8 = 42
            _ = v15
        }
    } else {
        v16 := GenS1{F0: (612 >= -2147483648)}
        v6 -= (v6 / ((v6) | 1))
        var v17 uint32 = 2147483648
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v6 = (fi18 + (v6 ^ fi18))
        if false {
            v19 := []int{fi18, (fi18 | v6)}
            _ = v19
        } else {
            v20 := GenS2{F0: 135}
            _ = v20
        }
        _ = fi18
    }
    fmt.Println(v6)
    _ = v6
    return ((("hello" + "\n\t") + (("" + ((("hello" + ("\n\t" + "\n\t")) + ("" + "foo bar")) + "hello")) + "a")) + "foo bar")
}

func main() {
    var v1 uint64 = 330
    var v2 uint8 = 255
    var v3 int = 128
    var v21 uint32 = 0
    var v22 float32 = 1e10
    v22 += v22
    for fi23 := 0; fi23 < 15; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            v22 -= float32(0.1)
            fmt.Println(v1, v22, fi23)
            _ = fi24
        }
        _ = fi23
    }
    fmt.Println(v21, v1)
    switch (-(v3)) {
    case 32767:
        var v25 uint16 = 2
        _ = v25
    }
    v26 := GenS3{F0: 2, F1: int32(v1), F2: (v21 + v21)}
    switch "hello" {
    case "foo bar":
        v1 += v1
    case "a":
        v27 := []int{}
        _ = v27
    case "\n\t":
        v28 := GenS4{F0: 0}
        _ = v28
    }
    switch ("" + "foo bar") {
    case "\n\t":
        v29 := []int{}
        _ = v29
    case "hello":
        v30 := GenS2{F0: v2}
        _ = v30
    case "foo bar":
        var v31 int32 = 0
        _ = v31
    }
    v3 = (-(v3))
    var v32 float32 = v22
    v33 := GenS4{F0: 29}
    var v34 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
