package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 bool; F2 int16 }
type GenS5 struct { F0 uint64; F1 int32 }
func genFn0(p0 int32, p1 int8) uint16 {
    switch 31 {
    case 127:
        if true {
            var v4 uint16 = 1
            v5 := []int{127, -1, 32767}
            for fi6 := 0; fi6 < 4; fi6++ {
                v7 := GenS1{F0: 154}
                _ = fi6
                _ = v7
            }
            _ = v4
            _ = v5
        } else {
            p1 += (p1 * p1)
            for fi8 := 0; fi8 < 11; fi8++ {
                p0 = 376
                if false {
                    p0 = (p0 % (((-(p0))) | 1))
                } else {
                    p1 = -110
                }
                _ = fi8
            }
        }
    case 745:
        p1 += 127
    case -190:
        var v9 uint64 = 4294967295
        _ = v9
    default:
        p1 += (p1 >> ((42) & 63))
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(p1, fi11)
            p0 -= p0
            switch 32767 {
            case 757:
                switch fi10 {
                case 1:
                    var v12 int8 = 127
                    _ = v12
                case 127:
                    v13 := GenS2{F0: -95, F1: (-(p0))}
                    var v14 int64 = 42
                    p1 += (^(p1))
                    fmt.Println(fi10, fi11)
                    _ = v13
                    _ = v14
                default:
                    switch "hello" {
                    case "a":
                        fmt.Println(fi11)
                    case "":
                        fmt.Println(p1, fi10, fi11)
                    }
                }
            case -9223372036854775808:
                var v15 bool = (!false)
                _ = v15
            default:
                v16 := GenS3{F0: 684}
                _ = v16
            }
            fmt.Println(p1, p1)
            if false {
                v17 := GenS1{F0: 2147483647}
                fmt.Println(p1, p1)
                _ = v17
            }
            var v18 int = int(p1)
            p1 -= 127
            _ = fi11
            _ = v18
        }
        _ = fi10
    }
    fmt.Println(p0, p0)
    if false {
        p0 += p0
    } else {
        fmt.Println(p1, p1)
    }
    v19 := GenS2{F0: 0, F1: -389}
    _ = p0
    _ = p1
    _ = v19
    return 613
}

func main() {
    v1 := []string{"\n\t", "a", "", "foo bar"}
    v2 := make(map[string]int)
    var v3 int = 0
    v3 -= v3
    fmt.Println(v3)
    fmt.Println(v3)
    var v20 uint = 42
    var v21 int = v3
    v20 = 175
    var v22 uint8 = 242
    if (!true) {
        switch ("hello" + "a") {
        case "a":
            v21 = 127
        case "\n\t":
            var v23 float32 = 1e10
            _ = v23
        case "foo bar":
            switch 0 {
            case 638:
                if (float64(((float32(-2.25) - -2.25) / float32((0.1 + (-2.25 - 0.0))))) == (float64(1.5) + (((1e10 / 1.0) / (1e10 * -2.25)) + (3.14159 / (((0.1 + 1.5) + -2.25) + 0.0))))) {
                    var v24 uint64 = 1
                    v22 = v22
                    var v25 float32 = 1.5
                    _ = v24
                    _ = v25
                } else {
                    v26 := GenS2{F0: -1, F1: 565}
                    _ = v26
                }
            case -1:
                fmt.Println(v20, v21, v3)
            }
        }
        v3 += -2147483648
    } else {
        if (!false) {
            v21 = (v21 / (((v21 & v3)) | 1))
            v27 := GenS4{F0: true, F1: false, F2: 42}
            _ = v27
        } else {
            v20 = v20
            var v28 uint16 = 32767
            _ = v28
        }
        if false {
            for fi29 := 0; fi29 < 16; fi29++ {
                v30 := GenS4{F0: (9223372036854775807 == 42), F1: (325 > 1), F2: 1000}
                v20 = 42
                switch ("\n\t" + "a") {
                case "":
                    for fi31 := 0; fi31 < 13; fi31++ {
                        if true {
                            v32 := GenS5{F0: 376, F1: 2147483647}
                            v2[(("a" + ("a" + "")) + "a")] = 4294967295
                            _ = v32
                        } else {
                            v21 -= fi29
                        }
                        switch "a" {
                        case "a":
                            switch fi31 {
                            case -380:
                                v33 := GenS4{F0: true, F1: true, F2: -1}
                                _ = v33
                            default:
                                var v34 uint = (v20 / ((v20) | 1))
                                _ = v34
                            }
                        case "hello":
                            if (v3 == -489) {
                                if (false && true) {
                                    fmt.Println(fi31, v21)
                                }
                            } else {
                                v1 = append(v1, "foo bar")
                            }
                        }
                        var v35 uint64 = 42
                        _ = fi31
                        _ = v35
                    }
                case "foo bar":
                    v3 = fi29
                }
                _ = fi29
                _ = v30
            }
            v21 = v21
        } else {
            if true {
                var v36 string = "\n\t"
                _ = v36
            }
            switch "foo bar" {
            case "\n\t":
                var v37 uint8 = 42
                _ = v37
            case "a":
                var v38 uint = v20
                _ = v38
            default:
                var v39 float32 = 0.0
                _ = v39
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
