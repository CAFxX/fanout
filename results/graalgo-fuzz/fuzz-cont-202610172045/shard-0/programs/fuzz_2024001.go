package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint; F1 bool; F2 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int64 }

func main() {
    v1 := []uint32{1, 1, 0}
    v2 := []string{"", ("foo bar" + "hello"), (("" + "") + "\n\t")}
    v3 := make(map[uint32]int)
    v4 := []uint32{}
    var v5 int = 127
    var v6 uint32 = 795
    var v7 int8 = 6
    switch v5 {
    case 407:
        switch "hello" {
        case "hello":
            fmt.Println(v7, v7, v5)
        case "foo bar":
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := GenS1{F0: 42, F1: v6}
                var v10 uint16 = 0
                for fi11 := 0; fi11 < 7; fi11++ {
                    var v12 uint32 = 447
                    if true {
                        if (false && (true || false)) {
                            fmt.Println(v12, v10, v6)
                            v6 += (v6 ^ (v12 % ((v6) | 1)))
                            v10 = (v10 % ((v10) | 1))
                        }
                        if (719 >= 30) {
                            var v13 int64 = -1
                            _ = v13
                        }
                        v14 := v3[v6]
                        _ = v14
                    } else {
                        fmt.Println(v12, v6, fi11)
                    }
                    v15 := GenS1{F0: 1, F1: (v6 | v12)}
                    fmt.Println(v6, v7)
                    if true {
                        v10 -= 764
                    }
                    _ = fi11
                    _ = v12
                    _ = v15
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        }
    case 798:
        v16 := v3[v6]
        _ = v16
    case 42:
        if true {
            for fi17 := 0; fi17 < 20; fi17++ {
                v18 := GenS1{F0: -819, F1: 477}
                v19 := GenS1{F0: 71, F1: v6}
                _ = fi17
                _ = v18
                _ = v19
            }
            v20 := GenS1{F0: 0, F1: 65535}
            _ = v20
        } else {
            v5 = 167
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 float64 = 0.1
                v5 += v5
                _ = fi21
                _ = v22
            }
            fmt.Println(v6)
        }
    }
    v23 := GenS1{F0: 32767, F1: 65535}
    fmt.Println(v5)
    var v24 string = "hello"
    fmt.Println(v7, v24, v7)
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        fmt.Println(v5, v24, v24)
        _ = fi25
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v26 := GenS2{F0: 42, F1: false, F2: 42}
    fmt.Println(v7, v7)
    if false {
        switch "a" {
        case "\n\t":
            v24 = v24
        case "hello":
            switch "" {
            case "foo bar":
                fmt.Println(v5, v24)
            case "hello":
                v27 := GenS3{F0: v6}
                _ = v27
            default:
                var v28 string = ((v24 + "foo bar") + "a")
                _ = v28
            }
        case "a":
            v2 = append(v2, ("a" + "a"))
        default:
            var v29 int = 32768
            _ = v29
        }
        if (!(!(65535 != v6))) {
            _ = len(v2)
        } else {
            if (!false) {
                if true {
                    v5 -= (v5 >> ((0) & 63))
                    var v30 float64 = 1e10
                    _ = v30
                } else {
                    v7 += v7
                    var v31 float64 = -2.25
                    _ = v31
                }
            }
            v32 := GenS4{F0: v7, F1: 1}
            delete(v3, (v6 * v6))
            _ = v32
        }
    } else {
        switch (v5 & v5) {
        case 2147483647:
            _ = len(v1)
        case 204:
            for fi33 := 0; fi33 < 11; fi33++ {
                var v34 string = "\n\t"
                if false {
                    var v35 int8 = -128
                    v36 := GenS4{F0: v35, F1: -9223372036854775808}
                    v37 := v3[2147483648]
                    _ = v35
                    _ = v36
                    _ = v37
                } else {
                    var v38 int = (^(v5))
                    _ = v38
                }
                _ = fi33
                _ = v34
            }
        }
        switch v24 {
        case "hello":
            v39 := GenS1{F0: 0, F1: (v6 - (v6 & v6))}
            _ = v39
        default:
            v7 = v7
        }
    }
    if (!true) {
        if (-749 <= 42) {
            _ = len(v1)
            if (!false) {
                fmt.Println(v5, v5)
                for fi40 := 0; fi40 < 17; fi40++ {
                    switch ("hello" + v24) {
                    case "foo bar":
                        fmt.Println(fi40, v7)
                    }
                    _ = fi40
                }
            } else {
                v24 = v24
            }
        } else {
            fmt.Println(v5, v6)
        }
        _ = len(v1)
        var v41 int16 = 32767
        _ = v41
    } else {
        var v42 uint32 = v6
        _ = v42
    }
    switch v24 {
    case "":
        var v43 int8 = v7
        _ = v43
    case "\n\t":
        fmt.Println(v5, v6, v6)
    default:
        if true {
            v6 += 65535
        }
    }
    v7 += -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
