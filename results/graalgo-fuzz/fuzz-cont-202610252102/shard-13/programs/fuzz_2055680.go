package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint }
type GenS2 struct { F0 int32; F1 bool }
func genFn0(p0 int64) string {
    var v5 int16 = -1
    v6 := GenS1{F0: 878, F1: 2147483647, F2: 42}
    fmt.Println(p0)
    _ = p0
    _ = v5
    _ = v6
    return ("foo bar" + ("" + "\n\t"))
}
func genFn1() uint8 {
    var v7 int8 = 1
    if true {
        var v8 int8 = -1
        _ = v8
    }
    _ = v7
    return 2
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    v3 := []int{32768, 2147483647}
    var v4 string = (("" + "hello") + "\n\t")
    switch v4 {
    case "hello":
        var v9 string = "\n\t"
        _ = v9
    case "a":
        if false {
            fmt.Println(v4)
        } else {
            if ("" < ((v4 + v4) + ("\n\t" + ("a" + ((("foo bar" + "\n\t") + v4) + (("hello" + ("a" + "\n\t")) + "\n\t")))))) {
                var v10 bool = false
                var v11 uint = 2
                _ = v10
                _ = v11
            }
        }
    case "":
        v1[v4] = ((v4 + v4) + v4)
    }
    var v12 int16 = 32767
    v13 := GenS2{F0: -2147483648, F1: false}
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(fi14)
        v3 = append(v3, (fi14 << ((255) & 63)))
        _ = len(v3)
        _ = fi14
    }
    v4 = ""
    if len(v3) > 0 {
        v3[0] = 42
    }
    v15 := GenS2{F0: 800, F1: ((2 != 933) && false)}
    var v16 int64 = -1
    var v17 uint8 = genFn1()
    v16 += (v16 ^ v16)
    var v18 uint32 = 42
    fmt.Println(v12, v4)
    v16 = (v16 * 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
