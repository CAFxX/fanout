package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS6 struct { F0 string }
func genFn0() bool {
    v7 := GenS1{F0: 1, F1: 14}
    var v8 uint32 = 0
    switch "\n\t" {
    case "foo bar":
        if (0 < 564) {
            for fi9 := 0; fi9 < 1; fi9++ {
                for fi10 := 0; fi10 < 9; fi10++ {
                    if ((!(((1.5 * -2.25) >= (0.1 * 0.0)) && true)) && (-258 == (fi9 % ((574) | 1)))) {
                        switch -9223372036854775808 {
                        case 1:
                            v11 := []int{(fi10 * 2147483647), -9223372036854775808, (fi10 ^ fi9)}
                            _ = v11
                        }
                    } else {
                        fmt.Println(v8, v8, v8)
                        v8 -= v8
                    }
                    _ = fi10
                }
                _ = fi9
            }
            v8 -= (v8 & 2)
        }
    case "":
        v12 := []int{219}
        _ = v12
    default:
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := []int{1, (fi13 >> ((42) & 63)), 9223372036854775807}
            if false {
                v8 = v8
            } else {
                if true {
                    fmt.Println(v8, fi13)
                    v15 := GenS2{F0: 389}
                    v8 = v8
                    _ = v15
                } else {
                    var v16 int = fi13
                    _ = v16
                }
            }
            _ = fi13
            _ = v14
        }
    }
    _ = v7
    _ = v8
    return ((true && false) && true)
}
func genFn1() int64 {
    var v17 uint = 0
    var v18 uint32 = 42
    if ((20 == 185) || true) {
        v19 := GenS3{F0: "\n\t"}
        var v20 int8 = 35
        _ = v19
        _ = v20
    } else {
        for fi21 := 0; fi21 < 5; fi21++ {
            fmt.Println(fi21, v18, fi21)
            v22 := []int{(fi21 % (((fi21 * (fi21 / ((fi21) | 1)))) | 1)), -442, fi21}
            v23 := GenS4{F0: 255}
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    _ = v17
    _ = v18
    return 113
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint16 = 925
    var v3 int8 = 108
    var v4 uint32 = 980
    v5 := []int{-246, 0, -9223372036854775808}
    v6 := make(map[int]int)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 835)
    v24 := GenS5{F0: v1, F1: 32767, F2: -9223372036854775808}
    for fi25 := 0; fi25 < 9; fi25++ {
        if false {
            v26 := GenS3{F0: "a"}
            for fi27 := 0; fi27 < 10; fi27++ {
                v4 = (^(v4))
                delete(v6, -9223372036854775808)
                var v28 uint16 = v2
                _ = fi27
                _ = v28
            }
            _ = v26
        } else {
            for fi29 := 0; fi29 < 18; fi29++ {
                v6[fi25] = 500
                delete(v6, fi25)
                v30 := GenS4{F0: 42}
                _ = fi29
                _ = v30
            }
            if true {
                for fi31 := 0; fi31 < 13; fi31++ {
                    var v32 uint = 639
                    if (v1 < v1) {
                        v32 += (v32 % (((v32 % ((1) | 1))) | 1))
                        v1 += v1
                        var v33 string = "hello"
                        _ = v33
                    } else {
                        var v34 string = ""
                        var v35 int64 = -1
                        fmt.Println(v32)
                        _ = v34
                        _ = v35
                    }
                    _ = fi31
                    _ = v32
                }
                v2 += (^(v2))
            }
            var v36 int = (fi25 << ((18446744073709551615) & 63))
            _ = v36
        }
        for fi37 := 0; fi37 < 3; fi37++ {
            switch fi37 {
            case 1:
                if true {
                    if false {
                        if (true || (false && ((^(v2)) >= v2))) {
                            v4 -= (v4 - (v4 % ((v4) | 1)))
                            v38 := GenS3{F0: "foo bar"}
                            _ = v38
                        }
                        if false {
                            v39 := GenS4{F0: 255}
                            _ = v39
                        } else {
                            var v40 int16 = 42
                            _ = v40
                        }
                    } else {
                        v3 += (v3 & 63)
                        if len(v5) > 0 {
                            v5[0] = (-(fi25))
                        }
                        fmt.Println(fi37)
                    }
                    fmt.Println(v3, fi25)
                }
            case 42:
                if (-779 == 0) {
                    v41 := v6[(fi37 | 0)]
                    var v42 bool = (!false)
                    if v42 {
                        fmt.Println(fi25)
                        if (0 == genFn1()) {
                            v43 := GenS4{F0: 100}
                            delete(v6, (fi25 - (fi37 * -9223372036854775808)))
                            _ = len(v5)
                            _ = v43
                        }
                        v41 += 2147483647
                    } else {
                        var v44 int64 = 42
                        v1 = (^(v1))
                        _ = v44
                    }
                    _ = v41
                    _ = v42
                }
            }
            var v45 uint16 = v2
            v2 = v45
            _ = fi37
            _ = v45
        }
        _ = fi25
    }
    switch ("\n\t" + ("hello" + "")) {
    case "a":
        var v46 int64 = 0
        _ = v46
    case "\n\t":
        fmt.Println(v4, v4, v3)
    default:
        var v47 uint64 = 18446744073709551615
        _ = v47
    }
    v48 := v6[-2147483648]
    if false {
        v3 = v3
        var v49 bool = (true || (18446744073709551615 > 18446744073709551615))
        switch 9223372036854775807 {
        case 127:
            v50 := GenS5{F0: v1, F1: 32767, F2: 1}
            _ = v50
        case -9223372036854775808:
            v49 = false
        default:
            var v51 int = v48
            _ = v51
        }
        _ = v49
    }
    switch (("a" + ("foo bar" + "a")) + "hello") {
    case "\n\t":
        v52 := GenS5{F0: v1, F1: v2, F2: genFn1()}
        _ = v52
    case "hello":
        if true {
            var v53 uint8 = 178
            var v54 int8 = v3
            _ = v53
            _ = v54
        }
    case "a":
        v5 = append(v5, int(v4))
    default:
        v48 = (v48 / ((-2147483648) | 1))
    }
    var v55 uint64 = v1
    switch "" {
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = (v48 >> ((0) & 63))
        }
    case "":
        var v56 int = (v48 % ((v48) | 1))
        _ = v56
    case "hello":
        var v57 uint16 = v2
        _ = v57
    }
    v58 := GenS6{F0: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v55
    _ = v58
    _ = v6
}
