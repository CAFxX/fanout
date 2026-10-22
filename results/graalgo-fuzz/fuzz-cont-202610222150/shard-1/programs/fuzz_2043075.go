package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint8) int64 {
    p1 += p0
    switch ("" + "foo bar") {
    case "hello":
        v5 := []int{-969}
        _ = v5
    case "":
        v6 := GenS1{F0: "a", F1: "a"}
        _ = v6
    case "a":
        switch "foo bar" {
        case "a":
            v7 := GenS1{F0: "", F1: ""}
            _ = v7
        default:
            v8 := []int{2147483647, 42, 128}
            _ = v8
        }
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int64, p1 int, p2 uint16) bool {
    v9 := GenS1{F0: ("foo bar" + "\n\t"), F1: "hello"}
    v10 := GenS2{F0: 42}
    var v11 string = ("hello" + "\n\t")
    p2 = (p2 >> ((2) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return false
}

func main() {
    var v1 float32 = 1.5
    var v2 bool = false
    var v3 string = "a"
    var v4 float32 = v1
    v3 = (v3 + "\n\t")
    if true {
        v1 = 0.0
        v2 = v2
    }
    v12 := GenS2{F0: -32768}
    switch ("\n\t" + (v3 + ("hello" + ""))) {
    case "\n\t":
        v13 := []int{32768, 9223372036854775807, 343, 32767}
        _ = v13
    case "a":
        v14 := GenS1{F0: v3, F1: (v3 + (((v3 + v3) + (("\n\t" + (v3 + "foo bar")) + "a")) + "hello"))}
        _ = v14
    }
    fmt.Println(v1, v2)
    var v15 int = 1
    v4 = (1e10 / 1.0)
    if (!(128 <= 42)) {
        v4 -= ((v4 - 1.5) * 3.14159)
        for fi16 := 0; fi16 < 1; fi16++ {
            fmt.Println(v3, fi16, v15)
            for fi17 := 0; fi17 < 12; fi17++ {
                v15 += fi17
                v18 := GenS2{F0: 32767}
                _ = fi17
                _ = v18
            }
            switch (v3 + (v3 + (("hello" + v3) + v3))) {
            case "":
                v19 := []int{}
                _ = v19
            case "foo bar":
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 bool = (!(65535 > 255))
                    _ = fi20
                    _ = v21
                }
            default:
                var v22 int64 = 0
                _ = v22
            }
            _ = fi16
        }
        var v23 string = v3
        _ = v23
    } else {
        for fi24 := 0; fi24 < 5; fi24++ {
            fmt.Println(v4)
            var v25 float64 = (-2.25 * (0.0 / -2.25))
            switch 1 {
            case 128:
                var v26 uint = 572
                _ = v26
            case 32767:
                if false {
                    for fi27 := 0; fi27 < 2; fi27++ {
                        fmt.Println(fi24, v15)
                        _ = fi27
                    }
                    var v28 float32 = float32((1.5 - (v25 / (float64(((-2.25 / v1) / v4)) / v25))))
                    for fi29 := 0; fi29 < 14; fi29++ {
                        v2 = (0.1 > float32((((((v25 / (v25 - v25)) / float64((v1 - v28))) * ((v25 + (v25 - (((v25 * v25) + (((((((float64(float32(v25)) - v25) + v25) + v25) + v25) * (1e10 / 1.0)) / (v25 / v25)) + (0.1 + v25))) + (((((v25 + ((float64(3.14159) + v25) * v25)) + v25) / (v25 + 1.5)) + 1.5) / v25)))) - v25)) - (float64(-2.25) / (float64(v1) - 1e10))) * 1e10)))
                        v30 := []int{(fi29 & fi29), 42, (fi29 % ((0) | 1)), 32768}
                        var v31 int32 = 42
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = v28
                }
            case 1:
                var v32 uint32 = 2147483647
                _ = v32
            }
            _ = fi24
            _ = v25
        }
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 float64 = 3.14159
            _ = fi33
            _ = v34
        }
    }
    fmt.Println(v3, v4, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
