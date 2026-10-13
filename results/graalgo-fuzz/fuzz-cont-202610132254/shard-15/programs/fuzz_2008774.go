package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint; F2 int }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    if false {
        var v4 uint8 = 255
        _ = v4
    }
    var v5 int8 = 7
    v6 := GenS1{F0: -1}
    fmt.Println(v5)
    if (4294967295 < 847) {
        v5 = v5
        if true {
            fmt.Println(v5)
            v7 := []int{4294967295}
            _ = v7
        }
        v5 = -1
    }
    var v8 uint32 = 65535
    v1[""] = ((("" + "foo bar") + (((("foo bar" + "") + ("\n\t" + ((("foo bar" + "a") + ("\n\t" + ("hello" + "a"))) + ""))) + "foo bar") + "")) + "foo bar")
    v9 := []int{}
    if len(v9) > 0 {
        v9[0] = 788
    }
    if ((v8 >= v8) && true) {
        fmt.Println(v5, v5)
        if true {
            var v10 int64 = 619
            _ = v10
        } else {
            _ = len(v2)
            v1["\n\t"] = "foo bar"
            v11 := GenS2{F0: 42, F1: 65535, F2: 1}
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 15; fi12++ {
            v1[("foo bar" + ("" + "a"))] = "a"
            delete(v1, "foo bar")
            _ = fi12
        }
        v8 -= (v8 >> ((833) & 63))
        if (255 != 42) {
            switch 128 {
            case -746:
                var v13 int64 = 9223372036854775807
                _ = v13
            case -1:
                var v14 string = "a"
                _ = v14
            default:
                if ((uint(v8) != 97) || true) {
                    var v15 int16 = 32767
                    _ = v15
                } else {
                    var v16 int32 = 2147483647
                    v17 := GenS2{F0: 452, F1: 112, F2: -2147483648}
                    _ = v16
                    _ = v17
                }
            }
            v8 = v8
        }
    }
    var v18 uint8 = 128
    if ((((float64(1e10) / 0.1) / ((1.5 - 0.1) / -2.25)) >= 1.5) || (!true)) {
        fmt.Println(v18, v18, v5)
        if false {
            if (true && true) {
                switch "foo bar" {
                case "":
                    var v19 int32 = 42
                    _ = v19
                case "foo bar":
                    v8 = (v8 - v8)
                default:
                    switch int(v8) {
                    case -289:
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    case 0:
                        if (!false) {
                            var v20 int64 = -9223372036854775808
                            v18 -= (v18 % ((v18) | 1))
                            v5 = 120
                            _ = v20
                        }
                    case 32768:
                        v21 := GenS2{F0: 0, F1: 1, F2: -9223372036854775808}
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 18; fi22++ {
                            switch 912 {
                            case 0:
                                if false {
                                    fmt.Println(v18, v8)
                                    var v23 uint64 = 2
                                    _ = v23
                                }
                            case 9223372036854775807:
                                _ = len(v2)
                            case 42:
                                fmt.Println(v8, v5, fi22)
                            }
                            _ = fi22
                        }
                    }
                }
                var v24 int32 = 0
                _ = v24
            }
            v25 := v2[v8]
            _ = v25
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            v8 += 1
            v27 := GenS3{F0: 2147483647}
            if (false || (v5 == (v5 << ((65535) & 63)))) {
                v18 -= v18
            }
            _ = fi26
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 int = fi28
        v5 += (-(v5))
        var v30 uint = 0
        _ = fi28
        _ = v29
        _ = v30
    }
    v18 += v18
    var v31 uint64 = 0
    v32 := GenS3{F0: -2147483648}
    if (1 < 1000) {
        fmt.Println(v18, v18)
        for fi33 := 0; fi33 < 10; fi33++ {
            _ = len(v3)
            _ = fi33
        }
        switch ("hello" + "hello") {
        case "a":
            var v34 uint64 = 483
            _ = v34
        case "\n\t":
            var v35 float64 = 0.0
            _ = v35
        default:
            var v36 float32 = -2.25
            _ = v36
        }
    }
    var v37 int64 = -950
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v37
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
