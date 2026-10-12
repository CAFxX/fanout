package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int64 }
type GenS2 struct { F0 string; F1 uint32; F2 string }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 int16 }
func genFn0(p0 int32, p1 int16, p2 int32) uint32 {
    p2 += p0
    switch "" {
    case "foo bar":
        v7 := []int{-643, 4294967295, 32767, -811}
        _ = v7
    case "":
        v8 := []int{}
        _ = v8
    case "a":
        var v9 int8 = -6
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 bool) uint16 {
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 string = (("" + "") + ("a" + "a"))
            v12 = v12
            p0 = p0
            v12 = (v12 + "hello")
            v13 := []int{}
            if (!((1.5 - 1.5) < 3.14159)) {
                v14 := GenS1{F0: 65535, F1: 1000, F2: 42}
                if p0 {
                    var v15 uint32 = 448
                    _ = v15
                }
                v16 := GenS1{F0: 54, F1: 42, F2: -9223372036854775808}
                _ = v14
                _ = v16
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        var v17 uint64 = 815
        _ = fi10
        _ = v17
    }
    if (-972 != -400) {
        v18 := GenS2{F0: "a", F1: genFn0(519, -530, -2147483648), F2: (("foo bar" + "hello") + "foo bar")}
        if p0 {
            switch 1 {
            case 32767:
                var v19 int8 = -128
                _ = v19
            case 127:
                var v20 uint8 = 42
                _ = v20
            }
        }
        var v21 int8 = 33
        _ = v18
        _ = v21
    } else {
        var v22 uint64 = 18446744073709551615
        fmt.Println(v22)
        _ = v22
    }
    var v23 int32 = 315
    v24 := GenS1{F0: 2, F1: 42, F2: 0}
    _ = p0
    _ = v23
    _ = v24
    return 649
}

func main() {
    var v1 uint = 197
    var v2 int8 = -128
    var v3 int64 = 69
    var v4 int16 = 1000
    var v5 uint8 = 42
    var v6 uint = (^(v1))
    if ((!true) && (!true)) {
        var v25 int8 = v2
        for fi26 := 0; fi26 < 16; fi26++ {
            if true {
                v27 := []int{int(v3)}
                if false {
                    for fi28 := 0; fi28 < 16; fi28++ {
                        v6 -= v6
                        _ = fi28
                    }
                }
                _ = v27
            }
            var v29 uint = v6
            _ = fi26
            _ = v29
        }
        _ = v25
    } else {
        v2 = (-(v2))
        for fi30 := 0; fi30 < 7; fi30++ {
            if (2 != 18446744073709551615) {
                var v31 uint8 = (v5 % (((v5 % ((v5) | 1))) | 1))
                v32 := []int{}
                _ = len(v32)
                _ = v31
                _ = v32
            }
            var v33 int64 = -35
            var v34 int16 = v4
            _ = fi30
            _ = v33
            _ = v34
        }
        fmt.Println(v2)
    }
    for fi35 := 0; fi35 < 13; fi35++ {
        v36 := []int{fi35}
        if (!(((v6 & 255) != (v1 % ((v6) | 1))) || false)) {
            fmt.Println(v2, v3)
        } else {
            switch "hello" {
            case "\n\t":
                v3 += (v3 | v3)
            case "a":
                var v37 int8 = (v2 * v2)
                _ = v37
            case "":
                var v38 string = "a"
                _ = v38
            default:
                if false {
                    v39 := GenS2{F0: "foo bar", F1: genFn0(2147483647, v4, -2147483648), F2: "hello"}
                    _ = v39
                } else {
                    v36 = append(v36, fi35)
                    v40 := GenS3{F0: "hello", F1: (!true)}
                    var v41 uint8 = (v5 | (v5 * 91))
                    _ = v40
                    _ = v41
                }
            }
            var v42 float32 = 0.0
            _ = v42
        }
        _ = fi35
        _ = v36
    }
    switch "\n\t" {
    case "foo bar":
        var v43 uint8 = v5
        _ = v43
    case "\n\t":
        var v44 float64 = 0.0
        _ = v44
    case "a":
        v45 := GenS4{F0: -32768}
        _ = v45
    }
    for fi46 := 0; fi46 < 8; fi46++ {
        v5 += v5
        v47 := GenS1{F0: 2, F1: (-(v4)), F2: int64(v4)}
        _ = fi46
        _ = v47
    }
    v48 := GenS2{F0: "foo bar", F1: 545, F2: "foo bar"}
    var v49 float64 = 3.14159
    v50 := GenS5{F0: 682}
    v51 := GenS3{F0: "\n\t", F1: false}
    v4 += (v4 >> ((v1) & 63))
    switch -463 {
    case 2147483647:
        v52 := []int{-2147483648, 947}
        _ = v52
    default:
        switch 997 {
        case 0:
            if ((v3 ^ v3) <= v3) {
                var v53 uint64 = uint64(v5)
                _ = v53
            } else {
                var v54 bool = true
                v55 := GenS5{F0: 1}
                var v56 uint16 = 65535
                _ = v54
                _ = v55
                _ = v56
            }
        case 2147483647:
            fmt.Println(v6, v2, v49)
        case 42:
            if (72 >= 1) {
                fmt.Println(v6, v6)
                v57 := GenS6{F0: (v5 - (v5 + (v5 << ((v6) & 63)))), F1: v4}
                _ = v57
            } else {
                for fi58 := 0; fi58 < 13; fi58++ {
                    if (202 == (-(fi58))) {
                        if true {
                            v59 := []int{1}
                            _ = v59
                        } else {
                            v1 += v1
                            v4 = (v4 / ((v4) | 1))
                        }
                        fmt.Println(v2, v4, v5)
                        v5 = 255
                    }
                    for fi60 := 0; fi60 < 10; fi60++ {
                        v61 := []int{int(v2), 32768, -9223372036854775808}
                        var v62 string = ("" + "")
                        _ = fi60
                        _ = v61
                        _ = v62
                    }
                    _ = fi58
                }
                v1 = 65535
            }
        }
    }
    v63 := []int{128, 4294967295, 2147483647, 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v63:", len(v63), v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v63
}
