package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 int64; F1 int32; F2 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 bool = false
    var v4 uint32 = 4294967295
    var v5 int64 = -989
    if v3 {
        for fi6 := 0; fi6 < 14; fi6++ {
            switch "hello" {
            case "hello":
                v7 := []int{(fi6 % ((32767) | 1)), fi6}
                _ = v7
            case "\n\t":
                for fi8 := 0; fi8 < 12; fi8++ {
                    fmt.Println(v4, v4, fi6)
                    var v9 float32 = float32(3.14159)
                    v10 := GenS1{F0: "hello", F1: -233, F2: 408}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            default:
                var v11 uint8 = 128
                _ = v11
            }
            v4 = 622
            _ = fi6
        }
        switch "hello" {
        case "hello":
            v3 = true
        case "":
            v3 = (v3 || v3)
        case "a":
            v5 = (v5 / ((v5) | 1))
        default:
            var v12 float64 = -2.25
            _ = v12
        }
        var v13 int32 = 2147483647
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            if (((3.14159 - (1.5 / 1.0)) + 0.1) != (-2.25 * -2.25)) {
                v4 = (^(v4))
            } else {
                v3 = (3.14159 < -2.25)
                var v16 int8 = 0
                v5 += 794
                _ = v16
            }
            v17 := v2[(fi15 << ((4294967295) & 63))]
            fmt.Println(v5, v3, v5)
            _ = fi15
            _ = v17
        }
        fmt.Println(v3, v4)
        _ = fi14
    }
    if v3 {
        v4 += (v4 / ((2147483648) | 1))
        var v18 float32 = 1e10
        _ = v18
    } else {
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 bool = ((0 <= 1) || ((int(v4) < (-(fi19))) || false))
            var v21 uint8 = 255
            _ = fi19
            _ = v20
            _ = v21
        }
        var v22 uint8 = 249
        fmt.Println(v5)
        _ = v22
    }
    if ((255 > 452) || (true && (v3 || v3))) {
        var v23 int32 = 2147483647
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        switch "" {
        case "a":
            for fi25 := 0; fi25 < 5; fi25++ {
                v3 = (("foo bar" + ("" + "\n\t")) < ("\n\t" + ("" + "a")))
                _ = fi25
            }
        }
        _ = fi24
    }
    if v3 {
        fmt.Println(v5)
        var v26 uint64 = 243
        _ = v26
    }
    switch ("" + "hello") {
    case "":
        var v27 uint8 = 128
        _ = v27
    default:
        v5 -= (v5 * v5)
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        switch ("\n\t" + ("" + "foo bar")) {
        case "a":
            switch "hello" {
            case "foo bar":
                v1[v4] = 1
            case "":
                for fi28 := 0; fi28 < 4; fi28++ {
                    v5 = 42
                    var v29 uint = 327
                    _ = fi28
                    _ = v29
                }
            }
        case "hello":
            v30 := GenS2{F0: 128, F1: 32767}
            _ = v30
        }
    case "foo bar":
        if v3 {
            v3 = (0.1 > 0.1)
            v3 = false
            if v3 {
                switch -9223372036854775808 {
                case 474:
                    _ = len(v2)
                default:
                    var v31 uint16 = 255
                    _ = v31
                }
                v32 := GenS2{F0: 42, F1: 1}
                _ = v32
            } else {
                fmt.Println(v5)
            }
        } else {
            var v33 bool = true
            if false {
                fmt.Println(v5)
                var v34 uint16 = 0
                _ = v34
            } else {
                for fi35 := 0; fi35 < 13; fi35++ {
                    fmt.Println(v5, fi35)
                    _ = fi35
                }
                v33 = false
            }
            v5 = v5
            _ = v33
        }
    }
    if (0 >= 4294967295) {
        if true {
            switch "" {
            case "":
                _ = len(v2)
            case "foo bar":
                var v36 string = "foo bar"
                _ = v36
            case "hello":
                v37 := GenS3{F0: 0, F1: 76, F2: (("hello" + "foo bar") + ((("foo bar" + "hello") + (("foo bar" + "") + ("a" + ""))) + ""))}
                _ = v37
            }
            switch "foo bar" {
            case "":
                for fi38 := 0; fi38 < 11; fi38++ {
                    v3 = ("a" != "hello")
                    var v39 uint32 = (v4 - 2147483648)
                    if false {
                        if false {
                            v40 := GenS4{F0: 65535}
                            v4 += 0
                            v41 := GenS2{F0: 128, F1: 582}
                            _ = v40
                            _ = v41
                        } else {
                            var v42 uint = 42
                            _ = v42
                        }
                        _ = len(v2)
                        v43 := []int{(fi38 >> ((0) & 63)), 128, 2147483647}
                        _ = v43
                    }
                    _ = fi38
                    _ = v39
                }
            case "a":
                if v3 {
                    fmt.Println(v3)
                } else {
                    v44 := GenS5{F0: v3, F1: -2147483648}
                    var v45 bool = v3
                    _ = v44
                    _ = v45
                }
            default:
                for fi46 := 0; fi46 < 14; fi46++ {
                    v5 = v5
                    switch "\n\t" {
                    case "":
                        v4 -= (v4 + v4)
                    }
                    _ = fi46
                }
            }
        } else {
            var v47 bool = true
            v3 = (0 == 1000)
            _ = v47
        }
        var v48 uint8 = 215
        _ = v48
    } else {
        v49 := []int{127}
        if v3 {
            var v50 int8 = int8(v4)
            _ = v50
        }
        var v51 float64 = -2.25
        _ = v49
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
