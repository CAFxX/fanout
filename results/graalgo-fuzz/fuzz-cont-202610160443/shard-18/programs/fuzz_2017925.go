package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 bool, p1 uint16, p2 uint8) uint64 {
    p1 -= p1
    p0 = p0
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1() int {
    if true {
        var v5 int64 = -1
        v6 := GenS1{F0: 389}
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 13; fi7++ {
            fmt.Println(fi7)
            _ = fi7
        }
        var v8 float64 = 1e10
        switch "foo bar" {
        case "a":
            if (65535 == 18446744073709551615) {
                v9 := []int{815, -2147483648, 379, 2147483647}
                v8 -= v8
                _ = v9
            }
        case "hello":
            var v10 string = "hello"
            _ = v10
        default:
            v11 := GenS1{F0: 1}
            _ = v11
        }
        _ = v8
    }
    if (0 == 2) {
        var v12 bool = (-1 >= -557)
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := []int{241, fi13}
        _ = fi13
        _ = v14
    }
    var v15 int16 = 0
    v15 = v15
    _ = v15
    return int(v15)
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = (1.5 + 0.0)
    var v3 int16 = -1
    v4 := make(map[uint32]string)
    v16 := GenS2{F0: "a", F1: 111, F2: 0}
    var v17 uint8 = 126
    v3 = v3
    var v18 int16 = (v3 & (v3 | 254))
    var v19 uint64 = 18446744073709551615
    v20 := GenS1{F0: 255}
    fmt.Println(v3, v3)
    switch "" {
    case "hello":
        v18 -= v18
    case "foo bar":
        fmt.Println(v17, v19)
    }
    v19 = 2
    v21 := GenS2{F0: "a", F1: (v17 | (v17 - (v17 ^ v17))), F2: 0}
    v22 := GenS1{F0: 511}
    switch 620 {
    case 9223372036854775807:
        v23 := GenS3{F0: (v19 << ((1) & 63))}
        _ = v23
    case -191:
        v3 = (-(v18))
    }
    fmt.Println(v19, v17)
    switch "foo bar" {
    case "\n\t":
        if false {
            switch ("hello" + "foo bar") {
            case "foo bar":
                v24 := []int{}
                _ = v24
            default:
                v25 := []int{}
                _ = v25
            }
            v17 = (v17 / (((v17 >> ((42) & 63))) | 1))
            v19 = v19
        } else {
            v18 -= v3
            if false {
                switch -96 {
                case 4294967295:
                    var v26 uint16 = 604
                    _ = v26
                default:
                    v27 := GenS1{F0: 65535}
                    _ = v27
                }
                fmt.Println(v2)
                v28 := []int{}
                _ = v28
            } else {
                v2 += v2
                for fi29 := 0; fi29 < 14; fi29++ {
                    var v30 string = "\n\t"
                    for fi31 := 0; fi31 < 13; fi31++ {
                        v3 = v3
                        v17 = (v17 >> ((255) & 63))
                        var v32 uint64 = 42
                        v18 += (v3 & 42)
                        v4[930] = "\n\t"
                        _ = fi31
                        _ = v32
                    }
                    var v33 uint = 42
                    _ = fi29
                    _ = v30
                    _ = v33
                }
            }
        }
    case "foo bar":
        fmt.Println(v3)
    case "hello":
        var v34 uint8 = v17
        _ = v34
    default:
        if (2 == 636) {
            if true {
                v17 = v17
                fmt.Println(v19, v18, v2)
                v35 := []int{}
                _ = v35
            } else {
                v3 -= int16(v19)
                v36 := GenS2{F0: "", F1: 42, F2: -266}
                _ = v36
            }
            var v37 int8 = 1
            var v38 string = ""
            _ = v37
            _ = v38
        }
    }
    var v39 int = 9223372036854775807
    v18 += v18
    switch (-(v39)) {
    case 32768:
        v2 -= 3.14159
    case 4294967295:
        var v40 uint64 = (^(v19))
        _ = v40
    case -570:
        v41 := v1[(v39 + (^(v39)))]
        _ = v41
    default:
        var v42 int8 = -128
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v39
    _ = v4
}
