package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint8, p1 uint32, p2 uint16) uint64 {
    if ((-467 >= -1) && ("a" != "hello")) {
        for fi5 := 0; fi5 < 7; fi5++ {
            p0 = 255
            p2 += (p2 >> ((840) & 63))
            _ = fi5
        }
    }
    v6 := []int{-9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    var v1 int = -549
    var v2 string = "\n\t"
    var v3 uint64 = 4294967295
    v4 := []int{v1}
    v3 -= 2
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 0, F1: 4294967295}
        if (-2.25 != float32(1.5)) {
            var v9 int8 = 1
            _ = v9
        } else {
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS2{F0: 42}
                _ = len(v4)
                _ = fi10
                _ = v11
            }
            if len(v4) > 0 {
                v4[0] = v1
            }
        }
        if ((v2 + ("" + ((v2 + "") + v2))) < (("a" + v2) + v2)) {
            var v12 int8 = 0
            _ = v12
        } else {
            if len(v4) > 0 {
                v4[0] = v1
            }
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v1, v2)
    var v13 bool = true
    v14 := GenS3{F0: 1}
    v2 = v2
    var v15 string = ((("a" + ("foo bar" + ("foo bar" + v2))) + (v2 + "hello")) + ((v2 + v2) + v2))
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS1{F0: -564, F1: 1}
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
