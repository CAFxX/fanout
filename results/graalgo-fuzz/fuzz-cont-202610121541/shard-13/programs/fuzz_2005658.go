package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 float64 = 0.0
    v2 := []int{128, 127}
    v3 := make(map[int]uint32)
    var v4 int8 = 0
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    if true {
        switch 4294967295 {
        case 127:
            switch "\n\t" {
            case "foo bar":
                for fi7 := 0; fi7 < 15; fi7++ {
                    var v8 int64 = 524
                    _ = fi7
                    _ = v8
                }
            default:
                fmt.Println(v4, v4)
            }
        default:
            v9 := GenS1{F0: 2}
            _ = v9
        }
    } else {
        v4 -= -77
        v10 := GenS1{F0: 42}
        if len(v2) > 0 {
            v2[0] = 32768
        }
        _ = v10
    }
    v1 = 1.5
    fmt.Println(v1, v1)
    var v11 uint = 2
    delete(v6, 755)
    v12 := GenS1{F0: 4294967295}
    if (!true) {
        switch 96 {
        case 1:
            switch -617 {
            case 65:
                switch -548 {
                case 32768:
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v11 = (^(v11))
                        _ = fi13
                    }
                default:
                    v14 := GenS2{F0: 128}
                    _ = v14
                }
            case -1:
                var v15 uint64 = 1
                _ = v15
            default:
                var v16 string = ""
                _ = v16
            }
        case -333:
            v11 += (v11 + 578)
        case 4294967295:
            var v17 uint32 = uint32(v11)
            _ = v17
        default:
            for fi18 := 0; fi18 < 8; fi18++ {
                v19 := GenS2{F0: 255}
                var v20 string = ""
                var v21 uint8 = 128
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
        }
        if (39 < 255) {
            v4 -= v4
        } else {
            delete(v3, -2147483648)
        }
    } else {
        var v22 bool = false
        var v23 uint8 = 0
        _ = v22
        _ = v23
    }
    v11 = 255
    v11 += (^(v11))
    var v24 int64 = 1
    if true {
        for fi25 := 0; fi25 < 9; fi25++ {
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 float32 = ((1e10 + float32(0.1)) * 0.1)
                fmt.Println(v1, v24)
                if true {
                    if false {
                        v28 := GenS1{F0: 42}
                        if false {
                            var v29 uint64 = 4294967295
                            v24 = (v24 * v24)
                            v4 += (-(v4))
                            _ = v29
                        } else {
                            v1 += (v1 * v1)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v11 -= (^(v11))
                        }
                        _ = len(v2)
                        v30 := GenS1{F0: 178}
                        v24 = v24
                        v4 += (v4 | 100)
                        v31 := v3[(fi25 << ((1) & 63))]
                        if (42 < 4294967295) {
                            var v32 int64 = 0
                            v27 = v27
                            _ = v32
                        }
                        if len(v2) > 0 {
                            v2[0] = fi26
                        }
                        var v33 bool = false
                        _ = v28
                        _ = v30
                        _ = v31
                        _ = v33
                    }
                    var v34 int16 = -32768
                    v5[""] = "hello"
                    _ = v34
                }
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    }
    switch ("" + "\n\t") {
    case "foo bar":
        switch (((("" + "hello") + "a") + "a") + ("a" + (((("\n\t" + ("\n\t" + ("foo bar" + ("\n\t" + "foo bar")))) + (("foo bar" + (("\n\t" + "a") + (("" + "\n\t") + "a"))) + (("" + "a") + ""))) + "hello") + "hello"))) {
        case "hello":
            switch "foo bar" {
            case "":
                for fi35 := 0; fi35 < 14; fi35++ {
                    if (("foo bar" <= ((("hello" + ("a" + "a")) + "foo bar") + "")) || true) {
                        v11 = v11
                    } else {
                        fmt.Println(v1)
                    }
                    v24 = (^(v24))
                    _ = fi35
                }
            default:
                var v36 int32 = 42
                _ = v36
            }
        }
    case "":
        var v37 float32 = 1.5
        _ = v37
    }
    var v38 bool = false
    v1 = 0.1
    fmt.Println(v1, v38)
    v11 += 255
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
