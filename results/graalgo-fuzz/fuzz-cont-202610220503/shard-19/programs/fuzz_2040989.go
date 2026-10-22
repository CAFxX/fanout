package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int; F1 uint32; F2 int64 }
type GenS5 struct { F0 bool; F1 uint; F2 int64 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 int64; F1 uint }
type GenS8 struct { F0 uint32; F1 int16; F2 int32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32; F1 string }

func main() {
    var v1 int32 = 580
    v2 := []int{}
    var v3 uint8 = 128
    var v4 uint8 = (^(v3))
    var v5 uint8 = (^(v3))
    if true {
        switch "hello" {
        case "a":
            v3 += (^(v3))
        }
    } else {
        for fi6 := 0; fi6 < 4; fi6++ {
            v3 -= (v5 >> ((760) & 63))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi6
        }
        if false {
            for fi7 := 0; fi7 < 9; fi7++ {
                v2 = append(v2, (fi7 ^ fi7))
                v4 = 255
                if false {
                    var v8 uint8 = v3
                    _ = v8
                } else {
                    v3 = (v5 ^ 1)
                }
                _ = fi7
            }
            v2 = append(v2, 0)
        } else {
            if ((^(v5)) > v5) {
                var v9 int32 = v1
                _ = v9
            } else {
                for fi10 := 0; fi10 < 11; fi10++ {
                    for fi11 := 0; fi11 < 1; fi11++ {
                        v12 := GenS1{F0: v3, F1: ("foo bar" + ("a" + "foo bar"))}
                        _ = fi11
                        _ = v12
                    }
                    if (2 == 4294967295) {
                        fmt.Println(v5, v1)
                        v5 = v5
                        v3 = (v5 >> ((754) & 63))
                    } else {
                        if false {
                            v13 := GenS2{F0: 967}
                            var v14 uint64 = 1
                            v2 = append(v2, fi10)
                            _ = v13
                            _ = v14
                        }
                    }
                    var v15 uint32 = 2147483647
                    _ = fi10
                    _ = v15
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            v16 := GenS3{F0: true, F1: (42 == 65535)}
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 9; fi18++ {
            if len(v2) > 0 {
                v2[0] = -215
            }
            _ = fi18
        }
        v19 := GenS4{F0: fi17, F1: 4294967295, F2: -9223372036854775808}
        v20 := GenS4{F0: (-(fi17)), F1: 186, F2: 9223372036854775807}
        _ = fi17
        _ = v19
        _ = v20
    }
    if (!true) {
        switch ("a" + "a") {
        case "foo bar":
            for fi21 := 0; fi21 < 10; fi21++ {
                v22 := GenS5{F0: (!false), F1: 1, F2: -9223372036854775808}
                _ = fi21
                _ = v22
            }
        case "a":
            if (127 >= 1) {
                var v23 float32 = (float32(3.14159) + -2.25)
                _ = v23
            } else {
                switch -1 {
                case -712:
                    v3 = (^(v5))
                }
                if false {
                    var v24 uint32 = 42
                    for fi25 := 0; fi25 < 15; fi25++ {
                        if (127 != 100) {
                            var v26 int32 = (v1 * v1)
                            v24 = v24
                            switch (fi25 | fi25) {
                            case -9223372036854775808:
                                _ = len(v2)
                            }
                            _ = v26
                        } else {
                            v27 := GenS4{F0: (fi25 + 886), F1: v24, F2: 0}
                            v5 -= (^(v5))
                            switch "" {
                            case "\n\t":
                                v3 -= (^(v5))
                            }
                            _ = v27
                        }
                        v1 = (v1 ^ (^(v1)))
                        _ = fi25
                    }
                    _ = v24
                }
            }
        default:
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 bool = false
                _ = fi28
                _ = v29
            }
        }
        var v30 int64 = -9223372036854775808
        var v31 uint16 = 952
        _ = v30
        _ = v31
    }
    switch -9223372036854775808 {
    case 636:
        var v32 int = -9223372036854775808
        _ = v32
    case -189:
        if false {
            switch ((("a" + "a") + "a") + "") {
            case "":
                fmt.Println(v5)
            case "a":
                var v33 uint64 = 18446744073709551615
                _ = v33
            case "foo bar":
                var v34 float64 = ((-2.25 / 1e10) + (0.0 + -2.25))
                _ = v34
            default:
                v3 = 1
            }
            v1 = -1
        }
    case 127:
        v1 = 134
    default:
        fmt.Println(v4)
    }
    v35 := GenS6{F0: "hello", F1: 521}
    var v36 uint16 = 65535
    v37 := GenS7{F0: 0, F1: 4294967295}
    v38 := GenS8{F0: 610, F1: 1000, F2: (v1 + v1)}
    v39 := GenS9{F0: -128}
    var v40 int32 = v1
    var v41 int64 = -9223372036854775808
    var v42 uint64 = 4294967295
    switch int(v40) {
    case -9223372036854775808:
        v43 := GenS10{F0: 4294967295, F1: "a"}
        _ = v43
    default:
        var v44 int32 = 1
        _ = v44
    }
    switch "" {
    case "foo bar":
        var v45 uint64 = 4294967295
        _ = v45
    case "hello":
        var v46 int64 = v41
        _ = v46
    default:
        fmt.Println(v3)
    }
    for fi47 := 0; fi47 < 17; fi47++ {
        fmt.Println(v36, v41)
        for fi48 := 0; fi48 < 15; fi48++ {
            v1 -= v40
            v41 -= (-(v41))
            if (0.0 <= 0.0) {
                v4 = 0
                _ = len(v2)
            }
            _ = fi48
        }
        _ = fi47
    }
    v5 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
