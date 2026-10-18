package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int8; F1 uint; F2 uint16 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 bool }
func genFn0() int64 {
    switch 930 {
    case 0:
        var v5 uint64 = 4294967295
        _ = v5
    case 2147483647:
        if true {
            v6 := []int{32768, 90, 0, 0}
            v7 := GenS1{F0: -98}
            var v8 uint8 = 1
            _ = v6
            _ = v7
            _ = v8
        }
    case 9223372036854775807:
        switch ("\n\t" + ("a" + ("hello" + ("a" + "\n\t")))) {
        case "":
            v9 := GenS2{F0: 127, F1: -1}
            _ = v9
        case "hello":
            switch ("foo bar" + ("" + (("hello" + "a") + ""))) {
            case "a":
                v10 := GenS1{F0: 2147483647}
                _ = v10
            case "\n\t":
                var v11 uint64 = 4294967295
                _ = v11
            case "hello":
                var v12 int = -671
                _ = v12
            default:
                if false {
                    var v13 int16 = 1
                    for fi14 := 0; fi14 < 8; fi14++ {
                        v15 := GenS3{F0: -128, F1: 0, F2: 42}
                        var v16 int8 = 100
                        v17 := GenS2{F0: (v16 / (((v16 / (((v16 & v16)) | 1))) | 1)), F1: -9223372036854775808}
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    _ = v13
                } else {
                    if false {
                        if true {
                            v18 := []int{}
                            var v19 bool = (9 <= -85)
                            var v20 uint32 = 323
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        var v21 float32 = (0.0 + (0.1 / (1.5 / (1.5 - 1e10))))
                        var v22 int32 = -839
                        _ = v21
                        _ = v22
                    }
                    var v23 bool = false
                    if v23 {
                        for fi24 := 0; fi24 < 1; fi24++ {
                            v25 := GenS1{F0: 2147483647}
                            _ = fi24
                            _ = v25
                        }
                    }
                    _ = v23
                }
            }
        }
    }
    v26 := GenS4{F0: 1000, F1: 0}
    var v27 int16 = 950
    v27 = v27
    _ = v26
    _ = v27
    return -937
}

func main() {
    var v1 uint16 = 1
    v2 := []int{-802, 128, 143}
    var v3 int16 = 790
    var v4 int = -30
    var v28 string = ""
    if (!(0 <= 1)) {
        var v29 uint16 = v1
        fmt.Println(v3, v4)
        _ = v29
    } else {
        if false {
            var v30 float32 = 0.0
            switch v4 {
            case 944:
                for fi31 := 0; fi31 < 8; fi31++ {
                    if len(v2) > 0 {
                        v2[0] = v4
                    }
                    fmt.Println(v4)
                    _ = fi31
                }
            case 32767:
                var v32 uint32 = 0
                _ = v32
            }
            switch v28 {
            case "a":
                switch "\n\t" {
                case "":
                    v28 = "a"
                case "a":
                    v33 := GenS5{F0: 0, F1: 2147483647}
                    _ = v33
                case "\n\t":
                    if (38 != 0) {
                        for fi34 := 0; fi34 < 17; fi34++ {
                            var v35 uint = 4294967295
                            if (genFn0() == 1) {
                                v4 = (v4 % (((^(v4))) | 1))
                            }
                            _ = fi34
                            _ = v35
                        }
                    }
                }
            }
            _ = v30
        }
        var v36 string = ("foo bar" + "")
        var v37 uint64 = 313
        _ = v36
        _ = v37
    }
    var v38 string = (((("" + v28) + (v28 + (v28 + v28))) + ((((("foo bar" + (v28 + (v28 + "a"))) + "\n\t") + (v28 + ("foo bar" + v28))) + (v28 + v28)) + (v28 + (("foo bar" + "") + v28)))) + "foo bar")
    var v39 uint16 = v1
    for fi40 := 0; fi40 < 2; fi40++ {
        for fi41 := 0; fi41 < 4; fi41++ {
            v38 = (v38 + (("" + "a") + v28))
            v28 = v38
            v42 := GenS5{F0: -9223372036854775808, F1: 464}
            var v43 string = ("foo bar" + "a")
            var v44 float32 = 0.1
            v1 += 974
            _ = fi41
            _ = v42
            _ = v43
            _ = v44
        }
        v2 = append(v2, 781)
        _ = fi40
    }
    if false {
        switch v38 {
        case "hello":
            for fi45 := 0; fi45 < 15; fi45++ {
                if len(v2) > 0 {
                    v2[0] = 32768
                }
                _ = fi45
            }
        case "":
            if false {
                var v46 int8 = 1
                _ = v46
            }
        case "a":
            v4 = 580
        default:
            v3 += (v3 & 42)
        }
        v4 = 128
    }
    var v47 int8 = -128
    v2 = append(v2, v4)
    fmt.Println(v38)
    var v48 float64 = 1e10
    v3 -= (v3 << ((42) & 63))
    var v49 int64 = genFn0()
    var v50 int32 = 141
    switch (v4 % (((v4 << ((255) & 63))) | 1)) {
    case 32768:
        if true {
            fmt.Println(v49, v4, v3)
            var v51 uint32 = 4294967295
            _ = v51
        } else {
            fmt.Println(v39, v38, v50)
        }
    case -38:
        v39 += 32767
    default:
        v4 = (v4 * v4)
    }
    fmt.Println(v50)
    var v52 string = "a"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        if ((v4 >> ((0) & 63)) < (v4 - int(v50))) {
            v53 := GenS6{F0: (659 <= 255)}
            _ = v53
        }
        var v54 bool = false
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v52
}
