package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 string; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[uint32]bool)
    var v3 float32 = float32((1e10 - -2.25))
    if true {
        switch "" {
        case "hello":
            delete(v2, v1)
        case "":
            if (1.5 > 0.1) {
                if true {
                    var v4 float32 = v3
                    _ = v4
                }
                delete(v2, v1)
            } else {
                switch "hello" {
                case "":
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        v5 := []int{127}
                        _ = v5
                    default:
                        var v6 uint = 18446744073709551615
                        _ = v6
                    }
                case "a":
                    _ = len(v2)
                case "\n\t":
                    v3 += ((v3 / (3.14159 + (0.1 + (v3 / float32(((float64(1.5) + (-2.25 - 3.14159)) - 1e10)))))) + v3)
                default:
                    fmt.Println(v1)
                }
                for fi7 := 0; fi7 < 10; fi7++ {
                    v1 = v1
                    switch "hello" {
                    case "hello":
                        v2[2] = true
                    case "\n\t":
                        v1 = (^(v1))
                    }
                    for fi8 := 0; fi8 < 14; fi8++ {
                        v1 -= 48
                        _ = fi8
                    }
                    _ = fi7
                }
                for fi9 := 0; fi9 < 10; fi9++ {
                    fmt.Println(v3, v3, v1)
                    _ = fi9
                }
            }
        default:
            v10 := []int{}
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            v1 -= (v1 | (v1 << ((114) & 63)))
            _ = fi11
        }
        var v12 int16 = -32768
        v13 := []int{127}
        _ = v12
        _ = v13
    }
    fmt.Println(v1, v3)
    switch 2147483647 {
    case 9223372036854775807:
        var v14 bool = ((false || true) && false)
        _ = v14
    case 32768:
        var v15 uint32 = v1
        _ = v15
    case 32767:
        if false {
            v16 := GenS1{F0: 843}
            _ = v16
        }
    }
    v17 := GenS1{F0: 0}
    if (((v1 << ((2) & 63)) < v1) || (false && (65535 != 0))) {
        if true {
            v1 += v1
            if false {
                fmt.Println(v1, v1)
                v3 = v3
                if false {
                    v18 := []int{-543, 32767, 319, -358}
                    var v19 string = ("a" + ("a" + "a"))
                    _ = v18
                    _ = v19
                }
            } else {
                fmt.Println(v3, v3)
                v3 = 1.5
                v1 = v1
            }
        }
    } else {
        delete(v2, v1)
    }
    _ = len(v2)
    v20 := GenS2{F0: 42, F1: ("foo bar" + "a"), F2: 431}
    var v21 string = "a"
    switch "hello" {
    case "a":
        v22 := []int{9223372036854775807, 32767}
        _ = v22
    }
    if false {
        if true {
            var v23 int32 = 0
            v24 := []int{}
            for fi25 := 0; fi25 < 5; fi25++ {
                for fi26 := 0; fi26 < 19; fi26++ {
                    if ((v1 | (v1 >> ((4294967295) & 63))) <= v1) {
                        v23 += v23
                        v27 := GenS3{F0: 0}
                        _ = v27
                    }
                    var v28 int = -9223372036854775808
                    var v29 uint = 1
                    _ = fi26
                    _ = v28
                    _ = v29
                }
                v30 := v2[v1]
                _ = fi25
                _ = v30
            }
            _ = v23
            _ = v24
        } else {
            if false {
                switch -297 {
                case 128:
                    delete(v2, v1)
                default:
                    v31 := GenS3{F0: 2}
                    _ = v31
                }
                for fi32 := 0; fi32 < 2; fi32++ {
                    var v33 float64 = -2.25
                    v3 += float32(float64(v3))
                    _ = fi32
                    _ = v33
                }
            }
            v1 = v1
            var v34 int32 = 1
            _ = v34
        }
        v35 := []int{153}
        fmt.Println(v1, v1)
        _ = v35
    } else {
        v3 = (-2.25 * 1e10)
    }
    var v36 uint8 = 128
    if (-1 <= -128) {
        v37 := GenS4{F0: v36}
        _ = v37
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        fmt.Println(v36)
        v39 := []int{485, 4294967295}
        v36 += 0
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v36
}
