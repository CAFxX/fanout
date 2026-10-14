package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 int32 }
func genFn0(p0 uint8, p1 string, p2 int) int32 {
    fmt.Println(p2)
    v4 := []int{}
    if ((0 > 42) && (!false)) {
        for fi5 := 0; fi5 < 7; fi5++ {
            if false {
                if (true || (true && true)) {
                    var v6 int32 = 1
                    var v7 int = p2
                    v4 = append(v4, 128)
                    _ = v6
                    _ = v7
                } else {
                    v4 = append(v4, (^(fi5)))
                    v8 := GenS1{F0: ((((("hello" + "a") + "a") + (p1 + p1)) + "\n\t") + p1), F1: -299}
                    fmt.Println(p2, fi5, p2)
                    _ = v8
                }
                var v9 uint32 = 1
                for fi10 := 0; fi10 < 5; fi10++ {
                    p0 -= (p0 & p0)
                    var v11 uint64 = 609
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                _ = len(v4)
                for fi12 := 0; fi12 < 10; fi12++ {
                    p2 -= p2
                    _ = fi12
                }
                switch "a" {
                case "\n\t":
                    v13 := GenS2{F0: 4294967295, F1: true}
                    _ = v13
                case "a":
                    p2 -= p2
                }
            }
            switch p2 {
            case -2147483648:
                for fi14 := 0; fi14 < 7; fi14++ {
                    p2 -= 4294967295
                    switch "a" {
                    case "":
                        v15 := GenS1{F0: (((p1 + p1) + p1) + p1), F1: 1}
                        _ = v15
                    case "foo bar":
                        p1 = p1
                    }
                    _ = fi14
                }
            case 0:
                p1 = p1
            default:
                var v16 uint = 42
                _ = v16
            }
            if true {
                for fi17 := 0; fi17 < 17; fi17++ {
                    if true {
                        _ = len(v4)
                        fmt.Println(fi17)
                    } else {
                        fmt.Println(p2, fi17)
                        var v18 int64 = 212
                        if (!false) {
                            var v19 bool = (v18 < 0)
                            _ = v19
                        }
                        _ = v18
                    }
                    _ = len(v4)
                    _ = fi17
                }
                for fi20 := 0; fi20 < 18; fi20++ {
                    var v21 bool = (3.14159 <= ((0.0 / 1e10) * 3.14159))
                    _ = fi20
                    _ = v21
                }
            } else {
                switch (fi5 & (p2 % ((-202) | 1))) {
                case 1:
                    fmt.Println(fi5, p0, p0)
                case -1:
                    var v22 int16 = -1
                    _ = v22
                case 0:
                    p0 += p0
                default:
                    var v23 int64 = 42
                    _ = v23
                }
            }
            _ = fi5
        }
    } else {
        var v24 uint16 = 551
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -145
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[string]bool)
    var v3 int64 = -682
    v25 := []int{-1}
    for fi26 := 0; fi26 < 4; fi26++ {
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 int8 = -89
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    switch ("hello" + "foo bar") {
    case "a":
        v1 = (v1 + float64(1e10))
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        if (!true) {
            var v30 int16 = 1
            _ = v30
        }
        v31 := GenS3{F0: false}
        v25 = append(v25, fi29)
        _ = fi29
        _ = v31
    }
    switch 32767 {
    case 0:
        var v32 int = 455
        _ = v32
    default:
        for fi33 := 0; fi33 < 4; fi33++ {
            v3 -= (v3 | (v3 ^ v3))
            v2[""] = (548 > 259)
            _ = fi33
        }
    }
    v34 := GenS4{F0: -839, F1: -1}
    switch 9223372036854775807 {
    case 4294967295:
        v3 -= v3
    case -470:
        var v35 uint16 = 2
        _ = v35
    }
    var v36 uint8 = 111
    if (!(370 < 783)) {
        fmt.Println(v3, v1)
        fmt.Println(v36)
    } else {
        switch "hello" {
        case "a":
            for fi37 := 0; fi37 < 3; fi37++ {
                for fi38 := 0; fi38 < 18; fi38++ {
                    var v39 uint64 = 42
                    _ = fi38
                    _ = v39
                }
                var v40 int16 = int16(fi37)
                var v41 uint = 711
                _ = fi37
                _ = v40
                _ = v41
            }
        }
    }
    if len(v25) > 0 {
        v25[0] = 110
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v36
}
