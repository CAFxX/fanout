package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int8; F1 uint32; F2 uint }
type GenS4 struct { F0 int64; F1 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 uint64 }
type GenS7 struct { F0 uint16; F1 bool }
func genFn0(p0 uint8, p1 uint, p2 uint) uint16 {
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(p1)
        var v6 uint8 = (p0 * 41)
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 4294967295}
    p0 -= (p0 >> ((p2) & 63))
    p2 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 int8 = 23
    var v2 int32 = -2147483648
    var v3 uint32 = 738
    var v4 int = 32767
    switch (-(v4)) {
    case 1:
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(v1, v3)
            _ = fi8
        }
    case 958:
        var v9 uint64 = 42
        _ = v9
    default:
        v10 := []int{}
        _ = v10
    }
    if false {
        switch v4 {
        case -28:
            var v11 int = -689
            _ = v11
        case -1:
            v12 := GenS2{F0: (^(v4)), F1: true}
            _ = v12
        default:
            v13 := GenS3{F0: 42, F1: 66, F2: 368}
            _ = v13
        }
    }
    var v14 uint8 = 255
    if false {
        var v15 int8 = (v1 << ((439) & 63))
        v16 := []int{(v4 >> ((uint(v3)) & 63)), v4, 4294967295, v4}
        switch v4 {
        case -9223372036854775808:
            if false {
                fmt.Println(v3)
            } else {
                if ((true || (int8(v14) != v15)) && false) {
                    for fi17 := 0; fi17 < 5; fi17++ {
                        fmt.Println(v2, v15)
                        v18 := GenS4{F0: 117, F1: (^(v2))}
                        switch fi17 {
                        case 128:
                            var v19 int8 = v15
                            _ = v19
                        case -2147483648:
                            _ = len(v16)
                        }
                        _ = fi17
                        _ = v18
                    }
                } else {
                    v20 := GenS1{F0: v4}
                    v21 := GenS5{F0: 1}
                    for fi22 := 0; fi22 < 7; fi22++ {
                        var v23 uint16 = 2
                        var v24 uint32 = 875
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = v20
                    _ = v21
                }
                var v25 string = (("" + "a") + ("" + "foo bar"))
                if true {
                    var v26 bool = false
                    _ = v26
                } else {
                    v1 += v1
                    v25 = v25
                    var v27 uint = 0
                    _ = v27
                }
                _ = v25
            }
        default:
            for fi28 := 0; fi28 < 1; fi28++ {
                if ((!(4294967295 > 181)) || ((v14 * 255) >= v14)) {
                    if len(v16) > 0 {
                        v16[0] = 32767
                    }
                    v15 = (v15 / ((-1) | 1))
                    v29 := GenS4{F0: -456, F1: v2}
                    _ = v29
                } else {
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    switch ("a" + "foo bar") {
                    case "a":
                        fmt.Println(v14)
                    case "foo bar":
                        switch (("hello" + ("" + "foo bar")) + "\n\t") {
                        case "\n\t":
                            v16 = append(v16, 42)
                        case "hello":
                            v3 = (v3 << ((uint(v14)) & 63))
                        }
                    }
                }
                v1 = (v1 << ((18446744073709551615) & 63))
                _ = fi28
            }
        }
        _ = v15
        _ = v16
    } else {
        for fi30 := 0; fi30 < 19; fi30++ {
            v31 := GenS4{F0: -1, F1: (^(v2))}
            _ = fi30
            _ = v31
        }
        if false {
            v32 := []int{}
            v33 := GenS6{F0: 255, F1: 4294967295}
            switch "a" {
            case "\n\t":
                fmt.Println(v3, v14, v1)
            default:
                var v34 uint64 = 477
                _ = v34
            }
            _ = v32
            _ = v33
        } else {
            v4 = v4
        }
    }
    if (139 >= 18446744073709551615) {
        v1 -= -102
    }
    v2 = 0
    for fi35 := 0; fi35 < 2; fi35++ {
        v2 = (v2 ^ (v2 * 2147483647))
        fmt.Println(v3, fi35, v4)
        fmt.Println(fi35)
        _ = fi35
    }
    var v36 uint32 = (v3 - v3)
    var v37 bool = (1 == 4294967295)
    var v38 uint64 = 42
    v1 += 100
    v39 := GenS7{F0: 267, F1: false}
    v3 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
