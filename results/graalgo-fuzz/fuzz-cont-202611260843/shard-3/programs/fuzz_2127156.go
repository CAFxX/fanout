package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int }
type GenS2 struct { F0 uint64; F1 int16; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS7 struct { F0 int64; F1 uint8 }
func genFn0() int32 {
    var v6 int8 = 0
    fmt.Println(v6)
    _ = v6
    return 793
}
func genFn1(p0 bool) uint8 {
    if p0 {
        v7 := GenS1{F0: 2147483647, F1: 988, F2: 827}
        for fi8 := 0; fi8 < 3; fi8++ {
            switch fi8 {
            case -9223372036854775808:
                var v9 string = ("hello" + "foo bar")
                _ = v9
            default:
                var v10 string = ("a" + ("a" + "\n\t"))
                _ = v10
            }
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 int64 = 42
                _ = fi11
                _ = v12
            }
            var v13 int32 = 0
            _ = fi8
            _ = v13
        }
        fmt.Println(p0)
        _ = v7
    }
    v14 := []int{-9223372036854775808}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    _ = p0
    _ = v14
    return 101
}

func main() {
    var v1 int = 2147483647
    var v2 uint16 = 42
    var v3 uint8 = 2
    v4 := make(map[string]int)
    var v5 int = (v1 * (v1 * (v1 - v1)))
    v15 := GenS1{F0: (v1 << ((476) & 63)), F1: 748, F2: v5}
    switch v1 {
    case 261:
        var v16 uint = 0
        _ = v16
    }
    switch (("hello" + ((((("\n\t" + "\n\t") + "\n\t") + "\n\t") + "\n\t") + ("" + "a"))) + "a") {
    case "a":
        var v17 float32 = 3.14159
        _ = v17
    case "\n\t":
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(v2)
            _ = fi18
        }
    case "hello":
        v3 = v3
    default:
        if (-285 < -9223372036854775808) {
            v19 := v4[""]
            v5 -= (v1 << ((0) & 63))
            for fi20 := 0; fi20 < 15; fi20++ {
                v19 += v5
                v21 := GenS1{F0: v5, F1: 525, F2: 128}
                _ = fi20
                _ = v21
            }
            _ = v19
        } else {
            switch ("hello" + "foo bar") {
            case "hello":
                fmt.Println(v5, v2, v3)
            case "":
                if (v3 <= (v3 >> ((255) & 63))) {
                    v3 += (v3 ^ (v3 % ((v3) | 1)))
                    for fi22 := 0; fi22 < 16; fi22++ {
                        for fi23 := 0; fi23 < 9; fi23++ {
                            switch "\n\t" {
                            case "foo bar":
                                v24 := GenS1{F0: v5, F1: 0, F2: (v1 % (((-(v1))) | 1))}
                                _ = v24
                            }
                            var v25 int64 = 469
                            _ = fi23
                            _ = v25
                        }
                        var v26 int = v1
                        _ = fi22
                        _ = v26
                    }
                    v27 := GenS2{F0: 2, F1: 87, F2: 168}
                    _ = v27
                } else {
                    v2 += (^(v2))
                    var v28 int = 35
                    _ = v28
                }
            default:
                var v29 int8 = 127
                _ = v29
            }
        }
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        switch (v1 & (-(v5))) {
        case -9223372036854775808:
            v31 := []int{fi30, 32767}
            _ = v31
        default:
            var v32 float64 = 0.0
            _ = v32
        }
        _ = fi30
    }
    v33 := GenS3{F0: "hello"}
    switch "foo bar" {
    case "foo bar":
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 int8 = 42
            for fi36 := 0; fi36 < 10; fi36++ {
                v37 := GenS4{F0: (v3 & v3), F1: v35}
                _ = fi36
                _ = v37
            }
            _ = fi34
            _ = v35
        }
    case "hello":
        switch "" {
        case "":
            delete(v4, "foo bar")
        case "a":
            for fi38 := 0; fi38 < 8; fi38++ {
                v3 = v3
                v5 = (v1 & (v1 & v1))
                _ = fi38
            }
        }
    }
    var v39 uint = uint(v2)
    v40 := []int{(v1 / (((v1 - 603)) | 1)), v5, v1, (v1 % (((v5 + (^(v1)))) | 1))}
    if false {
        for fi41 := 0; fi41 < 12; fi41++ {
            var v42 int = -1
            var v43 int16 = -32768
            switch fi41 {
            case 0:
                fmt.Println(v3)
            case 128:
                fmt.Println(v42, v3)
            case 4294967295:
                var v44 uint16 = (^(v2))
                _ = v44
            }
            _ = fi41
            _ = v42
            _ = v43
        }
        v39 = (v39 ^ v39)
    }
    v45 := GenS2{F0: 18446744073709551615, F1: int16(v1), F2: 42}
    var v46 uint8 = v3
    for fi47 := 0; fi47 < 1; fi47++ {
        v48 := v4["foo bar"]
        _ = fi47
        _ = v48
    }
    for fi49 := 0; fi49 < 7; fi49++ {
        if ((!true) && false) {
            if (int(v39) != (v1 / ((490) | 1))) {
                v39 = 276
            } else {
                if false {
                    for fi50 := 0; fi50 < 2; fi50++ {
                        v46 -= v46
                        _ = fi50
                    }
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v46, v39, v2)
                    case "hello":
                        var v51 bool = true
                        _ = v51
                    case "\n\t":
                        v2 = v2
                    }
                }
            }
            var v52 string = "a"
            if false {
                var v53 uint64 = 2
                delete(v4, ("foo bar" + (((v52 + "a") + v52) + v52)))
                var v54 float64 = (float64(3.14159) + 0.1)
                _ = v53
                _ = v54
            } else {
                _ = len(v40)
                v55 := GenS5{F0: genFn0()}
                v56 := GenS6{F0: 1, F1: 2147483647, F2: (v2 - v2)}
                _ = v55
                _ = v56
            }
            _ = v52
        } else {
            for fi57 := 0; fi57 < 7; fi57++ {
                if false {
                    var v58 uint32 = 42
                    var v59 uint8 = 0
                    var v60 bool = (-1 >= 2147483647)
                    _ = v58
                    _ = v59
                    _ = v60
                } else {
                    if (9223372036854775807 <= 0) {
                        v61 := GenS1{F0: (-(v1)), F1: 893, F2: fi57}
                        var v62 uint64 = 2
                        var v63 string = ((("hello" + ("foo bar" + "a")) + ("a" + "foo bar")) + "hello")
                        var v64 uint32 = 65535
                        _ = v61
                        _ = v62
                        _ = v63
                        _ = v64
                    }
                    var v65 uint64 = 4294967295
                    v66 := v4["\n\t"]
                    _ = v65
                    _ = v66
                }
                if (int16(v5) > 1) {
                    if (!true) {
                        var v67 uint64 = 946
                        fmt.Println(v5, fi49, v3)
                        _ = v67
                    }
                    var v68 uint32 = 817
                    _ = v68
                }
                _ = fi57
            }
            v1 -= (v5 & v5)
            v69 := GenS7{F0: 9223372036854775807, F1: (v3 % ((1) | 1))}
            _ = v69
        }
        _ = fi49
    }
    var v70 float32 = -2.25
    v4[(("a" + (("foo bar" + "") + (("" + "\n\t") + "hello"))) + "hello")] = (v1 ^ (v1 - int(v3)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v45
    _ = v46
    _ = v5
    _ = v70
}
