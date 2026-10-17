package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint32; F1 int8; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 bool; F2 uint64 }
func genFn0() int {
    var v6 uint = 4294967295
    v6 += v6
    v6 = v6
    v7 := []int{42, -9223372036854775808, -151, -9223372036854775808}
    _ = v6
    _ = v7
    return 128
}
func genFn1(p0 int64) string {
    for fi8 := 0; fi8 < 15; fi8++ {
        p0 += (p0 >> ((18446744073709551615) & 63))
        var v9 uint64 = 1
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 int64 = p0
            if (127 != 156) {
                fmt.Println(fi8, v11)
                if true {
                    p0 += p0
                    if (0 >= 2) {
                        v9 -= v9
                        if (1 > 0) {
                            p0 -= v11
                            var v12 bool = ((1 <= 1) && (true && false))
                            v12 = false
                            _ = v12
                        } else {
                            if ((!(4294967295 <= 0)) || true) {
                                v13 := GenS1{F0: 65535, F1: 2, F2: -863}
                                _ = v13
                            } else {
                                switch "" {
                                case "a":
                                    fmt.Println(fi10)
                                case "hello":
                                    var v14 float64 = 1e10
                                    _ = v14
                                }
                                var v15 bool = (!true)
                                v11 -= (v11 >> ((907) & 63))
                                _ = v15
                            }
                            fmt.Println(v9, fi10, v9)
                            v16 := []int{(fi8 & int(v9)), 9223372036854775807, fi8}
                            _ = v16
                        }
                    } else {
                        p0 += (p0 + 42)
                    }
                    v11 += 0
                }
                v11 += v11
            }
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    var v17 bool = false
    var v18 float64 = -2.25
    _ = p0
    _ = v17
    _ = v18
    return "hello"
}

func main() {
    v1 := []string{"hello"}
    var v2 float64 = 3.14159
    var v3 uint16 = 32767
    v4 := []string{"\n\t", ("" + ("a" + "\n\t")), "\n\t"}
    v5 := make(map[int]int)
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS2{F0: 1, F1: 2147483647}
        _ = fi19
        _ = v20
    }
    if (!true) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v3 -= 1
        if true {
            v2 = v2
            switch "foo bar" {
            case "hello":
                var v21 int = 4294967295
                _ = v21
            case "":
                switch "" {
                case "hello":
                    v3 += 920
                case "\n\t":
                    for fi22 := 0; fi22 < 20; fi22++ {
                        v23 := GenS3{F0: 0, F1: 127, F2: -780}
                        _ = fi22
                        _ = v23
                    }
                }
            }
        }
    } else {
        v2 = (((v2 * v2) + 3.14159) / (v2 + float64(0.0)))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := GenS4{F0: 1}
        fmt.Println(fi24, v3, fi24)
        var v26 int64 = -1
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := GenS5{F0: 42, F1: (!((528 >= -593) || (664 == -2147483648))), F2: 358}
    fmt.Println(v3, v2)
    var v28 uint8 = 0
    switch -252 {
    case 9223372036854775807:
        var v29 int8 = -1
        _ = v29
    }
    var v30 uint64 = uint64(v3)
    for fi31 := 0; fi31 < 3; fi31++ {
        var v32 int8 = 42
        var v33 int = int(v3)
        _ = fi31
        _ = v32
        _ = v33
    }
    v30 = 2
    var v34 int8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
