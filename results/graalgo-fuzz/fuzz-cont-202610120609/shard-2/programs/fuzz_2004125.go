package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 string; F1 int8; F2 uint }
type GenS3 struct { F0 uint; F1 string; F2 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16; F1 uint }
func genFn0(p0 uint, p1 int16) int16 {
    switch "a" {
    case "\n\t":
        v6 := []int{127}
        _ = v6
    }
    var v7 string = ""
    _ = p0
    _ = p1
    _ = v7
    return (-(p1))
}

func main() {
    var v1 int64 = 1
    var v2 int8 = 83
    v3 := []uint32{2}
    var v4 uint16 = 693
    v5 := make(map[string]uint32)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 -= 9223372036854775807
    for fi8 := 0; fi8 < 2; fi8++ {
        v4 = v4
        _ = len(v3)
        v3 = append(v3, 1)
        _ = fi8
    }
    switch "hello" {
    case "\n\t":
        var v9 uint32 = 42
        _ = v9
    case "":
        switch 0 {
        case 42:
            var v10 uint8 = 0
            _ = v10
        }
    }
    v11 := GenS1{F0: 2, F1: -32768}
    if true {
        if true {
            fmt.Println(v1)
        } else {
            if len(v3) > 0 {
                v3[0] = 734
            }
        }
        var v12 uint = 968
        _ = v12
    } else {
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS2{F0: "hello", F1: (v2 + v2), F2: 702}
            var v15 string = "hello"
            switch ("hello" + v15) {
            case "hello":
                v4 = 556
            case "foo bar":
                for fi16 := 0; fi16 < 14; fi16++ {
                    v2 += (v2 % ((-1) | 1))
                    _ = fi16
                }
            default:
                v17 := GenS3{F0: 255, F1: v15, F2: 2}
                _ = v17
            }
            _ = fi13
            _ = v14
            _ = v15
        }
        if (false && (true || (42 == 0))) {
            for fi18 := 0; fi18 < 12; fi18++ {
                v1 = 9223372036854775807
                _ = fi18
            }
            fmt.Println(v1, v2)
            for fi19 := 0; fi19 < 14; fi19++ {
                var v20 float32 = 0.1
                _ = fi19
                _ = v20
            }
        } else {
            if true {
                v21 := GenS3{F0: 1, F1: "\n\t", F2: 481}
                for fi22 := 0; fi22 < 14; fi22++ {
                    v1 = v1
                    _ = fi22
                }
                v23 := GenS4{F0: false}
                _ = v21
                _ = v23
            } else {
                v24 := v5[""]
                v25 := GenS5{F0: 228}
                v26 := GenS6{F0: "\n\t"}
                _ = v24
                _ = v25
                _ = v26
            }
            v4 += 42
        }
    }
    _ = len(v5)
    switch 32767 {
    case -2147483648:
        var v27 uint32 = 0
        _ = v27
    default:
        v28 := GenS4{F0: true}
        _ = v28
    }
    v2 += v2
    switch ("hello" + "a") {
    case "\n\t":
        var v29 int = 32768
        _ = v29
    case "foo bar":
        v30 := GenS7{F0: v4, F1: 0}
        _ = v30
    case "hello":
        var v31 int32 = 996
        _ = v31
    }
    switch "a" {
    case "foo bar":
        v2 += (-(v2))
    case "hello":
        for fi32 := 0; fi32 < 10; fi32++ {
            v1 = (-(v1))
            _ = fi32
        }
    default:
        fmt.Println(v1, v1, v1)
    }
    var v33 uint64 = 217
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
