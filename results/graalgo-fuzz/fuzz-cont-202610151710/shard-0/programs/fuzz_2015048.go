package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 uint; F1 uint; F2 uint }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 int = 283
    v2 := make(map[uint32]uint32)
    var v3 bool = (1 != -89)
    v4 := make(map[uint32]int)
    var v5 int32 = 42
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 18; fi7++ {
        if v3 {
            v5 = v5
            delete(v2, 0)
        } else {
            fmt.Println(v1, v3, v5)
            fmt.Println(fi7, v5, v5)
            v1 -= (fi7 % ((v1) | 1))
        }
        v8 := GenS1{F0: v3, F1: 2147483647}
        var v9 float32 = -2.25
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v5, v5)
    var v10 int32 = (v5 | 0)
    switch 32768 {
    case -735:
        for fi11 := 0; fi11 < 15; fi11++ {
            fmt.Println(v1, v3)
            var v12 uint8 = 128
            delete(v2, 2147483647)
            _ = fi11
            _ = v12
        }
    case 32768:
        if (!((2147483647 <= 356) && ((v1 / (((v1 | v1)) | 1)) > (v1 - v1)))) {
            var v13 int32 = v5
            switch ("hello" + "a") {
            case "foo bar":
                var v14 int32 = -1
                _ = v14
            }
            v1 = 0
            _ = v13
        }
    }
    switch "a" {
    case "\n\t":
        var v15 uint64 = 42
        _ = v15
    default:
        switch (("foo bar" + ("foo bar" + "hello")) + ("foo bar" + "foo bar")) {
        case "hello":
            v16 := GenS2{F0: 2, F1: 42}
            _ = v16
        case "a":
            var v17 int32 = -36
            _ = v17
        }
    }
    switch "hello" {
    case "hello":
        v18 := GenS1{F0: ((v3 || (1.5 >= (0.1 / 0.1))) && v3), F1: 635}
        _ = v18
    case "":
        var v19 uint8 = 173
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v1 = fi20
        for fi21 := 0; fi21 < 3; fi21++ {
            switch (v1 ^ fi21) {
            case -2147483648:
                v22 := GenS3{F0: 241, F1: 0}
                _ = v22
            case 9223372036854775807:
                fmt.Println(fi21, v5)
            default:
                switch ("foo bar" + ("" + ("" + "\n\t"))) {
                case "a":
                    switch "foo bar" {
                    case "a":
                        v1 = (-(fi21))
                    }
                }
            }
            _ = fi21
        }
        _ = fi20
    }
    v1 -= 32767
    fmt.Println(v3)
    var v23 int32 = 2147483647
    if v3 {
        v24 := v2[2]
        if false {
            v1 -= (v1 - 4294967295)
        } else {
            switch v1 {
            case 1:
                fmt.Println(v24)
            case 2147483647:
                v24 = v24
            }
        }
        _ = v24
    }
    if v3 {
        if (!(!true)) {
            if false {
                for fi25 := 0; fi25 < 19; fi25++ {
                    for fi26 := 0; fi26 < 15; fi26++ {
                        v27 := GenS1{F0: ("hello" > ("foo bar" + (("" + "hello") + "a"))), F1: 4294967295}
                        if (!((false || v3) || ((!true) && true))) {
                            switch fi25 {
                            case 4294967295:
                                v10 = -1
                            case 32767:
                                _ = len(v4)
                            default:
                                v28 := GenS2{F0: 844, F1: -9223372036854775808}
                                _ = v28
                            }
                        }
                        _ = fi26
                        _ = v27
                    }
                    var v29 int64 = 1
                    var v30 uint64 = 18446744073709551615
                    _ = fi25
                    _ = v29
                    _ = v30
                }
                switch "a" {
                case "":
                    if v3 {
                        var v31 bool = v3
                        _ = v31
                    }
                case "a":
                    for fi32 := 0; fi32 < 4; fi32++ {
                        var v33 int = fi32
                        _ = len(v2)
                        for fi34 := 0; fi34 < 10; fi34++ {
                            v35 := []int{}
                            _ = fi34
                            _ = v35
                        }
                        _ = fi32
                        _ = v33
                    }
                case "hello":
                    if v3 {
                        if v3 {
                            var v36 uint32 = 2147483647
                            _ = v36
                        }
                        var v37 int = v1
                        v5 = v23
                        _ = v37
                    }
                }
                if v3 {
                    switch 2147483647 {
                    case -722:
                        var v38 float32 = (-2.25 * 1e10)
                        _ = v38
                    case -668:
                        var v39 bool = (-128 < 100)
                        _ = v39
                    case 1:
                        switch "\n\t" {
                        case "a":
                            v40 := GenS4{F0: 2, F1: 65535, F2: 42}
                            _ = v40
                        case "\n\t":
                            var v41 uint16 = 1
                            _ = v41
                        default:
                            if false {
                                v1 = 4294967295
                                v23 = v5
                                var v42 int8 = -128
                                _ = v42
                            }
                        }
                    }
                    v43 := GenS3{F0: 42, F1: 907}
                    v44 := []int{v1, -185, 1}
                    _ = v43
                    _ = v44
                }
            } else {
                var v45 bool = v3
                v46 := GenS5{F0: 1}
                v5 = -1
                _ = v45
                _ = v46
            }
        }
        if false {
            v10 += v23
            _ = len(v2)
        } else {
            for fi47 := 0; fi47 < 6; fi47++ {
                v48 := GenS5{F0: 4294967295}
                fmt.Println(v3, v10, v3)
                _ = fi47
                _ = v48
            }
            if true {
                var v49 bool = (!v3)
                var v50 uint64 = 4294967295
                _ = v49
                _ = v50
            }
        }
        for fi51 := 0; fi51 < 12; fi51++ {
            fmt.Println(fi51, v10)
            var v52 bool = v3
            var v53 int32 = v5
            _ = fi51
            _ = v52
            _ = v53
        }
    }
    v54 := GenS2{F0: 1, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
