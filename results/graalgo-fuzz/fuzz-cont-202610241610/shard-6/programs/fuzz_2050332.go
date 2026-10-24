package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint64 }
func genFn0(p0 uint16) uint8 {
    p0 = (^(p0))
    if false {
        v6 := []int{116, 32767}
        var v7 uint16 = (p0 >> ((2) & 63))
        if (float64(((0.0 * -2.25) + (3.14159 * (0.0 * (3.14159 - 1.5))))) <= 1e10) {
            v7 -= p0
            v7 = 144
            switch "" {
            case "\n\t":
                p0 = p0
            default:
                switch 933 {
                case -473:
                    switch "" {
                    case "foo bar":
                        v8 := GenS1{F0: -1, F1: 2, F2: 1}
                        _ = v8
                    case "hello":
                        switch "hello" {
                        case "foo bar":
                            for fi9 := 0; fi9 < 2; fi9++ {
                                switch (("\n\t" + "a") + "foo bar") {
                                case "":
                                    for fi10 := 0; fi10 < 3; fi10++ {
                                        var v11 float64 = 3.14159
                                        if true {
                                            switch (-(fi10)) {
                                            case 148:
                                                var v12 int16 = -32768
                                                _ = v12
                                            case -9223372036854775808:
                                                p0 = 32767
                                            }
                                        }
                                        fmt.Println(v11)
                                        _ = fi10
                                        _ = v11
                                    }
                                case "foo bar":
                                    fmt.Println(fi9, p0, p0)
                                }
                                v6 = append(v6, fi9)
                                _ = fi9
                            }
                        }
                    }
                case 32768:
                    if (1 == 100) {
                        var v13 int16 = -588
                        _ = v13
                    }
                case 32767:
                    for fi14 := 0; fi14 < 9; fi14++ {
                        v15 := GenS1{F0: 42, F1: 2147483647, F2: 1}
                        v6 = append(v6, fi14)
                        _ = fi14
                        _ = v15
                    }
                }
            }
        }
        _ = v6
        _ = v7
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(p0)
        for fi17 := 0; fi17 < 5; fi17++ {
            fmt.Println(p0, fi17, p0)
            _ = fi17
        }
        for fi18 := 0; fi18 < 16; fi18++ {
            if false {
                var v19 int = -9223372036854775808
                p0 += (p0 % (((^(p0))) | 1))
                _ = v19
            } else {
                var v20 uint16 = 65535
                _ = v20
            }
            p0 = (p0 + (p0 + 42))
            p0 += p0
            _ = fi18
        }
        _ = fi16
    }
    fmt.Println(p0)
    v21 := []int{32767, -247, 0}
    _ = p0
    _ = v21
    return 61
}

func main() {
    var v1 int32 = -1
    v2 := make(map[uint32]string)
    var v3 uint8 = 100
    var v4 uint8 = 0
    var v5 uint32 = 2147483647
    var v22 int64 = -1
    fmt.Println(v1, v5)
    switch ("hello" + ("hello" + "foo bar")) {
    case "foo bar":
        for fi23 := 0; fi23 < 12; fi23++ {
            for fi24 := 0; fi24 < 6; fi24++ {
                v5 += v5
                _ = fi24
            }
            _ = fi23
        }
    case "hello":
        var v25 uint = 42
        _ = v25
    case "a":
        var v26 uint64 = 18446744073709551615
        _ = v26
    default:
        var v27 int = -201
        _ = v27
    }
    v4 -= v4
    for fi28 := 0; fi28 < 8; fi28++ {
        v29 := []int{4294967295, fi28, (^(fi28)), (fi28 * (fi28 % (((fi28 + -754)) | 1)))}
        _ = len(v2)
        var v30 uint32 = v5
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 uint32 = v5
    var v32 bool = false
    var v33 uint32 = v31
    for fi34 := 0; fi34 < 20; fi34++ {
        v2[42] = ""
        if true {
            v33 = v31
        }
        _ = fi34
    }
    switch "a" {
    case "":
        var v35 int32 = (v1 * (v1 % ((-890) | 1)))
        _ = v35
    }
    var v36 int32 = v1
    var v37 int8 = 127
    v1 = 2147483647
    switch 386 {
    case 127:
        v32 = v32
    }
    var v38 uint = 887
    v39 := GenS1{F0: -39, F1: 2, F2: 987}
    v40 := GenS1{F0: (v37 << ((77) & 63)), F1: v31, F2: 0}
    var v41 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
