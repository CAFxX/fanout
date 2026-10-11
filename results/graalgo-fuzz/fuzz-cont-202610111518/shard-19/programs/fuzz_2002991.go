package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }
type GenS2 struct { F0 int8; F1 uint32; F2 string }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 int64; F1 int64; F2 int16 }
type GenS6 struct { F0 int; F1 int8 }
func genFn0(p0 uint, p1 int8) uint16 {
    p0 += p0
    var v4 uint8 = 101
    p0 += 11
    p1 += (p1 >> ((p0) & 63))
    _ = p0
    _ = p1
    _ = v4
    return 2
}
func genFn1(p0 int64, p1 uint) int8 {
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            switch "foo bar" {
            case "hello":
                var v7 int16 = -1
                _ = v7
            case "foo bar":
                p1 = 0
            default:
                var v8 bool = false
                _ = v8
            }
            _ = fi6
        }
        var v9 uint16 = 32767
        var v10 uint64 = 18446744073709551615
        _ = fi5
        _ = v10
        _ = v9
    }
    var v11 int16 = -32768
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 bool = true
    v12 := GenS1{F0: 0, F1: 32767, F2: 110}
    v13 := v2[808]
    switch 42 {
    case 127:
        var v14 uint8 = 0
        _ = v14
    case 32767:
        v2[2147483647] = v13
    default:
        v3 = true
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 uint16 = 255
        switch -2147483648 {
        case -2147483648:
            v3 = (0.0 < (3.14159 - 1.5))
        default:
            var v17 uint8 = 0
            _ = v17
        }
        if true {
            fmt.Println(v3, v3, v13)
            v18 := GenS2{F0: 100, F1: 2147483648, F2: v13}
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 bool = true
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    switch 127 {
    case 127:
        v3 = false
    }
    v13 = "\n\t"
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 uint = 42
        switch (("\n\t" + (v13 + v13)) + v13) {
        case "":
            v22 += 2
        case "hello":
            for fi23 := 0; fi23 < 3; fi23++ {
                if false {
                    v13 = v13
                    v24 := GenS1{F0: 0, F1: 807, F2: v22}
                    switch "" {
                    case "hello":
                        v25 := v1[v13]
                        _ = v25
                    case "foo bar":
                        var v26 int32 = -2147483648
                        _ = v26
                    }
                    _ = v24
                } else {
                    v27 := []int{(fi21 * fi23)}
                    v28 := GenS2{F0: 0, F1: 4294967295, F2: "foo bar"}
                    if ((0.1 + 1.5) != -2.25) {
                        fmt.Println(v22, v3, v13)
                        v3 = (!(255 < 0))
                        var v29 int8 = 1
                        if len(v27) > 0 {
                            v27[0] = fi23
                        }
                        _ = v29
                    }
                    v30 := GenS3{F0: 491, F1: -128}
                    _ = v27
                    _ = v28
                    _ = v30
                }
                var v31 uint = 42
                _ = fi23
                _ = v31
            }
        case "\n\t":
            var v32 string = (v13 + (("foo bar" + v13) + ("a" + v13)))
            _ = v32
        }
        v13 = v13
        _ = fi21
        _ = v22
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        switch 42 {
        case 32768:
            v3 = (-128 >= 1)
        case 1:
            var v34 uint16 = 32767
            _ = v34
        default:
            var v35 float64 = 0.0
            _ = v35
        }
        v36 := GenS4{F0: -32768, F1: -668}
        var v37 int8 = 1
        _ = fi33
        _ = v36
        _ = v37
    }
    v1[v13] = 4294967295
    for fi38 := 0; fi38 < 10; fi38++ {
        fmt.Println(v13, v13, fi38)
        v39 := GenS2{F0: 0, F1: 4294967295, F2: v13}
        for fi40 := 0; fi40 < 12; fi40++ {
            v41 := []int{}
            _ = fi40
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    switch 0 {
    case 4294967295:
        for fi42 := 0; fi42 < 13; fi42++ {
            v43 := GenS5{F0: -1, F1: 42, F2: 42}
            if (!((((!v3) && (42 > 4294967295)) || (false || v3)) || v3)) {
                v3 = v3
                v13 = (v13 + "")
            } else {
                var v44 bool = (-672 > 2147483647)
                v45 := GenS6{F0: fi42, F1: 0}
                fmt.Println(v13, v3, v3)
                _ = v44
                _ = v45
            }
            _ = fi42
            _ = v43
        }
    case 32767:
        if (!false) {
            fmt.Println(v13, v13)
            var v46 float64 = 0.0
            _ = v46
        }
    default:
        for fi47 := 0; fi47 < 5; fi47++ {
            v1[""] = 701
            var v48 int16 = -1
            _ = fi47
            _ = v48
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
