package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 int16; F2 uint }
func genFn0(p0 uint8, p1 int64, p2 int) uint {
    for fi7 := 0; fi7 < 13; fi7++ {
        if false {
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 float64 = 3.14159
                v10 := GenS1{F0: "foo bar", F1: 359}
                var v11 uint = 1
                var v12 string = ("hello" + "a")
                switch v12 {
                case "foo bar":
                    var v13 int = -9223372036854775808
                    _ = v13
                case "hello":
                    v14 := GenS1{F0: "a", F1: 2147483648}
                    _ = v14
                default:
                    p1 -= (p1 / (((p1 >> (((v11 * v11)) & 63))) | 1))
                }
                v15 := []int{(fi7 ^ (fi7 % ((p2) | 1)))}
                _ = fi8
                _ = v10
                _ = v11
                _ = v12
                _ = v15
                _ = v9
            }
            v16 := GenS2{F0: 42}
            _ = v16
        } else {
            if false {
                for fi17 := 0; fi17 < 7; fi17++ {
                    p2 -= fi7
                    p2 += (fi17 - -9223372036854775808)
                    _ = fi17
                }
            }
            var v18 uint = 717
            var v19 uint8 = 163
            _ = v18
            _ = v19
        }
        v20 := GenS3{F0: p1, F1: 35, F2: 42}
        if false {
            var v21 int64 = (-(p1))
            v22 := []int{fi7, fi7, fi7}
            switch (fi7 & p2) {
            case 316:
                v23 := GenS1{F0: "hello", F1: 527}
                _ = v23
            case 4294967295:
                p2 -= (fi7 >> ((42) & 63))
            }
            _ = v21
            _ = v22
        }
        _ = fi7
        _ = v20
    }
    p1 -= 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    return 275
}
func genFn1(p0 int64, p1 uint32, p2 int64) int {
    var v24 int64 = -9223372036854775808
    var v25 int64 = (v24 - (p2 / (((-(v24))) | 1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    return 127
}

func main() {
    var v1 int32 = -2147483648
    v2 := []int{-9223372036854775808}
    var v3 int32 = 26
    var v4 int16 = -32768
    var v5 int = 0
    v6 := []string{}
    if ((^(v5)) < -9223372036854775808) {
        if true {
            for fi26 := 0; fi26 < 1; fi26++ {
                v2 = append(v2, 4294967295)
                v5 += (^(fi26))
                for fi27 := 0; fi27 < 11; fi27++ {
                    var v28 string = "foo bar"
                    switch fi27 {
                    case -986:
                        var v29 uint = 655
                        _ = v29
                    }
                    fmt.Println(fi26, v3, v1)
                    if ((-2.25 == -2.25) && (!true)) {
                        var v30 uint32 = 2
                        v5 = (fi27 % ((fi27) | 1))
                        var v31 bool = true
                        _ = v30
                        _ = v31
                    }
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
            for fi32 := 0; fi32 < 7; fi32++ {
                var v33 bool = (4294967295 < 2)
                _ = fi32
                _ = v33
            }
            v3 = (-(v1))
        } else {
            if len(v2) > 0 {
                v2[0] = (v5 / ((1) | 1))
            }
        }
        v5 += (v5 % (((^(v5))) | 1))
        v34 := GenS3{F0: 42, F1: (v4 % ((v4) | 1)), F2: 18446744073709551615}
        _ = v34
    } else {
        v4 = v4
    }
    var v35 uint32 = 2147483647
    v36 := GenS1{F0: ("\n\t" + "a"), F1: 2}
    var v37 uint16 = 2
    v5 += (v5 % ((v5) | 1))
    if false {
        v35 -= 65535
    }
    var v38 uint8 = 2
    switch ("hello" + ((("a" + "\n\t") + "\n\t") + "a")) {
    case "foo bar":
        var v39 int32 = 1
        _ = v39
    case "a":
        var v40 float32 = (((1.5 + 1.5) / (0.0 - 1e10)) - 1e10)
        _ = v40
    default:
        fmt.Println(v1, v5, v37)
    }
    for fi41 := 0; fi41 < 1; fi41++ {
        var v42 bool = true
        _ = fi41
        _ = v42
    }
    v43 := GenS1{F0: ("a" + "a"), F1: (v35 % (((^(v35))) | 1))}
    v44 := GenS1{F0: "", F1: (v35 << ((2) & 63))}
    var v45 float64 = 1e10
    for fi46 := 0; fi46 < 1; fi46++ {
        v47 := GenS2{F0: int16(fi46)}
        _ = fi46
        _ = v47
    }
    switch (v5 - v5) {
    case 127:
        if len(v6) > 0 {
            v6[0] = "hello"
        }
    default:
        fmt.Println(v35, v5)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi48 := 0; fi48 < 2; fi48++ {
        v37 += v37
        _ = fi48
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
