package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 bool; F2 uint }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint8 }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 1; fi5++ {
        switch "a" {
        case "a":
            for fi6 := 0; fi6 < 18; fi6++ {
                var v7 uint32 = 65535
                var v8 int64 = 9223372036854775807
                _ = fi6
                _ = v7
                _ = v8
            }
        }
        _ = fi5
    }
    var v9 int8 = 11
    if false {
        v9 += (v9 << ((538) & 63))
        var v10 int8 = (v9 % (((v9 % (((v9 & (v9 - (-(v9))))) | 1))) | 1))
        _ = v10
    } else {
        switch "" {
        case "hello":
            v11 := GenS1{F0: -360, F1: 42, F2: -1}
            _ = v11
        case "foo bar":
            var v12 float32 = 0.0
            _ = v12
        case "":
            fmt.Println(v9)
        default:
            v13 := []int{}
            _ = v13
        }
    }
    if false {
        v14 := []int{975}
        switch "\n\t" {
        case "hello":
            var v15 string = "hello"
            _ = v15
        case "\n\t":
            v9 += -127
        case "foo bar":
            var v16 int8 = (v9 & v9)
            _ = v16
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 uint16 = 42
            _ = fi17
            _ = v18
        }
        _ = v14
    }
    if (true && false) {
        switch -56 {
        case -2147483648:
            if true {
                var v19 uint64 = 18446744073709551615
                v20 := GenS1{F0: 0, F1: 836, F2: 0}
                fmt.Println(v9, v19)
                _ = v19
                _ = v20
            }
        case 32768:
            v21 := GenS2{F0: 446}
            _ = v21
        default:
            v9 -= v9
        }
    } else {
        if ((1e10 - 1.5) <= -2.25) {
            for fi22 := 0; fi22 < 6; fi22++ {
                v9 = v9
                fmt.Println(v9)
                _ = fi22
            }
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 int16 = -1
                _ = fi23
                _ = v24
            }
            v25 := GenS2{F0: 1}
            _ = v25
        }
        if ((false || (501 == 68)) || (0 > 2)) {
            if true {
                for fi26 := 0; fi26 < 1; fi26++ {
                    v27 := []int{(fi26 ^ (fi26 | (fi26 ^ (fi26 - (fi26 << ((0) & 63)))))), (fi26 << ((733) & 63)), 0, fi26}
                    _ = fi26
                    _ = v27
                }
            } else {
                fmt.Println(v9)
                if (-210 < -32768) {
                    switch ("hello" + (("hello" + "foo bar") + ("foo bar" + "\n\t"))) {
                    case "\n\t":
                        fmt.Println(v9)
                    case "a":
                        v9 -= v9
                    default:
                        v9 = v9
                    }
                    for fi28 := 0; fi28 < 1; fi28++ {
                        var v29 uint64 = 0
                        var v30 string = "\n\t"
                        switch "foo bar" {
                        case "a":
                            v31 := []int{fi28}
                            _ = v31
                        default:
                            switch (v30 + "hello") {
                            case "":
                                fmt.Println(fi28, v29)
                            }
                        }
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                } else {
                    for fi32 := 0; fi32 < 20; fi32++ {
                        fmt.Println(fi32, v9)
                        _ = fi32
                    }
                }
                if (!(!false)) {
                    var v33 int16 = 32767
                    fmt.Println(v9, v33)
                    v34 := []int{32768, -1}
                    _ = v33
                    _ = v34
                }
            }
            switch "foo bar" {
            case "a":
                v9 = (-(v9))
            case "foo bar":
                var v35 uint8 = 42
                _ = v35
            case "hello":
                var v36 string = (("\n\t" + "hello") + "")
                _ = v36
            default:
                v37 := []int{-1, -952}
                _ = v37
            }
        }
        fmt.Println(v9)
    }
    _ = v9
    return 2
}
func genFn1() int64 {
    var v38 uint32 = 1
    fmt.Println(v38)
    _ = v38
    return 0
}

func main() {
    v1 := []int{346, 32768}
    var v2 uint32 = 1
    v3 := make(map[string]int)
    var v4 string = ("\n\t" + "hello")
    var v39 uint = 65535
    fmt.Println(v4, v39)
    fmt.Println(v2, v39)
    v2 -= v2
    v40 := GenS2{F0: 996}
    fmt.Println(v2)
    v4 = v4
    if true {
        var v41 int32 = 42
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v41
    }
    v39 = 255
    switch 32768 {
    case -720:
        var v42 uint64 = 1
        _ = v42
    }
    v39 = v39
    v43 := GenS3{F0: "foo bar", F1: true, F2: (^(v39))}
    if false {
        v44 := GenS1{F0: 9223372036854775807, F1: genFn0(), F2: genFn1()}
        _ = v44
    }
    var v45 float64 = 1.5
    v46 := GenS3{F0: v4, F1: false, F2: v39}
    if false {
        fmt.Println(v45)
        v47 := GenS1{F0: 1, F1: 754, F2: 841}
        _ = v47
    }
    for fi48 := 0; fi48 < 11; fi48++ {
        v2 = v2
        var v49 uint8 = 74
        switch ("a" + ((v4 + v4) + "foo bar")) {
        case "\n\t":
            if true {
                v45 = float64(1.5)
                v50 := GenS3{F0: "hello", F1: (!(1 <= 2147483647)), F2: uint(v49)}
                for fi51 := 0; fi51 < 1; fi51++ {
                    if true {
                        switch v4 {
                        case "hello":
                            var v52 int16 = 1
                            _ = v52
                        case "foo bar":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        v45 = v45
                    }
                    v2 = (v2 ^ v2)
                    var v53 uint16 = 32767
                    v54 := GenS1{F0: 9223372036854775807, F1: v53, F2: -263}
                    _ = fi51
                    _ = v53
                    _ = v54
                }
                _ = v50
            }
        case "foo bar":
            v55 := GenS4{F0: (^(v2)), F1: -380}
            _ = v55
        default:
            switch 828 {
            case -966:
                v56 := GenS5{F0: (v49 * (v49 >> (((v39 & (^(v39)))) & 63)))}
                _ = v56
            }
        }
        _ = fi48
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v45
    _ = v46
}
