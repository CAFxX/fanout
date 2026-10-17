package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 uint8; F1 string }
func genFn0() string {
    v5 := []int{1}
    var v6 uint32 = 42
    if false {
        v7 := GenS1{F0: -1, F1: 0}
        v8 := GenS1{F0: 1000, F1: 849}
        _ = v7
        _ = v8
    }
    _ = v5
    _ = v6
    return ("hello" + "\n\t")
}

func main() {
    var v1 uint32 = 892
    var v2 bool = true
    var v3 uint = 65535
    var v4 int32 = -1
    fmt.Println(v3, v1)
    v9 := GenS2{F0: 127, F1: 127}
    v2 = true
    var v10 int64 = 272
    fmt.Println(v2, v3)
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 uint64 = 42
            v4 = (v4 >> ((v3) & 63))
            v3 = (v3 & (v3 * (v3 >> (((v3 & (v3 >> ((v3) & 63)))) & 63))))
            _ = fi12
            _ = v13
        }
        switch fi11 {
        case 42:
            v1 += 4294967295
        case 79:
            var v14 bool = false
            _ = v14
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            if v2 {
                fmt.Println(v1)
                v2 = v2
                fmt.Println(v1)
            } else {
                v10 += (v10 % ((v10) | 1))
            }
            if (1 != v10) {
                v1 -= v1
                switch (fi15 * fi15) {
                case 32768:
                    v3 -= v3
                    v16 := []int{0, (fi11 << (((v3 / ((v3) | 1))) & 63)), (fi15 * 32767)}
                    _ = v16
                case -2147483648:
                    v17 := GenS3{F0: "\n\t", F1: 0}
                    _ = v17
                default:
                    v18 := GenS2{F0: uint8(v3), F1: 2}
                    _ = v18
                }
                v19 := GenS4{F0: 1, F1: -37}
                _ = v19
            }
            fmt.Println(v2, v3)
            switch 4294967295 {
            case 42:
                var v20 uint16 = 53
                _ = v20
            case -12:
                fmt.Println(v3)
                fmt.Println(v2, fi15)
            default:
                switch "foo bar" {
                case "":
                    v3 += 610
                    v21 := []int{127, int(v10)}
                    _ = v21
                case "a":
                    var v22 float32 = -2.25
                    _ = v22
                case "foo bar":
                    v2 = false
                default:
                    v23 := []int{9223372036854775807, fi11, 0, (fi11 | 32767)}
                    _ = v23
                }
                var v24 uint32 = 4294967295
                _ = v24
            }
            _ = fi15
        }
        _ = fi11
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        v2 = v2
        v10 += (v10 & 0)
        _ = fi25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        v27 := []int{fi26, fi26, (fi26 - (fi26 * (fi26 * fi26))), fi26}
        for fi28 := 0; fi28 < 4; fi28++ {
            switch "foo bar" {
            case "":
                if (!v2) {
                    v29 := GenS5{F0: (!true), F1: fi28}
                    if false {
                        switch fi26 {
                        case 2147483647:
                            var v30 string = "\n\t"
                            _ = v30
                        case -9223372036854775808:
                            v31 := GenS3{F0: ("" + ("hello" + (("a" + "a") + ("foo bar" + "hello")))), F1: 232}
                            _ = v31
                        }
                        fmt.Println(v2, fi28, v2)
                        if v2 {
                            var v32 bool = v2
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                            var v33 int64 = 334
                            v3 -= v3
                            _ = v32
                            _ = v33
                        } else {
                            fmt.Println(v10, v4)
                            var v34 uint32 = 2147483647
                            _ = v34
                        }
                    }
                    _ = v29
                }
            case "\n\t":
                v3 = (v3 - v3)
            default:
                fmt.Println(v4, fi28, v1)
            }
            fmt.Println(v10, v3)
            v35 := GenS6{F0: 255, F1: ""}
            _ = fi28
            _ = v35
        }
        var v36 int64 = (v10 / ((v10) | 1))
        _ = fi26
        _ = v27
        _ = v36
    }
    v37 := GenS6{F0: 127, F1: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v9
}
