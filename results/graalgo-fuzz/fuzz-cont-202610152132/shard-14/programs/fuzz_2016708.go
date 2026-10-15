package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int, p1 uint, p2 uint8) int8 {
    v4 := []int{(p0 >> (((^(p1))) & 63)), (p0 + -1)}
    var v5 uint64 = 240
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -29
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 128
    v3 := make(map[string]int)
    var v6 float64 = (float64(1e10) - 0.0)
    v6 -= v6
    var v7 int64 = -9223372036854775808
    v7 = -619
    fmt.Println(v6)
    switch "a" {
    case "\n\t":
        v8 := GenS1{F0: 42, F1: 32767, F2: 32767}
        _ = v8
    case "":
        v9 := GenS2{F0: 892, F1: 932, F2: 524}
        _ = v9
    }
    v10 := GenS3{F0: 877}
    switch -1 {
    case 9223372036854775807:
        for fi11 := 0; fi11 < 19; fi11++ {
            switch ("a" + ((("foo bar" + "") + ("foo bar" + "")) + "a")) {
            case "hello":
                v12 := GenS1{F0: 578, F1: 32767, F2: (fi11 / ((fi11) | 1))}
                _ = v12
            case "a":
                var v13 float64 = ((1e10 + v6) - float64(((1e10 + 0.1) - 1e10)))
                _ = v13
            case "":
                v14 := GenS4{F0: 473}
                _ = v14
            }
            var v15 uint = 255
            _ = fi11
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        if (65535 < 65535) {
            v6 = 1.5
            v6 = (v6 / float64(3.14159))
        }
        _ = fi16
    }
    v6 -= (float64((1.5 / ((-2.25 + (3.14159 * 1.5)) + 3.14159))) / float64(3.14159))
    v6 -= 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
