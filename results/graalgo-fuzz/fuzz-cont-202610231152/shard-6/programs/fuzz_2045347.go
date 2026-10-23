package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int64 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 uint64; F1 string; F2 uint8 }
type GenS5 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 uint32; F1 int32; F2 int16 }
type GenS8 struct { F0 uint16 }
func genFn0(p0 bool) int32 {
    for fi6 := 0; fi6 < 14; fi6++ {
        switch fi6 {
        case -459:
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 uint8 = 255
                v9 := []int{}
                _ = fi7
                _ = v8
                _ = v9
            }
        case -644:
            p0 = (689 == -807)
        default:
            p0 = false
        }
        _ = fi6
    }
    p0 = p0
    var v10 int64 = 9223372036854775807
    _ = p0
    _ = v10
    return -225
}
func genFn1(p0 uint32, p1 int64) uint8 {
    v11 := GenS1{F0: 32767, F1: 2147483647, F2: (p1 | p1)}
    v12 := GenS2{F0: 1, F1: -1}
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(p0, p0, fi13)
        p1 += p1
        v14 := []int{fi13}
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 91
}

func main() {
    var v1 uint16 = 1
    var v2 uint = 18446744073709551615
    var v3 uint8 = 0
    var v4 string = (("\n\t" + "") + ("foo bar" + ("a" + ("\n\t" + ""))))
    var v5 int64 = -9223372036854775808
    v3 += v3
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := []int{fi15, int(v1), 1, 32768}
        if (!true) {
            fmt.Println(v5)
            v17 := GenS3{F0: v2, F1: 1}
            _ = v17
        } else {
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS1{F0: fi18, F1: 442, F2: v5}
                if false {
                    var v20 uint = v2
                    v21 := GenS4{F0: 873, F1: "\n\t", F2: 89}
                    v22 := GenS5{F0: 0, F1: v1, F2: (^(v5))}
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    v3 += (v3 | 127)
                    var v23 int = (-(fi15))
                    var v24 int8 = 0
                    _ = v23
                    _ = v24
                }
                var v25 uint64 = 2
                v26 := GenS5{F0: 832, F1: v1, F2: (v5 ^ v5)}
                switch "" {
                case "a":
                    switch int(v1) {
                    case -2147483648:
                        v25 = v25
                    case -653:
                        if len(v16) > 0 {
                            v16[0] = fi18
                        }
                    }
                case "hello":
                    var v27 int32 = -872
                    _ = v27
                case "":
                    switch fi18 {
                    case 59:
                        switch "\n\t" {
                        case "foo bar":
                            if true {
                                var v28 string = v4
                                v5 = -629
                                _ = v28
                            }
                        case "\n\t":
                            var v29 int16 = 1000
                            _ = v29
                        case "a":
                            if ((!true) || ((2147483647 <= 42) && (true || false))) {
                                fmt.Println(v1, fi18, v25)
                                v16 = append(v16, fi15)
                                v25 += (v25 / ((v25) | 1))
                            }
                        }
                    case 955:
                        var v30 int = fi18
                        _ = v30
                    default:
                        v31 := GenS6{F0: -275, F1: v5}
                        switch (v4 + "foo bar") {
                        case "foo bar":
                            v1 = v1
                        case "hello":
                            var v32 int16 = 1000
                            _ = v32
                        }
                        _ = v31
                    }
                }
                if false {
                    v33 := GenS6{F0: 420, F1: v5}
                    _ = v33
                }
                _ = fi18
                _ = v19
                _ = v25
                _ = v26
            }
            var v34 string = ""
            for fi35 := 0; fi35 < 8; fi35++ {
                var v36 float64 = ((3.14159 * -2.25) + (0.0 / 1.0))
                if ((true || (4294967295 > 4294967295)) || (v4 <= "")) {
                    if false {
                        var v37 int32 = 103
                        _ = v37
                    }
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    v2 = (v2 - (v2 / (((v2 % ((v2) | 1))) | 1)))
                } else {
                    v38 := GenS1{F0: 32767, F1: 398, F2: -198}
                    fmt.Println(v5, v36)
                    _ = v38
                }
                _ = fi35
                _ = v36
            }
            _ = v34
        }
        switch "\n\t" {
        case "hello":
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v5, v5, v3)
    var v39 bool = true
    var v40 int8 = 1
    var v41 int8 = v40
    var v42 int16 = 1
    var v43 int32 = 0
    var v44 float64 = 0.1
    fmt.Println(v2)
    if (!v39) {
        v45 := []int{-2147483648, 9223372036854775807, 0}
        _ = v45
    }
    var v46 uint8 = v3
    var v47 int32 = 0
    fmt.Println(v47, v42)
    for fi48 := 0; fi48 < 19; fi48++ {
        v49 := []int{fi48, fi48, fi48}
        for fi50 := 0; fi50 < 5; fi50++ {
            switch "hello" {
            case "\n\t":
                v4 = v4
            default:
                var v51 float32 = -2.25
                _ = v51
            }
            v52 := GenS7{F0: 740, F1: int32(v3), F2: (v42 / (((v42 | v42)) | 1))}
            if true {
                if len(v49) > 0 {
                    v49[0] = 1
                }
                v46 -= uint8(v1)
                fmt.Println(v2, v1, v42)
                _ = len(v49)
                var v53 int64 = (v5 * v5)
                if true {
                    v54 := GenS3{F0: v2, F1: v2}
                    if ((v47 + v47) < (^(v47))) {
                        v55 := GenS3{F0: v2, F1: v2}
                        v56 := GenS4{F0: 1, F1: "a", F2: v46}
                        _ = v55
                        _ = v56
                    } else {
                        var v57 float32 = ((((0.1 + -2.25) - 1.5) + ((0.1 - 0.1) * 1.5)) * 0.1)
                        _ = v57
                    }
                    _ = v54
                } else {
                    _ = len(v49)
                    v53 -= (v53 >> ((4294967295) & 63))
                    var v58 uint32 = 1
                    _ = v58
                }
                _ = v53
            } else {
                fmt.Println(v5, fi50, v47)
            }
            _ = fi50
            _ = v52
        }
        _ = fi48
        _ = v49
    }
    if v39 {
        var v59 uint8 = (v46 >> ((344) & 63))
        _ = v59
    }
    switch 4294967295 {
    case 128:
        v2 = (v2 >> (((v2 - v2)) & 63))
    case 42:
        fmt.Println(v3, v5)
    case 32767:
        var v60 uint = v2
        _ = v60
    }
    v61 := GenS8{F0: 245}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v46
    _ = v47
    _ = v5
    _ = v61
}
