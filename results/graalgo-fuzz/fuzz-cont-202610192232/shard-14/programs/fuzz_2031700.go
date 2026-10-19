package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() string {
    v6 := []int{-9223372036854775808, 4294967295, 42}
    var v7 int = 1
    var v8 uint8 = 2
    if (42 <= 42) {
        _ = len(v6)
    } else {
        v7 = 0
        fmt.Println(v7, v8)
    }
    v7 = -1
    _ = v6
    _ = v7
    _ = v8
    return "a"
}
func genFn1(p0 int64, p1 uint16, p2 int64) int32 {
    var v9 int16 = 32767
    fmt.Println(v9, v9)
    var v10 uint16 = (p1 & (^(p1)))
    var v11 float32 = (1e10 - -2.25)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 90
}

func main() {
    var v1 float32 = (0.1 - -2.25)
    var v2 uint8 = 158
    v3 := make(map[string]int)
    var v4 uint16 = 255
    v5 := []int{-9223372036854775808, 42, -500}
    switch ("hello" + "hello") {
    case "foo bar":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        fmt.Println(v4, v1)
    }
    v12 := GenS1{F0: 1}
    var v13 int8 = 100
    fmt.Println(v4, v4, v1)
    switch -2147483648 {
    case 273:
        _ = len(v3)
    case 2147483647:
        v14 := GenS1{F0: 18446744073709551615}
        _ = v14
    case 42:
        v3[""] = -2147483648
    }
    delete(v3, (("\n\t" + "hello") + "a"))
    if true {
        v2 += 0
        fmt.Println(v2, v13, v13)
        v2 -= v2
    } else {
        fmt.Println(v1)
    }
    var v15 uint16 = 32767
    var v16 int = -963
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 uint8 = 128
        v13 -= v13
        _ = fi17
        _ = v18
    }
    var v19 uint8 = 252
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
