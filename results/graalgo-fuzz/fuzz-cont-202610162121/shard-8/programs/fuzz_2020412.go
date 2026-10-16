package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int8; F2 int32 }
type GenS5 struct { F0 int32 }
func genFn0(p0 bool, p1 uint8, p2 uint8) uint16 {
    p2 = (p2 << ((2) & 63))
    p1 = 127
    if p0 {
        var v4 uint64 = 1
        for fi5 := 0; fi5 < 14; fi5++ {
            var v6 float64 = (((3.14159 * 1e10) * -2.25) + (float64(0.1) - -2.25))
            fmt.Println(v6)
            _ = fi5
            _ = v6
        }
        _ = v4
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 float32 = (1e10 + 0.1)
    v2 := []string{"a", "\n\t", "", "hello"}
    var v3 int8 = 42
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(fi7, v3, fi7)
        _ = fi7
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -950 {
    case 42:
        var v8 uint = 375
        _ = v8
    case -2147483648:
        fmt.Println(v1, v1)
    }
    var v9 int8 = v3
    v3 = v3
    if ((0 != -94) || false) {
        v1 = v1
        for fi10 := 0; fi10 < 5; fi10++ {
            switch fi10 {
            case -2147483648:
                for fi11 := 0; fi11 < 2; fi11++ {
                    v12 := GenS1{F0: (false && ("" == ("a" + ("foo bar" + "a")))), F1: -1}
                    _ = len(v2)
                    _ = fi11
                    _ = v12
                }
            case 128:
                v2 = append(v2, "a")
            }
            switch "" {
            case "hello":
                v13 := GenS2{F0: 255, F1: "\n\t"}
                _ = v13
            case "foo bar":
                var v14 uint = 642
                _ = v14
            case "a":
                for fi15 := 0; fi15 < 7; fi15++ {
                    fmt.Println(fi10, fi15, fi15)
                    v3 += (v9 - (v9 * (v3 - v9)))
                    _ = fi15
                }
            default:
                switch ("\n\t" + ("" + ("a" + (((("" + (((("" + "") + "") + "") + ("foo bar" + ("" + "foo bar")))) + "") + "a") + "hello")))) {
                case "\n\t":
                    switch "foo bar" {
                    case "a":
                        var v16 int8 = -128
                        _ = v16
                    case "foo bar":
                        _ = len(v2)
                    case "\n\t":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        if true {
                            var v17 int64 = 0
                            _ = v17
                        }
                    }
                }
            }
            if (false || true) {
                fmt.Println(v1)
            } else {
                if false {
                    var v18 uint64 = 0
                    var v19 int32 = -1
                    _ = v18
                    _ = v19
                }
            }
            _ = fi10
        }
    }
    if false {
        v3 = (v3 - (-(v3)))
        switch "a" {
        case "\n\t":
            if (!(878 > -2147483648)) {
                switch 42 {
                case 32767:
                    v3 += (v3 | v3)
                case 42:
                    for fi20 := 0; fi20 < 5; fi20++ {
                        var v21 int = fi20
                        _ = fi20
                        _ = v21
                    }
                case 32768:
                    var v22 uint32 = 917
                    _ = v22
                default:
                    _ = len(v2)
                }
                v1 = ((v1 - v1) + 1.5)
            } else {
                fmt.Println(v3)
                var v23 uint8 = 127
                _ = v23
            }
        }
        if len(v2) > 0 {
            v2[0] = ((((("\n\t" + "") + "") + "a") + "\n\t") + ((("foo bar" + "") + "a") + "a"))
        }
    } else {
        v24 := GenS3{F0: "a"}
        for fi25 := 0; fi25 < 14; fi25++ {
            for fi26 := 0; fi26 < 16; fi26++ {
                if true {
                    v3 = (v3 >> ((1) & 63))
                    v27 := GenS2{F0: 18446744073709551615, F1: ("foo bar" + "hello")}
                    _ = v27
                }
                _ = fi26
            }
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 int8 = -1
                _ = fi28
                _ = v29
            }
            v30 := GenS4{F0: "", F1: (v9 - (v3 + 100)), F2: 42}
            _ = fi25
            _ = v30
        }
        _ = v24
    }
    v3 += (v9 ^ v9)
    var v31 uint8 = 42
    switch (("a" + "hello") + "") {
    case "foo bar":
        switch -513 {
        case 1:
            var v32 float64 = 1e10
            _ = v32
        case 128:
            v31 += v31
        case -765:
            v9 = -91
        }
    }
    var v33 uint64 = 1
    if (false && (2147483647 < 65535)) {
        v31 = 1
        if (!false) {
            _ = len(v2)
            for fi34 := 0; fi34 < 10; fi34++ {
                if true {
                    v31 = 255
                }
                v33 += 6
                _ = fi34
            }
            v33 -= (v33 & (v33 ^ v33))
        } else {
            v2 = append(v2, (("hello" + ((("hello" + (("a" + "hello") + "hello")) + "a") + "foo bar")) + "foo bar"))
        }
    } else {
        _ = len(v2)
        v33 = v33
    }
    if ((v33 & (v33 ^ (^(v33)))) >= (v33 / ((v33) | 1))) {
        for fi35 := 0; fi35 < 13; fi35++ {
            v31 -= (v31 << ((2) & 63))
            _ = fi35
        }
        switch "a" {
        case "\n\t":
            v2 = append(v2, ("\n\t" + ("foo bar" + "\n\t")))
        case "foo bar":
            v36 := GenS5{F0: -183}
            _ = v36
        }
    }
    var v37 float32 = ((0.1 - v1) + -2.25)
    for fi38 := 0; fi38 < 12; fi38++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi39 := 0; fi39 < 12; fi39++ {
            switch (fi38 % ((fi38) | 1)) {
            case 9223372036854775807:
                v1 = float32(1e10)
            }
            v40 := GenS3{F0: ("hello" + "a")}
            var v41 uint64 = (^(v33))
            if true {
                v2 = append(v2, ("\n\t" + ("" + "a")))
                var v42 int8 = v3
                _ = v42
            } else {
                if false {
                    fmt.Println(v33, v3, v3)
                    v43 := GenS4{F0: "", F1: v3, F2: 0}
                    v2 = append(v2, "foo bar")
                    _ = v43
                }
            }
            v44 := GenS1{F0: true, F1: -156}
            _ = len(v2)
            var v45 int32 = 967
            if (42 > uint16(v41)) {
                v9 -= v3
                if ((!true) && (true && (2 >= 42))) {
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                    _ = len(v2)
                    if (42 != (v45 ^ int32(fi38))) {
                        if false {
                            if ((fi39 / ((fi38) | 1)) < fi39) {
                                v45 -= v45
                            } else {
                                switch 42 {
                                case 905:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    var v46 float64 = float64(v37)
                                    _ = v46
                                }
                                if true {
                                    v47 := GenS4{F0: "\n\t", F1: (v9 >> ((1) & 63)), F2: -677}
                                    _ = v47
                                }
                                if ((-1 < -909) || false) {
                                    switch (fi38 & int(v31)) {
                                    case 42:
                                        var v48 uint16 = 255
                                        _ = v48
                                    }
                                }
                            }
                            v9 += (v9 >> ((0) & 63))
                        } else {
                            v49 := GenS5{F0: -2147483648}
                            var v50 int32 = v45
                            _ = v49
                            _ = v50
                        }
                    }
                } else {
                    if ((fi39 > (fi39 ^ (fi38 >> ((1) & 63)))) || false) {
                        fmt.Println(v33, v33, v37)
                        var v51 bool = false
                        _ = v51
                    }
                }
                var v52 uint32 = 0
                var v53 uint64 = 390
                _ = v52
                _ = v53
            }
            var v54 uint64 = v41
            fmt.Println(v1, v31)
            fmt.Println(v41)
            _ = fi39
            _ = v40
            _ = v41
            _ = v44
            _ = v45
            _ = v54
        }
        var v55 int = fi38
        _ = fi38
        _ = v55
    }
    v31 += (v31 >> ((65535) & 63))
    for fi56 := 0; fi56 < 5; fi56++ {
        var v57 uint32 = 80
        if (!("hello" != "hello")) {
            v31 -= v31
            fmt.Println(v31, v1, v9)
            if (1 != 46) {
                switch (fi56 % ((-745) | 1)) {
                case 4294967295:
                    v58 := GenS1{F0: (uint16(v33) >= 780), F1: -2147483648}
                    _ = v58
                case 42:
                    v59 := GenS2{F0: 1, F1: "\n\t"}
                    _ = v59
                default:
                    fmt.Println(v33)
                }
                v3 = (v3 * (v3 ^ (v9 >> ((721) & 63))))
            } else {
                v60 := GenS3{F0: ("foo bar" + ("a" + ""))}
                v9 += v9
                if true {
                    v2 = append(v2, "hello")
                }
                _ = v60
            }
        }
        _ = fi56
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v37
    _ = v9
}
