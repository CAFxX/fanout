package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint16 }

func main() {
    v1 := []uint32{1, 2147483647, 42, 1}
    var v2 uint16 = 0
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    v5 := []uint32{50, 842}
    v6 := []int{127, 32768}
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 bool = false
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 42, F1: 127, F2: v2}
        var v11 int = fi7
        _ = fi7
        _ = v10
        _ = v11
    }
    v2 -= 0
    var v12 uint64 = 18446744073709551615
    if len(v5) > 0 {
        v5[0] = uint32(v2)
    }
    delete(v3, "foo bar")
    fmt.Println(v12, v12)
    for fi13 := 0; fi13 < 9; fi13++ {
        if true {
            var v14 uint = 255
            _ = v14
        } else {
            delete(v3, "a")
            if ((false && true) && (!(false && false))) {
                switch "a" {
                case "a":
                    var v15 int16 = 42
                    _ = v15
                case "foo bar":
                    v2 += v2
                case "":
                    if true {
                        for fi16 := 0; fi16 < 5; fi16++ {
                            v17 := GenS1{F0: 128, F1: 0, F2: v2}
                            v2 = (^(v2))
                            v4[42] = (("" + "\n\t") + "\n\t")
                            if false {
                                v2 -= (v2 + v2)
                                var v18 uint64 = 366
                                _ = v18
                            } else {
                                fmt.Println(fi16)
                                var v19 float64 = 1e10
                                v2 = 42
                                _ = v19
                            }
                            var v20 uint8 = 127
                            if (int16(v12) <= 32767) {
                                var v21 bool = false
                                _ = v21
                            } else {
                                _ = len(v3)
                                if ((32767 > -1) && (uint(v2) <= uint(v20))) {
                                    v2 += 0
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = 878
                                    }
                                    _ = len(v4)
                                    fmt.Println(v20)
                                }
                                v22 := GenS1{F0: 77, F1: 1, F2: v2}
                                _ = v22
                            }
                            _ = fi16
                            _ = v17
                            _ = v20
                        }
                        var v23 float32 = 1.5
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 18; fi24++ {
                            v4[1] = (("a" + ("\n\t" + "\n\t")) + ("a" + "hello"))
                            v12 = v12
                            var v25 int64 = 0
                            _ = fi24
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 9; fi26++ {
                            fmt.Println(fi26)
                            _ = fi26
                        }
                    }
                }
                v12 = v12
            }
            if ((1.5 * 1e10) < 1.5) {
                if (2 != 1) {
                    var v27 uint = 4294967295
                    if (-120 != 1) {
                        v28 := GenS1{F0: 202, F1: 142, F2: v2}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(fi13, fi13, v27)
                        _ = v28
                    }
                    var v29 uint16 = 65535
                    _ = v27
                    _ = v29
                } else {
                    fmt.Println(fi13)
                    var v30 uint32 = 65535
                    _ = v30
                }
            } else {
                v12 -= uint64(v2)
                var v31 string = ("a" + ("hello" + "foo bar"))
                _ = v31
            }
        }
        _ = fi13
    }
    v2 = 719
    var v32 float64 = 1.5
    for fi33 := 0; fi33 < 11; fi33++ {
        if len(v5) > 0 {
            v5[0] = 0
        }
        fmt.Println(fi33, v12)
        v2 = (v2 & v2)
        _ = fi33
    }
    switch "a" {
    case "foo bar":
        var v34 int = 242
        _ = v34
    case "hello":
        var v35 uint32 = 65535
        _ = v35
    case "":
        var v36 string = "a"
        _ = v36
    default:
        v37 := GenS1{F0: 255, F1: 2, F2: v2}
        _ = v37
    }
    v38 := GenS1{F0: uint8(v2), F1: 128, F2: 1}
    for fi39 := 0; fi39 < 10; fi39++ {
        var v40 float64 = ((v32 * 0.1) * v32)
        _ = fi39
        _ = v40
    }
    switch "" {
    case "hello":
        var v41 uint64 = 2
        _ = v41
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = uint32(v2)
        }
    case "\n\t":
        v2 += v2
    default:
        v42 := GenS1{F0: 2, F1: 33, F2: (v2 / ((uint16(v12)) | 1))}
        _ = v42
    }
    var v43 uint16 = 360
    switch "a" {
    case "":
        for fi44 := 0; fi44 < 2; fi44++ {
            if false {
                v43 += v43
            } else {
                v45 := GenS1{F0: 0, F1: 128, F2: (^(v2))}
                _ = v45
            }
            var v46 bool = false
            _ = fi44
            _ = v46
        }
    case "a":
        _ = len(v6)
    default:
        var v47 string = "\n\t"
        _ = v47
    }
    v6 = append(v6, -1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
