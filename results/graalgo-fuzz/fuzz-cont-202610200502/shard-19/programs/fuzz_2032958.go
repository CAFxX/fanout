package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int8 = 127
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    v4 := []int{}
    var v5 uint16 = 991
    var v6 int64 = 42
    switch -936 {
    case 9223372036854775807:
        var v7 string = ("" + "foo bar")
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        if true {
            for fi9 := 0; fi9 < 8; fi9++ {
                switch (fi9 | (^(fi8))) {
                case 1:
                    _ = len(v4)
                    switch "foo bar" {
                    case "a":
                        if false {
                            v5 = (v5 - v5)
                            v1 = 1
                            v10 := GenS1{F0: 67, F1: 78}
                            _ = v10
                        } else {
                            v1 = (-(v1))
                            v6 -= v6
                        }
                    case "foo bar":
                        var v11 uint64 = 0
                        _ = v11
                    }
                case 128:
                    var v12 int16 = 0
                    _ = v12
                case 32767:
                    var v13 uint32 = 968
                    _ = v13
                default:
                    v4 = append(v4, (fi8 / (((fi8 * fi9)) | 1)))
                }
                var v14 int32 = 503
                var v15 int32 = (v14 / ((v14) | 1))
                delete(v2, 42)
                _ = fi9
                _ = v14
                _ = v15
            }
            if (("hello" + "") > ("" + "\n\t")) {
                v6 -= 42
            } else {
                var v16 uint = 42
                v6 = (v6 ^ (^(v6)))
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 19; fi17++ {
                var v18 int = fi8
                _ = fi17
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS1{F0: 290, F1: 2147483647}
            _ = fi19
            _ = v20
        }
        _ = fi8
    }
    switch 1 {
    case -2147483648:
        if false {
            fmt.Println(v5, v1)
        } else {
            for fi21 := 0; fi21 < 2; fi21++ {
                v1 += v1
                for fi22 := 0; fi22 < 17; fi22++ {
                    v5 -= v5
                    var v23 int32 = 2147483647
                    _ = fi22
                    _ = v23
                }
                var v24 uint32 = 537
                _ = fi21
                _ = v24
            }
            v5 -= (v5 % ((32767) | 1))
        }
    case -9223372036854775808:
        if (2 > 877) {
            var v25 uint8 = 1
            v26 := GenS2{F0: (v25 + 1), F1: -2147483648}
            _ = v25
            _ = v26
        } else {
            v5 -= 1
            fmt.Println(v5, v1, v5)
            v1 = v1
        }
    }
    for fi27 := 0; fi27 < 8; fi27++ {
        fmt.Println(v5, fi27, fi27)
        v28 := GenS1{F0: 1, F1: 42}
        _ = fi27
        _ = v28
    }
    switch ("a" + ("\n\t" + "a")) {
    case "":
        fmt.Println(v1)
    case "foo bar":
        fmt.Println(v1, v1)
    default:
        if false {
            var v29 uint16 = (v5 | v5)
            var v30 string = "foo bar"
            v5 = (^(v5))
            _ = v29
            _ = v30
        } else {
            if (2 == 179) {
                switch 4294967295 {
                case 0:
                    var v31 float64 = (0.0 * 3.14159)
                    _ = v31
                default:
                    var v32 int = -2147483648
                    _ = v32
                }
                if ((-95 >= 42) || true) {
                    var v33 int64 = (v6 & 9223372036854775807)
                    v5 = (v5 >> ((178) & 63))
                    _ = v33
                } else {
                    fmt.Println(v6)
                }
                var v34 float64 = float64(3.14159)
                _ = v34
            } else {
                v35 := GenS3{F0: 223}
                _ = v35
            }
            v4 = append(v4, -806)
        }
    }
    v4 = append(v4, 0)
    var v36 string = ("" + "\n\t")
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
