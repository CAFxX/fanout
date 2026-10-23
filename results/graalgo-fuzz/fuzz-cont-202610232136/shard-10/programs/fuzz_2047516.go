package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int32, p2 bool) uint32 {
    fmt.Println(p2, p0, p0)
    fmt.Println(p0, p0, p0)
    fmt.Println(p0, p0, p2)
    v7 := []int{}
    var v8 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 int16, p1 uint8, p2 int) int16 {
    p1 = (p1 ^ p1)
    switch "" {
    case "":
        v9 := GenS1{F0: (!true)}
        _ = v9
    case "hello":
        v10 := GenS2{F0: 1000, F1: "a"}
        _ = v10
    case "foo bar":
        v11 := []int{(p2 >> ((255) & 63)), p2}
        _ = v11
    default:
        v12 := GenS2{F0: p0, F1: "hello"}
        _ = v12
    }
    var v13 float64 = -2.25
    var v14 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return p0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -31
    v3 := make(map[int]int)
    var v4 uint32 = 2
    v5 := make(map[int]string)
    var v6 int64 = (v2 & v2)
    var v15 bool = (!(true || (!(2 >= 255))))
    for fi16 := 0; fi16 < 6; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            if v15 {
                v2 = (v6 + 487)
                var v18 float64 = (3.14159 * (3.14159 / 1.0))
                switch "a" {
                case "hello":
                    v19 := GenS3{F0: 42, F1: 1}
                    _ = v19
                }
                _ = v18
            } else {
                v2 = v6
                var v20 uint32 = (v4 % (((v4 << ((2) & 63))) | 1))
                _ = v20
            }
            _ = fi17
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 uint64 = 0
            v23 := []int{}
            fmt.Println(v2)
            if v15 {
                v15 = v15
                fmt.Println(fi16, v6, v15)
                v24 := v1[("foo bar" + "hello")]
                _ = v24
            } else {
                switch "hello" {
                case "":
                    v6 += (v6 >> ((314) & 63))
                }
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi16
    }
    var v25 uint = 65535
    v26 := GenS4{F0: v15, F1: v25}
    v4 -= v4
    fmt.Println(v4, v15)
    if (true && (-1 > 1)) {
        for fi27 := 0; fi27 < 2; fi27++ {
            switch fi27 {
            case 32767:
                v28 := GenS5{F0: 2, F1: "hello"}
                _ = v28
            case 128:
                v29 := []int{0, fi27, int(v4)}
                _ = v29
            default:
                for fi30 := 0; fi30 < 14; fi30++ {
                    switch 324 {
                    case 32767:
                        v31 := []int{(fi27 >> ((v25) & 63)), -743, (fi27 % (((fi27 - fi30)) | 1)), fi30}
                        _ = v31
                    default:
                        fmt.Println(fi30, v15)
                    }
                    _ = fi30
                }
            }
            _ = fi27
        }
    } else {
        v15 = v15
    }
    var v32 float32 = ((float32(0.0) - 0.1) + 0.0)
    var v33 bool = v15
    v34 := GenS4{F0: false, F1: 18446744073709551615}
    fmt.Println(v25, v2)
    if ("\n\t" != ("\n\t" + "")) {
        var v35 uint32 = v4
        if (((v32 / v32) >= v32) && v33) {
            v33 = true
            switch ("hello" + "") {
            case "":
                delete(v1, "a")
            default:
                switch 32768 {
                case 128:
                    if (255 == 270) {
                        v32 += float32((1e10 / 0.1))
                        fmt.Println(v2)
                    }
                default:
                    fmt.Println(v15, v4)
                }
            }
            v6 = 9223372036854775807
        }
        delete(v5, 1)
        _ = v35
    }
    for fi36 := 0; fi36 < 15; fi36++ {
        v33 = false
        var v37 float64 = 3.14159
        v33 = (v15 || (-525 > 0))
        _ = fi36
        _ = v37
    }
    v38 := GenS6{F0: 128}
    v39 := GenS3{F0: (v6 * int64(v25)), F1: 0}
    switch 0 {
    case 42:
        delete(v5, -1)
    case -1:
        v15 = v33
    case -9223372036854775808:
        v40 := GenS4{F0: (!false), F1: 4294967295}
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
