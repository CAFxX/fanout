package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint; F1 bool; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int8 }

func main() {
    var v1 uint = 65535
    var v2 float64 = 3.14159
    v3 := make(map[int]string)
    var v4 uint64 = 2
    var v5 int = -144
    v4 -= (v4 % (((v4 | v4)) | 1))
    for fi6 := 0; fi6 < 4; fi6++ {
        if true {
            var v7 bool = ((!(!true)) || (false || true))
            var v8 int32 = -994
            _ = v7
            _ = v8
        } else {
            v5 += (^(v5))
        }
        _ = fi6
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(v4)
        v2 -= v2
        fmt.Println(v2, v1, v5)
        _ = fi9
    }
    if (true && (!false)) {
        switch "foo bar" {
        case "a":
            var v10 int32 = 42
            _ = v10
        default:
            v11 := GenS1{F0: (v4 % (((v4 % (((v4 - (v4 / ((v4) | 1)))) | 1))) | 1)), F1: true}
            _ = v11
        }
    }
    v3[(^(v5))] = "a"
    var v12 int = v5
    fmt.Println(v1)
    if ((v5 * 523) > (v5 ^ v12)) {
        switch "\n\t" {
        case "foo bar":
            v12 = 0
        case "hello":
            var v13 uint8 = 22
            _ = v13
        case "":
            v3[-891] = "\n\t"
        }
    }
    v1 += v1
    if (0 != 32767) {
        switch "\n\t" {
        case "hello":
            var v14 bool = true
            _ = v14
        case "":
            v15 := GenS2{F0: v1, F1: true, F2: 74}
            _ = v15
        case "foo bar":
            v12 -= v5
        }
    } else {
        v4 -= (v4 - v4)
        var v16 int64 = 42
        _ = v16
    }
    v12 += v5
    var v17 uint8 = 2
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int32 = -1
        _ = fi18
        _ = v19
    }
    switch ("\n\t" + "hello") {
    case "":
        v20 := GenS3{F0: v12}
        _ = v20
    default:
        switch "" {
        case "\n\t":
            fmt.Println(v4, v1)
        case "":
            if (42 >= 0) {
                var v21 int32 = 968
                v5 += v12
                _ = v21
            } else {
                for fi22 := 0; fi22 < 16; fi22++ {
                    v17 += 127
                    _ = len(v3)
                    if true {
                        var v23 bool = (42 > (^(v4)))
                        _ = v23
                    } else {
                        v24 := GenS4{F0: true, F1: 128}
                        v25 := GenS3{F0: fi22}
                        _ = v24
                        _ = v25
                    }
                    _ = fi22
                }
                v2 += float64((-2.25 - -2.25))
                fmt.Println(v4, v5)
            }
        }
    }
    if ((v5 & -1) != v5) {
        for fi26 := 0; fi26 < 6; fi26++ {
            switch "hello" {
            case "foo bar":
                v5 = 4294967295
            case "":
                delete(v3, v12)
            case "\n\t":
                v12 = v5
            default:
                v27 := []int{128}
                _ = v27
            }
            _ = fi26
        }
    } else {
        if (-158 != v12) {
            if true {
                fmt.Println(v1, v2, v12)
                v2 += float64(1e10)
                v28 := []int{v12, (v12 << ((65535) & 63)), (v12 + 4294967295), -596}
                _ = v28
            } else {
                var v29 float32 = 0.1
                for fi30 := 0; fi30 < 11; fi30++ {
                    v31 := []int{-552}
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
            fmt.Println(v12, v1)
            var v32 uint64 = v4
            _ = v32
        } else {
            v12 -= v5
            v33 := GenS5{F0: v4, F1: -2}
            for fi34 := 0; fi34 < 20; fi34++ {
                v17 = 127
                for fi35 := 0; fi35 < 9; fi35++ {
                    v12 = -664
                    if true {
                        var v36 uint32 = 2147483648
                        _ = v36
                    }
                    _ = fi35
                }
                _ = fi34
            }
            _ = v33
        }
        var v37 int = v5
        var v38 int8 = 0
        _ = v37
        _ = v38
    }
    v12 = (-(v12))
    _ = len(v3)
    for fi39 := 0; fi39 < 1; fi39++ {
        v40 := []int{}
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
