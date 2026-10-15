package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64; F1 int32; F2 string }

func main() {
    var v1 uint8 = 0
    var v2 int32 = 2147483647
    v3 := make(map[uint32]bool)
    v4 := []string{"foo bar", "", (("foo bar" + "foo bar") + ("" + "a")), "foo bar"}
    if false {
        for fi5 := 0; fi5 < 2; fi5++ {
            v1 -= v1
            if (718 == 364) {
                if (32767 <= -1) {
                    for fi6 := 0; fi6 < 18; fi6++ {
                        var v7 bool = false
                        var v8 uint8 = (v1 + v1)
                        _ = fi6
                        _ = v7
                        _ = v8
                    }
                    fmt.Println(v2, fi5, v1)
                    var v9 uint = 18446744073709551615
                    _ = v9
                } else {
                    fmt.Println(v1)
                    v10 := GenS1{F0: 127}
                    _ = v10
                }
                var v11 uint = 4294967295
                switch ("a" + "hello") {
                case "foo bar":
                    _ = len(v4)
                case "hello":
                    var v12 uint = v11
                    _ = v12
                }
                _ = v11
            } else {
                var v13 int32 = v2
                var v14 uint = 896
                v3[42] = false
                _ = v13
                _ = v14
            }
            var v15 uint64 = 1
            _ = fi5
            _ = v15
        }
        switch -2147483648 {
        case -9223372036854775808:
            var v16 int = -2147483648
            _ = v16
        case 32767:
            for fi17 := 0; fi17 < 10; fi17++ {
                for fi18 := 0; fi18 < 20; fi18++ {
                    v19 := GenS2{F0: "foo bar"}
                    _ = fi18
                    _ = v19
                }
                if true {
                    for fi20 := 0; fi20 < 8; fi20++ {
                        var v21 int8 = 127
                        _ = fi20
                        _ = v21
                    }
                    fmt.Println(v1, fi17, fi17)
                    var v22 uint32 = 4294967295
                    _ = v22
                }
                for fi23 := 0; fi23 < 4; fi23++ {
                    var v24 uint8 = 226
                    v1 = v1
                    if true {
                        fmt.Println(v1, v1, fi17)
                    } else {
                        var v25 int = 127
                        if (!(0 < -1)) {
                            _ = len(v3)
                            fmt.Println(fi17, v25)
                        }
                        _ = v25
                    }
                    _ = fi23
                    _ = v24
                }
                _ = fi17
            }
        case 1:
            if true {
                v1 -= (v1 + (v1 | v1))
                switch "\n\t" {
                case "hello":
                    fmt.Println(v1, v1)
                case "":
                    v1 = v1
                case "\n\t":
                    if false {
                        var v26 string = ("a" + ("" + "foo bar"))
                        _ = v26
                    }
                }
            } else {
                var v27 bool = true
                _ = v27
            }
        }
        v1 += (v1 / (((v1 + v1)) | 1))
    }
    v1 -= v1
    switch 9223372036854775807 {
    case 42:
        for fi28 := 0; fi28 < 4; fi28++ {
            v1 -= v1
            v2 += 1
            v1 -= 2
            _ = fi28
        }
    }
    v29 := GenS3{F0: -1, F1: 0}
    v30 := GenS1{F0: v1}
    v31 := GenS2{F0: ("" + "a")}
    v2 = int32(v1)
    switch 0 {
    case 32768:
        fmt.Println(v1)
    case 127:
        if true {
            var v32 string = ("" + "")
            switch 0 {
            case 32767:
                if ((4294967295 == 2) || (true && true)) {
                    var v33 uint64 = 2
                    _ = v33
                }
            default:
                v4 = append(v4, "hello")
            }
            var v34 uint16 = 255
            _ = v32
            _ = v34
        } else {
            var v35 uint8 = v1
            v2 = -984
            if (9223372036854775807 > -454) {
                if len(v4) > 0 {
                    v4[0] = ""
                }
            }
            _ = v35
        }
    }
    fmt.Println(v1)
    v36 := GenS4{F0: 42}
    var v37 uint32 = 4294967295
    v38 := GenS1{F0: (v1 ^ (v1 << ((0) & 63)))}
    for fi39 := 0; fi39 < 6; fi39++ {
        v3[v37] = true
        if true {
            var v40 string = "\n\t"
            _ = v40
        } else {
            if ((1.5 / (-2.25 - 3.14159)) < 0.1) {
                v1 = 2
                var v41 uint32 = 2
                _ = v41
            } else {
                v42 := GenS5{F0: 4294967295, F1: (v2 - v2), F2: ("a" + "")}
                fmt.Println(v2, v37)
                var v43 uint = 0
                _ = v42
                _ = v43
            }
            var v44 float64 = (0.1 + -2.25)
            _ = v44
        }
        if ((true && true) && true) {
            v1 += v1
        }
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
