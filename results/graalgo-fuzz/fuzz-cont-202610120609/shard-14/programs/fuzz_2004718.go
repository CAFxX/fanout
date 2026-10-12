package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint16; F1 int16 }

func main() {
    var v1 int32 = -537
    var v2 uint16 = 324
    var v3 float32 = (3.14159 / (-2.25 / (3.14159 * 0.1)))
    if (false || (32767 == 924)) {
        if true {
            var v4 bool = (42 != 4294967295)
            switch 32768 {
            case 9223372036854775807:
                var v5 float32 = (v3 * (0.1 * v3))
                _ = v5
            }
            var v6 uint64 = 18446744073709551615
            _ = v4
            _ = v6
        } else {
            if false {
                var v7 uint = uint(v2)
                v8 := []int{-9223372036854775808, 9223372036854775807, 127}
                _ = v7
                _ = v8
            } else {
                fmt.Println(v2, v1, v1)
                if ((v3 - ((((1e10 + float32(0.0)) - v3) + 0.1) - float32((0.1 - 0.0)))) != v3) {
                    if (false || true) {
                        v3 -= (0.1 + float32(0.0))
                    }
                } else {
                    v9 := GenS1{F0: -1, F1: ""}
                    _ = v9
                }
                for fi10 := 0; fi10 < 7; fi10++ {
                    switch ("\n\t" + "a") {
                    case "a":
                        v1 = v1
                    case "foo bar":
                        if (!(true || (-2.25 <= 1e10))) {
                            v11 := GenS2{F0: 1, F1: (2 == 0), F2: v2}
                            v2 += 290
                            switch "" {
                            case "foo bar":
                                switch "foo bar" {
                                case "foo bar":
                                    var v12 int8 = 42
                                    _ = v12
                                }
                            }
                            _ = v11
                        }
                    case "hello":
                        v13 := GenS1{F0: -44, F1: "hello"}
                        _ = v13
                    }
                    _ = fi10
                }
            }
            if false {
                v14 := []int{4294967295, 383, -2147483648, 127}
                if false {
                    fmt.Println(v2, v2)
                    var v15 int64 = 607
                    switch 32767 {
                    case 128:
                        if (((!false) && (false && (!((709 > 65535) && (((127 >= 127) && true) || false))))) || (!("" == "foo bar"))) {
                            v16 := GenS3{F0: 127, F1: 127}
                            _ = v16
                        }
                    case -309:
                        v3 += v3
                    case 9223372036854775807:
                        v17 := GenS1{F0: 101, F1: (("hello" + "foo bar") + "a")}
                        _ = v17
                    }
                    _ = v15
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        v3 = 0.0
                    case "hello":
                        fmt.Println(v3, v2, v3)
                    default:
                        v18 := GenS1{F0: -22, F1: ("\n\t" + "a")}
                        _ = v18
                    }
                }
                _ = v14
            }
            var v19 int64 = -318
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        v3 = v3
        var v21 uint16 = v2
        _ = fi20
        _ = v21
    }
    v22 := []int{}
    var v23 bool = (2147483648 <= 0)
    v23 = (v1 <= v1)
    v22 = append(v22, 9223372036854775807)
    v23 = true
    for fi24 := 0; fi24 < 8; fi24++ {
        for fi25 := 0; fi25 < 11; fi25++ {
            var v26 int = fi24
            var v27 bool = ((("\n\t" + (("foo bar" + (("a" + ("hello" + "hello")) + ("hello" + "foo bar"))) + ("a" + ("\n\t" + ("a" + "\n\t"))))) + "") == "hello")
            _ = fi25
            _ = v26
            _ = v27
        }
        var v28 int32 = (v1 + (v1 % (((v1 % (((v1 % ((v1) | 1))) | 1))) | 1)))
        switch "" {
        case "a":
            if (!(0.0 != 0.0)) {
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
                v29 := GenS4{F0: v2, F1: -32768}
                var v30 float64 = -2.25
                _ = v29
                _ = v30
            }
        }
        _ = fi24
        _ = v28
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        v23 = (v23 || v23)
        var v32 int8 = -20
        if v23 {
            v33 := GenS3{F0: 494, F1: (fi31 / ((fi31) | 1))}
            v1 += (v1 * (v1 / ((2147483647) | 1)))
            v3 -= v3
            _ = v33
        }
        _ = fi31
        _ = v32
    }
    v22 = append(v22, 92)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
