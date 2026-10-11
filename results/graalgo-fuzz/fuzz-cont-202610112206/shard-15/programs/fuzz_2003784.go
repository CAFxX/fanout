package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := []int{}
    var v2 int16 = 32767
    v3 := []int{}
    v4 := []string{}
    var v5 int8 = 51
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            fmt.Println(fi6, fi6, fi6)
            _ = fi7
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: 101, F1: fi8, F2: 42}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS2{F0: 1}
            var v12 int8 = v5
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi6
    }
    if (0 != (v5 & v5)) {
        for fi13 := 0; fi13 < 4; fi13++ {
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 float32 = -2.25
                fmt.Println(v2, fi13, v5)
                if len(v3) > 0 {
                    v3[0] = (^(fi13))
                }
                _ = fi14
                _ = v15
            }
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 uint64 = 396
                _ = fi16
                _ = v17
            }
            _ = fi13
        }
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        v5 = (v5 * v5)
        var v19 int16 = v2
        _ = fi18
        _ = v19
    }
    v3 = append(v3, 128)
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 int = 4294967295
        if (-2.25 > 0.1) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            switch "a" {
            case "":
                var v22 float64 = 1.5
                _ = v22
            default:
                var v23 int64 = 42
                _ = v23
            }
            v24 := GenS2{F0: 0}
            if len(v4) > 0 {
                v4[0] = ""
            }
            _ = v24
        }
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 uint16 = 786
            fmt.Println(v5)
            if (v26 != v26) {
                v21 -= 42
                v27 := GenS3{F0: 42, F1: 12}
                v5 -= (-(v5))
                var v28 int16 = (v2 >> ((787) & 63))
                _ = len(v4)
                var v29 int64 = 9223372036854775807
                v5 = 1
                v30 := GenS3{F0: 128, F1: 18446744073709551615}
                _ = v27
                _ = v28
                _ = v29
                _ = v30
            } else {
                var v31 float32 = 0.0
                _ = v31
            }
            var v32 uint16 = 1
            fmt.Println(fi25, v21, v32)
            _ = fi25
            _ = v26
            _ = v32
        }
        _ = fi20
        _ = v21
    }
    v2 -= 32767
    var v33 uint = 2
    var v34 float32 = 0.1
    v2 += v2
    for fi35 := 0; fi35 < 5; fi35++ {
        v36 := GenS3{F0: 0, F1: 0}
        if false {
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v34, fi35)
            case "foo bar":
                switch "a" {
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    var v37 int = (fi35 % ((fi35) | 1))
                    _ = v37
                }
            }
            v5 = (v5 >> ((v33) & 63))
        } else {
            v2 = v2
            _ = len(v4)
        }
        _ = fi35
        _ = v36
    }
    v2 += 42
    for fi38 := 0; fi38 < 15; fi38++ {
        v4 = append(v4, "a")
        fmt.Println(fi38, fi38)
        switch "" {
        case "hello":
            v3 = append(v3, (fi38 + fi38))
        case "\n\t":
            _ = len(v1)
        default:
            v5 += (-(v5))
        }
        _ = fi38
    }
    fmt.Println(v34, v5, v2)
    for fi39 := 0; fi39 < 14; fi39++ {
        v40 := GenS3{F0: 2, F1: 42}
        _ = fi39
        _ = v40
    }
    v41 := GenS2{F0: 1}
    var v42 bool = (false || true)
    for fi43 := 0; fi43 < 8; fi43++ {
        var v44 uint32 = 1
        _ = fi43
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v5
}
