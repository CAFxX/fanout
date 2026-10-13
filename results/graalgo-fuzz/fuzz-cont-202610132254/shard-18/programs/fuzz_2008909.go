package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int64) uint32 {
    fmt.Println(p1, p0)
    if (100 > 127) {
        p0 = p0
        p0 += 255
        switch "\n\t" {
        case "foo bar":
            if true {
                v5 := []int{1, -1, 42}
                _ = v5
            }
        }
    }
    if ((1.5 <= 0.1) || true) {
        switch ("hello" + ("foo bar" + ("foo bar" + "foo bar"))) {
        case "a":
            var v6 uint16 = (p0 << ((2) & 63))
            _ = v6
        case "":
            if (true || false) {
                var v7 int16 = 1000
                _ = v7
            }
        }
        if (((false && false) || (0 != 42)) || false) {
            p0 = 65535
        } else {
            p0 = (^(p0))
            v8 := []int{-547, 159, 127, -2147483648}
            _ = v8
        }
    } else {
        var v9 int64 = -1
        _ = v9
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 uint64, p1 int) int32 {
    v10 := []int{(p1 << ((42) & 63)), 2147483647, (p1 | (p1 / (((p1 * 2147483647)) | 1))), p1}
    _ = len(v10)
    v11 := GenS1{F0: 0, F1: 742, F2: 255}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 2147483647
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]string)
    var v3 int = 127
    v4 := make(map[string]bool)
    v3 += (v3 ^ v3)
    v12 := []int{(v3 / ((113) | 1)), v3, v3, 4294967295}
    delete(v4, ("\n\t" + v1))
    var v13 int = -9223372036854775808
    var v14 uint64 = 2
    switch v1 {
    case "a":
        for fi15 := 0; fi15 < 10; fi15++ {
            _ = len(v4)
            _ = fi15
        }
    case "hello":
        fmt.Println(v3, v14, v1)
    case "":
        for fi16 := 0; fi16 < 11; fi16++ {
            if (!false) {
                fmt.Println(v13, v14, v14)
                v3 += fi16
                _ = len(v12)
            }
            if (1 >= -9223372036854775808) {
                fmt.Println(v1, fi16, v14)
            } else {
                switch -2147483648 {
                case -311:
                    for fi17 := 0; fi17 < 17; fi17++ {
                        v18 := GenS2{F0: 42, F1: 112}
                        var v19 float64 = -2.25
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                default:
                    _ = len(v4)
                }
                var v20 bool = true
                v14 = v14
                _ = v20
            }
            _ = fi16
        }
    default:
        fmt.Println(v13)
    }
    v14 = 4294967295
    v12 = append(v12, (v3 % (((v3 * (v3 / ((v3) | 1)))) | 1)))
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 uint8 = 0
        switch (v3 | (fi21 / (((fi21 ^ v13)) | 1))) {
        case -149:
            v1 = v1
        case 42:
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := GenS3{F0: (^(v22))}
                _ = fi23
                _ = v24
            }
        }
        switch v1 {
        case "hello":
            for fi25 := 0; fi25 < 2; fi25++ {
                v3 = (^(v13))
                _ = fi25
            }
        }
        _ = fi21
        _ = v22
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        v14 = (v14 % (((v14 << ((42) & 63))) | 1))
        v14 += (v14 ^ (v14 + v14))
        var v27 int64 = 1
        _ = fi26
        _ = v27
    }
    switch 2147483647 {
    case -750:
        v12 = append(v12, v13)
    case 2147483647:
        var v28 float32 = float32(3.14159)
        _ = v28
    default:
        var v29 bool = (true && false)
        _ = v29
    }
    var v30 uint16 = 2
    switch ("a" + (v1 + (v1 + "hello"))) {
    case "\n\t":
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    var v31 float64 = 0.1
    var v32 bool = false
    if (0 == genFn1(v14, (v13 - 128))) {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        var v33 uint32 = 2147483647
        _ = v33
    } else {
        v32 = (true || v32)
        for fi34 := 0; fi34 < 15; fi34++ {
            for fi35 := 0; fi35 < 6; fi35++ {
                fmt.Println(v3)
                _ = fi35
            }
            _ = fi34
        }
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        var v37 uint64 = v14
        _ = fi36
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
