package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint32 = 1
    var v2 int8 = 100
    var v3 bool = true
    v4 := []int{}
    if (v3 || v3) {
        v3 = (770 == 0)
        v5 := GenS1{F0: 556, F1: "hello", F2: ((v1 % ((v1) | 1)) <= (v1 / (((v1 ^ (v1 / ((v1) | 1)))) | 1)))}
        _ = len(v4)
        _ = v5
    }
    v3 = (4294967295 > 0)
    for fi6 := 0; fi6 < 4; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            v1 -= v1
            var v8 int32 = -192
            if (755 <= 624) {
                v9 := GenS2{F0: -10}
                _ = v9
            }
            _ = fi7
            _ = v8
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            if v3 {
                var v11 uint8 = 255
                if v3 {
                    v12 := GenS1{F0: int(v2), F1: "a", F2: (343 != uint(v11))}
                    v11 = (v11 + (v11 & 255))
                    _ = v12
                }
                fmt.Println(v1, v2)
                _ = v11
            }
            v13 := GenS1{F0: (fi6 - (fi6 + fi10)), F1: "a", F2: (!("" > (("\n\t" + "hello") + "a")))}
            _ = fi10
            _ = v13
        }
        fmt.Println(fi6, v3)
        _ = fi6
    }
    var v14 float64 = (0.1 * (-2.25 / 1.0))
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 string = "\n\t"
        var v17 bool = (-162 > 42)
        fmt.Println(v14, v14)
        _ = fi15
        _ = v16
        _ = v17
    }
    v14 = (0.0 + v14)
    switch -9223372036854775808 {
    case 42:
        v3 = v3
    case 128:
        var v18 float64 = v14
        _ = v18
    case -799:
        _ = len(v4)
    }
    v14 -= float64(1e10)
    var v19 uint16 = 32767
    var v20 uint64 = 42
    switch ("a" + ("hello" + ("" + "a"))) {
    case "foo bar":
        switch 128 {
        case -9223372036854775808:
            if false {
                v21 := GenS1{F0: 128, F1: ((("a" + ("a" + "")) + "hello") + "\n\t"), F2: (149 <= (^(v19)))}
                _ = v21
            } else {
                if ((v1 ^ (v1 * 2147483648)) == (v1 / ((v1) | 1))) {
                    v4 = append(v4, 42)
                } else {
                    var v22 int8 = 116
                    if len(v4) > 0 {
                        v4[0] = -353
                    }
                    _ = v22
                }
                var v23 uint32 = v1
                fmt.Println(v23, v1)
                _ = v23
            }
        }
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = -242
        }
    case "a":
        var v24 bool = v3
        _ = v24
    default:
        for fi25 := 0; fi25 < 5; fi25++ {
            if true {
                var v26 int16 = 585
                _ = v26
            }
            v3 = (true && v3)
            v19 += 217
            _ = fi25
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 -= (v2 + v2)
    fmt.Println(v14, v19)
    switch ((("hello" + "a") + "a") + "") {
    case "foo bar":
        if (947 == 1) {
            v2 = -1
            var v27 uint8 = 0
            _ = v27
        }
    case "\n\t":
        var v28 uint64 = (v20 - v20)
        _ = v28
    default:
        if (-1 > 201) {
            for fi29 := 0; fi29 < 9; fi29++ {
                for fi30 := 0; fi30 < 3; fi30++ {
                    var v31 float32 = ((0.0 + 1.5) * (float32((0.1 * (v14 + 3.14159))) / float32(-2.25)))
                    fmt.Println(v3, fi30)
                    if v3 {
                        var v32 uint32 = v1
                        v33 := GenS3{F0: v1, F1: "\n\t"}
                        if len(v4) > 0 {
                            v4[0] = (fi30 / ((-9223372036854775808) | 1))
                        }
                        var v34 uint16 = (^(v19))
                        _ = v32
                        _ = v33
                        _ = v34
                    } else {
                        v35 := GenS4{F0: 42}
                        var v36 uint16 = v19
                        _ = v35
                        _ = v36
                    }
                    _ = fi30
                    _ = v31
                }
                _ = fi29
            }
            var v37 uint64 = (v20 % ((v20) | 1))
            v38 := GenS1{F0: 32767, F1: "", F2: false}
            _ = v37
            _ = v38
        }
    }
    v20 = (^(v20))
    v39 := GenS4{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
}
