package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64; F1 int; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int8; F1 uint }
func genFn0() string {
    if false {
        var v7 int32 = 42
        if (0.1 <= -2.25) {
            var v8 int32 = (v7 + v7)
            _ = v8
        } else {
            if (!true) {
                v9 := []int{727, 811}
                _ = v9
            } else {
                v10 := GenS1{F0: -895, F1: 545}
                switch "a" {
                case "\n\t":
                    for fi11 := 0; fi11 < 17; fi11++ {
                        var v12 int8 = 34
                        var v13 string = (("foo bar" + ("" + "")) + "hello")
                        var v14 float32 = 1e10
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                case "hello":
                    fmt.Println(v7)
                case "":
                    v15 := GenS1{F0: 915, F1: 0}
                    _ = v15
                default:
                    for fi16 := 0; fi16 < 4; fi16++ {
                        var v17 string = "\n\t"
                        _ = fi16
                        _ = v17
                    }
                }
                _ = v10
            }
            if (!false) {
                switch 9223372036854775807 {
                case -2147483648:
                    var v18 uint64 = 0
                    _ = v18
                case 601:
                    v19 := GenS1{F0: -260, F1: 0}
                    _ = v19
                }
                v20 := []int{929, -78}
                _ = len(v20)
                _ = v20
            } else {
                v7 -= v7
                v21 := []int{}
                switch "a" {
                case "a":
                    v22 := GenS1{F0: -124, F1: 1000}
                    _ = v22
                }
                _ = v21
            }
            v23 := GenS1{F0: int(v7), F1: 60}
            _ = v23
        }
        v7 += (v7 ^ v7)
        _ = v7
    } else {
        var v24 int64 = -510
        v25 := GenS2{F0: "foo bar"}
        fmt.Println(v24)
        _ = v24
        _ = v25
    }
    v26 := GenS3{F0: 790}
    _ = v26
    return "\n\t"
}
func genFn1(p0 uint, p1 int16, p2 int) uint8 {
    v27 := GenS2{F0: (("" + "\n\t") + "hello")}
    v28 := GenS3{F0: (p0 >> ((p0) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    return 1
}

func main() {
    var v1 uint32 = 788
    var v2 int = -89
    var v3 int16 = 1
    var v4 float64 = -2.25
    var v5 int = v2
    var v6 int16 = v3
    var v29 int8 = 127
    v1 = (v1 << ((4294967295) & 63))
    var v30 int32 = -71
    if false {
        var v31 int32 = v30
        _ = v31
    }
    fmt.Println(v6, v5)
    for fi32 := 0; fi32 < 13; fi32++ {
        var v33 bool = true
        _ = fi32
        _ = v33
    }
    switch (v2 << ((357) & 63)) {
    case 42:
        v34 := GenS4{F0: 240, F1: int(v1), F2: v1}
        _ = v34
    case 32768:
        if (true || (!true)) {
            switch ("hello" + "") {
            case "":
                if (false && true) {
                    switch (v5 - v2) {
                    case -2147483648:
                        v35 := []int{(v5 & v5), (v5 >> ((711) & 63))}
                        _ = v35
                    case 32768:
                        var v36 int64 = -9223372036854775808
                        _ = v36
                    case 128:
                        var v37 float64 = (v4 - float64((1e10 + 0.0)))
                        _ = v37
                    default:
                        fmt.Println(v1)
                    }
                } else {
                    fmt.Println(v3, v29, v2)
                    v38 := GenS3{F0: 2}
                    if false {
                        var v39 int8 = (v29 - v29)
                        var v40 uint32 = 1
                        v41 := GenS5{F0: v5}
                        _ = v39
                        _ = v40
                        _ = v41
                    }
                    _ = v38
                }
            case "foo bar":
                if false {
                    var v42 uint8 = 2
                    _ = v42
                }
            }
            v43 := []int{v5, int(v30), v5, v2}
            _ = v43
        } else {
            for fi44 := 0; fi44 < 6; fi44++ {
                v45 := []int{127, (v5 * int(v29))}
                for fi46 := 0; fi46 < 9; fi46++ {
                    v3 = v3
                    if len(v45) >= 2 {
                        v45 = v45[1:len(v45)]
                    }
                    v29 = (v29 + v29)
                    v47 := GenS6{F0: (v29 >> ((42) & 63)), F1: 2}
                    if (v4 >= 1.5) {
                        var v48 int8 = (v29 | 100)
                        switch ("\n\t" + ("\n\t" + "a")) {
                        case "":
                            fmt.Println(v3)
                        }
                        if true {
                            v49 := GenS2{F0: (("foo bar" + "a") + "\n\t")}
                            v6 -= v6
                            _ = v49
                        }
                        _ = v48
                    } else {
                        fmt.Println(v4, v30, v1)
                        if len(v45) >= 2 {
                            v45 = v45[1:len(v45)]
                        }
                    }
                    var v50 uint8 = 215
                    _ = fi46
                    _ = v47
                    _ = v50
                }
                var v51 uint16 = 1
                _ = fi44
                _ = v45
                _ = v51
            }
            fmt.Println(v4, v3, v3)
        }
    case 4294967295:
        if false {
            if ((v6 + v6) >= v6) {
                var v52 bool = (false && true)
                _ = v52
            }
            fmt.Println(v30)
            var v53 int64 = 9223372036854775807
            _ = v53
        } else {
            v54 := []int{-1, (v5 >> ((18446744073709551615) & 63)), v5}
            v5 = v5
            fmt.Println(v30, v3, v1)
            _ = v54
        }
    default:
        v55 := []int{(^(v5))}
        _ = v55
    }
    v56 := []int{(v2 ^ v2), (v2 + 4294967295), -637}
    fmt.Println(v1)
    var v57 int64 = -1
    v58 := GenS5{F0: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v6
}
