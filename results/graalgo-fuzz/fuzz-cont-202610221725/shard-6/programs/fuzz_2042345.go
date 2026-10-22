package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int; F1 string; F2 uint32 }
type GenS5 struct { F0 int8 }
func genFn0() bool {
    var v6 uint64 = 1
    switch 1 {
    case 2147483647:
        for fi7 := 0; fi7 < 7; fi7++ {
            if (!false) {
                v6 = (v6 & 795)
            }
            v8 := GenS1{F0: 65535, F1: 1}
            v6 = 1
            _ = fi7
            _ = v8
        }
    case -2147483648:
        v9 := []int{-2147483648, 480}
        _ = v9
    }
    _ = v6
    return (!false)
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 2147483647
    var v3 uint = 628
    var v4 int = 32768
    var v5 int = v4
    var v10 int8 = -1
    v11 := GenS1{F0: 255, F1: 0}
    v5 = -9223372036854775808
    switch "\n\t" {
    case "foo bar":
        v12 := GenS2{F0: 1}
        _ = v12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        if (!(v2 < (v2 & (^(v2))))) {
            v14 := GenS1{F0: 1, F1: 405}
            if false {
                var v15 float32 = 1.5
                for fi16 := 0; fi16 < 4; fi16++ {
                    v4 += v5
                    v2 += (v2 | (^(v2)))
                    var v17 uint64 = 42
                    _ = fi16
                    _ = v17
                }
                var v18 string = "hello"
                _ = v15
                _ = v18
            } else {
                switch (v5 % ((v5) | 1)) {
                case -985:
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v10 -= v10
                        v20 := GenS3{F0: ("a" + "\n\t"), F1: "foo bar"}
                        delete(v1, "")
                        _ = fi19
                        _ = v20
                    }
                case -674:
                    v2 = v2
                }
            }
            _ = len(v1)
            _ = v14
        }
        _ = fi13
    }
    switch 42 {
    case 477:
        switch ("foo bar" + "a") {
        case "foo bar":
            var v21 string = ("" + "foo bar")
            _ = v21
        default:
            v3 = (v3 << ((v3) & 63))
        }
    case 32767:
        v2 += (v2 + 357)
    }
    v4 += (v5 ^ v4)
    var v22 uint32 = (v2 >> ((255) & 63))
    var v23 bool = (!(!(true || true)))
    v5 += (v5 ^ v5)
    var v24 uint = uint(v22)
    switch "\n\t" {
    case "foo bar":
        if v23 {
            v25 := GenS3{F0: ("a" + "foo bar"), F1: ""}
            delete(v1, "\n\t")
            _ = v25
        } else {
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 uint32 = 2
                var v28 int64 = -334
                var v29 uint8 = 128
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            v30 := GenS2{F0: 1}
            v31 := []int{(v5 + 9223372036854775807), 831, v4}
            _ = v30
            _ = v31
        }
    case "hello":
        if v23 {
            switch (v5 | v4) {
            case 42:
                if v23 {
                    v32 := []int{-1, int(v2)}
                    _ = v32
                } else {
                    if (1 == v22) {
                        v33 := GenS2{F0: 42}
                        _ = len(v1)
                        fmt.Println(v23, v5)
                        _ = v33
                    } else {
                        var v34 uint16 = 32767
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 16; fi35++ {
                        v22 = (v22 >> (((v24 << (((^(v24))) & 63))) & 63))
                        v36 := GenS2{F0: 42}
                        _ = fi35
                        _ = v36
                    }
                }
            case -496:
                if (!v23) {
                    v24 = (v3 + v24)
                    if (!v23) {
                        var v37 float64 = 1e10
                        _ = v37
                    } else {
                        var v38 uint64 = 2
                        var v39 int64 = 9223372036854775807
                        _ = v38
                        _ = v39
                    }
                    v40 := []int{4294967295, (v5 >> ((v24) & 63)), 1, -126}
                    _ = v40
                } else {
                    v3 = (v3 >> (((v3 * (v3 ^ (v24 % ((v24) | 1))))) & 63))
                    var v41 float64 = -2.25
                    var v42 bool = (0 >= 160)
                    _ = v41
                    _ = v42
                }
            case 127:
                v22 += (v2 % ((42) | 1))
            }
            var v43 int32 = -847
            v44 := GenS4{F0: (v5 & (v5 * v4)), F1: (("a" + "\n\t") + "a"), F2: (v22 % (((v2 * (v2 * (v22 ^ v22)))) | 1))}
            _ = v43
            _ = v44
        }
    case "":
        fmt.Println(v22)
    default:
        switch v4 {
        case -2147483648:
            var v45 uint32 = (^(v22))
            _ = v45
        case 128:
            var v46 uint8 = 1
            _ = v46
        default:
            if false {
                v22 -= (v22 / ((v22) | 1))
                fmt.Println(v5)
                v2 -= 2147483648
            } else {
                var v47 uint8 = 2
                _ = v47
            }
        }
    }
    var v48 int16 = 1000
    v49 := GenS5{F0: v10}
    v48 = -85
    if false {
        v2 = 42
    } else {
        fmt.Println(v22, v2, v23)
    }
    v50 := []int{v5, (v5 + v4), -962, 688}
    if v23 {
        v10 = v10
    } else {
        v51 := GenS4{F0: v5, F1: "hello", F2: v22}
        fmt.Println(v2, v5)
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
