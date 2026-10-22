package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int }
func genFn0(p0 uint32, p1 uint64) int {
    if (!(42 > 32)) {
        if true {
            fmt.Println(p1)
            p1 = (^(p1))
        }
        var v6 uint64 = 332
        v7 := []int{1}
        _ = v6
        _ = v7
    } else {
        if false {
            switch 127 {
            case 32768:
                fmt.Println(p1)
            case -364:
                v8 := []int{9223372036854775807, int(p0), 32768}
                _ = v8
            case 0:
                var v9 int32 = -1
                _ = v9
            }
        }
        v10 := GenS1{F0: 2147483648, F1: true}
        _ = v10
    }
    var v11 string = ((((("\n\t" + "") + "hello") + "") + "\n\t") + "foo bar")
    var v12 float32 = 0.0
    v13 := GenS1{F0: p0, F1: false}
    var v14 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return -175
}

func main() {
    var v1 bool = (42 != 0)
    var v2 uint16 = 34
    v3 := make(map[string]bool)
    var v4 uint64 = 533
    var v5 bool = true
    if ((42 < 42) || v1) {
        switch 1 {
        case 32768:
            var v15 float32 = 0.1
            _ = v15
        }
    }
    var v16 string = "\n\t"
    var v17 uint64 = (v4 * v4)
    for fi18 := 0; fi18 < 17; fi18++ {
        switch v16 {
        case "a":
            fmt.Println(v4, fi18, v16)
        default:
            fmt.Println(v2, v16, v5)
        }
        switch fi18 {
        case 32767:
            v19 := GenS2{F0: 73, F1: 2}
            _ = v19
        }
        v20 := GenS3{F0: v2}
        _ = fi18
        _ = v20
    }
    v21 := GenS4{F0: 127}
    v16 = v16
    v17 = v4
    v5 = (65535 >= 2147483647)
    var v22 uint = 202
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
