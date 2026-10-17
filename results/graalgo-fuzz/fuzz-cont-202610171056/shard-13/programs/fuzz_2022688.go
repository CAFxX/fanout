package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int64 }
type GenS2 struct { F0 string }
func genFn0() bool {
    if (-1 != -2147483648) {
        v7 := []int{2147483647, 2147483647, 127, 761}
        if len(v7) > 0 {
            v7[0] = 9223372036854775807
        }
        _ = v7
    }
    if true {
        v8 := []int{42}
        var v9 bool = (512 > 65535)
        _ = v8
        _ = v9
    } else {
        var v10 float32 = (0.1 / 0.1)
        v11 := []int{-80}
        var v12 string = "a"
        _ = v10
        _ = v11
        _ = v12
    }
    v13 := []int{1, -1, 837, 127}
    switch "hello" {
    case "hello":
        var v14 bool = (18446744073709551615 < 65535)
        _ = v14
    case "\n\t":
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 uint32 = 827
            if true {
                switch 127 {
                case 9223372036854775807:
                    var v17 int16 = 42
                    _ = v17
                case -371:
                    var v18 int = (fi15 % ((fi15) | 1))
                    _ = v18
                case 486:
                    v19 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: -292}
                    _ = v19
                default:
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                if true {
                    v13 = append(v13, (fi15 ^ (fi15 + -9223372036854775808)))
                    var v20 int8 = -128
                    _ = v20
                }
            }
            _ = fi15
            _ = v16
        }
    default:
        var v21 int64 = 1
        _ = v21
    }
    _ = v13
    return ((2 <= 79) || false)
}

func main() {
    var v1 bool = true
    var v2 float32 = 1.5
    var v3 float32 = v2
    v4 := []uint32{49}
    var v5 uint8 = 42
    var v6 int = 898
    for fi22 := 0; fi22 < 6; fi22++ {
        for fi23 := 0; fi23 < 13; fi23++ {
            switch "hello" {
            case "hello":
                fmt.Println(v6)
                var v24 int32 = 1
                _ = v24
            }
            if (v1 || v1) {
                fmt.Println(v5, fi22)
                fmt.Println(v3)
                v25 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: 1}
                _ = v25
            } else {
                v26 := GenS1{F0: 18446744073709551615, F1: 2, F2: -1}
                _ = v26
            }
            if ((0.1 == 3.14159) || v1) {
                v5 = (v5 + v5)
                if v1 {
                    if (!v1) {
                        v5 -= 127
                    } else {
                        v3 -= v2
                    }
                }
            }
            _ = fi23
        }
        _ = fi22
    }
    v27 := GenS1{F0: 780, F1: 0, F2: -717}
    v2 = v2
    var v28 int16 = 32767
    var v29 int32 = 0
    fmt.Println(v1, v3)
    v30 := GenS1{F0: 427, F1: 42, F2: int64(v28)}
    v4 = append(v4, 647)
    var v31 string = ""
    if len(v4) > 0 {
        v4[0] = 2
    }
    for fi32 := 0; fi32 < 2; fi32++ {
        v28 -= (v28 | (v28 - v28))
        switch -13 {
        case -314:
            for fi33 := 0; fi33 < 5; fi33++ {
                fmt.Println(v2)
                v29 = v29
                _ = fi33
            }
        case 2147483647:
            v34 := GenS2{F0: v31}
            _ = v34
        }
        var v35 string = ("hello" + (("hello" + (v31 + v31)) + (("a" + ("a" + v31)) + (v31 + ("\n\t" + "foo bar")))))
        _ = fi32
        _ = v35
    }
    fmt.Println(v28)
    fmt.Println(v28, v31, v31)
    for fi36 := 0; fi36 < 4; fi36++ {
        var v37 int = (v6 >> ((1) & 63))
        v38 := GenS2{F0: ""}
        if false {
            switch (fi36 % (((v37 * fi36)) | 1)) {
            case -485:
                v3 -= v2
            case 32768:
                var v39 bool = (v1 || true)
                _ = v39
            case 4294967295:
                if v1 {
                    for fi40 := 0; fi40 < 5; fi40++ {
                        switch (v31 + v31) {
                        case "hello":
                            fmt.Println(v2, v3)
                        default:
                            var v41 uint16 = 2
                            _ = v41
                        }
                        v1 = (!false)
                        _ = fi40
                    }
                    v1 = (false && (29 != -128))
                }
            }
        } else {
            _ = len(v4)
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    var v42 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
