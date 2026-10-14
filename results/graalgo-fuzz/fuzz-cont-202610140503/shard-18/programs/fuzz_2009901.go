package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0() string {
    for fi6 := 0; fi6 < 17; fi6++ {
        var v7 int16 = 1
        _ = fi6
        _ = v7
    }
    v8 := []int{}
    _ = v8
    return "foo bar"
}

func main() {
    var v1 float32 = 0.0
    var v2 int32 = 0
    v3 := make(map[int]bool)
    var v4 int8 = 123
    var v5 uint = 42
    fmt.Println(v2)
    fmt.Println(v4, v2, v1)
    switch -1 {
    case -9223372036854775808:
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 string = (("\n\t" + "") + "hello")
            v11 := []int{1, fi9}
            for fi12 := 0; fi12 < 13; fi12++ {
                var v13 float64 = -2.25
                var v14 uint = v5
                var v15 int8 = (v4 >> ((42) & 63))
                v3[fi12] = true
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    case 9223372036854775807:
        _ = len(v3)
    case 0:
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int32 = v2
            _ = fi16
            _ = v17
        }
    }
    switch "a" {
    case "a":
        v2 = int32(v4)
    case "hello":
        var v18 bool = (!true)
        _ = v18
    }
    var v19 int64 = 147
    fmt.Println(v2, v2)
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := GenS1{F0: 1}
        for fi22 := 0; fi22 < 2; fi22++ {
            switch fi22 {
            case 32768:
                v23 := []int{4294967295, int(v4)}
                _ = v23
            case -9223372036854775808:
                v24 := GenS2{F0: 4294967295, F1: 127}
                _ = v24
            case -1:
                var v25 float64 = 1.5
                _ = v25
            }
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    fmt.Println(v4, v1)
    v2 = (v2 >> ((v5) & 63))
    v26 := GenS2{F0: 1, F1: 127}
    v4 -= 42
    var v27 int8 = -1
    v28 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
