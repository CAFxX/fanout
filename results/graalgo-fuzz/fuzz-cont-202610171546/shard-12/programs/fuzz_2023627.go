package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 uint8; F2 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 uint; F2 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64; F1 int; F2 uint8 }
type GenS10 struct { F0 string; F1 int16 }
type GenS11 struct { F0 uint; F1 int64 }
type GenS12 struct { F0 string }
func genFn0(p0 uint, p1 uint32) int32 {
    switch ("hello" + "") {
    case "\n\t":
        if true {
            var v4 int16 = -32768
            var v5 float64 = (0.1 - -2.25)
            switch (("hello" + "hello") + "hello") {
            case "":
                for fi6 := 0; fi6 < 20; fi6++ {
                    var v7 int32 = 42
                    p0 = (p0 + 65535)
                    switch "hello" {
                    case "hello":
                        p0 -= (p0 + (p0 & p0))
                    }
                    _ = fi6
                    _ = v7
                }
            case "a":
                v8 := GenS1{F0: "a"}
                _ = v8
            }
            _ = v4
            _ = v5
        } else {
            p0 += (^(p0))
        }
    }
    var v9 int64 = 42
    if false {
        var v10 int = -470
        _ = v10
    } else {
        if false {
            var v11 float64 = 1.5
            _ = v11
        }
    }
    fmt.Println(v9, p1, p1)
    var v12 uint = p0
    _ = p0
    _ = p1
    _ = v12
    _ = v9
    return 2147483647
}
func genFn1() bool {
    var v13 uint32 = 222
    if (true || (887 < 1)) {
        v14 := GenS2{F0: 804}
        switch ("foo bar" + "\n\t") {
        case "":
            v13 = (v13 & (v13 << ((4294967295) & 63)))
        case "foo bar":
            v13 -= (^(v13))
        }
        fmt.Println(v13)
        _ = v14
    }
    _ = v13
    return true
}

func main() {
    var v1 string = ""
    v2 := make(map[string]string)
    v3 := []string{"\n\t"}
    _ = len(v2)
    switch 323 {
    case 42:
        switch -9223372036854775808 {
        case -1:
            _ = len(v3)
        default:
            if false {
                v15 := GenS1{F0: "a"}
                v16 := GenS3{F0: 32767, F1: 90, F2: 1}
                var v17 string = (v1 + "\n\t")
                _ = v15
                _ = v16
                _ = v17
            } else {
                if true {
                    if false {
                        if (v1 == "") {
                            var v18 uint64 = 1
                            _ = v18
                        }
                    }
                }
                switch -1 {
                case 2147483647:
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                case -482:
                    switch -2147483648 {
                    case 42:
                        v19 := GenS2{F0: 0}
                        _ = v19
                    default:
                        if (!(4294967295 != 4294967295)) {
                            v1 = v1
                        } else {
                            var v20 int64 = 42
                            _ = v20
                        }
                    }
                case -785:
                    v21 := GenS4{F0: 4294967295}
                    _ = v21
                }
                if (4294967295 <= 244) {
                    var v22 int32 = genFn0(702, 42)
                    v1 = v1
                    _ = v22
                } else {
                    fmt.Println(v1)
                    if false {
                        fmt.Println(v1)
                    }
                }
            }
        }
    case 128:
        v23 := GenS5{F0: 42, F1: 111, F2: -407}
        _ = v23
    default:
        switch -9223372036854775808 {
        case 0:
            v1 = (("" + v1) + (v1 + ("" + ("\n\t" + v1))))
        case -1:
            v24 := GenS6{F0: 577}
            _ = v24
        }
    }
    var v25 string = v1
    switch 127 {
    case 835:
        v25 = v1
    default:
        var v26 int64 = -9223372036854775808
        _ = v26
    }
    delete(v2, "hello")
    v27 := GenS2{F0: 42}
    for fi28 := 0; fi28 < 19; fi28++ {
        if (-127 == 42) {
            if false {
                if (true || (-2.25 <= -2.25)) {
                    delete(v2, "foo bar")
                }
            }
            if (((true || (0.1 >= 3.14159)) && true) || true) {
                if false {
                    var v29 bool = false
                    v25 = v25
                    _ = v29
                }
                var v30 string = "hello"
                _ = v30
            }
            v31 := v2[("a" + v25)]
            _ = v31
        } else {
            var v32 uint64 = 79
            fmt.Println(fi28)
            _ = v32
        }
        _ = fi28
    }
    v33 := GenS7{F0: v25, F1: 18446744073709551615, F2: -1}
    fmt.Println(v25, v25)
    v34 := GenS7{F0: v1, F1: 721, F2: 377}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v25, v25)
    if (false && false) {
        v35 := GenS8{F0: 123}
        _ = v35
    } else {
        v36 := GenS5{F0: 1000, F1: 255, F2: 9223372036854775807}
        v1 = ("a" + v25)
        var v37 int64 = -9223372036854775808
        _ = v36
        _ = v37
    }
    v38 := GenS7{F0: v1, F1: 2, F2: 1}
    if false {
        var v39 string = ""
        _ = v39
    } else {
        if false {
            var v40 int8 = 1
            v41 := GenS9{F0: 0, F1: -9223372036854775808, F2: 127}
            _ = v40
            _ = v41
        }
        if false {
            for fi42 := 0; fi42 < 1; fi42++ {
                var v43 uint8 = 2
                v44 := GenS10{F0: "foo bar", F1: 1}
                fmt.Println(fi42)
                _ = fi42
                _ = v43
                _ = v44
            }
            switch 1 {
            case 32768:
                for fi45 := 0; fi45 < 12; fi45++ {
                    var v46 int16 = 0
                    var v47 int32 = 1
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            case 809:
                if (true && false) {
                    v48 := v2["hello"]
                    for fi49 := 0; fi49 < 19; fi49++ {
                        switch (fi49 - fi49) {
                        case 93:
                            fmt.Println(v25, fi49, v1)
                        case 834:
                            v50 := GenS4{F0: 65535}
                            _ = v50
                        }
                        var v51 string = v1
                        v25 = "foo bar"
                        _ = fi49
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 11; fi52++ {
                        switch "a" {
                        case "foo bar":
                            v25 = ("" + v25)
                        case "hello":
                            for fi53 := 0; fi53 < 8; fi53++ {
                                var v54 int8 = -5
                                delete(v2, ((((v48 + ((v1 + v25) + "")) + v1) + "") + v1))
                                _ = fi53
                                _ = v54
                            }
                        default:
                            var v55 int = fi52
                            _ = v55
                        }
                        if (32767 >= 32767) {
                            for fi56 := 0; fi56 < 20; fi56++ {
                                var v57 float32 = (1.5 / (float32(float64(1.5)) / 1.5))
                                _ = fi56
                                _ = v57
                            }
                            fmt.Println(v48, v25, v1)
                            v58 := GenS11{F0: 255, F1: -100}
                            _ = v58
                        }
                        _ = fi52
                    }
                    _ = v48
                }
            }
        } else {
            v59 := GenS12{F0: "\n\t"}
            _ = v59
        }
        v25 = ""
    }
    v25 = v25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v38
}
