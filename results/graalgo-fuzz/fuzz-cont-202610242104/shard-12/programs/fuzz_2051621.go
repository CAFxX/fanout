package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int64; F1 string; F2 bool }
type GenS3 struct { F0 string; F1 string; F2 bool }
type GenS4 struct { F0 int32; F1 int; F2 uint64 }
type GenS5 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint64, p1 int32) int64 {
    p0 = p0
    v5 := GenS1{F0: (p0 >> ((42) & 63)), F1: -1}
    v6 := []int{154, 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1() uint32 {
    var v7 uint = 731
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint32 = 361
        _ = fi8
        _ = v9
    }
    if true {
        v10 := GenS1{F0: 42, F1: 42}
        v11 := []int{85, 860, -2147483648}
        _ = v10
        _ = v11
    } else {
        if ((v7 <= 255) || true) {
            var v12 int64 = -326
            if true {
                for fi13 := 0; fi13 < 20; fi13++ {
                    v14 := []int{}
                    v12 = v12
                    var v15 float32 = (1e10 + 1.5)
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                if (!(-1 == -240)) {
                    if (true && false) {
                        for fi16 := 0; fi16 < 14; fi16++ {
                            v17 := []int{(fi16 | fi16)}
                            v7 = (v7 << ((4294967295) & 63))
                            switch (^(fi16)) {
                            case -667:
                                if true {
                                    var v18 uint8 = 100
                                    _ = v18
                                }
                            }
                            _ = fi16
                            _ = v17
                        }
                    } else {
                        v19 := GenS2{F0: (v12 & v12), F1: "foo bar", F2: true}
                        var v20 bool = true
                        _ = v19
                        _ = v20
                    }
                    var v21 uint16 = 853
                    v7 = (^(v7))
                    _ = v21
                } else {
                    v7 -= v7
                    var v22 float32 = 0.0
                    _ = v22
                }
            }
            _ = v12
        } else {
            switch 4294967295 {
            case 914:
                switch 9223372036854775807 {
                case 0:
                    v7 += v7
                }
            case 32767:
                v23 := []int{-844}
                _ = v23
            }
        }
        fmt.Println(v7)
        var v24 uint = 304
        _ = v24
    }
    if true {
        var v25 int8 = -128
        for fi26 := 0; fi26 < 1; fi26++ {
            if true {
                v7 += v7
                var v27 int16 = 66
                _ = v27
            } else {
                for fi28 := 0; fi28 < 6; fi28++ {
                    v29 := GenS3{F0: "foo bar", F1: ("" + "hello"), F2: false}
                    _ = fi28
                    _ = v29
                }
            }
            _ = fi26
        }
        v30 := GenS4{F0: 2147483647, F1: -835, F2: 42}
        _ = v25
        _ = v30
    }
    if ((1.5 >= 1e10) || false) {
        var v31 bool = (42 == 100)
        fmt.Println(v31, v7)
        var v32 int8 = 127
        _ = v31
        _ = v32
    } else {
        fmt.Println(v7)
        v7 += (v7 - v7)
    }
    _ = v7
    return 4294967295
}

func main() {
    var v1 float32 = 0.1
    var v2 bool = false
    var v3 int32 = 1
    v4 := make(map[int]bool)
    if v2 {
        var v33 uint16 = 42
        for fi34 := 0; fi34 < 12; fi34++ {
            var v35 float32 = v1
            var v36 float32 = (v1 - v1)
            switch "\n\t" {
            case "a":
                var v37 uint64 = 892
                _ = v37
            }
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = v33
    } else {
        switch ("\n\t" + "a") {
        case "hello":
            for fi38 := 0; fi38 < 7; fi38++ {
                v1 -= 1.5
                v39 := GenS2{F0: 884, F1: "a", F2: v2}
                for fi40 := 0; fi40 < 1; fi40++ {
                    if true {
                        var v41 int64 = -1
                        _ = v41
                    } else {
                        if v2 {
                            fmt.Println(v3)
                            v42 := v4[fi40]
                            _ = v42
                        }
                        var v43 float32 = 0.1
                        _ = v43
                    }
                    v1 += 0.0
                    v3 = (v3 ^ (v3 * (v3 & 1)))
                    _ = fi40
                }
                _ = fi38
                _ = v39
            }
        default:
            switch 42 {
            case 2147483647:
                v1 -= v1
            default:
                var v44 int8 = 100
                _ = v44
            }
        }
    }
    for fi45 := 0; fi45 < 2; fi45++ {
        for fi46 := 0; fi46 < 12; fi46++ {
            if v2 {
                switch "hello" {
                case "":
                    switch "hello" {
                    case "":
                        v3 = v3
                    case "a":
                        v3 = (v3 ^ v3)
                    }
                }
                v1 += (v1 + (1e10 + ((1.5 - -2.25) * 0.0)))
                if v2 {
                    v3 = (v3 * (v3 >> ((255) & 63)))
                    var v47 float32 = v1
                    _ = v47
                }
            } else {
                var v48 bool = v2
                _ = v48
            }
            _ = fi46
        }
        v3 = 0
        _ = fi45
    }
    fmt.Println(v3)
    if v2 {
        v49 := GenS2{F0: 777, F1: "a", F2: true}
        _ = v49
    }
    var v50 float32 = (1e10 / v1)
    v2 = v2
    var v51 uint16 = 2
    for fi52 := 0; fi52 < 1; fi52++ {
        switch "hello" {
        case "":
            v53 := GenS1{F0: 2, F1: -1}
            _ = v53
        default:
            fmt.Println(v1, v1)
        }
        v54 := []int{}
        var v55 bool = v2
        _ = fi52
        _ = v54
        _ = v55
    }
    fmt.Println(v3)
    var v56 uint8 = 198
    switch -951 {
    case 2147483647:
        var v57 float64 = (float64(-2.25) + float64((-2.25 - 1.5)))
        _ = v57
    default:
        v1 = 3.14159
    }
    v1 -= v50
    delete(v4, -9223372036854775808)
    if (!(!(false && false))) {
        v56 += v56
    } else {
        v58 := v4[128]
        for fi59 := 0; fi59 < 5; fi59++ {
            v51 += 42
            v3 += v3
            if (-840 <= 996) {
                v60 := GenS4{F0: v3, F1: 9223372036854775807, F2: 979}
                _ = v60
            }
            _ = fi59
        }
        _ = v58
    }
    v61 := GenS5{F0: 2, F1: genFn1()}
    v1 = v50
    for fi62 := 0; fi62 < 6; fi62++ {
        var v63 float32 = (((0.1 + -2.25) * v50) / (-2.25 + ((float32(3.14159) / (v1 / ((-2.25 - v1) - 0.0))) + (((0.0 + v1) / v1) + v50))))
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v56
    _ = v61
}
