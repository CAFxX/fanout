package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 int64, p1 int16, p2 bool) uint64 {
    var v7 int = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        switch "\n\t" {
        case "\n\t":
            p2 = p2
        case "a":
            var v9 int32 = 747
            _ = v9
        }
        _ = fi8
    }
    p0 = 1
    v10 := []int{v7, (v7 - (v7 + -2147483648))}
    switch "\n\t" {
    case "a":
        var v11 float64 = 1e10
        _ = v11
    case "":
        var v12 int16 = 0
        _ = v12
    case "\n\t":
        var v13 uint = 18446744073709551615
        _ = v13
    default:
        for fi14 := 0; fi14 < 16; fi14++ {
            v10 = append(v10, 4294967295)
            _ = fi14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 94
}
func genFn1(p0 int, p1 int16) uint16 {
    switch p0 {
    case 443:
        p0 -= 190
    }
    var v15 uint16 = 42
    p1 -= p1
    p1 = 0
    _ = p0
    _ = p1
    _ = v15
    return v15
}

func main() {
    var v1 bool = (0.0 <= (0.1 - -2.25))
    var v2 int = -9223372036854775808
    var v3 uint8 = 0
    var v4 int = v2
    var v5 uint32 = 2147483648
    var v6 uint = 2
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 int32 = 2147483647
        var v18 uint = v6
        fmt.Println(v2)
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            fmt.Println(fi20)
            fmt.Println(v4)
            v4 = fi19
            _ = fi20
        }
        _ = fi19
    }
    switch "hello" {
    case "hello":
        var v21 int32 = 42
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        for fi23 := 0; fi23 < 8; fi23++ {
            v2 += (v2 & fi23)
            _ = fi23
        }
        var v24 string = ("" + "a")
        _ = fi22
        _ = v24
    }
    v25 := GenS1{F0: "", F1: (^(v3)), F2: (v6 << (((v6 * v6)) & 63))}
    v26 := GenS1{F0: "\n\t", F1: v3, F2: (v6 ^ 454)}
    for fi27 := 0; fi27 < 6; fi27++ {
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 uint8 = 2
            fmt.Println(v3, v3)
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    var v30 int8 = int8(v2)
    switch v4 {
    case 32767:
        if v1 {
            for fi31 := 0; fi31 < 10; fi31++ {
                v1 = false
                v30 = (v30 >> ((v6) & 63))
                _ = fi31
            }
            fmt.Println(v1, v4, v1)
        } else {
            v32 := GenS1{F0: ((("foo bar" + ("foo bar" + "hello")) + "") + (("a" + "\n\t") + "a")), F1: (v3 ^ 0), F2: (v6 + v6)}
            _ = v32
        }
    case 127:
        if false {
            v4 = v2
            fmt.Println(v3, v3)
        }
    default:
        for fi33 := 0; fi33 < 3; fi33++ {
            fmt.Println(v1, v1, v5)
            _ = fi33
        }
    }
    v34 := GenS1{F0: "a", F1: 127, F2: v6}
    switch (v2 - v2) {
    case 32767:
        var v35 int32 = -78
        _ = v35
    }
    v36 := GenS2{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
