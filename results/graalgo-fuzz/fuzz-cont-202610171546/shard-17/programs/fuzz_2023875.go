package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 2
    var v3 float32 = 0.1
    var v4 uint32 = 2147483647
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            if true {
                var v7 int64 = 597
                _ = v7
            }
            var v8 float32 = 1e10
            _ = fi6
            _ = v8
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            if (!(915 > 0)) {
                v10 := []int{fi5, (fi5 - fi9), 203}
                _ = v10
            }
            v4 = uint32(v2)
            if (42 == 1) {
                v11 := GenS1{F0: ("\n\t" + "hello")}
                if false {
                    v12 := []int{(fi9 ^ fi5), fi5}
                    _ = v12
                }
                _ = v11
            } else {
                v13 := []int{(fi5 * -967), fi9, 9223372036854775807, fi5}
                fmt.Println(v4, v2)
                _ = v13
            }
            v14 := GenS1{F0: "foo bar"}
            var v15 int64 = 9223372036854775807
            _ = fi9
            _ = v14
            _ = v15
        }
        var v16 uint32 = v4
        _ = fi5
        _ = v16
    }
    var v17 uint = 42
    var v18 int16 = -84
    v19 := []int{}
    if (255 <= 255) {
        v17 = v17
    }
    if (true || (!(2 < 604))) {
        v17 += v17
    }
    var v20 string = "\n\t"
    var v21 int8 = -128
    v4 = 912
    _ = len(v19)
    v22 := GenS2{F0: 128, F1: 436}
    v23 := GenS1{F0: "\n\t"}
    var v24 int = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
