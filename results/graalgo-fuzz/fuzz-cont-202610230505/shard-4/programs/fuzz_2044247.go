package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint32 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 uint16 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{4294967295}
    var v3 int32 = 1
    var v4 int32 = (v3 - v3)
    var v5 uint32 = 734
    v5 -= v5
    v6 := GenS1{F0: (32768 == -9223372036854775808), F1: "foo bar", F2: (v5 ^ (v5 - v5))}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint8 = 128
        v9 := GenS2{F0: "hello", F1: -67}
        for fi10 := 0; fi10 < 6; fi10++ {
            v3 = (^(v3))
            _ = fi10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        switch fi11 {
        case -9223372036854775808:
            if (!false) {
                switch fi11 {
                case -902:
                    for fi12 := 0; fi12 < 13; fi12++ {
                        if false {
                            var v13 uint8 = 127
                            _ = v13
                        } else {
                            var v14 uint16 = 255
                            v1 = v1
                            var v15 float32 = 0.0
                            _ = v14
                            _ = v15
                        }
                        if (460 <= 9223372036854775807) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        } else {
                            _ = len(v2)
                            fmt.Println(fi12, v4, fi12)
                        }
                        _ = fi12
                    }
                case 32768:
                    var v16 uint32 = (^(v5))
                    _ = v16
                default:
                    v17 := GenS1{F0: false, F1: "hello", F2: v5}
                    _ = v17
                }
                switch (-(fi11)) {
                case -802:
                    switch ("a" + "\n\t") {
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "\n\t":
                        for fi18 := 0; fi18 < 4; fi18++ {
                            if ("\n\t" > "a") {
                                var v19 int8 = 0
                                v1 = v1
                                _ = v19
                            } else {
                                if (fi11 != (fi11 - fi11)) {
                                    switch fi11 {
                                    case 4294967295:
                                        var v20 string = "a"
                                        _ = v20
                                    case 254:
                                        v3 = v3
                                    }
                                    var v21 int32 = v3
                                    var v22 uint8 = 217
                                    _ = v21
                                    _ = v22
                                }
                            }
                            v5 -= (^(v5))
                            _ = fi18
                        }
                    default:
                        var v23 string = "a"
                        _ = v23
                    }
                case 9:
                    var v24 bool = true
                    _ = v24
                default:
                    fmt.Println(v3, fi11)
                }
            }
        default:
            switch (("foo bar" + "\n\t") + "") {
            case "hello":
                v25 := GenS3{F0: -111}
                _ = v25
            case "a":
                if (!(fi11 < fi11)) {
                    switch -646 {
                    case 128:
                        if (42 != 127) {
                            fmt.Println(v1, v1)
                            for fi26 := 0; fi26 < 2; fi26++ {
                                switch fi11 {
                                case 632:
                                    var v27 uint = 0
                                    _ = v27
                                case -496:
                                    var v28 uint8 = 158
                                    _ = v28
                                case -9223372036854775808:
                                    v4 -= v3
                                }
                                v2 = append(v2, 2147483648)
                                var v29 uint32 = 2147483647
                                switch fi26 {
                                case 4294967295:
                                    fmt.Println(v3, v29)
                                }
                                _ = fi26
                                _ = v29
                            }
                        }
                    case -788:
                        v30 := GenS3{F0: -110}
                        _ = v30
                    default:
                        var v31 float64 = float64(0.0)
                        _ = v31
                    }
                }
            case "":
                fmt.Println(v4, v1, v1)
            default:
                var v32 bool = true
                _ = v32
            }
        }
        _ = fi11
    }
    var v33 bool = true
    for fi34 := 0; fi34 < 4; fi34++ {
        fmt.Println(v4, v3, v5)
        var v35 float64 = 0.0
        _ = fi34
        _ = v35
    }
    _ = len(v2)
    v5 -= v5
    if ((1 < 1) || (-111 == 1)) {
        v36 := GenS4{F0: -353, F1: -32768}
        v33 = (0.1 == 1e10)
        var v37 int32 = (-(v3))
        _ = v36
        _ = v37
    } else {
        v1 = (^(v1))
        if (1.5 >= 3.14159) {
            for fi38 := 0; fi38 < 16; fi38++ {
                v5 = v5
                v33 = (v33 && (-32768 == -32768))
                _ = fi38
            }
            if (!v33) {
                if v33 {
                    var v39 int32 = (v4 | 1)
                    fmt.Println(v4, v39, v5)
                    _ = v39
                }
            }
            var v40 uint = 18446744073709551615
            _ = v40
        } else {
            fmt.Println(v3)
            v3 = (^(v3))
            switch "" {
            case "hello":
                v41 := GenS4{F0: 127, F1: -1}
                _ = v41
            default:
                fmt.Println(v33, v4, v1)
            }
        }
        if (255 == 1) {
            for fi42 := 0; fi42 < 4; fi42++ {
                for fi43 := 0; fi43 < 9; fi43++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi43
                }
                var v44 int16 = 42
                _ = fi42
                _ = v44
            }
            v45 := GenS3{F0: 1}
            _ = v45
        }
    }
    var v46 string = "a"
    _ = len(v2)
    v5 += 2147483648
    var v47 int64 = 0
    if (1 > v47) {
        v48 := GenS5{F0: 65535}
        if (257 != 32767) {
            var v49 uint16 = 32767
            for fi50 := 0; fi50 < 4; fi50++ {
                var v51 float64 = ((0.0 + -2.25) - (0.0 - 1.5))
                _ = fi50
                _ = v51
            }
            _ = v49
        }
        _ = v48
    } else {
        if v33 {
            for fi52 := 0; fi52 < 12; fi52++ {
                v1 -= v1
                v53 := GenS2{F0: "\n\t", F1: 22}
                var v54 int16 = -730
                _ = fi52
                _ = v53
                _ = v54
            }
            _ = len(v2)
        } else {
            v33 = false
            for fi55 := 0; fi55 < 10; fi55++ {
                _ = len(v2)
                _ = fi55
            }
            switch 302 {
            case -9223372036854775808:
                v56 := GenS5{F0: 32767}
                _ = v56
            case -2147483648:
                var v57 uint16 = 255
                _ = v57
            case 1:
                if len(v2) > 0 {
                    v2[0] = v5
                }
            }
        }
        fmt.Println(v47)
    }
    var v58 string = v46
    fmt.Println(v5, v1, v4)
    if v33 {
        var v59 float32 = 1e10
        v46 = "\n\t"
        _ = v59
    }
    var v60 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v58
    _ = v6
    _ = v60
}
