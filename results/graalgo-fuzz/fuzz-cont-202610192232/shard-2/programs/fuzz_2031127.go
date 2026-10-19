package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint32; F1 uint16; F2 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 int16; F1 int }
func genFn0(p0 uint16, p1 uint16) string {
    for fi6 := 0; fi6 < 11; fi6++ {
        p0 = (p0 % ((255) | 1))
        var v7 uint32 = 42
        v8 := []int{fi6, fi6, (fi6 % ((fi6) | 1))}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        p0 -= p0
        v10 := GenS1{F0: 127}
        _ = fi9
        _ = v10
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = v11
    return ("hello" + "")
}

func main() {
    v1 := []int{}
    var v2 int64 = 1
    var v3 uint8 = 218
    v4 := make(map[string]string)
    v5 := []string{""}
    var v12 uint32 = 965
    var v13 int64 = v2
    v14 := GenS1{F0: v3}
    var v15 int = 128
    var v16 uint64 = 4294967295
    v15 -= v15
    switch 332 {
    case 127:
        var v17 int16 = -69
        _ = v17
    case 225:
        v18 := GenS1{F0: 145}
        _ = v18
    case 42:
        for fi19 := 0; fi19 < 19; fi19++ {
            if false {
                switch 32768 {
                case -2147483648:
                    v20 := GenS1{F0: 1}
                    _ = v20
                case 32768:
                    if false {
                        for fi21 := 0; fi21 < 17; fi21++ {
                            v22 := GenS1{F0: v3}
                            _ = fi21
                            _ = v22
                        }
                        switch 2147483647 {
                        case 128:
                            _ = len(v4)
                        case 9223372036854775807:
                            v23 := GenS2{F0: 65535, F1: 4294967295}
                            _ = v23
                        case -2147483648:
                            fmt.Println(fi19, v12, v3)
                        }
                    }
                }
                for fi24 := 0; fi24 < 17; fi24++ {
                    var v25 uint = uint(v3)
                    fmt.Println(v3, v2)
                    fmt.Println(v2, v16, v2)
                    _ = fi24
                    _ = v25
                }
                v26 := GenS3{F0: -835, F1: (v3 - 128), F2: (v12 / ((v12) | 1))}
                _ = v26
            }
            v4["hello"] = ("hello" + "")
            switch (fi19 & 2147483647) {
            case -423:
                var v27 float64 = 0.0
                _ = v27
            case 1:
                var v28 uint32 = (v12 / ((v12) | 1))
                _ = v28
            case 4294967295:
                v15 -= 194
            }
            _ = fi19
        }
    default:
        v12 += (^(v12))
    }
    v16 -= (v16 * (v16 % ((74) | 1)))
    v2 += (v13 / ((v2) | 1))
    for fi29 := 0; fi29 < 9; fi29++ {
        for fi30 := 0; fi30 < 18; fi30++ {
            fmt.Println(v15, v13)
            var v31 int32 = 0
            _ = len(v4)
            _ = fi30
            _ = v31
        }
        v3 = (v3 & v3)
        fmt.Println(v15, v16)
        _ = fi29
    }
    if false {
        v32 := GenS2{F0: v12, F1: 255}
        _ = v32
    } else {
        if (32767 != 32767) {
            for fi33 := 0; fi33 < 5; fi33++ {
                delete(v4, "foo bar")
                var v34 int16 = int16(v12)
                _ = fi33
                _ = v34
            }
            for fi35 := 0; fi35 < 13; fi35++ {
                v36 := GenS4{F0: 880, F1: 273, F2: "hello"}
                for fi37 := 0; fi37 < 5; fi37++ {
                    v2 = v13
                    v38 := GenS1{F0: v3}
                    _ = fi37
                    _ = v38
                }
                if false {
                    fmt.Println(v16)
                    v15 -= v15
                    switch ("foo bar" + ("\n\t" + ("\n\t" + "hello"))) {
                    case "foo bar":
                        if false {
                            v16 = (v16 & v16)
                        }
                    case "\n\t":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case "":
                        v4[""] = ""
                    default:
                        if ((((0.0 * float64(3.14159)) * (1.5 + 1e10)) * 1e10) == 3.14159) {
                            fmt.Println(v12, v13)
                            if (v12 <= (v12 ^ (v12 ^ (^(v12))))) {
                                v39 := GenS2{F0: 863, F1: 1}
                                var v40 int = 2147483647
                                _ = v39
                                _ = v40
                            } else {
                                switch ("foo bar" + ("\n\t" + ("foo bar" + ("foo bar" + ("a" + ""))))) {
                                case "foo bar":
                                    for fi41 := 0; fi41 < 2; fi41++ {
                                        v42 := GenS5{F0: (^(v16))}
                                        _ = fi41
                                        _ = v42
                                    }
                                case "":
                                    v43 := GenS2{F0: (^(v12)), F1: 65535}
                                    _ = v43
                                default:
                                    var v44 uint64 = 338
                                    _ = v44
                                }
                            }
                            var v45 uint = 309
                            _ = v45
                        } else {
                            for fi46 := 0; fi46 < 8; fi46++ {
                                v13 = int64(v12)
                                v47 := GenS6{F0: 42, F1: 1}
                                _ = fi46
                                _ = v47
                            }
                        }
                    }
                }
                _ = fi35
                _ = v36
            }
            fmt.Println(v16, v13)
        }
        if (true || false) {
            switch "" {
            case "\n\t":
                if false {
                    var v48 int32 = 0
                    _ = v48
                }
            case "a":
                var v49 int8 = -6
                _ = v49
            case "foo bar":
                v13 -= v2
            }
        }
    }
    var v50 uint64 = 2
    for fi51 := 0; fi51 < 19; fi51++ {
        fmt.Println(v50, v50)
        var v52 int8 = -18
        _ = fi51
        _ = v52
    }
    v13 += (v13 ^ 36)
    _ = len(v4)
    switch (("foo bar" + "a") + "hello") {
    case "a":
        v53 := GenS5{F0: v16}
        _ = v53
    default:
        v54 := GenS7{F0: -1, F1: (v15 ^ 127)}
        _ = v54
    }
    for fi55 := 0; fi55 < 16; fi55++ {
        v15 -= v15
        v2 += int64(v16)
        for fi56 := 0; fi56 < 4; fi56++ {
            v57 := GenS4{F0: 42, F1: 54, F2: ""}
            if true {
                switch "foo bar" {
                case "\n\t":
                    var v58 uint8 = (v3 % (((v3 % ((255) | 1))) | 1))
                    _ = v58
                case "a":
                    var v59 string = ((("hello" + "foo bar") + "a") + "foo bar")
                    _ = v59
                }
            } else {
                v12 += 42
            }
            if (v3 >= v3) {
                v60 := GenS7{F0: 32767, F1: (-(fi55))}
                if false {
                    v61 := GenS3{F0: -750, F1: (v3 / ((42) | 1)), F2: v12}
                    _ = v61
                }
                _ = v60
            }
            fmt.Println(fi56, v2, v12)
            _ = fi56
            _ = v57
        }
        _ = fi55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
}
