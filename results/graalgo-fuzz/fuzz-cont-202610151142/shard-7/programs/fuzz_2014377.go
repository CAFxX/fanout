package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 string; F1 int16; F2 bool }
func genFn0(p0 int64, p1 int64, p2 string) int32 {
    p2 = (p2 + (p2 + (p2 + "")))
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 7; fi6++ {
        p1 += (^(p1))
        p0 += p0
        _ = fi6
    }
    if false {
        if (!true) {
            for fi7 := 0; fi7 < 10; fi7++ {
                p2 = p2
                v8 := GenS1{F0: 275, F1: 1}
                p2 = ("" + p2)
                _ = fi7
                _ = v8
            }
            v9 := GenS2{F0: "", F1: 32767}
            _ = v9
        }
        if false {
            var v10 uint64 = 42
            _ = v10
        } else {
            v11 := GenS3{F0: 108}
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1() uint64 {
    v12 := []int{}
    switch ("" + "\n\t") {
    case "":
        var v13 uint = 2
        _ = v13
    case "\n\t":
        if false {
            if len(v12) > 0 {
                v12[0] = 796
            }
            var v14 int64 = 9223372036854775807
            var v15 uint8 = 1
            _ = v14
            _ = v15
        } else {
            if (true || false) {
                if true {
                    var v16 uint32 = 0
                    fmt.Println(v16)
                    if true {
                        for fi17 := 0; fi17 < 4; fi17++ {
                            v16 = v16
                            for fi18 := 0; fi18 < 4; fi18++ {
                                v12 = append(v12, 644)
                                _ = fi18
                            }
                            _ = fi17
                        }
                        v19 := GenS2{F0: ("foo bar" + ""), F1: 498}
                        v12 = append(v12, -627)
                        _ = v19
                    } else {
                        if false {
                            if ((!(-1 >= -65)) && true) {
                                for fi20 := 0; fi20 < 6; fi20++ {
                                    v16 = 2
                                    for fi21 := 0; fi21 < 10; fi21++ {
                                        var v22 uint = 613
                                        _ = fi21
                                        _ = v22
                                    }
                                    v23 := GenS3{F0: 127}
                                    _ = fi20
                                    _ = v23
                                }
                                var v24 string = (("\n\t" + (("a" + "a") + "foo bar")) + "\n\t")
                                _ = v24
                            }
                            v16 = 2147483648
                            var v25 uint8 = 255
                            _ = v25
                        } else {
                            v16 = v16
                            if (!(-1 < -1)) {
                                v16 = (v16 - (v16 ^ 65535))
                            } else {
                                v16 -= 1
                            }
                            if len(v12) > 0 {
                                v12[0] = 0
                            }
                        }
                        v16 -= 449
                    }
                    _ = v16
                } else {
                    var v26 uint8 = 0
                    for fi27 := 0; fi27 < 12; fi27++ {
                        if false {
                            var v28 uint16 = 42
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            var v29 string = ""
                            _ = v28
                            _ = v29
                        }
                        _ = fi27
                    }
                    for fi30 := 0; fi30 < 7; fi30++ {
                        v26 = v26
                        _ = fi30
                    }
                    _ = v26
                }
                v31 := GenS4{F0: 9223372036854775807}
                v32 := GenS1{F0: 255, F1: -1}
                _ = v31
                _ = v32
            }
        }
    case "foo bar":
        var v33 uint = 4294967295
        _ = v33
    default:
        var v34 uint = 785
        _ = v34
    }
    var v35 uint32 = 770
    v36 := GenS1{F0: 32767, F1: 1}
    _ = v12
    _ = v35
    _ = v36
    return 0
}

func main() {
    var v1 bool = (-1 < 261)
    var v2 int64 = 0
    var v3 uint16 = 548
    var v4 uint16 = 177
    var v5 int8 = -128
    for fi37 := 0; fi37 < 17; fi37++ {
        var v38 uint64 = 911
        var v39 string = ("a" + "")
        fmt.Println(v1, v3, v3)
        _ = fi37
        _ = v38
        _ = v39
    }
    v1 = true
    v40 := GenS5{F0: 255, F1: -128}
    v41 := GenS3{F0: (v5 - 93)}
    fmt.Println(v4)
    for fi42 := 0; fi42 < 10; fi42++ {
        switch fi42 {
        case 2147483647:
            fmt.Println(v5, v3, v2)
        case -699:
            var v43 uint8 = 42
            _ = v43
        default:
            v5 -= int8(v2)
        }
        _ = fi42
    }
    var v44 uint8 = 1
    if ((!true) || (127 <= -352)) {
        for fi45 := 0; fi45 < 18; fi45++ {
            fmt.Println(v1, v5, v5)
            var v46 string = "a"
            _ = fi45
            _ = v46
        }
        v47 := GenS1{F0: v3, F1: -37}
        v48 := GenS6{F0: "hello", F1: 1000, F2: (!false)}
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v44
    _ = v5
}
