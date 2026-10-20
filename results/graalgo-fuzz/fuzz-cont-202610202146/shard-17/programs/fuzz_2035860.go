package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 bool }

func main() {
    var v1 string = "a"
    var v2 int16 = 42
    var v3 uint32 = 4294967295
    var v4 int32 = 42
    var v5 string = "a"
    v6 := []int{-2147483648, 0}
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v1)
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: 1, F1: 42}
            if false {
                var v10 uint64 = 0
                _ = v10
            } else {
                if (false || (true && true)) {
                    var v11 uint64 = 0
                    _ = v11
                } else {
                    if true {
                        if (0 != 0) {
                            v6 = append(v6, fi8)
                        } else {
                            _ = len(v6)
                        }
                        var v12 float32 = 3.14159
                        var v13 int16 = (v2 * v2)
                        v2 = (v13 | 42)
                        _ = v12
                        _ = v13
                    }
                }
                _ = len(v6)
            }
            _ = fi8
            _ = v9
        }
        v1 = v1
        _ = fi7
    }
    if ((!false) || true) {
        if false {
            v5 = v5
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 int8 = -128
                v16 := GenS1{F0: 65535, F1: 4294967295}
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 bool = (18446744073709551615 >= 352)
            _ = v17
        }
        var v18 bool = ((v3 - 2147483647) == (v3 - (v3 - v3)))
        var v19 int = -2147483648
        _ = v18
        _ = v19
    }
    if (!(false && false)) {
        if len(v6) > 0 {
            v6[0] = 0
        }
    }
    fmt.Println(v3, v1, v2)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS1{F0: 2, F1: 65535}
        fmt.Println(v4, fi20, v3)
        switch (fi20 << ((0) & 63)) {
        case 127:
            v22 := GenS1{F0: uint(v3), F1: 66}
            _ = v22
        }
        _ = fi20
        _ = v21
    }
    var v23 int16 = (v2 ^ v2)
    fmt.Println(v1, v4, v23)
    if false {
        if true {
            for fi24 := 0; fi24 < 20; fi24++ {
                for fi25 := 0; fi25 < 15; fi25++ {
                    if (!true) {
                        if false {
                            var v26 int32 = (v4 * (v4 << ((4294967295) & 63)))
                            _ = v26
                        }
                        v1 = "hello"
                    }
                    v1 = ""
                    v27 := GenS2{F0: 127, F1: 0}
                    _ = fi25
                    _ = v27
                }
                var v28 int64 = -281
                fmt.Println(v4, v1, v28)
                _ = fi24
                _ = v28
            }
        }
        switch -9223372036854775808 {
        case 846:
            v29 := GenS1{F0: 2, F1: 928}
            _ = v29
        }
    } else {
        v30 := GenS3{F0: (43 > 1)}
        v31 := GenS2{F0: 49, F1: 481}
        var v32 string = (("hello" + v1) + ("\n\t" + ("\n\t" + ("\n\t" + ("foo bar" + v5)))))
        _ = v30
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
