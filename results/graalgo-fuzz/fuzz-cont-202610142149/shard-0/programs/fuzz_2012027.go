package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 string; F2 uint }
type GenS4 struct { F0 uint32; F1 uint32; F2 string }
type GenS5 struct { F0 uint8; F1 bool }
func genFn0() int32 {
    var v4 uint8 = 127
    var v5 uint32 = 65535
    if (0 > 32768) {
        var v6 int = -201
        if (4294967295 >= 42) {
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := []int{}
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 8; fi9++ {
                v4 += (v4 - (^(v4)))
                _ = fi9
            }
        }
        var v10 uint64 = 1
        _ = v10
        _ = v6
    } else {
        var v11 int8 = 1
        v12 := []int{42, 9223372036854775807}
        _ = v11
        _ = v12
    }
    switch ("foo bar" + "a") {
    case "":
        switch "hello" {
        case "\n\t":
            v4 = v4
        }
    case "a":
        v13 := GenS1{F0: (v5 / ((2147483648) | 1)), F1: ("hello" < "hello"), F2: 42}
        _ = v13
    case "foo bar":
        v14 := GenS2{F0: 1}
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        if (457 <= 65535) {
            fmt.Println(v4, fi15, v4)
            switch fi15 {
            case -248:
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 int = -685
                    _ = fi16
                    _ = v17
                }
            case -2147483648:
                v18 := []int{fi15, fi15, (fi15 - (fi15 / ((fi15) | 1)))}
                _ = v18
            case 32767:
                v19 := []int{(fi15 % ((fi15) | 1))}
                _ = v19
            }
        } else {
            switch (^(fi15)) {
            case 42:
                fmt.Println(v4)
            case 9223372036854775807:
                var v20 int = (fi15 | fi15)
                _ = v20
            }
            var v21 uint64 = 0
            _ = v21
        }
        _ = fi15
    }
    _ = v4
    _ = v5
    return -27
}

func main() {
    var v1 string = "hello"
    var v2 float64 = (-2.25 - float64(((0.0 * -2.25) / 1.0)))
    var v3 uint64 = 18446744073709551615
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int8 = -56
        _ = fi22
        _ = v23
    }
    switch v1 {
    case "a":
        var v24 uint64 = v3
        _ = v24
    case "\n\t":
        if (32767 == -32768) {
            if true {
                v25 := GenS1{F0: 65535, F1: (!true), F2: 1}
                var v26 uint64 = 42
                _ = v25
                _ = v26
            }
            v27 := GenS3{F0: int64(v3), F1: "\n\t", F2: 18446744073709551615}
            _ = v27
        } else {
            var v28 int16 = 32767
            v3 -= 0
            v29 := GenS3{F0: 0, F1: v1, F2: 850}
            _ = v28
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        switch fi30 {
        case 42:
            var v31 uint = uint(v3)
            _ = v31
        case -9223372036854775808:
            for fi32 := 0; fi32 < 14; fi32++ {
                if false {
                    switch fi32 {
                    case -1:
                        if false {
                            v2 += v2
                            var v33 int32 = 2147483647
                            v34 := GenS4{F0: 2147483648, F1: 0, F2: v1}
                            _ = v33
                            _ = v34
                        } else {
                            v35 := GenS4{F0: 2, F1: 2147483647, F2: ((("foo bar" + (v1 + v1)) + v1) + ((v1 + (("foo bar" + v1) + ((v1 + v1) + ("a" + ("foo bar" + v1))))) + "a"))}
                            fmt.Println(v2)
                            if (!(!(true || true))) {
                                v36 := []int{816}
                                _ = v36
                            } else {
                                if false {
                                    v37 := GenS3{F0: -1, F1: v1, F2: 4294967295}
                                    var v38 float32 = 1e10
                                    var v39 float64 = ((v2 + float64((1.5 * 1.5))) - (((1e10 + float64((((v38 / -2.25) / 0.1) - v38))) - (((1.5 - v2) + (v2 / (0.1 / 1.0))) - v2)) * 3.14159))
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                } else {
                                    var v40 int16 = -1
                                    if (false && (0 == 127)) {
                                        if (!false) {
                                            v3 = (^(v3))
                                            v2 += 1e10
                                            v2 = 3.14159
                                        }
                                    }
                                    var v41 uint64 = (v3 & 18446744073709551615)
                                    _ = v40
                                    _ = v41
                                }
                            }
                            _ = v35
                        }
                    }
                    if (1 > 1) {
                        var v42 int64 = 42
                        _ = v42
                    }
                }
                v43 := []int{}
                _ = fi32
                _ = v43
            }
        }
        fmt.Println(v3, v1, v1)
        v44 := []int{(fi30 % ((fi30) | 1)), (fi30 + (fi30 % ((fi30) | 1)))}
        _ = fi30
        _ = v44
    }
    fmt.Println(v1, v1, v2)
    v45 := []int{127, 2147483647}
    v46 := GenS3{F0: 9223372036854775807, F1: "", F2: 42}
    v47 := GenS4{F0: 65535, F1: 1, F2: v1}
    if (2 >= (v3 * v3)) {
        for fi48 := 0; fi48 < 15; fi48++ {
            v45 = append(v45, (fi48 << ((42) & 63)))
            _ = fi48
        }
        for fi49 := 0; fi49 < 11; fi49++ {
            var v50 int64 = 42
            v51 := GenS1{F0: 4294967295, F1: (false && (18446744073709551615 != (v3 >> ((4294967295) & 63)))), F2: 255}
            v52 := GenS2{F0: fi49}
            _ = fi49
            _ = v50
            _ = v51
            _ = v52
        }
    } else {
        v3 -= v3
        v1 = v1
        var v53 int32 = -1
        _ = v53
    }
    var v54 int32 = 1
    _ = len(v45)
    var v55 uint = 65535
    for fi56 := 0; fi56 < 18; fi56++ {
        if false {
            for fi57 := 0; fi57 < 18; fi57++ {
                var v58 int32 = (v54 ^ v54)
                switch ((v1 + v1) + (v1 + (v1 + v1))) {
                case "foo bar":
                    v55 = 187
                case "a":
                    v2 = (float64(0.0) / (v2 * (v2 + (v2 + (3.14159 * v2)))))
                }
                fmt.Println(v1)
                _ = fi57
                _ = v58
            }
        } else {
            var v59 float32 = 3.14159
            _ = v59
        }
        for fi60 := 0; fi60 < 8; fi60++ {
            v2 += v2
            _ = fi60
        }
        v61 := GenS1{F0: 42, F1: false, F2: 32767}
        _ = fi56
        _ = v61
    }
    if (1.5 > float32(0.1)) {
        v62 := GenS2{F0: -193}
        var v63 int8 = int8(v3)
        _ = v62
        _ = v63
    } else {
        v64 := GenS5{F0: 42, F1: true}
        switch 1 {
        case 2147483647:
            for fi65 := 0; fi65 < 3; fi65++ {
                v3 += v3
                fmt.Println(v3, v55)
                _ = fi65
            }
        }
        v2 = 0.1
        _ = v64
    }
    if len(v45) >= 2 {
        v45 = v45[1:len(v45)]
    }
    v45 = append(v45, 128)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v47
    _ = v54
    _ = v55
}
