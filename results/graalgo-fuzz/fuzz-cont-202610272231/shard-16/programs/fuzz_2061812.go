package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int; F1 int8 }
func genFn0() uint8 {
    v6 := []int{-341, 2147483647, 4294967295}
    if true {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int16 = 847
            switch ("hello" + ("\n\t" + "a")) {
            case "a":
                v8 += v8
            default:
                switch (fi7 - 32767) {
                case 610:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case -2147483648:
                    if (!true) {
                        if false {
                            var v9 float64 = (1.5 + 0.1)
                            _ = v9
                        }
                        v8 = 792
                        switch 128 {
                        case 127:
                            v8 = (v8 + v8)
                        case 4294967295:
                            var v10 uint = 481
                            _ = v10
                        }
                    } else {
                        if true {
                            if len(v6) > 0 {
                                v6[0] = (-(fi7))
                            }
                            _ = len(v6)
                        }
                    }
                default:
                    if (!(!(!true))) {
                        switch -9223372036854775808 {
                        case 285:
                            var v11 string = (("\n\t" + "\n\t") + "foo bar")
                            _ = v11
                        default:
                            _ = len(v6)
                        }
                    }
                }
            }
            switch "" {
            case "a":
                v8 += (v8 ^ v8)
            case "\n\t":
                v8 = -1
            }
            _ = fi7
            _ = v8
        }
    } else {
        if (false && (!false)) {
            for fi12 := 0; fi12 < 6; fi12++ {
                var v13 uint = 42
                _ = fi12
                _ = v13
            }
            if len(v6) > 0 {
                v6[0] = 127
            }
        } else {
            var v14 int64 = 42
            var v15 int32 = -1
            _ = v14
            _ = v15
        }
        if (122 > 32768) {
            var v16 float64 = float64(1e10)
            _ = len(v6)
            v16 = (v16 / v16)
            _ = v16
        } else {
            var v17 uint = 255
            for fi18 := 0; fi18 < 13; fi18++ {
                switch "a" {
                case "a":
                    v17 -= v17
                }
                for fi19 := 0; fi19 < 15; fi19++ {
                    var v20 int64 = -9223372036854775808
                    fmt.Println(fi19)
                    v17 -= (v17 << ((37) & 63))
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
            _ = v17
        }
    }
    _ = v6
    return 41
}
func genFn1(p0 int16, p1 uint, p2 int8) int64 {
    var v21 bool = ((true || false) || true)
    for fi22 := 0; fi22 < 11; fi22++ {
        v23 := []int{2147483647, 32768, fi22}
        if v21 {
            var v24 int = fi22
            _ = v24
        }
        _ = fi22
        _ = v23
    }
    v25 := GenS1{F0: p0, F1: 176, F2: 739}
    fmt.Println(p2, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v25
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("\n\t" + "a"), ""}
    var v3 int = 112
    var v4 uint = 533
    var v5 bool = (2 > uint16(v4))
    v26 := GenS1{F0: 1000, F1: 1, F2: 843}
    fmt.Println(v5, v5)
    for fi27 := 0; fi27 < 14; fi27++ {
        for fi28 := 0; fi28 < 1; fi28++ {
            v29 := v1[618]
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v30 := GenS1{F0: 290, F1: genFn0(), F2: 32767}
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi27
    }
    var v31 uint16 = 941
    var v32 uint16 = (v31 | (^(v31)))
    if v5 {
        for fi33 := 0; fi33 < 17; fi33++ {
            switch ("" + "") {
            case "a":
                var v34 int16 = 42
                _ = v34
            case "hello":
                var v35 string = ("hello" + "")
                _ = v35
            case "foo bar":
                v36 := GenS1{F0: 0, F1: genFn0(), F2: v31}
                _ = v36
            }
            _ = len(v1)
            var v37 uint32 = 65535
            _ = fi33
            _ = v37
        }
        for fi38 := 0; fi38 < 20; fi38++ {
            _ = len(v2)
            _ = fi38
        }
    }
    v31 = v32
    v39 := GenS2{F0: (^(v3)), F1: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v5
}
