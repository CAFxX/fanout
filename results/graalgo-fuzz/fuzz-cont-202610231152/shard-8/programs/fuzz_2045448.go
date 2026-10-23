package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint32, p1 int, p2 int64) int {
    v5 := []int{p1, (-(p1))}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 := GenS1{F0: p2}
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: -1}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p1 - -1)
}
func genFn1(p0 uint8, p1 int8, p2 int32) string {
    var v9 uint32 = 65535
    if true {
        var v10 uint16 = 1
        switch ("a" + ("a" + ("hello" + ("" + ("foo bar" + (("foo bar" + "") + "\n\t")))))) {
        case "hello":
            v11 := []int{}
            _ = v11
        default:
            v12 := GenS1{F0: -9223372036854775808}
            _ = v12
        }
        _ = v10
    } else {
        p0 += 25
        var v13 int64 = 1
        _ = v13
    }
    if true {
        p1 = (p1 ^ 42)
        switch "\n\t" {
        case "foo bar":
            var v14 uint32 = 751
            _ = v14
        case "":
            v15 := []int{-1}
            _ = v15
        case "a":
            v9 -= (^(v9))
        default:
            v16 := []int{-87, genFn0((v9 % (((v9 % (((v9 >> ((1) & 63))) | 1))) | 1)), 694, -174), 127}
            _ = v16
        }
    }
    p1 = p1
    switch "a" {
    case "a":
        for fi17 := 0; fi17 < 4; fi17++ {
            p2 += p2
            v18 := GenS1{F0: 9223372036854775807}
            v19 := GenS2{F0: true}
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return ""
}

func main() {
    v1 := make(map[int]bool)
    var v2 int16 = -670
    var v3 int = 32768
    var v4 int8 = 42
    v20 := GenS3{F0: int32(v4), F1: 0, F2: 441}
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            var v23 float32 = float32(0.1)
            _ = fi22
            _ = v23
        }
        if false {
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 int64 = -1
                _ = fi24
                _ = v25
            }
            if true {
                for fi26 := 0; fi26 < 9; fi26++ {
                    v2 = 586
                    v2 = (v2 | -1)
                    _ = fi26
                }
            }
        }
        _ = fi21
    }
    fmt.Println(v2)
    switch "\n\t" {
    case "":
        var v27 float32 = ((float32(0.1) * (3.14159 + float32((-2.25 + ((-2.25 / -2.25) * ((1e10 - float64(0.0)) - -2.25)))))) * 1.5)
        _ = v27
    case "a":
        if ((v2 <= v2) && true) {
            var v28 uint = 0
            for fi29 := 0; fi29 < 17; fi29++ {
                v3 -= v3
                v3 = 127
                _ = fi29
            }
            _ = v28
        }
    default:
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 uint8 = 42
            var v32 int16 = (v2 % (((v2 - v2)) | 1))
            _ = fi30
            _ = v31
            _ = v32
        }
    }
    var v33 int8 = (v4 % ((100) | 1))
    var v34 float64 = -2.25
    var v35 int32 = 945
    v33 = v4
    delete(v1, (v3 * 257))
    var v36 bool = (true && (1 < 1))
    switch v3 {
    case 0:
        v1[443] = v36
    case 1:
        switch -9223372036854775808 {
        case 910:
            if false {
                v36 = v36
                v33 = v4
                v37 := GenS1{F0: -198}
                _ = v37
            }
        case 558:
            v3 = genFn0(810, (v3 + 1), int64(v2))
        default:
            fmt.Println(v3, v36, v4)
        }
    default:
        v38 := GenS4{F0: -562}
        _ = v38
    }
    if (!(!(!v36))) {
        var v39 float64 = v34
        v40 := GenS5{F0: -471, F1: v33}
        _ = v39
        _ = v40
    } else {
        if (746 == 26) {
            v36 = ((!(v4 <= (v33 | v33))) && true)
            switch (v3 << ((0) & 63)) {
            case 1:
                if v36 {
                    for fi41 := 0; fi41 < 1; fi41++ {
                        v42 := []int{v3, (fi41 | (fi41 << ((4294967295) & 63)))}
                        _ = fi41
                        _ = v42
                    }
                } else {
                    if v36 {
                        var v43 uint64 = 249
                        v36 = (v36 || v36)
                        _ = v43
                    }
                }
            case 2147483647:
                fmt.Println(v36, v4)
            }
            if (v36 || false) {
                v33 = -128
                var v44 uint16 = 2
                var v45 int = (^(v3))
                _ = v44
                _ = v45
            } else {
                if true {
                    fmt.Println(v33, v34)
                }
            }
        }
        v46 := GenS6{F0: 42}
        _ = v46
    }
    v35 = -297
    switch ("\n\t" + "\n\t") {
    case "hello":
        if v36 {
            if (203 != 569) {
                var v47 string = "foo bar"
                _ = v47
            } else {
                var v48 uint = 4294967295
                _ = v48
            }
        } else {
            if false {
                var v49 int16 = (v2 >> ((1) & 63))
                _ = v49
            } else {
                fmt.Println(v36, v33, v3)
                var v50 uint64 = 42
                var v51 float64 = v34
                _ = v50
                _ = v51
            }
            switch "" {
            case "a":
                for fi52 := 0; fi52 < 5; fi52++ {
                    if false {
                        if ((-(v35)) == -921) {
                            var v53 int8 = (v4 + int8(v3))
                            _ = v53
                        } else {
                            var v54 int16 = v2
                            switch v3 {
                            case -2147483648:
                                fmt.Println(v54, v33, v2)
                            case 32767:
                                if (v54 < v54) {
                                    for fi55 := 0; fi55 < 8; fi55++ {
                                        var v56 string = "foo bar"
                                        _ = fi55
                                        _ = v56
                                    }
                                    var v57 uint = 2
                                    for fi58 := 0; fi58 < 17; fi58++ {
                                        var v59 uint64 = 440
                                        v3 -= (v3 + (v3 ^ v3))
                                        _ = fi58
                                        _ = v59
                                    }
                                    _ = v57
                                } else {
                                    v60 := GenS4{F0: 1}
                                    _ = v60
                                }
                            case 4294967295:
                                var v61 string = ("foo bar" + "\n\t")
                                _ = v61
                            }
                            _ = v54
                        }
                        v62 := GenS1{F0: -1}
                        v63 := []int{(v3 - -9223372036854775808), v3, fi52}
                        _ = v62
                        _ = v63
                    }
                    switch "foo bar" {
                    case "\n\t":
                        var v64 bool = ("a" != ("foo bar" + ("" + (("" + "hello") + "\n\t"))))
                        _ = v64
                    case "hello":
                        switch (("foo bar" + ("\n\t" + "\n\t")) + "foo bar") {
                        case "hello":
                            v65 := GenS6{F0: 2147483648}
                            _ = v65
                        case "":
                            v3 += v3
                        default:
                            var v66 uint8 = 128
                            _ = v66
                        }
                    case "a":
                        var v67 int8 = v33
                        _ = v67
                    default:
                        fmt.Println(v3)
                    }
                    _ = fi52
                }
            default:
                if (785 >= (v35 + v35)) {
                    var v68 uint16 = 42
                    v69 := GenS1{F0: 1}
                    if (!v36) {
                        var v70 string = ("" + ((("hello" + ("\n\t" + ("" + "hello"))) + "a") + "foo bar"))
                        _ = v70
                    }
                    _ = v68
                    _ = v69
                }
            }
        }
    case "":
        for fi71 := 0; fi71 < 2; fi71++ {
            var v72 int = fi71
            _ = fi71
            _ = v72
        }
    default:
        _ = len(v1)
    }
    var v73 uint16 = 255
    if v36 {
        var v74 uint16 = (^(v73))
        if true {
            var v75 int16 = (v2 + (v2 / ((v2) | 1)))
            for fi76 := 0; fi76 < 2; fi76++ {
                v77 := GenS6{F0: 42}
                v78 := []int{}
                _ = fi76
                _ = v77
                _ = v78
            }
            _ = v75
        }
        var v79 float32 = -2.25
        _ = v74
        _ = v79
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v73
}
