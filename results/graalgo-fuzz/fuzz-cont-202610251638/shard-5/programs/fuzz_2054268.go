package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 bool; F2 int }
func genFn0(p0 string) int16 {
    var v6 uint64 = 565
    var v7 uint32 = 871
    switch (("" + p0) + "\n\t") {
    case "\n\t":
        fmt.Println(v6, p0, v7)
    case "foo bar":
        v7 += v7
    default:
        switch p0 {
        case "\n\t":
            if true {
                v8 := GenS1{F0: 127, F1: (v7 & 2147483647), F2: (v6 % (((v6 + v6)) | 1))}
                var v9 string = p0
                _ = v8
                _ = v9
            } else {
                v10 := GenS2{F0: (-9223372036854775808 <= 274), F1: v7, F2: 571}
                _ = v10
            }
        case "a":
            switch 128 {
            case -1:
                fmt.Println(p0)
            }
        case "foo bar":
            switch 0 {
            case 605:
                p0 = p0
            }
        default:
            var v11 int64 = 0
            _ = v11
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return -1
}
func genFn1(p0 bool, p1 string) uint8 {
    v12 := GenS2{F0: false, F1: 228, F2: 0}
    switch -338 {
    case 128:
        p1 = (p1 + (p1 + (p1 + (p1 + ((p1 + p1) + "foo bar")))))
    case 0:
        var v13 uint8 = 91
        _ = v13
    case 42:
        fmt.Println(p1)
    default:
        v14 := GenS3{F0: 2, F1: 2}
        _ = v14
    }
    p0 = (-1 <= -1)
    if p0 {
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 float64 = 1e10
            v17 := GenS1{F0: 1, F1: 1, F2: 42}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = v12
    return 1
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{-2147483648, 2147483647, 42}
    v3 := make(map[uint32]string)
    var v4 bool = (-2147483648 < 42)
    var v5 string = ""
    var v18 string = v5
    var v19 uint8 = 255
    v5 = (v5 + "a")
    if v4 {
        v4 = (!v4)
    } else {
        var v20 string = ("foo bar" + v5)
        v2 = append(v2, -288)
        _ = v20
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5 = ""
    v21 := GenS4{F0: 2}
    v4 = v4
    v5 = "\n\t"
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v22 int64 = 9223372036854775807
    v23 := GenS5{F0: 0, F1: (((v4 && v4) || (((false && v4) && (1 <= 20)) || (!v4))) || true), F2: 32768}
    if ((0.0 / 0.1) < (1.5 + 1e10)) {
        if (!(0.1 >= 0.1)) {
            if (!v4) {
                if v4 {
                    for fi24 := 0; fi24 < 8; fi24++ {
                        for fi25 := 0; fi25 < 7; fi25++ {
                            switch (fi24 % ((32768) | 1)) {
                            case 969:
                                var v26 uint8 = v19
                                _ = v26
                            default:
                                fmt.Println(v19, v5, v18)
                            }
                            if v4 {
                                var v27 float32 = 1e10
                                fmt.Println(v4)
                                switch "\n\t" {
                                case "hello":
                                    switch v18 {
                                    case "a":
                                        v28 := GenS5{F0: 18446744073709551615, F1: (2147483648 > 2147483648), F2: fi25}
                                        _ = v28
                                    case "foo bar":
                                        if true {
                                            var v29 uint64 = 2
                                            _ = v29
                                        }
                                    }
                                case "\n\t":
                                    v18 = v18
                                    v2 = append(v2, (fi24 + (fi25 / ((fi24) | 1))))
                                    _ = len(v2)
                                }
                                _ = v27
                            }
                            v2 = append(v2, (^(fi24)))
                            _ = fi25
                        }
                        v30 := GenS3{F0: 42, F1: 0}
                        _ = fi24
                        _ = v30
                    }
                }
                fmt.Println(v19)
                v4 = v4
            }
            if (-13 == -61) {
                var v31 int16 = -363
                switch "" {
                case "hello":
                    fmt.Println(v5)
                case "a":
                    for fi32 := 0; fi32 < 2; fi32++ {
                        var v33 int64 = v22
                        var v34 uint16 = 880
                        for fi35 := 0; fi35 < 18; fi35++ {
                            v33 = int64(fi32)
                            var v36 float32 = 0.0
                            _ = fi35
                            _ = v36
                        }
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                case "\n\t":
                    fmt.Println(v5, v4)
                }
                _ = v31
            }
            var v37 int = -2147483648
            _ = v37
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch (("a" + v18) + v18) {
        case "foo bar":
            switch v18 {
            case "":
                var v38 bool = true
                _ = v38
            case "foo bar":
                var v39 float32 = -2.25
                _ = v39
            case "hello":
                var v40 int8 = -128
                _ = v40
            }
        }
    }
    if (!((!(v18 < v18)) && v4)) {
        v2 = append(v2, 2147483647)
        switch -2147483648 {
        case 4294967295:
            var v41 uint8 = (v19 >> ((18446744073709551615) & 63))
            _ = v41
        }
    }
    if true {
        v5 = v5
        var v42 uint64 = 18446744073709551615
        v5 = ""
        _ = v42
    } else {
        var v43 int16 = genFn0(v18)
        fmt.Println(v4, v18)
        var v44 float32 = 1.5
        _ = v43
        _ = v44
    }
    switch 805 {
    case -1:
        switch v5 {
        case "":
            v19 = 2
        case "a":
            for fi45 := 0; fi45 < 16; fi45++ {
                for fi46 := 0; fi46 < 19; fi46++ {
                    if len(v2) > 0 {
                        v2[0] = fi46
                    }
                    var v47 string = v5
                    _ = fi46
                    _ = v47
                }
                v5 = (v5 + v5)
                delete(v3, 4294967295)
                _ = fi45
            }
        }
    }
    v4 = (!v4)
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
