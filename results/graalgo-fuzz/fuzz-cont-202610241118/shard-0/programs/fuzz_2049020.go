package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int32; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int; F2 string }
func genFn0() int32 {
    var v4 int = -115
    if (1e10 != 0.1) {
        if false {
            var v5 bool = false
            v6 := []int{v4, v4, (v4 << ((596) & 63)), v4}
            if v5 {
                fmt.Println(v4, v5)
                v7 := GenS1{F0: 299, F1: (v4 ^ -616)}
                switch "a" {
                case "foo bar":
                    for fi8 := 0; fi8 < 4; fi8++ {
                        var v9 float32 = float32(1.5)
                        _ = fi8
                        _ = v9
                    }
                case "a":
                    v10 := GenS2{F0: 42, F1: 2147483648, F2: 139}
                    _ = v10
                }
                _ = v7
            } else {
                if (32767 == 2) {
                    v4 = -9223372036854775808
                }
                fmt.Println(v4)
                fmt.Println(v4, v5)
            }
            _ = v5
            _ = v6
        }
    } else {
        switch v4 {
        case 0:
            var v11 int8 = -1
            _ = v11
        }
        for fi12 := 0; fi12 < 16; fi12++ {
            v13 := GenS3{F0: 0, F1: 18446744073709551615}
            if (false && ((fi12 - fi12) > fi12)) {
                if ((42 > 187) || (v4 != (-(v4)))) {
                    var v14 float32 = 0.0
                    if true {
                        v15 := []int{}
                        v16 := GenS2{F0: -928, F1: 283, F2: 42}
                        _ = v15
                        _ = v16
                    }
                    v14 = 0.1
                    _ = v14
                }
                v17 := GenS3{F0: 584, F1: 4294967295}
                if false {
                    v18 := []int{v4}
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v4 += fi19
                        _ = fi19
                    }
                    v4 = (^(v4))
                    _ = v18
                } else {
                    fmt.Println(fi12)
                    for fi20 := 0; fi20 < 17; fi20++ {
                        v21 := []int{}
                        var v22 int8 = -68
                        _ = len(v21)
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    fmt.Println(fi12)
                }
                _ = v17
            }
            _ = fi12
            _ = v13
        }
    }
    _ = v4
    return 1
}

func main() {
    var v1 int32 = 42
    var v2 string = "\n\t"
    var v3 int16 = -334
    v3 = (^(v3))
    v23 := GenS1{F0: 272, F1: 2147483647}
    if (true || (4294967295 <= 255)) {
        switch 9223372036854775807 {
        case 32767:
            v24 := []int{127, 995, 875}
            _ = v24
        case -692:
            for fi25 := 0; fi25 < 20; fi25++ {
                fmt.Println(fi25)
                _ = fi25
            }
        default:
            v26 := []int{1}
            _ = v26
        }
        v1 = v1
        if true {
            var v27 float64 = ((1.5 + 3.14159) - 1e10)
            v2 = v2
            for fi28 := 0; fi28 < 6; fi28++ {
                if false {
                    v29 := GenS1{F0: 42, F1: (fi28 & fi28)}
                    _ = v29
                }
                v30 := []int{42, (fi28 % ((fi28) | 1)), fi28}
                _ = fi28
                _ = v30
            }
            _ = v27
        } else {
            var v31 uint32 = 0
            if true {
                for fi32 := 0; fi32 < 4; fi32++ {
                    if false {
                        v1 = (v1 | v1)
                        v2 = v2
                    } else {
                        v31 += (v31 << ((2) & 63))
                        switch v2 {
                        case "\n\t":
                            v33 := []int{fi32, (fi32 % (((fi32 >> ((1) & 63))) | 1)), fi32}
                            _ = v33
                        case "hello":
                            v1 -= 2147483647
                        }
                    }
                    v34 := GenS3{F0: 32767, F1: 1}
                    _ = fi32
                    _ = v34
                }
                v35 := []int{-1}
                _ = v35
            } else {
                fmt.Println(v3)
            }
            fmt.Println(v2, v1, v3)
            _ = v31
        }
    } else {
        if false {
            v36 := GenS1{F0: 18446744073709551615, F1: 704}
            v37 := GenS3{F0: 255, F1: 0}
            var v38 uint = 0
            _ = v36
            _ = v37
            _ = v38
        }
        v1 -= genFn0()
        switch int(v1) {
        case -1:
            var v39 int = 9223372036854775807
            _ = v39
        default:
            v1 = 1
        }
    }
    for fi40 := 0; fi40 < 4; fi40++ {
        var v41 int8 = 0
        var v42 uint8 = 1
        _ = fi40
        _ = v41
        _ = v42
    }
    var v43 int = 128
    v44 := GenS4{F0: 150, F1: (^(v43)), F2: v2}
    v45 := GenS1{F0: 2, F1: (v43 >> ((115) & 63))}
    var v46 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
