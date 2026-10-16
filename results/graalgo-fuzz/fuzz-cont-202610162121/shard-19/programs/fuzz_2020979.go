package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int }
type GenS3 struct { F0 int16; F1 uint; F2 string }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 uint; F1 int16; F2 string }
func genFn0(p0 uint64) int {
    v4 := GenS1{F0: 1}
    if true {
        v5 := GenS2{F0: (p0 & 2), F1: 4294967295, F2: 32768}
        var v6 int32 = 42
        _ = v5
        _ = v6
    } else {
        var v7 uint = 1
        p0 = (p0 * p0)
        var v8 int = 42
        _ = v7
        _ = v8
    }
    if true {
        v9 := GenS2{F0: p0, F1: 959, F2: 42}
        _ = v9
    } else {
        p0 -= p0
        if false {
            v10 := GenS1{F0: 2147483647}
            v11 := GenS2{F0: 114, F1: 2147483647, F2: 42}
            v12 := []int{128, -548, 4294967295}
            _ = v10
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 2; fi13++ {
                for fi14 := 0; fi14 < 14; fi14++ {
                    fmt.Println(fi13, fi13)
                    switch "\n\t" {
                    case "\n\t":
                        v15 := []int{fi14, fi13, (fi13 + fi13), 210}
                        _ = v15
                    case "a":
                        var v16 int32 = 0
                        _ = v16
                    }
                    var v17 float32 = float32(3.14159)
                    _ = fi14
                    _ = v17
                }
                _ = fi13
            }
            v18 := GenS3{F0: 32767, F1: 18446744073709551615, F2: "foo bar"}
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 int32 = 42
                if (true || (true || false)) {
                    v21 := GenS4{F0: 154, F1: 1}
                    _ = v21
                } else {
                    switch 128 {
                    case 127:
                        v22 := []int{fi19, 128, (fi19 * (fi19 >> ((42) & 63)))}
                        _ = v22
                    default:
                        if false {
                            v23 := GenS5{F0: 255, F1: true}
                            v24 := []int{fi19, fi19}
                            _ = v23
                            _ = v24
                        } else {
                            p0 = p0
                            v25 := []int{}
                            if (42 == 1) {
                                p0 = p0
                            }
                            _ = v25
                        }
                    }
                    switch (-(fi19)) {
                    case 127:
                        var v26 int16 = -32768
                        _ = v26
                    default:
                        switch ("\n\t" + "\n\t") {
                        case "a":
                            if (90 == 2) {
                                fmt.Println(v20)
                                p0 = p0
                            } else {
                                v27 := []int{-2147483648}
                                _ = v27
                            }
                        case "hello":
                            p0 = 4294967295
                        case "\n\t":
                            p0 = p0
                        default:
                            var v28 uint = 682
                            _ = v28
                        }
                    }
                    var v29 uint64 = p0
                    _ = v29
                }
                var v30 uint16 = 874
                _ = fi19
                _ = v20
                _ = v30
            }
            _ = v18
        }
        var v31 int8 = 0
        _ = v31
    }
    p0 += 704
    var v32 int64 = 9223372036854775807
    _ = p0
    _ = v32
    _ = v4
    return 1
}
func genFn1() bool {
    v33 := GenS2{F0: 2, F1: 2, F2: 191}
    var v34 uint = 65535
    fmt.Println(v34)
    var v35 string = "hello"
    for fi36 := 0; fi36 < 10; fi36++ {
        if ("a" >= v35) {
            v37 := []int{(^(fi36)), fi36, (fi36 - (fi36 << ((v34) & 63)))}
            var v38 int64 = 359
            var v39 uint16 = 1
            _ = v37
            _ = v38
            _ = v39
        }
        v35 = (v35 + v35)
        _ = fi36
    }
    _ = v33
    _ = v34
    _ = v35
    return (true || (true || (true && ((0 >= 4294967295) || true))))
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    var v3 int64 = 1
    v3 = v3
    switch 4294967295 {
    case 32768:
        fmt.Println(v3)
    default:
        if (526 == 718) {
            var v40 float64 = (3.14159 + 0.1)
            v41 := GenS4{F0: 42, F1: 1}
            _ = v40
            _ = v41
        }
    }
    v42 := GenS4{F0: 690, F1: 42}
    for fi43 := 0; fi43 < 17; fi43++ {
        var v44 uint16 = 65535
        delete(v2, fi43)
        _ = fi43
        _ = v44
    }
    for fi45 := 0; fi45 < 5; fi45++ {
        v46 := GenS3{F0: 340, F1: 357, F2: "\n\t"}
        v3 = -9223372036854775808
        v3 = (v3 / ((0) | 1))
        _ = fi45
        _ = v46
    }
    for fi47 := 0; fi47 < 10; fi47++ {
        var v48 int8 = 42
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            switch "foo bar" {
            case "\n\t":
                v48 = (v48 + v48)
            case "":
                v49 := GenS3{F0: -1, F1: 772, F2: (("hello" + ((("hello" + "hello") + ("" + "")) + ("" + "a"))) + "\n\t")}
                _ = v49
            case "foo bar":
                for fi50 := 0; fi50 < 7; fi50++ {
                    v51 := GenS1{F0: 589}
                    switch ("hello" + "\n\t") {
                    case "a":
                        var v52 int = fi50
                        if (true || false) {
                            fmt.Println(fi47)
                        }
                        if true {
                            v53 := v1["hello"]
                            _ = v53
                        }
                        _ = v52
                    case "\n\t":
                        fmt.Println(v48)
                        v48 -= v48
                    }
                    _ = fi50
                    _ = v51
                }
            default:
                v3 -= (v3 % ((-709) | 1))
            }
        }
        v54 := GenS6{F0: 2, F1: 1, F2: ((("\n\t" + "a") + "\n\t") + "a")}
        _ = fi47
        _ = v48
        _ = v54
    }
    fmt.Println(v3)
    if true {
        var v55 float64 = ((-2.25 / float64(1.5)) - 0.0)
        switch "a" {
        case "":
            switch 186 {
            case 32767:
                _ = len(v2)
            default:
                if (!(true || (0 == 2))) {
                    v56 := []int{1, -818, 128}
                    for fi57 := 0; fi57 < 19; fi57++ {
                        for fi58 := 0; fi58 < 18; fi58++ {
                            v3 = v3
                            _ = fi58
                        }
                        delete(v1, "hello")
                        _ = fi57
                    }
                    if (-2147483648 < 42) {
                        v59 := v1["\n\t"]
                        switch (("\n\t" + "foo bar") + "") {
                        case "\n\t":
                            if (true && false) {
                                var v60 int16 = 0
                                _ = v60
                            }
                        case "a":
                            var v61 float32 = 0.0
                            _ = v61
                        }
                        _ = v59
                    }
                    _ = v56
                }
            }
        }
        _ = v55
    }
    var v62 uint64 = 0
    switch (("" + (("a" + "hello") + "")) + ((("hello" + "") + "") + (("foo bar" + ("\n\t" + "")) + ("" + "hello")))) {
    case "foo bar":
        for fi63 := 0; fi63 < 3; fi63++ {
            var v64 uint64 = v62
            fmt.Println(fi63, v64)
            fmt.Println(v3, fi63)
            _ = fi63
            _ = v64
        }
    case "":
        fmt.Println(v62, v3)
    case "hello":
        v62 = (v62 ^ (v62 >> ((611) & 63)))
    }
    if (!true) {
        v3 -= 940
    } else {
        for fi65 := 0; fi65 < 2; fi65++ {
            var v66 int64 = (v3 >> ((18446744073709551615) & 63))
            fmt.Println(fi65, v3)
            _ = fi65
            _ = v66
        }
        var v67 uint16 = 42
        _ = v67
    }
    var v68 uint16 = 903
    var v69 string = "hello"
    v62 -= (^(v62))
    var v70 string = v69
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v62
    _ = v68
    _ = v69
    _ = v70
}
