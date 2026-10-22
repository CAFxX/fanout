package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint32; F1 string; F2 int8 }
type GenS3 struct { F0 int; F1 uint16; F2 int32 }
type GenS4 struct { F0 int32; F1 uint32 }
func genFn0(p0 string, p1 string, p2 uint64) int16 {
    switch "a" {
    case "":
        p1 = "a"
    case "hello":
        var v7 int = 42
        _ = v7
    }
    p2 -= p2
    v8 := GenS1{F0: -554, F1: 74, F2: 2147483647}
    var v9 bool = false
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 float64 = 3.14159
        v9 = v9
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 int = fi10
            v13 -= (^(v13))
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 32767
}
func genFn1(p0 string) uint {
    v14 := GenS2{F0: 149, F1: (p0 + ""), F2: 7}
    v15 := []int{-2147483648, 32767, -9223372036854775808}
    p0 = p0
    _ = p0
    _ = v14
    _ = v15
    return 831
}

func main() {
    v1 := []string{""}
    var v2 uint = 65535
    v3 := []int{0, 127, 32768, 90}
    var v4 float32 = -2.25
    var v5 uint = (v2 >> ((v2) & 63))
    v6 := []string{("" + "\n\t"), "foo bar", ""}
    if (true && (!(!(!false)))) {
        v16 := GenS2{F0: 2147483648, F1: "\n\t", F2: 0}
        v5 = (v2 * (v2 + (v2 | v2)))
        _ = v16
    } else {
        if true {
            v3 = append(v3, 9223372036854775807)
        } else {
            v4 += v4
            if (!true) {
                v4 = (v4 / (0.1 + (v4 * v4)))
                fmt.Println(v2, v4, v4)
            }
            var v17 int16 = 0
            _ = v17
        }
    }
    v18 := GenS3{F0: 32768, F1: 42, F2: 198}
    fmt.Println(v5, v2)
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS1{F0: -516, F1: 1, F2: 608}
        _ = fi19
        _ = v20
    }
    var v21 uint16 = 32767
    v22 := GenS3{F0: 127, F1: v21, F2: 1}
    for fi23 := 0; fi23 < 19; fi23++ {
        v5 = 123
        _ = fi23
    }
    v24 := GenS4{F0: -1, F1: 65535}
    v1 = append(v1, ("" + (("a" + "\n\t") + "\n\t")))
    if (0 != -345) {
        v2 = v5
    }
    _ = len(v1)
    v3 = append(v3, -363)
    for fi25 := 0; fi25 < 14; fi25++ {
        fmt.Println(v2)
        if len(v3) > 0 {
            v3[0] = fi25
        }
        switch (("foo bar" + "a") + "") {
        case "\n\t":
            for fi26 := 0; fi26 < 2; fi26++ {
                _ = len(v1)
                _ = fi26
            }
        }
        _ = fi25
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v2 -= v2
        _ = fi27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
