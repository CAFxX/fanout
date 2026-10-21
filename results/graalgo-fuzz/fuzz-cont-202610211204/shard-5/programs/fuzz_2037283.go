package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int64; F2 string }
type GenS3 struct { F0 int8; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint8; F1 int32 }
func genFn0() uint16 {
    var v6 uint64 = 0
    if true {
        if false {
            var v7 int64 = 42
            _ = v7
        }
    }
    switch 40 {
    case 0:
        if false {
            v6 += (v6 << ((1) & 63))
            switch "\n\t" {
            case "a":
                fmt.Println(v6)
            case "hello":
                var v8 float32 = 1.5
                _ = v8
            case "\n\t":
                var v9 bool = (false || (false || (2147483647 >= -432)))
                _ = v9
            default:
                v6 -= (v6 / ((v6) | 1))
            }
        }
    case 949:
        var v10 int64 = -9223372036854775808
        _ = v10
    case 9223372036854775807:
        v11 := GenS1{F0: 255}
        _ = v11
    default:
        var v12 float64 = -2.25
        _ = v12
    }
    _ = v6
    return 0
}

func main() {
    var v1 int = -250
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 uint16 = 42
    var v5 int64 = 42
    v1 -= (v1 >> ((2) & 63))
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            if (836 != fi13) {
                delete(v3, 32767)
                var v15 int8 = -124
                _ = v15
            } else {
                var v16 int16 = 42
                v4 = v4
                _ = v16
            }
            switch fi14 {
            case 625:
                var v17 uint8 = 112
                _ = v17
            }
            _ = fi14
        }
        if (3.14159 <= 1.5) {
            var v18 float32 = 1.5
            _ = v18
        }
        _ = fi13
    }
    v19 := GenS2{F0: v1, F1: (v5 - (-(v5))), F2: ((("hello" + "hello") + ((((((("\n\t" + ("a" + "a")) + ("\n\t" + ("\n\t" + "a"))) + ("" + "hello")) + ("hello" + "\n\t")) + ("" + "hello")) + "hello") + "a")) + "\n\t")}
    v20 := []int{-9223372036854775808, v1, v1, (v1 >> ((644) & 63))}
    v21 := v2[2147483647]
    var v22 uint64 = 5
    var v23 uint = 850
    for fi24 := 0; fi24 < 7; fi24++ {
        fmt.Println(v23)
        if (-122 == -1) {
            v5 += (v5 - -690)
        } else {
            if true {
                var v25 int16 = -734
                _ = len(v20)
                _ = v25
            } else {
                var v26 int16 = 1000
                for fi27 := 0; fi27 < 10; fi27++ {
                    v28 := GenS2{F0: (v1 & int(v23)), F1: (^(v5)), F2: "hello"}
                    v1 = fi27
                    _ = fi27
                    _ = v28
                }
                var v29 int64 = (v5 | v5)
                _ = v26
                _ = v29
            }
            v30 := GenS3{F0: -118, F1: true, F2: 2}
            if false {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                var v31 uint = v23
                if (!false) {
                    switch "a" {
                    case "foo bar":
                        switch -9223372036854775808 {
                        case -9223372036854775808:
                            var v32 int16 = -1
                            _ = v32
                        case 42:
                            for fi33 := 0; fi33 < 17; fi33++ {
                                if (-2147483648 > -545) {
                                    fmt.Println(v31, v23)
                                    fmt.Println(fi24, v5, v1)
                                    switch "\n\t" {
                                    case "hello":
                                        v34 := GenS4{F0: 127, F1: int32(v4)}
                                        _ = v34
                                    case "a":
                                        if false {
                                            var v35 int16 = 726
                                            _ = v35
                                        }
                                    default:
                                        var v36 string = "a"
                                        _ = v36
                                    }
                                } else {
                                    delete(v3, -9223372036854775808)
                                    var v37 uint32 = 893
                                    var v38 uint = uint(v4)
                                    _ = v37
                                    _ = v38
                                }
                                _ = fi33
                            }
                        default:
                            fmt.Println(v4)
                        }
                    case "hello":
                        v20 = append(v20, fi24)
                    case "a":
                        for fi39 := 0; fi39 < 15; fi39++ {
                            switch 992 {
                            case 42:
                                v40 := v2[uint32(v31)]
                                _ = v40
                            }
                            _ = fi39
                        }
                    default:
                        var v41 uint64 = v22
                        _ = v41
                    }
                    _ = len(v3)
                }
                _ = v31
            } else {
                v2[2] = v21
                v21 += 2147483647
            }
            _ = v30
        }
        var v42 int = fi24
        _ = fi24
        _ = v42
    }
    if false {
        if len(v20) > 0 {
            v20[0] = -9223372036854775808
        }
    } else {
        switch v21 {
        case 2147483647:
            if true {
                for fi43 := 0; fi43 < 19; fi43++ {
                    var v44 uint16 = (v4 * 2)
                    _ = fi43
                    _ = v44
                }
            } else {
                var v45 uint8 = 128
                fmt.Println(v4, v5, v1)
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                _ = v45
            }
        case 4294967295:
            v5 -= (-(v5))
        case -9223372036854775808:
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
        default:
            v5 = v5
        }
        v4 += v4
    }
    if ("foo bar" > "a") {
        fmt.Println(v4, v4, v5)
        fmt.Println(v4, v23)
    } else {
        var v46 int8 = 100
        for fi47 := 0; fi47 < 14; fi47++ {
            var v48 uint = (v23 - (^(v23)))
            if true {
                fmt.Println(fi47, v23)
                fmt.Println(v21, v21)
                fmt.Println(v22, v46)
            } else {
                v46 += 0
            }
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                v22 = v22
            default:
                var v49 string = (("\n\t" + ("hello" + "hello")) + "foo bar")
                _ = v49
            }
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
