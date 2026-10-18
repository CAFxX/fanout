package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 int; F2 string }

func main() {
    var v1 float32 = 1e10
    var v2 int32 = -524
    v3 := []string{("hello" + "hello")}
    v4 := make(map[int]bool)
    var v5 uint32 = 4294967295
    v6 := make(map[int]bool)
    v2 = v2
    for fi7 := 0; fi7 < 1; fi7++ {
        fmt.Println(fi7, fi7, v2)
        fmt.Println(fi7)
        _ = fi7
    }
    var v8 int16 = 1000
    switch 0 {
    case -9223372036854775808:
        for fi9 := 0; fi9 < 4; fi9++ {
            switch fi9 {
            case 128:
                if (!false) {
                    fmt.Println(v2, v8, fi9)
                    fmt.Println(v1, fi9, v5)
                } else {
                    var v10 float64 = float64(((v1 * (v1 - v1)) - 0.0))
                    _ = v10
                }
            }
            _ = fi9
        }
    case 32767:
        if ((18446744073709551615 <= 76) || true) {
            switch 720 {
            case -2147483648:
                if ("\n\t" > "hello") {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        var v12 float64 = 1.5
                        for fi13 := 0; fi13 < 11; fi13++ {
                            var v14 uint16 = 65535
                            v15 := GenS1{F0: true}
                            var v16 int16 = 0
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                        v8 = (v8 | -1)
                        _ = fi11
                        _ = v12
                    }
                    v17 := GenS1{F0: (false || (!true))}
                    fmt.Println(v2, v2)
                    _ = v17
                }
            case 128:
                if true {
                    v18 := GenS2{F0: 0, F1: 0}
                    if (true || (true || (!(0 >= 255)))) {
                        var v19 int8 = -69
                        v19 = v19
                        _ = v19
                    } else {
                        if (255 != 79) {
                            for fi20 := 0; fi20 < 17; fi20++ {
                                var v21 float64 = -2.25
                                v22 := GenS1{F0: (127 != -125)}
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                            v8 = -501
                        } else {
                            if true {
                                fmt.Println(v8, v1, v2)
                            }
                            fmt.Println(v8, v8, v5)
                        }
                        v23 := GenS3{F0: 32767}
                        var v24 int16 = 0
                        _ = v23
                        _ = v24
                    }
                    var v25 int64 = 42
                    _ = v18
                    _ = v25
                } else {
                    fmt.Println(v2, v2)
                }
            case -9223372036854775808:
                v26 := v4[-327]
                _ = v26
            }
            if (((3.14159 / (float32(1e10) * v1)) * 1.5) != float32(1e10)) {
                if false {
                    var v27 uint16 = 156
                    for fi28 := 0; fi28 < 6; fi28++ {
                        var v29 float32 = v1
                        if (620 >= 18446744073709551615) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v1 = 0.0
                        } else {
                            var v30 int16 = v8
                            for fi31 := 0; fi31 < 1; fi31++ {
                                var v32 bool = (!false)
                                fmt.Println(v2)
                                var v33 bool = v32
                                if true {
                                    v5 += v5
                                    var v34 int32 = v2
                                    switch (fi31 | fi28) {
                                    case 0:
                                        if true {
                                            switch (fi31 - (fi31 & 352)) {
                                            case 32768:
                                                delete(v6, fi28)
                                            case 32767:
                                                var v35 string = "foo bar"
                                                _ = v35
                                            default:
                                                if v32 {
                                                    fmt.Println(fi31, v33, v34)
                                                    v33 = v32
                                                    var v36 int64 = -1
                                                    switch 42 {
                                                    case 1:
                                                        var v37 bool = (65535 == 1)
                                                        _ = v37
                                                    case 86:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    default:
                                                        var v38 uint8 = 128
                                                        _ = v38
                                                    }
                                                    if (v29 <= v1) {
                                                        fmt.Println(v27, v30, v2)
                                                    }
                                                    _ = v36
                                                } else {
                                                    fmt.Println(v29, v33, v29)
                                                    v8 = v8
                                                }
                                            }
                                        } else {
                                            _ = len(v3)
                                            v39 := GenS4{F0: -1, F1: 4294967295, F2: "hello"}
                                            var v40 float32 = ((0.1 * (v29 / float32(float64(float32(1.5))))) * v1)
                                            _ = v39
                                            _ = v40
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(fi28, v34)
                                    case -2147483648:
                                        var v41 int8 = -27
                                        _ = v41
                                    default:
                                        var v42 int = int(v27)
                                        _ = v42
                                    }
                                    _ = v34
                                } else {
                                    var v43 uint64 = uint64(v27)
                                    var v44 int16 = (v30 << ((2) & 63))
                                    fmt.Println(v5, v1, v27)
                                    _ = v43
                                    _ = v44
                                }
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                            v30 += (v30 >> ((255) & 63))
                            _ = v30
                        }
                        _ = fi28
                        _ = v29
                    }
                    _ = v27
                }
            }
            delete(v6, 0)
        }
    }
    var v45 int32 = v2
    v1 = v1
    v1 += 1.5
    var v46 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v8
}
