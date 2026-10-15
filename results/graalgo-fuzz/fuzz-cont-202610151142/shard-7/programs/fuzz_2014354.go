package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint }
func genFn0(p0 uint64) uint {
    p0 = 0
    fmt.Println(p0)
    _ = p0
    return 255
}

func main() {
    var v1 float32 = (0.1 + 0.1)
    v2 := []uint32{0, 65535, 2, 826}
    var v3 uint16 = 0
    var v4 float32 = v1
    v5 := make(map[string]int)
    for fi6 := 0; fi6 < 2; fi6++ {
        v3 += (v3 >> ((427) & 63))
        _ = fi6
    }
    v1 -= 1.5
    v7 := GenS1{F0: 2147483648, F1: (0 < -739), F2: 453}
    v3 = 186
    fmt.Println(v1, v3)
    v8 := GenS1{F0: 65535, F1: (42 != 930), F2: 4294967295}
    v9 := GenS2{F0: 432}
    if (308 < -32768) {
        if len(v2) > 0 {
            v2[0] = 637
        }
        v1 += v1
        switch (("hello" + "hello") + "hello") {
        case "foo bar":
            var v10 uint64 = uint64(v3)
            _ = v10
        case "hello":
            if (2 == 0) {
                v11 := GenS3{F0: true, F1: (v3 & 2), F2: 2}
                _ = v11
            } else {
                var v12 uint64 = 4294967295
                _ = v12
            }
        }
    }
    var v13 uint8 = 127
    var v14 string = (((((("hello" + ("" + ("" + (("a" + "") + "")))) + "foo bar") + "") + ("hello" + ("a" + ""))) + "a") + "hello")
    fmt.Println(v14, v1, v3)
    var v15 string = v14
    v5[v15] = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
