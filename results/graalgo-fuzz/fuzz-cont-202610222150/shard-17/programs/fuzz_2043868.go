package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 uint, p1 int, p2 int16) int8 {
    v6 := []int{}
    v7 := GenS1{F0: p0, F1: 1, F2: 0}
    if (42 < 0) {
        var v8 uint32 = 2
        if (-1 != -9223372036854775808) {
            if true {
                var v9 uint64 = 926
                _ = v9
            }
            p1 = (p1 | p1)
        } else {
            switch "hello" {
            case "hello":
                var v10 string = (("foo bar" + "hello") + "a")
                _ = v10
            case "a":
                for fi11 := 0; fi11 < 1; fi11++ {
                    fmt.Println(fi11, p2)
                    for fi12 := 0; fi12 < 7; fi12++ {
                        var v13 float32 = 0.0
                        p1 = -191
                        switch (-(p1)) {
                        case -521:
                            switch "foo bar" {
                            case "":
                                if false {
                                    if true {
                                        v14 := GenS2{F0: ("\n\t" + "hello"), F1: (v8 ^ 0)}
                                        var v15 uint64 = 42
                                        var v16 int16 = 32767
                                        _ = v14
                                        _ = v15
                                        _ = v16
                                    }
                                }
                            default:
                                var v17 uint16 = uint16(v8)
                                _ = v17
                            }
                        default:
                            switch fi12 {
                            case 4294967295:
                                if (false || (!true)) {
                                    v8 -= v8
                                }
                            case -740:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            case 128:
                                var v18 uint16 = 255
                                if true {
                                    p0 -= 1
                                    if true {
                                        v19 := GenS3{F0: 128}
                                        _ = v19
                                    }
                                    v20 := GenS3{F0: 128}
                                    var v21 bool = true
                                    switch "a" {
                                    case "a":
                                        var v22 uint8 = 15
                                        _ = v22
                                    default:
                                        v23 := GenS1{F0: p0, F1: 988, F2: 18446744073709551615}
                                        _ = v23
                                    }
                                    _ = v20
                                    _ = v21
                                }
                                switch 0 {
                                case 127:
                                    v13 = float32(3.14159)
                                default:
                                    fmt.Println(p1)
                                }
                                _ = v18
                            }
                        }
                        _ = fi12
                        _ = v13
                    }
                    _ = fi11
                }
            default:
                if (!((42 <= 42) && (81 >= 127))) {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = len(v6)
                }
            }
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 127
}
func genFn1() string {
    var v24 uint64 = 0
    for fi25 := 0; fi25 < 9; fi25++ {
        v26 := []int{-9223372036854775808}
        v24 = (v24 % (((v24 + v24)) | 1))
        switch (fi25 | (fi25 % ((fi25) | 1))) {
        case -9223372036854775808:
            v24 -= v24
        case 32767:
            var v27 uint8 = 2
            _ = v27
        }
        _ = fi25
        _ = v26
    }
    if ((188 == 2) && true) {
        v24 = (^(v24))
        switch "hello" {
        case "foo bar":
            v28 := GenS3{F0: 85}
            _ = v28
        case "\n\t":
            var v29 int16 = -318
            _ = v29
        case "hello":
            var v30 int = 602
            _ = v30
        }
        var v31 bool = (!(v24 != 2))
        _ = v31
    } else {
        for fi32 := 0; fi32 < 20; fi32++ {
            v33 := []int{fi32, (-(fi32))}
            _ = len(v33)
            _ = fi32
            _ = v33
        }
    }
    v34 := GenS3{F0: 1}
    fmt.Println(v24)
    _ = v24
    _ = v34
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]uint32)
    var v3 int64 = 238
    var v4 float32 = (0.0 / (1.5 + 1e10))
    var v5 uint16 = 0
    fmt.Println(v4, v5, v4)
    v5 = 1
    var v35 int = -433
    if true {
        v4 = (-2.25 + 0.1)
    } else {
        if false {
            v35 -= v35
        } else {
            v36 := []int{}
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
            _ = v36
        }
    }
    var v37 uint32 = 0
    for fi38 := 0; fi38 < 4; fi38++ {
        v35 = fi38
        var v39 int64 = v3
        fmt.Println(v4, v5, fi38)
        _ = fi38
        _ = v39
    }
    var v40 int64 = v3
    for fi41 := 0; fi41 < 17; fi41++ {
        for fi42 := 0; fi42 < 8; fi42++ {
            v35 = fi41
            v37 = v37
            var v43 int = fi42
            _ = fi42
            _ = v43
        }
        if true {
            v37 = v37
        }
        fmt.Println(v37)
        _ = fi41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v40
    _ = v5
}
