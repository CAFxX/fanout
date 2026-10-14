package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 int32 }

func main() {
    var v1 float32 = ((0.0 / ((1.5 * 1e10) * 3.14159)) + 0.1)
    v2 := []string{}
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    var v5 float32 = v1
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := v4[-9223372036854775808]
        _ = len(v4)
        if (v7 || ((^(fi6)) <= 767)) {
            v1 = v1
            v8 := GenS1{F0: true}
            v2 = append(v2, "\n\t")
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        v5 -= float32((1.5 / 1.0))
        _ = fi9
    }
    v2 = append(v2, "\n\t")
    v10 := GenS1{F0: true}
    for fi11 := 0; fi11 < 1; fi11++ {
        v12 := GenS2{F0: 0, F1: true}
        v5 -= (v5 + 3.14159)
        switch fi11 {
        case 356:
            _ = len(v2)
        case 351:
            switch (fi11 / (((fi11 * 32767)) | 1)) {
            case 812:
                for fi13 := 0; fi13 < 4; fi13++ {
                    switch "a" {
                    case "\n\t":
                        fmt.Println(fi13, fi11, fi11)
                    case "hello":
                        if (0 <= 1000) {
                            fmt.Println(v1)
                            v1 = (v1 / v1)
                        }
                    default:
                        var v14 float32 = 0.0
                        _ = v14
                    }
                    var v15 uint8 = 255
                    var v16 uint8 = (^(v15))
                    _ = fi13
                    _ = v15
                    _ = v16
                }
            case 32768:
                v1 -= 3.14159
            }
        }
        _ = fi11
        _ = v12
    }
    var v17 uint16 = 2
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 bool = (!true)
        if true {
            v19 = (!v19)
        } else {
            if v19 {
                var v20 bool = (65535 >= 2)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v20
            } else {
                v21 := GenS3{F0: 1}
                switch fi18 {
                case 42:
                    var v22 string = ("foo bar" + (("hello" + "hello") + ("hello" + ((("hello" + "a") + "a") + "foo bar"))))
                    _ = v22
                }
                v5 -= float32((-2.25 / 3.14159))
                _ = v21
            }
        }
        _ = fi18
        _ = v19
    }
    switch ("\n\t" + "hello") {
    case "hello":
        var v23 string = "a"
        _ = v23
    case "":
        switch -9223372036854775808 {
        case 4294967295:
            fmt.Println(v1)
        default:
            v5 += 1e10
        }
    case "\n\t":
        if true {
            var v24 uint32 = 152
            switch ("" + "a") {
            case "hello":
                switch "a" {
                case "a":
                    switch ("" + "hello") {
                    case "":
                        if (-9223372036854775808 <= 359) {
                            if true {
                                v25 := GenS4{F0: 0}
                                _ = v25
                            }
                            v26 := GenS3{F0: 32767}
                            fmt.Println(v5, v17)
                            _ = v26
                        } else {
                            var v27 int16 = -1
                            var v28 int32 = 2147483647
                            _ = v27
                            _ = v28
                        }
                    case "a":
                        var v29 int32 = -2147483648
                        _ = v29
                    case "\n\t":
                        for fi30 := 0; fi30 < 10; fi30++ {
                            switch fi30 {
                            case 152:
                                fmt.Println(v17, v1)
                            case 4294967295:
                                if len(v2) > 0 {
                                    v2[0] = ("\n\t" + "")
                                }
                            }
                            v24 = (v24 >> ((65535) & 63))
                            if (false || ((false || (-1 < -1)) && (fi30 >= 128))) {
                                switch "a" {
                                case "\n\t":
                                    for fi31 := 0; fi31 < 16; fi31++ {
                                        v5 -= v1
                                        v24 += 2147483647
                                        v2 = append(v2, ("" + ("hello" + "hello")))
                                        _ = fi31
                                    }
                                case "":
                                    v32 := GenS1{F0: true}
                                    _ = v32
                                case "a":
                                    var v33 uint16 = (^(v17))
                                    _ = v33
                                }
                            } else {
                                for fi34 := 0; fi34 < 10; fi34++ {
                                    var v35 int64 = 1
                                    _ = fi34
                                    _ = v35
                                }
                            }
                            _ = fi30
                        }
                    }
                default:
                    var v36 uint16 = v17
                    _ = v36
                }
            }
            _ = v24
        }
    default:
        var v37 string = "a"
        _ = v37
    }
    fmt.Println(v1, v1, v5)
    v5 += v1
    v1 = v1
    switch 0 {
    case 127:
        v38 := GenS3{F0: -339}
        _ = v38
    default:
        if (127 <= -622) {
            var v39 uint8 = 127
            _ = v39
        }
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        if true {
            v41 := GenS5{F0: 42, F1: -218}
            v2 = append(v2, "foo bar")
            var v42 int8 = 42
            _ = v41
            _ = v42
        }
        _ = fi40
    }
    v43 := GenS5{F0: 2, F1: 0}
    if true {
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
