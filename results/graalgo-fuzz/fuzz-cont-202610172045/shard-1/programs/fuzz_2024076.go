package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 string = (((((("hello" + (("" + "a") + ("a" + ""))) + "\n\t") + ("hello" + "\n\t")) + ("a" + "\n\t")) + (("hello" + ("hello" + ("" + "foo bar"))) + ("foo bar" + ((("a" + (("foo bar" + "foo bar") + "")) + "a") + ("hello" + "foo bar"))))) + "hello")
    var v2 int64 = 1
    var v3 bool = false
    v2 += (v2 * v2)
    v2 += -1
    v4 := []int{128, 2147483647, 115}
    _ = len(v4)
    var v5 int32 = 400
    v6 := GenS1{F0: (v5 >> ((302) & 63)), F1: 487, F2: 0}
    v7 := GenS1{F0: v5, F1: 1, F2: 2}
    for fi8 := 0; fi8 < 16; fi8++ {
        v4 = append(v4, (fi8 + fi8))
        v2 += v2
        switch 9223372036854775807 {
        case 128:
            fmt.Println(v1)
        }
        _ = fi8
    }
    switch 32767 {
    case 4294967295:
        v9 := GenS1{F0: v5, F1: v5, F2: 42}
        _ = v9
    case -2147483648:
        v4 = append(v4, 128)
    default:
        v2 -= v2
    }
    if (68 > 18446744073709551615) {
        for fi10 := 0; fi10 < 10; fi10++ {
            if v3 {
                v1 = v1
            }
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 string = v1
                v5 -= v5
                v13 := GenS1{F0: v5, F1: v5, F2: 127}
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
        switch (v1 + (v1 + v1)) {
        case "foo bar":
            fmt.Println(v5)
        }
    }
    v14 := GenS2{F0: -128}
    var v15 string = (((v1 + v1) + v1) + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
