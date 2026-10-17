package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int64; F2 int8 }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint) int8 {
    for fi4 := 0; fi4 < 17; fi4++ {
        for fi5 := 0; fi5 < 14; fi5++ {
            p0 -= p0
            p0 += p0
            _ = fi5
        }
        _ = fi4
    }
    var v6 float64 = (1.5 / 3.14159)
    _ = p0
    _ = v6
    return -79
}
func genFn1() int64 {
    var v7 int32 = 1
    v8 := GenS1{F0: 1000, F1: -32768, F2: 2}
    v9 := []int{127}
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]int)
    v3 := []string{"\n\t", ("a" + ""), ("a" + (("foo bar" + "") + "hello")), "a"}
    for fi10 := 0; fi10 < 8; fi10++ {
        v1 -= (v1 - (v1 >> ((42) & 63)))
        _ = fi10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        v1 = v1
        switch "hello" {
        case "":
            switch (fi11 % (((fi11 - (-(fi11)))) | 1)) {
            case 128:
                for fi12 := 0; fi12 < 5; fi12++ {
                    v13 := GenS1{F0: -236, F1: 32767, F2: 586}
                    if (2 >= 4294967295) {
                        v1 = v1
                        fmt.Println(fi12)
                        _ = len(v3)
                        v1 -= v1
                        _ = len(v2)
                    }
                    v14 := GenS1{F0: -970, F1: 0, F2: 65535}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            case 0:
                var v15 bool = false
                _ = v15
            case 251:
                var v16 int = (fi11 - fi11)
                _ = v16
            }
        case "foo bar":
            v17 := GenS2{F0: 42, F1: (0 == -128)}
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            v1 = v1
            var v19 int = fi11
            _ = fi18
            _ = v19
        }
        _ = fi11
    }
    var v20 uint32 = 2147483648
    for fi21 := 0; fi21 < 12; fi21++ {
        v1 += v1
        switch "hello" {
        case "hello":
            if true {
                v22 := GenS3{F0: 18446744073709551615}
                _ = v22
            }
        case "\n\t":
            var v23 uint = 2
            _ = v23
        case "a":
            if true {
                v3 = append(v3, "\n\t")
                v24 := GenS4{F0: "", F1: -833, F2: 100}
                v3 = append(v3, ("\n\t" + ("hello" + "hello")))
                _ = v24
            }
        default:
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        }
        if false {
            v3 = append(v3, ("foo bar" + "\n\t"))
        } else {
            v20 = (v20 & (v20 + (v20 * 2)))
        }
        _ = fi21
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        switch "foo bar" {
        case "hello":
            v1 = 121
        case "\n\t":
            for fi26 := 0; fi26 < 18; fi26++ {
                if true {
                    v20 -= (v20 / (((^(v20))) | 1))
                }
                v27 := GenS5{F0: 415}
                var v28 bool = (0.0 < 3.14159)
                _ = fi26
                _ = v27
                _ = v28
            }
        case "":
            for fi29 := 0; fi29 < 17; fi29++ {
                var v30 uint16 = 0
                var v31 int32 = 0
                var v32 string = "\n\t"
                v30 -= 300
                v33 := GenS1{F0: 648, F1: 271, F2: uint32(v1)}
                if (int8(v31) != 127) {
                    fmt.Println(v30)
                }
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
                _ = v33
            }
        }
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi34 := 0; fi34 < 16; fi34++ {
                v35 := GenS5{F0: -123}
                _ = fi34
                _ = v35
            }
            if false {
                v36 := GenS2{F0: 32767, F1: (false || (true || (((false || ((243 == 4294967295) && true)) && (2 != 505)) && (!true))))}
                _ = v36
            }
        }
        _ = fi25
    }
    var v37 int = -9223372036854775808
    fmt.Println(v1)
    var v38 int = (^(v37))
    v39 := GenS3{F0: 4294967295}
    var v40 int = (v38 % ((128) | 1))
    var v41 uint8 = v1
    v1 = 179
    for fi42 := 0; fi42 < 2; fi42++ {
        v37 = (v40 << ((312) & 63))
        var v43 uint32 = 123
        var v44 uint32 = 4294967295
        _ = fi42
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v41
}
