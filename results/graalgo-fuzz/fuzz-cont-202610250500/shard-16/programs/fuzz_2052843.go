package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int8 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
func genFn0() int {
    var v4 int8 = 127
    var v5 uint64 = 4294967295
    var v6 int = 32768
    _ = v4
    _ = v5
    _ = v6
    return (v6 * v6)
}
func genFn1(p0 int16, p1 int16, p2 uint8) uint16 {
    if false {
        p2 -= (p2 * 255)
        switch "" {
        case "\n\t":
            v7 := []int{127}
            _ = v7
        case "hello":
            for fi8 := 0; fi8 < 14; fi8++ {
                if false {
                    switch ("hello" + "") {
                    case "\n\t":
                        p0 = p0
                    }
                    v9 := []int{(fi8 << ((770) & 63))}
                    p0 = int16(fi8)
                    _ = v9
                }
                var v10 uint32 = 945
                _ = fi8
                _ = v10
            }
        }
    }
    p2 = p2
    v11 := []int{9223372036854775807, 214}
    for fi12 := 0; fi12 < 6; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            switch ("\n\t" + ("a" + ("a" + ""))) {
            case "":
                if false {
                    v14 := GenS1{F0: 100, F1: 484, F2: 103}
                    _ = v14
                }
            }
            if true {
                var v15 bool = false
                var v16 uint8 = (p2 % ((p2) | 1))
                fmt.Println(p0, p0, p1)
                _ = v15
                _ = v16
            }
            _ = fi13
        }
        _ = fi12
    }
    var v17 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v17
    return 281
}

func main() {
    var v1 uint = 655
    var v2 uint64 = 4294967295
    var v3 int32 = 207
    if (!(!(true && (v3 >= 2147483647)))) {
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 16; fi19++ {
                v20 := GenS2{F0: 32767, F1: 0}
                v2 -= (v2 / (((^(v2))) | 1))
                var v21 int16 = -32768
                _ = fi19
                _ = v20
                _ = v21
            }
            fmt.Println(v3, v3)
            _ = fi18
        }
    } else {
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 2; fi23++ {
                if ((!true) || false) {
                    var v24 bool = (genFn0() == 1)
                    _ = v24
                }
                v1 = v1
                v2 += v2
                v25 := []int{2147483647, (fi23 ^ 42), fi22}
                _ = fi23
                _ = v25
            }
            for fi26 := 0; fi26 < 10; fi26++ {
                v2 = v2
                _ = fi26
            }
            _ = fi22
        }
        fmt.Println(v2)
        if true {
            var v27 int = 127
            _ = v27
        } else {
            var v28 string = "hello"
            _ = v28
        }
    }
    if ("a" == ("\n\t" + "\n\t")) {
        var v29 uint8 = 115
        var v30 int = -9223372036854775808
        var v31 uint32 = 2147483647
        _ = v29
        _ = v30
        _ = v31
    } else {
        var v32 bool = (92 >= 2)
        var v33 uint = (v1 & (v1 << ((255) & 63)))
        for fi34 := 0; fi34 < 14; fi34++ {
            for fi35 := 0; fi35 < 10; fi35++ {
                if (1 > -520) {
                    switch "\n\t" {
                    case "hello":
                        v1 += (^(v33))
                    case "\n\t":
                        v1 += (v33 / ((v1) | 1))
                    case "":
                        if true {
                            v36 := []int{-2147483648}
                            var v37 uint16 = 964
                            v33 += (^(v33))
                            _ = v36
                            _ = v37
                        }
                    }
                    v1 = v33
                    var v38 bool = v32
                    v38 = (42 < 0)
                    _ = v38
                }
                v39 := []int{}
                v3 = -2147483648
                _ = fi35
                _ = v39
            }
            if ((445 != 345) || v32) {
                v40 := GenS3{F0: "hello"}
                v41 := []int{genFn0(), 0}
                _ = v40
                _ = v41
            }
            v42 := GenS3{F0: "foo bar"}
            _ = fi34
            _ = v42
        }
        _ = v32
        _ = v33
    }
    switch 32767 {
    case 185:
        v43 := GenS2{F0: -1, F1: 761}
        _ = v43
    case 395:
        fmt.Println(v1, v2)
    case -9223372036854775808:
        var v44 string = "hello"
        _ = v44
    }
    var v45 int32 = 892
    var v46 float32 = 0.1
    for fi47 := 0; fi47 < 12; fi47++ {
        var v48 float64 = 0.1
        switch "a" {
        case "a":
            v49 := []int{fi47, (fi47 - 801), (fi47 + 997), fi47}
            _ = v49
        case "":
            v1 = uint(v2)
        case "hello":
            for fi50 := 0; fi50 < 9; fi50++ {
                fmt.Println(v45, v46)
                _ = fi50
            }
        default:
            var v51 float32 = 0.0
            _ = v51
        }
        _ = fi47
        _ = v48
    }
    if false {
        v52 := GenS4{F0: (v1 | (v1 & (^(v1))))}
        _ = v52
    } else {
        v2 = 84
    }
    v53 := []int{4294967295, -9223372036854775808, 2147483647}
    v54 := GenS2{F0: 42, F1: 2}
    fmt.Println(v2, v3)
    v55 := GenS2{F0: 1, F1: 4294967295}
    v1 -= (v1 ^ v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v53:", len(v53), v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v53
    _ = v54
    _ = v55
}
