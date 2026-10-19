package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 int8; F2 uint }
type GenS3 struct { F0 string; F1 uint; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 int }
func genFn0(p0 int16, p1 bool, p2 int16) int64 {
    v4 := []int{}
    switch 4294967295 {
    case 0:
        switch 32767 {
        case 42:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 0:
            var v5 bool = p1
            _ = v5
        }
    case -2147483648:
        var v6 int8 = -1
        _ = v6
    case 128:
        if (!p1) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            p1 = p1
        }
    default:
        var v7 int = -1
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        p2 = p2
        var v9 uint32 = 56
        fmt.Println(v9, fi8)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    var v3 uint16 = 1
    if true {
        _ = len(v1)
        v1[("hello" + "foo bar")] = (42 < 65535)
    } else {
        for fi10 := 0; fi10 < 12; fi10++ {
            fmt.Println(fi10, fi10)
            _ = fi10
        }
        var v11 uint = 906
        var v12 int64 = -777
        _ = v11
        _ = v12
    }
    v13 := []int{-998, -9223372036854775808, -2147483648}
    for fi14 := 0; fi14 < 18; fi14++ {
        v3 = 32767
        delete(v1, "a")
        for fi15 := 0; fi15 < 18; fi15++ {
            fmt.Println(fi15, fi14, fi14)
            fmt.Println(fi14, fi14, fi15)
            _ = len(v1)
            var v16 uint16 = v3
            var v17 float64 = -2.25
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    if ((!false) || (!(340 != 42))) {
        if false {
            var v18 int8 = -29
            switch "hello" {
            case "\n\t":
                v19 := GenS1{F0: (v18 + v18)}
                _ = v19
            case "foo bar":
                for fi20 := 0; fi20 < 19; fi20++ {
                    var v21 int = (fi20 | (fi20 - 664))
                    v3 += v3
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v23 := GenS1{F0: (v18 ^ (v18 - (v18 - (v18 - (v18 & (v18 << ((65535) & 63)))))))}
                        _ = fi22
                        _ = v23
                    }
                    _ = fi20
                    _ = v21
                }
            }
            var v24 uint8 = 42
            _ = v18
            _ = v24
        } else {
            v25 := GenS2{F0: 2147483647, F1: 1, F2: 439}
            fmt.Println(v3)
            v26 := GenS3{F0: "", F1: 0, F2: 100}
            _ = v25
            _ = v26
        }
    }
    if true {
        var v27 bool = (!true)
        v28 := GenS3{F0: "hello", F1: 972, F2: 127}
        var v29 int8 = -14
        _ = v27
        _ = v28
        _ = v29
    } else {
        var v30 float32 = (0.0 + -2.25)
        v31 := GenS1{F0: -78}
        _ = v30
        _ = v31
    }
    _ = len(v2)
    v3 = (^(v3))
    var v32 string = "foo bar"
    v33 := GenS4{F0: 42}
    for fi34 := 0; fi34 < 16; fi34++ {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = fi34
    }
    v32 = v32
    var v35 string = v32
    v36 := GenS5{F0: false, F1: 0}
    var v37 float32 = float32((-2.25 / (1.5 - (3.14159 * 1.5))))
    var v38 int64 = -1
    if false {
        fmt.Println(v37)
        v35 = v35
        for fi39 := 0; fi39 < 4; fi39++ {
            v40 := GenS2{F0: -2147483648, F1: 42, F2: 1}
            _ = fi39
            _ = v40
        }
    } else {
        var v41 uint = 611
        _ = v41
    }
    v42 := GenS3{F0: (v32 + "\n\t"), F1: 255, F2: 127}
    switch -1 {
    case 0:
        v32 = (v32 + v32)
    case 32767:
        v3 -= (v3 * v3)
    case -2147483648:
        for fi43 := 0; fi43 < 5; fi43++ {
            for fi44 := 0; fi44 < 15; fi44++ {
                var v45 int16 = 984
                _ = fi44
                _ = v45
            }
            _ = fi43
        }
    default:
        var v46 uint = 255
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v42
}
