package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 string; F1 int32; F2 uint32 }
type GenS7 struct { F0 uint32; F1 int16; F2 int8 }
type GenS8 struct { F0 uint8 }
func genFn0() uint16 {
    v4 := GenS1{F0: 128}
    v5 := []int{}
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint64 = 751
        var v8 bool = ("" != "a")
        for fi9 := 0; fi9 < 12; fi9++ {
            switch (fi6 | 2147483647) {
            case -1:
                var v10 uint32 = 42
                _ = v10
            case -333:
                v11 := GenS2{F0: 42}
                _ = v11
            }
            v7 += (^(v7))
            _ = fi9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v12 int32 = -302
    var v13 uint16 = 32767
    _ = v12
    _ = v13
    _ = v4
    _ = v5
    return (v13 << ((4294967295) & 63))
}

func main() {
    v1 := []string{("hello" + ("" + (("foo bar" + ("" + "a")) + "hello"))), "", "hello"}
    var v2 bool = false
    var v3 uint8 = 0
    switch "" {
    case "hello":
        v3 = v3
    }
    if len(v1) > 0 {
        v1[0] = ""
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := GenS1{F0: (^(v3))}
            var v17 int32 = -670
            var v18 uint = 65535
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        var v19 int16 = -32768
        switch (fi14 ^ (fi14 | (fi14 >> ((uint(v3)) & 63)))) {
        case 736:
            switch ("hello" + "hello") {
            case "hello":
                v2 = false
            case "\n\t":
                if true {
                    fmt.Println(v3, fi14, v2)
                    v19 = (v19 - (v19 & v19))
                    v2 = v2
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch ("hello" + "\n\t") {
                    case "\n\t":
                        switch fi14 {
                        case 0:
                            fmt.Println(fi14, fi14, fi14)
                        }
                    case "hello":
                        if (32767 == 859) {
                            for fi20 := 0; fi20 < 15; fi20++ {
                                v21 := GenS3{F0: 18446744073709551615}
                                _ = fi20
                                _ = v21
                            }
                            if v2 {
                                v22 := GenS3{F0: 624}
                                _ = v22
                            } else {
                                v23 := GenS1{F0: v3}
                                var v24 int64 = 0
                                var v25 string = ("hello" + ("foo bar" + ("" + ("" + "hello"))))
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            v26 := GenS4{F0: 2147483648, F1: 42}
                            _ = v26
                        }
                    case "a":
                        var v27 int8 = int8(v19)
                        _ = v27
                    }
                    if true {
                        for fi28 := 0; fi28 < 2; fi28++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi28
                        }
                    }
                }
            default:
                for fi29 := 0; fi29 < 16; fi29++ {
                    v2 = v2
                    v30 := GenS5{F0: (^(v3)), F1: 417}
                    fmt.Println(v2)
                    _ = fi29
                    _ = v30
                }
            }
        }
        _ = fi14
        _ = v19
    }
    var v31 float64 = 0.0
    for fi32 := 0; fi32 < 17; fi32++ {
        var v33 uint = 255
        _ = fi32
        _ = v33
    }
    v34 := GenS6{F0: (((("hello" + "a") + "") + "") + "hello"), F1: -1, F2: 42}
    for fi35 := 0; fi35 < 8; fi35++ {
        var v36 string = "hello"
        v1 = append(v1, v36)
        _ = fi35
        _ = v36
    }
    v37 := GenS4{F0: 42, F1: 32767}
    v2 = (((false || (v2 || (134 > 1))) && (42 < -1)) && ((1.5 + float32(3.14159)) > 1.5))
    v31 = (v31 * v31)
    fmt.Println(v2, v31)
    v3 = v3
    for fi38 := 0; fi38 < 4; fi38++ {
        switch "\n\t" {
        case "hello":
            v39 := GenS2{F0: 1}
            _ = v39
        case "\n\t":
            if (2 >= 65535) {
                v2 = true
            } else {
                if v2 {
                    v3 -= v3
                    if (-1 != 42) {
                        var v40 int32 = -2147483648
                        _ = v40
                    }
                } else {
                    v3 -= (^(v3))
                    var v41 float64 = (v31 / v31)
                    _ = v41
                }
                v42 := GenS5{F0: v3, F1: 65535}
                _ = v42
            }
        }
        fmt.Println(v3, v31, v3)
        _ = fi38
    }
    v43 := GenS7{F0: 14, F1: 32767, F2: 12}
    fmt.Println(v2, v2, v2)
    if false {
        v44 := GenS6{F0: (("foo bar" + "foo bar") + ""), F1: 0, F2: 2147483648}
        v45 := GenS8{F0: 1}
        _ = v44
        _ = v45
    } else {
        v31 -= v31
        if (!v2) {
            for fi46 := 0; fi46 < 8; fi46++ {
                fmt.Println(v31)
                var v47 float32 = float32(v31)
                var v48 uint16 = 653
                _ = fi46
                _ = v47
                _ = v48
            }
            fmt.Println(v2, v31)
            v3 = (v3 >> ((351) & 63))
        } else {
            if (!(-768 == 49)) {
                var v49 uint32 = 2147483647
                v31 -= -2.25
                _ = v49
            }
            v50 := GenS5{F0: v3, F1: 255}
            var v51 float64 = v31
            _ = v50
            _ = v51
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v37
    _ = v43
}
