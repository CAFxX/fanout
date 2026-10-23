package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    var v3 uint32 = 2147483648
    var v4 bool = (!("\n\t" > "a"))
    var v5 string = "foo bar"
    var v6 int16 = 0
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := v1[71]
        for fi9 := 0; fi9 < 8; fi9++ {
            if true {
                var v10 string = ("\n\t" + ("hello" + "a"))
                if v4 {
                    switch (fi9 / (((-(fi7))) | 1)) {
                    case -947:
                        switch fi7 {
                        case 127:
                            v11 := GenS1{F0: -32768}
                            _ = v11
                        case 950:
                            fmt.Println(v6, fi7, v6)
                        case 128:
                            switch (fi7 ^ (fi7 - 32768)) {
                            case 9223372036854775807:
                                v2["a"] = 9223372036854775807
                            case 0:
                                v12 := []int{-503, fi9}
                                _ = v12
                            default:
                                if true {
                                    var v13 bool = (!(!((v4 && (128 != 152)) || false)))
                                    _ = v13
                                } else {
                                    v14 := []int{fi7, (-(fi7))}
                                    _ = v14
                                }
                            }
                        default:
                            v6 -= v6
                        }
                    case -2147483648:
                        v3 -= v3
                        v15 := GenS2{F0: 617, F1: (fi7 >> ((0) & 63))}
                        _ = v15
                    case 885:
                        v16 := GenS3{F0: 18446744073709551615, F1: 142}
                        _ = v16
                    }
                    switch (fi9 ^ fi9) {
                    case -797:
                        v8 = (v8 + v8)
                    case -28:
                        var v17 uint32 = v3
                        _ = v17
                    case -1:
                        if (v4 && v4) {
                            var v18 bool = v4
                            if (!(v8 < "foo bar")) {
                                v19 := GenS1{F0: v6}
                                var v20 int64 = 0
                                _ = v19
                                _ = v20
                            } else {
                                v21 := GenS4{F0: (fi9 | fi9)}
                                _ = v21
                            }
                            v5 = "foo bar"
                            _ = v18
                        }
                    }
                    v22 := GenS5{F0: -643, F1: (^(v3))}
                    _ = v22
                } else {
                    delete(v2, v10)
                    v23 := v1[(fi9 << ((18446744073709551615) & 63))]
                    fmt.Println(v3, fi7, v4)
                    _ = v23
                }
                delete(v1, 42)
                fmt.Println(v8)
                _ = v10
            }
            _ = fi9
        }
        var v24 int32 = -613
        _ = fi7
        _ = v24
        _ = v8
    }
    var v25 int64 = -1
    var v26 int8 = 42
    if v4 {
        v3 -= v3
        _ = len(v2)
        if ((4294967295 <= 42) || (!(int(v26) >= 32768))) {
            v27 := []int{32768, -9223372036854775808, -101, 4294967295}
            _ = v27
        }
    } else {
        for fi28 := 0; fi28 < 5; fi28++ {
            v29 := []int{(fi28 ^ (fi28 + fi28)), (fi28 / ((4294967295) | 1))}
            _ = fi28
            _ = v29
        }
        v2[""] = 665
    }
    for fi30 := 0; fi30 < 12; fi30++ {
        delete(v1, -2147483648)
        var v31 uint16 = 888
        v3 = 915
        _ = fi30
        _ = v31
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        fmt.Println(v4, v3, v26)
        _ = fi32
    }
    if v4 {
        var v33 int16 = (-(v6))
        _ = v33
    } else {
        if (v4 && v4) {
            v34 := GenS4{F0: -9223372036854775808}
            v4 = false
            v35 := GenS6{F0: 42, F1: (v25 | 9223372036854775807)}
            _ = v34
            _ = v35
        }
    }
    if v4 {
        v5 = v5
        v6 += (v6 % ((v6) | 1))
        v4 = (177 <= 255)
    }
    var v36 uint8 = 28
    v4 = v4
    fmt.Println(v4, v4, v36)
    v26 = 42
    v36 -= 164
    fmt.Println(v36)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
