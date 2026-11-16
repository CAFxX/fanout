package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int; F2 string }
func genFn0(p0 uint16, p1 int8, p2 uint) uint32 {
    fmt.Println(p1, p1, p1)
    v5 := []int{684, 0}
    v6 := GenS1{F0: "a"}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 uint16, p1 bool, p2 uint32) string {
    p0 += 817
    fmt.Println(p0)
    for fi7 := 0; fi7 < 7; fi7++ {
        switch ("a" + (("foo bar" + "\n\t") + ("hello" + "hello"))) {
        case "foo bar":
            switch "hello" {
            case "foo bar":
                p0 = (p0 ^ p0)
            case "a":
                var v8 uint32 = 65535
                _ = v8
            default:
                fmt.Println(p2)
            }
        default:
            var v9 uint16 = (p0 * p0)
            _ = v9
        }
        _ = fi7
    }
    p2 = (p2 << ((4294967295) & 63))
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    return ("\n\t" + ("foo bar" + "foo bar"))
}

func main() {
    var v1 uint = 42
    v2 := []uint32{888, 874, 2, 2}
    var v3 int8 = 81
    var v4 float64 = ((-2.25 - 0.1) + ((0.0 - 3.14159) - (1.5 + float64((-2.25 + -2.25)))))
    var v10 int16 = -32768
    var v11 float32 = (-2.25 / (-2.25 - 0.1))
    fmt.Println(v4, v3)
    v11 -= ((float32(((((3.14159 / 1.0) / 1.0) / float64(0.1)) + v4)) / (0.1 / 1.5)) - 1e10)
    v10 -= v10
    var v12 bool = false
    fmt.Println(v4, v4, v11)
    for fi13 := 0; fi13 < 17; fi13++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := GenS2{F0: 2147483647, F1: fi14, F2: ("\n\t" + "\n\t")}
        v12 = (!v12)
        if v12 {
            for fi16 := 0; fi16 < 15; fi16++ {
                v11 = (((v11 / float32((v4 + 0.1))) - 0.1) / (v11 - 0.0))
                var v17 uint8 = 255
                fmt.Println(v3, v3)
                _ = fi16
                _ = v17
            }
        }
        _ = fi14
        _ = v15
    }
    var v18 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
