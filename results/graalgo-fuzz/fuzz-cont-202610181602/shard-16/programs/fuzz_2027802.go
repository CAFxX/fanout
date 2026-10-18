package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 1
    v3 := make(map[string]int)
    var v4 int64 = 1
    var v5 uint32 = 65535
    var v6 int64 = (v4 & 42)
    if ((("hello" + "") < "\n\t") || (!true)) {
        v6 = (v4 - 9223372036854775807)
        if (!(-358 < -2147483648)) {
            fmt.Println(v5, v6)
            var v7 int64 = 0
            _ = v7
        }
    }
    v8 := GenS1{F0: "hello", F1: 2}
    v9 := []int{-9223372036854775808, 1, 32768, 2147483647}
    if true {
        var v10 int32 = 0
        _ = v10
    } else {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int16 = 1
            v13 := v1["\n\t"]
            var v14 uint32 = uint32(v2)
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        fmt.Println(v4)
        v6 = (v4 + (v6 * (v4 & v4)))
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int = (^(fi15))
        v17 := GenS1{F0: "foo bar", F1: v5}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 float32 = 0.0
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 float64 = 0.1
        var v21 int = fi19
        if (2 != 766) {
            if (v2 > v2) {
                for fi22 := 0; fi22 < 12; fi22++ {
                    v4 += 0
                    v21 = (fi19 | (fi19 << ((695) & 63)))
                    v23 := GenS1{F0: "a", F1: 712}
                    _ = fi22
                    _ = v23
                }
                if (1 != -68) {
                    if ("\n\t" <= "hello") {
                        if (42 != 10) {
                            fmt.Println(v5)
                            v4 = v6
                            fmt.Println(v20)
                        }
                    }
                }
                var v24 bool = true
                _ = v24
            }
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    v25 := GenS1{F0: (("a" + "hello") + "foo bar"), F1: v5}
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 string = (("a" + ("" + "a")) + "hello")
        v28 := GenS2{F0: -1}
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println(v18, v5)
    var v29 int64 = (v4 + 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
