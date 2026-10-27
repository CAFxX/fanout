package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 int; F2 int8 }
type GenS5 struct { F0 uint16; F1 string; F2 int8 }

func main() {
    var v1 int16 = 42
    var v2 int16 = v1
    v3 := []int{32767, -396, -874}
    var v4 int64 = 1
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]uint32)
    switch (("" + ("a" + "")) + "foo bar") {
    case "a":
        fmt.Println(v1)
    case "foo bar":
        fmt.Println(v2)
    case "":
        if true {
            for fi7 := 0; fi7 < 2; fi7++ {
                v4 -= 670
                _ = fi7
            }
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 uint64 = 0
                switch -9223372036854775808 {
                case -1:
                    switch fi8 {
                    case 32768:
                        var v10 int = (-(fi8))
                        _ = v10
                    case 32767:
                        v2 -= v2
                    default:
                        var v11 int8 = 127
                        _ = v11
                    }
                case 9223372036854775807:
                    var v12 uint32 = 2147483647
                    _ = v12
                }
                _ = fi8
                _ = v9
            }
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS1{F0: 740, F1: 2147483647}
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 int64 = (v4 - (v4 + (v4 & 804)))
                    v3 = append(v3, 42)
                    v17 := GenS2{F0: 594}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v18 := GenS2{F0: 2}
                _ = fi13
                _ = v14
                _ = v18
            }
        }
    }
    var v19 int8 = 127
    fmt.Println(v2)
    v19 -= v19
    for fi20 := 0; fi20 < 7; fi20++ {
        v21 := GenS3{F0: 0}
        v6["foo bar"] = 4294967295
        for fi22 := 0; fi22 < 15; fi22++ {
            v23 := GenS2{F0: 909}
            var v24 float64 = 3.14159
            v19 += (v19 >> ((1) & 63))
            if false {
                _ = len(v3)
            }
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi20
        _ = v21
    }
    v2 = (-(v2))
    switch 9223372036854775807 {
    case 127:
        v25 := GenS3{F0: (v2 | v2)}
        _ = v25
    case 1:
        _ = len(v5)
    default:
        var v26 uint16 = 319
        _ = v26
    }
    var v27 int16 = (v2 % ((v1) | 1))
    switch ("\n\t" + "\n\t") {
    case "":
        var v28 uint16 = 338
        _ = v28
    case "foo bar":
        var v29 int64 = v4
        _ = v29
    default:
        if ("a" > ("foo bar" + (("foo bar" + "\n\t") + "foo bar"))) {
            switch 42 {
            case 32767:
                v30 := GenS1{F0: 0, F1: 2147483647}
                _ = v30
            case 838:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            if len(v3) > 0 {
                v3[0] = 388
            }
        } else {
            var v31 int16 = 1
            _ = v31
        }
    }
    fmt.Println(v27, v1, v2)
    var v32 int8 = (v19 - 127)
    for fi33 := 0; fi33 < 8; fi33++ {
        switch "\n\t" {
        case "\n\t":
            if false {
                var v34 uint8 = 148
                v5[167] = 0
                _ = v34
            } else {
                if (!false) {
                    _ = len(v3)
                }
                v35 := GenS3{F0: (v2 >> ((993) & 63))}
                _ = v35
            }
        case "foo bar":
            if false {
                v36 := GenS2{F0: 280}
                _ = v36
            } else {
                var v37 uint16 = 255
                v38 := GenS4{F0: v4, F1: (^(fi33)), F2: -6}
                fmt.Println(v27, v27, v1)
                _ = v37
                _ = v38
            }
        case "a":
            var v39 float64 = 1e10
            _ = v39
        }
        v3 = append(v3, (fi33 * -2147483648))
        if true {
            var v40 int32 = -1
            if false {
                v41 := GenS4{F0: v4, F1: 32768, F2: (v32 * 0)}
                var v42 uint64 = 1
                var v43 int32 = v40
                _ = v41
                _ = v42
                _ = v43
            } else {
                v40 = (v40 | v40)
                fmt.Println(v27, v40, v40)
                v44 := GenS1{F0: 1, F1: (v40 | (v40 << ((399) & 63)))}
                _ = v44
            }
            _ = v40
        }
        _ = fi33
    }
    var v45 string = (("" + "\n\t") + "hello")
    if true {
        if false {
            if ((65535 <= 970) && false) {
                v46 := GenS5{F0: 11, F1: v45, F2: (^(v19))}
                v47 := GenS1{F0: 4294967295, F1: 497}
                _ = v46
                _ = v47
            }
            fmt.Println(v27, v32)
        }
        for fi48 := 0; fi48 < 16; fi48++ {
            _ = len(v5)
            _ = fi48
        }
        for fi49 := 0; fi49 < 7; fi49++ {
            fmt.Println(v2)
            _ = fi49
        }
    }
    v50 := GenS2{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v45
    _ = v5
    _ = v50
    _ = v6
}
