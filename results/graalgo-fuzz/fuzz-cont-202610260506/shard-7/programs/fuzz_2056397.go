package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0() string {
    var v4 int = 25
    for fi5 := 0; fi5 < 9; fi5++ {
        v4 = 1
        _ = fi5
    }
    var v6 int = v4
    v7 := GenS1{F0: 255}
    var v8 int8 = -51
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    return "hello"
}
func genFn1() uint {
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := []int{fi9, 984, fi9}
        var v11 float32 = 0.1
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 string = ("" + "hello")
    _ = v12
    return 325
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 3.14159
    v3 := make(map[uint32]uint32)
    if ((1 != 65535) || false) {
        var v13 int64 = -650
        _ = v13
    } else {
        for fi14 := 0; fi14 < 8; fi14++ {
            if true {
                switch fi14 {
                case -1:
                    for fi15 := 0; fi15 < 14; fi15++ {
                        v2 = v2
                        v16 := GenS1{F0: 0}
                        var v17 uint64 = 316
                        fmt.Println(v17)
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case 0:
                    fmt.Println(v2)
                case 32767:
                    v18 := GenS2{F0: 42}
                    _ = v18
                }
            } else {
                v2 -= 0.1
                v2 -= float64(0.0)
                delete(v1, -1)
            }
            var v19 string = "\n\t"
            for fi20 := 0; fi20 < 20; fi20++ {
                if false {
                    if (9223372036854775807 == 1) {
                        switch fi20 {
                        case -889:
                            if (2 <= 724) {
                                v19 = ""
                            }
                        default:
                            v21 := GenS1{F0: 272}
                            _ = v21
                        }
                        fmt.Println(fi14, v2)
                        var v22 int64 = -9223372036854775808
                        _ = v22
                    } else {
                        fmt.Println(v2)
                    }
                    v19 = ((v19 + (((v19 + (v19 + ("foo bar" + ("a" + "\n\t")))) + v19) + v19)) + v19)
                    v19 = (v19 + ((("\n\t" + v19) + "\n\t") + (("\n\t" + ((v19 + v19) + (v19 + v19))) + v19)))
                } else {
                    switch fi14 {
                    case 453:
                        var v23 float32 = ((3.14159 / 1.5) - float32(1.5))
                        _ = v23
                    }
                }
                var v24 int32 = 0
                v2 += v2
                _ = fi20
                _ = v24
            }
            _ = fi14
            _ = v19
        }
    }
    fmt.Println(v2)
    if (!true) {
        fmt.Println(v2)
        if false {
            for fi25 := 0; fi25 < 6; fi25++ {
                fmt.Println(fi25, fi25)
                _ = fi25
            }
        }
        v26 := []int{-1, 32767, -9223372036854775808, 940}
        _ = v26
    }
    v27 := GenS2{F0: 858}
    v2 += (v2 / float64(3.14159))
    var v28 int64 = 425
    switch 0 {
    case 9223372036854775807:
        switch -56 {
        case 32767:
            var v29 uint = 42
            _ = v29
        case 42:
            if false {
                fmt.Println(v28, v2)
            }
        }
    default:
        fmt.Println(v28, v28)
    }
    v30 := []int{128, 2147483647}
    var v31 float64 = (((v2 * float64(((1e10 + 0.1) - (0.0 - -2.25)))) / 0.1) * v2)
    for fi32 := 0; fi32 < 16; fi32++ {
        for fi33 := 0; fi33 < 10; fi33++ {
            var v34 uint8 = 2
            switch ("hello" + ("\n\t" + "hello")) {
            case "foo bar":
                v34 = v34
            case "":
                var v35 int8 = 42
                _ = v35
            case "a":
                _ = len(v3)
            default:
                fmt.Println(v2, v2)
            }
            _ = fi33
            _ = v34
        }
        delete(v3, 0)
        switch ("" + "a") {
        case "foo bar":
            for fi36 := 0; fi36 < 8; fi36++ {
                v1[-364] = fi32
                v31 -= (v31 + (-2.25 / v2))
                _ = fi36
            }
        case "a":
            fmt.Println(fi32, v2)
        }
        _ = fi32
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        for fi38 := 0; fi38 < 9; fi38++ {
            v31 -= v31
            _ = fi38
        }
        v30 = append(v30, (fi37 + 128))
        for fi39 := 0; fi39 < 2; fi39++ {
            if len(v30) > 0 {
                v30[0] = -2147483648
            }
            var v40 string = (("" + ("foo bar" + "\n\t")) + "hello")
            v31 -= v31
            _ = fi39
            _ = v40
        }
        _ = fi37
    }
    v30 = append(v30, 920)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
}
