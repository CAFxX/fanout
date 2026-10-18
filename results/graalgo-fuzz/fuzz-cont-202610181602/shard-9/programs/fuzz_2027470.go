package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }

func main() {
    v1 := []uint32{}
    v2 := make(map[string]uint32)
    var v3 uint16 = 915
    var v4 uint8 = 128
    fmt.Println(v4, v4)
    switch ("foo bar" + "hello") {
    case "hello":
        fmt.Println(v3, v4)
    }
    for fi5 := 0; fi5 < 6; fi5++ {
        v3 = 65535
        v6 := v2["hello"]
        var v7 int64 = -1
        _ = fi5
        _ = v6
        _ = v7
    }
    if (!(true || (65535 > 4294967295))) {
        delete(v2, ("hello" + ("foo bar" + "a")))
        switch "hello" {
        case "foo bar":
            switch -120 {
            case 680:
                v8 := v2["hello"]
                _ = v8
            default:
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 uint64 = 0
                    if len(v1) > 0 {
                        v1[0] = 17
                    }
                    _ = fi9
                    _ = v10
                }
            }
        case "":
            v4 += v4
        default:
            if true {
                for fi11 := 0; fi11 < 16; fi11++ {
                    var v12 int32 = 370
                    if (true && (true || true)) {
                        v3 -= v3
                    }
                    _ = fi11
                    _ = v12
                }
            }
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            switch fi13 {
            case 42:
                var v14 string = "foo bar"
                _ = v14
            case 128:
                if ((v3 << ((544) & 63)) != (v3 << ((677) & 63))) {
                    fmt.Println(fi13, fi13)
                    v3 += (v3 & 0)
                }
            case 2147483647:
                var v15 int32 = 42
                _ = v15
            }
            _ = fi13
        }
    }
    if (false || true) {
        if (-1 == 2147483647) {
            if (!true) {
                var v16 int64 = -1
                for fi17 := 0; fi17 < 14; fi17++ {
                    v3 -= (v3 - (v3 >> ((4294967295) & 63)))
                    _ = fi17
                }
                _ = v16
            }
            for fi18 := 0; fi18 < 4; fi18++ {
                for fi19 := 0; fi19 < 19; fi19++ {
                    _ = len(v1)
                    v20 := GenS1{F0: 4294967295, F1: 81}
                    var v21 float64 = -2.25
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                _ = fi18
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                var v23 float64 = float64(0.1)
                v23 = (v23 * float64(1e10))
                _ = fi22
                _ = v23
            }
        }
        var v24 uint16 = v3
        _ = v24
    }
    fmt.Println(v3, v3)
    if true {
        var v25 int = -22
        for fi26 := 0; fi26 < 18; fi26++ {
            var v27 int16 = 1000
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        if (-561 == -216) {
            fmt.Println(v4)
            var v29 int32 = -2147483648
            _ = v29
        }
        if (uint(v3) < 0) {
            for fi30 := 0; fi30 < 10; fi30++ {
                var v31 bool = (-1 <= 194)
                v32 := GenS1{F0: 807, F1: 216}
                _ = fi30
                _ = v31
                _ = v32
            }
        } else {
            fmt.Println(fi28, fi28)
            _ = len(v2)
            v33 := GenS1{F0: 627, F1: 2}
            _ = v33
        }
        _ = fi28
    }
    var v34 uint32 = 2147483647
    fmt.Println(v34, v3)
    var v35 int = 128
    var v36 bool = (("hello" + "hello") != "")
    switch "\n\t" {
    case "\n\t":
        var v37 uint32 = (^(v34))
        _ = v37
    case "foo bar":
        v4 = uint8(v34)
    case "hello":
        v34 = (v34 & (v34 ^ 2147483647))
    default:
        var v38 bool = v36
        _ = v38
    }
    fmt.Println(v35)
    v39 := GenS1{F0: 1, F1: 4294967295}
    v36 = (v36 && v36)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
}
