package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint }
type GenS2 struct { F0 int32; F1 uint; F2 int16 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 int = 2147483647
    v3 := []string{"foo bar", "hello", "\n\t"}
    v4 := GenS1{F0: -9223372036854775808, F1: "\n\t", F2: 4294967295}
    v3 = append(v3, "")
    if true {
        fmt.Println(v2)
        var v5 uint16 = 2
        var v6 float64 = 1.5
        _ = v5
        _ = v6
    } else {
        var v7 uint64 = 1
        switch "\n\t" {
        case "foo bar":
            var v8 uint64 = v7
            _ = v8
        case "hello":
            for fi9 := 0; fi9 < 15; fi9++ {
                for fi10 := 0; fi10 < 14; fi10++ {
                    v7 = (v7 - v7)
                    _ = fi10
                }
                _ = fi9
            }
        case "\n\t":
            var v11 int32 = -166
            _ = v11
        }
        _ = v7
    }
    if true {
        v3 = append(v3, "foo bar")
        if false {
            v2 -= v2
            fmt.Println(v2)
        }
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 int16 = 667
            var v14 int16 = v13
            if false {
                for fi15 := 0; fi15 < 6; fi15++ {
                    v14 = (v13 + v13)
                    _ = fi15
                }
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        v16 := GenS2{F0: 0, F1: 42, F2: 42}
        var v17 int32 = 1
        _ = v16
        _ = v17
    }
    v2 = -9223372036854775808
    if ("" == ("" + "")) {
        if (!(false && ("\n\t" == "hello"))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 = v2
            v18 := GenS1{F0: -129, F1: "\n\t", F2: 4294967295}
            _ = v18
        } else {
            var v19 float64 = ((0.0 * 0.1) + 1e10)
            var v20 int16 = 1000
            _ = v19
            _ = v20
        }
        v21 := GenS3{F0: 227, F1: (v2 ^ (v2 % ((128) | 1)))}
        var v22 int64 = 24
        _ = v21
        _ = v22
    }
    var v23 uint32 = 267
    fmt.Println(v2)
    if ("hello" > "foo bar") {
        var v24 int8 = 13
        if (42 > 646) {
            var v25 bool = true
            _ = v25
        } else {
            var v26 int32 = 452
            _ = v26
        }
        _ = v24
    } else {
        for fi27 := 0; fi27 < 15; fi27++ {
            for fi28 := 0; fi28 < 10; fi28++ {
                v29 := v1["\n\t"]
                fmt.Println(v2, v29, v23)
                v30 := GenS2{F0: 1, F1: 18446744073709551615, F2: 42}
                _ = fi28
                _ = v29
                _ = v30
            }
            if (!(!true)) {
                fmt.Println(v2)
                v31 := GenS4{F0: 42}
                _ = v31
            } else {
                v2 += fi27
                var v32 int = v2
                if false {
                    delete(v1, "a")
                    for fi33 := 0; fi33 < 12; fi33++ {
                        if (1 <= 32767) {
                            var v34 uint = 65535
                            v35 := v1["a"]
                            switch 127 {
                            case -318:
                                v2 -= (fi33 << ((v34) & 63))
                            case 204:
                                fmt.Println(fi33, v23, v35)
                            }
                            _ = v34
                            _ = v35
                        }
                        if false {
                            _ = len(v1)
                            v32 += -2147483648
                        } else {
                            var v36 int8 = -1
                            v37 := GenS5{F0: 527, F1: 818}
                            _ = v36
                            _ = v37
                        }
                        _ = fi33
                    }
                } else {
                    switch "foo bar" {
                    case "":
                        if false {
                            v38 := GenS1{F0: -891, F1: "foo bar", F2: 42}
                            v32 = 1
                            _ = v38
                        }
                    case "a":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "\n\t":
                        var v39 uint8 = 108
                        _ = v39
                    default:
                        if (1 < 9223372036854775807) {
                            v2 += fi27
                            fmt.Println(v2)
                            if ((-2.25 != 0.0) && true) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                var v40 int32 = -969
                                var v41 uint32 = 1
                                for fi42 := 0; fi42 < 18; fi42++ {
                                    v43 := GenS2{F0: v40, F1: 4294967295, F2: 1}
                                    _ = fi42
                                    _ = v43
                                }
                                _ = v40
                                _ = v41
                            }
                        } else {
                            v44 := GenS3{F0: 2147483648, F1: (v32 - v2)}
                            delete(v1, ("\n\t" + ""))
                            _ = v44
                        }
                    }
                    v45 := GenS6{F0: v23, F1: 42}
                    _ = v45
                }
                _ = v32
            }
            _ = len(v1)
            _ = fi27
        }
        var v46 float64 = -2.25
        if (v23 != v23) {
            var v47 int32 = 0
            v2 -= 0
            _ = v47
        }
        _ = v46
    }
    v48 := GenS2{F0: 2147483647, F1: 870, F2: 42}
    if false {
        var v49 int64 = 1
        _ = v49
    } else {
        fmt.Println(v2)
        v50 := GenS5{F0: 4294967295, F1: 610}
        v2 = (v2 - 32767)
        _ = v50
    }
    var v51 uint = 0
    fmt.Println(v51)
    v23 += v23
    if false {
        for fi52 := 0; fi52 < 16; fi52++ {
            _ = len(v3)
            v51 += 0
            _ = fi52
        }
        v53 := GenS7{F0: 0}
        v54 := v1["\n\t"]
        _ = v53
        _ = v54
    } else {
        v2 += v2
    }
    v1["a"] = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v48
    _ = v51
}
