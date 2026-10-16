package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint; F1 uint64; F2 bool }
func genFn0(p0 int, p1 int16, p2 uint8) uint32 {
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 int32 = 1
        var v6 int32 = v5
        if true {
            if false {
                if true {
                    for fi7 := 0; fi7 < 2; fi7++ {
                        fmt.Println(fi7, fi7)
                        _ = fi7
                    }
                    v6 -= -2147483648
                    fmt.Println(v6, v6)
                }
                v8 := GenS1{F0: 842, F1: 0}
                _ = v8
            } else {
                fmt.Println(v6)
                v5 = (v5 & v6)
            }
            if false {
                for fi9 := 0; fi9 < 4; fi9++ {
                    switch 9223372036854775807 {
                    case 32767:
                        if true {
                            var v10 uint = 451
                            if false {
                                p0 += fi9
                            }
                            _ = v10
                        }
                    }
                    switch 0 {
                    case 950:
                        var v11 int = -2147483648
                        fmt.Println(p2, fi9)
                        v12 := []int{(fi4 / (((p0 >> ((108) & 63))) | 1)), (v11 << ((65535) & 63)), 356}
                        _ = v11
                        _ = v12
                    }
                    _ = fi9
                }
                var v13 float32 = float32(0.0)
                _ = v13
            }
            fmt.Println(v6, v6, p0)
        } else {
            var v14 float32 = 0.1
            var v15 int32 = v5
            v16 := []int{(-(p0)), fi4, fi4}
            _ = v14
            _ = v15
            _ = v16
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    var v17 int32 = 1
    var v18 int32 = 2147483647
    if (!false) {
        v17 += (v18 * v18)
        v19 := []int{p0, (p0 >> ((1) & 63)), -9223372036854775808}
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS1{F0: 1, F1: 18446744073709551615}
            v17 += (^(v18))
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 2
}

func main() {
    var v1 bool = false
    v2 := make(map[int]uint32)
    var v3 uint8 = 1
    v22 := GenS2{F0: 919, F1: 0}
    v23 := GenS2{F0: 32767, F1: -717}
    var v24 uint64 = 1
    v3 -= 127
    v25 := GenS3{F0: (v3 | v3), F1: 2}
    var v26 string = (("\n\t" + ("hello" + ("foo bar" + "a"))) + "\n\t")
    var v27 uint64 = v24
    v27 = v24
    for fi28 := 0; fi28 < 18; fi28++ {
        if (!(false || v1)) {
            if (913 != uint32(v27)) {
                fmt.Println(fi28, v27)
            }
            switch v26 {
            case "hello":
                switch v26 {
                case "\n\t":
                    fmt.Println(v3, v1, fi28)
                case "a":
                    v3 += v3
                case "foo bar":
                    if false {
                        if v1 {
                            if false {
                                v29 := []int{}
                                v24 -= v27
                                v30 := GenS2{F0: 0, F1: 127}
                                _ = v29
                                _ = v30
                            } else {
                                v31 := []int{}
                                _ = v31
                            }
                            fmt.Println(v24, fi28)
                        }
                        v32 := GenS3{F0: (v3 / ((127) | 1)), F1: 1}
                        switch (fi28 >> ((18446744073709551615) & 63)) {
                        case 128:
                            for fi33 := 0; fi33 < 19; fi33++ {
                                _ = len(v2)
                                _ = fi33
                            }
                        case 618:
                            fmt.Println(v27, v1, v26)
                        }
                        _ = v32
                    } else {
                        v1 = true
                    }
                }
            case "":
                v34 := GenS3{F0: 0, F1: 2}
                _ = v34
            case "foo bar":
                if v1 {
                    for fi35 := 0; fi35 < 20; fi35++ {
                        var v36 uint64 = 2
                        _ = fi35
                        _ = v36
                    }
                    switch (v26 + (v26 + v26)) {
                    case "\n\t":
                        v27 -= v27
                    case "hello":
                        v2[(fi28 - fi28)] = genFn0(42, -32768, v3)
                    case "a":
                        v37 := GenS3{F0: v3, F1: 42}
                        _ = v37
                    default:
                        if v1 {
                            v38 := GenS4{F0: 18446744073709551615, F1: 4294967295, F2: false}
                            switch (fi28 + -1) {
                            case 4294967295:
                                v39 := GenS3{F0: v3, F1: 159}
                                _ = v39
                            case 339:
                                if (-142 < 2147483647) {
                                    v1 = (("foo bar" + "hello") < (v26 + "\n\t"))
                                    switch v26 {
                                    case "\n\t":
                                        for fi40 := 0; fi40 < 13; fi40++ {
                                            fmt.Println(fi28)
                                            fmt.Println(v27)
                                            _ = fi40
                                        }
                                    }
                                }
                            default:
                                v41 := GenS4{F0: 763, F1: v24, F2: (v1 && (0 < 18446744073709551615))}
                                _ = v41
                            }
                            _ = v38
                        }
                    }
                }
            default:
                fmt.Println(fi28, v1)
            }
        } else {
            var v42 bool = false
            v43 := GenS2{F0: 255, F1: -425}
            if (32767 != 386) {
                v42 = (2 != 0)
                v44 := GenS1{F0: 0, F1: 4294967295}
                switch (fi28 ^ 2147483647) {
                case -891:
                    v42 = v42
                case -305:
                    v24 = v24
                case -9223372036854775808:
                    var v45 string = (v26 + ("" + (v26 + (v26 + (v26 + "a")))))
                    _ = v45
                }
                _ = v44
            }
            _ = v42
            _ = v43
        }
        _ = fi28
    }
    v1 = v1
    var v46 int32 = -1
    var v47 int16 = -1
    for fi48 := 0; fi48 < 5; fi48++ {
        fmt.Println(fi48)
        v49 := []int{}
        _ = fi48
        _ = v49
    }
    v24 = (^(v24))
    fmt.Println(v47, v26)
    v1 = v1
    var v50 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v46
    _ = v47
    _ = v50
}
