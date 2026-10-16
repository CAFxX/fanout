package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint16; F2 int }
func genFn0(p0 int32) uint16 {
    fmt.Println(p0)
    var v5 bool = (2 < 255)
    if v5 {
        var v6 string = "hello"
        switch 42 {
        case 127:
            if false {
                v7 := []int{-220, -711, 32767, 498}
                var v8 float32 = 0.1
                _ = v7
                _ = v8
            }
        default:
            var v9 bool = v5
            _ = v9
        }
        var v10 uint32 = 42
        _ = v10
        _ = v6
    }
    _ = p0
    _ = v5
    return 255
}
func genFn1(p0 int8) uint8 {
    v11 := []int{127}
    v12 := GenS1{F0: ("" + "\n\t")}
    v11 = append(v11, 42)
    fmt.Println(p0)
    v13 := GenS2{F0: 950, F1: 255}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 string = ""
    v2 := make(map[int]bool)
    v3 := []uint32{959}
    var v4 uint8 = 91
    var v14 int8 = 90
    if false {
        if (841 > 755) {
            v15 := GenS1{F0: ""}
            switch -414 {
            case 4294967295:
                fmt.Println(v1, v4)
            case 817:
                switch "foo bar" {
                case "foo bar":
                    var v16 int64 = -1
                    _ = v16
                case "\n\t":
                    v17 := GenS3{F0: (v14 << ((0) & 63)), F1: genFn0(1), F2: 9223372036854775807}
                    _ = v17
                }
            default:
                var v18 int8 = v14
                _ = v18
            }
            _ = v15
        } else {
            _ = len(v2)
        }
        var v19 int = 128
        _ = v19
    }
    v14 += v14
    switch 32768 {
    case 385:
        if (!(!(42 != 2147483648))) {
            if true {
                for fi20 := 0; fi20 < 11; fi20++ {
                    fmt.Println(v1)
                    v21 := GenS2{F0: 491, F1: (^(v4))}
                    v22 := GenS2{F0: 42, F1: (^(v4))}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            } else {
                v1 = ("" + ("" + v1))
            }
            v1 = (v1 + v1)
        } else {
            v14 = (^(v14))
            for fi23 := 0; fi23 < 4; fi23++ {
                v4 -= 127
                switch (v1 + (v1 + v1)) {
                case "a":
                    var v24 float32 = float32(3.14159)
                    _ = v24
                case "foo bar":
                    for fi25 := 0; fi25 < 5; fi25++ {
                        v3 = append(v3, 259)
                        v26 := GenS1{F0: "a"}
                        _ = fi25
                        _ = v26
                    }
                default:
                    var v27 float64 = 3.14159
                    _ = v27
                }
                v2[fi23] = (!true)
                _ = fi23
            }
        }
    case 1:
        var v28 int32 = 42
        _ = v28
    case 128:
        if true {
            switch "" {
            case "":
                v4 = genFn1(-128)
            case "foo bar":
                fmt.Println(v4)
            case "\n\t":
                fmt.Println(v1, v4)
            }
            v14 = v14
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            var v29 uint = 561
            _ = v29
        }
    }
    switch 1 {
    case 128:
        if (42 < 255) {
            v2[-416] = (758 >= 2147483648)
            v30 := GenS2{F0: 156, F1: (v4 << ((255) & 63))}
            _ = v30
        }
    case 891:
        var v31 float32 = float32(1.5)
        _ = v31
    case 9223372036854775807:
        var v32 int16 = -32768
        _ = v32
    default:
        _ = len(v3)
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        for fi34 := 0; fi34 < 1; fi34++ {
            var v35 int8 = (-(v14))
            _ = fi34
            _ = v35
        }
        v1 = v1
        _ = fi33
    }
    var v36 int64 = 42
    v1 = (v1 + (("" + "foo bar") + ((v1 + v1) + v1)))
    v36 += v36
    for fi37 := 0; fi37 < 9; fi37++ {
        v36 += v36
        _ = fi37
    }
    v14 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v4
}
