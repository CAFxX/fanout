package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint8; F1 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (18446744073709551615 > 1)
    var v3 bool = (65535 <= 397)
    var v4 uint = 668
    var v5 float32 = -2.25
    v6 := []int{-1, 4294967295, -777, -9223372036854775808}
    switch "\n\t" {
    case "hello":
        if true {
            var v7 bool = ((!false) && (!(2 <= 674)))
            delete(v1, "hello")
            _ = v7
        } else {
            if (!v2) {
                v8 := GenS1{F0: 42, F1: 42}
                _ = v8
            }
        }
    case "foo bar":
        var v9 int64 = 9223372036854775807
        _ = v9
    case "\n\t":
        var v10 uint16 = 0
        _ = v10
    }
    _ = len(v6)
    for fi11 := 0; fi11 < 6; fi11++ {
        v4 += 773
        v5 += float32(1.5)
        _ = fi11
    }
    v12 := GenS2{F0: 118, F1: 127}
    var v13 uint = 484
    fmt.Println(v2, v13)
    v14 := v1[((("a" + "\n\t") + "\n\t") + "foo bar")]
    switch -9223372036854775808 {
    case -1:
        v3 = true
    case 42:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v15 := GenS3{F0: 0, F1: 42}
    var v16 uint64 = 209
    var v17 int = 32767
    var v18 int32 = -1
    var v19 int16 = -816
    v20 := GenS2{F0: 1, F1: v17}
    var v21 uint16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
