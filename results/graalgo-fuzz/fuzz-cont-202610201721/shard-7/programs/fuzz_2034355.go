package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint32; F2 uint16 }
func genFn0(p0 uint, p1 uint8, p2 uint64) int16 {
    v4 := GenS1{F0: 81, F1: 100, F2: p1}
    v5 := []int{0, 1, 2147483647, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}
func genFn1() bool {
    if (-196 == -1) {
        if ((65535 >= 4294967295) && false) {
            var v6 float32 = 1e10
            _ = v6
        } else {
            v7 := []int{2147483647, -9223372036854775808, 42, -312}
            var v8 uint32 = 65535
            _ = v7
            _ = v8
        }
    } else {
        var v9 int8 = 34
        _ = v9
    }
    if ((2147483647 != 165) && (2147483648 <= 123)) {
        var v10 uint = 2
        if (!(!false)) {
            v11 := GenS2{F0: 1, F1: 1}
            var v12 float64 = (float64(((3.14159 / -2.25) + 0.1)) - ((3.14159 / -2.25) * (0.0 - -2.25)))
            var v13 int64 = -933
            _ = v11
            _ = v12
            _ = v13
        } else {
            switch "a" {
            case "a":
                fmt.Println(v10)
            case "hello":
                for fi14 := 0; fi14 < 4; fi14++ {
                    fmt.Println(fi14, fi14)
                    v15 := []int{(fi14 * fi14), 103, -825}
                    _ = fi14
                    _ = v15
                }
            default:
                if (1.5 > -2.25) {
                    if false {
                        if (false || (2147483648 < 1)) {
                            var v16 uint = v10
                            var v17 string = "a"
                            var v18 float32 = (-2.25 / float32(1e10))
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        var v19 int32 = -2147483648
                        _ = v19
                    } else {
                        v10 += v10
                        fmt.Println(v10)
                    }
                    var v20 uint32 = 501
                    _ = v20
                } else {
                    v10 = (v10 & 243)
                    v21 := []int{4294967295, 42, 127, 32768}
                    fmt.Println(v10)
                    _ = v21
                }
            }
            v22 := GenS3{F0: 32767}
            _ = v22
        }
        var v23 uint16 = 70
        _ = v10
        _ = v23
    }
    return true
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]int)
    var v3 int16 = 1
    if false {
        v24 := GenS3{F0: -99}
        _ = v24
    } else {
        if false {
            v25 := GenS2{F0: uint64(v1), F1: 2}
            fmt.Println(v1, v1)
            v26 := GenS4{F0: false, F1: 131, F2: 286}
            _ = v25
            _ = v26
        } else {
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 uint8 = v1
                _ = fi27
                _ = v28
            }
            v29 := []int{}
            fmt.Println(v1, v1)
            _ = v29
        }
    }
    fmt.Println(v1, v1)
    fmt.Println(v3, v1)
    fmt.Println(v1, v3)
    var v30 float32 = 1.5
    if false {
        var v31 int8 = int8(v3)
        _ = v31
    } else {
        v32 := []int{32768, 0, 4294967295}
        _ = v32
    }
    v33 := GenS3{F0: 32768}
    v1 -= (v1 >> ((801) & 63))
    if false {
        var v34 int = 32768
        var v35 float64 = (float64(v30) / 1.0)
        _ = v34
        _ = v35
    }
    v30 = ((v30 + (v30 / v30)) / (v30 / (3.14159 / 0.1)))
    var v36 float32 = 1.5
    v36 += v30
    v36 = (3.14159 - v36)
    var v37 bool = (2 > 65535)
    switch "\n\t" {
    case "foo bar":
        if true {
            v38 := GenS1{F0: 798, F1: -1, F2: v1}
            var v39 uint = 42
            _ = v38
            _ = v39
        }
    case "\n\t":
        v40 := GenS1{F0: uint(v1), F1: -1, F2: v1}
        _ = v40
    case "hello":
        for fi41 := 0; fi41 < 8; fi41++ {
            var v42 uint16 = 148
            v36 = ((v30 / (1.5 + 1e10)) / (v36 - (0.0 + v36)))
            var v43 int = fi41
            _ = fi41
            _ = v42
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v36
    _ = v37
}
