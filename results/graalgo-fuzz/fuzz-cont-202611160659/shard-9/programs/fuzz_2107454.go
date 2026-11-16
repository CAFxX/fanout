package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int8; F2 uint8 }
type GenS3 struct { F0 int }
func genFn0() int16 {
    if (9223372036854775807 <= 0) {
        if (-631 < -2147483648) {
            var v7 int = 265
            v7 = v7
            _ = v7
        }
    }
    v8 := GenS1{F0: 103}
    var v9 uint32 = 2
    for fi10 := 0; fi10 < 13; fi10++ {
        v9 = (v9 + (v9 << ((42) & 63)))
        _ = fi10
    }
    _ = v8
    _ = v9
    return 73
}

func main() {
    var v1 string = ("hello" + (("a" + "foo bar") + "hello"))
    var v2 uint16 = 706
    var v3 string = ""
    var v4 float32 = 1.5
    v5 := make(map[uint32]uint32)
    v6 := []string{("a" + "a")}
    v11 := GenS1{F0: 745}
    v12 := GenS2{F0: -1, F1: -128, F2: 42}
    for fi13 := 0; fi13 < 5; fi13++ {
        if true {
            var v14 uint32 = 783
            var v15 int32 = -935
            var v16 string = (v3 + v3)
            _ = v14
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := GenS1{F0: 65535}
                v4 += (3.14159 - 3.14159)
                _ = fi17
                _ = v18
            }
        }
        v3 = v3
        _ = fi13
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v19 uint = 2
    var v20 float64 = (((1e10 * (-2.25 + 1e10)) / -2.25) / ((1.5 + (0.0 - (0.1 * ((0.0 / 3.14159) * 3.14159)))) / (((1.5 * (((1e10 + (3.14159 * (1e10 + ((1e10 - 0.0) + 0.0)))) - 3.14159) * 0.1)) * 0.0) - 3.14159)))
    switch ((((v1 + ((v3 + "\n\t") + "")) + v3) + (v1 + "a")) + v3) {
    case "a":
        v4 += v4
    case "":
        _ = len(v5)
    default:
        var v21 uint16 = (v2 + 773)
        _ = v21
    }
    var v22 uint = v19
    v20 += v20
    switch -2147483648 {
    case -9223372036854775808:
        var v23 int32 = int32(v19)
        _ = v23
    default:
        var v24 int16 = 1000
        _ = v24
    }
    if (!false) {
        v25 := GenS3{F0: 32767}
        var v26 uint16 = v2
        var v27 int64 = -9223372036854775808
        _ = v25
        _ = v26
        _ = v27
    } else {
        var v28 uint32 = 65535
        _ = v28
    }
    v29 := GenS3{F0: 4294967295}
    v30 := GenS1{F0: 2147483648}
    if false {
        v31 := v5[2147483647]
        _ = v31
    } else {
        for fi32 := 0; fi32 < 10; fi32++ {
            v33 := GenS1{F0: uint32(v19)}
            _ = fi32
            _ = v33
        }
        v2 = v2
        v1 = v3
    }
    v20 += 0.0
    fmt.Println(v2, v2)
    for fi34 := 0; fi34 < 16; fi34++ {
        switch (^(fi34)) {
        case 32768:
            fmt.Println(fi34)
        default:
            var v35 uint16 = 255
            _ = v35
        }
        var v36 int = fi34
        var v37 uint16 = v2
        _ = fi34
        _ = v36
        _ = v37
    }
    var v38 int = 392
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
