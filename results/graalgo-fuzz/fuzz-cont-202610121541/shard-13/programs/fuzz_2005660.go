package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 uint; F2 int8 }
type GenS5 struct { F0 int; F1 int8; F2 uint32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string; F1 int16 }
func genFn0(p0 uint) bool {
    var v7 int64 = 969
    v8 := GenS1{F0: 128}
    v9 := []int{}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return true
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 65535
    v3 := make(map[string]int)
    var v4 int32 = -891
    var v5 bool = false
    var v6 int8 = -128
    for fi10 := 0; fi10 < 8; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS1{F0: 128}
            if (!v5) {
                v5 = (128 < 2)
                delete(v1, "foo bar")
            } else {
                var v13 float64 = (float64(3.14159) * (((0.0 / 3.14159) + 3.14159) * -2.25))
                v5 = v5
                _ = v13
            }
            v14 := GenS2{F0: (fi10 << ((18446744073709551615) & 63)), F1: 132}
            _ = fi11
            _ = v12
            _ = v14
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            if true {
                var v16 bool = (127 != 0)
                var v17 bool = (0 >= -1)
                v18 := GenS2{F0: (fi15 - fi10), F1: 2}
                _ = v16
                _ = v17
                _ = v18
            } else {
                v6 = (v6 + 100)
                switch -2147483648 {
                case 127:
                    var v19 int = 0
                    _ = v19
                }
            }
            v6 = 116
            var v20 uint = (^(v2))
            v21 := GenS2{F0: int(v4), F1: 1}
            var v22 uint16 = 32767
            _ = fi15
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi10
    }
    switch int(v2) {
    case 0:
        v23 := GenS1{F0: 0}
        _ = v23
    case 128:
        var v24 int64 = 0
        _ = v24
    default:
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 uint = v2
            v2 -= (v2 + v2)
            if true {
                var v27 float32 = 1.5
                switch "hello" {
                case "foo bar":
                    var v28 uint64 = 422
                    _ = v28
                case "a":
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 uint64 = 94
                        v31 := GenS3{F0: 94}
                        if (!v5) {
                            if false {
                                var v32 float32 = v27
                                _ = v32
                            } else {
                                v33 := GenS3{F0: 1}
                                var v34 float64 = 1.5
                                v6 += -128
                                _ = v33
                                _ = v34
                            }
                            v30 = v30
                            v4 -= v4
                        } else {
                            var v35 uint16 = 42
                            v26 = v2
                            _ = v35
                        }
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                case "":
                    fmt.Println(v26, fi25)
                default:
                    v2 += 878
                }
                _ = v27
            } else {
                for fi36 := 0; fi36 < 18; fi36++ {
                    switch "" {
                    case "":
                        var v37 uint = (v2 & v26)
                        _ = v37
                    }
                    v38 := GenS2{F0: (fi36 / ((fi36) | 1)), F1: 255}
                    _ = fi36
                    _ = v38
                }
                v39 := []int{127, 1}
                _ = v39
            }
            _ = fi25
            _ = v26
        }
    }
    v5 = (((float32(1e10) / (1e10 - (0.0 - 1e10))) / ((3.14159 * ((((float32(1e10) - 0.1) + -2.25) * 0.0) + float32((1.5 / ((-2.25 + (1.5 - 1.5)) / 3.14159))))) + (1e10 + (3.14159 / (0.1 - float32(0.0)))))) <= -2.25)
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        for fi40 := 0; fi40 < 7; fi40++ {
            var v41 uint = (^(v2))
            _ = fi40
            _ = v41
        }
    case "":
        var v42 bool = true
        _ = v42
    }
    for fi43 := 0; fi43 < 15; fi43++ {
        switch ("" + ("a" + "\n\t")) {
        case "hello":
            fmt.Println(fi43)
        default:
            fmt.Println(v2, v6, v4)
        }
        _ = fi43
    }
    var v44 int = -801
    switch "\n\t" {
    case "\n\t":
        v4 -= (v4 / ((v4) | 1))
    }
    var v45 int64 = 1
    switch "a" {
    case "":
        v6 = -19
    case "\n\t":
        if (299 == 0) {
            if (float32(-2.25) <= 1e10) {
                var v46 uint8 = 1
                switch "\n\t" {
                case "a":
                    for fi47 := 0; fi47 < 4; fi47++ {
                        v48 := GenS2{F0: fi47, F1: (v46 | (v46 << ((v2) & 63)))}
                        v49 := GenS4{F0: "foo bar", F1: 347, F2: 60}
                        var v50 uint = uint(v46)
                        _ = fi47
                        _ = v48
                        _ = v49
                        _ = v50
                    }
                case "":
                    v2 = (^(v2))
                default:
                    v51 := GenS2{F0: (v44 & v44), F1: v46}
                    _ = v51
                }
                for fi52 := 0; fi52 < 18; fi52++ {
                    v44 = (fi52 * (fi52 | 4294967295))
                    _ = len(v1)
                    switch 32768 {
                    case 2147483647:
                        v1[""] = 4294967295
                    default:
                        var v53 uint16 = 255
                        _ = v53
                    }
                    _ = fi52
                }
                _ = v46
            } else {
                fmt.Println(v5, v5)
                var v54 string = "hello"
                _ = len(v3)
                _ = v54
            }
            v6 = v6
            fmt.Println(v44, v6)
        }
    case "hello":
        var v55 float64 = (3.14159 * ((1.5 - (1e10 * ((1.5 - 0.0) / 3.14159))) - -2.25))
        _ = v55
    default:
        var v56 float32 = 0.1
        _ = v56
    }
    v57 := []int{v44, 877}
    v4 += (^(v4))
    v58 := GenS5{F0: 32768, F1: (v6 ^ v6), F2: 0}
    _ = len(v1)
    switch ("" + "foo bar") {
    case "hello":
        if v5 {
            v5 = false
            v2 = (v2 + (^(v2)))
            v3[((("" + (("a" + (("a" + (("" + "hello") + (((("a" + ((("foo bar" + "") + "\n\t") + "hello")) + "foo bar") + (("" + ("hello" + (("\n\t" + ((("a" + "foo bar") + ("hello" + ("hello" + (("\n\t" + ("hello" + (("a" + "hello") + "a"))) + ("a" + ("hello" + "foo bar")))))) + "")) + "a"))) + "\n\t")) + "foo bar"))) + ("a" + ("a" + "a")))) + "hello")) + "foo bar") + "a")] = v44
        } else {
            var v59 uint16 = 32767
            v59 = v59
            _ = v59
        }
    case "foo bar":
        switch "foo bar" {
        case "a":
            v60 := GenS3{F0: 42}
            _ = v60
        default:
            for fi61 := 0; fi61 < 3; fi61++ {
                if len(v57) >= 2 {
                    v57 = v57[1:len(v57)]
                }
                for fi62 := 0; fi62 < 13; fi62++ {
                    v63 := GenS5{F0: (v44 / ((1) | 1)), F1: 0, F2: 202}
                    var v64 float64 = 1e10
                    v65 := GenS6{F0: 28}
                    _ = fi62
                    _ = v63
                    _ = v64
                    _ = v65
                }
                _ = fi61
            }
        }
    }
    fmt.Println(v2, v45)
    for fi66 := 0; fi66 < 19; fi66++ {
        switch v44 {
        case 2147483647:
            var v67 int8 = (v6 & v6)
            _ = v67
        case 485:
            var v68 uint16 = 0
            _ = v68
        }
        for fi69 := 0; fi69 < 11; fi69++ {
            if false {
                v70 := GenS7{F0: -1}
                var v71 uint = v2
                _ = v70
                _ = v71
            } else {
                fmt.Println(v4, v2)
                v72 := GenS8{F0: "a", F1: -982}
                _ = v72
            }
            _ = fi69
        }
        _ = fi66
    }
    for fi73 := 0; fi73 < 13; fi73++ {
        var v74 float32 = (0.0 + -2.25)
        _ = len(v3)
        _ = fi73
        _ = v74
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
