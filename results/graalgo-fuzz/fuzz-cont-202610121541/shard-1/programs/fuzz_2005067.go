package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }
type GenS2 struct { F0 int64; F1 uint32; F2 string }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint64; F2 bool }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool; F1 uint64; F2 int }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 int64; F1 uint; F2 int8 }
type GenS10 struct { F0 uint; F1 int32; F2 uint8 }

func main() {
    var v1 bool = (1 <= 96)
    var v2 string = (("foo bar" + "a") + ("\n\t" + "\n\t"))
    v3 := make(map[string]int)
    var v4 string = v2
    delete(v3, (("a" + v2) + ("hello" + ("foo bar" + v2))))
    var v5 int = 42
    var v6 int32 = 1
    var v7 int32 = v6
    if v1 {
        v8 := GenS1{F0: 1, F1: v4, F2: 0}
        var v9 uint = 18446744073709551615
        v3[v4] = (v5 + v5)
        _ = v8
        _ = v9
    } else {
        var v10 int32 = v6
        v4 = v4
        _ = v10
    }
    fmt.Println(v5)
    fmt.Println(v6, v4)
    v1 = ((!(1 < 789)) || ((633 < 1) || true))
    if v1 {
        fmt.Println(v5)
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        fmt.Println(v4, v4, v1)
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 uint64 = 1
            switch (fi12 / ((fi11) | 1)) {
            case -651:
                if false {
                    switch "" {
                    case "":
                        if false {
                            v14 := GenS1{F0: 1, F1: ((("hello" + v2) + v2) + (v2 + "a")), F2: -2147483648}
                            var v15 float64 = 0.0
                            v15 = 0.1
                            _ = v14
                            _ = v15
                        }
                    case "a":
                        v7 = (v6 * 42)
                    case "foo bar":
                        fmt.Println(v6, v6, v7)
                    }
                    _ = len(v3)
                    v3[(("foo bar" + (v4 + v2)) + v4)] = fi11
                } else {
                    var v16 int16 = 32767
                    v17 := GenS1{F0: 223, F1: "", F2: (v7 ^ v7)}
                    var v18 uint = 965
                    _ = v16
                    _ = v17
                    _ = v18
                }
            default:
                v19 := GenS2{F0: 1, F1: 2, F2: v4}
                _ = v19
            }
            if (-2.25 >= ((1e10 - (((float64((0.1 - (1e10 / 0.1))) + 1e10) - ((-2.25 + -2.25) * 3.14159)) - (1e10 - 3.14159))) + 1e10)) {
                v7 += (v7 / ((v7) | 1))
                if (!(v1 || true)) {
                    if true {
                        v3[v4] = fi12
                        v20 := GenS2{F0: -1, F1: 65535, F2: "a"}
                        v7 = 932
                        _ = v20
                    } else {
                        if v1 {
                            v21 := GenS2{F0: 0, F1: 605, F2: v2}
                            v22 := v3[v4]
                            _ = v21
                            _ = v22
                        } else {
                            if (204 != 1) {
                                fmt.Println(v6, v13, v5)
                                v23 := GenS3{F0: -182, F1: (v7 / (((v6 << ((339) & 63))) | 1))}
                                v6 -= (^(v7))
                                _ = v23
                            }
                            var v24 float32 = 1e10
                            v2 = v4
                            v13 += (v13 ^ (v13 ^ 18446744073709551615))
                            fmt.Println(v24)
                            _ = v24
                        }
                        v4 = v4
                    }
                    fmt.Println(v2, v1, v7)
                    fmt.Println(v5, fi12, v2)
                    v25 := v3[v2]
                    fmt.Println(v2, v7)
                    _ = v25
                }
            } else {
                v6 = (v7 >> ((661) & 63))
                if true {
                    switch v4 {
                    case "\n\t":
                        _ = len(v3)
                    case "":
                        if (2 > 993) {
                            v26 := GenS2{F0: 42, F1: 520, F2: v2}
                            _ = len(v3)
                            _ = v26
                        } else {
                            fmt.Println(v2)
                            _ = len(v3)
                            var v27 string = (("\n\t" + ("foo bar" + v2)) + "a")
                            _ = v27
                        }
                    case "foo bar":
                        v4 = v2
                        var v28 uint = 0
                        _ = v28
                    }
                    v29 := GenS4{F0: 65535, F1: v13, F2: v1}
                    _ = v29
                }
                if (2147483647 < fi12) {
                    var v30 float64 = 1e10
                    if v1 {
                        if true {
                            if v1 {
                                v31 := []int{(v5 + 127), fi12, (fi11 + fi11)}
                                v4 = ("hello" + v4)
                                var v32 float32 = ((float32(v30) - 3.14159) / 3.14159)
                                _ = v31
                                _ = v32
                            } else {
                                var v33 uint16 = 256
                                v30 = v30
                                _ = v33
                            }
                            v34 := GenS4{F0: 42, F1: (v13 + (^(v13))), F2: (100 > -128)}
                            if (v13 < (^(v13))) {
                                switch v2 {
                                case "foo bar":
                                    var v35 uint = 602
                                    _ = v35
                                }
                                v36 := GenS5{F0: 4294967295, F1: 4294967295, F2: 0}
                                fmt.Println(v1, v5, v1)
                                _ = v36
                            }
                            _ = v34
                        } else {
                            fmt.Println(v2)
                            v37 := GenS4{F0: 506, F1: 0, F2: v1}
                            if false {
                                if ((-128 != -1) || false) {
                                    switch fi12 {
                                    case 1:
                                        v3["hello"] = 0
                                    case 9223372036854775807:
                                        var v38 uint16 = 65535
                                        _ = v38
                                    }
                                    v39 := GenS6{F0: 1}
                                    v1 = v1
                                    _ = v39
                                } else {
                                    fmt.Println(v7)
                                    v40 := GenS7{F0: true, F1: (^(v13)), F2: fi12}
                                    _ = v40
                                }
                                _ = len(v3)
                                v13 = v13
                            } else {
                                switch "\n\t" {
                                case "a":
                                    delete(v3, v2)
                                case "hello":
                                    var v41 int64 = 9223372036854775807
                                    _ = v41
                                case "\n\t":
                                    if true {
                                        var v42 float64 = 3.14159
                                        v2 = ""
                                        v2 = ("hello" + "\n\t")
                                        _ = v42
                                    }
                                default:
                                    var v43 uint64 = (v13 + 18446744073709551615)
                                    _ = v43
                                }
                                v7 -= v7
                            }
                            _ = v37
                        }
                        fmt.Println(v2, v5, v6)
                        if (((((((v30 / v30) - v30) < 3.14159) && (!(1 != 1))) && (float32(v30) >= (1e10 + 3.14159))) || false) && (v1 && ((v13 >> ((780) & 63)) > v13))) {
                            v30 = v30
                            v44 := GenS5{F0: 2, F1: 4294967295, F2: 733}
                            _ = v44
                        }
                    }
                    var v45 float32 = -2.25
                    _ = v30
                    _ = v45
                } else {
                    var v46 uint64 = (v13 % (((v13 - 835)) | 1))
                    v47 := GenS3{F0: (v5 >> ((434) & 63)), F1: v7}
                    v48 := v3[("\n\t" + "foo bar")]
                    _ = v46
                    _ = v47
                    _ = v48
                }
                var v49 uint64 = (^(v13))
                _ = v49
            }
            v13 = 2
            v7 -= v7
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    if (!((!(-1 <= -1)) || v1)) {
        for fi50 := 0; fi50 < 3; fi50++ {
            for fi51 := 0; fi51 < 7; fi51++ {
                v4 = v2
                var v52 int = (fi50 + 42)
                delete(v3, (((v2 + v2) + (v2 + v2)) + v2))
                _ = fi51
                _ = v52
            }
            _ = fi50
        }
        fmt.Println(v5, v4)
        if (3.14159 == (3.14159 / 1.0)) {
            for fi53 := 0; fi53 < 7; fi53++ {
                switch fi53 {
                case -2147483648:
                    for fi54 := 0; fi54 < 9; fi54++ {
                        var v55 uint32 = 65535
                        _ = fi54
                        _ = v55
                    }
                }
                _ = fi53
            }
            var v56 int = (v5 * 689)
            _ = v56
        } else {
            if v1 {
                v5 = -222
            }
            var v57 bool = v1
            _ = v57
        }
    } else {
        v58 := []int{}
        switch "a" {
        case "":
            for fi59 := 0; fi59 < 9; fi59++ {
                v6 = v6
                _ = fi59
            }
        default:
            fmt.Println(v1, v5, v1)
        }
        _ = v58
    }
    v60 := GenS6{F0: (v7 >> ((65535) & 63))}
    v61 := []int{(v5 & v5), (v5 * (v5 / (((v5 << ((442) & 63))) | 1)))}
    for fi62 := 0; fi62 < 10; fi62++ {
        if (!v1) {
            switch fi62 {
            case 2147483647:
                for fi63 := 0; fi63 < 20; fi63++ {
                    fmt.Println(v6, v1)
                    var v64 int = (^(fi63))
                    if v1 {
                        var v65 float32 = float32(0.1)
                        fmt.Println(v7)
                        if (32767 == 65535) {
                            v66 := GenS8{F0: (fi62 * 32768), F1: (128 <= v5)}
                            v2 = "hello"
                            _ = v66
                        }
                        var v67 uint16 = 963
                        var v68 float32 = v65
                        v69 := GenS9{F0: -33, F1: 4294967295, F2: 1}
                        fmt.Println(v65)
                        fmt.Println(v6, fi62, v65)
                        _ = v65
                        _ = v67
                        _ = v68
                        _ = v69
                    } else {
                        v70 := GenS10{F0: 1, F1: v7, F2: 68}
                        var v71 int64 = -228
                        _ = v70
                        _ = v71
                    }
                    _ = len(v3)
                    _ = fi63
                    _ = v64
                }
            default:
                fmt.Println(fi62)
            }
        } else {
            var v72 float64 = 0.1
            v73 := v3[v4]
            var v74 bool = v1
            _ = v72
            _ = v73
            _ = v74
        }
        _ = fi62
    }
    var v75 int32 = (v6 - 42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v7:", v7)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
    _ = v7
    _ = v75
}
