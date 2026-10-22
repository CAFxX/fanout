package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint8) string {
    v5 := GenS1{F0: 0}
    v6 := GenS1{F0: 1}
    v7 := GenS2{F0: 42}
    fmt.Println(p0)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    var v4 uint32 = 65535
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 uint64 = 1
        _ = fi8
        _ = v9
    }
    _ = len(v2)
    v1 -= v1
    if (!true) {
        switch "" {
        case "":
            v4 = v4
        default:
            var v10 uint = uint(v1)
            _ = v10
        }
        fmt.Println(v1)
        switch 4294967295 {
        case 32768:
            if (!true) {
                v1 = v1
                switch 128 {
                case -287:
                    var v11 int8 = 127
                    _ = v11
                case 127:
                    switch -2147483648 {
                    case -9223372036854775808:
                        if false {
                            var v12 uint16 = 484
                            var v13 int = -1
                            var v14 int16 = -720
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    case 32768:
                        if (0 < -26) {
                            if false {
                                for fi15 := 0; fi15 < 1; fi15++ {
                                    var v16 int64 = 1
                                    v17 := GenS1{F0: 0}
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                }
                                var v18 int = 9223372036854775807
                                for fi19 := 0; fi19 < 14; fi19++ {
                                    var v20 int16 = -1
                                    var v21 uint64 = 18446744073709551615
                                    _ = fi19
                                    _ = v20
                                    _ = v21
                                }
                                _ = v18
                            }
                        } else {
                            var v22 uint32 = 2147483648
                            v22 += v4
                            _ = v22
                        }
                    }
                case 107:
                    v23 := GenS1{F0: 0}
                    _ = v23
                }
                switch "\n\t" {
                case "hello":
                    v24 := GenS1{F0: 0}
                    _ = v24
                default:
                    v2[("a" + "hello")] = ""
                }
            }
        case -9223372036854775808:
            for fi25 := 0; fi25 < 8; fi25++ {
                v4 = 65535
                _ = fi25
            }
        default:
            switch "hello" {
            case "a":
                var v26 uint64 = 4294967295
                _ = v26
            case "":
                v4 -= v4
            case "hello":
                for fi27 := 0; fi27 < 1; fi27++ {
                    if false {
                        v28 := []int{173, (fi27 >> ((1) & 63)), fi27}
                        var v29 uint32 = v4
                        _ = v28
                        _ = v29
                    }
                    _ = fi27
                }
            default:
                if true {
                    for fi30 := 0; fi30 < 4; fi30++ {
                        if (false || false) {
                            v4 += 474
                        }
                        for fi31 := 0; fi31 < 18; fi31++ {
                            v32 := []int{}
                            _ = fi31
                            _ = v32
                        }
                        _ = fi30
                    }
                    var v33 string = ""
                    _ = v33
                }
            }
        }
    }
    if (9223372036854775807 >= 32768) {
        v4 = (v1 & (v4 & (^(v1))))
    }
    var v34 int = -762
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 bool = true
        v37 := []int{fi35}
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 int16 = 1
    for fi39 := 0; fi39 < 10; fi39++ {
        if false {
            switch "foo bar" {
            case "foo bar":
                var v40 uint = 2
                _ = v40
            case "":
                var v41 int8 = 0
                _ = v41
            case "hello":
                v42 := []int{(v34 * fi39), (v34 / ((-9223372036854775808) | 1)), fi39}
                _ = v42
            default:
                var v43 int16 = v38
                _ = v43
            }
            if false {
                v44 := GenS3{F0: int32(v1)}
                _ = v44
            }
        }
        _ = fi39
    }
    v34 -= v34
    var v45 bool = true
    v46 := GenS2{F0: -32768}
    fmt.Println(v1, v45)
    switch "foo bar" {
    case "foo bar":
        v38 = (^(v38))
    }
    fmt.Println(v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v45
    _ = v46
}
