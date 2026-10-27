package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint8; F1 uint; F2 int32 }
func genFn0(p0 uint32) uint {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 uint16 = 1
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        if ((fi6 ^ -1) != fi6) {
            var v7 int = (-(fi6))
            if (true || (!true)) {
                fmt.Println(fi6)
                v8 := GenS1{F0: "hello"}
                _ = v8
            }
            var v9 float64 = 3.14159
            _ = v7
            _ = v9
        }
        if false {
            v10 := GenS1{F0: "a"}
            if (true || (("hello" + "\n\t") > "hello")) {
                v11 := GenS2{F0: "\n\t", F1: 42, F2: ((-9223372036854775808 > 1) || (!true))}
                v12 := []int{4294967295, 32767}
                v12 = append(v12, (fi6 * -1))
                _ = v11
                _ = v12
            }
            _ = v10
        } else {
            v13 := []int{(^(fi6)), 1, (fi6 / ((128) | 1)), fi6}
            switch fi6 {
            case 1:
                for fi14 := 0; fi14 < 5; fi14++ {
                    v13 = append(v13, (fi14 / (((^(fi14))) | 1)))
                    fmt.Println(p0)
                    _ = fi14
                }
            default:
                if (999 == 32767) {
                    var v15 uint8 = 144
                    p0 = (p0 * p0)
                    _ = len(v13)
                    _ = v15
                }
            }
            _ = v13
        }
        v16 := []int{}
        _ = fi6
        _ = v16
    }
    if (((4294967295 > 4294967295) && (4294967295 == 42)) && false) {
        var v17 uint64 = uint64(p0)
        v18 := GenS1{F0: "\n\t"}
        switch "" {
        case "a":
            v17 -= v17
        case "":
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 int32 = 2147483647
                switch "\n\t" {
                case "\n\t":
                    for fi21 := 0; fi21 < 20; fi21++ {
                        v22 := GenS2{F0: "\n\t", F1: 177, F2: true}
                        _ = fi21
                        _ = v22
                    }
                case "":
                    var v23 bool = ((!((((0 < 88) && false) || (1 == (^(v17)))) && false)) && ((-2147483648 <= v20) || true))
                    _ = v23
                }
                var v24 int8 = -128
                _ = fi19
                _ = v20
                _ = v24
            }
        case "\n\t":
            switch 32768 {
            case -2147483648:
                v17 = v17
            case 4294967295:
                if false {
                    v25 := []int{-1, 1}
                    for fi26 := 0; fi26 < 18; fi26++ {
                        var v27 string = ("foo bar" + (("foo bar" + "hello") + ("a" + "\n\t")))
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                } else {
                    v28 := GenS3{F0: (v17 / ((v17) | 1)), F1: 255}
                    v29 := GenS4{F0: 42}
                    _ = v28
                    _ = v29
                }
            }
        }
        _ = v17
        _ = v18
    } else {
        p0 -= p0
        v30 := GenS5{F0: 0, F1: 952}
        var v31 uint = 713
        _ = v30
        _ = v31
    }
    switch (("" + "hello") + "") {
    case "a":
        var v32 bool = true
        _ = v32
    default:
        v33 := GenS1{F0: ""}
        _ = v33
    }
    v34 := []int{0, 4294967295}
    _ = p0
    _ = v34
    return 42
}
func genFn1(p0 uint) bool {
    fmt.Println(p0)
    p0 -= (p0 << ((65535) & 63))
    fmt.Println(p0)
    _ = p0
    return true
}

func main() {
    v1 := []string{"a", "", (("" + "foo bar") + "a")}
    var v2 uint8 = 1
    v3 := make(map[string]bool)
    v35 := GenS1{F0: "hello"}
    for fi36 := 0; fi36 < 18; fi36++ {
        switch (fi36 - fi36) {
        case 2147483647:
            if false {
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            } else {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v2, fi36)
                case "":
                    v2 += 1
                default:
                    if (-111 == -128) {
                        v2 -= v2
                    }
                }
            }
        case 42:
            switch -471 {
            case 2147483647:
                var v37 uint8 = (^(v2))
                _ = v37
            case 4294967295:
                switch 32767 {
                case 128:
                    for fi38 := 0; fi38 < 4; fi38++ {
                        var v39 string = ("a" + (("\n\t" + (("hello" + "hello") + ("hello" + ("hello" + ((("foo bar" + (("\n\t" + "foo bar") + "hello")) + "hello") + "a"))))) + "\n\t"))
                        _ = fi38
                        _ = v39
                    }
                case 127:
                    if true {
                        v2 += 0
                    }
                case 42:
                    var v40 uint = 42
                    _ = v40
                }
            }
        case 32767:
            var v41 uint64 = 382
            _ = v41
        }
        _ = fi36
    }
    v2 = 128
    v2 = v2
    switch 1 {
    case 9223372036854775807:
        if true {
            fmt.Println(v2)
            if (66 > 2) {
                fmt.Println(v2)
            } else {
                if (!(false || true)) {
                    fmt.Println(v2)
                }
                var v42 uint8 = 1
                _ = v42
            }
        } else {
            _ = len(v3)
            if (18446744073709551615 != 0) {
                fmt.Println(v2)
                var v43 int8 = 127
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    var v44 bool = true
                    var v45 float64 = 3.14159
                    _ = v44
                    _ = v45
                }
                _ = v43
            } else {
                if (-2.25 >= 0.1) {
                    var v46 int16 = 0
                    _ = v46
                }
                v47 := v3[("hello" + "")]
                var v48 int8 = -128
                _ = v47
                _ = v48
            }
            v49 := GenS6{F0: 170, F1: 1, F2: 2147483647}
            _ = v49
        }
    default:
        var v50 float64 = 1e10
        _ = v50
    }
    fmt.Println(v2)
    v2 -= (v2 >> ((2) & 63))
    v2 = v2
    switch 110 {
    case 2147483647:
        v2 -= v2
    }
    var v51 int32 = 42
    v51 = v51
    var v52 int32 = v51
    var v53 bool = (true && false)
    var v54 float64 = float64(1.5)
    delete(v3, "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v51
    _ = v52
    _ = v53
    _ = v54
}
