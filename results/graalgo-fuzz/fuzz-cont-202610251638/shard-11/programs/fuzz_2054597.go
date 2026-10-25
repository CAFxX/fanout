package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 string; F1 int32; F2 int32 }
func genFn0(p0 bool) uint {
    var v6 uint64 = 18446744073709551615
    v7 := []int{-2147483648, -358}
    _ = p0
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 int32, p1 uint64, p2 int8) uint8 {
    p2 = p2
    var v8 int64 = 9223372036854775807
    if false {
        p2 -= p2
        v9 := GenS1{F0: 14, F1: p0}
        v10 := []int{128}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        switch ("hello" + "foo bar") {
        case "":
            var v12 int32 = (p0 / ((2147483647) | 1))
            _ = v12
        }
        if (p1 > p1) {
            p1 = (p1 * p1)
            v13 := []int{(fi11 - (fi11 & fi11))}
            v8 += v8
            _ = v13
        }
        _ = fi11
    }
    if false {
        v14 := GenS2{F0: "hello", F1: p0, F2: (p0 & 42)}
        if true {
            p2 -= p2
            fmt.Println(p0, v8, p0)
            if (!true) {
                v15 := []int{-2147483648, -1, 2147483647, 42}
                v15 = append(v15, 42)
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = v15
            }
        } else {
            var v16 float64 = 3.14159
            var v17 uint64 = p1
            _ = v16
            _ = v17
        }
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2
}

func main() {
    var v1 int = 0
    v2 := []string{"\n\t"}
    var v3 string = "foo bar"
    v4 := make(map[string]int)
    v5 := []int{127, (v1 - (v1 % ((2147483647) | 1)))}
    var v18 int = (v1 * (v1 + v1))
    _ = len(v2)
    for fi19 := 0; fi19 < 3; fi19++ {
        v1 -= (v1 & (fi19 >> ((505) & 63)))
        _ = fi19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        v21 := GenS2{F0: v3, F1: 42, F2: 2147483647}
        _ = fi20
        _ = v21
    }
    if true {
        v1 = (v18 - 128)
        var v22 int16 = -885
        _ = v22
    }
    if ((0.1 == 1e10) || ((float64(0.0) > -2.25) || (((343 == -568) || (!(170 >= 18446744073709551615))) && (1e10 != 0.0)))) {
        for fi23 := 0; fi23 < 18; fi23++ {
            v3 = (("foo bar" + (v3 + ("\n\t" + "a"))) + "")
            v24 := GenS2{F0: "\n\t", F1: -365, F2: 2147483647}
            _ = fi23
            _ = v24
        }
    }
    var v25 uint64 = 1
    if len(v5) > 0 {
        v5[0] = v1
    }
    v4["a"] = (v18 ^ (v1 + 1))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
