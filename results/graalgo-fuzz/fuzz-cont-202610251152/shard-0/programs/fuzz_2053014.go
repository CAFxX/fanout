package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint; F1 uint64; F2 string }
type GenS4 struct { F0 int32; F1 uint64; F2 uint8 }
func genFn0() int32 {
    v7 := []int{0, 4294967295, 32767, -2147483648}
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: 748}
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 string = "foo bar"
            if (false || (((-1 > 0) || false) || true)) {
                v11 = ""
                v12 := GenS1{F0: -1}
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        var v13 int64 = 9223372036854775807
        _ = fi8
        _ = v13
        _ = v9
    }
    _ = v7
    return -1
}
func genFn1(p0 uint, p1 int32, p2 uint8) int8 {
    fmt.Println(p0, p1)
    if true {
        fmt.Println(p1, p0, p0)
        if false {
            for fi14 := 0; fi14 < 20; fi14++ {
                p2 -= p2
                p1 = (p1 & (p1 - 2147483647))
                _ = fi14
            }
            switch 32768 {
            case 9223372036854775807:
                v15 := GenS1{F0: int32(p2)}
                _ = v15
            case 32768:
                if true {
                    var v16 uint = p0
                    _ = v16
                } else {
                    p0 -= p0
                    var v17 float32 = 1e10
                    var v18 uint64 = 677
                    _ = v17
                    _ = v18
                }
            default:
                fmt.Println(p0, p2, p0)
            }
            p0 = (p0 >> (((p0 << (((^(p0))) & 63))) & 63))
        } else {
            p1 += (p1 * 42)
            switch -9223372036854775808 {
            case 32767:
                v19 := GenS1{F0: genFn0()}
                _ = v19
            case 9223372036854775807:
                var v20 int8 = 127
                _ = v20
            case 127:
                var v21 int16 = 196
                _ = v21
            }
        }
        var v22 string = "foo bar"
        _ = v22
    } else {
        if false {
            v23 := []int{452, 196}
            fmt.Println(p2, p2, p0)
            var v24 int8 = 0
            _ = v23
            _ = v24
        }
        p0 -= p0
        var v25 float32 = -2.25
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := []uint32{}
    v2 := []int{}
    var v3 int = -588
    v4 := []string{("hello" + "hello"), ("foo bar" + "hello")}
    v5 := make(map[int]string)
    var v6 uint = 42
    if true {
        for fi26 := 0; fi26 < 6; fi26++ {
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 float64 = 1e10
                v29 := GenS1{F0: 1}
                v3 = (v3 + fi27)
                var v30 int8 = genFn1(204, 42, 127)
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
            _ = fi26
        }
    }
    v3 = v3
    v1 = append(v1, 4294967295)
    _ = len(v4)
    v31 := GenS1{F0: 42}
    for fi32 := 0; fi32 < 20; fi32++ {
        _ = len(v5)
        if (!(3.14159 > ((float32(1e10) - -2.25) - (1.5 + 3.14159)))) {
            var v33 uint32 = 42
            v34 := GenS2{F0: genFn0(), F1: 2}
            _ = v33
            _ = v34
        }
        _ = fi32
    }
    v35 := v5[v3]
    for fi36 := 0; fi36 < 17; fi36++ {
        if (v6 < (v6 + (v6 & v6))) {
            var v37 uint16 = 42
            switch ("\n\t" + "") {
            case "\n\t":
                var v38 int8 = 100
                _ = v38
            case "hello":
                for fi39 := 0; fi39 < 8; fi39++ {
                    fmt.Println(fi36, v6, fi39)
                    v35 = "foo bar"
                    fmt.Println(fi39, fi39, v35)
                    _ = fi39
                }
            case "a":
                fmt.Println(v3, v3, v3)
            }
            var v40 int = (v3 & fi36)
            _ = v37
            _ = v40
        } else {
            switch v3 {
            case -77:
                for fi41 := 0; fi41 < 11; fi41++ {
                    if (true && (42 == 0)) {
                        v35 = (v35 + "")
                        v35 = "a"
                        v6 = (v6 ^ v6)
                    } else {
                        fmt.Println(fi41)
                        v6 += (v6 + v6)
                    }
                    v42 := GenS3{F0: (v6 / (((v6 % ((65535) | 1))) | 1)), F1: 2, F2: (v35 + (v35 + (v35 + "")))}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi41
                    _ = v42
                }
            case -302:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            v43 := GenS4{F0: -783, F1: 18446744073709551615, F2: 95}
            _ = v43
        }
        v6 = 23
        for fi44 := 0; fi44 < 7; fi44++ {
            delete(v5, 1)
            _ = fi44
        }
        _ = fi36
    }
    var v45 uint8 = 1
    if (false || true) {
        _ = len(v1)
        v35 = v35
    }
    v5[(v3 << ((v6) & 63))] = v35
    switch ((("" + (v35 + v35)) + "foo bar") + "") {
    case "foo bar":
        switch -1 {
        case -1:
            _ = len(v1)
        case 4294967295:
            var v46 string = v35
            _ = v46
        case 42:
            fmt.Println(v3)
        }
    default:
        v6 += (v6 << ((2) & 63))
    }
    var v47 int32 = 245
    for fi48 := 0; fi48 < 13; fi48++ {
        v3 = 127
        v5[(fi48 | 830)] = v35
        for fi49 := 0; fi49 < 13; fi49++ {
            switch int(v6) {
            case 2147483647:
                if len(v1) > 0 {
                    v1[0] = 0
                }
            }
            if (false && (400 >= 1)) {
                var v50 float64 = 1.5
                _ = len(v5)
                _ = v50
            } else {
                v51 := GenS1{F0: (^(v47))}
                _ = v51
            }
            _ = fi49
        }
        _ = fi48
    }
    v3 = (v3 / ((v3) | 1))
    var v52 int8 = 1
    if true {
        for fi53 := 0; fi53 < 16; fi53++ {
            v35 = v35
            v54 := v5[-9223372036854775808]
            _ = fi53
            _ = v54
        }
        for fi55 := 0; fi55 < 16; fi55++ {
            var v56 uint64 = 370
            _ = fi55
            _ = v56
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v45
    _ = v47
    _ = v5
    _ = v52
    _ = v6
}
