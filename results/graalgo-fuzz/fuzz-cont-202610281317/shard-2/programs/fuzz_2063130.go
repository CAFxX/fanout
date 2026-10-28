package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 int; F2 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 int64; F2 int8 }
type GenS7 struct { F0 uint64; F1 int16; F2 uint16 }
func genFn0() uint {
    var v4 uint32 = 290
    v5 := []int{2147483647, 32768, 127, 1}
    _ = v4
    _ = v5
    return 494
}
func genFn1() bool {
    var v6 uint32 = 42
    var v7 int64 = 876
    _ = v6
    _ = v7
    return (v7 <= v7)
}

func main() {
    var v1 uint16 = 32767
    var v2 float64 = 1e10
    var v3 uint32 = 65535
    for fi8 := 0; fi8 < 4; fi8++ {
        if true {
            v9 := GenS1{F0: -128, F1: int8(v1)}
            fmt.Println(v3, v3)
            var v10 uint64 = 42
            _ = v10
            _ = v9
        }
        fmt.Println(v1)
        _ = fi8
    }
    switch -9223372036854775808 {
    case -162:
        fmt.Println(v2)
    case -818:
        if false {
            v3 -= (v3 << ((255) & 63))
        }
    }
    v1 -= 32767
    v11 := GenS1{F0: 100, F1: 0}
    var v12 uint8 = 2
    v13 := []int{90, 9223372036854775807, 1, 42}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v14 := GenS1{F0: 42, F1: 42}
    for fi15 := 0; fi15 < 10; fi15++ {
        if true {
            v16 := GenS1{F0: 127, F1: -1}
            var v17 int8 = 91
            _ = v16
            _ = v17
        } else {
            v2 = float64(float32((0.0 * ((v2 * v2) * v2))))
        }
        if false {
            var v18 string = (("\n\t" + ("hello" + "foo bar")) + "foo bar")
            var v19 float32 = ((float32(v2) + 1e10) * (0.1 - 3.14159))
            v20 := GenS2{F0: (v12 ^ 120), F1: -1}
            _ = v18
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: 0, F1: 100}
            v22 := GenS3{F0: fi15}
            if (true || false) {
                var v23 uint64 = 0
                _ = v23
            } else {
                for fi24 := 0; fi24 < 17; fi24++ {
                    var v25 string = "hello"
                    _ = fi24
                    _ = v25
                }
                var v26 uint32 = v3
                _ = v26
            }
            _ = v21
            _ = v22
        }
        _ = fi15
    }
    v27 := GenS4{F0: (-9223372036854775808 != int(v1)), F1: -1, F2: (true || true)}
    var v28 uint32 = (v3 - 65535)
    if (!false) {
        v29 := GenS1{F0: 1, F1: 127}
        v30 := GenS5{F0: 0}
        fmt.Println(v1, v3)
        _ = v29
        _ = v30
    }
    if ((v2 / (float64(1e10) / 0.1)) == (0.0 * 0.0)) {
        var v31 uint = uint(v3)
        if false {
            for fi32 := 0; fi32 < 10; fi32++ {
                v33 := GenS3{F0: 32768}
                for fi34 := 0; fi34 < 13; fi34++ {
                    v35 := GenS6{F0: v1, F1: -574, F2: 99}
                    _ = fi34
                    _ = v35
                }
                for fi36 := 0; fi36 < 13; fi36++ {
                    fmt.Println(v1, v3)
                    _ = fi36
                }
                _ = fi32
                _ = v33
            }
            if false {
                switch "hello" {
                case "":
                    var v37 int = 42
                    _ = v37
                case "\n\t":
                    for fi38 := 0; fi38 < 17; fi38++ {
                        v39 := GenS7{F0: 723, F1: 1000, F2: v1}
                        _ = fi38
                        _ = v39
                    }
                case "a":
                    switch 4294967295 {
                    case -89:
                        if (v1 >= 695) {
                            _ = len(v13)
                        } else {
                            var v40 bool = true
                            v31 -= (v31 & (^(v31)))
                            var v41 int8 = -73
                            _ = v40
                            _ = v41
                        }
                    case -2147483648:
                        for fi42 := 0; fi42 < 12; fi42++ {
                            var v43 uint16 = 255
                            _ = fi42
                            _ = v43
                        }
                    default:
                        v44 := GenS4{F0: ((v3 < (v28 & 42)) && true), F1: 2147483647, F2: true}
                        _ = v44
                    }
                }
                v1 = 2
            } else {
                switch 1 {
                case -2147483648:
                    var v45 uint16 = (v1 & (v1 >> ((genFn0()) & 63)))
                    _ = v45
                case 1:
                    _ = len(v13)
                }
                var v46 int32 = -750
                _ = v46
            }
            for fi47 := 0; fi47 < 9; fi47++ {
                v31 += (v31 & (v31 - v31))
                var v48 uint64 = 42
                v12 = (v12 >> (((^(v31))) & 63))
                _ = fi47
                _ = v48
            }
        }
        var v49 int32 = -546
        _ = v31
        _ = v49
    }
    _ = len(v13)
    if (687 >= 2) {
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
