package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    var v3 int32 = 1
    var v4 int32 = v3
    v5 := make(map[uint32]uint32)
    var v6 uint = 0
    for fi7 := 0; fi7 < 18; fi7++ {
        fmt.Println(fi7)
        for fi8 := 0; fi8 < 9; fi8++ {
            if false {
                var v9 float32 = 3.14159
                v10 := v2[42]
                _ = v10
                _ = v9
            } else {
                v6 += (v6 & 65535)
                fmt.Println(v4)
            }
            _ = fi8
        }
        _ = fi7
    }
    v11 := []int{-1, -218, 32768, -979}
    fmt.Println(v6)
    if (int(v6) > 1) {
        delete(v5, 0)
        if (462 != (^(v6))) {
            if (1 > 9223372036854775807) {
                v4 -= v4
            } else {
                v4 = (v4 | (v3 >> ((v6) & 63)))
                var v12 uint64 = 406
                _ = v12
            }
            v11 = append(v11, 42)
            switch ("\n\t" + "a") {
            case "a":
                if ((18446744073709551615 >= v6) || ((2147483647 >= 2) && (v6 != 2))) {
                    if (32767 >= int(v6)) {
                        for fi13 := 0; fi13 < 20; fi13++ {
                            var v14 float64 = (((0.1 / 1.5) + 0.0) - float64((((3.14159 / 1.0) - (0.0 * 1e10)) - float32((1e10 * -2.25)))))
                            v15 := GenS1{F0: 42, F1: (("hello" + ("" + ((("\n\t" + "\n\t") + ("foo bar" + "foo bar")) + ("foo bar" + "hello")))) + ""), F2: 42}
                            var v16 string = ("" + "hello")
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                    }
                }
            case "hello":
                switch ((("a" + "a") + "hello") + (("foo bar" + "hello") + ("a" + "foo bar"))) {
                case "hello":
                    var v17 uint32 = 2147483647
                    _ = v17
                }
            case "\n\t":
                fmt.Println(v3, v6, v3)
            }
        }
    } else {
        v4 += -2147483648
        if false {
            v3 -= 0
            var v18 string = "foo bar"
            _ = v18
        }
        v19 := GenS1{F0: -32768, F1: "", F2: 1}
        _ = v19
    }
    var v20 uint = v6
    v4 = (v3 & (^(v4)))
    if false {
        v21 := v5[787]
        var v22 float64 = (0.0 - float64(1e10))
        _ = len(v2)
        _ = v21
        _ = v22
    }
    switch 42 {
    case 4294967295:
        switch -2147483648 {
        case 4294967295:
            var v23 int16 = 1
            _ = v23
        }
    case -1:
        if ((true || (v20 >= v20)) && (-821 >= -32768)) {
            v24 := GenS2{F0: "a", F1: 2147483647}
            var v25 string = ""
            _ = v24
            _ = v25
        } else {
            if (0.1 > 1.5) {
                for fi26 := 0; fi26 < 12; fi26++ {
                    var v27 uint = 2
                    _ = fi26
                    _ = v27
                }
                v4 -= 2147483647
                fmt.Println(v20, v4, v4)
            } else {
                fmt.Println(v20, v3)
                switch 128 {
                case 42:
                    var v28 float64 = 0.1
                    _ = v28
                case -1:
                    _ = len(v5)
                case -9223372036854775808:
                    for fi29 := 0; fi29 < 10; fi29++ {
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        for fi30 := 0; fi30 < 11; fi30++ {
                            switch fi30 {
                            case 4294967295:
                                v20 -= (v6 - (^(v6)))
                            case 32767:
                                switch (-(fi30)) {
                                case 0:
                                    if ((-9223372036854775808 > 0) && (-128 >= 127)) {
                                        var v31 int32 = int32(fi29)
                                        _ = v31
                                    } else {
                                        fmt.Println(v4)
                                    }
                                }
                            }
                            var v32 int = (fi30 >> ((v20) & 63))
                            fmt.Println(v3, v32)
                            _ = fi30
                            _ = v32
                        }
                        _ = fi29
                    }
                }
            }
        }
    default:
        v33 := GenS3{F0: 9223372036854775807}
        _ = v33
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        fmt.Println(fi34, v3, fi34)
        _ = fi34
    }
    for fi35 := 0; fi35 < 8; fi35++ {
        if true {
            v36 := GenS1{F0: -912, F1: "foo bar", F2: 65535}
            if true {
                v3 = 0
                var v37 int64 = 0
                v4 -= v4
                _ = v37
            }
            _ = v36
        }
        _ = fi35
    }
    _ = len(v11)
    v4 -= (-(v4))
    if len(v11) > 0 {
        v11[0] = 1
    }
    var v38 float64 = 1e10
    v5[825] = uint32(v20)
    if (1 > 423) {
        v39 := v2[787]
        _ = v39
    }
    var v40 uint32 = 854
    var v41 uint8 = 221
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
