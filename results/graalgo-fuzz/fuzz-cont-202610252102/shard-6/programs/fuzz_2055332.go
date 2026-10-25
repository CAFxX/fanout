package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 bool; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 uint; F2 uint16 }
type GenS6 struct { F0 uint16; F1 string; F2 int8 }
type GenS7 struct { F0 int8; F1 int16; F2 uint64 }
type GenS8 struct { F0 int32 }

func main() {
    var v1 string = ("hello" + "\n\t")
    v2 := make(map[string]uint32)
    var v3 float32 = float32((3.14159 + float64(-2.25)))
    var v4 int16 = -1
    var v5 int = 32768
    var v6 uint8 = 0
    v3 = -2.25
    if false {
        v7 := GenS1{F0: 0}
        _ = v7
    }
    if true {
        var v8 uint32 = 2147483647
        _ = v8
    }
    if (4294967295 > (-(v5))) {
        if ((904 == 65535) && (false && false)) {
            switch "" {
            case "foo bar":
                var v9 uint64 = 922
                _ = v9
            default:
                v4 = (v4 - v4)
            }
        } else {
            v10 := GenS2{F0: 537, F1: true, F2: 42}
            _ = v10
        }
        v11 := GenS3{F0: uint(v6)}
        for fi12 := 0; fi12 < 18; fi12++ {
            if true {
                if (false || (v6 < 169)) {
                    v6 = v6
                } else {
                    v4 -= 1
                    fmt.Println(v3)
                }
                fmt.Println(v3)
            } else {
                var v13 float32 = 1e10
                if ((!(!false)) && true) {
                    if (-1 <= 119) {
                        v4 = v4
                        fmt.Println(v13, v4)
                        if (false && (true || (1e10 != 1e10))) {
                            if true {
                                var v14 uint8 = v6
                                delete(v2, v1)
                                _ = v14
                            }
                            switch v1 {
                            case "\n\t":
                                for fi15 := 0; fi15 < 16; fi15++ {
                                    v16 := GenS4{F0: 32767}
                                    v5 -= (fi12 + fi12)
                                    var v17 uint16 = 130
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                }
                            case "foo bar":
                                switch -295 {
                                case 128:
                                    var v18 uint64 = 711
                                    _ = v18
                                case 4294967295:
                                    v6 = 42
                                case 127:
                                    if true {
                                        v4 = (v4 >> ((870) & 63))
                                        var v19 string = "foo bar"
                                        _ = v19
                                    } else {
                                        for fi20 := 0; fi20 < 11; fi20++ {
                                            var v21 uint8 = v6
                                            v1 = "hello"
                                            v22 := GenS5{F0: 42, F1: 42, F2: 2}
                                            if (!true) {
                                                var v23 string = ("a" + v1)
                                                var v24 uint64 = 589
                                                var v25 uint8 = v6
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            } else {
                                                var v26 int = (v5 << ((692) & 63))
                                                v27 := []int{}
                                                _ = v26
                                                _ = v27
                                            }
                                            _ = fi20
                                            _ = v21
                                            _ = v22
                                        }
                                        for fi28 := 0; fi28 < 1; fi28++ {
                                            v6 += v6
                                            _ = fi28
                                        }
                                        v29 := GenS6{F0: 32767, F1: (("a" + v1) + v1), F2: -27}
                                        _ = v29
                                    }
                                }
                            default:
                                v30 := []int{}
                                _ = v30
                            }
                            var v31 uint32 = 42
                            _ = v31
                        } else {
                            switch (v1 + v1) {
                            case "\n\t":
                                switch v5 {
                                case 42:
                                    v4 = (v4 % (((^(v4))) | 1))
                                }
                            }
                        }
                    }
                }
                v32 := GenS7{F0: 0, F1: v4, F2: 2}
                _ = v13
                _ = v32
            }
            for fi33 := 0; fi33 < 19; fi33++ {
                var v34 uint64 = 18446744073709551615
                if false {
                    v35 := []int{}
                    v36 := GenS1{F0: 0}
                    v37 := GenS8{F0: -692}
                    _ = v35
                    _ = v36
                    _ = v37
                }
                v34 += (^(v34))
                _ = fi33
                _ = v34
            }
            v6 = 1
            _ = fi12
        }
        _ = v11
    }
    for fi38 := 0; fi38 < 9; fi38++ {
        var v39 uint = 763
        v40 := []int{v5, (fi38 & 128), (v5 << ((65535) & 63))}
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 int16 = v4
    v41 = v41
    var v42 uint = 460
    v41 -= v4
    var v43 int16 = (v41 >> ((v42) & 63))
    v41 = v43
    fmt.Println(v5, v1, v1)
    var v44 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
