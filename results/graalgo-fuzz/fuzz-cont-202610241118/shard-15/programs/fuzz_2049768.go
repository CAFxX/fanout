package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint16 }
type GenS2 struct { F0 string; F1 uint64; F2 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 uint64; F2 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8 }

func main() {
    var v1 bool = (true && (((0.1 + 0.1) < 3.14159) || true))
    var v2 uint = 1
    var v3 float64 = float64((1.5 / 0.1))
    var v4 uint8 = 255
    v5 := make(map[int]uint32)
    var v6 uint8 = 127
    var v7 int16 = 42
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            delete(v5, (fi8 | fi8))
            if v1 {
                v10 := GenS1{F0: (fi8 >> ((255) & 63)), F1: 412, F2: 0}
                _ = v10
            } else {
                v3 -= v3
                if true {
                    v11 := GenS2{F0: (("a" + "foo bar") + "a"), F1: 212, F2: 65535}
                    _ = v11
                } else {
                    var v12 float32 = ((3.14159 - 0.0) * (0.1 * 0.1))
                    v1 = (!true)
                    switch "" {
                    case "":
                        v3 = (3.14159 - float64(0.1))
                        fmt.Println(v4, fi8)
                        var v13 string = "a"
                        _ = v13
                    }
                    _ = v12
                }
                v14 := GenS1{F0: 32767, F1: 4294967295, F2: 42}
                _ = v14
            }
            v7 -= 1
            _ = fi9
        }
        _ = fi8
    }
    v2 = v2
    v4 += v6
    if (!((255 == 42) && ((v3 * v3) > (((v3 + float64(3.14159)) / float64(1e10)) / 3.14159)))) {
        if (!v1) {
            for fi15 := 0; fi15 < 20; fi15++ {
                if v1 {
                    v2 += v2
                    fmt.Println(v4, fi15)
                }
                if (v1 || (!v1)) {
                    var v16 float64 = v3
                    var v17 uint16 = 42
                    _ = v16
                    _ = v17
                }
                v3 = (((1e10 + v3) + v3) * (float64(0.0) + (v3 + v3)))
                _ = fi15
            }
        } else {
            var v18 uint = 837
            if ((v3 * v3) == (v3 + float64((0.1 + 0.1)))) {
                v19 := []int{}
                _ = v19
            }
            v20 := GenS2{F0: "foo bar", F1: 595, F2: 65535}
            _ = v18
            _ = v20
        }
    }
    var v21 uint32 = 0
    v22 := GenS1{F0: 0, F1: 4294967295, F2: 320}
    var v23 int16 = v7
    switch "" {
    case "a":
        if v1 {
            var v24 uint8 = (v4 / ((v6) | 1))
            var v25 uint64 = 42
            _ = v24
            _ = v25
        }
    default:
        var v26 string = ""
        _ = v26
    }
    if (156 != 445) {
        v1 = (v1 || true)
        v27 := GenS3{F0: 765}
        _ = v27
    } else {
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 float32 = -2.25
            _ = fi28
            _ = v29
        }
        v30 := GenS3{F0: 234}
        if (((!v1) && false) || (!((3.14159 - 1e10) <= (1.5 * 3.14159)))) {
            v5[32767] = v21
            var v31 float64 = 3.14159
            _ = v31
        } else {
            var v32 int16 = (v7 >> ((v2) & 63))
            v33 := GenS4{F0: 0, F1: 42, F2: (-1 < 1)}
            _ = v32
            _ = v33
        }
        _ = v30
    }
    var v34 uint8 = v4
    v35 := GenS5{F0: "foo bar"}
    v36 := []int{9223372036854775807, 128, 128}
    if ((!v1) && (true || false)) {
        var v37 int32 = -49
        _ = v37
    } else {
        var v38 uint8 = v6
        var v39 uint64 = 42
        _ = v38
        _ = v39
    }
    fmt.Println(v21, v3)
    v40 := GenS6{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
