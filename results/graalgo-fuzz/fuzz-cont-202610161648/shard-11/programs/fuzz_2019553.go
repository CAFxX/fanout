package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8 }
func genFn0(p0 int, p1 int32, p2 string) int16 {
    if (false && (28 < 1)) {
        v7 := GenS1{F0: 18446744073709551615}
        _ = v7
    }
    if false {
        v8 := []int{(p0 / ((-9223372036854775808) | 1)), p0, 0}
        if true {
            v9 := GenS2{F0: ((p1 & (p1 << ((65535) & 63))) >= 930)}
            v10 := GenS1{F0: 56}
            _ = v10
            _ = v9
        }
        _ = v8
    }
    v11 := GenS1{F0: 18446744073709551615}
    switch p0 {
    case 4294967295:
        switch (p2 + p2) {
        case "\n\t":
            var v12 string = (((p2 + "\n\t") + (p2 + ((p2 + (p2 + "hello")) + ("foo bar" + "foo bar")))) + "")
            _ = v12
        case "":
            switch 42 {
            case 32767:
                v13 := []int{(p0 - 32767), (-(p0))}
                _ = v13
            case 498:
                v14 := GenS1{F0: 1}
                _ = v14
            case 602:
                fmt.Println(p1, p0, p2)
            default:
                var v15 string = "foo bar"
                _ = v15
            }
        }
    case 1:
        p1 = 1
    }
    p2 = (p2 + "a")
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}
func genFn1() int16 {
    var v16 float32 = (0.0 / 1e10)
    v16 -= v16
    var v17 int8 = 42
    v17 += v17
    switch "hello" {
    case "":
        v18 := GenS3{F0: (-(v17))}
        _ = v18
    }
    _ = v16
    _ = v17
    return -1
}

func main() {
    var v1 uint = 339
    v2 := []int{-1, 2147483647, 127}
    var v3 bool = true
    var v4 uint16 = 0
    var v5 uint = (^(v1))
    var v6 int8 = -1
    for fi19 := 0; fi19 < 15; fi19++ {
        v20 := GenS1{F0: 2}
        if v3 {
            fmt.Println(fi19)
            switch (fi19 & (fi19 & -116)) {
            case 32767:
                fmt.Println(fi19, v1)
            }
        }
        switch fi19 {
        case 0:
            for fi21 := 0; fi21 < 11; fi21++ {
                var v22 int8 = v6
                _ = fi21
                _ = v22
            }
        case 252:
            var v23 uint64 = 42
            _ = v23
        }
        _ = fi19
        _ = v20
    }
    v5 = (^(v5))
    _ = len(v2)
    _ = len(v2)
    _ = len(v2)
    var v24 uint64 = 2
    v4 += uint16(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 0 {
    case -1:
        var v25 float64 = -2.25
        _ = v25
    case -491:
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
