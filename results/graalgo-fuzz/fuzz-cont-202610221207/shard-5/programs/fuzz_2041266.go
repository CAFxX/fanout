package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -1
    var v3 float32 = 3.14159
    var v4 int16 = -819
    for fi5 := 0; fi5 < 4; fi5++ {
        switch fi5 {
        case 476:
            v2 -= (v2 << ((1) & 63))
        case 32768:
            if ((v3 / v3) == ((((v3 / 0.1) * v3) + v3) / 3.14159)) {
                for fi6 := 0; fi6 < 14; fi6++ {
                    var v7 uint8 = 255
                    var v8 string = "hello"
                    if (true || true) {
                        v9 := []int{(fi6 * fi5), (fi6 / ((int(v4)) | 1)), (fi6 * (fi6 + (fi6 % (((fi6 << ((2) & 63))) | 1)))), 685}
                        _ = v9
                    } else {
                        v10 := GenS1{F0: 1, F1: 1, F2: v7}
                        _ = v10
                    }
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                if (false || false) {
                    delete(v1, (^(fi5)))
                    if true {
                        v11 := []int{4294967295}
                        _ = v11
                    } else {
                        fmt.Println(v3, v3, v3)
                    }
                }
                v3 -= (float32(0.0) * (v3 + (3.14159 + 1.5)))
            } else {
                v4 -= v4
            }
        }
        var v12 uint = 332
        _ = fi5
        _ = v12
    }
    v13 := []int{32768, 358}
    v14 := GenS2{F0: 42, F1: v2}
    v3 -= v3
    var v15 float32 = (1e10 + 0.1)
    var v16 float32 = v3
    switch ((("\n\t" + "foo bar") + ("a" + "")) + ("foo bar" + "hello")) {
    case "foo bar":
        v16 += (v3 + (v3 - (0.0 + (v15 / v3))))
    default:
        v17 := GenS2{F0: 2147483648, F1: v2}
        _ = v17
    }
    v4 = v4
    if (!(2 <= 42)) {
        v1[65] = 2
    }
    v18 := GenS1{F0: 1, F1: 2147483648, F2: 1}
    v19 := GenS2{F0: 2, F1: (v2 >> ((65535) & 63))}
    var v20 bool = (!true)
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
