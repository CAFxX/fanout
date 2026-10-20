package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int }
func genFn0(p0 uint16, p1 int16, p2 bool) uint8 {
    v4 := GenS1{F0: "hello", F1: -117, F2: 1}
    p2 = (true || true)
    switch 169 {
    case 27:
        if (p2 || true) {
            v5 := []int{42, -2147483648, -275}
            for fi6 := 0; fi6 < 8; fi6++ {
                var v7 uint32 = 2147483648
                _ = fi6
                _ = v7
            }
            if p2 {
                p1 -= (p1 - p1)
                p0 += (p0 ^ 356)
            }
            _ = v5
        }
    case -1:
        p0 = (p0 << ((2) & 63))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 239
}
func genFn1(p0 int) uint16 {
    v8 := GenS1{F0: "a", F1: 902, F2: (p0 ^ p0)}
    p0 -= p0
    _ = p0
    _ = v8
    return 1
}

func main() {
    var v1 bool = (!(false && true))
    var v2 uint8 = 42
    v3 := make(map[string]uint32)
    v9 := GenS1{F0: (((((("\n\t" + ((("" + "foo bar") + ((("\n\t" + "") + "") + ("" + ""))) + "\n\t")) + "foo bar") + "foo bar") + ("" + "a")) + "foo bar") + ("" + "\n\t")), F1: -2147483648, F2: 32767}
    switch 128 {
    case 127:
        v10 := GenS1{F0: "\n\t", F1: 2147483647, F2: -27}
        _ = v10
    case -235:
        var v11 int = 9223372036854775807
        _ = v11
    }
    if true {
        if v1 {
            v2 = (v2 ^ 218)
            var v12 int = 547
            fmt.Println(v1, v2)
            _ = v12
        } else {
            _ = len(v3)
            var v13 uint64 = 2
            _ = v13
        }
    } else {
        v2 += (^(v2))
    }
    v2 -= v2
    v14 := GenS1{F0: "", F1: 1, F2: 4294967295}
    if ("hello" > "hello") {
        fmt.Println(v1, v1)
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        v1 = (v1 && false)
        v16 := GenS1{F0: "hello", F1: -567, F2: 42}
        v17 := GenS1{F0: "", F1: 671, F2: (fi15 | (fi15 * (fi15 << ((4294967295) & 63))))}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
