package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int16, p1 uint32) uint8 {
    p0 = (p0 - p0)
    p1 -= p1
    var v7 float32 = (-2.25 + 1.5)
    var v8 int = 261
    v9 := []int{v8, v8, v8}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 45
}

func main() {
    var v1 int = 406
    v2 := []int{0, v1, v1}
    var v3 bool = true
    var v4 int16 = -32768
    v5 := []uint32{1, 42, 75, 221}
    v6 := make(map[uint32]uint32)
    var v10 uint16 = 42
    var v11 uint16 = (v10 | (^(v10)))
    v10 += 255
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            v3 = false
            _ = fi13
        }
        _ = fi12
    }
    v1 -= v1
    v10 = (v11 ^ v11)
    if (!(!(!(!v3)))) {
        v6[uint32(v10)] = 4294967295
    } else {
        v1 += (v1 | v1)
        v5 = append(v5, 1)
    }
    var v14 uint16 = 2
    v15 := GenS1{F0: 0}
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 4; fi17++ {
            v14 = (v14 / (((v11 / (((v11 / ((v10) | 1))) | 1))) | 1))
            _ = fi17
        }
        _ = fi16
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v6[4] = 2147483648
        _ = fi18
    }
    var v19 uint = 42
    var v20 uint32 = 4294967295
    switch ("foo bar" + "a") {
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "a":
        var v21 float32 = float32(1.5)
        _ = v21
    }
    v10 = (v11 & (v11 / ((505) | 1)))
    fmt.Println(v20, v10)
    if (v3 && (v1 > v1)) {
        if len(v5) > 0 {
            v5[0] = v20
        }
    } else {
        for fi22 := 0; fi22 < 6; fi22++ {
            switch "" {
            case "hello":
                for fi23 := 0; fi23 < 7; fi23++ {
                    v20 += (v20 / ((2147483648) | 1))
                    _ = fi23
                }
            case "":
                fmt.Println(v1, v11, v1)
            case "a":
                v10 = v10
            }
            var v24 int8 = int8(v20)
            switch ("" + "hello") {
            case "a":
                var v25 uint16 = uint16(v20)
                _ = v25
            case "foo bar":
                v6[v20] = 42
            }
            _ = fi22
            _ = v24
        }
    }
    var v26 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
