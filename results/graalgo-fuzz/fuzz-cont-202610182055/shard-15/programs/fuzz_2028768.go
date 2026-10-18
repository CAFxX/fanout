package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint64) int16 {
    if false {
        v4 := GenS1{F0: 100}
        _ = v4
    } else {
        switch "" {
        case "hello":
            v5 := GenS1{F0: -86}
            _ = v5
        case "":
            p0 += 2
        case "\n\t":
            v6 := GenS1{F0: 127}
            _ = v6
        }
        v7 := GenS2{F0: 1, F1: 1}
        var v8 uint64 = (p0 ^ 589)
        _ = v7
        _ = v8
    }
    v9 := GenS3{F0: ((0 != (p0 | p0)) || ((65535 > 65535) || (128 == -1))), F1: ("\n\t" + (("" + "a") + "\n\t"))}
    _ = p0
    _ = v9
    return 32767
}

func main() {
    var v1 uint8 = 128
    var v2 float32 = float32(((0.1 - 1e10) - 0.1))
    var v3 int8 = 123
    var v10 uint32 = 4294967295
    if (925 >= 783) {
        v11 := GenS1{F0: v3}
        _ = v11
    } else {
        fmt.Println(v2, v2)
    }
    v2 -= ((v2 + (v2 - v2)) + (v2 + v2))
    var v12 uint16 = 42
    v2 += 1.5
    v13 := []int{127}
    for fi14 := 0; fi14 < 19; fi14++ {
        v3 -= 1
        _ = fi14
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS4{F0: -1, F1: 65535}
        switch ("hello" + "") {
        case "hello":
            fmt.Println(v1, fi15, v3)
        case "\n\t":
            v17 := GenS3{F0: false, F1: "\n\t"}
            _ = v17
        case "":
            var v18 string = ""
            _ = v18
        default:
            v2 = v2
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
