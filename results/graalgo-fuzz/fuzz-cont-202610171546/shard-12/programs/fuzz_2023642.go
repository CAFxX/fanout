package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int64; F2 uint8 }
type GenS4 struct { F0 bool; F1 int64; F2 int32 }
func genFn0(p0 uint32) int {
    v7 := GenS1{F0: 100}
    var v8 int8 = -99
    for fi9 := 0; fi9 < 19; fi9++ {
        fmt.Println(v8, p0)
        v10 := GenS2{F0: 2147483648}
        var v11 uint16 = 70
        _ = fi9
        _ = v10
        _ = v11
    }
    v8 = v8
    _ = p0
    _ = v7
    _ = v8
    return int(p0)
}

func main() {
    v1 := []uint32{2147483647, 1, 65535, 577}
    var v2 float64 = 1e10
    var v3 string = "a"
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]int)
    var v6 uint32 = 542
    if true {
        for fi12 := 0; fi12 < 5; fi12++ {
            v5[(((v3 + (v3 + v3)) + v3) + v3)] = fi12
            _ = fi12
        }
        v2 += (v2 - v2)
        if true {
            v6 -= (v6 + 4294967295)
        } else {
            v13 := GenS1{F0: 1}
            var v14 float32 = 1e10
            _ = v13
            _ = v14
        }
    } else {
        var v15 bool = false
        fmt.Println(v3, v3, v15)
        _ = v15
    }
    delete(v4, (v6 % (((v6 | (v6 | v6))) | 1)))
    v3 = v3
    v16 := GenS3{F0: 42, F1: 0, F2: uint8(v6)}
    switch v3 {
    case "\n\t":
        v17 := GenS4{F0: (-1 == -9223372036854775808), F1: 9223372036854775807, F2: 42}
        _ = v17
    case "foo bar":
        for fi18 := 0; fi18 < 4; fi18++ {
            if (255 > 65535) {
                v19 := v5["hello"]
                _ = v19
            } else {
                v20 := v5["foo bar"]
                var v21 uint64 = 2
                _ = v20
                _ = v21
            }
            if ((-1 > int8(v6)) || (981 == 288)) {
                switch fi18 {
                case 128:
                    v1 = append(v1, (v6 / ((v6) | 1)))
                case 32768:
                    v3 = ("\n\t" + (v3 + "foo bar"))
                }
            } else {
                for fi22 := 0; fi22 < 3; fi22++ {
                    fmt.Println(v3)
                    v3 = v3
                    _ = fi22
                }
                var v23 float64 = 1.5
                fmt.Println(fi18, fi18)
                _ = v23
            }
            _ = fi18
        }
    default:
        if false {
            var v24 int64 = 0
            for fi25 := 0; fi25 < 5; fi25++ {
                v6 = 4294967295
                v26 := GenS4{F0: (genFn0(0) > genFn0((v6 - 0))), F1: (v24 | (v24 % ((v24) | 1))), F2: 1}
                _ = fi25
                _ = v26
            }
            _ = v24
        }
    }
    var v27 uint32 = (v6 * 258)
    if ((1000 >= 42) && false) {
        if (15 <= 200) {
            v28 := GenS3{F0: -9223372036854775808, F1: 0, F2: 128}
            var v29 uint8 = uint8(v6)
            var v30 int = int(v6)
            _ = v28
            _ = v29
            _ = v30
        }
        if true {
            var v31 int16 = 42
            var v32 string = "a"
            _ = v31
            _ = v32
        } else {
            v3 = "a"
        }
    }
    delete(v5, (v3 + ("\n\t" + (("" + ("hello" + "foo bar")) + v3))))
    for fi33 := 0; fi33 < 11; fi33++ {
        v1 = append(v1, v27)
        for fi34 := 0; fi34 < 11; fi34++ {
            v35 := GenS1{F0: 100}
            var v36 uint = 18446744073709551615
            if true {
                delete(v4, (v6 & 0))
            } else {
                var v37 string = "foo bar"
                _ = v37
            }
            v27 = v27
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi33
    }
    if (1 <= 879) {
        var v38 int8 = 1
        fmt.Println(v6)
        fmt.Println(v3, v2)
        _ = v38
    } else {
        for fi39 := 0; fi39 < 1; fi39++ {
            for fi40 := 0; fi40 < 8; fi40++ {
                if ("hello" >= v3) {
                    if false {
                        _ = len(v1)
                        switch ("" + (v3 + "foo bar")) {
                        case "":
                            v6 = v6
                        case "hello":
                            switch "" {
                            case "a":
                                var v41 uint32 = v6
                                _ = v41
                            }
                        }
                    }
                    fmt.Println(v27, fi39, fi40)
                    v3 = "foo bar"
                } else {
                    v42 := GenS3{F0: -9223372036854775808, F1: 0, F2: 11}
                    v2 -= v2
                    _ = v42
                }
                var v43 uint16 = 1
                v6 = v6
                _ = fi40
                _ = v43
            }
            v3 = v3
            v44 := GenS4{F0: (!false), F1: -1, F2: 557}
            _ = fi39
            _ = v44
        }
        for fi45 := 0; fi45 < 1; fi45++ {
            if len(v1) > 0 {
                v1[0] = (v27 & 4294967295)
            }
            _ = fi45
        }
    }
    var v46 int = 262
    var v47 uint16 = 1
    v47 = 1
    if false {
        v48 := GenS2{F0: v6}
        _ = v48
    } else {
        fmt.Println(v27)
    }
    var v49 uint32 = uint32(v47)
    v49 -= v27
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v49
    _ = v5
    _ = v6
}
