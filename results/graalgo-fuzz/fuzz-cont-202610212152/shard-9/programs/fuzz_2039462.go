package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 string; F2 int32 }
type GenS7 struct { F0 int64; F1 int64; F2 uint16 }
func genFn0(p0 uint64, p1 string) uint {
    switch "\n\t" {
    case "hello":
        v5 := []int{32767, 2147483647, 682, 1}
        _ = v5
    case "":
        fmt.Println(p0)
    case "foo bar":
        p1 = "foo bar"
    }
    if false {
        p0 += 42
        fmt.Println(p0, p0)
        p1 = "hello"
    } else {
        v6 := []int{}
        _ = v6
    }
    _ = p0
    _ = p1
    return 65535
}

func main() {
    v1 := []string{"", "\n\t", "\n\t"}
    var v2 int8 = -128
    var v3 string = (("foo bar" + ("hello" + "hello")) + "hello")
    var v4 float32 = 0.1
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint32 = 47
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if len(v1) > 0 {
        v1[0] = v3
    }
    var v10 int32 = 2147483647
    var v11 int = 603
    v10 = 1
    v3 = (v3 + v3)
    var v12 uint64 = 443
    switch "" {
    case "hello":
        var v13 uint64 = v12
        _ = v13
    case "foo bar":
        if false {
            var v14 int8 = (v2 + (v2 >> ((133) & 63)))
            v15 := GenS1{F0: (v11 >> ((42) & 63))}
            _ = v14
            _ = v15
        } else {
            switch v11 {
            case -1:
                _ = len(v1)
            case 32768:
                v16 := GenS1{F0: v11}
                _ = v16
            default:
                for fi17 := 0; fi17 < 4; fi17++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi17
                }
            }
            for fi18 := 0; fi18 < 11; fi18++ {
                var v19 int64 = 9223372036854775807
                switch (v11 >> ((0) & 63)) {
                case 445:
                    _ = len(v1)
                case -176:
                    switch fi18 {
                    case 4294967295:
                        if true {
                            v20 := GenS1{F0: (v11 & (v11 * (fi18 / ((fi18) | 1))))}
                            fmt.Println(v4, v2)
                            v21 := GenS1{F0: (fi18 + v11)}
                            _ = v20
                            _ = v21
                        } else {
                            if true {
                                var v22 float64 = (float64(v4) / -2.25)
                                for fi23 := 0; fi23 < 1; fi23++ {
                                    var v24 string = v3
                                    if (!false) {
                                        switch -858 {
                                        case -502:
                                            v10 += (v10 | -38)
                                        default:
                                            if (100 <= v2) {
                                                v1 = append(v1, ("foo bar" + "hello"))
                                                _ = len(v1)
                                            } else {
                                                if (17 < 9) {
                                                    switch int(v19) {
                                                    case 42:
                                                        v10 -= (v10 | v10)
                                                    }
                                                    v25 := GenS1{F0: (v11 >> ((65535) & 63))}
                                                    _ = v25
                                                }
                                                v4 = 1e10
                                                var v26 int8 = (v2 & v2)
                                                v19 = v19
                                                _ = v26
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        var v27 int = 2147483647
                                        _ = v27
                                    } else {
                                        _ = len(v1)
                                        var v28 uint = 4294967295
                                        fmt.Println(v24, v28, v28)
                                        _ = v28
                                    }
                                    var v29 uint8 = 255
                                    var v30 uint32 = 158
                                    v31 := GenS2{F0: 658, F1: 1000}
                                    _ = fi23
                                    _ = v24
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                                _ = v22
                            }
                            switch (v11 % ((v11) | 1)) {
                            case -80:
                                v19 += 1
                            case 127:
                                if (267 <= -919) {
                                    v32 := GenS1{F0: (-(v11))}
                                    v33 := GenS3{F0: false, F1: 72, F2: 65535}
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v34 int16 = 0
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v34
                                }
                            }
                        }
                    }
                }
                _ = fi18
                _ = v19
            }
        }
    case "\n\t":
        var v35 int8 = int8(v11)
        _ = v35
    }
    var v36 float32 = float32(1.5)
    v12 = 4294967295
    v37 := GenS4{F0: 0, F1: 4294967295}
    v10 += (-(v10))
    for fi38 := 0; fi38 < 7; fi38++ {
        switch "foo bar" {
        case "a":
            fmt.Println(v11, v4, v2)
        case "\n\t":
            switch v11 {
            case -1:
                var v39 int32 = (v10 - v10)
                _ = v39
            case 128:
                var v40 uint16 = 299
                _ = v40
            }
        }
        _ = fi38
    }
    for fi41 := 0; fi41 < 10; fi41++ {
        if (!false) {
            v4 = v36
            if true {
                v42 := GenS5{F0: ((^(v12)) <= v12)}
                if true {
                    v4 -= float32(0.1)
                    v3 = ("\n\t" + v3)
                } else {
                    v43 := GenS6{F0: v3, F1: "a", F2: v10}
                    _ = v43
                }
                switch (v11 + v11) {
                case 32767:
                    var v44 string = ("" + v3)
                    _ = v44
                case 127:
                    v45 := GenS3{F0: true, F1: 76, F2: 2}
                    _ = v45
                }
                _ = v42
            } else {
                if len(v1) > 0 {
                    v1[0] = v3
                }
                if len(v1) > 0 {
                    v1[0] = v3
                }
            }
        } else {
            switch 653 {
            case -845:
                fmt.Println(fi41, v4, v12)
            case 42:
                var v46 uint32 = 2
                _ = v46
            case 32767:
                for fi47 := 0; fi47 < 3; fi47++ {
                    var v48 bool = (0 <= 1)
                    if v48 {
                        v48 = true
                    }
                    _ = fi47
                    _ = v48
                }
            default:
                var v49 float64 = 1.5
                _ = v49
            }
            switch (fi41 - fi41) {
            case -314:
                fmt.Println(v4, v2)
            case -9223372036854775808:
                v50 := GenS1{F0: 9223372036854775807}
                _ = v50
            case 1:
                if len(v1) > 0 {
                    v1[0] = ((v3 + "hello") + v3)
                }
            }
            _ = len(v1)
        }
        v51 := GenS7{F0: int64(v12), F1: 42, F2: 1}
        _ = fi41
        _ = v51
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
