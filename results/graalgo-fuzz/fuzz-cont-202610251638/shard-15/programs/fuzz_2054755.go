package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 bool }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    v3 := []uint32{2147483647, 2}
    var v4 float64 = 1.5
    v5 := make(map[uint32]string)
    if (969 <= 0) {
        switch "\n\t" {
        case "foo bar":
            var v6 uint64 = 42
            _ = v6
        case "":
            var v7 int = 128
            _ = v7
        case "\n\t":
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 float64 = v4
                fmt.Println(v4, v9, fi8)
                v4 = v9
                _ = fi8
                _ = v9
            }
        }
        var v10 string = "foo bar"
        v10 = (v10 + "foo bar")
        _ = v10
    }
    v11 := GenS1{F0: 938}
    fmt.Println(v4)
    if (-1 <= -536) {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 uint32 = 65535
            _ = fi12
            _ = v13
        }
        if false {
            fmt.Println(v4)
            v4 = (-2.25 * 0.0)
        } else {
            var v14 uint16 = 495
            var v15 uint8 = 127
            _ = v14
            _ = v15
        }
    } else {
        v4 += (((v4 - (v4 + v4)) - v4) + (1e10 - (v4 - v4)))
        v16 := GenS2{F0: 733}
        _ = v16
    }
    var v17 string = "a"
    v17 = ("\n\t" + "\n\t")
    fmt.Println(v17, v4)
    v1[v17] = (42 > 42)
    v18 := GenS1{F0: -32768}
    var v19 int32 = -1
    var v20 int8 = -60
    var v21 uint64 = 1
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 uint = 4294967295
        v24 := GenS2{F0: fi22}
        fmt.Println(v21)
        _ = fi22
        _ = v23
        _ = v24
    }
    if (((v4 - float64(3.14159)) - 3.14159) != (v4 / (v4 - float64(0.1)))) {
        if (0 > 111) {
            for fi25 := 0; fi25 < 9; fi25++ {
                if (2 != 255) {
                    for fi26 := 0; fi26 < 20; fi26++ {
                        v27 := GenS2{F0: fi25}
                        if true {
                            v28 := GenS2{F0: (fi25 / ((fi26) | 1))}
                            _ = v28
                        } else {
                            var v29 bool = false
                            v30 := GenS1{F0: 32767}
                            _ = v29
                            _ = v30
                        }
                        v19 += (^(v19))
                        _ = fi26
                        _ = v27
                    }
                    v17 = v17
                }
                fmt.Println(v19, v4, v21)
                fmt.Println(v19, v4)
                _ = fi25
            }
            for fi31 := 0; fi31 < 6; fi31++ {
                v32 := GenS1{F0: 124}
                _ = fi31
                _ = v32
            }
        } else {
            v17 = "hello"
        }
        var v33 bool = (2 > 0)
        for fi34 := 0; fi34 < 13; fi34++ {
            v35 := GenS3{F0: 0, F1: false}
            v20 = (v20 + (v20 % (((v20 % ((v20) | 1))) | 1)))
            _ = fi34
            _ = v35
        }
        _ = v33
    }
    for fi36 := 0; fi36 < 19; fi36++ {
        var v37 float64 = 0.0
        if ((^(v20)) >= v20) {
            v38 := GenS2{F0: (^(fi36))}
            if (v20 < v20) {
                delete(v5, 311)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v38
        }
        for fi39 := 0; fi39 < 13; fi39++ {
            if false {
                fmt.Println(v20, v4, v20)
            } else {
                delete(v1, v17)
            }
            switch "\n\t" {
            case "":
                fmt.Println(fi36)
            case "foo bar":
                if (-421 >= -9223372036854775808) {
                    v1[v17] = (!(((((0.1 - -2.25) * 1e10) + 0.1) == (0.1 / 1.5)) || (1 < 272)))
                    if true {
                        fmt.Println(fi36, v20)
                    }
                    var v40 bool = (false && false)
                    var v41 int8 = v20
                    fmt.Println(v4, v20, v21)
                    _ = v40
                    _ = v41
                }
            case "hello":
                if false {
                    var v42 uint16 = 0
                    _ = len(v3)
                    _ = v42
                }
            default:
                v19 += (v19 * v19)
            }
            var v43 uint = 921
            _ = fi39
            _ = v43
        }
        _ = fi36
        _ = v37
    }
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
