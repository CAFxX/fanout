package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint, p1 int32) uint {
    v5 := []int{}
    var v6 string = "a"
    switch v6 {
    case "a":
        p1 = -1
    case "foo bar":
        var v7 string = (v6 + "")
        _ = v7
    case "hello":
        var v8 float64 = (1.5 - 3.14159)
        _ = v8
    }
    var v9 int32 = p1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v9
    return (p0 & 65535)
}
func genFn1(p0 int8, p1 uint8, p2 int32) int32 {
    if false {
        for fi10 := 0; fi10 < 5; fi10++ {
            v11 := []int{}
            _ = fi10
            _ = v11
        }
    }
    if (0 > -32768) {
        var v12 int32 = p2
        _ = v12
    } else {
        v13 := []int{4294967295, 42, 128}
        _ = v13
    }
    var v14 float64 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return (p2 << ((135) & 63))
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]bool)
    v3 := []string{""}
    var v4 float64 = 1e10
    var v15 int = -2147483648
    var v16 int16 = 32767
    v16 -= (v16 & v16)
    if (-1 < -128) {
        v17 := GenS1{F0: "foo bar", F1: -9223372036854775808, F2: 255}
        v18 := GenS1{F0: "hello", F1: 9223372036854775807, F2: 197}
        _ = len(v3)
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 string = ("foo bar" + "")
            if len(v3) > 0 {
                v3[0] = v20
            }
            _ = fi19
            _ = v20
        }
        v1 = genFn1(42, 6, (^(v1)))
    }
    var v21 uint8 = 128
    switch "\n\t" {
    case "a":
        v15 -= v15
    case "foo bar":
        switch (("hello" + ((("\n\t" + "\n\t") + "") + (("hello" + "hello") + "hello"))) + "\n\t") {
        case "a":
            var v22 uint = 42
            _ = v22
        case "":
            var v23 int64 = 498
            _ = v23
        default:
            v1 -= 0
        }
    case "hello":
        var v24 int = (v15 % ((v15) | 1))
        _ = v24
    default:
        for fi25 := 0; fi25 < 19; fi25++ {
            if (false || (false || true)) {
                v26 := GenS1{F0: "a", F1: 330, F2: 1}
                v27 := GenS1{F0: "", F1: 1, F2: 0}
                v28 := GenS1{F0: "foo bar", F1: 600, F2: 1}
                _ = v26
                _ = v27
                _ = v28
            } else {
                v1 = (-(v1))
                var v29 string = "foo bar"
                v2[fi25] = (true || (!true))
                _ = v29
            }
            v21 += (v21 << ((4294967295) & 63))
            v30 := GenS2{F0: (v21 | v21), F1: 680}
            _ = fi25
            _ = v30
        }
    }
    var v31 int8 = -128
    for fi32 := 0; fi32 < 12; fi32++ {
        for fi33 := 0; fi33 < 6; fi33++ {
            var v34 uint32 = 2147483648
            _ = fi33
            _ = v34
        }
        _ = fi32
    }
    v35 := GenS3{F0: 42, F1: v15}
    switch (v15 ^ v15) {
    case -9223372036854775808:
        for fi36 := 0; fi36 < 17; fi36++ {
            _ = len(v3)
            fmt.Println(v31, v21)
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
            _ = fi36
        }
    case -1:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v37 uint8 = v21
        _ = v37
    }
    if false {
        for fi38 := 0; fi38 < 1; fi38++ {
            if false {
                v16 = (v16 | (v16 | v16))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                if true {
                    for fi39 := 0; fi39 < 18; fi39++ {
                        v40 := GenS2{F0: (^(v21)), F1: 0}
                        _ = fi39
                        _ = v40
                    }
                }
            }
            switch (fi38 ^ (fi38 >> ((4294967295) & 63))) {
            case -265:
                v41 := GenS2{F0: 1, F1: 995}
                _ = v41
            default:
                var v42 bool = ((fi38 == -2147483648) && false)
                _ = v42
            }
            _ = fi38
        }
        var v43 uint16 = 32767
        _ = v43
    }
    switch -168 {
    case 128:
        var v44 int = (v15 >> ((36) & 63))
        _ = v44
    }
    v45 := GenS4{F0: 1, F1: v1}
    fmt.Println(v1, v4)
    fmt.Println(v15, v16)
    v1 = v1
    switch (^(v15)) {
    case 127:
        if (true && true) {
            var v46 int64 = -380
            var v47 int16 = int16(v46)
            _ = len(v2)
            _ = v46
            _ = v47
        } else {
            for fi48 := 0; fi48 < 4; fi48++ {
                fmt.Println(v31, fi48)
                var v49 float64 = (v4 - (0.0 * 1.5))
                for fi50 := 0; fi50 < 18; fi50++ {
                    v51 := GenS1{F0: ("" + ("\n\t" + ("foo bar" + ("a" + "a")))), F1: 0, F2: 101}
                    var v52 int16 = v16
                    v52 = (^(v52))
                    _ = fi50
                    _ = v51
                    _ = v52
                }
                _ = fi48
                _ = v49
            }
            fmt.Println(v21)
            if (((1 < 32767) && true) || true) {
                for fi53 := 0; fi53 < 3; fi53++ {
                    v2[(-(fi53))] = true
                    var v54 uint = 1
                    _ = fi53
                    _ = v54
                }
            } else {
                fmt.Println(v21, v4, v21)
                _ = len(v3)
                var v55 int8 = -74
                _ = v55
            }
        }
    case 42:
        v4 += -2.25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v45
}
