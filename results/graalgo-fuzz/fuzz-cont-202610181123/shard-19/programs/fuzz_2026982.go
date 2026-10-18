package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 2
    v3 := make(map[int]bool)
    for fi4 := 0; fi4 < 3; fi4++ {
        v2 -= (^(v2))
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 bool = true
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v2 = v2
    v7 := GenS1{F0: 2}
    if (!true) {
        delete(v3, -1)
    } else {
        v8 := []int{42, 2147483647, -696}
        _ = v8
    }
    fmt.Println(v2)
    v9 := []int{692, -1, 807}
    if false {
        v2 = v2
    } else {
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 float32 = 0.1
            switch "a" {
            case "a":
                var v12 int32 = 1
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        switch "a" {
        case "hello":
            if false {
                v2 += v2
                var v13 int32 = 1
                var v14 uint = 4294967295
                _ = v13
                _ = v14
            } else {
                if (!true) {
                    switch "\n\t" {
                    case "foo bar":
                        var v15 uint16 = 2
                        _ = v15
                    case "\n\t":
                        v3[576] = (false || true)
                    }
                    var v16 uint = 0
                    if false {
                        fmt.Println(v16, v16)
                    }
                    _ = v16
                } else {
                    v17 := v1[""]
                    _ = v17
                }
            }
        case "foo bar":
            v18 := GenS1{F0: 32767}
            _ = v18
        default:
            var v19 int8 = 30
            _ = v19
        }
    }
    if (!false) {
        for fi20 := 0; fi20 < 12; fi20++ {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            v2 += (v2 * v2)
            if (65535 != 2) {
                for fi21 := 0; fi21 < 18; fi21++ {
                    var v22 float32 = (1e10 / ((-2.25 * 1e10) * -2.25))
                    _ = len(v3)
                    var v23 uint16 = 32767
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                v2 -= (v2 / (((^(v2))) | 1))
            }
            _ = fi20
        }
        var v24 bool = true
        var v25 int64 = 176
        _ = v24
        _ = v25
    } else {
        v2 -= v2
        var v26 uint8 = v2
        _ = v26
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        fmt.Println(fi27)
        if (((-1 != -9223372036854775808) && (!(!true))) || (!(110 >= 2))) {
            var v28 int = (-(fi27))
            if true {
                v9 = append(v9, (fi27 % (((v28 | (fi27 >> ((0) & 63)))) | 1)))
                if ((float64(1e10) < 0.1) && false) {
                    if len(v9) > 0 {
                        v9[0] = v28
                    }
                } else {
                    v2 += (v2 * v2)
                    fmt.Println(v28, v28)
                }
            } else {
                fmt.Println(fi27, v2)
                var v29 string = (("" + "\n\t") + "hello")
                _ = v29
            }
            _ = v28
        } else {
            var v30 int16 = 140
            if false {
                var v31 bool = (1 > 65535)
                var v32 int32 = -1
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 11; fi33++ {
                v2 = v2
                v30 += -178
                _ = fi33
            }
            _ = v30
        }
        _ = fi27
    }
    if false {
        var v34 int64 = 1
        switch -543 {
        case 32767:
            if false {
                switch ("foo bar" + ((("foo bar" + (("hello" + ("a" + "a")) + "a")) + "") + ("\n\t" + "hello"))) {
                case "":
                    v2 -= 1
                default:
                    v2 = v2
                }
                var v35 string = "foo bar"
                _ = v35
            } else {
                v36 := GenS1{F0: 622}
                _ = v36
            }
        case 32768:
            if (1 > 100) {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                fmt.Println(v34, v2)
                if true {
                    _ = len(v3)
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v37 string = "\n\t"
                    _ = v37
                }
            } else {
                v34 = (v34 * 832)
                fmt.Println(v2, v34)
            }
        }
        _ = v34
    } else {
        var v38 uint = 2
        v39 := GenS2{F0: int32(v2), F1: ""}
        var v40 uint8 = (v2 / ((uint8(v38)) | 1))
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
