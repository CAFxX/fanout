package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 string; F2 uint16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint32 }

func main() {
    v1 := []uint32{}
    var v2 int32 = 42
    v3 := []uint32{755}
    v4 := make(map[uint32]bool)
    v5 := make(map[int]bool)
    v3 = append(v3, 2147483647)
    for fi6 := 0; fi6 < 13; fi6++ {
        switch (fi6 >> ((995) & 63)) {
        case 32768:
            v7 := GenS1{F0: 486, F1: 124, F2: 1}
            _ = v7
        case 32767:
            switch (("\n\t" + (((("\n\t" + (("hello" + "foo bar") + "a")) + "foo bar") + ((("foo bar" + (("" + (((("foo bar" + (("\n\t" + (("\n\t" + "\n\t") + "")) + "foo bar")) + "a") + ("hello" + "foo bar")) + "hello")) + ("hello" + "foo bar"))) + "hello") + "")) + ("\n\t" + "\n\t"))) + "hello") {
            case "a":
                v8 := GenS1{F0: 1, F1: 206, F2: -463}
                _ = v8
            }
        }
        _ = fi6
    }
    v2 = v2
    v9 := GenS1{F0: 4294967295, F1: 127, F2: 42}
    v10 := GenS2{F0: 91}
    for fi11 := 0; fi11 < 17; fi11++ {
        if true {
            var v12 int16 = 673
            _ = v12
        } else {
            fmt.Println(fi11, v2)
            v13 := GenS1{F0: 4294967295, F1: 181, F2: int64(fi11)}
            if false {
                v14 := GenS1{F0: 1, F1: 1, F2: -580}
                var v15 uint32 = 747
                fmt.Println(v2, fi11, fi11)
                _ = v14
                _ = v15
            }
            _ = v13
        }
        _ = fi11
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        if len(v3) > 0 {
            v3[0] = 557
        }
        _ = fi16
    }
    fmt.Println(v2)
    switch "" {
    case "":
        _ = len(v5)
    default:
        if ((0.1 != (3.14159 + -2.25)) && (true || true)) {
            v2 = v2
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 uint64 = 727
                _ = fi17
                _ = v18
            }
        }
    }
    if false {
        if (635 >= 745) {
            v5[9223372036854775807] = ((0 >= 887) || false)
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 bool = (!false)
            v20 = v20
            _ = fi19
            _ = v20
        }
        if true {
            v2 = v2
            var v21 uint16 = 0
            _ = v21
        } else {
            switch -972 {
            case 0:
                switch "" {
                case "foo bar":
                    if true {
                        var v22 uint16 = 255
                        _ = v22
                    } else {
                        var v23 float32 = (1.5 - -2.25)
                        v24 := GenS3{F0: 139, F1: (("hello" + "a") + (("foo bar" + "") + ("hello" + ("\n\t" + ("\n\t" + ""))))), F2: 255}
                        v2 = (v2 * -1)
                        _ = v23
                        _ = v24
                    }
                }
            default:
                if true {
                    var v25 bool = true
                    _ = v25
                } else {
                    var v26 uint32 = 2
                    var v27 int16 = 1
                    _ = v26
                    _ = v27
                }
            }
            var v28 string = "\n\t"
            for fi29 := 0; fi29 < 16; fi29++ {
                v2 += -381
                var v30 float64 = ((0.1 / 3.14159) * 0.0)
                _ = fi29
                _ = v30
            }
            _ = v28
        }
    }
    v31 := GenS2{F0: -128}
    for fi32 := 0; fi32 < 18; fi32++ {
        v33 := GenS4{F0: 274, F1: 143, F2: 2147483648}
        _ = fi32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
