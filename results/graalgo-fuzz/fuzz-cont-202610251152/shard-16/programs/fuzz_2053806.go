package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 string; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 bool }
func genFn0() string {
    v4 := []int{128, 32767, 2147483647}
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 string = "a"
        switch "\n\t" {
        case "a":
            for fi7 := 0; fi7 < 16; fi7++ {
                if true {
                    var v8 uint = 65535
                    switch (fi7 | fi7) {
                    case -2147483648:
                        fmt.Println(fi7, fi7)
                    default:
                        v9 := GenS1{F0: 2147483647, F1: 109}
                        _ = v9
                    }
                    var v10 int32 = -570
                    if (!((1 == -128) || false)) {
                        var v11 uint64 = 0
                        _ = v11
                    } else {
                        v12 := GenS1{F0: -2147483648, F1: -935}
                        _ = v12
                    }
                    _ = v10
                    _ = v8
                } else {
                    var v13 string = (v6 + ((v6 + v6) + v6))
                    v13 = v6
                    v14 := GenS1{F0: 0, F1: 42}
                    _ = v13
                    _ = v14
                }
                _ = fi7
            }
        case "\n\t":
            v6 = v6
        case "hello":
            v15 := GenS1{F0: -890, F1: 1}
            _ = v15
        default:
            if (true && (1 < 42)) {
                v16 := GenS2{F0: "foo bar", F1: 4294967295, F2: 1}
                if true {
                    v17 := GenS1{F0: 42, F1: 212}
                    var v18 int32 = 616
                    fmt.Println(v18)
                    _ = v17
                    _ = v18
                } else {
                    var v19 uint8 = 127
                    v19 += 0
                    _ = v19
                }
                v20 := GenS1{F0: -2147483648, F1: 9223372036854775807}
                _ = v16
                _ = v20
            }
        }
        _ = fi5
        _ = v6
    }
    var v21 uint32 = 291
    var v22 int64 = -1
    _ = v21
    _ = v22
    _ = v4
    return "hello"
}

func main() {
    var v1 uint32 = 46
    var v2 uint32 = 1
    v3 := []int{32767, 32767}
    if false {
        var v23 uint8 = 2
        _ = v23
    } else {
        var v24 int64 = 1
        v3 = append(v3, 9223372036854775807)
        _ = v24
    }
    v25 := GenS3{F0: -76}
    fmt.Println(v1, v1)
    v2 = 1
    v2 -= (^(v1))
    for fi26 := 0; fi26 < 20; fi26++ {
        if (!(false || (42 >= 321))) {
            for fi27 := 0; fi27 < 11; fi27++ {
                _ = len(v3)
                v2 = v2
                _ = fi27
            }
            v28 := GenS4{F0: 255, F1: fi26}
            _ = v28
        } else {
            if ((("" + ("\n\t" + "")) + (("foo bar" + ((("hello" + ("foo bar" + "a")) + "a") + "")) + "a")) < "") {
                _ = len(v3)
                switch fi26 {
                case 4294967295:
                    v2 += (v2 * v1)
                case -816:
                    for fi29 := 0; fi29 < 5; fi29++ {
                        var v30 float64 = 1.5
                        _ = fi29
                        _ = v30
                    }
                }
                switch "\n\t" {
                case "a":
                    if len(v3) > 0 {
                        v3[0] = -617
                    }
                }
            } else {
                v31 := GenS3{F0: 1}
                var v32 int = (fi26 / (((^(fi26))) | 1))
                _ = v31
                _ = v32
            }
        }
        _ = fi26
    }
    if false {
        if false {
            switch "" {
            case "\n\t":
                if true {
                    for fi33 := 0; fi33 < 13; fi33++ {
                        for fi34 := 0; fi34 < 15; fi34++ {
                            switch "foo bar" {
                            case "foo bar":
                                var v35 int = (fi33 + fi33)
                                _ = v35
                            case "":
                                var v36 int = (fi33 + -844)
                                _ = v36
                            case "a":
                                v37 := GenS1{F0: 0, F1: 1}
                                _ = v37
                            }
                            _ = fi34
                        }
                        _ = fi33
                    }
                    v38 := GenS5{F0: (255 >= 678)}
                    fmt.Println(v2, v1)
                    _ = v38
                }
            case "hello":
                _ = len(v3)
            }
            v1 = (v2 + v2)
        } else {
            switch 1 {
            case 32768:
                for fi39 := 0; fi39 < 14; fi39++ {
                    var v40 int = fi39
                    v1 -= v1
                    _ = fi39
                    _ = v40
                }
            case 1:
                v1 = v1
            default:
                var v41 int16 = -32768
                _ = v41
            }
            v1 = 366
        }
        v1 = (^(v1))
        var v42 int8 = -1
        _ = v42
    } else {
        fmt.Println(v2)
        v1 = (v1 & (^(v2)))
        var v43 int = 127
        _ = v43
    }
    var v44 uint = 256
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = v1
    var v45 uint16 = 42
    if true {
        fmt.Println(v2, v2, v1)
        switch "foo bar" {
        case "hello":
            for fi46 := 0; fi46 < 12; fi46++ {
                if len(v3) > 0 {
                    v3[0] = int(v44)
                }
                _ = fi46
            }
        case "":
            fmt.Println(v45, v1)
        }
    }
    for fi47 := 0; fi47 < 19; fi47++ {
        if false {
            v45 = v45
            var v48 float32 = 3.14159
            _ = v48
        } else {
            var v49 float64 = 3.14159
            _ = v49
        }
        v50 := GenS4{F0: 8, F1: (-(fi47))}
        _ = fi47
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v44
    _ = v45
}
