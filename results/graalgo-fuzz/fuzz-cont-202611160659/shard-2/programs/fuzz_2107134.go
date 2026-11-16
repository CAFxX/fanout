package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
func genFn0() uint {
    var v6 int16 = 1000
    for fi7 := 0; fi7 < 12; fi7++ {
        switch "\n\t" {
        case "hello":
            v8 := []int{(^(fi7)), fi7, 42, 436}
            _ = v8
        case "a":
            v6 += (v6 / (((-(v6))) | 1))
        default:
            v6 += (v6 & -756)
        }
        v6 = v6
        _ = fi7
    }
    switch (("hello" + "foo bar") + ("hello" + "")) {
    case "\n\t":
        v6 += v6
    default:
        v9 := []int{627}
        _ = v9
    }
    if ("" == "\n\t") {
        v10 := GenS1{F0: (!(255 == 175)), F1: 32767}
        _ = v10
    }
    v6 = (v6 & v6)
    _ = v6
    return 0
}
func genFn1() string {
    var v11 bool = false
    fmt.Println(v11)
    var v12 uint64 = 313
    _ = v11
    _ = v12
    return ""
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = 1.5
    v3 := make(map[string]bool)
    v4 := []string{"a", "hello", ""}
    v5 := make(map[string]int)
    v4 = append(v4, ("foo bar" + "foo bar"))
    _ = len(v4)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        if (42 == 128) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v2 -= (3.14159 - (v2 + 1e10))
            v2 = (3.14159 * v2)
        } else {
            for fi14 := 0; fi14 < 16; fi14++ {
                v3["a"] = false
                switch "" {
                case "":
                    if true {
                        v2 -= (((0.0 * v2) / v2) - ((float64(((float32(0.0) / ((1.5 + ((0.0 / 1e10) * 1.5)) / float32((float64((1e10 * -2.25)) + v2)))) - 1e10)) - v2) * (v2 + (v2 * v2))))
                    }
                case "hello":
                    v2 = -2.25
                case "foo bar":
                    fmt.Println(fi13, fi13)
                }
                var v15 string = "a"
                if (2147483647 != 1) {
                    if len(v4) > 0 {
                        v4[0] = (v15 + (((v15 + v15) + "foo bar") + ((((("" + (v15 + "a")) + v15) + v15) + (v15 + "foo bar")) + (v15 + "hello"))))
                    }
                }
                switch (v15 + v15) {
                case "foo bar":
                    if (false || (626 == 1)) {
                        _ = len(v4)
                    }
                case "":
                    v15 = v15
                case "hello":
                    var v16 bool = false
                    _ = v16
                default:
                    v17 := GenS1{F0: true, F1: 255}
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
        }
        _ = fi13
    }
    fmt.Println(v2)
    fmt.Println(v2)
    v2 -= v2
    delete(v3, ((("\n\t" + (("hello" + (("\n\t" + (("a" + "a") + ("" + ""))) + "")) + ("hello" + "foo bar"))) + "a") + "hello"))
    delete(v1, 1)
    var v18 uint = genFn0()
    switch "a" {
    case "a":
        v2 += v2
    }
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
