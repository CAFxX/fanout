package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int32; F2 int64 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 string; F1 int64; F2 string }

func main() {
    v1 := []string{"\n\t", "hello", "\n\t"}
    var v2 int8 = 105
    var v3 uint32 = 0
    for fi4 := 0; fi4 < 10; fi4++ {
        switch "foo bar" {
        case "a":
            switch ("foo bar" + "foo bar") {
            case "":
                v5 := GenS1{F0: v3, F1: 0, F2: 42}
                _ = v5
            case "foo bar":
                if false {
                    for fi6 := 0; fi6 < 4; fi6++ {
                        var v7 int8 = 100
                        _ = fi6
                        _ = v7
                    }
                    v2 = 42
                    var v8 uint = 961
                    _ = v8
                } else {
                    v2 = v2
                    fmt.Println(v3, v3)
                }
            default:
                if true {
                    var v9 uint8 = 238
                    var v10 bool = (true && false)
                    _ = v10
                    _ = v9
                } else {
                    var v11 uint = 324
                    _ = v11
                }
            }
        }
        var v12 int8 = v2
        var v13 int8 = v12
        _ = fi4
        _ = v12
        _ = v13
    }
    switch 995 {
    case -2147483648:
        fmt.Println(v2, v3)
    default:
        if (v3 <= 219) {
            if (1 >= 65535) {
                v1 = append(v1, ("hello" + ""))
                var v14 int8 = -1
                _ = v14
            }
            var v15 uint32 = (v3 & (v3 | v3))
            var v16 int8 = -128
            _ = v15
            _ = v16
        } else {
            if (!(352 <= 18446744073709551615)) {
                fmt.Println(v3)
            }
            if false {
                v3 += v3
                v17 := GenS1{F0: v3, F1: 42, F2: 2}
                _ = v17
            } else {
                if len(v1) > 0 {
                    v1[0] = (("" + "a") + "hello")
                }
                var v18 uint16 = 397
                _ = v18
            }
            v3 += 844
        }
    }
    var v19 int16 = 32767
    var v20 uint32 = v3
    v1 = append(v1, (("hello" + "") + "hello"))
    v20 += 2147483647
    _ = len(v1)
    switch 42 {
    case 32768:
        fmt.Println(v19, v3)
    case 32767:
        fmt.Println(v2)
    case 620:
        if len(v1) > 0 {
            v1[0] = "a"
        }
    }
    switch 1 {
    case 32768:
        switch int(v2) {
        case -891:
            v3 = (v20 % ((v3) | 1))
        case 1:
            v21 := GenS1{F0: (v3 >> ((448) & 63)), F1: 0, F2: 128}
            _ = v21
        }
    case 32767:
        v22 := GenS1{F0: (v20 * v20), F1: 1, F2: 128}
        _ = v22
    default:
        v3 += v3
    }
    var v23 int64 = 757
    v3 -= v20
    if true {
        for fi24 := 0; fi24 < 6; fi24++ {
            fmt.Println(v3, v3)
            var v25 int16 = (v19 & v19)
            _ = fi24
            _ = v25
        }
        _ = len(v1)
        if ((^(v20)) == 42) {
            if (-2147483648 > 2147483647) {
                var v26 int8 = (^(v2))
                _ = v26
            }
        } else {
            v27 := GenS2{F0: 240}
            _ = len(v1)
            _ = v27
        }
    } else {
        for fi28 := 0; fi28 < 16; fi28++ {
            if true {
                v1 = append(v1, "\n\t")
                var v29 int64 = v23
                for fi30 := 0; fi30 < 15; fi30++ {
                    var v31 int16 = v19
                    var v32 float64 = 0.0
                    v33 := GenS2{F0: v19}
                    _ = fi30
                    _ = v31
                    _ = v32
                    _ = v33
                }
                _ = v29
            }
            v34 := GenS2{F0: v19}
            _ = fi28
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        v36 := GenS2{F0: -368}
        fmt.Println(v19)
        for fi37 := 0; fi37 < 19; fi37++ {
            v3 += (^(v20))
            var v38 uint32 = (v20 >> ((2) & 63))
            fmt.Println(v3)
            _ = fi37
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    if ((v2 - (v2 | v2)) == 1) {
        switch 0 {
        case 1:
            if (v23 > 0) {
                switch "" {
                case "":
                    v39 := GenS3{F0: v20, F1: -1, F2: v23}
                    _ = v39
                case "\n\t":
                    switch -587 {
                    case -669:
                        v19 -= v19
                    default:
                        v19 = v19
                    }
                case "foo bar":
                    var v40 string = ""
                    _ = v40
                default:
                    _ = len(v1)
                }
            }
        case -337:
            switch ("foo bar" + "a") {
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                if false {
                    fmt.Println(v3, v19)
                    var v41 uint = 42
                    v20 = v20
                    _ = v41
                } else {
                    var v42 uint32 = 2
                    v42 -= (v42 / ((2147483647) | 1))
                    switch ((("" + "a") + "\n\t") + "foo bar") {
                    case "foo bar":
                        var v43 int16 = (-(v19))
                        _ = v43
                    default:
                        for fi44 := 0; fi44 < 12; fi44++ {
                            var v45 int64 = v23
                            _ = fi44
                            _ = v45
                        }
                    }
                    _ = v42
                }
            }
        default:
            var v46 uint32 = v20
            _ = v46
        }
        if true {
            for fi47 := 0; fi47 < 6; fi47++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v48 int8 = 1
                _ = fi47
                _ = v48
            }
            v49 := GenS4{F0: 921, F1: (^(v20))}
            _ = v49
        }
        switch 32768 {
        case 32768:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 658:
            var v50 int = 32767
            _ = v50
        default:
            v3 = 969
        }
    } else {
        switch 0 {
        case 32767:
            var v51 int8 = 100
            _ = v51
        case 68:
            var v52 int8 = v2
            _ = v52
        }
    }
    _ = len(v1)
    for fi53 := 0; fi53 < 4; fi53++ {
        switch ("" + "a") {
        case "foo bar":
            v54 := GenS5{F0: ("a" + "a"), F1: v23, F2: ("\n\t" + (("foo bar" + ((("a" + ("a" + ("" + "a"))) + ("a" + "\n\t")) + "\n\t")) + "\n\t"))}
            _ = v54
        case "a":
            for fi55 := 0; fi55 < 1; fi55++ {
                v1 = append(v1, "")
                _ = fi55
            }
        }
        _ = fi53
    }
    v23 += (v23 * (v23 * v23))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
