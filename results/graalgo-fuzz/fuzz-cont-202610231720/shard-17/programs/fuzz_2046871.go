package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 bool }

func main() {
    var v1 int64 = 0
    v2 := make(map[string]bool)
    var v3 string = ""
    var v4 uint = 519
    if ((v1 / ((42) | 1)) < v1) {
        v5 := GenS1{F0: 1, F1: 933, F2: (v3 + (v3 + ((("hello" + v3) + v3) + v3)))}
        _ = v5
    } else {
        var v6 bool = (true || false)
        _ = v6
    }
    var v7 uint32 = 2147483648
    v8 := GenS1{F0: 32767, F1: -1, F2: (v3 + "foo bar")}
    if false {
        if ((1e10 * 1.5) == float32((1e10 / 1.5))) {
            fmt.Println(v4)
            v3 = v3
            var v9 uint8 = 2
            _ = v9
        }
    } else {
        var v10 int64 = v1
        v7 = 531
        fmt.Println(v7)
        _ = v10
    }
    v3 = v3
    if true {
        if (42 < 1) {
            v4 = uint(v7)
            v11 := GenS1{F0: 779, F1: 69, F2: ""}
            var v12 uint16 = 255
            _ = v11
            _ = v12
        } else {
            var v13 uint32 = v7
            if true {
                v14 := []int{-9223372036854775808, -1}
                for fi15 := 0; fi15 < 20; fi15++ {
                    v16 := GenS2{F0: (!false)}
                    _ = fi15
                    _ = v16
                }
                if len(v14) > 0 {
                    v14[0] = -681
                }
                _ = v14
            }
            _ = v13
        }
        if false {
            v1 = 9223372036854775807
            if true {
                v17 := GenS3{F0: uint16(v4), F1: -101}
                _ = v17
            }
        }
        var v18 float64 = (0.1 / 1e10)
        _ = v18
    }
    fmt.Println(v4, v7, v3)
    if (!((true || (32767 == -9223372036854775808)) && (-1 > 42))) {
        if false {
            if true {
                var v19 uint64 = 1
                _ = v19
            }
            fmt.Println(v7)
            switch v3 {
            case "foo bar":
                _ = len(v2)
            case "\n\t":
                v3 = ""
            case "":
                v20 := GenS4{F0: v3}
                _ = v20
            default:
                var v21 uint16 = 1
                _ = v21
            }
        }
        var v22 uint = v4
        _ = v22
    }
    if true {
        v1 = 856
        v23 := GenS5{F0: true, F1: (!(128 > -1))}
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 uint8 = 128
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    v26 := GenS4{F0: "hello"}
    v27 := GenS5{F0: false, F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
