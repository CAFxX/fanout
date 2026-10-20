package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int64 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int16; F1 string }

func main() {
    var v1 int = 854
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 int8 = -99
    var v5 int8 = v4
    var v6 uint64 = 42
    if true {
        var v7 string = "hello"
        v6 = (^(v6))
        switch v1 {
        case 24:
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(v5, v6, v5)
                switch v7 {
                case "a":
                    var v9 int16 = -1
                    _ = v9
                case "\n\t":
                    for fi10 := 0; fi10 < 9; fi10++ {
                        v11 := []int{v1, -9223372036854775808}
                        _ = fi10
                        _ = v11
                    }
                default:
                    v5 = (v5 - v4)
                }
                _ = fi8
            }
        default:
            v5 = 24
        }
        _ = v7
    } else {
        v12 := []int{v1, v1, -9223372036854775808}
        _ = v12
    }
    var v13 uint64 = (v6 & v6)
    delete(v2, "foo bar")
    var v14 float32 = -2.25
    if false {
        var v15 string = ""
        for fi16 := 0; fi16 < 3; fi16++ {
            if (-913 < -66) {
                switch (v1 & (fi16 + (v1 ^ v1))) {
                case 2147483647:
                    var v17 int32 = -2147483648
                    _ = v17
                case 128:
                    if true {
                        switch v1 {
                        case -9223372036854775808:
                            switch fi16 {
                            case -495:
                                for fi18 := 0; fi18 < 5; fi18++ {
                                    v19 := []int{}
                                    _ = fi18
                                    _ = v19
                                }
                            case -9223372036854775808:
                                v20 := GenS1{F0: v1, F1: ((v15 + v15) > v15), F2: 42}
                                _ = v20
                            }
                        case -2147483648:
                            for fi21 := 0; fi21 < 17; fi21++ {
                                var v22 uint64 = (v13 & (v6 / (((v13 + 0)) | 1)))
                                var v23 bool = (false && (1 < 98))
                                var v24 uint32 = 633
                                _ = fi21
                                _ = v22
                                _ = v23
                                _ = v24
                            }
                        default:
                            delete(v2, ("" + ((((v15 + "a") + v15) + "hello") + v15)))
                        }
                        var v25 int64 = -386
                        _ = v25
                    } else {
                        var v26 int16 = -446
                        if true {
                            for fi27 := 0; fi27 < 14; fi27++ {
                                v28 := GenS1{F0: (fi27 / (((^(v1))) | 1)), F1: (true && true), F2: 0}
                                v29 := []int{v1}
                                _ = fi27
                                _ = v28
                                _ = v29
                            }
                        } else {
                            fmt.Println(v6, v6, v14)
                            v13 = v13
                        }
                        _ = v26
                    }
                case -9223372036854775808:
                    fmt.Println(v14, v5, v5)
                default:
                    var v30 bool = true
                    _ = v30
                }
                switch v1 {
                case 32767:
                    v31 := GenS1{F0: int(v4), F1: true, F2: -1}
                    _ = v31
                }
                switch v1 {
                case 9223372036854775807:
                    v3[2147483648] = -677
                case -863:
                    for fi32 := 0; fi32 < 3; fi32++ {
                        if false {
                            var v33 float32 = v14
                            _ = v33
                        }
                        _ = fi32
                    }
                }
            } else {
                v15 = ((v15 + v15) + ("hello" + v15))
                v34 := GenS1{F0: (fi16 ^ (fi16 - (-(fi16)))), F1: (!false), F2: 562}
                _ = v34
            }
            delete(v2, v15)
            _ = fi16
        }
        _ = v15
    } else {
        var v35 bool = false
        var v36 int16 = -243
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        var v38 int8 = v5
        _ = fi37
        _ = v38
    }
    v39 := GenS2{F0: 256, F1: 4294967295}
    switch "hello" {
    case "\n\t":
        for fi40 := 0; fi40 < 9; fi40++ {
            if true {
                var v41 float32 = ((float32(0.1) - ((1e10 * (v14 + v14)) / 0.1)) + v14)
                _ = v41
            } else {
                fmt.Println(v13, fi40, v4)
                fmt.Println(v4, fi40)
                v6 = (v6 ^ (^(v13)))
            }
            v42 := []int{(v1 ^ (-(v1))), fi40}
            _ = fi40
            _ = v42
        }
    case "":
        for fi43 := 0; fi43 < 17; fi43++ {
            if false {
                for fi44 := 0; fi44 < 17; fi44++ {
                    var v45 int16 = -356
                    v46 := GenS3{F0: v45, F1: ""}
                    _ = fi44
                    _ = v45
                    _ = v46
                }
            }
            _ = fi43
        }
    default:
        var v47 string = "\n\t"
        _ = v47
    }
    fmt.Println(v5, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
