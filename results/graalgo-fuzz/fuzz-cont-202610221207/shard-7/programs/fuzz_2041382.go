package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 string }
func genFn0(p0 int8, p1 int8, p2 bool) int8 {
    p2 = false
    var v4 int8 = (p0 + 42)
    if ("a" > "foo bar") {
        v5 := GenS1{F0: 32767}
        _ = v5
    }
    switch -375 {
    case 127:
        if (660 <= 107) {
            var v6 uint64 = 50
            v4 = v4
            _ = v6
        } else {
            if (!((0.0 > 0.0) && p2)) {
                var v7 uint = 42
                _ = v7
            }
            var v8 uint = 126
            _ = v8
        }
    default:
        p1 = (p1 & 100)
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS1{F0: fi9}
        for fi11 := 0; fi11 < 9; fi11++ {
            if p2 {
                v12 := []int{-2147483648, (fi9 / ((fi9) | 1)), fi11}
                fmt.Println(fi9, p2)
                var v13 int8 = p0
                _ = v12
                _ = v13
            }
            if p2 {
                v14 := []int{}
                v15 := GenS1{F0: -1}
                v16 := GenS2{F0: 42}
                _ = v14
                _ = v15
                _ = v16
            }
            var v17 int16 = 32767
            _ = fi11
            _ = v17
        }
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 uint32 = 2147483648
            v20 := []int{(fi18 ^ fi9), fi9}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 / (((p1 << ((0) & 63))) | 1))
}

func main() {
    var v1 int16 = 0
    var v2 int = 128
    var v3 int = 4294967295
    v3 = (v3 % ((-1) | 1))
    v21 := []int{(v3 + v3), v2}
    for fi22 := 0; fi22 < 12; fi22++ {
        v23 := GenS1{F0: (v2 ^ v2)}
        if true {
            for fi24 := 0; fi24 < 6; fi24++ {
                v3 = v2
                if false {
                    v25 := GenS3{F0: 137}
                    switch fi22 {
                    case 32768:
                        var v26 uint = 18446744073709551615
                        _ = v26
                    case 2147483647:
                        if false {
                            v27 := GenS2{F0: 1}
                            v3 -= fi22
                            v28 := GenS4{F0: -2147483648, F1: 1}
                            switch 32767 {
                            case -2147483648:
                                v3 = 32768
                            case -9223372036854775808:
                                fmt.Println(fi22)
                            case -943:
                                v2 -= (fi22 << ((0) & 63))
                            }
                            _ = v27
                            _ = v28
                        } else {
                            var v29 int64 = 42
                            v21 = append(v21, (v2 ^ 32767))
                            fmt.Println(v2)
                            _ = v29
                        }
                    case 0:
                        var v30 uint8 = 128
                        _ = v30
                    }
                    _ = len(v21)
                    if (42 == 1) {
                        var v31 int64 = 323
                        if len(v21) > 0 {
                            v21[0] = fi24
                        }
                        var v32 float64 = (3.14159 / 1e10)
                        _ = v31
                        _ = v32
                    } else {
                        var v33 string = "a"
                        v33 = ""
                        _ = v33
                    }
                    _ = v25
                } else {
                    v3 -= (v3 * -636)
                }
                _ = fi24
            }
            v34 := GenS5{F0: 255, F1: 2147483647}
            switch v3 {
            case 568:
                switch 127 {
                case 4294967295:
                    v35 := GenS3{F0: -1}
                    _ = v35
                default:
                    var v36 bool = false
                    _ = v36
                }
            }
            _ = v34
        }
        _ = fi22
        _ = v23
    }
    var v37 uint32 = 134
    if (1.5 >= 0.0) {
        v38 := GenS5{F0: 127, F1: 0}
        var v39 string = ("a" + "\n\t")
        _ = v38
        _ = v39
    } else {
        if false {
            v40 := GenS2{F0: 1}
            fmt.Println(v1, v2)
            _ = v40
        }
        var v41 uint16 = 1
        _ = v41
    }
    if true {
        v3 = v2
        for fi42 := 0; fi42 < 5; fi42++ {
            switch (v2 >> ((0) & 63)) {
            case -337:
                v2 -= (fi42 - fi42)
            case 32767:
                fmt.Println(v1)
            default:
                fmt.Println(fi42)
            }
            _ = fi42
        }
        v21 = append(v21, 127)
    }
    _ = len(v21)
    for fi43 := 0; fi43 < 17; fi43++ {
        for fi44 := 0; fi44 < 4; fi44++ {
            v3 -= -266
            v1 -= v1
            v21 = append(v21, v2)
            _ = fi44
        }
        _ = fi43
    }
    v45 := GenS6{F0: "\n\t"}
    switch 32768 {
    case -2147483648:
        v2 = (v3 % ((v2) | 1))
    case 2147483647:
        var v46 int16 = (v1 >> ((729) & 63))
        _ = v46
    case 836:
        var v47 uint16 = 65535
        _ = v47
    }
    var v48 uint = 159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v45
    _ = v48
}
