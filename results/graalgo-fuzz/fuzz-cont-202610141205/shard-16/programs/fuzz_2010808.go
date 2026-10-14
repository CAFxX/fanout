package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 uint16; F1 int8; F2 bool }
func genFn0(p0 uint) bool {
    v7 := GenS1{F0: -1, F1: p0}
    var v8 int32 = 42
    for fi9 := 0; fi9 < 7; fi9++ {
        if (((0.0 / (0.1 / ((((1e10 * -2.25) / -2.25) + 0.1) - (0.1 * 0.1)))) != (3.14159 / -2.25)) && (uint8(p0) < 127)) {
            switch ("hello" + "\n\t") {
            case "a":
                if false {
                    v10 := []int{(fi9 + 2147483647)}
                    fmt.Println(p0)
                    _ = v10
                }
            }
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := []int{}
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 10; fi13++ {
                v14 := GenS1{F0: v8, F1: p0}
                v15 := []int{648, (fi9 * (fi13 & (fi13 / ((128) | 1)))), (fi13 + (fi9 * fi9)), fi9}
                _ = fi13
                _ = v14
                _ = v15
            }
        } else {
            switch "foo bar" {
            case "":
                v16 := []int{(fi9 % (((fi9 * -86)) | 1))}
                _ = v16
            case "a":
                fmt.Println(fi9)
            default:
                if false {
                    var v17 int32 = (v8 ^ v8)
                    p0 -= 65535
                    _ = v17
                }
            }
        }
        _ = fi9
    }
    p0 -= 255
    _ = p0
    _ = v7
    _ = v8
    return ((2 <= int8(v8)) && (86 >= 128))
}
func genFn1(p0 bool, p1 string, p2 uint8) bool {
    var v18 string = p1
    switch v18 {
    case "a":
        for fi19 := 0; fi19 < 20; fi19++ {
            v20 := []int{-2147483648, fi19, (fi19 >> ((uint(p2)) & 63)), fi19}
            p0 = p0
            p1 = (p1 + p1)
            _ = fi19
            _ = v20
        }
    default:
        var v21 int64 = -894
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return true
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]string)
    var v3 uint8 = 127
    var v4 float64 = 3.14159
    v5 := make(map[string]bool)
    var v6 uint32 = v1
    var v22 float64 = float64(3.14159)
    if false {
        v23 := GenS2{F0: 42, F1: 40}
        v5[("\n\t" + "a")] = (1 > 2)
        v24 := GenS3{F0: 45}
        _ = v23
        _ = v24
    }
    var v25 float32 = float32(0.1)
    v26 := []int{127}
    for fi27 := 0; fi27 < 16; fi27++ {
        v28 := GenS4{F0: v3}
        if (1 > 4294967295) {
            var v29 uint = 756
            if (4294967295 != 4294967295) {
                var v30 int64 = 1
                _ = len(v26)
                switch int(v29) {
                case 251:
                    v31 := GenS5{F0: -1, F1: 42}
                    _ = v31
                default:
                    var v32 int = fi27
                    _ = v32
                }
                _ = v30
            } else {
                var v33 string = ((("hello" + "a") + ((("" + "\n\t") + "\n\t") + "a")) + "")
                _ = v33
            }
            _ = v29
        } else {
            for fi34 := 0; fi34 < 2; fi34++ {
                var v35 int16 = -454
                _ = fi34
                _ = v35
            }
            v36 := v2[(v6 ^ 61)]
            _ = len(v26)
            _ = v36
        }
        _ = fi27
        _ = v28
    }
    v37 := GenS4{F0: (v3 << ((65535) & 63))}
    var v38 uint64 = 198
    switch 464 {
    case -9223372036854775808:
        for fi39 := 0; fi39 < 18; fi39++ {
            var v40 int = -1
            var v41 string = "hello"
            _ = fi39
            _ = v40
            _ = v41
        }
    case 112:
        if len(v26) > 0 {
            v26[0] = 32768
        }
    }
    var v42 uint = 171
    v3 -= (v3 / (((v3 | 30)) | 1))
    if true {
        fmt.Println(v25, v6, v22)
        delete(v5, "foo bar")
    }
    var v43 float64 = 3.14159
    var v44 int = 672
    var v45 uint16 = 1
    v46 := GenS6{F0: 1, F1: -66, F2: false}
    v4 = v4
    for fi47 := 0; fi47 < 16; fi47++ {
        var v48 uint = v42
        _ = fi47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
