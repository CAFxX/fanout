package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 0.1
    var v3 uint64 = 917
    v4 := []uint32{4294967295, 42, 65535}
    v5 := []int{750, 32767, 38, 0}
    v6 := []string{"", ("" + "")}
    for fi7 := 0; fi7 < 13; fi7++ {
        v3 -= (v3 & v3)
        _ = fi7
    }
    if (false && (-1 <= -504)) {
        var v8 uint = 4294967295
        var v9 int32 = 42
        for fi10 := 0; fi10 < 4; fi10++ {
            switch (fi10 % ((fi10) | 1)) {
            case 313:
                v9 -= -768
            case -9223372036854775808:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi10
        }
        _ = v8
        _ = v9
    }
    v3 += 42
    v5 = append(v5, 32767)
    v11 := GenS1{F0: 11, F1: -9223372036854775808}
    for fi12 := 0; fi12 < 5; fi12++ {
        v5 = append(v5, (fi12 ^ fi12))
        var v13 float64 = 1e10
        _ = fi12
        _ = v13
    }
    var v14 int16 = 32767
    var v15 float64 = (3.14159 / 0.1)
    var v16 int16 = 1000
    v17 := GenS1{F0: 127, F1: -1}
    for fi18 := 0; fi18 < 15; fi18++ {
        v3 = 42
        switch "hello" {
        case "hello":
            var v19 int8 = -49
            _ = v19
        case "foo bar":
            var v20 float64 = 1e10
            _ = v20
        case "\n\t":
            fmt.Println(v14, v3)
        }
        v1[2] = 875
        _ = fi18
    }
    _ = len(v5)
    if ((v16 >> ((4294967295) & 63)) < (v16 % ((v16) | 1))) {
        var v21 int64 = 564
        for fi22 := 0; fi22 < 8; fi22++ {
            switch "" {
            case "":
                var v23 bool = (-1 != -965)
                _ = v23
            case "hello":
                for fi24 := 0; fi24 < 4; fi24++ {
                    v3 += v3
                    switch ("" + "a") {
                    case "\n\t":
                        if (false || false) {
                            v15 = 3.14159
                        }
                    case "hello":
                        delete(v1, 2147483648)
                    case "a":
                        v3 -= v3
                    default:
                        if (fi22 <= (fi24 % ((-1) | 1))) {
                            v3 -= 18446744073709551615
                        } else {
                            v15 = v15
                            switch (fi22 - (fi22 >> ((4294967295) & 63))) {
                            case -201:
                                var v25 uint16 = 871
                                _ = v25
                            case 0:
                                v26 := GenS2{F0: 0}
                                _ = v26
                            case 127:
                                switch (^(fi22)) {
                                case 127:
                                    v2 = (3.14159 - v2)
                                case 32768:
                                    var v27 int8 = -1
                                    _ = v27
                                case 42:
                                    v16 -= (v14 + v14)
                                }
                            default:
                                v14 = v14
                            }
                            v16 -= v16
                            v14 += v16
                            _ = len(v4)
                            if (760 > v3) {
                                var v28 bool = false
                                _ = v28
                            } else {
                                v29 := GenS1{F0: 100, F1: v21}
                                var v30 uint64 = (^(v3))
                                _ = v29
                                _ = v30
                            }
                            var v31 int32 = 851
                            if false {
                                if (!(0 != -128)) {
                                    _ = len(v1)
                                    fmt.Println(v15, v14, fi24)
                                } else {
                                    if true {
                                        v32 := GenS2{F0: 570}
                                        var v33 int16 = (^(v16))
                                        v34 := GenS3{F0: v31}
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    }
                                }
                                switch "a" {
                                case "a":
                                    var v35 float64 = v15
                                    _ = v35
                                case "foo bar":
                                    var v36 uint8 = 107
                                    _ = v36
                                default:
                                    v31 -= 42
                                }
                            }
                            var v37 int16 = (v14 ^ 1)
                            v38 := GenS4{F0: 1}
                            _ = v31
                            _ = v37
                            _ = v38
                        }
                    }
                    v39 := GenS5{F0: 0, F1: 1, F2: 65535}
                    var v40 int = (^(fi22))
                    if true {
                        _ = len(v1)
                        v41 := GenS3{F0: 2147483647}
                        var v42 bool = true
                        _ = v41
                        _ = v42
                    }
                    _ = fi24
                    _ = v39
                    _ = v40
                }
            }
            _ = fi22
        }
        _ = v21
    }
    if (42 >= 42) {
        v16 -= -32768
        for fi43 := 0; fi43 < 17; fi43++ {
            if true {
                var v44 uint32 = 47
                _ = v44
            } else {
                var v45 float64 = 3.14159
                v14 += (v14 | (v14 + v16))
                _ = v45
            }
            switch fi43 {
            case -549:
                var v46 int = (^(fi43))
                _ = v46
            case 2147483647:
                var v47 float64 = v15
                _ = v47
            case -329:
                if false {
                    fmt.Println(v16, fi43)
                    var v48 float32 = 0.1
                    _ = v48
                }
            default:
                v2 = v2
            }
            v14 -= v16
            _ = fi43
        }
        var v49 float32 = v2
        _ = v49
    }
    var v50 float32 = (-2.25 + (v2 - (v2 + v2)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
