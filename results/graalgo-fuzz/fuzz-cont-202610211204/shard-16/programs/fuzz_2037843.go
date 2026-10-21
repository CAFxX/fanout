package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 string; F2 bool }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 uint8 }
func genFn0() uint {
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 int16 = -1
        v6 := []int{(fi4 / (((fi4 - int(v5))) | 1)), (^(fi4))}
        var v7 uint64 = 2
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    if (false && true) {
        v8 := GenS1{F0: 1}
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint16 = 255
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        var v11 uint = 292
        for fi12 := 0; fi12 < 8; fi12++ {
            switch ((("foo bar" + (("foo bar" + ("hello" + ("" + ("hello" + "hello")))) + "foo bar")) + "\n\t") + "hello") {
            case "hello":
                switch fi12 {
                case 2147483647:
                    for fi13 := 0; fi13 < 17; fi13++ {
                        v11 -= v11
                        if false {
                            switch "" {
                            case "foo bar":
                                switch "a" {
                                case "foo bar":
                                    switch ("a" + "foo bar") {
                                    case "foo bar":
                                        if (!false) {
                                            fmt.Println(v11, fi12)
                                            fmt.Println(fi12, fi12, fi12)
                                        }
                                    case "":
                                        var v14 float64 = 0.1
                                        _ = v14
                                    case "\n\t":
                                        v15 := GenS2{F0: 128, F1: 0, F2: -2147483648}
                                        _ = v15
                                    default:
                                        v11 -= 123
                                    }
                                case "\n\t":
                                    var v16 float32 = 0.0
                                    _ = v16
                                }
                            case "\n\t":
                                if true {
                                    v11 += v11
                                    fmt.Println(fi12, fi12)
                                    v17 := []int{9223372036854775807, (^(fi12)), (-(fi13))}
                                    _ = v17
                                }
                            default:
                                var v18 uint32 = 65535
                                _ = v18
                            }
                            var v19 float64 = 0.0
                            if true {
                                v20 := []int{fi12}
                                var v21 uint8 = 42
                                _ = v20
                                _ = v21
                            }
                            _ = v19
                        }
                        switch 1 {
                        case -1:
                            v11 = v11
                        case 32768:
                            v22 := []int{(fi12 & (fi13 | (fi13 << (((v11 >> ((0) & 63))) & 63)))), fi13, (fi12 ^ (fi13 << ((42) & 63))), fi12}
                            _ = v22
                        default:
                            var v23 int32 = -794
                            if (0 > -32768) {
                                var v24 int16 = -1
                                _ = v24
                            }
                            v25 := GenS2{F0: 1, F1: (v11 % ((1) | 1)), F2: -2147483648}
                            _ = v23
                            _ = v25
                        }
                        v11 += v11
                        _ = fi13
                    }
                case -736:
                    v11 += 42
                }
            }
            v11 -= 35
            v11 = v11
            _ = fi12
        }
        _ = v11
    }
    switch "a" {
    case "hello":
        var v26 int16 = 895
        _ = v26
    case "":
        var v27 float64 = 1.5
        _ = v27
    case "foo bar":
        if true {
            var v28 int16 = 1000
            fmt.Println(v28)
            var v29 uint = 18446744073709551615
            _ = v28
            _ = v29
        } else {
            var v30 int = 32768
            v30 = (-(v30))
            switch (v30 * -1) {
            case 127:
                var v31 uint32 = 4294967295
                _ = v31
            case 905:
                v30 = -2147483648
            case 32768:
                v30 = v30
            default:
                v32 := []int{v30, (-(v30)), v30, -405}
                _ = v32
            }
            _ = v30
        }
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        var v34 float32 = -2.25
        var v35 float32 = ((v34 - v34) * 1e10)
        _ = fi33
        _ = v34
        _ = v35
    }
    return 42
}
func genFn1() bool {
    var v36 int16 = -288
    v37 := []int{666, -9223372036854775808, 9223372036854775807, 1}
    var v38 uint64 = 765
    _ = v36
    _ = v37
    _ = v38
    return (0 == 65535)
}

func main() {
    var v1 int16 = 627
    var v2 uint8 = 152
    var v3 uint8 = v2
    v2 -= (^(v3))
    if false {
        switch 4294967295 {
        case 32768:
            for fi39 := 0; fi39 < 5; fi39++ {
                v1 += v1
                switch "foo bar" {
                case "a":
                    v40 := GenS3{F0: 255}
                    _ = v40
                }
                _ = fi39
            }
        case 4294967295:
            v41 := GenS4{F0: v2, F1: "\n\t", F2: true}
            _ = v41
        case -9223372036854775808:
            v3 -= v2
        default:
            for fi42 := 0; fi42 < 2; fi42++ {
                for fi43 := 0; fi43 < 18; fi43++ {
                    var v44 float32 = (-2.25 * ((0.1 - 0.1) * float32(0.1)))
                    if false {
                        var v45 uint64 = 2
                        v46 := []int{(fi42 << ((721) & 63))}
                        v47 := GenS1{F0: 2147483647}
                        _ = v45
                        _ = v46
                        _ = v47
                    } else {
                        if (1 == 9223372036854775807) {
                            fmt.Println(v2, v1)
                        } else {
                            var v48 uint = 18446744073709551615
                            fmt.Println(v48)
                            _ = v48
                        }
                        fmt.Println(v2, fi42, v2)
                    }
                    var v49 uint16 = 0
                    _ = fi43
                    _ = v44
                    _ = v49
                }
                _ = fi42
            }
        }
    } else {
        v50 := []int{-916}
        for fi51 := 0; fi51 < 8; fi51++ {
            if (true || true) {
                if (!(!true)) {
                    for fi52 := 0; fi52 < 20; fi52++ {
                        var v53 float32 = 1e10
                        _ = fi52
                        _ = v53
                    }
                    var v54 uint64 = 4294967295
                    v1 -= -37
                    _ = v54
                }
                v55 := GenS5{F0: 2, F1: "hello"}
                switch "a" {
                case "foo bar":
                    v56 := GenS6{F0: 1, F1: -1}
                    _ = v56
                default:
                    v2 -= 255
                }
                _ = v55
            }
            _ = fi51
        }
        _ = v50
    }
    for fi57 := 0; fi57 < 19; fi57++ {
        v58 := []int{}
        _ = fi57
        _ = v58
    }
    fmt.Println(v2, v2, v1)
    if (!true) {
        var v59 float32 = (1e10 + (float32(float64(float32(float64((3.14159 * 0.1))))) - -2.25))
        fmt.Println(v2, v3)
        _ = v59
    }
    for fi60 := 0; fi60 < 7; fi60++ {
        v1 = (v1 >> ((18446744073709551615) & 63))
        v2 = v2
        var v61 int64 = 445
        _ = fi60
        _ = v61
    }
    v62 := GenS7{F0: 0, F1: 42}
    fmt.Println(v3, v2)
    fmt.Println(v2, v1, v3)
    var v63 int64 = -1
    var v64 uint16 = 225
    v65 := GenS8{F0: (v2 << ((65535) & 63))}
    fmt.Println(v1, v63, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
