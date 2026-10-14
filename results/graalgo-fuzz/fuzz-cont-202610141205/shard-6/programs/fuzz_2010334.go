package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 string; F1 int; F2 int16 }
type GenS4 struct { F0 int8; F1 int16; F2 int8 }
func genFn0(p0 int64, p1 uint8) string {
    if false {
        var v6 uint64 = 42
        v7 := []int{}
        _ = v6
        _ = v7
    } else {
        p0 = (^(p0))
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        if false {
            var v9 int = int(p1)
            _ = v9
        }
        fmt.Println(p0, p0)
        p1 = (^(p1))
        _ = fi8
    }
    if (65535 >= 2147483648) {
        v10 := GenS1{F0: 2, F1: (^(p1))}
        _ = v10
    } else {
        fmt.Println(p0, p0)
        for fi11 := 0; fi11 < 15; fi11++ {
            p1 = (p1 + p1)
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    return "a"
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 uint16 = 0
    var v4 uint8 = 255
    var v5 int32 = 2147483647
    v12 := []int{0}
    v4 += (v4 >> ((120) & 63))
    switch "a" {
    case "foo bar":
        if false {
            if true {
                var v13 uint = 492
                fmt.Println(v13, v4, v13)
                v14 := GenS1{F0: v4, F1: v4}
                _ = v13
                _ = v14
            }
            switch -9223372036854775808 {
            case 42:
                var v15 uint16 = (v3 * 65535)
                _ = v15
            case 4294967295:
                var v16 string = ""
                _ = v16
            case 9223372036854775807:
                v3 -= v3
            default:
                fmt.Println(v4, v3, v5)
            }
        }
    case "a":
        v17 := GenS2{F0: v5, F1: 0}
        _ = v17
    }
    _ = len(v12)
    var v18 int = -1
    if true {
        if true {
            for fi19 := 0; fi19 < 15; fi19++ {
                var v20 float64 = 3.14159
                var v21 uint = 1
                _ = fi19
                _ = v20
                _ = v21
            }
        } else {
            v22 := GenS2{F0: (v5 >> ((65535) & 63)), F1: 0}
            _ = v22
        }
        fmt.Println(v5, v5)
        var v23 uint16 = (v3 % (((^(v3))) | 1))
        _ = v23
    }
    v24 := v1[""]
    for fi25 := 0; fi25 < 1; fi25++ {
        v24 -= fi25
        v18 -= (v18 & -1)
        switch (("foo bar" + "") + "hello") {
        case "hello":
            for fi26 := 0; fi26 < 4; fi26++ {
                var v27 uint32 = 42
                v18 -= 9223372036854775807
                var v28 int32 = v5
                _ = fi26
                _ = v27
                _ = v28
            }
        default:
            v29 := GenS2{F0: (v5 * 42), F1: 42}
            _ = v29
        }
        _ = fi25
    }
    v3 = v3
    switch "foo bar" {
    case "hello":
        switch ("" + "hello") {
        case "\n\t":
            if true {
                v24 = (v24 | (v24 % ((4294967295) | 1)))
                v18 += (^(v18))
                var v30 int64 = 9223372036854775807
                _ = v30
            } else {
                var v31 uint32 = 2
                var v32 uint64 = 0
                v3 = (v3 % (((^(v3))) | 1))
                _ = v31
                _ = v32
            }
        case "a":
            v33 := v2[(v18 | (v18 & 4294967295))]
            _ = v33
        case "foo bar":
            if false {
                var v34 uint = 420
                v35 := GenS3{F0: "hello", F1: (v24 ^ (v24 ^ v24)), F2: 0}
                v4 = (v4 / (((v4 ^ v4)) | 1))
                _ = v34
                _ = v35
            }
        }
    case "\n\t":
        if true {
            var v36 int16 = 0
            _ = v36
        } else {
            v12 = append(v12, (v18 - (v18 / (((v18 - (-(v24)))) | 1))))
            v4 = (^(v4))
        }
    case "a":
        if (v3 < 1) {
            var v37 int = (v24 % (((v24 / ((32767) | 1))) | 1))
            var v38 int64 = 1
            _ = v37
            _ = v38
        }
    }
    if (true && false) {
        for fi39 := 0; fi39 < 17; fi39++ {
            var v40 bool = false
            for fi41 := 0; fi41 < 15; fi41++ {
                v2[fi41] = 42
                if (!(1 == int8(fi39))) {
                    v40 = ((("foo bar" + "a") + "\n\t") == "\n\t")
                    v4 = v4
                }
                _ = fi41
            }
            for fi42 := 0; fi42 < 3; fi42++ {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = fi42
            }
            _ = fi39
            _ = v40
        }
    }
    if false {
        if false {
            var v43 int16 = 32767
            fmt.Println(v5, v4, v5)
            v44 := GenS1{F0: (^(v4)), F1: v4}
            _ = v43
            _ = v44
        } else {
            fmt.Println(v18, v24, v3)
            var v45 int64 = 1
            switch v18 {
            case 2147483647:
                var v46 int8 = -1
                _ = v46
            case 695:
                v45 = (v45 & v45)
            }
            _ = v45
        }
        _ = len(v12)
        fmt.Println(v5)
    } else {
        v47 := GenS4{F0: -127, F1: -1, F2: 42}
        if true {
            v2[v18] = 65535
        } else {
            switch "\n\t" {
            case "":
                var v48 int32 = (v5 / (((v5 ^ v5)) | 1))
                _ = v48
            case "foo bar":
                switch (v18 ^ v18) {
                case 118:
                    _ = len(v12)
                case -2147483648:
                    var v49 string = "hello"
                    _ = v49
                default:
                    if false {
                        var v50 uint64 = 4294967295
                        switch "hello" {
                        case "":
                            for fi51 := 0; fi51 < 12; fi51++ {
                                if false {
                                    if true {
                                        v52 := GenS4{F0: 42, F1: -322, F2: -22}
                                        switch "a" {
                                        case "hello":
                                            v50 -= v50
                                        case "\n\t":
                                            if len(v12) > 0 {
                                                v12[0] = 32767
                                            }
                                        }
                                        _ = v52
                                    } else {
                                        v50 = (^(v50))
                                    }
                                    if (((true && true) || true) || (true || false)) {
                                        delete(v1, "")
                                    }
                                }
                                _ = fi51
                            }
                        }
                        _ = v50
                    } else {
                        v18 -= 128
                    }
                }
            }
            switch "" {
            case "a":
                v53 := GenS1{F0: (v4 >> ((255) & 63)), F1: (^(v4))}
                _ = v53
            case "hello":
                switch ("foo bar" + "") {
                case "\n\t":
                    for fi54 := 0; fi54 < 13; fi54++ {
                        v55 := GenS3{F0: ("hello" + "hello"), F1: v24, F2: -1}
                        v4 = (v4 | (v4 / ((v4) | 1)))
                        _ = fi54
                        _ = v55
                    }
                case "a":
                    fmt.Println(v3, v5)
                }
            }
            v18 = v18
        }
        _ = v47
    }
    var v56 uint16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v56
}
