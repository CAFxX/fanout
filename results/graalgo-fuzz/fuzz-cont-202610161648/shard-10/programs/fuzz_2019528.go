package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint; F1 int8; F2 uint16 }
type GenS5 struct { F0 bool; F1 int8 }
func genFn0(p0 uint64, p1 uint32, p2 string) int64 {
    var v7 uint = 1
    var v8 int64 = -1
    fmt.Println(v8, p2, p0)
    if true {
        v7 = v7
    } else {
        var v9 float64 = 1e10
        _ = v9
    }
    switch p2 {
    case "foo bar":
        var v10 uint32 = 1
        _ = v10
    case "a":
        v7 -= (v7 / (((v7 % ((802) | 1))) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (^(v8))
}

func main() {
    var v1 uint = 916
    v2 := make(map[uint32]uint32)
    v3 := []int{-658, 32768, 4294967295, 404}
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    v6 := make(map[string]bool)
    v1 -= 42
    if false {
        for fi11 := 0; fi11 < 3; fi11++ {
            switch (^(fi11)) {
            case -2147483648:
                var v12 float32 = -2.25
                _ = v12
            default:
                for fi13 := 0; fi13 < 16; fi13++ {
                    switch "" {
                    case "foo bar":
                        if (1 == -1) {
                            v14 := GenS1{F0: 115}
                            if false {
                                v1 += 255
                                v15 := GenS2{F0: -1, F1: uint8(v1), F2: 42}
                                _ = v15
                            }
                            _ = v14
                        } else {
                            v16 := GenS2{F0: (fi11 - fi13), F1: 234, F2: 2147483648}
                            v4[89] = 7
                            _ = v16
                        }
                    case "":
                        v1 -= (v1 * 4294967295)
                    default:
                        var v17 uint8 = 186
                        _ = v17
                    }
                    var v18 bool = (!true)
                    _ = fi13
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 2; fi19++ {
                v1 += v1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v20 uint = (v1 | 807)
                _ = fi19
                _ = v20
            }
            v1 = v1
            _ = fi11
        }
        v1 = (v1 >> ((0) & 63))
        var v21 int32 = 0
        _ = v21
    } else {
        if (!true) {
            for fi22 := 0; fi22 < 20; fi22++ {
                for fi23 := 0; fi23 < 18; fi23++ {
                    if (!(100 >= -1)) {
                        var v24 uint32 = 65535
                        var v25 uint8 = 197
                        _ = v24
                        _ = v25
                    }
                    _ = fi23
                }
                _ = fi22
            }
            v3 = append(v3, 256)
            v1 = v1
        } else {
            for fi26 := 0; fi26 < 8; fi26++ {
                for fi27 := 0; fi27 < 2; fi27++ {
                    var v28 int64 = 468
                    switch "" {
                    case "":
                        fmt.Println(v28)
                    case "foo bar":
                        var v29 uint16 = 42
                        _ = v29
                    case "\n\t":
                        var v30 int32 = -43
                        _ = v30
                    default:
                        v28 = -1
                    }
                    v28 = genFn0(1, 2147483648, "foo bar")
                    _ = fi27
                    _ = v28
                }
                var v31 float64 = 3.14159
                delete(v6, "foo bar")
                _ = fi26
                _ = v31
            }
            for fi32 := 0; fi32 < 15; fi32++ {
                fmt.Println(v1, fi32)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v33 := GenS3{F0: 0, F1: 0, F2: -1}
                _ = fi32
                _ = v33
            }
            var v34 int32 = 0
            _ = v34
        }
        _ = len(v3)
        v35 := GenS4{F0: 0, F1: 42, F2: 2}
        _ = v35
    }
    var v36 int = -2147483648
    if false {
        v37 := GenS1{F0: -1}
        fmt.Println(v1, v1)
        switch ("\n\t" + "") {
        case "hello":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "a":
            v38 := GenS3{F0: 0, F1: 42, F2: 42}
            _ = v38
        case "":
            var v39 bool = ((1 >= 127) || (42 > 2))
            _ = v39
        }
        _ = v37
    }
    for fi40 := 0; fi40 < 15; fi40++ {
        _ = len(v3)
        v36 = v36
        v1 = v1
        _ = fi40
    }
    var v41 string = "hello"
    delete(v4, 65535)
    if true {
        v42 := v5[-1]
        var v43 uint = v1
        _ = v42
        _ = v43
    }
    var v44 int16 = -1
    v1 += (v1 + 1)
    fmt.Println(v1, v1, v41)
    if (!false) {
        for fi45 := 0; fi45 < 4; fi45++ {
            v46 := GenS5{F0: false, F1: 42}
            _ = fi45
            _ = v46
        }
    }
    v47 := GenS3{F0: genFn0(4294967295, 687, v41), F1: 2, F2: -908}
    fmt.Println(v44)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v41
    _ = v44
    _ = v47
    _ = v5
    _ = v6
}
