package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int8; F1 uint16; F2 int32 }
type GenS6 struct { F0 int64; F1 uint16 }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 uint64; F1 int32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 string; F1 int8; F2 int8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 string = ("\n\t" + "hello")
    v3 := make(map[int]bool)
    v4 := []uint32{}
    var v5 uint64 = 2
    var v6 int = 58
    var v7 uint8 = 127
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int64 = -1
        v3[fi8] = false
        var v10 int8 = 121
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 42, F1: 32767, F2: 42}
    v1[0] = (v6 ^ v6)
    switch "foo bar" {
    case "a":
        v6 += v6
    case "hello":
        v2 = v2
    case "":
        for fi12 := 0; fi12 < 10; fi12++ {
            switch v2 {
            case "hello":
                var v13 int8 = 1
                _ = v13
            default:
                v5 += v5
            }
            v14 := GenS2{F0: -2147483648}
            _ = fi12
            _ = v14
        }
    }
    fmt.Println(v5, v6, v7)
    v7 += (^(v7))
    v15 := GenS3{F0: (v2 + ""), F1: (^(v5))}
    switch (v6 >> ((42) & 63)) {
    case 57:
        switch (v2 + "") {
        case "a":
            if ((v6 * v6) <= int(v7)) {
                if (127 < 127) {
                    v5 -= (^(v5))
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    switch (v6 | v6) {
                    case -2147483648:
                        for fi16 := 0; fi16 < 1; fi16++ {
                            var v17 uint8 = v7
                            v5 += v5
                            v6 = (fi16 % (((fi16 + 32768)) | 1))
                            _ = fi16
                            _ = v17
                        }
                    case 269:
                        for fi18 := 0; fi18 < 11; fi18++ {
                            for fi19 := 0; fi19 < 15; fi19++ {
                                fmt.Println(fi19, v5, fi19)
                                if false {
                                    fmt.Println(v6, fi18)
                                } else {
                                    var v20 uint64 = (^(v5))
                                    var v21 string = v2
                                    v7 -= (v7 * v7)
                                    _ = v20
                                    _ = v21
                                }
                                v6 += fi19
                                if ((v2 + (v2 + v2)) < "") {
                                    var v22 int32 = -2147483648
                                    v22 += -927
                                    _ = v22
                                } else {
                                    switch v6 {
                                    case 32768:
                                        if true {
                                            fmt.Println(v6, v2, fi19)
                                            if (1e10 == 3.14159) {
                                                v23 := GenS4{F0: true, F1: 42}
                                                delete(v1, 1)
                                                _ = v23
                                            }
                                        } else {
                                            switch (v2 + "\n\t") {
                                            case "foo bar":
                                                if (-71 != 79) {
                                                    fmt.Println(v6)
                                                    v4 = append(v4, 4294967295)
                                                    delete(v1, 65535)
                                                    if (243 <= 1) {
                                                        switch fi19 {
                                                        case 1:
                                                            v5 = v5
                                                            v7 -= 1
                                                            if (v5 > v5) {
                                                                fmt.Println(v5)
                                                                v24 := GenS5{F0: -128, F1: 42, F2: 2147483647}
                                                                _ = v24
                                                            }
                                                        case -1:
                                                            v2 = v2
                                                        }
                                                        v5 += 0
                                                        v25 := GenS1{F0: 1, F1: 2, F2: v6}
                                                        _ = v25
                                                    } else {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    }
                                                    v2 = "hello"
                                                    if false {
                                                        fmt.Println(v5, v5, v5)
                                                        v5 -= 278
                                                        fmt.Println(v5, v6, fi18)
                                                    }
                                                    v26 := GenS4{F0: (32767 == 1), F1: 0}
                                                    var v27 int16 = 1
                                                    if (182 < 2) {
                                                        if (true || true) {
                                                            v28 := GenS2{F0: -128}
                                                            _ = v28
                                                        } else {
                                                            v27 = 1
                                                        }
                                                        switch v2 {
                                                        case "hello":
                                                            v1[2147483647] = (fi19 << ((101) & 63))
                                                        }
                                                    } else {
                                                        v7 = 181
                                                    }
                                                    var v29 float32 = 3.14159
                                                    _ = v26
                                                    _ = v27
                                                    _ = v29
                                                }
                                            case "":
                                                if (fi18 > fi18) {
                                                    v30 := GenS3{F0: "", F1: v5}
                                                    _ = v30
                                                } else {
                                                    v2 = (v2 + "hello")
                                                }
                                                v3[(fi19 >> ((255) & 63))] = false
                                            default:
                                                var v31 string = (v2 + "foo bar")
                                                _ = v31
                                            }
                                            fmt.Println(fi18, v2, v2)
                                            switch fi18 {
                                            case 127:
                                                if len(v4) > 0 {
                                                    v4[0] = 839
                                                }
                                            case 927:
                                                var v32 uint = 0
                                                _ = v32
                                            default:
                                                var v33 float32 = (3.14159 / 1.5)
                                                _ = v33
                                            }
                                        }
                                    default:
                                        switch (fi18 | v6) {
                                        case 32767:
                                            var v34 string = v2
                                            _ = v34
                                        }
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                _ = fi19
                            }
                            v35 := GenS6{F0: 0, F1: 547}
                            _ = fi18
                            _ = v35
                        }
                    }
                    var v36 int = -1
                    _ = v36
                }
                v7 = 127
            }
        case "":
            fmt.Println(v5, v6, v6)
        case "\n\t":
            v37 := GenS7{F0: "foo bar", F1: 2}
            _ = v37
        }
    case -102:
        if true {
            fmt.Println(v7)
            var v38 uint16 = 42
            v39 := v3[-9223372036854775808]
            _ = v38
            _ = v39
        }
    case 32767:
        var v40 uint64 = 18446744073709551615
        _ = v40
    default:
        switch (v2 + v2) {
        case "\n\t":
            fmt.Println(v6)
        case "":
            var v41 int64 = -794
            _ = v41
        }
    }
    switch ((v2 + v2) + v2) {
    case "a":
        for fi42 := 0; fi42 < 11; fi42++ {
            if (!(v5 == (v5 ^ v5))) {
                var v43 float64 = 1e10
                v6 += -786
                _ = v43
            } else {
                if (!true) {
                    v44 := GenS2{F0: 0}
                    if true {
                        v4 = append(v4, 2)
                        v7 = 33
                        fmt.Println(v2, v2, v7)
                    }
                    _ = v44
                } else {
                    var v45 string = ((v2 + "") + v2)
                    switch v2 {
                    case "foo bar":
                        if (-128 == 1) {
                            v45 = v2
                            v1[2147483647] = v6
                        } else {
                            v46 := v3[fi42]
                            v47 := GenS8{F0: 2, F1: (v6 + 9223372036854775807)}
                            v46 = false
                            _ = v46
                            _ = v47
                        }
                    }
                    var v48 string = v2
                    _ = v45
                    _ = v48
                }
            }
            v49 := GenS9{F0: v5, F1: -31}
            _ = fi42
            _ = v49
        }
    case "\n\t":
        v6 += v6
    }
    var v50 uint32 = 659
    var v51 bool = true
    for fi52 := 0; fi52 < 9; fi52++ {
        var v53 int64 = 672
        for fi54 := 0; fi54 < 12; fi54++ {
            var v55 int8 = 58
            if (v2 >= v2) {
                var v56 uint64 = (^(v5))
                if true {
                    v51 = (v50 >= (^(v50)))
                    switch v6 {
                    case 2147483647:
                        v57 := v1[uint32(v5)]
                        _ = v57
                    case 32767:
                        fmt.Println(fi54, v55)
                        delete(v1, (v50 - (^(v50))))
                        if v51 {
                            v58 := GenS1{F0: 32, F1: 65535, F2: (^(fi52))}
                            v56 += 18446744073709551615
                            _ = v58
                        } else {
                            v2 = v2
                        }
                    case -9223372036854775808:
                        switch (^(v6)) {
                        case 1:
                            if true {
                                v7 -= (v7 + 128)
                                v59 := GenS10{F0: (v56 | v5)}
                                v50 = 2147483647
                                _ = v59
                            } else {
                                var v60 bool = false
                                _ = v60
                            }
                        }
                    }
                    fmt.Println(v55)
                    var v61 float32 = 1e10
                    _ = v61
                } else {
                    fmt.Println(v7, v6)
                    v55 = (v55 | (v55 % ((0) | 1)))
                    var v62 int = (fi52 >> ((255) & 63))
                    _ = v62
                }
                if v51 {
                    var v63 uint = 1
                    var v64 string = (v2 + (("\n\t" + v2) + "\n\t"))
                    fmt.Println(v55, fi52, v50)
                    _ = v63
                    _ = v64
                } else {
                    v56 -= 4294967295
                    v65 := v3[2147483647]
                    _ = v65
                }
                _ = v56
            }
            _ = fi54
            _ = v55
        }
        _ = fi52
        _ = v53
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v66 int16 = 32767
    v67 := GenS11{F0: v2, F1: 127, F2: 1}
    if ((0.1 / 1e10) <= 0.1) {
        var v68 float32 = 1.5
        var v69 bool = v51
        _ = v68
        _ = v69
    }
    switch (v6 >> ((456) & 63)) {
    case -1:
        var v70 uint8 = 0
        _ = v70
    case 42:
        for fi71 := 0; fi71 < 3; fi71++ {
            fmt.Println(v5)
            delete(v3, 1)
            var v72 uint32 = (v50 * uint32(v7))
            _ = fi71
            _ = v72
        }
    default:
        fmt.Println(v2, v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v66
    _ = v67
    _ = v7
}
