package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint16; F2 int64 }
type GenS4 struct { F0 int16; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 int16 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]int)
    var v3 float64 = 3.14159
    var v4 uint64 = 181
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 16; fi7++ {
            var v8 int16 = 858
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 1, F1: 462, F2: fi6}
        _ = fi6
        _ = v9
    }
    fmt.Println(v4)
    for fi10 := 0; fi10 < 17; fi10++ {
        switch "foo bar" {
        case "hello":
            fmt.Println(v4, fi10, v4)
        case "":
            var v11 uint = 1
            _ = v11
        case "foo bar":
            v12 := GenS2{F0: 0}
            _ = v12
        }
        v13 := GenS3{F0: -1, F1: 992, F2: -1}
        _ = fi10
        _ = v13
    }
    v4 = (v4 | 4294967295)
    switch 901 {
    case 1:
        v14 := GenS4{F0: 1, F1: 410, F2: v1}
        _ = v14
    case 179:
        v15 := []int{127, 298, 128}
        _ = v15
    case -2147483648:
        v3 += float64((3.14159 * float32(v3)))
    default:
        fmt.Println(v1)
    }
    v16 := GenS5{F0: 224}
    v17 := GenS6{F0: 0}
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 int64 = 1
        _ = fi18
        _ = v19
    }
    var v20 uint64 = v4
    switch "a" {
    case "\n\t":
        switch -9223372036854775808 {
        case 128:
            var v21 float64 = v3
            _ = v21
        case -881:
            if (false && ((!false) && false)) {
                var v22 uint16 = 0
                _ = v22
            } else {
                var v23 bool = (1 > -88)
                v24 := v2[("" + ("hello" + "foo bar"))]
                v3 -= v3
                _ = v23
                _ = v24
            }
        }
    case "foo bar":
        for fi25 := 0; fi25 < 14; fi25++ {
            for fi26 := 0; fi26 < 14; fi26++ {
                if false {
                    v20 = v4
                    _ = len(v5)
                }
                var v27 uint16 = 32767
                v28 := []int{(fi26 & fi25), (fi25 - fi26), fi26, (fi26 + -648)}
                v4 = v20
                _ = fi26
                _ = v27
                _ = v28
            }
            var v29 uint = 255
            v4 = 120
            _ = fi25
            _ = v29
        }
    case "a":
        for fi30 := 0; fi30 < 15; fi30++ {
            switch "foo bar" {
            case "foo bar":
                var v31 uint16 = 0
                _ = v31
            case "\n\t":
                if (true && (13 > 65535)) {
                    v1 = (^(v1))
                    if (true || (!(0 == 1))) {
                        v1 += v1
                        var v32 int32 = -959
                        _ = v32
                    } else {
                        var v33 float64 = (v3 - 1e10)
                        _ = v33
                    }
                } else {
                    if (!((^(fi30)) == (^(fi30)))) {
                        var v34 string = ("\n\t" + ("" + "foo bar"))
                        v35 := v2[v34]
                        _ = v34
                        _ = v35
                    } else {
                        var v36 int = fi30
                        v20 = 4294967295
                        if (!false) {
                            v37 := []int{-326}
                            var v38 int64 = 31
                            _ = v37
                            _ = v38
                        } else {
                            for fi39 := 0; fi39 < 5; fi39++ {
                                v36 -= fi39
                                v40 := []int{fi30, (fi39 ^ v36), fi39}
                                var v41 uint64 = v20
                                _ = fi39
                                _ = v40
                                _ = v41
                            }
                            v4 -= v20
                        }
                        _ = v36
                    }
                }
            case "hello":
                var v42 int8 = 100
                _ = v42
            default:
                for fi43 := 0; fi43 < 14; fi43++ {
                    var v44 float32 = (float32(float64(1e10)) * (1e10 + (-2.25 * -2.25)))
                    _ = fi43
                    _ = v44
                }
            }
            _ = fi30
        }
    default:
        var v45 int = 2147483647
        _ = v45
    }
    v3 = ((v3 - v3) * v3)
    v46 := GenS7{F0: (v20 / ((853) | 1)), F1: v1}
    v47 := GenS7{F0: (v20 >> ((42) & 63)), F1: (v1 | (v1 << ((1) & 63)))}
    v1 -= (v1 << ((42) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
