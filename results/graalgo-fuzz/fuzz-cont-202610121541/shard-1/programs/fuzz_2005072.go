package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0(p0 int32, p1 int32, p2 uint16) int {
    switch "" {
    case "":
        switch -1 {
        case 32768:
            v6 := GenS1{F0: 175}
            _ = v6
        case 128:
            var v7 uint64 = 2
            _ = v7
        }
    case "\n\t":
        var v8 uint8 = 2
        _ = v8
    case "a":
        var v9 uint16 = p2
        _ = v9
    default:
        var v10 int = -597
        _ = v10
    }
    var v11 int64 = 9223372036854775807
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(p0, fi12)
        var v13 float64 = 0.1
        fmt.Println(p0)
        _ = fi12
        _ = v13
    }
    p1 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 4294967295
}
func genFn1() uint {
    if (-32768 >= 1) {
        var v14 uint16 = 753
        v15 := []int{-1, -2147483648}
        _ = v14
        _ = v15
    }
    var v16 int32 = 1
    _ = v16
    return 4294967295
}

func main() {
    var v1 int64 = -1
    var v2 int16 = 42
    v3 := []uint32{232, 2}
    v4 := []string{("\n\t" + "hello"), "hello"}
    var v5 float64 = (1.5 / float64(3.14159))
    v17 := GenS2{F0: 0}
    v1 -= v1
    v18 := GenS3{F0: ("hello" + "foo bar"), F1: 4294967295}
    var v19 uint32 = 751
    switch "foo bar" {
    case "hello":
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 bool = (!false)
            v22 := GenS2{F0: (^(v2))}
            _ = fi20
            _ = v21
            _ = v22
        }
    case "foo bar":
        v5 = v5
    case "":
        var v23 int = 128
        _ = v23
    }
    var v24 int64 = 856
    v2 += (v2 * -1)
    if true {
        for fi25 := 0; fi25 < 19; fi25++ {
            v2 = v2
            v19 = 71
            _ = fi25
        }
        if false {
            v26 := GenS2{F0: v2}
            v19 -= 4294967295
            if (false || false) {
                var v27 int64 = 42
                v2 += 32767
                _ = v27
            }
            _ = v26
        }
        switch 32767 {
        case -9223372036854775808:
            for fi28 := 0; fi28 < 6; fi28++ {
                v24 += v24
                var v29 uint8 = 42
                _ = fi28
                _ = v29
            }
        }
    }
    var v30 uint64 = 1
    v30 = v30
    for fi31 := 0; fi31 < 4; fi31++ {
        v32 := GenS2{F0: (v2 ^ 1000)}
        var v33 int = fi31
        _ = fi31
        _ = v32
        _ = v33
    }
    v1 += int64(v30)
    fmt.Println(v24, v24, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
