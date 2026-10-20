package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint64; F1 int16; F2 bool }
type GenS4 struct { F0 int; F1 uint; F2 int16 }
type GenS5 struct { F0 int32; F1 int; F2 int32 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 bool; F1 int16; F2 int16 }
type GenS8 struct { F0 int8; F1 uint64; F2 int }

func main() {
    var v1 uint16 = 1
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    var v4 bool = ((-2.25 / ((1e10 + (0.1 - 0.1)) / 1.0)) == ((0.1 / (-2.25 - 1.5)) - 0.1))
    var v5 float64 = 3.14159
    var v6 float64 = (0.0 * (float64(1.5) / v5))
    if (340 < 630) {
        var v7 bool = (!false)
        v8 := GenS1{F0: "a", F1: -1}
        var v9 float32 = -2.25
        _ = v7
        _ = v8
        _ = v9
    } else {
        switch 127 {
        case 42:
            var v10 uint16 = (v1 | 916)
            _ = v10
        case 9223372036854775807:
            v6 += 3.14159
        case -9223372036854775808:
            for fi11 := 0; fi11 < 18; fi11++ {
                fmt.Println(v5, v4)
                v4 = (fi11 < (fi11 * fi11))
                _ = fi11
            }
        default:
            for fi12 := 0; fi12 < 20; fi12++ {
                if (fi12 != (fi12 ^ (fi12 * fi12))) {
                    v13 := v3[fi12]
                    v14 := GenS2{F0: 32767, F1: 42}
                    if v4 {
                        if v4 {
                            if ((true || v4) || (false && v4)) {
                                v15 := []int{}
                                if true {
                                    if (!(!(561 >= 821))) {
                                        v16 := GenS2{F0: -802, F1: -596}
                                        v6 = 0.1
                                        _ = v16
                                    }
                                    v6 = ((v5 + v6) * v6)
                                    var v17 int64 = 0
                                    _ = v17
                                }
                                if true {
                                    v18 := GenS2{F0: 32767, F1: -32768}
                                    v4 = v4
                                    v19 := GenS1{F0: "a", F1: 127}
                                    _ = v18
                                    _ = v19
                                } else {
                                    switch fi12 {
                                    case -9223372036854775808:
                                        v1 -= (^(v1))
                                    case 127:
                                        v4 = (-1 != 0)
                                    case 528:
                                        for fi20 := 0; fi20 < 18; fi20++ {
                                            v21 := v2[""]
                                            v21 = (!v4)
                                            if v4 {
                                                if (!v4) {
                                                    fmt.Println(v5, v6, v6)
                                                }
                                                v22 := v3[fi20]
                                                _ = v22
                                            }
                                            _ = fi20
                                            _ = v21
                                        }
                                    default:
                                        v5 -= v5
                                    }
                                    fmt.Println(v6, fi12)
                                }
                                _ = v15
                            } else {
                                if (!(!v4)) {
                                    _ = len(v3)
                                    var v23 string = ("" + ("hello" + "foo bar"))
                                    _ = v23
                                } else {
                                    var v24 uint64 = 0
                                    v25 := GenS3{F0: v24, F1: 1, F2: (!v4)}
                                    for fi26 := 0; fi26 < 9; fi26++ {
                                        var v27 uint32 = v13
                                        var v28 float32 = 0.0
                                        fmt.Println(fi12)
                                        _ = fi26
                                        _ = v27
                                        _ = v28
                                    }
                                    _ = v24
                                    _ = v25
                                }
                            }
                        } else {
                            if v4 {
                                delete(v3, (fi12 & 128))
                                v13 += v13
                            }
                            var v29 uint64 = 1
                            v30 := GenS2{F0: -1, F1: int16(v29)}
                            _ = v29
                            _ = v30
                        }
                        v31 := GenS3{F0: 0, F1: 435, F2: v4}
                        _ = v31
                    }
                    _ = v13
                    _ = v14
                } else {
                    if v4 {
                        var v32 uint8 = 127
                        var v33 string = "foo bar"
                        _ = v32
                        _ = v33
                    } else {
                        v1 = v1
                    }
                }
                fmt.Println(fi12, v1, v6)
                if (false && (!(2 <= 128))) {
                    v5 = v5
                } else {
                    var v34 bool = v4
                    _ = v34
                }
                _ = fi12
            }
        }
    }
    v35 := GenS4{F0: 127, F1: 0, F2: 42}
    var v36 uint8 = 127
    for fi37 := 0; fi37 < 17; fi37++ {
        v1 -= v1
        var v38 uint16 = (v1 & (v1 - v1))
        switch (^(fi37)) {
        case 42:
            switch "hello" {
            case "foo bar":
                v6 = v6
            }
        case 1:
            var v39 int8 = 1
            _ = v39
        case -89:
            v5 = v5
        }
        _ = fi37
        _ = v38
    }
    if v4 {
        v40 := []int{912, 0, -1}
        _ = v40
    } else {
        v41 := GenS5{F0: 0, F1: 127, F2: 893}
        _ = v41
    }
    fmt.Println(v36)
    v42 := GenS1{F0: "hello", F1: -128}
    if v4 {
        v43 := GenS6{F0: -983, F1: 65535}
        _ = v43
    } else {
        _ = len(v3)
        v44 := []int{-255}
        _ = v44
    }
    var v45 int = 32768
    switch "foo bar" {
    case "a":
        switch "a" {
        case "foo bar":
            v4 = v4
        }
    case "hello":
        v46 := GenS3{F0: 2, F1: 771, F2: ((0 == 0) && (v1 == (^(v1))))}
        _ = v46
    default:
        for fi47 := 0; fi47 < 2; fi47++ {
            _ = len(v2)
            if (v4 && v4) {
                switch (v45 + 977) {
                case 1:
                    var v48 bool = true
                    _ = v48
                case 0:
                    v36 += (v36 << ((2) & 63))
                case -327:
                    if v4 {
                        if (80 < 9) {
                            v49 := GenS7{F0: false, F1: 42, F2: -174}
                            v1 = 405
                            if true {
                                v36 = (^(v36))
                                v50 := GenS1{F0: "", F1: 100}
                                _ = v50
                            } else {
                                v1 += (v1 + 1)
                            }
                            _ = v49
                        } else {
                            v36 -= v36
                            fmt.Println(v45, v45)
                        }
                        if v4 {
                            delete(v3, 1)
                            v36 += v36
                        }
                    }
                default:
                    if ((18446744073709551615 > 868) || (!(!v4))) {
                        if v4 {
                            if ((v4 && v4) && true) {
                                v51 := []int{(^(fi47)), (fi47 * (v45 + fi47))}
                                _ = v51
                            } else {
                                v52 := GenS8{F0: 100, F1: 2, F2: v45}
                                _ = v52
                            }
                        } else {
                            var v53 uint = 279
                            for fi54 := 0; fi54 < 8; fi54++ {
                                var v55 bool = false
                                var v56 bool = (3.14159 != (float32((v6 / 3.14159)) * 0.0))
                                v2[("" + (((("\n\t" + (((("" + ((("a" + "\n\t") + "hello") + (("foo bar" + "foo bar") + "\n\t"))) + "a") + "\n\t") + ("\n\t" + ("hello" + "a")))) + "") + ("\n\t" + "foo bar")) + "\n\t"))] = v4
                                _ = fi54
                                _ = v55
                                _ = v56
                            }
                            _ = v53
                        }
                        fmt.Println(v1, v45)
                        var v57 uint16 = 0
                        _ = v57
                    }
                }
            } else {
                var v58 uint16 = v1
                _ = v58
            }
            if ((int32(v36) > 2147483647) || (24 > 1)) {
                var v59 float32 = -2.25
                v60 := []int{v45, (v45 % (((v45 | (fi47 * (v45 << ((255) & 63))))) | 1))}
                v61 := GenS5{F0: 2147483647, F1: fi47, F2: -194}
                _ = v59
                _ = v60
                _ = v61
            }
            _ = fi47
        }
    }
    for fi62 := 0; fi62 < 12; fi62++ {
        switch "hello" {
        case "a":
            v1 += (v1 * v1)
        }
        _ = fi62
    }
    var v63 uint = 42
    if false {
        for fi64 := 0; fi64 < 18; fi64++ {
            var v65 int16 = -711
            var v66 uint64 = 18446744073709551615
            fmt.Println(v66, fi64, v63)
            _ = fi64
            _ = v65
            _ = v66
        }
        fmt.Println(v4, v6)
        delete(v2, "hello")
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v45
    _ = v5
    _ = v6
    _ = v63
}
