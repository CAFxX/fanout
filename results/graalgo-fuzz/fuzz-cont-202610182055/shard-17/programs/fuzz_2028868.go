package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int16 {
    var v7 int = 32768
    for fi8 := 0; fi8 < 17; fi8++ {
        switch v7 {
        case 1:
            if false {
                v7 = (v7 / ((fi8) | 1))
            } else {
                if false {
                    var v9 float32 = 3.14159
                    _ = v9
                }
                v7 -= 660
            }
        case 32768:
            v7 = (v7 * v7)
        default:
            var v10 uint = 0
            _ = v10
        }
        _ = fi8
    }
    var v11 int = 0
    _ = v11
    _ = v7
    return 926
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{1, 2147483647, 404}
    v3 := []int{-559}
    var v4 string = ("foo bar" + "a")
    v5 := make(map[uint32]uint32)
    var v6 uint = 2
    fmt.Println(v6, v4)
    var v12 float32 = 1.5
    if len(v2) > 0 {
        v2[0] = 65535
    }
    if false {
        v13 := GenS1{F0: (v6 / ((v6) | 1))}
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint32 = 4294967295
        var v16 float32 = v12
        if ((874 <= 1) && (32767 < 1)) {
            v12 += -2.25
            switch v4 {
            case "hello":
                v12 -= 3.14159
            }
            for fi17 := 0; fi17 < 6; fi17++ {
                v4 = v4
                v6 += 255
                _ = fi17
            }
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    var v18 uint = v6
    var v19 string = v4
    v1[("hello" + v4)] = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
