package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
func genFn0(p0 int64, p1 int64, p2 int32) int64 {
    fmt.Println(p0)
    v7 := GenS1{F0: 42, F1: 168}
    p2 = p2
    if (p0 == p1) {
        var v8 uint32 = 2
        p2 = (p2 ^ (-(p2)))
        _ = v8
    } else {
        if false {
            p0 += p1
            v9 := GenS1{F0: 0, F1: 127}
            _ = v9
        }
        var v10 int64 = p0
        if false {
            v11 := GenS1{F0: (p2 + (p2 << ((218) & 63))), F1: 255}
            var v12 bool = true
            _ = v11
            _ = v12
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 int64, p1 uint16) string {
    p1 += p1
    fmt.Println(p1, p0)
    var v13 uint32 = 2
    _ = p0
    _ = p1
    _ = v13
    return (("foo bar" + ("foo bar" + (("" + (("a" + "\n\t") + ("hello" + ("hello" + "foo bar")))) + ((("a" + "hello") + "\n\t") + "a")))) + "\n\t")
}

func main() {
    v1 := []uint32{189, 86, 2147483648, 2}
    var v2 uint = 4294967295
    var v3 int16 = -694
    var v4 int = -2147483648
    v5 := make(map[string]bool)
    v6 := make(map[uint32]string)
    var v14 float32 = (float32(0.1) - float32((((1.5 * 0.1) / (-2.25 / 1e10)) * 0.0)))
    var v15 float32 = (v14 / (v14 + ((v14 * 3.14159) - (v14 + float32(float64(v14))))))
    v14 -= (1e10 + (v14 + 1.5))
    v1 = append(v1, 351)
    v15 += ((0.0 * float32(-2.25)) - float32(float64(0.0)))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v16 uint16 = 875
    v5["hello"] = true
    if false {
        v16 = (v16 << ((4294967295) & 63))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
