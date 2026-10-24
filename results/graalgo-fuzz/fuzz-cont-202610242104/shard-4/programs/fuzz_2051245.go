package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint64 }
type GenS2 struct { F0 int; F1 int8; F2 int8 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 string; F1 int8 }
func genFn0(p0 string, p1 uint8) uint {
    switch (p0 + "foo bar") {
    case "foo bar":
        p0 = ("\n\t" + p0)
    case "hello":
        for fi5 := 0; fi5 < 13; fi5++ {
            switch p0 {
            case "\n\t":
                for fi6 := 0; fi6 < 3; fi6++ {
                    var v7 uint64 = 2
                    var v8 uint16 = 32767
                    _ = fi6
                    _ = v7
                    _ = v8
                }
            }
            if ((634 > 45) && (!false)) {
                var v9 float64 = 1.5
                if true {
                    fmt.Println(p1, p0)
                    if (false || (4294967295 == 337)) {
                        v10 := GenS1{F0: 94, F1: ((v9 / (1e10 / (v9 * (v9 / float64(-2.25))))) == float64(-2.25)), F2: 18446744073709551615}
                        var v11 int16 = -706
                        v12 := []int{fi5}
                        _ = v10
                        _ = v11
                        _ = v12
                    } else {
                        if (1000 > 42) {
                            v9 += v9
                            v13 := GenS1{F0: 0, F1: ((p0 + "\n\t") < ((p0 + p0) + "a")), F2: 18446744073709551615}
                            _ = v13
                        } else {
                            for fi14 := 0; fi14 < 17; fi14++ {
                                var v15 bool = false
                                _ = fi14
                                _ = v15
                            }
                        }
                    }
                } else {
                    v16 := []int{fi5, (^(fi5)), (fi5 << ((700) & 63)), (-(fi5))}
                    var v17 int32 = -1
                    _ = v16
                    _ = v17
                }
                _ = v9
            }
            _ = fi5
        }
    }
    var v18 uint8 = (p1 % (((p1 / ((p1) | 1))) | 1))
    var v19 int8 = 1
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 0
}
func genFn1(p0 uint16, p1 int32) bool {
    fmt.Println(p0, p0)
    for fi20 := 0; fi20 < 3; fi20++ {
        fmt.Println(p1, fi20)
        _ = fi20
    }
    var v21 uint64 = 671
    fmt.Println(p1, v21, p0)
    v22 := GenS1{F0: 801, F1: true, F2: (^(v21))}
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    return true
}

func main() {
    v1 := []int{32767, 257, 0, 335}
    v2 := make(map[string]string)
    var v3 bool = ((42 > 0) && (42 > 93))
    var v4 bool = (v3 && (!(1000 >= -32768)))
    if v3 {
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
        switch "a" {
        case "hello":
            var v23 uint8 = 3
            _ = v23
        case "\n\t":
            delete(v2, ((("\n\t" + "hello") + "") + "foo bar"))
        }
    }
    switch -940 {
    case -9223372036854775808:
        for fi24 := 0; fi24 < 13; fi24++ {
            for fi25 := 0; fi25 < 15; fi25++ {
                if len(v1) > 0 {
                    v1[0] = fi25
                }
                v4 = (!v4)
                v26 := v2[(((("\n\t" + "hello") + "\n\t") + ("foo bar" + ("foo bar" + "\n\t"))) + "")]
                _ = fi25
                _ = v26
            }
            fmt.Println(v3, fi24, fi24)
            if (1 > 358) {
                var v27 string = "\n\t"
                _ = v27
            }
            _ = fi24
        }
    case 128:
        switch "\n\t" {
        case "":
            v28 := GenS2{F0: 127, F1: 127, F2: -86}
            _ = v28
        case "hello":
            fmt.Println(v3)
        case "foo bar":
            v3 = true
        default:
            var v29 uint = genFn0("foo bar", 2)
            _ = v29
        }
    case 32768:
        if true {
            v3 = (369 != 2)
        } else {
            var v30 int8 = 127
            _ = v30
        }
    }
    v3 = v3
    switch "" {
    case "":
        var v31 uint8 = 42
        _ = v31
    default:
        _ = len(v2)
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        v3 = false
        var v33 float64 = 0.1
        _ = fi32
        _ = v33
    }
    switch -2147483648 {
    case 32768:
        v34 := GenS3{F0: 42, F1: 14, F2: 2147483647}
        _ = v34
    default:
        if (0 > 2147483647) {
            var v35 uint16 = 235
            switch 109 {
            case 776:
                if (false && v4) {
                    v35 += (v35 - v35)
                    fmt.Println(v35, v3)
                }
            case 128:
                switch -590 {
                case 0:
                    v4 = false
                case 1:
                    if v3 {
                        var v36 float32 = 0.1
                        var v37 float32 = 3.14159
                        for fi38 := 0; fi38 < 6; fi38++ {
                            if len(v1) > 0 {
                                v1[0] = (fi38 >> ((217) & 63))
                            }
                            var v39 float32 = v36
                            _ = fi38
                            _ = v39
                        }
                        _ = v36
                        _ = v37
                    } else {
                        v40 := v2["foo bar"]
                        var v41 int8 = 109
                        _ = v40
                        _ = v41
                    }
                case -1:
                    v42 := GenS4{F0: 1, F1: "foo bar"}
                    _ = v42
                default:
                    fmt.Println(v4)
                }
            case 9223372036854775807:
                fmt.Println(v4, v3)
            }
            _ = v35
        } else {
            var v43 uint64 = 18446744073709551615
            _ = v43
        }
    }
    v44 := v2[""]
    var v45 string = v44
    if false {
        var v46 int32 = 0
        v44 = ""
        if (1000 > 1000) {
            var v47 int64 = 42
            v48 := GenS5{F0: ("a" + v45), F1: 127}
            v46 = (v46 - v46)
            _ = v47
            _ = v48
        }
        _ = v46
    }
    v45 = ""
    v2[v44] = v45
    delete(v2, v44)
    switch (v45 + "foo bar") {
    case "\n\t":
        fmt.Println(v3, v3, v3)
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "a":
        v45 = v45
    }
    v4 = v4
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v49 := v2[v44]
    var v50 uint16 = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v49
    _ = v50
}
