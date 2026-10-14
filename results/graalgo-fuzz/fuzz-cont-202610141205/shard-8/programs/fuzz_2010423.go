package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint16, p1 int8) int8 {
    var v7 int = 128
    if (!true) {
        p1 -= (-(p1))
    }
    _ = p0
    _ = p1
    _ = v7
    return (p1 ^ 0)
}
func genFn1(p0 uint32, p1 uint) uint8 {
    if (255 != 255) {
        v8 := GenS1{F0: 308}
        var v9 int16 = 0
        fmt.Println(v9, v9)
        _ = v8
        _ = v9
    } else {
        p1 -= p1
    }
    p0 = (p0 * p0)
    var v10 float64 = float64((-2.25 / 1.0))
    _ = p0
    _ = p1
    _ = v10
    return 2
}

func main() {
    var v1 bool = (4294967295 == 348)
    var v2 int8 = -41
    var v3 float32 = float32(1.5)
    v4 := make(map[int]uint32)
    var v5 uint = 817
    var v6 string = ((("hello" + ((("\n\t" + "hello") + ((((("foo bar" + "foo bar") + (("foo bar" + "\n\t") + "a")) + ("foo bar" + "\n\t")) + "") + "foo bar")) + "")) + "a") + "hello")
    var v11 uint = (v5 ^ (v5 * 306))
    var v12 string = v6
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint16 = 255
        _ = fi13
        _ = v14
    }
    fmt.Println(v5)
    var v15 uint32 = 0
    var v16 bool = (0.1 != float64(1.5))
    fmt.Println(v16, v12)
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 int = fi17
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 uint = v5
        _ = fi19
        _ = v20
    }
    var v21 uint8 = genFn1(v15, v11)
    var v22 bool = ((!(!(((3.14159 * 0.1) + 1e10) >= (0.1 - (0.1 + ((((3.14159 * 3.14159) + 1.5) + float64(((v3 * 1e10) / v3))) * 0.1)))))) && v1)
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 float32 = ((1e10 + 1e10) + v3)
        _ = fi23
        _ = v24
    }
    switch -9223372036854775808 {
    case 0:
        var v25 uint32 = 1
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
