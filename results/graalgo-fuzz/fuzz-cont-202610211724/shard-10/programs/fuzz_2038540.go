package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 uint }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16 }
func genFn0(p0 uint, p1 int8) uint64 {
    v5 := GenS1{F0: (p1 << ((42) & 63))}
    v6 := []int{128, -1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return uint64(p0)
}

func main() {
    var v1 uint64 = 882
    v2 := []int{1}
    v3 := []uint32{215}
    var v4 float64 = (0.0 + (0.0 - (3.14159 + 1.5)))
    switch 982 {
    case 127:
        switch "\n\t" {
        case "a":
            if (!true) {
                var v7 int8 = 1
                _ = v7
            } else {
                if ((genFn0(255, 42) < v1) || (!false)) {
                    v8 := GenS1{F0: -56}
                    switch ((("hello" + "\n\t") + "") + ("\n\t" + (("" + "foo bar") + ("" + "\n\t")))) {
                    case "\n\t":
                        v9 := GenS2{F0: ("foo bar" + "a"), F1: 0}
                        _ = v9
                    default:
                        var v10 float32 = 1.5
                        _ = v10
                    }
                    _ = v8
                } else {
                    fmt.Println(v4, v1)
                }
                v11 := GenS1{F0: 0}
                _ = v11
            }
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v12 := GenS2{F0: "", F1: 1}
            _ = v12
        }
    case 2147483647:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint8 = 150
            _ = fi13
            _ = v14
        }
    }
    v15 := GenS3{F0: 1, F1: -546, F2: 0}
    v16 := GenS4{F0: 18446744073709551615, F1: 42}
    _ = len(v3)
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 int16 = 1
        v4 -= float64(float32(v4))
        _ = fi17
        _ = v18
    }
    v1 = v1
    v19 := GenS1{F0: -128}
    v1 += genFn0(18446744073709551615, -64)
    v4 = (v4 - v4)
    var v20 int = -758
    switch "a" {
    case "hello":
        v20 = 9223372036854775807
    case "\n\t":
        var v21 uint = 632
        _ = v21
    case "":
        for fi22 := 0; fi22 < 18; fi22++ {
            if true {
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 uint8 = 1
                    _ = fi23
                    _ = v24
                }
                fmt.Println(fi22)
                var v25 int32 = 435
                _ = v25
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi22
        }
    }
    v1 = genFn0(65535, 100)
    if true {
        if true {
            fmt.Println(v4, v1, v4)
            var v26 string = "a"
            var v27 float32 = 1e10
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 1; fi28++ {
                v29 := GenS5{F0: 42}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi28
                _ = v29
            }
            v30 := GenS6{F0: 1}
            var v31 bool = true
            _ = v30
            _ = v31
        }
    } else {
        _ = len(v2)
        for fi32 := 0; fi32 < 13; fi32++ {
            var v33 uint8 = 1
            v33 += 128
            _ = fi32
            _ = v33
        }
    }
    switch "hello" {
    case "\n\t":
        v4 += float64(1e10)
    case "a":
        for fi34 := 0; fi34 < 2; fi34++ {
            v20 = (v20 << ((42) & 63))
            for fi35 := 0; fi35 < 13; fi35++ {
                var v36 string = "foo bar"
                fmt.Println(v36)
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        var v38 string = ("foo bar" + "")
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
