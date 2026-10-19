package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 610
    var v3 float64 = 0.1
    switch int(v2) {
    case 623:
        if true {
            fmt.Println(v3, v3)
        } else {
            v3 -= 1e10
            v4 := GenS1{F0: 9223372036854775807}
            _ = v4
        }
    }
    v5 := []int{}
    v3 = 1.5
    v6 := GenS2{F0: -1, F1: 255}
    if (v2 != v2) {
        for fi7 := 0; fi7 < 2; fi7++ {
            v3 = ((float64(0.1) * (0.1 + 1.5)) - (1e10 - v3))
            switch ("a" + "hello") {
            case "hello":
                var v8 uint8 = 1
                _ = v8
            default:
                fmt.Println(v3, v2)
            }
            v2 = -1
            _ = fi7
        }
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := GenS1{F0: 0}
            _ = fi9
            _ = v10
        }
        var v11 int64 = 0
        _ = v11
    } else {
        for fi12 := 0; fi12 < 19; fi12++ {
            _ = len(v1)
            var v13 uint32 = 224
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi12
            _ = v13
        }
        switch -1 {
        case 0:
            v2 -= -1
        }
        _ = len(v1)
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if true {
            v3 -= v3
            v2 = (^(v2))
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 int8 = 0
                var v17 uint8 = 42
                _ = fi15
                _ = v16
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 4; fi18++ {
                v3 = (float64((3.14159 + 1e10)) - ((v3 / v3) - v3))
                var v19 uint16 = 673
                switch "foo bar" {
                case "hello":
                    var v20 int32 = v2
                    _ = v20
                case "a":
                    var v21 int = fi18
                    switch "a" {
                    case "a":
                        var v22 int32 = v2
                        _ = v22
                    case "foo bar":
                        _ = len(v5)
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v21
                case "foo bar":
                    delete(v1, "a")
                default:
                    v23 := GenS1{F0: 1}
                    _ = v23
                }
                _ = fi18
                _ = v19
            }
            if (!false) {
                switch "\n\t" {
                case "a":
                    var v24 int64 = -9223372036854775808
                    _ = v24
                }
                var v25 int64 = 1
                for fi26 := 0; fi26 < 17; fi26++ {
                    v2 = -169
                    var v27 uint32 = 383
                    _ = fi26
                    _ = v27
                }
                _ = v25
            } else {
                fmt.Println(v2, v2, v2)
            }
            v28 := v1[("hello" + "\n\t")]
            _ = v28
        }
        fmt.Println(v2)
        _ = fi14
    }
    v1[("hello" + ("hello" + ("\n\t" + ("foo bar" + "hello"))))] = 2
    var v29 uint = 42
    var v30 int64 = 1
    for fi31 := 0; fi31 < 1; fi31++ {
        for fi32 := 0; fi32 < 11; fi32++ {
            v30 += v30
            var v33 int32 = (-(v2))
            fmt.Println(v3, fi32)
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    v2 -= v2
    for fi34 := 0; fi34 < 18; fi34++ {
        var v35 float64 = ((v3 / v3) + v3)
        _ = fi34
        _ = v35
    }
    var v36 float64 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v5
    _ = v6
}
