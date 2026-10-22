package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 bool; F2 int64 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 int32; F2 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint; F1 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 125
    v3 := []string{}
    var v4 int8 = (-(v2))
    v2 -= v4
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 int = (fi5 << ((42) & 63))
        var v7 uint = 65535
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint64 = 18446744073709551615
            switch ("foo bar" + "foo bar") {
            case "hello":
                var v10 int32 = 1
                _ = v10
            case "foo bar":
                v7 = v7
            case "\n\t":
                v11 := GenS1{F0: 127}
                _ = v11
            }
            var v12 uint = (v7 & 1)
            var v13 uint64 = 0
            _ = fi8
            _ = v12
            _ = v13
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    if (42 < 32767) {
        v2 = v2
        if (false && (!(1.5 >= ((0.1 / 0.1) / ((0.0 - 1.5) / ((1.5 + (1.5 - 1e10)) + 1e10)))))) {
            v2 += v2
        } else {
            if (255 <= 216) {
                if true {
                    _ = len(v1)
                } else {
                    fmt.Println(v4, v4)
                    var v14 bool = (42 < 244)
                    _ = v14
                }
                v15 := GenS1{F0: 0}
                _ = v15
            }
            switch 32767 {
            case -427:
                v2 = -1
            }
            fmt.Println(v2, v4)
        }
        var v16 uint64 = 393
        _ = v16
    } else {
        var v17 uint32 = 756
        _ = v17
    }
    var v18 string = ("" + ("hello" + ("a" + "a")))
    for fi19 := 0; fi19 < 1; fi19++ {
        if (0.1 >= -2.25) {
            var v20 int32 = 0
            _ = v20
        }
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 uint64 = 219
            switch v18 {
            case "foo bar":
                v4 = (v2 ^ 1)
            case "a":
                v23 := GenS1{F0: 1}
                _ = v23
            case "\n\t":
                if true {
                    var v24 uint8 = 255
                    _ = v24
                } else {
                    if true {
                        v25 := GenS1{F0: 127}
                        if false {
                            v26 := GenS2{F0: 243, F1: true, F2: -9223372036854775808}
                            v18 = ""
                            _ = v26
                        } else {
                            v2 = (v2 + (v2 >> ((2) & 63)))
                            var v27 float64 = 0.1
                            _ = v27
                        }
                        _ = v25
                    }
                }
            default:
                v1[(v18 + ("foo bar" + "a"))] = 65535
                v28 := GenS3{F0: 42, F1: 2147483647, F2: 192}
                _ = v28
            }
            var v29 float64 = 0.0
            _ = fi21
            _ = v22
            _ = v29
        }
        if true {
            switch (fi19 << ((55) & 63)) {
            case 32768:
                v30 := GenS4{F0: (-(v2))}
                _ = v30
            case -2147483648:
                if len(v3) > 0 {
                    v3[0] = ((v18 + v18) + v18)
                }
            case 0:
                v18 = (("a" + v18) + v18)
            default:
                for fi31 := 0; fi31 < 12; fi31++ {
                    var v32 bool = true
                    v33 := GenS5{F0: fi31, F1: 0, F2: 0}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            }
        } else {
            v34 := GenS1{F0: 128}
            v35 := GenS3{F0: 2147483647, F1: 4294967295, F2: 153}
            _ = v34
            _ = v35
        }
        _ = fi19
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        switch ("a" + ((("\n\t" + "a") + v18) + v18)) {
        case "a":
            v37 := GenS6{F0: 0}
            _ = v37
        case "":
            fmt.Println(fi36, v2)
        case "foo bar":
            if false {
                fmt.Println(fi36)
                if (4294967295 <= 0) {
                    var v38 uint64 = 2
                    var v39 uint = 335
                    _ = v38
                    _ = v39
                } else {
                    if false {
                        if (float32(-2.25) < 3.14159) {
                            v2 += (-(v4))
                        } else {
                            if false {
                                v4 = v2
                            }
                            delete(v1, (v18 + v18))
                            switch (v18 + "a") {
                            case "\n\t":
                                var v40 int8 = (v4 + v4)
                                _ = v40
                            case "foo bar":
                                if (!false) {
                                    fmt.Println(v4)
                                }
                            }
                        }
                        v18 = "hello"
                        switch ("hello" + (("hello" + "foo bar") + ("hello" + v18))) {
                        case "foo bar":
                            v2 = 100
                        case "hello":
                            v41 := GenS4{F0: (v2 >> ((0) & 63))}
                            _ = v41
                        case "a":
                            for fi42 := 0; fi42 < 7; fi42++ {
                                v18 = (v18 + (v18 + v18))
                                _ = fi42
                            }
                        }
                    }
                }
            } else {
                _ = len(v3)
            }
        default:
            for fi43 := 0; fi43 < 5; fi43++ {
                v44 := GenS2{F0: -668, F1: ((!true) || true), F2: 955}
                _ = fi43
                _ = v44
            }
        }
        _ = fi36
    }
    v45 := GenS7{F0: 355, F1: 1}
    v46 := v1[(v18 + v18)]
    var v47 int16 = 32767
    for fi48 := 0; fi48 < 5; fi48++ {
        if true {
            switch (fi48 - fi48) {
            case -846:
                var v49 int16 = int16(v2)
                _ = v49
            case 635:
                fmt.Println(v47, fi48, v46)
            }
            var v50 int16 = 605
            _ = v50
        } else {
            _ = len(v3)
        }
        var v51 int = (fi48 ^ fi48)
        _ = fi48
        _ = v51
    }
    if true {
        switch 718 {
        case -1:
            if (v18 != "foo bar") {
                fmt.Println(v2, v18)
                delete(v1, "a")
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
            }
        case -9223372036854775808:
            if true {
                v4 -= (v2 | (v4 ^ 127))
                var v52 uint = 65535
                _ = v52
            }
        case 2147483647:
            v46 -= (^(v46))
        }
        v4 -= v2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v53 int64 = 1
        v3 = append(v3, (v18 + v18))
        var v54 uint = 347
        _ = v53
        _ = v54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
}
