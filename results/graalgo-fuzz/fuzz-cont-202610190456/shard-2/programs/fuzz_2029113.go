package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint32; F2 int16 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 uint16; F1 string; F2 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    var v4 bool = false
    var v5 int16 = -32768
    for fi6 := 0; fi6 < 16; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            switch "hello" {
            case "":
                var v8 uint8 = 255
                _ = v8
            case "a":
                if false {
                    var v9 int32 = -350
                    fmt.Println(v5, v5)
                    var v10 int32 = v9
                    _ = v10
                    _ = v9
                } else {
                    fmt.Println(fi7, v4, fi7)
                    if (!(!(float64(float32(1e10)) > 3.14159))) {
                        v4 = v4
                        _ = len(v1)
                    } else {
                        var v11 uint16 = 260
                        _ = len(v1)
                        _ = v11
                    }
                }
            default:
                v12 := []int{}
                _ = v12
            }
            fmt.Println(v4, fi7)
            _ = fi7
        }
        _ = fi6
    }
    _ = len(v3)
    v4 = v4
    if v4 {
        var v13 float32 = 1.5
        _ = v13
    }
    v14 := GenS1{F0: 107}
    fmt.Println(v5, v4)
    switch ("a" + "") {
    case "hello":
        v15 := GenS1{F0: 42}
        _ = v15
    case "foo bar":
        var v16 int64 = -700
        _ = v16
    default:
        v17 := GenS2{F0: (9223372036854775807 <= -1), F1: 789, F2: 42}
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        if true {
            for fi19 := 0; fi19 < 15; fi19++ {
                if v4 {
                    if ((!true) || (!(true && v4))) {
                        v4 = v4
                        v20 := []int{9223372036854775807, 42, (fi18 & 9223372036854775807)}
                        _ = v20
                    }
                    v21 := GenS3{F0: "", F1: 42}
                    _ = v21
                }
                _ = fi19
            }
            fmt.Println(v4, fi18, v4)
        } else {
            switch "a" {
            case "":
                for fi22 := 0; fi22 < 1; fi22++ {
                    v4 = false
                    fmt.Println(v4, v4, fi22)
                    _ = len(v2)
                    _ = fi22
                }
            case "foo bar":
                switch ("\n\t" + ("hello" + "")) {
                case "hello":
                    switch (^(fi18)) {
                    case 128:
                        if (!(!v4)) {
                            for fi23 := 0; fi23 < 8; fi23++ {
                                var v24 int64 = -1
                                var v25 int8 = 1
                                v5 = v5
                                var v26 uint = 18446744073709551615
                                var v27 int8 = (v25 / ((v25) | 1))
                                var v28 int8 = (^(v25))
                                _ = fi23
                                _ = v24
                                _ = v25
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                            v29 := GenS1{F0: 2}
                            var v30 float32 = 1e10
                            _ = v29
                            _ = v30
                        } else {
                            if true {
                                fmt.Println(v4, v5)
                                v31 := v1[55]
                                if v4 {
                                    for fi32 := 0; fi32 < 19; fi32++ {
                                        var v33 float32 = 1.5
                                        switch fi18 {
                                        case 32768:
                                            var v34 int16 = (v5 >> ((42) & 63))
                                            _ = v34
                                        case -9223372036854775808:
                                            var v35 uint8 = 0
                                            _ = v35
                                        }
                                        _ = fi32
                                        _ = v33
                                    }
                                    v36 := v2[v31]
                                    _ = v36
                                }
                                _ = v31
                            } else {
                                var v37 int32 = -1
                                v5 = v5
                                _ = v37
                            }
                            v4 = v4
                        }
                    }
                case "":
                    var v38 int32 = 354
                    _ = v38
                default:
                    for fi39 := 0; fi39 < 9; fi39++ {
                        delete(v2, 48)
                        fmt.Println(v4)
                        _ = fi39
                    }
                }
            case "a":
                fmt.Println(v5, v4)
            }
            v40 := GenS1{F0: 2}
            v4 = ((false || true) || false)
            _ = v40
        }
        var v41 int8 = -120
        for fi42 := 0; fi42 < 1; fi42++ {
            switch fi42 {
            case 2147483647:
                v43 := GenS4{F0: 0, F1: "\n\t", F2: ("a" + "hello")}
                _ = v43
            }
            if (1.5 == ((1e10 + (3.14159 + (1e10 - (-2.25 / 1.0)))) * 1.5)) {
                v44 := GenS2{F0: v4, F1: 65535, F2: -32768}
                _ = v44
            }
            v45 := v1[2]
            _ = fi42
            _ = v45
        }
        _ = fi18
        _ = v41
    }
    v4 = (1 >= 243)
    for fi46 := 0; fi46 < 15; fi46++ {
        for fi47 := 0; fi47 < 11; fi47++ {
            var v48 uint = 42
            _ = fi47
            _ = v48
        }
        switch fi46 {
        case -553:
            v5 -= v5
        case 42:
            v4 = v4
        default:
            v3[2147483648] = "hello"
        }
        _ = fi46
    }
    v4 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
