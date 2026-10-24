package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint16; F1 int; F2 uint }
type GenS3 struct { F0 uint64; F1 string }

func main() {
    var v1 float32 = (0.0 / 1.0)
    v2 := make(map[uint32]string)
    var v3 uint64 = 74
    var v4 uint32 = 65535
    var v5 int64 = 42
    var v6 uint16 = 2
    v7 := []int{42, -760}
    v2[v4] = "a"
    v5 = -245
    if (v4 <= v4) {
        var v8 string = ("\n\t" + "\n\t")
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: 18446744073709551615, F1: "foo bar"}
            var v11 uint32 = 0
            _ = fi9
            _ = v10
            _ = v11
        }
        if (!(128 > uint8(v4))) {
            for fi12 := 0; fi12 < 1; fi12++ {
                v1 = v1
                v1 = v1
                _ = fi12
            }
            v8 = "\n\t"
        }
        _ = v8
    } else {
        v3 = (v3 / ((v3) | 1))
        var v13 int = 71
        _ = v13
    }
    var v14 bool = true
    if (0.0 == 0.1) {
        v5 += (-(v5))
    }
    v6 += (^(v6))
    switch "foo bar" {
    case "hello":
        v14 = ((v4 % ((v4) | 1)) > v4)
    case "a":
        if (v4 <= (v4 | v4)) {
            var v15 int8 = 0
            for fi16 := 0; fi16 < 20; fi16++ {
                if len(v7) > 0 {
                    v7[0] = 1
                }
                var v17 float32 = 1.5
                fmt.Println(v17)
                _ = fi16
                _ = v17
            }
            switch "hello" {
            case "hello":
                v18 := GenS2{F0: (v6 ^ 42), F1: -576, F2: 4294967295}
                _ = v18
            case "\n\t":
                if (255 == (v6 ^ v6)) {
                    fmt.Println(v5, v5)
                    var v19 string = ""
                    var v20 float64 = 0.1
                    _ = v19
                    _ = v20
                }
            }
            _ = v15
        } else {
            switch -1 {
            case 9223372036854775807:
                var v21 uint8 = 127
                _ = v21
            case 127:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            default:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            var v22 uint64 = v3
            if (-113 > 1) {
                switch ("foo bar" + "") {
                case "\n\t":
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                default:
                    v5 += (v5 ^ (v5 | v5))
                }
            }
            _ = v22
        }
    case "foo bar":
        v23 := GenS2{F0: 65535, F1: -75, F2: 2}
        _ = v23
    default:
        for fi24 := 0; fi24 < 20; fi24++ {
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 int32 = int32(v5)
                if v14 {
                    switch (fi24 | fi25) {
                    case 42:
                        fmt.Println(v3, v14)
                    }
                    v14 = (v14 || (!(2147483647 >= (fi24 >> ((244) & 63)))))
                    var v27 float32 = 1e10
                    _ = v27
                }
                v28 := GenS1{F0: 803, F1: "a"}
                fmt.Println(v14, v1)
                fmt.Println(fi25, v5)
                _ = fi25
                _ = v26
                _ = v28
            }
            v29 := GenS1{F0: 1, F1: ""}
            v6 = (v6 >> ((18446744073709551615) & 63))
            _ = fi24
            _ = v29
        }
    }
    var v30 int64 = v5
    var v31 float64 = float64(v1)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    var v32 int = -223
    for fi33 := 0; fi33 < 11; fi33++ {
        for fi34 := 0; fi34 < 13; fi34++ {
            fmt.Println(fi33)
            fmt.Println(fi33, v14, fi34)
            if (v3 != 1) {
                if v14 {
                    switch v32 {
                    case 9223372036854775807:
                        var v35 int64 = -1
                        _ = v35
                    case 4294967295:
                        var v36 bool = false
                        _ = v36
                    case 42:
                        fmt.Println(v30, v1, v4)
                    }
                    switch "foo bar" {
                    case "":
                        switch "\n\t" {
                        case "hello":
                            fmt.Println(v5, v31, v30)
                        }
                    case "foo bar":
                        if (!(!v14)) {
                            var v37 uint64 = v3
                            _ = len(v7)
                            _ = v37
                        } else {
                            v38 := GenS1{F0: 0, F1: ""}
                            v39 := GenS3{F0: (v3 + 0), F1: ("foo bar" + "hello")}
                            _ = v38
                            _ = v39
                        }
                    case "\n\t":
                        if false {
                            var v40 uint64 = 18446744073709551615
                            var v41 uint = 1
                            switch "\n\t" {
                            case "":
                                v42 := v2[v4]
                                _ = v42
                            case "a":
                                var v43 float32 = (0.0 / 3.14159)
                                _ = v43
                            default:
                                var v44 uint8 = 143
                                _ = v44
                            }
                            if ((v41 / ((v41) | 1)) <= v41) {
                                v14 = v14
                            } else {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                            _ = v40
                            _ = v41
                        }
                    default:
                        var v45 uint = 0
                        _ = v45
                    }
                    var v46 uint16 = (^(v6))
                    v14 = false
                    _ = v46
                } else {
                    var v47 uint64 = uint64(v6)
                    _ = v47
                }
                v48 := GenS3{F0: v3, F1: "hello"}
                v5 = v5
                _ = v48
            }
            _ = fi34
        }
        if v14 {
            fmt.Println(v5)
        }
        for fi49 := 0; fi49 < 3; fi49++ {
            var v50 int16 = 1000
            var v51 int32 = 713
            _ = fi49
            _ = v50
            _ = v51
        }
        _ = fi33
    }
    if true {
        v3 = v3
        for fi52 := 0; fi52 < 7; fi52++ {
            switch (fi52 - v32) {
            case 257:
                for fi53 := 0; fi53 < 6; fi53++ {
                    switch v32 {
                    case 1:
                        v54 := GenS1{F0: 0, F1: "\n\t"}
                        _ = v54
                    case 899:
                        v55 := GenS3{F0: (v3 - (^(v3))), F1: ""}
                        _ = v55
                    case -1:
                        v56 := GenS1{F0: 42, F1: "hello"}
                        _ = v56
                    }
                    _ = fi53
                }
            }
            v3 -= 442
            _ = fi52
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
