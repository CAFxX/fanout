package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int8 }
type GenS2 struct { F0 int64; F1 uint; F2 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 int16; F2 int8 }

func main() {
    var v1 uint = 42
    var v2 float32 = float32(-2.25)
    var v3 string = "hello"
    var v4 float64 = ((1.5 + -2.25) / ((0.1 - 0.0) / 3.14159))
    v5 := GenS1{F0: 42, F1: (((v3 + (v3 + v3)) + v3) + ""), F2: 42}
    v3 = v3
    switch v3 {
    case "foo bar":
        var v6 int = 9223372036854775807
        _ = v6
    case "":
        switch "" {
        case "foo bar":
            v7 := []int{}
            _ = v7
        }
    case "\n\t":
        switch v3 {
        case "foo bar":
            v8 := GenS1{F0: 1, F1: "hello", F2: 0}
            _ = v8
        case "hello":
            v9 := []int{}
            _ = v9
        default:
            switch 4294967295 {
            case 32767:
                var v10 int8 = 111
                _ = v10
            case 42:
                switch (v3 + v3) {
                case "foo bar":
                    v1 = 847
                default:
                    if ((301 > 32767) || true) {
                        if false {
                            v11 := []int{}
                            for fi12 := 0; fi12 < 7; fi12++ {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                _ = fi12
                            }
                            _ = v11
                        }
                    } else {
                        v2 = 1e10
                        if true {
                            v2 = v2
                            v13 := GenS2{F0: 42, F1: 42, F2: false}
                            var v14 int8 = -1
                            _ = v13
                            _ = v14
                        }
                        var v15 int16 = -1
                        _ = v15
                    }
                }
            }
        }
    default:
        var v16 uint = v1
        _ = v16
    }
    var v17 uint8 = 182
    v4 = ((v4 - (((v4 / v4) * v4) * v4)) + 0.1)
    var v18 int8 = -1
    v4 -= v4
    v19 := GenS3{F0: v1}
    var v20 int = 32767
    switch v3 {
    case "foo bar":
        for fi21 := 0; fi21 < 13; fi21++ {
            for fi22 := 0; fi22 < 3; fi22++ {
                var v23 uint = 329
                fmt.Println(v2)
                var v24 float32 = float32(((v4 - (v4 / (v4 - float64(v2)))) / v4))
                var v25 float32 = (v2 / v24)
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi21
        }
    case "a":
        if (v2 <= 0.0) {
            fmt.Println(v20, v3, v18)
            var v26 uint32 = 919
            if (!(0 != 1)) {
                fmt.Println(v20, v17, v20)
            }
            _ = v26
        }
    }
    if false {
        var v27 uint64 = 1
        v2 -= (v2 + (v2 * v2))
        _ = v27
    }
    switch 128 {
    case 9223372036854775807:
        v28 := GenS4{F0: 1, F1: 0, F2: v18}
        _ = v28
    case -1:
        if true {
            var v29 uint64 = 42
            v3 = "foo bar"
            _ = v29
        }
    case 1:
        v1 = v1
    default:
        var v30 int = 128
        _ = v30
    }
    var v31 int64 = 1
    v4 += float64((v2 / v2))
    var v32 int64 = 95
    var v33 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
