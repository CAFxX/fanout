package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS4 struct { F0 int16; F1 uint32; F2 int8 }
type GenS5 struct { F0 int64; F1 string; F2 uint8 }
type GenS6 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS7 struct { F0 uint32; F1 int64 }

func main() {
    var v1 bool = false
    v2 := []int{}
    var v3 int = 127
    var v4 uint8 = 2
    _ = len(v2)
    v2 = append(v2, 128)
    for fi5 := 0; fi5 < 18; fi5++ {
        if (-1 >= 0) {
            if (false || v1) {
                var v6 uint16 = 288
                var v7 int16 = 541
                for fi8 := 0; fi8 < 7; fi8++ {
                    switch fi5 {
                    case -2147483648:
                        var v9 uint = 2
                        _ = v9
                    case 2147483647:
                        var v10 uint8 = v4
                        _ = v10
                    default:
                        var v11 bool = v1
                        _ = v11
                    }
                    _ = fi8
                }
                _ = v6
                _ = v7
            } else {
                v12 := GenS1{F0: -9223372036854775808, F1: 1, F2: 4294967295}
                _ = v12
            }
            var v13 int64 = 1
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 int32 = 675
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        var v16 float32 = -2.25
        v2 = append(v2, v3)
        _ = fi5
        _ = v16
    }
    var v17 int32 = 0
    switch "\n\t" {
    case "a":
        v18 := GenS1{F0: 42, F1: 49, F2: 563}
        _ = v18
    case "\n\t":
        v19 := GenS2{F0: 911, F1: 0}
        _ = v19
    default:
        var v20 int8 = 0
        _ = v20
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        if len(v2) > 0 {
            v2[0] = v3
        }
        _ = fi21
    }
    for fi22 := 0; fi22 < 8; fi22++ {
        var v23 uint8 = v4
        var v24 uint32 = 1
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS2{F0: 42, F1: -128}
            v28 := GenS3{F0: 42, F1: 0, F2: 0}
            fmt.Println(v3, fi26)
            fmt.Println(v4)
            var v29 uint = 255
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
        v1 = true
        _ = fi25
    }
    v30 := GenS4{F0: -32768, F1: 2147483647, F2: 42}
    v3 = (-(v3))
    v1 = ((!false) || ((^(v17)) <= (v17 << ((353) & 63))))
    v31 := GenS4{F0: 32767, F1: 775, F2: 42}
    for fi32 := 0; fi32 < 2; fi32++ {
        v2 = append(v2, -9223372036854775808)
        _ = fi32
    }
    if v1 {
        if (!(0 < 0)) {
            fmt.Println(v1, v17)
        } else {
            for fi33 := 0; fi33 < 17; fi33++ {
                if len(v2) > 0 {
                    v2[0] = -640
                }
                if v1 {
                    fmt.Println(v1, v1)
                    var v34 bool = v1
                    _ = v34
                }
                _ = fi33
            }
            switch (("a" + "\n\t") + ("foo bar" + ("foo bar" + "foo bar"))) {
            case "hello":
                switch v3 {
                case -9223372036854775808:
                    switch -641 {
                    case 526:
                        _ = len(v2)
                    case 907:
                        v4 = v4
                    }
                case 128:
                    v2 = append(v2, (-(v3)))
                case 0:
                    var v35 uint64 = 0
                    _ = v35
                }
            case "foo bar":
                for fi36 := 0; fi36 < 5; fi36++ {
                    switch v3 {
                    case 0:
                        v4 += 127
                    }
                    var v37 int64 = 9223372036854775807
                    _ = fi36
                    _ = v37
                }
            case "\n\t":
                v38 := GenS5{F0: 0, F1: "", F2: 3}
                _ = v38
            }
            for fi39 := 0; fi39 < 8; fi39++ {
                switch -9223372036854775808 {
                case -9223372036854775808:
                    var v40 int32 = v17
                    _ = v40
                case 4294967295:
                    var v41 string = ""
                    _ = v41
                }
                if v1 {
                    fmt.Println(v17, v3)
                    var v42 bool = true
                    _ = v42
                }
                _ = fi39
            }
        }
        var v43 bool = (!v1)
        for fi44 := 0; fi44 < 8; fi44++ {
            v45 := GenS6{F0: v4, F1: (v4 + (^(v4))), F2: (v17 ^ 42)}
            var v46 int16 = -32768
            var v47 uint32 = 1
            _ = fi44
            _ = v45
            _ = v46
            _ = v47
        }
        _ = v43
    } else {
        v4 = v4
        if (1 == 1) {
            v48 := GenS7{F0: 828, F1: 42}
            _ = v48
        } else {
            v4 += v4
            var v49 uint16 = 0
            _ = v49
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
