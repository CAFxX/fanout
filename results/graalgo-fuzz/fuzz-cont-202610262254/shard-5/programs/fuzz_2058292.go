package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint16, p1 int64, p2 uint32) int {
    p1 -= (p1 + p1)
    switch 32768 {
    case 32767:
        fmt.Println(p2, p2, p0)
    case 121:
        if false {
            for fi5 := 0; fi5 < 11; fi5++ {
                var v6 int8 = -128
                _ = fi5
                _ = v6
            }
        } else {
            v7 := []int{852, 604, 42, 42}
            v8 := GenS1{F0: (true || false), F1: -1}
            v9 := GenS2{F0: 197}
            _ = v7
            _ = v8
            _ = v9
        }
    }
    var v10 int8 = -11
    fmt.Println(p0)
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return -133
}
func genFn1(p0 uint, p1 int8) int16 {
    var v12 uint = (p0 * (p0 / ((65535) | 1)))
    if true {
        var v13 int64 = 0
        _ = v13
    }
    fmt.Println(v12, p1)
    fmt.Println(p0, p1, v12)
    switch "a" {
    case "a":
        v14 := []int{2147483647, -245, 9223372036854775807}
        _ = v14
    case "":
        v15 := GenS1{F0: (216 <= 9223372036854775807), F1: p1}
        _ = v15
    case "\n\t":
        v16 := GenS2{F0: p0}
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v12
    return -32768
}

func main() {
    var v1 uint32 = 586
    var v2 int16 = -1
    var v3 int16 = (-(v2))
    var v4 uint16 = 1
    v1 = 212
    v17 := GenS1{F0: (false || false), F1: -1}
    var v18 string = "a"
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            if ((v2 ^ v3) != (^(v2))) {
                var v21 uint8 = 25
                _ = v21
            }
            var v22 uint = 0
            _ = fi20
            _ = v22
        }
        _ = fi19
    }
    if (1e10 > 1e10) {
        var v23 uint = uint(v4)
        _ = v23
    } else {
        v4 += v4
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 uint = 255
            if (!true) {
                fmt.Println(v18)
            } else {
                v26 := []int{(fi24 - 128)}
                _ = v26
            }
            _ = fi24
            _ = v25
        }
    }
    switch v18 {
    case "hello":
        fmt.Println(v1, v18, v4)
    case "":
        v27 := GenS2{F0: 366}
        _ = v27
    }
    var v28 int = 9223372036854775807
    switch genFn0(1, -1, (^(v1))) {
    case -246:
        if false {
            fmt.Println(v2, v1)
        }
    case 0:
        fmt.Println(v28, v18, v28)
    default:
        var v29 float32 = float32((float64(3.14159) + 1.5))
        _ = v29
    }
    var v30 uint32 = 65535
    v30 = (v1 * (v1 - (v30 | (v30 ^ (^(v30))))))
    if ((!false) && (!true)) {
        switch v18 {
        case "hello":
            if (2147483648 <= (^(v30))) {
                var v31 int = v28
                _ = v31
            }
        case "foo bar":
            var v32 float32 = -2.25
            _ = v32
        case "\n\t":
            v33 := GenS3{F0: 65535, F1: 9223372036854775807, F2: 56}
            _ = v33
        default:
            for fi34 := 0; fi34 < 6; fi34++ {
                if ((!(303 != -254)) || false) {
                    v35 := GenS4{F0: (v1 * (^(v30))), F1: true}
                    var v36 bool = true
                    var v37 float32 = -2.25
                    _ = v35
                    _ = v36
                    _ = v37
                } else {
                    switch (v18 + "") {
                    case "foo bar":
                        var v38 uint64 = 4294967295
                        _ = v38
                    case "a":
                        var v39 int = (fi34 - (v28 * 837))
                        _ = v39
                    case "\n\t":
                        v40 := GenS5{F0: 0, F1: (^(v4))}
                        _ = v40
                    default:
                        if (v18 == "") {
                            var v41 int32 = 1
                            v30 = (v30 - (^(v1)))
                            _ = v41
                        } else {
                            var v42 string = v18
                            _ = v42
                        }
                    }
                }
                var v43 float64 = 1e10
                _ = fi34
                _ = v43
            }
        }
    }
    var v44 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v44
}
