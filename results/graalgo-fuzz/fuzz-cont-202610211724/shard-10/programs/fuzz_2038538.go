package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint32 }
func genFn0() bool {
    v7 := []int{-731}
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 float64 = -2.25
        var v10 uint = 65535
        _ = fi8
        _ = v10
        _ = v9
    }
    switch "foo bar" {
    case "a":
        var v11 string = "hello"
        _ = v11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        switch (("foo bar" + "hello") + ("" + ("a" + "hello"))) {
        case "":
            if (true || (4294967295 < 4294967295)) {
                var v13 uint64 = 913
                _ = v13
            } else {
                for fi14 := 0; fi14 < 3; fi14++ {
                    if false {
                        v15 := GenS1{F0: 959, F1: (fi14 - -9223372036854775808), F2: 75}
                        fmt.Println(fi14, fi12)
                        if true {
                            var v16 int16 = -32768
                            _ = v16
                        } else {
                            if len(v7) > 0 {
                                v7[0] = (^(fi12))
                            }
                        }
                        fmt.Println(fi12, fi14)
                        var v17 uint32 = 2147483648
                        v17 = (v17 + v17)
                        _ = v15
                        _ = v17
                    }
                    _ = fi14
                }
                if false {
                    fmt.Println(fi12)
                    fmt.Println(fi12)
                    var v18 int8 = 127
                    _ = v18
                }
                var v19 uint = 18446744073709551615
                _ = v19
            }
        default:
            for fi20 := 0; fi20 < 12; fi20++ {
                var v21 int32 = -1
                if ((!(-128 <= -128)) && false) {
                    var v22 string = ""
                    v7 = append(v7, fi20)
                    _ = v22
                }
                _ = fi20
                _ = v21
            }
        }
        if (268 == 1) {
            var v23 uint16 = 125
            fmt.Println(v23, v23)
            _ = v23
        } else {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        _ = fi12
    }
    _ = v7
    return (206 >= 2)
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint64 = (v1 << ((756) & 63))
    var v3 uint8 = 255
    var v4 uint64 = (v1 >> ((173) & 63))
    v5 := []uint32{}
    var v6 int32 = 1
    v2 = v1
    fmt.Println(v2, v1, v1)
    v6 += (^(v6))
    var v24 int8 = 23
    v25 := GenS1{F0: v2, F1: 9223372036854775807, F2: v3}
    fmt.Println(v2)
    _ = len(v5)
    v3 -= uint8(v4)
    fmt.Println(v3)
    if true {
        fmt.Println(v24, v6)
        v4 += v1
    } else {
        v26 := GenS1{F0: v4, F1: 32768, F2: v3}
        var v27 int8 = -128
        v1 += 674
        _ = v26
        _ = v27
    }
    if false {
        v28 := GenS1{F0: (v4 >> ((622) & 63)), F1: 9223372036854775807, F2: (^(v3))}
        if false {
            if true {
                v29 := GenS2{F0: 81, F1: 0, F2: 515}
                var v30 uint32 = 42
                _ = v29
                _ = v30
            } else {
                for fi31 := 0; fi31 < 10; fi31++ {
                    var v32 float32 = (-2.25 + 3.14159)
                    var v33 int64 = 42
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                v3 = (v3 / ((v3) | 1))
            }
        } else {
            var v34 uint = 0
            v3 = v3
            for fi35 := 0; fi35 < 17; fi35++ {
                switch "foo bar" {
                case "foo bar":
                    switch fi35 {
                    case 32768:
                        var v36 int32 = 628
                        _ = v36
                    case 42:
                        switch fi35 {
                        case 4294967295:
                            v2 -= 0
                        case 899:
                            var v37 uint = v34
                            _ = v37
                        case 32767:
                            v24 = (-(v24))
                        }
                    case 0:
                        v1 = v2
                    }
                case "":
                    var v38 uint = v34
                    _ = v38
                case "a":
                    v3 -= 0
                default:
                    v39 := GenS3{F0: (v34 >> ((0) & 63)), F1: 65535}
                    _ = v39
                }
                if ((!false) || (v34 > v34)) {
                    v40 := GenS1{F0: 18446744073709551615, F1: fi35, F2: 42}
                    var v41 uint64 = 18446744073709551615
                    v2 -= (v1 + (^(v41)))
                    _ = v40
                    _ = v41
                }
                v42 := GenS3{F0: v34, F1: 329}
                _ = fi35
                _ = v42
            }
            _ = v34
        }
        if (v6 != (v6 / ((0) | 1))) {
            fmt.Println(v24, v6, v2)
            v2 = 1
            for fi43 := 0; fi43 < 7; fi43++ {
                var v44 float32 = 0.0
                _ = fi43
                _ = v44
            }
        } else {
            _ = len(v5)
            var v45 float64 = 1e10
            _ = v45
        }
        _ = v28
    } else {
        if (false && (32768 != 986)) {
            v46 := GenS2{F0: 1, F1: 2147483647, F2: v2}
            _ = len(v5)
            _ = v46
        }
    }
    var v47 int32 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
