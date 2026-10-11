package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int8; F1 uint8; F2 uint32 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint8 = 63
    v3 := []int{1, -9223372036854775808, -2147483648, 2147483647}
    var v4 int32 = 1
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 10; fi6++ {
            fmt.Println(fi6)
            var v7 int8 = 104
            _ = fi6
            _ = v7
        }
        var v8 int16 = 968
        var v9 uint64 = 357
        _ = fi5
        _ = v8
        _ = v9
    }
    var v10 string = "\n\t"
    v4 = int32(v2)
    for fi11 := 0; fi11 < 18; fi11++ {
        if false {
            v3 = append(v3, (^(fi11)))
            if len(v3) > 0 {
                v3[0] = fi11
            }
        }
        switch (fi11 >> ((980) & 63)) {
        case 2147483647:
            v12 := GenS1{F0: 0}
            _ = v12
        case -9223372036854775808:
            if len(v3) > 0 {
                v3[0] = fi11
            }
        case 4294967295:
            var v13 uint64 = 188
            _ = v13
        default:
            if false {
                if len(v3) > 0 {
                    v3[0] = -2147483648
                }
            } else {
                v4 = v4
            }
        }
        v1 += 0
        _ = fi11
    }
    v2 -= 0
    v14 := GenS2{F0: -107, F1: 2147483647, F2: 42}
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := GenS1{F0: 42}
        for fi17 := 0; fi17 < 16; fi17++ {
            if false {
                if (((!(1000 != -1)) || (!((!(1e10 >= 3.14159)) && false))) || false) {
                    if (false && ("\n\t" < v10)) {
                        if true {
                            if ((-2.25 + 3.14159) != (3.14159 - 1.5)) {
                                if false {
                                    v1 -= (-(v1))
                                    var v18 float32 = 1e10
                                    _ = v18
                                }
                                switch v10 {
                                case "\n\t":
                                    v19 := GenS2{F0: 42, F1: 0, F2: 125}
                                    _ = v19
                                case "foo bar":
                                    fmt.Println(v10, v2)
                                }
                            }
                            v20 := GenS2{F0: -1, F1: 612, F2: 100}
                            _ = v20
                        }
                        var v21 int16 = 1000
                        switch fi17 {
                        case -515:
                            if ((!(223 < int64(fi15))) && ((!(42 == 1)) && true)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v22 float64 = 0.1
                                fmt.Println(v4)
                                _ = v22
                            } else {
                                var v23 int = 0
                                switch v10 {
                                case "":
                                    _ = len(v3)
                                case "hello":
                                    switch v10 {
                                    case "a":
                                        var v24 int8 = -1
                                        _ = v24
                                    case "foo bar":
                                        fmt.Println(fi17, v21)
                                    case "hello":
                                        var v25 float32 = 0.1
                                        _ = v25
                                    }
                                case "a":
                                    v21 = 217
                                default:
                                    switch v10 {
                                    case "hello":
                                        v1 = 0
                                    }
                                }
                                _ = v23
                            }
                            v26 := GenS2{F0: 0, F1: 42, F2: 42}
                            _ = v26
                        }
                        _ = v21
                    }
                    _ = len(v3)
                } else {
                    switch ("" + v10) {
                    case "\n\t":
                        v1 += v1
                    default:
                        switch v10 {
                        case "hello":
                            v1 = -1
                        case "a":
                            v1 += (v1 * (v1 + (v1 * (v1 >> ((0) & 63)))))
                        case "\n\t":
                            var v27 uint8 = 128
                            _ = v27
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        v4 -= 1
                    }
                    fmt.Println(fi17, v10)
                    _ = len(v3)
                }
                var v28 float64 = (0.0 / (3.14159 + (0.0 * 0.0)))
                v1 += v1
                _ = v28
            }
            v1 = (-(v1))
            _ = fi17
        }
        v29 := GenS1{F0: 0}
        _ = fi15
        _ = v16
        _ = v29
    }
    v30 := GenS2{F0: -70, F1: 2147483648, F2: 28}
    fmt.Println(v10)
    switch 42 {
    case 2147483647:
        for fi31 := 0; fi31 < 15; fi31++ {
            v10 = v10
            _ = fi31
        }
    case 265:
        for fi32 := 0; fi32 < 1; fi32++ {
            if false {
                switch ("a" + (v10 + "a")) {
                case "hello":
                    switch 9223372036854775807 {
                    case 199:
                        switch v10 {
                        case "\n\t":
                            fmt.Println(v10, v2, v10)
                        case "hello":
                            v1 -= 42
                        case "foo bar":
                            v33 := GenS3{F0: (^(v2))}
                            _ = v33
                        }
                    case -878:
                        var v34 uint = 255
                        _ = v34
                    }
                case "a":
                    _ = len(v3)
                default:
                    if len(v3) > 0 {
                        v3[0] = 85
                    }
                }
            }
            _ = fi32
        }
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        var v36 int8 = -1
        var v37 string = v10
        for fi38 := 0; fi38 < 15; fi38++ {
            v3 = append(v3, fi38)
            v36 = (v36 ^ (v36 & v36))
            _ = fi38
        }
        _ = fi35
        _ = v36
        _ = v37
    }
    if true {
        if len(v3) > 0 {
            v3[0] = -1
        }
        switch 2147483647 {
        case 127:
            for fi39 := 0; fi39 < 1; fi39++ {
                v40 := GenS3{F0: (v2 << ((4294967295) & 63))}
                if (0.1 <= 1.5) {
                    _ = len(v3)
                    var v41 int64 = v1
                    _ = v41
                }
                _ = fi39
                _ = v40
            }
        default:
            v42 := GenS4{F0: 372, F1: 2147483647}
            _ = v42
        }
        v43 := GenS3{F0: (v2 | v2)}
        _ = v43
    } else {
        var v44 int = 1
        if false {
            var v45 float64 = 3.14159
            var v46 int16 = 1000
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v45
            _ = v46
        } else {
            v3 = append(v3, v44)
        }
        _ = v44
    }
    v47 := GenS3{F0: v2}
    fmt.Println(v1, v2, v10)
    var v48 uint16 = 681
    var v49 uint32 = 2
    v50 := GenS5{F0: 100, F1: (v2 % ((v2) | 1)), F2: v49}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
