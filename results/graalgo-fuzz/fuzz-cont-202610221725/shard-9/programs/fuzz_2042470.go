package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 uint32 = 1
    var v2 int32 = 928
    var v3 uint16 = 255
    v4 := make(map[int]int)
    v5 := make(map[int]int)
    v6 := make(map[string]bool)
    _ = len(v4)
    delete(v6, "foo bar")
    if (859 == 827) {
        v7 := []int{}
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS1{F0: 42, F1: 59}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi8
            _ = v9
        }
        if (true || false) {
            if false {
                v2 = v2
                var v10 uint16 = 0
                _ = v10
            }
            v3 += (v3 % (((v3 >> ((255) & 63))) | 1))
        } else {
            _ = len(v4)
            v1 -= v1
            v7 = append(v7, 76)
        }
        _ = v7
    } else {
        fmt.Println(v1, v1, v2)
    }
    if true {
        switch 32767 {
        case -9223372036854775808:
            if true {
                var v11 uint16 = (v3 ^ uint16(v1))
                for fi12 := 0; fi12 < 10; fi12++ {
                    v13 := []int{-293}
                    var v14 int16 = -1
                    switch "" {
                    case "":
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    case "a":
                        v14 -= -69
                    case "hello":
                        v15 := GenS1{F0: 0, F1: 0}
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                _ = v11
            } else {
                switch 360 {
                case -2147483648:
                    switch 127 {
                    case -890:
                        var v16 uint8 = 42
                        _ = v16
                    default:
                        if true {
                            var v17 uint32 = v1
                            _ = v17
                        } else {
                            for fi18 := 0; fi18 < 11; fi18++ {
                                if (((v3 * v3) >= 1) || true) {
                                    v3 -= 255
                                }
                                fmt.Println(v1, v3, v2)
                                _ = fi18
                            }
                            if true {
                                fmt.Println(v3, v2)
                            }
                        }
                    }
                case 42:
                    v19 := []int{}
                    _ = v19
                case 32767:
                    for fi20 := 0; fi20 < 20; fi20++ {
                        for fi21 := 0; fi21 < 10; fi21++ {
                            var v22 int32 = int32(v3)
                            _ = fi21
                            _ = v22
                        }
                        v23 := []int{}
                        _ = fi20
                        _ = v23
                    }
                }
                v1 = 86
            }
        case 32768:
            for fi24 := 0; fi24 < 14; fi24++ {
                for fi25 := 0; fi25 < 18; fi25++ {
                    v1 += uint32(v3)
                    var v26 int8 = 100
                    v27 := GenS1{F0: 2, F1: uint8(v3)}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = fi24
            }
        }
        var v28 float32 = (float32(0.0) * (0.1 / 0.1))
        _ = v28
    } else {
        for fi29 := 0; fi29 < 19; fi29++ {
            v2 = 42
            v30 := []int{}
            var v31 uint8 = 117
            _ = fi29
            _ = v30
            _ = v31
        }
        var v32 int32 = (v2 / (((v2 + v2)) | 1))
        _ = v32
    }
    if (!((4294967295 == 18446744073709551615) || true)) {
        var v33 string = "\n\t"
        switch 0 {
        case 128:
            var v34 bool = false
            _ = v34
        }
        _ = v33
    } else {
        _ = len(v5)
    }
    v1 += 2147483648
    for fi35 := 0; fi35 < 18; fi35++ {
        var v36 uint = 860
        fmt.Println(v36)
        _ = fi35
        _ = v36
    }
    fmt.Println(v3, v1, v3)
    if ((!((!(!(float64(0.0) != 1.5))) && true)) || true) {
        if (true && (0 > -1)) {
            fmt.Println(v1, v1)
        } else {
            fmt.Println(v2, v1)
        }
    } else {
        if true {
            v1 = v1
        }
        if (1.5 == 1e10) {
            v37 := []int{-687, 128, 538, 32768}
            v37 = append(v37, 0)
            var v38 int64 = -1
            _ = v37
            _ = v38
        } else {
            v3 = v3
        }
        fmt.Println(v1)
    }
    v39 := []int{}
    switch "" {
    case "\n\t":
        v40 := GenS1{F0: 0, F1: 49}
        _ = v40
    default:
        v2 = (v2 ^ v2)
    }
    fmt.Println(v2, v3)
    for fi41 := 0; fi41 < 16; fi41++ {
        v42 := GenS1{F0: 4294967295, F1: 42}
        v3 += (v3 + 255)
        _ = fi41
        _ = v42
    }
    if (true || ((v3 << ((42) & 63)) == v3)) {
        var v43 int8 = 0
        fmt.Println(v43, v2)
        _ = v43
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
