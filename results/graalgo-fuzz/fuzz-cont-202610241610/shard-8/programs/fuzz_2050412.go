package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int16; F2 bool }
type GenS4 struct { F0 int8; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 int; F1 uint; F2 int8 }
type GenS7 struct { F0 uint; F1 int64; F2 uint }
type GenS8 struct { F0 string; F1 uint64; F2 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 bool }
func genFn0(p0 int32, p1 uint32) int32 {
    v7 := GenS1{F0: ("\n\t" + "foo bar")}
    fmt.Println(p1)
    v8 := []int{-788}
    v8 = append(v8, 42)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (^(p0))
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 9223372036854775807
    var v3 string = "hello"
    var v4 bool = true
    v5 := make(map[string]int)
    v6 := make(map[uint32]int)
    v3 = (("a" + "hello") + v3)
    fmt.Println(v3, v3)
    if (!(127 != 42)) {
        v2 += 42
    } else {
        var v9 float64 = (-2.25 * 1e10)
        _ = v9
    }
    var v10 int8 = 1
    v3 = (v3 + v3)
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(v3)
        _ = fi11
    }
    if v4 {
        v12 := GenS1{F0: ("foo bar" + "")}
        v13 := []int{598, -1, 32768, 4294967295}
        var v14 uint8 = 2
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        v1[v3] = v4
        _ = fi15
    }
    switch v3 {
    case "":
        v3 = v3
    case "a":
        if v4 {
            v16 := []int{32768, 0, 32767, -1}
            _ = len(v16)
            var v17 uint16 = 879
            _ = v16
            _ = v17
        }
    case "foo bar":
        if v4 {
            fmt.Println(v2)
            var v18 uint16 = 2
            _ = v18
        }
    default:
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := GenS2{F0: 4294967295, F1: 18446744073709551615}
            for fi21 := 0; fi21 < 8; fi21++ {
                v22 := GenS3{F0: 922, F1: 1, F2: (128 >= 94)}
                _ = fi21
                _ = v22
            }
            _ = fi19
            _ = v20
        }
    }
    var v23 string = "hello"
    v24 := GenS3{F0: 1, F1: -32768, F2: (-2147483648 == 32768)}
    v2 -= (v2 % ((-1) | 1))
    v25 := GenS4{F0: (-(v10)), F1: -2147483648, F2: 242}
    v10 = v10
    switch v3 {
    case "foo bar":
        var v26 bool = v4
        _ = v26
    case "hello":
        if v4 {
            for fi27 := 0; fi27 < 4; fi27++ {
                v28 := GenS4{F0: v10, F1: -1, F2: 255}
                _ = len(v5)
                _ = fi27
                _ = v28
            }
            switch 9223372036854775807 {
            case 0:
                v29 := GenS5{F0: 18446744073709551615, F1: 28}
                _ = v29
            case 32768:
                fmt.Println(v4, v4)
            case -272:
                for fi30 := 0; fi30 < 4; fi30++ {
                    switch fi30 {
                    case 32767:
                        switch 128 {
                        case -978:
                            switch fi30 {
                            case 684:
                                var v31 uint16 = 2
                                _ = v31
                            case 32768:
                                fmt.Println(v2, v4)
                            case 0:
                                fmt.Println(v3, v4)
                            default:
                                var v32 int32 = 1
                                _ = v32
                            }
                        }
                    case -9223372036854775808:
                        v23 = v23
                    case -507:
                        fmt.Println(v4, v2)
                    default:
                        v2 = v2
                    }
                    switch v3 {
                    case "a":
                        var v33 string = ""
                        _ = v33
                    default:
                        v3 = ("\n\t" + v23)
                    }
                    fmt.Println(v23)
                    _ = fi30
                }
            }
            for fi34 := 0; fi34 < 20; fi34++ {
                v3 = ((v23 + v3) + "a")
                v3 = "foo bar"
                _ = fi34
            }
        } else {
            fmt.Println(v3, v4, v10)
            v5["hello"] = 957
            switch -9223372036854775808 {
            case 4294967295:
                fmt.Println(v3)
            case -9223372036854775808:
                var v35 uint16 = 1
                _ = v35
            case -1:
                for fi36 := 0; fi36 < 4; fi36++ {
                    for fi37 := 0; fi37 < 3; fi37++ {
                        var v38 int16 = 42
                        v39 := []int{}
                        v2 += (v2 % ((v2) | 1))
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    _ = fi36
                }
            }
        }
    case "\n\t":
        for fi40 := 0; fi40 < 5; fi40++ {
            if (v4 || (53 > v10)) {
                v2 -= 9223372036854775807
            } else {
                if (!false) {
                    fmt.Println(fi40, v2)
                } else {
                    v41 := []int{fi40}
                    for fi42 := 0; fi42 < 19; fi42++ {
                        var v43 uint64 = 888
                        if len(v41) > 0 {
                            v41[0] = fi42
                        }
                        _ = fi42
                        _ = v43
                    }
                    _ = v41
                }
                if ((((809 >= 0) && (v4 && v4)) || (true || false)) || v4) {
                    for fi44 := 0; fi44 < 20; fi44++ {
                        v3 = v23
                        v3 = (v3 + ("\n\t" + v3))
                        _ = fi44
                    }
                    var v45 float64 = 1e10
                    var v46 uint16 = 109
                    _ = v45
                    _ = v46
                } else {
                    var v47 bool = v4
                    switch (fi40 % ((fi40) | 1)) {
                    case -888:
                        var v48 float32 = 1.5
                        _ = v48
                    case 4294967295:
                        switch "foo bar" {
                        case "a":
                            var v49 int64 = v2
                            _ = v49
                        case "foo bar":
                            switch (v3 + "foo bar") {
                            case "":
                                v50 := GenS2{F0: 0, F1: 18446744073709551615}
                                _ = v50
                            default:
                                switch -9223372036854775808 {
                                case 2147483647:
                                    var v51 int16 = -833
                                    _ = v51
                                case 9223372036854775807:
                                    if v4 {
                                        for fi52 := 0; fi52 < 1; fi52++ {
                                            switch v3 {
                                            case "a":
                                                v53 := GenS6{F0: (fi40 << ((4294967295) & 63)), F1: 18446744073709551615, F2: v10}
                                                _ = v53
                                            case "foo bar":
                                                v47 = v47
                                            }
                                            var v54 float64 = (1e10 - 0.1)
                                            _ = fi52
                                            _ = v54
                                        }
                                        for fi55 := 0; fi55 < 3; fi55++ {
                                            v56 := GenS7{F0: 1, F1: (v2 | v2), F2: 2}
                                            var v57 uint16 = 339
                                            _ = fi55
                                            _ = v56
                                            _ = v57
                                        }
                                        v58 := GenS6{F0: fi40, F1: 4294967295, F2: -9}
                                        _ = v58
                                    } else {
                                        var v59 int = (fi40 << ((627) & 63))
                                        var v60 int8 = v10
                                        v2 = v2
                                        _ = v59
                                        _ = v60
                                    }
                                default:
                                    v61 := GenS8{F0: (v23 + "foo bar"), F1: 14, F2: (^(v2))}
                                    _ = v61
                                }
                            }
                        default:
                            switch fi40 {
                            case 9223372036854775807:
                                fmt.Println(v4, v23, v2)
                            }
                        }
                    case 249:
                        if v47 {
                            var v62 int64 = v2
                            if (v23 >= "hello") {
                                var v63 int = (fi40 << ((255) & 63))
                                _ = v63
                            } else {
                                if ("a" > v3) {
                                    switch fi40 {
                                    case 555:
                                        var v64 uint = 0
                                        _ = v64
                                    case 908:
                                        v23 = "\n\t"
                                    }
                                    var v65 int64 = (^(v62))
                                    switch (^(fi40)) {
                                    case 251:
                                        v66 := []int{}
                                        _ = v66
                                    default:
                                        v65 += v2
                                    }
                                    _ = v65
                                }
                                v62 = v62
                            }
                            _ = v62
                        }
                    default:
                        v4 = (!(827 < 18446744073709551615))
                    }
                    _ = v47
                }
                v67 := GenS9{F0: v3}
                _ = v67
            }
            var v68 float32 = (1e10 + -2.25)
            _ = fi40
            _ = v68
        }
    default:
        v69 := GenS10{F0: false}
        _ = v69
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
