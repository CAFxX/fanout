package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int16; F2 int16 }
type GenS5 struct { F0 int16; F1 uint; F2 uint32 }

func main() {
    var v1 string = "a"
    v2 := make(map[string]string)
    var v3 uint16 = 42
    var v4 uint64 = 0
    v5 := make(map[int]bool)
    var v6 uint32 = 42
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    switch v1 {
    case "":
        v6 -= 366
    case "\n\t":
        fmt.Println(v3, v6, v1)
    }
    v9 := GenS1{F0: 163}
    delete(v2, v1)
    for fi10 := 0; fi10 < 6; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            if (!false) {
                var v12 int = 2147483647
                var v13 bool = (true || true)
                _ = v12
                _ = v13
            } else {
                switch "hello" {
                case "a":
                    switch (fi10 - -9223372036854775808) {
                    case -2147483648:
                        var v14 int16 = -1
                        _ = v14
                    case 32767:
                        fmt.Println(v6, v3)
                    }
                case "foo bar":
                    v15 := v5[(-(fi10))]
                    _ = v15
                default:
                    fmt.Println(v1, v4)
                }
                if (v6 <= 2) {
                    var v16 uint = 1
                    _ = v16
                }
            }
            fmt.Println(v1, v4)
            _ = fi11
        }
        _ = fi10
    }
    v17 := GenS2{F0: -63, F1: v1}
    v1 = "a"
    if true {
        switch "" {
        case "\n\t":
            v18 := GenS2{F0: 0, F1: (v1 + ("foo bar" + (v1 + "hello")))}
            _ = v18
        }
        if (9223372036854775807 > -9223372036854775808) {
            for fi19 := 0; fi19 < 10; fi19++ {
                v20 := GenS1{F0: 42}
                if false {
                    switch v1 {
                    case "foo bar":
                        switch -664 {
                        case -9223372036854775808:
                            fmt.Println(v3, v4, fi19)
                        case 42:
                            v1 = ""
                        case -209:
                            v21 := v5[fi19]
                            _ = v21
                        default:
                            delete(v5, -2147483648)
                        }
                    case "hello":
                        v3 = (v3 << ((65535) & 63))
                    case "a":
                        var v22 bool = false
                        _ = v22
                    default:
                        var v23 uint16 = (v3 & v3)
                        _ = v23
                    }
                }
                v24 := v5[(fi19 ^ (fi19 % ((0) | 1)))]
                _ = fi19
                _ = v20
                _ = v24
            }
            if ("a" <= (v1 + "foo bar")) {
                switch ((("hello" + ("foo bar" + "foo bar")) + "\n\t") + "\n\t") {
                case "\n\t":
                    var v25 string = v1
                    _ = v25
                case "foo bar":
                    switch 0 {
                    case 1:
                        for fi26 := 0; fi26 < 16; fi26++ {
                            var v27 float32 = 1.5
                            _ = fi26
                            _ = v27
                        }
                    case 32768:
                        var v28 uint32 = (^(v6))
                        _ = v28
                    case -2147483648:
                        var v29 uint16 = v3
                        _ = v29
                    }
                case "":
                    v4 += 889
                }
                switch int(v4) {
                case 489:
                    var v30 int64 = 9223372036854775807
                    _ = v30
                case 42:
                    var v31 int8 = -28
                    _ = v31
                case 1:
                    fmt.Println(v6)
                default:
                    if ((v4 % ((uint64(v3)) | 1)) != (v4 | v4)) {
                        v32 := []int{4294967295}
                        for fi33 := 0; fi33 < 14; fi33++ {
                            v34 := GenS3{F0: 18446744073709551615}
                            fmt.Println(v6)
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    } else {
                        v35 := []int{0, 128}
                        v4 = v4
                        var v36 uint8 = 42
                        _ = v35
                        _ = v36
                    }
                }
            } else {
                if (0.0 <= 0.0) {
                    v37 := []int{386, 685, 239}
                    _ = v37
                }
                fmt.Println(v4, v4, v4)
            }
            v38 := GenS4{F0: int8(v3), F1: 1, F2: 32767}
            _ = v38
        }
    } else {
        var v39 uint64 = v4
        if (false && false) {
            v40 := GenS5{F0: 32767, F1: 2, F2: 594}
            v39 = v39
            v41 := []int{}
            _ = v40
            _ = v41
        } else {
            switch 1 {
            case 2147483647:
                switch -9223372036854775808 {
                case -698:
                    if true {
                        for fi42 := 0; fi42 < 16; fi42++ {
                            delete(v5, fi42)
                            var v43 uint32 = 2
                            v4 = 580
                            _ = fi42
                            _ = v43
                        }
                    }
                }
            case 32767:
                var v44 int16 = 439
                _ = v44
            case 163:
                var v45 uint8 = 0
                _ = v45
            default:
                fmt.Println(v6, v6, v1)
            }
            fmt.Println(v1)
        }
        v46 := GenS5{F0: 42, F1: 1, F2: v6}
        _ = v39
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
