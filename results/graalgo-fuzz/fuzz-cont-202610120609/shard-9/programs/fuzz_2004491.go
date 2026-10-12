package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int; F1 bool; F2 uint8 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint; F1 int64 }
func genFn0(p0 int8, p1 bool, p2 uint16) string {
    var v6 float64 = 0.1
    v7 := GenS1{F0: 2, F1: 42}
    for fi8 := 0; fi8 < 11; fi8++ {
        if p1 {
            p2 = p2
            v9 := GenS1{F0: 0, F1: 0}
            var v10 string = "hello"
            _ = v10
            _ = v9
        } else {
            var v11 float32 = (float32(v6) + (3.14159 * 3.14159))
            for fi12 := 0; fi12 < 2; fi12++ {
                if p1 {
                    var v13 int = -363
                    switch (("foo bar" + "\n\t") + "\n\t") {
                    case "foo bar":
                        var v14 uint64 = 0
                        _ = v14
                    default:
                        v15 := GenS2{F0: "a", F1: (p2 / ((1) | 1))}
                        v16 := GenS2{F0: "", F1: p2}
                        v17 := []int{v13}
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    v11 -= v11
                    var v18 uint16 = (^(p2))
                    var v19 int64 = 1
                    _ = v13
                    _ = v18
                    _ = v19
                }
                v20 := GenS1{F0: 2, F1: 0}
                _ = fi12
                _ = v20
            }
            var v21 int16 = -1
            _ = v11
            _ = v21
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            v23 := []int{127}
            _ = fi22
            _ = v23
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return ("\n\t" + ("foo bar" + ("a" + ("foo bar" + ("foo bar" + "\n\t")))))
}
func genFn1(p0 bool, p1 int8) bool {
    for fi24 := 0; fi24 < 16; fi24++ {
        for fi25 := 0; fi25 < 17; fi25++ {
            if (p1 <= 100) {
                p1 = (p1 / ((p1) | 1))
                var v26 bool = false
                v27 := GenS1{F0: 65535, F1: 2147483647}
                _ = v26
                _ = v27
            } else {
                switch 32768 {
                case -9223372036854775808:
                    var v28 int64 = 9223372036854775807
                    _ = v28
                case 4294967295:
                    fmt.Println(p0)
                }
            }
            var v29 int64 = 879
            fmt.Println(fi24, fi25)
            if p0 {
                p1 += p1
                if (42 != 42) {
                    v30 := []int{fi25}
                    _ = v30
                } else {
                    v31 := GenS3{F0: fi25, F1: (p0 && false), F2: 42}
                    _ = v31
                }
            }
            _ = fi25
            _ = v29
        }
        _ = fi24
    }
    var v32 string = "foo bar"
    if ((0.1 * (1e10 + (0.1 / float64(3.14159)))) > (-2.25 - 3.14159)) {
        v33 := []int{}
        p1 = (p1 + 0)
        _ = v33
    } else {
        var v34 int8 = (p1 >> ((18446744073709551615) & 63))
        if ((0.0 < 1.5) || p0) {
            for fi35 := 0; fi35 < 4; fi35++ {
                v36 := GenS4{F0: 1, F1: (599 != -1)}
                p1 -= v34
                var v37 int32 = 2147483647
                _ = fi35
                _ = v36
                _ = v37
            }
            v38 := []int{int(p1), 1, 42, -9223372036854775808}
            p1 = (p1 | (-(v34)))
            _ = v38
        }
        switch 2147483647 {
        case 32768:
            for fi39 := 0; fi39 < 14; fi39++ {
                var v40 uint64 = 4294967295
                v41 := GenS5{F0: 0, F1: 3}
                _ = fi39
                _ = v40
                _ = v41
            }
        default:
            p1 = v34
        }
        _ = v34
    }
    switch 32767 {
    case 0:
        v42 := GenS2{F0: (v32 + v32), F1: 1}
        _ = v42
    }
    _ = p0
    _ = p1
    _ = v32
    return p0
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"\n\t"}
    v3 := []string{"\n\t", ("hello" + "\n\t")}
    var v4 int8 = -62
    var v5 uint16 = 773
    v43 := GenS5{F0: 42, F1: 127}
    v44 := GenS1{F0: 108, F1: 0}
    var v45 int64 = 42
    switch 42 {
    case -2147483648:
        if false {
            fmt.Println(v5, v5, v4)
            for fi46 := 0; fi46 < 9; fi46++ {
                if false {
                    if (v45 >= v45) {
                        v47 := GenS6{F0: 1}
                        _ = v47
                    }
                }
                _ = fi46
            }
        }
    case -278:
        if (18446744073709551615 > 18446744073709551615) {
            switch (((("foo bar" + (("\n\t" + "\n\t") + "foo bar")) + "") + "") + "") {
            case "foo bar":
                var v48 int = -160
                _ = v48
            case "":
                if true {
                    v49 := GenS4{F0: -2147483648, F1: false}
                    v50 := GenS7{F0: 551}
                    if (-2147483648 > -1) {
                        for fi51 := 0; fi51 < 15; fi51++ {
                            v5 = (v5 - v5)
                            for fi52 := 0; fi52 < 8; fi52++ {
                                if ((v4 == v4) && false) {
                                    v4 -= v4
                                }
                                var v53 bool = false
                                switch fi52 {
                                case 2147483647:
                                    var v54 uint16 = v5
                                    _ = v54
                                default:
                                    var v55 uint = 0
                                    _ = v55
                                }
                                _ = fi52
                                _ = v53
                            }
                            switch "\n\t" {
                            case "":
                                var v56 bool = true
                                _ = v56
                            case "a":
                                var v57 uint32 = 4294967295
                                _ = v57
                            case "\n\t":
                                fmt.Println(v4, fi51)
                            }
                            _ = fi51
                        }
                        var v58 int64 = v45
                        _ = v58
                    }
                    _ = v49
                    _ = v50
                } else {
                    for fi59 := 0; fi59 < 6; fi59++ {
                        var v60 bool = true
                        _ = fi59
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 9; fi61++ {
                        for fi62 := 0; fi62 < 13; fi62++ {
                            var v63 string = "a"
                            var v64 int16 = 42
                            var v65 int8 = v4
                            delete(v1, -2)
                            _ = fi62
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                        _ = fi61
                    }
                }
            }
            delete(v1, 4294967295)
            if false {
                var v66 string = "\n\t"
                _ = v66
            }
        }
    case 1:
        v67 := GenS3{F0: 852, F1: false, F2: 127}
        _ = v67
    default:
        v68 := v1[127]
        _ = v68
    }
    _ = len(v1)
    if true {
        if ((v5 + (v5 / ((42) | 1))) <= 65535) {
            fmt.Println(v45, v45)
            v45 -= -177
            var v69 float32 = 0.0
            _ = v69
        }
    }
    v70 := GenS8{F0: 4294967295, F1: -9223372036854775808}
    switch 0 {
    case -2147483648:
        v45 -= 0
    case 127:
        v4 = (^(v4))
    case 128:
        var v71 uint8 = 0
        _ = v71
    }
    var v72 int64 = (v45 + (v45 - v45))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v70
    _ = v72
}
