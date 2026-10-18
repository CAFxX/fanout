package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int64, p1 int64) uint {
    v5 := GenS1{F0: 77, F1: 2147483648, F2: 32767}
    var v6 string = ("\n\t" + "hello")
    v7 := []int{-453, 32768, 128}
    v6 = (v6 + ((v6 + v6) + v6))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 255
}
func genFn1() int64 {
    v8 := []int{191, 0}
    var v9 int64 = -9223372036854775808
    _ = v8
    _ = v9
    return v9
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 127
    v3 := make(map[uint32]int)
    v4 := []uint32{42, 0, 0}
    switch 9223372036854775807 {
    case 0:
        switch ("" + "a") {
        case "foo bar":
            fmt.Println(v2)
        case "":
            var v10 float32 = (1e10 - -2.25)
            _ = v10
        }
    }
    if ((genFn1() <= -1) && (!(2147483647 >= -2147483648))) {
        v2 = (v2 << ((65535) & 63))
    } else {
        for fi11 := 0; fi11 < 10; fi11++ {
            v2 -= (^(v2))
            v2 = (v2 | (v2 - (v2 | v2)))
            fmt.Println(fi11)
            _ = fi11
        }
    }
    switch ("hello" + (("hello" + ("" + "a")) + "hello")) {
    case "hello":
        switch "hello" {
        case "":
            v2 = 100
        case "\n\t":
            v2 -= (v2 | (v2 & v2))
        default:
            v4 = append(v4, 2147483647)
        }
    case "foo bar":
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 uint = 4294967295
            if (v13 >= (v13 & 0)) {
                var v14 int = 902
                if (!(false || (true || false))) {
                    var v15 int = (fi12 % ((v14) | 1))
                    v16 := GenS1{F0: 1, F1: 4294967295, F2: 1000}
                    _ = v15
                    _ = v16
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    for fi17 := 0; fi17 < 2; fi17++ {
                        switch "foo bar" {
                        case "\n\t":
                            v18 := GenS1{F0: (v2 - 115), F1: 2, F2: 32767}
                            _ = v18
                        case "hello":
                            switch v14 {
                            case 4294967295:
                                if true {
                                    if ((false || (("" >= "hello") && (uint8(v13) <= 127))) && ((-2.25 * float64(((-2.25 + ((1e10 + float32((0.0 / 0.1))) * 3.14159)) / (-2.25 + 1.5)))) >= 0.1)) {
                                        v19 := GenS1{F0: v2, F1: 763, F2: 42}
                                        _ = v19
                                    } else {
                                        v2 += v2
                                        v20 := GenS1{F0: v2, F1: 1, F2: 481}
                                        _ = v20
                                    }
                                    _ = len(v1)
                                    if true {
                                        v14 = (fi12 % (((^(fi17))) | 1))
                                    }
                                } else {
                                    _ = len(v4)
                                    v4 = append(v4, 1)
                                    v21 := GenS2{F0: -917, F1: 450, F2: 65535}
                                    _ = v21
                                }
                            case 42:
                                if false {
                                    var v22 bool = true
                                    v23 := GenS3{F0: (v2 | v2)}
                                    _ = v22
                                    _ = v23
                                }
                            case -1:
                                var v24 uint64 = 1
                                _ = v24
                            }
                        case "":
                            if false {
                                fmt.Println(v2)
                                _ = len(v4)
                                if (0 <= 0) {
                                    if false {
                                        switch "a" {
                                        case "hello":
                                            var v25 int32 = -193
                                            _ = v25
                                        default:
                                            var v26 uint64 = 940
                                            _ = v26
                                        }
                                    }
                                    var v27 float64 = (1e10 * (-2.25 - 0.0))
                                    fmt.Println(v14, v13, fi12)
                                    _ = v27
                                }
                            } else {
                                var v28 uint16 = 42
                                switch ("\n\t" + ("hello" + "a")) {
                                case "\n\t":
                                    switch (fi17 + (-(fi12))) {
                                    case -49:
                                        switch "" {
                                        case "foo bar":
                                            switch "hello" {
                                            case "foo bar":
                                                var v29 int8 = v2
                                                _ = v29
                                            }
                                        }
                                    case 917:
                                        v30 := v1[2]
                                        _ = v30
                                    case 460:
                                        if ((true || (!true)) || (1 >= 1)) {
                                            var v31 uint8 = 0
                                            _ = v31
                                        }
                                    }
                                case "a":
                                    if true {
                                        if false {
                                            v13 -= (v13 << ((v13) & 63))
                                        }
                                        if true {
                                            _ = len(v3)
                                            var v32 uint8 = 42
                                            _ = v32
                                        }
                                    }
                                }
                                _ = v28
                            }
                        }
                        if (1 < 42) {
                            var v33 uint32 = 65535
                            _ = len(v3)
                            _ = v33
                        }
                        var v34 int16 = 42
                        v4 = append(v4, 1)
                        v34 += v34
                        _ = fi17
                        _ = v34
                    }
                }
                _ = v14
            }
            _ = fi12
            _ = v13
        }
    }
    if false {
        switch ("hello" + "") {
        case "foo bar":
            fmt.Println(v2)
        case "a":
            var v35 bool = (false || true)
            _ = v35
        default:
            for fi36 := 0; fi36 < 8; fi36++ {
                for fi37 := 0; fi37 < 13; fi37++ {
                    v2 = v2
                    var v38 int16 = 0
                    _ = fi37
                    _ = v38
                }
                _ = fi36
            }
        }
        for fi39 := 0; fi39 < 16; fi39++ {
            if false {
                if len(v4) > 0 {
                    v4[0] = 1
                }
                for fi40 := 0; fi40 < 12; fi40++ {
                    v41 := GenS1{F0: (v2 / (((v2 >> ((782) & 63))) | 1)), F1: 2, F2: 38}
                    if (genFn1() >= 1) {
                        _ = len(v4)
                        var v42 int64 = -9223372036854775808
                        v2 = -1
                        _ = v42
                    }
                    switch "" {
                    case "foo bar":
                        if (((^(fi39)) != (fi40 - (fi40 << ((65535) & 63)))) || true) {
                            if false {
                                v43 := GenS1{F0: (v2 | v2), F1: 780, F2: 1}
                                if ((false && (!(32767 < 1))) || (((!(!false)) && (!(fi39 > (fi40 | fi39)))) && true)) {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                } else {
                                    if (18446744073709551615 > 431) {
                                        v2 = (-(v2))
                                    } else {
                                        v1[2147483648] = ""
                                        _ = len(v3)
                                        v44 := GenS3{F0: 34}
                                        _ = v44
                                    }
                                    if true {
                                        var v45 uint = 18446744073709551615
                                        _ = v45
                                    }
                                }
                                _ = v43
                            }
                            var v46 uint16 = 32767
                            _ = v46
                        } else {
                            v2 = (v2 << ((0) & 63))
                        }
                    }
                    _ = fi40
                    _ = v41
                }
                for fi47 := 0; fi47 < 17; fi47++ {
                    fmt.Println(v2, fi39)
                    _ = len(v1)
                    fmt.Println(fi47, fi39)
                    _ = fi47
                }
            } else {
                var v48 uint32 = 639
                var v49 float32 = 1.5
                v4 = append(v4, (v48 - 752))
                _ = v48
                _ = v49
            }
            _ = fi39
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 = (v2 + (^(v2)))
    v50 := GenS4{F0: 42}
    v2 += (v2 * -38)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
}
