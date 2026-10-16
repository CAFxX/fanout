package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 string, p1 uint, p2 uint) bool {
    var v6 float64 = float64(3.14159)
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return true
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    v4 := []uint32{978, 1, 2}
    var v5 uint32 = 4294967295
    switch "hello" {
    case "a":
        for fi7 := 0; fi7 < 13; fi7++ {
            for fi8 := 0; fi8 < 10; fi8++ {
                v5 = 1
                _ = fi8
            }
            _ = fi7
        }
    case "\n\t":
        v5 = (v5 % (((v5 + v5)) | 1))
    }
    if (0 != 88) {
        var v9 float64 = 0.1
        _ = v9
    } else {
        if true {
            v1 += v1
            for fi10 := 0; fi10 < 9; fi10++ {
                v4 = append(v4, (^(v5)))
                for fi11 := 0; fi11 < 12; fi11++ {
                    fmt.Println(v5, fi11)
                    if false {
                        v1 -= v1
                        v5 = 2147483647
                        v2[1] = uint32(v1)
                    } else {
                        if (int8(v1) != -116) {
                            v5 += v5
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v1 -= 1
                    }
                    v1 -= 2
                    _ = fi11
                }
                _ = fi10
            }
            v4 = append(v4, v5)
        }
        v5 += 0
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := v2[v5]
            v13 -= v5
            _ = fi12
            _ = v13
        }
    }
    if (952 < 28) {
        fmt.Println(v1, v1)
        v14 := GenS1{F0: (("\n\t" + "") + "hello")}
        _ = v14
    }
    v5 -= 1
    switch 1 {
    case 328:
        v1 += v1
    default:
        var v15 uint64 = 0
        _ = v15
    }
    switch 32767 {
    case 127:
        if false {
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := GenS2{F0: 0}
                delete(v2, (^(v5)))
                _ = fi16
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 18; fi18++ {
                fmt.Println(v1, v5)
                _ = len(v4)
                _ = fi18
            }
            var v19 uint = (v1 ^ v1)
            _ = v19
        }
    }
    v1 = 18446744073709551615
    if (float32(-2.25) != 1.5) {
        for fi20 := 0; fi20 < 6; fi20++ {
            for fi21 := 0; fi21 < 15; fi21++ {
                fmt.Println(fi21)
                _ = fi21
            }
            _ = fi20
        }
    }
    fmt.Println(v5)
    var v22 uint32 = v5
    v1 += (v1 & 818)
    if (485 != 353) {
        v23 := GenS3{F0: 0}
        v4 = append(v4, 0)
        var v24 string = ("foo bar" + ("a" + "\n\t"))
        _ = v23
        _ = v24
    }
    v22 += v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v22 = v22
    switch 16 {
    case -9223372036854775808:
        v1 = 65535
    default:
        if len(v4) > 0 {
            v4[0] = 0
        }
    }
    v5 += (v22 | v5)
    if true {
        if false {
            v5 -= v5
        } else {
            var v25 uint = v1
            _ = v25
        }
        fmt.Println(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
