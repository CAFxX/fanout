package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int8; F1 bool; F2 bool }
type GenS3 struct { F0 bool; F1 int32; F2 string }
func genFn0() uint16 {
    var v6 bool = (!("foo bar" != "hello"))
    switch (("hello" + "a") + "a") {
    case "":
        v7 := GenS1{F0: 38, F1: 1, F2: 42}
        _ = v7
    }
    var v8 uint64 = 42
    _ = v6
    _ = v8
    return 299
}

func main() {
    var v1 uint8 = 161
    var v2 int32 = 467
    v3 := []string{"foo bar", "", "\n\t", ("foo bar" + ("a" + "hello"))}
    var v4 int = 0
    var v5 int32 = (v2 | v2)
    var v9 uint16 = 32767
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            v12 := GenS1{F0: -1, F1: genFn0(), F2: v9}
            if false {
                var v13 float64 = (3.14159 - (3.14159 - 1e10))
                var v14 string = ("foo bar" + ("hello" + ("" + ("\n\t" + "\n\t"))))
                _ = v13
                _ = v14
            } else {
                v2 -= (v2 << ((509) & 63))
                v4 = (v4 >> ((255) & 63))
            }
            _ = fi11
            _ = v12
        }
        if true {
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 uint8 = v1
                _ = fi15
                _ = v16
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v17 := GenS1{F0: 1, F1: v9, F2: (^(v9))}
            if true {
                v18 := GenS1{F0: -1, F1: (^(v9)), F2: genFn0()}
                _ = v18
            }
            var v19 int32 = v5
            _ = v17
            _ = v19
        }
        v1 = 127
        _ = fi10
    }
    switch 9223372036854775807 {
    case 1:
        for fi20 := 0; fi20 < 7; fi20++ {
            for fi21 := 0; fi21 < 14; fi21++ {
                var v22 int32 = v5
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    case -928:
        _ = len(v3)
    case 831:
        v23 := GenS1{F0: 100, F1: (v9 * 255), F2: 42}
        _ = v23
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v5 = v5
    for fi24 := 0; fi24 < 2; fi24++ {
        fmt.Println(v2, v9, v4)
        _ = fi24
    }
    v25 := GenS1{F0: 0, F1: 522, F2: 0}
    var v26 float32 = (((float32(0.0) + -2.25) / float32(1e10)) * 1.5)
    v27 := GenS1{F0: -79, F1: uint16(v1), F2: (v9 >> ((uint(v1)) & 63))}
    var v28 uint16 = v9
    switch v4 {
    case 128:
        for fi29 := 0; fi29 < 11; fi29++ {
            if false {
                for fi30 := 0; fi30 < 10; fi30++ {
                    var v31 bool = (((0.0 + (((1e10 + -2.25) * float64((v26 - 1.5))) / (0.1 / 3.14159))) - 0.1) < (((1e10 - 1e10) - 1e10) * 1e10))
                    _ = fi30
                    _ = v31
                }
            } else {
                if true {
                    var v32 int = 2147483647
                    v9 = 65535
                    _ = v32
                } else {
                    if true {
                        fmt.Println(v1, v9, v5)
                        v3 = append(v3, (("" + "hello") + "foo bar"))
                    }
                }
            }
            v26 = 1e10
            var v33 int64 = 916
            _ = fi29
            _ = v33
        }
    case 561:
        var v34 uint8 = 128
        _ = v34
    default:
        var v35 int8 = -1
        _ = v35
    }
    for fi36 := 0; fi36 < 2; fi36++ {
        v26 -= (v26 - float32(float64(v26)))
        _ = fi36
    }
    if ((v5 * v2) == (v2 ^ (^(v5)))) {
        if true {
            for fi37 := 0; fi37 < 4; fi37++ {
                var v38 int8 = 127
                if len(v3) > 0 {
                    v3[0] = (("hello" + "a") + "a")
                }
                v39 := GenS2{F0: v38, F1: true, F2: (!true)}
                _ = fi37
                _ = v38
                _ = v39
            }
            switch "foo bar" {
            case "a":
                for fi40 := 0; fi40 < 7; fi40++ {
                    v28 = v9
                    _ = fi40
                }
            case "hello":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            var v41 uint = 4294967295
            _ = v41
        } else {
            v9 = (^(v28))
        }
        v42 := GenS3{F0: false, F1: 0, F2: ("\n\t" + "\n\t")}
        v28 = genFn0()
        _ = v42
    } else {
        fmt.Println(v4, v2, v5)
        var v43 int32 = (v2 & v5)
        _ = v43
    }
    var v44 float32 = (v26 + (v26 - v26))
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v28 -= genFn0()
    v28 = v28
    var v45 uint16 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v9
}
