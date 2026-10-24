package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 uint; F1 int }
func genFn0(p0 bool, p1 string) uint32 {
    v6 := []int{937, -1}
    p0 = ((false || (402 == 193)) && (p0 || p0))
    switch 9223372036854775807 {
    case 1:
        for fi7 := 0; fi7 < 11; fi7++ {
            p0 = false
            for fi8 := 0; fi8 < 8; fi8++ {
                p1 = p1
                if false {
                    v9 := GenS1{F0: ""}
                    v10 := GenS2{F0: 0}
                    switch p1 {
                    case "hello":
                        v11 := GenS2{F0: 32767}
                        _ = v11
                    default:
                        var v12 float64 = 0.0
                        _ = v12
                    }
                    _ = v10
                    _ = v9
                }
                _ = fi8
            }
            _ = fi7
        }
    case 42:
        p0 = true
    default:
        for fi13 := 0; fi13 < 13; fi13++ {
            if len(v6) > 0 {
                v6[0] = fi13
            }
            fmt.Println(p1, p0, p1)
            v14 := GenS3{F0: 0, F1: 438, F2: 569}
            _ = fi13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS4{F0: 42, F1: 2147483647, F2: 65535}
        _ = fi15
        _ = v16
    }
    if true {
        switch p1 {
        case "hello":
            var v17 uint64 = 18446744073709551615
            _ = v17
        case "":
            v6 = append(v6, 2147483647)
        case "foo bar":
            fmt.Println(p0, p1)
        default:
            switch "" {
            case "hello":
                var v18 uint64 = 4294967295
                _ = v18
            }
        }
    } else {
        if false {
            p0 = (-9223372036854775808 == 127)
        } else {
            switch 69 {
            case 2147483647:
                var v19 uint16 = 0
                _ = v19
            case 32768:
                v20 := GenS1{F0: ((p1 + p1) + p1)}
                _ = v20
            case 1:
                p1 = "hello"
            }
            switch ("hello" + ((p1 + ("" + p1)) + "a")) {
            case "hello":
                v21 := GenS4{F0: 4294967295, F1: 2147483647, F2: 981}
                _ = v21
            case "\n\t":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case "":
                switch 32768 {
                case 385:
                    fmt.Println(p0)
                case 42:
                    for fi22 := 0; fi22 < 10; fi22++ {
                        var v23 bool = ((p0 || p0) && true)
                        v23 = (v23 || v23)
                        _ = fi22
                        _ = v23
                    }
                case 317:
                    fmt.Println(p1)
                default:
                    var v24 uint64 = 1
                    _ = v24
                }
            default:
                for fi25 := 0; fi25 < 12; fi25++ {
                    var v26 bool = p0
                    if len(v6) > 0 {
                        v6[0] = fi25
                    }
                    for fi27 := 0; fi27 < 10; fi27++ {
                        var v28 int16 = -800
                        _ = fi27
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                }
            }
            v6 = append(v6, -9223372036854775808)
        }
        switch 462 {
        case 127:
            var v29 int64 = -79
            _ = v29
        case 9223372036854775807:
            fmt.Println(p0, p0)
        }
        _ = len(v6)
    }
    _ = p0
    _ = p1
    _ = v6
    return 2147483647
}

func main() {
    var v1 string = ""
    v2 := make(map[string]string)
    var v3 int64 = 184
    v4 := []string{"foo bar"}
    var v5 int16 = 61
    if true {
        var v30 int32 = -1
        _ = v30
    } else {
        fmt.Println(v5, v3)
        delete(v2, v1)
        var v31 string = v1
        _ = v31
    }
    var v32 uint = 0
    switch 2147483647 {
    case 32767:
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v32, v1, v1)
        case "":
            v32 = (v32 - v32)
        case "a":
            var v33 int64 = -935
            _ = v33
        default:
            for fi34 := 0; fi34 < 10; fi34++ {
                if len(v4) > 0 {
                    v4[0] = v1
                }
                _ = fi34
            }
        }
    case -913:
        v35 := GenS1{F0: ("\n\t" + ((v1 + v1) + v1))}
        _ = v35
    default:
        var v36 uint64 = 521
        _ = v36
    }
    switch 600 {
    case 0:
        v1 = (v1 + v1)
    case 128:
        for fi37 := 0; fi37 < 8; fi37++ {
            var v38 int32 = -1
            for fi39 := 0; fi39 < 12; fi39++ {
                v1 = "\n\t"
                _ = fi39
            }
            switch fi37 {
            case -1:
                v32 = v32
            case 128:
                if true {
                    var v40 int16 = (v5 / ((int16(v3)) | 1))
                    _ = v40
                }
            case 32768:
                for fi41 := 0; fi41 < 13; fi41++ {
                    v42 := GenS5{F0: -65, F1: (v5 * (^(v5)))}
                    v43 := GenS6{F0: fi37}
                    var v44 string = v1
                    _ = fi41
                    _ = v42
                    _ = v43
                    _ = v44
                }
            default:
                if false {
                    v45 := GenS7{F0: 0, F1: ""}
                    _ = v45
                }
            }
            _ = fi37
            _ = v38
        }
    case -2147483648:
        v1 = (v1 + "\n\t")
    }
    v3 += (v3 | -9223372036854775808)
    v3 -= v3
    switch 9223372036854775807 {
    case -2147483648:
        v46 := GenS4{F0: genFn0(true, v1), F1: 0, F2: 959}
        _ = v46
    default:
        if true {
            v3 = v3
            v47 := GenS8{F0: v32, F1: 1}
            switch "a" {
            case "a":
                var v48 int8 = 94
                _ = v48
            case "":
                switch -596 {
                case 491:
                    v4 = append(v4, "")
                default:
                    if len(v4) > 0 {
                        v4[0] = v1
                    }
                }
            case "hello":
                var v49 int64 = 42
                _ = v49
            }
            _ = v47
        }
    }
    var v50 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v50
}
