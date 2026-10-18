package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int64 }
func genFn0(p0 uint, p1 uint16, p2 int64) int64 {
    switch -1 {
    case 691:
        for fi5 := 0; fi5 < 17; fi5++ {
            p2 = (p2 + (p2 - (p2 % ((512) | 1))))
            fmt.Println(fi5, p1, fi5)
            p1 = p1
            _ = fi5
        }
    case 794:
        var v6 uint16 = uint16(p0)
        _ = v6
    case 2147483647:
        var v7 uint64 = 1
        _ = v7
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[int]uint32)
    var v3 uint8 = 1
    v4 := make(map[uint32]uint32)
    v3 += uint8(v1)
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: 2147483647}
        v10 := []int{(^(fi8))}
        _ = fi8
        _ = v10
        _ = v9
    }
    if (65535 == 657) {
        fmt.Println(v1)
        switch "hello" {
        case "a":
            for fi11 := 0; fi11 < 6; fi11++ {
                var v12 uint64 = 42
                var v13 uint8 = 128
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        v1 = (^(v1))
    } else {
        v1 += v1
    }
    v14 := GenS1{F0: 623}
    switch ("foo bar" + "foo bar") {
    case "a":
        v3 -= v3
    case "hello":
        v1 += v1
    }
    v15 := GenS1{F0: 2147483648}
    v16 := GenS1{F0: 42}
    var v17 int32 = 0
    var v18 uint64 = 926
    for fi19 := 0; fi19 < 5; fi19++ {
        switch 976 {
        case -1:
            for fi20 := 0; fi20 < 13; fi20++ {
                switch fi20 {
                case 2147483647:
                    if true {
                        v1 += (v1 % ((65535) | 1))
                        if (-924 >= fi19) {
                            var v21 uint16 = (v1 ^ (v1 ^ (^(v1))))
                            _ = v21
                        }
                    } else {
                        v22 := GenS1{F0: 1000}
                        _ = v22
                    }
                }
                v23 := GenS1{F0: 2147483648}
                _ = fi20
                _ = v23
            }
        case 128:
            v17 = -1
        }
        _ = fi19
    }
    v4[0] = 1
    v24 := GenS2{F0: 2}
    for fi25 := 0; fi25 < 4; fi25++ {
        if (("a" < "") || (1 > 429)) {
            switch "hello" {
            case "a":
                switch "" {
                case "\n\t":
                    v26 := GenS2{F0: 42}
                    _ = v26
                case "a":
                    if true {
                        var v27 uint8 = (v3 << ((1) & 63))
                        _ = v27
                    } else {
                        if true {
                            switch (("foo bar" + "\n\t") + "\n\t") {
                            case "\n\t":
                                switch "hello" {
                                case "\n\t":
                                    v17 = (v17 + (v17 % ((1) | 1)))
                                case "hello":
                                    for fi28 := 0; fi28 < 12; fi28++ {
                                        var v29 int64 = genFn0(65535, v1, 0)
                                        _ = fi28
                                        _ = v29
                                    }
                                default:
                                    var v30 uint64 = (v18 | v18)
                                    _ = v30
                                }
                            case "foo bar":
                                if false {
                                    fmt.Println(fi25)
                                } else {
                                    v31 := []int{(fi25 % ((fi25) | 1)), fi25}
                                    _ = v31
                                }
                            default:
                                var v32 uint = 4294967295
                                _ = v32
                            }
                        } else {
                            var v33 string = "foo bar"
                            if (!(!(!(((v17 * 2147483647) < v17) || ((((^(v17)) > (v17 % ((v17) | 1))) || (!false)) || false))))) {
                                var v34 int16 = 42
                                _ = v34
                            } else {
                                var v35 uint64 = v18
                                var v36 string = v33
                                _ = v35
                                _ = v36
                            }
                            if true {
                                v18 += 1
                            } else {
                                var v37 float32 = 0.0
                                _ = len(v4)
                                _ = v37
                            }
                            _ = v33
                        }
                        v38 := GenS1{F0: 596}
                        _ = v38
                    }
                default:
                    var v39 uint16 = (v1 << ((1) & 63))
                    _ = v39
                }
            }
        } else {
            fmt.Println(fi25, v3, v1)
            v3 = 2
            switch ("" + ("foo bar" + ("hello" + "\n\t"))) {
            case "a":
                v17 = -1
            case "":
                delete(v4, 0)
            case "hello":
                if (42 == 2) {
                    var v40 int16 = -826
                    v1 = (v1 << ((uint(v1)) & 63))
                    for fi41 := 0; fi41 < 6; fi41++ {
                        if true {
                            if true {
                                switch "foo bar" {
                                case "a":
                                    v3 -= (v3 - v3)
                                case "":
                                    v1 = 398
                                }
                            }
                        }
                        _ = fi41
                    }
                    _ = v40
                } else {
                    v1 = (v1 >> ((676) & 63))
                }
            default:
                v42 := GenS3{F0: (fi25 - -1), F1: -423}
                _ = v42
            }
        }
        var v43 uint16 = 948
        var v44 uint32 = 2
        _ = fi25
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
