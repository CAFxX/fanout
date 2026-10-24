package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int64 }

func main() {
    v1 := []string{"a"}
    var v2 string = "foo bar"
    v3 := make(map[int]string)
    var v4 int16 = -1
    for fi5 := 0; fi5 < 10; fi5++ {
        switch (fi5 & -9223372036854775808) {
        case -1:
            var v6 uint8 = 2
            _ = v6
        case 9223372036854775807:
            var v7 uint64 = 4294967295
            _ = v7
        }
        _ = fi5
    }
    v4 -= (v4 / ((v4) | 1))
    switch 32767 {
    case -2147483648:
        fmt.Println(v4, v4)
    case 643:
        v2 = (v2 + ("" + v2))
    }
    fmt.Println(v4)
    v8 := GenS1{F0: -9223372036854775808}
    if true {
        delete(v3, -213)
        if false {
            fmt.Println(v4)
            if false {
                if ((53 < 0) && (!false)) {
                    var v9 float64 = 0.0
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v9
                }
            } else {
                for fi10 := 0; fi10 < 17; fi10++ {
                    for fi11 := 0; fi11 < 15; fi11++ {
                        var v12 int16 = v4
                        var v13 uint64 = 965
                        v14 := v3[fi10]
                        fmt.Println(v14, v14)
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    _ = fi10
                }
            }
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 uint = 799
                v2 = (((v2 + v2) + ("\n\t" + (v2 + "\n\t"))) + ("foo bar" + v2))
                _ = fi15
                _ = v16
            }
        } else {
            var v17 float32 = -2.25
            _ = v17
        }
    } else {
        if true {
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 int16 = 32767
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                fmt.Println(fi18)
                _ = fi18
                _ = v19
            }
            if true {
                _ = len(v3)
            }
        }
        v4 = v4
        switch ("hello" + "a") {
        case "a":
            if false {
                for fi20 := 0; fi20 < 17; fi20++ {
                    var v21 int8 = 42
                    var v22 uint64 = 2
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                for fi23 := 0; fi23 < 18; fi23++ {
                    if (2 <= 666) {
                        _ = len(v1)
                    }
                    _ = fi23
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi25 := 0; fi25 < 4; fi25++ {
                        v4 = (v4 << ((1) & 63))
                        switch (("a" + v2) + "\n\t") {
                        case "hello":
                            var v26 int64 = 867
                            _ = v26
                        case "a":
                            var v27 int64 = -9223372036854775808
                            _ = v27
                        default:
                            v28 := GenS2{F0: 1, F1: 42}
                            _ = v28
                        }
                        _ = fi25
                    }
                    _ = fi24
                }
            } else {
                switch 707 {
                case 128:
                    v1 = append(v1, v2)
                default:
                    v4 = v4
                }
            }
        case "hello":
            v29 := GenS2{F0: 42, F1: 42}
            _ = v29
        default:
            v2 = ("foo bar" + v2)
        }
    }
    var v30 int16 = (v4 + v4)
    for fi31 := 0; fi31 < 13; fi31++ {
        var v32 int32 = -2147483648
        var v33 float32 = 0.0
        switch (fi31 + -9223372036854775808) {
        case 4294967295:
            v32 = (v32 / ((v32) | 1))
        case 0:
            v4 = v4
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    v30 += (v30 >> ((1) & 63))
    for fi34 := 0; fi34 < 1; fi34++ {
        fmt.Println(v30)
        v4 = (v30 >> ((42) & 63))
        v4 -= v30
        _ = fi34
    }
    _ = len(v1)
    v30 = (v4 | (v30 & v4))
    for fi35 := 0; fi35 < 5; fi35++ {
        var v36 uint16 = 302
        if false {
            var v37 uint = 65535
            delete(v3, fi35)
            v38 := GenS2{F0: 2, F1: 591}
            _ = v37
            _ = v38
        } else {
            switch (fi35 * 4294967295) {
            case 0:
                v2 = ("hello" + ("" + v2))
            case 1:
                v4 = 35
            case 9223372036854775807:
                var v39 uint8 = 255
                _ = v39
            }
        }
        _ = fi35
        _ = v36
    }
    fmt.Println(v4, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v8
}
