package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int8 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 0}
        _ = fi6
        _ = v7
    }
    var v8 uint = 1
    v8 = v8
    v9 := GenS1{F0: -576}
    v10 := GenS2{F0: 42, F1: -80}
    _ = v10
    _ = v8
    _ = v9
    return 128
}
func genFn1() bool {
    for fi11 := 0; fi11 < 4; fi11++ {
        fmt.Println(fi11)
        var v12 int64 = 828
        _ = fi11
        _ = v12
    }
    var v13 float32 = 3.14159
    _ = v13
    return true
}

func main() {
    v1 := []int{42}
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    var v4 int = -2147483648
    var v5 int16 = 0
    fmt.Println(v5, v4)
    v14 := GenS3{F0: 1}
    var v15 float32 = 0.1
    for fi16 := 0; fi16 < 16; fi16++ {
        _ = len(v1)
        _ = fi16
    }
    fmt.Println(v4, v4)
    if false {
        switch "\n\t" {
        case "":
            var v17 uint32 = 2147483648
            _ = v17
        case "foo bar":
            var v18 int64 = 147
            _ = v18
        }
    }
    v15 = (0.1 * (v15 * 1e10))
    for fi19 := 0; fi19 < 6; fi19++ {
        if (42 == 1) {
            var v20 int64 = -399
            v5 = int16(v20)
            v21 := v2[((((("foo bar" + ("foo bar" + "a")) + "a") + "a") + ("a" + "hello")) + ("" + ((((("foo bar" + ("a" + "\n\t")) + "") + ("foo bar" + ("" + "a"))) + "") + "a")))]
            _ = v20
            _ = v21
        } else {
            v4 = 32767
            var v22 uint32 = 2
            _ = v22
        }
        _ = fi19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
