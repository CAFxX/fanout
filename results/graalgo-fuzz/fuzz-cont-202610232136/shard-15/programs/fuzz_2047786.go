package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 bool; F2 uint32 }
type GenS7 struct { F0 uint16; F1 uint }
func genFn0(p0 bool) int16 {
    v7 := GenS1{F0: 2}
    v8 := []int{9223372036854775807, 1}
    _ = p0
    _ = v7
    _ = v8
    return 877
}
func genFn1() int32 {
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint16 = 636
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := GenS1{F0: v10}
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    var v13 string = ("a" + "foo bar")
    v13 = v13
    v13 = v13
    v13 = ((v13 + v13) + v13)
    _ = v13
    return -108
}

func main() {
    var v1 int32 = 1
    var v2 int32 = (v1 + (v1 - v1))
    v3 := []uint32{2, 0, 582}
    var v4 int16 = -685
    v5 := make(map[int]int)
    v6 := make(map[uint32]int)
    switch -9223372036854775808 {
    case 32768:
        fmt.Println(v1, v1, v2)
    default:
        var v14 float32 = 1.5
        _ = v14
    }
    var v15 int = 1
    if false {
        if (((42 != 841) && false) || false) {
            for fi16 := 0; fi16 < 20; fi16++ {
                if (true || false) {
                    switch 0 {
                    case -466:
                        v17 := GenS1{F0: 42}
                        _ = v17
                    case 604:
                        v18 := GenS1{F0: 42}
                        _ = v18
                    case 127:
                        fmt.Println(v1, fi16)
                    }
                    switch (((("" + "\n\t") + "a") + ((("\n\t" + "foo bar") + "hello") + "")) + ("\n\t" + "hello")) {
                    case "a":
                        v19 := GenS1{F0: 42}
                        _ = v19
                    case "":
                        if (544 == 42) {
                            v5[-9223372036854775808] = (fi16 >> ((753) & 63))
                            var v20 string = ("foo bar" + "\n\t")
                            for fi21 := 0; fi21 < 1; fi21++ {
                                var v22 int = (fi16 >> ((255) & 63))
                                fmt.Println(fi21, fi16)
                                _ = fi21
                                _ = v22
                            }
                            _ = v20
                        } else {
                            v1 += (v2 * (v1 >> ((2) & 63)))
                        }
                    case "hello":
                        if (!(!true)) {
                            for fi23 := 0; fi23 < 8; fi23++ {
                                var v24 uint32 = 65535
                                fmt.Println(v15)
                                v2 -= (v1 % ((v1) | 1))
                                var v25 int16 = (-(v4))
                                _ = fi23
                                _ = v24
                                _ = v25
                            }
                        }
                    }
                    switch ("" + "a") {
                    case "\n\t":
                        fmt.Println(fi16, fi16, fi16)
                    case "":
                        var v26 uint16 = 921
                        _ = v26
                    case "a":
                        if (!true) {
                            v15 -= (fi16 >> ((1) & 63))
                            var v27 int32 = (v2 << ((306) & 63))
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v27
                        }
                    }
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi28 := 0; fi28 < 6; fi28++ {
                    fmt.Println(fi16, fi28)
                    fmt.Println(v15)
                    v1 = v2
                    _ = fi28
                }
                _ = fi16
            }
            v1 += -2147483648
            v1 -= (v1 / (((v1 >> ((13) & 63))) | 1))
        }
        switch "\n\t" {
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "foo bar":
            switch "hello" {
            case "\n\t":
                v4 = -1
            case "hello":
                v4 += (v4 >> ((900) & 63))
            case "foo bar":
                v29 := GenS2{F0: 1, F1: 2}
                _ = v29
            default:
                if false {
                    v2 = v1
                }
            }
        case "":
            _ = len(v3)
        }
        var v30 uint16 = 42
        _ = v30
    } else {
        fmt.Println(v1)
        v31 := GenS2{F0: 127, F1: 111}
        v32 := GenS3{F0: (v4 - (-(v4))), F1: genFn1()}
        _ = v31
        _ = v32
    }
    switch (("hello" + ("a" + "hello")) + "hello") {
    case "foo bar":
        v4 = v4
    case "":
        for fi33 := 0; fi33 < 2; fi33++ {
            var v34 float64 = 0.0
            for fi35 := 0; fi35 < 13; fi35++ {
                v36 := v6[0]
                v34 += ((v34 + float64(0.0)) / v34)
                _ = fi35
                _ = v36
            }
            var v37 uint8 = 255
            _ = fi33
            _ = v34
            _ = v37
        }
    }
    var v38 int8 = int8(v2)
    if (((v2 % ((-669) | 1)) >= v1) && (4294967295 >= (v15 + v15))) {
        v39 := GenS4{F0: 2}
        switch "hello" {
        case "foo bar":
            v4 += v4
        default:
            v1 = v2
        }
        v15 = (-(v15))
        _ = v39
    }
    switch "a" {
    case "":
        for fi40 := 0; fi40 < 6; fi40++ {
            for fi41 := 0; fi41 < 5; fi41++ {
                v2 = int32(v4)
                _ = fi41
            }
            for fi42 := 0; fi42 < 9; fi42++ {
                fmt.Println(v2, v38, fi40)
                var v43 int8 = (v38 & v38)
                if (1 > 4294967295) {
                    v44 := GenS5{F0: v15}
                    var v45 int8 = 100
                    v46 := GenS6{F0: (v4 >> ((2) & 63)), F1: true, F2: 855}
                    _ = v44
                    _ = v45
                    _ = v46
                }
                v15 = (v15 | (fi42 ^ 32768))
                _ = fi42
                _ = v43
            }
            _ = fi40
        }
    }
    v2 = -1
    v47 := GenS7{F0: 0, F1: 255}
    var v48 uint64 = 0
    var v49 uint = 250
    switch (v15 | (^(v15))) {
    case 492:
        _ = len(v3)
    default:
        fmt.Println(v15, v38)
    }
    v50 := GenS5{F0: 128}
    var v51 float32 = -2.25
    for fi52 := 0; fi52 < 7; fi52++ {
        _ = len(v3)
        fmt.Println(v2, v49)
        var v53 float64 = ((3.14159 + 1.5) + 0.1)
        _ = fi52
        _ = v53
    }
    var v54 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v54
    _ = v6
}
