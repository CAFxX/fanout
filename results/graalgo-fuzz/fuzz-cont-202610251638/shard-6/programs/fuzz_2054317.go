package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint64 }

func main() {
    var v1 float64 = (0.1 - 0.0)
    var v2 float64 = v1
    var v3 uint = 42
    var v4 uint = 18446744073709551615
    for fi5 := 0; fi5 < 17; fi5++ {
        switch "foo bar" {
        case "hello":
            v1 -= (-2.25 - v1)
        case "\n\t":
            v6 := []int{(fi5 / (((fi5 | (-(fi5)))) | 1)), (fi5 % ((fi5) | 1)), 0}
            _ = v6
        default:
            var v7 uint16 = 760
            _ = v7
        }
        var v8 int = (fi5 - (-(fi5)))
        _ = fi5
        _ = v8
    }
    var v9 int = -709
    if false {
        if (false && ("foo bar" >= "hello")) {
            v10 := GenS1{F0: 2, F1: (v3 + v3)}
            v11 := []int{}
            _ = v10
            _ = v11
        }
        v12 := GenS2{F0: 0}
        _ = v12
    } else {
        var v13 uint16 = 90
        v14 := GenS2{F0: v13}
        v15 := []int{}
        _ = v13
        _ = v14
        _ = v15
    }
    fmt.Println(v3)
    switch "a" {
    case "\n\t":
        var v16 int64 = -693
        _ = v16
    case "foo bar":
        v17 := GenS2{F0: 551}
        _ = v17
    }
    switch (v9 | 32767) {
    case 1:
        var v18 uint8 = 42
        _ = v18
    }
    v19 := []int{}
    var v20 bool = ("a" < ("a" + "foo bar"))
    if v20 {
        for fi21 := 0; fi21 < 7; fi21++ {
            for fi22 := 0; fi22 < 16; fi22++ {
                v2 -= 1.5
                switch "a" {
                case "":
                    if ((1.5 <= (3.14159 / 1.0)) && (true && (573 > 42))) {
                        if v20 {
                            var v23 int16 = 32767
                            v24 := GenS1{F0: 128, F1: v3}
                            _ = v23
                            _ = v24
                        } else {
                            v9 -= (v9 + (^(v9)))
                            v20 = true
                            v25 := GenS1{F0: 127, F1: (v4 + 2)}
                            _ = v25
                        }
                        var v26 uint64 = 894
                        var v27 float64 = (float64(-2.25) * (v1 - ((float64(((-2.25 / -2.25) - (1e10 * 0.0))) * v1) - ((1.5 * 0.1) * v1))))
                        v4 = 65535
                        v28 := GenS3{F0: 127, F1: 822, F2: (v26 + v26)}
                        fmt.Println(v2)
                        fmt.Println(fi22)
                        _ = v26
                        _ = v27
                        _ = v28
                    } else {
                        var v29 bool = false
                        if v20 {
                            v19 = append(v19, v9)
                            _ = len(v19)
                        }
                        _ = v29
                    }
                default:
                    var v30 uint = 292
                    _ = v30
                }
                v3 += (v4 | v4)
                _ = fi22
            }
            _ = fi21
        }
        v1 = ((v2 + 0.0) / ((3.14159 - v2) - 3.14159))
    }
    var v31 int8 = -1
    if true {
        v2 -= (1.5 * v2)
        fmt.Println(v9, v2)
        var v32 uint16 = 2
        _ = v32
    } else {
        if (true && v20) {
            v31 -= 100
            var v33 int32 = -1
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        _ = fi34
    }
    if v20 {
        var v35 uint = 4294967295
        for fi36 := 0; fi36 < 3; fi36++ {
            var v37 float64 = float64(3.14159)
            _ = fi36
            _ = v37
        }
        _ = v35
    }
    v38 := GenS1{F0: 127, F1: v4}
    v9 = v9
    if v20 {
        v31 += v31
        if v20 {
            var v39 uint32 = 42
            v1 = float64(1.5)
            _ = v39
        }
        v20 = ((v4 <= v3) && v20)
    }
    v3 = (v3 % ((118) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v38
    _ = v4
    _ = v9
}
