package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint64, p1 string) uint64 {
    v7 := []int{-101}
    fmt.Println(p0, p1)
    fmt.Println(p1, p1)
    v8 := GenS1{F0: -73, F1: "hello"}
    v9 := GenS1{F0: 42, F1: ""}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 uint16 = 2
    var v2 string = "foo bar"
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    var v5 string = "a"
    v6 := make(map[int]int)
    v5 = ("foo bar" + v5)
    fmt.Println(v2)
    v5 = (v5 + v2)
    for fi10 := 0; fi10 < 12; fi10++ {
        v2 = "hello"
        _ = fi10
    }
    v1 += (v1 * (^(v1)))
    if ("\n\t" >= "") {
        if (!false) {
            for fi11 := 0; fi11 < 13; fi11++ {
                if (42 != -32768) {
                    for fi12 := 0; fi12 < 13; fi12++ {
                        v13 := GenS1{F0: 42, F1: ("hello" + v2)}
                        var v14 uint64 = 31
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    v15 := []int{fi11}
                    _ = v15
                }
                var v16 uint = 4294967295
                v17 := GenS1{F0: 1, F1: (v5 + v2)}
                _ = fi11
                _ = v16
                _ = v17
            }
        } else {
            v18 := []int{9223372036854775807, -2147483648, 42, 42}
            v2 = "a"
            _ = v18
        }
    } else {
        fmt.Println(v1, v5)
        if false {
            v2 = (v2 + v2)
            var v19 int16 = 1
            _ = v19
        } else {
            v20 := GenS1{F0: 32767, F1: (v2 + "")}
            _ = v20
        }
    }
    var v21 int32 = -128
    for fi22 := 0; fi22 < 3; fi22++ {
        if (367 != 2147483647) {
            v23 := []int{}
            _ = v23
        }
        var v24 int8 = 41
        if false {
            switch fi22 {
            case 850:
                if true {
                    var v25 int64 = -1
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v1 = (^(v1))
                        fmt.Println(fi22, v25)
                        _ = fi26
                    }
                    if false {
                        for fi27 := 0; fi27 < 2; fi27++ {
                            var v28 int64 = (v25 | (v25 * (v25 + -749)))
                            v29 := []int{fi27, fi22}
                            var v30 int = 0
                            _ = fi27
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    } else {
                        v31 := GenS1{F0: fi22, F1: "a"}
                        if (false || true) {
                            v32 := GenS2{F0: (v25 << ((18446744073709551615) & 63)), F1: (v1 * (v1 ^ 0))}
                            for fi33 := 0; fi33 < 18; fi33++ {
                                v2 = v2
                                v25 += v25
                                var v34 float32 = -2.25
                                _ = fi33
                                _ = v34
                            }
                            var v35 uint8 = 127
                            _ = v32
                            _ = v35
                        } else {
                            var v36 int64 = 1
                            v37 := GenS3{F0: v25, F1: (639 < -32768)}
                            _ = v36
                            _ = v37
                        }
                        _ = v31
                    }
                    _ = v25
                } else {
                    var v38 float64 = float64((0.1 - 1e10))
                    var v39 string = v2
                    _ = v38
                    _ = v39
                }
            case 9223372036854775807:
                v40 := GenS3{F0: -1, F1: true}
                _ = v40
            }
            v24 += (v24 / (((v24 ^ 127)) | 1))
            v41 := []int{(fi22 - (-(fi22))), (fi22 / ((fi22) | 1)), fi22, fi22}
            _ = v41
        }
        _ = fi22
        _ = v24
    }
    switch int(v1) {
    case -1:
        if true {
            switch ((((("hello" + ("\n\t" + "\n\t")) + "\n\t") + v2) + (v2 + (v5 + v2))) + "a") {
            case "":
                if (991 == int64(v1)) {
                    fmt.Println(v2, v5)
                    if true {
                        for fi42 := 0; fi42 < 16; fi42++ {
                            v43 := []int{(^(fi42)), fi42, 42, (fi42 ^ fi42)}
                            v43 = append(v43, 127)
                            var v44 string = (v5 + "foo bar")
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        var v45 float32 = 0.0
                        v46 := GenS4{F0: 65535, F1: ((807 != (v1 - 0)) && true)}
                        _ = v45
                        _ = v46
                    } else {
                        var v47 uint8 = 123
                        _ = v47
                    }
                } else {
                    v5 = ((v5 + "hello") + (((((v2 + ((v5 + v2) + "foo bar")) + "") + (("hello" + v2) + v5)) + ("foo bar" + "")) + v2))
                }
            case "\n\t":
                var v48 int8 = -128
                _ = v48
            case "foo bar":
                if (!false) {
                    v49 := []int{-1, 127, -2147483648}
                    v5 = v2
                    var v50 uint16 = (^(v1))
                    _ = v49
                    _ = v50
                }
            }
            for fi51 := 0; fi51 < 13; fi51++ {
                fmt.Println(fi51)
                fmt.Println(v1, v2)
                _ = fi51
            }
        } else {
            for fi52 := 0; fi52 < 17; fi52++ {
                if true {
                    var v53 int64 = -1
                    _ = v53
                }
                _ = fi52
            }
            var v54 bool = (true && true)
            _ = v54
        }
    case 32768:
        for fi55 := 0; fi55 < 20; fi55++ {
            fmt.Println(fi55)
            _ = fi55
        }
    default:
        v56 := []int{0, 1}
        _ = v56
    }
    var v57 int = -2147483648
    v58 := GenS5{F0: 0}
    if false {
        for fi59 := 0; fi59 < 7; fi59++ {
            v2 = v2
            _ = fi59
        }
        var v60 int8 = -128
        _ = v60
    } else {
        v61 := v3[v2]
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
