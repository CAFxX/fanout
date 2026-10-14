package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 string }
func genFn0(p0 int32, p1 uint32) bool {
    v7 := []int{9223372036854775807, 32768, 814, 1}
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: (true || false)}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 127}
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS2{F0: (fi11 % ((1) | 1))}
        p1 = p1
        _ = len(v7)
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return (!true)
}
func genFn1() uint8 {
    switch "a" {
    case "a":
        var v13 int8 = -128
        _ = v13
    case "\n\t":
        switch (("foo bar" + "hello") + "a") {
        case "\n\t":
            var v14 int64 = 42
            _ = v14
        default:
            var v15 int32 = 42
            _ = v15
        }
    default:
        if (float64(1e10) < ((0.0 + 1e10) * (-2.25 / 1.0))) {
            switch -2147483648 {
            case 213:
                var v16 uint64 = 42
                _ = v16
            case 127:
                switch ("hello" + "hello") {
                case "":
                    v17 := []int{127, 154}
                    _ = v17
                default:
                    switch "\n\t" {
                    case "foo bar":
                        var v18 uint64 = 175
                        _ = v18
                    case "":
                        var v19 int32 = 1
                        _ = v19
                    }
                }
            case 79:
                var v20 uint = 42
                _ = v20
            }
        } else {
            var v21 uint64 = 42
            v22 := []int{-304, -2147483648, 2147483647, -2147483648}
            switch 328 {
            case 1:
                _ = len(v22)
            }
            _ = v21
            _ = v22
        }
    }
    if false {
        var v23 uint8 = 162
        _ = v23
    }
    v24 := GenS2{F0: 32767}
    _ = v24
    return 66
}

func main() {
    var v1 int32 = 42
    var v2 int64 = 503
    v3 := []int{}
    var v4 uint8 = 1
    var v5 uint = 42
    var v6 string = "a"
    for fi25 := 0; fi25 < 4; fi25++ {
        switch (fi25 & -2147483648) {
        case 1:
            var v26 bool = (v2 > v2)
            _ = v26
        }
        v1 -= (v1 & (-(v1)))
        _ = fi25
    }
    v27 := GenS1{F0: true}
    var v28 string = "foo bar"
    if true {
        switch 42 {
        case 32767:
            var v29 float64 = (-2.25 / 1.0)
            _ = v29
        case -780:
            _ = len(v3)
        }
    } else {
        var v30 float64 = (1e10 + 0.1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch v6 {
        case "":
            v5 = 202
        }
        _ = v30
    }
    if ((3.14159 > 3.14159) || ((4294967295 > 132) && false)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi31 := 0; fi31 < 7; fi31++ {
            v1 = -2147483648
            var v32 uint64 = 4294967295
            var v33 float32 = 1e10
            _ = fi31
            _ = v32
            _ = v33
        }
    } else {
        v3 = append(v3, 4294967295)
    }
    var v34 int32 = 948
    var v35 uint8 = 127
    if (3.14159 <= (1.5 + 1.5)) {
        var v36 uint = (v5 | 467)
        v3 = append(v3, -1)
        _ = v36
    } else {
        v6 = v6
        v37 := GenS3{F0: 32767, F1: v28}
        v1 += v1
        _ = v37
    }
    var v38 uint16 = 169
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
