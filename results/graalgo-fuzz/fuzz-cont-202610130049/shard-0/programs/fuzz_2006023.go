package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int16; F1 int64; F2 int8 }
type GenS4 struct { F0 uint8; F1 string; F2 uint }
func genFn0(p0 uint64, p1 uint64) uint {
    v7 := GenS1{F0: ("hello" + "")}
    v8 := GenS2{F0: (p0 + (p1 << ((1) & 63))), F1: 0}
    if ("hello" <= "a") {
        switch "\n\t" {
        case "hello":
            v9 := []int{-9223372036854775808, 605, 59}
            _ = v9
        case "":
            if true {
                for fi10 := 0; fi10 < 15; fi10++ {
                    v11 := []int{fi10, -355, 42}
                    fmt.Println(p0, p1)
                    if (!true) {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        v12 := GenS1{F0: "foo bar"}
                        _ = v12
                    } else {
                        switch 32768 {
                        case 127:
                            for fi13 := 0; fi13 < 9; fi13++ {
                                fmt.Println(fi13, fi10, fi13)
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                switch fi13 {
                                case 9223372036854775807:
                                    v14 := GenS1{F0: ("hello" + (((("\n\t" + (("\n\t" + "foo bar") + (("foo bar" + "foo bar") + ("hello" + "")))) + "") + "foo bar") + "a"))}
                                    _ = v14
                                case 0:
                                    fmt.Println(fi13, fi13)
                                default:
                                    var v15 uint = 18446744073709551615
                                    _ = v15
                                }
                                var v16 int8 = -37
                                _ = fi13
                                _ = v16
                            }
                        case 0:
                            var v17 float64 = -2.25
                            _ = v17
                        default:
                            p1 -= (p1 >> ((18446744073709551615) & 63))
                        }
                    }
                    _ = fi10
                    _ = v11
                }
            }
        }
        v18 := GenS2{F0: p1, F1: 1}
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 990
}
func genFn1(p0 uint64, p1 int8) int8 {
    v19 := []int{128, 128}
    v20 := GenS1{F0: ""}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 100
}

func main() {
    var v1 uint32 = 1
    v2 := []int{-319, 4294967295, -2147483648, 627}
    var v3 int32 = 1
    var v4 string = "hello"
    var v5 uint32 = v1
    var v6 bool = ((true || false) && true)
    fmt.Println(v5, v6, v1)
    var v21 int32 = (v3 - (-(v3)))
    if len(v2) > 0 {
        v2[0] = 0
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if true {
            for fi23 := 0; fi23 < 2; fi23++ {
                fmt.Println(v1, v5, fi23)
                if ((^(fi22)) > fi22) {
                    v4 = v4
                } else {
                    _ = len(v2)
                    fmt.Println(fi23, v6, v5)
                    v24 := GenS3{F0: 527, F1: 0, F2: 0}
                    _ = v24
                }
                _ = fi23
            }
        }
        v5 = v1
        v3 += -1
        _ = fi22
    }
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v25 int32 = -2147483648
        _ = v25
    } else {
        fmt.Println(v21, v1, v6)
        var v26 int8 = int8(v1)
        _ = v26
    }
    var v27 int = 127
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v28 int8 = 0
    var v29 uint16 = 42
    var v30 uint16 = (v29 ^ 608)
    for fi31 := 0; fi31 < 19; fi31++ {
        switch 127 {
        case 9223372036854775807:
            _ = len(v2)
        case 32768:
            if v6 {
                v4 = v4
                fmt.Println(v27, v28)
                for fi32 := 0; fi32 < 11; fi32++ {
                    v27 -= (v27 >> ((18446744073709551615) & 63))
                    _ = fi32
                }
            }
        }
        var v33 int64 = -9223372036854775808
        for fi34 := 0; fi34 < 5; fi34++ {
            v5 = v1
            var v35 float32 = -2.25
            v36 := GenS4{F0: 1, F1: v4, F2: 28}
            _ = len(v2)
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi31
        _ = v33
    }
    if v6 {
        for fi37 := 0; fi37 < 6; fi37++ {
            for fi38 := 0; fi38 < 19; fi38++ {
                v39 := GenS3{F0: 0, F1: 440, F2: (v28 << ((uint(v30)) & 63))}
                fmt.Println(v21, v30, v3)
                _ = fi38
                _ = v39
            }
            var v40 int = (fi37 & v27)
            v28 = -84
            _ = fi37
            _ = v40
        }
        var v41 int16 = 32767
        if false {
            v3 -= -570
            var v42 int = (v27 % (((^(v27))) | 1))
            var v43 uint16 = (v30 ^ (v30 & v29))
            _ = v42
            _ = v43
        }
        _ = v41
    } else {
        if len(v2) > 0 {
            v2[0] = (v27 - v27)
        }
        var v44 uint32 = v1
        _ = v44
    }
    switch (^(v27)) {
    case -1:
        v45 := GenS2{F0: 4294967295, F1: 1}
        _ = v45
    case -580:
        if (18446744073709551615 == 4294967295) {
            v5 = v5
            if len(v2) > 0 {
                v2[0] = v27
            }
        }
    case 296:
        var v46 int8 = v28
        _ = v46
    }
    var v47 int16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
