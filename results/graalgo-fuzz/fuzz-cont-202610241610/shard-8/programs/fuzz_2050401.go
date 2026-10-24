package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 int16 }
type GenS6 struct { F0 int64; F1 int; F2 uint32 }
type GenS7 struct { F0 uint64; F1 int64; F2 int8 }

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]string)
    v3 := []int{}
    for fi4 := 0; fi4 < 6; fi4++ {
        if false {
            v2[2147483647] = "foo bar"
        }
        _ = fi4
    }
    for fi5 := 0; fi5 < 16; fi5++ {
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 string = "foo bar"
            v8 := GenS1{F0: 1, F1: 42, F2: 2147483647}
            switch "a" {
            case "hello":
                var v9 uint64 = 42
                _ = v9
            case "a":
                if ((982 == 638) || false) {
                    v1 -= v1
                } else {
                    var v10 uint32 = 138
                    _ = v10
                }
            case "":
                if len(v3) > 0 {
                    v3[0] = fi6
                }
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        fmt.Println(fi5, fi5)
        _ = fi5
    }
    var v11 bool = true
    v1 = v1
    if (-225 >= -1) {
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 int16 = -32768
                var v15 uint16 = 2
                var v16 float32 = 3.14159
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi12
        }
        var v17 uint = 229
        _ = v17
    } else {
        v18 := GenS1{F0: 1, F1: 65535, F2: -493}
        v11 = v11
        if v11 {
            v19 := GenS2{F0: 128, F1: 32767, F2: 179}
            var v20 int = -516
            var v21 string = "hello"
            _ = v19
            _ = v20
            _ = v21
        } else {
            switch 127 {
            case -850:
                v1 -= (v1 & v1)
            }
            _ = len(v2)
        }
        _ = v18
    }
    if v11 {
        v11 = (-891 <= 667)
        v11 = v11
        var v22 float32 = 3.14159
        _ = v22
    } else {
        if true {
            v11 = true
        } else {
            var v23 uint32 = 0
            if false {
                v24 := GenS3{F0: (("hello" + "foo bar") + "")}
                _ = v24
            } else {
                if v11 {
                    delete(v2, v23)
                    if (1 < 644) {
                        v11 = (-93 <= 42)
                        v23 += (v23 & (v23 - (v23 * 72)))
                        v25 := GenS3{F0: ("" + "a")}
                        _ = v25
                    } else {
                        for fi26 := 0; fi26 < 20; fi26++ {
                            v27 := GenS4{F0: 128}
                            switch "hello" {
                            case "foo bar":
                                for fi28 := 0; fi28 < 5; fi28++ {
                                    if v11 {
                                        var v29 int16 = 1000
                                        if v11 {
                                            v30 := GenS5{F0: 42, F1: (v29 * (v29 << ((42) & 63)))}
                                            switch "foo bar" {
                                            case "hello":
                                                v11 = v11
                                            case "foo bar":
                                                v11 = v11
                                            default:
                                                var v31 bool = v11
                                                _ = v31
                                            }
                                            _ = v30
                                        }
                                        var v32 uint = 560
                                        _ = v29
                                        _ = v32
                                    }
                                    _ = fi28
                                }
                            case "":
                                for fi33 := 0; fi33 < 8; fi33++ {
                                    v1 -= -128
                                    v34 := GenS6{F0: 0, F1: 4294967295, F2: v23}
                                    if v11 {
                                        v35 := GenS7{F0: 2, F1: 42, F2: 1}
                                        _ = v35
                                    }
                                    _ = fi33
                                    _ = v34
                                }
                            case "a":
                                v1 += v1
                            }
                            for fi36 := 0; fi36 < 5; fi36++ {
                                _ = len(v3)
                                v1 = -25
                                v37 := GenS6{F0: 434, F1: (fi36 / ((fi26) | 1)), F2: v23}
                                _ = fi36
                                _ = v37
                            }
                            _ = fi26
                            _ = v27
                        }
                    }
                    var v38 float64 = 1.5
                    _ = v38
                }
                if v11 {
                    var v39 uint64 = 2
                    var v40 float32 = (0.0 / 1e10)
                    _ = v39
                    _ = v40
                }
            }
            if (!(true || v11)) {
                var v41 int64 = 9223372036854775807
                v41 = v41
                _ = v41
            }
            _ = v23
        }
        for fi42 := 0; fi42 < 12; fi42++ {
            fmt.Println(fi42)
            var v43 int64 = 1
            v43 = v43
            _ = fi42
            _ = v43
        }
    }
    var v44 bool = (-343 != 2147483647)
    if (1 > 255) {
        v45 := GenS1{F0: 1000, F1: 2, F2: 42}
        _ = v45
    } else {
        v46 := GenS4{F0: 1}
        _ = v46
    }
    var v47 uint32 = 42
    var v48 float32 = 1.5
    var v49 uint = 13
    for fi50 := 0; fi50 < 13; fi50++ {
        v51 := GenS2{F0: 64, F1: -790, F2: 127}
        _ = fi50
        _ = v51
    }
    fmt.Println(v44)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v44
    _ = v47
    _ = v48
    _ = v49
}
