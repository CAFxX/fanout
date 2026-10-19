package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint32 }
func genFn0(p0 int8, p1 uint64, p2 int64) string {
    if true {
        var v6 float32 = 3.14159
        _ = v6
    } else {
        fmt.Println(p1, p1)
        for fi7 := 0; fi7 < 14; fi7++ {
            if false {
                var v8 uint64 = p1
                _ = v8
            } else {
                for fi9 := 0; fi9 < 5; fi9++ {
                    p1 -= p1
                    var v10 uint16 = 1
                    fmt.Println(p0)
                    _ = fi9
                    _ = v10
                }
                var v11 uint32 = 65535
                if true {
                    switch ("a" + "a") {
                    case "foo bar":
                        var v12 int64 = (-(p2))
                        _ = v12
                    case "hello":
                        v13 := []int{fi7}
                        _ = v13
                    default:
                        p1 += (p1 >> ((21) & 63))
                    }
                }
                _ = v11
            }
            var v14 bool = false
            v15 := GenS1{F0: v14, F1: p0}
            _ = fi7
            _ = v14
            _ = v15
        }
        v16 := GenS1{F0: true, F1: int8(p2)}
        _ = v16
    }
    var v17 float32 = 1e10
    v18 := GenS2{F0: -667}
    var v19 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v19
    return ((("" + "foo bar") + ("a" + ("foo bar" + "hello"))) + "a")
}

func main() {
    var v1 int32 = -1
    var v2 uint = 2
    v3 := make(map[string]int)
    var v4 uint8 = 255
    v5 := make(map[string]uint32)
    if true {
        v20 := []int{}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 uint = 65535
            var v23 string = ("foo bar" + "a")
            if (!false) {
                _ = len(v3)
            }
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    v4 = v4
    fmt.Println(v1)
    v24 := GenS3{F0: (v1 | v1), F1: 32, F2: 316}
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 int8 = 1
        _ = fi25
        _ = v26
    }
    switch 4294967295 {
    case 500:
        v27 := []int{2147483647, 774}
        _ = v27
    case 32767:
        fmt.Println(v4, v1, v2)
    case 883:
        fmt.Println(v4)
    default:
        var v28 uint32 = 0
        _ = v28
    }
    var v29 int16 = 42
    var v30 uint16 = 1
    v31 := v3[("\n\t" + "")]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
