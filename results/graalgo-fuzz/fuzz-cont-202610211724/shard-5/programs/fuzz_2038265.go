package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint16) int32 {
    var v5 uint32 = 2147483648
    switch -312 {
    case 9223372036854775807:
        v6 := GenS1{F0: 128}
        _ = v6
    }
    v5 = 300
    _ = p0
    _ = v5
    return -1
}

func main() {
    var v1 int = 84
    var v2 uint64 = 2
    v3 := []int{}
    var v4 float64 = 0.1
    switch "hello" {
    case "foo bar":
        v7 := GenS2{F0: (!(1 < 100))}
        _ = v7
    case "hello":
        for fi8 := 0; fi8 < 13; fi8++ {
            fmt.Println(v2, v1)
            if false {
                for fi9 := 0; fi9 < 20; fi9++ {
                    var v10 uint8 = 42
                    if false {
                        v1 += fi9
                        var v11 int64 = 1
                        fmt.Println(v4, fi9, fi9)
                        _ = v11
                    } else {
                        v4 = (-2.25 + (-2.25 * 3.14159))
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch (-(v1)) {
                        case 128:
                            var v12 uint16 = 678
                            v13 := GenS1{F0: (fi8 >> ((65535) & 63))}
                            var v14 float64 = v4
                            _ = v12
                            _ = v13
                            _ = v14
                        default:
                            var v15 uint64 = v2
                            _ = v15
                        }
                    }
                    _ = fi9
                    _ = v10
                }
                var v16 int8 = -1
                _ = v16
            }
            _ = fi8
        }
    default:
        for fi17 := 0; fi17 < 8; fi17++ {
            fmt.Println(fi17, v2, v4)
            _ = fi17
        }
    }
    v18 := GenS3{F0: 1, F1: true, F2: (!(473 >= 18446744073709551615))}
    v3 = append(v3, v1)
    if len(v3) > 0 {
        v3[0] = (v1 >> ((771) & 63))
    }
    v19 := GenS3{F0: 18446744073709551615, F1: false, F2: false}
    if true {
        var v20 uint = 0
        v20 += (v20 << (((v20 << ((v20) & 63))) & 63))
        v2 = (v2 >> ((v20) & 63))
        _ = v20
    } else {
        v4 += 0.1
    }
    var v21 int = v1
    var v22 uint32 = 0
    if (42 == v21) {
        v23 := GenS3{F0: 209, F1: ((float32(float64(1.5)) + 3.14159) < (3.14159 - (0.0 / (-2.25 - float32(float64(3.14159)))))), F2: false}
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi24 := 0; fi24 < 1; fi24++ {
                var v25 uint8 = 255
                var v26 int32 = 0
                _ = fi24
                _ = v25
                _ = v26
            }
            for fi27 := 0; fi27 < 9; fi27++ {
                v28 := GenS4{F0: -19}
                var v29 int64 = -318
                _ = fi27
                _ = v28
                _ = v29
            }
        } else {
            v1 -= 128
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 uint8 = 1
                v4 += (v4 + -2.25)
                _ = fi30
                _ = v31
            }
            for fi32 := 0; fi32 < 17; fi32++ {
                v1 += (-(v21))
                _ = fi32
            }
        }
        var v33 float32 = (1.5 - 1.5)
        _ = v23
        _ = v33
    }
    v1 = 288
    if ((true || true) || true) {
        var v34 uint32 = v22
        _ = len(v3)
        switch "a" {
        case "\n\t":
            var v35 int16 = -191
            _ = v35
        case "a":
            var v36 uint64 = (v2 ^ v2)
            _ = v36
        }
        _ = v34
    } else {
        if (v4 == ((0.1 * v4) - v4)) {
            v37 := GenS1{F0: (v21 | -200)}
            var v38 uint64 = 18446744073709551615
            _ = v37
            _ = v38
        } else {
            for fi39 := 0; fi39 < 8; fi39++ {
                fmt.Println(v21, v1, v2)
                v22 = (v22 | 4294967295)
                var v40 int64 = -39
                _ = fi39
                _ = v40
            }
            var v41 uint64 = (v2 ^ v2)
            _ = v41
        }
        v4 -= 0.0
    }
    var v42 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v42
}
