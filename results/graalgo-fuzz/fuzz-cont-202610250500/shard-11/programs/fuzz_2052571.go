package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := []uint32{2147483648, 2}
    var v2 float32 = (0.0 / -2.25)
    v3 := []string{("a" + ((("hello" + ("a" + "hello")) + "foo bar") + "a")), "a", ("hello" + ("a" + "")), ""}
    v4 := make(map[string]uint32)
    if false {
        var v5 int32 = 1
        fmt.Println(v2, v5)
        v6 := v4["foo bar"]
        _ = v5
        _ = v6
    }
    switch "foo bar" {
    case "foo bar":
        if (42 < -169) {
            v1 = append(v1, 65535)
            v2 += v2
            var v7 uint8 = 2
            _ = v7
        }
    case "hello":
        var v8 float64 = 0.0
        _ = v8
    default:
        v2 = float32(0.1)
    }
    fmt.Println(v2)
    if ((!true) || false) {
        v9 := GenS1{F0: 322}
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 float32 = -2.25
            _ = fi10
            _ = v11
        }
        v12 := GenS2{F0: -804, F1: 2, F2: 1}
        _ = v12
        _ = v9
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 uint8 = 30
                switch fi14 {
                case 116:
                    v16 := GenS1{F0: 2}
                    _ = v16
                default:
                    var v17 int64 = 9223372036854775807
                    _ = v17
                }
                fmt.Println(v2, fi14)
                _ = fi14
                _ = v15
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v4)
            _ = fi13
        }
    }
    _ = len(v4)
    var v18 uint16 = 42
    if (1 == 0) {
        if false {
            var v19 uint = 1
            v2 -= (v2 * 1e10)
            if true {
                if (128 > 1) {
                    v4["a"] = 1
                    if ((false && (!true)) || (!(42 > -9223372036854775808))) {
                        var v20 int16 = 216
                        var v21 uint8 = 2
                        v20 -= v20
                        _ = v20
                        _ = v21
                    } else {
                        if (!(!true)) {
                            if true {
                                if ((^(v18)) != v18) {
                                    v18 += v18
                                    for fi22 := 0; fi22 < 1; fi22++ {
                                        switch fi22 {
                                        case 42:
                                            fmt.Println(fi22, v18)
                                        }
                                        v18 = v18
                                        var v23 float32 = v2
                                        _ = fi22
                                        _ = v23
                                    }
                                }
                                var v24 uint = (v19 % ((652) | 1))
                                var v25 uint8 = 2
                                _ = v24
                                _ = v25
                            }
                        } else {
                            fmt.Println(v18, v2)
                            var v26 uint16 = (v18 & (v18 / ((v18) | 1)))
                            _ = v26
                        }
                        var v27 uint16 = 42
                        switch "" {
                        case "hello":
                            var v28 float64 = 3.14159
                            _ = v28
                        case "foo bar":
                            if true {
                                if ((float32(float64(0.1)) == float32(3.14159)) && true) {
                                    fmt.Println(v19)
                                    var v29 string = "hello"
                                    _ = v29
                                }
                            }
                        case "a":
                            var v30 int64 = int64(v27)
                            _ = v30
                        }
                        _ = v27
                    }
                } else {
                    for fi31 := 0; fi31 < 5; fi31++ {
                        fmt.Println(v19)
                        var v32 uint64 = 18446744073709551615
                        _ = len(v4)
                        _ = fi31
                        _ = v32
                    }
                }
                v18 += v18
            }
            _ = v19
        } else {
            for fi33 := 0; fi33 < 20; fi33++ {
                v2 = float32(1e10)
                _ = fi33
            }
            if (!(1 < int64(v18))) {
                var v34 string = "a"
                fmt.Println(v18)
                _ = v34
            } else {
                for fi35 := 0; fi35 < 3; fi35++ {
                    v4["hello"] = 0
                    if len(v3) > 0 {
                        v3[0] = (("foo bar" + ("foo bar" + ("" + "\n\t"))) + ("hello" + ""))
                    }
                    _ = fi35
                }
                if (false || ((0 != 32767) || ((!false) && (4294967295 != 573)))) {
                    fmt.Println(v2, v18)
                }
            }
            for fi36 := 0; fi36 < 3; fi36++ {
                v18 -= v18
                var v37 uint8 = 1
                v18 = (^(v18))
                _ = fi36
                _ = v37
            }
        }
        v18 -= v18
    } else {
        if (!false) {
            v38 := GenS3{F0: -102, F1: 100}
            _ = len(v1)
            _ = v38
        }
        v4[("foo bar" + "")] = 279
        for fi39 := 0; fi39 < 17; fi39++ {
            var v40 string = "\n\t"
            _ = len(v1)
            v18 -= v18
            _ = fi39
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        var v42 int8 = 127
        _ = fi41
        _ = v42
    }
    var v43 int64 = -1
    v2 = (v2 * 0.0)
    v44 := GenS4{F0: 46, F1: -1}
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v45 uint8 = 2
            _ = v45
        } else {
            for fi46 := 0; fi46 < 1; fi46++ {
                for fi47 := 0; fi47 < 9; fi47++ {
                    if true {
                        fmt.Println(v2, fi46)
                        v48 := GenS5{F0: 1}
                        _ = v48
                    }
                    _ = fi47
                }
                _ = fi46
            }
        }
    } else {
        var v49 int32 = 42
        _ = v49
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
}
