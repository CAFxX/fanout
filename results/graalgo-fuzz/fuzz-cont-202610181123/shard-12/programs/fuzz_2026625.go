package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int32; F2 int64 }
type GenS3 struct { F0 int32; F1 int32; F2 uint32 }
func genFn0(p0 uint32) int64 {
    p0 = p0
    if (!false) {
        v6 := GenS1{F0: -9223372036854775808}
        _ = v6
    }
    var v7 float32 = (0.1 + 1.5)
    v8 := GenS1{F0: -9223372036854775808}
    v9 := []int{}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -685
}
func genFn1(p0 int32, p1 uint16, p2 bool) bool {
    var v10 uint = 4294967295
    v11 := GenS1{F0: genFn0(2)}
    v10 = v10
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return (p2 || p2)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 0.1
    v3 := []uint32{2147483648, 2, 4294967295}
    v4 := make(map[string]int)
    var v5 uint16 = 62
    for fi12 := 0; fi12 < 2; fi12++ {
        v5 -= 417
        var v13 uint = 167
        v14 := GenS1{F0: -374}
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v5, v5)
    v15 := GenS1{F0: 9223372036854775807}
    v4["hello"] = 32768
    v16 := GenS2{F0: ("foo bar" + ("\n\t" + "foo bar")), F1: 42, F2: 1}
    for fi17 := 0; fi17 < 14; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            _ = len(v3)
            _ = fi18
        }
        if false {
            if (!((("hello" + "foo bar") >= ("foo bar" + ((((("a" + ("\n\t" + "\n\t")) + "\n\t") + "hello") + "a") + "\n\t"))) && (42 > 1000))) {
                v5 += (^(v5))
                var v19 string = ("a" + "a")
                _ = v19
            } else {
                v20 := GenS3{F0: 0, F1: 0, F2: 2147483647}
                _ = v20
            }
        } else {
            var v21 int32 = 2147483647
            _ = v21
        }
        _ = fi17
    }
    v2 = (1.5 - v2)
    var v22 float32 = (float32(-2.25) / float32(v2))
    var v23 uint32 = 2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
