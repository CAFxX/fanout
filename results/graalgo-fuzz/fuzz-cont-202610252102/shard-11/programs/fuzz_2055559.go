package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 string }
func genFn0() int8 {
    var v7 int32 = -1
    if false {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 int = -178
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := GenS1{F0: "a"}
            var v12 float32 = (-2.25 / 3.14159)
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    _ = v7
    return 1
}
func genFn1() int32 {
    v13 := GenS1{F0: ("foo bar" + "")}
    var v14 float64 = -2.25
    _ = v13
    _ = v14
    return 567
}

func main() {
    var v1 uint64 = 42
    var v2 int = -9223372036854775808
    v3 := []string{("" + "foo bar"), ("\n\t" + "foo bar")}
    var v4 float32 = 1.5
    v5 := make(map[string]bool)
    v6 := make(map[string]bool)
    if ("hello" != "") {
        v15 := GenS2{F0: false, F1: ""}
        _ = v15
    }
    if true {
        fmt.Println(v1)
        v1 = (^(v1))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v16 float64 = 1.5
    fmt.Println(v16, v1)
    var v17 int32 = genFn1()
    v17 = genFn1()
    var v18 int64 = -285
    v16 -= v16
    v5["hello"] = true
    switch v2 {
    case 42:
        var v19 string = "foo bar"
        _ = v19
    case 9223372036854775807:
        var v20 int64 = (v18 << ((uint(v1)) & 63))
        _ = v20
    case 2147483647:
        if (191 > (v2 << ((4294967295) & 63))) {
            v4 -= (v4 * float32((v16 + v16)))
            var v21 int64 = v18
            _ = v21
        } else {
            fmt.Println(v17, v18)
            for fi22 := 0; fi22 < 16; fi22++ {
                var v23 int = -1
                v24 := GenS1{F0: "\n\t"}
                _ = fi22
                _ = v23
                _ = v24
            }
            switch "foo bar" {
            case "foo bar":
                v4 -= v4
            }
        }
    }
    v2 += (-(v2))
    fmt.Println(v4, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
