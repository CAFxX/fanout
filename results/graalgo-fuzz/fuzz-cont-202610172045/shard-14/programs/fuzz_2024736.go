package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 int32 }
func genFn0(p0 bool, p1 int16, p2 int) uint16 {
    var v5 uint64 = 42
    var v6 uint32 = 611
    p1 = p1
    for fi7 := 0; fi7 < 1; fi7++ {
        if p0 {
            if ((!p0) && (!p0)) {
                var v8 bool = (((3.14159 + 1.5) - -2.25) >= 0.0)
                v6 = (v6 - 2147483647)
                v9 := GenS1{F0: -90}
                _ = v8
                _ = v9
            }
            switch "a" {
            case "":
                var v10 float32 = (((1.5 / 1e10) - 1.5) * 3.14159)
                _ = v10
            case "hello":
                var v11 float32 = (1.5 / ((1.5 + (0.1 / ((((0.0 * 1.5) * (((-2.25 / 0.1) - 1.5) - 0.1)) / -2.25) + (float32(float64((((0.0 / 1e10) / 1e10) * (1.5 - 1e10)))) - ((float32((3.14159 / 1e10)) / 1.0) - 1.5))))) / -2.25))
                _ = v11
            case "foo bar":
                v12 := GenS1{F0: -8}
                _ = v12
            default:
                for fi13 := 0; fi13 < 15; fi13++ {
                    if p0 {
                        if (715 < 9223372036854775807) {
                            var v14 float64 = 0.0
                            if (p0 && false) {
                                p1 = (-(p1))
                                if false {
                                    if p0 {
                                        p0 = true
                                    }
                                    v15 := GenS1{F0: -1}
                                    var v16 uint = 4294967295
                                    _ = v15
                                    _ = v16
                                }
                                p2 += fi7
                            } else {
                                var v17 float32 = 0.0
                                v5 -= 0
                                if false {
                                    var v18 uint16 = 2
                                    _ = v18
                                } else {
                                    var v19 float32 = 3.14159
                                    _ = v19
                                }
                                if (!p0) {
                                    if p0 {
                                        v20 := GenS1{F0: 127}
                                        var v21 string = "a"
                                        _ = v20
                                        _ = v21
                                    }
                                    v17 -= v17
                                } else {
                                    p1 -= 0
                                    v6 -= (v6 + v6)
                                }
                                v22 := GenS2{F0: (p1 - (p1 / (((p1 | p1)) | 1)))}
                                _ = v17
                                _ = v22
                            }
                            v23 := []int{-1, -9223372036854775808, 32768, 626}
                            _ = v14
                            _ = v23
                        } else {
                            v24 := []int{127, -9223372036854775808, -230}
                            fmt.Println(p0, p2, v5)
                            _ = v24
                        }
                        v25 := GenS1{F0: -1}
                        var v26 int = 42
                        _ = v25
                        _ = v26
                    } else {
                        p1 = (p1 + (p1 / (((p1 << ((996) & 63))) | 1)))
                        v27 := GenS3{F0: uint32(v5)}
                        _ = v27
                    }
                    _ = fi13
                }
            }
        } else {
            switch "foo bar" {
            case "\n\t":
                for fi28 := 0; fi28 < 14; fi28++ {
                    var v29 int32 = int32(v5)
                    v30 := []int{(fi7 - (-(fi7)))}
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            default:
                var v31 int64 = 350
                _ = v31
            }
            v5 -= v5
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 255
}
func genFn1() int32 {
    var v32 int = 32767
    if true {
        v33 := GenS1{F0: -92}
        _ = v33
    }
    var v34 uint16 = 378
    v35 := GenS2{F0: -755}
    var v36 bool = (true || false)
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    return -1
}

func main() {
    var v1 uint8 = 129
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 709
    v4 := make(map[string]string)
    switch "foo bar" {
    case "":
        switch 128 {
        case 127:
            fmt.Println(v1, v1)
        case 0:
            var v37 int64 = 1
            _ = v37
        }
    default:
        var v38 bool = false
        _ = v38
    }
    var v39 uint = 895
    if false {
        fmt.Println(v3, v3, v1)
        _ = len(v4)
        for fi40 := 0; fi40 < 19; fi40++ {
            v41 := GenS3{F0: 2}
            v42 := GenS3{F0: 42}
            _ = fi40
            _ = v41
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 15; fi43++ {
        v39 = (v39 << ((v39) & 63))
        var v44 uint64 = (^(v3))
        switch "\n\t" {
        case "":
            v44 = 985
        case "a":
            v1 = v1
        case "\n\t":
            switch ("" + "foo bar") {
            case "foo bar":
                var v45 uint8 = v1
                _ = v45
            case "a":
                v46 := GenS1{F0: -128}
                _ = v46
            }
        }
        _ = fi43
        _ = v44
    }
    fmt.Println(v39, v39)
    for fi47 := 0; fi47 < 5; fi47++ {
        if false {
            if false {
                fmt.Println(v3)
                v39 = (v39 | v39)
                if true {
                    v48 := v4["\n\t"]
                    for fi49 := 0; fi49 < 4; fi49++ {
                        v3 += 2
                        var v50 int8 = 100
                        v51 := []int{(fi49 * fi49), (fi47 - (^(fi47)))}
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    _ = v48
                }
            } else {
                v39 -= v39
                var v52 uint8 = v1
                for fi53 := 0; fi53 < 1; fi53++ {
                    fmt.Println(v3)
                    var v54 int32 = -2147483648
                    _ = fi53
                    _ = v54
                }
                _ = v52
            }
            switch "" {
            case "":
                v55 := GenS4{F0: (^(v1)), F1: 866, F2: (v1 - v1)}
                _ = v55
            case "a":
                v56 := GenS5{F0: v1, F1: 255}
                _ = v56
            }
            v39 = v39
        } else {
            v57 := GenS6{F0: -32768}
            _ = v57
        }
        v58 := GenS7{F0: 2147483647, F1: int32(v39)}
        _ = fi47
        _ = v58
    }
    v3 = 766
    v59 := GenS3{F0: 2147483648}
    var v60 uint8 = (v1 - uint8(v3))
    var v61 uint64 = v3
    switch -577 {
    case 42:
        var v62 int16 = -32768
        _ = v62
    case 127:
        fmt.Println(v1)
    case -2147483648:
        var v63 int = 2147483647
        _ = v63
    default:
        switch (("" + "\n\t") + ("a" + "\n\t")) {
        case "hello":
            var v64 int16 = 1
            _ = v64
        }
    }
    v3 = v61
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v59
    _ = v60
    _ = v61
}
