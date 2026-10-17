package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 bool }

func main() {
    var v1 uint32 = 1
    var v2 uint32 = (v1 ^ v1)
    v3 := make(map[string]bool)
    var v4 int16 = 1000
    var v5 string = ("foo bar" + (((("foo bar" + "foo bar") + (("" + "hello") + "hello")) + ("" + "hello")) + ("foo bar" + ("foo bar" + (("" + "") + ((("foo bar" + "\n\t") + "hello") + ""))))))
    if (true && false) {
        v6 := GenS1{F0: uint16(v1), F1: 32768, F2: (("a" + "hello") + v5)}
        var v7 int32 = 667
        for fi8 := 0; fi8 < 12; fi8++ {
            switch (fi8 / (((fi8 >> ((4294967295) & 63))) | 1)) {
            case -685:
                fmt.Println(fi8)
            case 619:
                v9 := GenS1{F0: 65535, F1: -1, F2: (v5 + ((("" + "a") + (v5 + (v5 + v5))) + ((v5 + v5) + v5)))}
                _ = v9
            default:
                var v10 int16 = (v4 / (((v4 % (((^(v4))) | 1))) | 1))
                _ = v10
            }
            _ = fi8
        }
        _ = v6
        _ = v7
    } else {
        v11 := GenS1{F0: 2, F1: 32767, F2: (v5 + "foo bar")}
        for fi12 := 0; fi12 < 18; fi12++ {
            v4 = v4
            var v13 uint32 = 65535
            v2 += 65535
            _ = fi12
            _ = v13
        }
        fmt.Println(v4)
        _ = v11
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int = (fi14 | fi14)
        if false {
            switch v15 {
            case 4294967295:
                switch v15 {
                case 32768:
                    var v16 int8 = 100
                    _ = v16
                default:
                    switch v15 {
                    case 1:
                        if true {
                            v4 = v4
                            if false {
                                for fi17 := 0; fi17 < 8; fi17++ {
                                    v15 = 128
                                    var v18 uint = 255
                                    var v19 uint32 = (v2 - (v1 % ((v2) | 1)))
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                                for fi20 := 0; fi20 < 12; fi20++ {
                                    v4 -= (v4 / (((v4 - 0)) | 1))
                                    delete(v3, ("" + ("foo bar" + (v5 + (v5 + (v5 + v5))))))
                                    _ = len(v3)
                                    _ = fi20
                                }
                            }
                        }
                    case 128:
                        if (true && (true && (!false))) {
                            v21 := []int{}
                            v5 = v5
                            _ = v21
                        }
                    default:
                        v22 := v3[v5]
                        _ = v22
                    }
                }
            }
            v5 = (((v5 + v5) + ("foo bar" + v5)) + (v5 + "hello"))
            var v23 int64 = 1
            _ = v23
        } else {
            switch (v5 + ((v5 + v5) + (v5 + (v5 + v5)))) {
            case "":
                if false {
                    switch "" {
                    case "foo bar":
                        if (v4 == v4) {
                            v24 := []int{fi14, (fi14 << ((560) & 63))}
                            v25 := GenS1{F0: 2, F1: (v15 / ((42) | 1)), F2: v5}
                            _ = v24
                            _ = v25
                        }
                    case "hello":
                        v5 = ""
                    case "a":
                        var v26 float64 = 3.14159
                        _ = v26
                    }
                    v4 = 0
                    v5 = ""
                }
            case "\n\t":
                delete(v3, "hello")
            default:
                var v27 uint64 = 0
                _ = v27
            }
        }
        _ = fi14
        _ = v15
    }
    var v28 uint = 1
    if (((1e10 - 1.5) / 1.5) < 3.14159) {
        v29 := []int{}
        for fi30 := 0; fi30 < 11; fi30++ {
            var v31 uint32 = (v2 + 212)
            var v32 int = (fi30 << (((v28 | (v28 % (((v28 % (((v28 ^ 4294967295)) | 1))) | 1)))) & 63))
            _ = fi30
            _ = v31
            _ = v32
        }
        var v33 bool = (false && (611 == 0))
        _ = v29
        _ = v33
    }
    if true {
        switch ((("" + ("a" + v5)) + "") + "a") {
        case "\n\t":
            var v34 int64 = 9223372036854775807
            _ = v34
        }
        if (v5 <= v5) {
            v28 += 687
            switch "hello" {
            case "foo bar":
                v35 := GenS2{F0: 42}
                _ = v35
            }
            v36 := GenS3{F0: -9223372036854775808}
            _ = v36
        } else {
            for fi37 := 0; fi37 < 3; fi37++ {
                fmt.Println(v1, v4, v28)
                _ = fi37
            }
        }
        var v38 float32 = (0.0 - 0.1)
        _ = v38
    }
    v39 := GenS4{F0: 1, F1: ((-2147483648 <= 42) && true)}
    v40 := []int{-135, 32768}
    v41 := v3["hello"]
    v1 += (v2 % (((v1 ^ (v1 / ((42) | 1)))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
