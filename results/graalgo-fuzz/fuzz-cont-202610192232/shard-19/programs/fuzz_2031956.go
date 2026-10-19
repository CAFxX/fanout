package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint32; F2 bool }
type GenS4 struct { F0 int16; F1 uint8 }
func genFn0() int32 {
    v4 := []int{-68, -929}
    for fi5 := 0; fi5 < 4; fi5++ {
        switch (("a" + "hello") + "a") {
        case "a":
            _ = len(v4)
        case "\n\t":
            var v6 int64 = 1
            _ = v6
        case "":
            v4 = append(v4, 32767)
        }
        _ = fi5
    }
    switch "\n\t" {
    case "hello":
        switch 853 {
        case 32767:
            if true {
                v7 := GenS1{F0: 32767, F1: 4294967295}
                var v8 int64 = 9223372036854775807
                for fi9 := 0; fi9 < 18; fi9++ {
                    fmt.Println(v8, v8)
                    _ = fi9
                }
                _ = v7
                _ = v8
            }
        case -1:
            if false {
                for fi10 := 0; fi10 < 3; fi10++ {
                    var v11 uint32 = 2147483648
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v11)
                    case "foo bar":
                        fmt.Println(fi10, v11)
                    }
                    _ = fi10
                    _ = v11
                }
            }
        default:
            if (21 != 128) {
                if (-128 > 127) {
                    v12 := GenS1{F0: 391, F1: 2147483647}
                    v13 := GenS2{F0: 2147483648}
                    if true {
                        if (!true) {
                            v14 := GenS1{F0: 0, F1: 787}
                            _ = v14
                        } else {
                            if len(v4) > 0 {
                                v4[0] = -120
                            }
                        }
                    }
                    _ = v12
                    _ = v13
                } else {
                    var v15 bool = (!(42 == 9223372036854775807))
                    _ = v15
                }
                switch "foo bar" {
                case "\n\t":
                    v4 = append(v4, -196)
                }
            } else {
                v16 := GenS3{F0: 0, F1: 2, F2: true}
                var v17 int16 = 1000
                _ = len(v4)
                _ = v16
                _ = v17
            }
        }
    case "a":
        var v18 uint8 = 255
        _ = v18
    case "foo bar":
        if ((42 == 83) && ((!true) && (!true))) {
            var v19 int64 = 0
            var v20 int = -1
            switch (v20 >> ((45) & 63)) {
            case 1:
                v19 += 0
            }
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 16; fi21++ {
                var v22 string = "a"
                _ = fi21
                _ = v22
            }
        }
    default:
        for fi23 := 0; fi23 < 1; fi23++ {
            for fi24 := 0; fi24 < 9; fi24++ {
                fmt.Println(fi24, fi24)
                var v25 string = ("a" + ("hello" + ""))
                v25 = ((((("hello" + "hello") + v25) + (v25 + "a")) + (((v25 + "foo bar") + (v25 + "a")) + "a")) + v25)
                _ = fi24
                _ = v25
            }
            for fi26 := 0; fi26 < 17; fi26++ {
                var v27 uint8 = 2
                _ = fi26
                _ = v27
            }
            var v28 string = ("\n\t" + (("hello" + "hello") + ""))
            _ = fi23
            _ = v28
        }
    }
    var v29 int32 = 42
    _ = v29
    _ = v4
    return 0
}

func main() {
    var v1 int32 = -1
    var v2 float32 = 1e10
    v3 := []int{0, -662, 4294967295, 229}
    v30 := GenS4{F0: -1, F1: 2}
    switch ("\n\t" + ("a" + (("hello" + "a") + "a"))) {
    case "\n\t":
        _ = len(v3)
    }
    var v31 uint64 = 4294967295
    v1 = v1
    if (v31 != (v31 ^ v31)) {
        switch 9223372036854775807 {
        case 4294967295:
            for fi32 := 0; fi32 < 4; fi32++ {
                for fi33 := 0; fi33 < 19; fi33++ {
                    v1 -= (v1 / (((v1 % ((-781) | 1))) | 1))
                    var v34 bool = (127 != 0)
                    _ = fi33
                    _ = v34
                }
                v2 = 0.1
                for fi35 := 0; fi35 < 19; fi35++ {
                    v1 = v1
                    if true {
                        var v36 uint8 = 42
                        _ = v36
                    }
                    fmt.Println(v1)
                    _ = fi35
                }
                _ = fi32
            }
        default:
            fmt.Println(v31)
        }
        var v37 bool = false
        for fi38 := 0; fi38 < 4; fi38++ {
            v3 = append(v3, -189)
            for fi39 := 0; fi39 < 3; fi39++ {
                var v40 string = "\n\t"
                v41 := GenS2{F0: 2147483648}
                switch (v40 + ("" + v40)) {
                case "\n\t":
                    var v42 int8 = -1
                    _ = v42
                }
                _ = fi39
                _ = v40
                _ = v41
            }
            v43 := GenS1{F0: uint16(v31), F1: 4294967295}
            _ = fi38
            _ = v43
        }
        _ = v37
    }
    for fi44 := 0; fi44 < 7; fi44++ {
        v45 := GenS2{F0: 749}
        var v46 int = fi44
        var v47 string = "foo bar"
        _ = fi44
        _ = v45
        _ = v46
        _ = v47
    }
    fmt.Println(v1, v2, v1)
    _ = len(v3)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
}
