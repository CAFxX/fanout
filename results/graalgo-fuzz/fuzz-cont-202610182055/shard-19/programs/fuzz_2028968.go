package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint64 }

func main() {
    v1 := []string{"hello"}
    var v2 string = ("a" + "a")
    var v3 uint32 = 2
    var v4 uint16 = uint16(v3)
    v5 := []int{-9223372036854775808}
    var v6 int64 = 1
    switch v2 {
    case "foo bar":
        var v7 uint64 = uint64(v4)
        _ = v7
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "\n\t":
        v4 = (v4 - (v4 * v4))
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v6 -= v6
        v3 = (v3 % (((v3 & (^(v3)))) | 1))
        _ = fi8
    }
    v3 = (v3 * 42)
    v2 = (v2 + v2)
    var v9 int32 = 425
    if len(v5) > 0 {
        v5[0] = 60
    }
    if true {
        v9 += v9
    } else {
        v10 := GenS1{F0: v9, F1: -1, F2: 0}
        v11 := GenS2{F0: v3}
        _ = v10
        _ = v11
    }
    v9 += 1
    _ = len(v5)
    if (false || true) {
        var v12 int64 = v6
        v12 -= v12
        v13 := GenS2{F0: 283}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v15 := GenS3{F0: 127, F1: false}
        if (!((v4 / (((^(v4))) | 1)) > v4)) {
            v9 += v9
        }
        for fi16 := 0; fi16 < 6; fi16++ {
            switch (fi16 << ((18446744073709551615) & 63)) {
            case 4294967295:
                switch "hello" {
                case "a":
                    v17 := GenS4{F0: (v4 ^ (v4 & v4)), F1: true}
                    _ = v17
                case "hello":
                    var v18 int32 = (v9 >> ((65535) & 63))
                    _ = v18
                case "":
                    v3 = v3
                }
                v9 += -889
                if true {
                    fmt.Println(fi14, fi16, fi16)
                }
            case -947:
                v6 -= 42
            }
            var v19 bool = (127 >= -6)
            _ = fi16
            _ = v19
        }
        _ = fi14
        _ = v15
    }
    if false {
        if (true && (2 <= 209)) {
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
            for fi20 := 0; fi20 < 4; fi20++ {
                v6 += (v6 ^ (v6 + (v6 << ((2) & 63))))
                var v21 int8 = -1
                _ = fi20
                _ = v21
            }
        } else {
            v6 = v6
        }
        v2 = (((v2 + "hello") + ("foo bar" + v2)) + v2)
    } else {
        for fi22 := 0; fi22 < 3; fi22++ {
            for fi23 := 0; fi23 < 1; fi23++ {
                switch v2 {
                case "\n\t":
                    var v24 string = (v2 + "hello")
                    _ = v24
                case "":
                    if ((1e10 - 3.14159) == (-2.25 + 1.5)) {
                        v6 = v6
                        v4 -= v4
                    }
                default:
                    switch (fi22 - 81) {
                    case 9223372036854775807:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        fmt.Println(v3)
                    }
                }
                var v25 string = v2
                v1 = append(v1, "")
                _ = fi23
                _ = v25
            }
            _ = fi22
        }
    }
    switch "a" {
    case "a":
        v26 := GenS1{F0: 0, F1: -247, F2: -1}
        _ = v26
    case "\n\t":
        var v27 uint = 1
        _ = v27
    case "":
        v6 = -990
    }
    switch 927 {
    case 128:
        switch -854 {
        case -222:
            for fi28 := 0; fi28 < 14; fi28++ {
                v9 = 42
                var v29 int16 = 0
                switch ("foo bar" + "hello") {
                case "":
                    if (true || true) {
                        v6 -= v6
                        if len(v5) > 0 {
                            v5[0] = (fi28 / ((-405) | 1))
                        }
                        fmt.Println(fi28)
                    } else {
                        switch (fi28 & (fi28 ^ (fi28 | fi28))) {
                        case 1:
                            v30 := GenS2{F0: 65535}
                            _ = v30
                        case 32768:
                            v31 := GenS5{F0: 0}
                            _ = v31
                        case 2147483647:
                            v1 = append(v1, v2)
                        }
                    }
                case "\n\t":
                    var v32 bool = false
                    _ = v32
                }
                _ = fi28
                _ = v29
            }
        default:
            for fi33 := 0; fi33 < 4; fi33++ {
                for fi34 := 0; fi34 < 9; fi34++ {
                    var v35 bool = (42 != -1)
                    switch v2 {
                    case "hello":
                        v36 := GenS3{F0: (fi34 & (fi34 & 42)), F1: (!(!false))}
                        _ = v36
                    }
                    _ = fi34
                    _ = v35
                }
                _ = fi33
            }
        }
    }
    v37 := GenS5{F0: 432}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
