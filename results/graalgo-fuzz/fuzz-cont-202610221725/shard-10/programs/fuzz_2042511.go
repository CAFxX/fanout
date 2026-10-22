package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0(p0 int8) int {
    var v6 int16 = -738
    v7 := GenS1{F0: p0, F1: (v6 * (v6 << ((65535) & 63)))}
    v8 := GenS1{F0: p0, F1: v6}
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 string = "\n\t"
        if false {
            for fi11 := 0; fi11 < 13; fi11++ {
                v10 = ((v10 + "\n\t") + v10)
                var v12 bool = true
                v13 := []int{fi11, fi11, 42}
                _ = fi11
                _ = v12
                _ = v13
            }
            var v14 uint64 = 912
            _ = v14
        } else {
            var v15 int = fi9
            var v16 uint64 = 0
            v17 := []int{(-(fi9)), 4294967295}
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi9
        _ = v10
    }
    p0 = p0
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 32768
}
func genFn1() int64 {
    v18 := GenS1{F0: 1, F1: 32767}
    var v19 int32 = 0
    _ = v18
    _ = v19
    return 9223372036854775807
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]bool)
    var v3 int16 = -807
    var v4 string = ("" + "")
    v5 := []string{(v4 + v4), v4}
    var v20 int32 = 0
    _ = len(v2)
    var v21 int64 = 1
    v2[633] = false
    v22 := GenS2{F0: true}
    var v23 int8 = 1
    v24 := GenS3{F0: true, F1: 2}
    fmt.Println(v23, v4)
    var v25 uint8 = v1
    var v26 float64 = (1e10 / 1.0)
    var v27 string = (v4 + ((v4 + "") + ((("a" + "\n\t") + v4) + v4)))
    switch "" {
    case "foo bar":
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 int32 = v20
            var v30 float64 = v26
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 8; fi31++ {
        var v32 float64 = float64(1.5)
        var v33 float32 = -2.25
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 uint8 = v25
    var v35 int8 = (-(v23))
    var v36 int64 = (v21 + (v21 - 0))
    var v37 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
