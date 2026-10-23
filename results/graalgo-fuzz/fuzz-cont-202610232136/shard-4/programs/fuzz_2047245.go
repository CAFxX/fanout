package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int16; F2 uint16 }
type GenS3 struct { F0 string }
func genFn0(p0 int8, p1 int, p2 int) uint32 {
    fmt.Println(p0)
    v4 := GenS1{F0: 0}
    var v5 int64 = -9223372036854775808
    var v6 float32 = 1e10
    v7 := GenS1{F0: 255}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1() int32 {
    if (270 == 885) {
        var v8 uint32 = 42
        for fi9 := 0; fi9 < 10; fi9++ {
            v8 = 0
            _ = fi9
        }
        _ = v8
    } else {
        var v10 int = 32767
        v11 := GenS2{F0: (v10 / ((v10) | 1)), F1: 1000, F2: 363}
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 32767, F1: -88, F2: 468}
    _ = v12
    return -683
}

func main() {
    v1 := []string{"a", "foo bar", "foo bar", "hello"}
    var v2 uint32 = 822
    v3 := make(map[uint32]bool)
    if (1 <= 527) {
        if (false && ((9223372036854775807 == 324) && (v2 <= 65535))) {
            switch "\n\t" {
            case "":
                v13 := GenS2{F0: 127, F1: 42, F2: 0}
                _ = v13
            case "hello":
                if (1 != 2) {
                    var v14 int32 = 2147483647
                    _ = v14
                } else {
                    v15 := GenS3{F0: ""}
                    _ = v15
                }
            }
        } else {
            var v16 uint16 = 1
            v2 -= 4294967295
            _ = v16
        }
        var v17 int16 = 0
        _ = v17
    } else {
        switch 1 {
        case 1:
            var v18 int32 = 620
            _ = v18
        case -2147483648:
            v19 := GenS3{F0: "a"}
            _ = v19
        default:
            for fi20 := 0; fi20 < 5; fi20++ {
                for fi21 := 0; fi21 < 17; fi21++ {
                    var v22 uint = 843
                    switch (fi21 << ((4294967295) & 63)) {
                    case -1:
                        var v23 uint32 = v2
                        _ = v23
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
        }
        var v24 uint32 = genFn0(87, 611, 2147483647)
        _ = v24
    }
    v2 -= v2
    v25 := GenS2{F0: -1, F1: 785, F2: 886}
    if true {
        switch -959 {
        case 32768:
            _ = len(v3)
        case 374:
            v2 += 2147483647
        case 32767:
            if len(v1) > 0 {
                v1[0] = "foo bar"
            }
        default:
            for fi26 := 0; fi26 < 14; fi26++ {
                var v27 int = 128
                fmt.Println(fi26, v27, v27)
                v2 -= v2
                _ = fi26
                _ = v27
            }
        }
        if ((42 < 42) || true) {
            _ = len(v1)
        } else {
            for fi28 := 0; fi28 < 2; fi28++ {
                var v29 bool = (true || true)
                var v30 string = ("a" + ("\n\t" + "hello"))
                _ = len(v3)
                _ = fi28
                _ = v29
                _ = v30
            }
        }
    }
    v2 -= (v2 ^ (v2 >> ((707) & 63)))
    v2 = genFn0(0, -326, -300)
    v2 += v2
    fmt.Println(v2)
    switch "foo bar" {
    case "\n\t":
        if (!(976 > 4294967295)) {
            v3[(v2 * (v2 & (^(v2))))] = false
        } else {
            v2 = 2147483648
        }
    case "foo bar":
        if false {
            switch 2147483647 {
            case 127:
                var v31 float64 = (0.1 + (0.1 + 3.14159))
                _ = v31
            case 172:
                if (!((42 == -79) || false)) {
                    v2 -= (v2 | genFn0(100, 32767, 0))
                    if true {
                        var v32 float32 = 1.5
                        v2 -= genFn0(127, -620, 389)
                        var v33 int16 = 42
                        _ = v32
                        _ = v33
                    }
                }
            case -836:
                v2 += (v2 * (v2 << ((18446744073709551615) & 63)))
            default:
                var v34 int16 = 1
                _ = v34
            }
            if true {
                var v35 float64 = 1.5
                v36 := GenS3{F0: "hello"}
                _ = v35
                _ = v36
            }
            fmt.Println(v2)
        } else {
            v37 := GenS3{F0: "a"}
            v2 -= 0
            if ((!(!false)) && (0 >= 429)) {
                v2 = v2
            }
            _ = v37
        }
    case "":
        for fi38 := 0; fi38 < 1; fi38++ {
            v2 += v2
            for fi39 := 0; fi39 < 11; fi39++ {
                var v40 bool = true
                _ = fi39
                _ = v40
            }
            _ = fi38
        }
    }
    v2 += (v2 >> ((42) & 63))
    _ = len(v1)
    v2 = v2
    if true {
        for fi41 := 0; fi41 < 6; fi41++ {
            if true {
                v2 -= (v2 | 114)
            } else {
                var v42 int8 = 127
                _ = v42
            }
            var v43 uint64 = 684
            _ = fi41
            _ = v43
        }
    }
    if false {
        v2 += (v2 ^ v2)
        v44 := GenS3{F0: ("\n\t" + (("\n\t" + ("hello" + "hello")) + "\n\t"))}
        var v45 string = (("" + "") + "foo bar")
        _ = v44
        _ = v45
    }
    if (!true) {
        for fi46 := 0; fi46 < 10; fi46++ {
            var v47 bool = false
            _ = fi46
            _ = v47
        }
    } else {
        v2 -= (^(v2))
        var v48 uint8 = 1
        if false {
            v2 -= (v2 >> ((132) & 63))
        }
        _ = v48
    }
    fmt.Println(v2)
    v49 := GenS3{F0: ""}
    var v50 int64 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v49
    _ = v50
}
