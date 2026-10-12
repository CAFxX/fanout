package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
func genFn0() uint8 {
    var v6 int16 = -32768
    var v7 uint8 = 255
    v8 := []int{4294967295}
    v9 := GenS1{F0: 2}
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 float32 = (0.0 - 3.14159)
            v13 := GenS2{F0: 0}
            v7 = v7
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return v7
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 0.0
    var v3 uint32 = 1
    var v4 uint16 = 1
    var v5 int16 = 0
    if (!true) {
        v14 := GenS3{F0: (v3 ^ (v3 + v3))}
        _ = v14
    } else {
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS4{F0: 42}
            var v17 uint16 = v4
            fmt.Println(v5, fi15, v2)
            _ = fi15
            _ = v16
            _ = v17
        }
        switch ("hello" + "hello") {
        case "":
            for fi18 := 0; fi18 < 17; fi18++ {
                v5 += (v5 ^ (v5 >> ((uint(v4)) & 63)))
                v5 += 42
                v19 := []int{}
                _ = fi18
                _ = v19
            }
        }
        var v20 int16 = 42
        _ = v20
    }
    switch 1 {
    case 127:
        if ((9223372036854775807 != -1) || false) {
            for fi21 := 0; fi21 < 5; fi21++ {
                for fi22 := 0; fi22 < 7; fi22++ {
                    fmt.Println(fi21, v4)
                    _ = fi22
                }
                _ = fi21
            }
            var v23 float32 = (float32((float64(0.1) - v2)) * 0.0)
            _ = v23
        }
    case -1:
        var v24 int64 = -1
        _ = v24
    }
    fmt.Println(v5, v2)
    for fi25 := 0; fi25 < 5; fi25++ {
        fmt.Println(fi25, v5)
        var v26 int32 = 1
        _ = fi25
        _ = v26
    }
    var v27 string = "a"
    var v28 bool = (!(v3 <= (v3 >> ((255) & 63))))
    fmt.Println(v2, v28, v2)
    fmt.Println(v27)
    v28 = (!true)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
