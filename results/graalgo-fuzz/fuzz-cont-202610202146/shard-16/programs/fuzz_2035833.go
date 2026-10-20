package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint; F1 bool; F2 int }

func main() {
    var v1 uint16 = 32767
    var v2 int8 = 0
    var v3 int32 = -1
    var v4 bool = (!true)
    for fi5 := 0; fi5 < 11; fi5++ {
        v3 += v3
        var v6 float64 = (float64((-2.25 / 1e10)) + (-2.25 * 3.14159))
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: -32768}
    v8 := GenS2{F0: "a", F1: (0 == (^(v3)))}
    fmt.Println(v4)
    var v9 uint8 = 128
    v1 += uint16(v9)
    v10 := []int{990}
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS3{F0: 42, F1: v4, F2: int(v1)}
        _ = fi11
        _ = v12
    }
    v1 = 65535
    if ((v9 + (v9 + v9)) < v9) {
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 int64 = 9223372036854775807
            v10 = append(v10, fi13)
            _ = fi13
            _ = v14
        }
    } else {
        v15 := GenS3{F0: 255, F1: (732 <= 1), F2: 0}
        if false {
            v4 = false
        }
        _ = v15
    }
    var v16 uint = 818
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    fmt.Println(v1)
    var v17 bool = v4
    switch "a" {
    case "":
        var v18 bool = v4
        _ = v18
    case "foo bar":
        v17 = v4
    case "a":
        v10 = append(v10, 4294967295)
    }
    v19 := GenS1{F0: -721}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
