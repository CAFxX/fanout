package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
func genFn0(p0 string, p1 int, p2 int) int {
    p2 += (^(p1))
    v5 := GenS1{F0: 42, F1: 255}
    v6 := []int{}
    for fi7 := 0; fi7 < 9; fi7++ {
        p0 = "a"
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (^(p2))
}

func main() {
    var v1 int16 = 1000
    v2 := make(map[int]string)
    var v3 uint32 = 2
    v4 := []int{998, -1, 187, -9223372036854775808}
    var v8 float64 = (3.14159 / 3.14159)
    v8 -= v8
    v9 := GenS1{F0: (v1 & 32767), F1: 65535}
    _ = len(v2)
    var v10 float64 = float64((1.5 + 0.1))
    v1 = v1
    switch (("foo bar" + ((("" + "hello") + ("" + ("foo bar" + (((("hello" + ("a" + "\n\t")) + "\n\t") + (("a" + "a") + "")) + "\n\t")))) + "\n\t")) + "") {
    case "a":
        var v11 bool = false
        _ = v11
    case "foo bar":
        if (!(1 > -9223372036854775808)) {
            v12 := GenS1{F0: (v1 - 1000), F1: 979}
            var v13 int8 = 127
            _ = v12
            _ = v13
        }
    case "\n\t":
        v8 = 0.0
    default:
        v3 = 2147483647
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v1 -= (v1 & v1)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi14
    }
    v4 = append(v4, -9223372036854775808)
    v10 -= (1e10 - 1.5)
    fmt.Println(v1)
    if (42 == -757) {
        for fi15 := 0; fi15 < 19; fi15++ {
            v1 += v1
            for fi16 := 0; fi16 < 13; fi16++ {
                v1 -= (^(v1))
                var v17 uint = 2
                var v18 uint = v17
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 int = genFn0("", (fi15 - (fi15 / (((fi15 + fi15)) | 1))), (^(fi15)))
            _ = fi15
            _ = v19
        }
    }
    v20 := GenS1{F0: v1, F1: 825}
    fmt.Println(v1)
    v1 += v1
    var v21 string = ("" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
