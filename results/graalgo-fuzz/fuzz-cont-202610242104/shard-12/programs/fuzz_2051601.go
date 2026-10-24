package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 uint8 = 180
    v2 := []string{("" + "hello"), (("\n\t" + "a") + "\n\t")}
    var v3 uint16 = 630
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]string)
    switch "foo bar" {
    case "foo bar":
        switch ("hello" + ("a" + "\n\t")) {
        case "foo bar":
            var v6 bool = false
            _ = v6
        default:
            if (1000 == 859) {
                var v7 int32 = 2147483647
                v7 += (-(v7))
                var v8 int64 = int64(v3)
                _ = v7
                _ = v8
            }
        }
    }
    var v9 int64 = 42
    for fi10 := 0; fi10 < 12; fi10++ {
        switch -278 {
        case 1:
            v1 = v1
        }
        v3 = (v3 | (v3 * (v3 + v3)))
        var v11 int8 = -1
        _ = fi10
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 uint32 = 1
            var v14 uint8 = (v1 * 94)
            fmt.Println(fi12, fi12, v3)
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS1{F0: 42, F1: 1}
        _ = v15
    }
    v16 := GenS1{F0: 42, F1: 612}
    v2 = append(v2, "hello")
    v17 := v5[2147483648]
    for fi18 := 0; fi18 < 3; fi18++ {
        fmt.Println(v17, v3)
        _ = fi18
    }
    fmt.Println(v3, v1)
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v19 := GenS2{F0: 1, F1: v3}
    fmt.Println(v17)
    v9 -= (v9 + (v9 & v9))
    fmt.Println(v9)
    v20 := GenS2{F0: 1, F1: (v3 | v3)}
    v2 = append(v2, v17)
    for fi21 := 0; fi21 < 12; fi21++ {
        switch (((((v17 + ((v17 + ("hello" + "\n\t")) + (("\n\t" + (v17 + (((v17 + v17) + v17) + v17))) + v17))) + "") + (v17 + "")) + "foo bar") + v17) {
        case "":
            var v22 float64 = ((0.1 * (0.1 / 3.14159)) * 1.5)
            _ = v22
        case "a":
            fmt.Println(v1, v1)
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi21
    }
    var v23 int64 = (v9 & (v9 * v9))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
