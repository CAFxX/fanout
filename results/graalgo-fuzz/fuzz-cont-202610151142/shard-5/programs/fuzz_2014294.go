package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint32 }
func genFn0(p0 int, p1 uint8, p2 int8) bool {
    for fi4 := 0; fi4 < 12; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            v6 := []int{}
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    var v7 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return v7
}
func genFn1(p0 int8, p1 uint32, p2 bool) uint8 {
    fmt.Println(p2, p1, p0)
    if p2 {
        var v8 uint = 18446744073709551615
        _ = v8
    } else {
        switch ("" + "hello") {
        case "":
            var v9 uint16 = 255
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 148
}

func main() {
    var v1 uint32 = 2
    v2 := []uint32{65535, v1}
    var v3 uint8 = 128
    v2 = append(v2, v1)
    var v10 uint8 = genFn1(-1, (v1 >> ((0) & 63)), (42 > 532))
    v1 += (v1 / (((v1 - v1)) | 1))
    v1 += 2147483648
    var v11 int16 = 159
    var v12 int = -1
    v10 = 1
    v3 = (v10 % (((v10 << ((255) & 63))) | 1))
    v11 -= -506
    v11 -= (v11 + (-(v11)))
    var v13 bool = ((!false) || true)
    for fi14 := 0; fi14 < 20; fi14++ {
        var v15 uint32 = (v1 >> ((173) & 63))
        var v16 int32 = 0
        if (-1 < 42) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v2 = append(v2, (v15 << ((255) & 63)))
        } else {
            v13 = (!(0 > 0))
            for fi17 := 0; fi17 < 6; fi17++ {
                fmt.Println(v1, fi17)
                _ = fi17
            }
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(v1, v13, v12)
    fmt.Println(v13, v11)
    v18 := GenS1{F0: 32767, F1: -2147483648, F2: v1}
    v19 := GenS1{F0: 469, F1: 2147483647, F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
