package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint8; F1 uint }
func genFn0() uint8 {
    v4 := []int{}
    var v5 int = 42
    _ = v4
    _ = v5
    return 255
}
func genFn1(p0 string) int64 {
    p0 = "\n\t"
    v6 := GenS1{F0: 1}
    fmt.Println(p0)
    v7 := GenS2{F0: ("\n\t" + "a"), F1: 2147483647, F2: 1}
    fmt.Println(p0)
    _ = p0
    _ = v6
    _ = v7
    return 881
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 978
    v3 := []int{552, 4294967295, -702}
    switch (("\n\t" + "a") + "hello") {
    case "hello":
        switch 42 {
        case 0:
            v2 = v2
        }
    }
    if (true && true) {
        var v8 int64 = -1
        _ = v8
    }
    if len(v3) > 0 {
        v3[0] = 651
    }
    v2 += v2
    var v9 int32 = v2
    for fi10 := 0; fi10 < 6; fi10++ {
        _ = len(v3)
        switch 515 {
        case 32767:
            _ = len(v1)
        case 0:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi10
    }
    fmt.Println(v2, v9)
    v2 = (v9 & (v2 * v9))
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS3{F0: 100}
        v13 := GenS4{F0: 105, F1: 2147483647}
        _ = fi11
        _ = v12
        _ = v13
    }
    if true {
        var v14 int = -1
        v9 -= (v9 << ((2) & 63))
        var v15 float32 = 3.14159
        _ = v14
        _ = v15
    }
    v16 := GenS5{F0: 307, F1: 9223372036854775807}
    if (-1 <= -705) {
        var v17 float64 = 1.5
        _ = v17
    }
    var v18 int8 = 42
    delete(v1, "")
    v2 += 2147483647
    v9 = v2
    v19 := GenS6{F0: 78, F1: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
