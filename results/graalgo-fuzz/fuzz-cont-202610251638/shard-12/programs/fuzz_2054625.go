package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint8 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 0
    v3 := []int{128, -931, 4294967295}
    v4 := []int{32768, -535, -9223372036854775808, -2147483648}
    v5 := make(map[string]string)
    var v6 int32 = 705
    _ = len(v5)
    var v7 int32 = (^(v6))
    v7 -= 42
    for fi8 := 0; fi8 < 2; fi8++ {
        fmt.Println(v7, v7, v2)
        if len(v4) > 0 {
            v4[0] = (^(fi8))
        }
        _ = fi8
    }
    if (65535 == 756) {
        var v9 uint16 = 65535
        _ = v9
    }
    v7 = 0
    var v10 int32 = (v7 - (v6 >> ((255) & 63)))
    v11 := GenS1{F0: 42, F1: -9223372036854775808, F2: 127}
    if false {
        switch 2147483647 {
        case 127:
            fmt.Println(v10, v7)
        case 0:
            v10 -= 0
        case 2147483647:
            _ = len(v3)
        default:
            if true {
                switch 4294967295 {
                case 4294967295:
                    var v12 int8 = -125
                    _ = v12
                default:
                    var v13 int = -1
                    _ = v13
                }
                fmt.Println(v10)
                v7 = v7
            } else {
                v6 = (v10 - (v7 >> ((uint(v2)) & 63)))
                switch 128 {
                case 4294967295:
                    for fi14 := 0; fi14 < 7; fi14++ {
                        v2 = (v2 + (v2 & v2))
                        var v15 float32 = 1e10
                        _ = fi14
                        _ = v15
                    }
                case -2147483648:
                    if false {
                        var v16 int16 = 1
                        v6 -= v6
                        _ = v16
                    }
                case 9223372036854775807:
                    delete(v5, "a")
                default:
                    v2 = (^(v2))
                }
            }
        }
        if (98 <= 1) {
            switch ("a" + "") {
            case "a":
                v2 = 2
            case "":
                var v17 int32 = v7
                _ = v17
            }
            if (!true) {
                var v18 int64 = 42
                var v19 int16 = 1000
                for fi20 := 0; fi20 < 16; fi20++ {
                    v21 := GenS2{F0: fi20, F1: ("hello" + ("\n\t" + "\n\t"))}
                    fmt.Println(v19, v7, v10)
                    v6 = 0
                    _ = fi20
                    _ = v21
                }
                _ = v18
                _ = v19
            } else {
                if (872 <= -34) {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        v10 = v6
                        _ = fi22
                    }
                } else {
                    _ = len(v5)
                    if (-1 == -191) {
                        var v23 uint64 = v2
                        _ = v23
                    }
                }
                v3 = append(v3, 42)
            }
        } else {
            var v24 int64 = int64(v6)
            _ = v24
        }
    } else {
        var v25 int32 = v7
        if true {
            var v26 uint = 352
            for fi27 := 0; fi27 < 18; fi27++ {
                var v28 uint16 = 2
                var v29 uint16 = v28
                _ = fi27
                _ = v28
                _ = v29
            }
            v25 -= (v25 * 1)
            _ = v26
        }
        _ = v25
    }
    v30 := GenS3{F0: ""}
    var v31 int = int(v2)
    v6 -= v10
    v32 := GenS4{F0: false}
    var v33 float32 = 1e10
    var v34 uint8 = 177
    var v35 uint = 412
    var v36 uint64 = (v2 / ((v2) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
