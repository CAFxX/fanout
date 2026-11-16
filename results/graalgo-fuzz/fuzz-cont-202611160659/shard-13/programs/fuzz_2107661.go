package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint8) uint64 {
    var v5 uint16 = 32767
    if false {
        if (!((false || false) && false)) {
            v6 := GenS1{F0: 42}
            p0 = 128
            _ = v6
        }
        v7 := []int{128, 2147483647, -700}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int32 = 919
            v5 -= v5
            _ = fi8
            _ = v9
        }
        var v10 int32 = 42
        if (!(!false)) {
            v5 = (v5 | v5)
            v5 += 992
        }
        _ = v10
    }
    p0 += 128
    v5 -= v5
    var v11 int16 = -32768
    _ = p0
    _ = v11
    _ = v5
    return 4294967295
}

func main() {
    v1 := []string{"a", "a"}
    v2 := make(map[string]uint32)
    var v3 string = "a"
    var v4 int64 = -1
    if true {
        v1 = append(v1, (v3 + (v3 + v3)))
        v3 = ("" + v3)
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS2{F0: 532, F1: (("" + "\n\t") + (("" + v3) + (v3 + "a")))}
            _ = fi12
            _ = v13
        }
    }
    if (((!true) && (0 > 1)) && (!(-87 == 73))) {
        v14 := GenS3{F0: 4294967295}
        var v15 bool = false
        var v16 uint8 = 255
        _ = v14
        _ = v15
        _ = v16
    } else {
        delete(v2, v3)
        var v17 uint16 = 1
        _ = v17
    }
    if false {
        var v18 float32 = (float32(-2.25) * (0.0 * -2.25))
        var v19 string = v3
        _ = v18
        _ = v19
    }
    var v20 uint64 = 809
    v3 = ("" + "\n\t")
    v4 = -405
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 uint = 255
        _ = fi21
        _ = v22
    }
    v2[(v3 + ((("" + ((v3 + "hello") + v3)) + v3) + v3))] = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
