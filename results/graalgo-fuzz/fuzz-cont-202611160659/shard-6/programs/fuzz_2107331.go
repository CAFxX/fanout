package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int8 = -128
    v2 := []string{""}
    var v3 float32 = -2.25
    var v4 bool = (!(false && (false && (128 > 0))))
    for fi5 := 0; fi5 < 6; fi5++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = len(v2)
        var v6 float64 = (1e10 * ((1.5 - 3.14159) / 1.0))
        _ = fi5
        _ = v6
    }
    _ = len(v2)
    v2 = append(v2, "")
    if len(v2) > 0 {
        v2[0] = ("a" + "foo bar")
    }
    switch "" {
    case "a":
        v1 += v1
    case "":
        v1 = -3
    }
    v7 := GenS1{F0: (-2.25 < v3), F1: 1}
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 int16 = 770
        v10 := GenS1{F0: (!v4), F1: 0}
        var v11 bool = v4
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        v2 = append(v2, "foo bar")
        _ = fi12
    }
    var v13 float32 = v3
    for fi14 := 0; fi14 < 1; fi14++ {
        if v4 {
            for fi15 := 0; fi15 < 15; fi15++ {
                if v4 {
                    fmt.Println(v1, v13, v13)
                    if (1000 >= 83) {
                        v1 -= v1
                        v3 -= (v13 / 1.0)
                    } else {
                        v4 = (0 <= 68)
                        v13 = ((v13 / ((v13 + 0.0) / v3)) + 0.1)
                    }
                    if len(v2) > 0 {
                        v2[0] = ("" + "")
                    }
                }
                switch "" {
                case "foo bar":
                    if v4 {
                        if (!(-1 <= 42)) {
                            var v16 uint64 = 581
                            switch 2147483647 {
                            case 132:
                                if v4 {
                                    var v17 int64 = -1
                                    var v18 uint16 = 764
                                    v18 = v18
                                    _ = v17
                                    _ = v18
                                }
                            case 419:
                                v13 = 1.5
                            case 128:
                                var v19 uint32 = 1
                                _ = v19
                            }
                            _ = v16
                        }
                    } else {
                        if (-1 < -619) {
                            fmt.Println(v13, v3)
                            if v4 {
                                v20 := GenS1{F0: v4, F1: 1}
                                var v21 uint64 = 2
                                _ = v20
                                _ = v21
                            } else {
                                var v22 int8 = -1
                                v23 := GenS1{F0: v4, F1: 730}
                                var v24 uint64 = 728
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                            v25 := GenS2{F0: ("foo bar" + "")}
                            _ = v25
                        } else {
                            var v26 float32 = 1e10
                            if v4 {
                                fmt.Println(v26, v3)
                                v4 = (!v4)
                                _ = len(v2)
                            }
                            switch "" {
                            case "foo bar":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "\n\t":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            _ = v26
                        }
                        var v27 uint8 = 29
                        v1 = v1
                        _ = v27
                    }
                case "hello":
                    switch fi14 {
                    case 32768:
                        fmt.Println(fi15)
                    }
                case "\n\t":
                    var v28 uint16 = 65535
                    _ = v28
                }
                fmt.Println(fi14, fi15)
                _ = fi15
            }
        }
        var v29 int64 = -1
        _ = fi14
        _ = v29
    }
    if len(v2) > 0 {
        v2[0] = ""
    }
    if v4 {
        fmt.Println(v4, v4)
        v4 = ((3.14159 + float64(v13)) != 1e10)
    }
    v1 = v1
    var v30 int = 414
    v4 = (!v4)
    if (380 >= 0) {
        v1 = (v1 % (((v1 & -128)) | 1))
        fmt.Println(v3, v1, v30)
    } else {
        v13 -= float32((1e10 * 1.5))
        v31 := GenS1{F0: v4, F1: 2}
        v13 = v3
        _ = v31
    }
    v1 = v1
    if (!true) {
        v32 := GenS3{F0: -734, F1: 255}
        for fi33 := 0; fi33 < 8; fi33++ {
            switch (fi33 >> ((4294967295) & 63)) {
            case 127:
                for fi34 := 0; fi34 < 8; fi34++ {
                    v2 = append(v2, "")
                    fmt.Println(v3, v4)
                    _ = fi34
                }
            case 128:
                var v35 bool = ((v1 << ((0) & 63)) >= v1)
                _ = v35
            case 32768:
                v3 += (0.1 / v3)
            }
            _ = fi33
        }
        v36 := GenS2{F0: ("a" + "a")}
        _ = v32
        _ = v36
    } else {
        v37 := GenS3{F0: 0, F1: 1}
        var v38 uint16 = 386
        var v39 uint64 = 4294967295
        _ = v37
        _ = v38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v7
}
