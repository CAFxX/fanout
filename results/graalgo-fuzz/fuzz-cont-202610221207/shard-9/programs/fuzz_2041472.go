package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint16; F1 int; F2 int32 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 int8 = 117
    var v2 uint = 65535
    var v3 int16 = -599
    var v4 uint8 = 128
    var v5 int16 = v3
    if true {
        v6 := []int{-1, -241}
        v3 = -247
        var v7 int = 450
        _ = v6
        _ = v7
    } else {
        var v8 uint16 = 1
        if false {
            switch (((("hello" + "\n\t") + ("\n\t" + ("a" + ("foo bar" + ("\n\t" + "a"))))) + "\n\t") + ("" + "hello")) {
            case "hello":
                for fi9 := 0; fi9 < 15; fi9++ {
                    fmt.Println(v2, v4)
                    _ = fi9
                }
            }
            fmt.Println(v2, v3, v4)
        } else {
            var v10 float32 = 3.14159
            for fi11 := 0; fi11 < 16; fi11++ {
                if (!false) {
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v10 = ((v10 + 1.5) + v10)
                        fmt.Println(v3)
                        if false {
                            var v13 int32 = 1
                            switch "" {
                            case "\n\t":
                                var v14 int64 = 0
                                _ = v14
                            case "foo bar":
                                var v15 uint = (v2 - v2)
                                _ = v15
                            default:
                                if true {
                                    v16 := GenS1{F0: 4294967295, F1: v5}
                                    v17 := GenS1{F0: 18446744073709551615, F1: -32768}
                                    fmt.Println(v10, v8)
                                    _ = v16
                                    _ = v17
                                } else {
                                    if true {
                                        v18 := GenS1{F0: 42, F1: -32768}
                                        _ = v18
                                    }
                                    var v19 string = "foo bar"
                                    v5 -= v3
                                    _ = v19
                                }
                            }
                            v5 = (v3 >> ((65535) & 63))
                            fmt.Println(fi11)
                            if ((v2 << (((v2 & (^(v2)))) & 63)) <= (^(v2))) {
                                v20 := []int{fi11, 4294967295, (fi12 << ((v2) & 63)), fi11}
                                _ = v20
                            } else {
                                var v21 bool = (int(v8) > fi12)
                                _ = v21
                            }
                            fmt.Println(v1)
                            _ = v13
                        } else {
                            var v22 int64 = 1
                            _ = v22
                        }
                        _ = fi12
                    }
                    v8 = v8
                    v23 := GenS1{F0: 4294967295, F1: v3}
                    _ = v23
                } else {
                    v24 := GenS2{F0: 255, F1: 2147483647, F2: int32(v2)}
                    _ = v24
                }
                switch "a" {
                case "\n\t":
                    v25 := GenS3{F0: (v3 >> ((4294967295) & 63)), F1: 18446744073709551615}
                    _ = v25
                case "foo bar":
                    switch fi11 {
                    case 9223372036854775807:
                        if true {
                            v26 := []int{fi11, fi11}
                            if ((false || false) && (!(float32(1e10) < (v10 / v10)))) {
                                v3 -= (v3 & (^(v5)))
                            }
                            _ = v26
                        }
                    }
                case "a":
                    if false {
                        v4 += 0
                        var v27 float32 = 3.14159
                        v2 = v2
                        _ = v27
                    } else {
                        if true {
                            switch "a" {
                            case "foo bar":
                                if (true || (v8 != 65535)) {
                                    if false {
                                        var v28 string = "foo bar"
                                        _ = v28
                                    } else {
                                        for fi29 := 0; fi29 < 15; fi29++ {
                                            fmt.Println(v2, v1, v4)
                                            switch fi11 {
                                            case 9223372036854775807:
                                                v2 -= (v2 - uint(v8))
                                            default:
                                                v10 -= v10
                                            }
                                            v5 = v5
                                            if (42 <= 4294967295) {
                                                v1 += 127
                                                var v30 int8 = (v1 | v1)
                                                var v31 uint = v2
                                                _ = v30
                                                _ = v31
                                            } else {
                                                v2 = 843
                                                fmt.Println(v10, v5)
                                            }
                                            _ = fi29
                                        }
                                    }
                                    v3 = -1
                                } else {
                                    for fi32 := 0; fi32 < 13; fi32++ {
                                        var v33 int16 = v5
                                        var v34 float32 = float32(1.5)
                                        var v35 int = (fi32 / ((fi32) | 1))
                                        _ = fi32
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                    v4 += 0
                                }
                            }
                            v3 = v3
                        }
                        var v36 uint16 = v8
                        fmt.Println(v10, v10, v1)
                        _ = v36
                    }
                }
                v4 += (v4 << ((255) & 63))
                _ = fi11
            }
            switch "foo bar" {
            case "\n\t":
                var v37 bool = false
                _ = v37
            case "hello":
                v38 := GenS1{F0: uint64(v2), F1: v5}
                _ = v38
            }
            _ = v10
        }
        v39 := []int{0}
        _ = v39
        _ = v8
    }
    v4 += (v4 + (v4 % ((v4) | 1)))
    v1 = (-(v1))
    var v40 float32 = 1.5
    v1 += v1
    v41 := []int{118}
    var v42 float64 = (-2.25 * 0.0)
    fmt.Println(v4, v42)
    for fi43 := 0; fi43 < 10; fi43++ {
        var v44 int32 = 1
        v44 -= (v44 & v44)
        _ = fi43
        _ = v44
    }
    var v45 uint16 = 2
    v46 := GenS1{F0: 0, F1: v5}
    var v47 int32 = -857
    if ((false || true) || false) {
        if len(v41) >= 2 {
            v41 = v41[1:len(v41)]
        }
    } else {
        for fi48 := 0; fi48 < 9; fi48++ {
            if len(v41) >= 2 {
                v41 = v41[1:len(v41)]
            }
            _ = fi48
        }
    }
    if false {
        var v49 uint8 = uint8(v45)
        v47 += (v47 & v47)
        _ = v49
    } else {
        for fi50 := 0; fi50 < 7; fi50++ {
            v51 := GenS4{F0: (v3 + 32767)}
            v4 += 128
            _ = fi50
            _ = v51
        }
        switch ("hello" + "") {
        case "foo bar":
            v4 -= (v4 + 163)
        }
    }
    var v52 string = ((("" + "a") + "a") + "hello")
    var v53 int = 244
    var v54 uint16 = 389
    if true {
        switch v53 {
        case -62:
            v41 = append(v41, v53)
        case 307:
            v40 = v40
        default:
            var v55 string = ("a" + v52)
            _ = v55
        }
        v56 := GenS5{F0: (-(v47))}
        _ = len(v41)
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v52
    _ = v53
    _ = v54
}
