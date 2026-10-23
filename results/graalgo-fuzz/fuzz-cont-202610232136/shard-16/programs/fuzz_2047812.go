package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
func genFn0(p0 int16, p1 int16) string {
    v4 := GenS1{F0: 1, F1: "hello"}
    v5 := GenS1{F0: -2147483648, F1: ("" + (("\n\t" + ("foo bar" + "")) + ("hello" + ("foo bar" + ("hello" + "foo bar")))))}
    p1 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (("\n\t" + "a") + ("a" + (("a" + "") + "a")))
}
func genFn1() int16 {
    if (!true) {
        var v6 uint8 = 2
        _ = v6
    }
    var v7 string = "\n\t"
    _ = v7
    return -308
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[string]string)
    v3 := []int{-544, -9223372036854775808, -9223372036854775808, 2147483647}
    var v8 uint16 = 219
    v8 += v8
    fmt.Println(v8)
    v1 -= (v1 / ((v1) | 1))
    switch "" {
    case "hello":
        v9 := GenS1{F0: 160, F1: "foo bar"}
        _ = v9
    }
    delete(v2, "\n\t")
    var v10 uint8 = 1
    if (0 != 0) {
        v2["a"] = "a"
        for fi11 := 0; fi11 < 10; fi11++ {
            v1 -= v1
            v12 := v2[("hello" + (("" + ("" + "hello")) + "foo bar"))]
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi11
            _ = v12
        }
    } else {
        v1 = v1
        if true {
            v13 := GenS1{F0: int32(v1), F1: ("\n\t" + "foo bar")}
            _ = v13
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
}
