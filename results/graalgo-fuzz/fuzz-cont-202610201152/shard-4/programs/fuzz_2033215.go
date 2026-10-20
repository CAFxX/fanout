package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 int8; F2 int64 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 bool; F1 int8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string }
func genFn0() string {
    var v5 uint = 65535
    switch "" {
    case "a":
        var v6 uint64 = 6
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            if (42 > 102) {
                var v9 int16 = 1
                if true {
                    v9 = (v9 | 1000)
                }
                _ = v9
            }
            v10 := GenS1{F0: -1, F1: 42}
            if (42 > 65535) {
                switch "" {
                case "\n\t":
                    v11 := []int{}
                    _ = v11
                case "":
                    v12 := []int{fi7}
                    _ = v12
                }
                v13 := []int{fi7, (fi8 & (fi8 & -2147483648))}
                _ = v13
            }
            switch "foo bar" {
            case "hello":
                fmt.Println(v5, fi8, fi7)
            case "a":
                v5 = 255
            case "foo bar":
                switch fi8 {
                case 0:
                    v5 -= (v5 % (((v5 + 65535)) | 1))
                }
            default:
                var v14 bool = true
                v15 := GenS1{F0: 1000, F1: 707}
                v16 := GenS1{F0: 103, F1: -1}
                var v17 uint16 = 1
                _ = v14
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi8
            _ = v10
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            v5 = 2
            var v19 bool = true
            v20 := GenS1{F0: 1000, F1: 240}
            var v21 int8 = -128
            var v22 int8 = (v21 / ((v21) | 1))
            var v23 bool = true
            v24 := GenS2{F0: int64(v21)}
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
            _ = v22
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 8; fi25++ {
            if true {
                v26 := []int{-9223372036854775808, (fi25 << (((v5 ^ 584)) & 63)), fi7}
                if (1 >= -498) {
                    v5 = (v5 * v5)
                    fmt.Println(v5, fi25)
                    if (!false) {
                        var v27 bool = (false || ((!true) && (-721 <= -32768)))
                        _ = v27
                    }
                    v5 = 1
                }
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                _ = v26
            } else {
                var v28 uint64 = 1
                v5 += v5
                v29 := GenS3{F0: 0}
                if true {
                    var v30 float64 = (-2.25 + ((0.1 + -2.25) * -2.25))
                    _ = v30
                } else {
                    if (v5 <= (^(v5))) {
                        v28 = v28
                        switch "" {
                        case "hello":
                            if (!(((-(fi25)) == int(v28)) || (4294967295 < v28))) {
                                if (true && (!true)) {
                                    fmt.Println(fi25, v28)
                                    var v31 float32 = 0.1
                                    switch fi25 {
                                    case 42:
                                        fmt.Println(fi7, fi25)
                                    case 393:
                                        if ((127 >= 2) || false) {
                                            if (1 >= 0) {
                                                var v32 int8 = 100
                                                var v33 string = ("\n\t" + "\n\t")
                                                if true {
                                                    v34 := GenS4{F0: -1}
                                                    if (true && ((2 < 255) && (255 < 42))) {
                                                        v35 := []int{(fi25 + (fi7 / (((fi25 * (-(fi25)))) | 1)))}
                                                        var v36 string = ("\n\t" + v33)
                                                        fmt.Println(v5, v31, v32)
                                                        _ = v35
                                                        _ = v36
                                                    } else {
                                                        var v37 float64 = (((-2.25 / ((3.14159 - -2.25) * ((0.1 / 3.14159) / 0.1))) - -2.25) / 3.14159)
                                                        var v38 int = fi25
                                                        fmt.Println(v32)
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    if true {
                                                        var v39 bool = ((v28 / ((v28) | 1)) > v28)
                                                        _ = v39
                                                    } else {
                                                        v40 := GenS1{F0: -32768, F1: 1}
                                                        var v41 bool = false
                                                        switch "\n\t" {
                                                        case "hello":
                                                            fmt.Println(v5)
                                                        default:
                                                            v28 = (v28 / ((v28) | 1))
                                                        }
                                                        v42 := GenS2{F0: 0}
                                                        fmt.Println(v32)
                                                        _ = v40
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    var v43 uint8 = 128
                                                    var v44 int8 = v32
                                                    _ = v34
                                                    _ = v43
                                                    _ = v44
                                                } else {
                                                    if false {
                                                        var v45 int64 = 1
                                                        if (!false) {
                                                            v46 := []int{}
                                                            switch v33 {
                                                            case "hello":
                                                                v46 = append(v46, (fi7 ^ (fi25 - (fi7 >> (((v5 | 252)) & 63)))))
                                                            }
                                                            if len(v46) > 0 {
                                                                v46[0] = 42
                                                            }
                                                            fmt.Println(v31, fi7)
                                                            v5 -= v5
                                                            if (42 > 917) {
                                                                v33 = (v33 + v33)
                                                            } else {
                                                                if (!false) {
                                                                    v5 = v5
                                                                    fmt.Println(fi25, v45)
                                                                } else {
                                                                    var v47 string = "foo bar"
                                                                    switch (fi7 - fi7) {
                                                                    case -1:
                                                                        v47 = v33
                                                                    case 354:
                                                                        fmt.Println(v32, v47)
                                                                    }
                                                                    if len(v46) > 0 {
                                                                        v46[0] = fi7
                                                                    }
                                                                    _ = v47
                                                                }
                                                            }
                                                            var v48 uint32 = 2147483647
                                                            _ = v46
                                                            _ = v48
                                                        }
                                                        fmt.Println(v32, v32, v5)
                                                        _ = v45
                                                    }
                                                    var v49 string = (("hello" + ("\n\t" + "foo bar")) + (("\n\t" + v33) + "hello"))
                                                    v50 := GenS5{F0: int(v32), F1: (-(v32)), F2: -9223372036854775808}
                                                    _ = v49
                                                    _ = v50
                                                }
                                                _ = v32
                                                _ = v33
                                            }
                                            v31 -= ((float32((((0.1 / -2.25) / (1e10 - 3.14159)) / (3.14159 - (0.1 + 3.14159)))) - float32((1.5 / 1.0))) - ((-2.25 / v31) - 3.14159))
                                        }
                                    default:
                                        var v51 int = (fi25 + fi25)
                                        _ = v51
                                    }
                                    _ = v31
                                }
                                v5 -= v5
                                var v52 int = fi7
                                v53 := []int{}
                                var v54 float32 = 0.1
                                v55 := GenS2{F0: int64(fi25)}
                                v54 += ((v54 - float32(0.0)) - v54)
                                _ = v52
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                        case "foo bar":
                            v56 := []int{(fi25 | fi25), 0, (fi25 >> (((^(v5))) & 63))}
                            _ = v56
                        default:
                            v57 := GenS5{F0: (fi25 / (((^(fi25))) | 1)), F1: 37, F2: -9223372036854775808}
                            _ = v57
                        }
                        var v58 int16 = 242
                        _ = v58
                    } else {
                        switch fi25 {
                        case 418:
                            v5 = 4294967295
                        default:
                            v59 := []int{fi25}
                            _ = len(v59)
                            _ = v59
                        }
                    }
                    if (!(v5 == 2)) {
                        v60 := GenS6{F0: 1, F1: false}
                        v5 += 203
                        var v61 float32 = 1.5
                        _ = v60
                        _ = v61
                    }
                    var v62 float32 = (0.1 - 0.0)
                    _ = v62
                }
                _ = v28
                _ = v29
            }
            _ = fi25
        }
        _ = fi7
    }
    v63 := []int{32767}
    _ = v5
    _ = v63
    return "a"
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    var v3 uint = 975
    var v4 uint8 = 0
    switch "\n\t" {
    case "\n\t":
        for fi64 := 0; fi64 < 12; fi64++ {
            v65 := v1[int(v4)]
            switch (fi64 & 1) {
            case -1:
                v3 = (v3 << ((v3) & 63))
            case -462:
                var v66 uint = v3
                _ = v66
            case -9223372036854775808:
                v67 := v1[fi64]
                _ = v67
            }
            _ = fi64
            _ = v65
        }
    }
    v68 := GenS6{F0: -1, F1: false}
    for fi69 := 0; fi69 < 10; fi69++ {
        _ = len(v1)
        switch (^(fi69)) {
        case -792:
            if len(v2) > 0 {
                v2[0] = 0
            }
        }
        _ = fi69
    }
    var v70 int8 = 0
    var v71 int8 = (^(v70))
    _ = len(v2)
    var v72 uint16 = 65535
    v3 = (v3 | (v3 - uint(v72)))
    v73 := GenS1{F0: -1, F1: -733}
    switch 0 {
    case 127:
        delete(v1, -2147483648)
    case 32768:
        v72 -= 920
    case 0:
        switch 0 {
        case -429:
            v74 := v1[128]
            _ = v74
        case 128:
            for fi75 := 0; fi75 < 2; fi75++ {
                v70 = v71
                var v76 uint64 = 18446744073709551615
                _ = len(v2)
                _ = fi75
                _ = v76
            }
        case -9223372036854775808:
            if true {
                fmt.Println(v70, v71, v70)
                fmt.Println(v70, v72)
                v77 := GenS7{F0: false, F1: (v71 | v71)}
                _ = v77
            } else {
                v3 = v3
                var v78 uint32 = 2147483648
                var v79 float32 = 0.1
                _ = v78
                _ = v79
            }
        }
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    for fi80 := 0; fi80 < 20; fi80++ {
        var v81 int64 = -721
        var v82 uint8 = (v4 >> ((v3) & 63))
        v83 := GenS8{F0: "foo bar"}
        _ = fi80
        _ = v81
        _ = v82
        _ = v83
    }
    switch -2147483648 {
    case -1:
        var v84 uint32 = 1
        _ = v84
    case 1:
        if len(v2) > 0 {
            v2[0] = 2
        }
    }
    switch "" {
    case "":
        for fi85 := 0; fi85 < 11; fi85++ {
            if (true || true) {
                v4 = 127
                switch "foo bar" {
                case "foo bar":
                    v4 = 255
                default:
                    _ = len(v2)
                }
                var v86 int32 = int32(v4)
                _ = v86
            } else {
                v70 = (v70 << (((v3 >> ((v3) & 63))) & 63))
            }
            _ = fi85
        }
    case "a":
        var v87 uint16 = (v72 >> (((v3 / ((v3) | 1))) & 63))
        _ = v87
    default:
        v88 := GenS3{F0: -811}
        _ = v88
    }
    var v89 uint8 = v4
    v90 := GenS1{F0: 1, F1: -289}
    if false {
        v1[84] = false
        for fi91 := 0; fi91 < 19; fi91++ {
            _ = len(v1)
            _ = fi91
        }
        if true {
            v72 = v72
        }
    } else {
        if false {
            v92 := GenS9{F0: "hello"}
            _ = v92
        }
        v4 = v4
        var v93 uint32 = 42
        _ = v93
    }
    fmt.Println(v71, v72, v89)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v68
    _ = v70
    _ = v71
    _ = v72
    _ = v73
    _ = v89
    _ = v90
}
