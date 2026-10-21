package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 int8; F2 bool }
type GenS6 struct { F0 int64; F1 int8; F2 int8 }
type GenS7 struct { F0 int16; F1 int64; F2 int64 }
type GenS8 struct { F0 uint16; F1 uint32; F2 bool }
func genFn0() uint {
    switch "foo bar" {
    case "":
        var v4 uint8 = 127
        _ = v4
    default:
        for fi5 := 0; fi5 < 11; fi5++ {
            v6 := []int{fi5, fi5, 0, 379}
            _ = fi5
            _ = v6
        }
    }
    switch ("a" + ("" + "a")) {
    case "hello":
        v7 := []int{-78, -93, 0}
        _ = v7
    case "\n\t":
        var v8 float64 = 1.5
        _ = v8
    case "":
        if ((3.14159 * 0.1) != (1e10 * float64(-2.25))) {
            var v9 uint32 = 349
            _ = v9
        }
    default:
        if true {
            var v10 int32 = 42
            _ = v10
        }
    }
    var v11 int = 42
    for fi12 := 0; fi12 < 5; fi12++ {
        if (-2.25 != 1e10) {
            switch (^(v11)) {
            case 32767:
                var v13 uint8 = 243
                _ = v13
            case 9223372036854775807:
                switch ("" + ("hello" + "hello")) {
                case "a":
                    var v14 int = v11
                    _ = v14
                }
            case -565:
                v11 += (fi12 * fi12)
            default:
                v11 -= fi12
            }
            v11 = -950
            if true {
                v11 = v11
            }
        } else {
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 int16 = 432
                _ = fi15
                _ = v16
            }
        }
        switch "a" {
        case "a":
            v17 := []int{fi12, v11, -870, (^(fi12))}
            _ = v17
        default:
            switch fi12 {
            case 0:
                v18 := GenS1{F0: true}
                _ = v18
            default:
                for fi19 := 0; fi19 < 8; fi19++ {
                    v20 := GenS2{F0: 42, F1: -128}
                    _ = fi19
                    _ = v20
                }
            }
        }
        fmt.Println(fi12)
        _ = fi12
    }
    _ = v11
    return 558
}
func genFn1() uint8 {
    var v21 string = ("\n\t" + "a")
    for fi22 := 0; fi22 < 4; fi22++ {
        v21 = ((v21 + ((v21 + "\n\t") + v21)) + (v21 + ("foo bar" + "a")))
        if (!true) {
            v21 = v21
        } else {
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 uint32 = 2
                _ = fi23
                _ = v24
            }
            v21 = v21
            for fi25 := 0; fi25 < 3; fi25++ {
                var v26 int8 = 42
                v27 := GenS3{F0: true, F1: 41}
                v28 := GenS2{F0: 425, F1: v26}
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
        }
        var v29 uint32 = 42
        _ = fi22
        _ = v29
    }
    for fi30 := 0; fi30 < 7; fi30++ {
        fmt.Println(v21, v21)
        fmt.Println(v21, fi30)
        _ = fi30
    }
    _ = v21
    return 218
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 135
    var v3 uint16 = 654
    if (!(!true)) {
        if true {
            if (42 <= 1) {
                v31 := []int{32768, 42, 0}
                v32 := GenS2{F0: v2, F1: 1}
                v33 := GenS4{F0: -9223372036854775808}
                _ = v31
                _ = v32
                _ = v33
            } else {
                v34 := GenS5{F0: 9223372036854775807, F1: 73, F2: (true || false)}
                var v35 float64 = (0.0 + 1.5)
                _ = v34
                _ = v35
            }
            v1[-1] = 526
            v36 := GenS1{F0: (4294967295 != 826)}
            _ = v36
        }
        var v37 int = 0
        var v38 int16 = int16(v2)
        _ = v37
        _ = v38
    } else {
        v39 := GenS6{F0: -508, F1: 100, F2: 127}
        _ = len(v1)
        for fi40 := 0; fi40 < 8; fi40++ {
            v2 -= (v2 << ((v2) & 63))
            switch "" {
            case "":
                _ = len(v1)
            case "\n\t":
                for fi41 := 0; fi41 < 8; fi41++ {
                    v2 -= (^(v2))
                    v42 := v1[fi40]
                    _ = fi41
                    _ = v42
                }
            case "hello":
                switch "a" {
                case "hello":
                    v3 = 713
                case "a":
                    for fi43 := 0; fi43 < 4; fi43++ {
                        v2 = (v2 % ((genFn0()) | 1))
                        v2 += (v2 + 63)
                        _ = fi43
                    }
                case "foo bar":
                    v2 = (v2 >> ((genFn0()) & 63))
                }
            }
            _ = fi40
        }
        _ = v39
    }
    if ((false && true) && (4294967295 >= 1)) {
        for fi44 := 0; fi44 < 11; fi44++ {
            var v45 uint32 = 4294967295
            if true {
                fmt.Println(v45, fi44, v45)
                v2 = 376
            }
            _ = fi44
            _ = v45
        }
        var v46 uint64 = 42
        _ = v46
    }
    fmt.Println(v2)
    for fi47 := 0; fi47 < 13; fi47++ {
        v48 := GenS6{F0: 1, F1: -86, F2: -42}
        switch "" {
        case "a":
            if (42 > -647) {
                v3 -= uint16(v2)
            }
        }
        if (true && (!(60 != 1000))) {
            for fi49 := 0; fi49 < 15; fi49++ {
                if (true && true) {
                    v50 := []int{(fi49 % (((fi47 + fi49)) | 1)), (fi49 / ((-48) | 1)), -1, (fi49 & (fi47 & fi49))}
                    v3 = 32767
                    _ = v50
                }
                v2 += (^(v2))
                _ = fi49
            }
            delete(v1, fi47)
        }
        _ = fi47
        _ = v48
    }
    v51 := []int{-215, 32767, 32768, 127}
    var v52 bool = (130 <= 9223372036854775807)
    _ = len(v51)
    fmt.Println(v52, v2)
    for fi53 := 0; fi53 < 14; fi53++ {
        for fi54 := 0; fi54 < 11; fi54++ {
            switch ("a" + (("\n\t" + "foo bar") + "foo bar")) {
            case "foo bar":
                if (35 > 100) {
                    if ((^(fi54)) >= 127) {
                        if false {
                            v52 = (v2 <= v2)
                            var v55 bool = (228 != 2147483647)
                            _ = v55
                        } else {
                            v3 = (^(v3))
                            v2 = uint(v3)
                            v51 = append(v51, fi53)
                        }
                        var v56 uint64 = uint64(v2)
                        fmt.Println(v2, fi53)
                        _ = v56
                    }
                } else {
                    _ = len(v51)
                }
            case "a":
                v57 := GenS3{F0: true, F1: 244}
                _ = v57
            case "\n\t":
                var v58 bool = false
                _ = v58
            }
            v59 := GenS4{F0: 720}
            _ = fi54
            _ = v59
        }
        fmt.Println(v3, v52, v2)
        _ = fi53
    }
    v60 := GenS7{F0: 389, F1: 2, F2: 0}
    v52 = false
    v61 := GenS8{F0: v3, F1: 1, F2: (!(!(!(-1 != -19))))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v60
    _ = v61
}
