package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint }
type GenS4 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    var v4 bool = false
    switch "a" {
    case "":
        v5 := GenS1{F0: 639, F1: 48, F2: (v4 || (165 < 127))}
        _ = v5
    }
    v6 := []int{-9223372036854775808}
    if ((float32(0.1) - 1.5) != (0.0 * 3.14159)) {
        if ((!(!false)) || v4) {
            for fi7 := 0; fi7 < 18; fi7++ {
                v4 = (!false)
                switch "hello" {
                case "":
                    switch ((("foo bar" + "foo bar") + ("" + "a")) + ("\n\t" + ("" + ("\n\t" + "foo bar")))) {
                    case "foo bar":
                        fmt.Println(v1, fi7, fi7)
                    case "":
                        var v8 uint16 = 1
                        _ = v8
                    default:
                        v1 += v1
                    }
                default:
                    var v9 bool = v4
                    _ = v9
                }
                v10 := GenS1{F0: 1, F1: 65535, F2: (255 <= 2)}
                _ = fi7
                _ = v10
            }
            fmt.Println(v4, v1)
            for fi11 := 0; fi11 < 7; fi11++ {
                fmt.Println(v1, fi11)
                _ = fi11
            }
        } else {
            switch ("foo bar" + "\n\t") {
            case "hello":
                for fi12 := 0; fi12 < 5; fi12++ {
                    if ((3.14159 - 1e10) <= 3.14159) {
                        fmt.Println(v1, v1)
                    } else {
                        fmt.Println(v1)
                        var v13 int8 = 100
                        v13 = (v13 << ((224) & 63))
                        _ = v13
                    }
                    v1 += 0
                    v14 := GenS1{F0: 4294967295, F1: 255, F2: true}
                    _ = fi12
                    _ = v14
                }
            default:
                switch 32768 {
                case -306:
                    if v4 {
                        v15 := v2[32767]
                        switch ("a" + "\n\t") {
                        case "foo bar":
                            v16 := GenS2{F0: -1}
                            _ = v16
                        default:
                            if true {
                                var v17 int32 = 1
                                v15 = v15
                                v1 -= v1
                                _ = v17
                            }
                        }
                        v1 = v1
                        _ = v15
                    } else {
                        var v18 uint32 = 42
                        _ = v18
                    }
                case 168:
                    v19 := GenS2{F0: 42}
                    _ = v19
                default:
                    v20 := GenS2{F0: 616}
                    _ = v20
                }
            }
            switch 2147483647 {
            case 42:
                v21 := GenS1{F0: 166, F1: 684, F2: (-128 >= -1)}
                _ = v21
            case -160:
                v22 := GenS2{F0: -1}
                _ = v22
            case 683:
                v23 := GenS1{F0: 1, F1: 32767, F2: v4}
                _ = v23
            }
        }
    } else {
        if v4 {
            switch 1 {
            case 0:
                var v24 uint = 4294967295
                _ = v24
            case 1:
                if (0 <= -128) {
                    var v25 bool = (v4 || (42 <= 84))
                    v4 = false
                    for fi26 := 0; fi26 < 12; fi26++ {
                        if (42 == -128) {
                            var v27 uint16 = 42
                            v28 := GenS3{F0: -307, F1: 65535, F2: 1}
                            _ = v27
                            _ = v28
                        }
                        var v29 bool = (("foo bar" + "") <= "foo bar")
                        _ = fi26
                        _ = v29
                    }
                    _ = v25
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v30 := GenS3{F0: 0, F1: 32767, F2: 348}
                    _ = v30
                }
            default:
                for fi31 := 0; fi31 < 9; fi31++ {
                    if false {
                        fmt.Println(v1, v1, v1)
                    }
                    _ = fi31
                }
            }
            delete(v3, ("\n\t" + ""))
        } else {
            for fi32 := 0; fi32 < 10; fi32++ {
                switch int(v1) {
                case -2147483648:
                    if (!v4) {
                        var v33 float32 = ((1.5 - 0.0) + float32(0.1))
                        if (!((v1 << ((0) & 63)) == (^(v1)))) {
                            fmt.Println(fi32, v4)
                            for fi34 := 0; fi34 < 5; fi34++ {
                                if false {
                                    if true {
                                        switch "" {
                                        case "foo bar":
                                            var v35 int16 = 32767
                                            _ = v35
                                        case "a":
                                            var v36 int32 = 2147483647
                                            _ = v36
                                        case "hello":
                                            v37 := GenS3{F0: 42, F1: 1, F2: 42}
                                            _ = v37
                                        }
                                    } else {
                                        var v38 uint32 = 1
                                        v39 := GenS1{F0: (v38 | (v38 & v38)), F1: 255, F2: (float64(1.5) < 0.1)}
                                        _ = v38
                                        _ = v39
                                    }
                                } else {
                                    var v40 int16 = -745
                                    _ = v40
                                }
                                v33 = 0.0
                                fmt.Println(v4, v1, v1)
                                _ = fi34
                            }
                        }
                        _ = v33
                    } else {
                        if v4 {
                            v1 = (v1 * v1)
                        }
                    }
                case 318:
                    var v41 bool = v4
                    _ = v41
                default:
                    for fi42 := 0; fi42 < 5; fi42++ {
                        v6 = append(v6, (-(fi32)))
                        v4 = (true || v4)
                        _ = fi42
                    }
                }
                v43 := GenS4{F0: 0, F1: 2}
                _ = fi32
                _ = v43
            }
            v44 := GenS2{F0: -1}
            if len(v6) > 0 {
                v6[0] = 42
            }
            _ = v44
        }
    }
    fmt.Println(v4, v4)
    var v45 int16 = -32768
    fmt.Println(v4)
    v6 = append(v6, -9223372036854775808)
    v4 = v4
    var v46 int32 = 764
    if false {
        v2[1] = (!false)
        for fi47 := 0; fi47 < 8; fi47++ {
            var v48 uint32 = 904
            if len(v6) > 0 {
                v6[0] = 127
            }
            v49 := GenS2{F0: 0}
            _ = fi47
            _ = v48
            _ = v49
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v50 := GenS1{F0: 2147483647, F1: 32767, F2: false}
        _ = v50
    }
    v51 := GenS4{F0: 2, F1: 601}
    for fi52 := 0; fi52 < 14; fi52++ {
        var v53 int16 = (v45 >> ((42) & 63))
        _ = len(v6)
        switch "\n\t" {
        case "\n\t":
            if len(v6) > 0 {
                v6[0] = (fi52 - (^(fi52)))
            }
        case "a":
            v54 := GenS1{F0: 4294967295, F1: 17, F2: (-9223372036854775808 == -1)}
            _ = v54
        case "foo bar":
            var v55 float32 = 1.5
            _ = v55
        }
        _ = fi52
        _ = v53
    }
    v56 := GenS1{F0: 65535, F1: 2, F2: v4}
    fmt.Println(v46)
    v45 -= v45
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v51
    _ = v56
    _ = v6
}
