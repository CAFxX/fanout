package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint64 }
func genFn0(p0 int8, p1 int8, p2 int8) int64 {
    v4 := []int{-1, -462, 32767}
    _ = len(v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 uint32, p1 string) int8 {
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 uint64 = 2
        _ = fi5
        _ = v6
    }
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 uint64 = 42
    var v2 int = 4294967295
    var v3 uint32 = 2147483648
    if ("" <= ("" + (("" + "hello") + ("hello" + "")))) {
        for fi8 := 0; fi8 < 7; fi8++ {
            v1 -= v1
            _ = fi8
        }
        if true {
            switch 9223372036854775807 {
            case -2147483648:
                v1 = (v1 - v1)
            default:
                v2 = 127
            }
        }
        var v9 uint64 = (v1 / (((v1 << ((756) & 63))) | 1))
        _ = v9
    }
    v10 := GenS1{F0: v3}
    var v11 int8 = 42
    v2 = v2
    if false {
        var v12 uint = 65535
        v2 = (v2 | 2147483647)
        _ = v12
    } else {
        v1 -= (v1 << ((585) & 63))
    }
    v13 := GenS1{F0: (v3 * (v3 / (((v3 % ((v3) | 1))) | 1)))}
    for fi14 := 0; fi14 < 6; fi14++ {
        v3 += 567
        if (2 > 128) {
            v15 := []int{fi14}
            var v16 float32 = 3.14159
            for fi17 := 0; fi17 < 7; fi17++ {
                v18 := GenS1{F0: (^(v3))}
                v16 -= (float32(1.5) * v16)
                switch "foo bar" {
                case "foo bar":
                    v3 = 387
                case "hello":
                    if false {
                        if len(v15) >= 2 {
                            v15 = v15[1:len(v15)]
                        }
                        var v19 uint = 65535
                        _ = v19
                    }
                }
                _ = fi17
                _ = v18
            }
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        v11 += v11
        for fi21 := 0; fi21 < 10; fi21++ {
            switch "hello" {
            case "\n\t":
                switch (fi21 << ((663) & 63)) {
                case -2147483648:
                    var v22 int = fi21
                    _ = v22
                case -511:
                    var v23 uint = 42
                    v24 := []int{9223372036854775807, (fi20 | v2)}
                    _ = v23
                    _ = v24
                default:
                    if true {
                        var v25 uint16 = 425
                        _ = v25
                    } else {
                        var v26 bool = true
                        _ = v26
                    }
                }
            case "hello":
                if false {
                    switch (fi21 & 127) {
                    case 1:
                        switch fi21 {
                        case 2147483647:
                            var v27 uint8 = 255
                            _ = v27
                        case 42:
                            switch "foo bar" {
                            case "":
                                if false {
                                    var v28 string = ("hello" + "a")
                                    v3 -= 4294967295
                                    v29 := []int{(fi20 * 156), (v2 + int(v1)), (fi21 | (fi21 + v2))}
                                    v29 = append(v29, (fi21 >> ((18446744073709551615) & 63)))
                                    v30 := GenS1{F0: v3}
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                                v2 += (fi20 >> ((2) & 63))
                                v31 := GenS2{F0: v1, F1: false}
                                _ = v31
                            default:
                                if true {
                                    var v32 bool = (32767 != 712)
                                    var v33 int = fi21
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v34 float32 = -2.25
                                    v3 -= (v3 - (v3 ^ (v3 ^ (v3 ^ 2147483648))))
                                    _ = v34
                                }
                            }
                        }
                    case 732:
                        v35 := GenS3{F0: v2}
                        _ = v35
                    case -9223372036854775808:
                        v1 -= v1
                    }
                } else {
                    var v36 uint8 = 255
                    v37 := []int{int(v3), -521, (fi21 ^ (fi20 % ((fi21) | 1))), v2}
                    _ = v36
                    _ = v37
                }
            case "foo bar":
                if false {
                    var v38 int8 = v11
                    v39 := GenS3{F0: 1}
                    var v40 uint64 = v1
                    v41 := []int{}
                    _ = v38
                    _ = v39
                    _ = v40
                    _ = v41
                } else {
                    var v42 uint = 42
                    var v43 uint16 = 0
                    v44 := []int{}
                    _ = v42
                    _ = v43
                    _ = v44
                }
            }
            _ = fi21
        }
        if (!((!(v3 > v3)) && false)) {
            fmt.Println(v2, v11)
            v45 := []int{(fi20 + (fi20 - v2)), 4294967295}
            _ = v45
        } else {
            if (v3 > (v3 | v3)) {
                v46 := GenS2{F0: (v1 + v1), F1: (1 < 1)}
                switch "hello" {
                case "\n\t":
                    var v47 int8 = -128
                    _ = v47
                }
                fmt.Println(v1, v11, fi20)
                _ = v46
            }
        }
        _ = fi20
    }
    v48 := []int{v2, 127, v2, v2}
    _ = len(v48)
    var v49 bool = true
    for fi50 := 0; fi50 < 9; fi50++ {
        switch ("a" + ("" + "\n\t")) {
        case "\n\t":
            var v51 uint8 = 187
            _ = v51
        case "foo bar":
            for fi52 := 0; fi52 < 18; fi52++ {
                var v53 int8 = v11
                _ = fi52
                _ = v53
            }
        }
        _ = fi50
    }
    var v54 int16 = -401
    v55 := GenS4{F0: -1, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v54
    _ = v55
}
