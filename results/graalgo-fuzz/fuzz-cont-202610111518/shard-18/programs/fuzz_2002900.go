package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        var v5 uint8 = 0
        var v6 string = (("a" + "hello") + "")
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := []int{704, 42, 1}
    var v8 uint16 = 0
    var v9 uint64 = 970
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: uint(v8)}
        _ = fi10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint64, p1 uint32) int8 {
    if (-2.25 <= ((((3.14159 * ((1.5 / 1.5) - (1e10 + 1e10))) / 1e10) * (float64(3.14159) - ((0.1 + 0.0) * 1e10))) + 1.5)) {
        var v12 uint8 = 255
        _ = v12
    }
    p0 = p0
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 float64 = 0.1
    var v2 int64 = 422
    v3 := []string{"a", ""}
    switch 1 {
    case -1:
        v1 = ((1.5 / 3.14159) - 3.14159)
    default:
        if false {
            if (!false) {
                if len(v3) > 0 {
                    v3[0] = ("a" + ("hello" + "foo bar"))
                }
            } else {
                var v13 int8 = -1
                _ = v13
            }
            v2 += 1
        }
    }
    fmt.Println(v1)
    if len(v3) > 0 {
        v3[0] = (((("" + ("foo bar" + ("" + ("" + "")))) + "hello") + "hello") + "a")
    }
    v2 -= (v2 & v2)
    v2 = v2
    v1 += 0.1
    var v14 uint8 = 2
    var v15 int64 = v2
    fmt.Println(v1, v2, v2)
    var v16 uint8 = 111
    fmt.Println(v16)
    var v17 bool = false
    if (2 < 2) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v18 uint16 = 42
        _ = len(v3)
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 uint32 = 1
        _ = fi19
        _ = v20
    }
    var v21 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
}
