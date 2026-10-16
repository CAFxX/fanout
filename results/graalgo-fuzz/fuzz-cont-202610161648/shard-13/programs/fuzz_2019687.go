package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64; F1 string; F2 uint }
type GenS3 struct { F0 int64; F1 uint8; F2 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 uint32; F2 uint16 }
type GenS6 struct { F0 string; F1 uint64; F2 uint16 }
type GenS7 struct { F0 int32; F1 int; F2 uint64 }
type GenS8 struct { F0 int64; F1 uint16; F2 bool }

func main() {
    var v1 float64 = ((-2.25 / 0.1) + 3.14159)
    var v2 float64 = 1.5
    v3 := []string{}
    v4 := make(map[uint32]uint32)
    if (!true) {
        var v5 int = -1
        fmt.Println(v2)
        v6 := GenS1{F0: v5, F1: (false && (true && (18446744073709551615 == 255)))}
        _ = v5
        _ = v6
    } else {
        fmt.Println(v2, v1)
        switch (("foo bar" + "a") + "\n\t") {
        case "\n\t":
            var v7 int8 = 0
            _ = v7
        case "hello":
            for fi8 := 0; fi8 < 8; fi8++ {
                v9 := GenS2{F0: 590, F1: "hello", F2: 255}
                v2 = 0.0
                _ = fi8
                _ = v9
            }
        }
    }
    v10 := GenS3{F0: -131, F1: 255, F2: 100}
    if (0.1 >= 0.1) {
        _ = len(v4)
    } else {
        fmt.Println(v2, v2)
    }
    if (-9223372036854775808 >= 4294967295) {
        if ((!(false && (640 >= 440))) && false) {
            v11 := GenS4{F0: 0}
            _ = v11
        }
        var v12 uint32 = 2147483648
        fmt.Println(v12, v1)
        _ = v12
    }
    switch 298 {
    case 2147483647:
        v3 = append(v3, "\n\t")
    case -553:
        v2 += float64(((1.5 * (0.0 + (0.1 * 0.0))) / 1.0))
    case 9223372036854775807:
        var v13 uint16 = 32767
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v2 = v1
        _ = fi14
    }
    v15 := GenS5{F0: "hello", F1: 186, F2: 1}
    v2 = (v2 + float64((0.0 + -2.25)))
    v2 = ((0.0 / (((v1 + (v2 / 3.14159)) * ((((0.1 + 0.1) - v1) + 1.5) - 1.5)) - v2)) + (0.0 / (1e10 - 0.0)))
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint64 = 0
        if false {
            switch ((("a" + "foo bar") + ("\n\t" + ("hello" + "foo bar"))) + "foo bar") {
            case "":
                v2 -= ((1e10 * ((1e10 * v2) * ((float64((0.0 / float32((float64(0.0) + (v2 + 3.14159))))) / v2) * (float64(float32(-2.25)) - (float64((((3.14159 / (3.14159 - (3.14159 / 0.1))) - (1e10 + (-2.25 * 1.5))) + (1.5 * 0.0))) / float64(float32((1.5 * v2)))))))) / v1)
            case "hello":
                fmt.Println(v17)
            case "\n\t":
                v4[584] = 65535
            }
            if false {
                if false {
                    v1 = v1
                    if (!(42 <= 0)) {
                        fmt.Println(v1, v17)
                        if (1 > 1) {
                            var v18 float32 = 3.14159
                            if true {
                                for fi19 := 0; fi19 < 1; fi19++ {
                                    fmt.Println(v1, fi16)
                                    v3 = append(v3, "")
                                    var v20 uint32 = 4294967295
                                    _ = fi19
                                    _ = v20
                                }
                                if (3.14159 != 0.1) {
                                    var v21 bool = true
                                    delete(v4, 2147483647)
                                    v3 = append(v3, "\n\t")
                                    _ = v21
                                }
                            } else {
                                for fi22 := 0; fi22 < 7; fi22++ {
                                    switch 127 {
                                    case -1:
                                        _ = len(v3)
                                    default:
                                        var v23 uint = 42
                                        _ = v23
                                    }
                                    v17 = 966
                                    var v24 uint16 = 255
                                    _ = fi22
                                    _ = v24
                                }
                                switch ("hello" + "hello") {
                                case "\n\t":
                                    for fi25 := 0; fi25 < 20; fi25++ {
                                        v26 := GenS3{F0: -1, F1: 51, F2: -41}
                                        var v27 uint64 = 18446744073709551615
                                        fmt.Println(v1, fi16, fi25)
                                        _ = fi25
                                        _ = v26
                                        _ = v27
                                    }
                                case "":
                                    switch ("a" + "a") {
                                    case "":
                                        for fi28 := 0; fi28 < 10; fi28++ {
                                            v2 = 0.0
                                            v29 := GenS4{F0: -572}
                                            v30 := GenS2{F0: (^(v17)), F1: ("" + "foo bar"), F2: 4294967295}
                                            _ = fi28
                                            _ = v29
                                            _ = v30
                                        }
                                    }
                                }
                                if true {
                                    fmt.Println(v17, fi16, v1)
                                } else {
                                    var v31 int8 = 1
                                    switch fi16 {
                                    case 9223372036854775807:
                                        for fi32 := 0; fi32 < 9; fi32++ {
                                            var v33 int = (fi16 * fi32)
                                            switch (-(fi16)) {
                                            case 2147483647:
                                                var v34 bool = (!((fi16 * 32768) != (fi32 + (v33 / ((v33) | 1)))))
                                                fmt.Println(v33)
                                                _ = v34
                                            default:
                                                v1 += 0.0
                                            }
                                            _ = fi32
                                            _ = v33
                                        }
                                    }
                                    _ = v31
                                }
                            }
                            v35 := GenS3{F0: -759, F1: 42, F2: 67}
                            _ = v18
                            _ = v35
                        } else {
                            v2 -= (float64(0.1) / v2)
                            v1 = v1
                            for fi36 := 0; fi36 < 18; fi36++ {
                                var v37 bool = false
                                if len(v3) > 0 {
                                    v3[0] = "a"
                                }
                                _ = fi36
                                _ = v37
                            }
                        }
                    }
                    v1 = v2
                }
                if ((fi16 & fi16) > fi16) {
                    v17 = (^(v17))
                } else {
                    fmt.Println(v2)
                }
            }
            var v38 uint = 42
            _ = v38
        }
        switch ("\n\t" + "") {
        case "foo bar":
            for fi39 := 0; fi39 < 14; fi39++ {
                v1 = v1
                _ = fi39
            }
        case "a":
            v17 -= v17
        case "":
            if (127 >= 255) {
                for fi40 := 0; fi40 < 16; fi40++ {
                    v41 := GenS6{F0: "a", F1: 498, F2: uint16(v17)}
                    _ = fi40
                    _ = v41
                }
            } else {
                if true {
                    v42 := GenS6{F0: ("" + "foo bar"), F1: v17, F2: 255}
                    if true {
                        _ = len(v3)
                        v43 := GenS7{F0: 1, F1: fi16, F2: 42}
                        _ = v43
                    }
                    v44 := GenS8{F0: -241, F1: 0, F2: (9223372036854775807 == -220)}
                    _ = v42
                    _ = v44
                }
                fmt.Println(v1, v17, fi16)
            }
        }
        _ = fi16
        _ = v17
    }
    v2 += v1
    if false {
        if false {
            var v45 int = -2147483648
            for fi46 := 0; fi46 < 20; fi46++ {
                v1 += ((v2 + (v2 - v1)) * (v1 - 0.0))
                _ = fi46
            }
            _ = v45
        }
    }
    var v47 uint32 = 4294967295
    switch 636 {
    case 2147483647:
        var v48 uint32 = 131
        _ = v48
    case 127:
        if true {
            v2 = 0.0
            v3 = append(v3, "")
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            for fi49 := 0; fi49 < 19; fi49++ {
                _ = len(v3)
                _ = fi49
            }
        }
    case 0:
        switch -1 {
        case -9223372036854775808:
            fmt.Println(v47, v2, v1)
        case 42:
            v2 = -2.25
        case -189:
            for fi50 := 0; fi50 < 14; fi50++ {
                v51 := GenS2{F0: uint64(v47), F1: "a", F2: 42}
                for fi52 := 0; fi52 < 1; fi52++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi52
                }
                v53 := v4[(v47 | 507)]
                _ = fi50
                _ = v51
                _ = v53
            }
        default:
            v54 := GenS5{F0: "foo bar", F1: 2147483647, F2: 871}
            _ = v54
        }
    }
    fmt.Println(v47)
    v1 += v1
    v55 := v4[v47]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v55
}
