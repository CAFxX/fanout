package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 int) uint {
    v4 := []int{4294967295, p0, 42}
    var v5 uint8 = 128
    for fi6 := 0; fi6 < 17; fi6++ {
        v5 -= 255
        _ = fi6
    }
    var v7 float64 = (0.1 - 0.1)
    _ = p0
    _ = v4
    _ = v5
    _ = v7
    return 65535
}
func genFn1(p0 uint64) int32 {
    var v8 uint64 = p0
    v9 := []int{}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    var v10 int8 = 41
    for fi11 := 0; fi11 < 7; fi11++ {
        switch ("\n\t" + "foo bar") {
        case "":
            v12 := GenS1{F0: 628, F1: 667}
            _ = v12
        case "\n\t":
            v13 := GenS1{F0: (^(p0)), F1: 1}
            _ = v13
        case "hello":
            fmt.Println(v8, v10)
        default:
            var v14 int64 = 42
            _ = v14
        }
        v15 := GenS2{F0: 0}
        _ = fi11
        _ = v15
    }
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    v1 := []string{"\n\t", ("\n\t" + ("foo bar" + "a")), "foo bar"}
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    if (!false) {
        var v16 int32 = genFn1(1)
        v17 := GenS2{F0: 213}
        _ = v16
        _ = v17
    }
    var v18 uint8 = 255
    if true {
        _ = len(v1)
        v19 := GenS3{F0: -101}
        _ = v19
    }
    v18 = (v18 & (^(v18)))
    v18 = v18
    switch -623 {
    case 2147483647:
        switch "hello" {
        case "":
            fmt.Println(v18)
        case "hello":
            v18 = (^(v18))
        }
    }
    v18 -= v18
    v20 := v2["\n\t"]
    v21 := GenS3{F0: 1}
    for fi22 := 0; fi22 < 9; fi22++ {
        if true {
            var v23 int16 = 1000
            switch fi22 {
            case 691:
                v24 := GenS4{F0: 0}
                _ = v24
            default:
                for fi25 := 0; fi25 < 1; fi25++ {
                    var v26 uint64 = 852
                    v26 += v26
                    if (!true) {
                        if (100 < 1) {
                            v27 := GenS1{F0: (^(v26)), F1: 654}
                            _ = v27
                        } else {
                            var v28 uint16 = 639
                            v29 := v3[fi25]
                            delete(v3, 128)
                            var v30 uint64 = (^(v26))
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                    }
                    _ = fi25
                    _ = v26
                }
            }
            _ = v23
        } else {
            switch fi22 {
            case 32768:
                fmt.Println(fi22)
            case 128:
                fmt.Println(v18, fi22, v18)
            }
        }
        var v31 float32 = 1e10
        _ = fi22
        _ = v31
    }
    var v32 uint32 = (^(v20))
    for fi33 := 0; fi33 < 17; fi33++ {
        for fi34 := 0; fi34 < 2; fi34++ {
            if true {
                var v35 uint = 4294967295
                v35 = v35
                _ = v35
            } else {
                var v36 float32 = 1.5
                switch (fi33 * (fi34 << ((312) & 63))) {
                case -597:
                    v32 = v32
                case 127:
                    var v37 uint32 = (v20 >> ((2) & 63))
                    _ = v37
                case -9223372036854775808:
                    var v38 int32 = genFn1(2)
                    _ = v38
                }
                _ = v36
            }
            _ = len(v1)
            v18 -= 255
            _ = fi34
        }
        for fi39 := 0; fi39 < 15; fi39++ {
            var v40 bool = (!(v18 <= 117))
            switch fi33 {
            case -896:
                if (!(v20 != v32)) {
                    var v41 uint32 = (v32 - (v20 - 2147483647))
                    switch fi33 {
                    case -117:
                        v1 = append(v1, ((("" + (("foo bar" + "a") + ("a" + "foo bar"))) + "foo bar") + "hello"))
                    case 127:
                        v42 := GenS5{F0: 4294967295}
                        var v43 float32 = 3.14159
                        fmt.Println(v32)
                        _ = v42
                        _ = v43
                    case 32767:
                        v44 := GenS6{F0: (v32 - (^(v41)))}
                        _ = v44
                    }
                    _ = v41
                } else {
                    v32 = (^(v32))
                    var v45 float32 = (3.14159 / 3.14159)
                    _ = v45
                }
            case 32767:
                switch "hello" {
                case "a":
                    var v46 int = fi33
                    _ = v46
                case "hello":
                    v47 := GenS5{F0: 268}
                    _ = v47
                case "":
                    v40 = false
                }
            case -9223372036854775808:
                var v48 uint8 = (v18 % (((v18 | uint8(v32))) | 1))
                _ = v48
            default:
                switch ("a" + (("" + (((("foo bar" + "foo bar") + ("hello" + "foo bar")) + (("foo bar" + "") + "foo bar")) + ("\n\t" + "hello"))) + "hello")) {
                case "foo bar":
                    switch "a" {
                    case "a":
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    case "hello":
                        v32 += 42
                    case "foo bar":
                        v18 = 0
                    }
                }
            }
            if v40 {
                var v49 int64 = -348
                _ = v49
            } else {
                switch (^(fi33)) {
                case -391:
                    var v50 uint16 = 42
                    _ = v50
                }
                var v51 int16 = 42
                fmt.Println(fi39)
                _ = v51
            }
            _ = fi39
            _ = v40
        }
        _ = fi33
    }
    if true {
        var v52 int16 = -895
        _ = v52
    }
    v53 := v3[-264]
    for fi54 := 0; fi54 < 12; fi54++ {
        if len(v1) > 0 {
            v1[0] = ("foo bar" + "")
        }
        switch "" {
        case "\n\t":
            var v55 bool = true
            _ = v55
        case "foo bar":
            v56 := GenS1{F0: 42, F1: 65535}
            _ = v56
        default:
            if false {
                v20 -= 42
                v57 := GenS3{F0: -128}
                _ = v57
            } else {
                v53 = (^(v53))
            }
        }
        v58 := GenS4{F0: -1}
        _ = fi54
        _ = v58
    }
    var v59 int8 = 86
    var v60 float32 = (1e10 + (-2.25 / 1.5))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v53:", v53)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v53
    _ = v59
    _ = v60
}
