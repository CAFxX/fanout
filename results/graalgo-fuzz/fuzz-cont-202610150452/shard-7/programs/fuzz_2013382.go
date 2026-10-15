package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int64; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32; F1 bool; F2 int64 }
type GenS7 struct { F0 uint32; F1 uint8; F2 uint16 }
func genFn0() int64 {
    if (0.1 != -2.25) {
        var v7 uint8 = 42
        _ = v7
    }
    var v8 int8 = -19
    _ = v8
    return 1
}

func main() {
    var v1 int8 = -128
    var v2 string = ""
    v3 := []string{v2, v2}
    var v4 uint = 1
    v5 := make(map[int]string)
    var v6 int = 32768
    for fi9 := 0; fi9 < 1; fi9++ {
        v1 += (v1 % ((v1) | 1))
        v3 = append(v3, "\n\t")
        _ = fi9
    }
    v1 = v1
    var v10 bool = false
    if false {
        var v11 bool = true
        switch (v6 << ((v4) & 63)) {
        case 127:
            if v11 {
                switch v6 {
                case -2147483648:
                    delete(v5, (v6 & v6))
                }
                switch v2 {
                case "foo bar":
                    if ("" == ((((v2 + "\n\t") + v2) + "a") + (v2 + v2))) {
                        switch (v6 % ((v6) | 1)) {
                        case 32767:
                            v12 := GenS1{F0: "", F1: v4, F2: -910}
                            _ = v12
                        }
                        if (2147483647 < 802) {
                            delete(v5, (v6 ^ v6))
                        }
                        var v13 float32 = -2.25
                        _ = v13
                    } else {
                        delete(v5, (v6 % ((v6) | 1)))
                        for fi14 := 0; fi14 < 5; fi14++ {
                            switch v2 {
                            case "foo bar":
                                v15 := GenS1{F0: v2, F1: v4, F2: genFn0()}
                                _ = v15
                            case "\n\t":
                                switch v2 {
                                case "hello":
                                    v2 = "hello"
                                case "\n\t":
                                    for fi16 := 0; fi16 < 16; fi16++ {
                                        var v17 float64 = 1.5
                                        _ = fi16
                                        _ = v17
                                    }
                                case "a":
                                    if (869 == -365) {
                                        switch (v2 + (v2 + (v2 + v2))) {
                                        case "\n\t":
                                            delete(v5, (fi14 * (v6 * (fi14 >> (((v4 - v4)) & 63)))))
                                        }
                                    }
                                default:
                                    v1 = v1
                                }
                            case "hello":
                                if (0.1 <= float64((1e10 / 1.0))) {
                                    var v18 int64 = 1
                                    v1 = v1
                                    _ = v18
                                }
                            default:
                                fmt.Println(v6, v6)
                            }
                            _ = fi14
                        }
                        var v19 int8 = (v1 << ((v4) & 63))
                        _ = v19
                    }
                case "a":
                    if false {
                        _ = len(v3)
                        v2 = v2
                        for fi20 := 0; fi20 < 13; fi20++ {
                            if v11 {
                                v2 = v2
                                var v21 uint32 = 0
                                v11 = true
                                _ = v21
                            }
                            _ = fi20
                        }
                    }
                case "hello":
                    v22 := GenS2{F0: 1, F1: 1}
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v24 := v5[(fi23 + (-(v6)))]
                        v25 := GenS1{F0: (v24 + (v2 + (v2 + ("a" + (v2 + ("\n\t" + (v24 + (v2 + "foo bar")))))))), F1: 4294967295, F2: 1}
                        v26 := GenS1{F0: "a", F1: 630, F2: 0}
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                }
                var v27 float32 = 0.0
                _ = v27
            }
        case -1:
            if len(v3) > 0 {
                v3[0] = (v2 + "hello")
            }
        case 1:
            v28 := GenS1{F0: v2, F1: v4, F2: 672}
            _ = v28
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            if (!true) {
                v4 -= (v4 - 65535)
            }
            for fi30 := 0; fi30 < 8; fi30++ {
                fmt.Println(fi30, v11)
                _ = fi30
            }
            var v31 string = v2
            _ = fi29
            _ = v31
        }
        _ = v11
    } else {
        v32 := GenS3{F0: 42, F1: 42, F2: false}
        _ = v32
    }
    if v10 {
        v33 := GenS1{F0: v2, F1: (v4 - 1), F2: 1}
        v34 := GenS4{F0: v1}
        for fi35 := 0; fi35 < 20; fi35++ {
            _ = len(v3)
            switch (v2 + v2) {
            case "\n\t":
                _ = len(v5)
            case "a":
                var v36 uint16 = 2
                _ = v36
            }
            _ = fi35
        }
        _ = v33
        _ = v34
    } else {
        var v37 string = "foo bar"
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v1, v6)
        _ = v37
    }
    if v10 {
        v10 = (350 > 5)
        v38 := GenS5{F0: (^(v6))}
        _ = v38
    }
    fmt.Println(v2, v1, v2)
    _ = len(v5)
    for fi39 := 0; fi39 < 9; fi39++ {
        for fi40 := 0; fi40 < 10; fi40++ {
            if (4294967295 < 501) {
                switch ((("hello" + ((v2 + "\n\t") + "\n\t")) + "a") + ("\n\t" + "")) {
                case "hello":
                    var v41 float32 = 1e10
                    _ = v41
                case "a":
                    switch 929 {
                    case 2147483647:
                        fmt.Println(v10, v6, fi40)
                        var v42 uint = v4
                        v43 := GenS6{F0: 189, F1: v10, F2: 967}
                        _ = v42
                        _ = v43
                    }
                case "foo bar":
                    v44 := GenS7{F0: 2147483647, F1: 1, F2: 65535}
                    _ = v44
                default:
                    var v45 float32 = 3.14159
                    _ = v45
                }
            } else {
                fmt.Println(v1, fi39)
                var v46 float64 = -2.25
                _ = v46
            }
            var v47 string = (v2 + "foo bar")
            if (!(true && ((v47 + "\n\t") <= (("hello" + v47) + v2)))) {
                var v48 uint = 42
                var v49 uint = v48
                _ = v48
                _ = v49
            }
            _ = fi40
            _ = v47
        }
        _ = fi39
    }
    fmt.Println(v4)
    fmt.Println(v6, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
