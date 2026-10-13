package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 uint16 }

func main() {
    var v1 uint16 = 1
    var v2 uint64 = 777
    var v3 int32 = -1
    var v4 int = -953
    v5 := []string{""}
    var v6 string = (("foo bar" + "\n\t") + "a")
    v1 += 0
    v5 = append(v5, "a")
    var v7 int8 = -71
    v7 += -1
    v4 = (^(v4))
    fmt.Println(v7, v4)
    switch 1 {
    case -99:
        if (-1 > 0) {
            var v8 int32 = (v3 + v3)
            v8 = v8
            switch 9223372036854775807 {
            case 32767:
                if false {
                    var v9 uint8 = uint8(v2)
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        for fi11 := 0; fi11 < 14; fi11++ {
                            var v12 int16 = int16(v2)
                            fmt.Println(v7)
                            var v13 int32 = v3
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 14; fi14++ {
                            v6 = (v6 + (v6 + v6))
                            fmt.Println(v6, v8)
                            _ = fi14
                        }
                        _ = fi10
                    }
                    var v15 uint16 = v1
                    _ = v15
                }
            case -1:
                _ = len(v5)
            case 973:
                switch v6 {
                case "\n\t":
                    var v16 uint32 = 296
                    _ = v16
                case "hello":
                    var v17 int = -9223372036854775808
                    _ = v17
                case "foo bar":
                    var v18 int64 = 9223372036854775807
                    _ = v18
                }
            }
            _ = v8
        } else {
            if (int(v3) <= v4) {
                if true {
                    _ = len(v5)
                } else {
                    v7 -= v7
                    fmt.Println(v2, v2, v6)
                    v19 := GenS1{F0: (v3 >> ((2) & 63)), F1: (v3 >> ((839) & 63))}
                    _ = v19
                }
            } else {
                var v20 int16 = 547
                var v21 uint64 = (v2 / ((42) | 1))
                _ = v20
                _ = v21
            }
        }
    case -1:
        var v22 uint8 = 0
        _ = v22
    default:
        if true {
            var v23 uint8 = 2
            if true {
                fmt.Println(v23, v7, v2)
                v24 := GenS1{F0: v3, F1: 42}
                _ = v24
            } else {
                var v25 int64 = 1
                _ = v25
            }
            _ = v23
        }
    }
    var v26 uint8 = 0
    if true {
        for fi27 := 0; fi27 < 8; fi27++ {
            var v28 uint16 = v1
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 uint32 = 2
                var v31 int32 = -2147483648
                v32 := GenS1{F0: v3, F1: v31}
                v4 -= v4
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            _ = fi27
            _ = v28
        }
        v33 := GenS2{F0: (v26 << ((42) & 63))}
        if (4294967295 >= 42) {
            var v34 int8 = (v7 << ((164) & 63))
            var v35 int64 = 929
            _ = v34
            _ = v35
        }
        _ = v33
    } else {
        v36 := GenS2{F0: v26}
        _ = v36
    }
    var v37 int16 = -591
    switch (-(v4)) {
    case -9223372036854775808:
        v37 += v37
    case 128:
        fmt.Println(v26, v37)
    }
    var v38 int8 = (v7 - int8(v1))
    switch (v4 % (((v4 & v4)) | 1)) {
    case 32768:
        fmt.Println(v38, v4, v6)
    default:
        if false {
            for fi39 := 0; fi39 < 20; fi39++ {
                v3 += 2147483647
                var v40 int32 = v3
                _ = fi39
                _ = v40
            }
        } else {
            for fi41 := 0; fi41 < 13; fi41++ {
                for fi42 := 0; fi42 < 3; fi42++ {
                    v6 = v6
                    v4 += -1
                    var v43 int64 = -9223372036854775808
                    _ = fi42
                    _ = v43
                }
                switch (v6 + v6) {
                case "a":
                    var v44 uint32 = 42
                    _ = v44
                case "":
                    v5 = append(v5, v6)
                case "hello":
                    switch v6 {
                    case "":
                        var v45 int64 = 42
                        _ = v45
                    case "hello":
                        _ = len(v5)
                    default:
                        var v46 bool = true
                        _ = v46
                    }
                default:
                    if false {
                        v2 -= v2
                    }
                }
                var v47 uint8 = (v26 ^ v26)
                _ = fi41
                _ = v47
            }
            v48 := GenS3{F0: v3, F1: v1}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v48
        }
    }
    v1 = (v1 - 42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
