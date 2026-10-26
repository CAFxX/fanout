package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint; F2 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16 }
func genFn0(p0 uint8) int16 {
    var v4 uint32 = 4294967295
    p0 += 255
    if false {
        for fi5 := 0; fi5 < 20; fi5++ {
            var v6 int8 = 100
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: 4294967295}
        v8 := []int{-774, -1, 32768, 889}
        _ = v7
        _ = v8
    }
    fmt.Println(v4, v4)
    _ = p0
    _ = v4
    return int16(p0)
}

func main() {
    v1 := []string{("\n\t" + "\n\t")}
    v2 := []int{127, 0}
    v3 := []uint32{}
    if true {
        for fi9 := 0; fi9 < 8; fi9++ {
            v10 := GenS1{F0: fi9}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 14; fi11++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v12 := GenS2{F0: 32767}
            _ = fi11
            _ = v12
        }
    } else {
        var v13 float32 = ((-2.25 / float32(3.14159)) * float32(0.0))
        var v14 bool = (!false)
        var v15 int = 2147483647
        _ = v13
        _ = v14
        _ = v15
    }
    if true {
        v16 := GenS1{F0: 9223372036854775807}
        var v17 uint32 = 0
        for fi18 := 0; fi18 < 1; fi18++ {
            if true {
                v19 := GenS2{F0: -32768}
                v17 += 341
                _ = v19
            } else {
                v17 -= (v17 >> ((754) & 63))
                v3 = append(v3, (v17 * (v17 + 2)))
            }
            v17 += v17
            for fi20 := 0; fi20 < 8; fi20++ {
                v1 = append(v1, ("" + "\n\t"))
                var v21 int64 = 308
                _ = fi20
                _ = v21
            }
            _ = fi18
        }
        _ = v16
        _ = v17
    } else {
        var v22 int16 = 577
        v23 := GenS1{F0: 4294967295}
        var v24 int16 = 1000
        _ = v22
        _ = v23
        _ = v24
    }
    switch "a" {
    case "":
        var v25 string = "a"
        _ = v25
    case "\n\t":
        var v26 uint32 = 72
        _ = v26
    }
    var v27 int32 = -1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v28 int8 = 0
    v28 += 100
    fmt.Println(v27, v27)
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "a")
    }
    if (v27 >= (^(v27))) {
        switch "hello" {
        case "a":
            v27 = (v27 / (((v27 - (v27 + (v27 + (v27 << ((4294967295) & 63)))))) | 1))
        case "\n\t":
            var v29 float32 = float32(0.1)
            _ = v29
        }
        switch 2147483647 {
        case 533:
            for fi30 := 0; fi30 < 11; fi30++ {
                _ = len(v1)
                v27 += (v27 >> ((309) & 63))
                _ = fi30
            }
        case 128:
            for fi31 := 0; fi31 < 10; fi31++ {
                v32 := GenS1{F0: fi31}
                for fi33 := 0; fi33 < 12; fi33++ {
                    var v34 uint8 = 0
                    if false {
                        if ("hello" < "foo bar") {
                            v35 := GenS1{F0: (fi31 | (fi31 + (fi31 >> ((1) & 63))))}
                            var v36 uint64 = uint64(v34)
                            var v37 int = fi33
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        v38 := GenS3{F0: 32767, F1: 2, F2: (fi33 + fi31)}
                        v28 = -128
                        _ = v38
                    }
                    v39 := GenS1{F0: (fi33 + fi31)}
                    _ = fi33
                    _ = v34
                    _ = v39
                }
                fmt.Println(fi31, v28)
                _ = fi31
                _ = v32
            }
        default:
            var v40 uint8 = 0
            _ = v40
        }
        v41 := GenS4{F0: ("hello" + "\n\t")}
        _ = v41
    } else {
        var v42 float32 = (float32(-2.25) - 1e10)
        _ = v42
    }
    if false {
        for fi43 := 0; fi43 < 14; fi43++ {
            fmt.Println(fi43)
            _ = fi43
        }
    } else {
        if ((48 != 0) && false) {
            _ = len(v3)
            v27 = (v27 * (v27 & 42))
            switch "foo bar" {
            case "a":
                var v44 float64 = (float64((1e10 * float32(0.1))) * 1e10)
                _ = v44
            case "\n\t":
                v45 := GenS3{F0: 988, F1: 255, F2: 490}
                _ = v45
            case "hello":
                for fi46 := 0; fi46 < 12; fi46++ {
                    _ = len(v3)
                    _ = fi46
                }
            default:
                v27 += (^(v27))
            }
        }
        for fi47 := 0; fi47 < 1; fi47++ {
            v27 -= v27
            if (4294967295 <= 901) {
                v48 := GenS5{F0: false}
                if (false || (-2.25 < 1.5)) {
                    for fi49 := 0; fi49 < 20; fi49++ {
                        v28 += 100
                        if len(v1) > 0 {
                            v1[0] = ("hello" + "a")
                        }
                        v28 -= (v28 + (v28 >> ((0) & 63)))
                        _ = fi49
                    }
                    var v50 uint32 = 2
                    _ = v50
                }
                v28 = 1
                _ = v48
            }
            v28 = (v28 / ((0) | 1))
            _ = fi47
        }
    }
    for fi51 := 0; fi51 < 14; fi51++ {
        v1 = append(v1, ("hello" + "hello"))
        if ((0.1 / ((-2.25 * 3.14159) / float32((-2.25 * 0.1)))) < -2.25) {
            fmt.Println(fi51, v28)
            v27 = (v27 >> ((528) & 63))
        }
        _ = fi51
    }
    if true {
        for fi52 := 0; fi52 < 10; fi52++ {
            if (true && true) {
                if (false && false) {
                    switch (fi52 | 127) {
                    case 1:
                        fmt.Println(v28)
                    case 128:
                        if (!(((v27 | 0) != v27) || true)) {
                            var v53 int = fi52
                            for fi54 := 0; fi54 < 9; fi54++ {
                                v55 := GenS2{F0: 1}
                                v53 = fi54
                                v53 += (fi54 | 128)
                                var v56 int = (^(v53))
                                v28 += v28
                                var v57 uint32 = 65535
                                _ = fi54
                                _ = v55
                                _ = v56
                                _ = v57
                            }
                            v58 := GenS6{F0: (^(v53))}
                            _ = v53
                            _ = v58
                        }
                    case 9223372036854775807:
                        v59 := GenS1{F0: fi52}
                        _ = v59
                    default:
                        switch ("\n\t" + "") {
                        case "":
                            v60 := GenS7{F0: 0}
                            _ = v60
                        case "hello":
                            var v61 int32 = v27
                            _ = v61
                        }
                    }
                }
                var v62 uint8 = 127
                if (false || (true || (42 < v28))) {
                    v63 := GenS1{F0: fi52}
                    v27 += v27
                    v27 += (v27 % ((v27) | 1))
                    _ = v63
                } else {
                    var v64 uint = 255
                    v64 = (v64 * (v64 | v64))
                    _ = v64
                }
                _ = v62
            } else {
                if len(v1) > 0 {
                    v1[0] = ("foo bar" + ("foo bar" + ""))
                }
            }
            v28 = v28
            _ = fi52
        }
    } else {
        v27 = int32(v28)
    }
    v27 += (v27 ^ v27)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
