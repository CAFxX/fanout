package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 uint }
func genFn0(p0 string, p1 int8) uint16 {
    var v6 uint = 0
    p1 += p1
    _ = p0
    _ = p1
    _ = v6
    return 42
}
func genFn1() uint {
    v7 := []int{127}
    var v8 uint = 995
    var v9 uint32 = 615
    if (("foo bar" >= "hello") || (!true)) {
        if (float64(1e10) < (0.1 + float64(0.1))) {
            v10 := GenS1{F0: 127}
            _ = v10
        } else {
            v11 := GenS2{F0: 127}
            _ = v11
        }
    }
    _ = v7
    _ = v8
    _ = v9
    return 26
}

func main() {
    var v1 int = 32767
    var v2 string = ""
    v3 := make(map[string]int)
    v4 := []uint32{}
    v5 := make(map[string]string)
    switch v2 {
    case "foo bar":
        if (946 <= 1) {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 uint64 = 699
                _ = fi12
                _ = v13
            }
            switch ("" + (v2 + "a")) {
            case "":
                v14 := GenS3{F0: -76}
                _ = v14
            case "hello":
                if len(v4) > 0 {
                    v4[0] = 2
                }
            default:
                var v15 int16 = 228
                _ = v15
            }
        } else {
            var v16 uint = 4294967295
            _ = v16
        }
    default:
        for fi17 := 0; fi17 < 18; fi17++ {
            if len(v4) > 0 {
                v4[0] = 2147483648
            }
            v18 := GenS4{F0: 871, F1: 4294967295}
            _ = fi17
            _ = v18
        }
    }
    var v19 int32 = 42
    v2 = v2
    var v20 int64 = 0
    v2 = v2
    fmt.Println(v2, v1)
    v21 := GenS5{F0: 65535}
    v1 -= (v1 >> ((18446744073709551615) & 63))
    _ = len(v3)
    var v22 float64 = (1.5 - (1e10 + (-2.25 * 1.5)))
    var v23 uint = 692
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
