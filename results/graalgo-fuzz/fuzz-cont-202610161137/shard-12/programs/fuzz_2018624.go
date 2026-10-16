package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint }

func main() {
    var v1 uint32 = 1
    var v2 bool = (-9223372036854775808 > 0)
    var v3 uint = 42
    v4 := []uint32{v1, (v1 | v1), 42, v1}
    v5 := make(map[string]bool)
    if v2 {
        var v6 int32 = 0
        fmt.Println(v6, v1, v1)
        v7 := GenS1{F0: (^(v1)), F1: 127}
        _ = v6
        _ = v7
    } else {
        var v8 int8 = 111
        fmt.Println(v8, v1, v3)
        for fi9 := 0; fi9 < 10; fi9++ {
            v4 = append(v4, v1)
            _ = fi9
        }
        _ = v8
    }
    v3 = v3
    fmt.Println(v3, v1)
    var v10 int32 = -1
    var v11 uint64 = 30
    for fi12 := 0; fi12 < 14; fi12++ {
        v10 += (v10 ^ v10)
        if true {
            v4 = append(v4, (v1 / ((v1) | 1)))
        } else {
            var v13 int16 = 42
            if ((-1 > 128) && (v11 == (v11 >> ((v3) & 63)))) {
                delete(v5, "foo bar")
                switch "\n\t" {
                case "hello":
                    for fi14 := 0; fi14 < 13; fi14++ {
                        if (v2 && ("hello" != ("" + (("foo bar" + ("a" + "a")) + "\n\t")))) {
                            v2 = false
                            fmt.Println(fi12)
                        }
                        delete(v5, "a")
                        v15 := GenS1{F0: uint32(v11), F1: -58}
                        _ = fi14
                        _ = v15
                    }
                case "a":
                    v4 = append(v4, v1)
                case "\n\t":
                    v3 = 126
                }
                v5["\n\t"] = (9223372036854775807 != -298)
            }
            _ = v13
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            if v2 {
                v2 = v2
                switch fi16 {
                case 4294967295:
                    if len(v4) > 0 {
                        v4[0] = uint32(v11)
                    }
                    var v17 int64 = -9223372036854775808
                    v18 := GenS2{F0: v10, F1: v3, F2: v1}
                    v3 += v3
                    switch "a" {
                    case "a":
                        if (!v2) {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v19 uint64 = (v11 ^ (v11 >> ((v3) & 63)))
                            delete(v5, "\n\t")
                            _ = v19
                        } else {
                            v5["a"] = v2
                        }
                    case "hello":
                        if v2 {
                            v2 = ((!true) || (3.14159 <= 0.0))
                            v2 = v2
                            fmt.Println(v17)
                        }
                    default:
                        _ = len(v5)
                    }
                    _ = v17
                    _ = v18
                default:
                    if ((!(-81 == -1)) || true) {
                        v11 = 4294967295
                        var v20 uint16 = 976
                        v10 = 515
                        _ = v20
                    } else {
                        v21 := GenS1{F0: v1, F1: 1}
                        fmt.Println(v3, v3, v2)
                        v22 := GenS1{F0: v1, F1: 21}
                        _ = v21
                        _ = v22
                    }
                    v4 = append(v4, v1)
                    v2 = true
                }
            }
            v23 := GenS1{F0: (v1 << (((^(v3))) & 63)), F1: 42}
            v4 = append(v4, 65535)
            _ = fi16
            _ = v23
        }
        _ = fi12
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        v25 := v5["\n\t"]
        _ = fi24
        _ = v25
    }
    if (((-1 == 227) || (1.5 >= 0.0)) && (!((!(-74 <= 101)) || (uint64(v1) >= (^(v11)))))) {
        v4 = append(v4, v1)
        var v26 uint32 = (v1 | 697)
        _ = v26
    }
    switch ("hello" + ("\n\t" + "\n\t")) {
    case "":
        if v2 {
            v3 = uint(v11)
        } else {
            if (float64((1e10 - 1e10)) > 0.0) {
                v11 = v11
                if (1e10 != 1.5) {
                    for fi27 := 0; fi27 < 1; fi27++ {
                        v2 = (v11 >= v11)
                        _ = fi27
                    }
                    for fi28 := 0; fi28 < 18; fi28++ {
                        var v29 int = 32768
                        _ = fi28
                        _ = v29
                    }
                    v1 -= v1
                }
            }
            v4 = append(v4, v1)
        }
    default:
        switch "hello" {
        case "":
            v30 := GenS1{F0: (v1 << ((v3) & 63)), F1: 127}
            _ = v30
        }
    }
    v31 := GenS1{F0: v1, F1: 42}
    for fi32 := 0; fi32 < 3; fi32++ {
        if (-128 != -97) {
            for fi33 := 0; fi33 < 18; fi33++ {
                var v34 int32 = -2147483648
                var v35 int16 = 0
                _ = fi33
                _ = v34
                _ = v35
            }
        }
        _ = fi32
    }
    v2 = (!(v1 != 65535))
    var v36 int16 = 738
    var v37 int8 = 12
    v38 := GenS3{F0: v11, F1: 42, F2: (v36 << ((v3) & 63))}
    for fi39 := 0; fi39 < 17; fi39++ {
        if false {
            var v40 uint = (v3 - (v3 | (v3 + 65535)))
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v11 = 97
            _ = v40
        }
        for fi41 := 0; fi41 < 12; fi41++ {
            var v42 uint8 = 146
            fmt.Println(v42)
            _ = len(v5)
            switch 127 {
            case 9223372036854775807:
                if v2 {
                    v2 = true
                }
            case -1:
                v10 += v10
            }
            _ = fi41
            _ = v42
        }
        v43 := GenS4{F0: 65535, F1: v10, F2: 65535}
        _ = fi39
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
