package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := []int{0, 1, 4294967295, 1}
    var v4 string = "\n\t"
    var v5 int64 = 9223372036854775807
    switch v4 {
    case "\n\t":
        var v6 uint32 = 503
        _ = v6
    case "foo bar":
        v7 := GenS1{F0: 411, F1: ""}
        _ = v7
    }
    v8 := GenS2{F0: true}
    v5 = 1
    switch 42 {
    case 128:
        var v9 float64 = 1e10
        _ = v9
    default:
        switch (v4 + "foo bar") {
        case "foo bar":
            switch 440 {
            case -981:
                v4 = "a"
            default:
                v4 = "\n\t"
            }
        default:
            switch 2147483647 {
            case 1:
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 string = ""
                    _ = fi10
                    _ = v11
                }
            case 390:
                for fi12 := 0; fi12 < 8; fi12++ {
                    fmt.Println(v5, fi12, fi12)
                    v13 := GenS2{F0: (false && true)}
                    _ = fi12
                    _ = v13
                }
            case 923:
                var v14 uint = 1
                _ = v14
            }
        }
    }
    v5 -= v5
    var v15 int64 = v5
    var v16 int8 = 42
    v16 -= v16
    v17 := v2[42]
    if true {
        v18 := GenS1{F0: 42, F1: "hello"}
        _ = v18
    } else {
        v19 := v1[32767]
        for fi20 := 0; fi20 < 10; fi20++ {
            var v21 bool = false
            if (32767 >= 44) {
                var v22 int32 = -1
                v21 = ((v5 > v15) && v19)
                for fi23 := 0; fi23 < 10; fi23++ {
                    if true {
                        v22 += (v22 - 456)
                        var v24 float32 = ((0.1 - (0.0 * 0.1)) - (0.0 * 0.0))
                        v25 := GenS2{F0: true}
                        _ = v24
                        _ = v25
                    } else {
                        v21 = v19
                    }
                    var v26 uint = 0
                    fmt.Println(v15, v17, v21)
                    _ = fi23
                    _ = v26
                }
                _ = v22
            } else {
                v19 = (-659 != 42)
                for fi27 := 0; fi27 < 2; fi27++ {
                    _ = len(v2)
                    _ = fi27
                }
                if v19 {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        var v29 string = (((v4 + "a") + "foo bar") + "\n\t")
                        v30 := GenS2{F0: v21}
                        if len(v3) > 0 {
                            v3[0] = v17
                        }
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    if (v21 || true) {
                        if (!v19) {
                            var v31 int = fi20
                            for fi32 := 0; fi32 < 8; fi32++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if true {
                                    if (true || v19) {
                                        var v33 uint8 = 213
                                        _ = v33
                                    } else {
                                        switch (v4 + "hello") {
                                        case "\n\t":
                                            switch v4 {
                                            case "foo bar":
                                                v4 = "foo bar"
                                            default:
                                                switch "hello" {
                                                case "hello":
                                                    switch v4 {
                                                    case "hello":
                                                        v4 = v4
                                                    case "foo bar":
                                                        var v34 uint16 = 255
                                                        _ = v34
                                                    case "a":
                                                        var v35 float32 = (-2.25 / float32(0.1))
                                                        _ = v35
                                                    default:
                                                        var v36 string = v4
                                                        _ = v36
                                                    }
                                                default:
                                                    v37 := v2[65535]
                                                    _ = v37
                                                }
                                            }
                                        default:
                                            v38 := GenS2{F0: (0 < int32(v17))}
                                            _ = v38
                                        }
                                    }
                                } else {
                                    var v39 int64 = v5
                                    if false {
                                        v19 = v19
                                    }
                                    v15 -= v39
                                    _ = v39
                                }
                                switch fi32 {
                                case 4294967295:
                                    if ((-750 >= -1) && (!(-310 < 0))) {
                                        fmt.Println(v16)
                                        v4 = ("foo bar" + v4)
                                    }
                                default:
                                    v21 = v19
                                }
                                _ = fi32
                            }
                            _ = v31
                        } else {
                            if ((v16 ^ 42) <= v16) {
                                v40 := GenS3{F0: 100}
                                _ = v40
                            } else {
                                fmt.Println(fi20)
                            }
                            if true {
                                v19 = v21
                                v41 := GenS2{F0: false}
                                var v42 bool = v19
                                _ = v41
                                _ = v42
                            } else {
                                var v43 string = ((v4 + "") + ("\n\t" + v4))
                                for fi44 := 0; fi44 < 6; fi44++ {
                                    v43 = "\n\t"
                                    _ = fi44
                                }
                                v43 = ""
                                _ = v43
                            }
                            v4 = v4
                        }
                        v5 = v15
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    } else {
                        v19 = ((2 > 42) && v21)
                        var v45 uint64 = 612
                        _ = v45
                    }
                    v46 := GenS4{F0: 42}
                    _ = v46
                }
            }
            _ = len(v1)
            _ = fi20
            _ = v21
        }
        if v19 {
            var v47 int32 = 0
            v48 := GenS5{F0: v4}
            for fi49 := 0; fi49 < 12; fi49++ {
                v1[(v17 * (fi49 & (fi49 ^ (v17 ^ (v17 + (v17 ^ v17))))))] = true
                v50 := GenS3{F0: -1}
                v51 := GenS6{F0: 407}
                _ = fi49
                _ = v50
                _ = v51
            }
            _ = v47
            _ = v48
        } else {
            if false {
                var v52 uint16 = 0
                _ = v52
            }
            v3 = append(v3, (v17 << ((42) & 63)))
        }
        _ = v19
    }
    for fi53 := 0; fi53 < 18; fi53++ {
        fmt.Println(v17, v15)
        var v54 uint8 = 128
        var v55 float32 = 3.14159
        _ = fi53
        _ = v54
        _ = v55
    }
    if (56 <= v16) {
        var v56 int8 = 0
        _ = v56
    }
    var v57 uint16 = 866
    for fi58 := 0; fi58 < 10; fi58++ {
        var v59 uint = 0
        var v60 uint = v59
        _ = fi58
        _ = v59
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v8
}
