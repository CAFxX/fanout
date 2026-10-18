package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -32768
    v3 := make(map[string]bool)
    var v4 uint = 1
    var v5 float64 = 0.0
    var v6 float32 = (3.14159 * -2.25)
    v7 := GenS1{F0: (^(v4)), F1: 18446744073709551615}
    fmt.Println(v6, v6)
    var v8 int32 = 42
    var v9 uint = 553
    v2 += v2
    var v10 int32 = (v8 >> ((42) & 63))
    fmt.Println(v8, v10, v6)
    fmt.Println(v4, v9)
    v2 += v2
    fmt.Println(v2)
    fmt.Println(v6, v5)
    for fi11 := 0; fi11 < 3; fi11++ {
        if false {
            v12 := []int{}
            _ = v12
        }
        var v13 int32 = v10
        if false {
            for fi14 := 0; fi14 < 16; fi14++ {
                v2 -= 1000
                v13 = (v8 << ((v9) & 63))
                v15 := GenS1{F0: v9, F1: 18446744073709551615}
                _ = fi14
                _ = v15
            }
            for fi16 := 0; fi16 < 13; fi16++ {
                switch ("hello" + (("a" + "a") + "\n\t")) {
                case "":
                    var v17 int = fi16
                    v8 = 0
                    _ = v17
                case "foo bar":
                    fmt.Println(fi11)
                default:
                    var v18 uint = 255
                    _ = v18
                }
                _ = fi16
            }
            delete(v3, "\n\t")
        }
        _ = fi11
        _ = v13
    }
    v19 := []int{979}
    v20 := GenS2{F0: ""}
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
