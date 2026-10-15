package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 string; F1 bool; F2 uint32 }
func genFn0(p0 uint64, p1 int32) int64 {
    var v7 uint8 = 0
    v8 := GenS1{F0: 1, F1: -2147483648}
    var v9 uint64 = (p0 * (p0 << ((0) & 63)))
    v10 := []int{}
    if true {
        fmt.Println(p1, v7, v7)
        v11 := GenS1{F0: 65535, F1: p1}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []string{"a", "a", "foo bar", "hello"}
    var v2 uint32 = 2
    var v3 uint64 = 42
    v4 := make(map[int]uint32)
    var v5 uint64 = v3
    v6 := make(map[int]bool)
    v12 := GenS1{F0: 989, F1: 0}
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v3, fi13)
        for fi14 := 0; fi14 < 3; fi14++ {
            switch ("a" + "hello") {
            case "":
                v15 := v6[(fi13 / (((fi13 << ((4294967295) & 63))) | 1))]
                _ = v15
            }
            if false {
                v2 = v2
                var v16 int64 = genFn0((v3 * uint64(v2)), 2147483647)
                var v17 float32 = 1.5
                _ = v16
                _ = v17
            }
            var v18 string = "hello"
            _ = fi14
            _ = v18
        }
        _ = len(v4)
        _ = fi13
    }
    var v19 int = 42
    v19 = 33
    var v20 string = "foo bar"
    var v21 int = v19
    v19 += v19
    switch (v20 + v20) {
    case "\n\t":
        var v22 int64 = genFn0(v5, 1)
        _ = v22
    case "hello":
        var v23 int16 = 979
        _ = v23
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v24 string = ((v20 + v20) + "a")
        switch (^(v21)) {
        case 740:
            v25 := GenS2{F0: 1, F1: (v5 & v5)}
            _ = v25
        case 818:
            if ((false && (true && false)) && true) {
                switch (v21 / ((-611) | 1)) {
                case 0:
                    v20 = v20
                case 32767:
                    if false {
                        v26 := v6[32767]
                        for fi27 := 0; fi27 < 18; fi27++ {
                            var v28 int = (v21 - 9223372036854775807)
                            v3 += (v3 % (((v3 ^ (v3 / ((150) | 1)))) | 1))
                            if (!(0 >= 2147483647)) {
                                v5 -= 782
                                if (!v26) {
                                    switch v20 {
                                    case "a":
                                        var v29 uint = 0
                                        _ = v29
                                    case "\n\t":
                                        var v30 int16 = -1
                                        _ = v30
                                    }
                                    if v26 {
                                        for fi31 := 0; fi31 < 10; fi31++ {
                                            v19 += fi27
                                            _ = fi31
                                        }
                                    } else {
                                        fmt.Println(v28)
                                    }
                                }
                                if v26 {
                                    v3 -= v3
                                } else {
                                    var v32 int32 = -2147483648
                                    _ = v32
                                }
                            }
                            _ = fi27
                            _ = v28
                        }
                        _ = v26
                    } else {
                        v33 := GenS1{F0: 255, F1: 42}
                        _ = v33
                    }
                case -1:
                    fmt.Println(v24, v20)
                default:
                    var v34 int = -510
                    _ = v34
                }
            }
        case 2147483647:
            for fi35 := 0; fi35 < 7; fi35++ {
                v24 = v24
                v20 = "\n\t"
                for fi36 := 0; fi36 < 6; fi36++ {
                    if (true || false) {
                        if false {
                            fmt.Println(v2, v21)
                        }
                    }
                    v37 := GenS3{F0: v20, F1: true, F2: (v2 * (v2 * (^(v2))))}
                    _ = fi36
                    _ = v37
                }
                _ = fi35
            }
        default:
            if (1.5 > (1e10 / 0.1)) {
                fmt.Println(v20, v19, v20)
            } else {
                var v38 bool = (!false)
                var v39 float64 = 0.0
                var v40 uint8 = 127
                _ = v38
                _ = v39
                _ = v40
            }
        }
        _ = v24
    }
    if ((((!(-1 != 60)) && true) || true) || (false || (!true))) {
        if false {
            v2 += v2
            delete(v4, 32768)
        }
        switch ("foo bar" + "") {
        case "a":
            fmt.Println(v5, v3)
        case "":
            var v41 uint64 = 18446744073709551615
            _ = v41
        case "foo bar":
            var v42 int16 = 1000
            _ = v42
        }
    } else {
        if (v2 <= 4294967295) {
            v19 = (-(v19))
            fmt.Println(v5, v19)
        } else {
            v43 := GenS1{F0: 2, F1: -1}
            var v44 string = v20
            v19 -= 754
            _ = v43
            _ = v44
        }
        v45 := GenS3{F0: "foo bar", F1: true, F2: (v2 + v2)}
        _ = v45
    }
    for fi46 := 0; fi46 < 15; fi46++ {
        switch v20 {
        case "foo bar":
            var v47 float32 = 0.1
            _ = v47
        case "\n\t":
            _ = len(v4)
        case "a":
            var v48 uint = 151
            _ = v48
        default:
            switch (fi46 - (fi46 + 2147483647)) {
            case -721:
                var v49 int64 = 300
                _ = v49
            case 42:
                v3 = v5
            }
        }
        _ = fi46
    }
    for fi50 := 0; fi50 < 2; fi50++ {
        var v51 int = v19
        v19 -= (v19 / (((fi50 | (^(v21)))) | 1))
        var v52 string = v20
        _ = fi50
        _ = v51
        _ = v52
    }
    var v53 int32 = 159
    var v54 string = v20
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
