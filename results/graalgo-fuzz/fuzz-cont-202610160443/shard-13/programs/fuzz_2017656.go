package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := []int{}
    var v3 bool = (2 > 42)
    v4 := make(map[string]string)
    var v5 int64 = 42
    var v6 uint8 = 127
    v6 = (^(v6))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        if (!(42 <= 0)) {
            if (18446744073709551615 < 909) {
                var v8 string = ("" + "foo bar")
                v9 := v1[fi7]
                _ = v8
                _ = v9
            } else {
                v6 = v6
            }
        }
        var v10 float64 = (-2.25 + 1e10)
        fmt.Println(v10, v6, v6)
        _ = fi7
        _ = v10
    }
    fmt.Println(v5, v5)
    switch 0 {
    case 235:
        var v11 int32 = 42
        _ = v11
    case 127:
        fmt.Println(v5, v5)
    case 243:
        switch ("\n\t" + ("\n\t" + (("" + "\n\t") + (("hello" + "foo bar") + "a")))) {
        case "a":
            for fi12 := 0; fi12 < 9; fi12++ {
                if v3 {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v14 := GenS1{F0: 100}
                        switch "hello" {
                        case "\n\t":
                            switch ("a" + "foo bar") {
                            case "hello":
                                v6 -= (v6 << ((4294967295) & 63))
                            }
                        default:
                            delete(v1, -708)
                        }
                        if v3 {
                            var v15 int64 = (-(v5))
                            switch (-(fi13)) {
                            case 32767:
                                v16 := GenS2{F0: "a"}
                                _ = v16
                            default:
                                var v17 int64 = (v5 & 202)
                                _ = v17
                            }
                            _ = v15
                        }
                        _ = fi13
                        _ = v14
                    }
                    v6 = (v6 / (((v6 + v6)) | 1))
                } else {
                    if len(v2) > 0 {
                        v2[0] = (fi12 - fi12)
                    }
                    switch (("\n\t" + ("\n\t" + "")) + "a") {
                    case "hello":
                        var v18 uint64 = 644
                        _ = v18
                    default:
                        fmt.Println(v6)
                    }
                }
                _ = fi12
            }
        case "hello":
            fmt.Println(v3, v3)
        default:
            for fi19 := 0; fi19 < 1; fi19++ {
                switch "" {
                case "\n\t":
                    var v20 uint64 = 246
                    _ = v20
                case "a":
                    v3 = v3
                case "foo bar":
                    v3 = v3
                }
                for fi21 := 0; fi21 < 3; fi21++ {
                    if len(v2) > 0 {
                        v2[0] = (fi19 << ((0) & 63))
                    }
                    var v22 uint16 = 911
                    if (false || ((v6 + 1) == (^(v6)))) {
                        v23 := v1[127]
                        var v24 uint8 = 128
                        v1[700] = ""
                        _ = v23
                        _ = v24
                    }
                    v3 = false
                    if (0 != 0) {
                        if ((((v5 | (^(v5))) != -1) && (false || true)) && (-128 >= -84)) {
                            if (!v3) {
                                if v3 {
                                    switch 0 {
                                    case -2147483648:
                                        delete(v1, (fi21 & (-(fi19))))
                                    }
                                    v25 := GenS3{F0: fi21, F1: 2147483647}
                                    _ = v25
                                }
                            }
                            var v26 uint16 = v22
                            v27 := GenS4{F0: ("\n\t" + ("a" + "foo bar"))}
                            _ = v26
                            _ = v27
                        }
                        v5 = -1
                        var v28 uint32 = 65535
                        _ = v28
                    } else {
                        fmt.Println(v5)
                        switch (fi21 << ((4294967295) & 63)) {
                        case 9223372036854775807:
                            v3 = v3
                        case -2147483648:
                            v6 = 1
                        default:
                            if ((v22 - (v22 >> ((727) & 63))) > v22) {
                                v5 -= (v5 | v5)
                                if (1e10 >= float32(float64(-2.25))) {
                                    if v3 {
                                        var v29 float64 = 3.14159
                                        _ = v29
                                    } else {
                                        if (42 < -751) {
                                            v30 := GenS5{F0: v5}
                                            v31 := GenS3{F0: (-(fi19)), F1: -141}
                                            if (("hello" + "foo bar") <= "\n\t") {
                                                v2 = append(v2, (-(fi19)))
                                                var v32 int64 = (v5 ^ (v5 + 9223372036854775807))
                                                _ = v32
                                            } else {
                                                var v33 float32 = (3.14159 - -2.25)
                                                _ = v33
                                            }
                                            _ = v30
                                            _ = v31
                                        }
                                        v22 -= 0
                                    }
                                    v5 += 42
                                    v34 := GenS5{F0: v5}
                                    _ = v34
                                }
                                var v35 string = "hello"
                                _ = v35
                            }
                        }
                    }
                    if ((fi21 & fi21) >= (fi19 ^ 128)) {
                        var v36 int64 = 9223372036854775807
                        v4["\n\t"] = ""
                        v22 = (v22 % ((v22) | 1))
                        _ = v36
                    }
                    _ = fi21
                    _ = v22
                }
                if true {
                    v3 = (42 > 1)
                    var v37 uint32 = 2147483648
                    _ = v37
                } else {
                    v2 = append(v2, (fi19 - fi19))
                }
                _ = fi19
            }
        }
    default:
        for fi38 := 0; fi38 < 15; fi38++ {
            v5 -= (v5 / (((^(v5))) | 1))
            if (1 == 2147483647) {
                v39 := GenS2{F0: "hello"}
                var v40 int = (fi38 << ((uint(v6)) & 63))
                _ = v39
                _ = v40
            }
            _ = fi38
        }
    }
    var v41 float32 = float32((3.14159 / 1.0))
    v42 := v1[-851]
    for fi43 := 0; fi43 < 11; fi43++ {
        if false {
            switch "a" {
            case "foo bar":
                var v44 int = (-(fi43))
                _ = v44
            default:
                if ((((1.5 - 1.5) - 3.14159) - (((3.14159 / 0.1) / 0.1) - float64(((v41 + v41) / (0.1 / 1.0))))) > (1e10 - 1.5)) {
                    var v45 string = "foo bar"
                    v2 = append(v2, 437)
                    _ = v45
                }
            }
            var v46 int8 = int8(v5)
            for fi47 := 0; fi47 < 4; fi47++ {
                var v48 uint8 = (^(v6))
                if (-97 < 127) {
                    v3 = v3
                } else {
                    if v3 {
                        delete(v1, (fi43 - fi47))
                    } else {
                        var v49 int32 = 2147483647
                        var v50 bool = v3
                        var v51 uint64 = 381
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                }
                v5 = v5
                _ = fi47
                _ = v48
            }
            _ = v46
        }
        v3 = v3
        if ((2 != 0) || false) {
            for fi52 := 0; fi52 < 14; fi52++ {
                v41 -= v41
                v3 = true
                switch 4294967295 {
                case 2147483647:
                    switch v42 {
                    case "a":
                        v42 = v42
                    case "":
                        var v53 uint32 = 678
                        _ = v53
                    case "hello":
                        var v54 int32 = -1
                        _ = v54
                    default:
                        v42 = (((v42 + (v42 + ("\n\t" + ("hello" + "a")))) + v42) + v42)
                    }
                case -9223372036854775808:
                    var v55 string = v42
                    _ = v55
                case -1:
                    v41 += 0.0
                default:
                    var v56 string = v42
                    _ = v56
                }
                _ = fi52
            }
        }
        _ = fi43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
