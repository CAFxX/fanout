package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0() uint8 {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int16 = 0
        v8 := []int{fi6, fi6, (fi6 >> ((18446744073709551615) & 63)), fi6}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if false {
        if true {
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 uint16 = 32767
                switch "hello" {
                case "hello":
                    v11 := GenS1{F0: 2147483647, F1: 0}
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
            if (1 >= 0) {
                var v12 int8 = 1
                if (1 == 4294967295) {
                    v12 = v12
                    var v13 float32 = 0.1
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v13 -= (v13 * 0.0)
                        fmt.Println(v13, v12)
                        _ = fi14
                    }
                    _ = v13
                }
                v15 := []int{-1, 128, 42, 1}
                _ = v12
                _ = v15
            } else {
                for fi16 := 0; fi16 < 6; fi16++ {
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v18 := []int{1, fi17, (fi16 & (fi17 / ((fi17) | 1))), -1}
                        v19 := GenS1{F0: 1, F1: 65535}
                        if ((0 <= 2147483647) && ((((("a" + "a") + (("foo bar" + ("" + ("\n\t" + ""))) + ("foo bar" + ("" + "")))) + "") > (("foo bar" + "") + "hello")) || false)) {
                            var v20 int = 128
                            _ = v20
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    var v21 uint16 = 42
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        v21 += v21
                    case "a":
                        v21 += (^(v21))
                    default:
                        var v22 uint32 = 2
                        _ = v22
                    }
                    _ = fi16
                    _ = v21
                }
                var v23 int64 = 1
                var v24 float32 = 0.0
                _ = v23
                _ = v24
            }
            for fi25 := 0; fi25 < 4; fi25++ {
                v26 := []int{4294967295, fi25, -9223372036854775808, fi25}
                fmt.Println(fi25)
                _ = fi25
                _ = v26
            }
        }
        v27 := GenS2{F0: 42, F1: 32767}
        _ = v27
    }
    var v28 int = -1
    _ = v28
    return 39
}

func main() {
    var v1 uint8 = 89
    v2 := make(map[int]bool)
    var v3 uint8 = 127
    v4 := make(map[int]string)
    var v5 int16 = 1000
    v29 := GenS2{F0: 2, F1: (v5 - 1000)}
    v1 = v1
    v1 += 147
    var v30 int16 = (v5 >> ((185) & 63))
    fmt.Println(v30, v30)
    var v31 uint32 = 0
    v32 := GenS3{F0: 401, F1: (v5 ^ v30)}
    v33 := GenS3{F0: 1, F1: v30}
    var v34 int64 = 882
    var v35 bool = true
    switch -9223372036854775808 {
    case 2147483647:
        v3 = v1
    default:
        switch -1 {
        case 42:
            v30 = v5
        case 128:
            v36 := v4[216]
            _ = v36
        case 32767:
            if v35 {
                switch "" {
                case "\n\t":
                    fmt.Println(v5)
                default:
                    var v37 string = "a"
                    _ = v37
                }
                for fi38 := 0; fi38 < 13; fi38++ {
                    v3 = 128
                    switch ("" + "hello") {
                    case "a":
                        v30 = v30
                    }
                    _ = fi38
                }
                if v35 {
                    v34 += v34
                    fmt.Println(v35)
                }
            } else {
                if (151 != (v1 / ((v3) | 1))) {
                    v39 := GenS3{F0: 42, F1: 0}
                    for fi40 := 0; fi40 < 10; fi40++ {
                        var v41 int64 = 0
                        v42 := GenS1{F0: v31, F1: 65535}
                        for fi43 := 0; fi43 < 16; fi43++ {
                            _ = len(v2)
                            _ = fi43
                        }
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    if (87 == 354) {
                        v44 := GenS3{F0: 42, F1: (v5 % ((v5) | 1))}
                        _ = v44
                    }
                    _ = v39
                }
                if v35 {
                    switch 32768 {
                    case 826:
                        fmt.Println(v35, v35, v1)
                    }
                    v45 := []int{42, 0, 32768}
                    _ = v45
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
