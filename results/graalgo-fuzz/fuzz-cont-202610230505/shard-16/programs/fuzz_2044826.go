package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint64; F2 int }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 bool = (1e10 == -2.25)
    var v2 float32 = 3.14159
    var v3 int64 = 1
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := []int{(fi4 ^ (fi4 ^ fi4))}
        for fi6 := 0; fi6 < 3; fi6++ {
            switch ("" + "foo bar") {
            case "hello":
                v3 = (v3 ^ v3)
            case "a":
                if (!v1) {
                    var v7 uint8 = 2
                    var v8 uint8 = v7
                    _ = v7
                    _ = v8
                } else {
                    v9 := GenS1{F0: 42}
                    if v1 {
                        v3 += v3
                        v10 := GenS1{F0: 1}
                        var v11 uint8 = 42
                        _ = v10
                        _ = v11
                    } else {
                        _ = len(v5)
                        v1 = (!true)
                        v3 = (-(v3))
                        v12 := GenS1{F0: 1}
                        _ = v12
                    }
                    _ = v9
                }
            default:
                v3 += -1
            }
            v3 = (v3 + (v3 - -594))
            var v13 uint8 = 128
            _ = fi6
            _ = v13
        }
        _ = fi4
        _ = v5
    }
    if ("" <= ("\n\t" + "a")) {
        if v1 {
            for fi14 := 0; fi14 < 20; fi14++ {
                switch (("hello" + "a") + "\n\t") {
                case "\n\t":
                    switch 128 {
                    case -526:
                        switch "foo bar" {
                        case "\n\t":
                            v15 := GenS2{F0: "", F1: 18446744073709551615, F2: 128}
                            _ = v15
                        case "foo bar":
                            var v16 uint16 = 32767
                            _ = v16
                        case "a":
                            v1 = true
                        default:
                            v3 -= (v3 << ((65535) & 63))
                        }
                    case 0:
                        for fi17 := 0; fi17 < 20; fi17++ {
                            if (true || v1) {
                                var v18 bool = v1
                                v18 = (((-14 > 100) || (-929 >= 0)) && v18)
                                _ = v18
                            }
                            var v19 int32 = 42
                            _ = fi17
                            _ = v19
                        }
                    }
                case "":
                    switch "\n\t" {
                    case "":
                        var v20 uint64 = 497
                        _ = v20
                    case "foo bar":
                        if v1 {
                            for fi21 := 0; fi21 < 2; fi21++ {
                                v1 = (((fi21 - (fi21 >> ((14) & 63))) >= fi14) && (!(!v1)))
                                if true {
                                    fmt.Println(fi21, v3, fi14)
                                }
                                var v22 int16 = -955
                                _ = fi21
                                _ = v22
                            }
                        } else {
                            var v23 int32 = 2147483647
                            _ = v23
                        }
                    }
                default:
                    var v24 float32 = (v2 * (v2 - (3.14159 + 1e10)))
                    _ = v24
                }
                v25 := GenS3{F0: 1}
                _ = fi14
                _ = v25
            }
        } else {
            var v26 string = ""
            if (-2.25 > 1.5) {
                if v1 {
                    var v27 uint32 = 2
                    _ = v27
                }
                v26 = "foo bar"
                v28 := GenS1{F0: 605}
                _ = v28
            } else {
                for fi29 := 0; fi29 < 4; fi29++ {
                    var v30 float32 = float32((1e10 * float64(3.14159)))
                    _ = fi29
                    _ = v30
                }
            }
            v31 := GenS1{F0: 0}
            _ = v26
            _ = v31
        }
    }
    v3 -= 1
    var v32 bool = (0 < 42)
    v33 := GenS1{F0: 42}
    v3 = v3
    var v34 float32 = v2
    v3 = (v3 / ((v3) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
