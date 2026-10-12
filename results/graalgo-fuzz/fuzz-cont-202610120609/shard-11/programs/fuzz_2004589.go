package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int }
type GenS2 struct { F0 int32; F1 int8; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 uint; F2 int8 }
type GenS5 struct { F0 bool }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 int8 = -128
    v4 := make(map[uint32]bool)
    v5 := make(map[int]bool)
    var v6 string = "hello"
    switch 2147483647 {
    case 42:
        var v7 bool = (v6 == v6)
        _ = v7
    case 128:
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint8 = 0
            v10 := GenS1{F0: fi8, F1: v3, F2: 2147483647}
            _ = fi8
            _ = v10
            _ = v9
        }
    case 368:
        var v11 int8 = v3
        _ = v11
    default:
        var v12 int64 = 1
        _ = v12
    }
    delete(v5, 32767)
    v3 -= v3
    var v13 uint8 = 0
    for fi14 := 0; fi14 < 15; fi14++ {
        switch v6 {
        case "":
            switch v6 {
            case "\n\t":
                v3 = v3
            case "a":
                v15 := []int{(fi14 ^ (fi14 + (fi14 / ((fi14) | 1)))), (fi14 >> ((161) & 63)), fi14, (fi14 << ((741) & 63))}
                _ = v15
            default:
                for fi16 := 0; fi16 < 5; fi16++ {
                    switch (fi16 ^ (fi14 ^ (fi16 >> ((4294967295) & 63)))) {
                    case 4294967295:
                        fmt.Println(fi14)
                    case 297:
                        if (128 < (fi16 ^ 32768)) {
                            if true {
                                v17 := []int{32768, fi16}
                                _ = v17
                            } else {
                                v5[(fi14 & fi16)] = false
                            }
                            switch ("\n\t" + "hello") {
                            case "foo bar":
                                var v18 uint = 1
                                _ = v18
                            case "a":
                                v19 := []int{fi14, 32768, (^(fi14)), fi14}
                                _ = v19
                            default:
                                var v20 int8 = v3
                                v21 := GenS2{F0: -2147483648, F1: (^(v20)), F2: false}
                                var v22 uint32 = 244
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        } else {
                            v13 = (v13 << ((65535) & 63))
                            v23 := GenS3{F0: 42}
                            var v24 uint64 = 971
                            var v25 float64 = (0.0 / 1.0)
                            v6 = "hello"
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    case 127:
                        v26 := GenS2{F0: -420, F1: (^(v3)), F2: (!(((v6 + "foo bar") + (v6 + (v6 + v6))) != v6))}
                        _ = v26
                    }
                    if (!((false || false) && true)) {
                        var v27 int64 = -1
                        _ = v27
                    }
                    v28 := GenS1{F0: (fi14 << ((65535) & 63)), F1: v3, F2: (fi14 + fi16)}
                    _ = fi16
                    _ = v28
                }
            }
        }
        for fi29 := 0; fi29 < 5; fi29++ {
            v30 := GenS4{F0: 1, F1: 18446744073709551615, F2: 0}
            var v31 uint32 = 14
            _ = fi29
            _ = v30
            _ = v31
        }
        v3 = (-(v3))
        _ = fi14
    }
    v6 = v6
    switch ((v6 + v6) + "hello") {
    case "foo bar":
        var v32 float32 = ((0.0 + 0.0) + 3.14159)
        _ = v32
    default:
        fmt.Println(v13, v6, v6)
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        var v34 uint8 = (^(v13))
        var v35 float64 = (((0.0 - (-2.25 + 1e10)) * (0.0 + 1.5)) + (((-2.25 + (-2.25 * (1e10 + float64((-2.25 + -2.25))))) / (1.5 / 3.14159)) / float64(((float32(0.1) / (1e10 - 1.5)) + (3.14159 / float32(3.14159))))))
        v36 := GenS1{F0: fi33, F1: v3, F2: 1}
        _ = fi33
        _ = v34
        _ = v35
        _ = v36
    }
    fmt.Println(v3, v13, v6)
    var v37 uint64 = 4294967295
    v37 = 1
    var v38 float64 = (1.5 - 0.0)
    v39 := GenS5{F0: ((v37 * (v37 / ((v37) | 1))) >= (v37 % (((v37 + v37)) | 1)))}
    v40 := []int{-849}
    switch (v6 + (v6 + "\n\t")) {
    case "":
        if len(v40) >= 2 {
            v40 = v40[1:len(v40)]
        }
    case "a":
        if (2 >= uint(v13)) {
            for fi41 := 0; fi41 < 5; fi41++ {
                for fi42 := 0; fi42 < 18; fi42++ {
                    var v43 float32 = (((3.14159 - float32(v38)) * 0.0) + 0.0)
                    v1[v6] = (-(fi41))
                    var v44 string = "foo bar"
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                var v45 int32 = -841
                v46 := GenS3{F0: 65535}
                _ = fi41
                _ = v45
                _ = v46
            }
            var v47 int8 = (^(v3))
            if true {
                v40 = append(v40, -2147483648)
                v47 = v3
            }
            _ = v47
        }
    case "hello":
        v48 := GenS5{F0: (0 != 677)}
        _ = v48
    }
    v6 = "foo bar"
    var v49 int64 = 82
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v49
    _ = v5
    _ = v6
}
