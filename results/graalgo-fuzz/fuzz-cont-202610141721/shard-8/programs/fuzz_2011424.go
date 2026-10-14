package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int16 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []string{"a", (((("\n\t" + ("" + "a")) + "foo bar") + (((("a" + ("" + (("hello" + "hello") + ("" + ("a" + "\n\t"))))) + ((("a" + "hello") + ("\n\t" + "")) + ("" + ""))) + "a") + "")) + "\n\t"), "\n\t", "foo bar"}
    var v2 uint64 = 0
    var v3 int64 = 42
    var v4 uint16 = 100
    if false {
        var v5 int = 2147483647
        if ((v4 + v4) >= (v4 * 2)) {
            if (!(65535 >= 2147483648)) {
                var v6 uint = 0
                _ = v6
            } else {
                v1 = append(v1, ("" + ""))
                switch "" {
                case "\n\t":
                    if true {
                        v3 = v3
                    }
                }
            }
            v1 = append(v1, "a")
        } else {
            var v7 float64 = 0.0
            _ = v7
        }
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        _ = v5
    } else {
        fmt.Println(v2, v2)
    }
    var v8 bool = (-671 > 0)
    switch "a" {
    case "":
        v9 := GenS1{F0: 773, F1: (v4 | 529), F2: int16(v3)}
        _ = v9
    }
    v4 -= 0
    fmt.Println(v4, v2, v2)
    switch "" {
    case "foo bar":
        for fi10 := 0; fi10 < 10; fi10++ {
            fmt.Println(v3)
            _ = fi10
        }
    }
    var v11 int16 = 42
    v12 := GenS1{F0: v11, F1: (v4 + (v4 | v4)), F2: (-(v11))}
    for fi13 := 0; fi13 < 5; fi13++ {
        v2 = v2
        fmt.Println(v4, v8)
        if (!true) {
            fmt.Println(v11, v11)
            if len(v1) > 0 {
                v1[0] = "hello"
            }
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 bool = (0.0 > ((float64((-2.25 + float32(0.0))) * (0.0 * 0.1)) - 0.1))
                _ = fi14
                _ = v15
            }
        } else {
            v16 := GenS1{F0: (v11 & (v11 / ((v11) | 1))), F1: 32767, F2: 0}
            v4 -= 2
            for fi17 := 0; fi17 < 11; fi17++ {
                _ = len(v1)
                v8 = (v3 <= (v3 + v3))
                _ = fi17
            }
            _ = v16
        }
        _ = fi13
    }
    if v8 {
        var v18 string = "\n\t"
        _ = v18
    } else {
        v19 := GenS2{F0: (^(v3))}
        var v20 int16 = v11
        fmt.Println(v8, v20)
        _ = v19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
