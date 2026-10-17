package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 uint16 }
func genFn0(p0 int) uint64 {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 string = "foo bar"
        p0 -= fi7
        var v9 float32 = 3.14159
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        fmt.Println(fi10)
        switch p0 {
        case 1:
            if (2147483647 < 42) {
                v11 := []int{127, (fi10 / (((p0 << ((1) & 63))) | 1))}
                for fi12 := 0; fi12 < 5; fi12++ {
                    switch (fi12 - fi10) {
                    case -1:
                        var v13 int32 = -2147483648
                        _ = v13
                    case 2147483647:
                        v14 := GenS1{F0: 445, F1: 9223372036854775807}
                        _ = v14
                    }
                    _ = fi12
                }
                var v15 float64 = 1e10
                _ = v11
                _ = v15
            }
        case -293:
            var v16 float64 = 1.5
            _ = v16
        }
        _ = fi10
    }
    switch 0 {
    case 128:
        fmt.Println(p0)
    case 985:
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 float64 = 0.0
            switch "" {
            case "a":
                p0 -= p0
            case "foo bar":
                fmt.Println(p0, fi17, v18)
            }
            _ = fi17
            _ = v18
        }
    case 9223372036854775807:
        p0 -= (p0 * p0)
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        v20 := []int{(p0 >> ((255) & 63))}
        var v21 float64 = 0.0
        _ = fi19
        _ = v20
        _ = v21
    }
    p0 -= 1
    _ = p0
    return 783
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 4294967295
    var v3 uint8 = 2
    v4 := make(map[int]uint32)
    v5 := []string{"\n\t", ("a" + ("a" + ("hello" + "hello"))), "\n\t"}
    v6 := make(map[string]uint32)
    v22 := GenS1{F0: -406, F1: 0}
    if ((!(911 <= 0)) || false) {
        v3 -= (^(v3))
        _ = len(v1)
        var v23 int = 235
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        if false {
            switch fi24 {
            case -1:
                var v25 float32 = ((0.1 - 0.1) * (0.0 * (1.5 / float32((-2.25 + 3.14159)))))
                _ = v25
            }
            if true {
                v26 := GenS2{F0: (((true && false) || false) && (!false))}
                if true {
                    if false {
                        v2 -= (v2 & (^(v2)))
                        if false {
                            switch (fi24 | (fi24 | 2147483647)) {
                            case 32767:
                                _ = len(v5)
                            case 2147483647:
                                fmt.Println(fi24)
                            case 1:
                                fmt.Println(fi24, fi24, v2)
                            default:
                                var v27 uint16 = 42
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 18; fi28++ {
                                v2 += genFn0((fi24 >> ((42) & 63)))
                                _ = fi28
                            }
                            v1[2147483648] = "\n\t"
                        }
                        fmt.Println(v3)
                    }
                } else {
                    v5 = append(v5, (("\n\t" + "a") + (("\n\t" + "a") + (("a" + "hello") + (((("hello" + (("hello" + "a") + "a")) + ("foo bar" + "foo bar")) + "foo bar") + "hello")))))
                    for fi29 := 0; fi29 < 16; fi29++ {
                        var v30 bool = true
                        var v31 int32 = int32(v3)
                        v32 := v6["a"]
                        v33 := GenS3{F0: true}
                        if true {
                            v30 = v30
                            v31 -= 219
                            if v30 {
                                fmt.Println(v2)
                                var v34 uint16 = 0
                                if len(v5) > 0 {
                                    v5[0] = "a"
                                }
                                v35 := v6[""]
                                _ = v34
                                _ = v35
                            } else {
                                v2 = (v2 + v2)
                                v32 = (^(v32))
                            }
                        } else {
                            if v30 {
                                _ = len(v5)
                                var v36 bool = (255 != 65535)
                                v31 += v31
                                v31 = (v31 ^ (-(v31)))
                                _ = v36
                            }
                            if v30 {
                                v37 := GenS2{F0: (!false)}
                                _ = v37
                            } else {
                                v3 = (v3 + (^(v3)))
                                v38 := GenS2{F0: ((-2.25 * 0.1) <= -2.25)}
                                var v39 float32 = (((((-2.25 * -2.25) + 0.0) - 3.14159) + 0.0) - 0.1)
                                _ = v38
                                _ = v39
                            }
                        }
                        _ = fi29
                        _ = v30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                }
                _ = v26
            } else {
                fmt.Println(v2, v3, v3)
            }
            _ = len(v4)
        } else {
            var v40 float64 = 0.1
            var v41 float64 = (v40 * 0.1)
            if false {
                v5 = append(v5, "")
                v2 += (v2 * 4294967295)
                fmt.Println(fi24, v41)
            } else {
                var v42 float64 = (-2.25 + float64(3.14159))
                var v43 float64 = v41
                v44 := GenS4{F0: (v3 | v3)}
                _ = v42
                _ = v43
                _ = v44
            }
            _ = v40
            _ = v41
        }
        v5 = append(v5, "hello")
        v2 = v2
        _ = fi24
    }
    v2 = (v2 ^ (v2 ^ v2))
    fmt.Println(v3, v3)
    switch "hello" {
    case "foo bar":
        v45 := GenS5{F0: 65535, F1: 42, F2: 65535}
        _ = v45
    case "hello":
        v3 += (v3 & v3)
    case "":
        var v46 int32 = -1
        _ = v46
    default:
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    }
    switch 4294967295 {
    case -1:
        v47 := GenS3{F0: true}
        _ = v47
    default:
        var v48 uint = 850
        _ = v48
    }
    v2 = (v2 | v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
