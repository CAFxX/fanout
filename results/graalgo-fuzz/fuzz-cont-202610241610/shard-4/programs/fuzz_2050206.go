package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 int; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 int64; F1 int8; F2 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 0.0
    v3 := []int{1}
    v2 = v2
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 15; fi5++ {
            switch 4294967295 {
            case 9223372036854775807:
                v2 += (0.1 - v2)
            case 278:
                fmt.Println(v2)
            case -1:
                switch "hello" {
                case "\n\t":
                    if true {
                        v6 := GenS1{F0: false}
                        var v7 uint8 = 42
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch ("foo bar" + "") {
                        case "hello":
                            switch fi5 {
                            case 975:
                                v7 = 0
                            default:
                                var v8 int8 = -25
                                _ = v8
                            }
                        case "\n\t":
                            v2 -= v2
                            if true {
                                _ = len(v1)
                                var v9 uint32 = 42
                                v2 = float64(1e10)
                                _ = v9
                            }
                        case "":
                            var v10 int8 = -128
                            _ = v10
                        default:
                            v11 := GenS2{F0: 9223372036854775807}
                            _ = v11
                        }
                        _ = v6
                        _ = v7
                    } else {
                        v12 := GenS3{F0: false, F1: fi5, F2: 2147483647}
                        v13 := v1[""]
                        var v14 int32 = 2147483647
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                }
            }
            _ = len(v3)
            var v15 float32 = 0.0
            v15 = 1.5
            _ = fi5
            _ = v15
        }
        _ = len(v1)
        _ = fi4
    }
    if true {
        switch ("hello" + ("hello" + "a")) {
        case "":
            var v16 uint32 = 42
            _ = v16
        }
        if (2 > 2) {
            if false {
                var v17 int8 = 89
                _ = v17
            }
            if (-1 < -123) {
                if (true || true) {
                    var v18 int8 = 1
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 9; fi19++ {
                fmt.Println(v2, fi19)
                _ = fi19
            }
        }
    }
    v20 := GenS4{F0: true, F1: 215}
    if len(v3) > 0 {
        v3[0] = 0
    }
    _ = len(v3)
    delete(v1, (("a" + "a") + "foo bar"))
    delete(v1, "")
    if ((1.5 - 3.14159) == (1.5 * ((1.5 - -2.25) + 1.5))) {
        if ((!false) || (true && (-268 <= 1))) {
            v2 += 3.14159
            var v21 bool = true
            var v22 float64 = float64(float32(v2))
            _ = v21
            _ = v22
        } else {
            switch ("a" + "a") {
            case "\n\t":
                var v23 string = ("" + ("\n\t" + "a"))
                _ = v23
            case "a":
                if (!(173 >= 255)) {
                    var v24 int = -262
                    if true {
                        v24 = (v24 + v24)
                        v2 -= (v2 * (v2 + -2.25))
                    }
                    _ = v24
                }
            default:
                v2 += 1e10
            }
            v25 := GenS5{F0: 42, F1: 4294967295}
            for fi26 := 0; fi26 < 12; fi26++ {
                v2 = (v2 / 1.5)
                _ = fi26
            }
            _ = v25
        }
        if len(v3) > 0 {
            v3[0] = 32767
        }
    }
    v2 += v2
    for fi27 := 0; fi27 < 12; fi27++ {
        var v28 int8 = 17
        if (false || true) {
            v28 = (v28 >> ((2) & 63))
            if false {
                if (!(true || true)) {
                    switch ("foo bar" + (("a" + "foo bar") + "hello")) {
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                    case "hello":
                        v29 := v1["hello"]
                        _ = v29
                    default:
                        fmt.Println(fi27, v2, v28)
                    }
                    v30 := GenS2{F0: 0}
                    _ = v30
                }
                v31 := GenS4{F0: (584 < 4294967295), F1: 0}
                for fi32 := 0; fi32 < 19; fi32++ {
                    fmt.Println(fi27)
                    _ = fi32
                }
                _ = v31
            }
            var v33 int64 = -1
            _ = v33
        } else {
            delete(v1, (("a" + "a") + ""))
        }
        var v34 string = (("hello" + ("foo bar" + "")) + (("\n\t" + (("\n\t" + "") + "\n\t")) + "\n\t"))
        _ = fi27
        _ = v28
        _ = v34
    }
    var v35 string = "a"
    var v36 int8 = 100
    if true {
        var v37 int64 = -1
        var v38 int32 = -1
        _ = v37
        _ = v38
    } else {
        if len(v3) > 0 {
            v3[0] = 32767
        }
    }
    fmt.Println(v35, v2, v35)
    v36 -= -1
    switch v35 {
    case "a":
        switch ("a" + "hello") {
        case "hello":
            v35 = "hello"
        case "\n\t":
            v36 = (v36 ^ (-(v36)))
        case "":
            fmt.Println(v36)
        default:
            fmt.Println(v36, v35)
        }
    }
    v39 := GenS6{F0: 9223372036854775807, F1: (v36 & v36), F2: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v39
}
