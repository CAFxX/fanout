package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int64 }

func main() {
    var v1 float64 = 0.0
    var v2 int16 = 40
    var v3 uint64 = 183
    var v4 uint16 = 0
    var v5 uint8 = 255
    v6 := make(map[string]bool)
    for fi7 := 0; fi7 < 12; fi7++ {
        v2 += (-(v2))
        v8 := []int{fi7, 0, int(v5)}
        for fi9 := 0; fi9 < 12; fi9++ {
            switch 4294967295 {
            case 0:
                v4 -= 32767
            }
            switch fi9 {
            case 4294967295:
                switch "foo bar" {
                case "":
                    var v10 float64 = 1.5
                    _ = v10
                case "foo bar":
                    var v11 uint32 = 874
                    _ = v11
                case "\n\t":
                    if (!(!(329 >= 1))) {
                        v5 += uint8(v3)
                    } else {
                        var v12 int64 = -323
                        _ = v12
                    }
                default:
                    var v13 uint16 = (v4 - (^(v4)))
                    _ = v13
                }
                switch "hello" {
                case "foo bar":
                    var v14 int = fi9
                    _ = v14
                case "":
                    v15 := GenS1{F0: "a"}
                    _ = v15
                case "a":
                    fmt.Println(v4, v4)
                }
            case -879:
                v5 = (v5 >> ((4294967295) & 63))
            }
            v16 := v6["\n\t"]
            _ = fi9
            _ = v16
        }
        _ = fi7
        _ = v8
    }
    v17 := []int{398, 9223372036854775807, -456}
    v18 := GenS2{F0: 2147483647, F1: 42}
    v5 -= v5
    var v19 int64 = 1
    v20 := GenS1{F0: "hello"}
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 uint32 = 39
        _ = fi21
        _ = v22
    }
    var v23 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
