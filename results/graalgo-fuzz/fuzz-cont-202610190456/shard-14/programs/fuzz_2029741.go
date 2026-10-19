package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int16; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 0.0
    var v3 float32 = 3.14159
    var v4 string = (("a" + "a") + ((("" + "") + "\n\t") + "a"))
    var v5 int8 = 100
    var v6 int16 = -32768
    v3 = 0.1
    v4 = ""
    var v7 float32 = v2
    v2 -= v2
    switch "hello" {
    case "hello":
        v8 := GenS1{F0: 42, F1: -1}
        _ = v8
    case "a":
        v9 := GenS2{F0: 0}
        _ = v9
    }
    var v10 float64 = (1e10 * 3.14159)
    if false {
        var v11 int64 = -203
        switch (("\n\t" + (v4 + v4)) + ("hello" + "")) {
        case "hello":
            if true {
                v12 := GenS3{F0: 2, F1: v4}
                _ = v12
            } else {
                if false {
                    v7 = v3
                    v13 := GenS4{F0: 42, F1: 310, F2: 4294967295}
                    _ = v13
                } else {
                    var v14 uint8 = 128
                    _ = v14
                }
            }
        case "\n\t":
            v5 += (v5 & (v5 + (v5 >> ((255) & 63))))
        case "":
            v1[""] = (v4 + v4)
        default:
            v15 := []int{-380, 235}
            _ = v15
        }
        _ = v11
    } else {
        v16 := []int{}
        fmt.Println(v7, v10, v7)
        _ = v16
    }
    v4 = "hello"
    var v17 uint16 = 2
    v18 := GenS5{F0: (-(v6)), F1: 318, F2: int64(v6)}
    var v19 uint = 255
    switch 9223372036854775807 {
    case -268:
        var v20 uint64 = 1
        _ = v20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 string = (("foo bar" + v4) + v4)
            var v24 int8 = (v5 / ((v5) | 1))
            var v25 int = (fi21 / ((fi21) | 1))
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        switch (v4 + "a") {
        case "\n\t":
            for fi26 := 0; fi26 < 6; fi26++ {
                switch fi26 {
                case 42:
                    switch (("foo bar" + "foo bar") + ((v4 + ("\n\t" + ((("foo bar" + "hello") + v4) + v4))) + v4)) {
                    case "\n\t":
                        fmt.Println(v7, v19)
                    case "a":
                        v1[(v4 + "foo bar")] = v4
                    case "":
                        fmt.Println(v7)
                    default:
                        switch (fi21 / ((fi21) | 1)) {
                        case 128:
                            if (false || false) {
                                switch (fi21 - (fi26 + -305)) {
                                case 127:
                                    v27 := GenS2{F0: 255}
                                    _ = v27
                                case -1:
                                    fmt.Println(v17)
                                case 1:
                                    switch v4 {
                                    case "hello":
                                        var v28 bool = true
                                        _ = v28
                                    case "foo bar":
                                        v29 := []int{(fi26 - 110), fi26, -317}
                                        _ = v29
                                    case "\n\t":
                                        if false {
                                            var v30 uint8 = 127
                                            var v31 uint = v19
                                            _ = v30
                                            _ = v31
                                        } else {
                                            switch -361 {
                                            case 42:
                                                if false {
                                                    fmt.Println(v7)
                                                    v32 := []int{}
                                                    var v33 int16 = v6
                                                    if false {
                                                        v10 += (float64(v7) + (1.5 * (0.1 * (1.5 + float64((float32((3.14159 + 0.1)) / v3))))))
                                                        var v34 int16 = v33
                                                        v34 += 42
                                                        _ = v34
                                                    }
                                                    if ((false && true) || false) {
                                                        v35 := GenS6{F0: 2, F1: v19}
                                                        if len(v32) > 0 {
                                                            v32[0] = (fi26 << ((v19) & 63))
                                                        }
                                                        _ = v35
                                                    }
                                                    _ = v32
                                                    _ = v33
                                                }
                                            case 9223372036854775807:
                                                v10 -= 1e10
                                            }
                                        }
                                    }
                                }
                                var v36 uint32 = 42
                                v1[(v4 + ((v4 + v4) + v4))] = "hello"
                                v37 := GenS7{F0: (v5 - (v5 + (v5 << (((^(v19))) & 63)))), F1: 0}
                                if (v5 >= v5) {
                                    fmt.Println(v4, v17, v17)
                                }
                                _ = v36
                                _ = v37
                            } else {
                                var v38 bool = (2147483648 == 42)
                                _ = v38
                            }
                        }
                    }
                default:
                    v6 -= 503
                }
                _ = fi26
            }
        case "hello":
            switch v4 {
            case "hello":
                switch ("hello" + "") {
                case "hello":
                    v19 += v19
                }
            case "\n\t":
                if false {
                    var v39 bool = (true && true)
                    v40 := GenS5{F0: v6, F1: 42, F2: int64(v19)}
                    var v41 uint8 = 130
                    _ = v39
                    _ = v40
                    _ = v41
                }
            default:
                v42 := v1[("\n\t" + "foo bar")]
                _ = v42
            }
        }
        v43 := []int{(fi21 * (fi21 - (fi21 + -93))), -586, (fi21 ^ -2147483648)}
        _ = fi21
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
