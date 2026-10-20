package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 bool) string {
    var v7 uint32 = 1
    switch ("" + "a") {
    case "foo bar":
        var v8 uint8 = 1
        _ = v8
    }
    v9 := GenS1{F0: -2147483648}
    _ = p0
    _ = v7
    _ = v9
    return "foo bar"
}

func main() {
    var v1 int64 = -189
    var v2 float64 = 0.0
    var v3 uint32 = 4294967295
    var v4 int8 = 42
    v5 := make(map[string]string)
    var v6 int8 = -128
    v4 -= (v4 * v4)
    fmt.Println(v2, v4, v4)
    var v10 int = 29
    if (1 >= 127) {
        v4 -= 0
        if false {
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := GenS2{F0: -953, F1: (v10 + fi11)}
                fmt.Println(v1, v6, fi11)
                _ = fi11
                _ = v12
            }
            switch ("hello" + ("" + "a")) {
            case "":
                v13 := GenS3{F0: 247, F1: 1}
                _ = v13
            }
        }
    }
    v14 := GenS4{F0: uint16(v3)}
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := v5["a"]
        var v17 uint32 = v3
        v18 := GenS4{F0: 65535}
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 float64 = float64(3.14159)
    var v20 int64 = v1
    v21 := GenS3{F0: 127, F1: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
