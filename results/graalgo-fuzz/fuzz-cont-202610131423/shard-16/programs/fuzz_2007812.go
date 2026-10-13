package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[string]uint32)
    v3 := []string{}
    v4 := []int{127, 128, 128, 989}
    fmt.Println(v1)
    for fi5 := 0; fi5 < 7; fi5++ {
        v6 := v2["foo bar"]
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        delete(v2, "hello")
        v8 := GenS1{F0: -1, F1: 1, F2: ""}
        var v9 uint8 = 128
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int = fi10
        _ = fi10
        _ = v11
    }
    if (false && (0 >= -123)) {
        v2[""] = v1
        fmt.Println(v1)
        for fi12 := 0; fi12 < 2; fi12++ {
            v1 = 42
            _ = fi12
        }
    } else {
        v13 := GenS2{F0: 118}
        v1 = (v1 & (^(v1)))
        switch 42 {
        case -43:
            switch -641 {
            case -715:
                if (-387 <= -1) {
                    for fi14 := 0; fi14 < 9; fi14++ {
                        var v15 float32 = 1e10
                        _ = fi14
                        _ = v15
                    }
                    v1 += 65535
                    if false {
                        v16 := GenS2{F0: -1}
                        v17 := GenS1{F0: 0, F1: 32767, F2: ""}
                        var v18 int8 = -115
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        switch -347 {
                        case 128:
                            var v19 uint32 = v1
                            _ = v19
                        case -726:
                            if false {
                                v20 := v2[""]
                                for fi21 := 0; fi21 < 2; fi21++ {
                                    switch "hello" {
                                    case "":
                                        fmt.Println(fi21)
                                    }
                                    _ = fi21
                                }
                                _ = v20
                            } else {
                                fmt.Println(v1)
                                var v22 float32 = (((float32(3.14159) / 3.14159) - 0.0) + float32((1e10 - 0.1)))
                                _ = v22
                            }
                        case 0:
                            v1 = v1
                        default:
                            v23 := GenS1{F0: -102, F1: 451, F2: ((("a" + "a") + "a") + (("hello" + "foo bar") + ("a" + "foo bar")))}
                            _ = v23
                        }
                        var v24 int64 = -1
                        if (249 < 1) {
                            v24 += 42
                            v1 -= (v1 >> ((720) & 63))
                            var v25 int64 = v24
                            _ = v25
                        }
                        _ = v24
                    }
                }
            default:
                v26 := v2["\n\t"]
                _ = v26
            }
        }
        _ = v13
    }
    v1 = v1
    fmt.Println(v1)
    var v27 float32 = 1.5
    v28 := GenS2{F0: 0}
    switch "hello" {
    case "\n\t":
        switch ("" + "hello") {
        case "hello":
            for fi29 := 0; fi29 < 9; fi29++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi29
            }
        default:
            fmt.Println(v1, v1)
        }
    case "foo bar":
        v30 := GenS2{F0: int32(v1)}
        _ = v30
    default:
        v31 := GenS2{F0: -2147483648}
        _ = v31
    }
    _ = len(v4)
    var v32 int16 = 692
    v33 := GenS3{F0: -1}
    if (0.1 == 1e10) {
        for fi34 := 0; fi34 < 6; fi34++ {
            v27 += (v27 - ((-2.25 + v27) * 1e10))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi34
        }
        v32 = v32
        v35 := GenS1{F0: 195, F1: 65535, F2: ("foo bar" + "\n\t")}
        _ = v35
    } else {
        v27 = 3.14159
        v27 -= (float32((0.0 + (3.14159 / 1.5))) / -2.25)
        var v36 uint = 1
        _ = v36
    }
    var v37 uint = 341
    switch 127 {
    case 0:
        for fi38 := 0; fi38 < 20; fi38++ {
            v27 = (v27 - (float32((1.5 + 3.14159)) - -2.25))
            _ = fi38
        }
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        v37 = 65535
        v40 := GenS3{F0: 9223372036854775807}
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v4
}
