package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 string; F1 int8; F2 bool }

func main() {
    var v1 float32 = 1.5
    var v2 uint16 = 195
    var v3 uint = 0
    v4 := make(map[string]string)
    v1 = v1
    switch 950 {
    case 127:
        switch "hello" {
        case "foo bar":
            if true {
                v5 := GenS1{F0: 0, F1: 772, F2: v2}
                fmt.Println(v3)
                if (1 < 255) {
                    v6 := GenS1{F0: 128, F1: 42, F2: (v2 - 550)}
                    for fi7 := 0; fi7 < 20; fi7++ {
                        if (!true) {
                            var v8 uint16 = v2
                            _ = v8
                        }
                        var v9 bool = true
                        var v10 int64 = -1
                        _ = fi7
                        _ = v10
                        _ = v9
                    }
                    if (255 == (v2 | 127)) {
                        var v11 int64 = 42
                        v3 -= (^(v3))
                        _ = v11
                    } else {
                        switch "\n\t" {
                        case "hello":
                            if true {
                                fmt.Println(v3, v3, v2)
                            } else {
                                v12 := GenS2{F0: (-32768 < -29)}
                                switch (((("" + ("foo bar" + "foo bar")) + "a") + "a") + (("foo bar" + "") + "")) {
                                case "foo bar":
                                    for fi13 := 0; fi13 < 2; fi13++ {
                                        v4["hello"] = "a"
                                        var v14 uint8 = 244
                                        _ = fi13
                                        _ = v14
                                    }
                                case "":
                                    if (!false) {
                                        v4[(("foo bar" + "") + "")] = "foo bar"
                                        if false {
                                            var v15 float32 = v1
                                            _ = v15
                                        }
                                        delete(v4, ("hello" + ("\n\t" + (("hello" + ("a" + ("hello" + ("a" + "a")))) + "a"))))
                                    }
                                case "\n\t":
                                    for fi16 := 0; fi16 < 15; fi16++ {
                                        if false {
                                            fmt.Println(fi16)
                                            v3 = (^(v3))
                                            v3 -= uint(v2)
                                        } else {
                                            var v17 float64 = (0.0 / -2.25)
                                            if ((fi16 == 0) && false) {
                                                var v18 int32 = 0
                                                _ = v18
                                            }
                                            _ = v17
                                        }
                                        _ = fi16
                                    }
                                }
                                _ = v12
                            }
                        case "\n\t":
                            fmt.Println(v1, v3, v3)
                        default:
                            v19 := GenS3{F0: false, F1: 0}
                            _ = v19
                        }
                        if (false || (("foo bar" + "\n\t") < "")) {
                            _ = len(v4)
                        } else {
                            var v20 int = -1
                            var v21 int64 = int64(v20)
                            _ = v20
                            _ = v21
                        }
                    }
                    _ = v6
                }
                _ = v5
            }
        case "hello":
            v3 -= v3
        }
    }
    var v22 int16 = 454
    switch "foo bar" {
    case "hello":
        switch int(v22) {
        case -943:
            for fi23 := 0; fi23 < 9; fi23++ {
                v24 := GenS3{F0: false, F1: v22}
                _ = fi23
                _ = v24
            }
        case -1:
            v25 := []int{-459, 641, 812}
            _ = v25
        case 456:
            fmt.Println(v2, v2)
        }
    case "foo bar":
        v26 := GenS4{F0: ("" + "hello"), F1: -1, F2: true}
        _ = v26
    case "":
        var v27 int32 = -1
        _ = v27
    }
    var v28 uint = (v3 ^ 18446744073709551615)
    switch "a" {
    case "a":
        if false {
            var v29 uint64 = 42
            var v30 uint16 = v2
            _ = v29
            _ = v30
        }
    }
    switch "\n\t" {
    case "":
        fmt.Println(v28, v1)
    }
    var v31 string = "foo bar"
    switch (("" + v31) + v31) {
    case "hello":
        var v32 uint64 = 824
        _ = v32
    case "foo bar":
        fmt.Println(v2, v3)
    case "a":
        v28 -= (v28 - 42)
    }
    v22 += (v22 + v22)
    var v33 float32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v31
    _ = v33
    _ = v4
}
