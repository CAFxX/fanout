package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 uint16; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint8; F1 int16 }

func main() {
    v1 := []int{764}
    v2 := make(map[string]string)
    var v3 int32 = 841
    var v4 uint16 = 255
    v5 := []uint32{42}
    var v6 float64 = 1.5
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: 127, F1: 571}
        v9 := GenS1{F0: 42, F1: 255}
        switch (fi7 * fi7) {
        case 9223372036854775807:
            for fi10 := 0; fi10 < 6; fi10++ {
                if (fi10 >= 258) {
                    v5 = append(v5, 4294967295)
                    var v11 uint32 = 0
                    v11 -= 0
                    switch (((("a" + ("\n\t" + ("hello" + "hello"))) + ("\n\t" + "")) + "\n\t") + (("hello" + "foo bar") + "hello")) {
                    case "\n\t":
                        var v12 int32 = -1
                        _ = v12
                    default:
                        if false {
                            fmt.Println(v11, v4, v4)
                            var v13 uint64 = 4294967295
                            v14 := GenS1{F0: uint8(v13), F1: 2}
                            _ = v13
                            _ = v14
                        }
                    }
                    _ = v11
                }
                var v15 uint64 = 4294967295
                v16 := GenS2{F0: ("foo bar" + "")}
                _ = fi10
                _ = v15
                _ = v16
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = len(v2)
    if true {
        if (282 != 909) {
            var v17 float32 = (0.0 / 1e10)
            _ = len(v5)
            var v18 uint8 = 80
            _ = v17
            _ = v18
        }
        if false {
            for fi19 := 0; fi19 < 17; fi19++ {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v3, v4)
                case "":
                    for fi20 := 0; fi20 < 3; fi20++ {
                        var v21 float64 = v6
                        v6 -= 0.1
                        var v22 uint16 = v4
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
                v23 := GenS3{F0: 9223372036854775807, F1: (^(v4)), F2: 679}
                _ = fi19
                _ = v23
            }
            switch 949 {
            case 42:
                for fi24 := 0; fi24 < 13; fi24++ {
                    v4 = (v4 ^ v4)
                    _ = fi24
                }
            case 127:
                if (42 >= (^(v3))) {
                    var v25 int16 = 117
                    fmt.Println(v3, v25, v3)
                    _ = v25
                }
            }
            for fi26 := 0; fi26 < 20; fi26++ {
                v27 := GenS2{F0: "a"}
                _ = fi26
                _ = v27
            }
        } else {
            v2["\n\t"] = ("foo bar" + "foo bar")
            var v28 uint64 = 299
            v6 -= (v6 - float64((-2.25 * 0.0)))
            _ = v28
        }
    } else {
        switch "hello" {
        case "\n\t":
            var v29 int8 = 127
            _ = v29
        case "a":
            switch "a" {
            case "hello":
                var v30 int32 = v3
                _ = v30
            case "\n\t":
                v31 := GenS2{F0: "hello"}
                _ = v31
            case "a":
                v3 = v3
            }
        }
    }
    v4 += v4
    switch "foo bar" {
    case "":
        v32 := GenS3{F0: 200, F1: (v4 | v4), F2: -1}
        _ = v32
    case "a":
        var v33 float32 = (1.5 + (1.5 + float32(1.5)))
        _ = v33
    default:
        if false {
            var v34 string = ("a" + "a")
            var v35 int8 = 127
            _ = len(v2)
            _ = v34
            _ = v35
        } else {
            v4 += (^(v4))
            v6 = (0.1 - v6)
            v36 := GenS3{F0: -9223372036854775808, F1: (v4 % ((v4) | 1)), F2: 42}
            _ = v36
        }
    }
    if (0 <= 0) {
        switch -506 {
        case 128:
            v37 := GenS1{F0: 127, F1: 654}
            _ = v37
        case 32767:
            for fi38 := 0; fi38 < 11; fi38++ {
                var v39 uint64 = 1
                v40 := GenS4{F0: false}
                _ = fi38
                _ = v39
                _ = v40
            }
        case -892:
            if (-50 < -1) {
                v41 := GenS5{F0: 1, F1: v4, F2: -9223372036854775808}
                _ = v41
            } else {
                var v42 float32 = (float32(v6) + ((1e10 + -2.25) + (3.14159 * float32(v6))))
                if (1 < 4294967295) {
                    v6 = 3.14159
                } else {
                    fmt.Println(v3)
                    switch 0 {
                    case -1:
                        v2["\n\t"] = (("foo bar" + "") + "a")
                    }
                }
                _ = v42
            }
        }
    }
    switch 316 {
    case 32767:
        fmt.Println(v6, v4)
    case -9223372036854775808:
        if (976 >= 42) {
            if true {
                v3 += -2147483648
            } else {
                v3 = (-(v3))
                v3 = v3
            }
            v43 := GenS2{F0: "a"}
            _ = v43
        } else {
            v1 = append(v1, 42)
            var v44 int = -478
            for fi45 := 0; fi45 < 16; fi45++ {
                switch ("\n\t" + ("a" + "foo bar")) {
                case "foo bar":
                    var v46 float64 = (((1.5 / (float64(1.5) - (3.14159 * 1.5))) / (v6 / (3.14159 / v6))) * v6)
                    _ = v46
                }
                var v47 int = 2147483647
                var v48 int = (v47 / ((-9223372036854775808) | 1))
                _ = fi45
                _ = v47
                _ = v48
            }
            _ = v44
        }
    default:
        var v49 int = 42
        _ = v49
    }
    for fi50 := 0; fi50 < 19; fi50++ {
        delete(v2, "\n\t")
        switch ("hello" + "a") {
        case "":
            var v51 float32 = float32(3.14159)
            _ = v51
        case "\n\t":
            var v52 uint64 = 0
            _ = v52
        case "hello":
            switch "a" {
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "\n\t":
                if (-385 < -901) {
                    v6 = v6
                }
            default:
                if (false || (!false)) {
                    for fi53 := 0; fi53 < 15; fi53++ {
                        v3 = (v3 * int32(fi50))
                        _ = fi53
                    }
                    if false {
                        v6 += v6
                        v54 := GenS6{F0: 1}
                        _ = v54
                    } else {
                        var v55 int8 = 42
                        v4 += v4
                        if false {
                            v56 := GenS2{F0: "foo bar"}
                            v6 = v6
                            _ = v56
                        }
                        _ = v55
                    }
                }
            }
        }
        v57 := GenS4{F0: (true || false)}
        _ = fi50
        _ = v57
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v58 := GenS7{F0: 127, F1: 42}
    var v59 int32 = -1
    v2[(("hello" + "\n\t") + "hello")] = ("foo bar" + ("foo bar" + "foo bar"))
    for fi60 := 0; fi60 < 20; fi60++ {
        if false {
            v4 = (v4 | 42)
        } else {
            v61 := GenS2{F0: "a"}
            var v62 uint64 = 4294967295
            _ = v61
            _ = v62
        }
        _ = fi60
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
}
