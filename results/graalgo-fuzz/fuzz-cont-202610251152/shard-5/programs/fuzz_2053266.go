package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
func genFn0(p0 uint16, p1 int8) uint16 {
    if (true && (p0 == (^(p0)))) {
        fmt.Println(p1)
    } else {
        p0 += p0
        p1 += p1
        if true {
            p0 = p0
            v4 := []int{128}
            p0 = 566
            _ = v4
        }
    }
    v5 := GenS1{F0: 9223372036854775807, F1: -1}
    switch 32767 {
    case 127:
        for fi6 := 0; fi6 < 20; fi6++ {
            var v7 string = (("hello" + ((("" + "foo bar") + (("" + ("\n\t" + "\n\t")) + "")) + (("hello" + ("foo bar" + ("hello" + ("foo bar" + "hello")))) + ("" + ("a" + ""))))) + "")
            _ = fi6
            _ = v7
        }
    default:
        p1 += (p1 * (p1 + p1))
    }
    _ = p0
    _ = p1
    _ = v5
    return p0
}
func genFn1() int64 {
    v8 := []int{2147483647, 46, -779, 114}
    var v9 int32 = 2147483647
    for fi10 := 0; fi10 < 10; fi10++ {
        if len(v8) > 0 {
            v8[0] = (^(fi10))
        }
        _ = fi10
    }
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int64 = -825
    v2 := make(map[string]string)
    var v3 int64 = (v1 % ((0) | 1))
    delete(v2, "a")
    var v11 float32 = 1e10
    v12 := GenS1{F0: v3, F1: 42}
    v11 = v11
    v3 = (v1 * (-(v3)))
    v1 -= -262
    switch 1 {
    case 0:
        v3 = (v1 | 9223372036854775807)
    case 128:
        if ((-2.25 - (v11 / 1e10)) != -2.25) {
            v13 := GenS1{F0: v1, F1: 32767}
            _ = v13
        }
    default:
        v14 := GenS1{F0: v3, F1: 32767}
        _ = v14
    }
    var v15 uint16 = 0
    switch -575 {
    case 522:
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(v1, v15, v11)
            v15 = v15
            switch "hello" {
            case "foo bar":
                v17 := GenS1{F0: genFn1(), F1: -32768}
                _ = v17
            }
            _ = fi16
        }
    case -2147483648:
        v11 = 3.14159
    case 58:
        if true {
            var v18 float32 = v11
            _ = v18
        } else {
            v3 += 1
            switch -9223372036854775808 {
            case -82:
                v15 = genFn0((v15 / (((^(v15))) | 1)), 100)
            }
            var v19 int64 = v1
            _ = v19
        }
    }
    v3 -= 9223372036854775807
    v20 := v2["a"]
    var v21 int32 = 1
    v22 := GenS1{F0: (v3 - 326), F1: -107}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
