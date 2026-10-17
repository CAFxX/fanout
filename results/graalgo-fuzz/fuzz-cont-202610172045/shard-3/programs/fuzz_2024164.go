package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 bool }
type GenS2 struct { F0 uint64; F1 string; F2 uint16 }
type GenS3 struct { F0 uint; F1 int32 }
func genFn0() uint {
    var v6 uint64 = 874
    v6 -= 890
    _ = v6
    return 42
}

func main() {
    var v1 string = (("hello" + (("foo bar" + "foo bar") + "a")) + "hello")
    var v2 int = 42
    v3 := []int{(v2 * (v2 + v2)), 2147483647}
    var v4 uint64 = 18446744073709551615
    var v5 bool = false
    var v7 int16 = -32768
    if ((0.0 * (-2.25 * 0.0)) >= (-2.25 + 3.14159)) {
        v4 -= v4
        if v5 {
            v1 = v1
        } else {
            var v8 float32 = 0.1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v9 := GenS1{F0: 1, F1: 2147483648, F2: v5}
            _ = v8
            _ = v9
        }
    } else {
        var v10 float64 = 1.5
        fmt.Println(v2)
        _ = v10
    }
    if (!((-(v2)) == v2)) {
        var v11 int8 = 4
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 int = v2
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            v3 = append(v3, 246)
            v15 := GenS1{F0: 1, F1: 42, F2: (1 < 214)}
            var v16 uint8 = 42
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v11
    }
    v3 = append(v3, v2)
    fmt.Println(v7, v1)
    var v17 int = (v2 | v2)
    v4 = v4
    v4 = (v4 << ((319) & 63))
    v18 := GenS1{F0: 0, F1: 2147483647, F2: v5}
    v19 := GenS2{F0: (v4 ^ (^(v4))), F1: v1, F2: 42}
    var v20 uint32 = 1
    v5 = v5
    v1 = ("hello" + ("\n\t" + ((("a" + v1) + "foo bar") + "\n\t")))
    if (769 < 24) {
        v21 := GenS3{F0: 181, F1: -446}
        if (v5 || (!v5)) {
            for fi22 := 0; fi22 < 18; fi22++ {
                if ((true || v5) || ((((1.5 / float32(3.14159)) / -2.25) - (float32((1.5 / 1e10)) + float32(float64(1.5)))) <= ((0.0 + (((3.14159 + 3.14159) - 0.1) - (0.1 + 0.1))) + ((0.1 / float32(-2.25)) * (-2.25 * -2.25))))) {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        v5 = false
                        v24 := GenS3{F0: 1, F1: -12}
                        fmt.Println(v17, v5)
                        _ = fi23
                        _ = v24
                    }
                    var v25 float32 = ((0.1 - (float32(1e10) / ((3.14159 - 0.1) - 1.5))) - float32(0.0))
                    v26 := GenS3{F0: 255, F1: -1}
                    _ = v25
                    _ = v26
                }
                fmt.Println(v7, v17, v4)
                _ = fi22
            }
        }
        var v27 string = (((v1 + v1) + "a") + "a")
        _ = v21
        _ = v27
    } else {
        switch (v17 >> ((42) & 63)) {
        case 32768:
            var v28 bool = (v7 >= v7)
            _ = v28
        case -2147483648:
            var v29 uint16 = 43
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
