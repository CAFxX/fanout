package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int8 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int; F1 int8; F2 int8 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint16 }
func genFn0() uint {
    if (0 >= 323) {
        var v4 float32 = 0.1
        v5 := GenS1{F0: 32768, F1: 255, F2: -128}
        var v6 uint8 = 2
        _ = v4
        _ = v5
        _ = v6
    } else {
        v7 := []int{587, -696, 127, 2147483647}
        var v8 bool = true
        _ = len(v7)
        _ = v7
        _ = v8
    }
    var v9 bool = false
    fmt.Println(v9)
    switch ("" + "\n\t") {
    case "hello":
        v10 := []int{-609, 32767, 127}
        _ = v10
    default:
        v11 := []int{-872, 124, 32767, 32767}
        _ = v11
    }
    _ = v9
    return 56
}

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{2}
    var v3 uint32 = 4294967295
    v12 := GenS1{F0: 0, F1: 96, F2: -128}
    fmt.Println(v1)
    var v13 float32 = -2.25
    var v14 int = 128
    v14 = 32768
    var v15 float64 = -2.25
    switch v14 {
    case 32768:
        switch "" {
        case "":
            v1 += 1.5
        default:
            v3 = (v3 ^ v3)
        }
    case 128:
        fmt.Println(v15)
    case 9223372036854775807:
        fmt.Println(v14, v14, v15)
    default:
        v16 := GenS2{F0: 787, F1: ((("foo bar" + "a") + ("hello" + "")) + ("a" + "a"))}
        _ = v16
    }
    var v17 string = ("\n\t" + "")
    v18 := GenS3{F0: int(v3), F1: 98, F2: 100}
    v19 := GenS2{F0: (v14 >> ((715) & 63)), F1: v17}
    var v20 int16 = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v21 int16 = (v20 - (v20 << ((60) & 63)))
    v13 -= 0.0
    if true {
        var v22 uint = 255
        switch 0 {
        case 0:
            if len(v2) > 0 {
                v2[0] = v3
            }
        case -2147483648:
            var v23 float32 = v13
            _ = v23
        case 4294967295:
            for fi24 := 0; fi24 < 11; fi24++ {
                v25 := GenS4{F0: -1, F1: 0, F2: 0}
                _ = fi24
                _ = v25
            }
        }
        var v26 int16 = v20
        _ = v22
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
