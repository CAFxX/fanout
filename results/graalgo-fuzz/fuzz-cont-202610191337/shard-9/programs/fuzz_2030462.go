package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    var v3 bool = (889 != 2147483647)
    v4 := make(map[uint32]uint32)
    var v5 int32 = 0
    var v6 float64 = 1e10
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(v6, fi7, v6)
        _ = fi7
    }
    if v2 {
        if (42 == 32767) {
            for fi8 := 0; fi8 < 2; fi8++ {
                for fi9 := 0; fi9 < 19; fi9++ {
                    v6 = v6
                    v10 := []int{}
                    var v11 float32 = 3.14159
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = fi8
            }
            for fi12 := 0; fi12 < 2; fi12++ {
                for fi13 := 0; fi13 < 6; fi13++ {
                    var v14 float64 = ((0.1 / 1.0) + (3.14159 / (v6 * v6)))
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
        } else {
            fmt.Println(v3)
        }
    } else {
        v15 := []int{1}
        _ = v15
    }
    switch ("" + "hello") {
    case "a":
        var v16 uint16 = 1
        _ = v16
    case "":
        v6 = 3.14159
    case "hello":
        v17 := []int{-750, 42}
        _ = v17
    default:
        if ((1e10 / (0.1 + float32((((v6 - v6) - 3.14159) * (v6 / 1.0))))) >= 3.14159) {
            v6 = (v6 + 3.14159)
            v18 := GenS1{F0: 671, F1: 19, F2: 1}
            v19 := GenS2{F0: v5, F1: 883}
            _ = v18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 int16 = 1
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 int = fi22
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    if ("foo bar" < "hello") {
        v24 := []int{9223372036854775807, 316}
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 int16 = -32768
            _ = fi25
            _ = v26
        }
        switch "" {
        case "foo bar":
            var v27 int = -9223372036854775808
            _ = v27
        case "hello":
            v3 = (1 != 4294967295)
        default:
            if v3 {
                _ = len(v24)
                fmt.Println(v6, v5)
            }
        }
        _ = v24
    }
    if true {
        v28 := []int{}
        var v29 uint64 = 42
        var v30 float32 = (0.0 - 3.14159)
        _ = v28
        _ = v29
        _ = v30
    }
    if v2 {
        switch "hello" {
        case "a":
            var v31 string = "hello"
            _ = v31
        default:
            var v32 uint64 = 1
            _ = v32
        }
    } else {
        fmt.Println(v6)
    }
    for fi33 := 0; fi33 < 13; fi33++ {
        v3 = false
        if (0 > 1) {
            var v34 bool = v2
            delete(v4, 2)
            _ = v34
        }
        _ = fi33
    }
    v35 := GenS1{F0: 316, F1: 0, F2: 451}
    fmt.Println(v6, v3, v2)
    v36 := v1[96]
    var v37 bool = (0.0 <= (1e10 * 1e10))
    fmt.Println(v3, v5, v37)
    _ = len(v1)
    for fi38 := 0; fi38 < 15; fi38++ {
        for fi39 := 0; fi39 < 5; fi39++ {
            if (v5 == (v5 / ((-1) | 1))) {
                if (!false) {
                    v37 = v3
                    v5 = v5
                    switch "hello" {
                    case "a":
                        v6 += (float64(0.0) * v6)
                    case "":
                        switch (^(fi38)) {
                        case 32767:
                            fmt.Println(v37)
                            switch 898 {
                            case -2147483648:
                                delete(v1, 334)
                            case 32767:
                                v40 := GenS1{F0: 2147483648, F1: 290, F2: 1}
                                _ = v40
                            case 233:
                                var v41 string = (("a" + "hello") + ("a" + "foo bar"))
                                _ = v41
                            default:
                                var v42 uint = 65535
                                _ = v42
                            }
                            fmt.Println(v2)
                            var v43 int16 = 1000
                            _ = v43
                        }
                    case "\n\t":
                        v2 = v37
                    default:
                        v44 := GenS3{F0: true}
                        _ = v44
                    }
                } else {
                    v2 = v2
                }
                v37 = (4294967295 < 436)
                v6 = v6
            }
            v37 = (v37 && true)
            _ = fi39
        }
        var v45 int8 = 0
        v46 := []int{int(v45), -2147483648, -1}
        _ = fi38
        _ = v45
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
