package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 string }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint64, p1 uint16, p2 int) bool {
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint16 = (p1 ^ (p1 | (p1 - 42)))
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: p1, F1: p1, F2: ("\n\t" + ("hello" + "foo bar"))}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return (128 > 255)
}
func genFn1(p0 int8, p1 uint8) int32 {
    var v10 uint64 = 2
    v11 := []int{-415, 4294967295, 42}
    fmt.Println(p0, p1)
    var v12 uint16 = 250
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    var v3 uint = 60
    var v4 int16 = -1
    v5 := make(map[uint32]int)
    var v6 float32 = 3.14159
    v13 := v5[2147483647]
    v4 = v4
    v13 += v13
    if true {
        delete(v1, "\n\t")
        _ = len(v2)
    } else {
        v13 = -9223372036854775808
        switch ("foo bar" + "") {
        case "":
            v3 += (v3 << ((2) & 63))
        }
    }
    var v14 int = (v13 / ((264) | 1))
    if len(v2) > 0 {
        v2[0] = 65535
    }
    var v15 int64 = -1
    var v16 uint64 = 2
    var v17 string = (("a" + "") + "a")
    v6 = ((((float32(0.0) / float32(3.14159)) * 0.1) - float32((((0.1 - ((1.5 + float64((float32(3.14159) / ((v6 - -2.25) - v6)))) * 1.5)) + 0.1) / 1.0))) + 3.14159)
    v18 := GenS2{F0: (v16 + (^(v16)))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
