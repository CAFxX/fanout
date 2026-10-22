package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int; F2 int }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int; F1 uint32; F2 int64 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint; F1 uint16; F2 int }
type GenS8 struct { F0 bool; F1 uint64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (42 != 2147483647)
    var v3 bool = (1 <= 309)
    v4 := []uint32{1, 209, 2147483648, 368}
    v5 := []uint32{430}
    v3 = (-2.25 == (3.14159 - ((-2.25 / 0.1) - 0.1)))
    v6 := GenS1{F0: -914}
    switch "a" {
    case "\n\t":
        switch 42 {
        case 32768:
            var v7 int16 = 660
            _ = v7
        case 561:
            v8 := GenS2{F0: (489 >= -755), F1: 127, F2: -9223372036854775808}
            _ = v8
        }
    case "a":
        v9 := v1["\n\t"]
        _ = v9
    case "foo bar":
        v2 = (128 != 52)
    }
    v3 = (float32(0.0) < 0.0)
    switch 32768 {
    case 32768:
        var v10 uint = 65535
        _ = v10
    case 617:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        var v11 uint16 = 255
        _ = v11
    }
    v12 := v1["foo bar"]
    var v13 float64 = (1e10 / (0.1 / 1e10))
    var v14 float64 = (v13 / v13)
    switch ("a" + "\n\t") {
    case "":
        var v15 uint32 = 223
        _ = v15
    case "hello":
        v14 -= (v13 / float64(((((3.14159 / (-2.25 - ((float32(v13) * 0.1) / 1e10))) - 1e10) / (float32(float64(1.5)) * -2.25)) + 1.5)))
    case "\n\t":
        if (!v3) {
            var v16 string = "foo bar"
            _ = v16
        }
    default:
        v2 = v3
    }
    v13 = (v14 / v14)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        if (-1 < -9223372036854775808) {
            switch "\n\t" {
            case "foo bar":
                if v2 {
                    for fi18 := 0; fi18 < 16; fi18++ {
                        if (!((-2.25 + (1e10 * ((1e10 + 0.1) + 0.0))) >= 1e10)) {
                            var v19 uint16 = 2
                            fmt.Println(fi17, v14, fi18)
                            if ((255 >= 255) || (v14 == 3.14159)) {
                                var v20 float64 = 1.5
                                _ = v20
                            }
                            _ = v19
                        }
                        v3 = v2
                        _ = fi18
                    }
                    v14 += (v14 / 1.0)
                }
            case "a":
                switch ("\n\t" + "a") {
                case "hello":
                    v21 := GenS3{F0: 147, F1: "hello"}
                    _ = v21
                case "":
                    switch "\n\t" {
                    case "foo bar":
                        var v22 int64 = -9223372036854775808
                        _ = v22
                    default:
                        for fi23 := 0; fi23 < 4; fi23++ {
                            v14 = 3.14159
                            var v24 float32 = 0.0
                            _ = fi23
                            _ = v24
                        }
                    }
                case "foo bar":
                    v3 = ("hello" > "a")
                default:
                    v12 = 42
                }
            case "\n\t":
                if false {
                    switch (fi17 ^ (fi17 % ((1) | 1))) {
                    case -306:
                        v25 := GenS3{F0: 2, F1: ("\n\t" + "hello")}
                        _ = v25
                    }
                }
            }
        } else {
            var v26 int8 = 42
            if (1 == 255) {
                switch 2147483647 {
                case -2147483648:
                    if (-1 == 42) {
                        var v27 bool = false
                        v26 = (v26 % ((v26) | 1))
                        _ = v27
                    } else {
                        var v28 uint = 65535
                        if (127 >= 41) {
                            var v29 int64 = -9223372036854775808
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 15; fi30++ {
                            v31 := GenS4{F0: fi17, F1: (v12 + v12), F2: 863}
                            var v32 int64 = -9223372036854775808
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        _ = v28
                    }
                default:
                    for fi33 := 0; fi33 < 7; fi33++ {
                        if false {
                            var v34 int16 = -802
                            if (!true) {
                                v35 := GenS3{F0: 42, F1: ("\n\t" + "\n\t")}
                                switch ("foo bar" + (("" + "hello") + ((("foo bar" + "\n\t") + ("hello" + "foo bar")) + "a"))) {
                                case "a":
                                    var v36 int8 = (v26 / (((v26 / (((-(v26))) | 1))) | 1))
                                    _ = v36
                                case "":
                                    v26 += 1
                                    v37 := GenS1{F0: (fi33 ^ fi33)}
                                    _ = v37
                                case "foo bar":
                                    var v38 int16 = 746
                                    _ = v38
                                }
                                delete(v1, "hello")
                                var v39 uint64 = 170
                                _ = v35
                                _ = v39
                            }
                            v1[(("a" + ((("foo bar" + ("\n\t" + "")) + "hello") + "")) + "a")] = 2
                            _ = v34
                        }
                        var v40 uint = 0
                        _ = fi33
                        _ = v40
                    }
                }
                v41 := GenS3{F0: 128, F1: ("a" + "\n\t")}
                v42 := v1[("" + ("\n\t" + "hello"))]
                _ = v41
                _ = v42
            } else {
                for fi43 := 0; fi43 < 3; fi43++ {
                    v13 = 0.1
                    switch fi43 {
                    case 127:
                        v44 := GenS5{F0: "\n\t", F1: -9223372036854775808}
                        _ = v44
                    case 1:
                        v45 := GenS6{F0: -2147483648}
                        _ = v45
                    }
                    fmt.Println(v3, fi17, v26)
                    var v46 bool = v2
                    v12 += 248
                    _ = fi43
                    _ = v46
                }
            }
            var v47 int8 = (^(v26))
            _ = v26
            _ = v47
        }
        v48 := GenS7{F0: 18446744073709551615, F1: 1, F2: 32767}
        for fi49 := 0; fi49 < 16; fi49++ {
            switch -2147483648 {
            case 4294967295:
                if false {
                    switch "foo bar" {
                    case "hello":
                        v14 += (float64(1.5) / v14)
                    case "foo bar":
                        switch "" {
                        case "\n\t":
                            v50 := GenS8{F0: true, F1: 42}
                            _ = v50
                        default:
                            var v51 uint32 = 115
                            _ = v51
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case -764:
                fmt.Println(v13)
            }
            v3 = v2
            if ((46 < 1) || false) {
                if (-9223372036854775808 > -258) {
                    v2 = false
                    fmt.Println(v12, v12)
                }
                v2 = v2
            }
            _ = fi49
        }
        _ = fi17
        _ = v48
    }
    fmt.Println(v12)
    if (-497 < -32768) {
        if (v3 && (1e10 <= float32(v14))) {
            v52 := GenS2{F0: true, F1: 9223372036854775807, F2: -2147483648}
            _ = v52
        }
        v14 = (((v13 / v14) + float64((0.1 * 0.0))) - (v14 / v13))
    } else {
        for fi53 := 0; fi53 < 10; fi53++ {
            var v54 int64 = 9223372036854775807
            if v3 {
                var v55 string = "foo bar"
                _ = v55
            } else {
                var v56 uint32 = (^(v12))
                _ = v56
            }
            _ = fi53
            _ = v54
        }
    }
    var v57 float32 = -2.25
    v2 = v3
    fmt.Println(v57, v2)
    var v58 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
