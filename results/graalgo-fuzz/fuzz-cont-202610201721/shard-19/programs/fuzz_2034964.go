package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 string; F2 uint16 }
type GenS3 struct { F0 int16; F1 int8; F2 bool }
type GenS4 struct { F0 uint; F1 string; F2 int8 }
func genFn0(p0 uint64, p1 int32) int {
    for fi4 := 0; fi4 < 20; fi4++ {
        fmt.Println(p1)
        _ = fi4
    }
    fmt.Println(p0, p1)
    switch (("a" + "a") + "a") {
    case "":
        if (255 != 1) {
            fmt.Println(p1, p0)
            fmt.Println(p0)
        } else {
            p0 = (p0 | 18446744073709551615)
            fmt.Println(p0, p1)
        }
    case "\n\t":
        var v5 int16 = 1
        _ = v5
    default:
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    return 127
}
func genFn1(p0 int, p1 int64) bool {
    if false {
        if (100 > -88) {
            fmt.Println(p1, p0)
        }
        var v6 int64 = p1
        _ = v6
    }
    var v7 int = p0
    v7 += -330
    if (("" + "\n\t") >= ("a" + ("hello" + ("" + "")))) {
        v8 := []int{v7, p0, (p0 / ((genFn0(0, 0)) | 1))}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 bool = (true || (65535 < 4294967295))
            _ = fi9
            _ = v10
        }
        v7 = v7
    }
    _ = p0
    _ = p1
    _ = v7
    return true
}

func main() {
    var v1 uint32 = 0
    var v2 uint8 = 0
    v3 := []string{}
    if len(v3) > 0 {
        v3[0] = "a"
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int = (fi11 - (fi11 << ((1) & 63)))
        _ = fi11
        _ = v12
    }
    v1 += v1
    switch "" {
    case "\n\t":
        if true {
            if (((((679 != v1) || (v1 > 228)) || true) || true) && true) {
                for fi13 := 0; fi13 < 18; fi13++ {
                    v14 := GenS1{F0: 255}
                    var v15 uint16 = 1
                    switch (fi13 / ((fi13) | 1)) {
                    case -940:
                        _ = len(v3)
                    case 9223372036854775807:
                        v16 := GenS2{F0: genFn0(111, 42), F1: "foo bar", F2: (v15 >> ((1) & 63))}
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                fmt.Println(v1)
                v1 = 65535
            } else {
                v17 := GenS1{F0: 65535}
                _ = v17
            }
            fmt.Println(v2, v2)
            if (uint(v2) >= 2) {
                switch "\n\t" {
                case "hello":
                    for fi18 := 0; fi18 < 9; fi18++ {
                        v19 := GenS3{F0: -558, F1: -128, F2: (true && true)}
                        _ = fi18
                        _ = v19
                    }
                default:
                    switch -9223372036854775808 {
                    case 32768:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -9223372036854775808:
                        if ((0.1 - -2.25) > 0.0) {
                            v1 = (v1 % ((v1) | 1))
                            _ = len(v3)
                        } else {
                            v2 = v2
                            v1 -= 197
                            for fi20 := 0; fi20 < 3; fi20++ {
                                _ = len(v3)
                                var v21 float32 = 3.14159
                                _ = fi20
                                _ = v21
                            }
                        }
                    }
                }
                switch 2147483647 {
                case -9223372036854775808:
                    var v22 int = 127
                    _ = v22
                case 9223372036854775807:
                    v23 := GenS4{F0: 18446744073709551615, F1: ("foo bar" + "a"), F2: -88}
                    _ = v23
                }
                if true {
                    v2 = 127
                }
            } else {
                for fi24 := 0; fi24 < 17; fi24++ {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        v2 += v2
                        fmt.Println(fi24, v2)
                        fmt.Println(fi24, fi24, v1)
                        _ = fi25
                    }
                    v3 = append(v3, ("" + ("hello" + ("hello" + "hello"))))
                    _ = fi24
                }
                for fi26 := 0; fi26 < 13; fi26++ {
                    v2 = v2
                    _ = fi26
                }
            }
        } else {
            v2 = v2
            if false {
                v1 -= (v1 - v1)
                v2 = (v2 * (^(v2)))
                v2 = 42
            } else {
                v2 += v2
                for fi27 := 0; fi27 < 13; fi27++ {
                    v1 += 42
                    _ = fi27
                }
            }
        }
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        var v29 int32 = 0
        v1 += 94
        _ = fi28
        _ = v29
    }
    var v30 uint16 = uint16(v2)
    switch 9223372036854775807 {
    case 9223372036854775807:
        v1 = 282
    case 4294967295:
        var v31 int16 = -315
        _ = v31
    }
    _ = len(v3)
    v32 := GenS2{F0: 705, F1: ("hello" + (("" + ("hello" + "foo bar")) + (("foo bar" + ("" + "\n\t")) + (("a" + ("\n\t" + "")) + "\n\t")))), F2: (^(v30))}
    if (-2.25 == (0.1 + ((3.14159 / 1.0) / 3.14159))) {
        var v33 string = ""
        v30 = v30
        _ = v33
    }
    switch 32768 {
    case -2147483648:
        switch genFn0(699, int32(v2)) {
        case -2147483648:
            _ = len(v3)
        case 32768:
            v2 -= v2
        case 128:
            switch -11 {
            case 127:
                var v34 int = 32768
                _ = v34
            }
        default:
            for fi35 := 0; fi35 < 5; fi35++ {
                if false {
                    v1 = uint32(v2)
                } else {
                    v30 = uint16(v1)
                }
                v30 -= v30
                for fi36 := 0; fi36 < 17; fi36++ {
                    fmt.Println(v2)
                    _ = fi36
                }
                _ = fi35
            }
        }
    case 4294967295:
        _ = len(v3)
    }
    switch "a" {
    case "\n\t":
        v37 := GenS3{F0: 640, F1: -1, F2: (!false)}
        _ = v37
    case "hello":
        for fi38 := 0; fi38 < 2; fi38++ {
            fmt.Println(v2, v1, v30)
            _ = fi38
        }
    case "":
        _ = len(v3)
    default:
        v39 := GenS4{F0: 1, F1: ("a" + ("hello" + "")), F2: 94}
        _ = v39
    }
    var v40 uint32 = (v1 << ((18446744073709551615) & 63))
    for fi41 := 0; fi41 < 8; fi41++ {
        var v42 uint32 = uint32(v2)
        _ = fi41
        _ = v42
    }
    if false {
        v2 = (^(v2))
    }
    fmt.Println(v40, v30, v2)
    v3 = append(v3, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v40
}
