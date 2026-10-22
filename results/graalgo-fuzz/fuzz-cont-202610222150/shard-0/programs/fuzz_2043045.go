package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 string; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 bool; F2 int64 }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 uint64 }

func main() {
    v1 := []string{("hello" + "foo bar"), "hello"}
    v2 := make(map[uint32]string)
    var v3 string = ""
    var v4 uint = 18446744073709551615
    if (108 != 128) {
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 bool = true
            for fi7 := 0; fi7 < 8; fi7++ {
                v8 := GenS1{F0: 4294967295}
                switch "hello" {
                case "":
                    if (!(false && true)) {
                        v9 := GenS1{F0: 4294967295}
                        _ = v9
                    } else {
                        v4 += v4
                    }
                case "a":
                    if (-2147483648 > 0) {
                        v10 := GenS1{F0: 42}
                        fmt.Println(v6, fi5, fi5)
                        switch -529 {
                        case 128:
                            v2[4294967295] = "foo bar"
                        case 2147483647:
                            v3 = ((((v3 + ("" + "a")) + "") + "hello") + ((v3 + v3) + v3))
                        }
                        _ = v10
                    } else {
                        v11 := GenS2{F0: -242, F1: ("a" + (v3 + (v3 + v3))), F2: 0}
                        _ = v11
                    }
                }
                v12 := GenS3{F0: 4294967295}
                _ = fi7
                _ = v12
                _ = v8
            }
            var v13 int64 = -921
            _ = fi5
            _ = v13
            _ = v6
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            for fi15 := 0; fi15 < 3; fi15++ {
                _ = len(v1)
                v16 := GenS3{F0: 0}
                var v17 int32 = -1
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi14
        }
    }
    if (2147483647 != 2147483647) {
        if false {
            v18 := GenS4{F0: 416}
            _ = v18
        } else {
            if false {
                for fi19 := 0; fi19 < 11; fi19++ {
                    for fi20 := 0; fi20 < 7; fi20++ {
                        if (0 != 92) {
                            switch (v3 + "") {
                            case "":
                                var v21 uint32 = 188
                                _ = v21
                            default:
                                v3 = (v3 + (v3 + "hello"))
                            }
                        }
                        _ = fi20
                    }
                    switch -1 {
                    case 127:
                        v3 = (v3 + v3)
                    case 1:
                        v4 += v4
                    }
                    if true {
                        var v22 float32 = 0.1
                        _ = v22
                    } else {
                        v23 := GenS5{F0: 212, F1: false, F2: -361}
                        var v24 uint8 = 255
                        if (255 < 2) {
                            v1 = append(v1, (("a" + v3) + (v3 + v3)))
                        }
                        _ = v23
                        _ = v24
                    }
                    _ = fi19
                }
                delete(v2, 1)
            }
        }
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        v25 := GenS4{F0: 4}
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        v4 = (v4 + v4)
        var v27 string = v3
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        v3 = v3
        v3 = ("a" + v3)
        _ = fi28
    }
    fmt.Println(v4)
    v3 = ((v3 + (v3 + "a")) + ("a" + v3))
    if (668 > -469) {
        v3 = v3
        if ((("" + (v3 + v3)) <= (v3 + v3)) || true) {
            var v29 int = -779
            var v30 bool = true
            v29 -= (v29 | (v29 ^ v29))
            _ = v29
            _ = v30
        } else {
            switch -1 {
            case 32768:
                for fi31 := 0; fi31 < 1; fi31++ {
                    v4 = (v4 * 65535)
                    _ = fi31
                }
            case 91:
                var v32 int64 = 42
                _ = v32
            case -1:
                v33 := GenS6{F0: -661, F1: 4294967295}
                _ = v33
            default:
                v3 = v3
            }
        }
    }
    v3 = v3
    var v34 uint16 = uint16(v4)
    v1 = append(v1, v3)
    _ = len(v2)
    v35 := v2[42]
    if ((2147483647 >= 9223372036854775807) && ((1 > v4) && (-1 >= -22))) {
        fmt.Println(v3, v35, v3)
        v36 := GenS7{F0: 42}
        _ = v36
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        var v38 string = ("foo bar" + "\n\t")
        v38 = ((v38 + "foo bar") + "\n\t")
        _ = fi37
        _ = v38
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
