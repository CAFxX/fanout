package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int8 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint16; F1 uint16; F2 uint32 }
func genFn0() uint32 {
    var v5 string = ("foo bar" + "hello")
    v5 = v5
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint = 1
    var v2 bool = false
    var v3 uint8 = 255
    v4 := make(map[int]int)
    if v2 {
        switch "a" {
        case "\n\t":
            for fi6 := 0; fi6 < 9; fi6++ {
                switch "hello" {
                case "a":
                    var v7 int32 = 0
                    _ = v7
                }
                _ = fi6
            }
        case "a":
            if (("hello" + "\n\t") == "\n\t") {
                v3 += (v3 >> (((^(v1))) & 63))
            } else {
                fmt.Println(v2, v1, v3)
                var v8 float64 = ((0.0 / 0.1) - float64(0.0))
                fmt.Println(v2)
                _ = v8
            }
        case "foo bar":
            var v9 int = 32768
            _ = v9
        default:
            _ = len(v4)
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            switch (fi10 >> ((v1) & 63)) {
            case -1:
                v3 = v3
            default:
                var v11 string = ("hello" + ("a" + "foo bar"))
                _ = v11
            }
            var v12 int64 = 42
            _ = fi10
            _ = v12
        }
        if v2 {
            if ("hello" > "foo bar") {
                fmt.Println(v1, v2)
            } else {
                v13 := GenS1{F0: (v3 - v3), F1: -128, F2: 127}
                if v2 {
                    fmt.Println(v1)
                    v14 := GenS2{F0: -1, F1: 0}
                    v3 = 2
                    _ = v14
                }
                fmt.Println(v1)
                _ = v13
            }
            var v15 float64 = (1.5 + 3.14159)
            var v16 uint64 = 2
            _ = v15
            _ = v16
        }
    }
    v17 := GenS3{F0: v2}
    v18 := []int{127, -9223372036854775808}
    var v19 bool = (0.1 >= 0.1)
    switch 128 {
    case 42:
        v18 = append(v18, -854)
    case 4294967295:
        var v20 uint64 = 611
        _ = v20
    default:
        v19 = v2
    }
    v19 = ((v2 || false) || (32767 <= 255))
    v21 := GenS4{F0: 65535, F1: 0}
    fmt.Println(v2)
    if v2 {
        var v22 int8 = 1
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 int64 = -9223372036854775808
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    switch "\n\t" {
    case "foo bar":
        v1 -= (v1 | v1)
    case "\n\t":
        var v25 int32 = -688
        _ = v25
    default:
        if true {
            switch 32767 {
            case 1:
                switch -10 {
                case 512:
                    _ = len(v18)
                case 128:
                    v26 := GenS1{F0: (^(v3)), F1: -11, F2: 1}
                    _ = v26
                case 2147483647:
                    v1 += v1
                }
            default:
                switch "foo bar" {
                case "hello":
                    v3 = 35
                default:
                    v27 := GenS2{F0: -1, F1: 4294967295}
                    _ = v27
                }
            }
            if v2 {
                v19 = (-40 > -937)
                if v19 {
                    v19 = v2
                    v28 := GenS2{F0: 574, F1: 2}
                    _ = v28
                } else {
                    v19 = v19
                    for fi29 := 0; fi29 < 8; fi29++ {
                        var v30 uint16 = 944
                        v31 := GenS5{F0: 2, F1: 42}
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                }
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
            } else {
                var v32 float32 = (3.14159 / (0.0 - 0.1))
                if len(v18) > 0 {
                    v18[0] = 42
                }
                var v33 float64 = -2.25
                _ = v32
                _ = v33
            }
            fmt.Println(v2)
        } else {
            if (4294967295 >= -9223372036854775808) {
                v34 := GenS4{F0: uint(v3), F1: 2147483647}
                _ = v34
            } else {
                var v35 float64 = ((1e10 / -2.25) - float64(0.0))
                _ = v35
            }
            fmt.Println(v3, v1, v3)
            v36 := GenS6{F0: 255, F1: 1, F2: 2147483647}
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
