package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 int32; F2 int8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{}
    var v3 int8 = -128
    var v4 uint8 = 128
    v4 = (v4 & v4)
    var v5 uint64 = 4294967295
    if true {
        var v6 uint = 4294967295
        v4 += uint8(v6)
        if true {
            var v7 uint64 = 2
            var v8 uint = (v6 + (v6 / (((v6 & 65535)) | 1)))
            switch "hello" {
            case "\n\t":
                v7 += (^(v5))
            case "":
                v8 += uint(v7)
            }
            _ = v7
            _ = v8
        }
        _ = v6
    } else {
        var v9 uint = 64
        _ = v9
    }
    v10 := GenS1{F0: 0, F1: 42, F2: 2147483647}
    v5 = v5
    _ = len(v2)
    if true {
        var v11 int8 = (v3 - 100)
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 uint8 = 1
            _ = fi12
            _ = v13
        }
        _ = v11
    } else {
        var v14 string = ("" + "\n\t")
        v4 = v4
        _ = v14
    }
    v15 := GenS1{F0: -1, F1: 0, F2: 1}
    _ = len(v1)
    for fi16 := 0; fi16 < 1; fi16++ {
        switch "\n\t" {
        case "":
            fmt.Println(v3)
        case "foo bar":
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 int64 = 42
                _ = fi17
                _ = v18
            }
        case "a":
            if true {
                switch "a" {
                case "\n\t":
                    if false {
                        var v19 uint32 = 605
                        _ = v19
                    } else {
                        if false {
                            _ = len(v2)
                            var v20 uint64 = 42
                            _ = v20
                        }
                        v5 = (^(v5))
                    }
                case "hello":
                    if false {
                        fmt.Println(v3, v4)
                    } else {
                        var v21 uint32 = 65535
                        _ = v21
                    }
                case "":
                    if len(v2) > 0 {
                        v2[0] = fi16
                    }
                default:
                    v1[4294967295] = fi16
                }
            }
        }
        _ = fi16
    }
    v4 = 255
    v22 := GenS2{F0: -128, F1: 65535}
    for fi23 := 0; fi23 < 3; fi23++ {
        v24 := GenS2{F0: (^(v3)), F1: 0}
        v5 = (v5 & (v5 | (^(v5))))
        v25 := GenS1{F0: -1, F1: 42, F2: 65535}
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 int16 = 32767
    for fi27 := 0; fi27 < 3; fi27++ {
        switch (fi27 * fi27) {
        case -2147483648:
            for fi28 := 0; fi28 < 15; fi28++ {
                v29 := v1[2147483647]
                var v30 uint = 2
                var v31 float64 = 0.1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi28
                _ = v29
                _ = v30
                _ = v31
            }
        case 32768:
            v26 -= v26
        case -795:
            switch 9223372036854775807 {
            case -9223372036854775808:
                for fi32 := 0; fi32 < 18; fi32++ {
                    if true {
                        v3 += 127
                        if true {
                            v33 := GenS1{F0: 9223372036854775807, F1: (-(v26)), F2: 65535}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if (fi27 <= (fi27 << ((0) & 63))) {
                                if true {
                                    var v34 uint16 = uint16(v5)
                                    _ = v34
                                } else {
                                    var v35 float64 = float64(3.14159)
                                    var v36 uint16 = uint16(v4)
                                    var v37 int8 = v3
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                switch (fi27 * (fi32 & fi32)) {
                                case 32767:
                                    v38 := GenS2{F0: int8(v4), F1: 0}
                                    _ = v38
                                case 1:
                                    delete(v1, uint32(v4))
                                }
                            } else {
                                v4 = v4
                                switch "\n\t" {
                                case "":
                                    v4 = v4
                                case "a":
                                    var v39 uint16 = 1
                                    _ = v39
                                case "\n\t":
                                    fmt.Println(v26, fi27)
                                default:
                                    var v40 bool = ((255 >= 795) && (false || true))
                                    _ = v40
                                }
                            }
                            _ = v33
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        v41 := GenS1{F0: -14, F1: (-(v26)), F2: 2147483647}
                        _ = v41
                    }
                    v42 := GenS3{F0: 517}
                    switch fi32 {
                    case 0:
                        if ((fi32 / (((fi32 * (fi27 + fi32))) | 1)) >= fi27) {
                            var v43 string = "a"
                            _ = v43
                        }
                    case 308:
                        switch 2147483647 {
                        case 42:
                            var v44 int32 = 7
                            _ = v44
                        }
                    }
                    fmt.Println(v3, fi32, v26)
                    _ = fi32
                    _ = v42
                }
            }
        default:
            if (0 >= 42) {
                if ((-(fi27)) >= fi27) {
                    if (!((!(65535 == 923)) && false)) {
                        v45 := v1[1]
                        var v46 uint = 2
                        if true {
                            v5 = v5
                            var v47 int8 = (v3 & (-(v3)))
                            _ = v47
                        }
                        _ = v45
                        _ = v46
                    } else {
                        v5 = (v5 % ((0) | 1))
                        var v48 uint32 = 42
                        var v49 uint8 = v4
                        _ = v48
                        _ = v49
                    }
                }
            }
        }
        switch 1 {
        case 32768:
            _ = len(v1)
        case 32767:
            var v50 uint32 = 2
            _ = v50
        }
        _ = fi27
    }
    v51 := GenS4{F0: 1, F1: 0, F2: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v51
}
