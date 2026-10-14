package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int8 = -43
    v2 := make(map[int]int)
    var v3 string = "\n\t"
    switch (v3 + "\n\t") {
    case "a":
        v1 -= v1
    }
    var v4 int64 = -715
    v3 = v3
    var v5 float32 = float32(-2.25)
    if (!true) {
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 uint64 = 42
            _ = fi6
            _ = v7
        }
    } else {
        var v8 uint64 = 704
        v3 = v3
        fmt.Println(v8, v8)
        _ = v8
    }
    if true {
        v9 := GenS1{F0: 65535, F1: v1, F2: 242}
        v10 := []int{-935, 1, 650}
        _ = v10
        _ = v9
    } else {
        if ((1 >= 127) || (int(v1) < 2147483647)) {
            var v11 bool = true
            _ = v11
        } else {
            fmt.Println(v4, v5)
            v4 = -1
            var v12 uint8 = 0
            _ = v12
        }
        var v13 float64 = (3.14159 * float64((v5 + ((v5 / (v5 / (v5 - v5))) / v5))))
        switch v3 {
        case "":
            v14 := GenS2{F0: 65535, F1: -915}
            _ = v14
        case "\n\t":
            v15 := GenS3{F0: ("hello" + "foo bar")}
            _ = v15
        case "hello":
            if true {
                var v16 int64 = (v4 >> ((551) & 63))
                _ = v16
            } else {
                var v17 uint8 = 127
                v17 -= (v17 & (^(v17)))
                v18 := GenS1{F0: 506, F1: (v1 >> ((1) & 63)), F2: 25}
                _ = v17
                _ = v18
            }
        default:
            v19 := v2[0]
            _ = v19
        }
        _ = v13
    }
    delete(v2, -455)
    v20 := GenS2{F0: 65535, F1: 2147483647}
    if ("\n\t" != "a") {
        if (564 < 1) {
            v21 := GenS4{F0: (128 != 127)}
            if false {
                v1 = (v1 | (v1 & (v1 | v1)))
                var v22 int64 = v4
                var v23 int16 = 32767
                _ = v22
                _ = v23
            } else {
                v24 := GenS5{F0: 201}
                _ = v24
            }
            _ = v21
        }
        if false {
            v1 += (v1 * v1)
        }
    }
    var v25 int64 = 1
    v25 += -1
    if (88 == 103) {
        v26 := GenS5{F0: 255}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
