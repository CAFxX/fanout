package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 int; F2 int }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint32) int64 {
    v4 := GenS1{F0: 44}
    fmt.Println(p0)
    if true {
        p0 = (^(p0))
    } else {
        p0 -= 1
        var v5 float64 = (0.0 * ((0.0 - 0.1) / 1.0))
        p0 = 0
        _ = v5
    }
    v6 := []int{32768}
    v7 := GenS1{F0: 100}
    _ = p0
    _ = v4
    _ = v6
    _ = v7
    return 520
}

func main() {
    var v1 string = "a"
    v2 := []int{}
    var v3 float64 = (((float64((((1.5 + -2.25) / 1.0) * (0.1 * 0.0))) + (3.14159 + 1e10)) * (-2.25 / (1.5 * 3.14159))) + (0.1 / 0.1))
    switch v1 {
    case "\n\t":
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS2{F0: 32767, F1: 128}
            v10 := GenS1{F0: -19}
            var v11 float32 = 1.5
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    case "hello":
        var v12 bool = (true && false)
        _ = v12
    case "":
        if false {
            v13 := GenS2{F0: 32767, F1: 1}
            v1 = ((v1 + v1) + "\n\t")
            fmt.Println(v3, v3)
            _ = v13
        }
    default:
        v14 := GenS3{F0: 228, F1: "foo bar"}
        _ = v14
    }
    if false {
        if (84 >= 121) {
            v3 = v3
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS4{F0: v1}
                v3 -= v3
                _ = fi15
                _ = v16
            }
        } else {
            if len(v2) > 0 {
                v2[0] = -578
            }
        }
        switch v1 {
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = 32768
            }
        case "a":
            v17 := GenS5{F0: 4294967295, F1: 32767, F2: 4294967295}
            _ = v17
        default:
            v3 = (0.1 / 1e10)
        }
    } else {
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 uint32 = 2147483647
            var v20 uint16 = 255
            fmt.Println(v20)
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v1 = ""
    fmt.Println(v1, v3)
    for fi21 := 0; fi21 < 13; fi21++ {
        fmt.Println(v1, v1, v3)
        var v22 uint64 = 987
        fmt.Println(v3)
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        v1 = v1
        v24 := GenS5{F0: 435, F1: (fi23 * (fi23 | 127)), F2: 128}
        _ = fi23
        _ = v24
    }
    if (2 <= 2147483647) {
        switch ("a" + (v1 + "foo bar")) {
        case "a":
            _ = len(v2)
        case "hello":
            v25 := GenS6{F0: 42, F1: v1}
            _ = v25
        case "foo bar":
            if false {
                var v26 uint8 = 253
                v1 = ""
                _ = v26
            } else {
                if true {
                    v1 = (v1 + "")
                } else {
                    if false {
                        v1 = "hello"
                        for fi27 := 0; fi27 < 2; fi27++ {
                            v28 := GenS4{F0: "\n\t"}
                            _ = fi27
                            _ = v28
                        }
                        var v29 int32 = -1
                        _ = v29
                    } else {
                        v30 := GenS2{F0: 65535, F1: 0}
                        _ = v30
                    }
                    v1 = "a"
                    v1 = ("a" + ((("a" + ("\n\t" + (v1 + v1))) + "foo bar") + ("hello" + v1)))
                }
                switch (((("a" + ("foo bar" + v1)) + v1) + ("foo bar" + v1)) + "") {
                case "hello":
                    if (496 > 1) {
                        var v31 int = -2147483648
                        v32 := GenS7{F0: 18446744073709551615}
                        var v33 bool = ((true || (false && (3.14159 < 0.0))) && false)
                        _ = v31
                        _ = v32
                        _ = v33
                    } else {
                        var v34 uint32 = 1
                        v35 := GenS8{F0: (v34 + v34), F1: -1}
                        _ = v34
                        _ = v35
                    }
                }
                for fi36 := 0; fi36 < 19; fi36++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v37 uint16 = 1
                    _ = fi36
                    _ = v37
                }
            }
        }
    }
    v2 = append(v2, 9223372036854775807)
    if len(v2) > 0 {
        v2[0] = -783
    }
    if false {
        v3 -= -2.25
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = len(v2)
    } else {
        if true {
            var v38 int64 = -9223372036854775808
            _ = v38
        }
        for fi39 := 0; fi39 < 8; fi39++ {
            var v40 uint8 = 128
            _ = fi39
            _ = v40
        }
        for fi41 := 0; fi41 < 14; fi41++ {
            switch (fi41 * (fi41 - 805)) {
            case 597:
                fmt.Println(v1)
            }
            var v42 uint64 = 948
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi41
            _ = v42
        }
    }
    if false {
        v43 := GenS8{F0: 2147483648, F1: 814}
        fmt.Println(v1)
        var v44 int = 436
        _ = v43
        _ = v44
    }
    if ((3.14159 + (v3 / (v3 - float64(0.0)))) == ((v3 / ((v3 * v3) / v3)) - v3)) {
        switch 128 {
        case 4294967295:
            _ = len(v2)
        }
    }
    _ = len(v2)
    v2 = append(v2, 4294967295)
    fmt.Println(v1, v3)
    _ = len(v2)
    switch v1 {
    case "":
        fmt.Println(v1)
    }
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
