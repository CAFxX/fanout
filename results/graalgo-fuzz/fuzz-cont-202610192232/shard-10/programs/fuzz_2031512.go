package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int8; F1 uint; F2 bool }
type GenS4 struct { F0 string; F1 uint }

func main() {
    var v1 int32 = 42
    var v2 uint32 = 569
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    var v5 bool = false
    delete(v4, 761)
    var v6 string = ""
    if false {
        delete(v3, 42)
    } else {
        switch v6 {
        case "hello":
            var v7 int8 = -98
            _ = v7
        }
        v8 := []int{}
        var v9 uint = 0
        _ = v8
        _ = v9
    }
    if (100 >= 42) {
        fmt.Println(v6, v6, v2)
    } else {
        if (255 != 42) {
            _ = len(v4)
            v10 := []int{9223372036854775807, 9223372036854775807, 127, 42}
            _ = v10
        }
        v2 += (v2 / ((v2) | 1))
    }
    v11 := GenS1{F0: 255, F1: 2, F2: 157}
    v5 = v5
    if (!(v2 == v2)) {
        var v12 uint16 = 42
        v1 += 2147483647
        _ = v12
    } else {
        var v13 float32 = 0.0
        _ = v13
    }
    switch 1 {
    case 128:
        if (!v5) {
            var v14 int = 529
            v15 := []int{0, (^(v14))}
            v16 := GenS1{F0: 331, F1: 57, F2: 65535}
            _ = v14
            _ = v15
            _ = v16
        } else {
            var v17 uint16 = 0
            _ = v17
        }
    case 284:
        v5 = (1000 > 42)
    default:
        switch "a" {
        case "":
            fmt.Println(v5, v6, v1)
        case "hello":
            if (100 < 860) {
                v2 = 989
                var v18 uint64 = 1
                v19 := []int{2147483647, -2147483648, 633, int(v2)}
                _ = v18
                _ = v19
            }
        default:
            v20 := GenS1{F0: 65535, F1: 335, F2: 32767}
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        var v22 float32 = float32(-2.25)
        v6 = ("a" + "\n\t")
        _ = fi21
        _ = v22
    }
    var v23 string = "foo bar"
    fmt.Println(v2)
    if false {
        var v24 int64 = -9223372036854775808
        _ = v24
    }
    if v5 {
        for fi25 := 0; fi25 < 11; fi25++ {
            fmt.Println(v6)
            for fi26 := 0; fi26 < 1; fi26++ {
                v27 := GenS2{F0: v5, F1: v23}
                v28 := GenS1{F0: 1, F1: 0, F2: 338}
                switch v23 {
                case "\n\t":
                    v29 := GenS3{F0: 98, F1: 18446744073709551615, F2: (894 >= 1)}
                    _ = v29
                case "hello":
                    if v5 {
                        if (-32768 != 0) {
                            v30 := v3[(v2 % (((v2 | v2)) | 1))]
                            v30 = (v6 + (v30 + ((v23 + v23) + ("a" + (v30 + ("foo bar" + v6))))))
                            var v31 uint = 42
                            _ = v30
                            _ = v31
                        }
                    }
                default:
                    switch fi26 {
                    case 42:
                        if (v23 <= ("\n\t" + v23)) {
                            if false {
                                fmt.Println(v2, v23, v23)
                                v1 -= (v1 % (((v1 | 2147483647)) | 1))
                                var v32 int64 = 605
                                _ = v32
                            }
                            v33 := GenS2{F0: (true || (255 <= 1)), F1: "\n\t"}
                            _ = v33
                        }
                    case 418:
                        if (!v5) {
                            v23 = "a"
                            fmt.Println(fi26, v5, v1)
                            var v34 int64 = 255
                            _ = v34
                        }
                    }
                }
                _ = fi26
                _ = v27
                _ = v28
            }
            for fi35 := 0; fi35 < 15; fi35++ {
                var v36 int32 = (v1 / (((-(v1))) | 1))
                _ = fi35
                _ = v36
            }
            _ = fi25
        }
        v37 := []int{32767, 42}
        _ = v37
    } else {
        v38 := GenS3{F0: -1, F1: 2, F2: false}
        for fi39 := 0; fi39 < 18; fi39++ {
            v23 = v23
            for fi40 := 0; fi40 < 5; fi40++ {
                var v41 int = (fi39 * (-(fi40)))
                v42 := GenS4{F0: "", F1: 0}
                _ = fi40
                _ = v41
                _ = v42
            }
            for fi43 := 0; fi43 < 15; fi43++ {
                var v44 uint16 = 1
                _ = fi43
                _ = v44
            }
            _ = fi39
        }
        var v45 int = 475
        _ = v38
        _ = v45
    }
    for fi46 := 0; fi46 < 1; fi46++ {
        v47 := []int{}
        if (!false) {
            v5 = (110 < 1)
        }
        _ = fi46
        _ = v47
    }
    fmt.Println(v5, v23)
    if (!(!(v5 && v5))) {
        for fi48 := 0; fi48 < 19; fi48++ {
            if true {
                v3[(v2 % (((v2 >> ((0) & 63))) | 1))] = v23
                v2 -= (^(v2))
                for fi49 := 0; fi49 < 4; fi49++ {
                    v1 = (v1 - (v1 - (v1 >> ((65535) & 63))))
                    _ = fi49
                }
            } else {
                _ = len(v3)
            }
            if (!((fi48 & fi48) >= 9223372036854775807)) {
                fmt.Println(v5, v1)
            }
            _ = fi48
        }
    }
    for fi50 := 0; fi50 < 10; fi50++ {
        _ = len(v3)
        _ = fi50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
