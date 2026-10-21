package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
func genFn0() uint64 {
    v7 := []int{}
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: ("a" + "hello")}
        switch fi8 {
        case 1:
            fmt.Println(fi8)
        case 32767:
            var v10 int32 = -2147483648
            _ = v10
        }
        _ = len(v7)
        _ = fi8
        _ = v9
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(fi11)
        _ = len(v7)
        _ = fi11
    }
    var v12 int = -373
    _ = v12
    _ = v7
    return 567
}

func main() {
    var v1 int8 = 112
    var v2 string = "a"
    var v3 uint16 = 32767
    var v4 int32 = -1
    v5 := []uint32{}
    v6 := make(map[string]string)
    v1 += v1
    var v13 uint32 = 4294967295
    switch v2 {
    case "foo bar":
        if (-2.25 == 1e10) {
            for fi14 := 0; fi14 < 17; fi14++ {
                for fi15 := 0; fi15 < 19; fi15++ {
                    v16 := GenS2{F0: 2}
                    var v17 uint16 = (v3 & 32767)
                    switch fi15 {
                    case -935:
                        fmt.Println(fi14, fi14, v13)
                    case 42:
                        var v18 int64 = 1
                        _ = v18
                    default:
                        var v19 uint16 = uint16(v13)
                        _ = v19
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                if true {
                    fmt.Println(fi14)
                } else {
                    var v20 float32 = (0.1 + ((0.1 - 3.14159) - 0.0))
                    v21 := v6[v2]
                    _ = v20
                    _ = v21
                }
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 string = (("foo bar" + v2) + (((v2 + v2) + "\n\t") + v2))
                    _ = fi22
                    _ = v23
                }
                _ = fi14
            }
            var v24 uint8 = 245
            _ = v24
        } else {
            var v25 uint64 = 1
            v4 -= (v4 - (v4 >> ((0) & 63)))
            var v26 int = 32768
            _ = v25
            _ = v26
        }
    case "hello":
        var v27 int16 = 32767
        _ = v27
    case "":
        for fi28 := 0; fi28 < 12; fi28++ {
            var v29 int32 = v4
            v29 -= v29
            _ = fi28
            _ = v29
        }
    }
    switch v2 {
    case "":
        var v30 int32 = 42
        _ = v30
    case "hello":
        switch v2 {
        case "foo bar":
            switch ("\n\t" + v2) {
            case "foo bar":
                var v31 int8 = v1
                _ = v31
            default:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        }
    case "a":
        var v32 bool = (!true)
        _ = v32
    }
    v33 := GenS1{F0: (v2 + "foo bar")}
    v4 -= 0
    var v34 uint16 = v3
    switch 42 {
    case -226:
        if (v4 < -1) {
            var v35 string = (v2 + (((v2 + v2) + ("hello" + v2)) + "\n\t"))
            v35 = ("\n\t" + v2)
            v36 := GenS1{F0: v2}
            _ = v35
            _ = v36
        } else {
            var v37 uint32 = (v13 + 462)
            fmt.Println(v1, v3)
            _ = v37
        }
    case 637:
        v34 -= v3
    }
    if false {
        var v38 int = -2147483648
        switch ((v2 + (v2 + "")) + v2) {
        case "a":
            if false {
                v38 = v38
                fmt.Println(v38, v4)
                v38 = int(v13)
            } else {
                var v39 int32 = v4
                switch "foo bar" {
                case "hello":
                    if true {
                        v39 = (v39 * v39)
                    }
                case "foo bar":
                    if true {
                        v34 += v3
                        var v40 int32 = v4
                        var v41 int64 = 9223372036854775807
                        _ = v40
                        _ = v41
                    }
                default:
                    for fi42 := 0; fi42 < 12; fi42++ {
                        var v43 int32 = v4
                        var v44 int32 = (v4 % (((-(v4))) | 1))
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                }
                v38 += (v38 >> ((18446744073709551615) & 63))
                _ = v39
            }
        case "foo bar":
            v34 = uint16(v13)
        case "\n\t":
            v45 := GenS2{F0: 96}
            _ = v45
        }
        if ((-(v1)) < 0) {
            switch 127 {
            case 4294967295:
                v34 = v34
            }
        } else {
            if ((false && (!true)) || true) {
                v38 -= int(v13)
                v46 := GenS1{F0: "foo bar"}
                for fi47 := 0; fi47 < 6; fi47++ {
                    v4 -= v4
                    v48 := GenS2{F0: genFn0()}
                    v49 := GenS1{F0: (v2 + v2)}
                    _ = fi47
                    _ = v48
                    _ = v49
                }
                _ = v46
            }
        }
        _ = v38
    } else {
        v50 := GenS1{F0: v2}
        v51 := GenS2{F0: genFn0()}
        _ = v50
        _ = v51
    }
    fmt.Println(v3)
    v13 -= (v13 % ((692) | 1))
    var v52 int16 = 32767
    for fi53 := 0; fi53 < 12; fi53++ {
        var v54 uint32 = (v13 * 65535)
        _ = fi53
        _ = v54
    }
    if len(v5) > 0 {
        v5[0] = v13
    }
    var v55 int32 = int32(v52)
    v4 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v52
    _ = v55
    _ = v6
}
