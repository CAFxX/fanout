package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool; F1 int }
func genFn0(p0 string, p1 bool, p2 uint8) uint {
    var v4 float32 = float32(((1.5 - 1e10) / 1.0))
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 float32 = v4
        v7 := GenS1{F0: 42}
        fmt.Println(v4, p0)
        _ = fi5
        _ = v6
        _ = v7
    }
    p1 = ((p1 || ((-1 != 100) && false)) || (false && p1))
    var v8 float32 = v4
    if p1 {
        v9 := []int{32768, -9223372036854775808}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return 1
}
func genFn1(p0 uint16) uint16 {
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 uint = 1
        fmt.Println(fi10)
        v12 := []int{32768}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 bool = (!true)
    if (uint8(p0) >= 2) {
        v14 := []int{9223372036854775807, 128, 0, 128}
        if (!v13) {
            _ = len(v14)
        } else {
            var v15 int8 = 42
            _ = v15
        }
        v13 = v13
        _ = v14
    } else {
        if (!false) {
            for fi16 := 0; fi16 < 9; fi16++ {
                fmt.Println(p0, v13)
                var v17 int64 = -1
                v18 := GenS1{F0: p0}
                _ = fi16
                _ = v17
                _ = v18
            }
            v13 = true
        } else {
            v19 := []int{}
            switch -304 {
            case 127:
                if false {
                    var v20 string = ""
                    _ = v20
                }
            case 707:
                var v21 uint64 = 4294967295
                _ = v21
            }
            _ = v19
        }
    }
    _ = p0
    _ = v13
    return (p0 & p0)
}

func main() {
    var v1 int16 = -640
    v2 := make(map[uint32]string)
    var v3 float32 = (((1e10 + -2.25) - -2.25) - (1.5 + 1e10))
    v22 := GenS1{F0: 32767}
    fmt.Println(v1, v3)
    var v23 int32 = -2147483648
    v23 = 0
    if (v3 == v3) {
        v23 = v23
        v23 += v23
        v24 := []int{287}
        _ = v24
    } else {
        var v25 uint32 = 0
        var v26 bool = (255 <= 65535)
        if ((v1 << ((0) & 63)) <= v1) {
            if (2 >= 42) {
                for fi27 := 0; fi27 < 1; fi27++ {
                    var v28 int8 = 49
                    v29 := GenS1{F0: 756}
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                v25 += (v25 + (v25 | v25))
            }
            v30 := GenS1{F0: 42}
            v31 := v2[(v25 ^ (v25 | v25))]
            _ = v30
            _ = v31
        }
        _ = v25
        _ = v26
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        fmt.Println(fi32)
        if (33 == 2) {
            for fi33 := 0; fi33 < 17; fi33++ {
                switch "foo bar" {
                case "hello":
                    fmt.Println(v23, v1, v23)
                case "\n\t":
                    var v34 int = 32768
                    _ = v34
                case "a":
                    v1 = int16(v23)
                default:
                    if false {
                        var v35 string = ((("" + "foo bar") + "a") + "foo bar")
                        _ = v35
                    }
                }
                var v36 uint32 = 1
                _ = fi33
                _ = v36
            }
            fmt.Println(v1)
            v37 := GenS2{F0: 4294967295, F1: 973}
            _ = v37
        }
        _ = fi32
    }
    v3 -= v3
    delete(v2, 2147483648)
    fmt.Println(v23, v1, v3)
    var v38 int32 = v23
    for fi39 := 0; fi39 < 5; fi39++ {
        for fi40 := 0; fi40 < 13; fi40++ {
            switch 127 {
            case 127:
                var v41 uint64 = 42
                _ = v41
            case -144:
                v42 := GenS1{F0: 65535}
                var v43 uint64 = 1
                _ = v42
                _ = v43
            case -2147483648:
                fmt.Println(v3)
                var v44 int = 4294967295
                v45 := v2[1]
                _ = v44
                _ = v45
            default:
                v3 += 1e10
            }
            v46 := GenS3{F0: false, F1: fi39}
            v47 := GenS3{F0: ((fi40 < fi40) && true), F1: fi40}
            _ = fi40
            _ = v46
            _ = v47
        }
        fmt.Println(v1)
        var v48 float64 = (3.14159 + ((float64(0.0) + ((3.14159 - 1.5) / (0.0 + 1.5))) / 1.5))
        _ = fi39
        _ = v48
    }
    v49 := []int{0, 611, 32768, 32767}
    v3 += v3
    switch ("a" + "") {
    case "\n\t":
        fmt.Println(v38, v23, v38)
    default:
        var v50 float32 = float32(0.0)
        _ = v50
    }
    v38 += (v38 >> ((18446744073709551615) & 63))
    switch ("\n\t" + ("\n\t" + "hello")) {
    case "a":
        fmt.Println(v23, v38)
    case "":
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v49:", len(v49), v49)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v38
    _ = v49
}
