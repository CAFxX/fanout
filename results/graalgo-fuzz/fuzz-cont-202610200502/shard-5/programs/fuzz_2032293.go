package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 string }
type GenS2 struct { F0 int32; F1 int; F2 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 uint32; F2 bool }
type GenS6 struct { F0 uint }
func genFn0(p0 bool, p1 int32) string {
    var v4 int64 = 42
    for fi5 := 0; fi5 < 5; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 uint32 = 963
            v4 += (v4 ^ 42)
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint8 = 167
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: 0, F2: ("\n\t" + "a")}
    _ = p0
    _ = p1
    _ = v10
    _ = v4
    return "hello"
}
func genFn1(p0 string) int16 {
    v11 := GenS1{F0: 1, F1: 165, F2: ""}
    var v12 float64 = -2.25
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            if false {
                v15 := GenS2{F0: 42, F1: fi13, F2: p0}
                fmt.Println(fi13)
                v12 -= 1e10
                _ = v15
            }
            _ = fi14
        }
        _ = fi13
    }
    var v16 uint32 = 725
    _ = p0
    _ = v11
    _ = v12
    _ = v16
    return -1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    for fi17 := 0; fi17 < 4; fi17++ {
        if (1 < 69) {
            var v18 int = (fi17 * (fi17 * 1))
            v19 := GenS2{F0: 42, F1: 9223372036854775807, F2: "\n\t"}
            _ = v18
            _ = v19
        }
        var v20 int16 = -32768
        _ = fi17
        _ = v20
    }
    var v21 int = 42
    for fi22 := 0; fi22 < 17; fi22++ {
        if false {
            var v23 float64 = (3.14159 - (1.5 - 0.1))
            switch "a" {
            case "hello":
                var v24 uint32 = 2147483647
                _ = v24
            }
            _ = v23
        }
        _ = fi22
    }
    var v25 uint16 = 680
    for fi26 := 0; fi26 < 2; fi26++ {
        if ((2147483648 >= 42) && (!true)) {
            switch v21 {
            case -9223372036854775808:
                if false {
                    var v27 uint64 = 1
                    _ = v27
                } else {
                    fmt.Println(v21, fi26, v25)
                }
            case 9223372036854775807:
                switch fi26 {
                case -589:
                    for fi28 := 0; fi28 < 12; fi28++ {
                        var v29 uint64 = 0
                        var v30 uint64 = 4294967295
                        if ("\n\t" == "a") {
                            v2[973] = (fi26 - 550)
                        } else {
                            if ("a" >= "hello") {
                                var v31 float32 = 0.0
                                if (112 < 1) {
                                    var v32 int64 = 9223372036854775807
                                    _ = v32
                                }
                                v30 -= (v30 | v30)
                                _ = v31
                            }
                            v33 := GenS3{F0: true}
                            switch ("" + "foo bar") {
                            case "a":
                                fmt.Println(v29, v25, v21)
                            }
                            v30 = v30
                            _ = v33
                        }
                        v34 := GenS4{F0: 65535}
                        v35 := GenS5{F0: 0, F1: 42, F2: (false || true)}
                        _ = fi28
                        _ = v29
                        _ = v30
                        _ = v34
                        _ = v35
                    }
                case 720:
                    switch "" {
                    case "":
                        var v36 uint16 = (v25 | 65535)
                        _ = v36
                    case "\n\t":
                        for fi37 := 0; fi37 < 9; fi37++ {
                            if (2 != 4294967295) {
                                var v38 float64 = (-2.25 / 1.0)
                                _ = v38
                            }
                            _ = fi37
                        }
                    }
                }
            case 382:
                if false {
                    var v39 int16 = 0
                    var v40 float64 = -2.25
                    v40 += v40
                    _ = v39
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 14; fi41++ {
                        fmt.Println(fi26, v25, fi26)
                        var v42 uint16 = v25
                        _ = fi41
                        _ = v42
                    }
                    var v43 uint32 = 540
                    _ = v43
                }
            }
        }
        _ = fi26
    }
    v25 -= v25
    v44 := []int{int(v25), (v21 << ((18446744073709551615) & 63))}
    if false {
        v45 := GenS5{F0: -1, F1: 404, F2: ((0 <= 0) || false)}
        v25 = 65535
        var v46 int8 = 42
        _ = v45
        _ = v46
    }
    v47 := GenS4{F0: 65535}
    var v48 uint = 0
    var v49 int16 = -1
    if true {
        v50 := GenS2{F0: 1, F1: 87, F2: ""}
        if true {
            for fi51 := 0; fi51 < 16; fi51++ {
                fmt.Println(v48, v49)
                _ = fi51
            }
        }
        var v52 int32 = -2147483648
        _ = v50
        _ = v52
    } else {
        var v53 int32 = -504
        v54 := v3[2147483648]
        _ = v53
        _ = v54
    }
    if false {
        if len(v44) > 0 {
            v44[0] = v21
        }
        v21 -= v21
        var v55 string = "\n\t"
        _ = v55
    } else {
        fmt.Println(v25)
        var v56 uint64 = 618
        _ = v56
    }
    v57 := GenS6{F0: v48}
    v48 = v48
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v44
    _ = v47
    _ = v48
    _ = v49
    _ = v57
}
