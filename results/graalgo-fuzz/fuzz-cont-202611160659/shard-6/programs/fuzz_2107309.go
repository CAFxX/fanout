package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 bool; F1 int; F2 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"a", "foo bar"}
    v3 := make(map[uint32]uint32)
    var v4 int = 32767
    var v5 uint32 = 694
    var v6 bool = (1 > 32767)
    switch "hello" {
    case "":
        v6 = (v4 == v4)
    default:
        if v6 {
            var v7 string = "foo bar"
            _ = v7
        }
    }
    switch v4 {
    case -393:
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS1{F0: (^(fi8)), F1: 2, F2: 530}
            var v10 int8 = -103
            _ = fi8
            _ = v10
            _ = v9
        }
    case -786:
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS2{F0: -32768}
            if ((v6 && ((0.0 * 3.14159) != 0.1)) || v6) {
                fmt.Println(fi11, v5)
                var v13 uint8 = 128
                switch (fi11 << ((65535) & 63)) {
                case 200:
                    _ = len(v3)
                }
                _ = v13
            }
            _ = fi11
            _ = v12
        }
    case 783:
        v3[v5] = (^(v5))
    default:
        v6 = false
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        v4 += (fi14 * (fi14 - -9223372036854775808))
        _ = fi14
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v4 -= (v4 << ((553) & 63))
        _ = fi15
    }
    var v16 int32 = int32(v5)
    v17 := GenS1{F0: 0, F1: 2, F2: 4294967295}
    v16 = v16
    v16 = -51
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v6 = v6
    v4 -= (v4 << ((2) & 63))
    if v6 {
        switch (v4 / ((v4) | 1)) {
        case 32768:
            switch -342 {
            case 84:
                var v18 int16 = 32767
                _ = v18
            }
        case 4294967295:
            v4 = v4
        case 1:
            fmt.Println(v6, v6)
        }
    }
    var v19 bool = (1 != 42)
    v20 := GenS3{F0: 4294967295, F1: -9223372036854775808}
    v21 := GenS3{F0: 42, F1: -356}
    if false {
        var v22 int16 = -1
        _ = v22
    } else {
        v3[v5] = v5
        v23 := v1[v4]
        for fi24 := 0; fi24 < 20; fi24++ {
            if v23 {
                v23 = (42 <= 127)
                v6 = ((("\n\t" + "\n\t") + "") > ("foo bar" + ""))
            }
            _ = fi24
        }
        _ = v23
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        if true {
            v19 = (!v6)
        } else {
            v26 := GenS4{F0: v19, F1: 9223372036854775807, F2: -1}
            _ = v26
        }
        switch "foo bar" {
        case "":
            v3[v5] = (v5 & 0)
        case "foo bar":
            if v6 {
                for fi27 := 0; fi27 < 3; fi27++ {
                    v4 = (^(fi25))
                    v16 = v16
                    _ = fi27
                }
            }
        default:
            v28 := GenS1{F0: v4, F1: 18446744073709551615, F2: 0}
            _ = v28
        }
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
