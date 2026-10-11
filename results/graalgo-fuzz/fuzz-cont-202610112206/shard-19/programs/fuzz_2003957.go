package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 int }

func main() {
    v1 := make(map[int]int)
    var v2 int32 = 2147483647
    var v3 uint64 = 42
    var v4 uint8 = 163
    var v5 float64 = 0.1
    v6 := []int{-1}
    v1[32768] = 127
    var v7 int8 = 41
    if (1e10 > 1.5) {
        var v8 int8 = v7
        switch -1 {
        case -1:
            switch -2147483648 {
            case 4294967295:
                var v9 float32 = 1.5
                _ = v9
            case 42:
                v2 = (^(v2))
            case 0:
                var v10 int32 = -269
                _ = v10
            default:
                v1[-340] = 888
            }
        default:
            var v11 uint16 = 32767
            _ = v11
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            switch ("hello" + (("foo bar" + ("foo bar" + "foo bar")) + "hello")) {
            case "foo bar":
                v5 = (v5 * v5)
            case "\n\t":
                v13 := GenS1{F0: (v4 >> ((243) & 63)), F1: 433}
                _ = v13
            case "a":
                v7 -= (v7 >> ((789) & 63))
            }
            _ = fi12
        }
        _ = v8
    }
    v14 := GenS2{F0: v4}
    v3 -= v3
    _ = len(v6)
    switch "" {
    case "foo bar":
        var v15 uint64 = v3
        _ = v15
    default:
        for fi16 := 0; fi16 < 15; fi16++ {
            if false {
                var v17 uint16 = 42
                var v18 uint = uint(v3)
                switch (fi16 | 2147483647) {
                case 4294967295:
                    for fi19 := 0; fi19 < 7; fi19++ {
                        v17 -= v17
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if ((v3 >> ((v18) & 63)) == 4294967295) {
                            fmt.Println(v3)
                            var v20 string = ""
                            v21 := GenS1{F0: 0, F1: -9223372036854775808}
                            _ = v20
                            _ = v21
                        } else {
                            var v22 int = (fi16 + 671)
                            switch fi16 {
                            case 1:
                                fmt.Println(v4, v5)
                            }
                            v3 += v3
                            _ = v22
                        }
                        _ = fi19
                    }
                case 1:
                    fmt.Println(v2)
                }
                _ = v17
                _ = v18
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi16
        }
    }
    var v23 uint8 = 127
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := GenS3{F0: 919, F1: 0, F2: 65535}
        if true {
            var v26 int = fi24
            _ = v26
        }
        if false {
            var v27 bool = true
            _ = v27
        } else {
            v7 += (v7 | (v7 - (v7 % (((v7 + (v7 ^ v7))) | 1))))
        }
        _ = fi24
        _ = v25
    }
    fmt.Println(v5)
    switch ("hello" + "hello") {
    case "\n\t":
        var v28 uint64 = v3
        _ = v28
    case "foo bar":
        if true {
            if false {
                var v29 uint32 = 42
                var v30 uint8 = v23
                _ = v29
                _ = v30
            } else {
                switch -813 {
                case 561:
                    var v31 int32 = (v2 << ((4294967295) & 63))
                    _ = v31
                case 128:
                    for fi32 := 0; fi32 < 3; fi32++ {
                        v2 -= (-(v2))
                        _ = fi32
                    }
                case -264:
                    v2 = v2
                default:
                    var v33 int8 = (v7 >> ((1) & 63))
                    _ = v33
                }
                for fi34 := 0; fi34 < 13; fi34++ {
                    fmt.Println(v23, v3)
                    fmt.Println(v7)
                    _ = fi34
                }
            }
            if false {
                var v35 int16 = 1000
                for fi36 := 0; fi36 < 1; fi36++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi36
                }
                _ = v35
            }
        } else {
            if false {
                v6 = append(v6, 127)
                var v37 int64 = 456
                _ = v37
            } else {
                v38 := GenS4{F0: "foo bar"}
                _ = v38
            }
            fmt.Println(v4)
            var v39 string = "a"
            _ = v39
        }
    case "":
        v40 := GenS3{F0: 76, F1: 1, F2: 354}
        _ = v40
    default:
        fmt.Println(v23, v3, v23)
    }
    if false {
        v7 = (v7 | (-(v7)))
        for fi41 := 0; fi41 < 7; fi41++ {
            for fi42 := 0; fi42 < 9; fi42++ {
                if false {
                    fmt.Println(v2, v23, v7)
                    if (false && false) {
                        v3 = 18446744073709551615
                        var v43 uint16 = 32767
                        _ = v43
                    }
                }
                v5 = (1.5 - (float64(0.0) + v5))
                var v44 int8 = v7
                _ = fi42
                _ = v44
            }
            _ = fi41
        }
    } else {
        v45 := GenS4{F0: ("hello" + "foo bar")}
        v46 := GenS5{F0: (v2 >> ((657) & 63)), F1: 32767}
        _ = v45
        _ = v46
    }
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
