package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint8, p1 uint32) uint64 {
    fmt.Println(p0)
    if true {
        if true {
            fmt.Println(p1)
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 uint32 = (p1 & p1)
                switch 127 {
                case 32767:
                    v8 += (p1 + p1)
                }
                _ = fi7
                _ = v8
            }
            p0 = (p0 >> ((201) & 63))
        } else {
            var v9 uint16 = 354
            switch "" {
            case "":
                var v10 uint32 = 2
                _ = v10
            case "foo bar":
                switch "hello" {
                case "hello":
                    var v11 int = -1
                    _ = v11
                }
            case "\n\t":
                for fi12 := 0; fi12 < 16; fi12++ {
                    v13 := GenS1{F0: 786, F1: -730}
                    v14 := []int{(fi12 & 0), fi12}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            default:
                var v15 uint32 = p1
                _ = v15
            }
            _ = v9
        }
        var v16 int32 = 0
        v17 := GenS2{F0: 703, F1: 1}
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    return 4294967295
}
func genFn1(p0 int16) int8 {
    v18 := []int{32768}
    p0 = 32767
    _ = p0
    _ = v18
    return -1
}

func main() {
    var v1 int8 = 42
    var v2 string = "a"
    v3 := make(map[string]bool)
    var v4 uint16 = 42
    var v5 int32 = -1
    v6 := make(map[uint32]string)
    for fi19 := 0; fi19 < 4; fi19++ {
        v20 := GenS3{F0: 170, F1: 65535}
        switch ("a" + v2) {
        case "":
            fmt.Println(v1, v4)
        default:
            delete(v3, v2)
        }
        v21 := []int{(fi19 / (((fi19 + fi19)) | 1)), 2147483647, 9223372036854775807}
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        var v23 int16 = 0
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        v25 := GenS1{F0: 42, F1: v5}
        for fi26 := 0; fi26 < 3; fi26++ {
            var v27 int16 = 93
            v27 -= (v27 % ((v27) | 1))
            if false {
                var v28 int32 = v5
                if true {
                    v29 := []int{fi26, 128, 655, (fi26 >> ((338) & 63))}
                    var v30 uint8 = 217
                    _ = v29
                    _ = v30
                }
                _ = len(v6)
                _ = v28
            }
            _ = len(v3)
            _ = fi26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v2, v5, v1)
    if false {
        var v31 uint32 = 390
        var v32 string = v2
        v4 -= (^(v4))
        _ = v31
        _ = v32
    } else {
        fmt.Println(v5, v4)
    }
    v4 = v4
    if ("hello" < v2) {
        for fi33 := 0; fi33 < 18; fi33++ {
            if true {
                var v34 bool = false
                _ = v34
            } else {
                v35 := GenS4{F0: (v1 | -73), F1: 42}
                _ = v35
            }
            v36 := []int{128, fi33, fi33, (fi33 - 4294967295)}
            v1 = v1
            _ = fi33
            _ = v36
        }
        v37 := GenS1{F0: 4294967295, F1: v5}
        _ = v37
    } else {
        v38 := v3["hello"]
        _ = v38
    }
    if (true && false) {
        v39 := GenS3{F0: -1, F1: 234}
        for fi40 := 0; fi40 < 7; fi40++ {
            for fi41 := 0; fi41 < 1; fi41++ {
                v6[715] = "foo bar"
                if (!(2 >= 577)) {
                    var v42 int = fi41
                    v43 := GenS2{F0: 835, F1: (v4 | 32767)}
                    if (42 >= 286) {
                        var v44 int8 = (-(v1))
                        v44 = v1
                        v45 := GenS2{F0: 4294967295, F1: (v4 / ((v4) | 1))}
                        v5 -= v5
                        v6[967] = "foo bar"
                        v3[((v2 + v2) + "")] = false
                        v1 = 127
                        _ = v44
                        _ = v45
                    } else {
                        var v46 uint32 = 4294967295
                        _ = v46
                    }
                    _ = v42
                    _ = v43
                }
                _ = fi41
            }
            v47 := GenS4{F0: -128, F1: 1}
            v4 += v4
            _ = fi40
            _ = v47
        }
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
