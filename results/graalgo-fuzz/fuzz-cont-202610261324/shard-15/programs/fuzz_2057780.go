package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int64 }
func genFn0() uint16 {
    var v7 int64 = 950
    var v8 int8 = -1
    fmt.Println(v7)
    var v9 float32 = (0.0 / (3.14159 / -2.25))
    v9 -= v9
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := []string{"foo bar", "a"}
    var v2 float32 = -2.25
    var v3 uint16 = 65535
    var v4 uint8 = 195
    v5 := make(map[uint32]string)
    var v6 int64 = 278
    if true {
        switch "\n\t" {
        case "a":
            delete(v5, 1)
        case "foo bar":
            if true {
                fmt.Println(v6)
            } else {
                fmt.Println(v4)
                switch -794 {
                case 32767:
                    var v10 uint64 = 437
                    _ = v10
                case -763:
                    delete(v5, 0)
                case 504:
                    v11 := v5[2147483647]
                    _ = v11
                }
                if (100 != 100) {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        switch (fi12 / (((^(fi12))) | 1)) {
                        case -9223372036854775808:
                            if (!(-75 < 42)) {
                                _ = len(v1)
                            } else {
                                if (("a" + "a") != "foo bar") {
                                    v2 -= 3.14159
                                }
                                var v13 float32 = (float32(1.5) * float32(-2.25))
                                _ = v13
                            }
                        case -174:
                            for fi14 := 0; fi14 < 2; fi14++ {
                                v15 := GenS1{F0: 1, F1: 343}
                                _ = fi14
                                _ = v15
                            }
                        }
                        var v16 uint16 = v3
                        _ = fi12
                        _ = v16
                    }
                    fmt.Println(v4)
                }
            }
        }
        v6 += (v6 & (^(v6)))
        v17 := GenS2{F0: 42, F1: -1}
        _ = v17
    } else {
        var v18 uint = 883
        if (!false) {
            v19 := v5[0]
            var v20 int16 = 0
            if true {
                if ((^(v18)) > v18) {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        fmt.Println(v3)
                        v22 := v5[65535]
                        _ = fi21
                        _ = v22
                    }
                    switch 540 {
                    case 0:
                        v23 := GenS3{F0: 540, F1: 65535}
                        _ = v23
                    }
                } else {
                    var v24 uint8 = (v4 | v4)
                    var v25 bool = (!(false && false))
                    switch (v19 + "a") {
                    case "foo bar":
                        for fi26 := 0; fi26 < 3; fi26++ {
                            if v25 {
                                v24 = v24
                            }
                            _ = fi26
                        }
                    case "a":
                        _ = len(v5)
                    case "\n\t":
                        for fi27 := 0; fi27 < 3; fi27++ {
                            v28 := v5[1]
                            v29 := GenS4{F0: 9223372036854775807}
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    default:
                        for fi30 := 0; fi30 < 15; fi30++ {
                            v31 := v5[964]
                            switch (v19 + ((v31 + v31) + v19)) {
                            case "hello":
                                if v25 {
                                    switch v19 {
                                    case "":
                                        fmt.Println(v2)
                                    case "hello":
                                        var v32 bool = (-2.25 < 0.1)
                                        _ = v32
                                    }
                                }
                            case "":
                                var v33 uint16 = (^(v3))
                                _ = v33
                            }
                            var v34 int8 = 100
                            _ = fi30
                            _ = v31
                            _ = v34
                        }
                    }
                    _ = v24
                    _ = v25
                }
                fmt.Println(v20)
            }
            _ = v19
            _ = v20
        } else {
            if false {
                var v35 uint16 = genFn0()
                _ = v35
            }
            v6 += (v6 - (^(v6)))
        }
        _ = v18
    }
    v4 = (v4 % ((2) | 1))
    fmt.Println(v2)
    v4 = v4
    var v36 string = ""
    var v37 float64 = 0.1
    v6 = (v6 >> ((699) & 63))
    var v38 string = v36
    v39 := GenS3{F0: 1, F1: 2147483647}
    var v40 int32 = 42
    if (0 < 701) {
        for fi41 := 0; fi41 < 20; fi41++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v42 := GenS2{F0: 1, F1: 32767}
            _ = fi41
            _ = v42
        }
        var v43 int8 = -116
        _ = v43
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v36 = v38
        v2 -= v2
    }
    for fi44 := 0; fi44 < 7; fi44++ {
        switch v38 {
        case "a":
            var v45 string = ((v38 + v38) + v38)
            _ = v45
        case "\n\t":
            v6 = (v6 ^ -21)
        case "":
            v40 += v40
        default:
            var v46 string = v38
            _ = v46
        }
        _ = fi44
    }
    v38 = v38
    var v47 int32 = v40
    v38 = (((((((v38 + v36) + "\n\t") + v38) + "\n\t") + "foo bar") + "a") + v38)
    v47 -= (v47 << ((255) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v47
    _ = v5
    _ = v6
}
