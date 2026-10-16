package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int64; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint32; F1 uint }

func main() {
    var v1 int = 2147483647
    v2 := make(map[uint32]bool)
    v3 := []int{(v1 & v1)}
    v4 := []string{}
    v5 := []uint32{65535, 0}
    v6 := []string{"", ("a" + "\n\t"), "", "\n\t"}
    fmt.Println(v1)
    var v7 uint16 = 0
    fmt.Println(v7, v1)
    fmt.Println(v1, v7)
    switch (("hello" + "foo bar") + "\n\t") {
    case "\n\t":
        v8 := GenS1{F0: (4294967295 <= 2), F1: 65535}
        _ = v8
    default:
        fmt.Println(v1, v7)
    }
    var v9 uint32 = 0
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v10 := GenS1{F0: true, F1: 2147483648}
        v6 = append(v6, ("hello" + "\n\t"))
        _ = v10
    }
    switch ("\n\t" + "a") {
    case "a":
        v2[v9] = true
    case "\n\t":
        v11 := GenS1{F0: (true && true), F1: (^(v9))}
        _ = v11
    default:
        switch -2147483648 {
        case 4294967295:
            switch "" {
            case "":
                for fi12 := 0; fi12 < 20; fi12++ {
                    v7 -= (^(v7))
                    _ = fi12
                }
            case "a":
                if false {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        var v14 uint32 = v9
                        if false {
                            for fi15 := 0; fi15 < 15; fi15++ {
                                var v16 uint8 = 0
                                var v17 int32 = 1
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            v14 += v9
                        }
                        fmt.Println(v1)
                        _ = fi13
                        _ = v14
                    }
                    fmt.Println(v1, v7, v9)
                }
            default:
                switch ("hello" + "\n\t") {
                case "hello":
                    switch v1 {
                    case 343:
                        if false {
                            v18 := GenS1{F0: (-1 != 0), F1: (v9 / ((v9) | 1))}
                            _ = v18
                        } else {
                            var v19 uint8 = 42
                            v20 := GenS2{F0: (3.14159 == float64(1.5)), F1: v9, F2: (^(v7))}
                            _ = v19
                            _ = v20
                        }
                    case -2147483648:
                        v21 := GenS1{F0: false, F1: (v9 + 2147483648)}
                        _ = v21
                    default:
                        var v22 int = -591
                        _ = v22
                    }
                case "foo bar":
                    var v23 bool = (42 >= 828)
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 12; fi24++ {
                        fmt.Println(fi24, v1)
                        _ = fi24
                    }
                }
            }
        default:
            v2[(^(v9))] = true
        }
    }
    if (-94 > 0) {
        if (true && true) {
            v1 = 128
            v9 = (v9 - (v9 / (((^(v9))) | 1)))
            v1 += v1
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 int64 = 1
            _ = fi25
            _ = v26
        }
    } else {
        if (((v7 & 201) < v7) || false) {
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 bool = true
                var v29 int = -2147483648
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        if (97 == 127) {
            for fi30 := 0; fi30 < 4; fi30++ {
                var v31 int = v1
                _ = fi30
                _ = v31
            }
            v1 = -1
            v32 := GenS2{F0: false, F1: 2147483648, F2: 32767}
            _ = v32
        }
    }
    var v33 int16 = 1000
    v7 += 65535
    v1 += 2147483647
    if ("\n\t" == "") {
        switch v1 {
        case 1:
            _ = len(v4)
        default:
            switch (v1 * v1) {
            case -60:
                v34 := GenS3{F0: 1, F1: 518, F2: 2}
                _ = v34
            default:
                fmt.Println(v7, v9, v1)
            }
        }
        v35 := GenS4{F0: v7, F1: 9223372036854775807}
        v33 -= 42
        _ = v35
    }
    fmt.Println(v9, v7, v7)
    _ = len(v4)
    if true {
        v7 += 1
    } else {
        var v36 uint8 = 42
        v7 += v7
        v36 = (v36 & v36)
        _ = v36
    }
    v37 := GenS5{F0: 2147483647, F1: 42}
    switch "a" {
    case "a":
        v7 = (v7 & v7)
    case "foo bar":
        _ = len(v2)
    case "":
        if false {
            var v38 int8 = 25
            var v39 int32 = -2147483648
            v40 := GenS2{F0: (((v9 > 325) || true) || (((1.5 * 0.0) * (1.5 - (1.5 * 1e10))) <= 3.14159)), F1: (v9 / (((v9 << ((4294967295) & 63))) | 1)), F2: (v7 + (v7 % ((v7) | 1)))}
            _ = v38
            _ = v39
            _ = v40
        } else {
            v9 = v9
            v33 = v33
            v4 = append(v4, ("a" + ""))
        }
    default:
        v9 -= (v9 >> ((65535) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
