package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint64 }

func main() {
    v1 := []string{}
    v2 := []uint32{941, 2147483648}
    var v3 bool = false
    var v4 uint64 = 242
    v5 := make(map[uint32]uint32)
    v1 = append(v1, ("a" + "foo bar"))
    var v6 uint8 = 127
    v3 = v3
    if v3 {
        var v7 bool = (32767 != 109)
        v8 := GenS1{F0: 65535, F1: -32768, F2: v4}
        _ = v7
        _ = v8
    } else {
        if (("hello" + "\n\t") == "foo bar") {
            var v9 float32 = 0.0
            _ = v9
        } else {
            v4 = 0
            fmt.Println(v3)
            if ((true && v3) || ((2 < (v4 & (v4 << ((255) & 63)))) || v3)) {
                v6 = v6
            } else {
                for fi10 := 0; fi10 < 4; fi10++ {
                    for fi11 := 0; fi11 < 16; fi11++ {
                        var v12 uint = uint(v6)
                        var v13 int32 = 891
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 1, F1: -364, F2: v4}
                    var v15 uint64 = v4
                    _ = fi10
                    _ = v14
                    _ = v15
                }
                if v3 {
                    if v3 {
                        var v16 uint8 = (v6 * v6)
                        v16 = (v16 & 128)
                        _ = v16
                    }
                }
            }
        }
        var v17 uint = 69
        _ = len(v1)
        _ = v17
    }
    var v18 int = 376
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := v5[0]
        fmt.Println(fi19, v6, v3)
        for fi21 := 0; fi21 < 5; fi21++ {
            v4 -= 18446744073709551615
            v4 = (^(v4))
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    v18 = v18
    v22 := GenS1{F0: 32767, F1: -32768, F2: (v4 - 18446744073709551615)}
    var v23 string = "foo bar"
    var v24 float32 = float32(1e10)
    var v25 int8 = 100
    var v26 int8 = v25
    v1 = append(v1, ((((("\n\t" + (v23 + ((v23 + (v23 + "\n\t")) + v23))) + (v23 + v23)) + v23) + v23) + v23))
    if (4294967295 != 1) {
        v3 = v3
    }
    var v27 float32 = (v24 / (((((0.0 - float32(3.14159)) * v24) * 0.0) / 1.5) - v24))
    v28 := GenS1{F0: 32767, F1: 1000, F2: (^(v4))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
