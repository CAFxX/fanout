package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int8 }
func genFn0() uint16 {
    switch ("hello" + "a") {
    case "":
        var v6 bool = (!true)
        _ = v6
    case "a":
        switch -9223372036854775808 {
        case 32767:
            var v7 uint32 = 4294967295
            _ = v7
        case 102:
            var v8 int16 = 32767
            _ = v8
        case -9223372036854775808:
            v9 := GenS1{F0: 128, F1: 42}
            _ = v9
        default:
            v10 := []int{844}
            _ = v10
        }
    case "foo bar":
        if false {
            switch ("" + "a") {
            case "a":
                var v11 float64 = 0.1
                _ = v11
            case "":
                if false {
                    var v12 int64 = 797
                    _ = v12
                }
            case "\n\t":
                v13 := GenS1{F0: 42, F1: 2}
                _ = v13
            }
            var v14 uint32 = 2
            _ = v14
        }
    }
    switch -1 {
    case 4294967295:
        var v15 float32 = 1.5
        _ = v15
    case 42:
        v16 := []int{419, 32767}
        _ = v16
    case 885:
        v17 := GenS1{F0: 2, F1: 702}
        _ = v17
    default:
        v18 := GenS1{F0: 127, F1: 858}
        _ = v18
    }
    var v19 uint8 = 255
    switch 127 {
    case 389:
        var v20 int8 = 28
        _ = v20
    default:
        for fi21 := 0; fi21 < 3; fi21++ {
            v19 += v19
            v19 += v19
            fmt.Println(v19)
            _ = fi21
        }
    }
    _ = v19
    return 870
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = -2.25
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    v5 := make(map[string]bool)
    v22 := GenS2{F0: -128}
    v2 -= (-2.25 * v2)
    fmt.Println(v2)
    var v23 int64 = 9223372036854775807
    v2 -= (-2.25 - v2)
    if (-9223372036854775808 > (v23 - v23)) {
        v2 = 3.14159
        var v24 float32 = v2
        v24 = -2.25
        _ = v24
    }
    fmt.Println(v23)
    if false {
        for fi25 := 0; fi25 < 4; fi25++ {
            v2 = -2.25
            v2 = 3.14159
            _ = fi25
        }
        var v26 int64 = v23
        _ = v26
    }
    fmt.Println(v2)
    v1[("" + "a")] = 127
    var v27 float64 = (1.5 - (3.14159 * (-2.25 + float64(((float32(3.14159) * v2) - float32((float64((0.1 - (v2 * 1.5))) + 0.1)))))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
