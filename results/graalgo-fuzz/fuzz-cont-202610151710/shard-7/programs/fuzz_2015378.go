package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int16; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 int16 }

func main() {
    var v1 int = 0
    var v2 float32 = 1.5
    var v3 uint16 = 255
    v4 := []int{(v1 ^ int(v3)), (^(v1)), -436, 42}
    if len(v4) > 0 {
        v4[0] = (v1 % ((828) | 1))
    }
    var v5 float32 = float32(float64(v2))
    _ = len(v4)
    switch v1 {
    case -9223372036854775808:
        if (false || true) {
            switch (v1 | v1) {
            case 127:
                for fi6 := 0; fi6 < 17; fi6++ {
                    for fi7 := 0; fi7 < 12; fi7++ {
                        fmt.Println(v1, v1, v1)
                        _ = fi7
                    }
                    var v8 uint = 582
                    switch (fi6 - 495) {
                    case 128:
                        if false {
                            fmt.Println(fi6, v1)
                            for fi9 := 0; fi9 < 13; fi9++ {
                                fmt.Println(v8, v2)
                                var v10 float32 = v2
                                switch "a" {
                                case "":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "foo bar":
                                    var v11 int32 = -2147483648
                                    _ = v11
                                }
                                if true {
                                    fmt.Println(v1, v8)
                                    v4 = append(v4, (^(fi9)))
                                }
                                _ = fi9
                                _ = v10
                            }
                            v2 = float32(float64(0.0))
                        } else {
                            if true {
                                var v12 uint64 = 1
                                if len(v4) > 0 {
                                    v4[0] = v1
                                }
                                var v13 float64 = 1e10
                                _ = v12
                                _ = v13
                            }
                        }
                    case 1:
                        v14 := GenS1{F0: 2147483647, F1: 58}
                        _ = v14
                    case 809:
                        fmt.Println(v2)
                    }
                    _ = fi6
                    _ = v8
                }
            case 32767:
                var v15 uint8 = 2
                _ = v15
            }
        }
    case 32767:
        for fi16 := 0; fi16 < 9; fi16++ {
            v17 := GenS2{F0: 32767, F1: 2147483648, F2: 2147483647}
            _ = fi16
            _ = v17
        }
    case 42:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        if true {
            for fi18 := 0; fi18 < 19; fi18++ {
                v5 += ((v2 + v2) + float32(1e10))
                for fi19 := 0; fi19 < 3; fi19++ {
                    if (!(false || true)) {
                        v5 -= 0.0
                        var v20 int = fi19
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v21 bool = (!((("foo bar" + ((("hello" + "") + (("" + "hello") + "hello")) + "")) > (("" + (("\n\t" + "a") + ("foo bar" + "\n\t"))) + "\n\t")) && false))
                        _ = v20
                        _ = v21
                    } else {
                        v22 := GenS3{F0: v3}
                        v2 = v5
                        _ = v22
                    }
                    var v23 uint32 = 0
                    _ = fi19
                    _ = v23
                }
                _ = fi18
            }
            for fi24 := 0; fi24 < 20; fi24++ {
                switch 128 {
                case 32768:
                    v25 := GenS1{F0: 139, F1: 71}
                    _ = v25
                case -1:
                    for fi26 := 0; fi26 < 2; fi26++ {
                        fmt.Println(fi26, v5)
                        if true {
                            switch (^(fi26)) {
                            case 32767:
                                if true {
                                    v27 := GenS4{F0: "hello"}
                                    var v28 int8 = 100
                                    _ = v27
                                    _ = v28
                                } else {
                                    if false {
                                        switch (fi26 >> ((276) & 63)) {
                                        case 1:
                                            v29 := GenS4{F0: "hello"}
                                            _ = v29
                                        case 722:
                                            var v30 int16 = 42
                                            _ = v30
                                        }
                                        if (69 <= 127) {
                                            var v31 uint64 = 4294967295
                                            var v32 int = fi26
                                            var v33 bool = (!("hello" == "hello"))
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                        } else {
                                            v34 := GenS5{F0: 65535, F1: 2}
                                            switch fi24 {
                                            case 128:
                                                switch ("a" + "\n\t") {
                                                case "\n\t":
                                                    v3 = (v3 - v3)
                                                case "a":
                                                    _ = len(v4)
                                                default:
                                                    fmt.Println(fi26, v5)
                                                }
                                            case 0:
                                                if (127 < -1) {
                                                    var v35 float64 = float64(v5)
                                                    v35 -= (v35 / (v35 + (v35 + float64(float32((3.14159 - ((0.0 - v35) - v35)))))))
                                                    v2 = float32(1.5)
                                                    _ = v35
                                                }
                                            }
                                            _ = v34
                                        }
                                        if ((2 != 1) && true) {
                                            var v36 string = "foo bar"
                                            _ = v36
                                        }
                                    } else {
                                        v5 += 1.5
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            case -2147483648:
                                var v37 uint32 = 2147483648
                                _ = v37
                            case -9223372036854775808:
                                v38 := GenS6{F0: (float64(((0.1 / (v2 + (v2 + v2))) - v2)) >= 3.14159), F1: (fi26 << ((255) & 63))}
                                _ = v38
                            default:
                                v3 = (v3 << ((882) & 63))
                            }
                            v39 := GenS7{F0: 1000}
                            _ = v39
                        }
                        _ = fi26
                    }
                }
                _ = fi24
            }
        } else {
            switch "" {
            case "hello":
                if (849 < 1) {
                    fmt.Println(v5)
                    v5 += 0.0
                }
            case "\n\t":
                fmt.Println(v1, v5)
            default:
                v40 := GenS5{F0: uint(v3), F1: v3}
                _ = v40
            }
            for fi41 := 0; fi41 < 4; fi41++ {
                if len(v4) > 0 {
                    v4[0] = fi41
                }
                _ = fi41
            }
        }
    }
    switch (("hello" + "") + "foo bar") {
    case "foo bar":
        if ((uint8(v3) == 2) && (!true)) {
            for fi42 := 0; fi42 < 7; fi42++ {
                fmt.Println(v5, v5)
                v2 += v5
                if (((1e10 != ((3.14159 / 1.0) + v5)) || false) || true) {
                    for fi43 := 0; fi43 < 13; fi43++ {
                        _ = len(v4)
                        _ = fi43
                    }
                    var v44 float32 = (v2 + 0.1)
                    var v45 uint32 = 42
                    _ = v44
                    _ = v45
                } else {
                    v2 += (v5 * 1e10)
                    var v46 uint8 = 3
                    fmt.Println(v3)
                    _ = v46
                }
                _ = fi42
            }
            var v47 string = ""
            _ = v47
        }
    }
    if len(v4) > 0 {
        v4[0] = (^(v1))
    }
    var v48 uint64 = 660
    v48 = (v48 * v48)
    v4 = append(v4, -2147483648)
    for fi49 := 0; fi49 < 16; fi49++ {
        var v50 string = "a"
        v3 -= v3
        _ = len(v4)
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
}
