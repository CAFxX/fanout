package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 string }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 int16 }
func genFn0() uint {
    var v5 uint8 = 1
    v6 := GenS1{F0: 4294967295, F1: 42, F2: "hello"}
    v5 = v5
    var v7 bool = true
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = ("foo bar" + "\n\t")
    var v3 uint8 = 2
    var v4 int16 = -1
    v8 := []int{-228, -1, 246, 395}
    var v9 uint = 18446744073709551615
    v1[13] = ("hello" + (v2 + (v2 + "hello")))
    fmt.Println(v3, v9)
    v9 = (^(v9))
    switch v2 {
    case "a":
        var v10 uint64 = uint64(v9)
        _ = v10
    case "\n\t":
        var v11 float64 = 3.14159
        _ = v11
    case "foo bar":
        for fi12 := 0; fi12 < 14; fi12++ {
            v2 = "foo bar"
            fmt.Println(v2, v9, fi12)
            fmt.Println(v2, v4, v3)
            _ = fi12
        }
    }
    if (255 >= 592) {
        v13 := GenS1{F0: (v9 ^ v9), F1: 42, F2: "a"}
        var v14 uint64 = 1
        v3 -= 2
        _ = v13
        _ = v14
    }
    v2 = "foo bar"
    switch 4294967295 {
    case -2147483648:
        v15 := GenS2{F0: (float64(((1e10 + 0.1) + (3.14159 * ((float32((-2.25 - -2.25)) * 1e10) - (3.14159 * (1.5 * 3.14159)))))) >= float64(0.0)), F1: (!(0.1 >= 1.5))}
        _ = v15
    case -498:
        if true {
            for fi16 := 0; fi16 < 15; fi16++ {
                v9 += v9
                _ = fi16
            }
        } else {
            var v17 bool = (42 == 2147483647)
            for fi18 := 0; fi18 < 11; fi18++ {
                if (67 == -49) {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        var v20 float32 = 0.0
                        if v17 {
                            var v21 int64 = 0
                            var v22 int16 = v4
                            _ = v21
                            _ = v22
                        }
                        switch v2 {
                        case "\n\t":
                            var v23 int = fi18
                            _ = v23
                        default:
                            var v24 uint16 = 65535
                            _ = v24
                        }
                        _ = fi19
                        _ = v20
                    }
                }
                var v25 uint32 = 2147483647
                _ = fi18
                _ = v25
            }
            _ = v17
        }
    default:
        var v26 bool = (((-2.25 * 0.0) >= 0.0) || ((v2 + v2) > (((v2 + (((v2 + v2) + ("foo bar" + ("" + "a"))) + v2)) + "\n\t") + (v2 + (v2 + v2)))))
        _ = v26
    }
    v27 := GenS2{F0: false, F1: true}
    v2 = "hello"
    switch "foo bar" {
    case "a":
        var v28 uint32 = 809
        _ = v28
    default:
        if (0 <= 2147483647) {
            v29 := GenS2{F0: false, F1: (false || (!(32767 > v4)))}
            for fi30 := 0; fi30 < 15; fi30++ {
                fmt.Println(v9)
                var v31 int64 = 0
                _ = fi30
                _ = v31
            }
            _ = v29
        }
    }
    if false {
        var v32 int32 = 0
        _ = v32
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        v34 := GenS3{F0: -793}
        v3 -= v3
        v3 = v3
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        var v36 uint8 = 2
        _ = fi35
        _ = v36
    }
    if false {
        v9 -= (v9 - (v9 >> ((909) & 63)))
        fmt.Println(v9, v2, v2)
        for fi37 := 0; fi37 < 7; fi37++ {
            switch ("foo bar" + (v2 + v2)) {
            case "":
                var v38 float64 = float64((float32(1.5) + -2.25))
                _ = v38
            case "\n\t":
                for fi39 := 0; fi39 < 5; fi39++ {
                    if len(v8) > 0 {
                        v8[0] = (fi39 - (-(fi39)))
                    }
                    v3 -= 1
                    _ = fi39
                }
            case "a":
                switch fi37 {
                case 127:
                    var v40 uint = (^(v9))
                    _ = v40
                case 32767:
                    fmt.Println(v3, v4)
                case 128:
                    var v41 int32 = 2147483647
                    _ = v41
                default:
                    var v42 uint32 = 1
                    _ = v42
                }
            default:
                v1[0] = (v2 + v2)
            }
            _ = fi37
        }
    } else {
        delete(v1, 2)
        v43 := GenS4{F0: 1, F1: v4}
        v9 -= (v9 ^ genFn0())
        _ = v43
    }
    for fi44 := 0; fi44 < 7; fi44++ {
        var v45 uint32 = 4294967295
        _ = fi44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
