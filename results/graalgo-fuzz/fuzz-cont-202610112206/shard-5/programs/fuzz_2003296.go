package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 uint64 }
func genFn0(p0 uint32, p1 uint32) string {
    if true {
        var v6 float64 = (1.5 + 1.5)
        var v7 int64 = -9223372036854775808
        _ = v6
        _ = v7
    } else {
        p0 = (p1 >> ((872) & 63))
    }
    var v8 float32 = 0.0
    var v9 uint32 = (p0 << ((255) & 63))
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return ("hello" + "a")
}
func genFn1(p0 uint16) int {
    if (-358 < 42) {
        if false {
            fmt.Println(p0)
            v10 := []int{2147483647, -9223372036854775808}
            _ = v10
        } else {
            v11 := GenS1{F0: 258}
            _ = v11
        }
        switch 1 {
        case -1:
            p0 = 1
        case -2147483648:
            fmt.Println(p0)
        }
    } else {
        var v12 int = -1
        if (p0 < (p0 - 2)) {
            v13 := GenS1{F0: 1}
            _ = v13
        }
        if (42 == -128) {
            var v14 uint32 = 1
            var v15 int = (v12 / ((v12) | 1))
            for fi16 := 0; fi16 < 9; fi16++ {
                p0 += uint16(v14)
                v17 := GenS2{F0: 0}
                _ = fi16
                _ = v17
            }
            _ = v14
            _ = v15
        } else {
            switch "foo bar" {
            case "a":
                v18 := GenS1{F0: 4294967295}
                _ = v18
            case "hello":
                for fi19 := 0; fi19 < 8; fi19++ {
                    v20 := []int{(fi19 * 268)}
                    _ = fi19
                    _ = v20
                }
            }
            v21 := []int{v12, v12, -2147483648}
            _ = v21
        }
        _ = v12
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := []int{815, 4294967295}
        _ = fi22
        _ = v23
    }
    v24 := []int{-9223372036854775808, 42}
    switch 127 {
    case 2147483647:
        var v25 uint64 = 4294967295
        _ = v25
    default:
        var v26 uint16 = (^(p0))
        _ = v26
    }
    _ = p0
    _ = v24
    return 32768
}

func main() {
    var v1 uint16 = 1
    var v2 string = "hello"
    var v3 float64 = ((-2.25 + 3.14159) + 3.14159)
    var v4 int8 = 23
    v5 := make(map[uint32]bool)
    v27 := []int{212, 9223372036854775807, 32768}
    _ = len(v27)
    v2 = (v2 + ("foo bar" + (("foo bar" + ("foo bar" + (v2 + (((v2 + v2) + ("\n\t" + "a")) + v2)))) + (v2 + ("" + (v2 + v2))))))
    v5[2147483648] = (!(42 <= -1))
    if false {
        v1 = 42
    } else {
        fmt.Println(v2)
        v28 := GenS1{F0: 4294967295}
        _ = v28
    }
    var v29 uint = 2
    if true {
        var v30 float64 = v3
        var v31 uint8 = 201
        _ = v30
        _ = v31
    } else {
        switch 4294967295 {
        case 32768:
            v32 := GenS3{F0: v4}
            _ = v32
        default:
            if false {
                var v33 uint8 = 127
                _ = v33
            } else {
                fmt.Println(v2)
            }
        }
    }
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    fmt.Println(v29, v4, v29)
    _ = len(v27)
    var v34 int16 = -1
    for fi35 := 0; fi35 < 1; fi35++ {
        v36 := GenS4{F0: (fi35 & 914), F1: uint64(v29)}
        if false {
            v1 = v1
            v29 -= (^(v29))
        }
        _ = fi35
        _ = v36
    }
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    delete(v5, 2147483648)
    switch 9223372036854775807 {
    case -1:
        fmt.Println(v34, v34)
    case -9223372036854775808:
        v4 = 127
    case 4294967295:
        if len(v27) > 0 {
            v27[0] = -2147483648
        }
    }
    if true {
        var v37 uint32 = 608
        var v38 float32 = (1e10 / 1.0)
        v1 -= 255
        _ = v37
        _ = v38
    }
    if (uint16(v29) <= 42) {
        switch (v2 + v2) {
        case "hello":
            v1 = 32767
        }
    }
    v2 = (v2 + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
