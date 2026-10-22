package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint }
func genFn0(p0 bool, p1 int8) uint8 {
    v6 := GenS1{F0: int64(p1), F1: ((^(p1)) > (p1 & -109))}
    if (p1 == p1) {
        v7 := GenS1{F0: 1, F1: (!(0 >= 42))}
        p0 = (-32768 <= 730)
        v8 := []int{}
        _ = v7
        _ = v8
    } else {
        p0 = (false || (("a" + "hello") < (("\n\t" + "a") + "\n\t")))
        v9 := []int{-1, 2147483647, 127}
        p1 += p1
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    return 127
}
func genFn1() int {
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 float64 = 3.14159
            v13 := []int{fi11}
            _ = fi11
            _ = v12
            _ = v13
        }
        if (!true) {
            var v14 float32 = (float32(3.14159) / -2.25)
            _ = v14
        } else {
            fmt.Println(fi10)
            fmt.Println(fi10)
        }
        var v15 float64 = (-2.25 / (1e10 / (0.1 - (-2.25 + -2.25))))
        _ = fi10
        _ = v15
    }
    var v16 uint32 = 232
    v17 := []int{}
    var v18 int64 = 42
    _ = v16
    _ = v17
    _ = v18
    return 42
}

func main() {
    v1 := []int{}
    v2 := []int{-1, 1, -2147483648, -1}
    v3 := []int{}
    var v4 float64 = (1e10 + 1.5)
    v5 := make(map[uint32]string)
    v19 := GenS1{F0: 172, F1: false}
    var v20 float64 = (3.14159 / v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (!false) {
        if false {
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 int32 = 788
                _ = fi21
                _ = v22
            }
        }
    }
    if (true || (true || (false && (true || (0 >= 18446744073709551615))))) {
        v23 := GenS1{F0: 9223372036854775807, F1: (3.14159 < -2.25)}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v23
    } else {
        v20 -= (1.5 * float64(1e10))
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 uint32 = 4294967295
            if len(v2) > 0 {
                v2[0] = fi24
            }
            v4 = v20
            _ = fi25
            _ = v26
        }
        var v27 uint64 = 0
        switch "foo bar" {
        case "foo bar":
            v4 = (v20 - float64(1.5))
        case "\n\t":
            for fi28 := 0; fi28 < 14; fi28++ {
                switch "foo bar" {
                case "\n\t":
                    switch 127 {
                    case -291:
                        if (true && (-9223372036854775808 >= -2147483648)) {
                            var v29 float64 = v20
                            v27 = (v27 ^ (v27 + 797))
                            v30 := GenS2{F0: -9223372036854775808}
                            v31 := GenS3{F0: 32767, F1: false, F2: 65535}
                            v5[42] = "a"
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    case -2147483648:
                        var v32 uint8 = 1
                        _ = v32
                    case -1:
                        v33 := GenS2{F0: 1}
                        v2 = append(v2, genFn1())
                        _ = v33
                    default:
                        v34 := GenS1{F0: -9223372036854775808, F1: false}
                        _ = v34
                    }
                case "foo bar":
                    _ = len(v5)
                case "":
                    v27 = (v27 << ((18446744073709551615) & 63))
                default:
                    _ = len(v5)
                }
                v20 = float64(float32(v20))
                if false {
                    var v35 uint8 = 127
                    _ = v35
                }
                _ = fi28
            }
        case "a":
            var v36 float64 = (float64(float32(v4)) / (float64(3.14159) / 3.14159))
            _ = v36
        default:
            v20 += v20
        }
        _ = fi24
        _ = v27
    }
    var v37 int8 = 127
    fmt.Println(v4, v20)
    for fi38 := 0; fi38 < 1; fi38++ {
        v5[396] = "a"
        var v39 int = (fi38 | fi38)
        fmt.Println(v20, fi38)
        _ = fi38
        _ = v39
    }
    v4 += v20
    var v40 uint32 = 1
    v37 -= (v37 * (v37 % (((v37 & 42)) | 1)))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v20 = v20
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v5
}
