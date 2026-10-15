package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 int8 }
func genFn0() uint32 {
    v6 := []int{32768, 852, 128, 32767}
    var v7 int8 = 112
    if true {
        var v8 uint16 = 255
        switch "\n\t" {
        case "foo bar":
            for fi9 := 0; fi9 < 20; fi9++ {
                if true {
                    if true {
                        fmt.Println(v7, fi9, v8)
                        for fi10 := 0; fi10 < 11; fi10++ {
                            v8 += 864
                            _ = fi10
                        }
                    } else {
                        var v11 bool = (((v8 + (^(v8))) == v8) || (!(int64(v7) != 283)))
                        var v12 uint8 = 42
                        _ = v11
                        _ = v12
                    }
                    var v13 uint32 = 1
                    v8 -= 2
                    _ = v13
                }
                v7 -= 127
                if true {
                    for fi14 := 0; fi14 < 4; fi14++ {
                        if true {
                            if (!(152 != 18446744073709551615)) {
                                var v15 bool = (427 != 32767)
                                v16 := GenS1{F0: int16(fi9), F1: (v8 * 123)}
                                v15 = false
                                _ = v15
                                _ = v16
                            }
                            var v17 uint8 = 127
                            _ = v17
                        } else {
                            fmt.Println(v7, fi14)
                            switch (fi14 + (fi9 << ((762) & 63))) {
                            case 42:
                                var v18 bool = true
                                _ = v18
                            case -2147483648:
                                v7 += (-(v7))
                            case -176:
                                switch (fi14 | -1) {
                                case -9223372036854775808:
                                    if true {
                                        switch (("hello" + "") + "hello") {
                                        case "a":
                                            if ((false && (65535 < 2147483648)) || (false || (4294967295 > 18446744073709551615))) {
                                                v19 := GenS1{F0: -32768, F1: 985}
                                                _ = v19
                                            }
                                        }
                                        v8 -= v8
                                    }
                                case -833:
                                    v8 = v8
                                case 9223372036854775807:
                                    v20 := GenS1{F0: 1, F1: v8}
                                    _ = v20
                                default:
                                    v7 -= (v7 & (v7 * -1))
                                }
                                var v21 bool = (true || (!(!((v7 + (^(v7))) > 0))))
                                if false {
                                    v7 = v7
                                    switch fi9 {
                                    case 725:
                                        fmt.Println(v8, fi9)
                                    case 732:
                                        fmt.Println(fi9)
                                    case 2147483647:
                                        v7 += (v7 << ((1) & 63))
                                        v22 := GenS1{F0: 0, F1: (v8 / ((v8) | 1))}
                                        _ = v22
                                    }
                                    switch "" {
                                    case "":
                                        v23 := GenS1{F0: -805, F1: 221}
                                        _ = v23
                                    default:
                                        v21 = false
                                    }
                                }
                                if (0 >= 848) {
                                    v24 := GenS1{F0: 32767, F1: 32767}
                                    _ = len(v6)
                                    _ = v24
                                }
                                v21 = v21
                                fmt.Println(v21, v21)
                                if ((^(v7)) > (v7 ^ v7)) {
                                    var v25 int8 = (v7 >> ((890) & 63))
                                    _ = v25
                                }
                                v26 := GenS2{F0: 0, F1: "\n\t"}
                                var v27 int = (fi14 % (((^(fi9))) | 1))
                                _ = v21
                                _ = v26
                                _ = v27
                            default:
                                fmt.Println(v7, v8)
                            }
                            v7 = 1
                        }
                        _ = fi14
                    }
                }
                _ = fi9
            }
        case "\n\t":
            v7 = 42
        }
        _ = v8
    } else {
        v7 -= v7
        v28 := GenS3{F0: "foo bar"}
        if true {
            var v29 uint = 18446744073709551615
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if (true || true) {
                v29 += v29
            }
            _ = v29
        } else {
            v30 := GenS4{F0: 391, F1: 1}
            v31 := GenS3{F0: "\n\t"}
            switch "" {
            case "\n\t":
                v7 -= (v7 - v7)
            case "":
                var v32 uint8 = 106
                _ = v32
            }
            _ = v30
            _ = v31
        }
        _ = v28
    }
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 uint, p1 uint32, p2 int8) uint {
    var v33 uint64 = 103
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    return p0
}

func main() {
    var v1 uint = 668
    var v2 string = ("foo bar" + "")
    v3 := []int{127, 42, 42}
    v4 := make(map[uint32]uint32)
    var v5 int32 = 42
    var v34 uint32 = genFn0()
    switch v2 {
    case "hello":
        for fi35 := 0; fi35 < 19; fi35++ {
            switch "a" {
            case "\n\t":
                switch "" {
                case "a":
                    _ = len(v4)
                case "\n\t":
                    if false {
                        for fi36 := 0; fi36 < 18; fi36++ {
                            delete(v4, v34)
                            v2 = v2
                            _ = fi36
                        }
                        switch v2 {
                        case "hello":
                            fmt.Println(v2, v34, fi35)
                        case "foo bar":
                            v2 = "hello"
                        }
                        if true {
                            var v37 float64 = 1e10
                            _ = v37
                        } else {
                            for fi38 := 0; fi38 < 19; fi38++ {
                                var v39 string = v2
                                v1 = (v1 << ((v1) & 63))
                                _ = fi38
                                _ = v39
                            }
                            for fi40 := 0; fi40 < 19; fi40++ {
                                var v41 uint = (v1 - (v1 - v1))
                                var v42 uint16 = 0
                                _ = fi40
                                _ = v41
                                _ = v42
                            }
                        }
                    } else {
                        v34 = (v34 + v34)
                        fmt.Println(fi35, fi35, v1)
                        for fi43 := 0; fi43 < 3; fi43++ {
                            fmt.Println(fi35, v2, fi43)
                            _ = len(v3)
                            var v44 uint = genFn1((v1 | (^(v1))), 0, 84)
                            _ = fi43
                            _ = v44
                        }
                    }
                default:
                    var v45 float64 = 1e10
                    _ = v45
                }
            }
            _ = fi35
        }
    case "":
        v2 = ((v2 + ((("\n\t" + "hello") + (v2 + v2)) + v2)) + v2)
    default:
        v46 := GenS2{F0: 620, F1: ("hello" + ("foo bar" + "foo bar"))}
        _ = v46
    }
    v47 := GenS5{F0: 0}
    var v48 bool = (false || (true && (!true)))
    v34 -= (v34 / ((2) | 1))
    var v49 uint64 = 18446744073709551615
    v48 = (!v48)
    v50 := GenS6{F0: 2, F1: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
