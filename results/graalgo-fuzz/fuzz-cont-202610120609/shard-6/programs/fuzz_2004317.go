package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int64 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 2
    var v3 uint64 = 1
    v4 := []int{986, 659, 32768}
    if false {
        for fi5 := 0; fi5 < 2; fi5++ {
            v6 := GenS1{F0: -128, F1: (^(v2)), F2: 9223372036854775807}
            _ = fi5
            _ = v6
        }
        fmt.Println(v2)
        v2 += (v2 << ((304) & 63))
    } else {
        var v7 int = 0
        v8 := v1[(v7 * v7)]
        for fi9 := 0; fi9 < 15; fi9++ {
            if (!(1 != 42)) {
                fmt.Println(v2)
            } else {
                for fi10 := 0; fi10 < 5; fi10++ {
                    v11 := GenS1{F0: 42, F1: v2, F2: 1}
                    fmt.Println(v3, fi9)
                    v8 = (v8 >> ((uint(v3)) & 63))
                    _ = fi10
                    _ = v11
                }
                var v12 int16 = 0
                _ = v12
            }
            var v13 bool = ((1e10 / ((1e10 / (3.14159 / ((1.5 / ((3.14159 + (-2.25 - -2.25)) * 3.14159)) + ((1.5 - (0.1 + 0.1)) - (-2.25 + -2.25))))) / 1.5)) < 1.5)
            var v14 string = ("a" + ((("" + ("hello" + ("hello" + "hello"))) + ("" + "a")) + "hello"))
            _ = fi9
            _ = v13
            _ = v14
        }
        _ = v7
        _ = v8
    }
    if false {
        var v15 bool = (9223372036854775807 >= -1)
        v2 = v2
        _ = v15
    } else {
        for fi16 := 0; fi16 < 16; fi16++ {
            v2 -= (v2 * uint16(v3))
            v17 := GenS2{F0: (fi16 & (fi16 * fi16)), F1: -32768}
            _ = fi16
            _ = v17
        }
        var v18 uint16 = 42
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 uint32 = 65535
        _ = fi19
        _ = v20
    }
    fmt.Println(v3)
    var v21 int16 = 1000
    v22 := GenS3{F0: true, F1: 42}
    var v23 int32 = 42
    for fi24 := 0; fi24 < 2; fi24++ {
        v25 := v1[fi24]
        _ = fi24
        _ = v25
    }
    fmt.Println(v3)
    var v26 float64 = 1.5
    var v27 int32 = (v23 - v23)
    v23 += v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
