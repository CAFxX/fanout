package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string }

func main() {
    var v1 float64 = 0.1
    var v2 uint8 = 42
    var v3 int64 = -836
    v4 := []uint32{42, 450, 42, 65535}
    v5 := []int{1, 2147483647}
    var v6 float64 = (((v1 * ((v1 - v1) - (v1 + v1))) + v1) * 0.0)
    if (!(0 != -128)) {
        if (!true) {
            v6 += v6
        }
    }
    var v7 uint16 = 2
    fmt.Println(v1, v2)
    v2 = (v2 * 128)
    v6 += 3.14159
    var v8 bool = (128 > 9223372036854775807)
    var v9 uint = uint(v7)
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 3; fi11++ {
            switch "foo bar" {
            case "foo bar":
                v1 = -2.25
                switch "a" {
                case "foo bar":
                    if (!v8) {
                        var v12 int16 = 32767
                        fmt.Println(v8, fi11, v8)
                        _ = v12
                    }
                case "":
                    var v13 uint = v9
                    _ = v13
                case "a":
                    if (false || (855 > uint32(v9))) {
                        v6 += (0.0 - 1.5)
                    } else {
                        var v14 uint = v9
                        v15 := GenS1{F0: 973, F1: (v7 + v7)}
                        _ = v14
                        _ = v15
                    }
                default:
                    v16 := GenS2{F0: (v3 / ((-1) | 1)), F1: (fi11 ^ (fi11 + fi10))}
                    _ = v16
                }
            case "":
                var v17 int = fi11
                _ = v17
            case "\n\t":
                if v8 {
                    fmt.Println(v6, fi11, v3)
                    v18 := GenS3{F0: 0}
                    var v19 string = "hello"
                    _ = v18
                    _ = v19
                }
            default:
                v3 = 42
                if (float32(((float64(0.1) / float64(float32(v1))) - v6)) >= (1.5 + (((3.14159 / 3.14159) + -2.25) - 1e10))) {
                    if len(v4) > 0 {
                        v4[0] = 2147483648
                    }
                    v5 = append(v5, (fi10 * 127))
                } else {
                    var v20 uint32 = 125
                    switch "foo bar" {
                    case "a":
                        var v21 uint64 = 4294967295
                        _ = v21
                    case "hello":
                        fmt.Println(v7, v8, v2)
                    case "\n\t":
                        v1 += float64(1e10)
                    }
                    var v22 int64 = 0
                    _ = v20
                    _ = v22
                }
            }
            _ = fi11
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi23
        }
        _ = fi10
    }
    if (!v8) {
        if (true || v8) {
            switch "a" {
            case "a":
                if v8 {
                    v24 := GenS1{F0: 417, F1: v7}
                    _ = v24
                } else {
                    v1 -= (v6 * -2.25)
                    var v25 uint16 = (v7 / (((v7 - (v7 / ((42) | 1)))) | 1))
                    _ = v25
                }
            default:
                v7 = (v7 * 32767)
            }
            v2 = v2
        }
        v1 -= v1
        var v26 string = "a"
        _ = v26
    } else {
        v27 := GenS2{F0: v3, F1: 342}
        for fi28 := 0; fi28 < 20; fi28++ {
            var v29 int = fi28
            v30 := GenS4{F0: ("\n\t" + "\n\t")}
            switch 765 {
            case -455:
                var v31 float32 = 1e10
                _ = v31
            case 116:
                v2 = 255
            case 2147483647:
                fmt.Println(v8, v8, fi28)
            default:
                var v32 int32 = 196
                _ = v32
            }
            _ = fi28
            _ = v29
            _ = v30
        }
        var v33 uint32 = 4294967295
        _ = v27
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
