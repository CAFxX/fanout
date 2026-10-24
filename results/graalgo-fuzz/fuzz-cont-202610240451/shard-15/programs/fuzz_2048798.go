package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 bool; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint32; F2 uint }
type GenS4 struct { F0 int16; F1 uint64; F2 int8 }
type GenS5 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint; F1 uint8; F2 uint }
func genFn0() int32 {
    v4 := GenS1{F0: 42}
    switch 344 {
    case 332:
        switch 127 {
        case 4294967295:
            for fi5 := 0; fi5 < 14; fi5++ {
                switch ("hello" + "") {
                case "hello":
                    switch "hello" {
                    case "hello":
                        for fi6 := 0; fi6 < 8; fi6++ {
                            v7 := GenS1{F0: 4294967295}
                            var v8 int8 = -50
                            _ = fi6
                            _ = v7
                            _ = v8
                        }
                    default:
                        if (234 < 106) {
                            v9 := GenS1{F0: 0}
                            fmt.Println(fi5)
                            var v10 int16 = -32768
                            _ = v10
                            _ = v9
                        } else {
                            var v11 bool = false
                            _ = v11
                        }
                    }
                case "":
                    if true {
                        for fi12 := 0; fi12 < 16; fi12++ {
                            fmt.Println(fi12, fi12)
                            var v13 uint16 = 42
                            v14 := []int{(-(fi5)), fi12, 127, fi5}
                            _ = fi12
                            _ = v13
                            _ = v14
                        }
                        v15 := GenS2{F0: 42, F1: ((152 != -2147483648) || (!(!(150 != 0)))), F2: 100}
                        _ = v15
                    } else {
                        switch (fi5 % ((-947) | 1)) {
                        case 4294967295:
                            fmt.Println(fi5)
                        }
                    }
                default:
                    var v16 uint = 4294967295
                    _ = v16
                }
                switch 42 {
                case -1:
                    v17 := GenS3{F0: 127, F1: 2, F2: 42}
                    _ = v17
                }
                _ = fi5
            }
        default:
            v18 := GenS2{F0: 32767, F1: false, F2: 42}
            _ = v18
        }
    case 289:
        v19 := GenS4{F0: -595, F1: 456, F2: 100}
        _ = v19
    }
    v20 := []int{-9223372036854775808}
    for fi21 := 0; fi21 < 9; fi21++ {
        v20 = append(v20, (^(fi21)))
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 int = (fi21 / ((fi22) | 1))
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    _ = v20
    _ = v4
    return 0
}
func genFn1(p0 uint, p1 int8) int32 {
    if (float64(3.14159) > ((((float64(-2.25) - -2.25) - -2.25) + 3.14159) + ((3.14159 - (0.0 + 1e10)) - 1.5))) {
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := GenS1{F0: 1}
            _ = fi24
            _ = v25
        }
        if (!false) {
            for fi26 := 0; fi26 < 20; fi26++ {
                p1 -= (p1 << (((^(p0))) & 63))
                _ = fi26
            }
            p1 += p1
            v27 := GenS5{F0: 0, F1: 18446744073709551615, F2: 169}
            _ = v27
        } else {
            fmt.Println(p0, p1)
            p0 -= (p0 + 2)
            v28 := []int{-1}
            _ = v28
        }
        p1 = 0
    } else {
        if (1 < 42) {
            if false {
                p0 = (p0 & p0)
                for fi29 := 0; fi29 < 15; fi29++ {
                    v30 := GenS4{F0: -153, F1: 1, F2: p1}
                    _ = fi29
                    _ = v30
                }
                p1 -= (p1 >> ((p0) & 63))
            } else {
                v31 := GenS2{F0: -1, F1: (false || (!(964 < 128))), F2: (p1 / ((p1) | 1))}
                _ = v31
            }
        }
        var v32 bool = false
        _ = v32
    }
    p0 = (p0 & 42)
    v33 := GenS6{F0: 127}
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 string = ("foo bar" + (("\n\t" + "foo bar") + ("hello" + "foo bar")))
        p0 += 2
        _ = fi34
        _ = v35
    }
    p0 = (p0 << ((p0) & 63))
    _ = p0
    _ = p1
    _ = v33
    return 571
}

func main() {
    var v1 uint8 = 196
    var v2 uint8 = v1
    v3 := make(map[int]int)
    fmt.Println(v1, v2)
    v36 := GenS7{F0: 22, F1: -1}
    var v37 int32 = 2147483647
    v38 := GenS8{F0: 42, F1: 0}
    v37 += (v37 * -1)
    v39 := GenS9{F0: 2147483647}
    if false {
        v40 := GenS5{F0: 787, F1: 629, F2: 4294967295}
        _ = v40
    }
    if true {
        if true {
            if false {
                switch ((((("\n\t" + "\n\t") + ("" + ("a" + ("foo bar" + "\n\t")))) + "a") + "foo bar") + "\n\t") {
                case "foo bar":
                    v41 := []int{-783, 4294967295, 89, 2147483647}
                    _ = v41
                case "":
                    var v42 int8 = 31
                    _ = v42
                case "a":
                    if (-9223372036854775808 > -9223372036854775808) {
                        if (255 <= 0) {
                            var v43 int16 = 249
                            _ = v43
                        } else {
                            v37 = 1
                            v44 := []int{-557, -9223372036854775808, 32767}
                            _ = v44
                        }
                        v2 += v2
                    }
                default:
                    fmt.Println(v37, v2)
                }
                v45 := GenS6{F0: v1}
                v46 := GenS1{F0: 2}
                _ = v45
                _ = v46
            } else {
                var v47 uint64 = 220
                var v48 float64 = 1e10
                v49 := GenS9{F0: 342}
                _ = v47
                _ = v48
                _ = v49
            }
        } else {
            v37 = 1
            _ = len(v3)
        }
        var v50 int = 792
        _ = v50
    }
    v37 += -2147483648
    if true {
        fmt.Println(v37, v1, v1)
        for fi51 := 0; fi51 < 20; fi51++ {
            v37 += (-(v37))
            if ((100 == 42) && false) {
                for fi52 := 0; fi52 < 10; fi52++ {
                    if false {
                        v37 = (v37 - (v37 >> ((0) & 63)))
                        fmt.Println(v2, v2, fi51)
                        if (false || (true && false)) {
                            delete(v3, (fi51 | fi52))
                            v37 = 1
                            var v53 int64 = -731
                            _ = v53
                        }
                    }
                    _ = fi52
                }
            } else {
                var v54 float64 = 0.1
                delete(v3, (fi51 - (fi51 ^ fi51)))
                _ = v54
            }
            var v55 bool = (true || ((true || (true && true)) || true))
            _ = fi51
            _ = v55
        }
        if (42 != -32768) {
            var v56 int32 = -1
            _ = v56
        }
    } else {
        v2 += 172
        v1 += (v2 | (v1 << ((640) & 63)))
        if (true || false) {
            v1 += 2
        } else {
            v57 := []int{}
            _ = v57
        }
    }
    var v58 uint8 = v1
    var v59 bool = (false || false)
    if (float64(0.1) > float64((1.5 + (1.5 / 1.0)))) {
        v60 := GenS4{F0: 0, F1: 2, F2: 1}
        v37 += v37
        var v61 uint64 = 18446744073709551615
        _ = v60
        _ = v61
    } else {
        switch "a" {
        case "":
            v62 := GenS10{F0: 1, F1: v58, F2: 65535}
            _ = v62
        case "foo bar":
            fmt.Println(v59, v58, v58)
        }
        v2 = (v58 ^ 116)
    }
    if v59 {
        v1 = (v2 ^ v2)
    }
    var v63 float64 = -2.25
    var v64 uint16 = 536
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v58
    _ = v59
    _ = v63
    _ = v64
}
