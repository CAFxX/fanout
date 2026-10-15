package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint }
func genFn0() int8 {
    v7 := GenS1{F0: 2147483648, F1: 164}
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := GenS1{F0: 2147483647, F1: 0}
        fmt.Println(fi8)
        _ = fi8
        _ = v9
    }
    if false {
        var v10 string = ""
        v11 := GenS1{F0: 172, F1: 42}
        _ = v10
        _ = v11
    } else {
        v12 := GenS2{F0: 127}
        _ = v12
    }
    _ = v7
    return 48
}
func genFn1(p0 uint8) uint32 {
    v13 := GenS1{F0: 790, F1: 949}
    v14 := GenS3{F0: ("foo bar" + ""), F1: 4294967295}
    _ = p0
    _ = v13
    _ = v14
    return 2147483647
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"hello", "\n\t", ("" + (("hello" + ("" + "hello")) + "foo bar")), "a"}
    var v3 int64 = -9223372036854775808
    var v4 uint = 42
    v5 := make(map[uint32]int)
    var v6 int = 1
    v6 += 4294967295
    v4 += (^(v4))
    var v15 uint8 = 128
    fmt.Println(v3)
    if (!true) {
        var v16 bool = false
        v15 += 255
        _ = v16
    }
    var v17 string = "foo bar"
    v4 -= 255
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := GenS2{F0: genFn0()}
        switch v17 {
        case "hello":
            var v20 int8 = 75
            _ = v20
        case "":
            v21 := GenS2{F0: -1}
            _ = v21
        default:
            v22 := v1[uint32(v4)]
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    v4 = 889
    _ = len(v2)
    v2 = append(v2, "hello")
    v15 += v15
    var v23 int = v6
    for fi24 := 0; fi24 < 10; fi24++ {
        v4 -= 1
        for fi25 := 0; fi25 < 1; fi25++ {
            if false {
                if ((1 > 2147483647) && (!false)) {
                    v3 = -9223372036854775808
                } else {
                    var v26 float32 = 0.0
                    _ = v26
                }
            } else {
                if true {
                    var v27 int16 = -523
                    v6 = 181
                    _ = v27
                } else {
                    var v28 uint32 = uint32(v4)
                    var v29 uint8 = v15
                    var v30 int16 = 42
                    _ = v28
                    _ = v29
                    _ = v30
                }
                v31 := GenS2{F0: 0}
                _ = len(v2)
                _ = v31
            }
            v3 -= (v3 & 9223372036854775807)
            var v32 uint32 = 65535
            var v33 uint32 = v32
            fmt.Println(v15)
            v23 = (v23 & v6)
            _ = fi25
            _ = v32
            _ = v33
        }
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
