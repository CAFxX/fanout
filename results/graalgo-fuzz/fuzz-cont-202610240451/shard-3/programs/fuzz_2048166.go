package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0(p0 int32, p1 uint, p2 string) uint32 {
    p2 = p2
    p2 = p2
    var v6 uint64 = 0
    if true {
        var v7 uint = p1
        _ = v7
    }
    v8 := GenS1{F0: 255, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 127
    var v3 string = ("foo bar" + "foo bar")
    var v4 uint16 = 2
    var v5 uint16 = (v4 ^ v4)
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: (v4 * v4), F1: v4}
        fmt.Println(v5, v3)
        _ = fi9
        _ = v10
    }
    switch v3 {
    case "a":
        if (1000 == -32768) {
            for fi11 := 0; fi11 < 1; fi11++ {
                for fi12 := 0; fi12 < 18; fi12++ {
                    fmt.Println(v3)
                    if (-476 < 9223372036854775807) {
                        var v13 uint16 = 42
                        var v14 bool = false
                        _ = v13
                        _ = v14
                    }
                    _ = fi12
                }
                switch ("a" + v3) {
                case "":
                    for fi15 := 0; fi15 < 7; fi15++ {
                        var v16 float32 = float32(0.0)
                        if ((true || (42 >= 127)) && ((0.1 / float64(v16)) == float64(v16))) {
                            v4 = v4
                            v5 = 240
                        } else {
                            var v17 int32 = 847
                            v1[v3] = (192 == v2)
                            _ = v17
                        }
                        _ = fi15
                        _ = v16
                    }
                case "hello":
                    v3 = ("foo bar" + v3)
                }
                v4 = v4
                _ = fi11
            }
        } else {
            switch "" {
            case "\n\t":
                delete(v1, "hello")
            default:
                var v18 int32 = -1
                _ = v18
            }
            for fi19 := 0; fi19 < 1; fi19++ {
                fmt.Println(v2, fi19, v2)
                v4 = v5
                _ = fi19
            }
        }
    default:
        var v20 int64 = 0
        _ = v20
    }
    v21 := GenS1{F0: v5, F1: v5}
    switch "hello" {
    case "a":
        v3 = v3
    case "\n\t":
        var v22 string = v3
        _ = v22
    }
    var v23 int64 = 42
    v24 := []int{320}
    if ((v3 != "\n\t") || (255 < 65535)) {
        for fi25 := 0; fi25 < 18; fi25++ {
            for fi26 := 0; fi26 < 12; fi26++ {
                switch 32767 {
                case -79:
                    _ = len(v1)
                default:
                    if (!false) {
                        v1["hello"] = (v4 != uint16(v2))
                    } else {
                        if false {
                            v23 -= (v23 / ((v23) | 1))
                        } else {
                            v24 = append(v24, fi25)
                        }
                        v5 = v5
                        v3 = v3
                    }
                }
                v5 -= v4
                v24 = append(v24, (fi25 & fi26))
                _ = fi26
            }
            _ = len(v1)
            _ = fi25
        }
        var v27 uint = 2
        _ = v27
    } else {
        fmt.Println(v5, v4, v4)
        v28 := GenS2{F0: 2147483647, F1: 42, F2: v4}
        var v29 float64 = float64((((1e10 * 1.5) * 1.5) + 0.1))
        _ = v28
        _ = v29
    }
    if true {
        if (!false) {
            v5 = 2
            v30 := GenS1{F0: v5, F1: v5}
            if true {
                for fi31 := 0; fi31 < 1; fi31++ {
                    v3 = v3
                    _ = fi31
                }
                var v32 uint32 = 2147483647
                _ = v32
            }
            _ = v30
        } else {
            v5 = (v5 >> ((93) & 63))
            for fi33 := 0; fi33 < 3; fi33++ {
                var v34 int32 = 975
                v35 := GenS3{F0: (^(v4)), F1: (fi33 - (fi33 % (((fi33 >> ((655) & 63))) | 1)))}
                _ = fi33
                _ = v34
                _ = v35
            }
        }
        if true {
            for fi36 := 0; fi36 < 14; fi36++ {
                var v37 uint32 = uint32(v5)
                fmt.Println(v37)
                var v38 uint = uint(v2)
                _ = fi36
                _ = v37
                _ = v38
            }
        }
        var v39 int8 = -5
        _ = v39
    }
    fmt.Println(v23, v3)
    v40 := GenS2{F0: 30, F1: 104, F2: 575}
    if (-32768 >= 42) {
        var v41 int16 = -95
        v42 := GenS3{F0: 629, F1: 127}
        _ = v41
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
