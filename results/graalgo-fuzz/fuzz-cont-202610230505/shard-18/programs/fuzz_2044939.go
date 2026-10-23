package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int32 }
type GenS2 struct { F0 int64; F1 uint16; F2 bool }
type GenS3 struct { F0 int; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint32; F2 int }
type GenS5 struct { F0 int; F1 string; F2 string }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint; F1 uint16; F2 int16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int16; F1 int; F2 uint }
type GenS13 struct { F0 int }
type GenS14 struct { F0 bool; F1 int64; F2 bool }

func main() {
    var v1 bool = false
    var v2 int = -872
    var v3 int8 = -1
    var v4 bool = (!(42 < 1))
    v5 := GenS1{F0: int64(v2), F1: 42, F2: 0}
    switch "" {
    case "foo bar":
        v3 += -1
    default:
        if v1 {
            v6 := GenS1{F0: -1, F1: 65535, F2: -1}
            v7 := GenS2{F0: -1, F1: 199, F2: v4}
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS3{F0: 1, F1: 2, F2: 2}
                fmt.Println(v2, v4, fi8)
                _ = fi8
                _ = v9
            }
            _ = v6
            _ = v7
        }
    }
    v10 := []int{128}
    switch (v2 + v2) {
    case 127:
        if v1 {
            var v11 uint64 = 871
            var v12 uint32 = 0
            _ = v11
            _ = v12
        } else {
            v3 -= v3
            var v13 int32 = 2147483647
            v10 = append(v10, 127)
            _ = v13
        }
    case 1:
        v14 := GenS4{F0: -993, F1: 42, F2: -765}
        _ = v14
    }
    v3 += (^(v3))
    v10 = append(v10, 0)
    v1 = true
    for fi15 := 0; fi15 < 12; fi15++ {
        switch (("hello" + (("a" + "\n\t") + ((("" + "\n\t") + "foo bar") + ("foo bar" + "")))) + ("\n\t" + ("foo bar" + "a"))) {
        case "hello":
            var v16 string = ("hello" + "")
            _ = v16
        case "\n\t":
            var v17 float64 = 3.14159
            _ = v17
        case "foo bar":
            fmt.Println(v4, v3, v4)
        }
        var v18 int16 = 42
        var v19 int64 = 395
        _ = fi15
        _ = v18
        _ = v19
    }
    fmt.Println(v1)
    v20 := GenS5{F0: -1, F1: "a", F2: "a"}
    var v21 bool = ((v2 ^ v2) <= v2)
    v2 += (v2 ^ (v2 ^ int(v3)))
    switch ("hello" + ((("\n\t" + ("hello" + "")) + "hello") + ("foo bar" + ("\n\t" + "foo bar")))) {
    case "":
        if v1 {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            v22 := GenS6{F0: "\n\t", F1: "\n\t"}
            for fi23 := 0; fi23 < 11; fi23++ {
                if v21 {
                    _ = len(v10)
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v21 = (!v21)
                        _ = fi24
                    }
                    if true {
                        v25 := GenS7{F0: 561}
                        _ = v25
                    }
                }
                _ = fi23
            }
            _ = v22
        } else {
            if (923 <= -498) {
                v1 = false
            } else {
                fmt.Println(v21)
                var v26 int16 = 42
                _ = v26
            }
            if ((-860 >= 1000) && true) {
                var v27 int = 4294967295
                _ = v27
            } else {
                var v28 uint = 1
                fmt.Println(v4, v4)
                _ = v28
            }
            for fi29 := 0; fi29 < 11; fi29++ {
                switch "hello" {
                case "a":
                    if ((!true) && (true && (!true))) {
                        var v30 bool = (!v21)
                        _ = v30
                    } else {
                        switch (fi29 ^ (fi29 >> ((2) & 63))) {
                        case 231:
                            if len(v10) > 0 {
                                v10[0] = 2147483647
                            }
                        }
                        if false {
                            v21 = (v21 && v4)
                        }
                    }
                }
                fmt.Println(v4, v21, v4)
                _ = fi29
            }
        }
    }
    if (!(false && v21)) {
        var v31 string = ("a" + "foo bar")
        if v21 {
            if ((v3 & (v3 % (((-(v3))) | 1))) != v3) {
                if ((v1 && true) || v1) {
                    if (!v1) {
                        for fi32 := 0; fi32 < 19; fi32++ {
                            fmt.Println(v2)
                            _ = len(v10)
                            var v33 uint64 = 80
                            _ = fi32
                            _ = v33
                        }
                    }
                    var v34 uint32 = 42
                    v35 := GenS8{F0: 0}
                    _ = v34
                    _ = v35
                }
                if len(v10) > 0 {
                    v10[0] = v2
                }
                if (((v31 + v31) + v31) < "\n\t") {
                    if (!(!(961 >= 4294967295))) {
                        switch ("\n\t" + ((v31 + (v31 + "\n\t")) + "a")) {
                        case "hello":
                            if false {
                                v3 -= v3
                                v3 = (v3 + v3)
                                v36 := GenS9{F0: 18446744073709551615, F1: 32767, F2: 32767}
                                _ = v36
                            } else {
                                var v37 uint32 = 1
                                if ((!v4) && (v21 || v21)) {
                                    if v21 {
                                        var v38 uint16 = 2
                                        v39 := GenS10{F0: 0}
                                        _ = v38
                                        _ = v39
                                    }
                                } else {
                                    _ = len(v10)
                                    if (v4 || v1) {
                                        v2 = v2
                                        switch -2147483648 {
                                        case 9223372036854775807:
                                            var v40 int8 = v3
                                            _ = v40
                                        case 0:
                                            if len(v10) > 0 {
                                                v10[0] = (^(v2))
                                            }
                                        default:
                                            switch 9223372036854775807 {
                                            case 4294967295:
                                                v41 := GenS11{F0: int64(v37)}
                                                _ = v41
                                            case 0:
                                                v4 = false
                                            }
                                        }
                                        var v42 uint8 = 2
                                        _ = v42
                                    }
                                    if ((!true) || (128 < 152)) {
                                        v43 := GenS12{F0: 1000, F1: v2, F2: 1}
                                        _ = len(v10)
                                        var v44 string = (("\n\t" + v31) + (v31 + v31))
                                        _ = v43
                                        _ = v44
                                    } else {
                                        if ((v21 && v4) && v21) {
                                            _ = len(v10)
                                            if (0.1 > 1e10) {
                                                if len(v10) > 0 {
                                                    v10[0] = 163
                                                }
                                                v31 = ""
                                            } else {
                                                v45 := GenS13{F0: 912}
                                                _ = v45
                                            }
                                            v37 += 0
                                        }
                                        v46 := GenS9{F0: 32, F1: 32767, F2: 1000}
                                        v2 = (v2 & v2)
                                        _ = v46
                                    }
                                }
                                _ = v37
                            }
                        default:
                            v1 = (839 != 1)
                        }
                        fmt.Println(v1, v21)
                        var v47 float32 = 3.14159
                        _ = v47
                    }
                    var v48 int16 = -76
                    v48 += v48
                    _ = v48
                } else {
                    var v49 int = (v2 | (-(v2)))
                    _ = v49
                }
            }
            v50 := GenS13{F0: v2}
            if false {
                fmt.Println(v1, v21, v2)
                for fi51 := 0; fi51 < 14; fi51++ {
                    fmt.Println(v2, fi51)
                    _ = fi51
                }
            }
            _ = v50
        }
        _ = v31
    } else {
        var v52 int = (-(v2))
        for fi53 := 0; fi53 < 5; fi53++ {
            v54 := GenS6{F0: "hello", F1: "\n\t"}
            if v1 {
                var v55 float32 = 1e10
                var v56 uint = 840
                _ = v55
                _ = v56
            } else {
                v57 := GenS14{F0: v21, F1: 9223372036854775807, F2: (1000 <= -32768)}
                for fi58 := 0; fi58 < 10; fi58++ {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    _ = fi58
                }
                _ = v57
            }
            _ = fi53
            _ = v54
        }
        _ = v52
    }
    var v59 int8 = (v3 + v3)
    var v60 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
}
