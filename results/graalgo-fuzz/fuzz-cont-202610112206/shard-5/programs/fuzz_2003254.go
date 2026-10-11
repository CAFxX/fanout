package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int32; F1 uint; F2 uint }
type GenS3 struct { F0 bool; F1 uint32; F2 bool }
func genFn0() uint8 {
    if true {
        var v7 uint8 = 227
        v8 := []int{474}
        if (true && (true || true)) {
            var v9 int8 = -1
            v7 -= (^(v7))
            var v10 string = ((("a" + "\n\t") + "") + ("\n\t" + (("hello" + ("foo bar" + ("foo bar" + ("" + ("foo bar" + "foo bar"))))) + "foo bar")))
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 19; fi11++ {
                var v12 uint = 334
                _ = fi11
                _ = v12
            }
            var v13 int16 = 448
            v7 = v7
            _ = v13
        }
        _ = v7
        _ = v8
    } else {
        v14 := GenS1{F0: 1000, F1: -66}
        _ = v14
    }
    var v15 int = 128
    v16 := GenS1{F0: int16(v15), F1: 42}
    _ = v15
    _ = v16
    return 70
}
func genFn1() int16 {
    if (4294967295 == 42) {
        for fi17 := 0; fi17 < 8; fi17++ {
            var v18 uint = 216
            switch "a" {
            case "a":
                var v19 int = fi17
                _ = v19
            case "hello":
                v18 = v18
            case "":
                fmt.Println(fi17)
            default:
                v20 := GenS1{F0: -475, F1: 42}
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        var v21 string = "\n\t"
        for fi22 := 0; fi22 < 19; fi22++ {
            v21 = v21
            if true {
                v21 = v21
                var v23 int64 = 800
                _ = v23
            } else {
                fmt.Println(v21, fi22)
                var v24 uint32 = 616
                v24 = v24
                _ = v24
            }
            _ = fi22
        }
        _ = v21
    }
    for fi25 := 0; fi25 < 5; fi25++ {
        fmt.Println(fi25)
        _ = fi25
    }
    return 42
}

func main() {
    var v1 int8 = 0
    v2 := make(map[uint32]string)
    var v3 string = (("\n\t" + "\n\t") + "\n\t")
    var v4 uint32 = 2147483647
    var v5 int64 = 9223372036854775807
    var v6 uint32 = 65535
    var v26 int16 = 32767
    fmt.Println(v3, v3)
    v26 = (v26 >> ((504) & 63))
    fmt.Println(v5)
    for fi27 := 0; fi27 < 4; fi27++ {
        var v28 uint64 = 42
        v6 = 0
        var v29 int16 = 1000
        _ = fi27
        _ = v28
        _ = v29
    }
    switch v3 {
    case "a":
        v1 = v1
    default:
        switch (v3 + (v3 + v3)) {
        case "":
            v30 := []int{-2147483648}
            _ = v30
        case "\n\t":
            fmt.Println(v5, v26, v1)
        case "a":
            v6 = v4
        }
    }
    fmt.Println(v26)
    if (!(true && false)) {
        for fi31 := 0; fi31 < 8; fi31++ {
            if (true && (0 != 2)) {
                for fi32 := 0; fi32 < 20; fi32++ {
                    v33 := v2[v6]
                    var v34 int64 = v5
                    v26 = -32768
                    v4 += (v6 | 258)
                    if (-1 <= 2147483647) {
                        var v35 int8 = (v1 / (((v1 % (((v1 / (((v1 % (((v1 / (((v1 >> ((914) & 63))) | 1))) | 1))) | 1))) | 1))) | 1))
                        _ = v35
                    } else {
                        if (3.14159 != (((0.0 / 1e10) / 1.5) / 1.5)) {
                            v36 := GenS2{F0: 2147483647, F1: 1, F2: 18446744073709551615}
                            if (fi32 == 32768) {
                                if true {
                                    v33 = v3
                                    fmt.Println(v5)
                                }
                            } else {
                                v2[v6] = ("\n\t" + v33)
                                _ = len(v2)
                                v37 := []int{fi31, fi31, (fi31 + 4294967295), -1}
                                _ = v37
                            }
                            _ = v36
                        }
                        var v38 uint32 = v6
                        _ = v38
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            } else {
                v39 := GenS3{F0: true, F1: (^(v4)), F2: true}
                _ = v39
            }
            var v40 int = 4294967295
            _ = fi31
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
