package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int64 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int32; F1 string }
func genFn0(p0 int64, p1 uint8, p2 int) string {
    var v4 int16 = -1
    p0 = (p0 | (p0 - (p0 << ((255) & 63))))
    p0 = int64(v4)
    switch 128 {
    case 32768:
        for fi5 := 0; fi5 < 12; fi5++ {
            if (fi5 <= p2) {
                fmt.Println(v4, p1, p0)
                fmt.Println(p2)
                for fi6 := 0; fi6 < 13; fi6++ {
                    switch "" {
                    case "\n\t":
                        switch ("" + "a") {
                        case "":
                            var v7 int64 = 1
                            _ = v7
                        default:
                            p0 = p0
                        }
                    case "hello":
                        v8 := GenS1{F0: 4294967295, F1: v4, F2: p0}
                        _ = v8
                    case "":
                        v4 = (^(v4))
                    default:
                        var v9 float64 = ((1.5 - float64(0.1)) + float64(-2.25))
                        _ = v9
                    }
                    var v10 int32 = -1
                    _ = fi6
                    _ = v10
                }
            } else {
                var v11 string = ""
                _ = v11
            }
            fmt.Println(p1, p0, p0)
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 bool = false
                var v14 int64 = (p0 & p0)
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi5
        }
    default:
        var v15 int = (p2 & p2)
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ("" + "\n\t")
}

func main() {
    var v1 bool = false
    var v2 int8 = -1
    v3 := make(map[int]uint32)
    for fi16 := 0; fi16 < 3; fi16++ {
        v1 = (-956 > -540)
        v17 := GenS2{F0: 1, F1: -711}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(v2, v1)
        if ((!(v2 > v2)) || (0 != 18446744073709551615)) {
            v19 := []int{9223372036854775807, 559, fi18}
            var v20 int = 546
            fmt.Println(v1, v2)
            _ = v19
            _ = v20
        }
        _ = fi18
    }
    v21 := GenS3{F0: -136}
    for fi22 := 0; fi22 < 12; fi22++ {
        if v1 {
            v2 = v2
            fmt.Println(v1)
            var v23 int8 = (-(v2))
            _ = v23
        }
        if false {
            fmt.Println(v1)
            v2 = (^(v2))
            var v24 string = ("\n\t" + "\n\t")
            _ = v24
        }
        _ = fi22
    }
    if ((653 <= 840) && v1) {
        v25 := GenS2{F0: 32767, F1: 32767}
        _ = v25
    } else {
        v1 = v1
        var v26 bool = v1
        var v27 string = "foo bar"
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        v29 := GenS4{F0: (false && (!((((fi28 & (fi28 + (^(fi28)))) > fi28) && v1) && (1 != -1)))), F1: -542}
        v2 = (v2 | v2)
        _ = fi28
        _ = v29
    }
    var v30 float32 = 3.14159
    v31 := GenS5{F0: 1, F1: ("foo bar" + ("a" + "a"))}
    var v32 float32 = v30
    v32 += ((v30 * v32) - ((0.0 * float32(0.1)) / (v32 + (v32 * -2.25))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
