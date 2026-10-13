package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int64; F1 int; F2 uint64 }
type GenS4 struct { F0 uint; F1 bool }
func genFn0(p0 uint64, p1 string, p2 int64) uint16 {
    var v6 int64 = (p2 >> ((255) & 63))
    v7 := GenS1{F0: 42}
    var v8 int = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 uint, p1 int16, p2 string) int32 {
    var v9 int64 = 42
    v9 += 42
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 uint8 = 16
        var v12 int64 = v9
        switch fi10 {
        case -1:
            if true {
                if (-1 > 100) {
                    for fi13 := 0; fi13 < 17; fi13++ {
                        p1 = (^(p1))
                        _ = fi13
                    }
                    v14 := GenS2{F0: 232, F1: true}
                    var v15 int16 = -1
                    _ = v14
                    _ = v15
                } else {
                    v16 := []int{}
                    _ = v16
                }
                switch (fi10 * -632) {
                case 960:
                    switch int(v11) {
                    case 4294967295:
                        for fi17 := 0; fi17 < 12; fi17++ {
                            var v18 string = "\n\t"
                            _ = fi17
                            _ = v18
                        }
                    default:
                        v9 = -1
                    }
                case 436:
                    var v19 uint = p0
                    _ = v19
                default:
                    p2 = "\n\t"
                }
                v20 := GenS2{F0: 0, F1: false}
                _ = v20
            }
        case 32767:
            switch -9223372036854775808 {
            case -903:
                var v21 float64 = (0.1 + (0.0 + 1e10))
                _ = v21
            default:
                v22 := []int{(fi10 - fi10), (^(fi10)), (fi10 | 32767), (fi10 >> (((p0 ^ (p0 * uint(v11)))) & 63))}
                _ = v22
            }
        case 32768:
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 float64 = ((3.14159 - 3.14159) + (0.0 / 1.0))
                v24 -= (((v24 / ((v24 / 0.1) * (float64(((1e10 - 0.0) - ((0.0 * (3.14159 / (1e10 / 1.5))) / 3.14159))) / (float64(-2.25) - (1.5 - v24))))) * v24) - (float64((1e10 / -2.25)) * (((v24 - float64(3.14159)) / 1.0) + v24)))
                _ = fi23
                _ = v24
            }
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    var v25 float32 = float32(-2.25)
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v9
    return -724
}

func main() {
    v1 := []uint32{493, 0, 2}
    v2 := []string{"hello"}
    var v3 float64 = 1.5
    var v4 uint64 = 488
    v5 := []uint32{uint32(v4), 522, uint32(v4)}
    fmt.Println(v3, v4)
    var v26 int = -1
    for fi27 := 0; fi27 < 5; fi27++ {
        v28 := GenS2{F0: 42, F1: (!false)}
        var v29 int8 = int8(v4)
        switch 4294967295 {
        case 4294967295:
            fmt.Println(v26, v26)
        case 127:
            for fi30 := 0; fi30 < 19; fi30++ {
                v31 := GenS2{F0: 817, F1: true}
                var v32 float32 = (float32(0.1) - 0.1)
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    var v33 int = v26
    _ = len(v2)
    fmt.Println(v33, v26)
    fmt.Println(v33, v33, v3)
    v1 = append(v1, 2147483647)
    var v34 int16 = -32768
    v26 += v33
    switch v26 {
    case 1:
        for fi35 := 0; fi35 < 12; fi35++ {
            v33 = v26
            for fi36 := 0; fi36 < 20; fi36++ {
                fmt.Println(v4)
                _ = fi36
            }
            _ = fi35
        }
    case 627:
        for fi37 := 0; fi37 < 5; fi37++ {
            for fi38 := 0; fi38 < 2; fi38++ {
                if false {
                    v34 = v34
                }
                var v39 int16 = v34
                _ = fi38
                _ = v39
            }
            for fi40 := 0; fi40 < 19; fi40++ {
                switch (fi37 | 1) {
                case 127:
                    var v41 string = "a"
                    _ = v41
                case 2147483647:
                    var v42 float64 = v3
                    _ = v42
                case 42:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    if false {
                        v43 := GenS1{F0: 32767}
                        fmt.Println(fi40, v26)
                        var v44 uint16 = genFn0((v4 + v4), ("" + "hello"), 42)
                        _ = v43
                        _ = v44
                    } else {
                        v45 := GenS1{F0: genFn0(620, ("" + "hello"), -241)}
                        _ = v45
                    }
                    if (uint32(v4) <= 65535) {
                        var v46 float64 = v3
                        _ = v46
                    } else {
                        v2 = append(v2, (("hello" + "hello") + ("hello" + (("a" + "") + "a"))))
                        if false {
                            v47 := GenS3{F0: 448, F1: 738, F2: (v4 / ((v4) | 1))}
                            var v48 uint16 = genFn0(v4, ("a" + ("a" + ("hello" + "foo bar"))), -397)
                            fmt.Println(v3, v3)
                            _ = v47
                            _ = v48
                        } else {
                            fmt.Println(v34)
                            v26 = v26
                            if true {
                                if (9223372036854775807 > int64(v4)) {
                                    if true {
                                        var v49 int8 = 42
                                        var v50 uint16 = 837
                                        _ = v49
                                        _ = v50
                                    }
                                    v51 := GenS4{F0: 721, F1: true}
                                    var v52 int64 = -720
                                    v53 := GenS3{F0: (^(v52)), F1: 4294967295, F2: (^(v4))}
                                    fmt.Println(v33, v52)
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                } else {
                                    var v54 int16 = v34
                                    _ = v54
                                }
                                var v55 int64 = 42
                                v34 -= 0
                                var v56 string = ("hello" + "\n\t")
                                v34 -= 0
                                v56 = v56
                                var v57 bool = false
                                v26 = 4294967295
                                _ = v55
                                _ = v56
                                _ = v57
                            }
                        }
                        var v58 uint64 = (v4 | 0)
                        _ = v58
                    }
                    fmt.Println(v33)
                    var v59 uint32 = 65535
                    switch "foo bar" {
                    case "":
                        switch (("\n\t" + "a") + (("" + "hello") + "a")) {
                        case "a":
                            if (false && (v4 <= (v4 ^ 1))) {
                                v4 = (v4 - v4)
                            } else {
                                v4 -= (v4 + v4)
                                var v60 float32 = 3.14159
                                _ = v60
                            }
                        }
                    case "\n\t":
                        switch "" {
                        case "":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    case "a":
                        _ = len(v2)
                    }
                    _ = v59
                }
                _ = fi40
            }
            _ = fi37
        }
    case 315:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v61 uint64 = 224
    var v62 bool = true
    v63 := GenS4{F0: 447, F1: false}
    v64 := GenS4{F0: 42, F1: false}
    v34 += 0
    v1 = append(v1, 589)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v64
}
