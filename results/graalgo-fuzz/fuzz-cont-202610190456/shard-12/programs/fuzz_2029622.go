package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint8; F1 string }
func genFn0(p0 string, p1 int64) uint {
    p0 = (("a" + ((p0 + ((("foo bar" + "foo bar") + p0) + p0)) + (p0 + (p0 + (p0 + ""))))) + "foo bar")
    v5 := []int{9223372036854775807, 4294967295, 462}
    _ = p0
    _ = p1
    _ = v5
    return 334
}
func genFn1() int16 {
    v6 := []int{-343, -326, 426, 4294967295}
    var v7 int8 = -1
    _ = v6
    _ = v7
    return 1000
}

func main() {
    var v1 uint32 = 602
    var v2 uint = 831
    var v3 int32 = 42
    var v4 string = "a"
    switch 0 {
    case 32767:
        v3 -= v3
    }
    v8 := GenS1{F0: (-(v3))}
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint16 = 1
        for fi11 := 0; fi11 < 10; fi11++ {
            switch "hello" {
            case "foo bar":
                fmt.Println(fi11, v10, fi9)
            case "\n\t":
                var v12 float64 = (3.14159 + 1e10)
                v1 = (v1 * v1)
                v13 := []int{865, -9223372036854775808}
                _ = v12
                _ = v13
            default:
                v1 = v1
            }
            var v14 uint16 = (v10 + (^(v10)))
            _ = fi11
            _ = v14
        }
        switch v4 {
        case "\n\t":
            v15 := GenS2{F0: 42, F1: 703}
            _ = v15
        case "":
            switch (fi9 / ((fi9) | 1)) {
            case 32767:
                v16 := []int{}
                _ = v16
            }
        case "hello":
            v4 = (v4 + ("\n\t" + v4))
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v2, v4, v2)
    var v17 int16 = 42
    for fi18 := 0; fi18 < 2; fi18++ {
        v19 := GenS3{F0: -59, F1: (v3 & int32(v1))}
        _ = fi18
        _ = v19
    }
    var v20 uint = (v2 >> (((v2 * v2)) & 63))
    for fi21 := 0; fi21 < 18; fi21++ {
        v22 := []int{(fi21 % ((fi21) | 1)), fi21, int(v1)}
        v4 = v4
        v23 := GenS4{F0: 193, F1: "foo bar"}
        _ = fi21
        _ = v22
        _ = v23
    }
    if true {
        var v24 uint8 = uint8(v1)
        v25 := []int{}
        _ = v24
        _ = v25
    } else {
        if (!(42 == v1)) {
            v2 = 102
            var v26 uint = (v20 >> ((v20) & 63))
            _ = v26
        } else {
            v27 := []int{}
            _ = v27
        }
        var v28 int8 = 108
        _ = v28
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        v30 := GenS3{F0: 42, F1: v3}
        v2 -= v2
        _ = fi29
        _ = v30
    }
    if true {
        var v31 int = 2147483647
        _ = v31
    } else {
        var v32 int8 = -4
        _ = v32
    }
    var v33 string = (v4 + v4)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v4
    _ = v8
}
