package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int }
type GenS2 struct { F0 uint; F1 uint; F2 int16 }
type GenS3 struct { F0 int32; F1 uint }
func genFn0(p0 uint64) bool {
    var v7 int8 = 1
    switch -166 {
    case -9223372036854775808:
        if true {
            v7 = v7
            for fi8 := 0; fi8 < 13; fi8++ {
                switch ("a" + (("foo bar" + ("foo bar" + "")) + (("" + "\n\t") + "a"))) {
                case "foo bar":
                    for fi9 := 0; fi9 < 9; fi9++ {
                        var v10 uint = 1
                        var v11 uint32 = 65535
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                }
                _ = fi8
            }
        } else {
            v7 -= v7
        }
    case 0:
        var v12 int64 = -610
        _ = v12
    case -2147483648:
        var v13 uint64 = 0
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := []int{}
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = v7
    return (547 == 9223372036854775807)
}

func main() {
    var v1 uint64 = 363
    var v2 uint32 = 459
    var v3 int8 = 100
    var v4 uint16 = 940
    var v5 uint16 = (v4 + v4)
    v6 := make(map[uint32]string)
    for fi16 := 0; fi16 < 16; fi16++ {
        v3 += v3
        fmt.Println(v1, v5)
        _ = fi16
    }
    v3 -= v3
    for fi17 := 0; fi17 < 11; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            delete(v6, v2)
            if false {
                if (!false) {
                    if true {
                        switch fi18 {
                        case 32768:
                            v19 := GenS1{F0: 1000, F1: 1, F2: 127}
                            _ = v19
                        case -175:
                            var v20 uint8 = uint8(v5)
                            _ = v20
                        }
                        v2 += (v2 >> ((445) & 63))
                        var v21 bool = false
                        if v21 {
                            if (-306 <= -9223372036854775808) {
                                v4 = v5
                                v22 := GenS1{F0: 42, F1: 255, F2: fi18}
                                v23 := []int{}
                                _ = v22
                                _ = v23
                            }
                            var v24 uint16 = v5
                            _ = v24
                        } else {
                            switch (^(fi18)) {
                            case 4294967295:
                                v21 = v21
                                var v25 bool = (0 >= 1)
                                _ = v25
                            }
                            v26 := GenS1{F0: 32767, F1: 42, F2: fi17}
                            v4 = v4
                            switch -2147483648 {
                            case 0:
                                v2 = (v2 >> ((1) & 63))
                            case 32767:
                                if (211 == 852) {
                                    v27 := []int{}
                                    v28 := GenS1{F0: 1000, F1: 128, F2: fi17}
                                    _ = v27
                                    _ = v28
                                }
                            case 128:
                                v29 := GenS2{F0: 987, F1: 6, F2: 1000}
                                _ = v29
                            default:
                                v30 := GenS3{F0: -1, F1: 65535}
                                _ = v30
                            }
                            _ = v26
                        }
                        fmt.Println(v2, v3)
                        _ = v21
                    }
                } else {
                    v6[v2] = "hello"
                }
            }
            v3 = (^(v3))
            _ = fi18
        }
        for fi31 := 0; fi31 < 5; fi31++ {
            var v32 float64 = 3.14159
            var v33 int16 = 1000
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi17
    }
    fmt.Println(v5, v2, v3)
    var v34 int8 = v3
    v2 += (v2 << ((42) & 63))
    v35 := []int{}
    var v36 int32 = 2147483647
    var v37 string = ""
    v38 := GenS3{F0: (v36 * v36), F1: 629}
    if (!true) {
        switch v37 {
        case "foo bar":
            fmt.Println(v34, v4)
        case "\n\t":
            v3 -= v34
        case "a":
            v37 = v37
        default:
            var v39 float64 = (1e10 - 0.1)
            _ = v39
        }
        v37 = (v37 + v37)
        v37 = v37
    } else {
        fmt.Println(v1, v5)
        var v40 int32 = v36
        var v41 uint = 300
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
