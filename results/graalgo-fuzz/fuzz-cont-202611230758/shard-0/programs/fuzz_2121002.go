package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 int64; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64; F1 int64; F2 uint }
func genFn0(p0 uint64, p1 int16, p2 uint16) int16 {
    v5 := []int{32768, -733, -614}
    fmt.Println(p2, p1)
    if false {
        fmt.Println(p2, p1)
        p2 = (p2 - 1)
        fmt.Println(p2, p0, p2)
    }
    if len(v5) > 0 {
        v5[0] = 141
    }
    var v6 uint8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (p1 >> ((uint(p2)) & 63))
}

func main() {
    var v1 int32 = -1
    var v2 uint8 = 255
    var v3 uint8 = 1
    var v4 float32 = 3.14159
    switch "hello" {
    case "":
        v7 := GenS1{F0: 0}
        _ = v7
    case "a":
        v3 += 1
    case "foo bar":
        switch 9223372036854775807 {
        case 952:
            for fi8 := 0; fi8 < 5; fi8++ {
                for fi9 := 0; fi9 < 18; fi9++ {
                    v10 := GenS2{F0: 0}
                    _ = fi9
                    _ = v10
                }
                v11 := []int{(fi8 & fi8), 9223372036854775807}
                switch (fi8 << ((65535) & 63)) {
                case 0:
                    fmt.Println(v2)
                }
                _ = fi8
                _ = v11
            }
        default:
            if (231 > 1) {
                v12 := GenS2{F0: 1}
                if (42 < 2) {
                    v4 = v4
                } else {
                    fmt.Println(v3, v4, v4)
                    v4 += (float32((3.14159 - (((float64((((1.5 * v4) + v4) - v4)) / 3.14159) + 0.0) / ((-2.25 / 1.5) / float64(0.1))))) - float32(((((((-2.25 * -2.25) * 1e10) - 1.5) - 0.1) / float64(v4)) * 0.1)))
                }
                for fi13 := 0; fi13 < 15; fi13++ {
                    v14 := GenS2{F0: 42}
                    _ = fi13
                    _ = v14
                }
                _ = v12
            } else {
                v3 = v3
            }
        }
    }
    v15 := GenS1{F0: 111}
    if false {
        v4 -= 0.0
    } else {
        if (!(65535 < 2)) {
            switch -9223372036854775808 {
            case -9223372036854775808:
                if false {
                    fmt.Println(v3, v3)
                    var v16 uint = 18
                    _ = v16
                } else {
                    v17 := []int{}
                    for fi18 := 0; fi18 < 5; fi18++ {
                        var v19 int16 = 1000
                        var v20 int64 = -883
                        v2 += v3
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    if false {
                        v2 -= 211
                        if len(v17) > 0 {
                            v17[0] = -357
                        }
                    }
                    _ = v17
                }
            case -33:
                v4 += 3.14159
            case 9223372036854775807:
                if false {
                    for fi21 := 0; fi21 < 6; fi21++ {
                        var v22 string = ""
                        var v23 int8 = 127
                        var v24 int = fi21
                        _ = fi21
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                }
            default:
                var v25 uint16 = 65535
                _ = v25
            }
            switch 9223372036854775807 {
            case -9223372036854775808:
                if (!true) {
                    for fi26 := 0; fi26 < 9; fi26++ {
                        v27 := GenS3{F0: 32767}
                        v28 := GenS4{F0: ((("a" + "\n\t") + "foo bar") + (("\n\t" + "foo bar") + ((("a" + ("" + "a")) + "\n\t") + ("hello" + "")))), F1: 9223372036854775807, F2: v1}
                        var v29 uint8 = 197
                        _ = fi26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    fmt.Println(v3, v1)
                    v4 -= float32(0.1)
                }
            case 0:
                v3 = v2
            case 1:
                v30 := GenS3{F0: -579}
                _ = v30
            }
        }
        switch 0 {
        case 32768:
            var v31 string = ("foo bar" + "")
            _ = v31
        case -229:
            for fi32 := 0; fi32 < 4; fi32++ {
                if (false && true) {
                    v33 := GenS1{F0: 19}
                    _ = v33
                } else {
                    var v34 uint64 = 4294967295
                    v35 := GenS2{F0: -32768}
                    for fi36 := 0; fi36 < 14; fi36++ {
                        var v37 float64 = (3.14159 / 3.14159)
                        _ = fi36
                        _ = v37
                    }
                    _ = v34
                    _ = v35
                }
                if (true || ((-2.25 / 1.5) == (1e10 / 1e10))) {
                    if false {
                        var v38 float64 = 1.5
                        for fi39 := 0; fi39 < 15; fi39++ {
                            var v40 uint64 = 0
                            v41 := []int{}
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                        _ = v38
                    }
                    var v42 int64 = -1
                    _ = v42
                } else {
                    if false {
                        v2 -= 128
                        switch (fi32 * fi32) {
                        case 127:
                            if ((1.5 / float32(((-2.25 + -2.25) + 0.1))) > (v4 * (float32((0.0 * -2.25)) - (((3.14159 / (-2.25 - v4)) / (v4 + (v4 - 1.5))) * v4)))) {
                                v43 := GenS4{F0: ("\n\t" + "a"), F1: -1, F2: (v1 | v1)}
                                v1 += (v1 % ((v1) | 1))
                                _ = v43
                            } else {
                                for fi44 := 0; fi44 < 15; fi44++ {
                                    v3 -= 200
                                    _ = fi44
                                }
                            }
                        default:
                            v2 = v2
                        }
                    }
                    if (!true) {
                        v1 -= -2147483648
                        var v45 int = fi32
                        _ = v45
                    } else {
                        v46 := []int{(fi32 & fi32), (fi32 % ((fi32) | 1))}
                        var v47 uint32 = 40
                        fmt.Println(v4, fi32)
                        _ = v46
                        _ = v47
                    }
                }
                _ = fi32
            }
        case 4294967295:
            fmt.Println(v2, v1, v2)
        }
        var v48 float32 = v4
        _ = v48
    }
    if (-64 > 0) {
        fmt.Println(v1)
    }
    v49 := GenS3{F0: -1}
    if true {
        for fi50 := 0; fi50 < 10; fi50++ {
            if false {
                var v51 string = "foo bar"
                v51 = v51
                var v52 bool = (85 >= 0)
                _ = v51
                _ = v52
            } else {
                v1 += v1
                v4 -= (((v4 - v4) + v4) / v4)
            }
            _ = fi50
        }
        v1 += v1
        v53 := GenS5{F0: ("a" + "\n\t")}
        _ = v53
    }
    fmt.Println(v3)
    fmt.Println(v4, v2, v4)
    fmt.Println(v2, v3, v2)
    if false {
        var v54 int = 658
        for fi55 := 0; fi55 < 17; fi55++ {
            v56 := GenS3{F0: -652}
            _ = fi55
            _ = v56
        }
        v57 := GenS6{F0: 9223372036854775807, F1: 1, F2: 4294967295}
        _ = v54
        _ = v57
    } else {
        v3 = v3
    }
    for fi58 := 0; fi58 < 8; fi58++ {
        fmt.Println(v2, v3, v1)
        _ = fi58
    }
    var v59 int8 = 100
    v60 := []int{127}
    var v61 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v60:", len(v60), v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v59
    _ = v60
    _ = v61
}
