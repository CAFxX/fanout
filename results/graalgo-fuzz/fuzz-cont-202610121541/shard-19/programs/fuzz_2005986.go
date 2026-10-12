package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int16; F1 string; F2 uint32 }
func genFn0(p0 uint32, p1 uint32, p2 uint32) int8 {
    v7 := []int{-898, 32767, 32768, 127}
    fmt.Println(p2, p0)
    p2 += p2
    var v8 int = 42
    v7 = append(v7, (v8 ^ v8))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 100
}
func genFn1(p0 bool) bool {
    v9 := []int{12, 42}
    v10 := GenS1{F0: 65535, F1: 32767}
    _ = p0
    _ = v10
    _ = v9
    return (413 < 42)
}

func main() {
    var v1 uint = 4294967295
    var v2 int16 = 0
    var v3 float64 = ((1.5 / float64(0.1)) - 0.1)
    var v4 int16 = v2
    v5 := []int{32767, -368}
    v6 := make(map[string]string)
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: 4294967295, F1: 42}
        v2 += int16(v1)
        _ = fi11
        _ = v12
    }
    v1 = (v1 - (^(v1)))
    var v13 uint = v1
    for fi14 := 0; fi14 < 2; fi14++ {
        _ = len(v5)
        v1 = (v1 % (((v1 % ((v1) | 1))) | 1))
        _ = fi14
    }
    v1 += v13
    if true {
        if (((0.0 / 0.1) * (-2.25 * 3.14159)) >= 1e10) {
            var v15 float64 = (0.0 / ((v3 / float64(1.5)) / v3))
            _ = v15
        } else {
            v16 := GenS2{F0: (v4 - v4), F1: "hello", F2: 2147483648}
            v3 += 3.14159
            _ = v16
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            fmt.Println(fi17, fi17, v13)
            _ = fi17
        }
        switch (("\n\t" + "hello") + (("a" + ("hello" + "")) + ((("\n\t" + "a") + ((("hello" + ("a" + ("foo bar" + ("\n\t" + "hello")))) + "foo bar") + "a")) + ""))) {
        case "":
            for fi18 := 0; fi18 < 8; fi18++ {
                v3 -= -2.25
                v1 += v1
                _ = len(v6)
                _ = fi18
            }
        case "\n\t":
            var v19 uint16 = 32767
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        v3 = float64(1.5)
        _ = fi20
    }
    v1 -= (v13 ^ 0)
    var v21 uint = (v13 * v1)
    v22 := GenS2{F0: v4, F1: ("\n\t" + ("a" + "foo bar")), F2: 42}
    v4 = 300
    var v23 uint16 = 884
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
