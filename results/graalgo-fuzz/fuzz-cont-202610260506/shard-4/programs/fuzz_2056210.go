package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int8; F2 bool }
type GenS4 struct { F0 bool; F1 int32; F2 uint64 }

func main() {
    var v1 uint32 = 2
    var v2 int = 32768
    var v3 float64 = 0.0
    var v4 uint64 = 0
    var v5 int = v2
    v4 += 687
    var v6 string = "hello"
    if false {
        v7 := []int{(v5 - v5), v5}
        v8 := GenS1{F0: 2147483648}
        _ = v7
        _ = v8
    } else {
        v2 = (^(v2))
    }
    v1 = v1
    var v9 bool = (!true)
    if (-5 < -112) {
        v10 := []int{v2, v5, (v2 - (^(v2)))}
        if false {
            v2 += v5
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        } else {
            fmt.Println(v4)
        }
        v2 -= 127
        _ = v10
    }
    switch ("foo bar" + v6) {
    case "\n\t":
        for fi11 := 0; fi11 < 10; fi11++ {
            if v9 {
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 string = (((((("a" + (v6 + v6)) + "\n\t") + v6) + v6) + ((v6 + ((v6 + v6) + ("hello" + "\n\t"))) + (("\n\t" + v6) + (((((("" + (v6 + v6)) + "") + v6) + ("" + ("hello" + "hello"))) + ((("" + v6) + v6) + "\n\t")) + v6)))) + "a")
                    _ = fi12
                    _ = v13
                }
                v5 -= (v5 & (fi11 - v2))
            }
            v6 = v6
            _ = fi11
        }
    case "":
        for fi14 := 0; fi14 < 18; fi14++ {
            v2 = (v5 / ((v2) | 1))
            fmt.Println(v4, v5)
            var v15 int8 = 1
            _ = fi14
            _ = v15
        }
    case "hello":
        for fi16 := 0; fi16 < 3; fi16++ {
            if v9 {
                v1 = 4294967295
                switch v6 {
                case "foo bar":
                    var v17 uint8 = 95
                    _ = v17
                case "":
                    var v18 int16 = 562
                    _ = v18
                }
            }
            v3 = v3
            _ = fi16
        }
    default:
        if true {
            var v19 int16 = 0
            v4 += v4
            _ = v19
        } else {
            for fi20 := 0; fi20 < 15; fi20++ {
                v9 = (v9 || true)
                _ = fi20
            }
        }
    }
    var v21 int16 = 858
    if (("a" + "a") != v6) {
        v22 := GenS2{F0: 255}
        v3 += 1e10
        var v23 int8 = 0
        _ = v22
        _ = v23
    } else {
        v24 := GenS1{F0: v1}
        var v25 uint = 0
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        var v27 uint64 = v4
        _ = fi26
        _ = v27
    }
    v28 := GenS3{F0: 142, F1: -51, F2: ("\n\t" >= (v6 + "foo bar"))}
    switch v2 {
    case -93:
        v29 := GenS3{F0: 172, F1: 100, F2: v9}
        _ = v29
    case 127:
        fmt.Println(v5, v9)
    case 117:
        v3 = v3
    }
    switch v5 {
    case -342:
        v30 := GenS2{F0: 1}
        _ = v30
    }
    fmt.Println(v21, v6, v2)
    v31 := GenS4{F0: false, F1: 42, F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
