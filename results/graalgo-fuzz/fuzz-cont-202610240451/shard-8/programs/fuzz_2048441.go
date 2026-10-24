package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint }
func genFn0(p0 bool, p1 string, p2 int64) int8 {
    fmt.Println(p0, p1)
    var v5 string = ((p1 + p1) + (p1 + p1))
    v6 := []int{42, 98}
    switch 895 {
    case 301:
        v7 := GenS1{F0: 2, F1: 802, F2: 127}
        _ = v7
    case 9223372036854775807:
        var v8 string = ((p1 + v5) + v5)
        _ = v8
    case -9223372036854775808:
        for fi9 := 0; fi9 < 20; fi9++ {
            p2 = 1
            var v10 int64 = p2
            for fi11 := 0; fi11 < 18; fi11++ {
                fmt.Println(p0)
                v5 = "hello"
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
    default:
        for fi12 := 0; fi12 < 18; fi12++ {
            if len(v6) > 0 {
                v6[0] = (^(fi12))
            }
            _ = fi12
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{("foo bar" + ("\n\t" + "foo bar")), "\n\t", "a"}
    v3 := []int{-738, -803, 0, -2147483648}
    var v4 bool = (false || false)
    if false {
        var v13 float64 = 1e10
        _ = v13
    } else {
        var v14 int16 = 1
        v4 = false
        _ = v14
    }
    var v15 int16 = -1
    if false {
        var v16 uint32 = 65535
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint64 = 18446744073709551615
        v4 = false
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: 4294967295, F1: 648, F2: 127}
    switch "\n\t" {
    case "\n\t":
        if (42 < 1) {
            v20 := GenS2{F0: 65535}
            switch -263 {
            case 32767:
                v15 = (^(v15))
            }
            switch -126 {
            case -753:
                switch ("" + "hello") {
                case "foo bar":
                    v15 += 1
                }
            case 9223372036854775807:
                v4 = v4
            }
            _ = v20
        } else {
            v21 := GenS1{F0: 4294967295, F1: 86, F2: 30}
            _ = v21
        }
    }
    if v4 {
        switch "foo bar" {
        case "":
            switch 4294967295 {
            case 32768:
                var v22 uint64 = 0
                _ = v22
            default:
                v15 -= 42
            }
        case "\n\t":
            v15 = v15
        case "hello":
            var v23 string = ("\n\t" + "hello")
            _ = v23
        }
    } else {
        v24 := v1["a"]
        v15 += (v15 - (v15 & v15))
        fmt.Println(v4)
        _ = v24
    }
    if (v4 || true) {
        fmt.Println(v15, v4)
        v15 = (v15 << ((4294967295) & 63))
        for fi25 := 0; fi25 < 15; fi25++ {
            v26 := GenS1{F0: 42, F1: 4294967295, F2: 153}
            _ = fi25
            _ = v26
        }
    }
    fmt.Println(v15, v15)
    fmt.Println(v15, v4)
    if v4 {
        fmt.Println(v15)
        if true {
            var v27 bool = v4
            if (32767 > 4294967295) {
                if (v27 || false) {
                    v4 = ((0.0 == (((3.14159 + 0.0) * 1.5) / 1.5)) || true)
                    var v28 float32 = -2.25
                    _ = v28
                }
            } else {
                switch 32767 {
                case -213:
                    var v29 int64 = 1
                    _ = v29
                case -77:
                    var v30 uint32 = 4294967295
                    _ = v30
                default:
                    for fi31 := 0; fi31 < 6; fi31++ {
                        v15 -= (v15 * (-(v15)))
                        v32 := GenS1{F0: 597, F1: 0, F2: 42}
                        for fi33 := 0; fi33 < 17; fi33++ {
                            var v34 uint64 = 42
                            _ = fi33
                            _ = v34
                        }
                        _ = fi31
                        _ = v32
                    }
                }
                if (-128 == 127) {
                    var v35 string = ""
                    if true {
                        var v36 uint32 = 65535
                        fmt.Println(v4, v15)
                        _ = v36
                    } else {
                        var v37 string = (v35 + (v35 + v35))
                        v15 = v15
                        _ = v37
                    }
                    _ = v35
                } else {
                    switch -39 {
                    case 42:
                        var v38 int64 = 1
                        _ = v38
                    case 2147483647:
                        for fi39 := 0; fi39 < 9; fi39++ {
                            var v40 int16 = (v15 ^ v15)
                            _ = fi39
                            _ = v40
                        }
                    case 32768:
                        fmt.Println(v27, v4)
                    default:
                        if v4 {
                            v4 = v27
                            v4 = false
                        }
                    }
                }
            }
            _ = v27
        }
    } else {
        v15 += v15
        var v41 float64 = 3.14159
        _ = v41
    }
    v42 := GenS3{F0: 0}
    if true {
        fmt.Println(v4, v15)
        v4 = (!(!false))
        v43 := GenS2{F0: 748}
        _ = v43
    }
    var v44 bool = false
    var v45 int32 = 113
    var v46 uint64 = 134
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v45
    _ = v46
}
