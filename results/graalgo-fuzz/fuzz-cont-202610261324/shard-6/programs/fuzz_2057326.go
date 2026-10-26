package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }

func main() {
    v1 := []uint32{}
    var v2 uint8 = 1
    v3 := []string{"", ("a" + "foo bar")}
    v4 := []string{}
    var v5 string = ""
    v6 := make(map[string]string)
    fmt.Println(v5, v2)
    v7 := GenS1{F0: v2, F1: 346, F2: 4294967295}
    v5 = "foo bar"
    var v8 int = 127
    v5 = v5
    var v9 int32 = -319
    for fi10 := 0; fi10 < 3; fi10++ {
        if true {
            if (-9223372036854775808 < 9223372036854775807) {
                switch fi10 {
                case -9223372036854775808:
                    var v11 uint8 = v2
                    _ = v11
                case -492:
                    v12 := GenS2{F0: 128}
                    _ = v12
                }
            } else {
                var v13 float64 = float64(-2.25)
                _ = v13
            }
            for fi14 := 0; fi14 < 1; fi14++ {
                if (121 >= 65535) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v9 = (v9 - v9)
                    fmt.Println(v8, fi14, v9)
                    fmt.Println(v2)
                    if false {
                        var v15 int64 = 9223372036854775807
                        _ = v15
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v16 uint64 = uint64(v2)
                        _ = v16
                    }
                    switch 0 {
                    case 4294967295:
                        fmt.Println(v9, fi14, v9)
                        v17 := GenS3{F0: 0}
                        v18 := GenS3{F0: 771}
                        _ = v17
                        _ = v18
                    case 0:
                        var v19 int8 = 127
                        switch fi14 {
                        case 819:
                            if false {
                                if true {
                                    v20 := GenS1{F0: 1, F1: 0, F2: 332}
                                    v21 := GenS1{F0: (v2 << ((1) & 63)), F1: -1, F2: 675}
                                    _ = v20
                                    _ = v21
                                }
                                var v22 float64 = 0.0
                                _ = v22
                            }
                        case 945:
                            v23 := GenS3{F0: 44}
                            _ = v23
                        case 1:
                            fmt.Println(v19, fi14, v2)
                        }
                        _ = v19
                    default:
                        var v24 int = v8
                        _ = v24
                    }
                    var v25 float32 = ((1e10 * float32(1e10)) / 1.5)
                    _ = v25
                } else {
                    switch ((v5 + v5) + ("a" + ((((((v5 + ((v5 + v5) + v5)) + v5) + (v5 + "\n\t")) + (v5 + (v5 + v5))) + (("foo bar" + v5) + v5)) + ((v5 + v5) + v5)))) {
                    case "":
                        v26 := GenS1{F0: (v2 >> ((4294967295) & 63)), F1: 0, F2: 1}
                        _ = v26
                    }
                    fmt.Println(v8)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                if false {
                    var v27 float64 = 1e10
                    if true {
                        v28 := GenS2{F0: -727}
                        _ = v28
                    } else {
                        switch v5 {
                        case "foo bar":
                            fmt.Println(v27, v27, v5)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "hello":
                            v2 += (v2 & v2)
                        case "\n\t":
                            var v29 uint32 = 65535
                            _ = v29
                        }
                        v2 = (v2 - v2)
                    }
                    switch v5 {
                    case "\n\t":
                        _ = len(v4)
                    case "":
                        v30 := GenS2{F0: (fi14 * (v8 & fi14))}
                        _ = v30
                    case "foo bar":
                        if (false || false) {
                            var v31 int16 = -512
                            v32 := GenS4{F0: 42}
                            _ = v31
                            _ = v32
                        }
                    default:
                        v9 += v9
                    }
                    v8 = fi14
                    var v33 float64 = v27
                    _ = v27
                    _ = v33
                }
                _ = fi14
            }
            v5 = v5
        }
        var v34 uint32 = 387
        for fi35 := 0; fi35 < 12; fi35++ {
            v8 += fi10
            v2 -= v2
            _ = fi35
        }
        _ = fi10
        _ = v34
    }
    v36 := v6["\n\t"]
    fmt.Println(v9, v5, v5)
    v5 = (v36 + "a")
    var v37 int64 = -1
    v36 = "\n\t"
    for fi38 := 0; fi38 < 6; fi38++ {
        if true {
            fmt.Println(v2, v2, v9)
            if ((1 >= 1) && true) {
                v2 += v2
                _ = len(v6)
            }
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
        }
        if true {
            var v39 uint8 = 127
            var v40 int8 = -59
            var v41 uint8 = 42
            _ = v39
            _ = v40
            _ = v41
        } else {
            v37 = v37
            var v42 int64 = v37
            _ = v42
        }
        _ = fi38
    }
    if (0 < 100) {
        v43 := GenS2{F0: (v8 | (^(v8)))}
        _ = v43
    }
    v8 = (v8 / ((9223372036854775807) | 1))
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
