package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int8; F1 uint32; F2 uint16 }

func main() {
    var v1 int8 = 127
    v2 := make(map[string]uint32)
    var v3 uint16 = 2
    v4 := []string{("\n\t" + ("hello" + "a"))}
    var v5 float64 = (float64(1e10) + 1e10)
    v6 := v2["hello"]
    v2[""] = (^(v6))
    for fi7 := 0; fi7 < 14; fi7++ {
        if false {
            var v8 string = "\n\t"
            _ = v8
        }
        var v9 float64 = 1e10
        fmt.Println(v5)
        _ = fi7
        _ = v9
    }
    v10 := v2["\n\t"]
    v11 := GenS1{F0: v1, F1: -538}
    v1 -= (v1 - v1)
    if (!false) {
        fmt.Println(v3, v1)
        switch -1 {
        case 42:
            if false {
                for fi12 := 0; fi12 < 11; fi12++ {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        if (!false) {
                            fmt.Println(v3, v5, fi13)
                            v1 += (-(v1))
                            switch "foo bar" {
                            case "\n\t":
                                v1 = (v1 + v1)
                            }
                        } else {
                            if (!true) {
                                v10 = v10
                            }
                        }
                        v14 := GenS1{F0: 1, F1: (^(fi12))}
                        _ = fi13
                        _ = v14
                    }
                    var v15 uint16 = 2
                    _ = fi12
                    _ = v15
                }
                v16 := GenS2{F0: true, F1: 25}
                _ = v16
            } else {
                var v17 string = ""
                v18 := GenS2{F0: false, F1: 1}
                _ = v17
                _ = v18
            }
        case -2147483648:
            var v19 int8 = v1
            _ = v19
        }
    } else {
        v4 = append(v4, "")
        switch 98 {
        case -9223372036854775808:
            v1 = (v1 << ((65535) & 63))
        case 128:
            v20 := GenS3{F0: 65535}
            _ = v20
        default:
            fmt.Println(v3, v10, v6)
        }
    }
    if (true && (32768 != -1)) {
        var v21 float32 = 1e10
        var v22 int32 = 42
        _ = v21
        _ = v22
    }
    v10 = 65535
    v1 = (v1 * v1)
    for fi23 := 0; fi23 < 14; fi23++ {
        var v24 uint16 = v3
        _ = fi23
        _ = v24
    }
    var v25 int32 = -2147483648
    if true {
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 uint64 = 42
            var v28 uint32 = 4294967295
            v28 = (v6 >> ((450) & 63))
            _ = fi26
            _ = v27
            _ = v28
        }
        var v29 uint64 = 42
        v6 -= (v6 >> ((90) & 63))
        _ = v29
    } else {
        if true {
            if (false && true) {
                v30 := v2[""]
                var v31 uint = 126
                _ = v30
                _ = v31
            } else {
                fmt.Println(v5, v25, v25)
            }
            for fi32 := 0; fi32 < 6; fi32++ {
                v33 := GenS4{F0: uint8(v6), F1: v6, F2: 4294967295}
                _ = fi32
                _ = v33
            }
        } else {
            if true {
                if false {
                    for fi34 := 0; fi34 < 19; fi34++ {
                        fmt.Println(v6)
                        _ = fi34
                    }
                } else {
                    v35 := GenS3{F0: 98}
                    v1 -= v1
                    for fi36 := 0; fi36 < 14; fi36++ {
                        var v37 string = "foo bar"
                        _ = fi36
                        _ = v37
                    }
                    _ = v35
                }
                var v38 int8 = (v1 << ((1) & 63))
                _ = v38
            } else {
                var v39 int = -9223372036854775808
                fmt.Println(v10, v5)
                _ = v39
            }
            switch 751 {
            case 2147483647:
                var v40 float32 = 0.0
                _ = v40
            }
        }
        for fi41 := 0; fi41 < 8; fi41++ {
            if false {
                delete(v2, "a")
                if false {
                    v1 = 42
                    switch (("hello" + "hello") + "a") {
                    case "\n\t":
                        switch (fi41 - fi41) {
                        case 0:
                            v6 += v6
                        case 9223372036854775807:
                            var v42 int16 = -1
                            _ = v42
                        default:
                            v5 = v5
                        }
                    }
                    var v43 int64 = int64(v3)
                    _ = v43
                } else {
                    v25 = v25
                    if false {
                        var v44 string = ("" + "a")
                        fmt.Println(v3)
                        v5 -= 1.5
                        _ = v44
                    }
                }
                var v45 uint16 = 65535
                _ = v45
            }
            _ = len(v2)
            for fi46 := 0; fi46 < 19; fi46++ {
                switch (fi41 | 32767) {
                case 103:
                    var v47 uint16 = v3
                    _ = v47
                case -2147483648:
                    v1 = (v1 ^ int8(v6))
                default:
                    switch "a" {
                    case "foo bar":
                        v48 := GenS5{F0: (^(v1)), F1: 125, F2: v3}
                        _ = v48
                    case "hello":
                        switch 0 {
                        case -9223372036854775808:
                            v1 = v1
                        case 127:
                            v25 = v25
                        case 1:
                            v49 := GenS1{F0: (v1 & (v1 ^ (v1 | (-(v1))))), F1: (^(fi46))}
                            _ = v49
                        }
                    }
                }
                if len(v4) > 0 {
                    v4[0] = (((("hello" + ("" + "\n\t")) + "\n\t") + "a") + "hello")
                }
                _ = fi46
            }
            _ = fi41
        }
    }
    v50 := v2[("" + "hello")]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
