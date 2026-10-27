package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 string; F2 uint }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int16 }
func genFn0() uint32 {
    for fi7 := 0; fi7 < 12; fi7++ {
        switch 32767 {
        case 0:
            if (1.5 == 3.14159) {
                var v8 uint = 51
                _ = v8
            }
        }
        v9 := []int{}
        var v10 float32 = -2.25
        _ = fi7
        _ = v10
        _ = v9
    }
    var v11 int64 = 0
    v11 -= (v11 << ((4294967295) & 63))
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 int32 = -199
        _ = fi12
        _ = v13
    }
    var v14 int = 127
    _ = v11
    _ = v14
    return 65535
}
func genFn1() string {
    if (137 > 381) {
        var v15 uint64 = 658
        var v16 string = "hello"
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 65535, F1: true, F2: 4294967295}
    _ = v17
    return "a"
}

func main() {
    var v1 uint64 = 2
    v2 := []int{9223372036854775807, 32767}
    v3 := make(map[int]bool)
    v4 := []int{465, 1}
    var v5 string = "hello"
    v6 := make(map[string]bool)
    fmt.Println(v5, v5)
    v18 := GenS1{F0: 255, F1: false, F2: 144}
    for fi19 := 0; fi19 < 8; fi19++ {
        v1 = 835
        v20 := GenS2{F0: true}
        _ = fi19
        _ = v20
    }
    var v21 string = v5
    fmt.Println(v5)
    switch -2147483648 {
    case -7:
        v22 := GenS1{F0: 629, F1: ((!false) || (-1 < 1)), F2: 0}
        _ = v22
    case 4294967295:
        if true {
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 float32 = 1.5
                _ = fi23
                _ = v24
            }
        } else {
            if len(v2) > 0 {
                v2[0] = -2147483648
            }
            fmt.Println(v1, v21)
        }
    case -1:
        var v25 uint = 727
        _ = v25
    }
    var v26 int8 = 1
    var v27 string = (v5 + (v21 + (v5 + (v5 + v5))))
    var v28 uint64 = v1
    var v29 uint16 = 790
    for fi30 := 0; fi30 < 15; fi30++ {
        if (true || (783 < 2147483648)) {
            if (32767 < -32768) {
                v3[fi30] = ((fi30 & (fi30 ^ -9223372036854775808)) > (fi30 & (fi30 | fi30)))
            }
        } else {
            if false {
                var v31 int8 = v26
                _ = len(v6)
                v32 := GenS1{F0: 2, F1: false, F2: 18446744073709551615}
                _ = v31
                _ = v32
            } else {
                v33 := GenS2{F0: false}
                _ = v33
            }
            if (true && (v27 < v5)) {
                v26 -= -99
                v3[-571] = true
                v27 = (v27 + v5)
            }
            if true {
                switch 4294967295 {
                case 128:
                    for fi34 := 0; fi34 < 12; fi34++ {
                        _ = len(v4)
                        v5 = "foo bar"
                        if true {
                            v35 := GenS3{F0: (v28 / (((v28 + v1)) | 1)), F1: v27, F2: 65535}
                            _ = v35
                        }
                        fmt.Println(v5, v5, fi34)
                        _ = fi34
                    }
                case 765:
                    for fi36 := 0; fi36 < 11; fi36++ {
                        var v37 int32 = -1
                        _ = fi36
                        _ = v37
                    }
                case -9223372036854775808:
                    if false {
                        fmt.Println(v5)
                        if false {
                            v38 := GenS3{F0: v1, F1: (("\n\t" + v5) + v5), F2: 330}
                            var v39 uint32 = 42
                            _ = v38
                            _ = v39
                        } else {
                            v29 -= 854
                        }
                        v1 += (^(v1))
                    }
                default:
                    fmt.Println(v28)
                }
                v5 = (((v5 + (v27 + "a")) + v21) + "hello")
            }
        }
        var v40 int = fi30
        _ = fi30
        _ = v40
    }
    v1 -= (v28 / ((v1) | 1))
    v41 := GenS1{F0: 0, F1: true, F2: 2}
    if (2 != 127) {
        if (61 == v26) {
            if true {
                v42 := GenS3{F0: (^(v28)), F1: "a", F2: 621}
                for fi43 := 0; fi43 < 6; fi43++ {
                    v44 := GenS1{F0: 42, F1: (!false), F2: 1}
                    _ = fi43
                    _ = v44
                }
                switch 32767 {
                case 1:
                    v1 -= v1
                case 4294967295:
                    v6[v21] = false
                }
                _ = v42
            }
            fmt.Println(v21)
            v1 = (v28 + 1)
        } else {
            var v45 bool = (!(v29 >= (v29 / ((v29) | 1))))
            v1 += uint64(v29)
            _ = v45
        }
    }
    var v46 uint16 = v29
    for fi47 := 0; fi47 < 10; fi47++ {
        _ = len(v2)
        if true {
            v6[v5] = false
            v48 := GenS4{F0: 4294967295, F1: 2}
            _ = v48
        }
        _ = len(v6)
        _ = fi47
    }
    if false {
        v49 := GenS5{F0: 0}
        _ = v49
    } else {
        fmt.Println(v46)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v46
    _ = v5
    _ = v6
}
