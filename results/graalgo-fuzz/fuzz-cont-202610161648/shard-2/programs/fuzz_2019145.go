package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint32; F1 int }
func genFn0(p0 int64) uint {
    for fi4 := 0; fi4 < 2; fi4++ {
        if true {
            for fi5 := 0; fi5 < 3; fi5++ {
                v6 := GenS1{F0: ("" + "hello")}
                var v7 uint64 = 226
                _ = fi5
                _ = v6
                _ = v7
            }
        } else {
            switch ("hello" + ("" + "")) {
            case "\n\t":
                for fi8 := 0; fi8 < 6; fi8++ {
                    fmt.Println(p0, p0)
                    v9 := GenS2{F0: 2, F1: 32767}
                    _ = fi8
                    _ = v9
                }
            case "a":
                v10 := GenS2{F0: 42, F1: 42}
                _ = v10
            default:
                switch "foo bar" {
                case "hello":
                    var v11 float64 = 1.5
                    _ = v11
                default:
                    for fi12 := 0; fi12 < 18; fi12++ {
                        v13 := GenS3{F0: 394, F1: 9223372036854775807}
                        v14 := GenS2{F0: 2, F1: 613}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
            }
            for fi15 := 0; fi15 < 4; fi15++ {
                switch ("a" + "\n\t") {
                case "a":
                    if false {
                        v16 := []int{fi4, (fi15 | fi4)}
                        var v17 int32 = 0
                        var v18 float64 = 0.0
                        var v19 string = ("" + (("\n\t" + "foo bar") + "a"))
                        fmt.Println(fi4)
                        v17 = v17
                        _ = v16
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                case "hello":
                    switch -280 {
                    case 1:
                        p0 -= -451
                    }
                case "\n\t":
                    if (-745 != 0) {
                        switch "foo bar" {
                        case "hello":
                            v20 := []int{42, (fi4 - (fi4 - -9223372036854775808))}
                            _ = v20
                        }
                        p0 -= p0
                        if (true && false) {
                            p0 -= p0
                            v21 := []int{4294967295, fi4, fi4}
                            _ = v21
                        }
                    } else {
                        v22 := []int{}
                        _ = v22
                    }
                }
                _ = fi15
            }
            v23 := []int{(fi4 ^ fi4), int(p0), -1}
            _ = v23
        }
        _ = fi4
    }
    fmt.Println(p0)
    _ = p0
    return 65535
}

func main() {
    var v1 bool = (59 > 1)
    var v2 int16 = 488
    var v3 int8 = -4
    var v24 uint64 = 1
    v3 += v3
    v25 := []int{-9223372036854775808, 975, 4294967295, 128}
    v3 -= -1
    if false {
        v1 = v1
        v2 -= 32767
    } else {
        for fi26 := 0; fi26 < 8; fi26++ {
            v25 = append(v25, fi26)
            _ = fi26
        }
        if true {
            v25 = append(v25, 9223372036854775807)
        } else {
            v2 = v2
            if v1 {
                v25 = append(v25, 127)
                var v27 int32 = 841
                var v28 float32 = 3.14159
                _ = v27
                _ = v28
            } else {
                for fi29 := 0; fi29 < 17; fi29++ {
                    fmt.Println(v1, v2, v2)
                    v2 -= (v2 << ((0) & 63))
                    v30 := GenS3{F0: 2, F1: (^(fi29))}
                    _ = fi29
                    _ = v30
                }
                for fi31 := 0; fi31 < 17; fi31++ {
                    if (!v1) {
                        v24 = 2
                        var v32 uint32 = 0
                        _ = v32
                    } else {
                        v3 = (v3 >> ((18446744073709551615) & 63))
                        for fi33 := 0; fi33 < 5; fi33++ {
                            _ = len(v25)
                            v34 := GenS2{F0: 786, F1: (v2 >> ((917) & 63))}
                            _ = fi33
                            _ = v34
                        }
                    }
                    _ = fi31
                }
            }
        }
    }
    v35 := GenS2{F0: 255, F1: v2}
    var v36 int8 = v3
    v2 = 0
    v24 += v24
    v36 = (-(v36))
    v37 := GenS1{F0: "hello"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v36
    _ = v37
}
