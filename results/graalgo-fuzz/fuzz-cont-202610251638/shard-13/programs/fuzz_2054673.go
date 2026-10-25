package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
type GenS2 struct { F0 int; F1 uint8; F2 int64 }
type GenS3 struct { F0 string; F1 uint32 }
func genFn0() string {
    for fi4 := 0; fi4 < 5; fi4++ {
        for fi5 := 0; fi5 < 13; fi5++ {
            if (2 != 42) {
                switch ((("hello" + "\n\t") + ("" + "")) + "") {
                case "a":
                    var v6 float64 = 0.1
                    _ = v6
                case "foo bar":
                    var v7 uint8 = 42
                    _ = v7
                default:
                    v8 := []int{fi5}
                    _ = v8
                }
                v9 := []int{fi4, 377}
                if (1 != 100) {
                    fmt.Println(fi5, fi4)
                    switch "a" {
                    case "":
                        v10 := GenS1{F0: 138, F1: ((("a" + "") + "\n\t") + "hello"), F2: 282}
                        v11 := GenS1{F0: 1, F1: ("foo bar" + "a"), F2: 42}
                        _ = v10
                        _ = v11
                    case "hello":
                        if len(v9) > 0 {
                            v9[0] = fi5
                        }
                    case "a":
                        v12 := GenS1{F0: 128, F1: (("\n\t" + (("foo bar" + "hello") + "a")) + "hello"), F2: 255}
                        _ = v12
                    }
                    var v13 uint32 = 0
                    _ = len(v9)
                    _ = v13
                }
                _ = v9
            }
            if (908 <= 635) {
                switch 1 {
                case -407:
                    var v14 int32 = 1
                    _ = v14
                case 1000:
                    var v15 int16 = 0
                    _ = v15
                }
                if ((499 <= 2147483648) && true) {
                    fmt.Println(fi5, fi4)
                }
            }
            v16 := GenS1{F0: 1, F1: (("\n\t" + "a") + ("hello" + "foo bar")), F2: 930}
            _ = fi5
            _ = v16
        }
        _ = fi4
    }
    v17 := GenS2{F0: 4294967295, F1: 0, F2: -1}
    for fi18 := 0; fi18 < 11; fi18++ {
        if true {
            for fi19 := 0; fi19 < 14; fi19++ {
                fmt.Println(fi19, fi19)
                _ = fi19
            }
        }
        if true {
            for fi20 := 0; fi20 < 1; fi20++ {
                fmt.Println(fi20, fi20)
                var v21 uint32 = 759
                if true {
                    fmt.Println(fi20)
                }
                _ = fi20
                _ = v21
            }
            v22 := GenS2{F0: 128, F1: 95, F2: -226}
            for fi23 := 0; fi23 < 15; fi23++ {
                v24 := []int{fi18, 9223372036854775807, (fi18 + fi18), (fi18 | (fi23 ^ fi23))}
                _ = fi23
                _ = v24
            }
            _ = v22
        } else {
            if ((1e10 / 1.0) == -2.25) {
                for fi25 := 0; fi25 < 17; fi25++ {
                    var v26 int32 = 1
                    _ = fi25
                    _ = v26
                }
                fmt.Println(fi18)
            } else {
                var v27 int16 = 1
                v27 = (v27 >> ((0) & 63))
                _ = v27
            }
            var v28 int64 = -1
            _ = v28
        }
        for fi29 := 0; fi29 < 19; fi29++ {
            var v30 float64 = 3.14159
            v31 := GenS2{F0: fi29, F1: 1, F2: 0}
            switch "a" {
            case "hello":
                fmt.Println(fi18)
            case "\n\t":
                if false {
                    var v32 uint = 673
                    v33 := []int{fi18, fi18}
                    _ = v32
                    _ = v33
                } else {
                    v30 = v30
                    var v34 uint16 = 0
                    var v35 uint32 = 2147483647
                    _ = v34
                    _ = v35
                }
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi18
    }
    for fi36 := 0; fi36 < 17; fi36++ {
        v37 := []int{fi36, (fi36 & (fi36 + 32767))}
        var v38 uint = 2
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 uint = 228
    _ = v17
    _ = v39
    return "a"
}

func main() {
    var v1 uint32 = 42
    v2 := []int{}
    v3 := []string{"foo bar"}
    fmt.Println(v1)
    var v40 uint = 248
    fmt.Println(v40, v1)
    var v41 uint8 = uint8(v40)
    for fi42 := 0; fi42 < 13; fi42++ {
        if false {
            v43 := GenS3{F0: ("hello" + ("hello" + "hello")), F1: v1}
            v1 = 0
            _ = v43
        }
        _ = fi42
    }
    var v44 bool = (false && false)
    for fi45 := 0; fi45 < 2; fi45++ {
        v44 = true
        fmt.Println(v44, v44, v44)
        if v44 {
            fmt.Println(v40)
            var v46 int64 = -928
            _ = v46
        }
        _ = fi45
    }
    v1 -= v1
    var v47 int = 4294967295
    if len(v3) > 0 {
        v3[0] = "a"
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v44
    _ = v47
}
