package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int }
type GenS2 struct { F0 bool }
func genFn0(p0 string) int {
    var v5 uint64 = 18446744073709551615
    var v6 int = 128
    var v7 int = (-(v6))
    var v8 uint32 = 153
    var v9 uint8 = 151
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 32768
}
func genFn1(p0 int8) uint64 {
    var v10 int64 = 1
    v11 := GenS1{F0: "a", F1: 255, F2: 4294967295}
    _ = p0
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 float64 = (((3.14159 / 1.0) - (float64(((1e10 + -2.25) / 1e10)) * 3.14159)) + (0.1 + 3.14159))
    var v2 string = ""
    v3 := []int{2147483647, 0}
    var v4 float32 = (-2.25 + (-2.25 / 0.1))
    v2 = (v2 + v2)
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(v1)
        _ = fi12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        fmt.Println(v4, v1)
        _ = fi13
    }
    fmt.Println(v4, v1)
    switch 581 {
    case -9223372036854775808:
        v4 += (1e10 / v4)
    case 2147483647:
        v14 := GenS2{F0: (2 == 4294967295)}
        _ = v14
    case -640:
        fmt.Println(v1, v2)
    default:
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 bool = false
            fmt.Println(v2)
            switch (fi15 + fi15) {
            case 722:
                var v17 int8 = 34
                _ = v17
            case 484:
                var v18 float32 = ((v4 - (v4 + ((0.0 - (v4 * v4)) * float32(-2.25)))) / 0.1)
                _ = v18
            case -231:
                v2 = v2
            default:
                for fi19 := 0; fi19 < 8; fi19++ {
                    v16 = v16
                    _ = fi19
                }
            }
            _ = fi15
            _ = v16
        }
    }
    switch v2 {
    case "\n\t":
        if (genFn1(1) > 2) {
            v1 = 0.0
            var v20 string = "\n\t"
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v20
        } else {
            if false {
                if len(v3) > 0 {
                    v3[0] = 32767
                }
                var v21 uint8 = 255
                _ = v21
            }
            v1 -= v1
        }
    case "hello":
        var v22 uint32 = 1
        _ = v22
    }
    v23 := GenS2{F0: (0 <= 519)}
    var v24 uint64 = genFn1(127)
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    switch "\n\t" {
    case "\n\t":
        var v25 int16 = 32767
        _ = v25
    case "a":
        var v26 int32 = 1
        _ = v26
    default:
        var v27 uint32 = 967
        _ = v27
    }
    var v28 int = 42
    for fi29 := 0; fi29 < 3; fi29++ {
        v4 = v4
        _ = fi29
    }
    v30 := GenS1{F0: "", F1: 2, F2: genFn0(v2)}
    switch (v28 + 0) {
    case -9223372036854775808:
        var v31 int64 = 9223372036854775807
        _ = v31
    case 302:
        switch genFn0(("\n\t" + (v2 + v2))) {
        case -276:
            v4 += (v4 * float32(v1))
        case 2147483647:
            if (18446744073709551615 < 255) {
                var v32 int16 = -1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v33 int32 = -332
                _ = v32
                _ = v33
            } else {
                var v34 uint32 = 969
                var v35 int8 = 127
                var v36 uint8 = 255
                _ = v34
                _ = v35
                _ = v36
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v30
    _ = v4
}
