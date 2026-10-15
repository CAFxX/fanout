package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
func genFn0(p0 bool, p1 int16, p2 string) uint16 {
    p1 = 32767
    v6 := GenS1{F0: -765, F1: 2}
    p1 = (-(p1))
    switch p2 {
    case "foo bar":
        fmt.Println(p1, p2)
    default:
        p0 = p0
    }
    var v7 int64 = 183
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 32767
}
func genFn1(p0 int32) uint64 {
    var v8 int16 = -32768
    p0 = (p0 | p0)
    var v9 string = "a"
    v10 := GenS1{F0: -913, F1: 139}
    var v11 uint = 65535
    _ = p0
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 642
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    v3 := []uint32{2147483648, 42, 65535, 0}
    var v4 bool = (((1.5 * (float32(0.0) + 1.5)) > 0.0) || true)
    var v5 float32 = 1e10
    if (4294967295 < 2) {
        v4 = v4
    } else {
        var v12 uint8 = 0
        _ = len(v1)
        _ = v12
    }
    if v4 {
        v13 := GenS1{F0: 32767, F1: 2}
        if v4 {
            var v14 int32 = 2147483647
            _ = v14
        } else {
            var v15 int8 = -1
            var v16 uint8 = 255
            _ = v15
            _ = v16
        }
        _ = v13
    }
    v5 = v5
    if true {
        if v4 {
            if (v5 != ((v5 + ((v5 * 1.5) / 1.5)) / v5)) {
                for fi17 := 0; fi17 < 13; fi17++ {
                    var v18 int32 = 2147483647
                    _ = fi17
                    _ = v18
                }
                v5 -= v5
            }
        } else {
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 int = fi19
                _ = fi19
                _ = v20
            }
            switch "hello" {
            case "\n\t":
                for fi21 := 0; fi21 < 16; fi21++ {
                    fmt.Println(v4)
                    _ = fi21
                }
            case "":
                for fi22 := 0; fi22 < 20; fi22++ {
                    var v23 int32 = 42
                    _ = fi22
                    _ = v23
                }
            case "foo bar":
                v5 = v5
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = len(v2)
    }
    var v24 int32 = 0
    for fi25 := 0; fi25 < 19; fi25++ {
        v26 := GenS1{F0: (fi25 << ((0) & 63)), F1: 255}
        v3 = append(v3, 703)
        var v27 int16 = -32768
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        v29 := v1[0]
        _ = fi28
        _ = v29
    }
    var v30 int = 4294967295
    if (42 <= 65535) {
        var v31 string = "foo bar"
        _ = v31
    }
    var v32 uint8 = 2
    fmt.Println(v5)
    _ = len(v3)
    fmt.Println(v30, v30)
    v4 = (0 <= (v24 << ((463) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
}
