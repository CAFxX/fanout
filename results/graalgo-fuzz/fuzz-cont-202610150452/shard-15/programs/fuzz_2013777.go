package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 string; F1 int8 }

func main() {
    v1 := make(map[string]string)
    v2 := []string{"hello"}
    var v3 int32 = 0
    v4 := GenS1{F0: 910, F1: 65535, F2: 727}
    v3 += (v3 << ((1) & 63))
    v3 -= (v3 >> ((1) & 63))
    if true {
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            var v5 int8 = 124
            _ = v5
        case "":
            fmt.Println(v3)
        case "hello":
            delete(v1, "")
        default:
            switch -249 {
            case -874:
                for fi6 := 0; fi6 < 18; fi6++ {
                    var v7 float64 = ((3.14159 * -2.25) / 1.0)
                    v7 += v7
                    for fi8 := 0; fi8 < 17; fi8++ {
                        switch "a" {
                        case "hello":
                            var v9 uint8 = 2
                            if (float32(1.5) == 1.5) {
                                fmt.Println(fi8)
                                v7 = 0.0
                                var v10 uint16 = 7
                                _ = v10
                            } else {
                                var v11 uint64 = 18446744073709551615
                                _ = v11
                            }
                            var v12 int = (fi6 << ((470) & 63))
                            v3 += 1
                            var v13 float64 = v7
                            _ = v12
                            _ = v13
                            _ = v9
                        case "foo bar":
                            v3 += (v3 >> ((255) & 63))
                        default:
                            var v14 uint32 = 2
                            _ = v14
                        }
                        _ = fi8
                    }
                    _ = fi6
                    _ = v7
                }
            case 487:
                if true {
                    v2 = append(v2, ("" + "foo bar"))
                    var v15 int64 = -1
                    _ = v15
                }
            case 1:
                v1[("foo bar" + "\n\t")] = ("hello" + "foo bar")
            default:
                v16 := GenS2{F0: (!(((float64(1.5) / (1e10 * 3.14159)) != -2.25) && true))}
                _ = v16
            }
        }
    }
    v2 = append(v2, "\n\t")
    _ = len(v2)
    switch "\n\t" {
    case "hello":
        v3 = v3
    case "a":
        v1["a"] = ("\n\t" + "foo bar")
    default:
        if false {
            v3 -= (-(v3))
            v1["hello"] = "\n\t"
        } else {
            delete(v1, ("hello" + "a"))
            for fi17 := 0; fi17 < 1; fi17++ {
                switch "\n\t" {
                case "\n\t":
                    if (1000 <= 32767) {
                        v3 = v3
                    } else {
                        for fi18 := 0; fi18 < 2; fi18++ {
                            var v19 uint16 = 76
                            var v20 int32 = -2147483648
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        var v21 float32 = 0.0
                        v21 = 0.1
                        _ = v21
                    }
                default:
                    var v22 uint64 = 497
                    _ = v22
                }
                _ = fi17
            }
            fmt.Println(v3)
        }
    }
    var v23 string = "\n\t"
    for fi24 := 0; fi24 < 5; fi24++ {
        switch 32767 {
        case 618:
            var v25 int64 = -9223372036854775808
            _ = v25
        default:
            switch 0 {
            case 2147483647:
                v1[v23] = ""
            case 32767:
                v23 = v23
            default:
                if false {
                    if ((!false) && false) {
                        v26 := GenS3{F0: 1000, F1: -9223372036854775808}
                        _ = v26
                    }
                } else {
                    var v27 uint64 = 500
                    _ = v27
                }
            }
        }
        var v28 uint64 = 0
        _ = len(v2)
        _ = fi24
        _ = v28
    }
    v3 += v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        var v30 int16 = -32768
        if (false || false) {
            delete(v1, ("" + (("\n\t" + "hello") + v23)))
        } else {
            for fi31 := 0; fi31 < 4; fi31++ {
                fmt.Println(fi31, fi29, fi29)
                _ = fi31
            }
        }
        for fi32 := 0; fi32 < 11; fi32++ {
            v33 := GenS4{F0: "", F1: -77}
            if (!(2 < 659)) {
                fmt.Println(fi29, v23, v23)
            } else {
                v30 += v30
            }
            _ = fi32
            _ = v33
        }
        _ = fi29
        _ = v30
    }
    switch -1 {
    case 39:
        fmt.Println(v23)
    case 9223372036854775807:
        var v34 uint = 1
        _ = v34
    default:
        if false {
            v35 := GenS2{F0: true}
            if true {
                v36 := GenS3{F0: 42, F1: 42}
                if (true || true) {
                    var v37 int64 = 9223372036854775807
                    var v38 uint64 = 18446744073709551615
                    _ = v37
                    _ = v38
                } else {
                    _ = len(v2)
                    for fi39 := 0; fi39 < 6; fi39++ {
                        fmt.Println(fi39, v23, v3)
                        var v40 int64 = 0
                        _ = fi39
                        _ = v40
                    }
                    delete(v1, "a")
                }
                _ = v36
            } else {
                var v41 int8 = 42
                v41 += 100
                v23 = (v23 + "foo bar")
                _ = v41
            }
            var v42 int64 = 1
            _ = v35
            _ = v42
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
