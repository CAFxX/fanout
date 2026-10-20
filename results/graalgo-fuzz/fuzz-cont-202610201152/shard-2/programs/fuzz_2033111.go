package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 bool; F2 int }
func genFn0(p0 int32, p1 string, p2 uint32) int {
    p0 -= (p0 % (((p0 | (p0 + (p0 << ((18446744073709551615) & 63))))) | 1))
    var v4 uint = 65535
    var v5 int16 = 479
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 127
}
func genFn1(p0 uint, p1 int32, p2 int8) int8 {
    v6 := []int{}
    v7 := GenS1{F0: (p0 ^ p0), F1: (p1 * (p1 | p1))}
    var v8 uint16 = 0
    p0 = 675
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return p2
}

func main() {
    var v1 uint32 = 787
    v2 := make(map[int]bool)
    var v3 uint16 = 2
    fmt.Println(v1)
    if false {
        v2[4294967295] = true
        if false {
            if true {
                v3 = (v3 + (v3 << ((65535) & 63)))
                for fi9 := 0; fi9 < 4; fi9++ {
                    switch ("foo bar" + "") {
                    case "hello":
                        for fi10 := 0; fi10 < 8; fi10++ {
                            fmt.Println(v1)
                            v1 -= (v1 >> ((717) & 63))
                            v1 += 566
                            _ = fi10
                        }
                    }
                    _ = fi9
                }
            } else {
                var v11 int16 = 1000
                _ = v11
            }
            v12 := GenS1{F0: 42, F1: 42}
            if true {
                v13 := GenS1{F0: 65535, F1: -690}
                fmt.Println(v3, v3)
                switch 165 {
                case -754:
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v15 := GenS2{F0: 1, F1: 0}
                        v1 -= 42
                        var v16 float64 = (0.1 / (1.5 * 1.5))
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                case -851:
                    v3 += v3
                case -1:
                    switch "foo bar" {
                    case "\n\t":
                        v17 := GenS1{F0: 18446744073709551615, F1: -1}
                        _ = v17
                    default:
                        var v18 uint64 = 99
                        _ = v18
                    }
                default:
                    var v19 bool = (true && (true && false))
                    _ = v19
                }
                _ = v13
            } else {
                switch "foo bar" {
                case "foo bar":
                    v20 := []int{1, 128, 128, genFn0(-576, "a", v1)}
                    _ = v20
                case "a":
                    v1 -= 4294967295
                case "hello":
                    switch "a" {
                    case "a":
                        fmt.Println(v3, v1)
                    case "hello":
                        fmt.Println(v1, v1)
                    }
                default:
                    var v21 uint8 = 2
                    _ = v21
                }
                for fi22 := 0; fi22 < 3; fi22++ {
                    if true {
                        fmt.Println(v1)
                    } else {
                        if false {
                            v23 := GenS3{F0: 42}
                            v24 := []int{genFn0(-287, (("a" + (("hello" + ("" + "\n\t")) + ("foo bar" + "hello"))) + ""), (^(v1)))}
                            _ = v23
                            _ = v24
                        } else {
                            if false {
                                var v25 int8 = -1
                                _ = v25
                            }
                        }
                        if (!true) {
                            var v26 int8 = int8(fi22)
                            v1 += (v1 ^ uint32(v3))
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 1; fi27++ {
                                v28 := []int{fi22, (fi27 ^ fi27)}
                                v1 = v1
                                _ = fi27
                                _ = v28
                            }
                        }
                        if false {
                            if (42 > 0) {
                                fmt.Println(fi22, v1, v3)
                                v3 += v3
                                v29 := GenS4{F0: genFn0(-102, ("\n\t" + "\n\t"), (v1 / ((2147483647) | 1))), F1: true, F2: 32768}
                                _ = v29
                            } else {
                                var v30 float32 = ((0.0 * 0.1) / -2.25)
                                fmt.Println(v30, v1, v3)
                                for fi31 := 0; fi31 < 16; fi31++ {
                                    v3 += 42
                                    switch fi22 {
                                    case 1:
                                        var v32 int64 = 1
                                        _ = v32
                                    case 9223372036854775807:
                                        v33 := []int{(fi31 & (fi31 - fi31)), -1, 728, fi22}
                                        _ = v33
                                    default:
                                        var v34 uint16 = 632
                                        _ = v34
                                    }
                                    var v35 uint64 = 42
                                    _ = fi31
                                    _ = v35
                                }
                                _ = v30
                            }
                        }
                    }
                    v1 = v1
                    _ = fi22
                }
            }
            _ = v12
        }
        switch 32768 {
        case 4294967295:
            v1 = (v1 / ((v1) | 1))
        case 32768:
            if true {
                v3 += 127
                v1 -= v1
            }
        }
    } else {
        var v36 int16 = 42
        fmt.Println(v1, v1, v36)
        switch -1 {
        case -40:
            var v37 string = ("hello" + "")
            _ = v37
        case 32768:
            v36 -= (v36 << ((1) & 63))
        case 440:
            v3 += v3
        default:
            var v38 uint16 = 703
            _ = v38
        }
        _ = v36
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        var v40 uint64 = 2
        if true {
            fmt.Println(fi39)
            v3 = v3
            var v41 uint64 = (v40 + uint64(v1))
            _ = v41
        } else {
            v42 := GenS2{F0: genFn1(42, 42, genFn1(255, 42, 100)), F1: 1}
            switch "hello" {
            case "a":
                v40 += 1
            case "\n\t":
                v40 += v40
            case "":
                v40 += (v40 & 1)
            default:
                fmt.Println(v40, v1, fi39)
            }
            _ = v42
        }
        for fi43 := 0; fi43 < 14; fi43++ {
            switch fi39 {
            case -1:
                if (65535 <= (v3 & 65535)) {
                    fmt.Println(fi39, fi43)
                    var v44 uint32 = (^(v1))
                    v3 -= (v3 * 657)
                    _ = v44
                }
            }
            _ = fi43
        }
        _ = fi39
        _ = v40
    }
    var v45 int = -967
    for fi46 := 0; fi46 < 20; fi46++ {
        var v47 uint64 = 0
        var v48 uint = 588
        _ = fi46
        _ = v47
        _ = v48
    }
    switch 687 {
    case 127:
        for fi49 := 0; fi49 < 4; fi49++ {
            var v50 uint = 255
            _ = fi49
            _ = v50
        }
    case 921:
        v45 += (v45 ^ (v45 * 674))
    case -1:
        v45 -= v45
    default:
        v3 += (v3 ^ 0)
    }
    var v51 int32 = 416
    var v52 uint64 = 4294967295
    var v53 uint8 = 127
    var v54 uint8 = v53
    fmt.Println(v1, v53)
    var v55 string = "hello"
    switch (v45 / (((^(v45))) | 1)) {
    case 2147483647:
        v52 = 2
    case 9223372036854775807:
        var v56 int32 = v51
        _ = v56
    case 4294967295:
        for fi57 := 0; fi57 < 5; fi57++ {
            var v58 string = (v55 + v55)
            var v59 uint8 = v54
            _ = fi57
            _ = v58
            _ = v59
        }
    }
    var v60 uint16 = (^(v3))
    for fi61 := 0; fi61 < 15; fi61++ {
        for fi62 := 0; fi62 < 6; fi62++ {
            v54 = (^(v53))
            var v63 uint16 = 42
            v2[399] = true
            _ = fi62
            _ = v63
        }
        _ = fi61
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v60
}
