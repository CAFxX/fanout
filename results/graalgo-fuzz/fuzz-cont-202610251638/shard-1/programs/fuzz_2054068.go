package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint32 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 string; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint16 }
func genFn0() uint32 {
    v4 := []int{}
    var v5 uint = 936
    switch ("a" + "a") {
    case "":
        var v6 uint32 = uint32(v5)
        _ = v6
    case "foo bar":
        v5 += 0
    }
    var v7 int = 4294967295
    var v8 uint = 2
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    return 208
}
func genFn1(p0 uint32, p1 bool) uint32 {
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 float32 = 1e10
        _ = fi9
        _ = v10
    }
    p1 = false
    fmt.Println(p1, p1)
    var v11 int = 32767
    var v12 uint16 = 2
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return p0
}

func main() {
    var v1 float32 = float32(1e10)
    v2 := []int{}
    var v3 int32 = 2147483647
    fmt.Println(v3, v1)
    switch -922 {
    case -9223372036854775808:
        for fi13 := 0; fi13 < 3; fi13++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v1 += float32(3.14159)
            v14 := GenS1{F0: 2147483648, F1: (fi13 >> ((736) & 63)), F2: 2147483648}
            _ = fi13
            _ = v14
        }
    }
    v1 += v1
    fmt.Println(v3, v1)
    if (4294967295 >= 1) {
        var v15 int64 = 9223372036854775807
        switch -760 {
        case -979:
            var v16 float64 = (0.1 + 3.14159)
            _ = v16
        case 32768:
            v17 := GenS1{F0: 0, F1: -364, F2: genFn0()}
            _ = v17
        case 128:
            v18 := GenS1{F0: genFn1(0, (127 >= 100)), F1: -1, F2: 2}
            _ = v18
        }
        fmt.Println(v15, v3, v15)
        _ = v15
    } else {
        switch "a" {
        case "a":
            for fi19 := 0; fi19 < 8; fi19++ {
                v1 -= (v1 * (float32(3.14159) - v1))
                _ = fi19
            }
        case "\n\t":
            v1 -= v1
        default:
            for fi20 := 0; fi20 < 2; fi20++ {
                fmt.Println(v1, v1, fi20)
                if len(v2) > 0 {
                    v2[0] = 9223372036854775807
                }
                v21 := GenS1{F0: 2, F1: (-(fi20)), F2: 1}
                _ = fi20
                _ = v21
            }
        }
        v1 -= 1e10
    }
    var v22 uint8 = 0
    v1 -= v1
    var v23 uint16 = 149
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    for fi24 := 0; fi24 < 2; fi24++ {
        if (!(true || (int8(fi24) != 0))) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = len(v2)
            v25 := GenS1{F0: 340, F1: fi24, F2: 65535}
            _ = v25
        }
        if true {
            var v26 float32 = (float32(3.14159) - v1)
            _ = v26
        } else {
            v27 := GenS2{F0: 0, F1: (!(1 > 42))}
            if ((1 >= 0) && (!((-128 >= 78) || true))) {
                v3 = (v3 / ((v3) | 1))
                var v28 int64 = 1
                _ = v28
            }
            if (9 > -122) {
                fmt.Println(v23, v23)
                v29 := GenS2{F0: 0, F1: true}
                _ = v29
            }
            _ = v27
        }
        switch "\n\t" {
        case "a":
            for fi30 := 0; fi30 < 1; fi30++ {
                v31 := GenS3{F0: v22, F1: (((127 >= 47) || ("a" <= "hello")) && (uint(v22) >= 711))}
                switch (^(fi30)) {
                case -66:
                    if true {
                        v32 := GenS4{F0: "a", F1: -1, F2: 4294967295}
                        v3 -= (^(v3))
                        fmt.Println(v3, v3)
                        _ = v32
                    }
                case 0:
                    if (true || (!(65535 > 21))) {
                        _ = len(v2)
                        v33 := GenS5{F0: (^(v23)), F1: 2, F2: (v23 / ((v23) | 1))}
                        if len(v2) > 0 {
                            v2[0] = -759
                        }
                        _ = v33
                    }
                case 4294967295:
                    v3 = -136
                default:
                    v22 -= v22
                }
                v22 = (v22 * (v22 >> ((682) & 63)))
                _ = fi30
                _ = v31
            }
        }
        _ = fi24
    }
    v23 += (v23 ^ v23)
    var v34 float64 = (0.0 + ((-2.25 - ((0.0 * 0.1) / float64(v1))) + 0.0))
    var v35 uint8 = (v22 - (v22 - 128))
    switch "\n\t" {
    case "\n\t":
        v34 += 1.5
    }
    var v36 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
