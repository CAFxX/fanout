package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int32; F1 int32; F2 int64 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS8 struct { F0 int32; F1 bool; F2 bool }
type GenS9 struct { F0 int16; F1 uint }
type GenS10 struct { F0 int32; F1 bool }
type GenS11 struct { F0 uint }
func genFn0(p0 bool, p1 int8, p2 int) int8 {
    switch "hello" {
    case "a":
        if p0 {
            var v4 int = (p2 * p2)
            switch v4 {
            case -9223372036854775808:
                v5 := []int{42, p2, (^(v4)), (p2 << ((4294967295) & 63))}
                _ = v5
            case 42:
                var v6 float64 = ((3.14159 - ((1e10 - 1e10) + (-2.25 * (3.14159 - 1.5)))) * 0.1)
                _ = v6
            }
            var v7 float64 = float64(1e10)
            _ = v4
            _ = v7
        } else {
            fmt.Println(p1, p2, p2)
            for fi8 := 0; fi8 < 7; fi8++ {
                fmt.Println(fi8, p1)
                _ = fi8
            }
        }
    }
    if ((4294967295 <= 42) && false) {
        fmt.Println(p1, p1)
    } else {
        fmt.Println(p2, p2, p1)
        var v9 string = ("a" + "\n\t")
        _ = v9
    }
    v10 := []int{9223372036854775807}
    switch 449 {
    case 127:
        for fi11 := 0; fi11 < 4; fi11++ {
            if (true || (756 < 127)) {
                p0 = p0
                if false {
                    if len(v10) > 0 {
                        v10[0] = (^(fi11))
                    }
                    for fi12 := 0; fi12 < 12; fi12++ {
                        fmt.Println(fi11, p0, fi11)
                        _ = fi12
                    }
                } else {
                    if (-9223372036854775808 <= 836) {
                        _ = len(v10)
                        p2 += (-(p2))
                    }
                }
            }
            if ((381 > 266) && (p0 && ("" == "a"))) {
                v13 := GenS1{F0: -32768, F1: 255}
                v14 := GenS2{F0: 0, F1: 42, F2: -385}
                _ = v13
                _ = v14
            }
            _ = fi11
        }
    case 145:
        var v15 uint32 = 300
        _ = v15
    default:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS3{F0: 2147483647, F1: -240}
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 int64 = 0
            v20 := GenS1{F0: -1, F1: 1}
            p2 = fi18
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (p1 >> ((403) & 63))
}

func main() {
    v1 := []uint32{937, 954}
    v2 := make(map[int]bool)
    var v3 string = ""
    v2[0] = (true || true)
    switch 4294967295 {
    case 943:
        v3 = "foo bar"
    case 4294967295:
        v21 := GenS4{F0: 0}
        _ = v21
    case 9223372036854775807:
        v22 := GenS1{F0: -595, F1: 2}
        _ = v22
    default:
        v3 = "foo bar"
    }
    if (false || true) {
        v2[56] = true
    }
    if true {
        if (42 == 729) {
            if false {
                v23 := GenS4{F0: 0}
                fmt.Println(v3)
                if true {
                    fmt.Println(v3)
                    if (!true) {
                        for fi24 := 0; fi24 < 15; fi24++ {
                            v25 := GenS5{F0: 42}
                            _ = fi24
                            _ = v25
                        }
                        v3 = "foo bar"
                        var v26 uint8 = 128
                        _ = v26
                    }
                } else {
                    var v27 int64 = 9223372036854775807
                    fmt.Println(v27, v3)
                    _ = v27
                }
                _ = v23
            } else {
                v3 = (((("hello" + (("\n\t" + (("" + v3) + v3)) + (((v3 + "\n\t") + ((("" + v3) + "foo bar") + v3)) + v3))) + v3) + ("a" + "\n\t")) + v3)
                for fi28 := 0; fi28 < 16; fi28++ {
                    v3 = (((v3 + (v3 + "\n\t")) + ("hello" + v3)) + "")
                    if (77 >= -62) {
                        delete(v2, (-(fi28)))
                        v3 = v3
                    }
                    if false {
                        v29 := GenS6{F0: 1, F1: (true && true)}
                        _ = v29
                    }
                    _ = fi28
                }
                switch ((v3 + "foo bar") + ((v3 + v3) + v3)) {
                case "foo bar":
                    if true {
                        v3 = v3
                        if true {
                            if false {
                                var v30 uint = 4294967295
                                var v31 int64 = 42
                                _ = v30
                                _ = v31
                            }
                            _ = len(v2)
                        }
                        v3 = "hello"
                    }
                }
            }
        } else {
            v32 := GenS7{F0: 32767, F1: genFn0((!true), 127, 9223372036854775807), F2: 42}
            _ = v32
        }
        for fi33 := 0; fi33 < 12; fi33++ {
            if true {
                v3 = (v3 + "\n\t")
                var v34 uint16 = 42
                _ = v34
            }
            _ = fi33
        }
        v35 := GenS5{F0: 255}
        _ = v35
    } else {
        fmt.Println(v3)
        _ = len(v2)
        var v36 int32 = -65
        _ = v36
    }
    for fi37 := 0; fi37 < 12; fi37++ {
        _ = len(v2)
        _ = fi37
    }
    if false {
        v1 = append(v1, 65535)
        _ = len(v2)
        for fi38 := 0; fi38 < 10; fi38++ {
            if ((0.1 + 1.5) <= (0.0 + -2.25)) {
                switch "hello" {
                case "\n\t":
                    for fi39 := 0; fi39 < 3; fi39++ {
                        v3 = v3
                        _ = fi39
                    }
                default:
                    delete(v2, 4294967295)
                }
                var v40 int8 = genFn0(false, 0, 2147483647)
                var v41 string = v3
                _ = v40
                _ = v41
            } else {
                if true {
                    v42 := GenS8{F0: -134, F1: (!true), F2: true}
                    v2[-9223372036854775808] = (!((fi38 * fi38) >= (fi38 & fi38)))
                    _ = v42
                } else {
                    v3 = (("hello" + v3) + "\n\t")
                    v3 = (((v3 + "foo bar") + "") + "foo bar")
                    for fi43 := 0; fi43 < 14; fi43++ {
                        fmt.Println(fi43)
                        _ = fi43
                    }
                }
                delete(v2, (fi38 / ((-172) | 1)))
                for fi44 := 0; fi44 < 1; fi44++ {
                    fmt.Println(v3, fi44)
                    v3 = ("a" + "a")
                    v1 = append(v1, 75)
                    v45 := v2[9223372036854775807]
                    _ = fi44
                    _ = v45
                }
            }
            _ = fi38
        }
    } else {
        for fi46 := 0; fi46 < 3; fi46++ {
            for fi47 := 0; fi47 < 17; fi47++ {
                switch v3 {
                case "\n\t":
                    if false {
                        var v48 int = 128
                        _ = v48
                    }
                case "":
                    v49 := GenS7{F0: 2, F1: -128, F2: 255}
                    _ = v49
                default:
                    v50 := GenS9{F0: 0, F1: 598}
                    _ = v50
                }
                v3 = ((v3 + "") + "hello")
                var v51 float32 = 3.14159
                _ = fi47
                _ = v51
            }
            fmt.Println(v3)
            switch v3 {
            case "a":
                for fi52 := 0; fi52 < 2; fi52++ {
                    var v53 int32 = -1
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                    var v54 string = ((v3 + (((v3 + v3) + v3) + "a")) + v3)
                    v53 = 2147483647
                    _ = fi52
                    _ = v53
                    _ = v54
                }
            case "hello":
                if false {
                    v3 = ((v3 + (((v3 + "\n\t") + v3) + v3)) + "a")
                } else {
                    fmt.Println(fi46)
                }
            }
            _ = fi46
        }
        switch (v3 + "a") {
        case "":
            v55 := GenS10{F0: 42, F1: (255 == 905)}
            _ = v55
        case "\n\t":
            v3 = (v3 + v3)
        default:
            switch ("foo bar" + (v3 + "a")) {
            case "foo bar":
                if ((!false) && true) {
                    var v56 float64 = (0.1 * 0.0)
                    _ = v56
                } else {
                    fmt.Println(v3)
                    v3 = ("hello" + v3)
                    v57 := GenS3{F0: -2147483648, F1: 2147483647}
                    _ = v57
                }
            case "hello":
                var v58 int8 = -1
                _ = v58
            }
        }
        v59 := GenS11{F0: 18446744073709551615}
        _ = v59
    }
    fmt.Println(v3)
    var v60 float32 = (1.5 / 0.1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v60
}
