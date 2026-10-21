package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint32 }
func genFn0(p0 int8, p1 uint8, p2 uint64) uint32 {
    var v5 uint16 = 34
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}
func genFn1(p0 uint16, p1 bool, p2 int16) bool {
    v6 := []int{42}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return false
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 bool = true
    v4 := make(map[int]uint32)
    v3 = v3
    var v7 float64 = 3.14159
    fmt.Println(v7)
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        var v8 uint8 = 127
        _ = v8
    case "foo bar":
        var v9 float64 = (v7 / (v7 / (v7 - 3.14159)))
        _ = v9
    case "hello":
        for fi10 := 0; fi10 < 11; fi10++ {
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 float32 = ((float32(-2.25) + 1.5) / float32((((v7 - v7) / v7) / ((0.0 * (v7 * 1.5)) * 1e10))))
                v2[1] = ""
                var v13 uint32 = 2
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
    }
    var v14 int64 = 9223372036854775807
    var v15 float32 = float32(v7)
    v2[0] = "a"
    v15 = (v15 + (float32(v7) - float32(float64(v15))))
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            v7 = v7
            if (!(!v3)) {
                v4[(fi17 + 735)] = 802
                switch "\n\t" {
                case "a":
                    v18 := GenS1{F0: 42, F1: 2, F2: genFn0(100, 25, 976)}
                    _ = v18
                case "hello":
                    v15 = ((v15 - 1.5) / float32(3.14159))
                case "":
                    switch (((("foo bar" + "a") + "a") + (((("hello" + ("\n\t" + "\n\t")) + (("" + "a") + "a")) + "hello") + ("\n\t" + "foo bar"))) + "\n\t") {
                    case "hello":
                        v15 -= float32(v7)
                    case "":
                        var v19 uint16 = 0
                        _ = v19
                    }
                default:
                    var v20 float32 = -2.25
                    _ = v20
                }
            }
            delete(v4, fi17)
            v3 = (!false)
            v21 := []int{(fi16 | (fi17 << ((255) & 63))), (-(fi17)), -9223372036854775808, fi17}
            _ = fi17
            _ = v21
        }
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
