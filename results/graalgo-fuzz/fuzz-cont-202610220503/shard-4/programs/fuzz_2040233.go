package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 4294967295
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    v5 := []int{127, 4294967295}
    v6 := make(map[uint32]int)
    var v7 float32 = (1e10 * float32(3.14159))
    v7 = v7
    v7 += float32(((float64(1.5) / 0.1) / ((1.5 - (-2.25 / 3.14159)) + (0.1 + 1e10))))
    v7 += ((v7 - 1.5) - 3.14159)
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint64 = 18446744073709551615
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        if false {
            for fi11 := 0; fi11 < 11; fi11++ {
                v2 += v2
                var v12 uint16 = 375
                v7 = 1.5
                _ = fi11
                _ = v12
            }
        }
        v7 += (1e10 * (v7 / -2.25))
        if (true && ((!(!(2147483647 <= -160))) && (!true))) {
            switch (fi10 | fi10) {
            case 127:
                switch "a" {
                case "\n\t":
                    v2 += v2
                case "a":
                    v13 := GenS1{F0: 100}
                    _ = v13
                }
            case 958:
                v14 := GenS1{F0: 0}
                _ = v14
            default:
                v7 -= 1e10
            }
            v7 -= v7
        }
        _ = fi10
    }
    switch 9223372036854775807 {
    case 1:
        switch -9223372036854775808 {
        case 128:
            _ = len(v1)
        case 252:
            var v15 uint = 900
            _ = v15
        }
    case -1:
        if true {
            _ = len(v1)
            var v16 int8 = 127
            _ = v16
        }
    case 826:
        if false {
            _ = len(v1)
            if (32767 <= -404) {
                v2 = v2
            }
            var v17 uint8 = 255
            _ = v17
        }
    default:
        fmt.Println(v2)
    }
    v18 := GenS1{F0: 0}
    for fi19 := 0; fi19 < 17; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 string = "foo bar"
            if false {
                _ = len(v4)
            } else {
                v7 -= float32(0.0)
            }
            if (((667 > 2) && (4294967295 >= 0)) && false) {
                if (!(false || ((3.14159 + 0.0) == 1.5))) {
                    var v22 int8 = 0
                    _ = v22
                } else {
                    _ = len(v5)
                    v23 := GenS1{F0: 0}
                    v24 := GenS1{F0: 42}
                    _ = v23
                    _ = v24
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                switch 32768 {
                case 128:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case -1:
                    v7 = 0.0
                case -876:
                    var v25 int = fi20
                    _ = v25
                default:
                    v21 = v21
                }
            }
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v26 := GenS1{F0: 50}
    if true {
        v27 := GenS1{F0: 0}
        v5 = append(v5, 2147483647)
        var v28 uint8 = 183
        _ = v27
        _ = v28
    }
    var v29 int64 = -9223372036854775808
    v2 += v2
    if false {
        v30 := GenS2{F0: (v2 & 424), F1: 864, F2: 0}
        _ = v30
    } else {
        v7 -= float32(3.14159)
        var v31 int16 = 42
        v2 = 0
        _ = v31
    }
    if true {
        if true {
            v29 = v29
            switch "foo bar" {
            case "foo bar":
                v7 = (v7 + 1.5)
            default:
                v29 = v29
            }
        } else {
            var v32 int = 127
            for fi33 := 0; fi33 < 3; fi33++ {
                if (false && (42 == 65535)) {
                    var v34 uint64 = (^(v2))
                    _ = v34
                }
                var v35 float64 = 1e10
                delete(v4, "")
                _ = fi33
                _ = v35
            }
            _ = v32
        }
    } else {
        if true {
            for fi36 := 0; fi36 < 8; fi36++ {
                if ((2 < 2147483647) && false) {
                    var v37 int8 = -74
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v37
                } else {
                    if false {
                        fmt.Println(v7, v2, v7)
                        for fi38 := 0; fi38 < 7; fi38++ {
                            var v39 float32 = v7
                            if len(v5) > 0 {
                                v5[0] = (fi36 ^ 459)
                            }
                            _ = fi38
                            _ = v39
                        }
                        v40 := GenS3{F0: 610}
                        _ = v40
                    } else {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        fmt.Println(fi36, v7)
                        v7 += v7
                    }
                    v41 := GenS4{F0: (!(18446744073709551615 < uint(v2)))}
                    for fi42 := 0; fi42 < 10; fi42++ {
                        if false {
                            if false {
                                var v43 float32 = float32((((0.0 + 3.14159) / float64((v7 - 1e10))) * ((1.5 + (0.0 * 0.1)) * 1e10)))
                                var v44 uint16 = 65535
                                var v45 uint = 1
                                _ = v43
                                _ = v44
                                _ = v45
                            } else {
                                var v46 float32 = -2.25
                                _ = v46
                            }
                            v47 := GenS2{F0: 2, F1: 187, F2: 4294967295}
                            var v48 string = "hello"
                            _ = v47
                            _ = v48
                        }
                        v49 := GenS1{F0: 100}
                        _ = fi42
                        _ = v49
                    }
                    _ = v41
                }
                var v50 int8 = -114
                _ = fi36
                _ = v50
            }
            for fi51 := 0; fi51 < 14; fi51++ {
                for fi52 := 0; fi52 < 15; fi52++ {
                    fmt.Println(v2, v29)
                    _ = fi52
                }
                fmt.Println(fi51)
                _ = fi51
            }
        }
        v7 += v7
        switch 32767 {
        case -2147483648:
            if true {
                v53 := GenS4{F0: ((!(v2 <= v2)) && true)}
                fmt.Println(v2, v29, v29)
                var v54 uint32 = uint32(v2)
                _ = v53
                _ = v54
            }
        }
    }
    v29 = 9223372036854775807
    var v55 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v7
}
