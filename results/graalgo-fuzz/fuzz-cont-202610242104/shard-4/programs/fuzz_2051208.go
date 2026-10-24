package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int16 }
type GenS2 struct { F0 int16; F1 int8; F2 int }
type GenS3 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int16; F1 int64 }
func genFn0(p0 int8, p1 bool) int8 {
    v5 := []int{}
    p0 += (p0 + (^(p0)))
    p1 = (!(true && (((1e10 / 1.5) / -2.25) > -2.25)))
    var v6 string = "\n\t"
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}
func genFn1(p0 uint, p1 int, p2 uint8) int16 {
    var v7 int32 = -1
    if ("\n\t" > "foo bar") {
        p2 += 255
        switch p1 {
        case 32767:
            switch "foo bar" {
            case "hello":
                var v8 uint16 = 32767
                _ = v8
            case "":
                var v9 int = 32768
                _ = v9
            case "a":
                v10 := GenS1{F0: 18446744073709551615, F1: -2147483648, F2: 0}
                _ = v10
            default:
                v7 = v7
            }
        }
        var v11 bool = true
        _ = v11
    }
    var v12 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = (("a" + "a") < "\n\t")
    v3 := []int{814, 391, -687, 32768}
    var v4 uint8 = 172
    var v13 bool = (-128 < genFn0(-128, v2))
    v2 = v13
    v14 := GenS2{F0: 42, F1: -1, F2: 42}
    switch 4294967295 {
    case -36:
        if true {
            fmt.Println(v13)
        } else {
            if false {
                for fi15 := 0; fi15 < 15; fi15++ {
                    v2 = (1.5 >= (-2.25 / 1.0))
                    var v16 int = (fi15 * (fi15 + (fi15 + 4294967295)))
                    _ = fi15
                    _ = v16
                }
                for fi17 := 0; fi17 < 19; fi17++ {
                    v13 = (1e10 <= 1.5)
                    if v13 {
                        var v18 int16 = -32768
                        v19 := GenS2{F0: (v18 % ((-189) | 1)), F1: -44, F2: (fi17 + (fi17 - (fi17 - 9223372036854775807)))}
                        _ = v18
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 18; fi20++ {
                            v21 := GenS1{F0: 618, F1: -846, F2: genFn1(42, 9223372036854775807, (v4 >> ((2) & 63)))}
                            if ("foo bar" <= "a") {
                                v22 := GenS1{F0: 4294967295, F1: -81, F2: -669}
                                v4 = (v4 >> ((2) & 63))
                                switch ("a" + "") {
                                case "\n\t":
                                    _ = len(v1)
                                default:
                                    v2 = (!true)
                                }
                                _ = v22
                            }
                            v13 = v13
                            _ = fi20
                            _ = v21
                        }
                        fmt.Println(v2, v4)
                    }
                    _ = fi17
                }
            } else {
                for fi23 := 0; fi23 < 8; fi23++ {
                    v4 += v4
                    var v24 int64 = 42
                    _ = fi23
                    _ = v24
                }
                v25 := GenS3{F0: 759, F1: 18446744073709551615, F2: 1}
                _ = v25
            }
        }
    case 9223372036854775807:
        var v26 float32 = 0.1
        _ = v26
    case 127:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        for fi28 := 0; fi28 < 14; fi28++ {
            v29 := GenS4{F0: 4}
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    v2 = false
    if (!false) {
        switch 627 {
        case -707:
            fmt.Println(v4, v2)
        case -1:
            v30 := GenS5{F0: "foo bar", F1: -157}
            _ = v30
        case 127:
            if (!false) {
                v31 := v1[""]
                var v32 int8 = 0
                _ = v31
                _ = v32
            }
        }
        var v33 int8 = 100
        v2 = (((!false) || ("a" > "\n\t")) || (!true))
        _ = v33
    }
    var v34 uint8 = v4
    for fi35 := 0; fi35 < 20; fi35++ {
        if true {
            var v36 uint = 273
            v37 := GenS6{F0: -32768, F1: 80}
            fmt.Println(v2, v36)
            _ = v36
            _ = v37
        } else {
            switch (-(fi35)) {
            case -2147483648:
                var v38 int16 = 42
                _ = v38
            case 32767:
                if (25 < 361) {
                    v1["foo bar"] = (fi35 % ((fi35) | 1))
                }
            case -490:
                if v13 {
                    for fi39 := 0; fi39 < 16; fi39++ {
                        v40 := GenS5{F0: "a", F1: 32767}
                        _ = fi39
                        _ = v40
                    }
                }
            }
        }
        _ = fi35
    }
    v3 = append(v3, -269)
    v2 = (767 > 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v4
}
