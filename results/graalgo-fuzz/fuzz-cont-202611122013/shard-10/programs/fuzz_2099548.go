package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int8; F2 string }
type GenS3 struct { F0 uint16; F1 uint32 }
func genFn0(p0 string, p1 int32) uint8 {
    if true {
        p1 += p1
        var v4 bool = true
        _ = v4
    }
    if true {
        v5 := []int{}
        var v6 int = 608
        fmt.Println(p1, v6)
        _ = v5
        _ = v6
    }
    switch 128 {
    case 676:
        for fi7 := 0; fi7 < 2; fi7++ {
            p0 = "\n\t"
            _ = fi7
        }
    case 305:
        v8 := []int{-2147483648, -2147483648, -9223372036854775808, -471}
        _ = v8
    case 0:
        var v9 uint16 = 32767
        _ = v9
    }
    _ = p0
    _ = p1
    return 121
}

func main() {
    var v1 int = 4294967295
    v2 := make(map[uint32]bool)
    var v3 int32 = -2147483648
    v10 := v2[42]
    var v11 string = (("\n\t" + ("\n\t" + ("foo bar" + (("foo bar" + ("hello" + "a")) + "hello")))) + "\n\t")
    switch (v11 + "hello") {
    case "foo bar":
        switch "a" {
        case "":
            v2[4294967295] = ((v3 ^ (v3 | v3)) < v3)
        case "foo bar":
            v12 := []int{v1, -361, v1}
            _ = v12
        default:
            fmt.Println(v3, v1)
        }
    }
    var v13 float32 = 0.1
    v1 = (v1 & 2147483647)
    var v14 int16 = -247
    var v15 uint16 = 42
    var v16 uint = 0
    if (42 >= 0) {
        v17 := GenS1{F0: (!(!(1e10 == 0.1)))}
        v10 = false
        _ = v17
    } else {
        var v18 uint64 = 2
        v15 -= v15
        _ = v18
    }
    switch (v11 + (("a" + (v11 + ("hello" + ("\n\t" + ((v11 + v11) + "foo bar"))))) + (("hello" + v11) + (v11 + v11)))) {
    case "foo bar":
        for fi19 := 0; fi19 < 5; fi19++ {
            v16 -= v16
            v20 := GenS2{F0: 2147483648, F1: -68, F2: v11}
            _ = fi19
            _ = v20
        }
    case "\n\t":
        v21 := GenS2{F0: 4294967295, F1: 0, F2: v11}
        _ = v21
    }
    var v22 string = "foo bar"
    var v23 int64 = int64(v3)
    for fi24 := 0; fi24 < 3; fi24++ {
        v25 := GenS3{F0: v15, F1: 42}
        v16 = 107
        v22 = "a"
        _ = fi24
        _ = v25
    }
    var v26 uint = (^(v16))
    for fi27 := 0; fi27 < 11; fi27++ {
        v28 := GenS2{F0: 42, F1: 0, F2: v22}
        var v29 uint8 = 108
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
}
