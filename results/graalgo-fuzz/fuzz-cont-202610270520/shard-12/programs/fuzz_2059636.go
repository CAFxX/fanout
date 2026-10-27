package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 string, p1 int64, p2 uint) uint {
    v5 := GenS1{F0: -48}
    for fi6 := 0; fi6 < 10; fi6++ {
        p0 = p0
        p2 += p2
        var v7 float64 = (1e10 - (0.1 * (1.5 - (-2.25 * float64(float32(1e10))))))
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 9; fi8++ {
        p1 = 9223372036854775807
        _ = fi8
    }
    switch (p0 + p0) {
    case "":
        var v9 uint = (p2 >> ((42) & 63))
        _ = v9
    case "\n\t":
        p0 = (p0 + "")
    case "hello":
        var v10 float32 = float32(1e10)
        _ = v10
    default:
        var v11 uint32 = 2147483648
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p2 ^ (p2 ^ (p2 % (((p2 - 207)) | 1))))
}

func main() {
    v1 := []int{}
    v2 := make(map[string]uint32)
    var v3 uint8 = 255
    v4 := make(map[uint32]bool)
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            switch (fi13 + fi13) {
            case 127:
                v14 := GenS1{F0: 0}
                _ = v14
            case 810:
                v3 = (v3 * (v3 << ((42) & 63)))
            default:
                var v15 int8 = 83
                _ = v15
            }
            fmt.Println(v3, v3)
            switch "\n\t" {
            case "foo bar":
                v3 = 2
            default:
                switch "foo bar" {
                case "hello":
                    fmt.Println(fi13, fi12, fi13)
                default:
                    if true {
                        _ = len(v4)
                        var v16 uint16 = 65535
                        v16 -= (v16 + v16)
                        _ = v16
                    } else {
                        var v17 int16 = -32768
                        _ = v17
                    }
                }
            }
            var v18 uint = 205
            _ = fi13
            _ = v18
        }
        _ = fi12
    }
    switch 1 {
    case 32767:
        v3 -= v3
    }
    if true {
        switch 1 {
        case 0:
            var v19 int = 42
            _ = v19
        case 369:
            var v20 int16 = -32768
            _ = v20
        default:
            v3 -= 1
        }
        v3 = v3
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 int = fi21
            _ = fi21
            _ = v22
        }
    } else {
        var v23 uint32 = 168
        v23 -= v23
        _ = v23
    }
    v3 = v3
    fmt.Println(v3)
    if len(v1) > 0 {
        v1[0] = 127
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 int16 = int16(fi24)
        fmt.Println(v25)
        if false {
            if (!(111 == 0)) {
                switch fi24 {
                case -9223372036854775808:
                    fmt.Println(v3)
                case -2147483648:
                    for fi26 := 0; fi26 < 15; fi26++ {
                        fmt.Println(fi24)
                        if (!(true || (2 != uint32(v3)))) {
                            var v27 float64 = 1.5
                            switch (^(fi26)) {
                            case -69:
                                var v28 string = ("a" + (("" + ("\n\t" + "a")) + "a"))
                                _ = v28
                            case 963:
                                var v29 int8 = 0
                                _ = v29
                            }
                            _ = v27
                        }
                        _ = fi26
                    }
                case 9223372036854775807:
                    fmt.Println(fi24, fi24)
                }
                v30 := GenS1{F0: -1}
                _ = v30
            } else {
                var v31 bool = ((^(v25)) < v25)
                _ = v31
            }
            var v32 uint32 = 396
            _ = v32
        }
        _ = fi24
        _ = v25
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        v3 -= v3
        _ = fi33
    }
    switch "foo bar" {
    case "":
        v3 += v3
    default:
        for fi34 := 0; fi34 < 17; fi34++ {
            v3 = 0
            _ = fi34
        }
    }
    v35 := GenS1{F0: 0}
    v3 -= v3
    fmt.Println(v3)
    if (!(42 > 1)) {
        for fi36 := 0; fi36 < 8; fi36++ {
            var v37 uint = 4294967295
            fmt.Println(v3, v3, v37)
            _ = fi36
            _ = v37
        }
    } else {
        var v38 int16 = -1
        var v39 bool = (1 >= 0)
        for fi40 := 0; fi40 < 10; fi40++ {
            v3 += 127
            fmt.Println(fi40, v39, v38)
            _ = fi40
        }
        _ = v38
        _ = v39
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
