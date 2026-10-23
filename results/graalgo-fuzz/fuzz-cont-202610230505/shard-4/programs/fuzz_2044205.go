package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
func genFn0(p0 int32) int8 {
    v5 := []int{42}
    var v6 bool = (0 >= 502)
    switch 0 {
    case -9223372036854775808:
        v6 = true
    }
    _ = p0
    _ = v5
    _ = v6
    return 100
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]string)
    var v3 string = "\n\t"
    var v4 int8 = 100
    v3 = ""
    for fi7 := 0; fi7 < 20; fi7++ {
        switch "foo bar" {
        case "foo bar":
            switch ("hello" + "foo bar") {
            case "hello":
                if (false || (3.14159 > (0.1 / 1.0))) {
                    v3 = ((("\n\t" + ("\n\t" + "a")) + (((v3 + (v3 + (v3 + v3))) + (v3 + v3)) + "\n\t")) + "hello")
                } else {
                    delete(v2, v3)
                    var v8 int32 = -2147483648
                    _ = v8
                }
            case "foo bar":
                v3 = (v3 + "a")
            default:
                fmt.Println(v4, v1, v3)
            }
        case "":
            var v9 bool = (!(4294967295 < 437))
            _ = v9
        case "a":
            if (-347 > -375) {
                v1 += (^(v1))
            } else {
                for fi10 := 0; fi10 < 15; fi10++ {
                    v11 := []int{fi7, (fi7 >> ((255) & 63))}
                    _ = fi10
                    _ = v11
                }
            }
        }
        _ = fi7
    }
    if false {
        fmt.Println(v4)
        fmt.Println(v4, v1, v4)
        var v12 uint16 = v1
        _ = v12
    } else {
        fmt.Println(v3, v3, v1)
        var v13 uint = 65535
        v14 := GenS1{F0: 127, F1: true}
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v1 -= v1
        _ = fi15
    }
    var v16 uint32 = 2147483648
    if true {
        if ((false || false) || (false || ((v4 & (v4 | genFn0(0))) >= 0))) {
            v3 = v3
            var v17 string = ("foo bar" + (v3 + v3))
            var v18 int = 42
            _ = v17
            _ = v18
        }
    }
    _ = len(v2)
    var v19 float64 = -2.25
    var v20 int32 = -2147483648
    v19 += (3.14159 - v19)
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 int16 = -1
        _ = fi21
        _ = v22
    }
    var v23 int = -107
    for fi24 := 0; fi24 < 18; fi24++ {
        v23 -= (fi24 ^ (fi24 & v23))
        v25 := GenS1{F0: int8(v16), F1: (false || false)}
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        for fi27 := 0; fi27 < 7; fi27++ {
            v16 += v16
            var v28 float32 = ((3.14159 * 3.14159) + ((1e10 + 3.14159) / 3.14159))
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
}
