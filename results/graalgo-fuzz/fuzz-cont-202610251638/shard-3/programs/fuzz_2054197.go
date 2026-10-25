package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int16; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint32; F2 int16 }
type GenS5 struct { F0 bool; F1 int; F2 uint }
type GenS6 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 0
    var v3 int32 = 2147483647
    v4 := make(map[string]bool)
    var v5 int64 = -90
    v6 := []int{300, 4294967295, 9223372036854775807}
    v3 = v3
    v2 = v2
    if (0 >= 278) {
        var v7 string = "foo bar"
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 19; fi9++ {
                v5 = (-(v5))
                _ = fi9
            }
            v10 := GenS1{F0: 127}
            _ = fi8
            _ = v10
        }
        _ = v7
    } else {
        var v11 bool = true
        _ = v11
    }
    switch 635 {
    case 32767:
        for fi12 := 0; fi12 < 10; fi12++ {
            for fi13 := 0; fi13 < 2; fi13++ {
                fmt.Println(v5)
                _ = len(v4)
                _ = fi13
            }
            if true {
                for fi14 := 0; fi14 < 3; fi14++ {
                    if true {
                        if (48 != 4294967295) {
                            v1[2147483648] = (v5 == v5)
                            fmt.Println(v3, fi14)
                        }
                        if (true || (v3 == (-(v3)))) {
                            v6 = append(v6, fi14)
                            var v15 int32 = -1
                            fmt.Println(v3)
                            v16 := GenS1{F0: 2}
                            var v17 int64 = (v5 & (v5 - v5))
                            _ = v15
                            _ = v16
                            _ = v17
                        } else {
                            var v18 int32 = (v3 << ((18446744073709551615) & 63))
                            _ = v18
                        }
                        var v19 uint16 = 369
                        _ = v19
                    } else {
                        var v20 bool = true
                        _ = v20
                    }
                    fmt.Println(v2)
                    if (42 < 32767) {
                        v5 += (v5 >> ((4294967295) & 63))
                        fmt.Println(v3)
                        if (false && false) {
                            var v21 int8 = 42
                            _ = v21
                        } else {
                            v22 := GenS1{F0: 1}
                            switch fi12 {
                            case 9223372036854775807:
                                v3 -= v3
                            case 0:
                                v23 := GenS1{F0: 255}
                                _ = v23
                            case 42:
                                _ = len(v6)
                            }
                            if true {
                                fmt.Println(v3)
                                fmt.Println(v5)
                                v5 += (v5 + 841)
                            }
                            _ = v22
                        }
                    }
                    if false {
                        v2 = 1
                        v24 := GenS1{F0: 36}
                        _ = v24
                    }
                    v5 -= (v5 << ((2) & 63))
                    v3 += v3
                    v25 := GenS1{F0: 246}
                    _ = fi14
                    _ = v25
                }
                if (1000 <= -464) {
                    v2 = 18446744073709551615
                    switch (((("a" + "hello") + "") + "hello") + (("foo bar" + "\n\t") + "a")) {
                    case "hello":
                        v26 := GenS1{F0: 255}
                        _ = v26
                    }
                } else {
                    for fi27 := 0; fi27 < 1; fi27++ {
                        if ((false && ((!true) && false)) && true) {
                            if false {
                                if (!false) {
                                    fmt.Println(fi12, fi12, fi12)
                                    v28 := v1[840]
                                    v3 = -203
                                    _ = v28
                                }
                                v2 = (v2 + v2)
                                v3 += -1
                                v29 := GenS2{F0: 0}
                                v5 += -1
                                _ = v29
                            } else {
                                v3 = v3
                                var v30 uint16 = 559
                                _ = v30
                            }
                            var v31 int64 = v5
                            v2 = v2
                            _ = v31
                        } else {
                            var v32 uint8 = 42
                            _ = v32
                        }
                        v1[65535] = true
                        _ = fi27
                    }
                    if (42 <= 42) {
                        if (!false) {
                            v5 = (^(v5))
                            var v33 uint64 = (v2 >> ((1) & 63))
                            switch "" {
                            case "":
                                for fi34 := 0; fi34 < 13; fi34++ {
                                    var v35 int16 = 0
                                    var v36 int64 = int64(v33)
                                    if true {
                                        var v37 bool = (2147483648 <= 42)
                                        v38 := GenS2{F0: 1}
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v1[0] = ((v36 - (v5 - -288)) < (v36 ^ v5))
                                        v33 = v2
                                    }
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                }
                            default:
                                v39 := GenS1{F0: 127}
                                _ = v39
                            }
                            _ = v33
                        } else {
                            v40 := GenS3{F0: -1, F1: -32768, F2: 630}
                            _ = v40
                        }
                        var v41 int = (fi12 ^ (^(fi12)))
                        _ = v41
                    } else {
                        v42 := GenS1{F0: 0}
                        var v43 uint16 = 32767
                        v3 += 0
                        _ = v42
                        _ = v43
                    }
                }
                var v44 float32 = 1e10
                _ = v44
            }
            _ = fi12
        }
    case 128:
        v5 = (v5 | int64(v2))
    }
    switch ("hello" + "\n\t") {
    case "":
        var v45 int8 = 1
        _ = v45
    case "\n\t":
        v5 += v5
    case "a":
        v2 = 42
    }
    v46 := GenS3{F0: 1, F1: -900, F2: 1}
    fmt.Println(v3)
    if true {
        fmt.Println(v3, v2, v5)
    }
    switch -1 {
    case -2147483648:
        _ = len(v4)
    }
    v2 = v2
    v5 = v5
    v47 := GenS4{F0: 1000, F1: 0, F2: 0}
    var v48 int64 = v5
    if (!true) {
        var v49 uint64 = v2
        _ = v49
    }
    for fi50 := 0; fi50 < 8; fi50++ {
        var v51 int8 = 104
        var v52 uint32 = 517
        if true {
            switch -455 {
            case 2147483647:
                for fi53 := 0; fi53 < 15; fi53++ {
                    v54 := GenS5{F0: false, F1: 897, F2: 255}
                    if (("a" == (("\n\t" + "foo bar") + "")) || ((!(0 == v51)) || ("foo bar" <= "\n\t"))) {
                        var v55 string = "\n\t"
                        _ = v55
                    } else {
                        v56 := v1[v52]
                        v3 -= v3
                        _ = v56
                    }
                    _ = fi53
                    _ = v54
                }
            case 32767:
                if false {
                    var v57 uint16 = 718
                    _ = len(v6)
                    _ = v57
                } else {
                    v58 := GenS4{F0: 42, F1: (v52 ^ (^(v52))), F2: 42}
                    _ = v58
                }
            case 32768:
                v5 += (v5 >> ((uint(v52)) & 63))
            default:
                fmt.Println(fi50, v51)
            }
            v59 := GenS6{F0: 2}
            _ = v59
        }
        _ = fi50
        _ = v51
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
