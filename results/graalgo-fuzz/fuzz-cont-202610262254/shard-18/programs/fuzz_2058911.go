package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -128
    var v3 uint32 = 42
    v4 := make(map[int]int)
    var v5 int32 = 42
    var v6 float32 = 0.0
    v3 -= (v3 & (v3 - v3))
    var v7 int8 = v2
    v8 := []int{-377, 4294967295, -1, 164}
    switch 42 {
    case 42:
        var v9 int64 = 814
        _ = v9
    default:
        v6 = v6
    }
    var v10 uint64 = 436
    v8 = append(v8, 485)
    v7 = (v2 >> ((4294967295) & 63))
    v1[9223372036854775807] = "a"
    var v11 float32 = v6
    if (v6 < 1.5) {
        v12 := GenS1{F0: 128}
        _ = v12
    }
    var v13 float64 = float64(v11)
    var v14 uint8 = 42
    if (int32(v14) >= (v5 | v5)) {
        var v15 string = "a"
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        if true {
            var v17 string = ("" + ("" + "a"))
            v14 = 128
            _ = v17
        } else {
            if true {
                v1[fi16] = ("\n\t" + (((("a" + "\n\t") + "\n\t") + "a") + "foo bar"))
                var v18 uint32 = 599
                var v19 int = fi16
                _ = v18
                _ = v19
            }
            if false {
                if true {
                    v20 := GenS1{F0: 128}
                    v14 = (^(v14))
                    _ = v20
                } else {
                    v3 -= 0
                    var v21 uint16 = 42
                    _ = v21
                }
                _ = len(v4)
            } else {
                for fi22 := 0; fi22 < 7; fi22++ {
                    v23 := GenS1{F0: (v14 * (v14 >> ((207) & 63)))}
                    _ = fi22
                    _ = v23
                }
                v24 := v1[fi16]
                _ = v24
            }
            v13 = 0.1
        }
        _ = fi16
    }
    var v25 int64 = -1
    for fi26 := 0; fi26 < 3; fi26++ {
        switch (fi26 - fi26) {
        case 4294967295:
            v27 := v4[fi26]
            _ = v27
        }
        var v28 string = ("hello" + "a")
        v10 -= v10
        _ = fi26
        _ = v28
    }
    v7 -= 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
