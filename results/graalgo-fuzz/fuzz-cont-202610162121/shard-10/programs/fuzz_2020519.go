package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 uint8; F2 bool }

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]uint32)
    var v3 float32 = 1e10
    if false {
        _ = len(v2)
    } else {
        switch "\n\t" {
        case "\n\t":
            for fi4 := 0; fi4 < 14; fi4++ {
                delete(v2, "hello")
                v3 += float32((-2.25 + (3.14159 - -2.25)))
                _ = fi4
            }
        case "a":
            v5 := v2["\n\t"]
            _ = v5
        case "hello":
            var v6 uint16 = 65535
            _ = v6
        default:
            var v7 int64 = 42
            _ = v7
        }
        v8 := []int{}
        v9 := v2["foo bar"]
        _ = v8
        _ = v9
    }
    var v10 float32 = v3
    var v11 uint = 18446744073709551615
    v12 := v2[""]
    if true {
        v2["foo bar"] = (v12 << ((42) & 63))
        v1 += (v1 - (v1 * v1))
        fmt.Println(v11, v12)
    } else {
        v13 := GenS1{F0: 128, F1: 2147483647, F2: 100}
        var v14 string = "\n\t"
        for fi15 := 0; fi15 < 6; fi15++ {
            fmt.Println(fi15, v1, v3)
            switch v14 {
            case "hello":
                v12 = 2
            case "foo bar":
                _ = len(v2)
            case "a":
                var v16 uint16 = 824
                _ = v16
            default:
                for fi17 := 0; fi17 < 5; fi17++ {
                    v1 -= (v1 >> ((1) & 63))
                    var v18 int32 = -2147483648
                    v12 = (^(v12))
                    _ = fi17
                    _ = v18
                }
            }
            var v19 int16 = -599
            _ = fi15
            _ = v19
        }
        _ = v13
        _ = v14
    }
    v3 -= ((float32((0.1 / 1.5)) * (v10 / float32(((3.14159 * (((1.5 / 1.0) * ((1e10 - (float64(float32(3.14159)) * float64((3.14159 + 0.0)))) + 1.5)) / 1.0)) - -2.25)))) - 0.0)
    if true {
        for fi20 := 0; fi20 < 16; fi20++ {
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 int = 128
                if (false || (917 > 42)) {
                    v22 += v22
                }
                var v23 string = "a"
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        v24 := []int{637, 9223372036854775807, -2147483648}
        v10 -= v10
        _ = v24
    } else {
        if false {
            var v25 string = "\n\t"
            fmt.Println(v11, v25, v25)
            _ = v25
        } else {
            for fi26 := 0; fi26 < 16; fi26++ {
                fmt.Println(v12, v12, v1)
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 int = 1
                    _ = fi27
                    _ = v28
                }
                v2[("a" + ("\n\t" + ("foo bar" + "\n\t")))] = (v12 % (((v12 + (v12 | v12))) | 1))
                _ = fi26
            }
        }
        v29 := GenS1{F0: 249, F1: v12, F2: -43}
        _ = v29
    }
    if false {
        if false {
            fmt.Println(v1)
            v30 := GenS2{F0: ("foo bar" + "\n\t"), F1: 127}
            _ = v30
        } else {
            if true {
                v31 := GenS1{F0: 0, F1: (v12 << (((v11 << (((v11 * (v11 << (((v11 & (v11 & (v11 & 0)))) & 63)))) & 63))) & 63)), F2: 42}
                _ = v31
            } else {
                if (true || (!true)) {
                    var v32 int64 = 42
                    if ("\n\t" != "foo bar") {
                        v33 := GenS1{F0: -959, F1: v12, F2: -118}
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 4; fi34++ {
                        for fi35 := 0; fi35 < 18; fi35++ {
                            v3 = v10
                            v3 -= 3.14159
                            v12 = 580
                            _ = fi35
                        }
                        for fi36 := 0; fi36 < 5; fi36++ {
                            v37 := v2["hello"]
                            v38 := v2[("a" + "foo bar")]
                            v32 = v32
                            v39 := GenS2{F0: "foo bar", F1: 42}
                            v1 = v1
                            v40 := GenS1{F0: (fi34 >> ((903) & 63)), F1: v37, F2: -1}
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        var v41 uint32 = v12
                        _ = fi34
                        _ = v41
                    }
                    _ = v32
                }
                var v42 int32 = 42
                _ = v42
            }
            for fi43 := 0; fi43 < 1; fi43++ {
                for fi44 := 0; fi44 < 6; fi44++ {
                    v3 -= v10
                    var v45 int16 = 83
                    _ = fi44
                    _ = v45
                }
                _ = fi43
            }
            fmt.Println(v3, v10)
        }
    } else {
        switch 0 {
        case -2147483648:
            for fi46 := 0; fi46 < 10; fi46++ {
                var v47 int16 = 42
                v48 := GenS3{F0: v12}
                fmt.Println(v3)
                _ = fi46
                _ = v47
                _ = v48
            }
        default:
            var v49 string = ""
            _ = v49
        }
        v11 += (v11 | 2)
        _ = len(v2)
    }
    if false {
        for fi50 := 0; fi50 < 17; fi50++ {
            for fi51 := 0; fi51 < 11; fi51++ {
                v2["foo bar"] = v12
                v2[("foo bar" + (("foo bar" + "hello") + "hello"))] = (v12 & v12)
                _ = fi51
            }
            var v52 float64 = ((0.0 - ((1.5 / (((1.5 - (-2.25 * -2.25)) + 0.1) + 0.1)) / 1.5)) * (3.14159 - (-2.25 / 1.5)))
            v53 := GenS3{F0: 65535}
            _ = fi50
            _ = v52
            _ = v53
        }
        fmt.Println(v11)
        if (!(!(!((v3 - (v10 + v3)) < ((v3 / 1e10) / v3))))) {
            var v54 float32 = (v10 / (3.14159 / 3.14159))
            fmt.Println(v10, v54)
            _ = v54
        } else {
            if false {
                for fi55 := 0; fi55 < 10; fi55++ {
                    for fi56 := 0; fi56 < 18; fi56++ {
                        fmt.Println(v11, v1)
                        _ = fi56
                    }
                    fmt.Println(v10, v1)
                    var v57 int8 = -92
                    _ = fi55
                    _ = v57
                }
                var v58 int16 = -1
                _ = v58
            } else {
                var v59 int64 = 1
                v59 = v59
                delete(v2, ("a" + "\n\t"))
                _ = v59
            }
            v3 += (v10 + 1.5)
            v11 = v11
        }
    }
    var v60 uint = v11
    switch "hello" {
    case "foo bar":
        for fi61 := 0; fi61 < 15; fi61++ {
            v62 := GenS2{F0: "\n\t", F1: 100}
            if (!false) {
                v10 = 0.1
            } else {
                var v63 uint16 = 65535
                v64 := GenS3{F0: (v12 << (((v11 | (v11 + v11))) & 63))}
                for fi65 := 0; fi65 < 19; fi65++ {
                    var v66 string = ("\n\t" + "hello")
                    v11 = (v11 << (((v60 + v11)) & 63))
                    var v67 float64 = 3.14159
                    _ = fi65
                    _ = v66
                    _ = v67
                }
                _ = v63
                _ = v64
            }
            _ = fi61
            _ = v62
        }
    case "hello":
        v68 := GenS4{F0: -1, F1: 42, F2: (!true)}
        _ = v68
    default:
        v60 = 2
    }
    switch 1 {
    case -569:
        v60 = (v60 & v11)
    }
    var v69 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v60:", v60)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v60
    _ = v69
}
