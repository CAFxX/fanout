package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint8; F2 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 0
    v3 := make(map[int]uint32)
    v4 := []int{32768, 643}
    var v5 uint = 255
    if true {
        v2 = (v2 - v2)
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 float64 = (((1e10 / -2.25) - 0.1) / 1.0)
            v7 -= v7
            _ = fi6
            _ = v7
        }
    }
    v8 := GenS1{F0: (v5 * v5)}
    if (2147483648 != 65535) {
        for fi9 := 0; fi9 < 12; fi9++ {
            _ = len(v4)
            var v10 uint16 = 32767
            _ = fi9
            _ = v10
        }
        fmt.Println(v2, v2)
    }
    var v11 uint64 = 4294967295
    var v12 int16 = -624
    fmt.Println(v2)
    if false {
        var v13 uint16 = 65535
        var v14 int = 0
        v15 := GenS1{F0: (v5 - 65535)}
        _ = v13
        _ = v14
        _ = v15
    } else {
        v2 -= -2147483648
        v16 := GenS1{F0: (v5 * (^(v5)))}
        var v17 float64 = 3.14159
        _ = v16
        _ = v17
    }
    switch "foo bar" {
    case "":
        v18 := GenS2{F0: true, F1: 239, F2: 0}
        _ = v18
    }
    if true {
        if (!(99 < 127)) {
            if false {
                v12 = 42
                if true {
                    var v19 uint64 = 2
                    _ = v19
                } else {
                    var v20 uint = (v5 * (v5 * v5))
                    _ = v20
                }
            }
        }
        v12 = -32768
    }
    var v21 int16 = v12
    if false {
        var v22 int64 = -338
        v23 := GenS2{F0: true, F1: 42, F2: (v2 / ((-1) | 1))}
        v4 = append(v4, 547)
        _ = v22
        _ = v23
    } else {
        if (!true) {
            _ = len(v4)
            for fi24 := 0; fi24 < 12; fi24++ {
                fmt.Println(v5, v2, v11)
                v25 := GenS2{F0: false, F1: 42, F2: (v2 + -2147483648)}
                switch ("foo bar" + ("hello" + "foo bar")) {
                case "":
                    switch "foo bar" {
                    case "a":
                        for fi26 := 0; fi26 < 16; fi26++ {
                            var v27 int = int(v2)
                            v27 += (v27 << ((v5) & 63))
                            fmt.Println(fi26)
                            _ = fi26
                            _ = v27
                        }
                    case "\n\t":
                        for fi28 := 0; fi28 < 7; fi28++ {
                            fmt.Println(v12)
                            _ = fi28
                        }
                    case "hello":
                        v5 = 776
                    default:
                        for fi29 := 0; fi29 < 20; fi29++ {
                            fmt.Println(v21, fi29, v12)
                            _ = fi29
                        }
                    }
                case "\n\t":
                    _ = len(v4)
                }
                _ = fi24
                _ = v25
            }
        } else {
            if (1 >= 2) {
                var v30 uint64 = v11
                if (true || ("hello" <= "")) {
                    switch ("\n\t" + ("foo bar" + "hello")) {
                    case "foo bar":
                        var v31 string = "foo bar"
                        _ = v31
                    case "a":
                        for fi32 := 0; fi32 < 4; fi32++ {
                            for fi33 := 0; fi33 < 18; fi33++ {
                                switch (^(fi32)) {
                                case -962:
                                    switch "\n\t" {
                                    case "":
                                        var v34 uint = (v5 & 255)
                                        _ = v34
                                    }
                                case -92:
                                    delete(v3, (fi33 + (fi33 + 128)))
                                case 128:
                                    if false {
                                        var v35 uint = 545
                                        var v36 int = (fi32 & fi32)
                                        _ = v35
                                        _ = v36
                                    } else {
                                        var v37 int = 127
                                        if false {
                                            switch v37 {
                                            case -76:
                                                var v38 uint32 = 65535
                                                _ = v38
                                            default:
                                                delete(v1, (-(fi33)))
                                            }
                                        }
                                        _ = v37
                                    }
                                default:
                                    var v39 float32 = 3.14159
                                    _ = v39
                                }
                                var v40 int8 = 0
                                _ = fi33
                                _ = v40
                            }
                            if false {
                                v41 := GenS2{F0: ((true && (38 > (fi32 / (((^(fi32))) | 1)))) && ((0 < 65535) && (-9223372036854775808 != -1))), F1: 1, F2: v2}
                                for fi42 := 0; fi42 < 5; fi42++ {
                                    fmt.Println(fi42, v30, v12)
                                    _ = fi42
                                }
                                _ = v41
                            } else {
                                var v43 float64 = 1.5
                                fmt.Println(v43)
                                _ = v43
                            }
                            var v44 float64 = float64(1.5)
                            _ = fi32
                            _ = v44
                        }
                    case "hello":
                        var v45 int64 = -1
                        _ = v45
                    }
                    v46 := GenS1{F0: v5}
                    var v47 uint32 = 2
                    _ = v46
                    _ = v47
                }
                _ = v30
            }
        }
        var v48 uint64 = v11
        _ = v48
    }
    v49 := GenS2{F0: false, F1: 2, F2: 0}
    if (911 >= 170) {
        fmt.Println(v12)
        switch 42 {
        case 32767:
            v50 := GenS1{F0: v5}
            _ = v50
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v8
}
