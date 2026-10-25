package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int64; F2 bool }
type GenS3 struct { F0 int16; F1 int32; F2 int64 }
type GenS4 struct { F0 int8 }
func genFn0() int {
    var v7 string = ""
    var v8 int8 = 127
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 float32 = ((1e10 - 0.1) / 1.0)
        v8 = 0
        _ = fi9
        _ = v10
    }
    v8 += v8
    fmt.Println(v7, v7)
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]uint32)
    var v3 bool = (2147483648 <= 896)
    var v4 int = 1
    v5 := make(map[int]string)
    v6 := make(map[uint32]int)
    if (-435 <= 260) {
        switch v4 {
        case -2147483648:
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 uint8 = 204
                var v13 int64 = 0
                _ = fi11
                _ = v12
                _ = v13
            }
        case -945:
            v14 := GenS1{F0: 2147483648}
            _ = v14
        case 2147483647:
            switch ("foo bar" + "hello") {
            case "\n\t":
                var v15 string = ("a" + "\n\t")
                _ = v15
            }
        }
        v16 := GenS2{F0: 821, F1: 42, F2: (18446744073709551615 == 18446744073709551615)}
        v3 = (!(-9223372036854775808 > 0))
        _ = v16
    } else {
        var v17 int = (v4 / ((-2147483648) | 1))
        var v18 int8 = 113
        _ = v17
        _ = v18
    }
    v19 := []int{(v4 ^ (v4 % ((v4) | 1)))}
    if (!(v3 && (v3 || (!(611 < 65535))))) {
        var v20 int8 = 0
        v4 -= -9223372036854775808
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int32 = 878
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    var v23 bool = true
    if v23 {
        v24 := GenS1{F0: 2147483648}
        switch "" {
        case "\n\t":
            var v25 float32 = (1e10 / ((0.0 - -2.25) - float32(v1)))
            _ = v25
        case "foo bar":
            var v26 bool = (287 > 2)
            _ = v26
        }
        _ = v24
    } else {
        var v27 int = (-(v4))
        v4 -= (^(v4))
        _ = v27
    }
    v19 = append(v19, 1)
    v28 := GenS1{F0: 4294967295}
    fmt.Println(v23, v23, v4)
    fmt.Println(v3, v4, v4)
    v4 = 32768
    v29 := GenS3{F0: 0, F1: 0, F2: 163}
    if (!false) {
        var v30 int32 = 778
        for fi31 := 0; fi31 < 19; fi31++ {
            v23 = (244 < 32767)
            switch (fi31 + (fi31 ^ (v4 / (((v4 | v4)) | 1)))) {
            case -9223372036854775808:
                if (false && (v3 && v23)) {
                    var v32 int16 = 32767
                    fmt.Println(v30)
                    if false {
                        for fi33 := 0; fi33 < 6; fi33++ {
                            v34 := GenS4{F0: -1}
                            if v23 {
                                v35 := GenS1{F0: 42}
                                v36 := GenS4{F0: 42}
                                _ = v35
                                _ = v36
                            }
                            _ = fi33
                            _ = v34
                        }
                        v30 = (v30 << ((65535) & 63))
                    }
                    _ = v32
                } else {
                    v3 = true
                }
            }
            _ = fi31
        }
        delete(v5, (v4 - v4))
        _ = v30
    } else {
        for fi37 := 0; fi37 < 8; fi37++ {
            _ = len(v6)
            var v38 float64 = ((float64((((0.1 * 0.0) / 0.1) + 0.1)) / 3.14159) * v1)
            v4 = 0
            _ = fi37
            _ = v38
        }
        if (!(-1 >= -1)) {
            v4 = genFn0()
            var v39 string = (((("a" + "") + "hello") + "a") + ("" + "hello"))
            _ = v39
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
