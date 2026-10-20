package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint16, p1 bool, p2 uint64) int64 {
    v5 := []int{42, 2147483647, 127, 1}
    switch 0 {
    case 32768:
        if (!(4294967295 < 2)) {
            if len(v5) > 0 {
                v5[0] = 9223372036854775807
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi6 := 0; fi6 < 16; fi6++ {
                v7 := GenS1{F0: 0}
                _ = fi6
                _ = v7
            }
        }
    case 128:
        fmt.Println(p2, p0)
    case 32767:
        var v8 uint = uint(p0)
        _ = v8
    default:
        p1 = (p0 == p0)
    }
    v9 := GenS1{F0: 1}
    if (("hello" + ("hello" + "")) >= "") {
        p2 -= (^(p2))
        fmt.Println(p2, p0)
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        var v11 bool = false
        if false {
            fmt.Println(v11, p2, p0)
        }
        fmt.Println(p0, v11, fi10)
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return -1
}
func genFn1(p0 uint64) uint {
    v12 := []int{}
    var v13 uint32 = 2147483648
    for fi14 := 0; fi14 < 13; fi14++ {
        if (1.5 > (-2.25 - 1e10)) {
            p0 += (p0 / (((p0 | 494)) | 1))
        }
        _ = fi14
    }
    v12 = append(v12, 42)
    _ = p0
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 2
    v3 := []uint32{760, 0, 65535}
    v4 := []string{("foo bar" + "foo bar"), (("a" + "a") + (("hello" + "foo bar") + "a"))}
    delete(v1, "foo bar")
    var v15 int32 = -229
    v4 = append(v4, ((("foo bar" + ("" + ("foo bar" + (("hello" + "hello") + "a")))) + "\n\t") + "hello"))
    fmt.Println(v15, v15)
    var v16 int = -845
    v17 := GenS2{F0: 128}
    v18 := v1["foo bar"]
    var v19 float32 = ((float32(1.5) / (1.5 - (3.14159 / 1.0))) + 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
