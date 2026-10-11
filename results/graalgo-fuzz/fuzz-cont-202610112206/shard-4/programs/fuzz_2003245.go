package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 bool; F1 int32; F2 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 string; F2 int64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool; F1 int16 }
type GenS11 struct { F0 bool; F1 int16 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 string }
func genFn0(p0 uint16) uint32 {
    v7 := []int{468, -2147483648, -728, -1}
    var v8 float64 = 1.5
    _ = p0
    _ = v7
    _ = v8
    return 18
}
func genFn1() uint16 {
    v9 := GenS1{F0: 2147483647}
    var v10 uint32 = 442
    _ = v10
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{1, 379, 4294967295}
    var v2 int32 = -137
    var v3 float64 = (1e10 - 0.1)
    var v4 float32 = 0.1
    v5 := []int{32768, 0, 2147483647, 0}
    var v6 uint = 65535
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            switch ("a" + "") {
            case "\n\t":
                if true {
                    var v13 int32 = v2
                    _ = v13
                } else {
                    if true {
                        var v14 uint16 = 532
                        v15 := GenS1{F0: 621}
                        _ = v14
                        _ = v15
                    }
                }
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            _ = fi12
        }
        v16 := GenS1{F0: 1}
        v3 -= 3.14159
        _ = fi11
        _ = v16
    }
    var v17 float32 = v4
    v2 += v2
    var v18 int16 = 1000
    switch "a" {
    case "":
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS2{F0: (v18 - v18), F1: 1}
            _ = fi19
            _ = v20
        }
    case "\n\t":
        for fi21 := 0; fi21 < 16; fi21++ {
            if true {
                v18 = (v18 << ((v6) & 63))
                fmt.Println(v2, v17)
            } else {
                v22 := GenS1{F0: 42}
                v23 := GenS3{F0: 0, F1: 102}
                _ = v22
                _ = v23
            }
            if (!false) {
                var v24 int32 = v2
                _ = v24
            } else {
                var v25 uint = v6
                _ = v25
            }
            _ = fi21
        }
    case "hello":
        var v26 uint16 = 171
        _ = v26
    }
    var v27 int = -2147483648
    var v28 float32 = 1.5
    if true {
        var v29 int16 = v18
        _ = v29
    } else {
        if false {
            v30 := GenS4{F0: (v18 % (((^(v18))) | 1))}
            var v31 int8 = int8(v6)
            v6 = (^(v6))
            _ = v30
            _ = v31
        }
    }
    switch -9223372036854775808 {
    case 42:
        v32 := GenS5{F0: 2147483647, F1: 2147483647}
        _ = v32
    case 0:
        fmt.Println(v4, v4)
    default:
        fmt.Println(v27)
    }
    if ((^(v2)) >= v2) {
        v33 := GenS4{F0: -32768}
        v27 = -754
        switch ("\n\t" + ("a" + "a")) {
        case "hello":
            if false {
                v3 = (v3 * (0.0 + (3.14159 + v3)))
                fmt.Println(v4, v17)
            }
        case "":
            for fi34 := 0; fi34 < 11; fi34++ {
                var v35 bool = true
                _ = fi34
                _ = v35
            }
        case "\n\t":
            v36 := GenS6{F0: false, F1: v2, F2: 2}
            _ = v36
        default:
            v18 += v18
        }
        _ = v33
    } else {
        switch "a" {
        case "\n\t":
            v37 := GenS7{F0: 359}
            _ = v37
        case "a":
            switch "" {
            case "a":
                _ = len(v5)
            case "hello":
                fmt.Println(v28)
            case "\n\t":
                v38 := GenS7{F0: 42}
                _ = v38
            }
        case "hello":
            fmt.Println(v27, v6)
        }
        if (1 <= 0) {
            if ((true || true) || true) {
                for fi39 := 0; fi39 < 4; fi39++ {
                    switch "a" {
                    case "":
                        v1 = append(v1, 0)
                    default:
                        for fi40 := 0; fi40 < 3; fi40++ {
                            switch ("a" + "") {
                            case "hello":
                                if false {
                                    if false {
                                        var v41 uint32 = 4294967295
                                        v2 = v2
                                        if true {
                                            var v42 int32 = (v2 >> ((v6) & 63))
                                            if false {
                                                v6 -= (^(v6))
                                            } else {
                                                v42 -= (v2 / (((v42 / ((42) | 1))) | 1))
                                            }
                                            _ = v42
                                        }
                                        _ = v41
                                    } else {
                                        v5 = append(v5, (fi39 << ((v6) & 63)))
                                        if ((("hello" + "hello") + ("" + "")) == "foo bar") {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            if false {
                                                if (51 == 103) {
                                                    v43 := GenS8{F0: v6, F1: "hello", F2: -9223372036854775808}
                                                    switch "\n\t" {
                                                    case "":
                                                        var v44 uint8 = 20
                                                        var v45 int16 = v18
                                                        _ = v44
                                                        _ = v45
                                                    default:
                                                        v46 := GenS9{F0: 712}
                                                        _ = v46
                                                    }
                                                    var v47 uint = v6
                                                    _ = v43
                                                    _ = v47
                                                } else {
                                                    v48 := GenS10{F0: false, F1: v18}
                                                    _ = v48
                                                }
                                                var v49 int32 = 42
                                                fmt.Println(v4, v17)
                                                _ = v49
                                            }
                                        } else {
                                            v6 = (v6 - (v6 * v6))
                                            fmt.Println(v18, v17, v28)
                                        }
                                        if false {
                                            v50 := GenS11{F0: false, F1: (-(v18))}
                                            var v51 int = v27
                                            fmt.Println(v6, v3, fi40)
                                            var v52 uint16 = 255
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                    v6 -= v6
                                } else {
                                    v3 = 1e10
                                    v6 = (^(v6))
                                }
                            default:
                                switch (fi40 >> ((v6) & 63)) {
                                case 190:
                                    v53 := GenS10{F0: false, F1: v18}
                                    _ = v53
                                }
                            }
                            _ = fi40
                        }
                    }
                    _ = fi39
                }
            } else {
                v17 -= 0.0
                var v54 int32 = (v2 - (-(v2)))
                var v55 bool = (true && ((-(v2)) > (-(v2))))
                _ = v54
                _ = v55
            }
            for fi56 := 0; fi56 < 8; fi56++ {
                v18 += (v18 / ((v18) | 1))
                fmt.Println(v27, v6)
                _ = fi56
            }
        } else {
            v57 := GenS12{F0: 42}
            fmt.Println(v18)
            _ = v57
        }
    }
    v58 := GenS13{F0: "foo bar"}
    v4 = ((float32(v3) / (((v17 / ((0.0 + 1e10) - v28)) - v17) - (v4 - (v4 / float32((1e10 / (v3 / 1.5))))))) * (v28 - v4))
    if len(v5) > 0 {
        v5[0] = 859
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
}
