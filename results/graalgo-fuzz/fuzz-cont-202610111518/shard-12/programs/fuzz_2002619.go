package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int; F1 uint }

func main() {
    var v1 float32 = 0.1
    var v2 uint16 = 65535
    v3 := make(map[int]int)
    v4 := make(map[string]string)
    v5 := make(map[uint32]bool)
    var v6 int8 = 42
    fmt.Println(v1)
    v7 := GenS1{F0: 0, F1: 0, F2: 0}
    var v8 uint32 = uint32(v2)
    fmt.Println(v8, v2, v2)
    v1 = float32(0.1)
    v6 += v6
    fmt.Println(v1, v6)
    v9 := GenS2{F0: v6}
    if (-2147483648 > -293) {
        if false {
            switch 32768 {
            case 978:
                if false {
                    v2 = v2
                    v6 += v6
                    v10 := GenS3{F0: 42}
                    _ = v10
                } else {
                    v11 := v4["foo bar"]
                    var v12 uint64 = 336
                    var v13 int32 = 319
                    _ = v11
                    _ = v12
                    _ = v13
                }
            default:
                if (1 <= 2) {
                    switch 850 {
                    case 630:
                        v6 = (v6 + v6)
                    case 32768:
                        _ = len(v3)
                    case 32767:
                        switch 127 {
                        case -2147483648:
                            for fi14 := 0; fi14 < 20; fi14++ {
                                var v15 float32 = 0.0
                                _ = fi14
                                _ = v15
                            }
                        default:
                            for fi16 := 0; fi16 < 5; fi16++ {
                                var v17 int = (^(fi16))
                                for fi18 := 0; fi18 < 6; fi18++ {
                                    var v19 uint32 = v8
                                    _ = fi18
                                    _ = v19
                                }
                                _ = fi16
                                _ = v17
                            }
                        }
                    default:
                        if (-962 != -282) {
                            var v20 uint64 = 18446744073709551615
                            _ = v20
                        } else {
                            v1 = (v1 / (1e10 / v1))
                            v8 -= 2147483648
                            if true {
                                v21 := GenS4{F0: -683}
                                _ = v21
                            }
                        }
                    }
                }
            }
            if true {
                v1 += (v1 + float32(1.5))
                v1 = float32(0.1)
            } else {
                v22 := GenS5{F0: 4294967295, F1: 266}
                _ = v22
            }
        } else {
            _ = len(v4)
            var v23 string = ("\n\t" + "")
            _ = v23
        }
        v24 := []int{762}
        _ = v24
    } else {
        v2 -= v2
        switch -455 {
        case 894:
            v25 := GenS4{F0: -108}
            _ = v25
        }
    }
    v26 := GenS3{F0: -923}
    v6 = v6
    for fi27 := 0; fi27 < 16; fi27++ {
        if false {
            for fi28 := 0; fi28 < 1; fi28++ {
                v29 := []int{42, int(v2), 32768, 42}
                if len(v29) >= 2 {
                    v29 = v29[1:len(v29)]
                }
                var v30 float32 = 1.5
                _ = fi28
                _ = v29
                _ = v30
            }
            v1 += (0.0 * v1)
            if (false && ("foo bar" == ("\n\t" + "\n\t"))) {
                for fi31 := 0; fi31 < 3; fi31++ {
                    fmt.Println(fi31, v8, fi31)
                    fmt.Println(fi31, v6)
                    fmt.Println(v6, fi31, v8)
                    _ = fi31
                }
            }
        }
        _ = fi27
    }
    switch 128 {
    case 32767:
        var v32 int64 = 42
        _ = v32
    case 128:
        var v33 float64 = 1.5
        _ = v33
    }
    v34 := GenS5{F0: 128, F1: 42}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
