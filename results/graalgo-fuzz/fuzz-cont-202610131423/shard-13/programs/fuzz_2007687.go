package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 bool }
type GenS2 struct { F0 bool; F1 uint32; F2 int8 }
type GenS3 struct { F0 int }
func genFn0(p0 uint32, p1 int, p2 string) int64 {
    var v7 uint64 = 1
    var v8 string = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -278
}
func genFn1(p0 uint16, p1 int8) string {
    var v9 float64 = 3.14159
    for fi10 := 0; fi10 < 14; fi10++ {
        if (0 <= 65535) {
            if (!(2 >= 776)) {
                v11 := []int{}
                var v12 uint64 = 18446744073709551615
                _ = v11
                _ = v12
            } else {
                v9 = float64((0.1 - 1.5))
                var v13 int8 = 127
                var v14 int = -2147483648
                _ = v13
                _ = v14
            }
            var v15 bool = (false || (!((-1 == 1) || true)))
            _ = v15
        }
        v16 := GenS1{F0: -294, F1: true, F2: ((fi10 - 128) < fi10)}
        _ = fi10
        _ = v16
    }
    v9 = (float64(3.14159) + ((v9 * 3.14159) + v9))
    _ = p0
    _ = p1
    _ = v9
    return "hello"
}

func main() {
    var v1 string = "foo bar"
    v2 := []int{32768, 32768, -1}
    var v3 int32 = 42
    v4 := []int{971, 9223372036854775807, -669}
    var v5 int32 = v3
    var v6 uint8 = 42
    for fi17 := 0; fi17 < 19; fi17++ {
        if (2147483648 > 715) {
            if (622 < 42) {
                v18 := GenS2{F0: true, F1: 2, F2: 0}
                switch "foo bar" {
                case "hello":
                    var v19 uint32 = 0
                    _ = v19
                case "":
                    for fi20 := 0; fi20 < 19; fi20++ {
                        var v21 bool = false
                        switch 107 {
                        case 42:
                            var v22 string = ((v1 + "\n\t") + (v1 + (v1 + v1)))
                            _ = v22
                        }
                        if (42 < 893) {
                            v6 = (v6 | v6)
                            var v23 bool = v21
                            v1 = v1
                            _ = v23
                        } else {
                            v24 := GenS3{F0: (fi17 - (fi20 / ((128) | 1)))}
                            _ = v24
                        }
                        _ = fi20
                        _ = v21
                    }
                case "a":
                    switch "a" {
                    case "\n\t":
                        for fi25 := 0; fi25 < 13; fi25++ {
                            var v26 string = ((("foo bar" + (("\n\t" + "hello") + "\n\t")) + (v1 + (v1 + ""))) + (v1 + (v1 + v1)))
                            _ = fi25
                            _ = v26
                        }
                    case "a":
                        v1 = ("\n\t" + (("hello" + "a") + "hello"))
                    case "foo bar":
                        if false {
                            _ = len(v4)
                        } else {
                            var v27 uint32 = 2147483647
                            v2 = append(v2, fi17)
                            _ = v27
                        }
                    }
                default:
                    v2 = append(v2, (fi17 | -2147483648))
                }
                _ = v18
            }
        } else {
            var v28 float64 = (1.5 + float64((1.5 - 0.1)))
            fmt.Println(v1, v3, v5)
            _ = v28
        }
        _ = len(v4)
        if (-1 >= -9223372036854775808) {
            v29 := GenS3{F0: fi17}
            if true {
                v1 = "a"
                v30 := GenS2{F0: true, F1: 587, F2: 100}
                _ = v30
            }
            var v31 uint8 = 128
            _ = v29
            _ = v31
        }
        _ = fi17
    }
    var v32 uint32 = 2147483648
    fmt.Println(v6, v3, v32)
    if false {
        switch -1 {
        case 128:
            v33 := GenS3{F0: 128}
            _ = v33
        case 32767:
            switch v1 {
            case "foo bar":
                if (1 >= 206) {
                    var v34 string = "a"
                    _ = v34
                } else {
                    fmt.Println(v1, v5)
                }
            case "":
                v6 += 128
            case "a":
                fmt.Println(v5)
            }
        case -110:
            for fi35 := 0; fi35 < 11; fi35++ {
                var v36 int8 = 9
                v2 = append(v2, fi35)
                _ = fi35
                _ = v36
            }
        default:
            var v37 int8 = 1
            _ = v37
        }
    } else {
        v38 := GenS3{F0: 2147483647}
        if false {
            switch 32768 {
            case 585:
                if true {
                    var v39 float32 = 0.1
                    var v40 uint8 = v6
                    var v41 int32 = (v3 / ((0) | 1))
                    _ = v39
                    _ = v40
                    _ = v41
                } else {
                    v4 = append(v4, 9223372036854775807)
                }
            case -1:
                if (428 == 4294967295) {
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                }
            default:
                var v42 int64 = 0
                _ = v42
            }
            for fi43 := 0; fi43 < 20; fi43++ {
                fmt.Println(v5)
                _ = fi43
            }
        }
        switch 128 {
        case 9223372036854775807:
            var v44 uint8 = 118
            _ = v44
        }
        _ = v38
    }
    var v45 string = "hello"
    var v46 int8 = 100
    for fi47 := 0; fi47 < 14; fi47++ {
        v48 := GenS1{F0: -32768, F1: false, F2: true}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi47
        _ = v48
    }
    switch 32767 {
    case -2147483648:
        v3 = (v5 >> ((255) & 63))
    case 32768:
        var v49 int32 = -2147483648
        _ = v49
    }
    if len(v2) > 0 {
        v2[0] = 128
    }
    for fi50 := 0; fi50 < 11; fi50++ {
        v3 += (v3 * (^(v5)))
        var v51 float32 = 1.5
        _ = fi50
        _ = v51
    }
    var v52 string = ""
    var v53 uint64 = 4294967295
    switch "a" {
    case "\n\t":
        var v54 bool = (976 < 65535)
        _ = v54
    case "hello":
        v53 -= (v53 % (((v53 + (v53 << ((2) & 63)))) | 1))
    }
    var v55 float64 = 0.1
    if true {
        v46 -= (v46 >> ((65535) & 63))
    }
    var v56 uint32 = v32
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v52
    _ = v53
    _ = v55
    _ = v56
    _ = v6
}
