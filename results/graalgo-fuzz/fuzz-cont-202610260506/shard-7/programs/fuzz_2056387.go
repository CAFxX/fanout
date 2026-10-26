package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 bool; F1 int; F2 int8 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 int8) uint16 {
    p0 = 100
    p0 -= p0
    var v5 string = ""
    _ = p0
    _ = v5
    return 65535
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -1
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    v2 += -1
    if (1.5 <= float64((0.0 - 0.0))) {
        var v6 float32 = -2.25
        v7 := GenS1{F0: -86, F1: 365, F2: ("a" + ("hello" + "foo bar"))}
        v6 -= (v6 * v6)
        _ = v6
        _ = v7
    } else {
        if true {
            v2 -= 1
        }
        switch "\n\t" {
        case "hello":
            v8 := GenS1{F0: 0, F1: 636, F2: "a"}
            _ = v8
        default:
            if (!true) {
                v2 -= (-(v2))
                v2 -= (v2 & 42)
                v2 = v2
            } else {
                if (!true) {
                    var v9 float64 = 0.0
                    _ = v9
                } else {
                    v2 += 0
                }
                if true {
                    var v10 int = 316
                    var v11 uint32 = 545
                    for fi12 := 0; fi12 < 4; fi12++ {
                        if (192 > 230) {
                            switch v10 {
                            case -9223372036854775808:
                                for fi13 := 0; fi13 < 13; fi13++ {
                                    var v14 uint64 = 376
                                    _ = fi13
                                    _ = v14
                                }
                            default:
                                for fi15 := 0; fi15 < 6; fi15++ {
                                    v16 := []int{2147483647, fi12}
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    fmt.Println(v2)
                                    _ = fi15
                                    _ = v16
                                }
                            }
                        } else {
                            for fi17 := 0; fi17 < 12; fi17++ {
                                v18 := []int{v10, (v10 ^ v10)}
                                _ = fi17
                                _ = v18
                            }
                        }
                        v19 := []int{v10, (-(v10)), v10}
                        _ = fi12
                        _ = v19
                    }
                    _ = v10
                    _ = v11
                } else {
                    for fi20 := 0; fi20 < 1; fi20++ {
                        if (2 == 18446744073709551615) {
                            var v21 uint16 = genFn0(1)
                            if (1 != 0) {
                                fmt.Println(fi20)
                            }
                            _ = v21
                        } else {
                            if (-190 == -9223372036854775808) {
                                fmt.Println(fi20)
                                var v22 bool = true
                                _ = v22
                            } else {
                                var v23 bool = (42 > 1000)
                                var v24 uint8 = 128
                                _ = v23
                                _ = v24
                            }
                            for fi25 := 0; fi25 < 9; fi25++ {
                                v26 := []int{fi20, (-(fi25))}
                                fmt.Println(v2)
                                var v27 uint8 = 128
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            switch 248 {
                            case 2147483647:
                                if true {
                                    v2 += v2
                                }
                            }
                        }
                        v2 += v2
                        _ = fi20
                    }
                    var v28 int8 = 0
                    _ = v28
                }
                var v29 int = 0
                _ = v29
            }
        }
    }
    switch "a" {
    case "foo bar":
        var v30 string = ("" + "a")
        _ = v30
    case "":
        v31 := GenS2{F0: (!(65535 >= 42)), F1: 127, F2: -78}
        _ = v31
    }
    v32 := GenS2{F0: false, F1: 625, F2: 100}
    var v33 float64 = ((3.14159 + 0.0) - float64(1.5))
    _ = len(v3)
    for fi34 := 0; fi34 < 3; fi34++ {
        var v35 string = ("" + "a")
        v36 := GenS3{F0: 18446744073709551615, F1: 127}
        var v37 int8 = 42
        _ = fi34
        _ = v35
        _ = v36
        _ = v37
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        var v39 int32 = v2
        _ = fi38
        _ = v39
    }
    v40 := GenS4{F0: -50}
    if (232 > 127) {
        var v41 float32 = 1e10
        var v42 string = "hello"
        v43 := []int{9223372036854775807, 42, 2147483647, -816}
        _ = v41
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        v2 -= v2
        var v45 uint32 = 4294967295
        for fi46 := 0; fi46 < 7; fi46++ {
            v45 -= 2147483647
            var v47 float64 = float64(1.5)
            fmt.Println(v45, v2, v45)
            _ = fi46
            _ = v47
        }
        _ = fi44
        _ = v45
    }
    var v48 int = 0
    var v49 uint8 = 127
    fmt.Println(v49)
    switch "\n\t" {
    case "":
        for fi50 := 0; fi50 < 19; fi50++ {
            v49 += 255
            fmt.Println(v2, v48, v2)
            _ = fi50
        }
    case "\n\t":
        v2 -= -2147483648
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v40
    _ = v48
    _ = v49
}
