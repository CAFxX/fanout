package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
func genFn0() uint64 {
    switch (("a" + "hello") + (("" + ("foo bar" + "foo bar")) + (((("hello" + ("foo bar" + "a")) + ("foo bar" + "a")) + "hello") + "hello"))) {
    case "hello":
        var v7 uint = 1
        _ = v7
    case "a":
        var v8 bool = ((94 == 1) && false)
        _ = v8
    case "\n\t":
        v9 := []int{}
        _ = v9
    }
    v10 := []int{}
    v11 := GenS1{F0: 1, F1: 1}
    var v12 int16 = -32768
    _ = v10
    _ = v11
    _ = v12
    return 42
}
func genFn1() bool {
    var v13 bool = false
    v14 := []int{}
    var v15 uint32 = 2147483648
    _ = v13
    _ = v14
    _ = v15
    return v13
}

func main() {
    var v1 uint64 = 1
    var v2 uint16 = 61
    var v3 uint64 = (^(v1))
    v4 := make(map[uint32]string)
    var v5 uint8 = uint8(v1)
    var v6 uint = 2
    v16 := GenS1{F0: 1, F1: 1000}
    if true {
        switch 42 {
        case 128:
            v2 = (^(v2))
        case 42:
            var v17 int = 128
            _ = v17
        default:
            v18 := GenS1{F0: -414, F1: 392}
            _ = v18
        }
        var v19 uint16 = (v2 & (v2 << ((v6) & 63)))
        _ = v19
    } else {
        v1 -= (v3 >> ((817) & 63))
    }
    v20 := []int{int(v5), -696, 1, 128}
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 bool = (45 < -1)
        _ = fi21
        _ = v22
    }
    var v23 int = 9223372036854775807
    var v24 int8 = 100
    var v25 bool = (((v24 * (v24 * -1)) != int8(v1)) || (42 >= 9223372036854775807))
    v26 := GenS1{F0: 9223372036854775807, F1: 1000}
    for fi27 := 0; fi27 < 7; fi27++ {
        var v28 uint = 1
        v5 -= v5
        _ = fi27
        _ = v28
    }
    var v29 int8 = v24
    var v30 uint32 = 391
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
