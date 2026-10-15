package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint }
func genFn0() uint8 {
    v7 := []int{1, 0, 127, 2147483647}
    var v8 int16 = 644
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 float32 = (-2.25 + (-2.25 / (((0.0 * -2.25) - 3.14159) * 1.5)))
    var v2 string = "hello"
    v3 := []string{(v2 + "hello"), ("hello" + ((v2 + (v2 + (("foo bar" + v2) + ("\n\t" + "a")))) + "\n\t")), v2}
    var v4 uint16 = 328
    v5 := []uint32{1, 538, 917}
    var v6 int64 = 772
    v1 = (v1 + v1)
    var v9 int32 = -2147483648
    fmt.Println(v9, v9, v9)
    switch -580 {
    case 32767:
        for fi10 := 0; fi10 < 10; fi10++ {
            v2 = (("foo bar" + v2) + v2)
            switch fi10 {
            case 104:
                var v11 uint = 1
                _ = v11
            }
            v9 = v9
            _ = fi10
        }
    case 0:
        fmt.Println(v6)
    }
    var v12 uint16 = v4
    var v13 bool = true
    v14 := GenS1{F0: (v2 + (((v2 + v2) + ("" + ("hello" + "foo bar"))) + "foo bar"))}
    var v15 uint8 = 255
    v16 := GenS1{F0: v2}
    v12 += (v4 ^ v12)
    var v17 uint64 = 1
    v17 -= (v17 >> ((4294967295) & 63))
    switch 42 {
    case 0:
        if false {
            v18 := GenS1{F0: (((v2 + "hello") + v2) + (((v2 + (v2 + (v2 + v2))) + (v2 + ("" + ("a" + "")))) + "a"))}
            if (v13 || (!(v17 == (v17 / ((v17) | 1))))) {
                v19 := GenS1{F0: v2}
                _ = v19
            } else {
                v12 = 32767
                var v20 int = 1
                _ = v20
            }
            _ = v18
        } else {
            if v13 {
                var v21 uint8 = (v15 | (^(v15)))
                v15 += (v21 << ((2) & 63))
                _ = v21
            }
            v2 = v2
            v22 := GenS2{F0: 0}
            _ = v22
        }
    case 565:
        v23 := GenS3{F0: 801, F1: 4294967295}
        _ = v23
    case 32768:
        var v24 int = 32768
        _ = v24
    default:
        var v25 float32 = v1
        _ = v25
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        _ = len(v3)
        _ = fi26
    }
    var v27 int32 = v9
    v12 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
