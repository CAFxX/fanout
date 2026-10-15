package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 string; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 string; F1 uint }
func genFn0(p0 uint16) uint64 {
    var v7 float64 = 3.14159
    var v8 int16 = 42
    var v9 int8 = 42
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 960
}

func main() {
    var v1 uint8 = 42
    var v2 int16 = 1
    var v3 uint64 = uint64(v1)
    var v4 uint16 = 32767
    var v5 int8 = 100
    v6 := []uint32{2147483647, 2, 65535, 0}
    if len(v6) > 0 {
        v6[0] = 351
    }
    if false {
        var v10 float32 = (1.5 / (3.14159 - float32(-2.25)))
        if (int8(v3) <= (v5 | (v5 ^ 83))) {
            _ = len(v6)
            if true {
                var v11 uint64 = (v3 & v3)
                _ = v11
            }
            fmt.Println(v4, v4, v3)
        }
        _ = v10
    } else {
        _ = len(v6)
        v5 -= v5
        v1 = 42
    }
    var v12 float32 = -2.25
    v13 := GenS1{F0: (v2 / (((v2 + v2)) | 1))}
    if ("hello" <= "a") {
        v14 := GenS2{F0: 4294967295, F1: "\n\t", F2: 128}
        _ = v14
    }
    switch ("hello" + "\n\t") {
    case "hello":
        if false {
            for fi15 := 0; fi15 < 15; fi15++ {
                var v16 int = (fi15 | fi15)
                _ = fi15
                _ = v16
            }
            v3 -= (^(v3))
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            v12 -= (0.1 + (v12 - ((float32(0.0) / v12) - (v12 + 1e10))))
            if ((-1 != 0) && (2 > 42)) {
                var v17 int = -1
                _ = v17
            } else {
                v18 := GenS1{F0: v2}
                _ = v18
            }
        }
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS3{F0: v5, F1: 2}
        _ = fi19
        _ = v20
    }
    var v21 float64 = 1.5
    v22 := GenS4{F0: ((("hello" + ("hello" + ("foo bar" + "foo bar"))) + "\n\t") + "foo bar"), F1: 1}
    var v23 uint64 = genFn0(v4)
    var v24 bool = false
    v2 = 1
    v21 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
