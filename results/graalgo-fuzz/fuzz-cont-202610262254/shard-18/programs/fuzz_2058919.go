package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 int64; F1 int }
func genFn0(p0 int32) int64 {
    var v7 float64 = 0.1
    if (false || false) {
        var v8 uint16 = 32767
        var v9 int64 = -9223372036854775808
        v9 = -9223372036854775808
        _ = v8
        _ = v9
    } else {
        var v10 uint = 1
        if (0 < 384) {
            fmt.Println(p0, p0)
            switch 1 {
            case 1:
                var v11 int8 = -1
                _ = v11
            default:
                v10 -= (v10 ^ (v10 + 66))
            }
        }
        fmt.Println(v7)
        _ = v10
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        if ((560 < 0) || (false || false)) {
            v13 := GenS1{F0: 42, F1: 42}
            _ = v13
        }
        _ = fi12
    }
    _ = p0
    _ = v7
    return -1
}
func genFn1(p0 uint, p1 uint16) int32 {
    v14 := GenS1{F0: 9223372036854775807, F1: -1}
    var v15 int64 = 9223372036854775807
    var v16 uint = (p0 % ((p0) | 1))
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return int32(p1)
}

func main() {
    var v1 uint16 = 65535
    var v2 int = -137
    v3 := make(map[string]uint32)
    var v4 int64 = -968
    v5 := make(map[string]string)
    var v6 uint8 = 42
    if (-1 <= int16(v1)) {
        var v17 uint8 = v6
        v18 := GenS2{F0: -768}
        _ = v17
        _ = v18
    } else {
        var v19 uint64 = 1
        var v20 int16 = -912
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: 240, F1: (v1 & 363), F2: 105}
    fmt.Println(v6, v1, v2)
    v3[("a" + "foo bar")] = 0
    if true {
        v22 := GenS3{F0: 234, F1: (v1 << ((65535) & 63)), F2: 1}
        _ = v22
    } else {
        v23 := GenS1{F0: genFn0(974), F1: 1}
        if true {
            for fi24 := 0; fi24 < 3; fi24++ {
                v25 := []int{32768}
                v4 = genFn0(2147483647)
                for fi26 := 0; fi26 < 11; fi26++ {
                    v27 := GenS4{F0: (false || false), F1: uint8(v1)}
                    var v28 uint8 = 255
                    var v29 string = ("" + "a")
                    _ = fi26
                    _ = v27
                    _ = v28
                    _ = v29
                }
                _ = fi24
                _ = v25
            }
            var v30 uint64 = 882
            switch ("" + "hello") {
            case "a":
                v30 = (v30 & v30)
            }
            _ = v30
        } else {
            for fi31 := 0; fi31 < 10; fi31++ {
                fmt.Println(v6, v1)
                var v32 int16 = 0
                v33 := []int{}
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        _ = v23
    }
    var v34 uint = 1
    if ((v6 == v6) || ((1 != (v2 / ((v2) | 1))) && (false || (-1 != 1)))) {
        for fi35 := 0; fi35 < 2; fi35++ {
            if true {
                switch "hello" {
                case "":
                    v34 -= v34
                case "a":
                    v1 += 0
                default:
                    fmt.Println(v4, v4)
                }
            } else {
                if true {
                    v1 += 935
                    if (!(3.14159 == (0.0 - 0.0))) {
                        v6 -= 1
                    } else {
                        fmt.Println(v4, v34)
                    }
                    v4 = v4
                }
                var v36 bool = false
                v36 = false
                _ = v36
            }
            for fi37 := 0; fi37 < 15; fi37++ {
                v38 := GenS4{F0: true, F1: 127}
                var v39 string = "hello"
                var v40 uint64 = 4294967295
                _ = fi37
                _ = v38
                _ = v39
                _ = v40
            }
            _ = fi35
        }
        if false {
            var v41 int8 = 127
            v4 = int64(v34)
            var v42 int32 = 42
            _ = v41
            _ = v42
        } else {
            fmt.Println(v6, v2)
        }
    }
    v43 := GenS3{F0: 18446744073709551615, F1: v1, F2: 42}
    switch ("\n\t" + "hello") {
    case "foo bar":
        var v44 string = "foo bar"
        _ = v44
    case "\n\t":
        fmt.Println(v4, v2, v2)
    case "a":
        for fi45 := 0; fi45 < 2; fi45++ {
            v46 := GenS2{F0: 0}
            if false {
                v34 += v34
                v47 := GenS5{F0: -128}
                var v48 int64 = genFn0(0)
                _ = v47
                _ = v48
            }
            _ = fi45
            _ = v46
        }
    default:
        v49 := GenS6{F0: -2147483648, F1: 42}
        _ = v49
    }
    fmt.Println(v6)
    if true {
        if (-604 != 1) {
            for fi50 := 0; fi50 < 13; fi50++ {
                for fi51 := 0; fi51 < 18; fi51++ {
                    fmt.Println(v1, v34, v1)
                    if (42 < 2) {
                        v52 := GenS7{F0: v4, F1: (fi50 << ((2) & 63))}
                        var v53 uint8 = (^(v6))
                        fmt.Println(v4)
                        _ = v52
                        _ = v53
                    }
                    switch "\n\t" {
                    case "":
                        if true {
                            switch (v2 / ((fi50) | 1)) {
                            case 128:
                                if true {
                                    var v54 uint = v34
                                    var v55 uint16 = (^(v1))
                                    v5["hello"] = "hello"
                                    _ = v54
                                    _ = v55
                                }
                                var v56 int64 = v4
                                _ = v56
                            case -2147483648:
                                switch ((("a" + "a") + ("" + ("hello" + "\n\t"))) + "foo bar") {
                                case "foo bar":
                                    if (v34 <= (v34 << ((v34) & 63))) {
                                        v1 = (v1 - 42)
                                        if false {
                                            v57 := []int{fi51}
                                            var v58 float64 = 1e10
                                            _ = v57
                                            _ = v58
                                        } else {
                                            v2 -= fi50
                                            v6 += (v6 | 0)
                                            v2 -= fi51
                                        }
                                    }
                                case "":
                                    _ = len(v5)
                                default:
                                    v59 := []int{v2, fi51, 4294967295, fi51}
                                    _ = v59
                                }
                            default:
                                v6 -= uint8(v34)
                            }
                        }
                    }
                    _ = fi51
                }
                _ = fi50
            }
        }
        v60 := GenS7{F0: v4, F1: v2}
        v6 = 128
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
