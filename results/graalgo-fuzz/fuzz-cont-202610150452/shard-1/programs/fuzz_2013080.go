package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int32; F2 int32 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint }
type GenS4 struct { F0 int16; F1 int16 }
func genFn0() int32 {
    var v7 int32 = 1
    var v8 uint16 = 498
    v9 := GenS1{F0: 79}
    if true {
        var v10 float32 = 0.1
        v11 := []int{}
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(v8, fi12)
            v13 := GenS1{F0: 42}
            _ = fi12
            _ = v13
        }
        _ = v10
        _ = v11
    } else {
        var v14 int8 = 0
        _ = v14
    }
    switch (("foo bar" + "hello") + ((("foo bar" + "foo bar") + "") + "a")) {
    case "\n\t":
        v7 += (v7 * (-(v7)))
    case "hello":
        var v15 uint8 = 190
        _ = v15
    }
    _ = v7
    _ = v8
    _ = v9
    return (v7 + (v7 << ((289) & 63)))
}
func genFn1() uint8 {
    var v16 int32 = 1
    for fi17 := 0; fi17 < 3; fi17++ {
        switch -822 {
        case -9223372036854775808:
            var v18 int = -107
            _ = v18
        case 9223372036854775807:
            v19 := []int{(fi17 & fi17), fi17}
            _ = v19
        default:
            v20 := []int{}
            _ = v20
        }
        _ = fi17
    }
    _ = v16
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 42
    v3 := make(map[int]uint32)
    var v4 string = "foo bar"
    v5 := make(map[string]uint32)
    var v6 int = -20
    switch v6 {
    case -762:
        for fi21 := 0; fi21 < 14; fi21++ {
            for fi22 := 0; fi22 < 5; fi22++ {
                v23 := []int{}
                if len(v23) > 0 {
                    v23[0] = -1
                }
                v4 = "\n\t"
                if len(v23) > 0 {
                    v23[0] = (-(fi21))
                }
                if false {
                    v2 += v2
                    v24 := GenS2{F0: (v2 / (((v2 * genFn1())) | 1)), F1: 42, F2: -257}
                    _ = v24
                } else {
                    var v25 uint32 = 2
                    _ = v25
                }
                var v26 uint16 = 42
                _ = fi22
                _ = v23
                _ = v26
            }
            _ = fi21
        }
    case -738:
        var v27 uint16 = 2
        _ = v27
    case 858:
        v28 := GenS3{F0: (v2 << ((4294967295) & 63)), F1: 291, F2: 2}
        _ = v28
    default:
        fmt.Println(v4, v4, v2)
    }
    if false {
        v4 = v4
        for fi29 := 0; fi29 < 19; fi29++ {
            switch (v6 / (((fi29 % (((^(fi29))) | 1))) | 1)) {
            case 744:
                v6 = -1
            case 128:
                v30 := GenS4{F0: 625, F1: -809}
                _ = v30
            default:
                switch v4 {
                case "a":
                    fmt.Println(v6, v4, v2)
                case "foo bar":
                    var v31 float32 = 0.0
                    _ = v31
                }
            }
            _ = fi29
        }
    } else {
        v32 := GenS3{F0: (v2 + 42), F1: -272, F2: 2}
        fmt.Println(v2, v6, v6)
        v33 := []int{(v6 | (v6 << ((1) & 63)))}
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 1; fi34++ {
        fmt.Println(v6, v6)
        _ = fi34
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        v36 := []int{v6, (^(fi35)), (^(v6)), fi35}
        v6 += (^(v6))
        v4 = v4
        _ = fi35
        _ = v36
    }
    var v37 uint = 771
    var v38 int8 = -1
    switch v4 {
    case "":
        switch 32768 {
        case 234:
            if (!false) {
                v39 := []int{}
                for fi40 := 0; fi40 < 17; fi40++ {
                    for fi41 := 0; fi41 < 9; fi41++ {
                        switch int(v37) {
                        case 2147483647:
                            var v42 int16 = -32768
                            _ = v42
                        case -1:
                            v43 := GenS2{F0: genFn1(), F1: 890, F2: 29}
                            _ = v43
                        }
                        var v44 uint8 = (v2 >> (((^(v37))) & 63))
                        var v45 string = ""
                        _ = fi41
                        _ = v44
                        _ = v45
                    }
                    switch (fi40 | v6) {
                    case 32768:
                        var v46 int8 = (-(v38))
                        _ = v46
                    case 42:
                        v4 = v4
                    }
                    _ = fi40
                }
                if true {
                    fmt.Println(v2, v37, v4)
                    if true {
                        _ = len(v39)
                    }
                } else {
                    v47 := GenS3{F0: (^(v2)), F1: 9223372036854775807, F2: (v37 % ((v37) | 1))}
                    var v48 float32 = 0.1
                    _ = v47
                    _ = v48
                }
                _ = v39
            }
        default:
            for fi49 := 0; fi49 < 7; fi49++ {
                var v50 int16 = 1000
                _ = fi49
                _ = v50
            }
        }
    default:
        var v51 string = v4
        _ = v51
    }
    for fi52 := 0; fi52 < 19; fi52++ {
        var v53 uint8 = v2
        _ = fi52
        _ = v53
    }
    var v54 string = "hello"
    v55 := []int{int(v2), v6, (v6 | 4294967295)}
    if false {
        v2 += (v2 * (v2 << ((v37) & 63)))
        switch (v6 / ((42) | 1)) {
        case 42:
            v37 = v37
        case 128:
            for fi56 := 0; fi56 < 11; fi56++ {
                v2 = (v2 << ((430) & 63))
                for fi57 := 0; fi57 < 4; fi57++ {
                    delete(v1, v4)
                    var v58 uint64 = 4294967295
                    _ = fi57
                    _ = v58
                }
                _ = fi56
            }
        case 4294967295:
            var v59 int64 = 42
            _ = v59
        default:
            for fi60 := 0; fi60 < 18; fi60++ {
                switch v6 {
                case 4294967295:
                    if true {
                        var v61 float32 = 3.14159
                        switch -3 {
                        case 2147483647:
                            var v62 int16 = 1000
                            _ = v62
                        default:
                            var v63 uint8 = (^(v2))
                            _ = v63
                        }
                        for fi64 := 0; fi64 < 8; fi64++ {
                            var v65 string = "a"
                            if len(v55) > 0 {
                                v55[0] = v6
                            }
                            v54 = (v54 + "hello")
                            _ = fi64
                            _ = v65
                        }
                        _ = v61
                    }
                case 32768:
                    switch "\n\t" {
                    case "":
                        v66 := GenS1{F0: (v38 & v38)}
                        _ = v66
                    case "hello":
                        v67 := GenS2{F0: (v2 << (((v37 + v37)) & 63)), F1: -584, F2: -144}
                        _ = v67
                    case "foo bar":
                        var v68 float32 = 0.0
                        _ = v68
                    }
                default:
                    delete(v5, v54)
                }
                fmt.Println(v38, v37, v6)
                _ = fi60
            }
        }
    }
    v37 = (v37 >> ((v37) & 63))
    var v69 uint8 = (v2 / ((v2) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
    _ = v69
}
