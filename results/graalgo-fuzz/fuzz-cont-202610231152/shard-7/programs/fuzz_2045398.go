package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint8; F2 uint }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint32; F1 string }
func genFn0(p0 uint16, p1 uint64, p2 int16) uint64 {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int8 = 42
        _ = fi7
        _ = v8
    }
    switch int(p2) {
    case 396:
        v9 := []int{}
        _ = v9
    case 32767:
        for fi10 := 0; fi10 < 19; fi10++ {
            p0 = (p0 & p0)
            _ = fi10
        }
    default:
        var v11 int8 = -1
        _ = v11
    }
    if (true || (!("foo bar" != ("foo bar" + "a")))) {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 int8 = -51
            p0 = p0
            p2 = p2
            _ = fi12
            _ = v13
        }
    }
    var v14 bool = (-827 <= -1)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 42
}

func main() {
    var v1 int16 = 1
    var v2 uint16 = 32767
    v3 := make(map[int]string)
    var v4 uint32 = 1
    v5 := []string{"a"}
    v6 := make(map[string]uint32)
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            if len(v5) > 0 {
                v5[0] = ""
            }
            _ = fi16
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            v5 = append(v5, "foo bar")
            switch 775 {
            case -491:
                var v18 string = "a"
                _ = v18
            case -1:
                var v19 uint32 = (v4 * (v4 / (((v4 >> ((255) & 63))) | 1)))
                _ = v19
            }
            _ = fi17
        }
        v6["a"] = (v4 + v4)
        _ = fi15
    }
    var v20 uint32 = 65535
    v20 = (^(v20))
    var v21 int8 = 127
    v21 = -128
    var v22 int64 = -1
    if (false || (true && (true && (0 >= v4)))) {
        v22 += (v22 - (v22 - (v22 << ((18446744073709551615) & 63))))
        var v23 int32 = 109
        _ = v23
    }
    var v24 uint32 = (v20 >> ((593) & 63))
    switch "\n\t" {
    case "foo bar":
        var v25 int8 = (v21 & (^(v21)))
        _ = v25
    default:
        switch 127 {
        case 0:
            v3[0] = "foo bar"
        case -150:
            v21 = v21
        case 197:
            for fi26 := 0; fi26 < 11; fi26++ {
                switch ("" + "\n\t") {
                case "":
                    var v27 uint8 = 0
                    _ = v27
                case "a":
                    for fi28 := 0; fi28 < 14; fi28++ {
                        var v29 string = ("" + ((("\n\t" + ("foo bar" + "\n\t")) + "foo bar") + ("\n\t" + "a")))
                        var v30 string = v29
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                default:
                    for fi31 := 0; fi31 < 20; fi31++ {
                        _ = len(v3)
                        if (((!("" != "a")) && true) || false) {
                            var v32 uint = 2
                            switch ("" + "hello") {
                            case "hello":
                                v32 = 18446744073709551615
                            default:
                                v24 = (v20 / (((^(v20))) | 1))
                            }
                            v5 = append(v5, "a")
                            _ = v32
                        } else {
                            var v33 int32 = 2147483647
                            var v34 bool = false
                            switch ("" + ("foo bar" + "hello")) {
                            case "a":
                                v24 += (v24 * uint32(v2))
                            case "\n\t":
                                fmt.Println(fi26)
                                v24 += (v4 / ((v20) | 1))
                                if true {
                                    v22 = (v22 ^ (^(v22)))
                                } else {
                                    var v35 uint32 = 675
                                    var v36 int32 = v33
                                    var v37 uint16 = (v2 - (v2 % (((v2 + v2)) | 1)))
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                            case "hello":
                                if v34 {
                                    fmt.Println(v2, v22)
                                }
                                v38 := GenS1{F0: "a", F1: 185, F2: 0}
                                _ = v38
                            default:
                                v39 := GenS2{F0: fi26}
                                var v40 int64 = (v22 >> ((985) & 63))
                                _ = v39
                                _ = v40
                            }
                            _ = v33
                            _ = v34
                        }
                        if (!true) {
                            v41 := GenS3{F0: fi26, F1: 225, F2: 0}
                            _ = v41
                        } else {
                            if len(v5) > 0 {
                                v5[0] = "\n\t"
                            }
                            v21 = (-(v21))
                            v42 := GenS3{F0: (fi26 - (^(fi31))), F1: 42, F2: 255}
                            _ = v42
                        }
                        v43 := GenS4{F0: "foo bar", F1: fi31}
                        _ = fi31
                        _ = v43
                    }
                }
                _ = fi26
            }
        default:
            v44 := GenS5{F0: v4, F1: ""}
            _ = v44
        }
    }
    v24 -= (v24 * (v24 >> ((2) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
