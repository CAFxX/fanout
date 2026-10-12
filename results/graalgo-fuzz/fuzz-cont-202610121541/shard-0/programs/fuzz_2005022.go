package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int, p1 int16, p2 int) int16 {
    var v4 int = (p2 ^ -382)
    v5 := GenS1{F0: p1}
    v6 := []int{p0, -873, 32767, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return p1
}

func main() {
    var v1 int = 32768
    var v2 uint16 = 42
    v3 := make(map[uint32]uint32)
    var v7 uint32 = 396
    fmt.Println(v2, v1)
    if false {
        v8 := []int{v1, (v1 - -9223372036854775808), -580, v1}
        delete(v3, v7)
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        v1 = int(v7)
        v10 := GenS1{F0: 0}
        var v11 string = ("hello" + "foo bar")
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v2, v2, v1)
    var v12 uint8 = 42
    v7 = v7
    v13 := GenS1{F0: 1}
    if false {
        v7 -= 2147483648
        fmt.Println(v7, v2, v7)
    }
    var v14 uint64 = 169
    switch 2147483647 {
    case 0:
        v15 := []int{(v1 | v1), 4294967295, 901, 0}
        _ = v15
    case 922:
        var v16 float32 = float32((3.14159 - 3.14159))
        _ = v16
    default:
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := []int{}
            v18 = append(v18, (fi17 * (v1 % ((fi17) | 1))))
            if len(v18) > 0 {
                v18[0] = fi17
            }
            _ = fi17
            _ = v18
        }
    }
    if true {
        var v19 float64 = 0.1
        if ((v14 != v14) || true) {
            fmt.Println(v7, v14, v2)
        } else {
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 int = (fi20 + v1)
                _ = fi20
                _ = v21
            }
        }
        _ = v19
    }
    v2 = v2
    v1 += 32767
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
}
