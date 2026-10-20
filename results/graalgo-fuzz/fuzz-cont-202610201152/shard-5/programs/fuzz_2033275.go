package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int8 }
func genFn0() uint {
    var v5 uint8 = 0
    if true {
        var v6 uint64 = 54
        var v7 uint16 = 273
        _ = v6
        _ = v7
    }
    _ = v5
    return 1
}
func genFn1() int16 {
    var v8 int16 = 32767
    fmt.Println(v8)
    _ = v8
    return v8
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := make(map[string]int)
    v4 := make(map[uint32]bool)
    var v9 uint8 = 127
    switch ("a" + (("hello" + ("hello" + "hello")) + "foo bar")) {
    case "\n\t":
        fmt.Println(v9)
    case "":
        if true {
            v9 += (^(v9))
        }
    }
    v9 = v9
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            if (!false) {
                fmt.Println(v9, fi11, fi10)
                var v12 float64 = (-2.25 + 0.1)
                delete(v3, "a")
                _ = v12
            } else {
                fmt.Println(v9, v9)
                switch (fi10 - 1) {
                case 4294967295:
                    if false {
                        var v13 int = -9223372036854775808
                        _ = v13
                    }
                case 127:
                    fmt.Println(fi11, fi11, v9)
                }
                var v14 uint8 = 42
                fmt.Println(fi11)
                _ = v14
            }
            v15 := GenS1{F0: 42, F1: 236, F2: int64(fi10)}
            v9 = v9
            _ = fi11
            _ = v15
        }
        _ = fi10
    }
    v9 = (v9 / ((v9) | 1))
    if false {
        if (v9 != v9) {
            var v16 float32 = (float32(((3.14159 / 1e10) - 1.5)) * 1e10)
            _ = v16
        } else {
            var v17 int8 = -1
            var v18 int16 = -32768
            if (1 <= 0) {
                if true {
                    for fi19 := 0; fi19 < 17; fi19++ {
                        fmt.Println(v9, v9)
                        _ = fi19
                    }
                } else {
                    v18 += (-(v18))
                }
                v20 := []int{32768}
                _ = v20
            } else {
                for fi21 := 0; fi21 < 3; fi21++ {
                    fmt.Println(v18)
                    _ = fi21
                }
                for fi22 := 0; fi22 < 5; fi22++ {
                    v9 -= v9
                    v23 := v2[fi22]
                    var v24 int8 = (^(v17))
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            _ = v17
            _ = v18
        }
    } else {
        v25 := GenS1{F0: 1, F1: 4294967295, F2: 0}
        for fi26 := 0; fi26 < 12; fi26++ {
            fmt.Println(v9)
            switch fi26 {
            case 9223372036854775807:
                for fi27 := 0; fi27 < 2; fi27++ {
                    if false {
                        v9 -= (v9 * (v9 - (^(v9))))
                        var v28 int32 = 2147483647
                        v29 := GenS1{F0: (v9 - v9), F1: 238, F2: 1}
                        _ = v28
                        _ = v29
                    } else {
                        var v30 int32 = -630
                        v31 := []int{(fi27 ^ 416), (fi26 + fi26), (fi26 & (^(fi26)))}
                        if (0 <= 42) {
                            if (530 >= 42) {
                                switch "\n\t" {
                                case "hello":
                                    var v32 float64 = ((1e10 / 1.5) / 3.14159)
                                    _ = v32
                                case "":
                                    var v33 uint32 = 2147483648
                                    _ = v33
                                case "foo bar":
                                    fmt.Println(fi26, fi27, v30)
                                default:
                                    v3[((("\n\t" + (("hello" + "hello") + ("a" + "foo bar"))) + ("foo bar" + ("foo bar" + "a"))) + "")] = 717
                                }
                            }
                            switch 32768 {
                            case 42:
                                fmt.Println(v9, fi26)
                            case 127:
                                fmt.Println(v30, v30)
                            default:
                                var v34 float64 = 1e10
                                _ = v34
                            }
                        }
                        _ = v30
                        _ = v31
                    }
                    if (0 != 2) {
                        fmt.Println(fi26, fi26, v9)
                        if false {
                            v9 = 109
                            if false {
                                var v35 int8 = 84
                                v9 = v9
                                _ = v35
                            }
                        } else {
                            v36 := GenS2{F0: 1, F1: 255}
                            _ = v36
                        }
                        switch "a" {
                        case "\n\t":
                            var v37 float32 = 1e10
                            var v38 int16 = 32767
                            _ = v37
                            _ = v38
                        case "hello":
                            var v39 string = "\n\t"
                            _ = v39
                        default:
                            v9 = (v9 << ((255) & 63))
                        }
                    }
                    v9 = (^(v9))
                    if (32767 < 981) {
                        var v40 uint = 0
                        _ = v40
                    } else {
                        var v41 uint16 = 42
                        fmt.Println(v41)
                        var v42 float64 = (0.1 - (3.14159 * (((float64(3.14159) + float64(3.14159)) * -2.25) - 0.1)))
                        _ = v41
                        _ = v42
                    }
                    _ = fi27
                }
            case 32767:
                v9 += 206
            case 0:
                var v43 uint = 65535
                _ = v43
            default:
                var v44 uint16 = 1
                _ = v44
            }
            _ = fi26
        }
        if false {
            v45 := GenS1{F0: 0, F1: 0, F2: 496}
            _ = v45
        }
        _ = v25
    }
    var v46 uint16 = 255
    fmt.Println(v9, v46)
    switch (("foo bar" + "foo bar") + "foo bar") {
    case "a":
        var v47 float32 = 1.5
        _ = v47
    case "hello":
        var v48 uint32 = 65535
        _ = v48
    case "foo bar":
        v49 := GenS1{F0: v9, F1: 2, F2: 1}
        _ = v49
    default:
        delete(v1, -355)
    }
    for fi50 := 0; fi50 < 11; fi50++ {
        v51 := GenS3{F0: 642, F1: -323}
        fmt.Println(v9, v9)
        switch "hello" {
        case "\n\t":
            v46 -= v46
        case "a":
            fmt.Println(v46, v9, v46)
        case "hello":
            var v52 int64 = 1
            _ = v52
        }
        _ = fi50
        _ = v51
    }
    v46 = 65535
    var v53 float64 = float64(0.0)
    if (v46 <= v46) {
        fmt.Println(v46, v46)
        switch "a" {
        case "foo bar":
            for fi54 := 0; fi54 < 4; fi54++ {
                v53 = v53
                _ = fi54
            }
        case "hello":
            var v55 uint16 = (^(v46))
            _ = v55
        case "\n\t":
            v56 := GenS1{F0: (v9 & (v9 >> ((255) & 63))), F1: 61, F2: -1}
            _ = v56
        }
    } else {
        var v57 string = "hello"
        _ = len(v4)
        _ = v57
    }
    v58 := GenS4{F0: 100}
    var v59 uint32 = 990
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v53:", v53)
    fmt.Println("v59:", v59)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v53
    _ = v58
    _ = v59
    _ = v9
}
