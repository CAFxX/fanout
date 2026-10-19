package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }

func main() {
    var v1 int = 592
    v2 := []int{(v1 / (((^(v1))) | 1)), (v1 << ((4294967295) & 63)), (v1 / (((v1 & (v1 << ((4294967295) & 63)))) | 1))}
    var v3 string = "\n\t"
    v4 := make(map[string]bool)
    var v5 int8 = 127
    var v6 int64 = 720
    v2 = append(v2, 2147483647)
    fmt.Println(v6, v5)
    if ("" != "hello") {
        var v7 bool = (("" <= "a") || false)
        _ = v7
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi8 := 0; fi8 < 5; fi8++ {
            v1 = (fi8 / (((fi8 + 820)) | 1))
            _ = len(v4)
            _ = fi8
        }
        v5 = (v5 ^ 127)
    } else {
        var v9 int32 = 2147483647
        _ = v9
    }
    v1 += (v1 >> ((2) & 63))
    switch v3 {
    case "a":
        if len(v2) > 0 {
            v2[0] = (-(v1))
        }
    }
    var v10 int64 = v6
    if (!false) {
        var v11 string = (v3 + (("hello" + "\n\t") + "foo bar"))
        v1 = v1
        if (18446744073709551615 < 18446744073709551615) {
            fmt.Println(v6, v6, v10)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch ((v3 + ("\n\t" + v3)) + v11) {
            case "\n\t":
                var v12 bool = ((!true) || (42 <= 2))
                _ = v12
            }
        } else {
            delete(v4, "")
            for fi13 := 0; fi13 < 9; fi13++ {
                _ = len(v2)
                var v14 int64 = -1
                v15 := GenS1{F0: 992, F1: 0}
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        _ = v11
    }
    _ = len(v2)
    if (false && (!(v6 >= -242))) {
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 int16 = -32768
                _ = fi17
                _ = v18
            }
            v19 := v4[v3]
            var v20 int32 = 42
            _ = fi16
            _ = v19
            _ = v20
        }
    }
    switch (v3 + ((("a" + (((v3 + v3) + v3) + ((("\n\t" + "foo bar") + v3) + (v3 + "a")))) + v3) + "foo bar")) {
    case "hello":
        v2 = append(v2, (-(v1)))
    case "":
        switch (v1 + (-(v1))) {
        case 4294967295:
            var v21 float32 = float32(float64((0.0 / 1.0)))
            _ = v21
        case 9223372036854775807:
            if (float64(float32(1e10)) > -2.25) {
                v22 := GenS1{F0: 1, F1: 792}
                var v23 uint32 = 288
                _ = v22
                _ = v23
            } else {
                _ = len(v4)
                for fi24 := 0; fi24 < 10; fi24++ {
                    var v25 uint = 4294967295
                    v1 = fi24
                    _ = fi24
                    _ = v25
                }
                if (!true) {
                    _ = len(v4)
                    var v26 int8 = v5
                    v27 := GenS1{F0: 1, F1: 928}
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(v1, v6)
                }
            }
        default:
            for fi28 := 0; fi28 < 4; fi28++ {
                for fi29 := 0; fi29 < 2; fi29++ {
                    var v30 float32 = (0.0 * 3.14159)
                    _ = fi29
                    _ = v30
                }
                _ = fi28
            }
        }
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
