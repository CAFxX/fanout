package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int }
func genFn0(p0 uint64) int64 {
    v7 := []int{703, 947, 0, 32768}
    if (128 != 2) {
        if len(v7) > 0 {
            v7[0] = 32767
        }
    } else {
        var v8 int64 = 1
        v8 -= (v8 << ((2) & 63))
        _ = v8
    }
    p0 += (p0 % ((p0) | 1))
    _ = p0
    _ = v7
    return 389
}
func genFn1(p0 int, p1 uint32) bool {
    var v9 string = ("foo bar" + "\n\t")
    if (1 < 42) {
        var v10 int64 = -758
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 float32 = 1e10
            var v13 uint64 = 1
            _ = fi11
            _ = v12
            _ = v13
        }
        var v14 int16 = 42
        _ = v10
        _ = v14
    }
    var v15 float32 = ((1e10 + (3.14159 / 0.1)) + (0.0 - -2.25))
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 string = v9
        switch (p0 + (fi16 * (-(p0)))) {
        case -1:
            v9 = v17
        }
        _ = fi16
        _ = v17
    }
    v9 = v9
    _ = p0
    _ = p1
    _ = v15
    _ = v9
    return true
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]int)
    var v3 bool = false
    v4 := make(map[uint32]uint32)
    var v5 int64 = -581
    var v6 float32 = float32((float64(1e10) - (1.5 / (3.14159 * 1.5))))
    var v18 int8 = 42
    for fi19 := 0; fi19 < 3; fi19++ {
        v18 = (v18 << (((v1 + v1)) & 63))
        _ = fi19
    }
    v3 = ((v3 || v3) || (true || v3))
    switch 128 {
    case 32767:
        v20 := v4[0]
        _ = v20
    case 1:
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 string = "hello"
            v23 := GenS1{F0: 2, F1: 42, F2: fi21}
            _ = fi21
            _ = v22
            _ = v23
        }
    case 377:
        switch "\n\t" {
        case "foo bar":
            v1 -= v1
        }
    default:
        v3 = v3
    }
    var v24 uint = (v1 / (((v1 >> (((^(v1))) & 63))) | 1))
    if ((("hello" + "\n\t") + "a") > "hello") {
        fmt.Println(v6, v24, v6)
        if v3 {
            for fi25 := 0; fi25 < 10; fi25++ {
                if (int8(fi25) > 127) {
                    v4[2147483647] = 698
                }
                _ = fi25
            }
        }
        if (true || (v24 > (^(v24)))) {
            if v3 {
                var v26 bool = v3
                if (257 != 65535) {
                    v27 := GenS1{F0: 128, F1: uint8(v1), F2: 128}
                    var v28 uint = v24
                    var v29 uint = v1
                    _ = v27
                    _ = v28
                    _ = v29
                }
                fmt.Println(v1, v6)
                _ = v26
            }
        } else {
            v2["hello"] = -2147483648
        }
    } else {
        fmt.Println(v1)
    }
    var v30 uint = v24
    fmt.Println(v5, v18, v24)
    v30 -= (v1 << ((20) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
