package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint16, p1 uint16) uint {
    var v4 int16 = -32768
    var v5 float64 = (1.5 * float64(1.5))
    v6 := GenS1{F0: (p1 % (((p1 & 325)) | 1))}
    var v7 uint32 = 2147483648
    p1 += (p1 + (p1 % ((289) | 1)))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int32, p1 int8, p2 uint16) uint {
    v8 := []int{}
    var v9 uint16 = 2
    var v10 float32 = -2.25
    p1 = (p1 ^ p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{}
    var v2 int = -2147483648
    var v3 int64 = 1
    for fi11 := 0; fi11 < 10; fi11++ {
        if (42 >= 857) {
            fmt.Println(v3, v3, v2)
        } else {
            var v12 string = "a"
            v13 := GenS1{F0: 255}
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 bool = false
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch ("\n\t" + (((("hello" + ("" + "hello")) + "foo bar") + "\n\t") + "foo bar")) {
        case "":
            var v16 uint32 = 42
            _ = v16
        case "a":
            fmt.Println(fi14, v3, v3)
        case "foo bar":
            var v17 int64 = (v3 - (v3 >> ((4294967295) & 63)))
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    var v18 int64 = (v3 | v3)
    if true {
        v1 = append(v1, 659)
    }
    switch (v2 + (v2 & 2147483647)) {
    case 0:
        var v19 int = v2
        _ = v19
    case 2147483647:
        if (false || (0 == 42)) {
            v2 -= 42
            var v20 int64 = v18
            _ = v20
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v1 = append(v1, 4294967295)
    fmt.Println(v18, v18)
    v2 += v2
    var v21 int8 = -128
    switch "" {
    case "":
        var v22 uint16 = 65535
        _ = v22
    case "foo bar":
        switch (v2 % (((v2 + (v2 % (((v2 - (v2 << ((685) & 63)))) | 1)))) | 1)) {
        case -529:
            var v23 string = ("\n\t" + "a")
            _ = v23
        case 120:
            v3 += 1
        case 42:
            v24 := GenS1{F0: 255}
            _ = v24
        default:
            var v25 uint32 = 2
            _ = v25
        }
    }
    v3 = (-(v3))
    fmt.Println(v2, v3)
    v26 := GenS1{F0: 32767}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v26
    _ = v3
}
