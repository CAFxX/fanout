package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int32; F1 int32; F2 int16 }
type GenS5 struct { F0 int; F1 uint32; F2 bool }
func genFn0() uint32 {
    var v4 float32 = (1.5 + ((0.1 + ((0.1 + (float32(3.14159) - 0.0)) / 1.0)) - ((0.1 + 0.1) - (float32(((((((0.1 * (1.5 + 1e10)) - 0.0) + 3.14159) - -2.25) * 3.14159) * 3.14159)) * ((1e10 * (3.14159 / ((((0.0 - -2.25) - -2.25) * -2.25) - 1e10))) + float32((float64(-2.25) - 1e10)))))))
    v4 = 0.0
    for fi5 := 0; fi5 < 10; fi5++ {
        fmt.Println(v4, v4)
        fmt.Println(v4)
        _ = fi5
    }
    v4 += (v4 * float32(0.1))
    _ = v4
    return 530
}
func genFn1() uint32 {
    var v6 uint16 = 2
    var v7 int32 = int32(v6)
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    var v3 int64 = 1
    var v8 uint8 = 0
    v3 -= -9223372036854775808
    switch "a" {
    case "hello":
        var v9 int8 = -13
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 2; fi10++ {
            if false {
                for fi11 := 0; fi11 < 3; fi11++ {
                    if false {
                        switch (fi11 >> ((0) & 63)) {
                        case 128:
                            var v12 uint = 4294967295
                            _ = v12
                        case 9223372036854775807:
                            v13 := GenS1{F0: 42}
                            _ = v13
                        case -363:
                            v8 += v8
                        default:
                            v14 := v1[876]
                            _ = v14
                        }
                        v15 := GenS1{F0: 846}
                        v8 += 128
                        _ = v15
                    }
                    v3 += v3
                    v8 -= (v8 ^ (v8 - 2))
                    _ = fi11
                }
                if true {
                    var v16 int = (fi10 % ((-9223372036854775808) | 1))
                    for fi17 := 0; fi17 < 1; fi17++ {
                        fmt.Println(fi10, fi17)
                        var v18 int16 = 0
                        _ = fi17
                        _ = v18
                    }
                    v3 -= v3
                    _ = v16
                }
                v19 := GenS1{F0: 1}
                _ = v19
            } else {
                for fi20 := 0; fi20 < 6; fi20++ {
                    v8 = v8
                    _ = fi20
                }
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 string = (("a" + (("\n\t" + ("" + ("a" + "\n\t"))) + (("hello" + "a") + ""))) + ("hello" + ""))
                    _ = len(v1)
                    fmt.Println(fi21)
                    v8 = (^(v8))
                    _ = len(v2)
                    _ = fi21
                    _ = v22
                }
            }
            switch -775 {
            case -2147483648:
                v23 := GenS2{F0: 4294967295}
                _ = v23
            }
            switch ("" + "foo bar") {
            case "hello":
                for fi24 := 0; fi24 < 14; fi24++ {
                    v25 := GenS3{F0: 120, F1: 1}
                    var v26 int = -173
                    fmt.Println(v3)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            case "":
                v8 -= (v8 * 237)
            case "\n\t":
                if (fi10 <= -1) {
                    for fi27 := 0; fi27 < 10; fi27++ {
                        var v28 uint64 = 2
                        _ = fi27
                        _ = v28
                    }
                    var v29 int8 = 42
                    var v30 uint32 = genFn1()
                    _ = v29
                    _ = v30
                } else {
                    var v31 int64 = (v3 % ((v3) | 1))
                    v31 = v31
                    v32 := []int{818, fi10, (-(fi10)), (^(fi10))}
                    _ = v31
                    _ = v32
                }
            }
            _ = fi10
        }
    default:
        v3 = -532
    }
    v3 = v3
    if true {
        switch "foo bar" {
        case "\n\t":
            for fi33 := 0; fi33 < 8; fi33++ {
                fmt.Println(v8, fi33, v8)
                var v34 uint64 = 2
                v3 = -648
                _ = fi33
                _ = v34
            }
        case "foo bar":
            if true {
                switch 4294967295 {
                case 0:
                    var v35 uint64 = 988
                    _ = v35
                }
                fmt.Println(v3, v8)
                if (!((false && true) && false)) {
                    var v36 uint16 = 0
                    var v37 uint8 = v8
                    _ = v36
                    _ = v37
                } else {
                    v38 := []int{-9223372036854775808, -670, 2147483647, 32768}
                    var v39 uint32 = 1
                    _ = v38
                    _ = v39
                }
            }
        default:
            var v40 bool = true
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 11; fi41++ {
            var v42 uint32 = 0
            fmt.Println(v42)
            _ = len(v1)
            _ = fi41
            _ = v42
        }
        if (127 > 937) {
            fmt.Println(v8)
        }
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        switch "a" {
        case "a":
            var v44 uint16 = 2
            _ = v44
        case "":
            delete(v1, 550)
        default:
            v45 := GenS4{F0: 1, F1: -664, F2: 0}
            _ = v45
        }
        _ = fi43
    }
    var v46 uint = 280
    var v47 uint16 = 828
    v48 := GenS5{F0: -930, F1: 1, F2: (true || (986 >= uint(v47)))}
    var v49 string = "\n\t"
    v50 := []int{}
    v8 = 127
    if false {
        var v51 uint = v46
        _ = v51
    } else {
        v49 = (v49 + v49)
    }
    var v52 bool = (980 >= 2147483647)
    for fi53 := 0; fi53 < 13; fi53++ {
        switch (fi53 >> ((2) & 63)) {
        case 32767:
            fmt.Println(v8, v47)
        }
        for fi54 := 0; fi54 < 8; fi54++ {
            v47 += 463
            v49 = v49
            _ = fi54
        }
        var v55 bool = (1000 < 32767)
        _ = fi53
        _ = v55
    }
    v49 = "a"
    for fi56 := 0; fi56 < 7; fi56++ {
        var v57 uint8 = (v8 + v8)
        var v58 int64 = -1
        v59 := GenS4{F0: 917, F1: 42, F2: -32768}
        _ = fi56
        _ = v57
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 5; fi60++ {
        var v61 bool = ((true && (667 != 2147483647)) && (42 < 915))
        if v52 {
            v46 += v46
        } else {
            v47 += 2
            var v62 int32 = 2147483647
            v63 := GenS4{F0: v62, F1: 2147483647, F2: -256}
            _ = v62
            _ = v63
        }
        var v64 bool = (false || (!(!(42 < -54))))
        _ = fi60
        _ = v61
        _ = v64
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v52:", v52)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v52
    _ = v8
}
