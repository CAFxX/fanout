package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint16; F1 uint32; F2 int32 }
func genFn0(p0 int16, p1 bool, p2 uint16) uint32 {
    var v6 int = 810
    v7 := GenS1{F0: 255, F1: 1}
    if ((p1 || true) || true) {
        if ("\n\t" < "foo bar") {
            v6 -= 1
        } else {
            for fi8 := 0; fi8 < 15; fi8++ {
                v9 := []int{(fi8 << ((255) & 63)), v6, 838}
                for fi10 := 0; fi10 < 8; fi10++ {
                    v6 += v6
                    _ = fi10
                }
                if p1 {
                    for fi11 := 0; fi11 < 9; fi11++ {
                        var v12 int = fi11
                        _ = fi11
                        _ = v12
                    }
                    if p1 {
                        if p1 {
                            fmt.Println(p1, p0)
                        } else {
                            v6 = fi8
                            var v13 int = fi8
                            p2 = 155
                            _ = v13
                        }
                    }
                }
                _ = fi8
                _ = v9
            }
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 uint8 = 2
                v16 := []int{(fi14 - -2147483648)}
                _ = fi14
                _ = v15
                _ = v16
            }
            p0 -= 1
        }
        p0 += -32768
    }
    p0 = (p0 / ((p0) | 1))
    p2 += p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int64 = -1
    v2 := make(map[string]int)
    v3 := []string{"\n\t"}
    var v4 string = "\n\t"
    v5 := []int{33}
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS2{F0: ("" + v4), F1: -9223372036854775808}
            var v20 uint32 = 4294967295
            v20 += (^(v20))
            _ = fi18
            _ = v19
            _ = v20
        }
        if true {
            var v21 int64 = (v1 & (v1 | 0))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v21
        }
        var v22 uint = 622
        _ = fi17
        _ = v22
    }
    var v23 int32 = 42
    var v24 uint64 = 1
    var v25 int = 31
    var v26 uint = 42
    fmt.Println(v25, v26)
    var v27 string = v4
    if (false && true) {
        var v28 uint = v26
        for fi29 := 0; fi29 < 13; fi29++ {
            v30 := GenS2{F0: "hello", F1: 32767}
            _ = fi29
            _ = v30
        }
        for fi31 := 0; fi31 < 8; fi31++ {
            if true {
                v23 = 1
            } else {
                var v32 uint64 = v24
                fmt.Println(v25, fi31, v26)
                _ = v32
            }
            switch 127 {
            case 0:
                v24 += (^(v24))
            case -405:
                _ = len(v5)
            default:
                for fi33 := 0; fi33 < 5; fi33++ {
                    switch v4 {
                    case "hello":
                        v5 = append(v5, int(v26))
                    default:
                        switch v4 {
                        case "":
                            if len(v5) > 0 {
                                v5[0] = (v25 << ((v26) & 63))
                            }
                        case "a":
                            v34 := GenS2{F0: v27, F1: (fi31 % ((fi33) | 1))}
                            _ = v34
                        case "foo bar":
                            v4 = v27
                        }
                    }
                    v35 := GenS1{F0: (v28 % ((v26) | 1)), F1: -71}
                    _ = fi33
                    _ = v35
                }
            }
            if false {
                var v36 uint16 = 955
                v27 = (("hello" + "\n\t") + v27)
                v37 := GenS3{F0: 65535, F1: 677, F2: (v23 << (((^(v28))) & 63))}
                _ = v36
                _ = v37
            }
            _ = fi31
        }
        _ = v28
    }
    if true {
        v38 := GenS2{F0: "a", F1: (v25 & v25)}
        v27 = v27
        _ = v38
    } else {
        var v39 float32 = 0.1
        v5 = append(v5, (^(v25)))
        var v40 int8 = 0
        _ = v39
        _ = v40
    }
    v24 -= 1
    var v41 int8 = 42
    _ = len(v2)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
