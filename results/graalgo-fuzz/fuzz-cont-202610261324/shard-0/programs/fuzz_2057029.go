package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint16; F1 int; F2 uint16 }
type GenS3 struct { F0 bool; F1 string; F2 uint }
func genFn0(p0 string, p1 int32) int16 {
    var v4 int = 9223372036854775807
    var v5 uint8 = 255
    if (255 < 255) {
        switch ((p0 + p0) + "a") {
        case "\n\t":
            var v6 int8 = 6
            _ = v6
        case "foo bar":
            if false {
                p1 = -1
            }
        }
    }
    v7 := GenS1{F0: -678, F1: true}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v7
    return -229
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"a", (("" + "hello") + "foo bar"), "hello", "a"}
    v3 := make(map[int]bool)
    var v8 uint32 = 469
    switch "a" {
    case "foo bar":
        v9 := GenS2{F0: 2, F1: 4294967295, F2: 32767}
        _ = v9
    case "\n\t":
        v10 := GenS2{F0: 65535, F1: 32768, F2: 255}
        _ = v10
    }
    v11 := GenS2{F0: 984, F1: 503, F2: 32767}
    if true {
        if false {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
        } else {
            var v12 string = "\n\t"
            var v13 string = (((v12 + ("\n\t" + v12)) + (v12 + ((((v12 + v12) + "a") + v12) + v12))) + "hello")
            var v14 float64 = (float64(1.5) - 1e10)
            _ = v12
            _ = v13
            _ = v14
        }
        _ = len(v3)
    }
    delete(v3, 0)
    fmt.Println(v8)
    v8 -= 789
    fmt.Println(v8)
    v8 += v8
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS3{F0: (true || false), F1: "hello", F2: 475}
        v8 += (v8 * v8)
        _ = fi15
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v8
}
