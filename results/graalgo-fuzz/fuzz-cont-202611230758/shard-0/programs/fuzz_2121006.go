package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }

func main() {
    var v1 string = ("hello" + "a")
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    var v5 uint16 = 65535
    var v6 int32 = 1
    var v7 uint64 = 960
    if (true || true) {
        fmt.Println(v5, v5, v7)
        switch -1 {
        case 42:
            switch v1 {
            case "":
                switch 32767 {
                case 2147483647:
                    for fi8 := 0; fi8 < 9; fi8++ {
                        v5 += v5
                        fmt.Println(v5, v7, v6)
                        _ = fi8
                    }
                case 1:
                    var v9 uint = 2
                    _ = v9
                default:
                    v10 := GenS1{F0: 82}
                    _ = v10
                }
            case "foo bar":
                var v11 int64 = -9223372036854775808
                _ = v11
            case "hello":
                v12 := []int{-9223372036854775808, int(v5)}
                _ = v12
            default:
                switch 32767 {
                case 32767:
                    switch (v1 + "a") {
                    case "\n\t":
                        var v13 uint8 = 2
                        _ = v13
                    case "foo bar":
                        var v14 int8 = -117
                        _ = v14
                    case "hello":
                        var v15 uint32 = 548
                        _ = v15
                    default:
                        for fi16 := 0; fi16 < 2; fi16++ {
                            switch ((v1 + ((v1 + v1) + v1)) + "\n\t") {
                            case "\n\t":
                                var v17 int16 = -1
                                _ = v17
                            case "":
                                switch (("hello" + "") + v1) {
                                case "hello":
                                    var v18 uint8 = 128
                                    _ = v18
                                }
                            default:
                                v19 := []int{(fi16 * (fi16 - fi16)), -301, fi16, (fi16 * 32767)}
                                _ = v19
                            }
                            _ = fi16
                        }
                    }
                default:
                    var v20 bool = (false || (32768 > 32767))
                    _ = v20
                }
            }
        case 128:
            for fi21 := 0; fi21 < 19; fi21++ {
                v5 = v5
                switch v1 {
                case "\n\t":
                    var v22 float32 = 0.0
                    _ = v22
                case "hello":
                    fmt.Println(v6)
                case "":
                    var v23 int8 = -124
                    _ = v23
                }
                _ = fi21
            }
        }
    }
    v5 += 0
    v5 = (v5 >> ((uint(v7)) & 63))
    for fi24 := 0; fi24 < 11; fi24++ {
        switch -1 {
        case -1:
            var v25 float32 = (3.14159 + 1.5)
            _ = v25
        case -689:
            var v26 int16 = -572
            _ = v26
        case -2147483648:
            for fi27 := 0; fi27 < 12; fi27++ {
                fmt.Println(v5)
                if true {
                    var v28 float64 = ((0.0 + (0.1 / (((((-2.25 - 0.0) + 1.5) + 3.14159) / 1.5) / 0.1))) - (0.1 + (0.1 / (float64(float32(1e10)) - 0.1))))
                    var v29 uint = 2
                    _ = v28
                    _ = v29
                }
                if false {
                    if false {
                        var v30 uint16 = (^(v5))
                        var v31 int8 = 0
                        _ = v30
                        _ = v31
                    } else {
                        switch (fi24 << ((460) & 63)) {
                        case -9223372036854775808:
                            var v32 uint8 = 255
                            _ = v32
                        case 32767:
                            fmt.Println(v5, v1, fi27)
                        case 32768:
                            v33 := []int{fi24, fi27}
                            _ = v33
                        default:
                            fmt.Println(v5, v5)
                        }
                    }
                    if ((fi24 << ((18446744073709551615) & 63)) <= fi27) {
                        fmt.Println(v1)
                    }
                }
                v34 := GenS1{F0: -1}
                _ = fi27
                _ = v34
            }
        default:
            if (4294967295 >= (^(v7))) {
                var v35 float64 = 1e10
                v36 := GenS2{F0: false}
                switch (fi24 & fi24) {
                case -2147483648:
                    if false {
                        switch v1 {
                        case "a":
                            if (int32(fi24) <= v6) {
                                v6 = v6
                                if (0 >= 8) {
                                    v1 = v1
                                    var v37 int8 = 127
                                    _ = v37
                                }
                            } else {
                                switch (("foo bar" + "") + (("" + "hello") + v1)) {
                                case "foo bar":
                                    v38 := v4[v1]
                                    _ = v38
                                case "\n\t":
                                    v1 = "hello"
                                default:
                                    switch "hello" {
                                    case "foo bar":
                                        var v39 bool = (!true)
                                        _ = v39
                                    case "hello":
                                        v7 = (v7 | v7)
                                    }
                                }
                                v4[v1] = fi24
                            }
                        case "":
                            fmt.Println(v6, v35)
                        default:
                            v40 := []int{fi24}
                            _ = v40
                        }
                    } else {
                        for fi41 := 0; fi41 < 2; fi41++ {
                            v42 := GenS2{F0: false}
                            _ = fi41
                            _ = v42
                        }
                    }
                case 4294967295:
                    if (!(!((v6 + (v6 * (^(v6)))) == v6))) {
                        v6 += v6
                        fmt.Println(v7)
                    }
                default:
                    var v43 int64 = 9223372036854775807
                    _ = v43
                }
                _ = v35
                _ = v36
            }
        }
        _ = fi24
    }
    for fi44 := 0; fi44 < 10; fi44++ {
        v45 := v2[0]
        _ = fi44
        _ = v45
    }
    v1 = (v1 + "a")
    if false {
        var v46 uint32 = 63
        var v47 int64 = -1
        _ = v46
        _ = v47
    } else {
        var v48 int32 = (^(v6))
        v49 := GenS1{F0: 100}
        delete(v4, "\n\t")
        _ = v48
        _ = v49
    }
    v1 = (v1 + (v1 + v1))
    var v50 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v7
}
