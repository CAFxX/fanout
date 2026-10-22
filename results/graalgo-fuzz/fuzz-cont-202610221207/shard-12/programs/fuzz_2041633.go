package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int }
func genFn0(p0 int, p1 uint8) uint32 {
    var v5 uint = 65535
    for fi6 := 0; fi6 < 7; fi6++ {
        p1 -= 127
        var v7 string = "a"
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 2147483648
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 0.1
    var v3 uint = 0
    v4 := make(map[uint32]uint32)
    switch "\n\t" {
    case "a":
        if true {
            switch -9223372036854775808 {
            case 1:
                switch "" {
                case "hello":
                    var v8 float64 = (float64(1e10) / 3.14159)
                    _ = v8
                case "\n\t":
                    for fi9 := 0; fi9 < 5; fi9++ {
                        var v10 int8 = -30
                        _ = fi9
                        _ = v10
                    }
                case "a":
                    for fi11 := 0; fi11 < 5; fi11++ {
                        v12 := []int{fi11, fi11, fi11}
                        for fi13 := 0; fi13 < 18; fi13++ {
                            v14 := GenS1{F0: v3, F1: -1}
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            if true {
                                fmt.Println(v3, v2)
                                var v15 uint8 = 2
                                _ = v15
                            }
                            _ = fi13
                            _ = v14
                        }
                        _ = fi11
                        _ = v12
                    }
                }
            case 0:
                if false {
                    v2 = -2.25
                }
            case -1:
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 uint32 = 2
                    var v18 int64 = -1
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            default:
                for fi19 := 0; fi19 < 2; fi19++ {
                    v3 = v3
                    v2 -= (v2 - v2)
                    var v20 uint64 = 18446744073709551615
                    _ = fi19
                    _ = v20
                }
            }
        } else {
            switch 2147483647 {
            case 660:
                v3 += v3
            default:
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 uint16 = 65535
                    v23 := []int{(fi21 ^ (fi21 + (fi21 | fi21))), fi21}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            }
            var v24 uint8 = 2
            v25 := GenS1{F0: v3, F1: -817}
            _ = v24
            _ = v25
        }
    case "\n\t":
        fmt.Println(v2, v3)
    case "":
        v26 := GenS1{F0: 0, F1: -9223372036854775808}
        _ = v26
    }
    v27 := GenS2{F0: "hello", F1: 762}
    switch "" {
    case "foo bar":
        var v28 float64 = 0.1
        _ = v28
    case "a":
        fmt.Println(v2, v2)
    case "":
        switch "foo bar" {
        case "":
            if false {
                v3 -= (v3 ^ v3)
                v2 = ((((0.1 + v2) * v2) + ((v2 + 1e10) - 1e10)) + v2)
                if true {
                    v2 += v2
                    for fi29 := 0; fi29 < 17; fi29++ {
                        delete(v4, 0)
                        switch ("\n\t" + "hello") {
                        case "":
                            v2 = v2
                        case "\n\t":
                            v3 = (v3 + v3)
                        default:
                            if ("hello" != ("foo bar" + "hello")) {
                                fmt.Println(v3, fi29)
                                v3 += (v3 & (v3 % (((v3 / (((^(v3))) | 1))) | 1)))
                            } else {
                                if true {
                                    var v30 int = fi29
                                    _ = v30
                                }
                                for fi31 := 0; fi31 < 15; fi31++ {
                                    var v32 int64 = 42
                                    switch "foo bar" {
                                    case "hello":
                                        if (!true) {
                                            v2 = (3.14159 - v2)
                                            v32 = 916
                                            var v33 bool = false
                                            _ = v33
                                        }
                                    case "foo bar":
                                        v34 := GenS3{F0: 32}
                                        _ = v34
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                                v35 := GenS4{F0: -1}
                                _ = v35
                            }
                        }
                        var v36 int32 = 1
                        _ = fi29
                        _ = v36
                    }
                    fmt.Println(v2, v2)
                }
            }
        case "a":
            v37 := GenS3{F0: 2}
            _ = v37
        default:
            v3 += 591
        }
    default:
        var v38 float64 = (3.14159 + 0.0)
        _ = v38
    }
    var v39 float64 = (float64(float32(float64(v2))) - 3.14159)
    v40 := GenS1{F0: 65535, F1: 1}
    v3 += v3
    v41 := GenS1{F0: v3, F1: 1}
    v42 := GenS5{F0: 32767}
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
