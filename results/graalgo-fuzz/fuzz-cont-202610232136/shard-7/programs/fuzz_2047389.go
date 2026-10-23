package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 int16; F1 int16; F2 bool }
type GenS8 struct { F0 uint8; F1 uint16; F2 uint32 }
func genFn0(p0 uint8, p1 uint) bool {
    p0 -= 0
    var v7 uint32 = 304
    if false {
        v8 := []int{32768, -228}
        _ = v8
    } else {
        v7 = (v7 * (^(v7)))
        for fi9 := 0; fi9 < 3; fi9++ {
            if false {
                var v10 int8 = -128
                var v11 float64 = float64(float32(0.1))
                if (v10 >= v10) {
                    var v12 int16 = -804
                    v13 := GenS1{F0: "\n\t", F1: 9223372036854775807}
                    _ = v12
                    _ = v13
                } else {
                    v14 := GenS2{F0: 127}
                    v15 := GenS2{F0: (fi9 ^ 32768)}
                    _ = v14
                    _ = v15
                }
                _ = v10
                _ = v11
            } else {
                var v16 uint32 = (v7 ^ (v7 / (((v7 + 0)) | 1)))
                _ = v16
            }
            v17 := GenS1{F0: "a", F1: fi9}
            switch fi9 {
            case 240:
                v7 -= 2147483647
            case -832:
                var v18 int32 = 0
                _ = v18
            case 42:
                var v19 bool = (!(v7 != v7))
                _ = v19
            }
            _ = fi9
            _ = v17
        }
    }
    p1 = (p1 >> (((p1 + 0)) & 63))
    _ = p0
    _ = p1
    _ = v7
    return true
}
func genFn1(p0 uint32) string {
    v20 := []int{}
    var v21 int16 = 42
    v22 := GenS3{F0: "\n\t", F1: (-257 == 0)}
    switch -9223372036854775808 {
    case -125:
        for fi23 := 0; fi23 < 8; fi23++ {
            if len(v20) > 0 {
                v20[0] = 4294967295
            }
            if true {
                switch fi23 {
                case -1:
                    v24 := GenS1{F0: ((("foo bar" + ((("foo bar" + "a") + "hello") + "hello")) + (("" + ("hello" + "")) + ("\n\t" + (("foo bar" + "\n\t") + "hello")))) + "hello"), F1: -2147483648}
                    _ = v24
                case 32768:
                    for fi25 := 0; fi25 < 1; fi25++ {
                        fmt.Println(v21)
                        _ = fi25
                    }
                }
                var v26 uint = 113
                var v27 int = (fi23 * -2147483648)
                _ = v26
                _ = v27
            }
            _ = fi23
        }
    default:
        var v28 int64 = 0
        _ = v28
    }
    _ = p0
    _ = v20
    _ = v21
    _ = v22
    return "foo bar"
}

func main() {
    var v1 uint8 = 255
    var v2 int64 = 42
    v3 := []int{127, int(v2), -447, -576}
    v4 := make(map[uint32]string)
    var v5 int32 = 1
    var v6 float32 = 3.14159
    v2 = (-(v2))
    switch "\n\t" {
    case "a":
        v5 -= v5
    case "\n\t":
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 string = (("a" + "a") + ("\n\t" + "hello"))
            switch (((v30 + "\n\t") + v30) + v30) {
            case "\n\t":
                if ("foo bar" == v30) {
                    v5 = (v5 | v5)
                }
            case "":
                v2 = 1
            case "foo bar":
                v31 := GenS4{F0: 18446744073709551615, F1: 42}
                _ = v31
            default:
                var v32 string = v30
                _ = v32
            }
            for fi33 := 0; fi33 < 7; fi33++ {
                _ = len(v4)
                var v34 bool = (-128 != 42)
                _ = fi33
                _ = v34
            }
            _ = fi29
            _ = v30
        }
    case "hello":
        var v35 uint32 = 65535
        _ = v35
    default:
        for fi36 := 0; fi36 < 1; fi36++ {
            switch (fi36 & fi36) {
            case 32767:
                v5 = (v5 * v5)
            }
            _ = fi36
        }
    }
    v6 = v6
    for fi37 := 0; fi37 < 19; fi37++ {
        if (!false) {
            v38 := GenS5{F0: false}
            v5 = 1
            for fi39 := 0; fi39 < 20; fi39++ {
                var v40 float64 = 1e10
                var v41 uint16 = 65535
                if false {
                    v6 = v6
                    var v42 int16 = 0
                    _ = v42
                }
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = v38
        }
        for fi43 := 0; fi43 < 3; fi43++ {
            v1 = 0
            var v44 float32 = 0.0
            v1 = (v1 | v1)
            _ = fi43
            _ = v44
        }
        _ = fi37
    }
    v45 := GenS6{F0: v2, F1: (v1 % ((238) | 1))}
    var v46 uint16 = 9
    v47 := v4[2147483648]
    var v48 int16 = -1
    _ = len(v4)
    v49 := GenS7{F0: v48, F1: (v48 >> ((231) & 63)), F2: (931 == 42)}
    fmt.Println(v5, v47, v48)
    switch 878 {
    case 32767:
        var v50 string = ((v47 + v47) + "\n\t")
        _ = v50
    case -9223372036854775808:
        switch "\n\t" {
        case "foo bar":
            v51 := GenS1{F0: (("\n\t" + "a") + v47), F1: -739}
            _ = v51
        case "hello":
            v5 -= 1
        }
    case 820:
        v52 := GenS8{F0: v1, F1: 1, F2: 65535}
        _ = v52
    }
    var v53 uint32 = 42
    switch (((("" + v47) + "") + (v47 + v47)) + v47) {
    case "a":
        var v54 int8 = 127
        _ = v54
    case "":
        var v55 uint32 = (v53 + 550)
        _ = v55
    case "hello":
        if (!((v53 / ((v53) | 1)) >= v53)) {
            v48 -= v48
            v3 = append(v3, 854)
            for fi56 := 0; fi56 < 12; fi56++ {
                for fi57 := 0; fi57 < 16; fi57++ {
                    var v58 bool = ((-(v2)) < 0)
                    _ = fi57
                    _ = v58
                }
                _ = fi56
            }
        }
    }
    if true {
        switch 9223372036854775807 {
        case 1:
            var v59 bool = (!((v46 / (((v46 * (^(v46)))) | 1)) > v46))
            _ = v59
        case -2147483648:
            var v60 string = "a"
            _ = v60
        case 0:
            v61 := GenS2{F0: 4294967295}
            _ = v61
        default:
            v62 := GenS2{F0: -9223372036854775808}
            _ = v62
        }
        var v63 float64 = 0.1
        v47 = v47
        _ = v63
    } else {
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v53
    _ = v6
}
