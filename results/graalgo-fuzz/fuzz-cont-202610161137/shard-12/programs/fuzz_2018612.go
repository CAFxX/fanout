package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int32, p1 uint, p2 int16) uint64 {
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := []int{fi5, (fi5 * fi5)}
        _ = fi5
        _ = v6
    }
    var v7 float32 = ((-2.25 / 0.1) / float32((3.14159 / -2.25)))
    p2 += 720
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 int32, p1 string, p2 int32) uint64 {
    var v8 uint16 = 255
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            switch (fi9 >> ((4294967295) & 63)) {
            case -9223372036854775808:
                switch fi10 {
                case 1:
                    p2 -= 0
                case -306:
                    var v11 int32 = -545
                    _ = v11
                default:
                    switch "" {
                    case "\n\t":
                        var v12 uint32 = 1
                        _ = v12
                    case "foo bar":
                        var v13 int32 = (p2 >> ((4294967295) & 63))
                        _ = v13
                    case "":
                        var v14 string = "\n\t"
                        _ = v14
                    default:
                        v15 := []int{}
                        _ = v15
                    }
                }
            case -309:
                p1 = p1
            default:
                fmt.Println(p2, p1)
            }
            fmt.Println(p2)
            v16 := GenS1{F0: p2, F1: 42}
            v17 := []int{296, (fi9 ^ fi10), (-(fi10)), (fi10 << ((999) & 63))}
            _ = fi10
            _ = v16
            _ = v17
        }
        switch (fi9 | fi9) {
        case -9223372036854775808:
            v18 := []int{466, fi9, fi9}
            _ = v18
        case -1:
            p1 = "\n\t"
        case 32768:
            v8 = v8
        }
        switch fi9 {
        case 760:
            var v19 uint8 = 243
            _ = v19
        case 2147483647:
            v8 -= v8
        case 128:
            v20 := GenS1{F0: -1, F1: 1}
            _ = v20
        }
        _ = fi9
    }
    v21 := GenS2{F0: 2}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v8
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 256
    var v3 float64 = (float64(1e10) + ((3.14159 - 3.14159) + ((0.1 / 0.1) - 0.0)))
    var v4 int = 9223372036854775807
    if true {
        v22 := []int{127, v4}
        var v23 int32 = -1
        for fi24 := 0; fi24 < 20; fi24++ {
            fmt.Println(v4, v23)
            v2 = (v2 & v2)
            if (9223372036854775807 <= 1) {
                var v25 int32 = (v23 << ((255) & 63))
                var v26 uint8 = 42
                _ = v25
                _ = v26
            } else {
                var v27 int = (^(v4))
                for fi28 := 0; fi28 < 1; fi28++ {
                    var v29 uint16 = 65535
                    _ = fi28
                    _ = v29
                }
                v30 := GenS2{F0: (^(v2))}
                _ = v27
                _ = v30
            }
            _ = fi24
        }
        _ = v22
        _ = v23
    } else {
        var v31 int = v4
        _ = v31
    }
    switch 541 {
    case -9223372036854775808:
        var v32 float32 = (float32(v3) - -2.25)
        _ = v32
    case -721:
        if false {
            if ((128 != 20) && (42 >= 9223372036854775807)) {
                v4 = v4
            } else {
                for fi33 := 0; fi33 < 12; fi33++ {
                    for fi34 := 0; fi34 < 11; fi34++ {
                        v35 := GenS1{F0: 42, F1: 42}
                        _ = fi34
                        _ = v35
                    }
                    switch "a" {
                    case "":
                        v4 = v4
                    case "\n\t":
                        for fi36 := 0; fi36 < 7; fi36++ {
                            v37 := []int{}
                            v4 = (fi36 ^ (v4 ^ -641))
                            _ = fi36
                            _ = v37
                        }
                    case "foo bar":
                        for fi38 := 0; fi38 < 1; fi38++ {
                            v39 := GenS3{F0: 0}
                            if false {
                                v2 -= (v2 ^ (v2 % ((v2) | 1)))
                                v40 := []int{-541, 32767}
                                _ = v40
                            }
                            _ = fi38
                            _ = v39
                        }
                    default:
                        var v41 string = (("a" + "") + ("a" + "\n\t"))
                        _ = v41
                    }
                    _ = fi33
                }
                v42 := GenS4{F0: 929}
                var v43 bool = true
                _ = v42
                _ = v43
            }
            v44 := []int{(^(v4)), v4, (v4 & v4)}
            _ = v44
        } else {
            v4 += 0
            var v45 uint16 = uint16(v2)
            v46 := GenS5{F0: -1, F1: true}
            _ = v45
            _ = v46
        }
    case -886:
        if (1 >= -1) {
            v47 := GenS6{F0: v2, F1: 42}
            v4 = v4
            v48 := []int{(v4 + -1), (-(v4))}
            _ = v47
            _ = v48
        }
    }
    var v49 uint32 = v2
    var v50 int8 = 42
    fmt.Println(v50, v49)
    if false {
        fmt.Println(v3, v3)
        var v51 uint16 = 317
        var v52 int8 = -128
        _ = v51
        _ = v52
    }
    _ = len(v1)
    var v53 uint16 = 32767
    var v54 int8 = v50
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v50
    _ = v53
    _ = v54
}
