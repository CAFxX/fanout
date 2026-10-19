package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint64; F1 int; F2 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[string]string)
    var v2 string = ((("hello" + "a") + "") + "")
    v3 := []int{}
    var v4 int64 = -9223372036854775808
    if (255 >= 1) {
        for fi5 := 0; fi5 < 9; fi5++ {
            _ = len(v1)
            _ = fi5
        }
        v2 = ((v2 + v2) + (v2 + "a"))
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 string = (((v2 + v2) + v2) + v2)
            var v8 int64 = (v4 * -1)
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    v2 = "\n\t"
    for fi9 := 0; fi9 < 14; fi9++ {
        v2 = v2
        switch 1 {
        case 2147483647:
            switch "hello" {
            case "\n\t":
                if false {
                    v10 := GenS1{F0: 1}
                    var v11 bool = (false || (!(42 >= 630)))
                    _ = v10
                    _ = v11
                }
            }
        case -405:
            v2 = "foo bar"
        default:
            v4 -= v4
        }
        var v12 int32 = -2147483648
        _ = fi9
        _ = v12
    }
    v4 = v4
    var v13 int8 = 100
    v1["hello"] = v2
    for fi14 := 0; fi14 < 1; fi14++ {
        v13 = v13
        v13 -= v13
        _ = fi14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := GenS1{F0: 42}
        switch 32767 {
        case 32768:
            var v17 int = fi15
            _ = v17
        default:
            for fi18 := 0; fi18 < 14; fi18++ {
                var v19 float32 = 0.0
                switch int(v13) {
                case 128:
                    if false {
                        var v20 int = fi15
                        _ = v20
                    } else {
                        v21 := GenS2{F0: (v4 << ((255) & 63)), F1: 0, F2: 1}
                        v2 = ""
                        v2 = "foo bar"
                        _ = len(v1)
                        _ = v21
                    }
                case -365:
                    v22 := GenS3{F0: 56, F1: fi15, F2: (v13 > (v13 & v13))}
                    _ = v22
                case 42:
                    v23 := GenS2{F0: (^(v4)), F1: 42, F2: 1}
                    _ = v23
                }
                v24 := GenS4{F0: fi15}
                _ = fi18
                _ = v19
                _ = v24
            }
        }
        _ = fi15
        _ = v16
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 uint8 = 42
        for fi27 := 0; fi27 < 14; fi27++ {
            v26 -= v26
            if false {
                var v28 int = -1
                v28 += (v28 << ((18446744073709551615) & 63))
                _ = v28
            } else {
                var v29 int16 = 0
                _ = v29
            }
            _ = fi27
        }
        var v30 float32 = (float32((-2.25 - (-2.25 + 0.1))) * (0.0 - (0.1 - 1.5)))
        _ = fi25
        _ = v26
        _ = v30
    }
    fmt.Println(v4, v13)
    v31 := GenS5{F0: 1, F1: 4294967295}
    switch v2 {
    case "foo bar":
        if (true && false) {
            switch v2 {
            case "foo bar":
                fmt.Println(v2, v4)
            case "hello":
                for fi32 := 0; fi32 < 16; fi32++ {
                    switch (v2 + v2) {
                    case "":
                        delete(v1, v2)
                    }
                    v33 := GenS6{F0: -1, F1: v2}
                    _ = fi32
                    _ = v33
                }
            default:
                fmt.Println(v13)
            }
            v34 := v1[(v2 + v2)]
            _ = len(v1)
            _ = v34
        }
    case "":
        v13 = v13
    case "\n\t":
        var v35 uint = 927
        _ = v35
    }
    fmt.Println(v2, v13)
    var v36 int8 = v13
    _ = len(v1)
    if ((true || true) || (4294967295 == 18446744073709551615)) {
        if (((true && false) && (((((-2.25 / 1.5) / 1.0) - 1e10) > float32(0.0)) && true)) && false) {
            var v37 float64 = float64(((-2.25 - 0.0) + 0.1))
            if false {
                v37 += 1e10
            } else {
                var v38 float32 = 0.1
                v2 = v2
                _ = v38
            }
            _ = v37
        }
        v13 = 0
        v39 := v1[((v2 + v2) + "")]
        _ = v39
    }
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v4
}
