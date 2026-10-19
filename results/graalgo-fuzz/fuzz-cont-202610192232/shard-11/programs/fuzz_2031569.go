package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 bool, p1 uint16, p2 int8) int32 {
    p2 -= (p2 + p2)
    var v7 uint64 = 198
    for fi8 := 0; fi8 < 18; fi8++ {
        if p0 {
            for fi9 := 0; fi9 < 11; fi9++ {
                fmt.Println(p0)
                v10 := GenS1{F0: "a", F1: 40, F2: 128}
                p1 = p1
                _ = fi9
                _ = v10
            }
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1(p0 uint8) int32 {
    var v11 uint16 = 32767
    if false {
        fmt.Println(p0)
        p0 = 255
        v12 := GenS1{F0: "", F1: 65535, F2: 42}
        _ = v12
    }
    var v13 int16 = 170
    switch 0 {
    case 693:
        v14 := GenS1{F0: ("hello" + "foo bar"), F1: 2147483647, F2: uint8(v11)}
        _ = v14
    case -1:
        if (!(p0 >= 8)) {
            fmt.Println(p0, p0)
            v15 := []int{-2147483648, 32767, 0}
            v15 = append(v15, 127)
            _ = v15
        }
    default:
        v13 -= v13
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 int = -427
        fmt.Println(p0, p0, v13)
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = v11
    _ = v13
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 float32 = 1e10
    var v3 float64 = 3.14159
    var v4 uint16 = 32767
    var v5 int = -2147483648
    var v6 int16 = 32767
    switch "foo bar" {
    case "hello":
        var v18 int = v5
        _ = v18
    case "":
        v6 = (v6 * v6)
    case "a":
        var v19 bool = (358 < 4294967295)
        _ = v19
    default:
        for fi20 := 0; fi20 < 20; fi20++ {
            v4 += (^(v4))
            _ = fi20
        }
    }
    switch "\n\t" {
    case "a":
        delete(v1, ("hello" + "\n\t"))
    case "":
        v5 -= v5
    case "hello":
        var v21 float64 = float64(-2.25)
        _ = v21
    default:
        var v22 int32 = -2147483648
        _ = v22
    }
    switch "\n\t" {
    case "hello":
        v5 += (v5 + (v5 * 948))
    default:
        v23 := GenS1{F0: "", F1: 623, F2: 255}
        _ = v23
    }
    if ((1 < 65535) || (false || true)) {
        var v24 uint64 = 18446744073709551615
        v6 -= (v6 - (v6 - (v6 % ((v6) | 1))))
        v25 := GenS2{F0: -1}
        _ = v24
        _ = v25
    }
    if true {
        v26 := v1["foo bar"]
        _ = v26
    }
    v27 := v1[("hello" + ("" + ("\n\t" + "")))]
    fmt.Println(v3, v2, v3)
    for fi28 := 0; fi28 < 5; fi28++ {
        if false {
            switch ("hello" + "\n\t") {
            case "":
                fmt.Println(v3, v3)
            case "\n\t":
                v1["\n\t"] = ("foo bar" < "")
            case "hello":
                var v29 uint = 2
                _ = v29
            default:
                var v30 int64 = 9223372036854775807
                _ = v30
            }
        }
        _ = fi28
    }
    if v27 {
        v31 := GenS2{F0: 0}
        v32 := GenS1{F0: ("foo bar" + ((("a" + "hello") + "") + ("a" + ""))), F1: 0, F2: 191}
        if ((genFn1(1) <= 1) && (((float32((v3 - v3)) + ((v2 * v2) * -2.25)) + 3.14159) < 3.14159)) {
            _ = len(v1)
        } else {
            switch "hello" {
            case "":
                if ((!(4294967295 < 806)) || false) {
                    switch (v5 * -966) {
                    case -9223372036854775808:
                        var v33 int = v5
                        _ = v33
                    }
                    var v34 int16 = 736
                    if (((((v3 + (1.5 * (((v3 + (0.1 / ((v3 / (0.1 - 1e10)) - (((v3 + v3) - v3) - (v3 * (0.0 / 1.5)))))) * float64(3.14159)) / float64((v2 - v2))))) / (v3 * 1.5)) + ((v3 + (((v3 + (1.5 - ((v3 / -2.25) / (v3 + 0.1)))) + v3) - float64((v2 * ((float32(v3) * float32(1.5)) * v2))))) - v3)) * ((v3 * v3) - v3)) < ((v3 * v3) / v3)) {
                        v35 := v1[((((("hello" + ("hello" + "\n\t")) + "foo bar") + "") + "\n\t") + "\n\t")]
                        var v36 uint8 = 128
                        _ = v35
                        _ = v36
                    }
                    _ = v34
                }
            }
            var v37 float32 = float32((3.14159 - ((1e10 - -2.25) + (v3 - v3))))
            _ = v37
        }
        _ = v31
        _ = v32
    }
    v2 += (0.0 - v2)
    var v38 float64 = (1.5 - float64((v2 - (v2 + v2))))
    v38 += v38
    v4 = (v4 * v4)
    v2 = ((v2 - 0.1) - v2)
    var v39 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
