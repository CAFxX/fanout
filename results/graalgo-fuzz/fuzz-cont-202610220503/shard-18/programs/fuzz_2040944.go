package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    var v3 uint8 = 226
    v4 := []string{"hello"}
    var v5 float32 = 1e10
    for fi6 := 0; fi6 < 7; fi6++ {
        if (!false) {
            var v7 uint32 = 0
            _ = v7
        } else {
            v3 = 255
        }
        var v8 int32 = -1
        _ = fi6
        _ = v8
    }
    switch "a" {
    case "a":
        if (42 < 2) {
            var v9 int32 = 0
            v1[9223372036854775807] = 948
            fmt.Println(v3, v5, v5)
            _ = v9
        }
    case "hello":
        v3 += 127
    case "":
        v10 := GenS1{F0: 32767, F1: 2}
        _ = v10
    }
    if true {
        v5 = 0.1
        v3 += v3
    } else {
        for fi11 := 0; fi11 < 13; fi11++ {
            if (4294967295 >= uint(v3)) {
                var v12 uint64 = 0
                if (-1 != -1) {
                    var v13 string = "a"
                    var v14 int64 = 9223372036854775807
                    _ = v13
                    _ = v14
                }
                delete(v2, -287)
                _ = v12
            }
            var v15 int16 = -32768
            _ = fi11
            _ = v15
        }
        v5 = (v5 * v5)
        _ = len(v2)
    }
    var v16 int64 = -9223372036854775808
    var v17 uint = 162
    var v18 uint64 = 42
    _ = len(v2)
    v19 := GenS2{F0: 935, F1: -9223372036854775808, F2: v3}
    if true {
        var v20 float64 = 1.5
        var v21 float64 = v20
        _ = v20
        _ = v21
    }
    v5 -= ((float32((3.14159 - 3.14159)) + v5) * (-2.25 * ((v5 * (v5 - ((v5 + 0.1) - v5))) * v5)))
    if false {
        v22 := GenS1{F0: 128, F1: (v18 - 4294967295)}
        var v23 float64 = (0.0 / 3.14159)
        v24 := GenS2{F0: 1, F1: (v16 % ((v16) | 1)), F2: (v3 >> ((v17) & 63))}
        _ = v22
        _ = v23
        _ = v24
    }
    _ = len(v4)
    for fi25 := 0; fi25 < 10; fi25++ {
        v4 = append(v4, "hello")
        _ = fi25
    }
    var v26 string = "hello"
    if true {
        var v27 string = v26
        var v28 int64 = (v16 / ((v16) | 1))
        if true {
            v2[-1] = "a"
        } else {
            v18 = v18
            for fi29 := 0; fi29 < 12; fi29++ {
                v27 = (v27 + v26)
                _ = fi29
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = v27
        _ = v28
    } else {
        switch "" {
        case "":
            fmt.Println(v5)
        case "a":
            var v30 bool = true
            _ = v30
        }
        var v31 uint32 = 1
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
