package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int64; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint) int {
    if (128 <= 318) {
        switch "" {
        case "foo bar":
            v4 := []int{}
            _ = v4
        case "\n\t":
            for fi5 := 0; fi5 < 16; fi5++ {
                if true {
                    var v6 float32 = float32(0.1)
                    _ = v6
                }
                p0 = p0
                v7 := GenS1{F0: 1}
                _ = fi5
                _ = v7
            }
        }
        v8 := []int{1, -9223372036854775808, -179}
        _ = v8
    }
    v9 := []int{42}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = v9
    return -9223372036854775808
}
func genFn1(p0 uint64) uint32 {
    var v10 uint = 255
    var v11 uint16 = 891
    var v12 uint = v10
    v13 := []int{}
    fmt.Println(v11, v10)
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 2
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    v3 := []uint32{527, 65535}
    _ = len(v2)
    v14 := GenS1{F0: 1}
    v15 := GenS2{F0: 42}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v16 uint = 254
    v16 = (^(v16))
    if true {
        v17 := GenS2{F0: -9223372036854775808}
        if true {
            v18 := GenS2{F0: 715}
            if true {
                switch -288 {
                case -1:
                    v16 += v16
                default:
                    if false {
                        for fi19 := 0; fi19 < 16; fi19++ {
                            var v20 float32 = (((-2.25 / (0.1 - 1.5)) - 0.1) / 1.0)
                            var v21 float32 = float32(3.14159)
                            for fi22 := 0; fi22 < 1; fi22++ {
                                v21 -= v21
                                _ = fi22
                            }
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        if (0 != 1) {
                            v16 = (v16 - (v16 | 255))
                            for fi23 := 0; fi23 < 15; fi23++ {
                                switch "foo bar" {
                                case "foo bar":
                                    if false {
                                        v24 := GenS2{F0: 460}
                                        _ = v24
                                    }
                                case "hello":
                                    for fi25 := 0; fi25 < 18; fi25++ {
                                        var v26 string = ((("\n\t" + "") + ("" + "\n\t")) + "\n\t")
                                        if (38 < 2) {
                                            var v27 uint32 = 0
                                            _ = v27
                                        } else {
                                            var v28 uint16 = 0
                                            delete(v1, v26)
                                            _ = v28
                                        }
                                        _ = fi25
                                        _ = v26
                                    }
                                }
                                var v29 int = (^(fi23))
                                var v30 int64 = 1
                                _ = fi23
                                _ = v29
                                _ = v30
                            }
                        }
                        v31 := GenS3{F0: 476, F1: 311}
                        _ = v31
                    }
                }
                fmt.Println(v16)
                _ = len(v3)
            } else {
                v16 = v16
                for fi32 := 0; fi32 < 16; fi32++ {
                    var v33 bool = true
                    if len(v2) > 0 {
                        v2[0] = fi32
                    }
                    _ = fi32
                    _ = v33
                }
                for fi34 := 0; fi34 < 17; fi34++ {
                    v16 += 0
                    _ = fi34
                }
            }
            _ = v18
        }
        v35 := GenS3{F0: uint16(v16), F1: 4294967295}
        _ = v17
        _ = v35
    } else {
        v16 = 42
        delete(v1, "\n\t")
    }
    var v36 uint16 = 0
    var v37 uint64 = 4294967295
    v16 = v16
    for fi38 := 0; fi38 < 16; fi38++ {
        if (42 == 255) {
            v39 := GenS2{F0: -1}
            var v40 int32 = -211
            _ = v39
            _ = v40
        }
        v36 = (v36 % (((v36 >> ((v16) & 63))) | 1))
        _ = fi38
    }
    v37 = 1
    if true {
        v36 += 32767
        v36 = v36
    } else {
        v16 = v16
        for fi41 := 0; fi41 < 9; fi41++ {
            var v42 int16 = -838
            switch int(v37) {
            case -845:
                fmt.Println(v37, v42, v16)
            case -274:
                var v43 float64 = -2.25
                _ = v43
            default:
                for fi44 := 0; fi44 < 6; fi44++ {
                    v42 += v42
                    if (-446 == int16(fi44)) {
                        switch int(v37) {
                        case 2147483647:
                            v45 := GenS4{F0: 0, F1: 100, F2: v36}
                            _ = v45
                        default:
                            v42 = (v42 * v42)
                        }
                    } else {
                        var v46 int32 = -1
                        v42 -= (v42 & 1)
                        _ = v46
                    }
                    _ = fi44
                }
            }
            var v47 bool = (!true)
            _ = fi41
            _ = v42
            _ = v47
        }
        v48 := GenS5{F0: uint64(v36)}
        _ = v48
    }
    var v49 uint64 = v37
    switch (("hello" + "a") + "\n\t") {
    case "foo bar":
        for fi50 := 0; fi50 < 17; fi50++ {
            switch (fi50 | 32767) {
            case -573:
                for fi51 := 0; fi51 < 7; fi51++ {
                    var v52 int = (fi51 * (fi51 - (fi51 / (((fi50 & (fi50 - fi50))) | 1))))
                    _ = fi51
                    _ = v52
                }
            case 42:
                v1[("foo bar" + "a")] = fi50
            default:
                v16 = 441
            }
            fmt.Println(v37, v49, v16)
            _ = fi50
        }
    case "hello":
        for fi53 := 0; fi53 < 14; fi53++ {
            for fi54 := 0; fi54 < 17; fi54++ {
                _ = len(v1)
                _ = fi54
            }
            v49 = 42
            _ = fi53
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v49
}
