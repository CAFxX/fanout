package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16 }
func genFn0(p0 uint8) uint8 {
    p0 -= 128
    v7 := []int{}
    v8 := GenS1{F0: 0, F1: 836}
    p0 = (^(p0))
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(p0, p0)
        fmt.Println(p0)
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return (p0 | 2)
}

func main() {
    var v1 uint = 2
    var v2 float32 = -2.25
    v3 := make(map[string]bool)
    v4 := make(map[uint32]string)
    var v5 uint16 = 255
    v6 := make(map[int]string)
    if (!false) {
        var v10 int16 = -1
        var v11 string = ("" + "foo bar")
        _ = v10
        _ = v11
    } else {
        switch "a" {
        case "foo bar":
            v5 = (v5 % ((v5) | 1))
        }
        var v12 string = (("a" + "") + ("a" + "a"))
        _ = v12
    }
    v2 = (3.14159 / -2.25)
    if ((3.14159 >= 1.5) && (true || (18446744073709551615 < 2))) {
        for fi13 := 0; fi13 < 2; fi13++ {
            v14 := GenS2{F0: "hello"}
            for fi15 := 0; fi15 < 10; fi15++ {
                if true {
                    var v16 uint32 = 2
                    v17 := GenS2{F0: ((("hello" + (("hello" + "hello") + "foo bar")) + ("" + "\n\t")) + "")}
                    if false {
                        var v18 uint16 = v5
                        v18 += (^(v18))
                        _ = v18
                    }
                    _ = v16
                    _ = v17
                } else {
                    v5 += v5
                    fmt.Println(fi15)
                }
                switch 2147483647 {
                case 499:
                    if true {
                        v1 -= (v1 + v1)
                        v1 = 4294967295
                    } else {
                        v2 += v2
                    }
                case 1:
                    v19 := GenS1{F0: v5, F1: 901}
                    _ = v19
                case -2147483648:
                    v20 := []int{fi13, -908, (fi15 * (^(fi15))), (fi13 / ((fi15) | 1))}
                    _ = v20
                }
                _ = fi15
            }
            _ = fi13
            _ = v14
        }
    } else {
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 int = (fi21 >> (((^(v1))) & 63))
            var v23 int32 = -954
            v24 := GenS3{F0: (v23 - 0)}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
    }
    v25 := GenS4{F0: -722}
    for fi26 := 0; fi26 < 15; fi26++ {
        v27 := []int{(fi26 >> ((65535) & 63)), (fi26 & fi26)}
        _ = fi26
        _ = v27
    }
    var v28 string = ("a" + (("" + ("foo bar" + ("\n\t" + "hello"))) + ("\n\t" + "a")))
    for fi29 := 0; fi29 < 13; fi29++ {
        v30 := []int{(fi29 * (fi29 ^ (fi29 * 1))), 0, (fi29 % ((fi29) | 1)), (-(fi29))}
        v31 := GenS3{F0: 0}
        _ = fi29
        _ = v30
        _ = v31
    }
    v1 = v1
    if false {
        v2 = 0.0
        if false {
            var v32 int16 = 0
            _ = v32
        }
    } else {
        var v33 int = 32768
        var v34 int8 = -128
        _ = v33
        _ = v34
    }
    v35 := GenS4{F0: 870}
    var v36 uint8 = genFn0(166)
    v37 := GenS3{F0: int32(v5)}
    if (v1 <= v1) {
        fmt.Println(v1)
    } else {
        fmt.Println(v1, v36)
        v38 := GenS4{F0: int16(v36)}
        for fi39 := 0; fi39 < 11; fi39++ {
            for fi40 := 0; fi40 < 10; fi40++ {
                v2 = 0.0
                if (856 != 293) {
                    if false {
                        var v41 uint64 = 42
                        _ = v41
                    } else {
                        if (314 < 2147483648) {
                            v42 := v3["hello"]
                            fmt.Println(fi39)
                            _ = v42
                        } else {
                            switch ("a" + "foo bar") {
                            case "\n\t":
                                var v43 uint8 = (v36 << ((v1) & 63))
                                _ = v43
                            case "foo bar":
                                v28 = "a"
                            default:
                                v44 := GenS3{F0: 42}
                                _ = v44
                            }
                            switch -506 {
                            case 276:
                                v45 := v4[779]
                                _ = v45
                            case -9223372036854775808:
                                v46 := GenS3{F0: 0}
                                _ = v46
                            default:
                                v5 += (v5 % ((v5) | 1))
                            }
                        }
                    }
                    var v47 float32 = v2
                    v28 = v28
                    switch -984 {
                    case 9223372036854775807:
                        v1 = 42
                    }
                    var v48 float64 = (1.5 - (3.14159 / -2.25))
                    _ = v47
                    _ = v48
                } else {
                    var v49 int8 = 100
                    var v50 uint16 = v5
                    _ = v49
                    _ = v50
                }
                _ = fi40
            }
            _ = fi39
        }
        _ = v38
    }
    var v51 string = v28
    var v52 int8 = 81
    v53 := GenS4{F0: 1000}
    var v54 uint8 = (v36 >> (((v1 / ((1) | 1))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v6
}
