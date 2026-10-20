package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 string }
type GenS2 struct { F0 string; F1 uint64; F2 int32 }
func genFn0(p0 uint8, p1 int16, p2 uint) uint64 {
    if (!((42 != 4294967295) && (!true))) {
        p1 = (^(p1))
    }
    p2 = (p2 >> ((p2) & 63))
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 52
}

func main() {
    var v1 int8 = 42
    var v2 uint = 2
    var v3 bool = true
    switch -957 {
    case -909:
        for fi4 := 0; fi4 < 12; fi4++ {
            v3 = v3
            switch "" {
            case "foo bar":
                switch "a" {
                case "foo bar":
                    v5 := []int{(fi4 + fi4), (fi4 + 2147483647)}
                    _ = v5
                default:
                    if v3 {
                        var v6 uint8 = 128
                        _ = v6
                    } else {
                        v3 = v3
                        var v7 int64 = 198
                        _ = v7
                    }
                }
            default:
                fmt.Println(v1)
            }
            _ = fi4
        }
    default:
        if ("foo bar" <= "hello") {
            v8 := []int{-1, 32768, -2147483648, 32768}
            _ = v8
        }
    }
    var v9 int = 2147483647
    v3 = true
    switch (("hello" + "a") + "\n\t") {
    case "a":
        if false {
            var v10 float32 = -2.25
            _ = v10
        }
    case "\n\t":
        var v11 string = (("hello" + ("\n\t" + "")) + "")
        _ = v11
    }
    var v12 uint = 367
    v13 := GenS1{F0: 2, F1: 1, F2: "foo bar"}
    switch (("a" + "a") + ("\n\t" + (("foo bar" + ("foo bar" + ("\n\t" + "\n\t"))) + "a"))) {
    case "foo bar":
        v14 := GenS2{F0: ("a" + "hello"), F1: 0, F2: 0}
        _ = v14
    case "a":
        var v15 string = "\n\t"
        _ = v15
    case "":
        var v16 float64 = (float64(1e10) * float64((0.0 - ((-2.25 / (0.1 * 3.14159)) / -2.25))))
        _ = v16
    }
    var v17 bool = false
    var v18 float32 = 0.0
    v1 += (v1 % ((v1) | 1))
    var v19 uint = (v2 & (v12 | (^(v2))))
    for fi20 := 0; fi20 < 5; fi20++ {
        v17 = ((!false) && ((-2.25 / (-2.25 - (0.1 / 0.1))) <= 1.5))
        v2 += v19
        v19 -= (v19 - (v2 << ((868) & 63)))
        _ = fi20
    }
    if v3 {
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 bool = ((^(fi21)) > (-(fi21)))
            v12 -= 2
            _ = fi21
            _ = v22
        }
        v1 -= (v1 * -50)
        var v23 int16 = -32768
        _ = v23
    }
    fmt.Println(v17, v12)
    v3 = (!((v17 && false) || false))
    var v24 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v9
}
