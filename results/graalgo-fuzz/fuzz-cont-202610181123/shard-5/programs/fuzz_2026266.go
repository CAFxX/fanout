package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 bool; F1 string }
func genFn0() int16 {
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 int8 = 127
        var v6 uint = 1
        v6 = (v6 & (v6 * v6))
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int16 = 0
    var v8 int = -547
    fmt.Println(v7, v7)
    _ = v7
    _ = v8
    return (v7 ^ (v7 % ((v7) | 1)))
}

func main() {
    v1 := []string{"foo bar"}
    var v2 int64 = 901
    v3 := make(map[string]string)
    if (true && true) {
        switch ("" + "\n\t") {
        case "hello":
            fmt.Println(v2)
        case "a":
            var v9 float64 = 1e10
            _ = v9
        default:
            v10 := GenS1{F0: v2}
            _ = v10
        }
    }
    var v11 uint16 = 934
    switch 127 {
    case 32768:
        var v12 int32 = 42
        _ = v12
    }
    fmt.Println(v2, v2)
    if (!false) {
        var v13 uint16 = (v11 & v11)
        var v14 int = 127
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int8 = 102
        var v17 int64 = (v2 >> ((255) & 63))
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 int32 = 42
    v19 := GenS1{F0: v2}
    v20 := v3["foo bar"]
    delete(v3, v20)
    v18 += (v18 << ((1) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 uint64 = 42
        var v23 string = v20
        for fi24 := 0; fi24 < 20; fi24++ {
            if ((^(v11)) == v11) {
                if (-1 != 100) {
                    var v25 bool = false
                    v20 = ((("\n\t" + v20) + "") + (v23 + ("hello" + v23)))
                    v26 := GenS2{F0: 0, F1: v11}
                    _ = v25
                    _ = v26
                } else {
                    v22 += v22
                }
                fmt.Println(v20, v22)
            } else {
                if false {
                    if false {
                        v27 := v3[(v23 + "foo bar")]
                        if true {
                            switch "a" {
                            case "":
                                _ = len(v1)
                            }
                            _ = len(v1)
                            var v28 uint16 = v11
                            _ = v28
                        } else {
                            v29 := GenS2{F0: 1, F1: v11}
                            _ = v29
                        }
                        v27 = ("\n\t" + (v27 + v20))
                        switch (v23 + ("\n\t" + ((("\n\t" + "hello") + v20) + (v23 + "")))) {
                        case "foo bar":
                            v2 = -9223372036854775808
                        case "a":
                            fmt.Println(v23)
                        default:
                            if (1 >= -647) {
                                v30 := GenS1{F0: (v2 % (((v2 / ((v2) | 1))) | 1))}
                                v1 = append(v1, "foo bar")
                                _ = v30
                            } else {
                                if ((v22 ^ v22) != 2) {
                                    var v31 int8 = 0
                                    var v32 uint16 = v11
                                    var v33 bool = (true || true)
                                    switch fi24 {
                                    case 369:
                                        v34 := GenS1{F0: v2}
                                        _ = v34
                                    }
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v35 uint8 = 1
                                    _ = v35
                                }
                                switch (-(fi21)) {
                                case 127:
                                    v20 = ""
                                }
                                if (v11 <= 873) {
                                    var v36 bool = false
                                    v37 := GenS2{F0: 369, F1: (v11 & (^(v11)))}
                                    _ = v36
                                    _ = v37
                                }
                                v18 -= v18
                                if ((^(fi21)) > (fi24 >> ((1) & 63))) {
                                    var v38 uint16 = (v11 ^ v11)
                                    if ((v22 != (^(v22))) || (100 < 0)) {
                                        v11 -= (v11 - v11)
                                        var v39 int64 = (v2 % (((v2 & 9223372036854775807)) | 1))
                                        switch v23 {
                                        case "foo bar":
                                            fmt.Println(v23)
                                        default:
                                            switch (v27 + (v27 + "hello")) {
                                            case "":
                                                switch (v20 + (v27 + v20)) {
                                                case "a":
                                                    if true {
                                                        fmt.Println(v39)
                                                    }
                                                case "foo bar":
                                                    v2 = v2
                                                default:
                                                    v27 = "foo bar"
                                                }
                                            default:
                                                _ = len(v3)
                                            }
                                        }
                                        _ = v39
                                    } else {
                                        _ = len(v1)
                                    }
                                    var v40 int64 = 804
                                    _ = v38
                                    _ = v40
                                } else {
                                    switch v27 {
                                    case "foo bar":
                                        v41 := GenS1{F0: int64(fi21)}
                                        _ = v41
                                    }
                                    v20 = "foo bar"
                                }
                                v42 := GenS3{F0: false, F1: v27}
                                _ = v42
                            }
                            if false {
                                v23 = ""
                            } else {
                                var v43 int16 = -32768
                                if false {
                                    v1 = append(v1, v20)
                                    var v44 uint16 = v11
                                    var v45 float32 = (0.1 / 3.14159)
                                    v44 = (v44 & (v11 ^ (^(v44))))
                                    _ = v44
                                    _ = v45
                                }
                                _ = v43
                            }
                        }
                        var v46 int = fi24
                        switch v27 {
                        case "a":
                            var v47 uint64 = 768
                            _ = v47
                        default:
                            v48 := GenS3{F0: (v11 == (v11 * (v11 | v11))), F1: v20}
                            _ = v48
                        }
                        _ = v27
                        _ = v46
                    } else {
                        var v49 uint64 = 42
                        _ = v49
                    }
                    if (!(v18 == (v18 >> ((412) & 63)))) {
                        var v50 uint64 = 33
                        v51 := GenS3{F0: (2 != 239), F1: v23}
                        v11 = (v11 % ((1) | 1))
                        _ = v50
                        _ = v51
                    }
                } else {
                    var v52 int32 = v18
                    _ = v52
                }
                v22 = 42
            }
            var v53 uint32 = 2147483648
            var v54 string = v20
            _ = fi24
            _ = v53
            _ = v54
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v11)
    for fi55 := 0; fi55 < 20; fi55++ {
        fmt.Println(v11, v20)
        fmt.Println(v20, fi55, v11)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi55
    }
    for fi56 := 0; fi56 < 4; fi56++ {
        v18 -= (v18 - v18)
        v3[(((v20 + "foo bar") + v20) + "")] = ""
        v57 := GenS2{F0: 0, F1: v11}
        _ = fi56
        _ = v57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
