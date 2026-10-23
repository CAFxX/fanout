package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint; F1 bool }
func genFn0() uint64 {
    var v6 int = 32767
    fmt.Println(v6)
    _ = v6
    return 42
}
func genFn1() uint32 {
    v7 := []int{0, 128}
    var v8 int64 = 234
    v8 = 42
    _ = len(v7)
    v8 += (v8 % (((v8 & -1)) | 1))
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int32 = -2147483648
    var v2 int32 = v1
    var v3 int = -1
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]uint32)
    if (true && false) {
        v9 := GenS1{F0: 58, F1: 128}
        _ = v9
    } else {
        var v10 uint32 = genFn1()
        v2 -= v1
        if (!(-503 >= -1)) {
            var v11 float64 = 3.14159
            for fi12 := 0; fi12 < 11; fi12++ {
                v4[v10] = v10
                for fi13 := 0; fi13 < 5; fi13++ {
                    switch "" {
                    case "hello":
                        var v14 float32 = (3.14159 * 0.0)
                        _ = v14
                    default:
                        var v15 string = "hello"
                        _ = v15
                    }
                    v16 := []int{}
                    _ = fi13
                    _ = v16
                }
                _ = fi12
            }
            for fi17 := 0; fi17 < 17; fi17++ {
                v2 += v2
                if false {
                    var v18 uint64 = 266
                    switch "a" {
                    case "foo bar":
                        v18 = v18
                    case "":
                        v19 := GenS2{F0: (!false), F1: 255}
                        _ = v19
                    case "\n\t":
                        if true {
                            v20 := []int{fi17, v3, v3, 726}
                            _ = v20
                        } else {
                            v11 -= (3.14159 / 1e10)
                            switch (("\n\t" + "foo bar") + "foo bar") {
                            case "":
                                var v21 uint64 = v18
                                _ = v21
                            default:
                                var v22 int8 = -21
                                _ = v22
                            }
                        }
                    default:
                        v1 = v1
                    }
                    if false {
                        for fi23 := 0; fi23 < 8; fi23++ {
                            var v24 int32 = v2
                            fmt.Println(fi23, v10, v2)
                            delete(v4, (v10 & genFn1()))
                            _ = fi23
                            _ = v24
                        }
                        v25 := GenS2{F0: (false || (242 > 18446744073709551615)), F1: 75}
                        v26 := GenS2{F0: ((0.1 - ((0.1 + float64(1e10)) * v11)) <= -2.25), F1: 4294967295}
                        _ = v25
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 17; fi27++ {
                            v2 -= int32(v10)
                            v28 := GenS1{F0: genFn1(), F1: uint8(v18)}
                            _ = fi27
                            _ = v28
                        }
                        v10 -= genFn1()
                    }
                    _ = v18
                }
                fmt.Println(v2)
                _ = fi17
            }
            _ = v11
        } else {
            if (32767 <= 32767) {
                v10 -= genFn1()
                v29 := []int{}
                _ = v29
            } else {
                if (true && true) {
                    v10 -= genFn1()
                } else {
                    var v30 string = "a"
                    var v31 uint = 366
                    _ = v30
                    _ = v31
                }
                if (!true) {
                    switch v3 {
                    case 1:
                        v32 := GenS3{F0: 1, F1: 996}
                        _ = v32
                    case 32767:
                        v10 = genFn1()
                    default:
                        if (true && false) {
                            v1 = v2
                            if false {
                                var v33 uint = 553
                                _ = v33
                            }
                        } else {
                            for fi34 := 0; fi34 < 17; fi34++ {
                                v35 := GenS2{F0: (728 < 2), F1: 207}
                                _ = fi34
                                _ = v35
                            }
                            fmt.Println(v2)
                        }
                    }
                }
            }
            _ = len(v4)
        }
        _ = v10
    }
    var v36 int64 = 42
    v3 -= (v3 | v3)
    v36 = v36
    var v37 uint8 = 2
    var v38 int32 = 2147483647
    if false {
        var v39 uint64 = 42
        switch -96 {
        case 42:
            var v40 int16 = 1
            _ = v40
        case -9223372036854775808:
            v41 := []int{}
            _ = v41
        }
        v42 := v4[uint32(v37)]
        _ = v39
        _ = v42
    } else {
        var v43 int = v3
        var v44 bool = true
        v45 := []int{-2147483648, v43, (v3 >> ((65535) & 63))}
        _ = v43
        _ = v44
        _ = v45
    }
    var v46 float32 = float32(1.5)
    var v47 uint32 = 967
    for fi48 := 0; fi48 < 11; fi48++ {
        var v49 int8 = -88
        v38 += (v1 >> ((42) & 63))
        _ = fi48
        _ = v49
    }
    v47 += (v47 * genFn1())
    for fi50 := 0; fi50 < 10; fi50++ {
        for fi51 := 0; fi51 < 12; fi51++ {
            v36 = (v36 | v36)
            _ = fi51
        }
        if false {
            if false {
                var v52 float32 = float32(float64(0.0))
                var v53 int = v3
                v54 := []int{}
                _ = v52
                _ = v53
                _ = v54
            } else {
                if true {
                    v55 := GenS4{F0: 430, F1: true}
                    fmt.Println(v2)
                    _ = v55
                }
                var v56 float32 = v46
                _ = v56
            }
            v37 = v37
            v46 = v46
        } else {
            switch (("\n\t" + ((("a" + ("a" + "\n\t")) + ("hello" + "a")) + "hello")) + "a") {
            case "a":
                v4[v47] = 166
            case "foo bar":
                v47 = 65535
            case "hello":
                var v57 int = v3
                _ = v57
            default:
                v47 -= 2147483648
            }
        }
        _ = fi50
    }
    for fi58 := 0; fi58 < 4; fi58++ {
        v38 -= 0
        _ = fi58
    }
    for fi59 := 0; fi59 < 17; fi59++ {
        var v60 int16 = 0
        _ = fi59
        _ = v60
    }
    var v61 string = ("\n\t" + "a")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v61
}
