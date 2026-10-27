package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint8; F2 bool }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 18446744073709551615
    var v3 uint16 = 587
    var v4 uint64 = 18446744073709551615
    v3 = (v3 >> ((580) & 63))
    if ("foo bar" == (("hello" + ("hello" + ("foo bar" + "foo bar"))) + ("foo bar" + "foo bar"))) {
        if true {
            switch "hello" {
            case "hello":
                var v5 uint16 = v3
                _ = v5
            case "":
                if (923 == -2147483648) {
                    switch "foo bar" {
                    case "foo bar":
                        v6 := []int{127}
                        _ = v6
                    }
                    fmt.Println(v4)
                } else {
                    if (((-2.25 / (3.14159 - ((3.14159 + 1.5) - (((1e10 - 0.0) * (1.5 * 0.1)) + float64(-2.25))))) - 3.14159) > (0.0 / 3.14159)) {
                        v1[int(v2)] = ("hello" + "foo bar")
                    }
                    fmt.Println(v4, v3, v3)
                }
            }
        } else {
            v7 := []int{-9223372036854775808, 127}
            _ = v7
        }
    }
    var v8 int64 = -1
    if true {
        v9 := GenS1{F0: v3, F1: uint8(v3)}
        _ = v9
    }
    if ((v8 + 342) >= v8) {
        if false {
            if false {
                for fi10 := 0; fi10 < 12; fi10++ {
                    for fi11 := 0; fi11 < 14; fi11++ {
                        if true {
                            if true {
                                v3 += 442
                                _ = len(v1)
                                v12 := GenS1{F0: v3, F1: 2}
                                _ = v12
                            }
                            fmt.Println(v8, v8, v8)
                            var v13 int32 = 1
                            _ = v13
                        }
                        var v14 uint16 = v3
                        v4 = 673
                        _ = fi11
                        _ = v14
                    }
                    var v15 uint8 = 0
                    _ = fi10
                    _ = v15
                }
                switch ("a" + "foo bar") {
                case "\n\t":
                    var v16 float32 = (3.14159 - (float32(((1e10 + 0.0) * 0.0)) + (0.1 * 1e10)))
                    _ = v16
                case "foo bar":
                    fmt.Println(v3, v2, v3)
                }
            }
        } else {
            v2 = (v2 << ((v2) & 63))
            if (false || (!(false || ((4294967295 < 896) || (4294967295 < v2))))) {
                var v17 int16 = 1000
                var v18 uint32 = 853
                _ = v17
                _ = v18
            } else {
                var v19 uint64 = (v4 << ((0) & 63))
                v2 = (v2 + v2)
                var v20 string = ""
                _ = v19
                _ = v20
            }
            switch ("foo bar" + ("\n\t" + "hello")) {
            case "a":
                v21 := GenS2{F0: -149, F1: 127, F2: false}
                _ = v21
            case "\n\t":
                v8 = (v8 << ((56) & 63))
            case "":
                if (!true) {
                    v3 += (v3 % (((v3 - 402)) | 1))
                    for fi22 := 0; fi22 < 12; fi22++ {
                        switch "hello" {
                        case "hello":
                            v4 = (^(v4))
                        case "a":
                            if false {
                                v23 := GenS1{F0: v3, F1: 1}
                                _ = v23
                            }
                        }
                        var v24 uint32 = 1
                        _ = fi22
                        _ = v24
                    }
                } else {
                    fmt.Println(v8, v2)
                    v25 := GenS3{F0: v8, F1: (v8 > v8)}
                    var v26 uint = 470
                    _ = v25
                    _ = v26
                }
            }
        }
    }
    var v27 int32 = 0
    v28 := GenS4{F0: ((0.1 / 1.5) > 3.14159)}
    var v29 uint = 0
    v29 = v29
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v8
}
