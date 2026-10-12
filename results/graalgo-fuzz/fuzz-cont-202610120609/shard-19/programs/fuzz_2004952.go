package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint32, p1 bool, p2 int32) int64 {
    var v5 string = ""
    if p1 {
        switch 32768 {
        case 32768:
            var v6 string = "hello"
            _ = v6
        case -838:
            if ((p0 > (p0 + p0)) && ((float32(((1.5 + (3.14159 / 1.5)) * 3.14159)) + (0.1 - float32((1.5 + 0.1)))) == 3.14159)) {
                v7 := GenS1{F0: v5, F1: ((18446744073709551615 >= 42) || (p2 < p2)), F2: p1}
                _ = v7
            }
        default:
            v8 := []int{1, 0, -1}
            _ = v8
        }
    } else {
        switch (v5 + v5) {
        case "":
            switch 32767 {
            case 4294967295:
                var v9 float32 = 1.5
                _ = v9
            case 128:
                v5 = v5
            }
        case "\n\t":
            v10 := []int{-951}
            _ = v10
        case "a":
            var v11 uint32 = p0
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}
func genFn1(p0 uint64) uint32 {
    v12 := []int{42, -608, 128}
    switch "" {
    case "":
        if len(v12) > 0 {
            v12[0] = 127
        }
    case "foo bar":
        var v13 uint64 = (p0 + 629)
        _ = v13
    default:
        p0 = (p0 * 352)
    }
    _ = p0
    _ = v12
    return 2147483647
}

func main() {
    var v1 int16 = 32767
    var v2 bool = false
    v3 := []int{128}
    var v4 uint64 = 0
    v4 = 4294967295
    _ = len(v3)
    switch "a" {
    case "":
        for fi14 := 0; fi14 < 3; fi14++ {
            v1 = -32768
            var v15 int32 = 1
            v16 := GenS1{F0: ("foo bar" + ("hello" + "foo bar")), F1: v2, F2: (true || (770 != 147))}
            _ = fi14
            _ = v15
            _ = v16
        }
    case "\n\t":
        v4 -= 42
    case "a":
        var v17 uint8 = 2
        _ = v17
    }
    switch int(v4) {
    case -1:
        for fi18 := 0; fi18 < 11; fi18++ {
            if v2 {
                if ((v2 && v2) || v2) {
                    switch (fi18 % ((9223372036854775807) | 1)) {
                    case 4294967295:
                        var v19 uint64 = v4
                        _ = v19
                    }
                    var v20 string = (("foo bar" + ("hello" + ("" + "hello"))) + "a")
                    _ = v20
                } else {
                    fmt.Println(v4, fi18)
                    fmt.Println(v1, v4)
                }
            }
            var v21 int32 = 0
            _ = fi18
            _ = v21
        }
    }
    switch 4294967295 {
    case -9223372036854775808:
        if v2 {
            v22 := GenS2{F0: 52}
            _ = v22
        } else {
            var v23 uint32 = uint32(v4)
            if (v23 > v23) {
                _ = len(v3)
                fmt.Println(v23, v2)
                v24 := GenS2{F0: 1}
                _ = v24
            } else {
                v23 = 2147483648
                var v25 uint8 = 255
                v26 := GenS2{F0: v25}
                _ = v25
                _ = v26
            }
            _ = v23
        }
    case 127:
        switch 236 {
        case -2147483648:
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                v27 := GenS2{F0: 131}
                _ = v27
            case "a":
                if (0 > 1) {
                    v3 = append(v3, -988)
                }
            case "\n\t":
                if (!v2) {
                    for fi28 := 0; fi28 < 11; fi28++ {
                        var v29 int16 = 0
                        _ = fi28
                        _ = v29
                    }
                    var v30 int = 2147483647
                    v31 := GenS3{F0: (v1 * (v1 << ((0) & 63)))}
                    _ = v30
                    _ = v31
                } else {
                    v4 = v4
                }
            default:
                var v32 int = 42
                _ = v32
            }
        default:
            for fi33 := 0; fi33 < 14; fi33++ {
                if len(v3) > 0 {
                    v3[0] = fi33
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi33
            }
        }
    case 1:
        v34 := GenS2{F0: 1}
        _ = v34
    default:
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v2, v1)
        case "foo bar":
            var v35 string = ("hello" + ("" + ("\n\t" + "a")))
            _ = v35
        case "hello":
            v36 := GenS4{F0: 65535}
            _ = v36
        default:
            v37 := GenS3{F0: (v1 & -487)}
            _ = v37
        }
    }
    v38 := GenS3{F0: v1}
    for fi39 := 0; fi39 < 19; fi39++ {
        v4 = 0
        var v40 uint8 = 0
        var v41 int64 = 9223372036854775807
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println(v2)
    v42 := GenS4{F0: 2147483647}
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 += v1
    }
    v1 -= v1
    for fi43 := 0; fi43 < 20; fi43++ {
        v2 = (!false)
        var v44 uint64 = 686
        for fi45 := 0; fi45 < 18; fi45++ {
            var v46 uint8 = 0
            var v47 uint16 = 740
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
}
