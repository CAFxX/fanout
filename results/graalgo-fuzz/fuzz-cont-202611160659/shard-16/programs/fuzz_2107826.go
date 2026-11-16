package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 int8; F2 string }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 2147483647
    var v5 uint8 = 5
    v6 := make(map[int]bool)
    var v7 string = ("" + ((("a" + ("a" + "hello")) + ((("hello" + ("" + "hello")) + "") + "hello")) + "foo bar"))
    for fi8 := 0; fi8 < 12; fi8++ {
        if ((1e10 < 3.14159) && false) {
            var v9 int64 = (v1 - (^(v1)))
            _ = v9
        } else {
            if false {
                for fi10 := 0; fi10 < 12; fi10++ {
                    switch fi10 {
                    case 0:
                        _ = len(v2)
                    case 42:
                        v11 := v6[fi8]
                        _ = v11
                    }
                    if (!((^(v4)) <= v4)) {
                        var v12 int = fi10
                        if ((((v1 / ((-1) | 1)) == -9223372036854775808) || false) && (("" > ("" + "")) || (0 >= 0))) {
                            v4 = v4
                        }
                        var v13 uint64 = 411
                        v12 = (fi8 ^ (^(fi8)))
                        var v14 int32 = 2147483647
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        fmt.Println(v4, v4)
                    }
                    switch "\n\t" {
                    case "hello":
                        var v15 uint16 = 792
                        _ = v15
                    default:
                        v16 := GenS1{F0: false, F1: 42}
                        _ = v16
                    }
                    var v17 float64 = 3.14159
                    v4 = v4
                    v18 := []int{fi8, (fi10 % ((fi8) | 1)), (fi8 + -1)}
                    _ = fi10
                    _ = v17
                    _ = v18
                }
                v5 = v5
            } else {
                v19 := []int{(fi8 - (^(fi8))), -1, fi8}
                _ = v19
            }
        }
        if true {
            v20 := GenS1{F0: (18446744073709551615 == 42), F1: 32767}
            for fi21 := 0; fi21 < 18; fi21++ {
                v22 := GenS1{F0: false, F1: 0}
                switch (v7 + ((v7 + v7) + "foo bar")) {
                case "":
                    v5 -= v5
                default:
                    fmt.Println(v7, v7, v1)
                }
                var v23 uint8 = v5
                _ = fi21
                _ = v22
                _ = v23
            }
            v24 := GenS1{F0: false, F1: 65535}
            _ = v20
            _ = v24
        } else {
            for fi25 := 0; fi25 < 1; fi25++ {
                if (403 <= 2) {
                    if (uint64(v5) == 784) {
                        var v26 bool = false
                        if v26 {
                            v27 := GenS2{F0: v5, F1: 255, F2: v1}
                            var v28 bool = ((true && v26) || v26)
                            fmt.Println(v28, fi8, v1)
                            _ = v27
                            _ = v28
                        }
                        _ = v26
                    }
                    v29 := []int{(fi25 >> ((42) & 63)), fi25, (fi25 + fi8)}
                    v5 = 1
                    var v30 int8 = -45
                    var v31 uint = 430
                    _ = v29
                    _ = v30
                    _ = v31
                } else {
                    fmt.Println(fi25)
                    var v32 uint8 = v5
                    delete(v6, (fi25 * fi25))
                    _ = v32
                }
                switch 4294967295 {
                case 32767:
                    v7 = v7
                case 659:
                    v33 := []int{(fi25 >> ((247) & 63)), fi25, 0, -1}
                    _ = v33
                case -416:
                    v34 := GenS1{F0: ((!(0 < 813)) && (int16(v5) >= 66)), F1: 42}
                    _ = v34
                default:
                    var v35 int8 = int8(fi25)
                    _ = v35
                }
                _ = fi25
            }
            fmt.Println(v4, v1, v4)
            if true {
                fmt.Println(fi8, v1, fi8)
                fmt.Println(fi8, v1, v5)
            }
        }
        _ = fi8
    }
    v4 -= (^(v4))
    v5 = v5
    v7 = v7
    var v36 float64 = float64(1.5)
    v7 = "\n\t"
    v7 = v7
    switch "foo bar" {
    case "foo bar":
        v7 = v7
    case "\n\t":
        var v37 int64 = v1
        _ = v37
    default:
        if true {
            if (!false) {
                fmt.Println(v1)
                if (1 < -1) {
                    v5 += (v5 / (((v5 % (((v5 << ((0) & 63))) | 1))) | 1))
                    var v38 uint8 = (v5 ^ 0)
                    _ = v38
                }
                if (!(444 <= 42)) {
                    fmt.Println(v5, v5, v4)
                    if false {
                        v39 := GenS1{F0: (!true), F1: 2}
                        _ = v39
                    }
                    if true {
                        var v40 int32 = 830
                        _ = v40
                    }
                }
            } else {
                switch 128 {
                case -739:
                    v7 = v7
                default:
                    delete(v2, v4)
                }
                v41 := GenS3{F0: 1}
                v7 = ""
                _ = v41
            }
            v42 := []int{302, 127}
            v43 := v2[(^(v4))]
            _ = v42
            _ = v43
        }
    }
    v5 = (v5 << ((65535) & 63))
    delete(v6, -1)
    var v44 uint16 = 0
    for fi45 := 0; fi45 < 16; fi45++ {
        v44 -= 177
        v46 := GenS4{F0: fi45, F1: 127, F2: (v7 + (v7 + v7))}
        fmt.Println(v7, v36)
        _ = fi45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v7
}
