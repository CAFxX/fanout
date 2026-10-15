package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint64, p1 int16, p2 uint16) int16 {
    var v6 string = ("" + ("\n\t" + "a"))
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int = fi7
        var v9 uint8 = 2
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p1 | p1)
}
func genFn1(p0 uint) uint8 {
    if false {
        var v10 int8 = 127
        fmt.Println(p0)
        _ = v10
    }
    if (!((!false) || (4294967295 >= 2147483648))) {
        p0 = p0
        var v11 int8 = 100
        p0 = p0
        _ = v11
    }
    if true {
        var v12 bool = true
        var v13 int32 = -2147483648
        v13 = (v13 & (-(v13)))
        _ = v12
        _ = v13
    } else {
        var v14 int8 = -1
        _ = v14
    }
    var v15 uint8 = 1
    _ = p0
    _ = v15
    return v15
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int = 32768
    var v3 int = v2
    v4 := make(map[int]string)
    var v5 string = "hello"
    v16 := GenS1{F0: -1}
    _ = len(v4)
    v17 := GenS1{F0: -1}
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := []int{v2, (fi18 | (fi18 + 2147483647)), (v2 << ((520) & 63)), v2}
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 0}
    v2 = (v2 & v2)
    if true {
        v21 := GenS1{F0: 396}
        switch (v5 + (v5 + "a")) {
        case "a":
            var v22 uint32 = v1
            _ = v22
        case "hello":
            var v23 int64 = -9223372036854775808
            _ = v23
        default:
            v24 := GenS1{F0: genFn0(807, -32768, 42)}
            _ = v24
        }
        _ = v21
    } else {
        v2 = (v2 * (v2 >> ((4294967295) & 63)))
        v2 = 4294967295
    }
    var v25 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
