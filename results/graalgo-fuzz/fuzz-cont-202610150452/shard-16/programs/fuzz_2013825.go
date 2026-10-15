package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
func genFn0(p0 int16, p1 uint64) uint16 {
    p0 += 330
    fmt.Println(p0)
    var v5 int8 = 42
    var v6 uint16 = 32767
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return v6
}
func genFn1() int32 {
    v7 := []int{-9223372036854775808}
    if (!((-48 != 100) || (false || false))) {
        v8 := GenS1{F0: 1, F1: 32767, F2: "a"}
        _ = len(v7)
        var v9 int8 = 127
        _ = v8
        _ = v9
    } else {
        switch "hello" {
        case "foo bar":
            var v10 bool = true
            _ = v10
        case "hello":
            if (953 < 0) {
                v11 := GenS1{F0: 42, F1: 32767, F2: ("" + "foo bar")}
                var v12 bool = false
                _ = v11
                _ = v12
            }
        case "a":
            var v13 uint16 = 366
            _ = v13
        }
        v7 = append(v7, -94)
    }
    var v14 float32 = (0.0 - (float32(1.5) - (1e10 * float32(1e10))))
    _ = v14
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]bool)
    var v3 uint8 = (v1 >> ((206) & 63))
    var v4 uint8 = (^(v1))
    var v15 int = 9223372036854775807
    v16 := []int{(v15 * (^(v15)))}
    v16 = append(v16, v15)
    var v17 int = 32767
    v4 -= (v3 % ((v4) | 1))
    if (!(-32768 < 32767)) {
        var v18 uint = 65535
        _ = v18
    } else {
        var v19 uint64 = 434
        var v20 uint16 = genFn0(1000, 0)
        var v21 uint16 = v20
        _ = v19
        _ = v20
        _ = v21
    }
    v1 = v4
    var v22 uint32 = 743
    v23 := GenS1{F0: 1, F1: 105, F2: "a"}
    var v24 uint32 = 65535
    var v25 int8 = -88
    if (v25 >= (v25 | (-(v25)))) {
        if (genFn0(-32768, 18446744073709551615) <= 0) {
            var v26 uint = 42
            _ = v26
        } else {
            if (42 >= 1) {
                for fi27 := 0; fi27 < 14; fi27++ {
                    if len(v16) > 0 {
                        v16[0] = 32768
                    }
                    _ = fi27
                }
                if true {
                    v28 := GenS1{F0: (v25 % (((v25 + (v25 ^ (v25 * int8(v1))))) | 1)), F1: 42, F2: ""}
                    _ = v28
                } else {
                    var v29 int32 = 2147483647
                    v16 = append(v16, (v17 / (((^(v15))) | 1)))
                    v30 := GenS1{F0: v25, F1: 310, F2: ""}
                    _ = v29
                    _ = v30
                }
                for fi31 := 0; fi31 < 7; fi31++ {
                    v17 += (v15 | fi31)
                    v16 = append(v16, (fi31 << ((255) & 63)))
                    _ = fi31
                }
            }
            switch (-(v15)) {
            case 2147483647:
                switch -760 {
                case -1:
                    if false {
                        v4 = 255
                    }
                case 285:
                    switch 9223372036854775807 {
                    case 1:
                        switch v15 {
                        case 9223372036854775807:
                            for fi32 := 0; fi32 < 17; fi32++ {
                                var v33 int64 = 1
                                _ = fi32
                                _ = v33
                            }
                        case -9223372036854775808:
                            switch ((("foo bar" + ("foo bar" + "\n\t")) + ("a" + (("\n\t" + ("\n\t" + (("a" + (("\n\t" + "\n\t") + (("" + ("foo bar" + ("hello" + "a"))) + ""))) + ("a" + ("a" + "\n\t"))))) + "a"))) + "foo bar") {
                            case "a":
                                _ = len(v16)
                            default:
                                var v34 uint = 995
                                _ = v34
                            }
                        }
                    case 9223372036854775807:
                        switch ("hello" + "hello") {
                        case "":
                            v35 := GenS1{F0: 95, F1: 30, F2: (("\n\t" + "a") + "")}
                            _ = v35
                        }
                    }
                case 127:
                    v36 := GenS1{F0: (v25 + (v25 ^ v25)), F1: 436, F2: "\n\t"}
                    _ = v36
                }
            case -485:
                _ = len(v2)
            }
            var v37 uint8 = 67
            _ = v37
        }
    } else {
        var v38 int32 = -1
        v25 = (^(v25))
        switch (-(v17)) {
        case -2147483648:
            fmt.Println(v3)
        }
        _ = v38
    }
    for fi39 := 0; fi39 < 10; fi39++ {
        var v40 int32 = 979
        v4 = (v3 % ((v1) | 1))
        switch -1 {
        case 42:
            var v41 bool = (0.0 == float64(1e10))
            _ = v41
        case 127:
            v16 = append(v16, (v17 << ((65535) & 63)))
        default:
            switch "" {
            case "hello":
                for fi42 := 0; fi42 < 8; fi42++ {
                    var v43 uint32 = 42
                    _ = fi42
                    _ = v43
                }
            default:
                v25 -= v25
            }
        }
        _ = fi39
        _ = v40
    }
    if false {
        for fi44 := 0; fi44 < 18; fi44++ {
            v22 += (v22 % ((v22) | 1))
            _ = fi44
        }
    }
    v17 -= v17
    delete(v2, ((("a" + "") + (("foo bar" + "") + "hello")) + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
