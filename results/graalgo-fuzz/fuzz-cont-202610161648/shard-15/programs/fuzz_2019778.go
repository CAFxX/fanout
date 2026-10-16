package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 string }
type GenS2 struct { F0 int32; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int8, p1 uint16, p2 uint) int32 {
    var v5 float64 = ((0.0 - 3.14159) * 3.14159)
    var v6 uint16 = p1
    v7 := GenS1{F0: (p1 | v6), F1: 941, F2: "\n\t"}
    fmt.Println(p1)
    var v8 int16 = -32768
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return -834
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]int)
    var v3 int32 = 2147483647
    var v4 float64 = 3.14159
    if false {
        if (!false) {
            var v9 string = "\n\t"
            v10 := []int{2147483647, 4294967295, -2147483648}
            _ = len(v10)
            _ = v10
            _ = v9
        } else {
            fmt.Println(v3, v4)
            v11 := []int{-9223372036854775808, -1, 128, 32768}
            for fi12 := 0; fi12 < 17; fi12++ {
                switch "hello" {
                case "":
                    v13 := GenS1{F0: 2, F1: fi12, F2: "hello"}
                    _ = v13
                case "a":
                    for fi14 := 0; fi14 < 19; fi14++ {
                        v15 := GenS2{F0: v3, F1: v3, F2: 279}
                        _ = fi14
                        _ = v15
                    }
                default:
                    if (!false) {
                        if true {
                            for fi16 := 0; fi16 < 18; fi16++ {
                                v17 := GenS1{F0: 0, F1: fi12, F2: "\n\t"}
                                _ = fi16
                                _ = v17
                            }
                        }
                        for fi18 := 0; fi18 < 20; fi18++ {
                            var v19 uint16 = 42
                            _ = fi18
                            _ = v19
                        }
                        switch "hello" {
                        case "":
                            v4 = v4
                        case "foo bar":
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        case "\n\t":
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        default:
                            var v20 bool = true
                            _ = v20
                        }
                    } else {
                        var v21 float64 = v4
                        var v22 string = ((("\n\t" + "foo bar") + "hello") + "hello")
                        fmt.Println(v22, v4, v21)
                        _ = v21
                        _ = v22
                    }
                }
                v4 -= (((v4 / float64(1.5)) * v4) * 0.1)
                delete(v2, "a")
                _ = fi12
            }
            _ = v11
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            for fi24 := 0; fi24 < 11; fi24++ {
                v25 := []int{(fi23 / (((fi23 | fi23)) | 1)), 0, (fi23 + fi23), fi24}
                var v26 int8 = 127
                if (true || true) {
                    v26 += v26
                    v3 += v3
                    v2[("foo bar" + "\n\t")] = fi23
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            v3 = 1
            if false {
                for fi27 := 0; fi27 < 18; fi27++ {
                    fmt.Println(v3, fi23)
                    v28 := GenS2{F0: v3, F1: v3, F2: 58}
                    _ = fi27
                    _ = v28
                }
                v3 -= -2147483648
            }
            _ = fi23
        }
        v29 := []int{}
        _ = v29
    }
    v3 += v3
    v30 := GenS1{F0: 0, F1: 4294967295, F2: "\n\t"}
    if (float64((3.14159 + float32(v4))) == 1e10) {
        fmt.Println(v4)
    } else {
        v3 = genFn0(127, 42, 18446744073709551615)
        v31 := []int{-885}
        _ = v31
    }
    fmt.Println(v3, v3)
    v3 = 1
    v3 = v3
    switch 127 {
    case -9223372036854775808:
        var v32 int32 = v3
        _ = v32
    }
    switch "" {
    case "hello":
        v33 := GenS3{F0: 613}
        _ = v33
    case "":
        v34 := GenS1{F0: 32767, F1: 743, F2: ""}
        _ = v34
    case "a":
        switch 42 {
        case -501:
            var v35 int = -626
            _ = v35
        case 461:
            v3 = v3
        default:
            var v36 uint32 = 882
            _ = v36
        }
    default:
        if false {
            for fi37 := 0; fi37 < 2; fi37++ {
                var v38 uint16 = 396
                _ = fi37
                _ = v38
            }
            v39 := GenS3{F0: 2147483647}
            _ = v39
        } else {
            for fi40 := 0; fi40 < 17; fi40++ {
                fmt.Println(v4, v3)
                _ = fi40
            }
            v41 := GenS1{F0: 65535, F1: 32767, F2: "hello"}
            v3 -= (v3 << ((42) & 63))
            _ = v41
        }
    }
    var v42 string = "foo bar"
    v4 -= (v4 - v4)
    v3 = v3
    for fi43 := 0; fi43 < 10; fi43++ {
        fmt.Println(v4)
        var v44 int32 = (v3 & (v3 + v3))
        delete(v2, v42)
        _ = fi43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v42
}
