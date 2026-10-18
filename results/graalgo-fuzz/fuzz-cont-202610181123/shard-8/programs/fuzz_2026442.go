package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int8, p1 uint8) bool {
    var v6 int8 = -128
    v7 := []int{-823, 473}
    p0 -= p0
    p0 += v6
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = -70
    v3 := []uint32{2147483647}
    v4 := make(map[string]int)
    var v5 uint = 63
    if len(v3) > 0 {
        v3[0] = 65535
    }
    delete(v4, ("hello" + "foo bar"))
    v8 := GenS1{F0: (v2 + (^(v2)))}
    delete(v4, ("foo bar" + (("a" + "a") + "hello")))
    var v9 int32 = 1
    switch "hello" {
    case "hello":
        if (127 != 1) {
            v10 := GenS1{F0: v2}
            switch ((("a" + ("\n\t" + (("" + "") + (("foo bar" + "foo bar") + (("foo bar" + (("foo bar" + "foo bar") + "\n\t")) + ("a" + "")))))) + "foo bar") + "foo bar") {
            case "a":
                var v11 bool = (239 > 127)
                _ = v11
            default:
                var v12 uint8 = 1
                _ = v12
            }
            _ = v10
        } else {
            var v13 int32 = 2147483647
            var v14 uint32 = 42
            switch (("a" + "a") + ("a" + ("a" + ("a" + "foo bar")))) {
            case "":
                v14 -= uint32(v5)
            default:
                v13 = v13
            }
            _ = v13
            _ = v14
        }
    case "\n\t":
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 16; fi16++ {
                _ = len(v4)
                fmt.Println(fi16)
                var v17 int = (fi16 / (((fi16 / ((fi15) | 1))) | 1))
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
    case "a":
        var v18 uint = v5
        _ = v18
    default:
        var v19 string = ("a" + "foo bar")
        _ = v19
    }
    if (0 == -9223372036854775808) {
        fmt.Println(v2, v9, v9)
        v5 -= (v5 | 42)
        v20 := GenS1{F0: (v2 + v2)}
        _ = v20
    }
    fmt.Println(v9, v9, v2)
    var v21 bool = true
    fmt.Println(v2, v5, v21)
    var v22 uint16 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
