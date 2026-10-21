package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint16, p1 int16, p2 uint8) uint32 {
    v4 := []int{4294967295, 32768, -218, 0}
    switch ((("" + "\n\t") + ((("hello" + ("\n\t" + "hello")) + "") + "\n\t")) + ("a" + ("a" + ""))) {
    case "hello":
        v5 := GenS1{F0: uint64(p2)}
        _ = v5
    case "foo bar":
        fmt.Println(p0, p2)
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    p0 = p0
    v4 = append(v4, 127)
    v6 := GenS2{F0: 2, F1: -1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return 2147483648
}

func main() {
    v1 := make(map[string]string)
    var v2 string = ("\n\t" + "a")
    var v3 int16 = -32768
    v7 := GenS2{F0: 1, F1: 42}
    if false {
        v8 := GenS2{F0: 2, F1: 42}
        var v9 int = 32767
        delete(v1, "a")
        _ = v8
        _ = v9
    }
    fmt.Println(v2, v3)
    var v10 int16 = v3
    v2 = v2
    var v11 uint8 = 85
    v12 := GenS1{F0: 118}
    v13 := GenS2{F0: 136, F1: 1}
    v14 := GenS2{F0: 0, F1: 2147483647}
    var v15 int = -9223372036854775808
    switch (v2 + ("a" + v2)) {
    case "hello":
        v16 := GenS2{F0: 599, F1: 736}
        _ = v16
    case "foo bar":
        v17 := GenS2{F0: 660, F1: 1}
        _ = v17
    default:
        var v18 int16 = (v10 << ((42) & 63))
        _ = v18
    }
    var v19 int16 = v3
    v3 -= v10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v7
}
