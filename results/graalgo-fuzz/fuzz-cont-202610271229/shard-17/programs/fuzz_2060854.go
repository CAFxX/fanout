package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int32; F2 uint32 }
func genFn0(p0 int64) int32 {
    switch "foo bar" {
    case "\n\t":
        if (1 != 296) {
            for fi4 := 0; fi4 < 5; fi4++ {
                var v5 int32 = 302
                var v6 int32 = (v5 + (v5 & (v5 >> ((255) & 63))))
                switch ((("foo bar" + "foo bar") + "") + "foo bar") {
                case "":
                    var v7 int8 = 45
                    _ = v7
                case "a":
                    var v8 uint32 = 0
                    _ = v8
                case "hello":
                    for fi9 := 0; fi9 < 3; fi9++ {
                        v6 -= 667
                        var v10 int64 = p0
                        v11 := GenS1{F0: 2}
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                }
                _ = fi4
                _ = v5
                _ = v6
            }
            switch -572 {
            case -155:
                if (42 != 230) {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        if true {
                            var v13 int64 = 9223372036854775807
                            v14 := []int{42, (fi12 << ((198) & 63)), (fi12 ^ fi12)}
                            _ = v13
                            _ = v14
                        } else {
                            if (-1 > 687) {
                                var v15 int8 = -34
                                _ = v15
                            } else {
                                var v16 uint8 = 123
                                v16 = (v16 & 153)
                                if (false && (!(0 > 400))) {
                                    v17 := []int{}
                                    if (18446744073709551615 <= 2) {
                                        var v18 string = "hello"
                                        _ = v18
                                    }
                                    var v19 uint = 4294967295
                                    _ = v17
                                    _ = v19
                                }
                                _ = v16
                            }
                            v20 := []int{(fi12 << ((42) & 63)), (fi12 * fi12)}
                            if (!true) {
                                fmt.Println(fi12, p0)
                            }
                            _ = v20
                        }
                        _ = fi12
                    }
                    v21 := []int{}
                    _ = v21
                } else {
                    v22 := []int{0, -1, 0}
                    _ = v22
                }
            case -1:
                p0 = (p0 - p0)
            case 32768:
                var v23 uint64 = 1
                _ = v23
            }
        } else {
            if (p0 == (p0 * p0)) {
                var v24 uint8 = 103
                var v25 uint8 = v24
                _ = v24
                _ = v25
            }
            p0 += (p0 - p0)
        }
    case "hello":
        var v26 bool = false
        _ = v26
    }
    v27 := []int{-411, 32767, 0, 0}
    p0 = (p0 | (p0 ^ p0))
    _ = p0
    _ = v27
    return 2147483647
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint32 = 1
    v3 := make(map[uint32]uint32)
    fmt.Println(v2, v1)
    v1 = (v1 & v1)
    v28 := GenS2{F0: uint16(v2)}
    var v29 string = ("hello" + ("\n\t" + ""))
    v30 := GenS3{F0: (v29 + (v29 + v29)), F1: 42, F2: 679}
    for fi31 := 0; fi31 < 18; fi31++ {
        fmt.Println(v2)
        _ = fi31
    }
    v1 = (v1 << ((794) & 63))
    for fi32 := 0; fi32 < 7; fi32++ {
        v33 := []int{fi32, 127, fi32, (^(fi32))}
        _ = fi32
        _ = v33
    }
    fmt.Println(v1, v1, v2)
    switch "" {
    case "a":
        var v34 uint64 = 42
        _ = v34
    case "hello":
        v35 := GenS2{F0: 42}
        _ = v35
    default:
        for fi36 := 0; fi36 < 15; fi36++ {
            var v37 uint8 = 0
            _ = fi36
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
