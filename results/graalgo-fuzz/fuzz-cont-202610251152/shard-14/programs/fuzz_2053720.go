package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int64 }
func genFn0() uint16 {
    var v6 int = 821
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int32 = 42
        _ = fi7
        _ = v8
    }
    var v9 uint8 = 255
    _ = v6
    _ = v9
    return 35
}
func genFn1() int32 {
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 float64 = 0.0
        if (1e10 < 1e10) {
            if (!(805 <= 9223372036854775807)) {
                var v12 uint32 = 42
                if true {
                    var v13 bool = (!false)
                    v13 = v13
                    _ = v13
                } else {
                    v14 := GenS1{F0: 156}
                    _ = v14
                }
                _ = v12
            }
        } else {
            if ((1 <= 0) && (!false)) {
                var v15 uint64 = 2
                _ = v15
            } else {
                var v16 int16 = 0
                if false {
                    switch -934 {
                    case 128:
                        for fi17 := 0; fi17 < 14; fi17++ {
                            fmt.Println(v11, fi17, v16)
                            v18 := []int{(fi17 >> ((18446744073709551615) & 63)), (fi17 % ((127) | 1)), 128}
                            _ = fi17
                            _ = v18
                        }
                    case -888:
                        if true {
                            v16 -= (v16 * (v16 - v16))
                            v19 := []int{fi10, fi10, fi10}
                            _ = v19
                        }
                    case 1:
                        v20 := GenS2{F0: 42, F1: 1}
                        _ = v20
                    }
                }
                if (!("\n\t" == "")) {
                    var v21 int64 = -9223372036854775808
                    v16 = v16
                    _ = v21
                }
                _ = v16
            }
            if ((false && false) || false) {
                var v22 uint8 = 1
                var v23 uint64 = 377
                for fi24 := 0; fi24 < 3; fi24++ {
                    if false {
                        v25 := GenS3{F0: "foo bar", F1: true}
                        if false {
                            fmt.Println(v23)
                            v23 -= (^(v23))
                        } else {
                            v26 := []int{1, (fi10 - (^(fi24))), fi10, (fi10 >> ((18446744073709551615) & 63))}
                            var v27 float64 = (v11 - v11)
                            _ = v26
                            _ = v27
                        }
                        v28 := []int{}
                        _ = v25
                        _ = v28
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            var v29 uint64 = 841
                            _ = v29
                        default:
                            v30 := []int{(fi24 - fi10)}
                            _ = v30
                        }
                    }
                    _ = fi24
                }
                _ = v22
                _ = v23
            }
            if true {
                v31 := GenS2{F0: 32767, F1: 204}
                v11 += v11
                _ = v31
            }
        }
        _ = fi10
        _ = v11
    }
    v32 := GenS1{F0: 184}
    var v33 uint = 255
    _ = v32
    _ = v33
    return -2147483648
}

func main() {
    var v1 int32 = 1
    var v2 int = -909
    var v3 uint64 = 18446744073709551615
    var v4 uint8 = 0
    v5 := make(map[string]string)
    for fi34 := 0; fi34 < 16; fi34++ {
        for fi35 := 0; fi35 < 1; fi35++ {
            var v36 uint8 = (v4 + (v4 << ((1) & 63)))
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    v2 += v2
    if (false || false) {
        v1 = (v1 | -1)
    }
    var v37 int16 = -1
    switch "foo bar" {
    case "\n\t":
        if (true || (!(-44 < 0))) {
            v1 = (v1 | (v1 * v1))
            fmt.Println(v37)
            _ = len(v5)
        } else {
            fmt.Println(v37)
            switch v2 {
            case -638:
                v38 := []int{(v2 | (-(v2))), 9223372036854775807, v2}
                _ = v38
            default:
                delete(v5, "\n\t")
            }
        }
    case "a":
        v39 := GenS4{F0: 9223372036854775807}
        _ = v39
    case "foo bar":
        var v40 uint = 87
        _ = v40
    }
    switch -2147483648 {
    case 127:
        for fi41 := 0; fi41 < 19; fi41++ {
            switch ("a" + "foo bar") {
            case "\n\t":
                if false {
                    fmt.Println(fi41, v37, fi41)
                } else {
                    v42 := []int{-442, 911, (fi41 * (fi41 >> ((2) & 63)))}
                    if ((fi41 - int(v37)) <= -9223372036854775808) {
                        v43 := GenS3{F0: ("hello" + "hello"), F1: true}
                        _ = v43
                    }
                    _ = v42
                }
            case "hello":
                v37 += (^(v37))
            default:
                _ = len(v5)
            }
            _ = fi41
        }
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        var v45 int = v2
        var v46 int64 = -84
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println(v4, v3, v37)
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
