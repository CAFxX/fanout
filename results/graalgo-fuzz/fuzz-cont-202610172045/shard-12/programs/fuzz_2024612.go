package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 int; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int) int8 {
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            p0 -= (fi8 << ((0) & 63))
            var v9 uint8 = 127
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    p0 = p0
    var v10 int64 = -9223372036854775808
    _ = p0
    _ = v10
    return 123
}

func main() {
    var v1 uint32 = 65535
    var v2 string = "foo bar"
    var v3 int8 = 0
    v4 := make(map[string]string)
    v5 := []string{("hello" + v2)}
    var v6 uint64 = 4294967295
    fmt.Println(v6)
    var v11 uint = 255
    fmt.Println(v2)
    if false {
        switch 1 {
        case 128:
            if (!false) {
                var v12 int8 = v3
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            v14 := GenS1{F0: v3}
            _ = fi13
            _ = v14
        }
    }
    v15 := GenS1{F0: v3}
    v16 := GenS2{F0: (^(v6)), F1: 127, F2: 0}
    var v17 float32 = 0.1
    v18 := GenS3{F0: v1, F1: 1}
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 bool = (42 <= -771)
        _ = fi19
        _ = v20
    }
    var v21 uint64 = v6
    delete(v4, "a")
    var v22 int = int(v6)
    for fi23 := 0; fi23 < 4; fi23++ {
        switch ((v2 + (v2 + v2)) + ("" + v2)) {
        case "":
            var v24 uint16 = 255
            _ = v24
        case "foo bar":
            var v25 uint8 = 1
            _ = v25
        default:
            var v26 string = "a"
            _ = v26
        }
        _ = fi23
    }
    switch ("\n\t" + v2) {
    case "\n\t":
        v5 = append(v5, "")
    default:
        var v27 int = (v22 | v22)
        _ = v27
    }
    var v28 int8 = (v3 * v3)
    var v29 bool = false
    if v29 {
        var v30 int = v22
        switch v2 {
        case "foo bar":
            fmt.Println(v11)
        }
        v31 := GenS4{F0: 1}
        _ = v30
        _ = v31
    }
    var v32 int8 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
