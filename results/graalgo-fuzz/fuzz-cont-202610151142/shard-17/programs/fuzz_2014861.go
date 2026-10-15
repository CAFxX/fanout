package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint32; F1 bool; F2 int8 }
type GenS5 struct { F0 bool }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    var v3 int16 = -32768
    var v4 uint = 4294967295
    switch "hello" {
    case "":
        var v5 int8 = -1
        _ = v5
    case "foo bar":
        v6 := GenS1{F0: 42}
        _ = v6
    }
    var v7 int32 = -391
    var v8 int = 2147483647
    switch "hello" {
    case "foo bar":
        if (false && true) {
            if true {
                switch "" {
                case "hello":
                    v4 = 1
                case "foo bar":
                    v9 := GenS2{F0: (v3 / ((42) | 1))}
                    _ = v9
                default:
                    var v10 float64 = (-2.25 - 1.5)
                    _ = v10
                }
                switch "" {
                case "":
                    for fi11 := 0; fi11 < 5; fi11++ {
                        for fi12 := 0; fi12 < 10; fi12++ {
                            var v13 float64 = 3.14159
                            v8 -= v8
                            _ = fi12
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 9; fi14++ {
                            v3 += v3
                            v15 := GenS1{F0: 791}
                            if false {
                                v16 := []int{32768}
                                v17 := v2[fi11]
                                _ = v16
                                _ = v17
                            }
                            _ = fi14
                            _ = v15
                        }
                        if (69 == 1) {
                            var v18 uint = 18446744073709551615
                            _ = v18
                        } else {
                            v3 = v3
                            var v19 uint8 = 209
                            v8 = fi11
                            _ = v19
                        }
                        _ = fi11
                    }
                case "foo bar":
                    v4 += v4
                }
            } else {
                var v20 string = ""
                v2[v8] = v8
                _ = v20
            }
            v7 = (-(v7))
        }
    }
    v4 = v4
    var v21 int = v8
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 int8 = -128
        if (v4 <= (v4 | 65535)) {
            if true {
                v21 = 4294967295
            } else {
                var v24 bool = true
                if v24 {
                    v8 = (^(v21))
                    v25 := []int{fi22, fi22, v8}
                    switch "a" {
                    case "\n\t":
                        v7 += 42
                    case "hello":
                        if (!(652 < 480)) {
                            var v26 int = int(v4)
                            v27 := GenS2{F0: 32767}
                            _ = v26
                            _ = v27
                        } else {
                            if v24 {
                                fmt.Println(v21, v23, v7)
                                fmt.Println(v4, v7)
                                if ((1e10 * 0.0) == (3.14159 + 1e10)) {
                                    v24 = (297 == 433)
                                    if (0.1 != -2.25) {
                                        if ("foo bar" < "a") {
                                            delete(v1, 2147483647)
                                            fmt.Println(v7, v23, v23)
                                            v28 := GenS3{F0: 65535, F1: 1, F2: 465}
                                            _ = v28
                                        } else {
                                            var v29 float32 = 0.1
                                            v29 += 1.5
                                            fmt.Println(v4)
                                            _ = v29
                                        }
                                        _ = len(v2)
                                    } else {
                                        v24 = false
                                    }
                                    var v30 int16 = (v3 - (v3 * (v3 * -68)))
                                    _ = v30
                                } else {
                                    if (1.5 > (1.5 + 0.1)) {
                                        if ((-2.25 + (((0.0 / 3.14159) + (3.14159 * 1.5)) / 1e10)) <= 3.14159) {
                                            v23 -= -1
                                            var v31 uint16 = 1
                                            _ = v31
                                        }
                                        var v32 uint16 = 42
                                        _ = v32
                                    } else {
                                        switch (v8 % ((127) | 1)) {
                                        case -1:
                                            v3 = (v3 >> ((65535) & 63))
                                        default:
                                            var v33 bool = v24
                                            _ = v33
                                        }
                                    }
                                }
                            } else {
                                v7 += (v7 | v7)
                                for fi34 := 0; fi34 < 6; fi34++ {
                                    v21 = (-(fi34))
                                    _ = fi34
                                }
                            }
                            _ = len(v25)
                        }
                    case "":
                        var v35 int = (v8 % (((v21 * int(v3))) | 1))
                        _ = v35
                    default:
                        switch (fi22 << ((0) & 63)) {
                        case -506:
                            if true {
                                if ((v7 / ((951) | 1)) == v7) {
                                    if v24 {
                                        v25 = append(v25, fi22)
                                        fmt.Println(v3, v21)
                                    } else {
                                        var v36 int = (fi22 >> ((v4) & 63))
                                        switch (^(v8)) {
                                        case 108:
                                            v37 := GenS2{F0: (-(v3))}
                                            _ = v37
                                        default:
                                            for fi38 := 0; fi38 < 5; fi38++ {
                                                var v39 int16 = -32768
                                                switch fi38 {
                                                case -1:
                                                    v36 -= fi38
                                                case 0:
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v24 = false
                                                    case "\n\t":
                                                        if len(v25) >= 2 {
                                                            v25 = v25[1:len(v25)]
                                                        }
                                                    case "":
                                                        var v40 uint16 = 779
                                                        _ = v40
                                                    }
                                                default:
                                                    v41 := GenS1{F0: 823}
                                                    _ = v41
                                                }
                                                _ = fi38
                                                _ = v39
                                            }
                                        }
                                        _ = v36
                                    }
                                    switch int(v3) {
                                    case 134:
                                        fmt.Println(fi22, v23, fi22)
                                    case -52:
                                        var v42 int16 = (^(v3))
                                        _ = v42
                                    case 0:
                                        for fi43 := 0; fi43 < 1; fi43++ {
                                            v4 = (v4 - (v4 | v4))
                                            if false {
                                                v44 := GenS1{F0: 2}
                                                _ = v44
                                            } else {
                                                v23 -= (v23 - (v23 << ((821) & 63)))
                                            }
                                            v45 := v1[1]
                                            _ = fi43
                                            _ = v45
                                        }
                                    }
                                }
                            } else {
                                for fi46 := 0; fi46 < 20; fi46++ {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v47 := GenS4{F0: 1, F1: (true && (!false)), F2: (v23 % ((-1) | 1))}
                                        _ = v47
                                    case "a":
                                        v8 = (v8 * (^(v8)))
                                    }
                                    _ = fi46
                                }
                                v48 := GenS5{F0: (!false)}
                                var v49 int64 = 42
                                _ = v48
                                _ = v49
                            }
                        }
                    }
                    _ = v25
                }
                _ = v24
            }
        }
        _ = len(v2)
        _ = fi22
        _ = v23
    }
    if false {
        var v50 uint32 = 350
        v51 := []int{v21, v21}
        _ = v50
        _ = v51
    } else {
        v21 += v8
    }
    if true {
        var v52 int32 = 1
        v52 = v7
        for fi53 := 0; fi53 < 16; fi53++ {
            if true {
                for fi54 := 0; fi54 < 11; fi54++ {
                    v21 += v21
                    switch 32768 {
                    case -9223372036854775808:
                        fmt.Println(v21)
                    }
                    v55 := GenS5{F0: false}
                    _ = fi54
                    _ = v55
                }
            } else {
                fmt.Println(v52)
                delete(v1, 877)
            }
            var v56 uint16 = 0
            _ = fi53
            _ = v56
        }
        _ = v52
    } else {
        for fi57 := 0; fi57 < 10; fi57++ {
            var v58 int = 32767
            var v59 float64 = 0.1
            v60 := GenS2{F0: (v3 << ((v4) & 63))}
            _ = fi57
            _ = v58
            _ = v59
            _ = v60
        }
        var v61 int32 = int32(v3)
        _ = v61
    }
    var v62 uint16 = 723
    for fi63 := 0; fi63 < 2; fi63++ {
        if true {
            delete(v2, -679)
            for fi64 := 0; fi64 < 14; fi64++ {
                v65 := GenS1{F0: 1}
                var v66 uint8 = 127
                var v67 string = ""
                v68 := v1[526]
                _ = fi64
                _ = v65
                _ = v66
                _ = v67
                _ = v68
            }
        } else {
            switch "a" {
            case "\n\t":
                v62 = v62
            case "foo bar":
                var v69 uint = v4
                _ = v69
            case "hello":
                v62 -= v62
            default:
                fmt.Println(fi63, v3, v62)
            }
            v62 = 32767
        }
        if true {
            v62 += (^(v62))
            for fi70 := 0; fi70 < 15; fi70++ {
                switch v21 {
                case -1:
                    v71 := []int{(v8 | (-(v8))), fi70, (v21 - (v8 + 32768)), fi63}
                    _ = v71
                default:
                    v72 := GenS3{F0: 161, F1: v62, F2: 0}
                    _ = v72
                }
                _ = fi70
            }
        }
        _ = fi63
    }
    v4 = v4
    v8 += -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v62:", v62)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v62
    _ = v7
    _ = v8
}
