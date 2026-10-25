package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int32; F2 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 uint16; F1 bool }

func main() {
    v1 := []string{"", "\n\t", ("foo bar" + ("\n\t" + "\n\t"))}
    var v2 int8 = -128
    var v3 uint64 = 1
    if false {
        var v4 uint64 = (v3 << ((558) & 63))
        if ((42 < -312) || false) {
            fmt.Println(v2, v4, v3)
            for fi5 := 0; fi5 < 13; fi5++ {
                switch fi5 {
                case -1:
                    var v6 string = "\n\t"
                    _ = v6
                case 2147483647:
                    fmt.Println(v3, v2, v4)
                case 32767:
                    fmt.Println(v4, v4, fi5)
                }
                switch fi5 {
                case -1:
                    v1 = append(v1, "a")
                case 9223372036854775807:
                    switch ("a" + "hello") {
                    case "a":
                        var v7 float32 = float32(-2.25)
                        _ = v7
                    case "":
                        if (-37 == 42) {
                            fmt.Println(v3, v4)
                        } else {
                            var v8 uint16 = 208
                            for fi9 := 0; fi9 < 20; fi9++ {
                                if true {
                                    v10 := GenS1{F0: 0, F1: fi9, F2: true}
                                    if true {
                                        var v11 int32 = 2147483647
                                        var v12 int32 = (-(v11))
                                        v13 := GenS1{F0: -1, F1: (fi5 / ((fi5) | 1)), F2: true}
                                        _ = v11
                                        _ = v12
                                        _ = v13
                                    }
                                    switch 2147483647 {
                                    case 913:
                                        v14 := GenS2{F0: (^(v2))}
                                        _ = v14
                                    case -994:
                                        var v15 float32 = -2.25
                                        _ = v15
                                    case 32767:
                                        if ((18446744073709551615 <= 143) && true) {
                                            if (true && (!true)) {
                                                var v16 bool = (!true)
                                                _ = len(v1)
                                                _ = v16
                                            } else {
                                                var v17 uint64 = v3
                                                _ = v17
                                            }
                                        }
                                        var v18 string = ""
                                        _ = len(v1)
                                        var v19 float64 = (-2.25 * ((0.0 + 1e10) / (0.1 + (0.1 - (float64(0.0) - -2.25)))))
                                        _ = v18
                                        _ = v19
                                    }
                                    switch fi5 {
                                    case 32768:
                                        var v20 int8 = (v2 << ((214) & 63))
                                        _ = v20
                                    default:
                                        var v21 uint = 65535
                                        _ = v21
                                    }
                                    _ = v10
                                }
                                var v22 uint16 = 1
                                _ = fi9
                                _ = v22
                            }
                            _ = v8
                        }
                    }
                default:
                    if true {
                        var v23 int8 = v2
                        v24 := GenS3{F0: 4294967295}
                        var v25 string = ("foo bar" + (("" + "a") + ("foo bar" + ("hello" + (("hello" + "hello") + ("\n\t" + "foo bar"))))))
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                }
                v4 = (v3 * v4)
                _ = fi5
            }
            _ = len(v1)
        } else {
            fmt.Println(v4, v2, v2)
        }
        _ = v4
    } else {
        if true {
            v3 += v3
        }
        if true {
            v2 = v2
            fmt.Println(v3, v3)
            v26 := GenS4{F0: 32767, F1: 1, F2: ((float64(float32((1.5 - float64((1.5 + (-2.25 * 0.1)))))) + 3.14159) != float64(1e10))}
            _ = v26
        } else {
            fmt.Println(v2, v3)
            switch "" {
            case "\n\t":
                v27 := GenS5{F0: -588}
                _ = v27
            }
            v28 := GenS3{F0: 42}
            _ = v28
        }
        v29 := GenS6{F0: 2, F1: v2, F2: 500}
        _ = v29
    }
    var v30 uint64 = v3
    for fi31 := 0; fi31 < 9; fi31++ {
        for fi32 := 0; fi32 < 14; fi32++ {
            v33 := GenS7{F0: "", F1: v30}
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    if false {
        for fi34 := 0; fi34 < 14; fi34++ {
            var v35 int16 = 14
            _ = fi34
            _ = v35
        }
        switch 127 {
        case 500:
            for fi36 := 0; fi36 < 16; fi36++ {
                fmt.Println(v3, v3, fi36)
                switch (fi36 | 929) {
                case 853:
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                case -699:
                    if (false && true) {
                        v3 = v3
                    } else {
                        var v37 bool = true
                        _ = v37
                    }
                }
                _ = fi36
            }
        }
    }
    var v38 uint16 = 524
    v3 = v3
    var v39 uint16 = (v38 << ((0) & 63))
    var v40 uint16 = 2
    fmt.Println(v39)
    v41 := GenS6{F0: (v38 >> ((42) & 63)), F1: 40, F2: 65535}
    var v42 int8 = v2
    fmt.Println(v40, v2, v38)
    v43 := GenS8{F0: (v38 + 65535), F1: false}
    var v44 float64 = (0.0 - 0.0)
    var v45 int = 32768
    v3 = 543
    fmt.Println(v2, v44)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
}
