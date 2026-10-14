package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 string; F2 int8 }
type GenS5 struct { F0 int32; F1 string; F2 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 bool; F2 int8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS11 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int16, p1 uint, p2 string) int {
    v7 := GenS1{F0: uint64(p1)}
    p1 = (p1 - p1)
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 int32 = 108
    var v2 uint = 42
    v3 := []string{"foo bar", "", "foo bar", ""}
    var v4 int16 = 0
    v5 := []int{930, 9223372036854775807}
    var v6 int32 = v1
    v8 := GenS2{F0: 65535, F1: 226, F2: 1}
    switch 4294967295 {
    case 1:
        for fi9 := 0; fi9 < 12; fi9++ {
            if true {
                for fi10 := 0; fi10 < 3; fi10++ {
                    var v11 uint16 = 594
                    v1 = (v1 * 0)
                    fmt.Println(v1, fi10, v4)
                    _ = fi10
                    _ = v11
                }
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                if false {
                    v4 -= (v4 - v4)
                    v13 := GenS3{F0: fi9}
                    v6 = v6
                    _ = v13
                }
                v14 := GenS4{F0: fi9, F1: "foo bar", F2: 0}
                fmt.Println(fi12, v1)
                _ = fi12
                _ = v14
            }
            _ = fi9
        }
    case 0:
        var v15 uint = 255
        _ = v15
    case -9223372036854775808:
        v16 := GenS3{F0: -9223372036854775808}
        _ = v16
    }
    var v17 string = "foo bar"
    v18 := GenS5{F0: v1, F1: (v17 + "foo bar"), F2: 4294967295}
    var v19 int32 = (v1 * 42)
    var v20 string = ((("foo bar" + v17) + (v17 + v17)) + "hello")
    if len(v5) > 0 {
        v5[0] = 863
    }
    v17 = v17
    v21 := GenS6{F0: -634}
    switch ((v17 + v17) + v20) {
    case "":
        switch "foo bar" {
        case "\n\t":
            v19 = 2147483647
        case "":
            v2 = 593
        case "foo bar":
            v20 = ("" + ("" + "foo bar"))
        default:
            var v22 int32 = v19
            _ = v22
        }
    case "\n\t":
        if ("a" < (v17 + v17)) {
            v20 = "a"
            switch -161 {
            case 0:
                v23 := GenS7{F0: v4, F1: false, F2: 100}
                _ = v23
            default:
                _ = len(v3)
            }
            if (v4 > (v4 + v4)) {
                for fi24 := 0; fi24 < 5; fi24++ {
                    var v25 int = fi24
                    for fi26 := 0; fi26 < 15; fi26++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v27 uint64 = 2
                        _ = len(v3)
                        _ = fi26
                        _ = v27
                    }
                    v2 -= (v2 | (v2 / (((^(v2))) | 1)))
                    _ = fi24
                    _ = v25
                }
                v28 := GenS8{F0: -1}
                if (-2.25 < (0.1 / 1.5)) {
                    if false {
                        v2 = (v2 / ((186) | 1))
                        switch genFn0(-32768, v2, ("hello" + "a")) {
                        case 9223372036854775807:
                            v29 := GenS2{F0: 730, F1: 255, F2: 255}
                            _ = v29
                        default:
                            if (((1e10 - (((-2.25 / 1e10) + 1.5) * (3.14159 + 3.14159))) / -2.25) < (0.0 + 3.14159)) {
                                _ = len(v5)
                            }
                        }
                        v2 = (v2 << (((v2 >> ((v2) & 63))) & 63))
                    } else {
                        if false {
                            fmt.Println(v1, v17, v20)
                            var v30 int = 865
                            for fi31 := 0; fi31 < 12; fi31++ {
                                var v32 int32 = (^(v6))
                                var v33 int64 = 9223372036854775807
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                            _ = v30
                        }
                        var v34 int = 668
                        v4 += (v4 / ((1000) | 1))
                        _ = v34
                    }
                }
                _ = v28
            }
        } else {
            v20 = v20
            var v35 uint32 = 42
            var v36 uint = v2
            _ = v35
            _ = v36
        }
    case "hello":
        for fi37 := 0; fi37 < 9; fi37++ {
            for fi38 := 0; fi38 < 17; fi38++ {
                if len(v3) > 0 {
                    v3[0] = v20
                }
                fmt.Println(v17)
                if false {
                    switch "a" {
                    case "a":
                        v4 += -32768
                    case "foo bar":
                        v17 = v20
                        v39 := GenS9{F0: 435}
                        _ = v39
                    case "hello":
                        v5 = append(v5, fi37)
                    }
                    v3 = append(v3, (v20 + v17))
                } else {
                    v5 = append(v5, fi38)
                    var v40 int64 = 9223372036854775807
                    v41 := GenS10{F0: (v2 | 1), F1: 255, F2: 320}
                    _ = v40
                    _ = v41
                }
                _ = fi38
            }
            v42 := GenS11{F0: 498, F1: 42}
            var v43 uint16 = 2
            _ = fi37
            _ = v42
            _ = v43
        }
    default:
        var v44 uint64 = 18446744073709551615
        _ = v44
    }
    if (!((-2.25 >= 1e10) || (uint8(v2) >= 242))) {
        _ = len(v3)
        v45 := GenS9{F0: 1}
        var v46 uint64 = 117
        _ = v45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
