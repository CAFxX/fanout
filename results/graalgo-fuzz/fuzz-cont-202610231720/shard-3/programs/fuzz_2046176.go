package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint16 }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]uint32)
    v4 := []string{}
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 string = ("hello" + ("" + "foo bar"))
        if true {
            if (393 != 18446744073709551615) {
                var v7 uint = 966
                switch (^(fi5)) {
                case 709:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 2147483647:
                    v2 += -9223372036854775808
                default:
                    fmt.Println(v6, v7)
                }
                _ = v7
            } else {
                v8 := GenS1{F0: 928, F1: (fi5 ^ (fi5 * (fi5 * (-(fi5))))), F2: 127}
                var v9 int8 = 42
                _ = v8
                _ = v9
            }
            v6 = (v6 + (v6 + v6))
            v6 = (v6 + "\n\t")
        }
        _ = fi5
        _ = v6
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            v2 = 42
            switch -615 {
            case -941:
                if false {
                    var v12 int = (fi10 - fi11)
                    var v13 bool = (!false)
                    if (true && v13) {
                        v4 = append(v4, (("foo bar" + "foo bar") + "a"))
                        fmt.Println(fi11)
                        if (false || v13) {
                            v14 := GenS1{F0: 1, F1: (fi10 ^ 32767), F2: 150}
                            var v15 uint16 = 32767
                            v16 := v1[(fi10 + (fi10 * (fi11 - -534)))]
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    } else {
                        v17 := GenS2{F0: -9223372036854775808}
                        var v18 uint32 = 2147483648
                        var v19 uint64 = 833
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                    _ = v12
                    _ = v13
                }
            case 309:
                v2 -= (v2 - v2)
            case 32767:
                var v20 int16 = 1
                _ = v20
            }
            _ = fi11
        }
        _ = fi10
    }
    switch -1 {
    case 127:
        fmt.Println(v2)
    case -559:
        var v21 uint64 = 1
        _ = v21
    }
    if true {
        var v22 int32 = -701
        if (!false) {
            v23 := GenS1{F0: 65535, F1: 3, F2: 128}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v23
        } else {
            v24 := GenS2{F0: (v2 / (((v2 * 1)) | 1))}
            switch 9223372036854775807 {
            case 4294967295:
                switch 540 {
                case 9223372036854775807:
                    if false {
                        var v25 int16 = -1
                        _ = v25
                    }
                case 32768:
                    if (true && true) {
                        delete(v3, "foo bar")
                        fmt.Println(v2)
                    } else {
                        var v26 int = 128
                        if false {
                            var v27 float32 = 0.1
                            v28 := GenS3{F0: 128, F1: 555}
                            _ = v27
                            _ = v28
                        } else {
                            for fi29 := 0; fi29 < 4; fi29++ {
                                v2 = v2
                                fmt.Println(v26, v2, fi29)
                                _ = fi29
                            }
                            _ = len(v3)
                            var v30 uint = 4294967295
                            _ = v30
                        }
                        for fi31 := 0; fi31 < 12; fi31++ {
                            if (2147483647 < (v22 - (v22 | v22))) {
                                v22 = (v22 - (v22 % ((int32(v2)) | 1)))
                            }
                            v4 = append(v4, "\n\t")
                            if ((float32(-2.25) - 1.5) == (0.1 - (-2.25 / 1.0))) {
                                _ = len(v3)
                                fmt.Println(v22, v26, v22)
                            }
                            _ = fi31
                        }
                        _ = v26
                    }
                case 42:
                    if ((653 > 204) || false) {
                        if false {
                            delete(v1, 2147483647)
                        }
                    } else {
                        for fi32 := 0; fi32 < 15; fi32++ {
                            v3[("" + ("\n\t" + ""))] = 0
                            v33 := GenS4{F0: 2}
                            v34 := v3["foo bar"]
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                    }
                default:
                    var v35 float32 = (1.5 / 1e10)
                    _ = v35
                }
            case 9223372036854775807:
                for fi36 := 0; fi36 < 18; fi36++ {
                    var v37 int16 = 42
                    _ = fi36
                    _ = v37
                }
            case 1:
                var v38 float32 = 0.1
                _ = v38
            }
            _ = v24
        }
        _ = v22
    }
    if false {
        var v39 uint8 = 0
        _ = v39
    }
    var v40 int64 = v2
    v40 -= v2
    v2 = v2
    switch ("\n\t" + "hello") {
    case "hello":
        fmt.Println(v2)
    }
    if (42 <= 1000) {
        v41 := GenS5{F0: false}
        if len(v4) > 0 {
            v4[0] = "a"
        }
        var v42 int16 = 486
        _ = v41
        _ = v42
    }
    var v43 float64 = ((3.14159 / 1.5) / 1e10)
    v40 += 9223372036854775807
    v44 := GenS6{F0: "", F1: 32767}
    v45 := GenS1{F0: 2147483647, F1: 0, F2: 187}
    var v46 int64 = (v2 + v40)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
