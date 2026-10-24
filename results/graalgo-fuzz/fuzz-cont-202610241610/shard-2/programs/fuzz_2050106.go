package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int64; F2 int64 }
type GenS6 struct { F0 uint; F1 int16 }
func genFn0(p0 int32, p1 uint32, p2 bool) uint16 {
    fmt.Println(p1, p2, p0)
    fmt.Println(p2, p0, p0)
    v4 := GenS1{F0: 255}
    if true {
        var v5 uint8 = 0
        p2 = (!true)
        v5 = (^(v5))
        _ = v5
    } else {
        v6 := []int{32767, 428}
        v7 := GenS2{F0: 42, F1: 255, F2: 1}
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 32767
}
func genFn1(p0 string, p1 int, p2 int8) string {
    p1 -= 737
    var v8 int8 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return "hello"
}

func main() {
    v1 := []string{"\n\t", "hello"}
    v2 := make(map[uint32]int)
    var v3 uint32 = 2147483647
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            if (!(((1e10 / 1.5) < -2.25) && false)) {
                v3 = 65535
                if (fi9 != (fi10 - (^(fi9)))) {
                    var v11 bool = (!(true || (-9223372036854775808 <= -1)))
                    var v12 int = (fi9 * fi9)
                    v13 := GenS3{F0: 0, F1: (3.14159 <= 1.5)}
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    var v14 int32 = -918
                    _ = v14
                }
            } else {
                v3 -= (v3 >> ((450) & 63))
                if true {
                    var v15 uint32 = v3
                    var v16 int = -693
                    _ = v15
                    _ = v16
                } else {
                    if false {
                        v3 = v3
                    } else {
                        v3 -= (v3 + v3)
                        var v17 int64 = 9223372036854775807
                        _ = v17
                    }
                    v3 += 4294967295
                    fmt.Println(v3, v3)
                    fmt.Println(v3, fi10, fi9)
                }
                if false {
                    var v18 uint8 = 2
                    _ = v18
                } else {
                    v19 := GenS4{F0: 591}
                    _ = v19
                }
            }
            _ = fi10
        }
        _ = fi9
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(fi20, fi20)
        v3 = 0
        v3 -= 1
        _ = fi20
    }
    if true {
        if (!(255 <= 0)) {
            if false {
                if false {
                    var v21 string = (("hello" + "\n\t") + "a")
                    if true {
                        v2[v3] = 32767
                    }
                    v21 = "hello"
                    _ = v21
                } else {
                    if false {
                        v3 += (v3 << ((2) & 63))
                    } else {
                        var v22 uint = 1
                        v2[(v3 % ((117) | 1))] = -746
                        _ = v22
                    }
                }
            } else {
                v23 := GenS4{F0: 65535}
                var v24 string = "hello"
                for fi25 := 0; fi25 < 16; fi25++ {
                    v26 := GenS2{F0: 704, F1: 1, F2: 255}
                    if true {
                        switch fi25 {
                        case -892:
                            var v27 int = (fi25 - fi25)
                            _ = v27
                        default:
                            fmt.Println(v24, fi25)
                        }
                        fmt.Println(v3)
                    } else {
                        fmt.Println(fi25)
                        for fi28 := 0; fi28 < 4; fi28++ {
                            var v29 int8 = 127
                            if true {
                                var v30 int8 = (v29 % (((v29 << ((79) & 63))) | 1))
                                v31 := GenS5{F0: (v29 / ((1) | 1)), F1: 9223372036854775807, F2: 827}
                                if (v24 > v24) {
                                    var v32 uint16 = 947
                                    _ = v32
                                }
                                _ = v30
                                _ = v31
                            }
                            _ = fi28
                            _ = v29
                        }
                    }
                    _ = fi25
                    _ = v26
                }
                _ = v23
                _ = v24
            }
            v3 += v3
            var v33 uint8 = 120
            _ = v33
        }
        var v34 int16 = 0
        _ = v34
    }
    switch -2147483648 {
    case -1:
        v3 -= 687
    }
    v2[(v3 >> ((255) & 63))] = -9223372036854775808
    v1 = append(v1, "foo bar")
    v35 := GenS6{F0: 0, F1: -32768}
    var v36 uint = 2
    fmt.Println(v36)
    v1 = append(v1, "")
    for fi37 := 0; fi37 < 5; fi37++ {
        if false {
            v38 := GenS1{F0: 255}
            _ = v38
        }
        _ = fi37
    }
    switch "hello" {
    case "foo bar":
        switch -833 {
        case 4294967295:
            var v39 uint64 = 4294967295
            _ = v39
        case -2147483648:
            v3 += 1
        case 1:
            for fi40 := 0; fi40 < 1; fi40++ {
                for fi41 := 0; fi41 < 16; fi41++ {
                    fmt.Println(fi41, fi41, fi41)
                    v36 += (v36 ^ v36)
                    _ = fi41
                }
                var v42 uint16 = 42
                _ = fi40
                _ = v42
            }
        default:
            if false {
                var v43 uint8 = 1
                _ = v43
            }
        }
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
}
