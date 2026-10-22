package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int }
type GenS5 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS6 struct { F0 string; F1 uint16 }
func genFn0(p0 int, p1 int8) uint8 {
    fmt.Println(p1)
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: 152}
    v8 := GenS2{F0: (520 != 2), F1: 2}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 bool, p1 uint64, p2 uint32) int32 {
    if true {
        v9 := GenS3{F0: 87, F1: p1, F2: 1}
        switch "hello" {
        case "\n\t":
            fmt.Println(p2, p2)
        case "":
            switch 128 {
            case 566:
                if (255 > 255) {
                    p2 -= p2
                }
            case 0:
                for fi10 := 0; fi10 < 12; fi10++ {
                    if (42 >= 255) {
                        var v11 int8 = -128
                        v12 := []int{(^(fi10)), (fi10 + int(p2))}
                        fmt.Println(p1, p0, p2)
                        _ = v11
                        _ = v12
                    } else {
                        var v13 int16 = 1
                        var v14 uint32 = (p2 & 2147483647)
                        var v15 uint = 255
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    var v16 uint = 255
                    _ = fi10
                    _ = v16
                }
            case 1:
                if true {
                    p2 = p2
                } else {
                    var v17 int16 = 0
                    v18 := []int{}
                    _ = v17
                    _ = v18
                }
            default:
                v19 := GenS4{F0: p2, F1: 32767, F2: 576}
                _ = v19
            }
        case "a":
            v20 := GenS1{F0: 0}
            _ = v20
        default:
            fmt.Println(p1, p0, p1)
        }
        v21 := []int{-204, -1, -9223372036854775808}
        _ = v21
        _ = v9
    }
    if (128 <= 1) {
        fmt.Println(p1, p0, p2)
        var v22 bool = (("foo bar" <= "a") || p0)
        v22 = (!v22)
        _ = v22
    } else {
        switch 523 {
        case -621:
            var v23 int64 = -9223372036854775808
            _ = v23
        case -957:
            p1 += (p1 & p1)
        default:
            for fi24 := 0; fi24 < 7; fi24++ {
                for fi25 := 0; fi25 < 14; fi25++ {
                    v26 := []int{fi24}
                    var v27 uint64 = p1
                    v28 := GenS3{F0: 0, F1: uint64(p2), F2: 100}
                    v29 := GenS3{F0: 221, F1: v27, F2: 67}
                    var v30 uint16 = 65535
                    _ = fi25
                    _ = v26
                    _ = v27
                    _ = v28
                    _ = v29
                    _ = v30
                }
                _ = fi24
            }
        }
        switch "foo bar" {
        case "foo bar":
            if ((p0 && (-2.25 < (1e10 + 0.0))) && (!p0)) {
                if (("\n\t" + "") == "a") {
                    p2 = 46
                } else {
                    fmt.Println(p1)
                }
                switch "" {
                case "a":
                    var v31 float64 = (-2.25 / 1.5)
                    _ = v31
                default:
                    fmt.Println(p1)
                }
                switch 9223372036854775807 {
                case 2147483647:
                    if (("" + "foo bar") < ("hello" + "")) {
                        for fi32 := 0; fi32 < 17; fi32++ {
                            for fi33 := 0; fi33 < 15; fi33++ {
                                var v34 uint = 1
                                var v35 int16 = 42
                                v36 := []int{}
                                _ = fi33
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                            _ = fi32
                        }
                    } else {
                        if true {
                            fmt.Println(p0, p0)
                        } else {
                            p1 = 4294967295
                            v37 := []int{-503, 32767, 4294967295}
                            _ = v37
                        }
                    }
                }
            }
        case "hello":
            p1 -= (p1 % ((42) | 1))
        case "a":
            var v38 int64 = 9223372036854775807
            _ = v38
        default:
            p0 = ((("foo bar" + "hello") + "hello") >= "a")
        }
    }
    v39 := []int{-593, -171}
    v39 = append(v39, int(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    return -1
}

func main() {
    var v1 uint16 = 32767
    v2 := []int{32768, 42}
    var v3 int64 = 881
    var v4 int16 = 0
    v5 := []string{("\n\t" + "a"), "\n\t"}
    for fi40 := 0; fi40 < 7; fi40++ {
        switch ("\n\t" + "\n\t") {
        case "\n\t":
            v1 = (v1 << ((42) & 63))
        case "hello":
            var v41 uint32 = 42
            _ = v41
        default:
            var v42 int64 = v3
            _ = v42
        }
        v3 = (v3 - v3)
        v43 := GenS5{F0: v4, F1: 2, F2: (^(v1))}
        _ = fi40
        _ = v43
    }
    var v44 int64 = v3
    fmt.Println(v4, v1)
    var v45 string = "a"
    v3 += v3
    for fi46 := 0; fi46 < 12; fi46++ {
        _ = len(v5)
        for fi47 := 0; fi47 < 18; fi47++ {
            fmt.Println(fi46)
            fmt.Println(v45, fi47)
            v48 := GenS6{F0: "a", F1: v1}
            _ = fi47
            _ = v48
        }
        var v49 uint64 = 18446744073709551615
        _ = fi46
        _ = v49
    }
    for fi50 := 0; fi50 < 8; fi50++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v44, v3, v4)
        fmt.Println(v44, v3, fi50)
        _ = fi50
    }
    var v51 uint64 = 18446744073709551615
    for fi52 := 0; fi52 < 5; fi52++ {
        if true {
            switch fi52 {
            case 128:
                if true {
                    if (false && (255 <= 255)) {
                        fmt.Println(fi52, v45)
                        var v53 float64 = 1e10
                        _ = v53
                    } else {
                        var v54 int32 = 2147483647
                        var v55 bool = ((fi52 ^ int(v4)) != (fi52 ^ fi52))
                        v55 = true
                        _ = v54
                        _ = v55
                    }
                }
            }
        } else {
            for fi56 := 0; fi56 < 16; fi56++ {
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
                _ = fi56
            }
        }
        v45 = v45
        _ = fi52
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v51
}
