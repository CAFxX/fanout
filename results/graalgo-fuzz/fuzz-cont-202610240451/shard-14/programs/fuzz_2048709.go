package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int64; F2 bool }

func main() {
    var v1 int32 = -2147483648
    v2 := []string{"hello", ""}
    var v3 bool = true
    var v4 uint32 = 335
    v5 := []int{-608, 2147483647}
    var v6 int8 = -3
    v7 := GenS1{F0: 0, F1: 65535}
    switch 42 {
    case 9223372036854775807:
        v8 := GenS2{F0: 1}
        _ = v8
    case 32768:
        var v9 bool = v3
        _ = v9
    case 788:
        if v3 {
            var v10 string = "a"
            v10 = v10
            var v11 int64 = 1
            _ = v10
            _ = v11
        } else {
            v1 -= (v1 & 42)
            fmt.Println(v3)
            v5 = append(v5, -9223372036854775808)
        }
    }
    v12 := GenS2{F0: 255}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v4 = (^(v4))
    v6 -= v6
    var v13 bool = v3
    var v14 uint8 = 255
    v15 := GenS1{F0: 759, F1: 42}
    switch "\n\t" {
    case "a":
        switch 127 {
        case 9223372036854775807:
            v16 := GenS3{F0: (v4 * (v4 | v4)), F1: 2, F2: 551}
            _ = v16
        }
    }
    var v17 uint16 = 0
    switch -767 {
    case -2147483648:
        if (!(!("foo bar" != ("" + "\n\t")))) {
            v18 := GenS1{F0: 1, F1: 65535}
            fmt.Println(v6, v4)
            _ = v18
        } else {
            fmt.Println(v3, v17, v6)
        }
    default:
        v19 := GenS4{F0: 42}
        _ = v19
    }
    var v20 int8 = -1
    for fi21 := 0; fi21 < 3; fi21++ {
        if ((v4 >> ((252) & 63)) < v4) {
            for fi22 := 0; fi22 < 1; fi22++ {
                v23 := GenS5{F0: v6, F1: 42, F2: v3}
                v14 -= v14
                var v24 uint64 = 0
                _ = fi22
                _ = v23
                _ = v24
            }
            v25 := GenS3{F0: v4, F1: 206, F2: 32767}
            fmt.Println(v4, v3, v1)
            _ = v25
        } else {
            fmt.Println(v3, v6)
        }
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
