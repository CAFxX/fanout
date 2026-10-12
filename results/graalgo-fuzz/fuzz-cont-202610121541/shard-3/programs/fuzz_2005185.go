package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 uint8 }
func genFn0(p0 uint8, p1 int8, p2 int64) int32 {
    v7 := GenS1{F0: 0}
    fmt.Println(p1)
    v8 := GenS1{F0: -32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 uint64) bool {
    v9 := []int{616, 9223372036854775807}
    switch "hello" {
    case "a":
        fmt.Println(p0)
    default:
        _ = len(v9)
    }
    var v10 int32 = 0
    _ = len(v9)
    switch 9223372036854775807 {
    case 961:
        if len(v9) > 0 {
            v9[0] = 622
        }
    case 4294967295:
        p0 += (p0 << ((4294967295) & 63))
    case 9223372036854775807:
        if false {
            if len(v9) > 0 {
                v9[0] = 127
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            v11 := GenS1{F0: 1000}
            _ = v11
        }
    }
    _ = p0
    _ = v10
    _ = v9
    return false
}

func main() {
    var v1 uint16 = 42
    var v2 string = "foo bar"
    var v3 uint = 564
    v4 := make(map[uint32]uint32)
    v5 := []uint32{692, 65535, 0, 2147483648}
    var v6 string = v2
    var v12 uint32 = 2147483648
    v3 -= uint(v1)
    _ = len(v5)
    v13 := GenS1{F0: 42}
    switch 538 {
    case -2147483648:
        var v14 uint64 = uint64(v1)
        _ = v14
    case 32767:
        v4[993] = 0
    case -9223372036854775808:
        fmt.Println(v12)
    }
    v4[v12] = 472
    v5 = append(v5, 1)
    if (!false) {
        var v15 uint16 = v1
        _ = v15
    } else {
        if true {
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 uint64 = 499
                for fi18 := 0; fi18 < 12; fi18++ {
                    var v19 uint32 = v12
                    fmt.Println(v12, v17, v3)
                    _ = fi18
                    _ = v19
                }
                v20 := v4[704]
                _ = fi16
                _ = v17
                _ = v20
            }
        } else {
            var v21 float32 = (1.5 - (float32(1.5) / 1e10))
            v4[v12] = (v12 * (v12 << ((v3) & 63)))
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 int64 = -9223372036854775808
                _ = fi22
                _ = v23
            }
            _ = v21
        }
    }
    switch 2147483647 {
    case -168:
        if false {
            switch 127 {
            case -720:
                v12 -= v12
            case 0:
                v6 = "hello"
            }
        }
    case 42:
        switch v2 {
        case "":
            switch (v2 + (v2 + v2)) {
            case "":
                if (1 <= 32768) {
                    v24 := v4[v12]
                    _ = v24
                }
            case "\n\t":
                v6 = (v2 + (v6 + (v2 + "hello")))
            case "a":
                for fi25 := 0; fi25 < 6; fi25++ {
                    if true {
                        v26 := GenS1{F0: -846}
                        if ((0.0 - 1e10) > 0.1) {
                            switch fi25 {
                            case 32767:
                                fmt.Println(v6, v1, v3)
                            default:
                                for fi27 := 0; fi27 < 12; fi27++ {
                                    switch (fi27 | fi27) {
                                    case -1:
                                        if (!(!false)) {
                                            if false {
                                                v12 = 42
                                            } else {
                                                v4[v12] = v12
                                            }
                                        } else {
                                            v28 := GenS2{F0: 108}
                                            switch (^(fi27)) {
                                            case 32768:
                                                v2 = "a"
                                            case 1:
                                                v29 := GenS2{F0: 2147483647}
                                                _ = v29
                                            }
                                            _ = v28
                                        }
                                    case -400:
                                        var v30 float64 = -2.25
                                        _ = v30
                                    case -669:
                                        v2 = (v6 + (v6 + v2))
                                    }
                                    if true {
                                        _ = len(v5)
                                        v12 = v12
                                        var v31 float64 = ((0.1 / 1.0) - 1.5)
                                        _ = v31
                                    } else {
                                        fmt.Println(fi25)
                                    }
                                    fmt.Println(v1, fi25, v6)
                                    v32 := GenS3{F0: 128}
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    switch fi27 {
                                    case -640:
                                        v33 := GenS4{F0: (fi27 / ((fi27) | 1)), F1: 153}
                                        _ = v33
                                    }
                                    v6 = (v2 + v6)
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    v2 = (v2 + (("hello" + "a") + (v6 + v2)))
                                    v3 = (v3 / ((545) | 1))
                                    if true {
                                        v34 := GenS3{F0: 9223372036854775807}
                                        var v35 int = int(v3)
                                        v4[v12] = v12
                                        var v36 float64 = 0.0
                                        v12 = 511
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                    v37 := GenS3{F0: 128}
                                    _ = fi27
                                    _ = v32
                                    _ = v37
                                }
                            }
                        } else {
                            v38 := v4[v12]
                            switch fi25 {
                            case 9223372036854775807:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case 253:
                                v39 := GenS1{F0: -234}
                                _ = v39
                            default:
                                var v40 int32 = 993
                                _ = v40
                            }
                            _ = v38
                        }
                        fmt.Println(v1, v2)
                        _ = v26
                    } else {
                        v1 += 461
                    }
                    _ = fi25
                }
            }
        default:
            fmt.Println(v12, v1)
        }
    case -1:
        for fi41 := 0; fi41 < 7; fi41++ {
            v6 = v6
            switch fi41 {
            case 513:
                _ = len(v4)
            }
            switch fi41 {
            case -2147483648:
                if (false && false) {
                    switch 450 {
                    case -2147483648:
                        v2 = v6
                    case 9223372036854775807:
                        for fi42 := 0; fi42 < 1; fi42++ {
                            var v43 float64 = (0.0 + 3.14159)
                            switch 42 {
                            case 42:
                                var v44 uint8 = 230
                                _ = v44
                            default:
                                delete(v4, uint32(v1))
                            }
                            var v45 uint64 = 18446744073709551615
                            _ = fi42
                            _ = v43
                            _ = v45
                        }
                    case 4294967295:
                        fmt.Println(v1)
                    default:
                        v4[(v12 + v12)] = v12
                    }
                } else {
                    v2 = v2
                    var v46 int64 = -571
                    _ = v46
                }
            case -1:
                v47 := GenS3{F0: (fi41 + (^(fi41)))}
                _ = v47
            }
            _ = fi41
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi48 := 0; fi48 < 18; fi48++ {
        for fi49 := 0; fi49 < 1; fi49++ {
            fmt.Println(v6, fi48, fi48)
            _ = fi49
        }
        _ = fi48
    }
    v50 := GenS2{F0: v12}
    delete(v4, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
