package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int64; F2 uint16 }
func genFn0() string {
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: (("a" + "a") + (("\n\t" + ("a" + "foo bar")) + "a"))}
        if (42 == 229) {
            var v7 uint64 = 925
            _ = v7
        }
        v8 := GenS2{F0: 230, F1: 0, F2: 2}
        _ = fi5
        _ = v6
        _ = v8
    }
    v9 := []int{1, 9223372036854775807}
    _ = v9
    return "a"
}
func genFn1(p0 int64, p1 int16) uint32 {
    var v10 string = ("a" + "\n\t")
    v11 := GenS2{F0: 128, F1: (^(p0)), F2: 255}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 26
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 1
    v3 := []string{((("hello" + "") + (("hello" + "\n\t") + "\n\t")) + "foo bar"), "foo bar"}
    v4 := make(map[int]string)
    var v12 int64 = v2
    var v13 uint16 = 2
    if false {
        v2 += v12
        v2 = v12
        fmt.Println(v12)
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v13 = (v13 & v13)
        v15 := v4[(fi14 * fi14)]
        _ = fi14
        _ = v15
    }
    v13 = 2
    v16 := GenS1{F0: ""}
    for fi17 := 0; fi17 < 6; fi17++ {
        _ = len(v3)
        var v18 int32 = -1
        _ = fi17
        _ = v18
    }
    v4[641] = (("foo bar" + ("\n\t" + "")) + ("" + ""))
    v19 := v4[-2147483648]
    var v20 int = -114
    v1[v19] = genFn1((v2 / ((-9223372036854775808) | 1)), 1)
    if false {
        var v21 int8 = 0
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
