package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int8 }
type GenS2 struct { F0 int16; F1 int32; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 300
    var v3 uint64 = 534
    var v4 uint32 = 2147483647
    v5 := make(map[int]string)
    v2 = (v2 - v2)
    switch 9223372036854775807 {
    case 32767:
        v6 := GenS1{F0: 127, F1: (v2 & v2), F2: -55}
        _ = v6
    case 127:
        if false {
            var v7 int32 = -1
            _ = v7
        } else {
            v8 := GenS1{F0: 83, F1: (v2 * (v2 ^ (v2 + (^(v2))))), F2: -1}
            var v9 string = "hello"
            switch 32768 {
            case 32767:
                for fi10 := 0; fi10 < 4; fi10++ {
                    var v11 uint32 = 65535
                    _ = fi10
                    _ = v11
                }
            default:
                v12 := GenS1{F0: 100, F1: (^(v2)), F2: -1}
                _ = v12
            }
            _ = v8
            _ = v9
        }
    default:
        fmt.Println(v4, v2, v3)
    }
    v4 = (^(v4))
    var v13 int64 = -57
    v2 -= v2
    fmt.Println(v13)
    v13 -= -1
    if (!(false || true)) {
        if (576 > 255) {
            for fi14 := 0; fi14 < 2; fi14++ {
                v4 = 2
                _ = fi14
            }
        } else {
            v3 -= v3
            var v15 string = ("foo bar" + "foo bar")
            v4 = 2
            _ = v15
        }
        var v16 float32 = 0.1
        fmt.Println(v13)
        _ = v16
    } else {
        v17 := GenS1{F0: -122, F1: v2, F2: 43}
        _ = v17
    }
    var v18 uint8 = 42
    if false {
        for fi19 := 0; fi19 < 19; fi19++ {
            v13 -= v13
            fmt.Println(v2, v2)
            _ = fi19
        }
        fmt.Println(v18)
    }
    var v20 int = -524
    v1[(v4 << ((4294967295) & 63))] = (^(v20))
    for fi21 := 0; fi21 < 13; fi21++ {
        v22 := GenS2{F0: 32767, F1: -1, F2: (^(v18))}
        v20 = (^(fi21))
        if (220 == (v18 / ((v18) | 1))) {
            v23 := GenS2{F0: 1, F1: -559, F2: 127}
            v2 += v2
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    if true {
        switch v20 {
        case 42:
            for fi24 := 0; fi24 < 19; fi24++ {
                var v25 int = (fi24 / ((v20) | 1))
                var v26 uint16 = (v2 | v2)
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        for fi27 := 0; fi27 < 3; fi27++ {
            if (728 <= v4) {
                var v28 uint16 = (^(v2))
                for fi29 := 0; fi29 < 1; fi29++ {
                    if ((1e10 - 3.14159) == (3.14159 - (0.1 * 0.0))) {
                        switch 4294967295 {
                        case 0:
                            v28 += uint16(v3)
                        }
                        var v30 float32 = 0.1
                        v31 := []int{-1}
                        _ = v30
                        _ = v31
                    } else {
                        var v32 int32 = -2147483648
                        _ = v32
                    }
                    v13 += (v13 % ((-1) | 1))
                    var v33 uint = 18446744073709551615
                    _ = fi29
                    _ = v33
                }
                _ = v28
            } else {
                v13 += v13
            }
            if (22 <= v18) {
                for fi34 := 0; fi34 < 9; fi34++ {
                    var v35 uint32 = 2147483647
                    if false {
                        v36 := GenS3{F0: 158, F1: (v3 ^ (v3 | 225))}
                        _ = v36
                    }
                    _ = fi34
                    _ = v35
                }
                fmt.Println(v13, fi27, v18)
                for fi37 := 0; fi37 < 1; fi37++ {
                    fmt.Println(v18, v13, v2)
                    _ = fi37
                }
            } else {
                fmt.Println(v13, v13, v3)
            }
            v20 = (fi27 ^ (v20 & fi27))
            _ = fi27
        }
        v38 := []int{v20, (-(v20)), v20, v20}
        _ = v38
    } else {
        v5[v20] = "foo bar"
        var v39 int = (v20 & (v20 + (v20 * (v20 + (v20 * (^(v20)))))))
        _ = v39
    }
    var v40 uint16 = uint16(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
