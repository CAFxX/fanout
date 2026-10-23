package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
func genFn0(p0 int64) int32 {
    switch -1 {
    case 1:
        v5 := []int{}
        _ = v5
    case 254:
        var v6 uint = 360
        _ = v6
    case 42:
        p0 += p0
    default:
        var v7 uint64 = 42
        _ = v7
    }
    switch ("foo bar" + ("hello" + "")) {
    case "":
        for fi8 := 0; fi8 < 13; fi8++ {
            if false {
                if ((1e10 / 1e10) < 3.14159) {
                    p0 += (p0 << ((42) & 63))
                    var v9 string = "foo bar"
                    p0 = (p0 % ((8) | 1))
                    _ = v9
                } else {
                    var v10 float64 = 0.1
                    _ = v10
                }
            }
            var v11 float32 = 1.5
            _ = fi8
            _ = v11
        }
    case "a":
        v12 := GenS1{F0: true, F1: -36}
        _ = v12
    case "hello":
        if true {
            v13 := GenS1{F0: true, F1: 46}
            p0 = (^(p0))
            var v14 int16 = 547
            _ = v13
            _ = v14
        } else {
            v15 := []int{-1, 2147483647, 42, 128}
            if false {
                var v16 uint8 = 2
                _ = v16
            } else {
                for fi17 := 0; fi17 < 10; fi17++ {
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                    fmt.Println(fi17)
                    fmt.Println(fi17, p0)
                    _ = fi17
                }
            }
            var v18 string = "hello"
            _ = v15
            _ = v18
        }
    }
    v19 := []int{}
    _ = p0
    _ = v19
    return -2147483648
}
func genFn1(p0 bool, p1 int32) uint8 {
    var v20 uint16 = 255
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 uint8 = 228
        var v23 uint = 347
        _ = fi21
        _ = v22
        _ = v23
    }
    if (127 == 1) {
        v24 := []int{32768}
        fmt.Println(p0, v20, p0)
        var v25 uint32 = 2147483648
        _ = v24
        _ = v25
    } else {
        p0 = (999 > -9223372036854775808)
    }
    _ = p0
    _ = p1
    _ = v20
    return 128
}

func main() {
    var v1 int16 = 42
    var v2 int8 = -128
    var v3 int32 = -2147483648
    var v4 int8 = 127
    var v26 int64 = -1
    v27 := []int{330, -2147483648, 42}
    var v28 float64 = -2.25
    fmt.Println(v2, v26, v28)
    v27 = append(v27, 42)
    var v29 int8 = v4
    v3 += 1
    var v30 uint16 = 2
    var v31 int8 = 0
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
