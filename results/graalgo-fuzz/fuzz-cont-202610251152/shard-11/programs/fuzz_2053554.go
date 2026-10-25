package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string; F1 uint }
func genFn0(p0 uint, p1 uint16) uint {
    if true {
        if true {
            if false {
                var v6 bool = false
                for fi7 := 0; fi7 < 3; fi7++ {
                    for fi8 := 0; fi8 < 18; fi8++ {
                        v9 := []int{-1, (fi8 * fi7), (fi7 - fi8), (fi8 % (((-(fi8))) | 1))}
                        v6 = (1000 == -32768)
                        _ = fi8
                        _ = v9
                    }
                    switch "hello" {
                    case "hello":
                        var v10 float64 = 3.14159
                        _ = v10
                    }
                    var v11 uint32 = 2147483647
                    _ = fi7
                    _ = v11
                }
                v12 := GenS1{F0: (("" + "foo bar") + ("hello" + "\n\t")), F1: 0}
                _ = v12
                _ = v6
            }
            var v13 float32 = 0.1
            _ = v13
        } else {
            fmt.Println(p0)
        }
        if (2 != 474) {
            p0 = 1
            p0 += p0
            v14 := []int{127, 4294967295, 1, -1}
            _ = v14
        } else {
            switch "a" {
            case "\n\t":
                p0 = p0
            default:
                var v15 int32 = 42
                _ = v15
            }
            for fi16 := 0; fi16 < 18; fi16++ {
                p1 += p1
                if true {
                    p1 = 255
                } else {
                    switch 2147483647 {
                    case 9223372036854775807:
                        for fi17 := 0; fi17 < 20; fi17++ {
                            var v18 uint64 = 1
                            switch ((("hello" + (("\n\t" + ("a" + "\n\t")) + ("hello" + "foo bar"))) + "") + "foo bar") {
                            case "a":
                                switch 42 {
                                case 127:
                                    fmt.Println(p1, p1, fi17)
                                case 387:
                                    v19 := []int{}
                                    _ = v19
                                default:
                                    if (fi16 > (fi17 >> (((^(p0))) & 63))) {
                                        var v20 int64 = 42
                                        switch ("" + "foo bar") {
                                        case "a":
                                            fmt.Println(v20)
                                        }
                                        _ = v20
                                    }
                                }
                            case "\n\t":
                                var v21 int16 = 42
                                _ = v21
                            case "":
                                p0 += (^(p0))
                                p1 -= (^(p1))
                            default:
                                var v22 uint32 = 491
                                _ = v22
                            }
                            var v23 int = 825
                            var v24 int8 = -104
                            _ = fi17
                            _ = v18
                            _ = v23
                            _ = v24
                        }
                    case 1:
                        if (!(!true)) {
                            v25 := []int{-2147483648}
                            var v26 int16 = 42
                            v27 := GenS1{F0: "\n\t", F1: 0}
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    case -9223372036854775808:
                        if false {
                            fmt.Println(p1, fi16, p1)
                            var v28 int64 = int64(p1)
                            _ = v28
                        } else {
                            p1 = 331
                            p1 = 255
                            var v29 int32 = 42
                            _ = v29
                        }
                    }
                    for fi30 := 0; fi30 < 18; fi30++ {
                        p0 = p0
                        _ = fi30
                    }
                }
                _ = fi16
            }
        }
        v31 := []int{}
        _ = v31
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        var v33 int8 = 1
        var v34 string = "foo bar"
        _ = fi32
        _ = v33
        _ = v34
    }
    p0 = 2
    v35 := GenS1{F0: "\n\t", F1: 138}
    var v36 int = 128
    _ = p0
    _ = p1
    _ = v35
    _ = v36
    return 1
}

func main() {
    var v1 int8 = -1
    var v2 int16 = 1000
    var v3 uint = 1
    v4 := make(map[int]string)
    v5 := make(map[int]int)
    fmt.Println(v2, v1, v1)
    var v37 uint8 = 86
    v2 = v2
    for fi38 := 0; fi38 < 4; fi38++ {
        delete(v4, (fi38 & -9223372036854775808))
        var v39 uint = v3
        var v40 uint8 = v37
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 uint = genFn0((v3 | (v3 << ((v3) & 63))), uint16(v37))
    v42 := v5[-9223372036854775808]
    switch (v42 ^ (v42 * 662)) {
    case 283:
        var v43 uint = (v3 | v41)
        _ = v43
    case -494:
        v37 += v37
    default:
        switch 645 {
        case 650:
            fmt.Println(v42, v2)
        }
    }
    v44 := GenS1{F0: "foo bar", F1: 1}
    var v45 uint64 = 42
    switch ("foo bar" + ("" + "a")) {
    case "\n\t":
        v46 := GenS2{F0: ("foo bar" + "\n\t"), F1: v3}
        _ = v46
    case "":
        for fi47 := 0; fi47 < 3; fi47++ {
            v3 = v3
            _ = fi47
        }
    }
    if false {
        var v48 uint32 = 0
        _ = v48
    } else {
        fmt.Println(v42, v1, v2)
    }
    fmt.Println(v45, v45, v37)
    var v49 uint8 = v37
    var v50 int32 = 902
    var v51 uint8 = (v37 * (v37 & v37))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v44
    _ = v45
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
