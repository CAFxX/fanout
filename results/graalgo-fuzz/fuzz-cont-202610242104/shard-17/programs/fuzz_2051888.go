package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint; F1 int64; F2 string }
type GenS4 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS5 struct { F0 string }

func main() {
    var v1 uint16 = 42
    var v2 int = 893
    var v3 int64 = 1
    v4 := make(map[string]int)
    v5 := GenS1{F0: (^(v1))}
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int32 = 0
            var v9 int64 = (v3 * (v3 + v3))
            _ = fi7
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            fmt.Println(fi6)
            if (-2147483648 <= 0) {
                var v11 int32 = -862
                _ = v11
            } else {
                if false {
                    v3 = (v3 & 164)
                    var v12 int64 = v3
                    v2 = -104
                    _ = v12
                } else {
                    if false {
                        fmt.Println(fi10)
                    } else {
                        v1 = (v1 - 65535)
                        switch -860 {
                        case -984:
                            v3 += v3
                        case 32767:
                            var v13 int64 = v3
                            _ = v13
                        default:
                            if false {
                                fmt.Println(fi6)
                                var v14 int16 = -692
                                fmt.Println(v1, v1)
                                _ = v14
                            }
                        }
                    }
                    v15 := GenS1{F0: v1}
                    v2 = fi6
                    _ = v15
                }
            }
            _ = fi10
        }
        v3 = v3
        _ = fi6
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        switch "\n\t" {
        case "\n\t":
            if false {
                v1 = (^(v1))
                v4[""] = v2
            }
        case "foo bar":
            v17 := GenS2{F0: 4294967295, F1: -138}
            _ = v17
        default:
            var v18 uint16 = v1
            _ = v18
        }
        v19 := GenS3{F0: 2, F1: (v3 % ((42) | 1)), F2: ""}
        if (-9223372036854775808 <= -1) {
            var v20 uint64 = 2
            if (!((-128 <= 100) && false)) {
                switch ("hello" + "hello") {
                case "hello":
                    v2 += -801
                }
            } else {
                v21 := []int{v2, (v2 % (((v2 & v2)) | 1))}
                switch "\n\t" {
                case "":
                    var v22 bool = false
                    _ = v22
                case "\n\t":
                    var v23 uint32 = 176
                    _ = v23
                case "a":
                    for fi24 := 0; fi24 < 11; fi24++ {
                        v25 := GenS2{F0: 348, F1: -891}
                        _ = fi24
                        _ = v25
                    }
                }
                _ = v21
            }
            v2 = fi16
            _ = v20
        } else {
            var v26 uint8 = 1
            if (-1 == 1) {
                v26 -= (^(v26))
            }
            _ = v26
        }
        _ = fi16
        _ = v19
    }
    fmt.Println(v2)
    for fi27 := 0; fi27 < 6; fi27++ {
        switch (v2 * (v2 ^ (v2 + (fi27 | 128)))) {
        case 706:
            fmt.Println(v3)
        }
        _ = fi27
    }
    fmt.Println(v2, v2, v3)
    if false {
        switch "a" {
        case "":
            v28 := GenS3{F0: 1, F1: 0, F2: "\n\t"}
            _ = v28
        default:
            var v29 float64 = 3.14159
            _ = v29
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            v31 := GenS3{F0: 42, F1: v3, F2: ((("\n\t" + "\n\t") + "\n\t") + "a")}
            v2 += (v2 % ((-958) | 1))
            var v32 uint32 = 42
            _ = fi30
            _ = v31
            _ = v32
        }
        fmt.Println(v3)
    } else {
        for fi33 := 0; fi33 < 7; fi33++ {
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                for fi34 := 0; fi34 < 6; fi34++ {
                    if (true || (true && (true && (214 <= 0)))) {
                        switch ("foo bar" + "\n\t") {
                        case "hello":
                            var v35 float64 = (((-2.25 + 1e10) + (1e10 / 1.0)) / 1e10)
                            _ = v35
                        }
                        var v36 uint16 = 2
                        var v37 uint = 4294967295
                        _ = v36
                        _ = v37
                    } else {
                        var v38 bool = false
                        v2 = fi34
                        v3 = (^(v3))
                        _ = v38
                    }
                    if ("\n\t" >= "\n\t") {
                        fmt.Println(fi33, v3, fi33)
                        switch (v2 / ((-34) | 1)) {
                        case 2147483647:
                            v39 := GenS4{F0: 65535, F1: 65535, F2: 32767}
                            _ = v39
                        case -9223372036854775808:
                            v2 -= fi34
                        case -479:
                            v40 := []int{(fi33 * fi33), v2, v2}
                            _ = v40
                        default:
                            if (float64((1.5 + 3.14159)) == (1.5 - (((-2.25 + 0.0) * (-2.25 * 0.0)) + float64(0.1)))) {
                                fmt.Println(fi34, v3)
                            } else {
                                v4[""] = (^(fi33))
                            }
                        }
                    } else {
                        v41 := v4["\n\t"]
                        if (-1 == 1) {
                            v41 += 42
                        } else {
                            delete(v4, (("" + "foo bar") + "a"))
                            var v42 float64 = -2.25
                            _ = v42
                        }
                        _ = v41
                    }
                    _ = fi34
                }
            case "\n\t":
                switch ("" + ("a" + ("hello" + "hello"))) {
                case "\n\t":
                    v43 := GenS5{F0: "a"}
                    _ = v43
                case "foo bar":
                    if (127 == 1) {
                        v44 := GenS4{F0: 42, F1: 4294967295, F2: v1}
                        _ = v44
                    }
                default:
                    v2 = (v2 ^ (-(v2)))
                }
            }
            switch "hello" {
            case "a":
                var v45 uint16 = (v1 >> ((936) & 63))
                _ = v45
            }
            _ = fi33
        }
        var v46 string = "a"
        _ = v46
    }
    for fi47 := 0; fi47 < 19; fi47++ {
        v1 = (v1 & (v1 * (v1 & (v1 << ((984) & 63)))))
        v3 = v3
        _ = fi47
    }
    for fi48 := 0; fi48 < 12; fi48++ {
        for fi49 := 0; fi49 < 17; fi49++ {
            switch -9 {
            case 0:
                var v50 float32 = 1.5
                _ = v50
            case 32767:
                v2 += fi48
            }
            switch fi49 {
            case 128:
                v2 -= fi49
            case 42:
                v3 = v3
            }
            _ = fi49
        }
        v1 = v1
        _ = fi48
    }
    v51 := GenS4{F0: 2147483647, F1: 4294967295, F2: (v1 + 32767)}
    if (!false) {
        v52 := []int{(v2 - (v2 | 1)), 2147483647, (v2 % ((v2) | 1)), -1}
        _ = v52
    } else {
        switch (("\n\t" + ("\n\t" + "hello")) + ("a" + (("hello" + ("" + (("\n\t" + "\n\t") + ""))) + "foo bar"))) {
        case "":
            var v53 uint = 2
            _ = v53
        }
        v54 := []int{-2147483648, v2, v2}
        _ = v54
    }
    for fi55 := 0; fi55 < 6; fi55++ {
        v56 := GenS2{F0: 0, F1: -587}
        _ = fi55
        _ = v56
    }
    switch "a" {
    case "\n\t":
        var v57 uint = 768
        _ = v57
    default:
        fmt.Println(v2, v3, v2)
    }
    var v58 uint8 = 255
    var v59 bool = (false || true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v58
    _ = v59
}
