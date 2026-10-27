package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint32; F1 bool; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 int8 = -111
    var v2 int = 442
    var v3 int16 = 0
    switch v2 {
    case 1:
        v4 := GenS1{F0: 0, F1: "\n\t"}
        _ = v4
    }
    if false {
        for fi5 := 0; fi5 < 8; fi5++ {
            if false {
                fmt.Println(v2, v3, v1)
            }
            if true {
                v6 := GenS2{F0: 2147483648, F1: ((((1e10 + ((0.0 - 1e10) + float64((float32(0.0) * 0.0)))) * 0.0) - (1.5 / 3.14159)) != (1e10 + (float64(0.1) * (1e10 / (0.0 - 1e10))))), F2: "hello"}
                v1 += (v1 - 42)
                _ = v6
            } else {
                var v7 float64 = (0.1 + 3.14159)
                v8 := GenS3{F0: ("" + "\n\t")}
                v2 += (v2 << ((2) & 63))
                _ = v7
                _ = v8
            }
            _ = fi5
        }
    } else {
        var v9 uint64 = 103
        if true {
            for fi10 := 0; fi10 < 8; fi10++ {
                switch (-(v2)) {
                case 128:
                    v11 := GenS4{F0: 65535, F1: (v3 + (v3 - int16(v1))), F2: (^(v9))}
                    _ = v11
                case 32768:
                    switch (("" + "foo bar") + ("foo bar" + ("" + "hello"))) {
                    case "foo bar":
                        switch (((("hello" + "a") + "hello") + "a") + "a") {
                        case "foo bar":
                            switch (-(fi10)) {
                            case 0:
                                if (-1 >= -2147483648) {
                                    var v12 int = (fi10 & (^(fi10)))
                                    v3 = 0
                                    v12 = (fi10 - (v12 * (^(v12))))
                                    _ = v12
                                }
                            case 32768:
                                fmt.Println(v2, v1)
                            }
                        case "\n\t":
                            fmt.Println(v2, v1)
                        case "":
                            v13 := []int{32767, 32767}
                            _ = v13
                        }
                    case "\n\t":
                        fmt.Println(v9, fi10, v3)
                    case "hello":
                        fmt.Println(v9)
                    }
                case -2147483648:
                    var v14 int = fi10
                    _ = v14
                }
                fmt.Println(v1, v9)
                _ = fi10
            }
        } else {
            for fi15 := 0; fi15 < 20; fi15++ {
                fmt.Println(v2, fi15, v9)
                _ = fi15
            }
            fmt.Println(v2, v1)
            switch "" {
            case "a":
                var v16 float32 = float32(float64(1.5))
                _ = v16
            default:
                switch "a" {
                case "foo bar":
                    v17 := []int{(v2 - (v2 >> ((42) & 63))), 42}
                    _ = v17
                case "hello":
                    v18 := GenS3{F0: ""}
                    _ = v18
                default:
                    switch (("a" + "hello") + ("foo bar" + ("foo bar" + "\n\t"))) {
                    case "hello":
                        if true {
                            v19 := GenS1{F0: 0, F1: ((("foo bar" + ((((("" + "foo bar") + (("hello" + (("hello" + "hello") + "")) + "")) + "foo bar") + "\n\t") + ("\n\t" + "a"))) + "\n\t") + "\n\t")}
                            var v20 int64 = 1
                            v21 := []int{}
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            v2 += v2
                            v22 := []int{(v2 << ((2) & 63)), (-(v2))}
                            _ = v22
                        }
                    case "":
                        v23 := []int{v2, (v2 / ((v2) | 1)), 549, (v2 >> ((255) & 63))}
                        _ = v23
                    case "a":
                        var v24 float32 = ((-2.25 + ((-2.25 / 0.1) + 0.1)) + (-2.25 * (1.5 + float32(0.0))))
                        _ = v24
                    }
                }
            }
        }
        _ = v9
    }
    v25 := GenS3{F0: ("\n\t" + "foo bar")}
    var v26 string = (("" + ("foo bar" + ("\n\t" + "\n\t"))) + ("a" + "\n\t"))
    v27 := GenS1{F0: -9223372036854775808, F1: ("a" + ((("hello" + (v26 + ((v26 + v26) + (v26 + ((v26 + ("a" + "")) + v26))))) + v26) + v26))}
    if false {
        if ((^(v1)) < (^(v1))) {
            v26 = "hello"
            v28 := []int{353, (^(v2))}
            for fi29 := 0; fi29 < 15; fi29++ {
                v2 -= -2147483648
                switch "" {
                case "\n\t":
                    switch v26 {
                    case "a":
                        v28 = append(v28, (fi29 << ((644) & 63)))
                    default:
                        v30 := GenS2{F0: 2, F1: (-460 == -325), F2: v26}
                        _ = v30
                    }
                default:
                    var v31 uint64 = 18446744073709551615
                    _ = v31
                }
                v28 = append(v28, (-(v2)))
                _ = fi29
            }
            _ = v28
        } else {
            v32 := GenS4{F0: 255, F1: v3, F2: 4294967295}
            v2 -= v2
            _ = v32
        }
        v1 -= v1
    }
    fmt.Println(v1, v2)
    var v33 int64 = 9223372036854775807
    var v34 uint64 = 790
    switch 4294967295 {
    case 42:
        for fi35 := 0; fi35 < 18; fi35++ {
            var v36 string = "\n\t"
            _ = fi35
            _ = v36
        }
    case -212:
        if false {
            v37 := GenS4{F0: 42, F1: -1, F2: (v34 % ((v34) | 1))}
            _ = v37
        }
    case 973:
        v2 -= (v2 << ((88) & 63))
    }
    v38 := []int{298, (v2 & (^(v2)))}
    v2 = v2
    switch (v2 | int(v1)) {
    case 622:
        var v39 bool = false
        _ = v39
    case -1:
        v40 := GenS5{F0: -761}
        _ = v40
    }
    v41 := GenS4{F0: 69, F1: (-(v3)), F2: 445}
    if len(v38) >= 2 {
        v38 = v38[1:len(v38)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v38:", len(v38), v38)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v38
    _ = v41
}
