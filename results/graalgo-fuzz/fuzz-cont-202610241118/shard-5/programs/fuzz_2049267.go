package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint; F1 uint64 }
func genFn0() uint {
    var v5 float64 = float64(1.5)
    v6 := []int{}
    v5 += v5
    for fi7 := 0; fi7 < 20; fi7++ {
        switch 32767 {
        case 1:
            if true {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v8 bool = false
                _ = v8
            }
        case 42:
            var v9 bool = (!false)
            _ = v9
        }
        v10 := GenS1{F0: (-(fi7)), F1: 2}
        _ = fi7
        _ = v10
    }
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 uint64 = 136
    var v2 uint8 = 33
    v3 := []string{"hello", "a"}
    v4 := make(map[int]bool)
    v2 = (v2 | 69)
    if (65535 <= uint16(v1)) {
        if (18446744073709551615 == genFn0()) {
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := GenS2{F0: 538, F1: 4294967295, F2: 42}
                _ = fi11
                _ = v12
            }
        }
        var v13 bool = (true || (!(3.14159 <= 1e10)))
        _ = v13
    } else {
        v3 = append(v3, "\n\t")
        var v14 uint64 = 2
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            switch (fi15 % ((fi15) | 1)) {
            case 4294967295:
                v2 -= v2
            case 873:
                v2 = (v2 | v2)
            default:
                v2 = v2
            }
            v4[(fi16 - (fi15 & (-(fi16))))] = ((fi16 * (fi15 % (((fi16 % (((fi16 + (^(fi15)))) | 1))) | 1))) > -537)
            if (!(true && (1 >= 65535))) {
                switch ("hello" + "hello") {
                case "\n\t":
                    if (("" + "a") != "\n\t") {
                        _ = len(v4)
                        v1 -= v1
                        v4[(fi15 * -1)] = ((true && true) && true)
                    } else {
                        v3 = append(v3, ("foo bar" + ("hello" + "foo bar")))
                    }
                }
            } else {
                v1 += (^(v1))
                fmt.Println(fi15)
            }
            _ = fi16
        }
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 int32 = 2147483647
            _ = fi17
            _ = v18
        }
        _ = fi15
    }
    v19 := GenS1{F0: 2147483647, F1: 510}
    if true {
        v20 := GenS3{F0: v1, F1: 833}
        var v21 string = ("a" + ("\n\t" + ""))
        if len(v3) > 0 {
            v3[0] = (("hello" + "hello") + (v21 + ("a" + v21)))
        }
        _ = v20
        _ = v21
    }
    if false {
        for fi22 := 0; fi22 < 20; fi22++ {
            v23 := GenS4{F0: 255, F1: 2}
            if false {
                var v24 string = "foo bar"
                fmt.Println(v24)
                _ = v24
            } else {
                if (!(-1 >= 42)) {
                    var v25 uint8 = v2
                    _ = v25
                }
            }
            _ = fi22
            _ = v23
        }
        v26 := GenS1{F0: 1, F1: 2}
        for fi27 := 0; fi27 < 3; fi27++ {
            if false {
                var v28 uint16 = 20
                _ = v28
            } else {
                var v29 string = "foo bar"
                var v30 bool = true
                _ = v29
                _ = v30
            }
            _ = fi27
        }
        _ = v26
    } else {
        v1 = v1
        for fi31 := 0; fi31 < 11; fi31++ {
            v32 := GenS1{F0: fi31, F1: uint(v1)}
            var v33 uint64 = (v1 % ((v1) | 1))
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    delete(v4, 42)
    var v34 uint = 42
    var v35 uint64 = (v1 % ((v1) | 1))
    switch "hello" {
    case "hello":
        for fi36 := 0; fi36 < 16; fi36++ {
            switch "foo bar" {
            case "\n\t":
                v4[fi36] = (!true)
            default:
                v1 = 424
            }
            var v37 float32 = -2.25
            _ = fi36
            _ = v37
        }
    case "foo bar":
        v34 += (v34 >> (((v34 - v34)) & 63))
    default:
        var v38 float64 = 0.1
        _ = v38
    }
    v39 := GenS3{F0: (v1 - v35), F1: 304}
    var v40 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v40
}
