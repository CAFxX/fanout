package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 int32 }
func genFn0(p0 int, p1 uint64) int8 {
    p1 -= 42
    p1 += p1
    if false {
        var v7 int64 = 519
        var v8 uint16 = 952
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    return 74
}
func genFn1(p0 uint, p1 uint32) bool {
    var v9 int32 = -242
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 bool = true
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 int32 = (-(v9))
        p1 -= p1
        var v14 int64 = -9223372036854775808
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v9
    return true
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"\n\t", "", "", "hello"}
    v3 := []uint32{161, 0}
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    var v6 int32 = -699
    fmt.Println(v6)
    fmt.Println(v6)
    _ = len(v3)
    fmt.Println(v6)
    var v15 int32 = (v6 + (v6 << ((949) & 63)))
    if ((!(v6 > (v15 & -1))) && (!(-2147483648 == 9223372036854775807))) {
        v6 = 2147483647
        var v16 float32 = 3.14159
        _ = v16
    } else {
        _ = len(v5)
        if true {
            v17 := GenS1{F0: 18446744073709551615}
            switch "\n\t" {
            case "":
                switch 128 {
                case 221:
                    v18 := GenS1{F0: 859}
                    _ = v18
                case -2147483648:
                    switch -1 {
                    case 2147483647:
                        v19 := GenS1{F0: 65535}
                        _ = v19
                    case 32767:
                        switch (("\n\t" + "\n\t") + ("a" + "foo bar")) {
                        case "hello":
                            for fi20 := 0; fi20 < 12; fi20++ {
                                var v21 float64 = (1.5 + float64(((3.14159 * float32(float64((0.0 + (0.1 - (float32((0.1 - 0.0)) / 0.1)))))) + 1e10)))
                                var v22 uint32 = 1
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        case "\n\t":
                            fmt.Println(v6, v6)
                        }
                    }
                case -267:
                    v6 = v6
                default:
                    var v23 int8 = -128
                    _ = v23
                }
            }
            switch "\n\t" {
            case "a":
                var v24 uint8 = 128
                _ = v24
            case "foo bar":
                switch ("foo bar" + "a") {
                case "a":
                    switch 1 {
                    case 0:
                        v25 := GenS1{F0: 666}
                        _ = v25
                    case -1:
                        v26 := GenS1{F0: 916}
                        _ = v26
                    case -2147483648:
                        fmt.Println(v15, v6)
                    default:
                        v27 := GenS2{F0: 1000, F1: 128}
                        _ = v27
                    }
                case "\n\t":
                    var v28 uint16 = 255
                    _ = v28
                }
            case "":
                var v29 int16 = -481
                _ = v29
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            _ = v17
        }
        switch "a" {
        case "hello":
            fmt.Println(v15, v6)
        case "foo bar":
            for fi30 := 0; fi30 < 6; fi30++ {
                var v31 string = "a"
                var v32 int = (fi30 & fi30)
                _ = fi30
                _ = v31
                _ = v32
            }
        case "\n\t":
            v33 := GenS3{F0: 9223372036854775807}
            _ = v33
        }
    }
    switch 32768 {
    case -556:
        if true {
            for fi34 := 0; fi34 < 1; fi34++ {
                for fi35 := 0; fi35 < 1; fi35++ {
                    v36 := GenS4{F0: 1, F1: false}
                    v6 += v15
                    switch -2147483648 {
                    case -145:
                        if false {
                            if (0 <= 4294967295) {
                                v15 += v15
                                var v37 uint32 = 0
                                _ = v37
                            }
                            fmt.Println(v15, v6, v15)
                            if (!(480 >= 42)) {
                                v15 += v15
                                fmt.Println(fi35, fi35)
                                v38 := GenS5{F0: fi34, F1: true}
                                _ = v38
                            } else {
                                v6 = v15
                                var v39 bool = false
                                _ = v39
                            }
                            v6 -= v6
                        } else {
                            var v40 int16 = 0
                            _ = v40
                        }
                    case 9223372036854775807:
                        fmt.Println(fi34)
                    }
                    _ = fi35
                    _ = v36
                }
                _ = fi34
            }
        }
    case 0:
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    v41 := GenS3{F0: -606}
    for fi42 := 0; fi42 < 10; fi42++ {
        if false {
            var v43 float32 = 1.5
            _ = v43
        } else {
            for fi44 := 0; fi44 < 16; fi44++ {
                delete(v4, fi44)
                v6 -= v6
                v45 := GenS2{F0: -1, F1: (fi44 & fi42)}
                _ = fi44
                _ = v45
            }
            v6 += (v15 & 2147483647)
        }
        var v46 int16 = 42
        if true {
            for fi47 := 0; fi47 < 1; fi47++ {
                _ = len(v2)
                _ = fi47
            }
            for fi48 := 0; fi48 < 12; fi48++ {
                v49 := GenS6{F0: 1}
                _ = fi48
                _ = v49
            }
            if (!((!((-2.25 - 1.5) > 3.14159)) || (!true))) {
                for fi50 := 0; fi50 < 7; fi50++ {
                    switch "\n\t" {
                    case "":
                        v51 := GenS6{F0: 42}
                        _ = v51
                    case "\n\t":
                        v15 = (v6 - (v6 % ((1) | 1)))
                    }
                    var v52 float32 = (-2.25 * (0.1 - 0.0))
                    _ = fi50
                    _ = v52
                }
                v53 := GenS1{F0: 4294967295}
                var v54 string = (("a" + ("a" + "a")) + "a")
                _ = v53
                _ = v54
            }
        }
        _ = fi42
        _ = v46
    }
    for fi55 := 0; fi55 < 16; fi55++ {
        fmt.Println(v6)
        var v56 bool = true
        delete(v4, (fi55 >> ((2) & 63)))
        _ = fi55
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
