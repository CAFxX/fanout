package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int64 }

func main() {
    var v1 uint32 = 159
    var v2 int32 = -2147483648
    var v3 uint32 = v1
    if (9223372036854775807 != 0) {
        v4 := GenS1{F0: 0}
        _ = v4
    } else {
        if (!(670 <= 262)) {
            v5 := GenS2{F0: 42, F1: "foo bar"}
            v3 = 2147483648
            fmt.Println(v1, v1, v1)
            _ = v5
        }
        var v6 uint32 = v1
        _ = v6
    }
    v7 := []int{}
    var v8 int64 = 0
    if true {
        switch "" {
        case "a":
            if (4294967295 != -9223372036854775808) {
                switch 9223372036854775807 {
                case 0:
                    var v9 uint64 = 0
                    _ = v9
                case 32768:
                    if (2 < 2) {
                        for fi10 := 0; fi10 < 2; fi10++ {
                            _ = len(v7)
                            if true {
                                v11 := GenS1{F0: 880}
                                var v12 int8 = 1
                                if false {
                                    var v13 bool = (!(!(true && ((v2 * v2) > (v2 * (v2 % ((419) | 1)))))))
                                    var v14 int8 = (v12 << ((680) & 63))
                                    fmt.Println(v13, v8)
                                    _ = v13
                                    _ = v14
                                }
                                _ = v11
                                _ = v12
                            }
                            switch "foo bar" {
                            case "hello":
                                switch 9223372036854775807 {
                                case 531:
                                    var v15 uint32 = (v3 + (^(v3)))
                                    _ = v15
                                case 32767:
                                    if true {
                                        for fi16 := 0; fi16 < 16; fi16++ {
                                            var v17 float32 = 1e10
                                            _ = fi16
                                            _ = v17
                                        }
                                        for fi18 := 0; fi18 < 6; fi18++ {
                                            if len(v7) > 0 {
                                                v7[0] = (-(fi18))
                                            }
                                            _ = fi18
                                        }
                                        var v19 int32 = v2
                                        _ = v19
                                    }
                                case 127:
                                    var v20 int64 = v8
                                    _ = v20
                                default:
                                    var v21 int16 = 1000
                                    _ = v21
                                }
                            case "":
                                if len(v7) > 0 {
                                    v7[0] = (fi10 / (((fi10 ^ (fi10 - fi10))) | 1))
                                }
                            case "a":
                                v1 = (v3 * (^(v1)))
                            }
                            _ = fi10
                        }
                        _ = len(v7)
                    } else {
                        if ((v8 % ((42) | 1)) >= (-(v8))) {
                            v22 := GenS1{F0: 1}
                            _ = len(v7)
                            var v23 float32 = 0.0
                            _ = v22
                            _ = v23
                        }
                        v24 := GenS3{F0: 4294967295, F1: 29}
                        for fi25 := 0; fi25 < 6; fi25++ {
                            v1 = (v3 ^ v3)
                            var v26 int16 = 187
                            _ = fi25
                            _ = v26
                        }
                        _ = v24
                    }
                default:
                    for fi27 := 0; fi27 < 9; fi27++ {
                        v7 = append(v7, (fi27 | fi27))
                        switch "foo bar" {
                        case "":
                            for fi28 := 0; fi28 < 3; fi28++ {
                                var v29 int16 = 1
                                if false {
                                    v1 += (v3 ^ (v1 & (v3 - v1)))
                                }
                                if (42 <= 2) {
                                    v8 = v8
                                }
                                _ = fi28
                                _ = v29
                            }
                        }
                        if (-128 >= 127) {
                            var v30 float32 = 3.14159
                            v30 -= v30
                            _ = v30
                        } else {
                            v31 := GenS3{F0: 0, F1: fi27}
                            var v32 string = "foo bar"
                            for fi33 := 0; fi33 < 17; fi33++ {
                                if len(v7) > 0 {
                                    v7[0] = fi33
                                }
                                _ = fi33
                            }
                            _ = v31
                            _ = v32
                        }
                        _ = fi27
                    }
                }
                v2 -= (v2 + (v2 ^ (v2 * v2)))
            }
        case "hello":
            var v34 uint16 = uint16(v3)
            _ = v34
        default:
            if len(v7) > 0 {
                v7[0] = -1
            }
        }
    } else {
        v35 := GenS2{F0: 2, F1: ""}
        _ = v35
    }
    var v36 int64 = v8
    var v37 int64 = -9223372036854775808
    v1 += v3
    var v38 bool = true
    v1 -= v1
    switch -557 {
    case 1:
        var v39 uint64 = 326
        _ = v39
    default:
        v38 = (436 != (-(v2)))
    }
    for fi40 := 0; fi40 < 18; fi40++ {
        for fi41 := 0; fi41 < 3; fi41++ {
            fmt.Println(fi41, v3, v36)
            _ = fi41
        }
        var v42 uint16 = 0
        var v43 uint8 = 2
        _ = fi40
        _ = v42
        _ = v43
    }
    var v44 int8 = int8(v2)
    fmt.Println(v2)
    fmt.Println(v38, v2, v8)
    for fi45 := 0; fi45 < 15; fi45++ {
        _ = len(v7)
        _ = fi45
    }
    v2 = (v2 | (v2 + v2))
    if v38 {
        if (!(!(!(v38 || false)))) {
            if false {
                fmt.Println(v1, v3, v8)
                v46 := GenS4{F0: v8}
                _ = v46
            }
            for fi47 := 0; fi47 < 5; fi47++ {
                for fi48 := 0; fi48 < 9; fi48++ {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    var v49 int = 32768
                    var v50 uint8 = 127
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                if (1000 != -1) {
                    var v51 string = ""
                    _ = v51
                }
                _ = fi47
            }
        }
    }
    v7 = append(v7, int(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v44
    _ = v7
    _ = v8
}
