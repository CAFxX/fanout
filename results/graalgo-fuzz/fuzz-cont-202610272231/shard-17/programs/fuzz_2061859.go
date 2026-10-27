package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 string; F2 int16 }
func genFn0() int {
    var v5 uint64 = 493
    v6 := []int{-1, -9223372036854775808}
    fmt.Println(v5)
    v5 = 715
    if true {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 float32 = 0.0
            _ = fi7
            _ = v8
        }
    }
    _ = v5
    _ = v6
    return -208
}
func genFn1(p0 string, p1 bool, p2 string) uint8 {
    v9 := []int{128}
    for fi10 := 0; fi10 < 4; fi10++ {
        v9 = append(v9, fi10)
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: 42, F1: 1}
            _ = fi11
            _ = v12
        }
        switch -124 {
        case 128:
            p0 = "a"
        case -1:
            p0 = ((p2 + p2) + ((p2 + p0) + (p0 + "\n\t")))
        case 402:
            var v13 bool = ((!(p1 || p1)) && (!(!p1)))
            _ = v13
        }
        _ = fi10
    }
    fmt.Println(p0, p2)
    switch p2 {
    case "foo bar":
        var v14 uint32 = 2147483648
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 136
}

func main() {
    var v1 uint8 = 255
    var v2 int = 127
    var v3 int16 = -32768
    var v4 float32 = 0.1
    var v15 uint64 = 856
    var v16 float32 = ((v4 * (((0.0 + 1.5) + float32(1.5)) - ((v4 * v4) * v4))) + float32((1.5 * -2.25)))
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 float32 = v16
        if (true && false) {
            var v19 int8 = -1
            var v20 int8 = 1
            _ = v19
            _ = v20
        }
        var v21 string = ((("" + ("" + "foo bar")) + ("\n\t" + ("hello" + ("a" + "a")))) + "")
        _ = fi17
        _ = v18
        _ = v21
    }
    switch ("foo bar" + "a") {
    case "\n\t":
        v22 := GenS2{F0: 100}
        _ = v22
    default:
        for fi23 := 0; fi23 < 9; fi23++ {
            for fi24 := 0; fi24 < 19; fi24++ {
                switch -9223372036854775808 {
                case -694:
                    var v25 int32 = 696
                    _ = v25
                case 4294967295:
                    v4 = v4
                default:
                    v3 = -1
                }
                var v26 uint8 = 128
                v27 := []int{127}
                _ = fi24
                _ = v26
                _ = v27
            }
            var v28 int64 = -139
            _ = fi23
            _ = v28
        }
    }
    v29 := []int{genFn0(), (v2 | (v2 & 9223372036854775807))}
    if false {
        var v30 int = (v2 % ((v2) | 1))
        var v31 string = "foo bar"
        var v32 uint16 = 32767
        _ = v30
        _ = v31
        _ = v32
    } else {
        var v33 int8 = 100
        v34 := GenS2{F0: v33}
        _ = v33
        _ = v34
    }
    switch "a" {
    case "":
        if (!(false || true)) {
            v3 = v3
        } else {
            var v35 uint = 42
            v36 := GenS2{F0: -1}
            v37 := GenS1{F0: 1, F1: v1}
            _ = v35
            _ = v36
            _ = v37
        }
    case "foo bar":
        for fi38 := 0; fi38 < 2; fi38++ {
            for fi39 := 0; fi39 < 4; fi39++ {
                switch (fi39 + 0) {
                case 0:
                    v40 := GenS2{F0: -1}
                    _ = v40
                case -2147483648:
                    _ = len(v29)
                case 4294967295:
                    if (42 < 0) {
                        if (false && false) {
                            if true {
                                if (((fi39 | fi38) != v2) && (false && (true && (1 > 168)))) {
                                    fmt.Println(v1, fi38, fi38)
                                    v41 := GenS1{F0: 9223372036854775807, F1: 146}
                                    var v42 float32 = v4
                                    _ = v41
                                    _ = v42
                                }
                                v43 := GenS2{F0: 127}
                                if len(v29) >= 2 {
                                    v29 = v29[1:len(v29)]
                                }
                                _ = v43
                            }
                            v16 += (0.1 * v4)
                            v2 = (fi39 + (v2 / ((v2) | 1)))
                        } else {
                            switch "\n\t" {
                            case "hello":
                                if len(v29) > 0 {
                                    v29[0] = v2
                                }
                            case "":
                                v44 := GenS3{F0: v1}
                                var v45 uint16 = 0
                                _ = v44
                                _ = v45
                            case "a":
                                v46 := GenS4{F0: -243, F1: "hello", F2: (^(v3))}
                                _ = v46
                            default:
                                fmt.Println(fi38, v16, v4)
                                v2 = (fi38 / ((genFn0()) | 1))
                            }
                            var v47 string = "a"
                            _ = v47
                        }
                    } else {
                        switch (fi39 & fi38) {
                        case 1:
                            if (false && (0.1 == 3.14159)) {
                                if len(v29) >= 2 {
                                    v29 = v29[1:len(v29)]
                                }
                                v3 -= (v3 ^ v3)
                                var v48 int8 = 100
                                _ = v48
                            }
                        default:
                            v15 = v15
                        }
                        v1 = v1
                        fmt.Println(v2)
                    }
                    var v49 bool = (42 == 2)
                    _ = v49
                }
                v15 = 939
                fmt.Println(fi39, v15, v15)
                _ = fi39
            }
            _ = fi38
        }
    default:
        fmt.Println(v3)
    }
    var v50 float32 = 3.14159
    if len(v29) > 0 {
        v29[0] = 580
    }
    for fi51 := 0; fi51 < 20; fi51++ {
        v1 = v1
        _ = fi51
    }
    var v52 uint = 65535
    var v53 int16 = (v3 << (((^(v52))) & 63))
    var v54 string = ""
    v50 += float32(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v50
    _ = v52
    _ = v53
    _ = v54
}
