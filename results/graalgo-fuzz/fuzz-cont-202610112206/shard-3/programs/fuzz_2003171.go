package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 string; F2 int }
type GenS3 struct { F0 uint }
func genFn0(p0 int64) int16 {
    v6 := []int{594, 712, 9223372036854775807, -1}
    var v7 float32 = 3.14159
    for fi8 := 0; fi8 < 11; fi8++ {
        if true {
            var v9 int64 = p0
            switch -2147483648 {
            case 4294967295:
                v7 += 0.1
            default:
                v10 := GenS1{F0: 966}
                _ = v10
            }
            _ = v9
        }
        var v11 int = fi8
        _ = fi8
        _ = v11
    }
    var v12 int16 = 667
    v13 := GenS1{F0: 570}
    _ = p0
    _ = v12
    _ = v13
    _ = v6
    _ = v7
    return v12
}
func genFn1(p0 int8, p1 int64) uint {
    v14 := []int{501}
    p0 += (p0 | p0)
    switch 32767 {
    case 4294967295:
        v15 := GenS2{F0: 2, F1: (("foo bar" + "foo bar") + "\n\t"), F2: 193}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v14
    return 0
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]uint32)
    v3 := []int{128, 324}
    var v4 int64 = 783
    var v5 uint32 = 65535
    v4 -= (^(v4))
    if ((false || false) || ((42 == uint8(v5)) && true)) {
        delete(v2, "a")
        if (!true) {
            if false {
                var v16 int16 = -1
                _ = v16
            }
            v5 -= 2147483647
        } else {
            var v17 string = ("" + "hello")
            var v18 int8 = 49
            _ = v17
            _ = v18
        }
    }
    v1 -= v1
    var v19 uint8 = 255
    var v20 int32 = (v1 / ((v1) | 1))
    fmt.Println(v5, v20)
    v5 -= v5
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v4 = v4
        var v21 uint = 42
        _ = v21
    } else {
        if false {
            switch -11 {
            case -9223372036854775808:
                var v22 int8 = 127
                _ = v22
            default:
                if true {
                    if (!false) {
                        var v23 uint8 = 0
                        switch ("hello" + "\n\t") {
                        case "\n\t":
                            var v24 float32 = -2.25
                            _ = v24
                        }
                        _ = v23
                    } else {
                        v5 = (v5 / ((v5) | 1))
                    }
                    var v25 float64 = -2.25
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 8; fi26++ {
                        var v27 uint16 = 1
                        var v28 uint16 = uint16(v19)
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    var v29 uint32 = v5
                    if len(v3) > 0 {
                        v3[0] = 32767
                    }
                    _ = v29
                }
            }
            fmt.Println(v4, v4)
            var v30 uint32 = (v5 + v5)
            _ = v30
        }
        var v31 int64 = (v4 & (v4 ^ (v4 | v4)))
        for fi32 := 0; fi32 < 3; fi32++ {
            for fi33 := 0; fi33 < 5; fi33++ {
                var v34 int = (^(fi32))
                v35 := GenS3{F0: 129}
                _ = fi33
                _ = v34
                _ = v35
            }
            _ = fi32
        }
        _ = v31
    }
    var v36 uint = 255
    if false {
        if (2147483647 != int32(v4)) {
            v20 += v1
        }
    } else {
        var v37 int = 2147483647
        fmt.Println(v37)
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
