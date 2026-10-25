package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int32; F1 bool; F2 uint }
type GenS5 struct { F0 string; F1 uint16; F2 uint16 }
func genFn0(p0 uint64, p1 uint8, p2 uint64) string {
    p2 = (p0 % ((p0) | 1))
    if false {
        v4 := []int{32768, -1, 32767}
        p0 -= (p2 + p0)
        var v5 uint32 = 430
        _ = v4
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return (("" + "a") + "foo bar")
}

func main() {
    var v1 float32 = (-2.25 / 1.5)
    var v2 float64 = 0.0
    v3 := []int{-981}
    var v6 int64 = 1
    fmt.Println(v1, v2)
    for fi7 := 0; fi7 < 11; fi7++ {
        if false {
            switch "" {
            case "\n\t":
                fmt.Println(v6, v6, v6)
            case "":
                v8 := GenS1{F0: 128}
                _ = v8
            case "a":
                var v9 string = ("hello" + "foo bar")
                _ = v9
            default:
                v2 -= (v2 - v2)
            }
            switch fi7 {
            case 9223372036854775807:
                fmt.Println(v1, v6, v2)
            case -1:
                fmt.Println(v2, v2)
            case 2147483647:
                v1 += (v1 + (0.0 - 1e10))
            }
            for fi10 := 0; fi10 < 9; fi10++ {
                if (!true) {
                    var v11 float32 = v1
                    _ = v11
                } else {
                    switch "hello" {
                    case "hello":
                        v2 = v2
                        if len(v3) > 0 {
                            v3[0] = (fi10 / ((fi7) | 1))
                        }
                    case "":
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch "" {
                            case "hello":
                                if false {
                                    v1 = v1
                                    if true {
                                        var v12 float32 = (v1 * v1)
                                        v13 := GenS2{F0: 52}
                                        _ = v12
                                        _ = v13
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = (fi7 & 128)
                                        }
                                        v2 -= v2
                                    }
                                }
                            }
                        } else {
                            v14 := GenS3{F0: 255, F1: (!true)}
                            if true {
                                switch "foo bar" {
                                case "hello":
                                    v1 = v1
                                }
                                v6 = v6
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v14
                        }
                    }
                }
                v1 += (v1 + v1)
                var v15 uint64 = 876
                _ = fi10
                _ = v15
            }
        }
        var v16 uint8 = 127
        if false {
            for fi17 := 0; fi17 < 12; fi17++ {
                fmt.Println(v2, v6)
                var v18 uint64 = 18446744073709551615
                var v19 uint = 18446744073709551615
                v20 := GenS4{F0: 1, F1: true, F2: 19}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
        } else {
            if false {
                var v21 int8 = 100
                for fi22 := 0; fi22 < 1; fi22++ {
                    var v23 uint = 0
                    fmt.Println(v23, v6)
                    _ = fi22
                    _ = v23
                }
                var v24 int = -883
                _ = v21
                _ = v24
            } else {
                var v25 float64 = 1.5
                _ = v25
            }
        }
        _ = fi7
        _ = v16
    }
    switch 0 {
    case -9223372036854775808:
        v26 := GenS2{F0: 100}
        _ = v26
    case 2147483647:
        for fi27 := 0; fi27 < 15; fi27++ {
            for fi28 := 0; fi28 < 6; fi28++ {
                v29 := GenS3{F0: 42, F1: true}
                var v30 uint16 = 496
                v2 = float64(v1)
                _ = fi28
                _ = v29
                _ = v30
            }
            var v31 int32 = -29
            var v32 string = ("" + "\n\t")
            _ = fi27
            _ = v31
            _ = v32
        }
    case -248:
        if (42 <= -32768) {
            var v33 uint32 = 65535
            var v34 bool = true
            _ = v33
            _ = v34
        }
    default:
        var v35 bool = (!false)
        _ = v35
    }
    v1 += (1e10 + (1.5 / 1.0))
    var v36 int16 = -32768
    v3 = append(v3, 807)
    v3 = append(v3, 42)
    v37 := GenS5{F0: (("\n\t" + "a") + ("\n\t" + "a")), F1: 65535, F2: 413}
    var v38 int16 = (^(v36))
    for fi39 := 0; fi39 < 20; fi39++ {
        v6 -= v6
        v2 -= 1e10
        _ = fi39
    }
    v6 += (v6 - -9223372036854775808)
    switch 42 {
    case -15:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 9223372036854775807:
        if len(v3) > 0 {
            v3[0] = 32767
        }
    default:
        v2 -= float64(((v1 / (3.14159 + v1)) / (v1 * (3.14159 * 1e10))))
    }
    v3 = append(v3, 127)
    var v40 float32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v40
    _ = v6
}
