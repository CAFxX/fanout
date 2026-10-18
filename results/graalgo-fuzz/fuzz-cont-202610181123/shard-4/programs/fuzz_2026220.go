package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 float64 = (0.0 - 0.1)
    v2 := []int{4294967295, -2147483648, -2147483648}
    v3 := make(map[int]string)
    switch ("a" + ("foo bar" + "foo bar")) {
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
    case "a":
        v3[32767] = "hello"
    }
    var v4 int8 = 55
    v1 -= (float64((1e10 / -2.25)) + v1)
    if (-314 == -1) {
        if (0 > -32768) {
            v3[128] = ""
            v4 = (v4 << ((0) & 63))
            switch "\n\t" {
            case "\n\t":
                var v5 int8 = v4
                _ = v5
            case "hello":
                v2 = append(v2, -17)
            }
        }
        v4 += (v4 << ((18446744073709551615) & 63))
        if false {
            var v6 bool = (v1 > float64((1e10 / (3.14159 + (0.0 * 1e10)))))
            v2 = append(v2, 4294967295)
            var v7 uint64 = 887
            _ = v6
            _ = v7
        } else {
            v1 += (1e10 + float64(-2.25))
            fmt.Println(v1, v4)
        }
    } else {
        var v8 uint8 = 255
        switch -644 {
        case 32768:
            switch ("hello" + "hello") {
            case "foo bar":
                switch "" {
                case "foo bar":
                    _ = len(v3)
                }
            case "":
                v8 = (v8 ^ 42)
            case "hello":
                if (!(float64(0.1) > v1)) {
                    var v9 int = 32768
                    switch (^(v9)) {
                    case -1:
                        for fi10 := 0; fi10 < 4; fi10++ {
                            var v11 int32 = 0
                            v12 := GenS1{F0: (fi10 % ((int(v11)) | 1)), F1: 117}
                            var v13 bool = true
                            _ = fi10
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                    case 32767:
                        var v14 uint32 = 765
                        _ = v14
                    }
                    v15 := v3[-80]
                    _ = v15
                    _ = v9
                } else {
                    var v16 bool = (!false)
                    _ = v16
                }
            default:
                var v17 int16 = 38
                _ = v17
            }
        default:
            v1 -= (((v1 - (v1 + (v1 - v1))) * (0.0 * v1)) * (float64(1.5) * v1))
        }
        _ = v8
    }
    if (false || false) {
        for fi18 := 0; fi18 < 14; fi18++ {
            if false {
                for fi19 := 0; fi19 < 8; fi19++ {
                    var v20 int64 = 1
                    v21 := v3[fi19]
                    var v22 int64 = -872
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                switch -1 {
                case 32767:
                    v4 -= (v4 % ((v4) | 1))
                default:
                    for fi23 := 0; fi23 < 4; fi23++ {
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        switch "hello" {
                        case "hello":
                            fmt.Println(fi23)
                            switch (fi18 >> ((786) & 63)) {
                            case -668:
                                var v24 bool = true
                                _ = v24
                            case 32767:
                                fmt.Println(fi23)
                            case -237:
                                v4 = -109
                            default:
                                v25 := GenS2{F0: -9223372036854775808}
                                _ = v25
                            }
                            if (v4 > (v4 - v4)) {
                                v1 = -2.25
                                if (777 < 1) {
                                    var v26 float64 = v1
                                    var v27 int = (fi23 ^ (-(fi23)))
                                    _ = v26
                                    _ = v27
                                }
                                v4 = 75
                            } else {
                                v1 += v1
                                v4 = (v4 << ((4294967295) & 63))
                                v4 = 127
                            }
                        case "\n\t":
                            v2 = append(v2, int(v4))
                        case "foo bar":
                            v28 := GenS2{F0: 617}
                            _ = v28
                        default:
                            v29 := v3[127]
                            _ = v29
                        }
                        _ = fi23
                    }
                }
            }
            _ = fi18
        }
    } else {
        if false {
            v30 := GenS2{F0: -220}
            v31 := GenS3{F0: 748}
            v1 += v1
            _ = v30
            _ = v31
        } else {
            fmt.Println(v1, v1)
            v4 = -123
            if (true || (true || true)) {
                v1 -= (v1 + v1)
                for fi32 := 0; fi32 < 3; fi32++ {
                    var v33 int32 = -969
                    for fi34 := 0; fi34 < 15; fi34++ {
                        v33 += (-(v33))
                        v33 = (v33 << ((2) & 63))
                        _ = fi34
                    }
                    _ = fi32
                    _ = v33
                }
            } else {
                var v35 uint16 = 1
                v1 -= v1
                var v36 int32 = -2147483648
                _ = v35
                _ = v36
            }
        }
        var v37 int8 = 42
        _ = v37
    }
    v38 := GenS1{F0: -362, F1: v4}
    var v39 int16 = 1000
    v4 += -80
    v3[286] = ("hello" + "")
    var v40 float64 = 3.14159
    if false {
        v1 = -2.25
    }
    var v41 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
