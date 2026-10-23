package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 string; F1 uint; F2 int16 }
type GenS11 struct { F0 int }
func genFn0() int8 {
    for fi4 := 0; fi4 < 2; fi4++ {
        switch "hello" {
        case "\n\t":
            v5 := GenS1{F0: 0}
            _ = v5
        }
        var v6 float64 = 3.14159
        _ = fi4
        _ = v6
    }
    var v7 int16 = 1000
    _ = v7
    return 17
}
func genFn1(p0 int8) int32 {
    var v8 int32 = 0
    v8 -= (v8 & 2147483647)
    _ = p0
    _ = v8
    return 42
}

func main() {
    var v1 float64 = 1e10
    var v2 int32 = -520
    v3 := []int{-1, 42, 1}
    var v9 uint64 = 24
    if (!(true || (float32(v1) < (float32((v1 / -2.25)) + (float32(float64(1e10)) - 0.1))))) {
        v2 -= v2
    } else {
        if (true || ((1e10 - (3.14159 / 1.0)) < float32(((((float64((3.14159 - 1e10)) / v1) - (v1 - (v1 * v1))) * (v1 + 0.1)) * v1)))) {
            if true {
                fmt.Println(v1, v2, v1)
                for fi10 := 0; fi10 < 15; fi10++ {
                    v11 := GenS2{F0: 1, F1: false, F2: 42}
                    switch fi10 {
                    case -2147483648:
                        switch ("\n\t" + (("a" + "\n\t") + "\n\t")) {
                        case "foo bar":
                            if false {
                                for fi12 := 0; fi12 < 13; fi12++ {
                                    v2 = 2147483647
                                    var v13 uint8 = 42
                                    if (fi12 >= fi12) {
                                        switch (("a" + ("\n\t" + ((("" + "") + "a") + "a"))) + "foo bar") {
                                        case "a":
                                            var v14 uint16 = 32767
                                            _ = v14
                                        default:
                                            fmt.Println(v2, v2)
                                        }
                                        var v15 string = "hello"
                                        v1 -= v1
                                        _ = v15
                                    }
                                    if false {
                                        v2 += v2
                                        var v16 string = ""
                                        _ = v16
                                    } else {
                                        var v17 uint32 = 0
                                        _ = v17
                                    }
                                    if true {
                                        v18 := GenS1{F0: v9}
                                        var v19 string = (("a" + (("\n\t" + ("a" + "a")) + "\n\t")) + "hello")
                                        _ = v18
                                        _ = v19
                                    }
                                    var v20 int16 = 0
                                    _ = fi12
                                    _ = v13
                                    _ = v20
                                }
                                fmt.Println(v2)
                                var v21 float32 = (0.0 - (((1.5 * 1e10) + 3.14159) / (1.5 + float32(float64((1e10 - ((((1.5 - 1.5) * -2.25) / 1e10) - 1e10)))))))
                                _ = v21
                            }
                        case "hello":
                            switch (("a" + "foo bar") + ("foo bar" + ("foo bar" + "foo bar"))) {
                            case "hello":
                                var v22 int64 = 1
                                _ = v22
                            }
                        default:
                            v23 := GenS3{F0: (fi10 + (fi10 * fi10))}
                            _ = v23
                        }
                    case 0:
                        v1 += 1.5
                    case 2147483647:
                        var v24 uint32 = 0
                        _ = v24
                    }
                    v25 := GenS4{F0: (v2 - -111), F1: 86}
                    _ = fi10
                    _ = v11
                    _ = v25
                }
                if len(v3) > 0 {
                    v3[0] = 1
                }
            }
            if false {
                fmt.Println(v1)
                var v26 int8 = -44
                _ = v26
            }
            v27 := GenS5{F0: -32768, F1: 502}
            _ = v27
        } else {
            v3 = append(v3, 250)
            switch "foo bar" {
            case "foo bar":
                if true {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        var v29 uint64 = 1
                        v2 = genFn1(genFn0())
                        switch fi28 {
                        case 128:
                            fmt.Println(fi28)
                        default:
                            v30 := GenS1{F0: (v29 | 4294967295)}
                            _ = v30
                        }
                        _ = fi28
                        _ = v29
                    }
                    v2 = 1
                    v2 -= genFn1(100)
                }
            default:
                fmt.Println(v1, v9)
            }
        }
    }
    v31 := GenS6{F0: 255, F1: 586}
    for fi32 := 0; fi32 < 13; fi32++ {
        for fi33 := 0; fi33 < 13; fi33++ {
            var v34 int8 = -128
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    v9 -= v9
    switch ((("\n\t" + ("foo bar" + "a")) + "") + "hello") {
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "foo bar":
        var v35 uint = 787
        _ = v35
    case "a":
        var v36 bool = ((1.5 * 1.5) == 1e10)
        _ = v36
    }
    switch ("" + "hello") {
    case "foo bar":
        var v37 int = 128
        _ = v37
    case "hello":
        for fi38 := 0; fi38 < 4; fi38++ {
            switch "a" {
            case "hello":
                v3 = append(v3, -453)
            case "a":
                switch (^(fi38)) {
                case 4294967295:
                    fmt.Println(v1)
                case -9223372036854775808:
                    v39 := GenS2{F0: -32768, F1: ((((((1e10 + 1e10) / (0.1 + 0.1)) + 3.14159) + 1.5) - (3.14159 - 0.1)) < 0.1), F2: 514}
                    _ = v39
                case 459:
                    v3 = append(v3, fi38)
                }
            default:
                fmt.Println(v1, fi38)
            }
            for fi40 := 0; fi40 < 20; fi40++ {
                v41 := GenS7{F0: 4294967295}
                _ = fi40
                _ = v41
            }
            var v42 float64 = v1
            _ = fi38
            _ = v42
        }
    case "\n\t":
        v43 := GenS8{F0: 221}
        _ = v43
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        if len(v3) > 0 {
            v3[0] = 32768
        }
        switch ("" + ((("foo bar" + "\n\t") + ("hello" + "\n\t")) + "a")) {
        case "hello":
            var v45 bool = false
            _ = v45
        case "a":
            if false {
                if false {
                    var v46 uint64 = 18446744073709551615
                    var v47 uint16 = 255
                    _ = v46
                    _ = v47
                } else {
                    v1 += 3.14159
                }
                if true {
                    var v48 string = "foo bar"
                    _ = v48
                }
            }
        case "":
            v9 -= v9
        }
        var v49 uint8 = 128
        _ = fi44
        _ = v49
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v50 bool = (3.14159 <= 0.1)
    var v51 uint8 = 2
    for fi52 := 0; fi52 < 7; fi52++ {
        var v53 bool = (!(2147483648 >= 2))
        for fi54 := 0; fi54 < 19; fi54++ {
            fmt.Println(v51, v51)
            if ((!(1.5 != v1)) && v50) {
                switch "a" {
                case "a":
                    fmt.Println(v51, v51, v1)
                case "":
                    if ((v50 && v53) && ((v51 << ((2) & 63)) < v51)) {
                        if (false && (v1 >= float64(1.5))) {
                            v55 := GenS9{F0: true}
                            if ("hello" != ((("hello" + "") + "hello") + ("\n\t" + ""))) {
                                var v56 int64 = 9223372036854775807
                                v57 := GenS2{F0: 717, F1: v50, F2: 255}
                                _ = v56
                                _ = v57
                            } else {
                                var v58 uint8 = v51
                                _ = v58
                            }
                            _ = v55
                        } else {
                            v59 := GenS10{F0: "\n\t", F1: uint(v51), F2: 1000}
                            v3 = append(v3, fi52)
                            v60 := GenS4{F0: v2, F1: -60}
                            _ = v59
                            _ = v60
                        }
                        v1 = v1
                        fmt.Println(v51, v1, v2)
                    } else {
                        var v61 bool = true
                        _ = v61
                    }
                case "hello":
                    if len(v3) > 0 {
                        v3[0] = fi52
                    }
                    v53 = true
                    if ((v51 | (^(v51))) >= (v51 * (v51 ^ v51))) {
                        var v62 bool = true
                        v2 -= (v2 & v2)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v62
                    }
                }
                v63 := GenS11{F0: fi54}
                var v64 string = "hello"
                _ = v63
                _ = v64
            }
            _ = fi54
        }
        var v65 float32 = 0.0
        _ = fi52
        _ = v53
        _ = v65
    }
    var v66 uint8 = (v51 - (v51 % ((255) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v66:", v66)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v50
    _ = v51
    _ = v66
    _ = v9
}
