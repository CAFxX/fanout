package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 uint64; F1 bool }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint16 = 1
    var v3 string = (("foo bar" + "a") + "hello")
    var v4 int = -2147483648
    var v5 int16 = 42
    var v6 int32 = 538
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 float64 = (-2.25 / 1.5)
        if false {
            v9 := GenS1{F0: v2, F1: 1, F2: v6}
            var v10 float32 = (float32(v8) * 1e10)
            _ = v10
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    v1 -= 42
    var v11 int64 = 158
    for fi12 := 0; fi12 < 4; fi12++ {
        if true {
            v6 += v6
        } else {
            v3 = v3
            for fi13 := 0; fi13 < 18; fi13++ {
                v11 = (v11 + (-(v11)))
                if (270 > 0) {
                    v14 := []int{}
                    _ = v14
                } else {
                    var v15 uint8 = 42
                    var v16 uint32 = 862
                    _ = v15
                    _ = v16
                }
                _ = fi13
            }
        }
        switch v3 {
        case "a":
            v17 := []int{int(v1), (^(fi12)), (v4 << ((926) & 63))}
            _ = v17
        case "foo bar":
            fmt.Println(v3, v3, fi12)
        case "hello":
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 uint32 = 146
                v20 := GenS2{F0: v11, F1: 4294967295, F2: 255}
                var v21 float32 = ((float32(((1e10 * (-2.25 / 1.0)) + 3.14159)) * 1.5) * float32(0.0))
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        default:
            fmt.Println(v6, fi12, v11)
        }
        if true {
            if true {
                var v22 bool = false
                if v22 {
                    if true {
                        var v23 int8 = 1
                        fmt.Println(v22, v4, v5)
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 1; fi24++ {
                            var v25 uint64 = v1
                            _ = fi24
                            _ = v25
                        }
                        v26 := GenS3{F0: 66}
                        for fi27 := 0; fi27 < 13; fi27++ {
                            v28 := []int{fi27, fi27, (fi27 << ((65535) & 63)), (v4 | int(v1))}
                            if v22 {
                                if false {
                                    switch "" {
                                    case "hello":
                                        v29 := GenS4{F0: 18446744073709551615, F1: (v11 % ((1) | 1))}
                                        _ = v29
                                    }
                                    fmt.Println(v4, v11, v1)
                                } else {
                                    v28 = append(v28, (fi27 | fi12))
                                    v28 = append(v28, (-(fi27)))
                                    _ = len(v28)
                                }
                            } else {
                                var v30 int8 = 0
                                _ = v30
                            }
                            _ = fi27
                            _ = v28
                        }
                        _ = v26
                    }
                } else {
                    if (-2.25 == 1.5) {
                        for fi31 := 0; fi31 < 5; fi31++ {
                            if ((v1 < 18446744073709551615) || (1 > (v2 % (((v2 | v2)) | 1)))) {
                                fmt.Println(v1, v1)
                                v22 = (!true)
                                var v32 bool = false
                                _ = v32
                            }
                            var v33 uint8 = 2
                            _ = fi31
                            _ = v33
                        }
                        fmt.Println(v4)
                    }
                    v34 := []int{(-(fi12)), v4, 942}
                    _ = v34
                }
                v3 = ("foo bar" + ((v3 + v3) + "\n\t"))
                _ = v22
            } else {
                v11 = v11
                if false {
                    if false {
                        for fi35 := 0; fi35 < 15; fi35++ {
                            v36 := []int{}
                            var v37 uint32 = 42
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                        var v38 int16 = (v5 - (^(v5)))
                        switch "a" {
                        case "hello":
                            for fi39 := 0; fi39 < 7; fi39++ {
                                var v40 int16 = 128
                                _ = fi39
                                _ = v40
                            }
                        case "":
                            v41 := []int{(fi12 >> ((65535) & 63)), (fi12 - fi12), int(v6), 32767}
                            _ = v41
                        case "\n\t":
                            var v42 int8 = 127
                            _ = v42
                        }
                        _ = v38
                    }
                }
            }
            var v43 uint = 4294967295
            _ = v43
        } else {
            switch v3 {
            case "hello":
                switch ("\n\t" + (v3 + "\n\t")) {
                case "\n\t":
                    if true {
                        for fi44 := 0; fi44 < 14; fi44++ {
                            v45 := GenS3{F0: 1}
                            var v46 uint32 = uint32(v1)
                            _ = fi44
                            _ = v45
                            _ = v46
                        }
                        for fi47 := 0; fi47 < 12; fi47++ {
                            switch (fi12 - fi47) {
                            case 128:
                                switch v3 {
                                case "foo bar":
                                    if (1.5 != (1e10 / -2.25)) {
                                        v48 := GenS5{F0: 1, F1: (-(v5))}
                                        _ = v48
                                    }
                                }
                            case 127:
                                fmt.Println(v11)
                            case 9223372036854775807:
                                if false {
                                    var v49 uint32 = 2
                                    _ = v49
                                } else {
                                    switch "" {
                                    case "":
                                        fmt.Println(fi12, v1, v5)
                                    case "a":
                                        switch v3 {
                                        case "foo bar":
                                            v6 = int32(v2)
                                        case "\n\t":
                                            v50 := []int{}
                                            _ = v50
                                        case "hello":
                                            var v51 int16 = v5
                                            _ = v51
                                        }
                                    case "\n\t":
                                        var v52 int16 = int16(v6)
                                        _ = v52
                                    default:
                                        v53 := GenS5{F0: -38, F1: 42}
                                        _ = v53
                                    }
                                }
                            default:
                                v54 := GenS3{F0: 42}
                                _ = v54
                            }
                            _ = fi47
                        }
                        v55 := GenS6{F0: v11, F1: v5}
                        _ = v55
                    } else {
                        fmt.Println(v2, v1)
                        var v56 int32 = v6
                        v57 := GenS7{F0: v1, F1: true}
                        _ = v56
                        _ = v57
                    }
                }
            }
        }
        _ = fi12
    }
    if false {
        v58 := GenS3{F0: 127}
        var v59 float64 = (1e10 + 1.5)
        _ = v58
        _ = v59
    }
    var v60 string = (v3 + v3)
    for fi61 := 0; fi61 < 18; fi61++ {
        v2 += v2
        v60 = "foo bar"
        v62 := []int{-973}
        _ = fi61
        _ = v62
    }
    var v63 float32 = -2.25
    v4 -= v4
    var v64 uint64 = 18446744073709551615
    v1 -= (v1 >> ((42) & 63))
    var v65 uint64 = 0
    v66 := []int{v4}
    v67 := GenS3{F0: 1}
    v68 := GenS6{F0: 9223372036854775807, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v66:", len(v66), v66)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v68
}
