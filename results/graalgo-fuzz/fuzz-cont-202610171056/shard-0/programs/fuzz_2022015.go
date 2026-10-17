package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int; F1 int32; F2 int8 }
func genFn0() uint8 {
    if (438 >= 2) {
        var v4 uint32 = 1
        for fi5 := 0; fi5 < 2; fi5++ {
            v4 += 2147483647
            _ = fi5
        }
        _ = v4
    } else {
        switch ("a" + "hello") {
        case "\n\t":
            if (65535 < 1) {
                for fi6 := 0; fi6 < 18; fi6++ {
                    var v7 uint = 2
                    for fi8 := 0; fi8 < 1; fi8++ {
                        fmt.Println(fi6, fi6)
                        v7 = (v7 / ((v7) | 1))
                        v9 := []int{1, fi6, (fi8 - (fi8 >> ((236) & 63)))}
                        var v10 float64 = float64(0.0)
                        fmt.Println(v7, v10)
                        if (255 == 42) {
                            v7 = v7
                            v7 = (v7 ^ v7)
                            switch (fi8 ^ (fi8 ^ fi6)) {
                            case 4294967295:
                                fmt.Println(v7)
                            case 32768:
                                v11 := GenS1{F0: 91}
                                _ = v11
                            case 42:
                                var v12 bool = (42 >= 2)
                                _ = v12
                            default:
                                var v13 float32 = (((-2.25 - 0.0) + (((1e10 * 1.5) + 0.0) - 3.14159)) - 0.0)
                                _ = v13
                            }
                            fmt.Println(v10, v7, fi8)
                        } else {
                            v14 := GenS1{F0: -9223372036854775808}
                            _ = v14
                        }
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    _ = fi6
                    _ = v7
                }
            }
        case "hello":
            v15 := []int{-690}
            _ = v15
        default:
            var v16 int64 = 1
            _ = v16
        }
    }
    v17 := GenS2{F0: -992}
    _ = v17
    return 1
}
func genFn1(p0 bool) int8 {
    if p0 {
        v18 := GenS1{F0: -423}
        _ = v18
    } else {
        if true {
            for fi19 := 0; fi19 < 6; fi19++ {
                fmt.Println(fi19, p0)
                for fi20 := 0; fi20 < 7; fi20++ {
                    fmt.Println(fi20, fi20, fi19)
                    var v21 uint32 = 2
                    _ = fi20
                    _ = v21
                }
                _ = fi19
            }
        }
        var v22 string = ("" + "\n\t")
        _ = v22
    }
    if p0 {
        for fi23 := 0; fi23 < 12; fi23++ {
            p0 = ((!p0) || p0)
            p0 = (0.0 >= float64(((0.0 / -2.25) * 1e10)))
            for fi24 := 0; fi24 < 10; fi24++ {
                v25 := GenS3{F0: 1, F1: p0}
                _ = fi24
                _ = v25
            }
            _ = fi23
        }
        p0 = (!(96 <= 0))
        if (42 > 0) {
            var v26 float64 = 1.5
            if (!true) {
                var v27 float64 = (1e10 + (v26 - 1.5))
                v28 := GenS4{F0: 32768, F1: 274, F2: 127}
                _ = v27
                _ = v28
            }
            var v29 int16 = 42
            _ = v26
            _ = v29
        } else {
            var v30 uint16 = 670
            var v31 float32 = (1.5 + 0.1)
            _ = v30
            _ = v31
        }
    } else {
        for fi32 := 0; fi32 < 13; fi32++ {
            var v33 uint16 = 42
            var v34 float32 = 1e10
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    fmt.Println(p0)
    v35 := []int{-9223372036854775808}
    var v36 uint8 = 2
    _ = p0
    _ = v35
    _ = v36
    return 0
}

func main() {
    v1 := []string{((("\n\t" + "foo bar") + ("hello" + ((("\n\t" + "foo bar") + "") + ("hello" + "")))) + "a"), "hello", "foo bar"}
    var v2 string = (("hello" + (("hello" + "hello") + (("foo bar" + ("a" + "\n\t")) + ("\n\t" + "a")))) + ("a" + "foo bar"))
    var v3 uint16 = 1
    if len(v1) > 0 {
        v1[0] = v2
    }
    if true {
        var v37 int16 = 430
        fmt.Println(v2, v37)
        if true {
            _ = len(v1)
            switch ((((v2 + ((v2 + "a") + "foo bar")) + v2) + "a") + (v2 + "a")) {
            case "hello":
                v2 = ((v2 + "") + v2)
            case "\n\t":
                v3 = (^(v3))
            case "a":
                if false {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi38 := 0; fi38 < 18; fi38++ {
                        var v39 bool = false
                        v3 -= 1
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi38
                        _ = v39
                    }
                    v40 := GenS3{F0: 68, F1: (genFn1((!(!(float64((1.5 - -2.25)) > 0.1)))) > -3)}
                    _ = v40
                } else {
                    v41 := GenS3{F0: genFn1(((false || ((2147483648 < 749) || true)) && (968 != 2))), F1: (!(v37 != (-(v37))))}
                    _ = v41
                }
            }
            var v42 uint16 = v3
            _ = v42
        }
        _ = v37
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 = v2
    }
    v3 = (v3 >> ((1) & 63))
    var v43 int8 = 0
    v44 := GenS1{F0: 962}
    if false {
        var v45 int64 = -9223372036854775808
        _ = v45
    } else {
        var v46 float64 = 0.1
        _ = v46
    }
    for fi47 := 0; fi47 < 15; fi47++ {
        fmt.Println(v3, v3, v2)
        v48 := GenS2{F0: -2147483648}
        _ = fi47
        _ = v48
    }
    var v49 uint16 = 207
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v49
}
