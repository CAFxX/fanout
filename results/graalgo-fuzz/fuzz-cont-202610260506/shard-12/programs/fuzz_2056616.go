package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int64; F2 uint }
type GenS3 struct { F0 bool; F1 uint; F2 int16 }
func genFn0(p0 int, p1 uint16, p2 bool) bool {
    v6 := GenS1{F0: "hello"}
    p1 = p1
    if ("\n\t" <= "\n\t") {
        p2 = (!(!false))
    } else {
        for fi7 := 0; fi7 < 7; fi7++ {
            switch (fi7 % ((-1) | 1)) {
            case 1:
                p1 -= 65535
            case -989:
                p1 = (p1 << ((2) & 63))
            case -2147483648:
                var v8 int16 = 1
                _ = v8
            }
            _ = fi7
        }
        var v9 bool = (-9223372036854775808 != 9223372036854775807)
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (1 != 1)
}
func genFn1(p0 uint8) uint64 {
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := GenS2{F0: -9223372036854775808, F1: 42, F2: 0}
        _ = fi10
        _ = v11
    }
    p0 -= 0
    for fi12 := 0; fi12 < 4; fi12++ {
        v13 := GenS2{F0: 9223372036854775807, F1: 0, F2: 65535}
        p0 = 42
        _ = fi12
        _ = v13
    }
    if (!false) {
        var v14 int = 127
        p0 = (p0 + p0)
        _ = v14
    } else {
        v15 := GenS2{F0: -143, F1: -58, F2: 2}
        v16 := GenS2{F0: -1, F1: -656, F2: 18446744073709551615}
        _ = v15
        _ = v16
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{("foo bar" + "\n\t"), "foo bar", "a", "foo bar"}
    var v3 uint16 = 496
    v4 := make(map[uint32]string)
    var v5 uint = 0
    for fi17 := 0; fi17 < 14; fi17++ {
        if true {
            v3 -= (v3 - (v3 ^ (^(v3))))
            switch "hello" {
            case "hello":
                v18 := GenS2{F0: -833, F1: -1, F2: 854}
                _ = v18
            case "a":
                for fi19 := 0; fi19 < 14; fi19++ {
                    var v20 int = fi17
                    fmt.Println(v20, v20, v5)
                    _ = fi19
                    _ = v20
                }
            case "foo bar":
                var v21 int8 = 100
                _ = v21
            }
            var v22 uint64 = 104
            _ = v22
        }
        var v23 uint8 = 1
        var v24 int = 127
        _ = fi17
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        _ = len(v1)
        for fi26 := 0; fi26 < 14; fi26++ {
            var v27 int16 = -32768
            fmt.Println(v27)
            _ = fi26
            _ = v27
        }
        var v28 float64 = 3.14159
        _ = fi25
        _ = v28
    }
    if true {
        var v29 int8 = 109
        _ = v29
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3)
    if true {
        fmt.Println(v3)
    } else {
        if true {
            v5 = (v5 % ((42) | 1))
            switch 4294967295 {
            case -427:
                v30 := GenS2{F0: -9223372036854775808, F1: -1, F2: v5}
                _ = v30
            }
        }
        if true {
            var v31 int32 = 42
            v3 = 2
            v1[-1] = 411
            _ = v31
        } else {
            switch "foo bar" {
            case "a":
                var v32 float64 = 1.5
                _ = v32
            case "":
                switch ("foo bar" + "a") {
                case "\n\t":
                    var v33 int32 = -2147483648
                    _ = v33
                case "":
                    var v34 string = "hello"
                    _ = v34
                default:
                    v5 = 952
                }
            case "\n\t":
                if false {
                    fmt.Println(v3, v3)
                    fmt.Println(v3, v5)
                } else {
                    var v35 uint = 0
                    var v36 float32 = 3.14159
                    if true {
                        v36 -= v36
                    }
                    _ = v35
                    _ = v36
                }
            }
        }
        v3 -= (v3 + v3)
    }
    var v37 float64 = 1e10
    delete(v4, 382)
    for fi38 := 0; fi38 < 2; fi38++ {
        var v39 uint16 = (v3 - 392)
        _ = fi38
        _ = v39
    }
    if ((^(v3)) <= v3) {
        v37 = ((v37 + 3.14159) - ((float64((1e10 + 1.5)) * (-2.25 * -2.25)) + v37))
        v40 := GenS3{F0: ((((717 >= 0) && false) || ((v3 | 398) >= (v3 * (v3 - v3)))) && (true && ((65535 >= 2) || true))), F1: 65535, F2: 42}
        _ = v40
    }
    v2 = append(v2, ("hello" + ("a" + ("hello" + "foo bar"))))
    var v41 uint16 = (v3 >> (((v5 % (((v5 - (v5 << ((uint(v3)) & 63)))) | 1))) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v5
}
