package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 string; F2 bool }
func genFn0(p0 uint16, p1 int8, p2 string) int32 {
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 1, F1: 0}
        p0 -= p0
        p0 = 2
        _ = fi6
        _ = v7
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 string = "a"
    var v2 uint16 = 255
    var v3 uint32 = 1
    v4 := []string{v1, ("" + v1)}
    v5 := make(map[int]bool)
    if (float64((0.0 * (3.14159 / 1.0))) > 1.5) {
        var v8 string = v1
        _ = v8
    }
    if (828 > 0) {
        var v9 uint64 = 1
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        if (false && (0 > 42)) {
            v3 += (v3 >> ((42) & 63))
            delete(v5, fi10)
            v2 = v2
        } else {
            fmt.Println(fi10)
        }
        switch (fi10 & (^(fi10))) {
        case -2147483648:
            if (!(1 != 18446744073709551615)) {
                v3 = v3
            }
        case 128:
            v2 -= v2
        case 32768:
            v1 = "a"
        }
        if false {
            if (fi10 > (fi10 ^ (^(fi10)))) {
                var v11 uint = 123
                var v12 int64 = -18
                _ = v11
                _ = v12
            } else {
                switch (fi10 << ((2) & 63)) {
                case -2147483648:
                    v13 := GenS2{F0: false}
                    _ = v13
                case 2147483647:
                    var v14 float32 = 0.0
                    _ = v14
                default:
                    fmt.Println(v1, fi10, fi10)
                }
            }
            for fi15 := 0; fi15 < 14; fi15++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v16 uint32 = (v3 % ((v3) | 1))
                _ = fi15
                _ = v16
            }
        }
        _ = fi10
    }
    fmt.Println(v3, v1, v1)
    var v17 uint = 2
    if len(v4) > 0 {
        v4[0] = ((v1 + ("" + v1)) + "foo bar")
    }
    fmt.Println(v17, v17)
    var v18 uint64 = 18446744073709551615
    if true {
        fmt.Println(v3)
        v18 -= (v18 / ((v18) | 1))
    } else {
        fmt.Println(v2, v17, v2)
    }
    var v19 uint = (^(v17))
    for fi20 := 0; fi20 < 10; fi20++ {
        switch ("\n\t" + (v1 + v1)) {
        case "a":
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 uint32 = (v3 / (((v3 & v3)) | 1))
                v18 = 18446744073709551615
                var v23 float64 = -2.25
                _ = fi21
                _ = v22
                _ = v23
            }
        case "foo bar":
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 float32 = ((1e10 / 1e10) + (1e10 * (-2.25 * 0.0)))
                v3 = v3
                if (!true) {
                    var v26 int64 = 42
                    switch v1 {
                    case "":
                        var v27 float64 = (((1e10 - -2.25) - ((0.0 / (((0.1 * float64((v25 * v25))) / 1.0) / 0.1)) - -2.25)) + -2.25)
                        _ = v27
                    }
                    switch 128 {
                    case -9223372036854775808:
                        if false {
                            v2 -= v2
                            v3 = v3
                        }
                    case 32767:
                        var v28 uint16 = v2
                        _ = v28
                    case 127:
                        v1 = v1
                    default:
                        fmt.Println(v26, v2)
                    }
                    _ = v26
                }
                _ = fi24
                _ = v25
            }
        case "\n\t":
            v4 = append(v4, (v1 + v1))
        }
        if (((0 >= int8(v2)) && (v3 < (v3 << ((255) & 63)))) || true) {
            var v29 bool = true
            fmt.Println(v18, fi20)
            var v30 int32 = 2147483647
            _ = v29
            _ = v30
        } else {
            fmt.Println(v18, v3, v17)
        }
        v31 := GenS3{F0: -521, F1: v1, F2: true}
        _ = fi20
        _ = v31
    }
    if false {
        var v32 uint8 = 0
        var v33 int64 = -9223372036854775808
        _ = v32
        _ = v33
    }
    if (true && false) {
        var v34 int8 = -26
        var v35 uint16 = 42
        v1 = v1
        _ = v34
        _ = v35
    } else {
        if true {
            var v36 uint64 = 963
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
