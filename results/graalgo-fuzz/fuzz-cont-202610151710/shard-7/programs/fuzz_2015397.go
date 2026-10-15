package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint }
func genFn0() int64 {
    for fi6 := 0; fi6 < 4; fi6++ {
        fmt.Println(fi6)
        switch ("\n\t" + ("\n\t" + "a")) {
        case "foo bar":
            if (0 >= 32767) {
                fmt.Println(fi6)
                v7 := []int{(fi6 & 0)}
                _ = v7
            }
        case "hello":
            var v8 int16 = 0
            _ = v8
        default:
            var v9 bool = true
            _ = v9
        }
        var v10 uint8 = 0
        _ = fi6
        _ = v10
    }
    v11 := []int{32768, 0, -2147483648}
    _ = v11
    return 1
}

func main() {
    var v1 uint = 0
    var v2 float32 = 1.5
    v3 := make(map[int]uint32)
    v4 := make(map[string]string)
    v5 := make(map[int]int)
    switch 2147483647 {
    case -275:
        for fi12 := 0; fi12 < 18; fi12++ {
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := []int{fi12}
                _ = fi13
                _ = v14
            }
            var v15 uint64 = 94
            v16 := GenS1{F0: 0, F1: -1}
            _ = fi12
            _ = v15
            _ = v16
        }
    case -2147483648:
        var v17 int32 = 1
        _ = v17
    default:
        var v18 int64 = 1
        _ = v18
    }
    fmt.Println(v1, v2)
    switch "" {
    case "a":
        delete(v5, -156)
    case "":
        switch "hello" {
        case "":
            fmt.Println(v2)
        case "a":
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 float32 = 0.1
                for fi21 := 0; fi21 < 2; fi21++ {
                    if (false && (!false)) {
                        v1 = v1
                        v20 += v20
                    }
                    var v22 int16 = -32768
                    _ = fi21
                    _ = v22
                }
                if true {
                    for fi23 := 0; fi23 < 13; fi23++ {
                        fmt.Println(v20, v20, v20)
                        _ = fi23
                    }
                    v24 := GenS2{F0: (^(v1))}
                    v20 -= float32((0.1 * -2.25))
                    _ = v24
                } else {
                    fmt.Println(v1, fi19)
                    v25 := GenS2{F0: (^(v1))}
                    _ = v25
                }
                _ = fi19
                _ = v20
            }
        default:
            if false {
                v26 := []int{0}
                v2 += (v2 - 3.14159)
                _ = v26
            }
        }
    }
    fmt.Println(v1, v2)
    v1 += 566
    switch 128 {
    case 4294967295:
        var v27 uint64 = 1
        _ = v27
    }
    if (!(0.0 < 1.5)) {
        var v28 uint16 = 1
        var v29 int16 = -1
        _ = v28
        _ = v29
    }
    var v30 uint64 = 0
    for fi31 := 0; fi31 < 17; fi31++ {
        fmt.Println(v30, fi31)
        _ = fi31
    }
    var v32 int32 = -1
    switch "a" {
    case "a":
        if false {
            var v33 uint32 = 42
            _ = v33
        }
    }
    var v34 int8 = 127
    v34 -= (v34 ^ v34)
    v2 = (v2 + v2)
    fmt.Println(v32)
    v2 += 0.0
    v35 := GenS1{F0: 1000, F1: v32}
    v36 := []int{32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
