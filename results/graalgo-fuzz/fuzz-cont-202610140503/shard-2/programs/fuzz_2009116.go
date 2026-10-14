package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 int; F1 bool; F2 int16 }

func main() {
    var v1 float32 = float32(-2.25)
    var v2 int = -9223372036854775808
    v3 := []int{v2, v2}
    var v4 int64 = -860
    var v5 int64 = v4
    switch ("\n\t" + "hello") {
    case "":
        switch ("foo bar" + ("hello" + "")) {
        case "":
            var v6 uint = 42
            _ = v6
        }
    case "hello":
        v2 = 41
    default:
        v7 := GenS1{F0: 1, F1: v2, F2: 532}
        _ = v7
    }
    v2 = v2
    var v8 int16 = -589
    v3 = append(v3, (v2 | (v2 - v2)))
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint16 = 65535
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 bool = false
        if v12 {
            v8 = v8
            v2 -= -555
            var v13 uint64 = 889
            _ = v13
        } else {
            _ = len(v3)
        }
        _ = fi11
        _ = v12
    }
    v14 := GenS2{F0: 4294967295, F1: 0, F2: 65535}
    v3 = append(v3, -45)
    var v15 int = 411
    v16 := GenS3{F0: (^(v2)), F1: false, F2: (v8 << ((65535) & 63))}
    var v17 float64 = 0.1
    if (!true) {
        if len(v3) > 0 {
            v3[0] = v2
        }
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 int32 = -48
        _ = fi18
        _ = v19
    }
    v3 = append(v3, (v2 << ((65535) & 63)))
    v4 += v5
    v20 := GenS3{F0: int(v4), F1: (42 <= 127), F2: 1000}
    switch v2 {
    case -1:
        switch "\n\t" {
        case "":
            var v21 float64 = v17
            _ = v21
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
