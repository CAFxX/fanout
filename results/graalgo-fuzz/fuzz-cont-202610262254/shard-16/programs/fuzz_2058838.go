package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint16) uint64 {
    for fi4 := 0; fi4 < 18; fi4++ {
        if (true || true) {
            switch ("hello" + "\n\t") {
            case "hello":
                for fi5 := 0; fi5 < 18; fi5++ {
                    p1 = 505
                    if true {
                        p1 += 42
                    }
                    if (-73 > 127) {
                        if true {
                            p1 = p1
                        } else {
                            v6 := GenS1{F0: 128}
                            _ = v6
                        }
                        p1 = p1
                        v7 := GenS2{F0: p0, F1: -72}
                        v8 := []int{(-(fi4)), (fi5 << ((65535) & 63)), fi4, (fi5 - fi4)}
                        p0 += 9223372036854775807
                        _ = v7
                        _ = v8
                    }
                    _ = fi5
                }
            }
        } else {
            if false {
                p0 += (p0 << ((499) & 63))
            } else {
                if (int(p0) < (-(fi4))) {
                    if false {
                        var v9 uint8 = 42
                        v10 := []int{fi4, fi4, (fi4 & fi4)}
                        _ = v10
                        _ = v9
                    } else {
                        var v11 int8 = 1
                        if true {
                            var v12 int = fi4
                            if (0 >= 42) {
                                v11 -= v11
                                if true {
                                    for fi13 := 0; fi13 < 20; fi13++ {
                                        var v14 float32 = 3.14159
                                        var v15 float32 = float32((0.0 * (0.1 / (float64((float32(3.14159) - 1.5)) / 1.5))))
                                        v16 := GenS2{F0: 0, F1: 127}
                                        _ = fi13
                                        _ = v14
                                        _ = v15
                                        _ = v16
                                    }
                                } else {
                                    v17 := []int{}
                                    p0 = 677
                                    v12 = (fi4 * (v12 & fi4))
                                    _ = v17
                                }
                                var v18 uint = 65535
                                _ = v18
                            } else {
                                var v19 int32 = -1
                                p1 += p1
                                v19 -= (v19 & int32(fi4))
                                _ = v19
                            }
                            v12 += fi4
                            _ = v12
                        }
                        fmt.Println(p0)
                        _ = v11
                    }
                    v20 := []int{}
                    _ = v20
                }
                v21 := GenS3{F0: "\n\t", F1: 1}
                _ = v21
            }
            for fi22 := 0; fi22 < 5; fi22++ {
                p1 -= (^(p1))
                fmt.Println(fi22, p0)
                var v23 uint = 165
                _ = fi22
                _ = v23
            }
        }
        _ = fi4
    }
    var v24 uint32 = 2
    var v25 int = 127
    v26 := GenS4{F0: v24}
    _ = p0
    _ = p1
    _ = v24
    _ = v25
    _ = v26
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 0
    v3 := []string{"foo bar", ""}
    fmt.Println(v2)
    v27 := GenS4{F0: 618}
    switch 42 {
    case 0:
        var v28 uint = 116
        _ = v28
    }
    var v29 int16 = 0
    v30 := GenS2{F0: 0, F1: 100}
    var v31 bool = false
    for fi32 := 0; fi32 < 14; fi32++ {
        v29 = (v29 + v29)
        _ = fi32
    }
    if ((18446744073709551615 > 2) && (658 == 0)) {
        _ = len(v1)
        var v33 float64 = (-2.25 - 0.0)
        var v34 uint64 = 4294967295
        _ = v33
        _ = v34
    }
    v29 = v29
    if ((576 != 2147483647) || ((false && true) || v31)) {
        for fi35 := 0; fi35 < 16; fi35++ {
            var v36 string = "a"
            _ = fi35
            _ = v36
        }
    } else {
        var v37 uint32 = 4294967295
        if false {
            switch v2 {
            case 2147483647:
                switch "a" {
                case "a":
                    v38 := GenS1{F0: 179}
                    _ = v38
                case "\n\t":
                    var v39 string = (("" + "\n\t") + "a")
                    _ = v39
                case "":
                    v2 = (v2 * (v2 * 32768))
                }
            default:
                v2 = (v2 >> ((154) & 63))
            }
        }
        for fi40 := 0; fi40 < 20; fi40++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 -= (fi40 + fi40)
            var v41 uint16 = 951
            _ = fi40
            _ = v41
        }
        _ = v37
    }
    var v42 int16 = (v29 + 0)
    fmt.Println(v29, v31)
    v42 -= v42
    if (42 == 46) {
        v29 = v29
        v43 := GenS1{F0: 2}
        _ = v43
    }
    fmt.Println(v29, v2, v31)
    var v44 float64 = 1e10
    switch 986 {
    case 32767:
        var v45 int8 = 42
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v42
    _ = v44
}
