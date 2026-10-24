package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 bool; F1 int64 }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]string)
    v4 := []uint32{}
    var v5 uint64 = 1
    var v6 uint16 = 32767
    v6 = 580
    var v7 uint = 65535
    switch (("hello" + v1) + ("hello" + v1)) {
    case "\n\t":
        var v8 float32 = 1.5
        _ = v8
    case "hello":
        _ = len(v2)
    case "foo bar":
        fmt.Println(v1, v1)
    default:
        v9 := GenS1{F0: ("" + "foo bar")}
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v2[0] = v1
        v5 += 2
        if (!true) {
            if (false && true) {
                var v11 bool = false
                _ = v11
            } else {
                switch v1 {
                case "foo bar":
                    v5 = (^(v5))
                default:
                    var v12 int8 = -1
                    _ = v12
                }
                var v13 bool = false
                var v14 uint64 = 18446744073709551615
                _ = v13
                _ = v14
            }
            var v15 int = 1
            _ = v15
        }
        _ = fi10
    }
    if false {
        switch (("hello" + (v1 + "\n\t")) + "foo bar") {
        case "a":
            if (255 <= 42) {
                var v16 bool = true
                v17 := GenS2{F0: v6}
                _ = v16
                _ = v17
            } else {
                for fi18 := 0; fi18 < 7; fi18++ {
                    fmt.Println(v6)
                    _ = fi18
                }
                v19 := GenS3{F0: 1000, F1: v5}
                _ = v19
            }
        case "foo bar":
            switch v1 {
            case "":
                for fi20 := 0; fi20 < 5; fi20++ {
                    var v21 int = (^(fi20))
                    _ = fi20
                    _ = v21
                }
            }
        case "hello":
            for fi22 := 0; fi22 < 15; fi22++ {
                v23 := GenS1{F0: (v1 + (((v1 + ("foo bar" + (v1 + v1))) + v1) + "hello"))}
                var v24 uint64 = v5
                _ = fi22
                _ = v23
                _ = v24
            }
        }
    }
    v1 = "hello"
    if (4294967295 != 42) {
        var v25 float32 = 3.14159
        _ = v25
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        if false {
            var v27 uint8 = 59
            if false {
                if false {
                    var v28 float32 = ((-2.25 * (1e10 * (1e10 - float32(1e10)))) + 0.1)
                    _ = v28
                } else {
                    if false {
                        delete(v2, 65535)
                        if true {
                            _ = len(v4)
                        }
                    } else {
                        var v29 bool = true
                        var v30 int64 = -905
                        _ = v29
                        _ = v30
                    }
                    var v31 uint = 4294967295
                    _ = v31
                }
                switch ((("a" + "a") + "a") + v1) {
                case "hello":
                    if false {
                        v32 := GenS2{F0: v6}
                        _ = v32
                    } else {
                        v5 = 18446744073709551615
                        v1 = "hello"
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                }
            } else {
                if (false && true) {
                    v1 = (v1 + "")
                    _ = len(v4)
                    var v33 uint64 = 42
                    _ = v33
                }
                v6 = v6
                var v34 int16 = 0
                _ = v34
            }
            _ = v27
        }
        fmt.Println(v6, v7)
        _ = fi26
    }
    v35 := GenS2{F0: (v6 << ((18446744073709551615) & 63))}
    v36 := GenS1{F0: v1}
    v5 = v5
    if true {
        for fi37 := 0; fi37 < 2; fi37++ {
            var v38 uint8 = 128
            fmt.Println(fi37, v5)
            if true {
                var v39 uint64 = (v5 ^ 42)
                if true {
                    v40 := GenS3{F0: 32767, F1: v5}
                    _ = v40
                }
                _ = v39
            } else {
                v41 := GenS4{F0: true, F1: 1}
                v1 = v1
                _ = v41
            }
            _ = fi37
            _ = v38
        }
    } else {
        v42 := v2[1]
        _ = v42
    }
    var v43 uint64 = 42
    for fi44 := 0; fi44 < 5; fi44++ {
        v45 := v2[65535]
        var v46 float32 = 3.14159
        _ = fi44
        _ = v45
        _ = v46
    }
    if (v1 >= ((v1 + v1) + (v1 + ("foo bar" + ("" + v1))))) {
        if true {
            var v47 float32 = float32(0.0)
            for fi48 := 0; fi48 < 12; fi48++ {
                switch (fi48 - int(v7)) {
                case 2147483647:
                    v47 = v47
                case -432:
                    var v49 int = fi48
                    _ = v49
                case -9223372036854775808:
                    v1 = v1
                default:
                    v5 = v43
                }
                if (!(v43 == (v5 % ((v43) | 1)))) {
                    if false {
                        var v50 uint16 = 42
                        var v51 uint = (v7 & 489)
                        _ = v50
                        _ = v51
                    } else {
                        v47 += -2.25
                        var v52 uint = 255
                        _ = v52
                    }
                } else {
                    fmt.Println(v7, v7, v43)
                }
                _ = fi48
            }
            for fi53 := 0; fi53 < 20; fi53++ {
                var v54 int16 = 779
                v55 := GenS3{F0: (v54 + 0), F1: v5}
                var v56 bool = (-2.25 > (v47 / 3.14159))
                _ = fi53
                _ = v54
                _ = v55
                _ = v56
            }
            _ = v47
        }
        v5 -= v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
