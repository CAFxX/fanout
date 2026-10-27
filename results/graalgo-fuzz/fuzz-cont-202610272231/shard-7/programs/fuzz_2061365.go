package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 string; F2 bool }
type GenS4 struct { F0 uint; F1 string; F2 uint16 }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 uint16 }
func genFn0(p0 uint8, p1 int16) int16 {
    for fi5 := 0; fi5 < 9; fi5++ {
        p1 -= (-(p1))
        v6 := GenS1{F0: 255, F1: -128}
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        if (18446744073709551615 >= 18446744073709551615) {
            v8 := GenS1{F0: 972, F1: -58}
            v9 := GenS1{F0: 42, F1: -76}
            var v10 int16 = (p1 ^ p1)
            _ = v10
            _ = v8
            _ = v9
        }
        var v11 uint32 = uint32(p0)
        p1 += (p1 % ((0) | 1))
        _ = fi7
        _ = v11
    }
    p1 = p1
    fmt.Println(p1)
    if false {
        fmt.Println(p1)
        if false {
            switch -1 {
            case 128:
                if false {
                    var v12 float64 = float64(1.5)
                    var v13 int16 = p1
                    v14 := GenS2{F0: 127}
                    _ = v12
                    _ = v13
                    _ = v14
                }
            case 4294967295:
                for fi15 := 0; fi15 < 7; fi15++ {
                    v16 := GenS2{F0: -128}
                    _ = fi15
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 2; fi17++ {
                v18 := []int{128, (fi17 + 128), (-(fi17))}
                switch 12 {
                case -733:
                    _ = len(v18)
                default:
                    var v19 int16 = -32768
                    _ = v19
                }
                fmt.Println(p1, fi17, fi17)
                _ = fi17
                _ = v18
            }
        }
        v20 := GenS3{F0: "a", F1: ("\n\t" + "hello"), F2: ((!true) || false)}
        _ = v20
    }
    _ = p0
    _ = p1
    return -1
}
func genFn1(p0 uint16, p1 int8, p2 int) int64 {
    var v21 string = (("hello" + ("" + "hello")) + ("hello" + "hello"))
    var v22 int8 = p1
    v23 := []int{p2, (p2 - (p2 + p2))}
    fmt.Println(v22, v21)
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    _ = v23
    return int64(p1)
}

func main() {
    v1 := []int{-1, -1, 128}
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 uint16 = 0
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 int32 = 420
        _ = fi24
        _ = v25
    }
    v4 = 0
    switch "" {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = 127
        }
    case "":
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS4{F0: 2, F1: ("foo bar" + "\n\t"), F2: (v4 << ((523) & 63))}
            if len(v1) > 0 {
                v1[0] = 1
            }
            _ = fi26
            _ = v27
        }
    default:
        v28 := GenS4{F0: 65535, F1: "foo bar", F2: (^(v4))}
        _ = v28
    }
    v4 = (v4 << ((42) & 63))
    fmt.Println(v4)
    if true {
        v4 = 847
        var v29 string = "foo bar"
        _ = v29
    }
    if (true || true) {
        v4 -= (v4 - v4)
        fmt.Println(v4)
    } else {
        v4 = v4
        v4 -= 1
        if true {
            if true {
                var v30 string = "\n\t"
                if (507 < 2) {
                    fmt.Println(v4, v4)
                }
                _ = v30
            }
            v31 := GenS2{F0: 42}
            _ = v31
        } else {
            switch 0 {
            case 74:
                v32 := GenS5{F0: -1, F1: 42}
                _ = v32
            case 4294967295:
                fmt.Println(v4)
            }
            v33 := GenS6{F0: true, F1: 1}
            var v34 string = "\n\t"
            _ = v33
            _ = v34
        }
    }
    v4 = 238
    v4 += 32767
    switch ("\n\t" + ("a" + ("" + "\n\t"))) {
    case "":
        fmt.Println(v4)
    case "foo bar":
        v35 := GenS4{F0: 2, F1: ("hello" + (("a" + "\n\t") + ("hello" + "a"))), F2: (v4 & (^(v4)))}
        _ = v35
    }
    var v36 uint = 0
    if ((v36 % ((v36) | 1)) != (v36 ^ v36)) {
        for fi37 := 0; fi37 < 20; fi37++ {
            v38 := GenS7{F0: 255}
            if false {
                var v39 uint = 18446744073709551615
                delete(v2, (-(fi37)))
                v39 -= v39
                _ = v39
            } else {
                v36 += v36
            }
            for fi40 := 0; fi40 < 20; fi40++ {
                if true {
                    fmt.Println(v36, fi37, fi40)
                    if false {
                        if false {
                            var v41 uint8 = 2
                            if true {
                                if false {
                                    v4 += v4
                                }
                            } else {
                                v36 += (v36 / (((v36 / (((v36 | v36)) | 1))) | 1))
                                fmt.Println(v4, v36, fi40)
                                v42 := v3[1]
                                _ = v42
                            }
                            _ = v41
                        } else {
                            if false {
                                v4 -= v4
                            }
                            v4 = 1
                        }
                        var v43 int = -2147483648
                        _ = v43
                    }
                    var v44 int8 = 42
                    if true {
                        if len(v1) > 0 {
                            v1[0] = int(v4)
                        }
                        v4 -= (v4 + v4)
                    } else {
                        v36 -= v36
                        var v45 int8 = -1
                        v46 := GenS2{F0: v44}
                        _ = v45
                        _ = v46
                    }
                    _ = v44
                } else {
                    fmt.Println(v4)
                    fmt.Println(v4)
                    var v47 uint16 = 0
                    _ = v47
                }
                _ = fi40
            }
            _ = fi37
            _ = v38
        }
        var v48 float64 = 0.0
        _ = v48
    } else {
        for fi49 := 0; fi49 < 19; fi49++ {
            for fi50 := 0; fi50 < 19; fi50++ {
                _ = len(v1)
                var v51 bool = ("foo bar" < "a")
                var v52 bool = v51
                _ = fi50
                _ = v51
                _ = v52
            }
            v4 += 1
            if true {
                for fi53 := 0; fi53 < 15; fi53++ {
                    var v54 int8 = -1
                    v54 -= (v54 << ((255) & 63))
                    _ = fi53
                    _ = v54
                }
            } else {
                switch ("" + "\n\t") {
                case "\n\t":
                    fmt.Println(v36, fi49, fi49)
                case "":
                    if true {
                        var v55 bool = (false || true)
                        var v56 uint16 = 0
                        var v57 uint8 = 0
                        _ = v55
                        _ = v56
                        _ = v57
                    }
                case "foo bar":
                    v4 = (v4 - 32767)
                default:
                    for fi58 := 0; fi58 < 3; fi58++ {
                        if false {
                            v59 := GenS6{F0: true, F1: 42}
                            _ = v59
                        }
                        var v60 string = ("\n\t" + ("foo bar" + "hello"))
                        _ = fi58
                        _ = v60
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi49
        }
        v61 := v3[401]
        _ = v61
    }
    delete(v2, 42)
    var v62 uint16 = (v4 & (^(v4)))
    if false {
        v62 = v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v62
}
