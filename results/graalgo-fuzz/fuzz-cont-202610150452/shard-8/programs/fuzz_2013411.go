package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint) uint {
    p0 = (p0 / ((435) | 1))
    switch 4294967295 {
    case 32768:
        var v4 uint64 = 310
        _ = v4
    case 60:
        var v5 uint16 = 901
        _ = v5
    default:
        v6 := []int{}
        _ = v6
    }
    var v7 uint16 = 1
    _ = p0
    _ = v7
    return 86
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]string)
    v3 := []string{}
    switch "foo bar" {
    case "":
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(v1, v1, fi9)
                v10 := GenS1{F0: 0, F1: -2147483648}
                var v11 uint64 = 2
                v12 := GenS2{F0: 71}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi8
        }
    case "foo bar":
        switch ("hello" + (("foo bar" + ("hello" + "hello")) + "a")) {
        case "":
            if (65535 != 2147483648) {
                fmt.Println(v1)
            } else {
                switch 2147483647 {
                case -502:
                    _ = len(v2)
                }
                var v13 int = 0
                var v14 int32 = 0
                _ = v13
                _ = v14
            }
        }
    }
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        fmt.Println(v1)
    case "":
        fmt.Println(v1)
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v15 float32 = 3.14159
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 int64 = 124
        fmt.Println(fi16)
        for fi18 := 0; fi18 < 2; fi18++ {
            v17 = v17
            fmt.Println(v17)
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v1 = v1
        _ = fi19
    }
    fmt.Println(v1)
    var v20 int8 = -30
    var v21 int16 = 32767
    var v22 int = 0
    for fi23 := 0; fi23 < 20; fi23++ {
        for fi24 := 0; fi24 < 7; fi24++ {
            var v25 uint32 = 4294967295
            switch (^(v22)) {
            case 128:
                if false {
                    _ = len(v3)
                } else {
                    v21 = v21
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v26 float32 = ((1.5 * (3.14159 / (-2.25 - 0.0))) / ((1.5 * -2.25) * 3.14159))
                    _ = v26
                }
            default:
                var v27 string = ("" + "\n\t")
                _ = v27
            }
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    v21 = v21
    fmt.Println(v22, v22, v20)
    var v28 int = 0
    if true {
        v29 := v2[1]
        switch v29 {
        case "a":
            for fi30 := 0; fi30 < 11; fi30++ {
                v22 = 9223372036854775807
                var v31 int8 = (-(v20))
                fmt.Println(v28)
                _ = fi30
                _ = v31
            }
        case "":
            v20 += v20
        }
        _ = v29
    } else {
        fmt.Println(v28, v28, v1)
        switch (v28 * (v28 - v28)) {
        case 128:
            if (true && false) {
                if false {
                    fmt.Println(v20, v20)
                    var v32 uint32 = 2
                    v20 = (v20 ^ (v20 << ((255) & 63)))
                    _ = v32
                } else {
                    v33 := GenS3{F0: 2, F1: (v21 >> ((409) & 63))}
                    if true {
                        for fi34 := 0; fi34 < 17; fi34++ {
                            for fi35 := 0; fi35 < 4; fi35++ {
                                _ = len(v2)
                                switch v28 {
                                case 42:
                                    var v36 uint = 732
                                    _ = v36
                                default:
                                    fmt.Println(fi34, fi34)
                                }
                                var v37 uint64 = 572
                                _ = fi35
                                _ = v37
                            }
                            v20 = 100
                            v3 = append(v3, (("foo bar" + "") + ("a" + "a")))
                            _ = fi34
                        }
                    }
                    _ = v33
                }
                v38 := GenS2{F0: 1}
                _ = v38
            } else {
                if (2 == 1) {
                    v21 = 176
                }
                v28 = 32768
                switch (v28 - int(v20)) {
                case 130:
                    if (((!false) && (42 > genFn0(4294967295))) && (0.0 <= 1e10)) {
                        if false {
                            var v39 int64 = -786
                            _ = v39
                        }
                    }
                case -729:
                    var v40 uint32 = 2147483647
                    _ = v40
                case 125:
                    var v41 float64 = (-2.25 / float64(float32(0.1)))
                    _ = v41
                default:
                    fmt.Println(v22, v28)
                }
            }
        case 127:
            var v42 uint64 = 1
            _ = v42
        }
        v1 -= (v1 >> ((4294967295) & 63))
    }
    v28 += (-(v28))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
}
