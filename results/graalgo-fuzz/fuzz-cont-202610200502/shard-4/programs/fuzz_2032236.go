package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint32; F1 uint }
func genFn0(p0 uint16) int64 {
    v5 := GenS1{F0: -1}
    v6 := []int{9223372036854775807}
    var v7 uint64 = 4294967295
    v6 = append(v6, 128)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return int64(p0)
}
func genFn1(p0 uint32, p1 bool) uint {
    p1 = (42 >= 0)
    v8 := []int{1, 69, -685, 32767}
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint = 1
        var v11 int64 = 341
        var v12 string = ("a" + ("foo bar" + (("" + "foo bar") + "\n\t")))
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 42, F1: 42}
    _ = p0
    _ = p1
    _ = v13
    _ = v8
    return 1
}

func main() {
    v1 := []string{"", ("hello" + (("\n\t" + "hello") + "hello")), ("\n\t" + ((("hello" + "foo bar") + "hello") + ("hello" + "\n\t"))), "foo bar"}
    var v2 int16 = 325
    var v3 string = ("foo bar" + (((("hello" + "") + "\n\t") + ("" + "hello")) + "hello"))
    var v4 float64 = 1e10
    v4 = v4
    v2 = -32768
    v1 = append(v1, "\n\t")
    v2 = 915
    v4 -= (0.1 + float64(0.0))
    if (-1 != 9223372036854775807) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v14 := GenS2{F0: -609, F1: 42}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v14
    } else {
        if true {
            var v15 int8 = 0
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 float64 = 0.0
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v17 += 1e10
                _ = fi16
                _ = v17
            }
            _ = v15
        }
    }
    v4 -= v4
    var v18 uint64 = 2
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        if false {
            var v20 uint = 255
            var v21 bool = (42 <= 42)
            _ = v20
            _ = v21
        } else {
            var v22 int16 = (v2 / (((v2 << ((619) & 63))) | 1))
            _ = v22
        }
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 float32 = 1e10
            _ = fi23
            _ = v24
        }
        if (!((v2 / ((v2) | 1)) == -573)) {
            if len(v1) > 0 {
                v1[0] = v3
            }
        } else {
            if (-48 < 45) {
                v25 := GenS3{F0: -1, F1: 1}
                if len(v1) > 0 {
                    v1[0] = v3
                }
                if (!true) {
                    v1 = append(v1, v3)
                    var v26 float64 = v4
                    v26 = v4
                    _ = v26
                } else {
                    switch fi19 {
                    case 127:
                        var v27 uint32 = 0
                        _ = v27
                    case -9223372036854775808:
                        for fi28 := 0; fi28 < 19; fi28++ {
                            switch "\n\t" {
                            case "foo bar":
                                _ = len(v1)
                            case "\n\t":
                                fmt.Println(v18, v4)
                            }
                            switch fi19 {
                            case 871:
                                var v29 string = v3
                                _ = v29
                            case 4294967295:
                                var v30 uint64 = v18
                                _ = v30
                            case -1:
                                v31 := GenS1{F0: -9223372036854775808}
                                _ = v31
                            default:
                                switch ("hello" + "") {
                                case "hello":
                                    v4 -= v4
                                }
                            }
                            _ = fi28
                        }
                    case 32768:
                        var v32 string = "a"
                        _ = v32
                    }
                    _ = len(v1)
                }
                _ = v25
            }
            var v33 float64 = (v4 / -2.25)
            if true {
                for fi34 := 0; fi34 < 5; fi34++ {
                    v35 := GenS4{F0: 0, F1: 203}
                    switch (fi34 / ((-898) | 1)) {
                    case 32768:
                        var v36 uint64 = (v18 / ((v18) | 1))
                        _ = v36
                    case 32767:
                        v37 := GenS5{F0: 44, F1: 0}
                        _ = v37
                    case 877:
                        v38 := GenS6{F0: (v2 * -422)}
                        _ = v38
                    }
                    var v39 int = fi34
                    _ = fi34
                    _ = v35
                    _ = v39
                }
                v2 = (v2 + v2)
                if (!true) {
                    v33 += (v4 - (v33 + 1.5))
                } else {
                    for fi40 := 0; fi40 < 18; fi40++ {
                        var v41 uint32 = 0
                        if (!false) {
                            if (true || false) {
                                var v42 uint = 18446744073709551615
                                v41 += (v41 << ((v42) & 63))
                                fmt.Println(v42, v18)
                                _ = v42
                            }
                        }
                        _ = fi40
                        _ = v41
                    }
                    v43 := GenS4{F0: 1, F1: 255}
                    v44 := GenS7{F0: 588, F1: 173}
                    _ = v43
                    _ = v44
                }
            }
            _ = v33
        }
        _ = fi19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
