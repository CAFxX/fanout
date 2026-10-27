package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 string; F2 int16 }
type GenS4 struct { F0 string; F1 uint8; F2 string }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 int; F1 uint32; F2 uint64 }
type GenS7 struct { F0 int16; F1 bool; F2 int }

func main() {
    var v1 uint32 = 442
    var v2 int = -1
    var v3 int = -873
    var v4 int32 = 837
    v5 := []int{v2, 0}
    v6 := GenS1{F0: 42, F1: 655}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi7 := 0; fi7 < 3; fi7++ {
        _ = len(v5)
        var v8 bool = true
        _ = fi7
        _ = v8
    }
    switch v2 {
    case 2147483647:
        for fi9 := 0; fi9 < 12; fi9++ {
            switch fi9 {
            case 625:
                v10 := GenS2{F0: ("\n\t" + "hello")}
                _ = v10
            case 740:
                if len(v5) > 0 {
                    v5[0] = v2
                }
            default:
                v11 := GenS2{F0: "hello"}
                _ = v11
            }
            if (0 <= 255) {
                if (fi9 == (^(v3))) {
                    v4 -= v4
                }
                var v12 string = "foo bar"
                var v13 int8 = -110
                _ = v12
                _ = v13
            } else {
                for fi14 := 0; fi14 < 3; fi14++ {
                    v2 += 385
                    v15 := GenS3{F0: 65535, F1: "", F2: 42}
                    _ = fi14
                    _ = v15
                }
                for fi16 := 0; fi16 < 5; fi16++ {
                    if false {
                        var v17 bool = ("foo bar" < "\n\t")
                        var v18 int = v3
                        _ = v17
                        _ = v18
                    } else {
                        var v19 int8 = 0
                        _ = v19
                    }
                    _ = fi16
                }
                var v20 float64 = (1e10 - (0.1 / 3.14159))
                _ = v20
            }
            switch (("a" + "") + "foo bar") {
            case "a":
                for fi21 := 0; fi21 < 14; fi21++ {
                    switch "\n\t" {
                    case "":
                        var v22 int8 = 1
                        _ = v22
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    fmt.Println(v1, v4)
                    var v23 uint = 335
                    v24 := GenS4{F0: "foo bar", F1: 217, F2: ""}
                    v3 = 4294967295
                    var v25 uint64 = 18446744073709551615
                    v2 += (fi21 >> ((18446744073709551615) & 63))
                    _ = fi21
                    _ = v23
                    _ = v24
                    _ = v25
                }
            case "hello":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "foo bar":
                if false {
                    v26 := GenS5{F0: 128, F1: (v4 & v4)}
                    _ = v26
                } else {
                    var v27 uint32 = (^(v1))
                    var v28 int16 = 1
                    _ = v27
                    _ = v28
                }
            }
            _ = fi9
        }
    default:
        var v29 int = v2
        _ = v29
    }
    v2 -= v2
    var v30 string = "\n\t"
    switch 32767 {
    case 272:
        if (!false) {
            if (!true) {
                var v31 float32 = 3.14159
                var v32 uint8 = 128
                v33 := GenS1{F0: v32, F1: -1}
                _ = v31
                _ = v32
                _ = v33
            }
            for fi34 := 0; fi34 < 12; fi34++ {
                v35 := GenS6{F0: 128, F1: 2147483647, F2: 342}
                _ = fi34
                _ = v35
            }
            if ((v30 >= "foo bar") || (-9223372036854775808 >= -1)) {
                if false {
                    fmt.Println(v3, v2)
                    for fi36 := 0; fi36 < 17; fi36++ {
                        switch v30 {
                        case "":
                            var v37 string = ((v30 + v30) + v30)
                            _ = v37
                        case "\n\t":
                            var v38 uint32 = v1
                            _ = v38
                        case "foo bar":
                            fmt.Println(v3)
                        default:
                            var v39 int64 = 9223372036854775807
                            _ = v39
                        }
                        if (!true) {
                            v40 := GenS1{F0: 52, F1: 42}
                            _ = v40
                        } else {
                            _ = len(v5)
                            v41 := GenS5{F0: 1, F1: int32(fi36)}
                            _ = v41
                        }
                        _ = fi36
                    }
                }
                if true {
                    var v42 uint32 = v1
                    for fi43 := 0; fi43 < 2; fi43++ {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi43
                    }
                    _ = v42
                } else {
                    v4 = v4
                    v3 += (v2 - v2)
                }
                fmt.Println(v1, v2)
            } else {
                var v44 int = v2
                if (77 != 0) {
                    switch ("foo bar" + "\n\t") {
                    case "foo bar":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "":
                        v45 := GenS7{F0: -1, F1: ((255 == 133) && (65535 > 2)), F2: (v44 + (v2 * v2))}
                        _ = v45
                    case "hello":
                        for fi46 := 0; fi46 < 14; fi46++ {
                            v1 = 392
                            v44 = v44
                            switch 9223372036854775807 {
                            case 127:
                                var v47 int64 = 0
                                _ = v47
                            case 128:
                                _ = len(v5)
                            case 1:
                                fmt.Println(v44, v4, v2)
                            }
                            _ = fi46
                        }
                    default:
                        v1 -= (v1 & v1)
                    }
                    v30 = v30
                }
                _ = v44
            }
        } else {
            fmt.Println(v3, v2)
        }
    case 2147483647:
        if len(v5) > 0 {
            v5[0] = v3
        }
    case -9223372036854775808:
        fmt.Println(v1)
    }
    var v48 float64 = 1e10
    v5 = append(v5, (^(v2)))
    v30 = v30
    if (0 == v1) {
        var v49 int16 = 714
        _ = v49
    } else {
        if len(v5) > 0 {
            v5[0] = (-(v2))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
