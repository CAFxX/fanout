package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int8, p1 int16, p2 bool) uint64 {
    p0 += (p0 >> ((2) & 63))
    for fi5 := 0; fi5 < 17; fi5++ {
        if p2 {
            var v6 uint8 = 0
            _ = v6
        } else {
            fmt.Println(p1)
            v7 := GenS1{F0: 44}
            var v8 int = (-(fi5))
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            if p2 {
                p2 = (!(false || p2))
                var v10 int64 = -1
                _ = v10
            }
            _ = fi9
        }
        var v11 uint64 = 4294967295
        _ = fi5
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]bool)
    v3 := []int{-852, 32768, 42}
    var v4 int8 = 127
    v4 -= -107
    v12 := v2[42]
    if v12 {
        var v13 int32 = 1
        _ = v13
    } else {
        for fi14 := 0; fi14 < 14; fi14++ {
            fmt.Println(fi14)
            switch "foo bar" {
            case "a":
                for fi15 := 0; fi15 < 10; fi15++ {
                    _ = len(v3)
                    var v16 string = ((("foo bar" + "a") + "\n\t") + "\n\t")
                    _ = fi15
                    _ = v16
                }
            case "foo bar":
                for fi17 := 0; fi17 < 14; fi17++ {
                    v12 = (!false)
                    v18 := GenS2{F0: 0, F1: v12}
                    var v19 bool = v12
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            default:
                var v20 uint32 = 2147483647
                _ = v20
            }
            fmt.Println(fi14, v12)
            _ = fi14
        }
        v4 = 0
        var v21 string = "\n\t"
        _ = v21
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    v12 = v12
    fmt.Println(v4, v4)
    var v22 int8 = (^(v4))
    switch ("hello" + ("" + ("foo bar" + "a"))) {
    case "hello":
        v22 = v22
    case "":
        if (!v12) {
            switch -2147483648 {
            case 32768:
                fmt.Println(v4)
            case 675:
                for fi23 := 0; fi23 < 2; fi23++ {
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v2[32768] = (4294967295 < 65535)
                        var v25 string = "hello"
                        _ = fi24
                        _ = v25
                    }
                    for fi26 := 0; fi26 < 19; fi26++ {
                        if (42 < 213) {
                            fmt.Println(fi23, v12)
                            var v27 int8 = (v22 + -128)
                            _ = len(v3)
                            _ = v27
                        } else {
                            v28 := GenS3{F0: -9223372036854775808}
                            _ = v28
                        }
                        v29 := GenS4{F0: 9223372036854775807}
                        _ = fi26
                        _ = v29
                    }
                    _ = fi23
                }
            case 32767:
                if false {
                    switch 9223372036854775807 {
                    case 32767:
                        if false {
                            var v30 int = 9223372036854775807
                            _ = v30
                        } else {
                            v1["a"] = 42
                            var v31 float32 = 1.5
                            delete(v2, 2147483647)
                            _ = v31
                        }
                    case 127:
                        fmt.Println(v4)
                    case -581:
                        var v32 int16 = 32767
                        _ = v32
                    }
                    fmt.Println(v22)
                } else {
                    var v33 float64 = float64(float32(float64(float32(1.5))))
                    _ = v33
                }
            default:
                delete(v2, -712)
            }
            v34 := GenS3{F0: 954}
            if (!(!((899 != 673) || (1 != 2147483647)))) {
                delete(v2, -119)
                v35 := v2[359]
                _ = v35
            }
            _ = v34
        } else {
            var v36 int = 128
            for fi37 := 0; fi37 < 2; fi37++ {
                fmt.Println(v36, v36)
                v38 := GenS3{F0: 42}
                v12 = (!true)
                _ = fi37
                _ = v38
            }
            _ = v36
        }
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        v4 -= 100
        for fi40 := 0; fi40 < 16; fi40++ {
            v4 = (v22 / (((v22 % ((v4) | 1))) | 1))
            _ = fi40
        }
        var v41 float32 = -2.25
        _ = fi39
        _ = v41
    }
    v22 = (v4 + v22)
    var v42 float32 = float32((1e10 - (1.5 * (0.1 * 1e10))))
    var v43 int32 = 1
    for fi44 := 0; fi44 < 11; fi44++ {
        fmt.Println(v42, v22)
        for fi45 := 0; fi45 < 7; fi45++ {
            fmt.Println(v42, v43)
            _ = fi45
        }
        _ = fi44
    }
    if v12 {
        v43 -= 0
        var v46 uint64 = 2
        _ = v46
    }
    v47 := GenS5{F0: 89}
    if ((201 != 2) || false) {
        v48 := GenS4{F0: 1}
        v49 := GenS6{F0: 0}
        v4 = v4
        _ = v48
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v47
}
