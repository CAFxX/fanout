package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() string {
    var v5 float64 = 1e10
    if (((!(-466 != -32768)) || (3.14159 < (((1.5 / -2.25) + ((-2.25 / float32((v5 / 1.0))) + (0.1 / 1.0))) / 0.1))) || (0.0 != (0.0 + (1e10 * 0.0)))) {
        v6 := GenS1{F0: 56}
        v7 := []int{1, 4294967295}
        _ = v6
        _ = v7
    } else {
        v5 += (0.0 - -2.25)
    }
    v5 = (v5 * (-2.25 + float64((1.5 / float32((0.1 * ((((v5 + v5) + v5) / v5) - 1.5)))))))
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int8 = -128
        for fi10 := 0; fi10 < 17; fi10++ {
            fmt.Println(v9, v9, v9)
            fmt.Println(v5, v9, v5)
            if (1 == 255) {
                var v11 uint = 2
                v12 := []int{(fi8 ^ (fi8 * -2147483648)), fi10, fi8}
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    _ = v5
    return (("hello" + "\n\t") + "\n\t")
}
func genFn1(p0 bool, p1 bool) int16 {
    switch "" {
    case "foo bar":
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 int16 = 0
            fmt.Println(p1, fi13, fi13)
            v14 = (v14 % ((-1) | 1))
            _ = fi13
            _ = v14
        }
    case "hello":
        var v15 int16 = 504
        _ = v15
    default:
        p0 = (18446744073709551615 > 46)
    }
    p0 = (917 == 2147483648)
    var v16 bool = true
    var v17 int64 = 1
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return -902
}

func main() {
    var v1 uint8 = 127
    var v2 uint32 = 2
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    v18 := v4[v2]
    fmt.Println(v18)
    v1 += v1
    var v19 uint16 = 32767
    var v20 uint = 65535
    fmt.Println(v19)
    v21 := []int{546, 32768, -413, -833}
    var v22 uint8 = (v1 >> ((v20) & 63))
    fmt.Println(v19, v1)
    if (true || false) {
        v1 += (v1 / ((190) | 1))
        var v23 int = 42
        _ = v23
    } else {
        if false {
            var v24 uint = (v20 / ((v20) | 1))
            v25 := GenS1{F0: 127}
            _ = v24
            _ = v25
        } else {
            if len(v21) >= 2 {
                v21 = v21[1:len(v21)]
            }
        }
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        v2 -= v2
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
