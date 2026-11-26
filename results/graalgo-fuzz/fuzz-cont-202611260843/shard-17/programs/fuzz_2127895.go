package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint8; F2 int }
type GenS3 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 int16 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[string]uint32)
    var v3 bool = true
    var v4 int32 = 50
    var v5 uint16 = 42
    v6 := make(map[uint32]string)
    if true {
        v4 -= (v4 ^ 1)
        var v7 uint = 171
        switch ("" + "") {
        case "foo bar":
            var v8 int64 = -298
            _ = v8
        default:
            var v9 int = 1
            _ = v9
        }
        _ = v7
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint64 = (v1 % (((v1 >> ((203) & 63))) | 1))
        var v12 string = "foo bar"
        _ = fi10
        _ = v11
        _ = v12
    }
    switch 32768 {
    case -188:
        switch "\n\t" {
        case "":
            var v13 uint32 = 2147483648
            _ = v13
        }
    case 42:
        switch ("a" + "a") {
        case "foo bar":
            for fi14 := 0; fi14 < 9; fi14++ {
                if (!(1 < (fi14 / (((fi14 / ((fi14) | 1))) | 1)))) {
                    if (!false) {
                        var v15 int = -1
                        _ = v15
                    } else {
                        v1 = v1
                        for fi16 := 0; fi16 < 12; fi16++ {
                            var v17 int8 = 42
                            switch ("hello" + "") {
                            case "hello":
                                var v18 string = "foo bar"
                                _ = v18
                            case "a":
                                v19 := GenS1{F0: 681}
                                _ = v19
                            }
                            var v20 uint8 = 137
                            _ = fi16
                            _ = v17
                            _ = v20
                        }
                    }
                    v21 := GenS2{F0: 142, F1: 2, F2: (^(fi14))}
                    _ = v21
                } else {
                    v4 -= -750
                    v22 := GenS2{F0: -9223372036854775808, F1: 42, F2: (^(fi14))}
                    var v23 bool = v3
                    _ = v22
                    _ = v23
                }
                v24 := GenS2{F0: 9223372036854775807, F1: 114, F2: (fi14 - fi14)}
                for fi25 := 0; fi25 < 13; fi25++ {
                    var v26 string = "\n\t"
                    fmt.Println(v1)
                    v27 := GenS1{F0: 65535}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                _ = fi14
                _ = v24
            }
        case "":
            if (true && v3) {
                var v28 uint = 16
                v29 := GenS3{F0: 829, F1: 1, F2: v4}
                if v3 {
                    var v30 int32 = -366
                    _ = v30
                }
                _ = v28
                _ = v29
            } else {
                fmt.Println(v1, v4)
                v31 := []int{32767, 527, -1, -922}
                v5 = 468
                _ = v31
            }
        }
    }
    if (1.5 <= 3.14159) {
        for fi32 := 0; fi32 < 18; fi32++ {
            for fi33 := 0; fi33 < 17; fi33++ {
                if v3 {
                    var v34 uint64 = (v1 << ((37) & 63))
                    _ = v34
                }
                _ = fi33
            }
            switch ("\n\t" + "") {
            case "":
                var v35 uint16 = (v5 + v5)
                _ = v35
            }
            for fi36 := 0; fi36 < 4; fi36++ {
                var v37 int = 32767
                v5 = v5
                _ = fi36
                _ = v37
            }
            _ = fi32
        }
        v38 := GenS1{F0: 611}
        fmt.Println(v4)
        _ = v38
    } else {
        fmt.Println(v4, v4)
    }
    var v39 uint16 = v5
    var v40 int8 = 125
    if v3 {
        v41 := v2[("\n\t" + "")]
        v42 := GenS4{F0: 0, F1: "foo bar"}
        var v43 int8 = (v40 >> ((42) & 63))
        _ = v41
        _ = v42
        _ = v43
    } else {
        switch "hello" {
        case "\n\t":
            if (136 != 65535) {
                v6[2147483648] = "foo bar"
            } else {
                fmt.Println(v5)
            }
        case "foo bar":
            fmt.Println(v40)
        case "hello":
            if true {
                switch int(v1) {
                case 128:
                    var v44 uint64 = 18446744073709551615
                    _ = v44
                case 1:
                    v45 := GenS5{F0: 1}
                    _ = v45
                case -1:
                    for fi46 := 0; fi46 < 3; fi46++ {
                        delete(v6, 542)
                        var v47 int8 = v40
                        switch (fi46 % (((^(fi46))) | 1)) {
                        case 32767:
                            delete(v6, 2147483648)
                        case 127:
                            v1 = 1
                        case 128:
                            v1 += 0
                        }
                        _ = fi46
                        _ = v47
                    }
                default:
                    fmt.Println(v39, v39)
                }
                fmt.Println(v39)
            }
        default:
            var v48 float64 = (float64(0.1) + (1.5 + 1.5))
            _ = v48
        }
    }
    for fi49 := 0; fi49 < 12; fi49++ {
        v50 := GenS3{F0: 834, F1: 1, F2: int32(fi49)}
        _ = fi49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
