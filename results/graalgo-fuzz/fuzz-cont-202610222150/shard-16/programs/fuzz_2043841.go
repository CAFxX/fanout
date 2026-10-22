package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int64 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint8; F1 int; F2 int16 }
type GenS5 struct { F0 uint8 }
func genFn0() uint16 {
    var v5 uint64 = 4294967295
    v5 -= (v5 / ((v5) | 1))
    switch "" {
    case "foo bar":
        var v6 int64 = 42
        _ = v6
    case "":
        v7 := GenS1{F0: false, F1: 0, F2: -1}
        _ = v7
    }
    v5 -= (v5 | v5)
    if (!true) {
        v5 -= 4294967295
        if false {
            var v8 uint8 = 0
            var v9 uint = uint(v5)
            fmt.Println(v9, v8)
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := GenS1{F0: (32767 >= 255), F1: 2, F2: 1}
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 uint16 = 2
                v13 += v13
                _ = fi12
                _ = v13
            }
        }
        v14 := []int{-2147483648, 873, -2147483648}
        _ = v14
    } else {
        var v15 int64 = 490
        if true {
            v15 += (v15 / ((-968) | 1))
        } else {
            v16 := []int{902, -75, 1}
            _ = v16
        }
        if false {
            var v17 uint32 = 2147483648
            _ = v17
        }
        _ = v15
    }
    _ = v5
    return 65535
}

func main() {
    var v1 uint16 = 0
    var v2 uint = 825
    v3 := make(map[uint32]uint32)
    var v4 string = "hello"
    var v18 int = 9223372036854775807
    if (127 < -80) {
        for fi19 := 0; fi19 < 1; fi19++ {
            v1 += 255
            if false {
                if (363 <= 42) {
                    var v20 int64 = 42
                    for fi21 := 0; fi21 < 12; fi21++ {
                        var v22 uint32 = 426
                        _ = len(v3)
                        if false {
                            fmt.Println(v2, v22)
                            var v23 string = ((v4 + ("\n\t" + v4)) + v4)
                            fmt.Println(v1, v1, v20)
                            switch int(v1) {
                            case 4294967295:
                                v24 := []int{v18}
                                _ = v24
                            case 9223372036854775807:
                                v25 := GenS2{F0: v20, F1: 0}
                                _ = v25
                            default:
                                var v26 uint32 = 192
                                _ = v26
                            }
                            _ = v23
                        } else {
                            var v27 bool = true
                            _ = v27
                        }
                        _ = fi21
                        _ = v22
                    }
                    _ = v20
                }
                fmt.Println(v18, v1, fi19)
            } else {
                var v28 int64 = 0
                _ = len(v3)
                for fi29 := 0; fi29 < 8; fi29++ {
                    var v30 uint = (v2 - v2)
                    var v31 int = (v18 ^ -891)
                    v32 := GenS2{F0: (v28 + v28), F1: (v1 ^ (v1 - (v1 - (^(v1)))))}
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                _ = v28
            }
            var v33 uint64 = 1
            _ = fi19
            _ = v33
        }
    } else {
        delete(v3, 1)
        v4 = "\n\t"
    }
    v34 := GenS1{F0: true, F1: 1, F2: 357}
    if (false && true) {
        v4 = (v4 + (v4 + ((v4 + (v4 + v4)) + v4)))
        v35 := []int{-2147483648}
        _ = v35
    } else {
        v36 := GenS3{F0: 32767, F1: v2}
        v18 = (v18 % ((0) | 1))
        v37 := []int{v18, (v18 << ((550) & 63))}
        _ = v36
        _ = v37
    }
    switch v4 {
    case "a":
        if ((!("" == "")) || true) {
            var v38 uint32 = 4294967295
            _ = v38
        } else {
            var v39 string = ("a" + "")
            fmt.Println(v1, v4, v2)
            _ = v39
        }
    case "foo bar":
        var v40 uint32 = 447
        _ = v40
    case "hello":
        v18 = v18
    default:
        v41 := GenS4{F0: 42, F1: (v18 / ((4294967295) | 1)), F2: -541}
        _ = v41
    }
    fmt.Println(v1, v18, v1)
    var v42 uint64 = 1
    v4 = v4
    _ = len(v3)
    for fi43 := 0; fi43 < 4; fi43++ {
        var v44 uint32 = 4294967295
        if (3.14159 <= 1e10) {
            if (!false) {
                for fi45 := 0; fi45 < 1; fi45++ {
                    var v46 int16 = -32768
                    fmt.Println(v44, v4, v44)
                    v47 := GenS2{F0: -1, F1: (v1 ^ v1)}
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                v2 -= v2
            } else {
                var v48 int64 = -9223372036854775808
                v49 := GenS5{F0: 149}
                fmt.Println(fi43)
                _ = v48
                _ = v49
            }
            v42 -= (^(v42))
        } else {
            _ = len(v3)
        }
        _ = fi43
        _ = v44
    }
    v2 = 542
    v50 := []int{v18, (v18 >> ((124) & 63))}
    for fi51 := 0; fi51 < 3; fi51++ {
        v1 = genFn0()
        _ = fi51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v50
}
