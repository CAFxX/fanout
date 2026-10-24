package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint8 }
func genFn0(p0 uint8) int32 {
    p0 = p0
    var v4 uint16 = 0
    switch -9223372036854775808 {
    case -757:
        if false {
            if (v4 > (^(v4))) {
                v5 := []int{}
                if ("hello" != "") {
                    var v6 int8 = 93
                    fmt.Println(p0)
                    _ = v6
                } else {
                    v7 := GenS1{F0: 32768}
                    var v8 bool = (42 != 4294967295)
                    v4 = (v4 << ((uint(v4)) & 63))
                    _ = v7
                    _ = v8
                }
                _ = v5
            }
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 int8 = -128
                if false {
                    v11 := []int{}
                    _ = v11
                } else {
                    var v12 uint64 = 711
                    _ = v12
                }
                _ = fi9
                _ = v10
            }
        } else {
            var v13 uint64 = 0
            var v14 uint8 = (^(p0))
            if false {
                p0 = v14
                var v15 int = -2147483648
                _ = v15
            } else {
                fmt.Println(v14)
                v16 := []int{}
                _ = v16
            }
            _ = v13
            _ = v14
        }
    case -2147483648:
        fmt.Println(v4, v4)
    }
    switch -1 {
    case -9223372036854775808:
        fmt.Println(p0)
    case 1:
        var v17 int32 = 2147483647
        _ = v17
    default:
        for fi18 := 0; fi18 < 20; fi18++ {
            v19 := []int{}
            var v20 int64 = 650
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    _ = p0
    _ = v4
    return 0
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int32 = 2147483647
    v3 := []string{"foo bar", "", ("\n\t" + "foo bar"), "\n\t"}
    if ((v2 | v2) > (v2 | (^(v2)))) {
        v1 -= (^(v1))
        for fi21 := 0; fi21 < 7; fi21++ {
            v1 += v1
            var v22 uint16 = 65535
            if false {
                fmt.Println(v22)
            } else {
                var v23 int = (fi21 * fi21)
                _ = v23
            }
            _ = fi21
            _ = v22
        }
    } else {
        if ((349 < 2) || true) {
            v24 := GenS1{F0: -644}
            var v25 uint64 = 4294967295
            _ = v24
            _ = v25
        } else {
            switch 9223372036854775807 {
            case 32768:
                if ((!(2 <= 18446744073709551615)) || false) {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        v1 -= (v1 | v1)
                        _ = fi26
                    }
                    v27 := GenS1{F0: -2147483648}
                    v3 = append(v3, "a")
                    _ = v27
                }
            case -1:
                var v28 uint16 = 255
                _ = v28
            }
        }
        var v29 int = 32767
        _ = v29
    }
    v1 -= v1
    var v30 bool = (0.0 > (0.1 * 1e10))
    fmt.Println(v1, v2, v30)
    fmt.Println(v30, v30)
    v30 = true
    for fi31 := 0; fi31 < 12; fi31++ {
        for fi32 := 0; fi32 < 16; fi32++ {
            switch fi31 {
            case 1:
                if v30 {
                    v33 := GenS2{F0: v1, F1: 790}
                    _ = len(v3)
                    _ = v33
                } else {
                    v2 = (^(v2))
                    v34 := GenS3{F0: 24, F1: 255, F2: 107}
                    v2 -= v2
                    _ = v34
                }
            }
            v35 := GenS1{F0: fi31}
            _ = fi32
            _ = v35
        }
        _ = fi31
    }
    v2 -= 2147483647
    var v36 int16 = -32768
    var v37 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v37
}
