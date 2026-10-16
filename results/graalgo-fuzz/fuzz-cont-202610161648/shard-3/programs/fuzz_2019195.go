package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint }
type GenS3 struct { F0 int; F1 int8; F2 uint64 }
type GenS4 struct { F0 string }
func genFn0(p0 uint8, p1 int16, p2 string) uint {
    switch -27 {
    case 32768:
        var v6 float32 = 3.14159
        _ = v6
    case -725:
        fmt.Println(p0)
    case 42:
        p1 += (^(p1))
    }
    v7 := []int{390}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 83
}
func genFn1(p0 int8) int8 {
    v8 := []int{863}
    var v9 int64 = 0
    v10 := GenS1{F0: "foo bar", F1: 18446744073709551615}
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return -114
}

func main() {
    v1 := []int{128, 32767}
    var v2 uint = 65535
    var v3 string = "hello"
    var v4 int32 = 42
    var v5 float32 = float32(float64(float32(-2.25)))
    v4 += v4
    var v11 int32 = (^(v4))
    v3 = v3
    v12 := GenS2{F0: 647, F1: 4294967295, F2: (v2 >> ((genFn0(62, -1, (("a" + v3) + "foo bar"))) & 63))}
    _ = len(v1)
    for fi13 := 0; fi13 < 7; fi13++ {
        if true {
            fmt.Println(v11, fi13, v11)
            if ((false || (false && true)) || (!(v5 != 0.1))) {
                for fi14 := 0; fi14 < 20; fi14++ {
                    switch int(v2) {
                    case 128:
                        switch v3 {
                        case "hello":
                            var v15 int16 = -1
                            _ = v15
                        default:
                            v3 = "a"
                        }
                    case 4294967295:
                        if false {
                            fmt.Println(fi13, fi14, v5)
                            if false {
                                v2 += v2
                                fmt.Println(v11, v5)
                            } else {
                                v16 := GenS3{F0: fi14, F1: 42, F2: 549}
                                var v17 float32 = (v5 - (v5 / 1e10))
                                _ = v16
                                _ = v17
                            }
                            var v18 float32 = (v5 + float32(0.0))
                            _ = v18
                        }
                    }
                    var v19 string = ""
                    fmt.Println(v4)
                    _ = fi14
                    _ = v19
                }
                var v20 string = "a"
                _ = v20
            }
        }
        var v21 float64 = 3.14159
        _ = fi13
        _ = v21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        v5 -= 1e10
        _ = fi22
    }
    fmt.Println(v5, v5, v4)
    var v23 uint8 = 233
    switch ("a" + v3) {
    case "foo bar":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "hello":
        var v24 int8 = 0
        _ = v24
    }
    v25 := GenS4{F0: v3}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
