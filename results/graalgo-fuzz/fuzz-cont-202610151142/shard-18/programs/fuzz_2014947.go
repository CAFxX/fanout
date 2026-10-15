package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
type GenS2 struct { F0 string; F1 int8; F2 int8 }
type GenS3 struct { F0 string; F1 int32 }
func genFn0(p0 uint, p1 int, p2 uint8) uint16 {
    var v7 float64 = 0.0
    fmt.Println(v7)
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := []int{}
        if (!(!false)) {
            if true {
                fmt.Println(p0, v7, p1)
                v10 := GenS1{F0: -9223372036854775808, F1: 1, F2: 97}
                _ = v10
            }
            v7 -= (0.0 * 1.5)
            v7 -= ((3.14159 / v7) * float64(0.0))
        }
        p2 = p2
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 943
}

func main() {
    v1 := []int{-1, -1}
    v2 := make(map[string]bool)
    var v3 uint8 = 127
    var v4 uint = 625
    var v5 uint64 = 1
    v6 := []int{697}
    _ = len(v2)
    fmt.Println(v4, v5, v3)
    v4 = (v4 % ((v4) | 1))
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 int16 = -32768
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 int8 = 0
        if false {
            if true {
                if (0 > 128) {
                    switch "hello" {
                    case "":
                        for fi15 := 0; fi15 < 3; fi15++ {
                            v16 := GenS1{F0: -9223372036854775808, F1: 2147483647, F2: v14}
                            _ = fi15
                            _ = v16
                        }
                    default:
                        v5 -= v5
                    }
                    switch fi13 {
                    case -9223372036854775808:
                        v3 -= v3
                    case 976:
                        var v17 string = "foo bar"
                        _ = v17
                    }
                }
                for fi18 := 0; fi18 < 1; fi18++ {
                    v5 -= v5
                    v6 = append(v6, (fi18 ^ fi18))
                    var v19 float32 = -2.25
                    var v20 string = ("a" + "\n\t")
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            } else {
                v6 = append(v6, 0)
                var v21 int16 = -1
                var v22 uint64 = uint64(v3)
                _ = v21
                _ = v22
            }
        }
        if false {
            var v23 uint64 = (v5 >> (((v4 % (((v4 * 2)) | 1))) & 63))
            _ = v23
        }
        _ = fi13
        _ = v14
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v4 = v4
        _ = fi24
    }
    if ((3.14159 - float32(float64(float32(-2.25)))) == 1e10) {
        _ = len(v2)
        v4 = (v4 << (((v4 * v4)) & 63))
    } else {
        if ("a" < "foo bar") {
            fmt.Println(v3, v4, v4)
            var v25 uint8 = (v3 - v3)
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 int64 = 42
                var v28 string = "hello"
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = v25
        } else {
            var v29 int16 = 153
            v5 = (v5 % ((v5) | 1))
            _ = v29
        }
    }
    if (!false) {
        fmt.Println(v3)
    } else {
        if false {
            v30 := GenS1{F0: -1, F1: -2147483648, F2: 0}
            v31 := GenS2{F0: ("\n\t" + ((("\n\t" + "a") + "hello") + "hello")), F1: -13, F2: -1}
            _ = v30
            _ = v31
        }
        var v32 uint8 = 2
        _ = v32
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        switch 32767 {
        case -519:
            if (!false) {
                v34 := GenS2{F0: (("foo bar" + ("hello" + "")) + "\n\t"), F1: int8(v5), F2: -1}
                fmt.Println(v3, v4, v3)
                for fi35 := 0; fi35 < 11; fi35++ {
                    fmt.Println(v5, v4, v4)
                    v36 := v2[""]
                    _ = fi35
                    _ = v36
                }
                _ = v34
            }
        case 2147483647:
            for fi37 := 0; fi37 < 12; fi37++ {
                v38 := GenS3{F0: "a", F1: 1}
                v3 = (^(v3))
                var v39 string = "\n\t"
                _ = fi37
                _ = v38
                _ = v39
            }
        default:
            for fi40 := 0; fi40 < 17; fi40++ {
                var v41 uint32 = 2147483648
                _ = fi40
                _ = v41
            }
        }
        switch "" {
        case "hello":
            var v42 int = (^(fi33))
            _ = v42
        }
        v4 = v4
        _ = fi33
    }
    var v43 float32 = -2.25
    v43 = 1e10
    var v44 bool = (!((1e10 != 0.1) || false))
    v44 = (42 >= int16(v3))
    switch "hello" {
    case "":
        var v45 int64 = -1
        _ = v45
    }
    if v44 {
        fmt.Println(v5, v3, v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
