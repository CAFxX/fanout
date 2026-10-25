package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 uint = 958
    var v4 uint8 = 2
    var v5 uint16 = 65535
    v6 := make(map[int]int)
    v4 -= v4
    v7 := GenS1{F0: 23}
    fmt.Println(v3, v3, v4)
    var v8 int16 = 1
    v8 = v8
    v6[-2147483648] = 658
    if true {
        fmt.Println(v4, v5, v5)
    } else {
        v9 := GenS1{F0: 153}
        var v10 string = ("hello" + "hello")
        delete(v1, 659)
        _ = v10
        _ = v9
    }
    if false {
        if true {
            if false {
                v11 := GenS1{F0: 0}
                _ = v11
            } else {
                var v12 float32 = 1.5
                var v13 uint8 = 0
                _ = v12
                _ = v13
            }
            var v14 uint32 = 2
            _ = v14
        }
        if ((v8 ^ 32767) <= 42) {
            for fi15 := 0; fi15 < 17; fi15++ {
                v4 = (v4 * v4)
                _ = fi15
            }
        } else {
            v16 := GenS2{F0: 4294967295}
            switch "hello" {
            case "foo bar":
                if false {
                    v17 := GenS3{F0: 4294967295}
                    _ = v17
                } else {
                    switch 0 {
                    case 32768:
                        v5 = (^(v5))
                    }
                }
            case "hello":
                for fi18 := 0; fi18 < 5; fi18++ {
                    v19 := v2[(fi18 << (((v3 | uint(v5))) & 63))]
                    var v20 int16 = (^(v8))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            default:
                switch 32768 {
                case 0:
                    v21 := GenS3{F0: 0}
                    _ = v21
                case 32767:
                    switch 4294967295 {
                    case -1:
                        fmt.Println(v8, v4)
                    case 1:
                        for fi22 := 0; fi22 < 11; fi22++ {
                            if false {
                                switch (-(fi22)) {
                                case 128:
                                    var v23 int64 = 1
                                    _ = v23
                                }
                            }
                            for fi24 := 0; fi24 < 20; fi24++ {
                                v25 := []int{fi22}
                                if true {
                                    switch 1 {
                                    case 2147483647:
                                        v26 := GenS4{F0: (v3 << (((v3 & (^(v3)))) & 63))}
                                        _ = v26
                                    }
                                    if false {
                                        switch "\n\t" {
                                        case "a":
                                            switch "\n\t" {
                                            case "hello":
                                                if true {
                                                    var v27 int = fi22
                                                    v8 += v8
                                                    if true {
                                                        fmt.Println(v8, v8)
                                                        _ = len(v6)
                                                        var v28 uint8 = (^(v4))
                                                        _ = v28
                                                    }
                                                    _ = v27
                                                } else {
                                                    delete(v1, 4294967295)
                                                    switch "a" {
                                                    case "":
                                                        switch "hello" {
                                                        case "":
                                                            if ((false || false) && (0 < v8)) {
                                                                var v29 uint = (v3 / (((v3 & v3)) | 1))
                                                                if (!false) {
                                                                    v30 := GenS1{F0: 102}
                                                                    _ = v30
                                                                }
                                                                v31 := v2[-1]
                                                                _ = v29
                                                                _ = v31
                                                            }
                                                        case "a":
                                                            if false {
                                                                if (101 != 42) {
                                                                    v3 = v3
                                                                    if len(v25) >= 2 {
                                                                        v25 = v25[1:len(v25)]
                                                                    }
                                                                }
                                                            }
                                                        case "\n\t":
                                                            var v32 uint = v3
                                                            _ = v32
                                                        default:
                                                            var v33 int8 = -1
                                                            _ = v33
                                                        }
                                                    case "hello":
                                                        switch ("foo bar" + "") {
                                                        case "hello":
                                                            if (2 <= 2) {
                                                                v5 += v5
                                                            }
                                                        }
                                                    }
                                                    _ = len(v2)
                                                    var v34 int64 = 0
                                                    _ = v34
                                                }
                                            case "a":
                                                var v35 uint64 = 42
                                                _ = v35
                                            default:
                                                var v36 bool = ((false && (32767 > (v5 - 255))) || true)
                                                _ = v36
                                            }
                                        default:
                                            v4 -= v4
                                        }
                                        var v37 int16 = v8
                                        var v38 float64 = 1.5
                                        var v39 float64 = (v38 / 1.0)
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    } else {
                                        var v40 uint = v3
                                        fmt.Println(v5)
                                        _ = v40
                                    }
                                    _ = len(v25)
                                } else {
                                    if ((v8 != (^(v8))) && (1 < 1)) {
                                        switch fi24 {
                                        case 127:
                                            v41 := v6[(-(fi22))]
                                            _ = v41
                                        default:
                                            v42 := GenS3{F0: 465}
                                            _ = v42
                                        }
                                        v5 += 195
                                        switch "hello" {
                                        case "hello":
                                            if false {
                                                switch 9223372036854775807 {
                                                case -9223372036854775808:
                                                    if ((^(v3)) < 255) {
                                                        v43 := GenS5{F0: fi24}
                                                        if true {
                                                            v5 -= 498
                                                            var v44 int = fi24
                                                            _ = v44
                                                        }
                                                        _ = v43
                                                    }
                                                case -620:
                                                    v45 := GenS2{F0: 104}
                                                    v46 := GenS6{F0: 2147483647, F1: (^(v4))}
                                                    var v47 float32 = -2.25
                                                    _ = v45
                                                    _ = v46
                                                    _ = v47
                                                case 4294967295:
                                                    v48 := GenS6{F0: 1, F1: 1}
                                                    _ = v48
                                                default:
                                                    if (!true) {
                                                        fmt.Println(fi24, v8)
                                                        _ = len(v6)
                                                    } else {
                                                        fmt.Println(v8)
                                                        switch fi24 {
                                                        case 0:
                                                            v5 += v5
                                                        case 488:
                                                            v5 -= 945
                                                        case 1:
                                                            fmt.Println(v5, v3, fi24)
                                                        default:
                                                            switch fi22 {
                                                            case 1:
                                                                switch fi24 {
                                                                case 1:
                                                                    v5 = (v5 >> ((v3) & 63))
                                                                }
                                                            case 4294967295:
                                                                v49 := GenS4{F0: (v3 & 4294967295)}
                                                                _ = v49
                                                            case 42:
                                                                switch (fi24 << (((v3 * v3)) & 63)) {
                                                                case -41:
                                                                    v50 := GenS5{F0: (^(fi22))}
                                                                    if false {
                                                                        delete(v2, fi24)
                                                                    }
                                                                    _ = v50
                                                                case 42:
                                                                    var v51 uint32 = 42
                                                                    _ = v51
                                                                case 852:
                                                                    if false {
                                                                        fmt.Println(fi24, fi22)
                                                                        v8 = (v8 - (-(v8)))
                                                                        if len(v25) >= 2 {
                                                                            v25 = v25[1:len(v25)]
                                                                        }
                                                                    }
                                                                default:
                                                                    fmt.Println(v8, v8)
                                                                }
                                                            default:
                                                                v8 = (v8 % (((v8 * v8)) | 1))
                                                            }
                                                        }
                                                    }
                                                }
                                                if (3.14159 == 0.1) {
                                                    v52 := GenS1{F0: 2}
                                                    _ = v52
                                                } else {
                                                    switch -259 {
                                                    case -346:
                                                        var v53 int16 = (v8 | -32768)
                                                        _ = v53
                                                    }
                                                }
                                                v8 = (v8 & (v8 + 676))
                                            }
                                        default:
                                            if len(v25) >= 2 {
                                                v25 = v25[1:len(v25)]
                                            }
                                        }
                                    } else {
                                        _ = len(v25)
                                        v3 = (v3 - (v3 << (((v3 + 879)) & 63)))
                                    }
                                    var v54 uint32 = 0
                                    switch ("\n\t" + "foo bar") {
                                    case "a":
                                        v5 = 2
                                    case "hello":
                                        fmt.Println(v8, v3, v4)
                                    }
                                    _ = v54
                                }
                                v5 -= 372
                                var v55 float32 = (float32((0.0 * 1.5)) / 0.1)
                                _ = fi24
                                _ = v25
                                _ = v55
                            }
                            _ = fi22
                        }
                    }
                default:
                    var v56 float64 = (0.1 / float64(1e10))
                    _ = v56
                }
            }
            _ = v16
        }
        var v57 bool = true
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
