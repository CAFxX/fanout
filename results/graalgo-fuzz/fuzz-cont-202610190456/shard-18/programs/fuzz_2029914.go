package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 int64; F2 string }
type GenS3 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int64, p1 int8, p2 bool) uint16 {
    var v4 uint32 = 0
    v5 := GenS1{F0: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 280
}
func genFn1(p0 uint8, p1 int32) uint16 {
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 float64 = 3.14159
        if false {
            for fi8 := 0; fi8 < 11; fi8++ {
                p0 = (^(p0))
                v9 := []int{}
                switch (-(fi8)) {
                case 0:
                    var v10 bool = true
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
        } else {
            p1 = 42
            p1 = p1
            v11 := []int{}
            _ = v11
        }
        _ = fi6
        _ = v7
    }
    v12 := GenS1{F0: (-(p1))}
    _ = p0
    _ = p1
    _ = v12
    return 42
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]bool)
    var v3 int = 32767
    v13 := GenS1{F0: 633}
    v1 += v1
    v1 = (v1 / ((v1) | 1))
    var v14 bool = ((("" + "a") + "\n\t") > ("hello" + ""))
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int32 = -851
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := v2["foo bar"]
        if ((2 == 0) || v18) {
            v1 = -128
            var v19 float64 = -2.25
            switch "a" {
            case "a":
                v19 -= v19
            case "hello":
                v20 := GenS2{F0: 42, F1: 185, F2: "\n\t"}
                _ = v20
            case "":
                v21 := GenS3{F0: 32767, F1: 128, F2: 4294967295}
                _ = v21
            default:
                fmt.Println(v3, fi17)
            }
            _ = v19
        } else {
            switch (v3 % ((v3) | 1)) {
            case 4294967295:
                switch 9 {
                case 1:
                    var v22 int16 = 0
                    _ = v22
                default:
                    for fi23 := 0; fi23 < 19; fi23++ {
                        var v24 int = fi17
                        if (0 > 0) {
                            switch "" {
                            case "a":
                                var v25 uint64 = 304
                                _ = v25
                            }
                            var v26 int16 = 1
                            fmt.Println(fi23)
                            _ = v26
                        }
                        _ = fi23
                        _ = v24
                    }
                }
            }
            if v18 {
                for fi27 := 0; fi27 < 2; fi27++ {
                    var v28 int8 = v1
                    v29 := []int{fi17, fi27, (v3 & fi17)}
                    v3 -= 278
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            }
        }
        _ = fi17
        _ = v18
    }
    switch (v3 + -9223372036854775808) {
    case 32768:
        v30 := []int{(v3 & (-(v3))), 32768, v3}
        _ = v30
    case 301:
        v31 := GenS4{F0: 637}
        _ = v31
    case 470:
        v32 := []int{(v3 | (v3 * (v3 | int(v1)))), (v3 >> ((396) & 63)), v3, 128}
        _ = v32
    }
    var v33 string = ((((("\n\t" + "a") + "") + "a") + (("a" + ("\n\t" + "")) + ("\n\t" + (("\n\t" + "foo bar") + ((("foo bar" + "") + "hello") + ("foo bar" + "hello")))))) + ((("a" + ("\n\t" + "\n\t")) + "hello") + (((((("a" + ("a" + "hello")) + "") + "a") + "") + "foo bar") + ("a" + "foo bar"))))
    for fi34 := 0; fi34 < 10; fi34++ {
        v35 := GenS3{F0: -32768, F1: 128, F2: 42}
        _ = fi34
        _ = v35
    }
    var v36 int8 = v1
    if (!v14) {
        var v37 uint16 = 255
        for fi38 := 0; fi38 < 5; fi38++ {
            if v14 {
                v39 := GenS5{F0: 128}
                switch fi38 {
                case 1:
                    v2[v33] = v14
                }
                _ = v39
            }
            _ = fi38
        }
        _ = v37
    } else {
        var v40 uint16 = 42
        var v41 int8 = (v36 >> ((255) & 63))
        v42 := []int{v3, (v3 + (v3 % ((v3) | 1)))}
        _ = v40
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v36
}
