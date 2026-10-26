package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 bool }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint }
func genFn0(p0 uint32) uint16 {
    p0 = 2147483647
    var v5 bool = (-2.25 >= ((1e10 - 0.0) - float32(((-2.25 - (-2.25 + 0.0)) - 1e10))))
    v6 := GenS1{F0: 42, F1: -2147483648, F2: v5}
    var v7 uint8 = 0
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint16, p1 uint32) int32 {
    if (!true) {
        switch "\n\t" {
        case "a":
            switch 127 {
            case 911:
                var v8 int8 = -1
                _ = v8
            case 2147483647:
                p0 -= (p0 & genFn0(p1))
            default:
                switch "hello" {
                case "hello":
                    v9 := GenS2{F0: 340, F1: 0}
                    _ = v9
                case "foo bar":
                    p0 += (p0 & p0)
                case "":
                    var v10 string = "foo bar"
                    _ = v10
                default:
                    var v11 uint = 233
                    _ = v11
                }
            }
        case "hello":
            if (-753 > 758) {
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 uint = 1
                    if false {
                        for fi14 := 0; fi14 < 5; fi14++ {
                            var v15 float64 = 3.14159
                            var v16 int64 = -626
                            var v17 float64 = (0.0 * v15)
                            fmt.Println(p0)
                            _ = fi14
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        switch fi12 {
                        case -557:
                            for fi18 := 0; fi18 < 14; fi18++ {
                                p0 = (^(p0))
                                p1 = p1
                                var v19 int16 = -827
                                _ = fi18
                                _ = v19
                            }
                        case 127:
                            p0 = 42
                        default:
                            switch "hello" {
                            case "a":
                                var v20 int16 = -1
                                _ = v20
                            }
                        }
                    }
                    p1 -= p1
                    _ = fi12
                    _ = v13
                }
            }
        default:
            var v21 float32 = 0.1
            _ = v21
        }
        for fi22 := 0; fi22 < 13; fi22++ {
            p1 += (p1 & 42)
            p1 = (p1 << ((1) & 63))
            if true {
                switch "\n\t" {
                case "hello":
                    var v23 bool = false
                    _ = v23
                case "foo bar":
                    switch fi22 {
                    case -9223372036854775808:
                        for fi24 := 0; fi24 < 2; fi24++ {
                            p0 = (p0 % ((1) | 1))
                            p1 = p1
                            p0 = genFn0(p1)
                            var v25 int8 = -1
                            var v26 float32 = 1.5
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    case 4294967295:
                        for fi27 := 0; fi27 < 13; fi27++ {
                            switch (("\n\t" + "a") + "hello") {
                            case "":
                                v28 := []int{-639, 127, (-(fi22)), fi22}
                                _ = v28
                            default:
                                switch (fi27 | fi22) {
                                case 32767:
                                    var v29 int64 = -749
                                    _ = v29
                                case 486:
                                    v30 := []int{}
                                    _ = v30
                                }
                            }
                            var v31 int32 = 2147483647
                            _ = fi27
                            _ = v31
                        }
                    default:
                        v32 := GenS3{F0: 1, F1: (^(p0)), F2: -128}
                        _ = v32
                    }
                }
                var v33 uint32 = (p1 ^ (p1 >> ((uint(p0)) & 63)))
                _ = v33
            }
            _ = fi22
        }
    }
    p0 = (p0 / ((42) | 1))
    var v34 int8 = -128
    _ = p0
    _ = p1
    _ = v34
    return 1
}

func main() {
    var v1 string = ("hello" + "\n\t")
    var v2 int32 = -1
    v3 := make(map[int]bool)
    var v4 int64 = 42
    v35 := GenS1{F0: 100, F1: 0, F2: true}
    var v36 uint64 = 2
    v1 = v1
    v1 = "a"
    var v37 uint16 = 65535
    v38 := v3[32767]
    for fi39 := 0; fi39 < 17; fi39++ {
        var v40 int32 = 2147483647
        v38 = v38
        var v41 uint64 = 657
        _ = fi39
        _ = v40
        _ = v41
    }
    v38 = (0.0 > 3.14159)
    if (0 != 2) {
        v42 := GenS4{F0: 105}
        _ = v42
    }
    v43 := []int{}
    v4 += v4
    var v44 string = ("foo bar" + v1)
    fmt.Println(v2)
    v45 := GenS3{F0: 18446744073709551615, F1: (v37 & v37), F2: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
