package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0() int64 {
    switch 739 {
    case 128:
        switch "foo bar" {
        case "":
            v4 := GenS1{F0: -2147483648, F1: 65535}
            _ = v4
        }
    case -564:
        v5 := []int{9223372036854775807, 32768, -9223372036854775808, -2147483648}
        _ = v5
    case 919:
        var v6 uint64 = 700
        _ = v6
    default:
        switch "a" {
        case "hello":
            var v7 int = 128
            _ = v7
        case "foo bar":
            var v8 int16 = 42
            _ = v8
        case "\n\t":
            switch -422 {
            case 605:
                if (3.14159 == 1e10) {
                    if (((!false) && false) || (true || true)) {
                        var v9 int = 32768
                        _ = v9
                    } else {
                        var v10 int = 1
                        v11 := []int{v10}
                        _ = v10
                        _ = v11
                    }
                    var v12 int16 = 0
                    _ = v12
                }
            case 1:
                var v13 int64 = -337
                _ = v13
            }
        default:
            var v14 int = 1
            _ = v14
        }
    }
    var v15 int64 = 823
    _ = v15
    return (^(v15))
}
func genFn1(p0 uint) uint32 {
    v16 := GenS1{F0: -385, F1: 40}
    var v17 float64 = (0.1 * ((3.14159 * (3.14159 * (-2.25 / -2.25))) / (0.1 * float64(1e10))))
    v18 := GenS1{F0: 128, F1: 4294967295}
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    return 65535
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 1
    var v3 uint8 = 42
    var v19 float32 = -2.25
    var v20 uint = 42
    var v21 int64 = 9223372036854775807
    var v22 uint32 = 42
    v23 := GenS1{F0: 127, F1: (v22 ^ (v22 >> (((v20 & 255)) & 63)))}
    for fi24 := 0; fi24 < 17; fi24++ {
        fmt.Println(v3)
        _ = fi24
    }
    switch 157 {
    case 127:
        switch (("foo bar" + ("hello" + (("foo bar" + "") + "\n\t"))) + "foo bar") {
        case "hello":
            v20 -= uint(v3)
        default:
            if false {
                for fi25 := 0; fi25 < 5; fi25++ {
                    switch 1 {
                    case 32767:
                        delete(v1, genFn1(v20))
                    case -659:
                        v2 = (v2 / ((42) | 1))
                    }
                    v20 = uint(v22)
                    _ = fi25
                }
                v26 := GenS1{F0: 32768, F1: uint32(v20)}
                _ = v26
            } else {
                for fi27 := 0; fi27 < 14; fi27++ {
                    v20 = 255
                    for fi28 := 0; fi28 < 16; fi28++ {
                        var v29 string = "\n\t"
                        _ = fi28
                        _ = v29
                    }
                    _ = fi27
                }
                var v30 uint16 = 2
                if false {
                    _ = len(v1)
                } else {
                    delete(v1, 42)
                    v20 -= 255
                    v31 := GenS1{F0: 468, F1: uint32(v3)}
                    _ = v31
                }
                _ = v30
            }
        }
    case 9223372036854775807:
        var v32 uint16 = 0
        _ = v32
    default:
        for fi33 := 0; fi33 < 11; fi33++ {
            if ((1.5 / (-2.25 * 1e10)) < 0.0) {
                if false {
                    if (v20 == (v20 ^ (^(v20)))) {
                        if false {
                            switch "" {
                            case "foo bar":
                                v3 -= (v3 % ((2) | 1))
                            case "a":
                                var v34 uint16 = uint16(v3)
                                _ = v34
                            }
                        } else {
                            v35 := GenS1{F0: -920, F1: uint32(v20)}
                            var v36 uint = 255
                            _ = v35
                            _ = v36
                        }
                        v20 += v20
                    }
                } else {
                    var v37 uint8 = (v3 + 255)
                    _ = v37
                }
            }
            var v38 int32 = 42
            v22 = v22
            _ = fi33
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 5; fi39++ {
        switch fi39 {
        case -983:
            v2 -= int8(v21)
        case 128:
            if false {
                v19 = v19
            }
        case 2147483647:
            var v40 int = (-(fi39))
            _ = v40
        }
        var v41 int8 = -1
        _ = fi39
        _ = v41
    }
    v42 := GenS1{F0: -2147483648, F1: uint32(v20)}
    v43 := GenS1{F0: 42, F1: (v22 & v22)}
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        switch ("hello" + "") {
        case "\n\t":
            _ = len(v1)
        }
    case "":
        for fi44 := 0; fi44 < 3; fi44++ {
            var v45 uint = (v20 >> (((v20 + v20)) & 63))
            var v46 bool = ("" <= ("foo bar" + "foo bar"))
            _ = len(v1)
            _ = fi44
            _ = v45
            _ = v46
        }
    default:
        var v47 string = ("" + ("" + ("" + (("foo bar" + "") + "\n\t"))))
        _ = v47
    }
    v21 = v21
    v48 := GenS1{F0: 9223372036854775807, F1: (v22 << (((v20 % ((uint(v3)) | 1))) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v42
    _ = v43
    _ = v48
}
