package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }

func main() {
    var v1 bool = (!(1000 < 1))
    var v2 uint = 572
    v3 := make(map[string]int)
    var v4 uint64 = 2
    var v5 uint64 = (v4 >> ((v2) & 63))
    v5 = 4294967295
    if v1 {
        delete(v3, "foo bar")
        v6 := GenS1{F0: "a", F1: 1}
        var v7 uint16 = 2
        _ = v6
        _ = v7
    } else {
        v8 := []int{-498, 1, 42, 448}
        fmt.Println(v4, v5, v1)
        if v1 {
            if v1 {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 uint8 = 127
                    _ = fi9
                    _ = v10
                }
            } else {
                for fi11 := 0; fi11 < 14; fi11++ {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        if v1 {
                            switch "a" {
                            case "hello":
                                var v13 int = 9223372036854775807
                                _ = v13
                            default:
                                v2 -= (^(v2))
                            }
                            fmt.Println(fi12)
                        }
                        _ = fi12
                    }
                    switch "" {
                    case "hello":
                        v14 := GenS1{F0: "hello", F1: 42}
                        _ = v14
                    }
                    var v15 int32 = -2147483648
                    _ = fi11
                    _ = v15
                }
            }
        }
        _ = v8
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 int64 = 0
        _ = fi16
        _ = v17
    }
    v18 := v3[""]
    v19 := v3[(("a" + "foo bar") + "\n\t")]
    v20 := []int{}
    switch 2147483647 {
    case 128:
        var v21 bool = (v18 < v18)
        _ = v21
    case -451:
        v19 = (v18 - 42)
    case 127:
        var v22 string = ("" + ("" + ""))
        _ = v22
    default:
        var v23 int16 = -693
        _ = v23
    }
    var v24 bool = (653 > 295)
    v3["\n\t"] = int(v5)
    v1 = ((0 >= 9223372036854775807) || true)
    v25 := GenS1{F0: "", F1: 255}
    v24 = v24
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
