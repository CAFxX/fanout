package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint64 }
func genFn0() uint64 {
    var v5 float32 = 3.14159
    var v6 int8 = 42
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{}
    var v3 int8 = -1
    v4 := make(map[string]uint32)
    v3 -= 100
    v3 -= v3
    var v7 uint32 = 2147483647
    var v8 float32 = (0.1 * -2.25)
    fmt.Println(v7, v7, v7)
    v9 := GenS1{F0: false, F1: -2147483648}
    var v10 int8 = (v3 - v3)
    if (1 != v3) {
        if true {
            v10 -= v10
            var v11 uint8 = 145
            _ = v11
        } else {
            fmt.Println(v8, v8)
            var v12 int32 = -479
            fmt.Println(v10)
            _ = v12
        }
        var v13 uint32 = 0
        _ = v13
    } else {
        v10 += (v3 & (v3 + v3))
        if (3.14159 != float64(v8)) {
            if false {
                if true {
                    if true {
                        var v14 uint = 769
                        _ = len(v4)
                        _ = v14
                    }
                }
                v2 = append(v2, (v7 + (v7 / (((v7 * v7)) | 1))))
            } else {
                var v15 uint64 = 0
                _ = v15
            }
            for fi16 := 0; fi16 < 18; fi16++ {
                v8 += v8
                if true {
                    v2 = append(v2, (v7 % ((v7) | 1)))
                    v17 := GenS1{F0: (!(true || ((42 != 127) && (-204 != -2147483648)))), F1: 0}
                    _ = v17
                } else {
                    var v18 uint32 = (v7 - (v7 % ((v7) | 1)))
                    if true {
                        fmt.Println(fi16, v7)
                        if len(v2) > 0 {
                            v2[0] = 2
                        }
                        fmt.Println(v18, v8, v8)
                    } else {
                        var v19 bool = ((v7 ^ v7) >= v7)
                        for fi20 := 0; fi20 < 5; fi20++ {
                            v19 = v19
                            v3 = v10
                            _ = fi20
                        }
                        v21 := GenS1{F0: false, F1: 2147483647}
                        _ = v19
                        _ = v21
                    }
                    var v22 uint8 = 2
                    _ = v18
                    _ = v22
                }
                _ = fi16
            }
        } else {
            switch 4294967295 {
            case -9223372036854775808:
                switch 9223372036854775807 {
                case 9223372036854775807:
                    v10 = (v10 % ((v3) | 1))
                case 799:
                    fmt.Println(v10, v8, v7)
                }
            case -518:
                if true {
                    v3 -= v10
                    v23 := GenS1{F0: true, F1: -254}
                    fmt.Println(v7, v7, v8)
                    _ = v23
                }
            case 32767:
                switch 32767 {
                case -465:
                    var v24 uint16 = 2
                    _ = v24
                }
            }
        }
    }
    if true {
        v25 := GenS1{F0: true, F1: -2147483648}
        _ = v25
    } else {
        v1[v7] = (("foo bar" + ("foo bar" + ("foo bar" + "\n\t"))) <= "a")
        if (!true) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v2 = append(v2, v7)
            if (0 != 255) {
                fmt.Println(v7, v3, v10)
                var v26 int64 = -1
                if (42 >= 132) {
                    v8 -= float32(-2.25)
                    for fi27 := 0; fi27 < 16; fi27++ {
                        fmt.Println(fi27)
                        _ = fi27
                    }
                    v28 := GenS1{F0: true, F1: 42}
                    _ = v28
                }
                _ = v26
            } else {
                if (-9223372036854775808 == 128) {
                    var v29 uint = 1
                    _ = v29
                }
                v30 := GenS2{F0: 0}
                _ = v30
            }
        }
        var v31 float64 = 1.5
        _ = v31
    }
    if len(v2) > 0 {
        v2[0] = 41
    }
    var v32 int32 = -589
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v8, v8)
    } else {
        if (!false) {
            var v33 uint64 = 42
            _ = v33
        }
    }
    var v34 bool = false
    switch 4294967295 {
    case 0:
        var v35 bool = ((4294967295 > 42) || (true || true))
        _ = v35
    default:
        delete(v1, (v7 % (((v7 & v7)) | 1)))
    }
    fmt.Println(v32, v34, v10)
    fmt.Println(v10, v7, v10)
    var v36 float32 = ((v8 - 0.1) / (3.14159 + 0.0))
    v10 += 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v36
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
