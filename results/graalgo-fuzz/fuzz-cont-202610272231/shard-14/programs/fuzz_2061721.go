package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
func genFn0() int64 {
    var v5 uint8 = 42
    v6 := []int{-252, 1, 127}
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []string{"a", "\n\t", ""}
    var v2 int16 = 0
    v3 := []uint32{276}
    v4 := []uint32{895}
    var v7 int = -1
    var v8 uint8 = 36
    if (1 <= -128) {
        if true {
            v2 = (^(v2))
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            v2 -= v2
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v9 := GenS1{F0: int8(v8)}
            _ = v9
        }
        v7 = 1
        switch 9223372036854775807 {
        case 9223372036854775807:
            v8 = (^(v8))
        case -365:
            v10 := GenS1{F0: 42}
            _ = v10
        case 4294967295:
            var v11 int64 = genFn0()
            _ = v11
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        v1 = append(v1, "a")
    } else {
        for fi12 := 0; fi12 < 17; fi12++ {
            v8 -= 128
            fmt.Println(v2)
            _ = fi12
        }
        var v13 int8 = -1
        for fi14 := 0; fi14 < 13; fi14++ {
            v13 += (v13 / ((v13) | 1))
            var v15 int32 = -1
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    if len(v4) > 0 {
        v4[0] = 749
    }
    switch v7 {
    case 410:
        v8 = (v8 >> ((69) & 63))
    case -782:
        v8 = 237
    case -550:
        v16 := GenS2{F0: "\n\t"}
        _ = v16
    default:
        for fi17 := 0; fi17 < 6; fi17++ {
            v7 -= (fi17 << ((255) & 63))
            v7 += v7
            for fi18 := 0; fi18 < 7; fi18++ {
                v19 := GenS1{F0: -128}
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    }
    fmt.Println(v2, v2)
    switch ((("foo bar" + "foo bar") + ("a" + ("foo bar" + ("" + "hello")))) + "hello") {
    case "":
        v7 = v7
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (v8 <= (^(v8))) {
        fmt.Println(v7)
        v3 = append(v3, 86)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println(v7, v8)
    var v20 uint8 = 182
    v7 = (v7 + -2147483648)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
