package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int16; F1 uint8; F2 uint }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint8) int8 {
    var v5 int32 = 42
    var v6 int8 = -97
    _ = p0
    _ = v5
    _ = v6
    return -1
}
func genFn1() uint8 {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint16 = 2
        switch (fi7 + (fi7 ^ fi7)) {
        case -259:
            var v9 uint8 = 2
            _ = v9
        case -907:
            var v10 string = (("foo bar" + (("\n\t" + ("a" + "foo bar")) + ("foo bar" + "foo bar"))) + ("foo bar" + "a"))
            _ = v10
        case -2147483648:
            if false {
                fmt.Println(v8, v8)
                v8 = (^(v8))
            } else {
                for fi11 := 0; fi11 < 18; fi11++ {
                    if false {
                        if (0 <= v8) {
                            v12 := []int{}
                            _ = v12
                        } else {
                            v13 := GenS1{F0: "", F1: (!(100 == 1))}
                            _ = v13
                        }
                        v14 := GenS1{F0: ("a" + ("hello" + (((("foo bar" + "a") + ("\n\t" + "a")) + "a") + "a"))), F1: (-653 > 9223372036854775807)}
                        if false {
                            var v15 uint32 = 4294967295
                            var v16 float64 = (float64((1e10 - 3.14159)) * 3.14159)
                            _ = v15
                            _ = v16
                        }
                        _ = v14
                    }
                    _ = fi11
                }
            }
        }
        _ = fi7
        _ = v8
    }
    v17 := GenS2{F0: true}
    _ = v17
    return 128
}

func main() {
    var v1 uint64 = 380
    var v2 int8 = 1
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]string)
    v3 = v3
    if ((((1e10 * 0.0) * ((float64(float32(0.0)) / 1e10) + 0.0)) / 1.0) <= 1e10) {
        for fi18 := 0; fi18 < 17; fi18++ {
            v4[(^(fi18))] = "a"
            if (-2.25 <= ((((0.1 / -2.25) / 1.5) * -2.25) - 1e10)) {
                var v19 int64 = (v3 ^ v3)
                _ = v19
            }
            v20 := GenS2{F0: ((v3 * -9223372036854775808) != -9223372036854775808)}
            _ = fi18
            _ = v20
        }
        for fi21 := 0; fi21 < 19; fi21++ {
            if (32767 < -32768) {
                v1 += v1
                fmt.Println(v2, v2)
                var v22 int = -9223372036854775808
                _ = v22
            }
            _ = fi21
        }
        switch "hello" {
        case "a":
            switch ((("a" + (("a" + "\n\t") + "")) + "foo bar") + ("a" + ("foo bar" + ("\n\t" + "\n\t")))) {
            case "":
                v1 = (^(v1))
            case "a":
                if true {
                    v23 := GenS3{F0: 880, F1: 384}
                    _ = v23
                } else {
                    var v24 int32 = 0
                    _ = v24
                }
            }
        case "hello":
            for fi25 := 0; fi25 < 3; fi25++ {
                for fi26 := 0; fi26 < 13; fi26++ {
                    v27 := []int{1, (fi25 | (fi25 << ((0) & 63)))}
                    _ = fi26
                    _ = v27
                }
                v28 := []int{(fi25 - (fi25 >> ((893) & 63))), fi25, fi25}
                _ = fi25
                _ = v28
            }
        case "\n\t":
            switch 128 {
            case 1:
                var v29 uint32 = 473
                _ = v29
            case -2147483648:
                v30 := GenS2{F0: true}
                _ = v30
            case -757:
                v31 := GenS3{F0: -542, F1: 0}
                _ = v31
            }
        default:
            v32 := GenS4{F0: 255, F1: 42}
            _ = v32
        }
    } else {
        if false {
            var v33 int16 = 1000
            v1 += (v1 ^ v1)
            _ = v33
        }
        if (100 > (v2 + genFn0(0))) {
            var v34 int8 = (v2 - v2)
            for fi35 := 0; fi35 < 14; fi35++ {
                v2 -= v34
                var v36 bool = false
                _ = fi35
                _ = v36
            }
            _ = v34
        }
    }
    _ = len(v4)
    v37 := GenS2{F0: false}
    v3 = v3
    v38 := GenS5{F0: 42, F1: 0, F2: 2}
    v2 = v2
    v39 := GenS6{F0: 255}
    delete(v4, 32767)
    var v40 float32 = (3.14159 / (1.5 / 3.14159))
    v40 -= v40
    if false {
        v2 = v2
        var v41 float32 = 0.1
        v1 = (v1 ^ v1)
        _ = v41
    }
    if (0.0 >= (1.5 - 0.0)) {
        for fi42 := 0; fi42 < 20; fi42++ {
            for fi43 := 0; fi43 < 2; fi43++ {
                v44 := GenS6{F0: 2}
                var v45 float64 = (1.5 / 1.0)
                var v46 float32 = float32(v45)
                v47 := []int{-1}
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
                _ = v47
            }
            var v48 uint = 4294967295
            _ = fi42
            _ = v48
        }
    } else {
        v49 := GenS1{F0: "", F1: ((0 != (^(v1))) || false)}
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
}
