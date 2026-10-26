package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint64) string {
    switch 773 {
    case 32767:
        var v4 bool = false
        _ = v4
    case 483:
        switch "hello" {
        case "a":
            p0 += (p0 & p0)
        }
    case 0:
        var v5 float64 = (-2.25 - (((1e10 - -2.25) * (0.0 + 3.14159)) + 0.0))
        _ = v5
    default:
        v6 := []int{127, 312}
        _ = v6
    }
    v7 := GenS1{F0: -32768, F1: 65535, F2: false}
    p0 -= (p0 ^ (^(p0)))
    var v8 uint16 = 2
    switch int(p0) {
    case -855:
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 int16 = 32767
            var v11 int = fi9
            if false {
                v12 := GenS2{F0: -85}
                _ = v12
            } else {
                v13 := GenS2{F0: -115}
                _ = v13
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    var v1 int = 746
    v2 := make(map[string]bool)
    var v3 uint64 = 855
    var v14 int8 = 42
    v14 = v14
    switch (^(v1)) {
    case 9223372036854775807:
        fmt.Println(v3, v3, v1)
    case -2147483648:
        if true {
            var v15 string = ""
            for fi16 := 0; fi16 < 7; fi16++ {
                v3 = (v3 ^ 18446744073709551615)
                _ = fi16
            }
            _ = v15
        }
    case 164:
        v1 -= v1
    default:
        for fi17 := 0; fi17 < 18; fi17++ {
            delete(v2, ("a" + "a"))
            v3 -= (v3 % (((v3 << ((584) & 63))) | 1))
            var v18 uint8 = 127
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v20 := GenS3{F0: (v14 >= 1), F1: 255}
        if true {
            switch "" {
            case "foo bar":
                fmt.Println(v14, fi19)
            case "\n\t":
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 uint16 = 888
                    _ = fi21
                    _ = v22
                }
            default:
                delete(v2, "hello")
            }
            if (32767 == int16(v14)) {
                switch ("a" + "hello") {
                case "\n\t":
                    fmt.Println(v3, fi19)
                case "foo bar":
                    for fi23 := 0; fi23 < 17; fi23++ {
                        switch ("\n\t" + ((("\n\t" + "foo bar") + "foo bar") + "hello")) {
                        case "\n\t":
                            v3 = v3
                        }
                        v14 += (-(v14))
                        switch "hello" {
                        case "foo bar":
                            if true {
                                var v24 uint64 = (^(v3))
                                _ = v24
                            } else {
                                v25 := []int{}
                                fmt.Println(v1)
                                _ = v25
                            }
                        case "":
                            fmt.Println(v1, v14, v1)
                        default:
                            var v26 uint = 1
                            _ = v26
                        }
                        _ = fi23
                    }
                default:
                    if (!(v1 > v1)) {
                        if (1 != 1) {
                            v27 := GenS4{F0: fi19, F1: (v14 < (v14 | 127))}
                            var v28 uint8 = 139
                            v29 := GenS4{F0: v1, F1: ("" >= "foo bar")}
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            var v30 bool = (true && (-128 >= v14))
                            switch (v1 * fi19) {
                            case 127:
                                if v30 {
                                    var v31 string = "\n\t"
                                    var v32 float32 = float32(1e10)
                                    _ = v31
                                    _ = v32
                                } else {
                                    v33 := []int{fi19}
                                    _ = v33
                                }
                            case -945:
                                v30 = (255 > uint16(v3))
                            case 9223372036854775807:
                                fmt.Println(fi19, v1)
                            }
                            _ = v30
                        }
                        var v34 uint64 = v3
                        _ = v34
                    }
                }
                v35 := []int{v1, (fi19 >> ((0) & 63)), 79, v1}
                _ = v35
            }
        }
        v14 = v14
        _ = fi19
        _ = v20
    }
    v36 := GenS5{F0: "a"}
    for fi37 := 0; fi37 < 20; fi37++ {
        var v38 uint16 = 831
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 20; fi39++ {
        delete(v2, "\n\t")
        if (729 != 2147483647) {
            v40 := []int{(fi39 / ((v1) | 1)), v1, (fi39 % (((^(fi39))) | 1)), fi39}
            _ = v40
        } else {
            delete(v2, "hello")
            v41 := GenS6{F0: 2}
            _ = v41
        }
        _ = fi39
    }
    v2[("foo bar" + "foo bar")] = true
    var v42 int64 = -1
    _ = len(v2)
    var v43 bool = (128 <= 128)
    var v44 int64 = (^(v42))
    for fi45 := 0; fi45 < 4; fi45++ {
        var v46 string = ""
        var v47 float32 = 0.0
        v48 := GenS2{F0: v14}
        _ = fi45
        _ = v46
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v42
    _ = v43
    _ = v44
}
