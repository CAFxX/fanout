package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 bool; F2 int32 }
type GenS4 struct { F0 bool; F1 int16; F2 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64; F1 bool; F2 uint64 }
func genFn0(p0 uint) int {
    var v6 int8 = 1
    if true {
        var v7 uint32 = 42
        v8 := GenS1{F0: 1, F1: 32767, F2: 277}
        _ = v7
        _ = v8
    } else {
        var v9 float32 = 1e10
        v6 += v6
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(p0, fi10)
        _ = fi10
    }
    _ = p0
    _ = v6
    return 32767
}

func main() {
    v1 := []uint32{2147483648, 183}
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    v4 := []uint32{42, 0}
    v5 := []int{}
    _ = len(v1)
    var v11 int16 = -32768
    if false {
        if true {
            v11 = v11
        } else {
            delete(v3, 2147483647)
            if (0 >= 255) {
                v11 -= (v11 % ((-1) | 1))
            }
        }
        if (42 < -128) {
            v11 = (v11 * (v11 - 1))
        }
    } else {
        if len(v5) > 0 {
            v5[0] = -2147483648
        }
        for fi12 := 0; fi12 < 20; fi12++ {
            fmt.Println(v11, v11)
            v2[fi12] = "\n\t"
            v11 = v11
            _ = fi12
        }
        v13 := GenS2{F0: 0}
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        switch ("\n\t" + "foo bar") {
        case "a":
            var v15 uint8 = 0
            _ = v15
        }
        var v16 uint8 = 127
        if false {
            v17 := GenS2{F0: 42}
            v16 = (v16 + (v16 - 169))
            if false {
                fmt.Println(fi14, v16)
            }
            _ = v17
        } else {
            if true {
                var v18 uint8 = (v16 << ((0) & 63))
                fmt.Println(fi14)
                _ = v18
            } else {
                for fi19 := 0; fi19 < 10; fi19++ {
                    v16 += v16
                    var v20 uint16 = 65535
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 2; fi21++ {
                    switch (("" + (("hello" + ("a" + "\n\t")) + "a")) + "hello") {
                    case "\n\t":
                        var v22 int32 = 678
                        _ = v22
                    case "foo bar":
                        var v23 string = "a"
                        _ = v23
                    case "hello":
                        var v24 uint8 = v16
                        _ = v24
                    default:
                        v25 := GenS3{F0: -505, F1: false, F2: 2147483647}
                        _ = v25
                    }
                    _ = fi21
                }
            }
        }
        _ = fi14
        _ = v16
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        switch "hello" {
        case "\n\t":
            for fi27 := 0; fi27 < 5; fi27++ {
                var v28 uint8 = 42
                _ = fi27
                _ = v28
            }
        }
        _ = fi26
    }
    if true {
        var v29 int32 = 42
        if false {
            var v30 int32 = (v29 ^ v29)
            v30 = v29
            _ = v30
        }
        if false {
            fmt.Println(v29, v29)
            v11 -= v11
            v11 = v11
        } else {
            v11 = v11
        }
        _ = v29
    }
    v11 += 32767
    var v31 int8 = -80
    v31 = (v31 - -128)
    var v32 int32 = int32(v31)
    v31 = int8(v11)
    if ((!false) || ((v31 & (v31 | v31)) == (v31 + (v31 * (v31 * v31))))) {
        if false {
            if true {
                v32 += -889
                v33 := GenS2{F0: 0}
                _ = v33
            } else {
                for fi34 := 0; fi34 < 10; fi34++ {
                    v11 = (v11 % ((v11) | 1))
                    switch "foo bar" {
                    case "hello":
                        var v35 int = fi34
                        _ = v35
                    case "":
                        if (209 <= 51) {
                            var v36 int8 = (v31 + (v31 * 1))
                            var v37 uint64 = 0
                            _ = v36
                            _ = v37
                        }
                    case "a":
                        v32 += v32
                    default:
                        var v38 uint = 4294967295
                        _ = v38
                    }
                    switch "\n\t" {
                    case "hello":
                        v39 := v2[genFn0(255)]
                        _ = v39
                    case "a":
                        fmt.Println(fi34, v11, v32)
                    case "":
                        if (824 <= 2147483647) {
                            _ = len(v2)
                        }
                    default:
                        v40 := GenS4{F0: (!(100 == 115)), F1: (v11 + v11), F2: 255}
                        _ = v40
                    }
                    _ = fi34
                }
                v41 := GenS4{F0: false, F1: (-(v11)), F2: 0}
                _ = v41
            }
            v11 = v11
            v31 += v31
        } else {
            if false {
                v42 := GenS5{F0: ((("a" + "foo bar") + "hello") + ((("\n\t" + "hello") + ("" + ("foo bar" + "foo bar"))) + "foo bar"))}
                if true {
                    if true {
                        var v43 int16 = v11
                        _ = v43
                    } else {
                        var v44 uint32 = 344
                        fmt.Println(v32, v44)
                        _ = v44
                    }
                } else {
                    var v45 int32 = v32
                    _ = len(v2)
                    for fi46 := 0; fi46 < 2; fi46++ {
                        v47 := GenS4{F0: false, F1: (-(v11)), F2: 0}
                        v48 := GenS5{F0: "\n\t"}
                        var v49 uint32 = 974
                        _ = fi46
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                    _ = v45
                }
                v50 := GenS6{F0: 128}
                _ = v42
                _ = v50
            }
            for fi51 := 0; fi51 < 14; fi51++ {
                var v52 bool = false
                if (!(v52 && (101 > 4294967295))) {
                    for fi53 := 0; fi53 < 19; fi53++ {
                        v31 -= v31
                        v11 = 42
                        _ = fi53
                    }
                    fmt.Println(fi51, v11)
                } else {
                    var v54 bool = (-33 != -9223372036854775808)
                    if (42 >= 0) {
                        v31 = (v31 >> ((2) & 63))
                        v3[107] = ((("foo bar" + ("foo bar" + "")) + "") >= "\n\t")
                        switch fi51 {
                        case 42:
                            for fi55 := 0; fi55 < 2; fi55++ {
                                var v56 uint8 = 255
                                var v57 int64 = 411
                                v1 = append(v1, 2147483647)
                                _ = fi55
                                _ = v56
                                _ = v57
                            }
                        case 871:
                            if (v32 != (v32 + (v32 + v32))) {
                                v52 = true
                            } else {
                                if (42 < 18446744073709551615) {
                                    v58 := GenS7{F0: 18446744073709551615, F1: (((-(v32)) > -2147483648) || v52), F2: 1}
                                    if v54 {
                                        for fi59 := 0; fi59 < 3; fi59++ {
                                            var v60 bool = v52
                                            _ = fi59
                                            _ = v60
                                        }
                                    } else {
                                        v11 += (-(v11))
                                    }
                                    _ = v58
                                }
                                var v61 int8 = -122
                                var v62 int16 = v11
                                _ = v61
                                _ = v62
                            }
                        default:
                            var v63 uint32 = 4294967295
                            _ = v63
                        }
                    } else {
                        v32 = (v32 ^ 0)
                        fmt.Println(v31, v52, v32)
                    }
                    _ = v54
                }
                _ = fi51
                _ = v52
            }
            v64 := GenS7{F0: 44, F1: true, F2: 18446744073709551615}
            _ = v64
        }
    }
    v4 = append(v4, 1)
    fmt.Println(v11, v11)
    if len(v4) > 0 {
        v4[0] = 654
    }
    v32 += -1
    var v65 int32 = v32
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v65
}
