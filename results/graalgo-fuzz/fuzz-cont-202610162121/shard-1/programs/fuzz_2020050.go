package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int32; F1 bool; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int8 }

func main() {
    var v1 uint = 1
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    var v4 uint = (v1 & v1)
    var v5 int8 = -93
    v6 := []int{}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint64 = 18446744073709551615
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "\n\t":
        var v9 float32 = 0.1
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS1{F0: 65535, F1: (v4 - v4)}
        delete(v2, "a")
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 int32 = 0
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v14 := GenS1{F0: 1, F1: (v4 ^ v1)}
    var v15 float64 = 1e10
    _ = len(v3)
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 bool = true
        v5 += v5
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: -1, F1: false, F2: v1}
    var v19 uint16 = 255
    switch 128 {
    case 1:
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 bool = ((false || (2 <= 1)) || (!false))
            _ = fi20
            _ = v21
        }
    case 109:
        for fi22 := 0; fi22 < 5; fi22++ {
            _ = len(v6)
            fmt.Println(v19, v1, v1)
            v23 := GenS1{F0: v19, F1: (v1 + v1)}
            _ = fi22
            _ = v23
        }
    case 0:
        if false {
            v24 := GenS3{F0: 2}
            _ = v24
        } else {
            for fi25 := 0; fi25 < 6; fi25++ {
                fmt.Println(v5, v15, v15)
                _ = fi25
            }
        }
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 int8 = (v5 | 67)
            switch "foo bar" {
            case "a":
                if false {
                    v4 += (v4 * v1)
                    v6 = append(v6, fi27)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            default:
                var v29 bool = (-1 != 42)
                _ = v29
            }
            _ = fi27
            _ = v28
        }
        switch int(v19) {
        case 0:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case 32768:
            switch "\n\t" {
            case "a":
                if (511 >= v4) {
                    fmt.Println(v15)
                    v30 := GenS4{F0: 6, F1: v5}
                    _ = v30
                } else {
                    v19 = (v19 + (v19 % (((v19 & v19)) | 1)))
                }
            }
        case -890:
            if (!true) {
                var v31 bool = ("hello" <= "foo bar")
                _ = v31
            }
        default:
            for fi32 := 0; fi32 < 14; fi32++ {
                if (2 >= (v19 << (((v1 % ((v4) | 1))) & 63))) {
                    var v33 uint16 = 390
                    fmt.Println(v19)
                    _ = v33
                }
                _ = fi32
            }
        }
        if false {
            v2[("" + "\n\t")] = "a"
        } else {
            fmt.Println(v4, v5)
            v15 -= v15
        }
        _ = fi26
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        for fi35 := 0; fi35 < 8; fi35++ {
            _ = len(v2)
            _ = fi35
        }
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
