package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int; F1 string; F2 int16 }

func main() {
    var v1 float64 = ((1e10 + 1e10) * 3.14159)
    v2 := make(map[string]int)
    var v3 int16 = -152
    v4 := make(map[int]uint32)
    var v5 float64 = v1
    switch 32768 {
    case 770:
        var v6 uint8 = 255
        _ = v6
    }
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int = fi7
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        fmt.Println(v5, v3)
        if true {
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 int8 = 100
                v1 = (v1 / v1)
                fmt.Println(v1)
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 18; fi12++ {
                switch ("" + "foo bar") {
                case "foo bar":
                    v1 = v1
                }
                fmt.Println(v1, v5, v1)
                _ = fi12
            }
        } else {
            for fi13 := 0; fi13 < 15; fi13++ {
                v3 += 42
                var v14 int = (^(fi13))
                _ = fi13
                _ = v14
            }
            v15 := GenS1{F0: 382}
            for fi16 := 0; fi16 < 18; fi16++ {
                var v17 string = (("a" + "") + ("\n\t" + "\n\t"))
                switch (fi16 << ((65535) & 63)) {
                case 32767:
                    v3 += v3
                default:
                    for fi18 := 0; fi18 < 13; fi18++ {
                        v5 -= (((float64(3.14159) - v5) - v1) + v5)
                        switch v17 {
                        case "a":
                            v19 := GenS1{F0: 1}
                            _ = v19
                        case "hello":
                            if false {
                                v4[(fi16 << ((42) & 63))] = 65535
                            }
                        default:
                            v17 = "hello"
                        }
                        switch (-(fi16)) {
                        case 42:
                            v20 := GenS2{F0: 2147483647, F1: 1}
                            _ = v20
                        default:
                            if false {
                                switch v17 {
                                case "\n\t":
                                    v1 = 1e10
                                case "hello":
                                    if (!true) {
                                        v3 -= -663
                                        v1 = v1
                                        v21 := GenS3{F0: (fi18 % (((fi18 & -2147483648)) | 1))}
                                        fmt.Println(v5)
                                        _ = v21
                                    } else {
                                        var v22 uint = 0
                                        v17 = v17
                                        v2["a"] = fi16
                                        _ = v22
                                    }
                                case "foo bar":
                                    var v23 int64 = 1
                                    _ = v23
                                }
                                fmt.Println(v5, v17, v3)
                            }
                        }
                        _ = fi18
                    }
                }
                _ = fi16
                _ = v17
            }
            _ = v15
        }
    } else {
        fmt.Println(v1, v1, v3)
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 int8 = 100
            v26 := []int{(fi24 << ((2) & 63)), 746, fi24}
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := v2[("\n\t" + "hello")]
            if (-46 != (fi27 & 344)) {
                v5 -= v1
                v29 := []int{9223372036854775807}
                _ = v29
            } else {
                for fi30 := 0; fi30 < 2; fi30++ {
                    v31 := GenS4{F0: -9223372036854775808}
                    fmt.Println(v1, v3)
                    var v32 float32 = (3.14159 / (-2.25 * 1.5))
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
            var v33 uint32 = 2
            _ = fi27
            _ = v28
            _ = v33
        }
    }
    switch ("\n\t" + ("\n\t" + "a")) {
    case "\n\t":
        var v34 int8 = 0
        _ = v34
    default:
        var v35 int64 = 291
        _ = v35
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        var v37 uint16 = 1
        v38 := []int{(^(fi36))}
        switch int(v3) {
        case 32768:
            v37 += (^(v37))
        case 32767:
            var v39 int64 = -1
            _ = v39
        case 1:
            v40 := GenS4{F0: 1}
            _ = v40
        default:
            v41 := GenS5{F0: "hello", F1: 1, F2: 128}
            _ = v41
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    v42 := []int{9223372036854775807}
    for fi43 := 0; fi43 < 17; fi43++ {
        switch 32767 {
        case 933:
            v3 -= int16(fi43)
        }
        _ = fi43
    }
    for fi44 := 0; fi44 < 4; fi44++ {
        if (0 == 503) {
            v1 -= (float64((3.14159 / (1.5 + float32((3.14159 * float64(1.5)))))) - (v5 - (float64(((1.5 + 1e10) - (0.0 + 1.5))) * (v5 / 1.5))))
            var v45 uint16 = 1
            _ = v45
        }
        v46 := GenS3{F0: fi44}
        _ = fi44
        _ = v46
    }
    var v47 int = 418
    fmt.Println(v1, v1, v47)
    if (true && (0 > 100)) {
        for fi48 := 0; fi48 < 16; fi48++ {
            var v49 uint8 = 1
            v50 := GenS2{F0: 397, F1: -1}
            if (!false) {
                for fi51 := 0; fi51 < 2; fi51++ {
                    v52 := GenS4{F0: 1}
                    var v53 int32 = -1
                    _ = fi51
                    _ = v52
                    _ = v53
                }
            }
            _ = fi48
            _ = v49
            _ = v50
        }
        var v54 uint = 932
        _ = v54
    } else {
        for fi55 := 0; fi55 < 6; fi55++ {
            v42 = append(v42, (v47 >> ((42) & 63)))
            for fi56 := 0; fi56 < 10; fi56++ {
                v1 -= float64((3.14159 + ((-2.25 / float32((v1 - v5))) - 0.0)))
                v5 += (v5 - (float64((1.5 / 1.0)) / -2.25))
                if false {
                    switch 42 {
                    case 4294967295:
                        v57 := GenS6{F0: (v47 % (((-(fi55))) | 1)), F1: "a", F2: (^(v3))}
                        _ = v57
                    case -196:
                        v3 -= v3
                    case -1:
                        _ = len(v4)
                    }
                    var v58 int32 = 961
                    _ = v58
                }
                _ = fi56
            }
            v1 -= ((v5 * (v1 - 0.1)) + ((v1 / v5) * v5))
            _ = fi55
        }
        delete(v2, ("a" + "hello"))
        v59 := GenS4{F0: int64(v3)}
        _ = v59
    }
    if false {
        v3 = v3
    }
    if false {
        var v60 uint32 = 42
        v1 = v1
        _ = v60
    }
    _ = len(v4)
    for fi61 := 0; fi61 < 8; fi61++ {
        switch "a" {
        case "a":
            var v62 uint32 = 1
            _ = v62
        case "foo bar":
            var v63 string = "foo bar"
            _ = v63
        default:
            var v64 uint8 = 255
            _ = v64
        }
        _ = fi61
    }
    for fi65 := 0; fi65 < 9; fi65++ {
        v66 := GenS6{F0: 42, F1: "\n\t", F2: v3}
        _ = fi65
        _ = v66
    }
    v67 := GenS3{F0: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v47
    _ = v5
    _ = v67
}
