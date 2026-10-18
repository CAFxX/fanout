package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool) uint {
    v6 := []int{32767, 931, 42, 32768}
    switch ("foo bar" + "hello") {
    case "":
        p0 = (false && (2 == 162))
    }
    _ = p0
    _ = v6
    return 109
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = 2147483647
    v3 := make(map[string]string)
    var v4 float32 = 1e10
    v5 := make(map[string]string)
    for fi7 := 0; fi7 < 7; fi7++ {
        if true {
            v8 := GenS1{F0: 9223372036854775807, F1: 1000, F2: ("a" + ((("hello" + "a") + "hello") + ("foo bar" + (("" + "a") + "a"))))}
            var v9 bool = true
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    v10 := []int{-920, 32767, 2147483647, 758}
    var v11 string = "a"
    v4 = v4
    v12 := GenS2{F0: -1}
    v13 := GenS3{F0: 22}
    v4 -= v4
    switch 602 {
    case 32767:
        v2 -= 2147483647
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
